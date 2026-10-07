---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_waf_bot_signatures."
xcsh_docs: {"aliases": ["waf bot signatures"], "body_bytes": 3240, "body_sha256": "sha256:bc8ee5a1d4cd48050c806163e44326feaabe4c40d7a558dc5ba2fc8a3a9b1298", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:waf_bot_signatures:properties:bot_signatures"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:waf_bot_signatures:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:waf_bot_signatures:reference", "parent_id": "xcsh-docs:data-sources:waf_bot_signatures:fundamentals", "path": "documentation/data-sources/waf_bot_signatures/properties/index.md", "product": "distributed-cloud", "provider_name": "waf_bot_signatures", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-0202302311100230-1213232002201200-1112333313213300-1120210203201311-3310221133022103-2132301023122012-0230333001330220-0113223112221303", "registry_path": "docs/guides/data-sources--waf_bot_signatures--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["bot signatures"], "anchor": "section", "description": "Bot Signatures. A list of all supported bot signatures.", "document_id": "xcsh-docs:data-sources:waf_bot_signatures:properties:bot_signatures", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["bot_signatures"], "syntax": "attribute", "type": "object"}, {"aliases": ["known version"], "anchor": "schema-known_version", "description": "Version of the bot signatures list that the client currently has, can be used for caching.", "document_id": "xcsh-docs:data-sources:waf_bot_signatures:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["known_version"], "syntax": "attribute", "type": "string"}, {"aliases": ["not modified"], "anchor": "schema-not_modified", "description": "Indicates if the attack signatures list has not been modified since the last version.", "document_id": "xcsh-docs:data-sources:waf_bot_signatures:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["not_modified"], "syntax": "attribute", "type": "bool"}, {"aliases": ["version"], "anchor": "schema-version", "description": "Version of the attack signatures list, can be used for caching.", "document_id": "xcsh-docs:data-sources:waf_bot_signatures:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["version"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/waf_bot_signatures/properties/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Property reference for xcsh_waf_bot_signatures.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": [], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_waf_bot_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_bot_signatures/)
- Property reference

## Direct properties

- [bot_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_bot_signatures/properties/bot_signatures/): complete subsection reference.

<a id="schema-known_version"></a>

### known_version property

Type: `"string"`. Optional.

Version of the bot signatures list that the client currently has, can be used for caching.

<a id="schema-not_modified"></a>

### not_modified property

Type: `"bool"`. Computed.

Indicates if the attack signatures list has not been modified since the last version.

<a id="schema-version"></a>

### version property

Type: `"string"`. Computed.

Version of the attack signatures list, can be used for caching.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `bot_signatures` | [bot_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_bot_signatures/properties/bot_signatures/#section) |
| `bot_signatures.bot_class` | [bot_signatures.bot_class](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_bot_signatures/properties/bot_signatures/#schema-bot_signatures--bot_class) |
| `bot_signatures.bot_name` | [bot_signatures.bot_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_bot_signatures/properties/bot_signatures/#schema-bot_signatures--bot_name) |
| `bot_signatures.category` | [bot_signatures.category](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_bot_signatures/properties/bot_signatures/#schema-bot_signatures--category) |
| `bot_signatures.hostnames` | [bot_signatures.hostnames](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_bot_signatures/properties/bot_signatures/#schema-bot_signatures--hostnames) |
| `bot_signatures.id` | [bot_signatures.id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_bot_signatures/properties/bot_signatures/#schema-bot_signatures--id) |
| `bot_signatures.last_updated` | [bot_signatures.last_updated](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_bot_signatures/properties/bot_signatures/#schema-bot_signatures--last_updated) |
| `bot_signatures.risk` | [bot_signatures.risk](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_bot_signatures/properties/bot_signatures/#schema-bot_signatures--risk) |
| `known_version` | [known_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_bot_signatures/properties/#schema-known_version) |
| `not_modified` | [not_modified](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_bot_signatures/properties/#schema-not_modified) |
| `version` | [version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_bot_signatures/properties/#schema-version) |
