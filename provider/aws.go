package provider

import (
	"encoding/json"
	"io/ioutil"
	"path"
	"runtime"
)

type EC2 struct {
	RegionList       []map[string]string `json:"region"`
	ZoneList         []map[string]string `json:"zone"`
	InstanceTypeList []map[string]string `json:"instance-type"`
}

type AWS struct {
	*Provider
	EC2 *EC2 `json:"ec2"`
}

func NewAWS(pvd *Provider) *AWS {
	_, filename, _, _ := runtime.Caller(1)
	filePath := path.Join(filename, sourceDir, "aws.json")
	f, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	aws := &AWS{pvd, nil}
	json.Unmarshal(f, &aws)
	return aws
}

func (aws *AWS) RenderEC2Region() {
	aws.render(aws.EC2.RegionList)
}

func (aws *AWS) RenderEC2Zone() {
	aws.render(aws.EC2.ZoneList)
}

func (aws *AWS) RenderEC2InstanceType() {
	aws.render(aws.EC2.InstanceTypeList)
}
