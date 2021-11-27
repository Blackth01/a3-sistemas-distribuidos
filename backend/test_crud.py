import requests

#################### ADDING RAW MATERIAL #############
raw = {"name":"Madeira da boa", "inventory":15}
r = requests.post(url+"/raw_material", json=raw)
print(r.text)
######################################################



################## GETTING ALL RAW MATERIALS #############
r = requests.get(url+"/raw_material")
print(r.text)
##########################################################


##################### GETTING ONE RAW MATERIAL ################
raw = {"name":"Plastico", "inventory":11}
r = requests.post(url+"/raw_material", json=raw) #ADDING ONE MORE RAW MATERIAL

r = requests.get(url+"/raw_material/2")

print(r.text)
################################################################
