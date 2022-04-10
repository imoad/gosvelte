package gosvelte

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type serveRes struct {
	HtmlAttrs template.HTMLAttr
	Head      template.HTML
	Body      template.HTML
}

func isFileExist(filePath string) error {
	_, err := os.Stat(filePath)
	return err
}

func getSvelteOptionsTag(svelteFilePath string) (string, error) {
	svelteFileBytes, err := ioutil.ReadFile(svelteFilePath)
	if err != nil {
		return "", err
	}
	r := regexp.MustCompile(`<svelte:options .*tag.*=.*"(.*?)".*/>`)
	svelteFileStr := string(svelteFileBytes)
	svelteOptionsTagRegxRes := r.FindStringSubmatch(svelteFileStr)
	if len(svelteOptionsTagRegxRes) == 0 || len(svelteOptionsTagRegxRes[1]) == 0 {
		return "", fmt.Errorf("no svelte options tag found in file %s", svelteFilePath)
	}
	return svelteOptionsTagRegxRes[1], nil
}

func setServeRes(goSvelteCfg *GoSvelte, serveReq ServeReq) (serveRes, error) {
	htmlAttrsStr := ""
	headStr := ""
	dataStr := ""
	bodyStr := ""

	//htmlAttrsStr
	if len(serveReq.HtmlGlobalLangAttribute) > 0 {
		htmlAttrsStr = fmt.Sprintf(`%s lang='%s'`, htmlAttrsStr, serveReq.HtmlGlobalLangAttribute)
	}
	if len(serveReq.HtmlGlobalDirAttribute) > 0 {
		htmlAttrsStr = fmt.Sprintf(`%s dir='%s'`, htmlAttrsStr, serveReq.HtmlGlobalDirAttribute)
	}

	//headStr
	if len(serveReq.Title) > 0 {
		headStr = fmt.Sprintf(`%s<title>%s</title>`, headStr, serveReq.Title)
	}

	faviconPngFilePath := filepath.Join(goSvelteCfg.InternPublicPath, "/favicon.png")
	if isFileExist(faviconPngFilePath) == nil {
		headStr = fmt.Sprintf(`%s<link rel='icon' type='image/png' href='%s/favicon.png'>`, headStr, goSvelteCfg.ForwardPublicPath)
	}

	globalCssFilePath := filepath.Join(goSvelteCfg.InternPublicPath, "/global.css")
	if isFileExist(globalCssFilePath) == nil {
		headStr = fmt.Sprintf(`%s<link rel='stylesheet' href='%s/global.css'>`, headStr, goSvelteCfg.ForwardPublicPath)
	}

	bundleCssFilePath := filepath.Join(goSvelteCfg.InternPublicPath, goSvelteCfg.SvelteOutputPath, strings.ToLower(serveReq.SvelteFileName), "/bundle.css")
	if isFileExist(bundleCssFilePath) == nil {
		headStr = fmt.Sprintf(`%s<link rel='stylesheet' href='%s%s/%s/bundle.css'>`, headStr, goSvelteCfg.ForwardPublicPath, goSvelteCfg.SvelteOutputPath, strings.ToLower(serveReq.SvelteFileName))
	}

	globalJsFilePath := filepath.Join(goSvelteCfg.InternPublicPath, "/global.js")
	if isFileExist(globalJsFilePath) == nil {
		headStr = fmt.Sprintf(`%s<script defer src='%s/global.js'></script>`, headStr, goSvelteCfg.ForwardPublicPath)
	}

	bundleJsFilePath := filepath.Join(goSvelteCfg.InternPublicPath, goSvelteCfg.SvelteOutputPath, strings.ToLower(serveReq.SvelteFileName), "/bundle.js")
	if err := isFileExist(bundleJsFilePath); os.IsNotExist(err) {
		return serveRes{}, fmt.Errorf("%s not found, did you forget to run {{ yarn install, yarn dev || npm install, npm run dev }} ? ", bundleJsFilePath)
	}
	headStr = fmt.Sprintf(`%s<script defer src='%s%s/%s/bundle.js'></script>`, headStr, goSvelteCfg.ForwardPublicPath, goSvelteCfg.SvelteOutputPath, strings.ToLower(serveReq.SvelteFileName))

	//dataStr
	for k, v := range serveReq.Data {
		dataStr = fmt.Sprintf(`%s %s="%s"`, dataStr, k, v)
	}

	//bodyStr
	bodyStr = fmt.Sprintf(`%s<%s %s></%s>`, bodyStr, serveReq.SvelteOptionsTag, dataStr, serveReq.SvelteOptionsTag)

	goSvelteRes := serveRes{
		HtmlAttrs: template.HTMLAttr(htmlAttrsStr),
		Head:      template.HTML(headStr) + serveReq.HeadElements,
		Body:      serveReq.BodyTopElements + template.HTML(bodyStr) + serveReq.BodyBottomElements,
	}

	return goSvelteRes, nil
}
