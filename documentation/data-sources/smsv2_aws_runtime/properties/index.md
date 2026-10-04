---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_smsv2_aws_runtime."
xcsh_docs: {"aliases": ["smsv2 aws runtime"], "body_bytes": 4990, "body_sha256": "sha256:fedcb1c89cc9dc7ed10a8dd0f4312935a723b6ae92af7da3d398ea432c796060", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:smsv2_aws_runtime:properties:interfaces", "xcsh-docs:data-sources:smsv2_aws_runtime:properties:nodes"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:smsv2_aws_runtime:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:smsv2_aws_runtime:reference", "parent_id": "xcsh-docs:data-sources:smsv2_aws_runtime:fundamentals", "path": "documentation/data-sources/smsv2_aws_runtime/properties/index.md", "product": "distributed-cloud", "provider_name": "smsv2_aws_runtime", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1131203033100032-2003330011020220-3022010022132201-2022013302210203-1333333332220120-3012021230132301-0322030310002231-3122323132101032", "registry_path": "docs/guides/data-sources--smsv2_aws_runtime--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["healthy"], "anchor": "schema-healthy", "description": "healthy", "document_id": "xcsh-docs:data-sources:smsv2_aws_runtime:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["healthy"], "syntax": "attribute", "type": "bool"}, {"aliases": ["id"], "anchor": "schema-id", "description": "id", "document_id": "xcsh-docs:data-sources:smsv2_aws_runtime:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["interfaces"], "anchor": "section", "description": "interfaces", "document_id": "xcsh-docs:data-sources:smsv2_aws_runtime:properties:interfaces", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "map", "relationships": [], "schema_path": ["interfaces"], "syntax": "attribute", "type": "object"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "namespace", "document_id": "xcsh-docs:data-sources:smsv2_aws_runtime:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["nodes"], "anchor": "section", "description": "nodes", "document_id": "xcsh-docs:data-sources:smsv2_aws_runtime:properties:nodes", "flags": ["required"], "max_items": null, "min_items": null, "nesting": "map", "relationships": [], "schema_path": ["nodes"], "syntax": "attribute", "type": "object"}, {"aliases": ["poll interval seconds"], "anchor": "schema-poll_interval_seconds", "description": "poll interval seconds", "document_id": "xcsh-docs:data-sources:smsv2_aws_runtime:reference", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["poll_interval_seconds"], "syntax": "attribute", "type": "number"}, {"aliases": ["site"], "anchor": "schema-site", "description": "site", "document_id": "xcsh-docs:data-sources:smsv2_aws_runtime:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["site"], "syntax": "attribute", "type": "string"}, {"aliases": ["duration", "timeout seconds"], "anchor": "schema-timeout_seconds", "description": "timeout seconds", "document_id": "xcsh-docs:data-sources:smsv2_aws_runtime:reference", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["timeout_seconds"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/smsv2_aws_runtime/properties/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Property reference for xcsh_smsv2_aws_runtime.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": [], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_smsv2_aws_runtime](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_aws_runtime/)
- Property reference

## Direct properties

<a id="schema-healthy"></a>

### healthy property

Type: `"bool"`. Computed.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

- [interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_aws_runtime/properties/interfaces/): complete subsection reference.

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{stringvalidator.OneOf("system")}
```

- [nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_aws_runtime/properties/nodes/): complete subsection reference.

<a id="schema-poll_interval_seconds"></a>

### poll_interval_seconds property

Type: `"number"`. Optional, Computed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{int64validator.Between(1, 60)}
```

<a id="schema-site"></a>

### site property

Type: `"string"`. Required.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{stringvalidator.LengthAtLeast(1)}
```

<a id="schema-timeout_seconds"></a>

### timeout_seconds property

Type: `"number"`. Optional, Computed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{int64validator.Between(1, 7200)}
```

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `healthy` | [healthy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_aws_runtime/properties/#schema-healthy) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_aws_runtime/properties/#schema-id) |
| `interfaces` | [interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_aws_runtime/properties/interfaces/#section) |
| `interfaces.healthy` | [interfaces.healthy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_aws_runtime/properties/interfaces/#schema-interfaces--healthy) |
| `interfaces.interface_name` | [interfaces.interface_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_aws_runtime/properties/interfaces/#schema-interfaces--interface_name) |
| `interfaces.mac` | [interfaces.mac](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_aws_runtime/properties/interfaces/#schema-interfaces--mac) |
| `interfaces.mtu` | [interfaces.mtu](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_aws_runtime/properties/interfaces/#schema-interfaces--mtu) |
| `interfaces.node` | [interfaces.node](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_aws_runtime/properties/interfaces/#schema-interfaces--node) |
| `interfaces.role` | [interfaces.role](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_aws_runtime/properties/interfaces/#schema-interfaces--role) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_aws_runtime/properties/#schema-namespace) |
| `nodes` | [nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_aws_runtime/properties/nodes/#section) |
| `nodes.mac` | [nodes.mac](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_aws_runtime/properties/nodes/#schema-nodes--mac) |
| `nodes.node` | [nodes.node](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_aws_runtime/properties/nodes/#schema-nodes--node) |
| `nodes.role` | [nodes.role](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_aws_runtime/properties/nodes/#schema-nodes--role) |
| `poll_interval_seconds` | [poll_interval_seconds](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_aws_runtime/properties/#schema-poll_interval_seconds) |
| `site` | [site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_aws_runtime/properties/#schema-site) |
| `timeout_seconds` | [timeout_seconds](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_aws_runtime/properties/#schema-timeout_seconds) |

## Next pages

- [interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_aws_runtime/properties/interfaces/)
- [nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_aws_runtime/properties/nodes/)
- [xcsh_smsv2_aws_runtime](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_aws_runtime/)
