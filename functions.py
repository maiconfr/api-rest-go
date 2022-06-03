import json
import requests

# Função para retornar somente números
def limpaCPF(cpf):
    cpfNovo = ""
    for i in cpf:
        if i.isdigit():
            cpfNovo = cpfNovo + i
    return cpfNovo

# Requisição post para a API
def reqPost(data):
    url = "http://localhost:5000/patient"

    payload = json.dumps(data)
    headers = {
    'Content-Type': 'application/json'
    }

    response = requests.request("POST", url, headers=headers, data=payload)

    print(response.text)
    return response.text