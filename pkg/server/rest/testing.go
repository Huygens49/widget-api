package rest

import (
	"fmt"

	"github.com/Huygens49/widget-api/pkg/saving"
)

func Okay() {
	s := saving.Widget{Description: "Test"}
	fmt.Println(s)
}
