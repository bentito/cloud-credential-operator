package gcp

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"google.golang.org/api/iam/v1"

	"github.com/openshift/cloud-credential-operator/pkg/cmd/provisioning"
	"github.com/openshift/cloud-credential-operator/pkg/gcp"
)

var (
	// CreateWorkloadIdentityPoolOpts captures the options that affect creation of the workload identity pool
	CreateWorkloadIdentityPoolOpts = options{
		Name:      "",
		Project:   "",
		TargetDir: "",
	}
	// createdByCcoctl is a standard description for the Google cloud resources created by ccoctl
	createdByCcoctl = "Created By OpenShift ccoctl"
	// createIdentityPoolScriptName is the name of the script used to create workload identity pool
	createIdentityPoolScriptName = "01-create-workload-identity-pool.sh"
	// createIdentityPoolCmd is a gcloud cli command to create workload identity pool
	createIdentityPoolCmd = "gcloud iam workload-identity-pools create %s --location=global --description=\"Created by OpenShift ccoctl\" --display-name=%s"
)

const (
	// fileModeCcoctlDryRun represents a mode and permission bits of the files created by ccoctl in dry run
	fileModeCcoctlDryRun = 0744
)

func createWorkloadIdentityPoolCmd(cmd *cobra.Command, args []string) {
	ctx := context.Background()

	gcpClient, err := gcp.NewClient(CreateWorkloadIdentityPoolOpts.Project, nil)
	if err != nil {
		log.Fatalf("Failed to setup GCP client: %s", err)
	}

	err = createWorkloadIdentityPool(ctx, gcpClient, CreateWorkloadIdentityPoolOpts.Name, CreateWorkloadIdentityPoolOpts.Project, CreateWorkloadIdentityPoolOpts.TargetDir, CreateWorkloadIdentityPoolOpts.DryRun)
	if err != nil {
		log.Fatal(err)
	}
}

// validationForCreateWorkloadIdentityPoolCmd will validate the arguments to the command, ensure the destination directory
// is ready to receive the generated files, and will create the directory if necessary.
func validationForCreateWorkloadIdentityPoolCmd(cmd *cobra.Command, args []string) {
	if len(CreateWorkloadIdentityPoolOpts.Name) > 32 {
		log.Fatalf("Name can be at most 32 characters long")
	}

	if CreateWorkloadIdentityPoolOpts.TargetDir == "" {
		pwd, err := os.Getwd()
		if err != nil {
			log.Fatalf("Failed to get current directory: %s", err)
		}

		CreateWorkloadIdentityPoolOpts.TargetDir = pwd
	}

	fPath, err := filepath.Abs(CreateWorkloadIdentityPoolOpts.TargetDir)
	if err != nil {
		log.Fatalf("Failed to resolve full path: %s", err)
	}

	// create target dir if necessary
	err = provisioning.EnsureDir(fPath)
	if err != nil {
		log.Fatalf("failed to create target directory at %s", fPath)
	}
}

func createWorkloadIdentityPool(ctx context.Context, client gcp.Client, name, project, targetDir string, generateOnly bool) error {
	if generateOnly {
		createWorkloadIdentityPoolScript := provisioning.CreateShellScript([]string{
			fmt.Sprintf(createIdentityPoolCmd, name, name),
		})
		createIdentityPoolScriptFullPath := filepath.Join(targetDir, createIdentityPoolScriptName)
		log.Printf("Saving script to create workload identity pool %s locally at %s", name, createIdentityPoolScriptFullPath)
		if err := ioutil.WriteFile(createIdentityPoolScriptFullPath, []byte(createWorkloadIdentityPoolScript), fileModeCcoctlDryRun); err != nil {
			return errors.Wrap(err, fmt.Sprintf("Failed to save script to create workload identity pool %s locally at %s", name, createIdentityPoolScriptFullPath))
		}
	} else {
		pool := &iam.WorkloadIdentityPool{
			Name:        name,
			DisplayName: name,
			Description: createdByCcoctl,
			State:       "ACTIVE",
			Disabled:    false,
		}

		parentResourceForPool := fmt.Sprintf("projects/%s/locations/global", project)
		_, err := client.CreateWorkloadIdentityPool(ctx, parentResourceForPool, name, pool)
		if err != nil {
			return errors.Wrapf(err, "failed to create workload identity pool %s", name)
		}
		log.Printf("workload identity Pool created with name %s", name)
	}

	return nil
}

// NewCreateWorkloadIdentityPool provides the "create-workload-identity-pool" subcommand
func NewCreateWorkloadIdentityPool() *cobra.Command {
	createWorkloadIdentityPoolCmd := &cobra.Command{
		Use:              "create-workload-identity-pool",
		Short:            "Create workload identity pool",
		Run:              createWorkloadIdentityPoolCmd,
		PersistentPreRun: validationForCreateWorkloadIdentityPoolCmd,
	}

	createWorkloadIdentityPoolCmd.PersistentFlags().StringVar(&CreateWorkloadIdentityPoolOpts.Name, "name", "", "User-defined name for all created Google cloud resources (can be separate from the cluster's infra-id)")
	createWorkloadIdentityPoolCmd.MarkPersistentFlagRequired("name")
	createWorkloadIdentityPoolCmd.PersistentFlags().StringVar(&CreateWorkloadIdentityPoolOpts.Project, "project", "", "ID of the Google cloud project")
	createWorkloadIdentityPoolCmd.MarkPersistentFlagRequired("project")
	createWorkloadIdentityPoolCmd.PersistentFlags().BoolVar(&CreateWorkloadIdentityPoolOpts.DryRun, "dry-run", false, "Skip creating objects, and just save what would have been created into files")
	createWorkloadIdentityPoolCmd.PersistentFlags().StringVar(&CreateWorkloadIdentityPoolOpts.TargetDir, "output-dir", "", "Directory to place generated files (defaults to current directory)")

	return createWorkloadIdentityPoolCmd
}