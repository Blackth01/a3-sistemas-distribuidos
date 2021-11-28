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



################################## ADDING PRODUCT #################################
product = {"name":"Caneta", "value":1900.00}
r = requests.post(url+"/product", json=product)

print(r.text)
###################################################################################

################################### GETTING ALL PRODUCTS ###########################
product = {"name":"Lápis", "value":10.50}
r = requests.post(url+"/product", json=product) # ADDING ANOTHER PRODUCT

r = requests.get(url+"/product")
print(r.text)
####################################################################################

################################## GETTING SECOND PRODUCT #############################
r = requests.get(url+"/product/2")
print(r.text)
#######################################################################################

#################################### UPDATING FIRST PRODUCT ############################
product = {"name":"Caneta", "value":30.90}
r = requests.put(url+"/product/1", json=product)
print(r.text)
########################################################################################

#################################### REMOVING SECOND PRODUCT ############################
r = requests.delete(url+"/product/2")
print(r.text)
#########################################################################################


######################### ADDING INPUT #################################################
input = {"product":1, "material":1, "quantity":30}
r = requests.post(url+"/input", json=input)

print(r.text)
#######################################################################################


############################## UPDATING INPUT ##########################################
input = {"product":1, "material":1, "quantity":25}
r = requests.put(url+"/input", json=input)

print(r.text)
##############################################################################################

########################### DELETING INPUT #####################################################
input = {"product":1, "material":1}
r = requests.delete(url+"/input", json=input)
print(r.text)
################################################################################################




######################## GETTING RAW MATERIALS BY PRODUCT #########################################
r = requests.get(url+"/raw_material/by_product/1")
print(r.text)
################################################################################################

