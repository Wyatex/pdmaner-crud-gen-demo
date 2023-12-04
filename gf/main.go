package main

import (
	_ "pdmaner-test/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "pdmaner-test/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"pdmaner-test/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
