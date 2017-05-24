class Program:
	"""Attributs de la classe :
	- name // Name
	- cmd // Command to launch the program
	- numprocs // Numbers of processes
	- autostart // Launch at start
	- restart // Program should be restarted always, never, or on unexpected exits
	- tts // Time to success
	- nbres // Numbers of restarts before abort
	- signal // signal to stop the program
	- stoptime // Time waiting for clean exit before killing the program
	- env // Environment variables to set before launching the program
	- wdir // Working directory to set before launching the program
	- stdout // File to redirect standard output
	- stderr // File to redirect error output
	- umask // Umask to set before launching the program"""

	def __init__(self, program, key):
			try:
				program["cmd"]
			except KeyError:
				program["cmd"] = None
			try:
				program["numprocs"]
			except KeyError:
				program["numprocs"] = None
			try:
				program["autostart"]
			except KeyError:
				program["autostart"] = None
			try:
				program["restart"]
			except KeyError:
				program["restart"] = None
			try:
				program["tts"]
			except KeyError:
				program["tts"] = None
			try:
				program["nbres"]
			except KeyError:
				program["nbres"] = None
			try:
				program["signal"]
			except KeyError:
				program["signal"] = None
			try:
				program["stoptime"]
			except KeyError:
				program["stoptime"] = None
			try:
				program["env"]
			except KeyError:
				program["env"] = None
			try:
				program["wdir"]
			except KeyError:
				program["wdir"] = None
			try:
				program["stdout"]
			except KeyError:
				program["stdout"] = None
			try:
				program["stderr"]
			except KeyError:
				program["stderr"] = None
			try:
				program["umask"]
			except KeyError:
				program["umask"] = None
			self.name = key
			self.cmd = program["cmd"]
			self.numprocs = program["numprocs"]
			self.autostart = program["autostart"]
			self.restart = program["restart"]
			self.tts = program["tts"]
			self.nbres = program["nbres"]
			self.signal = program["signal"]
			self.stoptime = program["stoptime"]
			self.env = program["env"]
			self.wdir = program["wdir"]
			self.stdout = program["stdout"]
			self.stderr = program["stderr"]
			self.umask = program["umask"]
