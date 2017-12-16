package main

func main() {
	var programs []program;
	var program_database map[string]program;
	var process_database map[string][]process;

	parse_config("./config.xml", &programs);
	init_database(programs, &program_database, &process_database);
}
