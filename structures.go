package main

import "encoding/xml";

type Element struct {
	XMLName         xml.Name `xml:"Element"`;
	Name            string  `xml:"Name"`;
	Cmd             string  `xml:"Cmd"`;
	Num_procs		string	`xml:"Num_procs"`;
	Autostart       string  `xml:"Autostart"`;
	Restart         string  `xml:"Restart"`;
	Exit_codes      string  `xml:"Exit_codes"`;
	Starting_time   string  `xml:"Starting_time"`;
	Retry_abort     string  `xml:"Retry_abort"`;
	Legit_signal    string  `xml:"Legit_signal"`;
	Time_kill       string  `xml:"Time_kill"`;
	Stdout          string  `xml:"Stdout"`;
	Stdout_file     string  `xml:"Stdout_file"`;
	Stderr          string  `xml:"Stderr"`;
	Stderr_file     string  `xml:"Stderr_file"`;
	Env_vars        string  `xml:"Env_vars"`;
	Work_dir        string  `xml:"Work_dir"`;
	Umask           string  `xml:"Umask"`;

}

type Elements struct {
	XMLName     xml.Name    `xml:"Elements"`;
	Elements    []Element   `xml:"Element"`;
}

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
