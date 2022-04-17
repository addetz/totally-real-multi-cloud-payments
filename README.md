# totally-real-multi-cloud-payments
Playing around with a demo for my talk "Choosing cloud native technologies for the journey to multi-cloud"

Enter at your own peril. 🙈

## Setup
1. Install [`golang-migrate`](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
1. Complete [Step 1](https://www.cockroachlabs.com/docs/v21.2/orchestrate-a-local-cluster-with-kubernetes.html#step-1-start-kubernetes) and [Step 2](https://www.cockroachlabs.com/docs/v21.2/orchestrate-a-local-cluster-with-kubernetes.html#step-2-start-cockroachdb) from the official docs to set up locally on Kubernetes.

1. Setup the Cockroach SQL client.
```bash
kubectl create -f https://raw.githubusercontent.com/cockroachdb/cockroach-operator/master/examples/client-secure-operator.yaml
```

1. Open a shell into the Cockroach SQL client pod 
```bash
kubectl exec -it cockroachdb-client-secure -- ./cockroach sql --certs-dir=/cockroach/cockroach-certs --host=cockroachdb-public
```

1. In the SQL shell, create database and user. Then, quit the shell using `\q`
```sql
CREATE DATABASE payments;
CREATE USER roach WITH PASSWORD 'payments-r-us';
GRANT admin TO roach;
```

1. Set up port forwarding locally 
```bash
kubectl port-forward service/cockroachdb-public 26257
```
1. Run the payments app. `go run main.go`

1. Turn everything off with `minikube stop` or `minikube delete`.

1. Build client
```bash
docker build -t go-client .  
```

1. Run client 
```bash 
docker run -d -p 8080:8080 go-client
```

## Troubleshoot 
Dirty database version: open SQL terminal and run 
```sql
drop table schema_migrations;
```

