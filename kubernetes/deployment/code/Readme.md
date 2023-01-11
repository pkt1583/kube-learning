az acr login -n myacrmyrg

make

k run curl -it --rm --image=curlimages/curl /bin/sh


1) create a pod
2) curl on pod
3) schedule it on node
3) curl again
4) create a service

kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.5.1/deploy/static/provider/cloud/deploy.yaml