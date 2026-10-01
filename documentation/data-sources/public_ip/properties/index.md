---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_public_ip."
xcsh_docs: {"aliases": [], "body_bytes": 3951, "body_sha256": "sha256:e6658b60c95eaf28447855a092c17fe8a28d15c2efb259389af87efb30dc347c", "child_ids": ["xcsh-docs:data-sources:public_ip:properties:virtual_sites"], "collection_id": "xcsh-docs:data-sources:public_ip:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:public_ip:reference", "parent_id": "xcsh-docs:data-sources:public_ip:fundamentals", "path": "documentation/data-sources/public_ip/properties/index.md", "provider_name": "public_ip", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/public_ip/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_public_ip.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/public_ip/)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Computed.

Annotations.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier.

<a id="schema-ip"></a>

### ip property

Type: `"string"`. Computed.

IPv4 address for this object. An empty string indicates no IPv4 address is configured.

<a id="schema-ipv6"></a>

### ipv6 property

Type: `"string"`. Computed.

IPv6 address for this object. An empty string indicates no IPv6 address is configured.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Computed.

Labels.

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the PublicIP to look up.

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace of the PublicIP.

- [virtual_sites](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/public_ip/properties/virtual_sites/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/public_ip/properties/#schema-annotations) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/public_ip/properties/#schema-description) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/public_ip/properties/#schema-id) |
| `ip` | [ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/public_ip/properties/#schema-ip) |
| `ipv6` | [ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/public_ip/properties/#schema-ipv6) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/public_ip/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/public_ip/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/public_ip/properties/#schema-namespace) |
| `virtual_sites` | [virtual_sites](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/public_ip/properties/virtual_sites/#section) |
| `virtual_sites.kind` | [virtual_sites.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/public_ip/properties/virtual_sites/#schema-virtual_sites--kind) |
| `virtual_sites.name` | [virtual_sites.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/public_ip/properties/virtual_sites/#schema-virtual_sites--name) |
| `virtual_sites.namespace` | [virtual_sites.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/public_ip/properties/virtual_sites/#schema-virtual_sites--namespace) |
| `virtual_sites.tenant` | [virtual_sites.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/public_ip/properties/virtual_sites/#schema-virtual_sites--tenant) |
| `virtual_sites.uid` | [virtual_sites.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/public_ip/properties/virtual_sites/#schema-virtual_sites--uid) |

## Next pages

- [virtual_sites](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/public_ip/properties/virtual_sites/)
- [xcsh_public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/public_ip/)
