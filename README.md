# go-expert-client-server-api

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