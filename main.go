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
	oldKey, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\NET Framework Setup\NDP`, registry.ENUMERATE_SUB_KEYS|registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer oldKey.Close()

	//	Enumerate the subkeys to find the versions installed
	oldVersions, err := oldKey.ReadSubKeyNames(-1)
	if err != nil {
		log.Fatal(err)
	}

	//	If we have old versions, indicate that we found some
	if len(oldVersions) > 0 {
		fmt.Println("\nOlder .NET versions installed:\n============================")
	}

	//	If we have old versions, indicate what versions are installed:
	for _, oldVersion := range oldVersions {
		if oldVersion != "CDF" {
			fmt.Printf("%v\n", oldVersion)
		}
	}

	//	Next, check to see if we have version 4+ of .NET installed
	newKey, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\NET Framework Setup\NDP\v4\Full`, registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer newKey.Close()

	releaseValue, _, _ := newKey.GetIntegerValue("Release")

	if releaseValue > 0 {
		fmt.Println("\nNewer .NET versions installed:\n============================")

		switch releaseValue {
		case 378389:
			fmt.Println("v4.5")
		case 378675:
			fmt.Println("v4.5.1 (installed with Windows 8.1 or Windows Server 2012 R2)")
		case 378758:
			fmt.Println("v4.5.1 (installed on Windows 8, Windows 7 SP1, or Windows Vista SP2)")
		case 379893:
			fmt.Println("v4.5.2")
		case 393295:
			fmt.Println("v4.6 (on Windows 10)")
		case 393297:
			fmt.Println("v4.6 (on an operating system other than Windows 10)")
		default:
			fmt.Printf(".NET version unknown.  Release version: %v", releaseValue)
		}
	}

	fmt.Println("\nFinished.")
}
