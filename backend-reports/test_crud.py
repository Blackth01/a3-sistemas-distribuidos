import requests

url = "http://localhost:9001"

#################### REPORT ##########################
r = requests.get(url+"/report")
print(r.text)
######################################################