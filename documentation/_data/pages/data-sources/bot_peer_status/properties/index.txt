---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_bot_peer_status."
xcsh_docs: {"aliases": ["bot peer status"], "body_bytes": 1120, "body_sha256": "sha256:a74b3cb9fd5fc0954ea936c74420a6c0ba6176511841e843e7eed8ee854d35fc", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_peer_status:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_peer_status:reference", "parent_id": "xcsh-docs:data-sources:bot_peer_status:fundamentals", "path": "documentation/data-sources/bot_peer_status/properties/index.md", "product": "distributed-cloud", "provider_name": "bot_peer_status", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-2233112331322323-3212110322020321-3202221110020001-3200220303102201-3002320020113100-1010310321312301-2132001113233301-3312110023102211", "registry_path": "docs/guides/data-sources--bot_peer_status--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["has peers"], "anchor": "schema-has_peers", "description": "Has peers status. The tenat has peers or not.", "document_id": "xcsh-docs:data-sources:bot_peer_status:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["has_peers"], "syntax": "attribute", "type": "bool"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "Namespace. namespace is used to scope the query. Only virtual_host in given namespace will be considered.", "document_id": "xcsh-docs:data-sources:bot_peer_status:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_peer_status/properties/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Property reference for xcsh_bot_peer_status.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": [], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_bot_peer_status](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_status/)
- Property reference

## Direct properties

<a id="schema-has_peers"></a>

### has_peers property

Type: `"bool"`. Computed.

Has peers status. The tenat has peers or not.

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace. namespace is used to scope the query. Only virtual\_host in given namespace will be
considered.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `has_peers` | [has_peers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_status/properties/#schema-has_peers) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_status/properties/#schema-namespace) |
