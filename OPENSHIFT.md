# Enable HTTP/2 on a single Ingress Controller
```
oc -n openshift-ingress-operator annotate ingresscontrollers/<ingresscontroller_name> ingress.operator.openshift.io/default-enable-http2=true
```

# Describe Default Ingress Controller
```
oc describe --namespace=openshift-ingress-operator ingresscontroller/default
```

# View status of ingress operator
```
oc describe clusteroperators/ingress
```

# View Ingress Controller Logs
```
oc logs --namespace=openshift-ingress-operator deployments/ingress-operator
```

# Get the hostname
```
oc get ingresses.config/cluster -o jsonpath='{.spec.domain}'
```

https://cloud.redhat.com/blog/grpc-or-http/2-ingress-connectivity-in-openshift