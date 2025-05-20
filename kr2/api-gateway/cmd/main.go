package main

import (
	"fmt"

	"github.com/Babushkin05/software-dev-course/kr2/api-gateway/internal/config"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)
}
