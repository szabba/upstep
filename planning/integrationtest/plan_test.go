package integrationtest

import "testing"

func Test_Plan_CanBeRetrieved(t *testing.T) {
	// given
	var srv Server
	srv.Start(t)
	defer srv.Kill(t)

	// when

	// then
}
