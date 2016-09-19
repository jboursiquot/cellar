//************************************************************************//
// API "cellar": Application Resource Href Factories
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/jboursiquot/cellar/design
// --out=$(GOPATH)/src/github.com/jboursiquot/cellar
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"fmt"
	"strings"
)

// BottleHref returns the resource href.
func BottleHref(bottleID interface{}) string {
	parambottleID := strings.TrimLeftFunc(fmt.Sprintf("%v", bottleID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/bottles/%v", parambottleID)
}
