---
page_title: "items.object.spec.gc_spec.infra"
subcategory: ""
description: "InfraMetadata stores information about instance infrastructure."
xcsh_docs: {"aliases": ["items object spec gc spec infra"], "body_bytes": 6826, "body_sha256": "sha256:9f5013c46924fd6e520a5c53439d1186b2def3b41ad8ea785e1970fc4e401acf", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec:infra:bond_config", "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec:infra:hugepages", "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec:infra:hw_info", "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec:infra:interfaces", "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec:infra:internet_proxy", "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec:infra:sw_info"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations_by_state:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec:infra", "parent_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec", "path": "documentation/data-sources/site_registrations_by_state/properties/items/object/spec/gc_spec/infra/index.md", "product": "distributed-cloud", "provider_name": "site_registrations_by_state", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-2311032302313011-3133111000003211-0232321221321111-3110112201113112-3121220013111032-0001133110213113-3310011221132020-0131221232203131", "registry_path": "docs/guides/data-sources--site_registrations_by_state--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "object", "spec", "gc_spec", "infra"], "schema_version": 1, "sections": [{"aliases": ["items object spec gc spec infra availability zone"], "anchor": "schema-items--object--spec--gc_spec--infra--availability_zone", "description": "Availability Zone is a high-availability offering that protects your applications and data from datacenter failures.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec:infra", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "availability_zone"], "syntax": "attribute", "type": "string"}, {"aliases": ["items object spec gc spec infra bond config"], "anchor": "section", "description": "Bond device configuration for VPM registration.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec:infra:bond_config", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "bond_config"], "syntax": "attribute", "type": "object"}, {"aliases": ["items object spec gc spec infra certified hw"], "anchor": "schema-items--object--spec--gc_spec--infra--certified_hw", "description": "Certified HW name used to map with F5XC certified_hardware definition.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec:infra", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "certified_hw"], "syntax": "attribute", "type": "string"}, {"aliases": ["items object spec gc spec infra domain"], "anchor": "schema-items--object--spec--gc_spec--infra--domain", "description": "Machine domain. It's used for Kubernetes cloud provider when domain must be different than F5 Distributed Cloud.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec:infra", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "domain"], "syntax": "attribute", "type": "string"}, {"aliases": ["items object spec gc spec infra hostname"], "anchor": "schema-items--object--spec--gc_spec--infra--hostname", "description": "Must be unique in entire cluster and same as OS settings. '.' (dots) are not allowed in hostname.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec:infra", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hostname"], "syntax": "attribute", "type": "string"}, {"aliases": ["items object spec gc spec infra hugepages"], "anchor": "section", "description": "Hugepage settings for CE on K8s SMV2 site.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec:infra:hugepages", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hugepages"], "syntax": "attribute", "type": "object"}, {"aliases": ["items object spec gc spec infra hw info"], "anchor": "section", "description": "OsInfo holds information about host OS and HW.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec:infra:hw_info", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info"], "syntax": "attribute", "type": "object"}, {"aliases": ["items object spec gc spec infra instance id"], "anchor": "schema-items--object--spec--gc_spec--infra--instance_id", "description": "Instance ID (assigned by infrastructure provider).", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec:infra", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "instance_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["items object spec gc spec infra interfaces"], "anchor": "section", "description": "Machine interfaces present during registration time.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec:infra:interfaces", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "interfaces"], "syntax": "attribute", "type": "object"}, {"aliases": ["items object spec gc spec infra internet proxy"], "anchor": "section", "description": "Proxy describes OPTIONS for HTTP or HTTPS proxy configurations.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec:infra:internet_proxy", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "internet_proxy"], "syntax": "attribute", "type": "object"}, {"aliases": ["items object spec gc spec infra is slo static"], "anchor": "schema-items--object--spec--gc_spec--infra--is_slo_static", "description": "Is SLO Static. Indicates whether the SLO is static.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec:infra", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "is_slo_static"], "syntax": "attribute", "type": "bool"}, {"aliases": ["items object spec gc spec infra machine id"], "anchor": "schema-items--object--spec--gc_spec--infra--machine_id", "description": "Machine ID - generated by operating system.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec:infra", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "machine_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["items object spec gc spec infra provider ref"], "anchor": "schema-items--object--spec--gc_spec--infra--provider_ref", "description": "Infrastructure provider enum for registration. It describes where is instance running. Provider was not detected AWS cloud instance Google cloud instance Azure cloud instance VMWare VM KVM VM Other provider, which was not identified by system. Possible values are `UNKNOWN`, `AWS`, `GOOGLE`, `AZURE`, `VMWARE`, `KVM`,", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec:infra", "enum_extraction_complete": true, "enum_validators": [{"case_sensitive": true, "complete": true, "source": "ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf", "validator": "OneOf", "values": ["AWS", "AWS_K8S", "AZURE", "AZURE_K8S", "EQUINIX", "F5OS", "GCP_K8S", "GOOGLE", "IBMCLOUD", "IBMCLOUD_K8S", "KUBERNETES", "KVM", "KVM_K8S", "NUTANIX", "OCI", "OPENSHIFT_VIRTUALIZATION", "OPENSTACK", "OTHER", "OTHER_K8S", "RSERIES", "UNKNOWN", "UNKNOWN_K8S", "VMWARE", "VMWARE_K8S", "VOLTERRA", "VOLTERRA_K8S"], "version": 1}], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "provider_ref"], "syntax": "attribute", "type": "string"}, {"aliases": ["items object spec gc spec infra sw info"], "anchor": "section", "description": "SWInfo holds information about sw version.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec:infra:sw_info", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "sw_info"], "syntax": "attribute", "type": "object"}, {"aliases": ["items object spec gc spec infra timestamp"], "anchor": "schema-items--object--spec--gc_spec--infra--timestamp", "description": "It's used to verify machine have acceptable time difference from server.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec:infra", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "timestamp"], "syntax": "attribute", "type": "string"}, {"aliases": ["items object spec gc spec infra zone"], "anchor": "schema-items--object--spec--gc_spec--infra--zone", "description": "Instance zone (or region), depends on provider.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec:infra", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "zone"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_state/properties/items/object/spec/gc_spec/infra/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "InfraMetadata stores information about instance infrastructure.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": [], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.object.spec.gc_spec.infra

Breadcrumbs:

- [xcsh_site_registrations_by_state](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/)
- [items.object](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/)
- [items.object.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/spec/)
- [items.object.spec.gc_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/spec/gc_spec/)
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

- [bond_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/spec/gc_spec/infra/bond_config/): complete subsection reference.

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
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$`),
    ""),
}
```

- [hugepages](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/spec/gc_spec/infra/hugepages/): complete subsection reference.

- [hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/spec/gc_spec/infra/hw_info/): complete subsection reference.

<a id="schema-items--object--spec--gc_spec--infra--instance_id"></a>

### instance_id property

Type: `"string"`. Computed.

Instance ID (assigned by infrastructure provider).

- [interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/spec/gc_spec/infra/interfaces/): complete subsection reference.

- [internet_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/spec/gc_spec/infra/internet_proxy/): complete subsection reference.

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
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["AWS","AWS_K8S","AZURE","AZURE_K8S","EQUINIX","F5OS","GCP_K8S","GOOGLE","IBMCLOUD","IBMCLOUD_K8S","KUBERNETES","KVM","KVM_K8S","NUTANIX","OCI","OPENSHIFT_VIRTUALIZATION","OPENSTACK","OTHER","OTHER_K8S","RSERIES","UNKNOWN","UNKNOWN_K8S","VMWARE","VMWARE_K8S","VOLTERRA","VOLTERRA_K8S"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
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

- [sw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/spec/gc_spec/infra/sw_info/): complete subsection reference.

<a id="schema-items--object--spec--gc_spec--infra--timestamp"></a>

### timestamp property

Type: `"string"`. Computed.

It's used to verify machine have acceptable time difference from server.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
