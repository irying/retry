package utils

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func readTemplateFile(filename string) (string, error) {
	_, err := os.Stat(filename)
	if err != nil {
		return "", fmt.Errorf("State%s error: %s", filename, err.Error())
	}

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("ReadFile %s error: %s", filename, err.Error())
	}

	return string(contents), nil
}

func evaluateTemplate(temp *template.Template, fillData interface{}) (string, error) {
	buffer := new(bytes.Buffer)
	err := temp.Execute(buffer, fillData)
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}

func evaluateTemplateFile(name, path string, fillData interface{}) (string, error) {
	yamlTemplateStr, err := readTemplateFile(path)
	if err != nil {
		return "", err
	}

	daemonSetTemplate, err := template.New(name).Funcs(templateFuncMaps).Parse(yamlTemplateStr)
	if err != nil {
		return "", err
	}
	return evaluateTemplate(daemonSetTemplate, fillData)

}

// EvaluateHealthCheckerTemplate generate health-checker daemonset k8s spec file
// from template file
func EvaluateHealthCheckerTemplate(tempDir string, data interface{}) (string, error) {
	name := "healthchecker.template.yaml"
	path := filepath.Join(tempDir, name)
	return evaluateTemplateFile(name, path, data)
}

// EvaluateDeploymentTemplate generate deployment resource k8s spec file from
// template file
func EvaluateDeploymentTemplate(tempDir string, data interface{}) (string, error) {
	name := "deployment.template.yaml"
	path := filepath.Join(tempDir, name)
	return evaluateTemplateFile(name, path, data)
}

// EvaluateServiceTemplate generate service resource k8s spec file from template
// file
func EvaluateServiceTemplate(tempDir string, data interface{}) (string, error) {
	name := "service.template.yaml"
	path := filepath.Join(tempDir, name)
	return evaluateTemplateFile(name, path, data)

}

// EvaluateStatefulSetTemplate generate statefulset resource k8s spec file from
// template file
func EvaluateStatefulSetTemplate(tempDir string, data interface{}) (string, error) {
	name := "statefulset.template.yaml"
	path := filepath.Join(tempDir, name)
	return evaluateTemplateFile(name, path, data)
}

// EvaluateConfigMapTemplate generate configmap resource k8s spec file from
// template file
func EvaluateConfigMapTemplate(tempDir string, data interface{}) (string, error) {
	name := "configmap.template.yaml"
	path := filepath.Join(tempDir, name)
	return evaluateTemplateFile(name, path, data)
}

// EvaluateIngressControllerTemplate generate ingressController resource k8s
// spec file from template file
func EvaluateIngressControllerTemplate(tempDir string, data interface{}) (string, error) {
	name := "ingress-controller.template.yaml"
	path := filepath.Join(tempDir, name)
	return evaluateTemplateFile(name, path, data)
}

// EvaluateIngressTemplate generate ingress resource k8s spec file from
// template file
func EvaluateIngressTemplate(tempDir string, data interface{}) (string, error) {
	name := "ingress.template.yaml"
	path := filepath.Join(tempDir, name)
	return evaluateTemplateFile(name, path, data)
}

func substring(begin int, str string) string {
	return str[begin:]
}

func length(s []string) int {
	return len(s)
}

func base64Encoding(str string) (result string) {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

var (
	templateFuncMaps = map[string]interface{}{
		"toupper":        strings.ToUpper,
		"tolower":        strings.ToLower,
		"replace":        strings.Replace,
		"substr":         substring,
		"joinstring":     strings.Join,
		"length":         length,
		"base64Encoding": base64Encoding,
	}
)
