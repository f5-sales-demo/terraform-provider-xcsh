// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/url"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	xcsherrors "github.com/f5-sales-demo/terraform-provider-xcsh/internal/errors"
)

var (
	_                                   resource.Resource                   = &Smsv2KVMRuntimeInterfaceResource{}
	_                                   resource.ResourceWithConfigure      = &Smsv2KVMRuntimeInterfaceResource{}
	_                                   resource.ResourceWithValidateConfig = &Smsv2KVMRuntimeInterfaceResource{}
	errSMSv2KVMRuntimeInterfaceNotFound                                     = errors.New("KVM runtime interface not found")
)

type Smsv2KVMRuntimeInterfaceResource struct {
	client *client.Client
}

type Smsv2KVMRuntimeInterfaceResourceModel struct {
	ID              types.String `tfsdk:"id"`
	Namespace       types.String `tfsdk:"namespace"`
	Site            types.String `tfsdk:"site"`
	InterfaceName   types.String `tfsdk:"interface_name"`
	ExpectedMAC     types.String `tfsdk:"expected_mac"`
	Hostname        types.String `tfsdk:"hostname"`
	Device          types.String `tfsdk:"device"`
	IPv4CIDR        types.String `tfsdk:"ipv4_cidr"`
	OwnerUID        types.String `tfsdk:"owner_uid"`
	ResourceVersion types.String `tfsdk:"resource_version"`
	Configured      types.Bool   `tfsdk:"configured"`
}

type smsv2KVMRuntimeInterfaceTarget struct {
	Namespace   string
	Site        string
	Name        string
	ExpectedMAC string
	Hostname    string
	Device      string
	OwnerUID    string
}

type smsv2KVMRuntimeInterfaceObject struct {
	Name            string
	Namespace       string
	Site            string
	OwnerUID        string
	Hostname        string
	Device          string
	MAC             string
	ResourceVersion string
	Spec            map[string]interface{}
	Raw             map[string]interface{}
}

func NewSmsv2KvmRuntimeInterfaceResource() resource.Resource {
	return &Smsv2KVMRuntimeInterfaceResource{}
}

func (r *Smsv2KVMRuntimeInterfaceResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_smsv2_kvm_runtime_interface"
}

func (r *Smsv2KVMRuntimeInterfaceResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	replace := []planmodifier.String{stringplanmodifier.RequiresReplace()}
	resp.Schema = schema.Schema{
		MarkdownDescription: "Adopts one existing XC-owned KVM Secure Mesh Site v2 SLI child and manages only its DHCP/static IPv4 mode. It never creates or deletes the runtime child.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{Computed: true, MarkdownDescription: "Stable namespace/interface identity."},
			"namespace": schema.StringAttribute{
				Optional: true, Computed: true, Default: stringdefault.StaticString("system"),
				Validators: []validator.String{stringvalidator.OneOf("system")}, PlanModifiers: replace,
				MarkdownDescription: "Namespace containing the site and runtime child. KVM SMSv2 supports only `system`.",
			},
			"site":             schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}, PlanModifiers: replace, MarkdownDescription: "Secure Mesh Site v2 configuration name."},
			"interface_name":   schema.StringAttribute{Computed: true, MarkdownDescription: "Exact platform-generated SLI child name resolved from live ownership. Platform names may exceed 64 characters."},
			"expected_mac":     schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}, PlanModifiers: replace, MarkdownDescription: "Terraform-owned SLI MAC used for live registration correlation."},
			"hostname":         schema.StringAttribute{Computed: true, MarkdownDescription: "Exact live registration hostname resolved from the expected MAC."},
			"device":           schema.StringAttribute{Computed: true, MarkdownDescription: "Exact live registration device resolved from the expected MAC."},
			"ipv4_cidr":        schema.StringAttribute{Required: true, MarkdownDescription: "Static IPv4 host address and prefix to configure on the SLI."},
			"owner_uid":        schema.StringAttribute{Computed: true, MarkdownDescription: "Secure Mesh Site v2 UID that owns the runtime child."},
			"resource_version": schema.StringAttribute{Computed: true, MarkdownDescription: "Latest XC concurrency version observed after reconciliation."},
			"configured":       schema.BoolAttribute{Computed: true, MarkdownDescription: "Whether the exact owned SLI currently uses static IPv4 configuration."},
		},
	}
}

