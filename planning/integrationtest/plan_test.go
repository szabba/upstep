package integrationtest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/szabba/assert"
)

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
	assertHTTPStatus(resp, 200, t.Fatalf)

	assert.That(
		resp.Header.Get("Content-Type") == "application/json",
		t.Errorf, "got content type %s, want %s",
		resp.Header.Get("Content-Type"), "application/json")

	var dto PlanDTO
	dec := json.NewDecoder(resp.Body)
	dec.DisallowUnknownFields()
	err = dec.Decode(&dto)
	assert.That(err == nil, t.Fatalf, "error parsing response: %s", err)
}

func assertHTTPStatus(resp *http.Response, wanted int, onErr assert.ErrorFunc) {
	assert.That(
		resp.StatusCode == wanted,
		onErr, "got status %s, want %s",
		NiceStatus(resp.StatusCode), NiceStatus(wanted))
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
