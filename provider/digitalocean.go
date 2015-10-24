package provider

import (
	"encoding/json"
	"io/ioutil"
	"path"
	"runtime"
)

type DigitalOceanImage struct {
	DistList []map[string]string `json:"dist"`
	AppList  []map[string]string `json:"app"`
}

type DigitalOcean struct {
	*Provider
	RegionList []map[string]string `json:"region"`
	SizeList   []map[string]string `json:"size"`
	Image      *DigitalOceanImage  `json:"image"`
}

func NewDigitalOcean(pvd *Provider) *DigitalOcean {
	_, filename, _, _ := runtime.Caller(1)
	filePath := path.Join(filename, sourceDir, "digitalocean.json")
	f, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	do := &DigitalOcean{pvd, nil, nil, nil}
	json.Unmarshal(f, &do)
	return do
}

func (do *DigitalOcean) RenderRegion() {
	do.render(do.RegionList)
}

func (do *DigitalOcean) RenderSize() {
	do.render(do.SizeList)
}

func (do *DigitalOcean) RenderDistImage() {
	do.render(do.Image.DistList)
}

func (do *DigitalOcean) RenderAppImage() {
	do.render(do.Image.AppList)
}
