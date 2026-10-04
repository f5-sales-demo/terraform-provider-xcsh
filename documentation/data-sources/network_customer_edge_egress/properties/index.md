---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_network_customer_edge_egress."
xcsh_docs: {"aliases": ["network customer edge egress"], "body_bytes": 3048, "body_sha256": "sha256:5b9234542ef988285763b5f9d16eeedeebb9ec7e1734668ebeb94885de62c6f6", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_customer_edge_egress:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_customer_edge_egress:reference", "parent_id": "xcsh-docs:data-sources:network_customer_edge_egress:fundamentals", "path": "documentation/data-sources/network_customer_edge_egress/properties/index.md", "product": "distributed-cloud", "provider_name": "network_customer_edge_egress", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3330330333211022-3210001223230133-2222131201231022-2130313303130101-2303311233122033-0012211213111320-0130220120201232-0201122200232020", "registry_path": "docs/guides/data-sources--network_customer_edge_egress--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["api release tag"], "anchor": "schema-api_release_tag", "description": "Pinned api-specs-enriched release tag compiled into this provider.", "document_id": "xcsh-docs:data-sources:network_customer_edge_egress:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_release_tag"], "syntax": "attribute", "type": "string"}, {"aliases": ["domains"], "anchor": "schema-domains", "description": "Published Secure Mesh v2 egress domains in source order. Leading dots are preserved.", "document_id": "xcsh-docs:data-sources:network_customer_edge_egress:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["domains"], "syntax": "attribute", "type": "list"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Stable identifier derived from the pinned source digest, data-source group, and selected regions.", "document_id": "xcsh-docs:data-sources:network_customer_edge_egress:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["manifest generated at"], "anchor": "schema-manifest_generated_at", "description": "Generation timestamp reported by the published network allowlist manifest.", "document_id": "xcsh-docs:data-sources:network_customer_edge_egress:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["manifest_generated_at"], "syntax": "attribute", "type": "string"}, {"aliases": ["registration addresses"], "anchor": "schema-registration_addresses", "description": "Published Secure Mesh v2 registration and update IPv4 addresses.", "document_id": "xcsh-docs:data-sources:network_customer_edge_egress:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["registration_addresses"], "syntax": "attribute", "type": "list"}, {"aliases": ["source sha256"], "anchor": "schema-source_sha256", "description": "SHA-256 digest of the canonical published network allowlist manifest.", "document_id": "xcsh-docs:data-sources:network_customer_edge_egress:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["source_sha256"], "syntax": "attribute", "type": "string"}, {"aliases": ["source url"], "anchor": "schema-source_url", "description": "Published F5 network allowlist source URL.", "document_id": "xcsh-docs:data-sources:network_customer_edge_egress:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["source_url"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_customer_edge_egress/properties/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Property reference for xcsh_network_customer_edge_egress.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": [], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_network_customer_edge_egress](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_customer_edge_egress/)
- Property reference

## Direct properties

<a id="schema-api_release_tag"></a>

### api_release_tag property

Type: `"string"`. Computed.

Pinned api-specs-enriched release tag compiled into this provider.

<a id="schema-domains"></a>

### domains property

Type: `["list", "string"]`. Computed.

Published Secure Mesh v2 egress domains in source order. Leading dots are preserved.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Stable identifier derived from the pinned source digest, data-source group, and selected regions.

<a id="schema-manifest_generated_at"></a>

### manifest_generated_at property

Type: `"string"`. Computed.

Generation timestamp reported by the published network allowlist manifest.

<a id="schema-registration_addresses"></a>

### registration_addresses property

Type: `["list", "string"]`. Computed.

Published Secure Mesh v2 registration and update IPv4 addresses.

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
| `api_release_tag` | [api_release_tag](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_customer_edge_egress/properties/#schema-api_release_tag) |
| `domains` | [domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_customer_edge_egress/properties/#schema-domains) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_customer_edge_egress/properties/#schema-id) |
| `manifest_generated_at` | [manifest_generated_at](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_customer_edge_egress/properties/#schema-manifest_generated_at) |
| `registration_addresses` | [registration_addresses](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_customer_edge_egress/properties/#schema-registration_addresses) |
| `source_sha256` | [source_sha256](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_customer_edge_egress/properties/#schema-source_sha256) |
| `source_url` | [source_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_customer_edge_egress/properties/#schema-source_url) |

## Next pages

- [xcsh_network_customer_edge_egress](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_customer_edge_egress/)
