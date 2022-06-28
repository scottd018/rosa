/*
Copyright (c) 2020 Red Hat, Inc.

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

package addon

import (
	"fmt"
	"os"
	"text/tabwriter"

	cmv1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
	"github.com/spf13/cobra"

	"github.com/openshift/rosa/pkg/ocm"
	"github.com/openshift/rosa/pkg/rosa"
)

var args struct {
	clusterKey string
}

var Cmd = &cobra.Command{
	Use:     "addons",
	Aliases: []string{"addon", "add-ons", "add-on"},
	Short:   "List add-on installations",
	Long:    "List add-ons installed on a cluster.",
	Example: `  # List all add-on installations on a cluster named "mycluster"
  rosa list addons --cluster=mycluster`,
	Run: run,
}

func init() {
	flags := Cmd.Flags()

	flags.StringVarP(
		&args.clusterKey,
		"cluster",
		"c",
		"",
		"Name or ID of the cluster to list the add-ons of (required).",
	)
}

func run(_ *cobra.Command, _ []string) {
	r := rosa.NewRuntime().WithAWS().WithOCM()
	defer r.Cleanup()

	// Check that the cluster key (name, identifier or external identifier) given by the user
	// is reasonably safe so that there is no risk of SQL injection:
	clusterKey := args.clusterKey
	if clusterKey != "" && !ocm.IsValidClusterKey(clusterKey) {
		r.Reporter.Errorf(
			"Cluster name, identifier or external identifier '%s' isn't valid: it "+
				"must contain only letters, digits, dashes and underscores",
			clusterKey,
		)
		os.Exit(1)
	}

	if clusterKey == "" {
		r.Reporter.Debugf("Fetching all available add-ons")
		addOnResources, err := r.OCMClient.GetAvailableAddOns()
		if err != nil {
			r.Reporter.Errorf("Failed to fetch add-ons: %v", err)
			os.Exit(1)
		}
		if len(addOnResources) == 0 {
			r.Reporter.Infof("There are no add-ons available")
			os.Exit(0)
		}

		// Create the writer that will be used to print the tabulated results:
		writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintf(writer, "ID\t\tNAME\t\tAVAILABILITY\n")
		for _, addOnResource := range addOnResources {
			availability := "unavailable"
			if addOnResource.Available {
				availability = "available"
			}
			fmt.Fprintf(writer, "%s\t\t%s\t\t%s\n", addOnResource.AddOn.ID(), addOnResource.AddOn.Name(), availability)
		}
		writer.Flush()

		os.Exit(0)
	}

	// Try to find the cluster:
	r.Reporter.Debugf("Loading cluster '%s'", clusterKey)
	cluster, err := r.OCMClient.GetCluster(clusterKey, r.Creator)
	if err != nil {
		r.Reporter.Errorf("Failed to get cluster '%s': %v", clusterKey, err)
		os.Exit(1)
	}

	if cluster.State() != cmv1.ClusterStateReady {
		r.Reporter.Errorf("Cluster '%s' is not yet ready", clusterKey)
		os.Exit(1)
	}

	// Load any existing Add-Ons for this cluster
	r.Reporter.Debugf("Loading add-ons installations for cluster '%s'", clusterKey)
	clusterAddOns, err := r.OCMClient.GetClusterAddOns(cluster)
	if err != nil {
		r.Reporter.Errorf("Failed to get add-ons for cluster '%s': %v", clusterKey, err)
		os.Exit(1)
	}

	if len(clusterAddOns) == 0 {
		r.Reporter.Infof("There are no add-ons installed on cluster '%s'", clusterKey)
		os.Exit(0)
	}

	// Create the writer that will be used to print the tabulated results:
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintf(writer, "ID\t\tNAME\t\tSTATE\n")
	for _, clusterAddOn := range clusterAddOns {
		fmt.Fprintf(writer, "%s\t\t%s\t\t%s\n", clusterAddOn.ID, clusterAddOn.Name, clusterAddOn.State)
	}
	writer.Flush()
}
