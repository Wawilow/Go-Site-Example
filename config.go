package main

type cnf struct {
	port int
}

func config() cnf {
	return cnf{50}
}
