package main

import (
	"fmt"
	"log"

	"golang.org/x/sys/windows/registry"
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
		case 394254:
			fmt.Println("v4.6.1 (on Windows 10)")
		case 394271:
			fmt.Println("v4.6.1 (on an operating system other than Windows 10)")
		case 394747:
			fmt.Println(".NET Framework 4.6.2 Preview installed on Windows 10 RS1 Preview")
		case 394748:
			fmt.Println(".NET Framework 4.6.2 Preview installed on all other Windows OS versions")
		case 394802:
			fmt.Println(".NET Framework 4.6.2 installed on Windows 10 Anniversary Update")
		case 394806:
			fmt.Println(".NET Framework 4.6.2 installed on all other Windows OS versions")
		case 460798:
			fmt.Println(".NET Framework 4.7 installed on Windows 10 Creators Update")
		case 460805:
			fmt.Println(".NET Framework 4.7 installed on all other OS versions")
		case 461308:
			fmt.Println(".NET Framework 4.7.1 installed on Windows 10 Fall Creators Update")
		case 461310:
			fmt.Println(".NET Framework 4.7.1 installed on all other OS versions")
		case 461808:
			fmt.Println(".NET Framework 4.7.2 Windows 10 April 2018 Update")
		case 461814:
			fmt.Println(".NET Framework 4.7.2 installed on all other OS versions")
		case 528040:
			fmt.Println(".NET Framework 4.8 Windows 10 May 2019 Update")
		case 528049:
			fmt.Println(".NET Framework 4.8 installed on all other OS versions")
		default:
			fmt.Printf(".NET version unknown.  Release version: %v", releaseValue)
		}
	}

	fmt.Println("\nFinished. Press Enter to close.")
	var input string
	fmt.Scanln(&input)
}