func (r *Smsv2KVMRuntimeInterfaceResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Resource Configure Type", "Expected *client.Client")
		return
	}
	r.client = c
}

func (r *Smsv2KVMRuntimeInterfaceResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data Smsv2KVMRuntimeInterfaceResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if !data.InterfaceName.IsNull() && !data.InterfaceName.IsUnknown() {
		if err := validateKVMRuntimeInterfaceName(data.InterfaceName.ValueString()); err != nil {
			resp.Diagnostics.AddAttributeError(path.Root("interface_name"), "Invalid KVM Runtime Interface Name", err.Error())
		}
	}
	if !data.ExpectedMAC.IsNull() && !data.ExpectedMAC.IsUnknown() {
		if _, err := normalizeSMSv2MAC(data.ExpectedMAC.ValueString()); err != nil {
			resp.Diagnostics.AddAttributeError(path.Root("expected_mac"), "Invalid KVM Runtime Interface MAC", err.Error())
		}
	}
	if !data.IPv4CIDR.IsNull() && !data.IPv4CIDR.IsUnknown() {
		if _, err := canonicalKVMRuntimeInterfaceCIDR(data.IPv4CIDR.ValueString()); err != nil {
			resp.Diagnostics.AddAttributeError(path.Root("ipv4_cidr"), "Invalid KVM Runtime Interface CIDR", err.Error())
		}
	}
}

func runtimeInterfaceTarget(data Smsv2KVMRuntimeInterfaceResourceModel) smsv2KVMRuntimeInterfaceTarget {
	target := smsv2KVMRuntimeInterfaceTarget{Namespace: data.Namespace.ValueString(), Site: data.Site.ValueString(), ExpectedMAC: data.ExpectedMAC.ValueString()}
	if !data.InterfaceName.IsNull() && !data.InterfaceName.IsUnknown() {
		target.Name = data.InterfaceName.ValueString()
	}
	if !data.Hostname.IsNull() && !data.Hostname.IsUnknown() {
		target.Hostname = data.Hostname.ValueString()
	}
	if !data.Device.IsNull() && !data.Device.IsUnknown() {
		target.Device = data.Device.ValueString()
	}
	if !data.OwnerUID.IsNull() && !data.OwnerUID.IsUnknown() {
		target.OwnerUID = data.OwnerUID.ValueString()
	}
	return target
}

func validateKVMRuntimeInterfaceName(name string) error {
	name = strings.TrimSpace(name)
	if name == "" || len(name) > 253 {
		return fmt.Errorf("platform-generated interface name must contain 1 to 253 characters")
	}
	if strings.ContainsAny(name, "/\\ \t\r\n") {
		return fmt.Errorf("platform-generated interface name contains a path separator or whitespace")
	}
	return nil
}

func canonicalKVMRuntimeInterfaceCIDR(value string) (string, error) {
	ip, network, err := net.ParseCIDR(strings.TrimSpace(value))
	if err != nil || ip.To4() == nil {
		return "", fmt.Errorf("static address must be an IPv4 CIDR")
	}
	ones, bits := network.Mask.Size()
	if bits != 32 || ones < 1 || ones > 32 {
		return "", fmt.Errorf("static address must have a valid IPv4 prefix")
	}
	canonicalIP := ip.To4()
	if ones < 32 && canonicalIP.Equal(network.IP.To4()) {
		return "", fmt.Errorf("static address must be a host address, not the network address")
	}
	if ones < 31 {
		broadcast := append(net.IP(nil), network.IP.To4()...)
		for index := range broadcast {
			broadcast[index] |= ^network.Mask[index]
		}
		if canonicalIP.Equal(broadcast) {
			return "", fmt.Errorf("static address must not be the broadcast address")
		}
	}
	return canonicalIP.String() + fmt.Sprintf("/%d", ones), nil
}

