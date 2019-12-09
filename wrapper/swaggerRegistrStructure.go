package wrapper

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

type s struct {}

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

var (
	jsonStr = ""
	sInfo = swaggerInfo{}
)

// SwaggerInfo holds exported Swagger Info so clients can modify it
func setSwaggerInfo(version, host, basePath, titile, description string, schema []string) {
	sInfo =  swaggerInfo{
		Version:     version,
		Host:        host,
		BasePath:    basePath,
		Schemes:     schema,
		Title:       titile,
		Description: description,
	}
}

func (s *s) ReadDoc() string {
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(jsonStr)
	if err != nil {
		return jsonStr
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return jsonStr
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}