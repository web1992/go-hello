package main

import (
	ostest "web1992/ostest"
	testgosemver "web1992/testgosemver"
	vartest "web1992/var"
)

func main() {
	println("hello go world.")

	ostest.OsTest()
	println(vartest.GlobaVar)
	testgosemver.TestGoSemver()
}