func deepCopySMSv2Value(value interface{}) interface{} {
	switch typed := value.(type) {
	case map[string]interface{}:
		return deepCopySMSv2Map(typed)
	case []interface{}:
		copyValue := make([]interface{}, len(typed))
		for index := range typed {
			copyValue[index] = deepCopySMSv2Value(typed[index])
		}
		return copyValue
	default:
		return typed
	}
}

func deepCopySMSv2Map(value map[string]interface{}) map[string]interface{} {
	copyValue := make(map[string]interface{}, len(value))
	for key, item := range value {
		copyValue[key] = deepCopySMSv2Value(item)
	}
	return copyValue
}

func smsv2RuntimeInterfaceItemIdentity(item map[string]interface{}) (string, string) {
	name, namespace := stringField(item, "name"), stringField(item, "namespace")
	if metadata, ok := nestedMap(item, "metadata"); ok {
		if name == "" {
			name = stringField(metadata, "name")
		}
		if namespace == "" {
			namespace = stringField(metadata, "namespace")
		}
	}
	return name, namespace
}

func smsv2RuntimeInterfaceSpec(item map[string]interface{}) (map[string]interface{}, bool) {
	if spec, ok := nestedMap(item, "get_spec"); ok {
		return spec, true
	}
	return nestedMap(item, "spec")
}

