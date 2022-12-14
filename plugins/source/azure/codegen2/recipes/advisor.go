// Code generated by codegen1; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/advisor/armadvisor"

func init() {
	tables := []Table{
		{
			Service:        "armadvisor",
			Name:           "recommendation_metadata",
			Struct:         &armadvisor.MetadataEntity{},
			ResponseStruct: &armadvisor.RecommendationMetadataClientListResponse{},
			Client:         &armadvisor.RecommendationMetadataClient{},
			ListFunc:       (&armadvisor.RecommendationMetadataClient{}).NewListPager,
			NewFunc:        armadvisor.NewRecommendationMetadataClient,
			URL:            "/providers/Microsoft.Advisor/metadata",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Advisor)`,
			ExtraColumns:   DefaultExtraColumns,
		},
		{
			Service:        "armadvisor",
			Name:           "recommendations",
			Struct:         &armadvisor.ResourceRecommendationBase{},
			ResponseStruct: &armadvisor.RecommendationsClientListResponse{},
			Client:         &armadvisor.RecommendationsClient{},
			ListFunc:       (&armadvisor.RecommendationsClient{}).NewListPager,
			NewFunc:        armadvisor.NewRecommendationsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Advisor/recommendations",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Advisor)`,
			ExtraColumns:   DefaultExtraColumns,
		},
		{
			Service:        "armadvisor",
			Name:           "suppressions",
			Struct:         &armadvisor.SuppressionContract{},
			ResponseStruct: &armadvisor.SuppressionsClientListResponse{},
			Client:         &armadvisor.SuppressionsClient{},
			ListFunc:       (&armadvisor.SuppressionsClient{}).NewListPager,
			NewFunc:        armadvisor.NewSuppressionsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Advisor/suppressions",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Advisor)`,
			ExtraColumns:   DefaultExtraColumns,
		},
	}
	Tables = append(Tables, tables...)
}