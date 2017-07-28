# goCep

Projeto em Golang para buscar dados do CEP e armazenar em cache 


# Compilar

Além de ter o Go instalado no sistema operacional é necessário executar:

	export GOPATH=/path/goCep
	go get -u github.com/golang/dep/cmd/dep
	dep ensure
	go build


Se você está compilando este projeto no Ubuntu, ou no Debian, verifique se você possuí 
uma das bibliotecas abaixo instaladas:

* libcurl4-gnutls-dev
* libcurl4-openssl-dev


# Executar

O binário chamado goCep será criado. Basta executá-lo e ele ficará ouvindo na porta 3000 por novas requisições


# Testando
Para rodar os tests, você pode executar:

	go test


# Uso

Basta acessar a URL como no exemplo abaixo

	http://localhost:3000/cep/89201405

O retorno será um JSON com o conteúdo 
