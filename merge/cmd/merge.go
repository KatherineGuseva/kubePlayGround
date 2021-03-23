/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	bluemix "github.com/IBM-Cloud/bluemix-go"
	v1 "github.com/IBM-Cloud/bluemix-go/api/container/containerv1"
	"github.com/IBM-Cloud/bluemix-go/session"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path/filepath"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/kubernetes"
	//"k8s.io/client-go/pkg/api/v1"

)
const apikey = ""
func NewRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "merge",
		Short: " A merge command that call the merge on two files cogconfiguration into cogstartup",
		Long: ` A merge command takes two paramters. 

		Example to call the tool is:

		`,
		Run: func(cmd *cobra.Command, args []string) {
			//source, _ := cmd.Flags().GetString("source")
			InitializeLogging("log.txt")

			fmt.Printf("%+v\n", "Getting the secrets from the cluster")
		},
	}
}

func getOwnedResources(cmd *cobra.Command) *corev1.SecretList {
	sess, err := session.New(&bluemix.Config{BluemixAPIKey: apikey})
	if err != nil {
		log.WithError(err).Error("unable to create an IKS session")
		panic(err)
	}

	clusterAPI, err := v1.New(sess)
	if err != nil {
		panic(err)
	}
	th := v1.ClusterTargetHeader{
		Region: "us-south",
	}
	// TODO: would it help to get account, group, org, space ?
	dir, err := ioutil.TempDir("", "tempFolder")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(dir)

	clConf, err := clusterAPI.Clusters().GetClusterConfig("katya-test", dir, false, th)
	fmt.Println(clConf)

	config, err := clientcmd.BuildConfigFromFlags("", clConf)

	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)

	if clientset == nil {
		fmt.Printf("Building config from flags")
	}

	selector, err := cmd.Flags().GetString("selector")
	if err != nil {
		fmt.Printf("Error")
	}
	listSecrets, err := clientset.CoreV1().Secrets("default").List(context.Background(), metav1.ListOptions{LabelSelector: selector})
	if listSecrets == nil {
		fmt.Printf("Error")
	}
	return listSecrets
}


func fileExists(dir string, fileName string) bool {
	_, err := os.Stat(filepath.Join(dir, fileName))
	if os.IsNotExist(err) {
		log.Fatalf("fileName does not exist")
		return false
	}
	return true

}

func InitializeLogging(logFile string) {

	var file, err = os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Could Not Open Log File : " + err.Error())
	}
	log.SetOutput(file)

	log.SetFormatter(&log.JSONFormatter{})
}

func init() {

	cmd := NewRootCmd()
	cmd.Execute()

}
