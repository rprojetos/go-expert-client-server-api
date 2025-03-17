# go-expert-client-server-api

## Desafio:
Consiste na aplicação de conhecimentos sobre webserver http, contextos,
banco de dados e manipulação de arquivos com Go.

## Tópicos de abrangência:

- webserver http
- contextos
- banco de dados
- manipulação de arquivos

## Desafio:
Criação de dois sistemas em Go:
- client.go
- server.go

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

## Escopo do projeto:

<img title="a title" alt="Alt text" src="/escopo/clienteServer.svg">
