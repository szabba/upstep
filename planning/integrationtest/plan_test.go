package integrationtest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/szabba/assert"
)

const ApplicationJson = "application/json"

func TestPlanCanBeRetrieved(t *testing.T) {
	// given
	srv := startServer(t)
	defer srv.Close()
	defer bufferLogs(t)()

	loadFixture(t, "GET_plan_id.json")
	planID := "plan-id"

	url := fmt.Sprintf("%s/plan/%s", srv.URL, planID)

	// when
	resp, err := http.Get(url)

	// then
	assert.That(err == nil, t.Errorf, "error on GET: %s", err)
	defer resp.Body.Close()
	assertHTTPStatus(resp, http.StatusOK, t.Fatalf)
	assertContentType(resp, ApplicationJson, t.Errorf)

	var dto PlanDTO
	decodeStrictly(resp.Body, &dto, t.Fatalf)
	assert.That(dto.ID == planID, t.Errorf, "got plan ID %q, want %q", dto.Name, planID)
	assert.That(dto.Name == "Learn Go", t.Errorf, "got plan name %q, want %q", dto.Name, "Learn Go")
	assert.That(dto.Status == "Active", t.Errorf, "got plan status %q, want %q", dto.Status, "Active")
}

func TestNonexistentPlanCannotBeRetrieved(t *testing.T) {
	// given
	srv := startServer(t)
	defer srv.Close()

	loadFixture(t, "GET_plan_id.json")
	planID := "nonexistent-plan-id"

	url := fmt.Sprintf("%s/plan/%s", srv.URL, planID)

	// when
	resp, err := http.Get(url)

	// then
	assert.That(err == nil, t.Errorf, "error on GET: %s", err)
	defer resp.Body.Close()
	assertHTTPStatus(resp, http.StatusNotFound, t.Fatalf)
}

func bufferLogs(t *testing.T) func() {
	var buf bytes.Buffer
	log.Logger = log.Logger.Output(&buf)
	return func() {
		log.Logger = log.Logger.Output(os.Stderr)
		for _, line := range strings.Split(buf.String(), "\n") {
			t.Log(line)
		}
	}
}

func assertHTTPStatus(resp *http.Response, wanted int, onErr assert.ErrorFunc) {
	assert.That(
		resp.StatusCode == wanted,
		onErr, "got status %s, want %s",
		NiceStatus(resp.StatusCode), NiceStatus(wanted))
}

func assertContentType(resp *http.Response, wanted string, onErr assert.ErrorFunc) {
	assert.That(
		resp.Header.Get("Content-Type") == wanted,
		onErr, "got content type %q, want %q",
		resp.Header.Get("Content-Type"), wanted)
}

func decodeStrictly(r io.Reader, into interface{}, onErr assert.ErrorFunc) {
	if r == nil {
		onErr("cannot decode from nil Reader")
		return
	}
	dec := json.NewDecoder(r)
	err := dec.Decode(into)
	assert.That(err == nil, onErr, "error decoding JSON: %s", err)
}

type PlanDTO struct {
	ID     string
	Name   string
	Status string
	Steps  []PlanStepDTO
}

type PlanStepDTO struct {
	StepID      string
	Name        string
	Goal        bool
	Taken       bool
	NextStepIDs []string
}
