# Demo app
Online Boutique is a cloud-first microservices demo application. Online Boutique consists of an 11-tier microservices application.

# How-to run:
1. Deploy it to Kubernetes:
    ```shell
    kubectl apply -k .
    ```
2. Forward service port to access UI locally:
    ```shell
    kubectl port-forward -n demo-app svc/frontend-external 4000:80
    ```
3. Visit http://localhost:4000 in browser

Source: https://github.com/GoogleCloudPlatform/microservices-demo