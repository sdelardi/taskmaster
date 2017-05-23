import yaml

def parseConfig(file):
	toto = yaml.load(file)
	print(toto)

def openConfig(file_name):
	try:
		file = open(file_name, "r")
	except FileNotFoundError:
		print("File not found")
		exit(2)
	parseConfig(file)

openConfig("config.conf")
