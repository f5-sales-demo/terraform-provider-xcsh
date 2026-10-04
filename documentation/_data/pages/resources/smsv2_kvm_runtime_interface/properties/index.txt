---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_smsv2_kvm_runtime_interface."
xcsh_docs: {"aliases": ["smsv2 kvm runtime interface"], "body_bytes": 4467, "body_sha256": "sha256:f49651f9584fe28ee72121d384813e7ae77670af80d19277bd11c41f7107c776", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:smsv2_kvm_runtime_interface:collection", "completeness": "complete", "id": "xcsh-docs:resources:smsv2_kvm_runtime_interface:reference", "parent_id": "xcsh-docs:resources:smsv2_kvm_runtime_interface:fundamentals", "path": "documentation/resources/smsv2_kvm_runtime_interface/properties/index.md", "product": "distributed-cloud", "provider_name": "smsv2_kvm_runtime_interface", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3012122202201230-2332231032003301-1332213012132300-3230003001013213-3113211202002223-1111020221310310-3232001000230113-0323321321201001", "registry_path": "docs/guides/resources--smsv2_kvm_runtime_interface--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["configured"], "anchor": "schema-configured", "description": "Whether the exact owned SLI currently uses static IPv4 configuration.", "document_id": "xcsh-docs:resources:smsv2_kvm_runtime_interface:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["configured"], "syntax": "attribute", "type": "bool"}, {"aliases": ["device"], "anchor": "schema-device", "description": "Exact live registration device resolved from the expected MAC.", "document_id": "xcsh-docs:resources:smsv2_kvm_runtime_interface:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["device"], "syntax": "attribute", "type": "string"}, {"aliases": ["expected mac"], "anchor": "schema-expected_mac", "description": "Terraform-owned SLI MAC used for live registration correlation.", "document_id": "xcsh-docs:resources:smsv2_kvm_runtime_interface:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["expected_mac"], "syntax": "attribute", "type": "string"}, {"aliases": ["hostname"], "anchor": "schema-hostname", "description": "Exact live registration hostname resolved from the expected MAC.", "document_id": "xcsh-docs:resources:smsv2_kvm_runtime_interface:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["hostname"], "syntax": "attribute", "type": "string"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Stable namespace/interface identity.", "document_id": "xcsh-docs:resources:smsv2_kvm_runtime_interface:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["interface name"], "anchor": "schema-interface_name", "description": "Exact platform-generated SLI child name resolved from live ownership. Platform names may exceed 64 characters.", "document_id": "xcsh-docs:resources:smsv2_kvm_runtime_interface:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["interface_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["ipv4 cidr"], "anchor": "schema-ipv4_cidr", "description": "Static IPv4 host address and prefix to configure on the SLI.", "document_id": "xcsh-docs:resources:smsv2_kvm_runtime_interface:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ipv4_cidr"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "Namespace containing the site and runtime child. KVM SMSv2 supports only `system`.", "document_id": "xcsh-docs:resources:smsv2_kvm_runtime_interface:reference", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["owner uid"], "anchor": "schema-owner_uid", "description": "Secure Mesh Site v2 UID that owns the runtime child.", "document_id": "xcsh-docs:resources:smsv2_kvm_runtime_interface:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["owner_uid"], "syntax": "attribute", "type": "string"}, {"aliases": ["resource version"], "anchor": "schema-resource_version", "description": "Latest XC concurrency version observed after reconciliation.", "document_id": "xcsh-docs:resources:smsv2_kvm_runtime_interface:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["resource_version"], "syntax": "attribute", "type": "string"}, {"aliases": ["site"], "anchor": "schema-site", "description": "Secure Mesh Site v2 configuration name.", "document_id": "xcsh-docs:resources:smsv2_kvm_runtime_interface:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["site"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/smsv2_kvm_runtime_interface/properties/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Property reference for xcsh_smsv2_kvm_runtime_interface.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": [], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_smsv2_kvm_runtime_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/smsv2_kvm_runtime_interface/)
- Property reference

## Direct properties

<a id="schema-configured"></a>

### configured property

Type: `"bool"`. Computed.

Whether the exact owned SLI currently uses static IPv4 configuration.

<a id="schema-device"></a>

### device property

Type: `"string"`. Computed.

Exact live registration device resolved from the expected MAC.

<a id="schema-expected_mac"></a>

### expected_mac property

Type: `"string"`. Required.

Terraform-owned SLI MAC used for live registration correlation.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{stringvalidator.LengthAtLeast(1)}
```

<a id="schema-hostname"></a>

### hostname property

Type: `"string"`. Computed.

Exact live registration hostname resolved from the expected MAC.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Stable namespace/interface identity.

<a id="schema-interface_name"></a>

### interface_name property

Type: `"string"`. Computed.

Exact platform-generated SLI child name resolved from live ownership. Platform names may exceed 64
characters.

<a id="schema-ipv4_cidr"></a>

### ipv4_cidr property

Type: `"string"`. Required.

Static IPv4 host address and prefix to configure on the SLI.

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Optional, Computed.

Namespace containing the site and runtime child. KVM SMSv2 supports only \`system\`.

Provider validators and defaults (from schema source):

```go
Default: stringdefault.StaticString("system")
Validators: []validator.String{stringvalidator.OneOf("system")}
```

<a id="schema-owner_uid"></a>

### owner_uid property

Type: `"string"`. Computed.

Secure Mesh Site v2 UID that owns the runtime child.

<a id="schema-resource_version"></a>

### resource_version property

Type: `"string"`. Computed.

Latest XC concurrency version observed after reconciliation.

<a id="schema-site"></a>

### site property

Type: `"string"`. Required.

Secure Mesh Site v2 configuration name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{stringvalidator.LengthAtLeast(1)}
```

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `configured` | [configured](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/smsv2_kvm_runtime_interface/properties/#schema-configured) |
| `device` | [device](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/smsv2_kvm_runtime_interface/properties/#schema-device) |
| `expected_mac` | [expected_mac](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/smsv2_kvm_runtime_interface/properties/#schema-expected_mac) |
| `hostname` | [hostname](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/smsv2_kvm_runtime_interface/properties/#schema-hostname) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/smsv2_kvm_runtime_interface/properties/#schema-id) |
| `interface_name` | [interface_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/smsv2_kvm_runtime_interface/properties/#schema-interface_name) |
| `ipv4_cidr` | [ipv4_cidr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/smsv2_kvm_runtime_interface/properties/#schema-ipv4_cidr) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/smsv2_kvm_runtime_interface/properties/#schema-namespace) |
| `owner_uid` | [owner_uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/smsv2_kvm_runtime_interface/properties/#schema-owner_uid) |
| `resource_version` | [resource_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/smsv2_kvm_runtime_interface/properties/#schema-resource_version) |
| `site` | [site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/smsv2_kvm_runtime_interface/properties/#schema-site) |

## Next pages

- [xcsh_smsv2_kvm_runtime_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/smsv2_kvm_runtime_interface/)
