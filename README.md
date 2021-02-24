#Calipsa tasks

This repo is having 2 tasks - 

1. simple golang webapp which 
- displays header in table, 
- sort by header name on port 8080 

This also contains Docker file to build webapp image. Container start can auto-start webapp

2. A simple helm chart with deployment and service with Default using 8080 for http


Go_Webapp instructions - 

To run go application - 
go run calipsa_goapp.go

To check the Web app - 
Sorted header file output can be obtained in http://localhost:8080
A simple hello app output can be obtained in http://localhost:8080/hello

To run Dockerfile - 
$ docker run -it golang:1.15.8-alpine3.13 bash
$ docker build -t calipsa-go .

Helm chart instructions -

To run helm chart - 

$ export POD_NAME=$(kubectl get pods -l "app.kubernetes.io/name=cherry-awesome-app,app.kubernetes.io/instance=my-helm-chart" -o jsonpath="{.items[0].metadata.name}")

$ kubectl port-forward $POD_NAME 8080:80
