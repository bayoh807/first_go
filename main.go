package main

import (
	_ "gf_0620/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"gf_0620/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
