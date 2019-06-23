---
to: main.go
---

package main

import (
	"gitlab.appsflyer.com/Architecture/<%= name %>/cmd"
)

func main() {
	cmd.Execute()
}
