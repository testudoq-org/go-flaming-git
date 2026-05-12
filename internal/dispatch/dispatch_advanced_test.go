package dispatch_test

import (
	"errors"
	"testing"

	"github.com/stephenhstewart/fgit/internal/dispatch"
)

func TestDispatch_Shipit_RequiresMessage(t *testing.T) {
	stub := &stubRun{}
	d := dispatch.New(stub.run)

	err := d.Dispatch([]string{"shipit"})
	if err == nil {
		t.Fatal("expected error when shipit is called without a commit message")
	}
}

func TestDispatch_Blaze_RequiresMessage(t *testing.T) {
	stub := &stubRun{}
	d := dispatch.New(stub.run)

	err := d.Dispatch([]string{"blaze"})
	if err == nil {
		t.Fatal("expected error when blaze is called without a commit message")
	}
}

func TestDispatch_Gcm_RequiresMessage(t *testing.T) {
	stub := &stubRun{}
	d := dispatch.New(stub.run)

	err := d.Dispatch([]string{"gcm"})
	if err == nil {
		t.Fatal("expected error when gcm is called without a commit message")
	}
}

func TestDispatch_Burn_ReturnsError_WhenRunFails(t *testing.T) {
	stub := &stubRun{err: errors.New("git error")}
	d := dispatch.New(stub.run)

	err := d.Dispatch([]string{"burn"})
	if err == nil {
		t.Fatal("expected error to propagate from run")
	}
}

func TestDispatch_CompoundAlias_StopsOnFirstError(t *testing.T) {
	callCount := 0
	run := func(args []string) error {
		callCount++
		return errors.New("step failed")
	}
	d := dispatch.New(run)

	_ = d.Dispatch([]string{"shipit", "msg"})
	if callCount != 1 {
		t.Errorf("expected compound to stop after first error, got %d calls", callCount)
	}
}