func selectSMSv2KVMRuntimeInterface(
	configuration client.SMSv2Observation,
	registrations *client.RegistrationListResponse,
	interfaces client.SMSv2Observation,
	target smsv2KVMRuntimeInterfaceTarget,
) (smsv2KVMRuntimeInterfaceObject, error) {
	if target.Namespace != "system" {
		return smsv2KVMRuntimeInterfaceObject{}, fmt.Errorf("KVM runtime interfaces are supported only in namespace system")
	}
	if target.Name != "" {
		if err := validateKVMRuntimeInterfaceName(target.Name); err != nil {
			return smsv2KVMRuntimeInterfaceObject{}, err
		}
	}
	mac, err := normalizeSMSv2MAC(target.ExpectedMAC)
	if err != nil {
		return smsv2KVMRuntimeInterfaceObject{}, err
	}
	metadata, _ := nestedMap(map[string]interface{}(configuration), "metadata")
	system, _ := nestedMap(map[string]interface{}(configuration), "system_metadata")
	site, namespace, uid := stringField(metadata, "name"), stringField(metadata, "namespace"), stringField(system, "uid")
	if site == "" || namespace == "" || uid == "" {
		return smsv2KVMRuntimeInterfaceObject{}, fmt.Errorf("site name, namespace, or UID is unavailable")
	}
	if site != target.Site || namespace != target.Namespace {
		return smsv2KVMRuntimeInterfaceObject{}, fmt.Errorf("site configuration identity does not match the requested runtime interface")
	}
	if target.OwnerUID != "" && target.OwnerUID != uid {
		return smsv2KVMRuntimeInterfaceObject{}, fmt.Errorf("runtime interface state belongs to a different Secure Mesh Site v2 UID")
	}
	if registrations == nil || len(registrations.Errors) != 0 {
		return smsv2KVMRuntimeInterfaceObject{}, fmt.Errorf("registration discovery is unavailable or contains partial errors")
	}
	type registrationMatch struct{ hostname, device string }
	registrationMatches := []registrationMatch{}
	for _, registration := range registrations.Items {
		hostname := strings.TrimSpace(registration.GetSpec.Infra.Hostname)
		if isTerminalRegistrationState(registration.Object.Status.CurrentState) || registration.GetSpec.Passport.ClusterName != target.Site || kvmRegistrationProvider(registration.GetSpec.Infra) != "KVM" || hostname == "" || (target.Hostname != "" && hostname != target.Hostname) {
			continue
		}
		for _, network := range registration.GetSpec.Infra.HWInfo.Network {
			observedMAC, normalizeErr := normalizeSMSv2MAC(network.MACAddress)
			device := strings.TrimSpace(network.Name)
			if normalizeErr == nil && observedMAC == mac && device != "" && (target.Device == "" || device == target.Device) {
				registrationMatches = append(registrationMatches, registrationMatch{hostname: hostname, device: device})
			}
		}
	}
	if len(registrationMatches) != 1 {
		return smsv2KVMRuntimeInterfaceObject{}, fmt.Errorf("expected one live KVM registration match; observed %d", len(registrationMatches))
	}
	if rawErrors, exists := interfaces["errors"]; exists && rawErrors != nil {
		entries, ok := rawErrors.([]interface{})
		if !ok || len(entries) != 0 {
			return smsv2KVMRuntimeInterfaceObject{}, fmt.Errorf("network_interface discovery contains partial errors")
		}
	}
	items, ok := interfaces["items"].([]interface{})
	if !ok {
		return smsv2KVMRuntimeInterfaceObject{}, fmt.Errorf("network_interface discovery has no items array")
	}
	var selected smsv2KVMRuntimeInterfaceObject
	matches := 0
	for _, raw := range items {
		item, itemOK := raw.(map[string]interface{})
		if !itemOK {
			return smsv2KVMRuntimeInterfaceObject{}, fmt.Errorf("network_interface discovery contains a malformed item")
		}
		name, itemNamespace := smsv2RuntimeInterfaceItemIdentity(item)
		if itemNamespace != target.Namespace || (target.Name != "" && name != target.Name) {
			continue
		}
		owner, _ := nestedMap(item, "owner_view")
		if len(owner) == 0 {
			owner, _ = nestedMap(item, "system_metadata", "owner_view")
		}
		if stringField(owner, "kind") != "securemesh_site_v2" || stringField(owner, "name") != site || stringField(owner, "namespace") != namespace || stringField(owner, "uid") != uid {
			if target.Name != "" {
				return smsv2KVMRuntimeInterfaceObject{}, fmt.Errorf("runtime interface ownership does not match the current Secure Mesh Site v2 configuration")
			}
			continue
		}
		if err := validateKVMRuntimeInterfaceName(name); err != nil {
			return smsv2KVMRuntimeInterfaceObject{}, err
		}
		spec, specOK := smsv2RuntimeInterfaceSpec(item)
		ethernet, ethernetOK := nestedMap(spec, "ethernet_interface")
		if !specOK || !ethernetOK {
			return smsv2KVMRuntimeInterfaceObject{}, fmt.Errorf("runtime interface has no ethernet_interface spec")
		}
		if stringField(ethernet, "node") != registrationMatches[0].hostname || stringField(ethernet, "device") != registrationMatches[0].device {
			if target.Name != "" {
				return smsv2KVMRuntimeInterfaceObject{}, fmt.Errorf("runtime interface hostname or device does not match the live registration")
			}
			continue
		}
		_, slo := ethernet["site_local_network"]
		_, sli := ethernet["site_local_inside_network"]
		if !sli || slo {
			if target.Name != "" {
				return smsv2KVMRuntimeInterfaceObject{}, fmt.Errorf("runtime interface is not exactly one site-local-inside interface")
			}
			continue
		}
		matches++
		selected = smsv2KVMRuntimeInterfaceObject{
			Name: name, Namespace: itemNamespace, Site: site, OwnerUID: uid, Hostname: registrationMatches[0].hostname,
			Device: registrationMatches[0].device, MAC: mac, ResourceVersion: stringField(item, "resource_version"), Spec: deepCopySMSv2Map(spec), Raw: deepCopySMSv2Map(item),
		}
	}
	if matches != 1 {
		if matches == 0 && target.Name != "" {
			return smsv2KVMRuntimeInterfaceObject{}, errSMSv2KVMRuntimeInterfaceNotFound
		}
		return smsv2KVMRuntimeInterfaceObject{}, fmt.Errorf("expected one exact owned KVM runtime interface; observed %d", matches)
	}
	return selected, nil
}

func normalizeSMSv2KVMRuntimeInterfaceObservation(observation client.SMSv2Observation) map[string]interface{} {
	item := deepCopySMSv2Map(map[string]interface{}(observation))
	if metadata, ok := nestedMap(item, "metadata"); ok {
		item["name"] = stringField(metadata, "name")
		item["namespace"] = stringField(metadata, "namespace")
	}
	if spec, ok := nestedMap(item, "spec"); ok {
		item["get_spec"] = spec
	}
	return item
}

