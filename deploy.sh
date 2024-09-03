#!/bin/bash

#Script para encerrar o minikube por completo, criar nova build e realizar o start, assim criando uma implantação limpa

#Remover linha abaixo ao entregar
docker build -t wellingt0npires/fase3automoveis:latest . && docker push wellingt0npires/fase3automoveis:latest

minikube delete
minikube start --ports=8080:8080
sudo chmod 777 /var/run/docker.sock
minikube addons enable metrics-server
kubectl apply -f fase3automoveis-deployment.yaml,fase3automoveis-service.yaml,postgres-initdb-config.yaml,postgres-claim0-persistentvolumeclaim.yaml,postgres-deployment.yaml,postgres-service.yaml
kubectl autoscale deployment fase3automoveis --cpu-percent=80 --min=1 --max=10
kubectl autoscale deployment postgres --cpu-percent=80 --min=1 --max=10