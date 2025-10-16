package departemen_controller

import (
	"regexp"
)

// Validasi Format Kode Pos
func validateKodePos(kodePos string) bool {
	regex := regexp.MustCompile(`^\d{5}$`)
	return regex.MatchString(kodePos)
}
