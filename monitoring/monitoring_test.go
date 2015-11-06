package monitoring

import (
	/*	"database/sql"
		"encoding/base64"
		"encoding/json"*/
	"testing"
)

func TestExec(t *testing.T) {
	status := exec()
	if status != 0 {
		t.Error("Exit status is not 0!")
	}

}
