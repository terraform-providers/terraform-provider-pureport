package pureport

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/pureport/pureport-sdk-go/pureport/client"
)

const testAccResourceSiteVPNConnectionConfig_common = `
data "pureport_accounts" "main" {
  name_regex = "Terraform"
}

data "pureport_locations" "main" {
  name_regex = "^Sea*"
}

data "pureport_networks" "main" {
  account_href = "${data.pureport_accounts.main.accounts.0.href}"
  name_regex = "Bansh.*"
}
`

const testAccResourceSiteVPNConnectionConfig_route_based_bgp = testAccResourceSiteVPNConnectionConfig_common + `
resource "pureport_site_vpn_connection" "main" {
  name = "SiteVPN_RouteBasedBGP"
  speed = "100"
  high_availability = true

  location_href = "${data.pureport_locations.main.locations.0.href}"
  network_href = "${data.pureport_networks.main.networks.0.href}"

  ike_version = "V2"

  routing_type = "ROUTE_BASED_BGP"
  customer_asn = 30000

  primary_customer_router_ip = "123.123.123.123"
  secondary_customer_router_ip = "124.124.124.124"
}
`

const testAccResourceSiteVPNConnectionConfig_with_ike_config = testAccResourceSiteVPNConnectionConfig_common + `
resource "pureport_site_vpn_connection" "main" {
  name = "SiteVPN_RouteBasedBGP"
  speed = "100"
  high_availability = true
  enable_bgp_password = true

  location_href = "${data.pureport_locations.main.locations.0.href}"
  network_href = "${data.pureport_networks.main.networks.0.href}"

  ike_version = "V2"

  ike_config {
    esp {
      dh_group   = "MODP_2048"
      encryption = "AES_128"
      integrity  = "SHA256_HMAC"
    }

    ike {
      dh_group   = "MODP_2048"
      encryption = "AES_128"
      integrity  = "SHA256_HMAC"
    }
  }

  routing_type = "ROUTE_BASED_BGP"
  customer_asn = 30000

  primary_customer_router_ip = "123.123.123.123"
  secondary_customer_router_ip = "124.124.124.124"
}
`

const testAccResourceSiteVPNConnectionConfig_route_based_static = testAccResourceSiteVPNConnectionConfig_common + `
resource "pureport_site_vpn_connection" "main" {
  name = "SiteVPN_RouteBasedStatic"
  description = "Some Description"
  speed = "100"
  high_availability = true

  location_href = "${data.pureport_locations.main.locations.0.href}"
  network_href = "${data.pureport_networks.main.networks.0.href}"

  customer_networks {
    name = "Customer#1"
    address = "12.12.12.12/32"
  }

  customer_networks {
    name = "Customer#2"
    address = "34.34.34.34/32"
  }

  ike_version = "V2"

  routing_type = "ROUTE_BASED_STATIC"

  primary_customer_router_ip = "111.111.111.111"
  secondary_customer_router_ip = "222.222.222.222"
}
`

const testAccResourceSiteVPNConnectionConfig_policy_based = testAccResourceSiteVPNConnectionConfig_common + `
resource "pureport_site_vpn_connection" "main" {
  name = "SiteVPN_PolicyBased"
  speed = "100"
  high_availability = true

  location_href = "${data.pureport_locations.main.locations.0.href}"
  network_href = "${data.pureport_networks.main.networks.0.href}"

  ike_version = "V2"

  routing_type = "ROUTE_BASED_BGP"
  customer_asn = 30000

  primary_customer_router_ip = "123.123.123.123"
  secondary_customer_router_ip = "124.124.124.124"
}
`

