package build

import (
	"os"
	"os/exec"
	"strings"

	"github.com/reonardoleis/nargas/internal/transpiler"
)

func Generate(code, outputFileName string, outputType transpiler.OutputType) {
	fileName := "./generated/" + outputFileName + "." + strings.ToLower(string(outputType))

	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	_, err = f.WriteString(code)
	if err != nil {
		panic(err)
	}

	err = exec.Command("go", "build", "-o", "./generated/a.out", fileName).Run()
	if err != nil {
		panic(err)
	}
}
