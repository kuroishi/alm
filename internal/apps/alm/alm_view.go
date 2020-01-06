package alm_view

import (
	"fmt"

	"../../../pkg/alm_common"
)

func Display_month(months []alm_common_type.Month) string {

	for idx, elt := range months {
		fmt.Print(elt)
		fmt.Print("\n")
		fmt.Print("idx: %d", idx)
	}

	return "ok"
}
