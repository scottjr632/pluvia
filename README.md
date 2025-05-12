# Overview

Pluvia is a tool that will:
 - Assess the current stack using the pulumi state file and find the cost of the deployed infrastructure
 - Allow the user to create a new plan and identify the new cost of the proposed changes
 - Present the user with the cost Delta prior to confirming that application of the upgrades
## Getting Started

Pluvia can be imported as any other pulumi extension.
Currently Pluvia is run from the root directory with `go run main.go` or `./pluvia`

## Architecture
Pluvia is designed to ingest an existing local state file, make it's assessments, and intercept the proposed upgrades to present cost deltas.
