package internal

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/stoewer/go-strcase"
)

const TEMPLATES = "templates"

type Data struct {
	Key        string
	Desciption string
}

type Template struct {
	Template   *template.Template
	Desciption string
	Arguments  []Data
}

type Templates map[string]Template

type File struct {
	Template string
	Path     string
}

func nilCase(f func(string) string) func(any) string {
	return func(value any) string {
		return f(fmt.Sprintf("%v", value))
	}
}

func LoadTemplates(dir string, writeFile func(tmpl string, path string)) (Templates, error) {
	templates := Templates{}
	var current Template

	var funcMap = template.FuncMap{
		"pascal":     nilCase(strcase.UpperCamelCase),
		"camel":      nilCase(strcase.LowerCamelCase),
		"snake":      nilCase(strcase.SnakeCase),
		"kebab":      nilCase(strcase.KebabCase),
		"trim":       strings.Trim,
		"trimSuffix": strings.TrimSuffix,
		"trimPrefix": strings.TrimPrefix,
		"split":      strings.Split,
		"basename":   path.Base,
		"dirname":    path.Dir,
		"pathJoin":   path.Join,
		"file": func(template string, path string) string {
			writeFile(template, path)
			return ""
		},
		"argument": func(name, desciption string) string {
			return ""
		},
		"description": func(desciption string) string {
			return ""
		},
	}

	var initMap = template.FuncMap{
		"pascal":     nilCase(strcase.UpperCamelCase),
		"camel":      nilCase(strcase.LowerCamelCase),
		"snake":      nilCase(strcase.SnakeCase),
		"kebab":      nilCase(strcase.KebabCase),
		"trim":       strings.Trim,
		"trimSuffix": strings.TrimSuffix,
		"trimPrefix": strings.TrimPrefix,
		"split":      strings.Split,
		"basename":   path.Base,
		"dirname":    path.Dir,
		"pathJoin":   path.Join,
		"file": func(template string, path string) string {
			// we don't want to mess with files during init.
			return ""
		},
		"argument": func(name, desciption string) string {
			current.Arguments = append(current.Arguments, Data{
				Key:        name,
				Desciption: desciption,
			})
			return ""
		},
		"description": func(desciption string) string {
			current.Desciption = desciption
			return ""
		},
	}

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		base := filepath.Base(path)
		ext := filepath.Ext(path)
		folder := strings.TrimSuffix(path, ext)

		if ext == ".tmpl" {
			tmpl, err := template.New(base).Funcs(initMap).ParseFiles(path)
			if err != nil {
				return err
			}

			keys := strings.Split(folder, string(os.PathSeparator))
			key := strings.Join(keys[1:], ".")

			current = Template{
				Template: tmpl,
			}

			tmpl.Execute(ioutil.Discard, nil)

			tmpl, err = template.New(base).Funcs(funcMap).ParseFiles(path)
			if err != nil {
				return err
			}
			current.Template = tmpl

			templates[key] = current
		}

		return nil
	})
	return templates, err
}
