# go-expert-client-server-api

## Servidor
#### configuração:
Essa configuração é realizada a partir de um arquivo yaml
local no arquivo: [server/internal/config/config.yaml](https://github.com/rprojetos/go-expert-client-server-api/blob/main/server/internal/config/config.yaml)

### Comando para iniciar o servidor:
A partir do diretório raiz do repositório entrar no diretório server
Então execute o comando no terminal:
> ***go run ./cmd/server***

Então, o servidor será iniciado.

## Cliente
#### configuração:
Essa configuração é realizada a partir de um arquivo yaml
local no arquivo: [client/internal/config/config.yaml](https://github.com/rprojetos/go-expert-client-server-api/blob/main/client/internal/config/config.yaml)

### Comando para iniciar o servidor:
A partir do diretório raiz do repositório entrar no diretório client
Então execute o comando no terminal:
> ***go run ./cmd/client***

Então, o client realizará uma requisição de cotação do dollar para o server.


## Resumo do desafio implementado:
Este desafio consiste no desenvolvimento de duas aplicações relativas a server/client
Sendo o client responsavel por buscar no servidor a cotação atual do dolar salvando esta em um arquivo de texto.
O servidor é responsável por buscar a cotação em uma api externa, persistindo a respctiva consulta em um bando de dados sqlite.

## Tópicos de abrangência:

- webserver http
- contextos
- banco de dados
- manipulação de arquivos

## Sistemas criados para o desafio:
- server.go
- client.go

## Requisitos do server.go:
- Endereço do servidor / endpoint para cotação do dólar:
    - Porta:
    `8080`
    - Endpoint:
    `/cotacao`
- Consumir a API contendo o câmbio de Dólar e Real:
    - Endereço: 
    `https://economia.awesomeapi.com.br/json/last/USD-BRL`
    - Exemplo JSON retornado:
    ```json
    code	"USD"
    codein	"BRL"
    name	"Dólar Americano/Real Brasileiro"
    high	"5.7452"
    low	"5.7442"
    varBid	"0.001"
    pctChange	"0.017405"
    bid	"5.7452"
    ask	"5.7472"
    timestamp	"1741987800"
    create_date	"2025-03-14 18:30:00"
    ```
    - Context / Timeout: 
    `200 milissegundos`
    `O contexto deverá retornar erro nos logs caso o tempo de execução seja insuficiente.`
- Registrar cada cotação recebida no banco de dados SQLite.
    - Context / Timeout:
    `10 milissegundos`
    `O contexto deverá retornar erro nos logs caso o tempo de execução seja insuficiente.`
- Retornar no formato JSON o resultado para o cliente.
    - O conteúdo do JSON deve ser:
    `Apenas o campo "bid"`

## Requisitos do client.go:
- Requisição HTTP no server.go solicitando a cotação do dólar.
    - Porta:
    `8080`
    - Endpoint:
    `/cotacao`
- Context / Timeout:
    `300 milissegundos`
    `O contexto deverá retornar erro nos logs caso o tempo de execução seja insuficiente.`
- A resposta deve ser no formato JSON. 
    - O conteúdo do JSON deve ser:
    `Apenas o campo "bid"`

- Salvar o valor da cotação em um arquivo `cotacao.txt`
    - O conteúdo do arquivo deve ser:
    `Dólar: {valor da cotação}`

## Escopo do projeto:

<img title="escopo" alt="escopo do projeto" src="/escopo/clienteServer.svg">

## Telas: App Server / App Client 

<img title="App Server / App Client" alt="App Server / App Client" src="/escopo/tela.svg">
