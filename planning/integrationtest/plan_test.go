package integrationtest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/szabba/assert"
)

const ApplicationJson = "application/json"

func TestPlanCanBeRetrieved(t *testing.T) {
	// given
	srv := startServer(t)
	defer srv.Close()

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
	Steps  []PlanStepDTO
	Status string
}

type PlanStepDTO struct {
	StepID      string
	Name        string
	Goal        bool
	Taken       bool
	NextStepIDs []string
}
