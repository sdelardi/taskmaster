package main

func main() {
	var programs []program;
	parse_config("./config.xml", &programs);
	var program_database map[string]program;
	var process_database map[string][]process;
	init_database(programs, &program_database, &process_database);
}