func TestSiteVPNConnection_route_based_bgp(t *testing.T) {

	resourceName := "pureport_site_vpn_connection.main"
	var instance client.SiteIpSecVpnConnection
	var respawn_instance client.SiteIpSecVpnConnection

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSiteVPNConnectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSiteVPNConnectionConfig_route_based_bgp,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckResourceSiteVPNConnection(resourceName, &instance),
					resource.TestCheckResourceAttrPtr(resourceName, "id", &instance.Id),
					resource.TestCheckResourceAttr(resourceName, "name", "SiteVPN_RouteBasedBGP"),
					resource.TestCheckResourceAttr(resourceName, "description", ""),
					resource.TestCheckResourceAttr(resourceName, "speed", "100"),
					resource.TestCheckResourceAttr(resourceName, "high_availability", "true"),

					resource.TestCheckResourceAttr(resourceName, "enable_bgp_password", "false"),
					resource.TestCheckResourceAttr(resourceName, "routing_type", "ROUTE_BASED_BGP"),
					resource.TestCheckResourceAttr(resourceName, "primary_customer_router_ip", "123.123.123.123"),
					resource.TestCheckResourceAttr(resourceName, "primary_key", ""),
					resource.TestCheckResourceAttr(resourceName, "secondary_customer_router_ip", "124.124.124.124"),
					resource.TestCheckResourceAttr(resourceName, "secondary_key", ""),

					resource.TestCheckResourceAttr(resourceName, "gateways.#", "2"),

					resource.TestCheckResourceAttr(resourceName, "gateways.0.availability_domain", "PRIMARY"),
					resource.TestCheckResourceAttr(resourceName, "gateways.0.name", "SITE_IPSEC_VPN"),
					resource.TestCheckResourceAttr(resourceName, "gateways.0.description", ""),
					resource.TestCheckResourceAttr(resourceName, "gateways.0.link_state", "PENDING"),
					resource.TestCheckResourceAttr(resourceName, "gateways.0.customer_asn", "30000"),
					resource.TestMatchResourceAttr(resourceName, "gateways.0.customer_ip", regexp.MustCompile("169.254.[0-9]{1,3}.[0-9]{1,3}")),
					resource.TestCheckResourceAttr(resourceName, "gateways.0.pureport_asn", "394351"),
					resource.TestMatchResourceAttr(resourceName, "gateways.0.pureport_ip", regexp.MustCompile("169.254.[0-9]{1,3}.[0-9]{1,3}")),
					resource.TestCheckResourceAttr(resourceName, "gateways.0.bgp_password", ""),
					resource.TestMatchResourceAttr(resourceName, "gateways.0.peering_subnet", regexp.MustCompile("169.254.[0-9]{1,3}.[0-9]{1,3}")),
					resource.TestCheckResourceAttr(resourceName, "gateways.0.public_nat_ip", ""),
					resource.TestCheckResourceAttr(resourceName, "gateways.0.customer_gateway_ip", "123.123.123.123"),
					resource.TestMatchResourceAttr(resourceName, "gateways.0.customer_vti_ip", regexp.MustCompile("169.254.[0-9]{1,3}.[0-9]{1,3}")),
					resource.TestMatchResourceAttr(resourceName, "gateways.0.pureport_gateway_ip", regexp.MustCompile("45.56.[0-9]{1,3}.[0-9]{1,3}")),
					resource.TestMatchResourceAttr(resourceName, "gateways.0.pureport_vti_ip", regexp.MustCompile("169.254.[0-9]{1,3}.[0-9]{1,3}")),
					resource.TestCheckResourceAttr(resourceName, "gateways.0.vpn_auth_type", "PSK"),
					resource.TestCheckResourceAttrSet(resourceName, "gateways.0.vpn_auth_key"),

					resource.TestCheckResourceAttr(resourceName, "gateways.1.availability_domain", "SECONDARY"),
					resource.TestCheckResourceAttr(resourceName, "gateways.1.name", "SITE_IPSEC_VPN 2"),
					resource.TestCheckResourceAttr(resourceName, "gateways.1.description", ""),
					resource.TestCheckResourceAttr(resourceName, "gateways.1.link_state", "PENDING"),
					resource.TestCheckResourceAttr(resourceName, "gateways.1.customer_asn", "30000"),
					resource.TestMatchResourceAttr(resourceName, "gateways.1.customer_ip", regexp.MustCompile("169.254.[0-9]{1,3}.[0-9]{1,3}")),
					resource.TestCheckResourceAttr(resourceName, "gateways.1.pureport_asn", "394351"),
					resource.TestMatchResourceAttr(resourceName, "gateways.1.pureport_ip", regexp.MustCompile("169.254.[0-9]{1,3}.[0-9]{1,3}")),
					resource.TestCheckResourceAttr(resourceName, "gateways.1.bgp_password", ""),
					resource.TestMatchResourceAttr(resourceName, "gateways.1.peering_subnet", regexp.MustCompile("169.254.[0-9]{1,3}.[0-9]{1,3}")),
					resource.TestCheckResourceAttr(resourceName, "gateways.1.public_nat_ip", ""),
					resource.TestCheckResourceAttr(resourceName, "gateways.1.customer_gateway_ip", "124.124.124.124"),
					resource.TestMatchResourceAttr(resourceName, "gateways.1.customer_vti_ip", regexp.MustCompile("169.254.[0-9]{1,3}.[0-9]{1,3}")),
					resource.TestMatchResourceAttr(resourceName, "gateways.1.pureport_gateway_ip", regexp.MustCompile("45.56.[0-9]{1,3}.[0-9]{1,3}")),
					resource.TestMatchResourceAttr(resourceName, "gateways.1.pureport_vti_ip", regexp.MustCompile("169.254.[0-9]{1,3}.[0-9]{1,3}")),
					resource.TestCheckResourceAttr(resourceName, "gateways.1.vpn_auth_type", "PSK"),
					resource.TestCheckResourceAttrSet(resourceName, "gateways.1.vpn_auth_key"),
				),
			},
			{
				Config: testAccResourceSiteVPNConnectionConfig_route_based_static,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckResourceSiteVPNConnection(resourceName, &respawn_instance),
					resource.TestCheckResourceAttrPtr(resourceName, "id", &respawn_instance.Id),
					resource.TestCheckResourceAttr(resourceName, "name", "SiteVPN_RouteBasedStatic"),
					resource.TestCheckResourceAttr(resourceName, "description", "Some Description"),
					resource.TestCheckResourceAttr(resourceName, "speed", "100"),
					resource.TestCheckResourceAttr(resourceName, "high_availability", "true"),

					resource.TestCheckResourceAttr(resourceName, "enable_bgp_password", "false"),
					resource.TestCheckResourceAttr(resourceName, "routing_type", "ROUTE_BASED_STATIC"),
					resource.TestCheckResourceAttr(resourceName, "primary_customer_router_ip", "111.111.111.111"),
					resource.TestCheckResourceAttr(resourceName, "primary_key", ""),
					resource.TestCheckResourceAttr(resourceName, "secondary_customer_router_ip", "222.222.222.222"),
					resource.TestCheckResourceAttr(resourceName, "secondary_key", ""),

					resource.TestCheckResourceAttr(resourceName, "gateways.#", "2"),

					resource.TestCheckResourceAttr(resourceName, "gateways.0.availability_domain", "PRIMARY"),
					resource.TestCheckResourceAttr(resourceName, "gateways.0.name", "SITE_IPSEC_VPN"),
					resource.TestCheckResourceAttr(resourceName, "gateways.0.description", ""),
					resource.TestCheckResourceAttr(resourceName, "gateways.0.link_state", "PENDING"),
					resource.TestCheckResourceAttr(resourceName, "gateways.0.customer_asn", "0"),
					resource.TestCheckResourceAttr(resourceName, "gateways.0.customer_ip", ""),
					resource.TestCheckResourceAttr(resourceName, "gateways.0.pureport_asn", "0"),
					resource.TestCheckResourceAttr(resourceName, "gateways.0.pureport_ip", ""),
					resource.TestCheckResourceAttr(resourceName, "gateways.0.bgp_password", ""),
					resource.TestCheckResourceAttr(resourceName, "gateways.0.peering_subnet", ""),
					resource.TestCheckResourceAttr(resourceName, "gateways.0.public_nat_ip", ""),
					resource.TestCheckResourceAttr(resourceName, "gateways.0.customer_gateway_ip", "111.111.111.111"),
					resource.TestMatchResourceAttr(resourceName, "gateways.0.customer_vti_ip", regexp.MustCompile("169.254.[0-9]{1,3}.[0-9]{1,3}")),
					resource.TestMatchResourceAttr(resourceName, "gateways.0.pureport_gateway_ip", regexp.MustCompile("45.56.[0-9]{1,3}.[0-9]{1,3}")),
					resource.TestMatchResourceAttr(resourceName, "gateways.0.pureport_vti_ip", regexp.MustCompile("169.254.[0-9]{1,3}.[0-9]{1,3}")),
					resource.TestCheckResourceAttr(resourceName, "gateways.0.vpn_auth_type", "PSK"),
					resource.TestCheckResourceAttrSet(resourceName, "gateways.0.vpn_auth_key"),

					resource.TestCheckResourceAttr(resourceName, "gateways.1.availability_domain", "SECONDARY"),
					resource.TestCheckResourceAttr(resourceName, "gateways.1.name", "SITE_IPSEC_VPN 2"),
					resource.TestCheckResourceAttr(resourceName, "gateways.1.description", ""),
					resource.TestCheckResourceAttr(resourceName, "gateways.1.link_state", "PENDING"),
					resource.TestCheckResourceAttr(resourceName, "gateways.1.customer_asn", "0"),
					resource.TestCheckResourceAttr(resourceName, "gateways.1.customer_ip", ""),
					resource.TestCheckResourceAttr(resourceName, "gateways.1.pureport_asn", "0"),
					resource.TestCheckResourceAttr(resourceName, "gateways.1.pureport_ip", ""),
					resource.TestCheckResourceAttr(resourceName, "gateways.1.bgp_password", ""),
					resource.TestCheckResourceAttr(resourceName, "gateways.1.peering_subnet", ""),
					resource.TestCheckResourceAttr(resourceName, "gateways.1.public_nat_ip", ""),
					resource.TestCheckResourceAttr(resourceName, "gateways.1.customer_gateway_ip", "222.222.222.222"),
					resource.TestMatchResourceAttr(resourceName, "gateways.1.customer_vti_ip", regexp.MustCompile("169.254.[0-9]{1,3}.[0-9]{1,3}")),
					resource.TestMatchResourceAttr(resourceName, "gateways.1.pureport_gateway_ip", regexp.MustCompile("45.56.[0-9]{1,3}.[0-9]{1,3}")),
					resource.TestMatchResourceAttr(resourceName, "gateways.1.pureport_vti_ip", regexp.MustCompile("169.254.[0-9]{1,3}.[0-9]{1,3}")),
					resource.TestCheckResourceAttr(resourceName, "gateways.1.vpn_auth_type", "PSK"),
					resource.TestCheckResourceAttrSet(resourceName, "gateways.1.vpn_auth_key"),
				),
			},
		},
	})
}

