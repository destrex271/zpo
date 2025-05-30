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

