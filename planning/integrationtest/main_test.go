package integrationtest

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	err := buildApp()
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	os.Exit(m.Run())
}

const _Executable = "app"

func buildApp() error {
	cmd := exec.Command("go", "build", "-o", _Executable, "github.com/szabba/upstep/cmd/upstep-planning")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("\n%s", out)
		log.Print(err)
		return err
	}
	return nil
}

type Server struct {
	cmd *exec.Cmd
	out []byte
	err error
}

func (srv *Server) Start(t *testing.T) {
	path, err := filepath.Abs(_Executable)
	if err != nil {
		t.Fatalf("failed to resolve %s: %s", _Executable, err)
	}

	srv.cmd = exec.Command(path)
	go srv.run(t)
}

func (srv *Server) run(t *testing.T) {
	srv.out, srv.err = srv.cmd.CombinedOutput()
	if srv.err != nil {
		t.Errorf("background server failure: %s", srv.err)
	}
}

func (srv *Server) Kill(t *testing.T) {
	time.Sleep(100 * time.Millisecond) // scheduler hack

	if srv.cmd == nil || srv.cmd.Process == nil {
		t.Fatalf("cannot kill process that was not started")
	}

	err := srv.cmd.Process.Kill()
	if err != nil {
		t.Logf("failed to kill %s (PID = %d): %s", _Executable, srv.cmd.Process.Pid, err)
	}

	err = srv.cmd.Wait()
	if err != nil {
		t.Logf("failed to wait for %s (PID = %d): %s", _Executable, srv.cmd.Process.Pid, err)
	}

	if srv.cmd.ProcessState.ExitCode() != 0 {
		t.Logf("%s", srv.out)
	}
}
