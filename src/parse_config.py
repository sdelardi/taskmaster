import yaml
import program

#Parsing du fichier de config
def parseConfig(file):
	datas = yaml.load(file)
	programs = []
	for (key, pro) in datas.items():
		programs.append(program.Program(pro, key))
	return programs

#Opening config file and parse it / Exit in case of file not found
def parsing(file_name):
	try:
		file = open(file_name, "r")
	except FileNotFoundError:
		print("Config file not found")
		exit(2)
	programs = parseConfig(file)
	file.close()
	return programs

parsing("config.conf")
