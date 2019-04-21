package arquitetura

import (
	"runtime"
	"testing"
)

func TestDependente(t *testing.T) {
	t.Parallel() // rodar testes em paralelo
	if runtime.GOARCH == "amd64" {
		t.Skip("Não funciona em arquitetura amd64")
	}
	//..
	t.Fail()
}
