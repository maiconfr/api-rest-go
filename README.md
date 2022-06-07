# api-rest-go

## Executar servidor
O servidor da API RESTful encontra-se na pasta "rest", e para iniciá-lo utilize o comando:
```
go run main.go
```
## Executar código python
Para executar o arquivo python, permaneça na pasta raiz e execute:
```
python3 main.py
```
Veja o resultado da inserção do json no banco de dados "rest/patient.db"
O arquivo "main.py" deve responder o valor do CPF modificado conforme a tarefa, e o JSON do dados caso a inserção no banco de dados tenha sido realizada.
Resposta semelhante conforme abaixo:
```
PS path> python3 main.py
CPF limpo: 12345678901
Resposta da API: {"fullName":"JoÃ£o da Silva","cpf":"12345678901","address":"Avenida Brasil","city":"FlorianÃ³polis","state":"Santa Catarina","phone":"","email":"joaosilva@gmail.com","hospital":"hospital baia sul","cardNo":"7844481124110331","appointmentDate":"2025-06-02"}
```
