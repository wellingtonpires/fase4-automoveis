#!/bin/bash

#Implementação automática da aplicação
sudo echo ""

#Realiza build e envia para o DockerHub
docker build -t wellingt0npires/fase4automoveis:latest . && docker push wellingt0npires/fase4automoveis:latest
docker build -f vendas/DockerfileVendas -t wellingt0npires/vendas-fase4automoveis:latest . && docker push wellingt0npires/vendas-fase4automoveis:latest

#Encerra e limpa o minikube
minikube stop
minikube delete
minikube cache delete

#Inicia uma nova instância do Minikube
minikube start --cpus 4 --memory 8192
sudo chmod 777 /var/run/docker.sock
minikube addons enable metrics-server

#Implementação das APIs e DB
kubectl apply -f vendas-fase4automoveis-deployment.yaml,vendas-fase4automoveis-service.yaml,fase4automoveis-deployment.yaml,fase4automoveis-service.yaml,postgres-initdb-config.yaml,postgres-claim0-persistentvolumeclaim.yaml,postgres-deployment.yaml,postgres-service.yaml

#Configura HPA
kubectl autoscale deployment vendas-fase4automoveis --cpu-percent=80 --min=1 --max=10
kubectl autoscale deployment fase4automoveis --cpu-percent=80 --min=1 --max=10
kubectl autoscale deployment postgres --cpu-percent=80 --min=1 --max=10