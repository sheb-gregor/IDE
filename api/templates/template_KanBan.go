// File generated by Gopher Sauce
// DO NOT EDIT!!
package templates

import (
	"github.com/thestrukture/IDE/types"
	"log"
)

// Template path
var templateIDKanBan = "tmpl/ui/kanban.tmpl"

//
// Renders HTML of template
// KanBan with struct types.Dex
func KanBan(d types.Dex) string {
	return netbKanBan(d)
}

// Render template with JSON string as
// data.
func netKanBan(args ...interface{}) string {

	// Get data from JSON
	var d = netcKanBan(args...)
	return netbKanBan(d)

}

// template render function
func netbKanBan(d types.Dex) string {
	localid := templateIDKanBan
	name := "KanBan"
	defer templateRecovery(name, localid, &d)

	// render and return template result
	return executeTemplate(name, localid, &d)
}

// Unmarshal a json string to the template's struct
// type
func netcKanBan(args ...interface{}) (d types.Dex) {

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
