package main

import "encoding/xml";
import "io/ioutil";
import "errors";
import "strconv";
import "strings";

func parse_config(path string, programs *[]program) {
	e := Elements{};
	dat, err := ioutil.ReadFile(path);
	rise_error(err);
	err = xml.Unmarshal([]byte(dat), &e);
	rise_error(err);
	fill_program_array(programs, &e);
}

func check_name(Name string) string {
	if Name == "" {
		rise_error(errors.New("Name is empty"));
	}
	for i := 0; i < len(Name); i++ {
		c := Name[i];
		if !((c >= 48 && c <= 57) || (c >= 65 && c <= 90) || (c >= 97 && c <= 122)) {
			rise_error(errors.New("Name is not correctly formatted"));
		}
	}
	return Name;
}

func check_cmd(Cmd string) string {
	if Cmd == "" {
		rise_error(errors.New("Cmd is empty"))
	}
	return Cmd;
}

func check_num_procs(Num_procs string) uint {
	var num_procs int;
	var err error;

	if Num_procs == "" {
		rise_error(errors.New("Num_procs is empty"));
	}
	num_procs, err = strconv.Atoi(Num_procs);
	rise_error(err);
	if num_procs <= 0 {
		rise_error(errors.New("Num_procs is <= 0"));
	}
	return uint(num_procs);
}

func check_autostart(Autostart string) bool {
	if Autostart == "true" {
		return true;
	} else if Autostart == "false" || Autostart == "" {
		return false;
	} else {
		rise_error(errors.New("Autostart has wrong value"));
	}
	return false;
}

func check_restart(Restart string) string {
	if Restart == "" {
		rise_error(errors.New("Restart is empty"));
	}
	if !(Restart == "unexpected" || Restart == "always" || Restart == "never") {
		rise_error(errors.New("Restart has wrong value"));
	}
	return Restart;
}

func check_exit_codes(Exit_codes string) []int {
	var codes []string;
	var int_codes []int;
	codes = strings.Split(Exit_codes, ",");
	int_codes = make([]int, len(codes));
	for i := 0; i < len(codes); i++ {
		res, err := strconv.Atoi(codes[i]);
		rise_error(err);
		int_codes[i] = res;
	}
	return int_codes;
}

func check_starting_time(Starting_time string) uint {
	var starting_time int;
	var err error;

	if Starting_time == "" {
		return 0;
	}
	starting_time, err = strconv.Atoi(Starting_time);
	rise_error(err);
	if starting_time < 0 {
		rise_error(errors.New("Starting_time is negative"));
	}
	return uint(starting_time);
}

func check_retry_abort(Retry_abort string) uint {
	var trys int;
	var err error;
	if Retry_abort == "" {
		rise_error(errors.New("Retry_abort is empty"));
	}
	trys, err = strconv.Atoi(Retry_abort);
	rise_error(err);
	if trys < 0 {
		rise_error(errors.New("Retry_abort is < 0"));
	}
	return uint(trys);
}

func check_legit_signal(Legit_signal string) string {
	if Legit_signal == "" {
		rise_error(errors.New("Legit_signal is empty"));
	}
	for i := 0; i < len(Legit_signal); i++ {
		c := Legit_signal[i];
		if !((c >= 48 && c <= 57) || (c >= 65 && c <= 90)) {
			rise_error(errors.New("Legit_signal is not correctly formatted"));
		}
	}
	return Legit_signal;
}

func check_time_kill(Time_kill string) uint {
	var time_kill int;
	var err error;

	if Time_kill == "" {
		return 0;
	}
	time_kill, err = strconv.Atoi(Time_kill);
	rise_error(err);
	if time_kill < 0 {
		rise_error(errors.New("Time_kill is < 0"));
	}
	return uint(time_kill);
}

func check_stdout(Stdout string) string {
	if !(Stdout == "discard" || Stdout == "file" || Stdout == "") {
		rise_error(errors.New("Stdout has wrong value"));
	}
	return Stdout;
}

func check_stdout_file(Stdout_file string) string {
	return Stdout_file;
}

func check_stderr(Stderr string) string {
	if !(Stderr == "discard" || Stderr == "file" || Stderr == "") {
		rise_error(errors.New("Stderr has wrong value"));
	}
	return Stderr;
}

func check_stderr_file(Stderr_file string) string {
	return Stderr_file;
}

func check_env_vars(Env_vars string) []string {
	var env []string;
	if Env_vars == "" {
		return nil;
	}
	env = strings.Split(Env_vars, ",");
	return env;
}

func check_work_dir(Work_dir string) string {
	return Work_dir;
}

func check_umask(Umask string) uint {
	var umask int;
	var err error;
	if Umask == "" {
		return 8;
	}
	for i := 0; i < len(Umask); i++ {
		c := Umask[i];
		if !(c >= 48 && c <= 55) {
			rise_error(errors.New("Umask is not valid"));
		}
	}
	umask, err = strconv.Atoi(Umask);
	rise_error(err);
	return uint(umask);
}

func fill_program_array(programs *[]program, e *Elements) {
	*programs = make([]program, len(e.Elements));
	for i := 0; i < len(e.Elements); i++ {
		(*programs)[i].name = check_name(e.Elements[i].Name);
		(*programs)[i].cmd = check_cmd(e.Elements[i].Cmd);
		(*programs)[i].num_procs = check_num_procs(e.Elements[i].Num_procs);
		(*programs)[i].autostart = check_autostart(e.Elements[i].Autostart);
		(*programs)[i].restart = check_restart(e.Elements[i].Restart);
		(*programs)[i].exit_codes = check_exit_codes(e.Elements[i].Exit_codes);
		(*programs)[i].starting_time = check_starting_time(e.Elements[i].Starting_time);
		(*programs)[i].retry_abort = check_retry_abort(e.Elements[i].Retry_abort);
		(*programs)[i].legit_signal = check_legit_signal(e.Elements[i].Legit_signal);
		(*programs)[i].time_kill = check_time_kill(e.Elements[i].Time_kill);
		(*programs)[i].stdout = check_stdout(e.Elements[i].Stdout);
		(*programs)[i].stdout_file = check_stdout_file(e.Elements[i].Stdout_file);
		(*programs)[i].stderr = check_stderr(e.Elements[i].Stderr);
		(*programs)[i].stderr_file = check_stderr_file(e.Elements[i].Stderr_file);
		(*programs)[i].env_vars = check_env_vars(e.Elements[i].Env_vars);
		(*programs)[i].work_dir = check_work_dir(e.Elements[i].Work_dir);
		(*programs)[i].umask = check_umask(e.Elements[i].Umask);
	}
}
