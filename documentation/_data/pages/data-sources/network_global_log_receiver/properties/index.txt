---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_network_global_log_receiver."
xcsh_docs: {"aliases": ["network global log receiver"], "body_bytes": 3027, "body_sha256": "sha256:adfa51fd5be99643bb187c2eef1d7459fc3770a1a0059a4d2cb6d4b826bff7fe", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_global_log_receiver:reference", "parent_id": "xcsh-docs:data-sources:network_global_log_receiver:fundamentals", "path": "documentation/data-sources/network_global_log_receiver/properties/index.md", "product": "distributed-cloud", "provider_name": "network_global_log_receiver", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "data-sources", "registry_anchor": "canonical-1302231310000103-2202232011203310-1003230012120012-3310013102103212-0201101202113030-3013331321011231-2133210011232031-2132201131331033", "registry_path": "docs/guides/data-sources--network_global_log_receiver--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["api release tag"], "anchor": "schema-api_release_tag", "description": "Pinned api-specs-enriched release tag compiled into this provider.", "document_id": "xcsh-docs:data-sources:network_global_log_receiver:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_release_tag"], "syntax": "attribute", "type": "string"}, {"aliases": ["cidr blocks"], "anchor": "schema-cidr_blocks", "description": "Sorted unique IPv4 CIDRs. Individual IPv4 addresses are normalized to /32.", "document_id": "xcsh-docs:data-sources:network_global_log_receiver:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cidr_blocks"], "syntax": "attribute", "type": "list"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Stable identifier derived from the pinned source digest, data-source group, and selected regions.", "document_id": "xcsh-docs:data-sources:network_global_log_receiver:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["manifest generated at"], "anchor": "schema-manifest_generated_at", "description": "Generation timestamp reported by the published network allowlist manifest.", "document_id": "xcsh-docs:data-sources:network_global_log_receiver:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["manifest_generated_at"], "syntax": "attribute", "type": "string"}, {"aliases": ["source entries"], "anchor": "schema-source_entries", "description": "IPv4 entries exactly as represented by the published group, preserving source order.", "document_id": "xcsh-docs:data-sources:network_global_log_receiver:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["source_entries"], "syntax": "attribute", "type": "list"}, {"aliases": ["source sha256"], "anchor": "schema-source_sha256", "description": "SHA-256 digest of the canonical published network allowlist manifest.", "document_id": "xcsh-docs:data-sources:network_global_log_receiver:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["source_sha256"], "syntax": "attribute", "type": "string"}, {"aliases": ["source url"], "anchor": "schema-source_url", "description": "Published F5 network allowlist source URL.", "document_id": "xcsh-docs:data-sources:network_global_log_receiver:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["source_url"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_global_log_receiver/properties/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Property reference for xcsh_network_global_log_receiver.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_network_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_global_log_receiver/)
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
| `api_release_tag` | [api_release_tag](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_global_log_receiver/properties/#schema-api_release_tag) |
| `cidr_blocks` | [cidr_blocks](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_global_log_receiver/properties/#schema-cidr_blocks) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_global_log_receiver/properties/#schema-id) |
| `manifest_generated_at` | [manifest_generated_at](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_global_log_receiver/properties/#schema-manifest_generated_at) |
| `source_entries` | [source_entries](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_global_log_receiver/properties/#schema-source_entries) |
| `source_sha256` | [source_sha256](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_global_log_receiver/properties/#schema-source_sha256) |
| `source_url` | [source_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_global_log_receiver/properties/#schema-source_url) |

## Next pages

- [xcsh_network_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_global_log_receiver/)
