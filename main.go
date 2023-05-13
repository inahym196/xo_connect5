package main

import (
	"xo_connect5/infra"
)

func main() {
	infra.Router().Run("127.0.0.1:8080")
}
