package main

import (
	//"bytes"
	"fmt"
	"os"

	//"encoding/json"
	"github.com/codegangsta/cli"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/fields"
	//"k8s.io/kubernetes/pkg/labels"

	"k8s.io/kubernetes/pkg/client/restclient"
	"k8s.io/kubernetes/pkg/client/unversioned"
	// "k8s.io/kubernetes/pkg/runtime"
	// "k8s.io/kubernetes/pkg/watch/json"
	// "k8s.io/kubernetes/pkg/field	s"
)

func checkKubePods(c *cli.Context) {

	clientConfig := restclient.Config{}
	clientConfig.Host = "127.0.0.1:8080"
	clientU, err := unversioned.New(&clientConfig)
	if err != nil {
		fmt.Println("new unverioned err!")
	}
	pods, err := clientU.Pods("").List(api.ListOptions{})
	if err != nil {
		fmt.Printf("list pods err! err := %v\n", err)
	}
	for _, pod := range pods.Items {
		for _, cond := range pod.Status.Conditions {
			fmt.Printf("pod.Name := %v pod.Type := %v, pod.Status := %v ,pod.Status.Phase := %v\n", pod.Name, cond.Type, cond.Status, pod.Status.Phase)
		}
	}
}

func watchPod(c *cli.Context) {

	clientConfig := restclient.Config{}
	clientConfig.Host = "127.0.0.1:8080"
	clientU, err := unversioned.New(&clientConfig)
	if err != nil {
		fmt.Println("new unverioned err!")
	}

	/*
		opts := api.ListOptions{
			LabelSelector: "",
			FieldSelector: "name=nginx",
		}
	*/

	/*	selector := fields.Set{
		"kind": "pod",
	}.AsSelector()*/

	opts := api.ListOptions{FieldSelector: fields.Set{"metadata.name": "nginx"}.AsSelector()}

	w, err := clientU.Pods("").Watch(opts)
	if err != nil {
		fmt.Println("get watch interface err")
	} else {
		for {
			event, ok := <-w.ResultChan()
			if !ok {
				fmt.Errorf("closed early\n")
				break
			} else {
				fmt.Println("11111")
				fmt.Println(event.Type)
				//pp := api.Pod(event.Object)
				/*a, _ := json.Marshal(event.Object)
				var ap api.Pod
				json.Unmarshal(a, &ap)
				fmt.Println(ap)
				fmt.Println(ap.GetObjectMeta().GetName())
				for a, _ := range ap.Spec.Containers {
					fmt.Println(api.Container(a).Name)
				}*/
				a := event.Object.(*api.Pod)
				fmt.Println(a)
				fmt.Println(a.GetObjectMeta().GetName())
				// fmt.Println(a.Spec.Containers.(api.Container))
				for _, a := range a.Spec.Containers {
					fmt.Println(a.Name)
				}
				// obj.(*api.Pod).Spec)
				//b, err := runtime.Encode(event, event.Object)
				//fmt.Println(b)
				//api.Pod(event.Object)
				//fmt.Println(event.Object)
				//bytes.NewBuffer(event.Object)
				//a := fmt.Sprintf(string(event.Object))
				//fmt.Println(a)
				//pp := new(api.Pod)
				//json.Unmarshal([]byte(a), pp)
				//fmt.Println(pp.GetName())
				//fmt.Println(pp.GetUID())

			}
		}
	}
	/*
		pods, err := clientU.Pods("").List(api.ListOptions{})
		if err != nil {
			fmt.Printf("list pods err! err := %v\n", err)
		}
		for _, pod := range pods.Items {
			for _, cond := range pod.Status.Conditions {
				fmt.Printf("pod.Name := %v pod.Type := %v, pod.Status := %v ,pod.Status.Phase := %v\n", pod.Name, cond.Type, cond.Status, pod.Status.Phase)
			}
		}
	*/
}

func main() {
	app := cli.NewApp()
	app.Name = "check_kube_nodes"
	app.HelpName = app.Name
	app.Usage = "Nagios check to verify Kubernetes resources status"
	app.Version = "1.0"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "api-endpoint",
			Value: "",
			Usage: "Kubernetes API Endpoint",
		},
		cli.StringFlag{
			Name:  "username",
			Value: "",
			Usage: "Kubernetes API Username",
		},
		cli.StringFlag{
			Name:  "password",
			Value: "",
			Usage: "Kubernetes API Password",
		},
		cli.BoolFlag{
			Name:  "skip-tls-verify",
			Usage: "Skip TLS certificate verification",
		},
	}

	app.Commands = []cli.Command{
		cli.Command{
			Name:    "pod",
			Aliases: []string{"p"},
			Usage:   "check pod status",
			Action: func(c *cli.Context) {
				checkKubePods(c)
			},
		},
		cli.Command{
			Name:    "watch",
			Aliases: []string{"w"},
			Usage:   "watch the pod",
			Action: func(c *cli.Context) {
				watchPod(c)
			},
		},
	}

	app.Run(os.Args)
}
