# go-winjob

[![Go Reference](https://pkg.go.dev/badge/github.com/aperturerobotics/go-winjob.svg)](https://pkg.go.dev/github.com/aperturerobotics/go-winjob)
[![Tests](https://github.com/aperturerobotics/go-winjob/actions/workflows/tests.yml/badge.svg?branch=master)](https://github.com/aperturerobotics/go-winjob/actions/workflows/tests.yml)

Go bindings for [Windows Job Objects](https://learn.microsoft.com/en-us/windows/win32/procthread/job-objects).

**This is a hard fork of the [upstream project].**

[upstream project]: https://github.com/kolesnikovae/go-winjob

At the time this fork was created, the upstream repository's latest release was
`v1.0.0` from June 30, 2020. This repository carries ongoing Aperture Robotics
maintenance, modern CI, and the `github.com/aperturerobotics/go-winjob` module
path.

The package provides a high-level Go API for creating and managing Windows job
objects. The [`jobapi`](./jobapi) sub-package exposes the lower-level Win32
types and calls when you need direct control over the underlying APIs.

## Features

- Create, open, terminate, and inspect Windows job objects
- Start processes directly inside a job object
- Apply basic, extended, UI, CPU, and network rate limits
- Subscribe to job object completion-port notifications
- Query aggregate accounting counters for a job and its processes

## Installation

Add the module to your `go.mod`:

```sh
go get github.com/aperturerobotics/go-winjob@latest
```

## Usage

### Start a process inside a job object

```go
cmd := exec.Command("app.exe")

job, err := winjob.Start(
	cmd,
	winjob.WithKillOnJobClose(),
	winjob.WithBreakawayOK(),
)
if err != nil {
	// ...
}
defer job.Close()

if err := cmd.Wait(); err != nil {
	// ...
}
```

`WithKillOnJobClose` is useful for lifecycle ownership: when the last handle to
the job closes, Windows terminates the associated processes and destroys the job
object.

### Apply and inspect limits

```go
limits := []winjob.Limit{
	winjob.WithKillOnJobClose(),
	winjob.WithWorkingSetLimit(1<<20, 8<<20),
	winjob.WithCPUHardCapLimit(5000),
	winjob.WithDSCPTag(0x14),
}

if err := job.SetLimit(limits...); err != nil {
	// ...
}

if err := job.QueryLimits(); err != nil {
	// ...
}

cpu := winjob.LimitCPU.LimitValue(job)
// cpu == winjob.CPURate{HardCap: 5000}
```

Limit values are populated by `QueryLimits()` and are also accessible through
`job.JobInfo` for callers that want the raw state.

### Subscribe to notifications

```go
c := make(chan winjob.Notification, 1)
s, err := winjob.Notify(c, job)
if err != nil {
	// ...
}
defer s.Close()

go func() {
	for n := range c {
		switch n.Type {
		case winjob.NotificationNewProcess:
			// ...
		case winjob.NotificationExitProcess:
			// ...
		default:
			log.Println(n.Type, n.PID)
		}
	}
}()
```

Most completion-port notifications are best-effort delivery. Notification-limit
messages set through the dedicated notification information classes are the
exception and are guaranteed by Windows.

### Query accounting counters

```go
var counters winjob.Counters
if err := job.QueryCounters(&counters); err != nil {
	// ...
}
```

`job.Counters()` is a convenience wrapper when you do not need to reuse the
same `Counters` allocation.

## Examples

- [Package example](./example_test.go)
- [Interactive demo](./examples)
- [Low-level Win32 surface](./jobapi)

## Development

Run tests on Windows:

```sh
go test ./...
```

From another platform, compile the Windows test binaries without executing
them:

```sh
GOOS=windows GOARCH=amd64 go test -exec=true ./...
```

## License

[MIT](./LICENSE)
