### Commands
```
GITHUB_URL=https://github.com/kubernetes/dashboard/releases
VERSION_KUBE_DASHBOARD=$(curl -w '%{url_effective}' -I -L -s -S ${GITHUB_URL}/latest -o /dev/null | sed -e 's|.*/||')
kubectl create -f https://raw.githubusercontent.com/kubernetes/dashboard/${VERSION_KUBE_DASHBOARD}/aio/deploy/recommended.yaml
```

```
kubectl create -f dashboard.admin-user.yml -f dashboard.admin-user-role.yml
kubectl -n kubernetes-dashboard describe secret admin-user-token | grep ^token
kubectl proxy
```

Open http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/  
Sign In with the admin-user Bearer Token


### Source
https://rancher.com/docs/k3s/latest/en/installation/kube-dashboard/
