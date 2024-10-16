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

// import (
// 	"fmt"
// 	"strconv"

// 	"github.com/apache/cloudstack-go/v2/cloudstack"
// 	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
// )

// func serviceOfferingMergeCommonSchema(s1 map[string]*schema.Schema) map[string]*schema.Schema {
// 	common := map[string]*schema.Schema{
// 		// required
// 		"display_text": {
// 			Type:     schema.TypeString,
// 			Required: true,
// 		},
// 		"name": {
// 			Type:     schema.TypeString,
// 			Required: true,
// 		},
// 		// optional
// 		"deployment_planner": {
// 			Type:     schema.TypeString,
// 			Optional: true,
// 			ForceNew: true,
// 		},
// 		"disk_offering_id": {
// 			Type:     schema.TypeString,
// 			Optional: true,
// 			ForceNew: true,
// 		},
// 		"domain_id": {
// 			Type:     schema.TypeList,
// 			Optional: true,
// 			Elem: &schema.Schema{
// 				Type: schema.TypeString,
// 			},
// 		},
// 		"dynamic_scaling_enabled": {
// 			Type:     schema.TypeBool,
// 			Optional: true,
// 			ForceNew: true,
// 		},
// 		"host_tags": {
// 			Description: "The host tag for this service offering",
// 			Type:        schema.TypeString,
// 			Optional:    true,
// 		},
// 		"is_volatile": {
// 			Type:     schema.TypeBool,
// 			Optional: true,
// 			ForceNew: true,
// 		},
// 		"limit_cpu_use": {
// 			Description: "Restrict the CPU usage to committed service offering",
// 			Type:        schema.TypeBool,
// 			Optional:    true,
// 			ForceNew:    true,
// 		},
// 		"network_rate": {
// 			Type:     schema.TypeInt,
// 			Optional: true,
// 			ForceNew: true,
// 		},
// 		"offer_ha": {
// 			Description: "The HA for the service offering",
// 			Type:        schema.TypeBool,
// 			Optional:    true,
// 			ForceNew:    true,
// 		},
// 		"tags": {
// 			Type:     schema.TypeString,
// 			Optional: true,
// 		},
// 		"zone_id": {
// 			Type:     schema.TypeSet,
// 			Optional: true,
// 			Elem: &schema.Schema{
// 				Type: schema.TypeString,
// 			},
// 		},
// 		// "disk_hypervisor": serviceOfferingDiskQosHypervisor(),
// 		// "disk_offering":   serviceOfferingDisk(),
// 		// "disk_storage": serviceOfferingDiskQosStorage(),
// 		//
// 		"disk_storage": {
// 			Type:     schema.TypeList,
// 			Optional: true,
// 			Elem:     serviceOfferingDiskQosStorage2(),
// 		},
// 	}

// 	for k, v := range s1 {
// 		common[k] = v
// 	}

// 	return common

// }

// // #################################################################
// // #################################################################
// func serviceOfferingDiskQosStorage2() *schema.Resource {
// 	return &schema.Resource{
// 		Schema: map[string]*schema.Schema{
// 			"customized_iops": {
// 				Type:     schema.TypeBool,
// 				Optional: true,
// 				Computed: true,
// 				ForceNew: true,
// 			},
// 			"hypervisor_snapshot_reserve": {
// 				Type:     schema.TypeInt,
// 				Optional: true,
// 				ForceNew: true,
// 			},
// 			"max_iops": {
// 				Type:     schema.TypeInt,
// 				Optional: true,
// 				Computed: true,
// 				ForceNew: true,
// 			},
// 			"min_iops": {
// 				Type:     schema.TypeInt,
// 				Optional: true,
// 				Computed: true,
// 				ForceNew: true,
// 			},
// 		},
// 	}
// }

