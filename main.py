import json
import requests

from functions import *

# Abrindo arquivo JSON
f = open('dados.json')

data = json.load(f)

#Concatenação
fullName = data["firstName"]+" "+data["lastName"]
data["fullName"] = fullName
# Excluindo 'firstName' e 'lastName'
data.pop("firstName")
data.pop("lastName")

# Limpando o campo CPF
data["cpf"] = limpaCPF(data["cpf"])
print("CPF limpo:", data["cpf"])

# Enviando dados para a API
resposta = reqPost(data)
print("Resposta da API: "+resposta)

# Fechando arquivo JSON
f.close()
