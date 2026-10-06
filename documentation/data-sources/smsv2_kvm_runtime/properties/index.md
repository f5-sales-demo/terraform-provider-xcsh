---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_smsv2_kvm_runtime."
xcsh_docs: {"aliases": ["smsv2 kvm runtime"], "body_bytes": 4694, "body_sha256": "sha256:b156a9806b6b8e08e46a7ed0264ea771d2eca9326874b7dfecb5de88e3cb4298", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:smsv2_kvm_runtime:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:smsv2_kvm_runtime:reference", "parent_id": "xcsh-docs:data-sources:smsv2_kvm_runtime:fundamentals", "path": "documentation/data-sources/smsv2_kvm_runtime/properties/index.md", "product": "distributed-cloud", "provider_name": "smsv2_kvm_runtime", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-3001120101110302-3200332330032022-1211131122311113-3223200200201030-0201120302233010-3313330010210020-0022220332221202-3020123013210220", "registry_path": "docs/guides/data-sources--smsv2_kvm_runtime--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["device"], "anchor": "schema-device", "description": "Guest device reported by the matching live KVM registration.", "document_id": "xcsh-docs:data-sources:smsv2_kvm_runtime:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["device"], "syntax": "attribute", "type": "string"}, {"aliases": ["expected mac"], "anchor": "schema-expected_mac", "description": "Terraform-owned KVM CE MAC used to select one registration device.", "document_id": "xcsh-docs:data-sources:smsv2_kvm_runtime:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["expected_mac"], "syntax": "attribute", "type": "string"}, {"aliases": ["hostname"], "anchor": "schema-hostname", "description": "Hostname reported by the matching live KVM registration.", "document_id": "xcsh-docs:data-sources:smsv2_kvm_runtime:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["hostname"], "syntax": "attribute", "type": "string"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Stable lookup identity composed from the site and normalized expected MAC.", "document_id": "xcsh-docs:data-sources:smsv2_kvm_runtime:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["interface name"], "anchor": "schema-interface_name", "description": "Exact owned XC network_interface object name.", "document_id": "xcsh-docs:data-sources:smsv2_kvm_runtime:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["interface_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["mac"], "anchor": "schema-mac", "description": "Normalized six-octet MAC address.", "document_id": "xcsh-docs:data-sources:smsv2_kvm_runtime:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["mac"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "Namespace containing the site and realized interface. Defaults to `system`.", "document_id": "xcsh-docs:data-sources:smsv2_kvm_runtime:reference", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["online"], "anchor": "schema-online", "description": "Whether the matching registration currently reports ONLINE.", "document_id": "xcsh-docs:data-sources:smsv2_kvm_runtime:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["online"], "syntax": "attribute", "type": "bool"}, {"aliases": ["poll interval seconds"], "anchor": "schema-poll_interval_seconds", "description": "Polling interval. Defaults to 10 seconds.", "document_id": "xcsh-docs:data-sources:smsv2_kvm_runtime:reference", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["poll_interval_seconds"], "syntax": "attribute", "type": "number"}, {"aliases": ["registration state"], "anchor": "schema-registration_state", "description": "Current state of the matching registration.", "document_id": "xcsh-docs:data-sources:smsv2_kvm_runtime:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["registration_state"], "syntax": "attribute", "type": "string"}, {"aliases": ["site"], "anchor": "schema-site", "description": "Secure Mesh Site v2 configuration name.", "document_id": "xcsh-docs:data-sources:smsv2_kvm_runtime:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["site"], "syntax": "attribute", "type": "string"}, {"aliases": ["duration", "timeout seconds"], "anchor": "schema-timeout_seconds", "description": "Bounded runtime discovery timeout. Defaults to 7200 seconds.", "document_id": "xcsh-docs:data-sources:smsv2_kvm_runtime:reference", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["timeout_seconds"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/smsv2_kvm_runtime/properties/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Property reference for xcsh_smsv2_kvm_runtime.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_smsv2_kvm_runtime](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_kvm_runtime/)
- Property reference

## Direct properties

<a id="schema-device"></a>

### device property

Type: `"string"`. Computed.

Guest device reported by the matching live KVM registration.

<a id="schema-expected_mac"></a>

### expected_mac property

Type: `"string"`. Required.

Terraform-owned KVM CE MAC used to select one registration device.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{stringvalidator.LengthAtLeast(1)}
```

<a id="schema-hostname"></a>

### hostname property

Type: `"string"`. Computed.

Hostname reported by the matching live KVM registration.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Stable lookup identity composed from the site and normalized expected MAC.

<a id="schema-interface_name"></a>

### interface_name property

Type: `"string"`. Computed.

Exact owned XC network\_interface object name.

<a id="schema-mac"></a>

### mac property

Type: `"string"`. Computed.

Normalized six-octet MAC address.

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Optional, Computed.

Namespace containing the site and realized interface. Defaults to \`system\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{stringvalidator.OneOf("system")}
```

<a id="schema-online"></a>

### online property

Type: `"bool"`. Computed.

Whether the matching registration currently reports ONLINE.

<a id="schema-poll_interval_seconds"></a>

### poll_interval_seconds property

Type: `"number"`. Optional, Computed.

Polling interval. Defaults to 10 seconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{int64validator.Between(1, 60)}
```

<a id="schema-registration_state"></a>

### registration_state property

Type: `"string"`. Computed.

Current state of the matching registration.

<a id="schema-site"></a>

### site property

Type: `"string"`. Required.

Secure Mesh Site v2 configuration name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{stringvalidator.LengthAtLeast(1)}
```

<a id="schema-timeout_seconds"></a>

### timeout_seconds property

Type: `"number"`. Optional, Computed.

Bounded runtime discovery timeout. Defaults to 7200 seconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{int64validator.Between(1, 7200)}
```

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `device` | [device](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_kvm_runtime/properties/#schema-device) |
| `expected_mac` | [expected_mac](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_kvm_runtime/properties/#schema-expected_mac) |
| `hostname` | [hostname](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_kvm_runtime/properties/#schema-hostname) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_kvm_runtime/properties/#schema-id) |
| `interface_name` | [interface_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_kvm_runtime/properties/#schema-interface_name) |
| `mac` | [mac](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_kvm_runtime/properties/#schema-mac) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_kvm_runtime/properties/#schema-namespace) |
| `online` | [online](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_kvm_runtime/properties/#schema-online) |
| `poll_interval_seconds` | [poll_interval_seconds](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_kvm_runtime/properties/#schema-poll_interval_seconds) |
| `registration_state` | [registration_state](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_kvm_runtime/properties/#schema-registration_state) |
| `site` | [site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_kvm_runtime/properties/#schema-site) |
| `timeout_seconds` | [timeout_seconds](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_kvm_runtime/properties/#schema-timeout_seconds) |
