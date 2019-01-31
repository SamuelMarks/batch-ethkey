package main

import "fmt"

type arrayHosts []string

func (a *arrayHosts) String() string {
	panic(fmt.Errorf("not implemented"))
}

func (a *arrayHosts) Set(value string) error {
	*a = append(*a, value)
	return nil
}
