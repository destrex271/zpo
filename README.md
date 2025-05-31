# ZPO: A CLI tool to interact with Zalando PostgreSQL Operator

`zpo` allows the following features:
 - list: List all postgresql clusters either in provided namespace or all namespaces
 - describe: Display full postgres custom resource for a given cluster

## Usage 

### List: 
 - `go run ./cmd list --namespace <namespace>`
 - `go run ./cmd list`

### Describe:
 - `go run ./cmd describe --namespace <namespace> --output <yaml|json> <cluster-name>`

## Setup

 - Ensure you have <a href="https://minikube.sigs.k8s.io/docs/start">minikube</a> installed on your system
 - Also ensure that you have `kubectl` installed on your system
 - After installing minikube run the following:
   ```bash
   minikube start
   ```
 - Install postgres operator:
   ```bash
   git clone https://github.com/zalando/postgres-operator.git
   cd postgres-operator
   ./run_operator_locally.sh
   ```

 - Verify Postgres Operator Status
   ```bash
   kubectl get pod -l name=postgres-operator
   ```
   This should list the postgres operator with Status as `Running`
 - Now lets build the `zpo` client:
   ```bash
   make build
   ```
 - To list resources:
   ```bash
   ./zpo list --namespace <namespace|empty for all>
   ./zpo describe --namespace <namespace> --output <yaml|json> <cluster-name>
   ```

