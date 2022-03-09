# Milestones 1

## Tarefa 1.2

- Fazer um servidor HTTP que receba um XML (NFe v4.00) via POST
  - Retornar em json se o documento foi criado ou se já existia
- Parsear XML numa estrutura que contém os seguintes campos
  - ID `nfeProc→NFe→infNFe→Id (como atributo)`
  - CNPJ do Emissor `nfeProc→NFe→infNFe→emit→CNPJ`
  - Valor total da nota `nfeProc→NFe→infNFe→total→ICMSTot→vNF`
- Salvar estrutura em memória usando mapa

**Material de apoio:** [servidor HTTP](https://pkg.go.dev/net/http), [parser de XML](https://pkg.go.dev/encoding/xml), [banco em memória usando mapa](https://gobyexample.com/maps)

## Tarefa 1.3

- Fazer o servidor HTTP devolver os dados de um XML via GET, através de seu ID

**Material de apoio:** [servidor HTTP](https://pkg.go.dev/net/http)

## Dúvidas

- [X] Como realizar um request com um xml ou json no body pelo Postman? [Resposta](https://stackoverflow.com/questions/47295675/how-do-i-post-xml-data-to-a-webservice-with-postman)
- [X] Como enviar um response com um xml ou json no body pela aplicação? [Resposta](https://golangbyexample.com/json-response-body-http-go/)

## Sugestões

- Criar tarefas ao fim das milestones aplicando o [Coding Style Guide](https://www.notion.so/arquiveiofficial/31735ff16956484a99363e3894d06289?v=035ce194e781401e8d2b8baee6b8a18e) da Casa Stark de forma gradual.