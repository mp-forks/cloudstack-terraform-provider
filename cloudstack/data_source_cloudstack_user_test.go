//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

package cloudstack

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccUserDataSource_basic(t *testing.T) {
	resourceName := "cloudstack_user.user-resource"
	datasourceName := "data.cloudstack_user.user-data-source"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testUserDataSourceConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(datasourceName, "first_name", resourceName, "first_name"),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

const testUserDataSourceConfig_basic = `
resource "cloudstack_user" "user-resource" {
  account = "admin"
  email         = "jon.doe@gmail.com"
  first_name    = "jon"
  last_name     = "doe"
  password      = "password"
  username      = "jon123"
}

data "cloudstack_user" "user-data-source"{
    filter{
    name = "first_name"
    value= "jon"
    }
    depends_on = [
	  cloudstack_user.user-resource
	]
  }

output "user-output" {
  value = "${data.cloudstack_user.user-data-source}"
}
  `
