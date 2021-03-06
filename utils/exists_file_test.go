package utils_test

import (
	"testing"

	"github.com/coffemanfp/beppin-server/utils"
)

func TestExistsFile(t *testing.T) {
	exists, err := utils.ExistsFile("../main.go")
	if err != nil {
		t.Errorf("unexpected error:\n%s", err)
	}

	if !exists {
		t.Errorf("exists value (%t) invalid, expected (%t)", exists, true)
	}
}

func TestFailExistsFile(t *testing.T) {
	exists, err := utils.ExistsFile("../main.go2")
	if err != nil {
		t.Errorf("unexpected error:\n%s", err)
	}

	if exists {
		t.Errorf("exists value (%t) invalid, expected (%t)", exists, false)
	}
}
