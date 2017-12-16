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

func parse_config(path string) {
	v := Elements{};
	dat, err := ioutil.ReadFile(path);
	check(err);
	e := xml.Unmarshal([]byte(dat), &v);
	check(e);
}
