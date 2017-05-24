#to stop taskmaster manually
def quit():
	#ask user to confirm program closure
	if(input("/!\\ Information : Taskmaster is about to close, press y to quit /!\\\n--> ") == "y"):
		exit(0);
	else:
		print("/!\\ Information : Closure canceled /!\\ ")

#status of all the programs described in the config file, or a specific program, or a specific instance of a program
def status(program=None, pid=None):
	print(program)
	print(pid);

#start a new instance of a program
def start(program):
	print("TODO");

#stop a program, or a specific instance of a program
def stop(program, pid=None):
	print("TODO");

#stop a program or a specific instance of a program, then restart a new instance of himself
def restart(program, pid=None):
	stop(program, pid)
	start(program);

#reload the configuration file
def reload():
	print("TODO");

#If the written command is unknown
def default():
	print("Unknown command");