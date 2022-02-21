// Dette er en enkel test.

package event

import "testing"

func TestPut(t *testing.T) {
// Hva forventer jeg?
    wanted := "[kylling og rev W\\_mann+korn_/________/ Ø"
// Hva fikk jeg?
    got := Put("korn")

    if got != wanted {
             t.Errorf("Feil, fikk %q, ønsket %q.", got, wanted)
    }
}