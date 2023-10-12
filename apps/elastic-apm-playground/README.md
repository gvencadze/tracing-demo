# Elastic APM playground

#### Note:
There's no image and prepared manifests to deploy this app to Kubernetes,
so it's much easier to run all required stuff locally via docker-compose

# Prerequisites
1. Go 1.18+
2. Docker

# Prepare environment:
### Run on Kubernetes:
1. Deploy Elastic APM in K8s following by following steps in [ECK readme](../../elastic-apm/README.md)
2. Forward service port to be able to send traces to APM
   ```shell
   kubectl port-forward -n apm svc/apm 8200
   ```
3. Get credentials for Kibana admin
   ```shell
   kubectl get secrets -n apm elasticsearch-es-elastic-user -o=jsonpath='{.data.elastic}' | base64 --decode; echo
   ```
4. Forward service port to access Kibana UI locally
   ```shell
   kubectl port-forward -n apm svc/kibana-kb-http 5601
   ```
### Run on Docker:
Execute the following commands:
```shell
   # this command will run Elasticsearch, Kibana, APM server and Postgres in Docker
   docker compose up -d 
    
   # load indexes, in real environment you probably don't need this
   docker-compose exec apm-server ./apm-server setup
```

# How-to run:
1. Run service:
    ```shell
    go run ./cmd/playground/main.go
    ```
2. Send HTTP request:
    ```shell
    curl "http://localhost:8080/user?id=1" -H "X-Request-Id: hello4playground"
    ```
    will return:
    ```shell
    {"FirstName":"Alexey", "LastName":"Akulovich"}
    ```
3. Open Kibana UI at http://localhost:5601, enter the following credentials (if deploy all stuff to Kubernetes. Otherwise, there are no credentials):
   ```yaml
   username: admin
   password: # credentials that were retrieved at a third step of `Prepare environment` section
   ```
   then select APM server integration, scroll to the end and click the last button -
   `Launch APM` and you will see your services