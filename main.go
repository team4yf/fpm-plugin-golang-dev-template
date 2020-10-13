package main

import (
	"github.com/team4yf/yf-fpm-server-go/fpm"

	_ "github.com/team4yf/fpm-go-plugin-{{name}}/plugin"
)

func main() {
	fpmApp := fpm.New()

	fpmApp.Init()

	fpmApp.Run()
}
