import requests

url = "http://localhost:9000"

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

######################### UPDATING FIRST RAW MATERIAL #######################
raw = {"name":"Madeira da boa", "inventory":90}
r = requests.put(url+"/raw_material/1", json=raw)
print(raw.text)

r = requests.get(url+"/raw_material/1")
print(r.text)
###############################################################################

########################## REMOVING SECOND RAW MATERIAL ###########################
r = requests.delete(url+"/raw_material/2")

print(r.text)
###################################################################################
