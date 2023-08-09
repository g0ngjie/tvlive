package main

import (
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type BaseTarget struct {
	Name string `yaml:"name"`
	Logo string `yaml:"logo"`
	Url  string `yaml:"url"`
}

type Config struct {
	Type  string       `yaml:"type"`
	Group string       `yaml:"group"`
	List  []BaseTarget `yarml:"list"`
}

const YamlPath = "config.yml"

func loadYaml() []Config {
	content, _ := os.ReadFile(YamlPath)
	config := &[]Config{}
	yaml.Unmarshal(content, &config)

	return *config
}

const Programme = "#EXTM3U x-tvg-url=\"https://live.g0ngjie.com/e.xml\"\n"
const WriteM3UPath = "./index.m3u"

func pkgContent(cfg []Config) (content string) {
	content = Programme
	for _, target := range cfg {
		// #EXTINF:-1 tvg-id="CCTV1" tvg-name="CCTV1" tvg-logo="https://live.g0ngjie.com/tv/CCTV1.png" group-title="央视",CCTV-1
		// https://cntv.sbs/live?auth=2307280808&id=cctv1
		for _, current := range target.List {
			var extinf strings.Builder
			extinf.WriteString("#EXTINF:-1 tvg-id=\"")
			extinf.WriteString(current.Name)
			extinf.WriteString("\" tvg-name=\"")
			extinf.WriteString(current.Name)
			extinf.WriteString("\" tvg-logo=\"https://live.g0ngjie.com/")
			extinf.WriteString(current.Logo)
			extinf.WriteString("\" group-title=\"")
			extinf.WriteString(target.Group)
			extinf.WriteString("\",")
			// group
			extinf.WriteString(current.Name)
			extinf.WriteString("\n")
			// url
			extinf.WriteString(current.Url)
			extinf.WriteString("\n")
			content += extinf.String()
		}
	}
	return
}

func writeM3UFile(list string) {
	os.WriteFile(WriteM3UPath, []byte(list), 0777)
}

func main() {
	cfgs := loadYaml()
	content := pkgContent(cfgs)
	writeM3UFile(content)
}