// // #################################################################
// // #################################################################
// func serviceOfferingDiskQosHypervisor() *schema.Schema {
// 	return &schema.Schema{
// 		Type:     schema.TypeMap,
// 		Optional: true,
// 		Elem: &schema.Resource{
// 			Schema: map[string]*schema.Schema{
// 				"hypervisor": {
// 					Type:     schema.TypeMap,
// 					Optional: true,
// 					Elem: &schema.Resource{
// 						Schema: map[string]*schema.Schema{
// 							"bytes_read_rate": {
// 								Type:     schema.TypeInt,
// 								Optional: true,
// 								ForceNew: true,
// 							},
// 							"bytes_read_rate_max": {
// 								Type:     schema.TypeInt,
// 								Optional: true,
// 								ForceNew: true,
// 							},
// 							"bytes_read_rate_max_length": {
// 								Type:     schema.TypeInt,
// 								Optional: true,
// 								ForceNew: true,
// 							},
// 							"bytes_write_rate": {
// 								Type:     schema.TypeInt,
// 								Optional: true,
// 								ForceNew: true,
// 							},
// 							"bytes_write_rate_max": {
// 								Type:     schema.TypeInt,
// 								Optional: true,
// 								ForceNew: true,
// 							},
// 							"bytes_write_rate_max_length": {
// 								Type:     schema.TypeInt,
// 								Optional: true,
// 								ForceNew: true,
// 							},
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}
// }

// func serviceOfferingDisk() *schema.Schema {
// 	return &schema.Schema{
// 		Type:     schema.TypeMap,
// 		Optional: true,
// 		Elem: &schema.Resource{
// 			Schema: map[string]*schema.Schema{
// 				"cache_mode": {
// 					Type:     schema.TypeInt,
// 					Required: true,
// 				},
// 				"disk_offering_strictness": {
// 					Type:     schema.TypeBool,
// 					Optional: true,
// 					ForceNew: true,
// 				},
// 				"provisioning_type": {
// 					Type:     schema.TypeInt,
// 					Required: true,
// 				},
// 				"root_disk_size": {
// 					Type:     schema.TypeInt,
// 					Optional: true,
// 					Computed: true,
// 					ForceNew: true,
// 				},
// 				"storage_type": {
// 					Description: "The storage type of the service offering. Values are local and shared",
// 					Type:        schema.TypeString,
// 					Required:    true,
// 					ForceNew:    true,
// 					Default:     "shared",
// 					ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
// 						v := val.(string)

// 						if v == "local" || v == "shared" {
// 							return
// 						}

// 						errs = append(errs, fmt.Errorf("storage type should be either local or shared, got %s", v))

// 						return
// 					},
// 				},
// 			},
// 		},
// 	}
// }

// func serviceOfferingDiskQosStorage() *schema.Schema {
// 	return &schema.Schema{
// 		Type:     schema.TypeMap,
// 		Optional: true,
// 		Elem: &schema.Resource{
// 			Schema: map[string]*schema.Schema{
// 				"storage": {
// 					Type:     schema.TypeMap,
// 					Optional: true,
// 					Elem: &schema.Resource{
// 						Schema: map[string]*schema.Schema{
// 							"customized_iops": {
// 								Type:     schema.TypeBool,
// 								Optional: true,
// 								Computed: true,
// 								ForceNew: true,
// 							},
// 							"hypervisor_snapshot_reserve": {
// 								Type:     schema.TypeInt,
// 								Optional: true,
// 								ForceNew: true,
// 							},
// 							"max_iops": {
// 								Type:     schema.TypeInt,
// 								Optional: true,
// 								Computed: true,
// 								ForceNew: true,
// 							},
// 							"min_iops": {
// 								Type:     schema.TypeInt,
// 								Optional: true,
// 								Computed: true,
// 								ForceNew: true,
// 							},
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}
// }

// func serviceOfferingCreateParams(p *cloudstack.CreateServiceOfferingParams, d *schema.ResourceData) *cloudstack.CreateServiceOfferingParams {
// 	// other
// 	if v, ok := d.GetOk("host_tags"); ok {
// 		p.SetHosttags(v.(string))
// 	}
// 	if v, ok := d.GetOk("network_rate"); ok {
// 		p.SetNetworkrate(v.(int))
// 	}
// 	if v, ok := d.GetOk("deployment_planner"); ok {
// 		p.SetDeploymentplanner(v.(string))
// 	}
// 	if v, ok := d.GetOk("disk_offering_id"); ok {
// 		p.SetDiskofferingid(v.(string))
// 	}
// 	if v, ok := d.GetOk("tags"); ok {
// 		p.SetTags(v.(string))
// 	}

// 	// Features flags
// 	p.SetDynamicscalingenabled(d.Get("dynamic_scaling_enabled").(bool))
// 	p.SetIsvolatile(d.Get("is_volatile").(bool))
// 	p.SetLimitcpuuse(d.Get("limit_cpu_use").(bool))
// 	p.SetOfferha(d.Get("offer_ha").(bool))

