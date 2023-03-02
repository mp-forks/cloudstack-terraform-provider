package cloudstack

import (
	"fmt"
	"testing"

	"github.com/apache/cloudstack-go/v2/cloudstack"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccCloudStackConfiguration_basic(t *testing.T) {
	var configuration cloudstack.ListConfigurationsResponse

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConfiguration(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudStackConfigurationExists("cloudstack_configuration.test", &configuration),
					testAccCheckCloudStackConfigurationAttributes(&configuration),
				),
			},
		},
	})
}

func TestAccCloudStackConfiguration_update(t *testing.T) {
	var configuration cloudstack.ListConfigurationsResponse
	// var configuration_update cloudstack.UpdateConfigurationResponse

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConfiguration(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudStackConfigurationExists("cloudstack_configuration.test", &configuration),
					testAccCheckCloudStackConfigurationAttributes(&configuration),
					resource.TestCheckResourceAttr("cloudstack_configuration.test", "value", "test_host"),
				),
			},

			{
				Config: testAccResourceConfiguration_update(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudStackConfigurationExists("cloudstack_configuration.test", &configuration),
					// testAccCheckCloudStackConfigurationUpdate(&configuration),
					resource.TestCheckResourceAttr("cloudstack_configuration.test", "value", "new_test_host"),
				),
			},
		},
	})
}

func TestAccCloudStackConfiguration_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConfiguration(),
			},

			{
				ResourceName:      "cloudstack_configuration.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckCloudStackConfigurationExists(n string, configuration *cloudstack.ListConfigurationsResponse) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("configuration ID not set")
		}

		cs := testAccProvider.Meta().(*cloudstack.CloudStackClient)
		p := cs.Configuration.NewListConfigurationsParams()
		p.SetName(rs.Primary.ID)

		cfg, err := cs.Configuration.ListConfigurations(p)
		if err != nil {
			return err
		}

		*configuration = *cfg

		return nil
	}
}

func testAccCheckCloudStackConfigurationAttributes(configuration *cloudstack.ListConfigurationsResponse) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if configuration.Configurations[0].Name != "host" {
			return fmt.Errorf("Bad name: %s", configuration.Configurations[0].Name)
		}

		if configuration.Configurations[0].Value != "test_host" {
			return fmt.Errorf("Bad name: %s", configuration.Configurations[0].Name)
		}

		return nil
	}
}

func testAccCheckCloudStackConfigurationUpdate(configuration *cloudstack.ListConfigurationsResponse) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if configuration.Configurations[0].Name != "host" {
			return fmt.Errorf("Bad name: %s", configuration.Configurations[0].Name)
		}

		if configuration.Configurations[0].Value != "new_test_host" {
			return fmt.Errorf("Bad name: %s", configuration.Configurations[0].Name)
		}

		return nil
	}
}

func testAccResourceConfiguration() string {
	return fmt.Sprintf(`
	resource "cloudstack_configuration" "test" {
		name  = "host"
		value = "test_host"
	}
`)
}

func testAccResourceConfiguration_update() string {
	return fmt.Sprintf(`
	resource "cloudstack_configuration" "test" {
		name  = "host"
		value = "new_test_host"
	}
`)
}
