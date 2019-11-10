package testgosemver

import (
	"fmt"
	"os"

	"github.com/coreos/go-semver/semver"
)

/**
 * 测试 semver
 * go run main.go 1.2.1 2.3.1
 */
func TestGoSemver() {
	vA, err := semver.NewVersion(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
	}
	vB, err := semver.NewVersion(os.Args[2])
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%s < %s == %t\n", vA, vB, vA.LessThan(*vB))
}
