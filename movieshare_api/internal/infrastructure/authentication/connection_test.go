package authentication

import (
	"testing"
)

func TestFirebaseConnection(t *testing.T) {
	app:= getFirebaseConnection()
	// if err != nil {
	// 	t.Fatalf("%v", err)
	// }
	t.Logf("%v", app)
}

