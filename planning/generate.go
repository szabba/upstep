// +build generate

package main

import "github.com/szabba/upstep/codegen"

func main() {
	codegen.GenerateAggregateBoilerplate("planner_boilerplate.go", "Planner")
	codegen.GenerateAggregateBoilerplate("plan_boilerplate.go", "Plan")
	codegen.GenerateAggregateBoilerplate("step_boilerplate.go", "Step")
}
