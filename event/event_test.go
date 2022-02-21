// Dette er en test

package event

import "testing"

func TestPut(t *testing.T) {
// Hva forventer jeg?
    wanted := "[snill mann+dyr og korn"
// Hva fikk jeg?
    got := Put("dyr og korn")

    if got != wanted {
             t.Errorf("Feil, fikk %q, ønsket %q.", got, wanted)
    }
}