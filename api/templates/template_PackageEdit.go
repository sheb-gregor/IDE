// File generated by Gopher Sauce
// DO NOT EDIT!!
package templates

import (
	"github.com/thestrukture/IDE/api/assets"
	"github.com/thestrukture/IDE/types"
	"bytes"
	"encoding/json"
	"fmt"
	"html"
	"html/template"
	"log"

	"github.com/fatih/color"
)

//
// Renders HTML of template
// PackageEdit with struct types.SPackageEdit
func PackageEdit(d types.SPackageEdit) string {
	return netbPackageEdit(d)
}

// recovery function used to log a
// panic.
func templateFNPackageEdit(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (ui/user/panel/package) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDPackageEdit = "tmpl/ui/user/panel/package.tmpl"

// Render template with JSON string as
// data.
func netPackageEdit(args ...interface{}) string {

	localid := templateIDPackageEdit
	var d *types.SPackageEdit
	defer templateFNPackageEdit(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &types.SPackageEdit{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := assets.Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("PackageEdit")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", localid))
		DebugTemplatePath(localid, d)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = nil
	output.Reset()
	output = nil
	args = nil
	return outpescaped

}

// alias of template render function.
func bPackageEdit(d types.SPackageEdit) string {
	return netbPackageEdit(d)
}

//

// template render function
func netbPackageEdit(d types.SPackageEdit) string {
	localid := templateIDPackageEdit
	defer templateFNPackageEdit(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := assets.Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("PackageEdit")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		log.Println(erro)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = types.SPackageEdit{}
	output.Reset()
	output = nil
	return outpescaped
}

// Unmarshal a json string to the template's struct
// type
func netcPackageEdit(args ...interface{}) (d types.SPackageEdit) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = types.SPackageEdit{}
	}
	return
}

// Create a struct variable of template.
func cPackageEdit(args ...interface{}) (d types.SPackageEdit) {
	if len(args) > 0 {
		d = netcPackageEdit(args[0])
	} else {
		d = netcPackageEdit()
	}
	return
}