func buildSMSv2KVMRuntimeInterfaceUpdate(item map[string]interface{}, cidr string, restoreDHCP bool) (map[string]interface{}, error) {
	name, namespace := smsv2RuntimeInterfaceItemIdentity(item)
	spec, ok := smsv2RuntimeInterfaceSpec(item)
	if name == "" || namespace == "" || !ok {
		return nil, fmt.Errorf("runtime interface response is missing metadata or spec")
	}
	resourceVersion := stringField(item, "resource_version")
	if resourceVersion == "" {
		return nil, fmt.Errorf("runtime interface response has no resource version")
	}
	updatedSpec := deepCopySMSv2Map(spec)
	ethernet, ok := nestedMap(updatedSpec, "ethernet_interface")
	if !ok {
		return nil, fmt.Errorf("runtime interface response has no ethernet_interface spec")
	}
	if restoreDHCP {
		delete(ethernet, "static_ip")
		ethernet["dhcp_client"] = map[string]interface{}{}
	} else {
		canonical, err := canonicalKVMRuntimeInterfaceCIDR(cidr)
		if err != nil {
			return nil, err
		}
		delete(ethernet, "dhcp_client")
		ethernet["static_ip"] = map[string]interface{}{"node_static_ip": map[string]interface{}{"ip_address": canonical}}
	}
	metadata := map[string]interface{}{}
	if existing, ok := nestedMap(item, "metadata"); ok {
		metadata = deepCopySMSv2Map(existing)
	}
	metadata["name"], metadata["namespace"] = name, namespace
	return map[string]interface{}{"metadata": metadata, "spec": updatedSpec, "resource_version": resourceVersion}, nil
}

func currentSMSv2KVMRuntimeInterfaceCIDR(object smsv2KVMRuntimeInterfaceObject) (string, bool, error) {
	ethernet, ok := nestedMap(object.Spec, "ethernet_interface")
	if !ok {
		return "", false, fmt.Errorf("runtime interface response has no ethernet_interface spec")
	}
	if staticIP, exists := ethernet["static_ip"]; exists {
		staticMap, mapOK := staticIP.(map[string]interface{})
		nodeStatic, nodeOK := nestedMap(staticMap, "node_static_ip")
		cidr := stringField(nodeStatic, "ip_address")
		if !mapOK || !nodeOK || cidr == "" {
			return "", false, fmt.Errorf("runtime interface has malformed static IPv4 configuration")
		}
		canonical, err := canonicalKVMRuntimeInterfaceCIDR(cidr)
		return canonical, true, err
	}
	if _, exists := ethernet["dhcp_client"]; exists {
		return "", false, nil
	}
	return "", false, fmt.Errorf("runtime interface has neither DHCP nor static IPv4 configuration")
}

func isSMSv2KVMRuntimeInterfaceNotFound(err error) bool {
	if errors.Is(err, errSMSv2KVMRuntimeInterfaceNotFound) {
		return true
	}
	var apiErr *xcsherrors.XCSHError
	return errors.As(err, &apiErr) && apiErr.IsNotFound()
}

func isAmbiguousSMSv2KVMRuntimeInterfacePUT(err error) bool {
	var apiErr *xcsherrors.XCSHError
	if !errors.As(err, &apiErr) {
		return false
	}
	return apiErr.Code == xcsherrors.ErrCodeNetworkError || apiErr.Code == xcsherrors.ErrCodeTimeout
}

