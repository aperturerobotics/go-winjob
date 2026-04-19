//go:build windows

package winjob_test

import (
	"os/exec"
	"testing"

	"github.com/aperturerobotics/go-winjob"
)

func TestStart(t *testing.T) {
	job, err := winjob.Start(exec.Command(commandName), winjob.WithKillOnJobClose())
	requireNoError(t, err)
	requireNoError(t, job.QueryLimits())
	if !winjob.LimitKillOnJobClose.IsSet(job) {
		t.Fatalf("Limit is not set after Start")
	}
}
