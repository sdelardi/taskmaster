#!/usr/bin/python
import commands
import argparse

def switchMenu(choice):

	#split input into a list. Command will be in argv[0] and parameters after
	argv = choice.split(" ");
	argc = len(argv)

	if(argv[0] == "exit"):
		if(argc == 1):
			commands.quit();
		else:
			print("/!\\ Error : exit command doesn't take parameters /!\\")
	
	elif(argv[0] == "status"):
		if(argc <= 3):
			commands.status(argv[1], argv[2]);
		else:
			print("/!\\ Error : status command doesn't take more than 2 parameters /!\\")
		
	elif(argv[0] == "start"):
		commands.start();

	elif(argv[0] == "stop"):
		commands.stop();

	elif(argv[0] == "restart"):
		commands.restart();

	elif(argv[0] == "reload"):
		commands.reload();

	else:
		commands.default();

##########################################################################

print(" -- Start ! -- ")

while True:
	
    switchMenu(input("--> "))



