package command

import "time"

// Custom fake was written for this under customv2fakes

// UI is the interface to STDOUT
type UI interface {
	DisplayBoolPrompt(prompt string, defaultResponse bool) (bool, error)
	DisplayError(err error)
	DisplayHeader(text string)
	DisplayNewline()
	DisplayOK()
	DisplayPair(attribute string, formattedString string, keys ...map[string]interface{})
	DisplayTable(prefix string, table [][]string, padding int)
	DisplayText(template string, data ...map[string]interface{})
	DisplayTextWithFlavor(text string, keys ...map[string]interface{})
	DisplayWarning(formattedString string, keys ...map[string]interface{})
	DisplayWarnings(warnings []string)
	TranslateText(template string, data ...map[string]interface{}) string
	UserFriendlyDate(input time.Time) string
}