// 	// access
// 	if v, ok := d.GetOk("domain_id"); ok {
// 		domain_id := v.([]interface{})
// 		items := make([]string, len(domain_id))
// 		for i, raw := range domain_id {
// 			items[i] = raw.(string)
// 		}
// 		p.SetDomainid(items)
// 	}

// 	if v, ok := d.GetOk("zone_id"); ok {
// 		zone_id := v.(*schema.Set).List()
// 		items := make([]string, len(zone_id))
// 		for i, raw := range zone_id {
// 			items[i] = raw.(string)
// 		}
// 		p.SetZoneid(items)
// 	}

// 	// disk offering
// 	if v, ok := d.GetOk("disk_offering"); ok {
// 		offering := v.(map[string]interface{})

// 		if v2, ok2 := offering["storage_type"]; ok2 {
// 			p.SetStoragetype(v2.(string))
// 		}
// 		if v2, ok2 := offering["provisioning_type"]; ok2 {
// 			p.SetProvisioningtype(v2.(string))
// 		}
// 		if v2, ok2 := offering["cache_mode"]; ok2 {
// 			p.SetCachemode(v2.(string))
// 		}
// 		if v2, ok2 := offering["root_disk_size"]; ok2 {
// 			tmp, _ := strconv.Atoi(v2.(string))
// 			p.SetRootdisksize(int64(tmp))
// 		}
// 		if v2, ok2 := offering["disk_offering_strictness"]; ok2 {
// 			tmp, _ := strconv.ParseBool(v2.(string))
// 			p.SetDiskofferingstrictness(tmp)
// 		}
// 	}

// 	// hypervisor qos
// 	if v, ok := d.GetOk("disk_hypervisor"); ok {
// 		hypervisor := v.(map[string]interface{})

// 		if v2, ok2 := hypervisor["bytes_read_rate"]; ok2 {
// 			tmp, _ := strconv.Atoi(v2.(string))
// 			p.SetBytesreadrate(int64(tmp))
// 		}
// 		if v2, ok2 := hypervisor["bytes_read_rate_max"]; ok2 {
// 			tmp, _ := strconv.Atoi(v2.(string))
// 			p.SetBytesreadrate(int64(tmp))
// 		}
// 		if v2, ok2 := hypervisor["bytes_read_rate_max_length"]; ok2 {
// 			tmp, _ := strconv.Atoi(v2.(string))
// 			p.SetBytesreadrate(int64(tmp))
// 		}
// 		if v2, ok2 := hypervisor["bytes_write_rate"]; ok2 {
// 			tmp, _ := strconv.Atoi(v2.(string))
// 			p.SetBytesreadrate(int64(tmp))
// 		}
// 		if v2, ok2 := hypervisor["bytes_write_rate_max"]; ok2 {
// 			tmp, _ := strconv.Atoi(v2.(string))
// 			p.SetBytesreadrate(int64(tmp))
// 		}
// 		if v2, ok2 := hypervisor["bytes_write_rate_max_length"]; ok2 {
// 			tmp, _ := strconv.Atoi(v2.(string))
// 			p.SetBytesreadrate(int64(tmp))
// 		}
// 	}

// 	// storage qos
// 	if v, ok := d.GetOk("disk_storage"); ok {
// 		storage := v.(map[string]interface{})

// 		if v2, ok2 := storage["min_iops"]; ok2 {
// 			tmp, _ := strconv.Atoi(v2.(string))
// 			p.SetMiniops(int64(tmp))
// 		}
// 		if v2, ok2 := storage["max_iops"]; ok2 {
// 			tmp, _ := strconv.Atoi(v2.(string))
// 			p.SetMaxiops(int64(tmp))
// 		}
// 		if v2, ok2 := storage["customized_iops"]; ok2 {
// 			tmp, _ := strconv.ParseBool(v2.(string))
// 			p.SetCustomizediops(tmp)
// 		}
// 		if v2, ok2 := storage["hypervisor_snapshot_reserve"]; ok2 {
// 			tmp, _ := strconv.Atoi(v2.(string))
// 			p.SetHypervisorsnapshotreserve(tmp)
// 		}
// 	}

// 	return p
// }
