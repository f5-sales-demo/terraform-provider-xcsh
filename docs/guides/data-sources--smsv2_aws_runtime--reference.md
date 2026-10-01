---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_smsv2_aws_runtime."
xcsh_docs: {"aliases": [], "body_bytes": 3805, "body_sha256": "sha256:41f593b4e0be13e578006691afd48064bffdc42eea43cbbaec61bfd02dd7aa64", "canonical_id": "xcsh-docs:data-sources:smsv2_aws_runtime:reference", "child_ids": ["xcsh-docs:data-sources:smsv2_aws_runtime:properties:interfaces", "xcsh-docs:data-sources:smsv2_aws_runtime:properties:nodes"], "collection_id": "xcsh-docs:data-sources:smsv2_aws_runtime:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:smsv2_aws_runtime:reference", "parent_id": "xcsh-docs:data-sources:smsv2_aws_runtime:fundamentals", "path": "docs/guides/data-sources--smsv2_aws_runtime--reference.md", "provider_name": "smsv2_aws_runtime", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/smsv2_aws_runtime/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_smsv2_aws_runtime.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_smsv2_aws_runtime](../data-sources/smsv2_aws_runtime.md)
- Property reference

## Direct properties

<a id="schema-healthy"></a>

### healthy property

Type: `"bool"`. Computed.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

- [interfaces](data-sources--smsv2_aws_runtime--properties--interfaces.md): complete subsection reference.

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{stringvalidator.OneOf("system")}
```

- [nodes](data-sources--smsv2_aws_runtime--properties--nodes.md): complete subsection reference.

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
| `healthy` | [healthy](data-sources--smsv2_aws_runtime--reference.md#schema-healthy) |
| `id` | [id](data-sources--smsv2_aws_runtime--reference.md#schema-id) |
| `interfaces` | [interfaces](data-sources--smsv2_aws_runtime--properties--interfaces.md#section) |
| `interfaces.healthy` | [interfaces.healthy](data-sources--smsv2_aws_runtime--properties--interfaces.md#schema-interfaces--healthy) |
| `interfaces.interface_name` | [interfaces.interface_name](data-sources--smsv2_aws_runtime--properties--interfaces.md#schema-interfaces--interface_name) |
| `interfaces.mac` | [interfaces.mac](data-sources--smsv2_aws_runtime--properties--interfaces.md#schema-interfaces--mac) |
| `interfaces.mtu` | [interfaces.mtu](data-sources--smsv2_aws_runtime--properties--interfaces.md#schema-interfaces--mtu) |
| `interfaces.node` | [interfaces.node](data-sources--smsv2_aws_runtime--properties--interfaces.md#schema-interfaces--node) |
| `interfaces.role` | [interfaces.role](data-sources--smsv2_aws_runtime--properties--interfaces.md#schema-interfaces--role) |
| `namespace` | [namespace](data-sources--smsv2_aws_runtime--reference.md#schema-namespace) |
| `nodes` | [nodes](data-sources--smsv2_aws_runtime--properties--nodes.md#section) |
| `nodes.mac` | [nodes.mac](data-sources--smsv2_aws_runtime--properties--nodes.md#schema-nodes--mac) |
| `nodes.node` | [nodes.node](data-sources--smsv2_aws_runtime--properties--nodes.md#schema-nodes--node) |
| `nodes.role` | [nodes.role](data-sources--smsv2_aws_runtime--properties--nodes.md#schema-nodes--role) |
| `poll_interval_seconds` | [poll_interval_seconds](data-sources--smsv2_aws_runtime--reference.md#schema-poll_interval_seconds) |
| `site` | [site](data-sources--smsv2_aws_runtime--reference.md#schema-site) |
| `timeout_seconds` | [timeout_seconds](data-sources--smsv2_aws_runtime--reference.md#schema-timeout_seconds) |

## Next pages

- [interfaces](data-sources--smsv2_aws_runtime--properties--interfaces.md)
- [nodes](data-sources--smsv2_aws_runtime--properties--nodes.md)
- [xcsh_smsv2_aws_runtime](../data-sources/smsv2_aws_runtime.md)
