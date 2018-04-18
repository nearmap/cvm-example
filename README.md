# CVM-Example

[![CircleCI](https://circleci.com/gh/nearmap/cvm-example.svg?style=svg&circle-token=66a948aa48d0605a294898519f1fc15eba899cc2)](https://circleci.com/gh/nearmap/cvm-example)[![Go Report Card](https://goreportcard.com/badge/github.com/nearmap/cvm-example)](https://goreportcard.com/report/github.com/nearmap/cvm-example)

CVM-Example is a sample application to demonstrate the capability of [Container Version Manager (CVM)](https://github.com/nearmap/cvmanager/): A CI/CD tool for Kubernetes. 



```
    export AWS_ACC_ID=<Your AWS Account ID>  
    sed "s/<AWS_ACC_ID>/$AWS_ACC_ID/g" k8s.yaml | kubectl apply --record -f - 
```


```    
    host=$(kubectl get service myapp -o=jsonpath='{.status.loadBalancer.ingress[].hostname}')
    curl http://$host/echo?msg=Hello
```
