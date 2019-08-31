package integrationtest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/szabba/assert"
)

func TestPlanCanBeRetrieved(t *testing.T) {
	t.Skip()

	// given
	srv := startServer(t)
	defer srv.Close()

	loadFixture(t, "plan")
	planID := "plan-id"

	url := fmt.Sprintf("%s/plan/%s", srv.URL, planID)

	// when
	resp, err := http.Get(url)
	assert.That(err == nil, t.Errorf, "error on GET: %s", err)
	defer resp.Body.Close()
	assert.That(
		resp.StatusCode == http.StatusOK,
		t.Fatalf, "got status %s, want %s",
		NiceStatus(resp.StatusCode), NiceStatus(http.StatusOK))

	assert.That(
		resp.Header.Get("Content-Type") == "application/json",
		t.Errorf, "got content type %s, want %s",
		resp.Header.Get("Content-Type"), "application/json")

	var dto PlanDTO
	err = json.NewDecoder(resp.Body).Decode(&dto)
	assert.That(err == nil, t.Fatalf, "error parsing response: %s", err)
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
