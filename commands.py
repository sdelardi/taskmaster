#to stop taskmaster manually
def quit():
	#ask user to confirm program closure
	if(input("/!\\ Information : Taskmaster is about to close, press y to quit /!\\\n--> ") == "y"):
		exit(0)
	else:
		print("/!\\ Information : Closure canceled /!\\ ")

#status of all the programs described in the config file, or a specific program, or a specific instance of a program
def status(program=None, pid=None):
	print("-Status-")
	print("Program: " + str(program or ""))
	print("Pid: " + str(pid or ""))
	print("TODO")

def statusHelp():
	print("status command usage : status \{program\} \{pid\}")

#start a new instance of a program
def start(program):
	print("-Start-")
	print("Program: " + str(program or ""))
	print("TODO")

def startHelp():
	print("start command usage : start [program]")

#stop a program, or a specific instance of a program
def stop(program, pid=None):
	print("-Stop-")
	print("Program: " + str(program or ""))
	print("Pid: " + str(pid or ""))
	print("TODO")

def stopHelp():
	print("stop command usage : stop [program] \{pid\}")

#stop a program or a specific instance of a program, then restart a new instance of himself
def restart(program, pid=None):
	print("-Restart-")
	stop(program, pid)
	start(program)

def restartHelp():
	print("restart command usage : restart [program] \{pid\}")

#reload the configuration file
def reload():
	print("-Reload-")
	print("TODO")

#If the written command is unknown
def default():
	print("Unknown command")