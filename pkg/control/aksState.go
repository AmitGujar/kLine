package control

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/AmitGujar/kLine/pkg/callbacks"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice"
)

func GetClusterState(clustername string) {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	subscription := os.Getenv("AZURE_SUBSCRIPTION_ID")
	client, err := armcontainerservice.NewManagedClustersClient(subscription, cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	resourceGroupName := os.Getenv("AZURE_RESOURCE_GROUP_NAME")
	res, err := client.Get(ctx,
		resourceGroupName,
		clustername,
		nil)
	if err != nil {
		callbacks.PrintErrorMessage(err)
		log.Fatalf("failed to pull the result: %v", err)
	}

	if res.ManagedCluster.Properties.PowerState != nil {
		fmt.Printf("AKS cluster '%v' is %s\n", clustername, *res.ManagedCluster.Properties.PowerState.Code)
	} else {
		fmt.Println("Cluster power state is not available")
	}
}
