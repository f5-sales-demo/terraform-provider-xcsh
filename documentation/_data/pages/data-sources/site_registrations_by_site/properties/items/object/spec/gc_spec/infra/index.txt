---
page_title: "items.object.spec.gc_spec.infra"
subcategory: ""
description: "InfraMetadata stores information about instance infrastructure."
xcsh_docs: {"aliases": ["items object spec gc spec infra"], "body_bytes": 7738, "body_sha256": "sha256:d2963cfb58ac886a8570e1d8853bd4780677be80580eacf1016c8f5b27f8fe60", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:bond_config", "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:hugepages", "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:hw_info", "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:interfaces", "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:internet_proxy", "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:sw_info"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations_by_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra", "parent_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec", "path": "documentation/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/infra/index.md", "product": "distributed-cloud", "provider_name": "site_registrations_by_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0332223233032330-1222012203221213-1222022200103221-3001233232000203-2023021210320110-3323010003030202-0211323013123320-3111121211311331", "registry_path": "docs/guides/data-sources--site_registrations_by_site--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "object", "spec", "gc_spec", "infra"], "schema_version": 1, "sections": [{"aliases": ["availability zone"], "anchor": "schema-items--object--spec--gc_spec--infra--availability_zone", "description": "Availability Zone is a high-availability offering that protects your applications and data from datacenter failures.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "availability_zone"], "syntax": "attribute", "type": "string"}, {"aliases": ["bond config"], "anchor": "section", "description": "Bond device configuration for VPM registration.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:bond_config", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "bond_config"], "syntax": "attribute", "type": "object"}, {"aliases": ["certified hw"], "anchor": "schema-items--object--spec--gc_spec--infra--certified_hw", "description": "Certified HW name used to map with F5XC certified_hardware definition.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "certified_hw"], "syntax": "attribute", "type": "string"}, {"aliases": ["domain"], "anchor": "schema-items--object--spec--gc_spec--infra--domain", "description": "Machine domain. It's used for Kubernetes cloud provider when domain must be different than F5 Distributed Cloud.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "domain"], "syntax": "attribute", "type": "string"}, {"aliases": ["hostname"], "anchor": "schema-items--object--spec--gc_spec--infra--hostname", "description": "Must be unique in entire cluster and same as OS settings. '.' (dots) are not allowed in hostname.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hostname"], "syntax": "attribute", "type": "string"}, {"aliases": ["hugepages"], "anchor": "section", "description": "Hugepage settings for CE on K8s SMV2 site.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:hugepages", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hugepages"], "syntax": "attribute", "type": "object"}, {"aliases": ["hw info"], "anchor": "section", "description": "OsInfo holds information about host OS and HW.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:hw_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info"], "syntax": "attribute", "type": "object"}, {"aliases": ["instance id"], "anchor": "schema-items--object--spec--gc_spec--infra--instance_id", "description": "Instance ID (assigned by infrastructure provider).", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "instance_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["interfaces"], "anchor": "section", "description": "Machine interfaces present during registration time.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:interfaces", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "interfaces"], "syntax": "attribute", "type": "object"}, {"aliases": ["internet proxy"], "anchor": "section", "description": "Proxy describes OPTIONS for HTTP or HTTPS proxy configurations.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:internet_proxy", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "internet_proxy"], "syntax": "attribute", "type": "object"}, {"aliases": ["is slo static"], "anchor": "schema-items--object--spec--gc_spec--infra--is_slo_static", "description": "Is SLO Static. Indicates whether the SLO is static.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "is_slo_static"], "syntax": "attribute", "type": "bool"}, {"aliases": ["machine id"], "anchor": "schema-items--object--spec--gc_spec--infra--machine_id", "description": "Machine ID - generated by operating system.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "machine_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["provider ref"], "anchor": "schema-items--object--spec--gc_spec--infra--provider_ref", "description": "Infrastructure provider enum for registration. It describes where is instance running. Provider was not detected AWS cloud instance Google cloud instance Azure cloud instance VMWare VM KVM VM Other provider, which was not identified by system. Possible values are `UNKNOWN`, `AWS`, `GOOGLE`, `AZURE`, `VMWARE`, `KVM`,", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "provider_ref"], "syntax": "attribute", "type": "string"}, {"aliases": ["sw info"], "anchor": "section", "description": "SWInfo holds information about sw version.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:sw_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "sw_info"], "syntax": "attribute", "type": "object"}, {"aliases": ["timestamp"], "anchor": "schema-items--object--spec--gc_spec--infra--timestamp", "description": "It's used to verify machine have acceptable time difference from server.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "timestamp"], "syntax": "attribute", "type": "string"}, {"aliases": ["zone"], "anchor": "schema-items--object--spec--gc_spec--infra--zone", "description": "Instance zone (or region), depends on provider.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "zone"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/infra/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "InfraMetadata stores information about instance infrastructure.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.object.spec.gc_spec.infra

Breadcrumbs:

- [xcsh_site_registrations_by_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/)
- [items.object](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/)
- [items.object.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/spec/)
- [items.object.spec.gc_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/)
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

- [bond_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/infra/bond_config/): complete subsection reference.

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

- [hugepages](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/infra/hugepages/): complete subsection reference.

- [hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/infra/hw_info/): complete subsection reference.

<a id="schema-items--object--spec--gc_spec--infra--instance_id"></a>

### instance_id property

Type: `"string"`. Computed.

Instance ID (assigned by infrastructure provider).

- [interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/infra/interfaces/): complete subsection reference.

- [internet_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/infra/internet_proxy/): complete subsection reference.

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

- [sw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/infra/sw_info/): complete subsection reference.

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

- [items.object.spec.gc_spec.infra.bond_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/infra/bond_config/)
- [items.object.spec.gc_spec.infra.hugepages](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/infra/hugepages/)
- [items.object.spec.gc_spec.infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/infra/hw_info/)
- [items.object.spec.gc_spec.infra.interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/infra/interfaces/)
- [items.object.spec.gc_spec.infra.internet_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/infra/internet_proxy/)
- [items.object.spec.gc_spec.infra.sw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/infra/sw_info/)
- [items.object.spec.gc_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/)
- [xcsh_site_registrations_by_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/)
