import json
import requests

# Opening JSON file
f = open('dados.json')

data = json.load(f)

fullName = data["firstName"]+" "+data["lastName"]

data["fullName"] = fullName
data.pop("firstName")
data.pop("lastName")

cpfNovo = ""
for i in data["cpf"]:
    if i.isdigit():
        cpfNovo = cpfNovo + i
print("Find numbers from string:",cpfNovo)
data["cpf"] = cpfNovo
for i in data:
    print(i)

print(data["cpf"])

url = "http://localhost:5000/teste"

payload = json.dumps(data)
headers = {
  'Content-Type': 'application/json'
}

response = requests.request("POST", url, headers=headers, data=payload)

print(response.text)
# Closing file
f.close()
