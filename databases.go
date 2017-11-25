package main

func create_program_database(program_database *map[string]program) {
	*program_database = make(map[string]program);
}
func create_process_database(process_database *map[string][]process) {
	*process_database = make(map[string][]process);
}

func fill_program_database(programs []program, program_database *map[string]program) {
	for i := 0; i < len(programs); i++ {
		(*program_database)[programs[i].name] = programs[i];
	}
}

func fill_process_database(programs []program , process_database *map[string][]process) {
	for  i := 0; i < len(programs); i++ {
		(*process_database)[programs[i].name] = make([]process, 0, programs[i].num_procs);
	}
}

func init_database(programs []program, program_database *map[string]program, process_database *map[string][]process) {
	create_program_database(program_database);
	create_process_database(process_database);
	fill_program_database(programs, program_database);
	fill_process_database(programs, process_database);
}
