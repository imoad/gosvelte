package gosvelte

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"go.uber.org/fx"
	"gopkg.in/yaml.v3"
)

type Data map[string]interface{}

type ServeReq struct {
	SvelteFileName          string
	SvelteOptionsTag        string
	HtmlGlobalLangAttribute string
	HtmlGlobalDirAttribute  string
	Title                   string
	Data                    Data
	HeadElements            template.HTML
	BodyTopElements         template.HTML
	BodyBottomElements      template.HTML
}

type GoSvelte struct {
	InternPublicPath    string `yaml:"internpublicpath"`
	ForwardPublicPath   string `yaml:"forwardpublicpath"`
	SvelteWorkspacePath string `yaml:"svelteworkspacepath"`
	SvelteOutputPath    string `yaml:"svelteoutputpath"`
	SvelteExtension     string `yaml:"svelteextension"`
}

type config struct {
	GoSvelte	`yaml:"gosvelte"`
}

func (goSvelteCfg *GoSvelte) Serve(w io.Writer, serveReq ServeReq) error {

	svelteFilePath := filepath.Join(goSvelteCfg.SvelteWorkspacePath, fmt.Sprintf(`%s%s`, serveReq.SvelteFileName, goSvelteCfg.SvelteExtension))
	if err := isFileExist(svelteFilePath); os.IsNotExist(err) {
		return err
	}

	if (len(serveReq.SvelteOptionsTag) == 0) {
	svelteOptionsTag, err := getSvelteOptionsTag(svelteFilePath)
	if err != nil {
		return err
	}
	serveReq.SvelteOptionsTag = svelteOptionsTag
	}

	goSvelteRes, err := setServeRes(goSvelteCfg, serveReq)
	if err != nil {
		return err
	}

	tmplPath := filepath.Join("system", "gosvelte", "src", "base.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return err
	}

	tmpl.ExecuteTemplate(w, "gosvlete", goSvelteRes)
	return nil
}

func provideGoSvelte() *GoSvelte {
	config := config{}
	data, err := ioutil.ReadFile("system/gosvelte/src/gosvelte.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		panic(err)
	}

	return &config.GoSvelte
	//return new("www/public", "/www/public", "resources/views", "/gosvelte", ".svelte")
}

var Module = fx.Options(
	fx.Provide(provideGoSvelte),
)
