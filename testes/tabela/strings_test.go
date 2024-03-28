package strings

import (
	"strings"
	"testing"
)

const msgIndex = "%s (parte: %s) - índices: esperado (%d) <> encontrado (%d)."

func TestIndex(t *testing.T) {
	testes := [] struct {
		texto string
		parte string
		esperado int
	} {
       {"Go é show", "Go", 0},
	   {"", "", 0},
	   {"Opa", "opa", -1},
	   {"David", "i", 3},
	}

	for _, teste := range testes {
		t.Logf("Massa: %v", teste)
		atual := strings.Index(teste.texto, teste.parte)

		if atual != teste.esperado {
			t.Errorf(msgIndex, teste.parte, teste.parte, teste.esperado, atual)
		}
	}
}
