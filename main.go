package main

import (
	"xo_connect5/infra/http"
)

func main() {
	http.Router().Run("127.0.0.1:8080")
}
