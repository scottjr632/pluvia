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

## Using the Kubernetes features

 You can create a basic kind cluster using the kind-config.yaml stored in this repo you will require the following for this to work.

### Prerequisities
    - [kind](https://kind.sigs.k8s.io/)
    - A container runtime engine. Pluvia is tested with [Docker](https://docs.docker.com/get-started/)
    - [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/)

    *Ensure docker is running and that kind installed correctly. verify kind with `kind version`*

### Creating the development cluster

To create a kind cluster using the provided config which will generate a two node cluster consisting of one control-plane and two worker node.
Run the following command from the root of the pluvia directory

`kind create cluster --config=config/kind-config.yaml`

By default kind will update the location that kubeconfig you have the `KUBECONFIG` environment variable pointing to.
If not set, it will use the default kubernetes kubeconfig location of `$HOME/.kube/config`.

If for some reason you've never built a cluster and don't have the directory `$HOME/.kube`. Don't worry kind will generate the necessary directories and files for you!

### Destroying the development cluster

By default `kind delete cluster` and for that matter `kind create cluster` ran without arguments will create and delete clusters called `kind`.
This will also exit with a `0` status regardless of whether or not the cluster exists. 

To delete the development cluster we created run: `kind delete cluster --name pluvia`
