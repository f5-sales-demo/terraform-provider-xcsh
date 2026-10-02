---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_network_customer_edge_defaults."
xcsh_docs: {"aliases": ["network customer edge defaults"], "body_bytes": 2975, "body_sha256": "sha256:dbbf7c08154080f38b28bbed134e5f90d189269d9cca3fa3d7c222b79e94bd81", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_customer_edge_defaults:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_customer_edge_defaults:reference", "parent_id": "xcsh-docs:data-sources:network_customer_edge_defaults:fundamentals", "path": "documentation/data-sources/network_customer_edge_defaults/properties/index.md", "product": "distributed-cloud", "provider_name": "network_customer_edge_defaults", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2131013332023213-0010232033113021-0202033330201320-2220210231031321-0211222320311011-1100301313023322-0211102303121020-0101122212031210", "registry_path": "docs/guides/data-sources--network_customer_edge_defaults--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["api release tag"], "anchor": "schema-api_release_tag", "description": "Pinned api-specs-enriched release tag compiled into this provider.", "document_id": "xcsh-docs:data-sources:network_customer_edge_defaults:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_release_tag"], "syntax": "attribute", "type": "string"}, {"aliases": ["dns servers"], "anchor": "schema-dns_servers", "description": "Published default DNS server IPv4 addresses.", "document_id": "xcsh-docs:data-sources:network_customer_edge_defaults:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dns_servers"], "syntax": "attribute", "type": "list"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Stable identifier derived from the pinned source digest, data-source group, and selected regions.", "document_id": "xcsh-docs:data-sources:network_customer_edge_defaults:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["manifest generated at"], "anchor": "schema-manifest_generated_at", "description": "Generation timestamp reported by the published network allowlist manifest.", "document_id": "xcsh-docs:data-sources:network_customer_edge_defaults:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["manifest_generated_at"], "syntax": "attribute", "type": "string"}, {"aliases": ["ntp servers"], "anchor": "schema-ntp_servers", "description": "Published default NTP server IPv4 addresses.", "document_id": "xcsh-docs:data-sources:network_customer_edge_defaults:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ntp_servers"], "syntax": "attribute", "type": "list"}, {"aliases": ["source sha256"], "anchor": "schema-source_sha256", "description": "SHA-256 digest of the canonical published network allowlist manifest.", "document_id": "xcsh-docs:data-sources:network_customer_edge_defaults:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["source_sha256"], "syntax": "attribute", "type": "string"}, {"aliases": ["source url"], "anchor": "schema-source_url", "description": "Published F5 network allowlist source URL.", "document_id": "xcsh-docs:data-sources:network_customer_edge_defaults:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["source_url"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_customer_edge_defaults/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_network_customer_edge_defaults.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_network_customer_edge_defaults](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_customer_edge_defaults/)
- Property reference

## Direct properties

<a id="schema-api_release_tag"></a>

### api_release_tag property

Type: `"string"`. Computed.

Pinned api-specs-enriched release tag compiled into this provider.

<a id="schema-dns_servers"></a>

### dns_servers property

Type: `["list", "string"]`. Computed.

Published default DNS server IPv4 addresses.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Stable identifier derived from the pinned source digest, data-source group, and selected regions.

<a id="schema-manifest_generated_at"></a>

### manifest_generated_at property

Type: `"string"`. Computed.

Generation timestamp reported by the published network allowlist manifest.

<a id="schema-ntp_servers"></a>

### ntp_servers property

Type: `["list", "string"]`. Computed.

Published default NTP server IPv4 addresses.

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
| `api_release_tag` | [api_release_tag](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_customer_edge_defaults/properties/#schema-api_release_tag) |
| `dns_servers` | [dns_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_customer_edge_defaults/properties/#schema-dns_servers) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_customer_edge_defaults/properties/#schema-id) |
| `manifest_generated_at` | [manifest_generated_at](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_customer_edge_defaults/properties/#schema-manifest_generated_at) |
| `ntp_servers` | [ntp_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_customer_edge_defaults/properties/#schema-ntp_servers) |
| `source_sha256` | [source_sha256](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_customer_edge_defaults/properties/#schema-source_sha256) |
| `source_url` | [source_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_customer_edge_defaults/properties/#schema-source_url) |

## Next pages

- [xcsh_network_customer_edge_defaults](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_customer_edge_defaults/)
