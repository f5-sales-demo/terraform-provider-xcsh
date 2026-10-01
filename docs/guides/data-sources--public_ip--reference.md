---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_public_ip."
xcsh_docs: {"aliases": [], "body_bytes": 3017, "body_sha256": "sha256:c5421b608bf924ea2f26c0a0d575bdca0c4f9f9d043d9d43dfc1ef96625e3e0c", "canonical_id": "xcsh-docs:data-sources:public_ip:reference", "child_ids": ["xcsh-docs:data-sources:public_ip:properties:virtual_sites"], "collection_id": "xcsh-docs:data-sources:public_ip:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:public_ip:reference", "parent_id": "xcsh-docs:data-sources:public_ip:fundamentals", "path": "docs/guides/data-sources--public_ip--reference.md", "provider_name": "public_ip", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/public_ip/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_public_ip.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_public_ip](../data-sources/public_ip.md)
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

- [virtual_sites](data-sources--public_ip--properties--virtual_sites.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--public_ip--reference.md#schema-annotations) |
| `description` | [description](data-sources--public_ip--reference.md#schema-description) |
| `id` | [id](data-sources--public_ip--reference.md#schema-id) |
| `ip` | [ip](data-sources--public_ip--reference.md#schema-ip) |
| `ipv6` | [ipv6](data-sources--public_ip--reference.md#schema-ipv6) |
| `labels` | [labels](data-sources--public_ip--reference.md#schema-labels) |
| `name` | [name](data-sources--public_ip--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--public_ip--reference.md#schema-namespace) |
| `virtual_sites` | [virtual_sites](data-sources--public_ip--properties--virtual_sites.md#section) |
| `virtual_sites.kind` | [virtual_sites.kind](data-sources--public_ip--properties--virtual_sites.md#schema-virtual_sites--kind) |
| `virtual_sites.name` | [virtual_sites.name](data-sources--public_ip--properties--virtual_sites.md#schema-virtual_sites--name) |
| `virtual_sites.namespace` | [virtual_sites.namespace](data-sources--public_ip--properties--virtual_sites.md#schema-virtual_sites--namespace) |
| `virtual_sites.tenant` | [virtual_sites.tenant](data-sources--public_ip--properties--virtual_sites.md#schema-virtual_sites--tenant) |
| `virtual_sites.uid` | [virtual_sites.uid](data-sources--public_ip--properties--virtual_sites.md#schema-virtual_sites--uid) |

## Next pages

- [virtual_sites](data-sources--public_ip--properties--virtual_sites.md)
- [xcsh_public_ip](../data-sources/public_ip.md)
