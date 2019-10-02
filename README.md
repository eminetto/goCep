# goCep

Projeto em Golang para buscar dados do CEP e armazenar em cache 


# Compilar

Além de ter o Go instalado no sistema operacional é necessário executar:

	go build

# Executar

O binário chamado goCep será criado. Basta executá-lo e ele ficará ouvindo na porta 3000 por novas requisições


# Testando
Para rodar os tests, você pode executar:

	go test

# Uso

Basta acessar a URL como no exemplo abaixo

	http://localhost:3000/cep/89201405

O retorno será um JSON com o conteúdo 
