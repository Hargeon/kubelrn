# Kube setup

## Start minikube

virtual box
```bash
minikube start --driver=virtualbox
```

docker
```bash
minikube start --driver=docker
```

start with not default settings
```bash
minikube start --driver=docker --cpus=4 --memory=8gb --disk-size=25gb
```

connect to minikube
```bash
minikube ssh
```

## Get Info
cluster info
```bash
kubectl cluster-info
```

get nodes
```bash
kubectl get nodes
```

get component statuses
```bash
kubectl get componentstatuses
```

## Stop minikube

```bash
minikube stop
```

```bash
minikube delete
```

## Kubernetes components

- Pod(Container works in Pod)
- Deployment(Set of identical pods)
- Service(Provide access to Deployment)
- Nodes(Servers)
- Clusters(Group of Nodes)

## Apply manifest file

Apply
```bash
kubectl apply -f goecho-v1.yaml
```

Delete
```bash
kubectl delete -f goecho-v1.yaml
```

## Pods

Get Pods
```bash
kubectl get pods
```

Create pod(goecho is pod name, --image is docker image)
```bash
kubectl run goecho --image=hargeon/goecho --port=3000
```

Delete pod(goecho is pod name)
```bash
kubectl delete pods goech
```

Get pods info(goecho is pod name)
```bash
kubectl describe pods goecho
```

Get logs(goecho is pod name)
```bash
kubectl logs goecho
```

Port forward(goecho is pod name, 3001 is local port)
```bash
kubectl port-forward goecho 3001:3000
```

## Deplpyments

Get deployments
```bash
kubectl get deployments
```

Describe(goecho-depl - name)
```bash
kubectl describe deployment goecho-depl
```

Create deployment(goecho-depl - name)
```bash
kubectl create deployment goecho-depl --image hargeon/goecho:latest
```

Scale(goecho-depl - name)
```bash
kubectl scale deployment goecho-depl --replicas 4
```

Get replica set
```bash
kubectl get rs
```

Autoscale(goecho-depl - name)
```bash
kubectl autoscale deployment goecho-depl --min=4 --max=6 --cpu-percent=80 
```

Get horizontal pod autoscaling
```bash
kubectl get hpa
```

Get history(goecho-depl - name)
```bash
kubectl rollout history deployment/goecho-depl
```

Get status(goecho-depl - name)
```bash
kubectl rollout status deployment/goecho-depl
```

Change image(goecho-depl - name, goecho - name of container,)
```bash
kubectl set image deployment/goecho-depl goecho=hargeon/goecho:0.0.2 --record
```

Rollback to previos version(goecho-depl - name)
```bash
kubectl rollout undo deployment/goecho-depl
```

Rollback to any version
```bash
kubectl rollout undo deployment/goecho-depl --to-revision=4
```

Update current image(goecho-depl - name)
```bash
kubectl rollout restart deployment/goecho-depl
```

## Services

- ClusterIP
- NodePort
- ExternalName
- LoadBalancer(Only in cloud clusters)

Create with ClusterIP type(goecho-depl - name of deployment)
```bash
kubectl expose deployment goecho-depl --type=ClusterIP --port 3000
```

Create with NodePort type(goecho-depl - name of deployment)
```bash
kubectl expose deployment goecho-depl --type=NodePort --port 3000
```

Create with LoadBalancer type(goecho-depl - name of deployment)
```bash
kubectl expose deployment goecho-depl --type=LoadBalancer --port 3000
```

Get minikube URL(goecho-depl - name of service)
```bash
minikube service goecho-depl --url
```

Get
```bash
kubectl get services
```

Delete
```bash
kubectl delete service goecho-depl
```

## Ingress in minikube

Enable addon
```bash
minikube addons enable ingress
```

Check
```bash
kubectl get pods -n ingress-nginx
```

Get
```bash
kubectl get ingress
```

Minikube ip
```bash
minikube ip
```

## Helm

```bash
helm install goecho helm_chart/
```

```bash
helm list
```

```bash
helm upgrade goecho helm_chart/
```

```bash
helm uninstall goecho
```

## Under the hood

- etcd - the main storage in cluster. store current and desired state. Each spec goes instant to etcd
- scheduler - planning placement our services(apps) in a cluster. it doesn't place our containers. it complements spec.
- controller manager - it consists of multiple controllers. it make current state desired.
- dns

- kube-proxy - it is a process that subscribes to a change in a node and sets up network.
- kubelet - works with docker engine, container engine 

## Tips

- depoyment/containers[].resources.request - need check the params. Use LimitRange
- cpu governor mode. set up perfomance. /sys/devices/system/cpu/cpu0/cpufreq/scaling_governor - perfomance