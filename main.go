package main

import (
	"fmt"
	"golang.org/x/sys/windows/registry"
	"log"
)

func main() {
	fmt.Println("Checking to see what versions of .NET are installed...")

	//	Using proceedure outlined here:
	//	https://msdn.microsoft.com/en-us/library/hh925568%28v=vs.110%29.aspx?f=255&MSPPError=-2147217396

	//	First, check to see if we have version 1-4 of .NET installed
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\NET Framework Setup\NDP`, registry.ENUMERATE_SUB_KEYS|registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	subkeys, err := k.ReadSubKeyNames(-1)
	if err != nil {
		log.Fatal(err)
	}

	//	If we have subkeys, indicate what versions are installed:
	for _, subkey := range subkeys {
		fmt.Printf("Found .NET %v\n", subkey)
	}

	fmt.Println("Finished.")
}
