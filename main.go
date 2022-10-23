package falcata

import (
	"strings"
)

func CutFields(matl *string) {
	*matl = strings.Join(strings.Fields(*matl), " ")
}