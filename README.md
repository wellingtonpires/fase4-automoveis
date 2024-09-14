# FASE 3 - Revenda de veículos

### SOBRE:

Backend para uma plataforma para revenda de veículos, implementada utilizando Golang, Postgres e Kubernetes através do Minikube. Toda a parte de autenticação foi desenvolvida de forma customizada e separada das funcionalidades do negócio, implementada utilizando JWT.

### DEPLOY:

**Pré requisitos:**\
kubectl v1.29.0\
minikube v1.32.0\
go 1.21.3

**Como implementar:**\
Executar o script deploy.sh

Ao executar o script, será deletado o cluster existente do Minikube e criado um novo, ajuste nas permissões do docker.sock, habilita o metrics-server, realiza o deploy da aplicação no cluster e configuração do HPA.

Importante aguardar alguns minutos para conclusão da criação dos pods antes de testar.

**Comando para acompanhar o status da implantação:**\
watch minikube kubectl get pods

### COMO TESTAR:

Importar o arquivo 'Sub Fase 3 - Venda de veículos.postman_collection.json' para o Postman.

Para utilizar as APIs, é necessário cadastrar um usuário na API '/usuario/cadastra-usuario' e em seguida gerar um token jwt utilizando a API '/usuario/cria-token?login=LOGINUSUARIO&senha=SENHAUSUARIO', preenchendo o login e senha com os dados do usuário criado anteriormente.

Para testar as APIs, é necessário informar o token gerado no campo 'Authorization'. 

Apenas usuários criados com a flag "admin = true" tem permissão para utilizar APIs de cadastro, atualização, deleção e consulta de veículos vendidos.
