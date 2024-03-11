package control

import (
	"context"
	"log"
	"os"

	"github.com/AmitGujar/kLine/pkg/callbacks"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice"
)

func GetAdminCredentials(clustername string) {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")
	client, err := armcontainerservice.NewManagedClustersClient(subscriptionID, cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	resourceGroupName := os.Getenv("AZURE_RESOURCE_GROUP_NAME")
	res, err := client.ListClusterAdminCredentials(ctx,
		resourceGroupName,
		clustername,
		&armcontainerservice.ManagedClustersClientListClusterAdminCredentialsOptions{ServerFqdn: nil})
	if err != nil {
		callbacks.PrintErrorMessage(err)
	}
	_ = res
}
