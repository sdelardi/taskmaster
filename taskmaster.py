#!/usr/bin/python
import commandsController
import parse_config
def set_programs():
	global programs
	programs = parse_config.parsing("config.conf")

def switchMenu(choice):

	#User input is splitted into a command and its parameters
	choice = choice.split(" ")

	command = choice[0]
	parametersSize = len(choice) - 1
	parameters = []

	if(parametersSize > 0):
		for i in range(1, parametersSize + 1):
			parameters.append(choice[i])

	#Debug only
	print("Command: " + command)
	print("Parameters: " + str(parameters or ""))

	#Switch
	if(command == "exit"):
		commandsController.quitController()
		
	elif(command == "status"):
		commandsController.statusController(parametersSize, parameters)
		
	elif(command == "start"):
		commandsController.startController(parametersSize, parameters)

	elif(command == "stop"):
		commandsController.stopController(parametersSize, parameters)

	elif(command == "restart"):
		commandsController.restartController(parametersSize, parameters)

	elif(command == "reload"):
		commandsController.reloadController()

	else:
		commandsController.defaultController()

##########################################################################

print(" -- Start ! -- ")
set_programs()
while True:
	
    switchMenu(input("--> "))



