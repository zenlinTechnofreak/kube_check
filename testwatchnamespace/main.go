package main

import (
	//"bytes"
	"fmt"
	// "github.com/containerops/vessel/models"
	client "github.com/containerops/vessel/module/kubernetes"
)

/*type PipelineVersion struct {
	Id            int64    `json:"id,omitempty"`
	WorkspaceId   int64    `json:"workspaceId, omitempty"`
	ProjectId     int64    `json:"projectId, omitempty"`
	PipelineId    int64    `json:"pipelineId, omitempty"`
	Namespace     string   `json:"namespace, omitempty"`
	SelfLink      string   `json:"selfLink, omitempty"`
	Created       int64    `json:"created, omitempty"`
	Updated       int64    `json:"updated, omitempty"`
	Labels        string   `json:"labels, omitempty"`
	Annotations   string   `json:"annotations, omitempty"`
	Detail        string   `json:"detail, omitempty"`
	StageVersions []string `json:"stageVersions, omitempty"`
	Log           string   `json:"log, omitempty"`
	Status        int64    `json:"state, omitempty"` // 0 not start    1 working    2 success     3 failed
	MetaData      string   `json:"metadata, omitempty"`
	Spec          string   `json:"spec, omitempty"`
}

func (p *PipelineVersion) GetMetadata() string {
	return p.MetaData
}

func (p *PipelineVersion) GetSpec() string {
	return p.Spec
}
*/
/*
type piplineSpec struct {
	name                string `json:"name, omitempty"`
	replicas            int    `json:"replicas, omitempty"`
	dependencies        string `json:"dependencies, omitempty"`
	kind                string `json:"kind, omitempty"`
	statusCheckLink     string `json:"statusCheckLink, omitempty"`
	statusCheckInterval int64  `json:"statusCheckInterval, omitempty"`
	statusCheckCount    int64  `json:"statusCheckCount, omitempty"`
	imageName           string `json:"imagename, omitempty"`
	port                int    `json:"port, omitempty"`
}
*/

/*
type piplineMetadata struct {
	name            string            `json:"name, omitempty"`
	namespace       string            `json:"namespace, omitempty"`
	selfLink        string            `json:"selflink, omitempty"`
	uid             types.UID         `json:"uid, omitempty"`
	createTimestamp unversioned.Time  `json:"createTimestamp, omitempty"`
	deleteTimestamp unversioned.Time  `json:"deleteTimestamp, omitempty"`
	timeoutDuration int64             `json:"timeoutDuration, omitempty"`
	labels          map[string]string `json:"labels, omitempty"`
	annotations     map[string]string `json:"annotations, omitempty"`
}
*/

func main() {

	client.New("127.0.0.1:8080")
	res, err := client.WatchNamespaceStatus("app", "zenlin", 40, client.Added)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)

	/*// pv := new(models.PipelineVersion)
	// pv.Detail = ""

	// pv.Detail = ""
	pv.MetaData.Name = "zenlin-test"
	pv.MetaData.Namespace = "zenlin"
	// pv.Labels[""]

	pv.StageSpecs = make([]models.StageSpec, 1)
	pv.StageSpecs[0].Name = "zenlintestxx"
	pv.StageSpecs[0].Image = "kubernetes/redis-slave:v2"
	pv.StageSpecs[0].Replicas = 2
	pv.StageSpecs[0].Port = 6379

	err := client.StartK8SResource(pv)
	if err != nil {
		fmt.Printf("star err %v\n", err)
	} else {
		fmt.Println("successful")
	}

	// client.NewClient("127.0.0.1:8080")

	// x := <-c
	// fmt.Println(c)*/
}
