#!/usr/bin/python
import commands

def switchMenu(choice):
	if(choice == "exit"):
		commands.quit();
	else:
		commands.default();

print(" -- Start ! -- ")

while True:
    switchMenu(input("--> "))



