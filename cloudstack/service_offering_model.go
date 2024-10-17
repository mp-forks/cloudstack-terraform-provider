package cloudstack

import "github.com/hashicorp/terraform-plugin-framework/types"

type serviceOfferingConstrainedResourceModel struct {
	serviceOfferingResourceModel
	CpuSpeed     types.String `tfsdk:"cpu_speed"`
	MaxCpuNumber types.String `tfsdk:"max_cpu_number"`
	MaxMemory    types.String `tfsdk:"max_memory"`
	MinCpuNumber types.String `tfsdk:"min_cpu_number"`
	MinMemory    types.String `tfsdk:"min_memory"`
	// customized types.String `tfsdk:"Iscustomized"`
}

type serviceOfferingFixedResourceModel struct {
	serviceOfferingResourceModel
	CpuNumber types.String `tfsdk:"cpu_number"`
	CpuSpeed  types.String `tfsdk:"cpu_speed"`
	Memory    types.String `tfsdk:"memory"`
}

type serviceOfferingUnconstrainedResourceModel struct {
	serviceOfferingResourceModel
}

type serviceOfferingResourceModel struct {
	DeploymentPlanner     types.String `tfsdk:"deployment_planner"`
	DiskOfferingId        types.String `tfsdk:"disk_offering_id"`
	DisplayText           types.String `tfsdk:"display_text"`
	DomainIds             types.Set    `tfsdk:"domain_ids"`
	DynamicScalingEnabled types.Bool   `tfsdk:"dynamic_scaling_enabled"`
	HostTags              types.String `tfsdk:"host_tags"`
	ID                    types.String `tfsdk:"id"`
	IsVolatile            types.Bool   `tfsdk:"is_volatile"`
	LimitCpuUse           types.Bool   `tfsdk:"limit_cpu_use"`
	Name                  types.String `tfsdk:"name"`
	NetworkRate           types.Int64  `tfsdk:"network_rate"`
	OfferHa               types.Bool   `tfsdk:"offer_ha"`
	StorageTags           types.String `tfsdk:"storage_tags"`
	ZoneIds               types.Set    `tfsdk:"zone_ids"`
	//
	disk_hypervisor types.Map `tfsdk:"disk_qos_storage"`
	disk_offering   types.Map `tfsdk:"disk_offering"`
	disk_storage    types.Map `tfsdk:"disk_storage"`
}

type serviceOfferingDiskQosHypervisor struct {
	DiskBytesReadRate           types.Int32 `tfsdk:"bytes_read_rate"`
	DiskBytesReadRateMax        types.Int32 `tfsdk:"bytes_read_rate_max"`
	DiskBytesReadRateMaxLength  types.Int32 `tfsdk:"bytes_read_rate_max_length"`
	DiskBytesWriteRate          types.Int32 `tfsdk:"bytes_write_rate"`
	DiskBytesWriteRateMax       types.Int32 `tfsdk:"bytes_write_rate_max"`
	DiskBytesWriteRateMaxLength types.Int32 `tfsdk:"bytes_write_rate_max_length"`
}

type serviceOfferingDisk struct {
	CacheMode              types.Int32  `tfsdk:"cache_mode"`
	DiskOfferingStrictness types.Bool   `tfsdk:"disk_offering_strictness"`
	ProvisionType          types.Int32  `tfsdk:"provisioning_type"`
	RootDiskSize           types.Int32  `tfsdk:"root_disk_size"`
	StorageType            types.String `tfsdk:"storage_type"`
}

type serviceOfferingDiskQosStorage struct {
	CustomizedIops            types.Bool  `tfsdk:"customized_iops"`
	HypervisorSnapshotReserve types.Int32 `tfsdk:"hypervisor_snapshot_reserve"`
	MaxIops                   types.Int32 `tfsdk:"max_iops"`
	MinIops                   types.Int32 `tfsdk:"min_iops"`
}
