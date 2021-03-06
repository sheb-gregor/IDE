// File generated by Gopher Sauce
// DO NOT EDIT!!
package templates

import (
	"github.com/thestrukture/IDE/types"
	"log"
)

// Template path
var templateIDxButton = "tmpl/ui/sbutton.tmpl"

//
// Renders HTML of template
// xButton with struct types.SButton
func XButton(d types.SButton) string {
	return netbxButton(d)
}

// Render template with JSON string as
// data.
func netxButton(args ...interface{}) string {

	// Get data from JSON
	var d = netcxButton(args...)
	return netbxButton(d)

}

// template render function
func netbxButton(d types.SButton) string {
	localid := templateIDxButton
	name := "xButton"
	defer templateRecovery(name, localid, &d)

	// render and return template result
	return executeTemplate(name, localid, &d)
}

// Unmarshal a json string to the template's struct
// type
func netcxButton(args ...interface{}) (d types.SButton) {

	if len(args) > 0 {
		jsonData := args[0].(string)
		err := parseJSON(jsonData, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	}

	return
}
