# Setup EKS 


## Create a EKS cluster with eksctl

```
eksctl create cluster --name demo-cluster --region us-east-1 
```

## Apply All manifests files

```
kubectl apply -f k8s/manifests/deployment.yaml 
kubectl apply -f k8s/manifests/service.yaml 
kubectl apply -f k8s/manifests/ingress.yaml
```

## Check all pods are runing 

```
kubectl get deploy
kubectl get ingress
kubectl get svc
```

## Edit service 

```
kubectl edit svc go-web-app-devops
```
Change ClusterIP to NodePort



## Delete the cluster

```
eksctl delete cluster --name demo-cluster --region us-east-1
```