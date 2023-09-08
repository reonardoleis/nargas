package main

import (
	"encoding/json"
	"os"
	"os/exec"
	"testing"

	"github.com/reonardoleis/nargas/internal/ast"
	"github.com/reonardoleis/nargas/internal/transpiler"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func cleanup() {
	os.Remove("../tests/main.go")
}

func TestTranspile(t *testing.T) {
	defer cleanup()
	assert := assert.New(t)
	require := require.New(t)

	tests := []struct {
		name           string
		inputFilePath  string
		expectedOutput string
	}{
		{"sum", "../rinha/sum.json", "15\n"},
		{"fib", "../rinha/fib.json", "55\n"},
		{"combination", "../rinha/combination.json", "45\n"},
	}

	for _, test := range tests {

		inputFile, err := os.ReadFile(test.inputFilePath)
		require.NoError(err)

		ast := ast.AST{}
		err = json.Unmarshal(inputFile, &ast)
		if err != nil {
			panic(err)
		}

		transpiler := transpiler.NewTranspiler(&ast, transpiler.Go)
		out := transpiler.Transpile()

		outputFileName := "../tests/main.go"
		outputFile, err := os.Create(outputFileName)
		require.NoError(err)
		defer outputFile.Close()

		_, err = outputFile.WriteString(out)
		require.NoError(err)

		programOutput, err := exec.Command("go", "run", "main.go").Output()
		require.NoError(err)

		assert.Equal(test.expectedOutput, string(programOutput))

	}
}
