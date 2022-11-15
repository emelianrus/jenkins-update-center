package main

import (
	"fmt"

	"github.com/emelianrus/jenkins-update-center/pkg/updateCenter"
)

func main() {
	updateCenter := updateCenter.Get("")

	fmt.Println(updateCenter.Plugins["blueocean"].Labels)

}
