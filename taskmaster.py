#!/usr/bin/python
import commands

def switchMenu(choice):

	if(choice == "exit"):
		commands.quit();
	
	elif(choice == "status"):
		commands.status();
	
	elif(choice == "start"):
		commands.start();

	elif(choice == "stop"):
	commands.stop();

	elif(choice == "restart"):
	commands.restart();

	elif(choice == "reload"):
		commands.reload();

	else:
		commands.default();

##########################################################################

print(" -- Start ! -- ")

while True:
    switchMenu(input("--> "))