func TestSiteVPNConnection_with_ikeconfig(t *testing.T) {

	resourceName := "pureport_site_vpn_connection.main"
	var instance client.SiteIpSecVpnConnection

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSiteVPNConnectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSiteVPNConnectionConfig_with_ike_config,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckResourceSiteVPNConnection(resourceName, &instance),
					resource.TestCheckResourceAttrPtr(resourceName, "id", &instance.Id),
					resource.TestCheckResourceAttr(resourceName, "name", "SiteVPN_RouteBasedBGP"),
					resource.TestCheckResourceAttr(resourceName, "description", ""),
					resource.TestCheckResourceAttr(resourceName, "speed", "100"),
					resource.TestCheckResourceAttr(resourceName, "high_availability", "true"),

					resource.TestCheckResourceAttr(resourceName, "enable_bgp_password", "true"),
					resource.TestCheckResourceAttr(resourceName, "routing_type", "ROUTE_BASED_BGP"),
					resource.TestCheckResourceAttr(resourceName, "primary_customer_router_ip", "123.123.123.123"),
					resource.TestCheckResourceAttr(resourceName, "primary_key", ""),
					resource.TestCheckResourceAttr(resourceName, "secondary_customer_router_ip", "124.124.124.124"),
					resource.TestCheckResourceAttr(resourceName, "secondary_key", ""),

					resource.TestCheckResourceAttr(resourceName, "gateways.#", "2"),

					resource.TestCheckResourceAttr(resourceName, "gateways.0.availability_domain", "PRIMARY"),
					resource.TestCheckResourceAttr(resourceName, "gateways.0.name", "SITE_IPSEC_VPN"),
					resource.TestCheckResourceAttr(resourceName, "gateways.0.description", ""),
					resource.TestCheckResourceAttr(resourceName, "gateways.0.link_state", "PENDING"),
					resource.TestCheckResourceAttr(resourceName, "gateways.0.customer_asn", "30000"),
					resource.TestMatchResourceAttr(resourceName, "gateways.0.customer_ip", regexp.MustCompile("169.254.[0-9]{1,3}.[0-9]{1,3}")),
					resource.TestCheckResourceAttr(resourceName, "gateways.0.pureport_asn", "394351"),
					resource.TestMatchResourceAttr(resourceName, "gateways.0.pureport_ip", regexp.MustCompile("169.254.[0-9]{1,3}.[0-9]{1,3}")),
					resource.TestCheckResourceAttrSet(resourceName, "gateways.0.bgp_password"),
					resource.TestMatchResourceAttr(resourceName, "gateways.0.peering_subnet", regexp.MustCompile("169.254.[0-9]{1,3}.[0-9]{1,3}")),
					resource.TestCheckResourceAttr(resourceName, "gateways.0.public_nat_ip", ""),
					resource.TestCheckResourceAttr(resourceName, "gateways.0.customer_gateway_ip", "123.123.123.123"),
					resource.TestMatchResourceAttr(resourceName, "gateways.0.customer_vti_ip", regexp.MustCompile("169.254.[0-9]{1,3}.[0-9]{1,3}")),
					resource.TestMatchResourceAttr(resourceName, "gateways.0.pureport_gateway_ip", regexp.MustCompile("45.56.[0-9]{1,3}.[0-9]{1,3}")),
					resource.TestMatchResourceAttr(resourceName, "gateways.0.pureport_vti_ip", regexp.MustCompile("169.254.[0-9]{1,3}.[0-9]{1,3}")),
					resource.TestCheckResourceAttr(resourceName, "gateways.0.vpn_auth_type", "PSK"),
					resource.TestCheckResourceAttrSet(resourceName, "gateways.0.vpn_auth_key"),

					resource.TestCheckResourceAttr(resourceName, "gateways.1.availability_domain", "SECONDARY"),
					resource.TestCheckResourceAttr(resourceName, "gateways.1.name", "SITE_IPSEC_VPN 2"),
					resource.TestCheckResourceAttr(resourceName, "gateways.1.description", ""),
					resource.TestCheckResourceAttr(resourceName, "gateways.1.link_state", "PENDING"),
					resource.TestCheckResourceAttr(resourceName, "gateways.1.customer_asn", "30000"),
					resource.TestMatchResourceAttr(resourceName, "gateways.1.customer_ip", regexp.MustCompile("169.254.[0-9]{1,3}.[0-9]{1,3}")),
					resource.TestCheckResourceAttr(resourceName, "gateways.1.pureport_asn", "394351"),
					resource.TestMatchResourceAttr(resourceName, "gateways.1.pureport_ip", regexp.MustCompile("169.254.[0-9]{1,3}.[0-9]{1,3}")),
					resource.TestCheckResourceAttrSet(resourceName, "gateways.1.bgp_password"),
					resource.TestMatchResourceAttr(resourceName, "gateways.1.peering_subnet", regexp.MustCompile("169.254.[0-9]{1,3}.[0-9]{1,3}")),
					resource.TestCheckResourceAttr(resourceName, "gateways.1.public_nat_ip", ""),
					resource.TestCheckResourceAttr(resourceName, "gateways.1.customer_gateway_ip", "124.124.124.124"),
					resource.TestMatchResourceAttr(resourceName, "gateways.1.customer_vti_ip", regexp.MustCompile("169.254.[0-9]{1,3}.[0-9]{1,3}")),
					resource.TestMatchResourceAttr(resourceName, "gateways.1.pureport_gateway_ip", regexp.MustCompile("45.56.[0-9]{1,3}.[0-9]{1,3}")),
					resource.TestMatchResourceAttr(resourceName, "gateways.1.pureport_vti_ip", regexp.MustCompile("169.254.[0-9]{1,3}.[0-9]{1,3}")),
					resource.TestCheckResourceAttr(resourceName, "gateways.1.vpn_auth_type", "PSK"),
					resource.TestCheckResourceAttrSet(resourceName, "gateways.1.vpn_auth_key"),
				),
			},
		},
	})
}

