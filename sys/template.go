package sys

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"strings"
)

var (
	templates *template.Template
)

// ReturnTemplate 根据名称回传模板
func ReturnTemplate(templateName string) *template.Template {
	templates = analysisTemplateFiles(analysisTemplateDirs("../web/template")...)
	return templates.Lookup(templateName)
}

// analysisTemplateFiles 解析模板
func analysisTemplateFiles(templateFiles ...string) *template.Template {
	templates, err := template.ParseFiles(templateFiles...)
	if err != nil {
		return nil
	}
	return templates
}

// analysisTemplateDirs 解析模板目录获取所有的模板地址
func analysisTemplateDirs(templateDirs ...string) []string {
	var templateFiles []string
	for index0 := range templateDirs {
		files, err0 := ioutil.ReadDir(templateDirs[index0])
		if err0 != nil {
			fmt.Printf("读取模板目录出错！错误信息：%s", err0.Error())
			continue
		}
		for index1 := range files {
			switch files[index1].IsDir() {
			case true:
				templateFiles = append(templateFiles, analysisTemplateDirs(templateDirs[index0]+"/"+files[index1].Name())...)
			case false:
				if strings.Contains(files[index1].Name(), ".scfy") {
					templateFiles = append(templateFiles, templateDirs[index0]+"/"+files[index1].Name())
				}
				continue
			}
		}
	}
	return templateFiles
}
func init() {
	templates = analysisTemplateFiles(analysisTemplateDirs("../web/template")...)
}
