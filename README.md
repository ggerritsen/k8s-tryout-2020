# K8s try out repo 2020

#### How to run locally

1. `cd api && go run main.go`  
1. `cd customersvc && go run main.go`  
1. `cd greetsvc && go run main.go`  

#### How to run 1 docker container locally

1. `docker run -t -d <image>`

#### How to run locally with docker-compose

1. `docker-compose up`  
Check with:
- `docker exec -it <container> /bin/bash`

#### How to start a k3s cluster
1. `k3d create --api-port 6550 --publish 8090:80 --workers 3` # 80 is the ingress port that traefik listens on
1. `export KUBECONFIG="$(k3d get-kubeconfig --name='k3s-default')"`
1. `kubectl cluster-info`
1. `kubectl get all -o wide`
1. For login credentials, see `cat $KUBECONFIG`

#### How to deploy to k3s
1. Make sure latest images have been built: `docker-compose build`
1. And have been tagged: `docker tag k8s-tryout-2020_api:latest ggerritsen1/k8s-tryout-2020_api:latest`
1. And have been pushed: `docker push ggerritsen1/k8s-tryout-2020_api:latest`  

Then:  
1. `kubectl apply -f api-deployment.yaml`  
Check with:
- `kubectl get pods -o wide`
- `kubectl logs -f -lapp=k8s-tryout-2020 --all-containers=true --max-log-requests=10`

1. Open http://api.localhost:8090/hello to see that it works

#### Troubleshoot
1. Run bash on a specific container `kubectl exec -it pod/k8s-tryout-2020-api-deployment-54d587f5bf-qqpwv -- /bin/bash`
1. Run bash on a helper container `kubectl --namespace=default run -it --image=alpine helper-container`, then `wget -SO- k8s-tryout-2020-api/hello`
1. See traefik dashboard:
  - `kubectl -n kube-system edit configmap traefik`, then add: 
  ```
    [api]
        dashboard = true 
```
  - `kubectl -n kube-system port-forward deployment/traefik 8080` and open http://localhost:8080


#### Next steps

- Deploy to k3s
- Make smaller containers (from scratch)
- Try https://k8slens.dev/
- Add a database
- Add integration with an external endpoint/system



##### Sources
- [Docker compose docs](https://docs.docker.com/compose/compose-file/)
- https://levelup.gitconnected.com/complete-guide-to-create-docker-container-for-your-golang-application-80f3fb59a15e
- https://nickjanetakis.com/blog/docker-tip-10-project-structure-with-multiple-dockerfiles-and-docker-compose
- https://medium.com/burak-tasci/full-stack-monorepo-part-i-go-services-967bb3527bb8
- https://itnext.io/run-kubernetes-on-your-machine-7ee463af21a2
- https://multipass.run/
- https://katacoda.com/courses/kubernetes/playground
- https://github.com/rancher/k3d
- https://medium.com/@zhimin.wen/running-k3s-with-multipass-on-mac-fbd559966f7c
- https://sysadmins.co.za/develop-build-and-deploy-a-golang-app-to-k3s/
- https://medium.com/google-cloud/kubernetes-101-pods-nodes-containers-and-clusters-c1509e409e16
- https://medium.com/google-cloud/kubernetes-110-your-first-deployment-bf123c1d3f8
- https://medium.com/@geraldcroes/kubernetes-traefik-101-when-simplicity-matters-957eeede2cf8
- https://github.com/rancher/k3d/issues/103
- https://github.com/rancher/k3s/issues/350
- https://rancher.com/docs/k3s/latest/en/installation/kube-dashboard/
