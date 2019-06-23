---
to: main.go
---

package main

import (
	"<%= repo_path %>/<%= name %>/cmd"
)

func main() {
	cmd.Execute()
}
