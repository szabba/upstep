package planning_test

import "testing"

func AssertNoError(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
}

func Assert(t *testing.T, cond bool, msg string, args ...interface{}) {
	if !cond {
		t.Fatalf(msg, args...)
	}
}