func (r *Smsv2KVMRuntimeInterfaceResource) observe(ctx context.Context, target smsv2KVMRuntimeInterfaceTarget) (smsv2KVMRuntimeInterfaceObject, error) {
	configuration, err := r.client.GetSMSv2Configuration(ctx, target.Namespace, target.Site)
	if err != nil {
		return smsv2KVMRuntimeInterfaceObject{}, err
	}
	registrations, err := r.client.ListRegistrationsBySite(ctx, target.Namespace, target.Site)
	if err != nil {
		return smsv2KVMRuntimeInterfaceObject{}, err
	}
	interfaces, err := r.client.ListSMSv2NetworkInterfaces(ctx, target.Namespace)
	if err != nil {
		return smsv2KVMRuntimeInterfaceObject{}, err
	}
	selected, err := selectSMSv2KVMRuntimeInterface(configuration, registrations, interfaces, target)
	if err != nil {
		return smsv2KVMRuntimeInterfaceObject{}, err
	}
	var exact client.SMSv2Observation
	exactPath := fmt.Sprintf("/api/config/namespaces/%s/network_interfaces/%s", url.PathEscape(target.Namespace), url.PathEscape(selected.Name))
	if err = r.client.Get(ctx, exactPath, &exact); err != nil {
		return smsv2KVMRuntimeInterfaceObject{}, err
	}
	exactItem := normalizeSMSv2KVMRuntimeInterfaceObservation(exact)
	exactTarget := target
	exactTarget.Name, exactTarget.Hostname, exactTarget.Device, exactTarget.OwnerUID = selected.Name, selected.Hostname, selected.Device, selected.OwnerUID
	return selectSMSv2KVMRuntimeInterface(configuration, registrations, client.SMSv2Observation{"items": []interface{}{exactItem}}, exactTarget)
}

func (r *Smsv2KVMRuntimeInterfaceResource) putAndReconcile(ctx context.Context, target smsv2KVMRuntimeInterfaceTarget, object smsv2KVMRuntimeInterfaceObject, cidr string, restoreDHCP bool) (smsv2KVMRuntimeInterfaceObject, error) {
	request, err := buildSMSv2KVMRuntimeInterfaceUpdate(object.Raw, cidr, restoreDHCP)
	if err != nil {
		return smsv2KVMRuntimeInterfaceObject{}, err
	}
	target.Name, target.Hostname, target.Device, target.OwnerUID = object.Name, object.Hostname, object.Device, object.OwnerUID
	path := fmt.Sprintf("/api/config/namespaces/%s/network_interfaces/%s", url.PathEscape(object.Namespace), url.PathEscape(object.Name))
	var response client.SMSv2Observation
	putErr := r.client.PutOnce(ctx, path, request, &response)
	if putErr != nil && !isAmbiguousSMSv2KVMRuntimeInterfacePUT(putErr) {
		return smsv2KVMRuntimeInterfaceObject{}, putErr
	}
	observed, readErr := r.observe(ctx, target)
	if readErr != nil {
		if putErr != nil {
			return smsv2KVMRuntimeInterfaceObject{}, fmt.Errorf("runtime interface PUT outcome could not be reconciled by exact-name read: %w", putErr)
		}
		return smsv2KVMRuntimeInterfaceObject{}, readErr
	}
	actual, static, stateErr := currentSMSv2KVMRuntimeInterfaceCIDR(observed)
	if stateErr != nil {
		return smsv2KVMRuntimeInterfaceObject{}, stateErr
	}
	if restoreDHCP {
		if static {
			return smsv2KVMRuntimeInterfaceObject{}, fmt.Errorf("runtime interface exact-name read did not confirm DHCP restoration")
		}
	} else {
		desired, desiredErr := canonicalKVMRuntimeInterfaceCIDR(cidr)
		if desiredErr != nil {
			return smsv2KVMRuntimeInterfaceObject{}, desiredErr
		}
		if !static || actual != desired {
			return smsv2KVMRuntimeInterfaceObject{}, fmt.Errorf("runtime interface exact-name read did not confirm requested static IPv4 state")
		}
	}
	if putErr != nil {
		tflog.Warn(ctx, "KVM runtime interface PUT returned a transport error but exact-name reconciliation proved the requested state", map[string]interface{}{"interface_name": target.Name, "namespace": target.Namespace})
	}
	return observed, nil
}

