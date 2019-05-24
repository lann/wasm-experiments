package main

import "C"
import (
	"io/ioutil"
	"log"

	"github.com/perlin-network/life/exec"
)

const payloadFile = "../payload/target/wasm32-unknown-unknown/debug/payload.wasm"

func main() {
	code, err := ioutil.ReadFile(payloadFile)
	fatalErr(err, "Readfile failed")

	vm, err := exec.NewVirtualMachine(code, exec.VMConfig{}, resolver{}, nil)
	fatalErr(err, "NewVirtualMachine failed")

	entryID, ok := vm.GetFunctionExport("exported")
	if !ok {
		log.Fatal("no exported function!")
	}

	ret, err := vm.Run(entryID)
	fatalErr(err, "Run failed")
	if ret != 0 {
		log.Printf("ret = %v (%T)", ret, ret)
	}
}

type resolver struct {
	*exec.NopResolver
}

func (resolver) ResolveFunc(module, field string) exec.FunctionImport {
	var hostFunc exec.FunctionImport
	if module == "env" {
		hostFunc = hostFuncs[field]
	}
	if hostFunc == nil {
		log.Panicf("cannot resolve func %s::%s", module, field)
	}
	return hostFunc
}

func fatalErr(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %v", msg, err)
	}
}
