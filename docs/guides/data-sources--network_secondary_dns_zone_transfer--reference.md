---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_network_secondary_dns_zone_transfer."
xcsh_docs: {"aliases": [], "body_bytes": 2543, "body_sha256": "sha256:d884b504a96fe44cabbe546df2cf575f9099dab3ad81b2d52e1624114d6dd11a", "canonical_id": "xcsh-docs:data-sources:network_secondary_dns_zone_transfer:reference", "child_ids": [], "collection_id": "xcsh-docs:data-sources:network_secondary_dns_zone_transfer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_secondary_dns_zone_transfer:reference", "parent_id": "xcsh-docs:data-sources:network_secondary_dns_zone_transfer:fundamentals", "path": "docs/guides/data-sources--network_secondary_dns_zone_transfer--reference.md", "provider_name": "network_secondary_dns_zone_transfer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_secondary_dns_zone_transfer/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_network_secondary_dns_zone_transfer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_network_secondary_dns_zone_transfer](../data-sources/network_secondary_dns_zone_transfer.md)
- Property reference

## Direct properties

<a id="schema-api_release_tag"></a>

### api_release_tag property

Type: `"string"`. Computed.

Pinned api-specs-enriched release tag compiled into this provider.

<a id="schema-cidr_blocks"></a>

### cidr_blocks property

Type: `["list", "string"]`. Computed.

Sorted unique IPv4 CIDRs. Individual IPv4 addresses are normalized to /32.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Stable identifier derived from the pinned source digest, data-source group, and selected regions.

<a id="schema-manifest_generated_at"></a>

### manifest_generated_at property

Type: `"string"`. Computed.

Generation timestamp reported by the published network allowlist manifest.

<a id="schema-source_entries"></a>

### source_entries property

Type: `["list", "string"]`. Computed.

IPv4 entries exactly as represented by the published group, preserving source order.

<a id="schema-source_sha256"></a>

### source_sha256 property

Type: `"string"`. Computed.

SHA-256 digest of the canonical published network allowlist manifest.

<a id="schema-source_url"></a>

### source_url property

Type: `"string"`. Computed.

Published F5 network allowlist source URL.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `api_release_tag` | [api_release_tag](data-sources--network_secondary_dns_zone_transfer--reference.md#schema-api_release_tag) |
| `cidr_blocks` | [cidr_blocks](data-sources--network_secondary_dns_zone_transfer--reference.md#schema-cidr_blocks) |
| `id` | [id](data-sources--network_secondary_dns_zone_transfer--reference.md#schema-id) |
| `manifest_generated_at` | [manifest_generated_at](data-sources--network_secondary_dns_zone_transfer--reference.md#schema-manifest_generated_at) |
| `source_entries` | [source_entries](data-sources--network_secondary_dns_zone_transfer--reference.md#schema-source_entries) |
| `source_sha256` | [source_sha256](data-sources--network_secondary_dns_zone_transfer--reference.md#schema-source_sha256) |
| `source_url` | [source_url](data-sources--network_secondary_dns_zone_transfer--reference.md#schema-source_url) |

## Next pages

- [xcsh_network_secondary_dns_zone_transfer](../data-sources/network_secondary_dns_zone_transfer.md)
