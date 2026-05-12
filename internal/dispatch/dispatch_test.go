package dispatch_test

import (
	"testing"

	"github.com/stephenhstewart/fgit/internal/dispatch"
)

// stubRun records calls without executing anything.
type stubRun struct {
	calls [][]string
	err   error
}

func (s *stubRun) run(args []string) error {
	s.calls = append(s.calls, append([]string(nil), args...))
	return s.err
}

func TestDispatch_SimpleAlias(t *testing.T) {
	stub := &stubRun{}
	d := dispatch.New(stub.run)

	if err := d.Dispatch([]string{"gs"}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(stub.calls) != 1 {
		t.Fatalf("expected 1 git call, got %d", len(stub.calls))
	}
	want := []string{"status"}
	for i, v := range stub.calls[0] {
		if v != want[i] {
			t.Errorf("arg[%d] = %q, want %q", i, v, want[i])
		}
	}
}

func TestDispatch_SimpleAlias_AppendExtraArgs(t *testing.T) {
	stub := &stubRun{}
	d := dispatch.New(stub.run)

	if err := d.Dispatch([]string{"gcm", "my message"}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(stub.calls) != 1 {
		t.Fatalf("expected 1 call, got %d", len(stub.calls))
	}
	if stub.calls[0][len(stub.calls[0])-1] != "my message" {
		t.Errorf("expected last arg to be the user message, got %v", stub.calls[0])
	}
}

func TestDispatch_Passthrough_UnknownCommand(t *testing.T) {
	stub := &stubRun{}
	d := dispatch.New(stub.run)

	if err := d.Dispatch([]string{"log", "--follow", "file.go"}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(stub.calls) != 1 {
		t.Fatalf("expected 1 call, got %d", len(stub.calls))
	}
	if stub.calls[0][0] != "log" {
		t.Errorf("expected passthrough arg log, got %v", stub.calls[0])
	}
}

func TestDispatch_CompoundAlias_Shipit(t *testing.T) {
	stub := &stubRun{}
	d := dispatch.New(stub.run)

	if err := d.Dispatch([]string{"shipit", "my commit message"}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// shipit has 3 steps: add -A, commit -m <msg>, push
	if len(stub.calls) != 3 {
		t.Fatalf("expected 3 git calls for shipit, got %d: %v", len(stub.calls), stub.calls)
	}
	// First step: add -A
	if stub.calls[0][0] != "add" {
		t.Errorf("step 1 should be add, got %v", stub.calls[0])
	}
	// Second step: commit -m <msg>
	if stub.calls[1][0] != "commit" {
		t.Errorf("step 2 should be commit, got %v", stub.calls[1])
	}
	if stub.calls[1][len(stub.calls[1])-1] != "my commit message" {
		t.Errorf("step 2 should end with user message, got %v", stub.calls[1])
	}
	// Third step: push
	if stub.calls[2][0] != "push" {
		t.Errorf("step 3 should be push, got %v", stub.calls[2])
	}
}

func TestDispatch_CompoundAlias_Sync(t *testing.T) {
	stub := &stubRun{}
	d := dispatch.New(stub.run)

	if err := d.Dispatch([]string{"sync"}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(stub.calls) != 2 {
		t.Fatalf("expected 2 calls for sync, got %d", len(stub.calls))
	}
}

func TestDispatch_EmptyArgs_Passthrough(t *testing.T) {
	stub := &stubRun{}
	d := dispatch.New(stub.run)

	if err := d.Dispatch([]string{}); err != nil {
		t.Fatalf("unexpected error on empty args: %v", err)
	}
	if len(stub.calls) != 1 {
		t.Fatalf("expected 1 call, got %d", len(stub.calls))
	}
	if len(stub.calls[0]) != 0 {
		t.Errorf("expected empty args pass-through, got %v", stub.calls[0])
	}
}
