// File generated by Gopher Sauce
// DO NOT EDIT!!
package templates

import (
	"github.com/thestrukture/IDE/types"
	"log"
)

// Template path
var templateIDDocker = "tmpl/ui/docker.tmpl"

//
// Renders HTML of template
// Docker with struct types.Dex
func Docker(d types.Dex) string {
	return netbDocker(d)
}

// Render template with JSON string as
// data.
func netDocker(args ...interface{}) string {

	// Get data from JSON
	var d = netcDocker(args...)
	return netbDocker(d)

}

// template render function
func netbDocker(d types.Dex) string {
	localid := templateIDDocker
	name := "Docker"
	defer templateRecovery(name, localid, &d)

	// render and return template result
	return executeTemplate(name, localid, &d)
}

// Unmarshal a json string to the template's struct
// type
func netcDocker(args ...interface{}) (d types.Dex) {

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
