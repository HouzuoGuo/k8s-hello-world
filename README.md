# Gentle start
1. Compile the app

    go build -o my-app

2. Build the app's container image

    sudo docker build -t hzgl/k8s-hello-world .

3. Test run

    sudo docker run --rm -p 20000:20000 hzgl/k8s-hello-world
    curl localhost:20000

# Cloud-on-my-desktop
1. Start [minikube](https://kubernetes.io/docs/setup/learning-environment/minikube/)

    sudo ~/minikube start --vm-driver=none --apiserver-port 10000

2. Start application cluster with load balancer

    sudo kubectl apply -f k8s.yaml

3. See the live cluster action

    sudo kubectl get pods
    sudo kubectl get services
    curl LOAD-BALANCER-IP:20000

4. Tear down

    sudo kubectl delete -f k8s.yaml
    sudo ~/minikube stop

# Deploy to Azure
1. Push container image to a common registry

    sudo docker push hzgl/k8s-hello-world

2. Login to cloud

    az login
    az aks get-credentials --resource-group houzuo-guo --name hzgl-k8s-practice

3. Deploy

    kubectl --context=hzgl-k8s-practice apply -f ./k8s.yaml
