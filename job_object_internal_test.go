//go:build windows

package winjob

import "testing"

func TestCopyCounters(t *testing.T) {
	job := &JobObject{}
	job.AccountingInfo.TotalUserTime = 1
	job.AccountingInfo.TotalKernelTime = 2
	job.AccountingInfo.ThisPeriodTotalUserTime = 3
	job.AccountingInfo.ThisPeriodTotalKernelTime = 4
	job.AccountingInfo.TotalPageFaultCount = 5
	job.AccountingInfo.TotalProcesses = 6
	job.AccountingInfo.ActiveProcesses = 7
	job.AccountingInfo.TotalTerminatedProcesses = 8
	job.AccountingInfo.ReadOperationCount = 9
	job.AccountingInfo.WriteOperationCount = 10
	job.AccountingInfo.OtherOperationCount = 11
	job.AccountingInfo.ReadTransferCount = 12
	job.AccountingInfo.WriteTransferCount = 13
	job.AccountingInfo.OtherTransferCount = 14

	var counters Counters
	job.copyCounters(&counters)

	if counters.TotalUserTime != 1 {
		t.Fatalf("expected TotalUserTime 1, got %d", counters.TotalUserTime)
	}
	if counters.TotalKernelTime != 2 {
		t.Fatalf("expected TotalKernelTime 2, got %d", counters.TotalKernelTime)
	}
	if counters.ThisPeriodTotalUserTime != 3 {
		t.Fatalf("expected ThisPeriodTotalUserTime 3, got %d", counters.ThisPeriodTotalUserTime)
	}
	if counters.ThisPeriodTotalKernelTime != 4 {
		t.Fatalf("expected ThisPeriodTotalKernelTime 4, got %d", counters.ThisPeriodTotalKernelTime)
	}
	if counters.TotalPageFaultCount != 5 {
		t.Fatalf("expected TotalPageFaultCount 5, got %d", counters.TotalPageFaultCount)
	}
	if counters.TotalProcesses != 6 {
		t.Fatalf("expected TotalProcesses 6, got %d", counters.TotalProcesses)
	}
	if counters.ActiveProcesses != 7 {
		t.Fatalf("expected ActiveProcesses 7, got %d", counters.ActiveProcesses)
	}
	if counters.TotalTerminatedProcesses != 8 {
		t.Fatalf("expected TotalTerminatedProcesses 8, got %d", counters.TotalTerminatedProcesses)
	}
	if counters.ReadOperationCount != 9 {
		t.Fatalf("expected ReadOperationCount 9, got %d", counters.ReadOperationCount)
	}
	if counters.WriteOperationCount != 10 {
		t.Fatalf("expected WriteOperationCount 10, got %d", counters.WriteOperationCount)
	}
	if counters.OtherOperationCount != 11 {
		t.Fatalf("expected OtherOperationCount 11, got %d", counters.OtherOperationCount)
	}
	if counters.ReadTransferCount != 12 {
		t.Fatalf("expected ReadTransferCount 12, got %d", counters.ReadTransferCount)
	}
	if counters.WriteTransferCount != 13 {
		t.Fatalf("expected WriteTransferCount 13, got %d", counters.WriteTransferCount)
	}
	if counters.OtherTransferCount != 14 {
		t.Fatalf("expected OtherTransferCount 14, got %d", counters.OtherTransferCount)
	}
}
