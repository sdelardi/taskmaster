package main

import "encoding/xml";
import "io/ioutil";
import "os";
import "fmt";

func check(e error) {
	if e != nil {
		fmt.Printf("error: %v\n", e);
		os.Exit(1);
	}
}

func parse_config(path string, programs *[]program) {
	e := Elements{};
	dat, err := ioutil.ReadFile(path);
	check(err);
	err = xml.Unmarshal([]byte(dat), &e);
	check(err);
	fill_program_array(programs, &e);
}

func check_name(Name string) string {
	//TODO
	return "";
}

func check_cmd(Cmd string) string {
	//TODO
	return "";
}

func check_num_procs(Num_procs string) uint {
	//TODO
	return 0;
}

func check_autostart(Autostart string) bool {
	//TODO
	return true;
}

func check_restart(Restart string) string {
	//TODO
	return "";
}

func check_exit_codes(Exit_codes string) []int {
	//TODO
	return nil;
}

func check_starting_time(Starting_time string) uint {
	//TODO
	return 0;
}

func check_retry_abort(Retry_abort string) uint {
	//TODO
	return 0;
}

func check_legit_signal(Legit_signal string) string {
	//TODO
	return "";
}

func check_time_kill(Time_kill string) uint {
	//TODO
	return 0;
}

func check_stdout(Stdout string) string {
	//TODO
	return "";
}

func check_stdout_file(Stdout_file string) string {
	//TODO
	return "";
}

func check_stderr(Stderr string) string {
	//TODO
	return "";
}

func check_stderr_file(Stderr_file string) string {
	//TODO
	return "";
}

func check_env_vars(Env_vars string) []string {
	//TODO
	return nil;
}

func check_work_dir(Work_dir string) string {
	//TODO
	return "";
}

func check_umask(Umask string) uint {
	//TODO
	return 0;
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
