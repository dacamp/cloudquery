// Code generated by codegen1; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"

func init() {
	tables := []Table{
		{
			Service:        "armautomation",
			Name:           "account",
			Struct:         &armautomation.Account{},
			ResponseStruct: &armautomation.AccountClientListResponse{},
			Client:         &armautomation.AccountClient{},
			ListFunc:       (&armautomation.AccountClient{}).NewListPager,
			NewFunc:        armautomation.NewAccountClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Automation/automationAccounts",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Automation)`,
			ExtraColumns:   DefaultExtraColumns,
		},
	}
	Tables = append(Tables, tables...)
}