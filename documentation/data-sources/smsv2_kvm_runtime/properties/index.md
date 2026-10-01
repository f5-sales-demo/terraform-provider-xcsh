---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_smsv2_kvm_runtime."
xcsh_docs: {"aliases": [], "body_bytes": 4826, "body_sha256": "sha256:d06c5e83edda148907919cae610f5a8b87252d7bf247d6bea0cc59ce5a7ab8ac", "child_ids": [], "collection_id": "xcsh-docs:data-sources:smsv2_kvm_runtime:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:smsv2_kvm_runtime:reference", "parent_id": "xcsh-docs:data-sources:smsv2_kvm_runtime:fundamentals", "path": "documentation/data-sources/smsv2_kvm_runtime/properties/index.md", "provider_name": "smsv2_kvm_runtime", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/smsv2_kvm_runtime/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_smsv2_kvm_runtime.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

## Next pages

- [xcsh_smsv2_kvm_runtime](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_kvm_runtime/)
