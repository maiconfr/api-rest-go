import json
import requests

from functions import *

# Abrindo arquivo JSON
f = open('dados.json')

data = json.load(f)

#Concateação
fullName = data["firstName"]+" "+data["lastName"]
data["fullName"] = fullName
# Excluindo 'firstName' e 'lastName'
data.pop("firstName")
data.pop("lastName")

# Limpando o campo CPF
data["cpf"] = limpaCPF(data["cpf"])
print("CPF limpo:", data["cpf"])

resposta = reqPost(data)

# Fechando arquivo JSON
f.close()