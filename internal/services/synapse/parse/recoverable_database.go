package parse

// NOTE: this file is generated via 'go:generate' - manual changes will be overwritten

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

type RecoverableDatabaseId struct {
	SubscriptionId string
	ResourceGroup  string
	WorkspaceName  string
	Name           string
}

func NewRecoverableDatabaseID(subscriptionId, resourceGroup, workspaceName, name string) RecoverableDatabaseId {
	return RecoverableDatabaseId{
		SubscriptionId: subscriptionId,
		ResourceGroup:  resourceGroup,
		WorkspaceName:  workspaceName,
		Name:           name,
	}
}

func (id RecoverableDatabaseId) String() string {
	segments := []string{
		fmt.Sprintf("Name %q", id.Name),
		fmt.Sprintf("Workspace Name %q", id.WorkspaceName),
		fmt.Sprintf("Resource Group %q", id.ResourceGroup),
	}
	segmentsStr := strings.Join(segments, " / ")
	return fmt.Sprintf("%s: (%s)", "Recoverable Database", segmentsStr)
}

func (id RecoverableDatabaseId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Synapse/workspaces/%s/recoverableDatabases/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroup, id.WorkspaceName, id.Name)
}

// RecoverableDatabaseID parses a RecoverableDatabase ID into an RecoverableDatabaseId struct
func RecoverableDatabaseID(input string) (*RecoverableDatabaseId, error) {
	id, err := resourceids.ParseAzureResourceID(input)
	if err != nil {
		return nil, err
	}

	resourceId := RecoverableDatabaseId{
		SubscriptionId: id.SubscriptionID,
		ResourceGroup:  id.ResourceGroup,
	}

	if resourceId.SubscriptionId == "" {
		return nil, fmt.Errorf("ID was missing the 'subscriptions' element")
	}

	if resourceId.ResourceGroup == "" {
		return nil, fmt.Errorf("ID was missing the 'resourceGroups' element")
	}

	if resourceId.WorkspaceName, err = id.PopSegment("workspaces"); err != nil {
		return nil, err
	}
	if resourceId.Name, err = id.PopSegment("recoverableDatabases"); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &resourceId, nil
}
