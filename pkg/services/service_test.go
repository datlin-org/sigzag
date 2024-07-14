package services

import "testing"

func TestRunService(t *testing.T) {
	err := RunService(":7500")
	if err != nil {
		return
	}
}
