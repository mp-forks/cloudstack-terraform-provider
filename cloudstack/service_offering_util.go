package cloudstack

import (
	"context"
	"strings"

	"github.com/apache/cloudstack-go/v2/cloudstack"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func createServiceOffering(_ context.Context, client *cloudstack.CloudStackClient, plan serviceOfferingResourceModel) (serviceOfferingResourceModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Set params
	p := client.ServiceOffering.NewCreateServiceOfferingParams(plan.DisplayText.ValueString(), plan.Name.ValueString())

	if !plan.DeploymentPlanner.IsNull() && !plan.DeploymentPlanner.IsUnknown() {
		p.SetDeploymentplanner(plan.DeploymentPlanner.ValueString())
	} else {
		plan.DeploymentPlanner = types.StringNull()
	}
	if !plan.DiskOfferingId.IsNull() {
		p.SetDiskofferingid(plan.DiskOfferingId.ValueString())
	}
	if !plan.DomainIds.IsNull() {
		domainids := make([]string, len(plan.DomainIds.Elements()))
		for i, v := range plan.DomainIds.Elements() {
			domainids[i] = v.String()
		}
		p.SetDomainid(domainids)
	}
	if !plan.DynamicScalingEnabled.IsNull() {
		p.SetDynamicscalingenabled(plan.DynamicScalingEnabled.ValueBool())
	}
	if !plan.HostTags.IsNull() {
		p.SetHosttags(plan.HostTags.ValueString())
	}
	if !plan.IsVolatile.IsNull() {
		p.SetIsvolatile(plan.IsVolatile.ValueBool())
	}
	if !plan.LimitCpuUse.IsNull() {
		p.SetLimitcpuuse(plan.LimitCpuUse.ValueBool())
	}
	if !plan.NetworkRate.IsNull() {
		p.SetNetworkrate(int(plan.NetworkRate.ValueInt64()))
	}
	if !plan.OfferHa.IsNull() {
		p.SetOfferha(plan.OfferHa.ValueBool())
	}
	if !plan.StorageTags.IsNull() {
		p.SetTags(plan.StorageTags.ValueString())
	}
	if !plan.ZoneIds.IsNull() {
		zids := make([]string, len(plan.ZoneIds.Elements()))
		for i, v := range plan.ZoneIds.Elements() {
			zids[i] = v.String()
		}
		p.SetZoneid(zids)
	}

	// Createing offering
	cs, err := client.ServiceOffering.CreateServiceOffering(p)
	if err != nil {
		diags.AddError(
			"Error creating order",
			"Could not create order, unexpected error: "+err.Error(),
		)
		return plan, diags
	}

	// set computed values
	plan.ID = types.StringValue(cs.Id)

	return plan, diags
}

func readServiceOffering(ctx context.Context, client *cloudstack.CloudStackClient, statea interface{}) (serviceOfferingResourceModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	state, _ := statea.(serviceOfferingResourceModel)
	//
	cs, _, err := client.ServiceOffering.GetServiceOfferingByID(state.ID.ValueString())
	if err != nil {
		diags.AddError(
			"Error reading Service Offering",
			"Could not Service Offering, unexpected error: "+err.Error(),
		)
		return state, diags
	}

	// Update state
	if cs.Deploymentplanner != "" {
		state.DeploymentPlanner = types.StringValue(cs.Deploymentplanner)
	}
	if cs.Diskofferingid != "" {
		state.DiskOfferingId = types.StringValue(cs.Diskofferingid)
	}
	if cs.Displaytext != "" {
		state.DisplayText = types.StringValue(cs.Displaytext)
	}
	if cs.Domainid != "" {
		state.DomainIds, _ = types.SetValueFrom(ctx, types.StringType, strings.Split(cs.Domainid, ","))
	}
	if cs.Hosttags != "" {
		state.HostTags = types.StringValue(cs.Hosttags)
	}
	if cs.Name != "" {
		state.Name = types.StringValue(cs.Name)
	}
	if cs.Networkrate > 0 {
		state.NetworkRate = types.Int64Value(int64(cs.Networkrate))
	}
	if cs.Storagetags != "" {
		state.StorageTags = types.StringValue(cs.Storagetags)
	}
	if cs.Zoneid != "" {
		state.ZoneIds, _ = types.SetValueFrom(ctx, types.StringType, strings.Split(cs.Zoneid, ","))
	}

	//
	state.DynamicScalingEnabled = types.BoolValue(cs.Dynamicscalingenabled)
	state.IsVolatile = types.BoolValue(cs.Isvolatile)
	state.LimitCpuUse = types.BoolValue(cs.Limitcpuuse)
	state.OfferHa = types.BoolValue(cs.Offerha)

	return state, diags
}
