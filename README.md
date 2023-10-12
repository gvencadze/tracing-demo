# Tracing demo

### Overview

A set of instructions and manifests to deploy different tracing tools and demo-apps

### Contents
1. [Prerequisites](#prerequisites)
2. [Tools](#tracing-tools)
3. [How to deploy](#how-to-deploy)

### Prerequisites:
1. Kubernetes cluster (I've tested on Selectel MKS v1.27.3). Cluster in cloud preferable,\
   cause launching eBPF-based tools on [Kind](https://kind.sigs.k8s.io) and similar projects can be tricky enough 
2. [kubectl](https://kubernetes.io/docs/tasks/tools/#kubectl)

### Tracing tools
1. [coroot](./coroot/README.md)
2. [pixie](./pixie/README.md)
3. [elastic-apm](./elastic-apm/README.md)

### How-to deploy
1. Select tracing system from the list above
2. Follow steps in README to install
3. Install/Run demo apps:
   1. [Online boutique](apps/online-boutique/README.md) (for eBPF based tools)
   2. [APM Playground](apps/elastic-apm-playground/README.md) (for Elastic APM)
