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

func ClusterStop(clustername string) {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")
	if subscriptionID == "" {
		log.Fatal("AZURE_SUBSCRIPTION_ID environment variable is not set")
	}
	client, err := armcontainerservice.NewManagedClustersClient(subscriptionID, cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	resourceGroupName := os.Getenv("AZURE_RESOURCE_GROUP_NAME")
	if resourceGroupName == "" {
		log.Fatal("AZURE_RESOURCE_GROUP_NAME environment variable is not set")
	}
	clusterName := clustername
	if clusterName == "" {
		fmt.Println("Cluster name argument is required")
	}
	poller, err := client.BeginStop(ctx, resourceGroupName, clusterName, nil)
	if err != nil {
		callbacks.PrintErrorMessage(err)
		log.Fatalf("failed to pull the result: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

func ClusterStart(clustername string) {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")
	if subscriptionID == "" {
		log.Fatal("AZURE_SUBSCRIPTION_ID environment variable is not set")
	}
	client, err := armcontainerservice.NewManagedClustersClient(subscriptionID, cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	resourceGroupName := os.Getenv("AZURE_RESOURCE_GROUP_NAME")
	if resourceGroupName == "" {
		log.Fatal("AZURE_RESOURCE_GROUP_NAME environment variable is not set")
	}
	clusterName := clustername
	if clusterName == "" {
		fmt.Println("Cluster name argument is required")
	}
	poller, err := client.BeginStart(ctx, resourceGroupName, clusterName, nil)
	if err != nil {
		callbacks.PrintErrorMessage(err)
		log.Fatalf("failed to pull the result: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
