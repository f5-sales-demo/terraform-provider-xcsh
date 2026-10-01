---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_smsv2_kvm_runtime_interface."
xcsh_docs: {"aliases": [], "body_bytes": 4467, "body_sha256": "sha256:f49651f9584fe28ee72121d384813e7ae77670af80d19277bd11c41f7107c776", "child_ids": [], "collection_id": "xcsh-docs:resources:smsv2_kvm_runtime_interface:collection", "completeness": "complete", "id": "xcsh-docs:resources:smsv2_kvm_runtime_interface:reference", "parent_id": "xcsh-docs:resources:smsv2_kvm_runtime_interface:fundamentals", "path": "documentation/resources/smsv2_kvm_runtime_interface/properties/index.md", "provider_name": "smsv2_kvm_runtime_interface", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/smsv2_kvm_runtime_interface/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_smsv2_kvm_runtime_interface.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
