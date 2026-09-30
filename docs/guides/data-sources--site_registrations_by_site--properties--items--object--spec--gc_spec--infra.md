---
page_title: "items.object.spec.gc_spec.infra"
subcategory: ""
description: "items.object.spec.gc_spec.infra for xcsh_site_registrations_by_site."
xcsh_docs: {"aliases": [], "body_bytes": 6686, "body_sha256": "sha256:6a4005a95a081808d943c2459231d81b970cd1d2d84b19cb81adebd6693b0077", "canonical_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra", "child_ids": ["xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:bond_config", "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:hugepages", "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:hw_info", "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:interfaces", "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:internet_proxy", "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:sw_info"], "collection_id": "xcsh-docs:data-sources:site_registrations_by_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra", "parent_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec", "path": "docs/guides/data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec--infra.md", "provider_name": "site_registrations_by_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["items", "object", "spec", "gc_spec", "infra"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/infra/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "items.object.spec.gc_spec.infra for xcsh_site_registrations_by_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# items.object.spec.gc_spec.infra

Breadcrumbs:

- [xcsh_site_registrations_by_site](../data-sources/site_registrations_by_site.md)
- [Property reference](data-sources--site_registrations_by_site--reference.md)
- [items](data-sources--site_registrations_by_site--properties--items.md)
- [items.object](data-sources--site_registrations_by_site--properties--items--object.md)
- [items.object.spec](data-sources--site_registrations_by_site--properties--items--object--spec.md)
- [items.object.spec.gc_spec](data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec.md)
- items.object.spec.gc_spec.infra

<a id="section"></a>

Type: `"single"`. Computed.

InfraMetadata stores information about instance infrastructure.

## Direct properties

<a id="schema-items--object--spec--gc_spec--infra--availability_zone"></a>

### availability_zone property

Type: `"string"`. Computed.

Availability Zone is a high-availability offering that protects your applications and data from
datacenter failures.

- [bond_config](data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec--infra--bond_config.md): complete subsection reference.

<a id="schema-items--object--spec--gc_spec--infra--certified_hw"></a>

### certified_hw property

Type: `"string"`. Computed.

Certified HW name used to map with F5XC certified\_hardware definition.

<a id="schema-items--object--spec--gc_spec--infra--domain"></a>

### domain property

Type: `"string"`. Computed.

Machine domain. It's used for Kubernetes cloud provider when domain must be different than F5
Distributed Cloud.

<a id="schema-items--object--spec--gc_spec--infra--hostname"></a>

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

- [hugepages](data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec--infra--hugepages.md): complete subsection reference.

- [hw_info](data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec--infra--hw_info.md): complete subsection reference.

<a id="schema-items--object--spec--gc_spec--infra--instance_id"></a>

### instance_id property

Type: `"string"`. Computed.

Instance ID (assigned by infrastructure provider).

- [interfaces](data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec--infra--interfaces.md): complete subsection reference.

- [internet_proxy](data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec--infra--internet_proxy.md): complete subsection reference.

<a id="schema-items--object--spec--gc_spec--infra--is_slo_static"></a>

### is_slo_static property

Type: `"bool"`. Computed.

Is SLO Static. Indicates whether the SLO is static.

<a id="schema-items--object--spec--gc_spec--infra--machine_id"></a>

### machine_id property

Type: `"string"`. Computed.

Machine ID - generated by operating system.

<a id="schema-items--object--spec--gc_spec--infra--provider_ref"></a>

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

- [sw_info](data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec--infra--sw_info.md): complete subsection reference.

<a id="schema-items--object--spec--gc_spec--infra--timestamp"></a>

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

<a id="schema-items--object--spec--gc_spec--infra--zone"></a>

### zone property

Type: `"string"`. Computed.

Instance zone (or region), depends on provider.

## Next pages

- [items.object.spec.gc_spec.infra.bond_config](data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec--infra--bond_config.md)
- [items.object.spec.gc_spec.infra.hugepages](data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec--infra--hugepages.md)
- [items.object.spec.gc_spec.infra.hw_info](data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec--infra--hw_info.md)
- [items.object.spec.gc_spec.infra.interfaces](data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec--infra--interfaces.md)
- [items.object.spec.gc_spec.infra.internet_proxy](data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec--infra--internet_proxy.md)
- [items.object.spec.gc_spec.infra.sw_info](data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec--infra--sw_info.md)
- [items.object.spec.gc_spec](data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec.md)
- [xcsh_site_registrations_by_site](../data-sources/site_registrations_by_site.md)
