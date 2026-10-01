---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_smsv2_aws_runtime."
xcsh_docs: {"aliases": [], "body_bytes": 4990, "body_sha256": "sha256:fedcb1c89cc9dc7ed10a8dd0f4312935a723b6ae92af7da3d398ea432c796060", "child_ids": ["xcsh-docs:data-sources:smsv2_aws_runtime:properties:interfaces", "xcsh-docs:data-sources:smsv2_aws_runtime:properties:nodes"], "collection_id": "xcsh-docs:data-sources:smsv2_aws_runtime:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:smsv2_aws_runtime:reference", "parent_id": "xcsh-docs:data-sources:smsv2_aws_runtime:fundamentals", "path": "documentation/data-sources/smsv2_aws_runtime/properties/index.md", "provider_name": "smsv2_aws_runtime", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/smsv2_aws_runtime/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_smsv2_aws_runtime.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
