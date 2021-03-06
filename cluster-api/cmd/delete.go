/*
Copyright 2017 The Kubernetes Authors.

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
	"github.com/spf13/cobra"
	"k8s.io/kube-deploy/cluster-api/deploy"
	"log"
)

type DeleteOptions struct {
	Cluster string
	Machine string
}

var do = &DeleteOptions{}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Simple kubernetes cluster manager",
	Long:  `Delete a kubernetes cluster with one command`,
	Run: func(cmd *cobra.Command, args []string) {
		if do.Cluster == "" {
			log.Fatal("Please provide yaml file for cluster definition.")
		}
		if do.Machine == "" {
			log.Fatal("Please provide yaml file for machine definition.")
		}

		if err := RunDelete(do); err != nil {
			log.Fatal(err)
		}
	},
}

func RunDelete(do *DeleteOptions) error {
	cluster, err := parseClusterYaml(do.Cluster)
	if err != nil {
		return err
	}

	if _, err := parseMachinesYaml(do.Machine); err != nil {
		return err
	}
	d := deploy.NewDeployer()

	return d.DeleteCluster(cluster)
}

func init() {
	deleteCmd.Flags().StringVarP(&do.Cluster, "cluster", "c", "", "cluster yaml file")
	deleteCmd.Flags().StringVarP(&do.Machine, "machines", "m", "", "machine yaml file")

	RootCmd.AddCommand(deleteCmd)
}
