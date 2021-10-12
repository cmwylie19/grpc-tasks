# Helm 

```
helm repo add bitnami https://charts.bitnami.com/bitnami

helm repo update

# list repos
helm repo list

k create ns metrics

helm install kube-state-metrics bitnami/kube-state-metrics -n metrics

# list charts
helm ls -n metrics

# Show  chart info [repo/chart]
helm show chart bitnami/kube-state-metrics

# Show  all info [repo/chart]
helm show all bitnami/kube-state-metrics

# Show helm overrides
helm show values bitnami/kube-state-metrics

# Lint Chart
helm lint [chart]

# Pull Chart
helm pull haproxy-ingress/haproxy-ingress
helm pull haproxy-ingress/haproxy-ingress --untar --untardir /tmp 


# Upgrading a helm chart [release/chart/flags]
helm upgrade kube-state-metrics bitnami/kube-state-metrics --version 2.1.3 

# Delete a chart in a namespace
helm delete challenge-metrics-server -n challenge

# Install a local helm chart
helm install grpc-todos .

# Check template before install
helm template grpc-todos .

# Upgrade a chart 
helm upgrade grpc-todos .

# History of a chart
helm history grpc-todos

# Rollback a revision
helm rollback grpc-todos

# Rollback to a specific revision
helm rollback grpc-todos 4
```