# FASE 3 - Revenda de veículos

### SOBRE:

### IMPLEMENTAÇÃO:

### DEPLOY:

**Pré requisitos:**
kubectl v1.29.0
minikube v1.32.0
go 1.21.3

**Como implementar:**
Executar o script deploy.sh

Ao executar o script, será deletado o cluster existente do Minikube e criado um novo, ajuste nas permissões do docker.sock, habilita o metrics-server, realiza o deploy da aplicação no cluster e configuração do HPA.

Importante aguardar alguns minutos para conclusão da criação dos pods antes de testar.

**Comando para acompanhar o status da implantação:**
watch minikube kubectl get pods

### COMO TESTAR:
