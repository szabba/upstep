package integrationtest

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"github.com/szabba/assert"
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
	cmd  *exec.Cmd
	port int
	out  []byte
	err  error
}

func (srv *Server) Start(t *testing.T) {
	path, err := filepath.Abs(_Executable)
	assert.That(err == nil, t.Fatalf, "failed to resolve %s: %s", _Executable, err)

	srv.port = srv.randomPort(t)

	srv.cmd = exec.Command(path)
	srv.cmd.Env = append(os.Environ(), fmt.Sprintf("PORT=%d", srv.port))
	go srv.run(t)
}

func (srv *Server) Addr(t *testing.T) string {
	assert.That(srv.port != 0, t.Fatalf, "server has not started yet, port is unknown")
	return fmt.Sprintf("http://localhos:%d", srv.port)
}

func (srv *Server) randomPort(t *testing.T) int {
	l, err := net.Listen("tcp", "localhost:0")
	assert.That(err == nil, t.Fatalf, "cannot get random port: %s", err)

	addr := l.Addr()
	err = l.Close()
	assert.That(err == nil, t.Fatalf, "cannot close socket listener: %s", err)

	switch addr := addr.(type) {
	case *net.TCPAddr:
		return addr.Port
	default:
		t.Fatalf("cannot get port of %#v, unexpected type %T", addr, addr)
		// unreachable
		return 0
	}
}

func (srv *Server) run(t *testing.T) {
	srv.out, srv.err = srv.cmd.CombinedOutput()
	assert.That(srv.err == nil, t.Errorf, "background server failure: %s", srv.err)
}

func (srv *Server) Kill(t *testing.T) {
	time.Sleep(100 * time.Millisecond) // scheduler hack

	assert.That(
		srv.cmd != nil && srv.cmd.Process != nil,
		t.Fatalf, "cannot kill process that was not started")

	err := srv.cmd.Process.Kill()
	assert.That(
		err == nil,
		t.Logf, "failed to kill %s (PID = %d): %s", _Executable, srv.cmd.Process.Pid, err)

	err = srv.cmd.Wait()
	assert.That(
		err == nil,
		t.Logf, "failed to wait for %s (PID = %d): %s", _Executable, srv.cmd.Process.Pid, err)

	assert.That(
		srv.runOK() && !t.Failed(),
		t.Logf, "%s", srv.out)
}

func (srv *Server) runOK() bool {
	return srv.cmd.ProcessState.ExitCode() == 0
}
