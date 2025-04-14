package runtime

import (
	"fmt"

	lua "github.com/yuin/gopher-lua"
)

func Run(outputPath string) error {
	L := lua.NewState()
	defer L.Close()

	if err := L.DoFile(outputPath); err != nil {
		fmt.Println("Erro:", err)
	}

	return nil
}
