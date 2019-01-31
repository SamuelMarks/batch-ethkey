package main

type arrayHosts []string

func (a *arrayHosts) String() string {
	return "list of hosts"
}

func (a *arrayHosts) Set(value string) error {
	*a = append(*a, value)
	return nil
}
