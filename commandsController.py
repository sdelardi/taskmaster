import commands

#Control commands.quit()
def quitController():
	commands.quit()

#Control commands.status(program=None, pid=None)
def statusController(parametersSize, parameters):

	if(parametersSize == 0):
		commands.status()
	elif(parametersSize == 1):
		commands.status(parameters[0])
	elif(parametersSize == 2):
		commands.status(parameters[0], parameters[1])
	else:
		commands.statusHelp()

#Control commands.start(program)
def startController(parametersSize, parameters):
	if(parametersSize == 1):
		commands.start(parameters[0])
	else:
		commands.startHelp()

#Control commands.stop(program, pid=None)
def stopController(parametersSize, parameters):
	if(parametersSize == 1):
		commands.stop(parameters[0])
	elif(parametersSize == 2):
		commands.stop(parameters[0], parameters[1])
	else:
		commands.stopHelp()

#Control commands.restart(program, pid=None)
def restartController(parametersSize, parameters):
	if(parametersSize == 1):
		commands.restart(parameters[0])
	elif(parametersSize == 2):
		commands.restart(parameters[0], parameters[1])
	else:
		commands.restartHelp()

#Control commands.reload()
def reloadController():
	commands.reload()

#Control commands.default()
def defaultController():
	commands.default()
