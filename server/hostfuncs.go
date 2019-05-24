package main

import (
	"log"

	"github.com/perlin-network/life/exec"
)

var hostFuncs = map[string]exec.FunctionImport{
	"println": hostPrintln,
}

func hostPrintln(vm *exec.VirtualMachine) int64 {
	var msg string
	fatalErr(wasmArgs(vm, &msg), "hostPrintln args")
	log.Printf(">>> %q", msg)
	return 0
}
