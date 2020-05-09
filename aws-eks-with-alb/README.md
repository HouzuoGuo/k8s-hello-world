## Synopsis
A short exercise to demo usage of AWS Elastic Kubernetes Service for hosting a horizontally scalable application and exposing it to public Internet via Application Load Balancer.

## Cluster setup
`eksctl create cluster --name my-dev-eks-cluster --fargate --version 1.15 --region eu-west-1 --profile my-aws-profile-name`

## Give the cluster capability to run Application Load Balancer
Follow this guide: https://docs.aws.amazon.com/eks/latest/userguide/alb-ingress.html

```
eksctl utils associate-iam-oidc-provider --cluster my-dev-eks-cluster --region eu-west-1 --profile my-aws-profile-name --approve

aws --profile my-aws-profile-name iam create-policy \
    --policy-name ALBIngressControllerIAMPolicy \
    --policy-document https://raw.githubusercontent.com/kubernetes-sigs/aws-alb-ingress-controller/v1.1.4/docs/examples/iam-policy.json

kubectl apply -f https://raw.githubusercontent.com/kubernetes-sigs/aws-alb-ingress-controller/v1.1.4/docs/examples/rbac-role.yaml

eksctl create iamserviceaccount --region eu-west-1 --profile my-aws-profile-name \
    --name alb-ingress-controller \
    --namespace kube-system \
    --cluster my-dev-eks-cluster \
    --attach-policy-arn arn:aws:iam::862724539066:policy/ALBIngressControllerIAMPolicy \
    --override-existing-serviceaccounts \
    --approve

kubectl apply -f https://raw.githubusercontent.com/kubernetes-sigs/aws-alb-ingress-controller/v1.1.4/docs/examples/alb-ingress-controller.yaml
```

Then run `kubectl edit deployment.apps/alb-ingress-controller -n kube-system` and write down:

```
spec:
  containers:
  - args:
    - --ingress-class=alb
    - --cluster-name=my-dev-eks-cluster
    - --aws-vpc-id=vpc-05bedf23757408c52
    - --aws-region=eu-west-1
    image: docker.io/amazon/aws-alb-ingress-controller:v1.1.4
```

Run `kubectl get pods -n kube-system` to verify that the ingress controller has started successfully.

## Install the application
Run `kubectl apply -f ./k8s.yaml` to install the application. The application service is given to ingress controller (AWS Application Load Balancer) and exposed to the public Internet.

Read more about ALB from this guide: https://docs.aws.amazon.com/eks/latest/userguide/alb-ingress.html

P.S. I have not managed to make AWS ELB (classic load balancer) and NLB (network load balancer) work with this app. ELB's public Internet endpoint does not respond anything at all, and NLB refuses to start due to AWS internal error.