func testAccCheckResourceSiteVPNConnection(name string, instance *client.SiteIpSecVpnConnection) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		config, ok := testAccProvider.Meta().(*Config)
		if !ok {
			return fmt.Errorf("Error getting Pureport client")
		}

		// Find the state object
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Can't find SiteVPN Connection resource: %s", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		id := rs.Primary.ID

		ctx := config.Session.GetSessionContext()
		found, resp, err := config.Session.Client.ConnectionsApi.GetConnection(ctx, id)

		if err != nil {
			return fmt.Errorf("receive error when requesting SiteVPN Connection %s", id)
		}

		if resp.StatusCode != 200 {
			return fmt.Errorf("Error getting SiteVPN Connection ID %s: %s", id, err)
		}

		*instance = found.(client.SiteIpSecVpnConnection)

		return nil
	}
}

func testAccCheckSiteVPNConnectionDestroy(s *terraform.State) error {

	config, ok := testAccProvider.Meta().(*Config)
	if !ok {
		return fmt.Errorf("Error getting Pureport client")
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "pureport_aws_connection" {
			continue
		}

		id := rs.Primary.ID

		ctx := config.Session.GetSessionContext()
		_, resp, err := config.Session.Client.ConnectionsApi.GetConnection(ctx, id)

		if err != nil && resp.StatusCode != 404 {
			return fmt.Errorf("should not get error for SiteVPN Connection with ID %s after delete: %s", id, err)
		}
	}

	return nil
}
