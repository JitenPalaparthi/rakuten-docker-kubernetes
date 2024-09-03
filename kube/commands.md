# Kubernetes Basic Commands

## General Commands
- **View Kubernetes version**
  ```bash
  kubectl version
  ```

- **View cluster information**
  ```bash
  kubectl cluster-info
  ```

- **View available nodes**
  ```bash
  kubectl get nodes
  ```

## Pod Management
- **List all pods in the current namespace**
  ```bash
  kubectl get pods
  ```

- **Describe a pod**
  ```bash
  kubectl describe pod <pod-name>
  ```

- **Create a pod from a YAML file**
  ```bash
  kubectl apply -f <pod-file>.yaml
  ```

- **Delete a pod**
  ```bash
  kubectl delete pod <pod-name>
  ```

- **Get pod logs**
  ```bash
  kubectl logs <pod-name>
  ```

- **Execute a command inside a pod**
  ```bash
  kubectl exec -it <pod-name> -- /bin/bash
  ```

## Deployment Management
- **List all deployments**
  ```bash
  kubectl get deployments
  ```

- **Create a deployment from a YAML file**
  ```bash
  kubectl apply -f <deployment-file>.yaml
  ```

- **Delete a deployment**
  ```bash
  kubectl delete deployment <deployment-name>
  ```

- **Scale a deployment**
  ```bash
  kubectl scale deployment <deployment-name> --replicas=<number>
  ```

## Service Management
- **List all services**
  ```bash
  kubectl get services
  ```

- **Create a service from a YAML file**
  ```bash
  kubectl apply -f <service-file>.yaml
  ```

- **Delete a service**
  ```bash
  kubectl delete service <service-name>
  ```

## ConfigMap and Secret Management
- **List all ConfigMaps**
  ```bash
  kubectl get configmaps
  ```

- **Create a ConfigMap from a YAML file**
  ```bash
  kubectl apply -f <configmap-file>.yaml
  ```

- **List all Secrets**
  ```bash
  kubectl get secrets
  ```

- **Create a Secret from a YAML file**
  ```bash
  kubectl apply -f <secret-file>.yaml
  ```

## Namespace Management
- **List all namespaces**
  ```bash
  kubectl get namespaces
  ```

- **Create a namespace**
  ```bash
  kubectl create namespace <namespace-name>
  ```

- **Delete a namespace**
  ```bash
  kubectl delete namespace <namespace-name>
  ```

## Persistent Volume (PV) and Persistent Volume Claim (PVC) Management
- **List all Persistent Volumes (PVs)**
  ```bash
  kubectl get pv
  ```

- **List all Persistent Volume Claims (PVCs)**
  ```bash
  kubectl get pvc
  ```

- **Create a Persistent Volume from a YAML file**
  ```bash
  kubectl apply -f <pv-file>.yaml
  ```

- **Create a Persistent Volume Claim from a YAML file**
  ```bash
  kubectl apply -f <pvc-file>.yaml
  ```

## Ingress Management
- **List all Ingress resources**
  ```bash
  kubectl get ingress
  ```

- **Create an Ingress resource from a YAML file**
  ```bash
  kubectl apply -f <ingress-file>.yaml
  ```

- **Delete an Ingress resource**
  ```bash
  kubectl delete ingress <ingress-name>
  ```

## Troubleshooting
- **Get events in the cluster**
  ```bash
  kubectl get events
  ```

- **Check resource usage (CPU/Memory) of pods**
  ```bash
  kubectl top pod
  ```

- **View detailed resource information**
  ```bash
  kubectl describe <resource-type> <resource-name>
  ```

## Apply Changes and Manage Resources
- **Apply changes from a YAML file**
  ```bash
  kubectl apply -f <file>.yaml
  ```

- **Delete a resource from a YAML file**
  ```bash
  kubectl delete -f <file>.yaml
  ```

- **View all resources in a namespace**
  ```bash
  kubectl get all -n <namespace>
  ```

## Other Useful Commands
- **View all API resources**
  ```bash
  kubectl api-resources
  ```

- **View the current context**
  ```bash
  kubectl config current-context
  ```

- **Switch context**
  ```bash
  kubectl config use-context <context-name>
  ```

- **View all contexts**
  ```bash
  kubectl config get-contexts
  ```

- **Delete a resource by label**
  ```bash
  kubectl delete <resource-type> -l <label-key>=<label-value>
  ```
```