func updateSMSv2KVMRuntimeInterfaceModel(data *Smsv2KVMRuntimeInterfaceResourceModel, object smsv2KVMRuntimeInterfaceObject) error {
	cidr, static, err := currentSMSv2KVMRuntimeInterfaceCIDR(object)
	if err != nil {
		return err
	}
	data.ID = types.StringValue(object.Namespace + "/" + object.Name)
	data.Namespace = types.StringValue(object.Namespace)
	data.Site = types.StringValue(object.Site)
	data.InterfaceName = types.StringValue(object.Name)
	data.Hostname = types.StringValue(object.Hostname)
	data.Device = types.StringValue(object.Device)
	data.OwnerUID = types.StringValue(object.OwnerUID)
	data.ResourceVersion = stringOrNull(object.ResourceVersion)
	data.Configured = types.BoolValue(static)
	if static {
		data.IPv4CIDR = types.StringValue(cidr)
	} else {
		data.IPv4CIDR = types.StringNull()
	}
	return nil
}

func (r *Smsv2KVMRuntimeInterfaceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data Smsv2KVMRuntimeInterfaceResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	target := runtimeInterfaceTarget(data)
	object, err := r.observe(ctx, target)
	if err == nil {
		actual, static, stateErr := currentSMSv2KVMRuntimeInterfaceCIDR(object)
		desired, desiredErr := canonicalKVMRuntimeInterfaceCIDR(data.IPv4CIDR.ValueString())
		if stateErr != nil {
			err = stateErr
		} else if desiredErr != nil {
			err = desiredErr
		} else if !static || actual != desired {
			object, err = r.putAndReconcile(ctx, target, object, desired, false)
		}
	}
	if err != nil {
		resp.Diagnostics.AddError("KVM Runtime Interface Adoption Failed", err.Error())
		return
	}
	if err = updateSMSv2KVMRuntimeInterfaceModel(&data, object); err != nil {
		resp.Diagnostics.AddError("KVM Runtime Interface State Failed", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Smsv2KVMRuntimeInterfaceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data Smsv2KVMRuntimeInterfaceResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	object, err := r.observe(ctx, runtimeInterfaceTarget(data))
	if err != nil {
		if isSMSv2KVMRuntimeInterfaceNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("KVM Runtime Interface Read Failed", err.Error())
		return
	}
	if err = updateSMSv2KVMRuntimeInterfaceModel(&data, object); err != nil {
		resp.Diagnostics.AddError("KVM Runtime Interface State Failed", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Smsv2KVMRuntimeInterfaceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data Smsv2KVMRuntimeInterfaceResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	target := runtimeInterfaceTarget(data)
	object, err := r.observe(ctx, target)
	if err == nil {
		object, err = r.putAndReconcile(ctx, target, object, data.IPv4CIDR.ValueString(), false)
	}
	if err != nil {
		resp.Diagnostics.AddError("KVM Runtime Interface Update Failed", err.Error())
		return
	}
	if err = updateSMSv2KVMRuntimeInterfaceModel(&data, object); err != nil {
		resp.Diagnostics.AddError("KVM Runtime Interface State Failed", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Smsv2KVMRuntimeInterfaceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data Smsv2KVMRuntimeInterfaceResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	target := runtimeInterfaceTarget(data)
	object, err := r.observe(ctx, target)
	if err != nil {
		if isSMSv2KVMRuntimeInterfaceNotFound(err) {
			return
		}
		resp.Diagnostics.AddError("KVM Runtime Interface Destroy Failed", err.Error())
		return
	}
	actual, static, err := currentSMSv2KVMRuntimeInterfaceCIDR(object)
	if err != nil {
		resp.Diagnostics.AddError("KVM Runtime Interface Destroy Failed", err.Error())
		return
	}
	if !static {
		return
	}
	desired, err := canonicalKVMRuntimeInterfaceCIDR(data.IPv4CIDR.ValueString())
	if err != nil || actual != desired {
		resp.Diagnostics.AddError("KVM Runtime Interface Destroy Refused", "The exact owned interface no longer matches the Terraform-managed static IPv4 state; DHCP was not restored.")
		return
	}
	if _, err = r.putAndReconcile(ctx, target, object, "", true); err != nil {
		resp.Diagnostics.AddError("KVM Runtime Interface Destroy Failed", err.Error())
	}
}
