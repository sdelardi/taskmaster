package main

type process struct {
	pid int;
	ppid int;
}

type program struct {
	name string;
	cmd string;
	num_procs uint;
	autostart bool;
	restart string;
	exit_codes []int;
	starting_time uint;
	retry_abort uint;
	legit_signal string;
	time_kill uint;
	stdout string;
	stdout_file string;
	stderr string;
	stderr_file string;
	env_vars []string;
	work_dir string;
	umask uint;
}
