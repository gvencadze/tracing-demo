# Elastic APM

### Contents:
1. CRDs (v2.5.0)
2. Operator (v2.5.0)
3. Components (Fleet APM integration + NS)
   1. Kibana (v8.4.2)
   2. Elasticsearch (v8.4.2)
   3. Elastic-agent (v8.4.2)

### Note:
If your cluster isn't located in ru-7,
then you need to change storageClassName 
in kustomize patch located at [components/pvc_patch.yaml](./components/pvc_patch_es.yaml)

### Installation:
```shell
kubectl apply -k elastic
```

## Links:
https://github.com/elastic/cloud-on-k8s/blob/2.5/config/crds/v1/all-crds.yaml