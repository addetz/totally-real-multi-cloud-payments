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

1. Run client & server without Kubernetes
    ```bash 
    docker-compose up
    ```

1. Run client & server with Kubernetes
    ```bash
    kubectl apply -f k8s
    kubectl port-forward service/totally-real-multi-cloud-payments-client-service 8080:8080
    kubectl port-forward service/nats 4222:4222
    kubectl get deployments
    kubectl get pods
    ```
## Publish images to DockerHub

1. Build the docker images
```bash
docker build -t totally-real-multi-cloud-payments-client -f ./docker/client/Dockerfile .
docker build -t totally-real-multi-cloud-payments-server -f ./docker/server/Dockerfile .
```
1. Tag the image
```bash
docker tag totally-real-multi-cloud-payments-client classicaddetz/totally-real-multi-cloud-payments-client:1.0.0
docker tag totally-real-multi-cloud-payments-server classicaddetz/totally-real-multi-cloud-payments-server:1.0.0
```

1. Login to DockerHub
```bash 
docker login
```
1. Push the image to DockerHub
```bash
docker push classicaddetz/totally-real-multi-cloud-payments-client:1.0.0
docker push classicaddetz/totally-real-multi-cloud-payments-server:1.0.0
```

## Troubleshoot 
Dirty database version: open SQL terminal and run 
```sql
drop table schema_migrations;
```

## Resources
- [Deploying a containerized Go app on Kubernetes](https://www.callicoder.com/deploy-containerized-go-app-kubernetes/)
- [Deploy a Local Cluster with Kubernetes](https://www.cockroachlabs.com/docs/stable/orchestrate-a-local-cluster-with-kubernetes.html)
- [Networking in Kubernetes](https://kubebyexample.com/en/learning-paths/application-development-kubernetes/lesson-3-networking-kubernetes)

