package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccExampleDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: testAccExampleDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.encode_base36.test", "id", "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"),
					resource.TestCheckResourceAttr("data.encode_base36.test", "result", "5PZCSZU7"),
					resource.TestCheckResourceAttr("data.encode_base36.test_lowercase", "id", "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"),
					resource.TestCheckResourceAttr("data.encode_base36.test_lowercase", "result", "5pzcszu7"),
				),
			},
		},
	})
}

const testAccExampleDataSourceConfig = `
data "encode_base36" "test" {
  value     = "hello"
}
data "encode_base36" "test_lowercase" {
  value     = "hello"
  lowercase = true
}
`
