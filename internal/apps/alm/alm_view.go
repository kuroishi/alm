package alm_view

import (
    "fmt"

    "../../../pkg/alm_common"
)

func Display_month(month alm_common_type.Month) string {
    fmt.Print(month)
    return "ok"
}
