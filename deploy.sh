#!/bin/bash

#Implementação automática da aplicação
sudo echo ""

#Realiza build e envia para o DockerHub
docker build -t wellingt0npires/fase4automoveis:latest . && docker push wellingt0npires/fase4automoveis:latest

#Encerra e limpa o minikube
minikube stop
minikube delete
minikube cache delete

#Inicia uma nova instância do Minikube
minikube start --ports=8080:8080
sudo chmod 777 /var/run/docker.sock
minikube addons enable metrics-server

#Implementação das APIs e DB

kubectl apply -f fase4automoveis-deployment.yaml,fase4automoveis-service.yaml,postgres-initdb-config.yaml,postgres-claim0-persistentvolumeclaim.yaml,postgres-deployment.yaml,postgres-service.yaml

#Configura HPA
kubectl autoscale deployment fase4automoveis --cpu-percent=80 --min=1 --max=10
kubectl autoscale deployment postgres --cpu-percent=80 --min=1 --max=10