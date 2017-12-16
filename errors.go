package main

import "fmt";
import "os";

func rise_error(err error) {
	if err != nil {
		fmt.Printf("error: %v\n", err);
		os.Exit(1);
	}
}
