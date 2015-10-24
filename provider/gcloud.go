package provider

import (
	"encoding/json"
	"io/ioutil"
	"path"
	"runtime"
)

type GCE struct {
	ZoneList        []map[string]string `json:"zone"`
	MachineTypeList []map[string]string `json:"machine-type"`
	DiskTypeList    []map[string]string `json:"disk-type"`
	ImageList       []map[string]string `json:"image"`
}

type GCloud struct {
	*Provider
	GCE *GCE `json:"gce"`
}

func NewGCloud(pvd *Provider) *GCloud {
	_, filename, _, _ := runtime.Caller(1)
	filePath := path.Join(filename, sourceDir, "gcloud.json")
	f, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	gcloud := &GCloud{pvd, nil}
	json.Unmarshal(f, &gcloud)
	return gcloud
}

func (gcld *GCloud) RenderGCEZone() {
	gcld.render(gcld.GCE.ZoneList)
}

func (gcld *GCloud) RenderGCEMachineType() {
	gcld.render(gcld.GCE.MachineTypeList)
}

func (gcld *GCloud) RenderGCEDiskType() {
	gcld.render(gcld.GCE.DiskTypeList)
}

func (gcld *GCloud) RenderGCEImage() {
	gcld.render(gcld.GCE.ImageList)
}
