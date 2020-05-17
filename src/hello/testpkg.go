package testpkg

import "fmt"

func PrintPkgName() {
	fmt.Printf("package: testpkg")
}

func privateName() {
	fmt.Printf("whatever")
}