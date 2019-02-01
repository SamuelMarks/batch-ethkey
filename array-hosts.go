package main

import (
	"strings"
)

type arrayHosts []string

func (a *arrayHosts) String() string {
	return strings.Join(*a,"\t")
}

func (a *arrayHosts) Set(value string) error {
	*a = append(*a, value)
	return nil
}
