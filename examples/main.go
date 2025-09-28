package main

import (
	"fmt"

	"github.com/mohar9h/irvalidate/formatter"
	"github.com/mohar9h/irvalidate/locales"
)

func main() {
	// load fa & en messages
	_ = locales.LoadMessages("fa")
	_ = locales.LoadMessages("en")

	// fetch message
	raw, _ := locales.GetMessage("fa", "persian_date_between")

	// format
	msg := formatter.FormatMessage(raw.(string), map[string]interface{}{
		"attribute": "تاریخ تولد",
		"startDate": "1400/01/01",
		"endDate":   "1402/12/29",
	}, "، ") // Persian uses "،" instead of ","

	fmt.Println(msg)

	// example with list
	rawDays, _ := locales.GetMessage("fa", "persian_day")
	msgDays := formatter.FormatMessage(rawDays.(string), map[string]interface{}{
		"attribute": "روز",
		"days": []string{
			"شنبه", "یکشنبه", "دوشنبه",
		},
	}, "، ")

	fmt.Println(msgDays)
}
