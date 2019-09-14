// +build not

package main

import "github.com/szabba/upstep/codegen"

func main() {
	codegen.DocstoreRepositoryBoilerplate{
		AggregateName: "Planner",
		ImportPath:    "github.com/szabba/upstep/planning",
		SourcePackage: "planning",
	}.Generate("planner_repo.go")
	codegen.DocstoreRepositoryBoilerplate{
		AggregateName: "Plan",
		ImportPath:    "github.com/szabba/upstep/planning",
		SourcePackage: "planning",
	}.Generate("plan_repo.go")
	codegen.DocstoreRepositoryBoilerplate{
		AggregateName: "Step",
		ImportPath:    "github.com/szabba/upstep/planning",
		SourcePackage: "planning",
	}.Generate("step_repo.go")
}
