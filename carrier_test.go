package goshopify

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestCarrierList(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://fooshop.myshopify.com/%s/carrier_services.json", client.pathPrefix),
		httpmock.NewStringResponder(200, `{
  "carrier_services": [
    {
      "id": 1,
      "name": "Shipping Rate Provider",
      "active": true,
      "service_discovery": true,
      "carrier_service_type": "api",
      "admin_graphql_api_id": "gid://shopify/DeliveryCarrierService/1",
      "format": "json",
      "callback_url": "https://fooshop.example.com/shipping"
    }
  ]
}`))

	carriers, err := client.Carrier.List()
	if err != nil {
		t.Errorf("Carrier.List returned error: %v", err)
	}

	expected := []CarrierResource{
		{
			ID:                 1,
			Name:               "Shipping Rate Provider",
			Active:             true,
			ServiceDiscovery:   true,
			CarrierServiceType: "api",
			AdminGraphqlAPIID:  "gid://shopify/DeliveryCarrierService/1",
			Format:             "json",
			CallbackURL:        "https://fooshop.example.com/shipping",
		},
	}
	if !reflect.DeepEqual(carriers, expected) {
		t.Errorf("Carrier.List returned %+v, expected %+v", carriers, expected)
	}
}
