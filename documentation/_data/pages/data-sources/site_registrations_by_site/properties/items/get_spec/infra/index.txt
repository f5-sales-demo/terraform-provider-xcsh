---
page_title: "items.get_spec.infra"
subcategory: ""
description: "items.get_spec.infra for xcsh_site_registrations_by_site."
xcsh_docs: {"aliases": [], "body_bytes": 6957, "body_sha256": "sha256:7de3be729212f027ef4127af8c13242e9f4de90c4528467585931b4a30b4b493", "child_ids": ["xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:bond_config", "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hugepages", "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hw_info", "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:interfaces", "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:internet_proxy", "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:sw_info"], "collection_id": "xcsh-docs:data-sources:site_registrations_by_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra", "parent_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec", "path": "documentation/data-sources/site_registrations_by_site/properties/items/get_spec/infra/index.md", "provider_name": "site_registrations_by_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["items", "get_spec", "infra"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_site/properties/items/get_spec/infra/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "items.get_spec.infra for xcsh_site_registrations_by_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# items.get_spec.infra

Breadcrumbs:

- [xcsh_site_registrations_by_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/)
- [items.get_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/get_spec/)
- items.get_spec.infra

<a id="section"></a>

Type: `"single"`. Computed.

InfraMetadata stores information about instance infrastructure.

## Direct properties

<a id="schema-items--get_spec--infra--availability_zone"></a>

### availability_zone property

Type: `"string"`. Computed.

Availability Zone is a high-availability offering that protects your applications and data from
datacenter failures.

- [bond_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/get_spec/infra/bond_config/): complete subsection reference.

<a id="schema-items--get_spec--infra--certified_hw"></a>

### certified_hw property

Type: `"string"`. Computed.

Certified HW name used to map with F5XC certified\_hardware definition.

<a id="schema-items--get_spec--infra--domain"></a>

### domain property

Type: `"string"`. Computed.

Machine domain. It's used for Kubernetes cloud provider when domain must be different than F5
Distributed Cloud.

<a id="schema-items--get_spec--infra--hostname"></a>

### hostname property

Type: `"string"`. Computed.

Must be unique in entire cluster and same as OS settings. '.' (dots) are not allowed in hostname.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$`),
    ""),
}
```

- [hugepages](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/get_spec/infra/hugepages/): complete subsection reference.

- [hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/get_spec/infra/hw_info/): complete subsection reference.

<a id="schema-items--get_spec--infra--instance_id"></a>

### instance_id property

Type: `"string"`. Computed.

Instance ID (assigned by infrastructure provider).

- [interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/get_spec/infra/interfaces/): complete subsection reference.

- [internet_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/get_spec/infra/internet_proxy/): complete subsection reference.

<a id="schema-items--get_spec--infra--is_slo_static"></a>

### is_slo_static property

Type: `"bool"`. Computed.

Is SLO Static. Indicates whether the SLO is static.

<a id="schema-items--get_spec--infra--machine_id"></a>

### machine_id property

Type: `"string"`. Computed.

Machine ID - generated by operating system.

<a id="schema-items--get_spec--infra--provider_ref"></a>

### provider_ref property

Type: `"string"`. Computed.

\[Enum:
UNKNOWN|AWS|GOOGLE|AZURE|VMWARE|KVM|OTHER|VOLTERRA|IBMCLOUD|UNKNOWN\_K8S|AWS\_K8S|GCP\_K8S|AZURE\_K8S|VMWARE\_K8S|KVM\_K8S|OTHER\_K8S|VOLTERRA\_K8S|IBMCLOUD\_K8S|F5OS|RSERIES|OCI|NUTANIX|OPENSTACK|EQUINIX|OPENSHIFT\_VIRTUALIZATION|KUBERNETES\]
Infrastructure provider enum for registration. It describes where is instance running. Provider was
not detected AWS cloud instance Google cloud instance Azure cloud instance VMWare VM KVM VM Other
provider, which was not identified by system. Possible values are \`UNKNOWN\`, \`AWS\`, \`GOOGLE\`,
\`AZURE\`, \`VMWARE\`, \`KVM\`, \`OTHER\`, \`VOLTERRA\`, \`IBMCLOUD\`, \`UNKNOWN\_K8S\`,
\`AWS\_K8S\`, \`GCP\_K8S\`, \`AZURE\_K8S\`, \`VMWARE\_K8S\`, \`KVM\_K8S\`, \`OTHER\_K8S\`,
\`VOLTERRA\_K8S\`, \`IBMCLOUD\_K8S\`, \`F5OS\`, \`RSERIES\`, \`OCI\`, \`NUTANIX\`, \`OPENSTACK\`,
\`EQUINIX\`, \`OPENSHIFT\_VIRTUALIZATION\`, \`KUBERNETES\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("UNKNOWN",
    "AWS",
    "GOOGLE",
    "AZURE",
    "VMWARE",
    "KVM",
    "OTHER",
    "VOLTERRA",
    "IBMCLOUD",
    "UNKNOWN_K8S",
    "AWS_K8S",
    "GCP_K8S",
    "AZURE_K8S",
    "VMWARE_K8S",
    "KVM_K8S",
    "OTHER_K8S",
    "VOLTERRA_K8S",
    "IBMCLOUD_K8S",
    "F5OS",
    "RSERIES",
    "OCI",
    "NUTANIX",
    "OPENSTACK",
    "EQUINIX",
    "OPENSHIFT_VIRTUALIZATION",
    "KUBERNETES"),
}
```

- [sw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/get_spec/infra/sw_info/): complete subsection reference.

<a id="schema-items--get_spec--infra--timestamp"></a>

### timestamp property

Type: `"string"`. Computed.

It's used to verify machine have acceptable time difference from server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(20, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d{3})?Z?$`),
    ""),
}
```

<a id="schema-items--get_spec--infra--zone"></a>

### zone property

Type: `"string"`. Computed.

Instance zone (or region), depends on provider.

## Next pages

- [items.get_spec.infra.bond_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/get_spec/infra/bond_config/)
- [items.get_spec.infra.hugepages](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/get_spec/infra/hugepages/)
- [items.get_spec.infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/get_spec/infra/hw_info/)
- [items.get_spec.infra.interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/get_spec/infra/interfaces/)
- [items.get_spec.infra.internet_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/get_spec/infra/internet_proxy/)
- [items.get_spec.infra.sw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/get_spec/infra/sw_info/)
- [items.get_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/get_spec/)
- [xcsh_site_registrations_by_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/)
