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

func TestAccSshKeyPairDataSource_basic(t *testing.T) {
	resourceName := "cloudstack_ssh_keypair.ssh-keypair-resource"
	datasourceName := "data.cloudstack_ssh_keypair.ssh-keypair-data"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSshKeyPairDataSourceConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(datasourceName, "id", resourceName, "id"),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

const testAccSshKeyPairDataSourceConfig_basic = `
	resource "cloudstack_ssh_keypair" "ssh-keypair-resource"{
	name       = "myKey"
  }
  
  data "cloudstack_ssh_keypair" "ssh-keypair-data" {
	  filter {
	  name = "name" 
	  value = "myKey"
	}
	  depends_on = [
	  cloudstack_ssh_keypair.ssh-keypair-resource
	]
  
  }
  `
