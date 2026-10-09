---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_waf_attack_signatures."
xcsh_docs: {"aliases": ["waf attack signatures"], "body_bytes": 4068, "body_sha256": "sha256:80646321fc9d3705abf5f4e4b1dbfd5d802a54cbb4f01ecd2ebdfb48123f0e1d", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:waf_attack_signatures:properties:attack_signatures"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:waf_attack_signatures:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:waf_attack_signatures:reference", "parent_id": "xcsh-docs:data-sources:waf_attack_signatures:fundamentals", "path": "documentation/data-sources/waf_attack_signatures/properties/index.md", "product": "distributed-cloud", "provider_name": "waf_attack_signatures", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-3120200000301112-2000001222010231-1030000101003120-0002131010003101-1121111122231111-2121223333310231-1312321220022310-2131323003202332", "registry_path": "docs/guides/data-sources--waf_attack_signatures--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["attack signatures"], "anchor": "section", "description": "List of all supported attack signatures.", "document_id": "xcsh-docs:data-sources:waf_attack_signatures:properties:attack_signatures", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["attack_signatures"], "syntax": "attribute", "type": "object"}, {"aliases": ["known version"], "anchor": "schema-known_version", "description": "Version of the attack signatures list that the client currently has, can be used for caching.", "document_id": "xcsh-docs:data-sources:waf_attack_signatures:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["known_version"], "syntax": "attribute", "type": "string"}, {"aliases": ["not modified"], "anchor": "schema-not_modified", "description": "Indicates if the attack signatures list has not been modified since the last version.", "document_id": "xcsh-docs:data-sources:waf_attack_signatures:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["not_modified"], "syntax": "attribute", "type": "bool"}, {"aliases": ["version"], "anchor": "schema-version", "description": "Version of the attack signatures list, can be used for caching.", "document_id": "xcsh-docs:data-sources:waf_attack_signatures:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["version"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/waf_attack_signatures/properties/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Property reference for xcsh_waf_attack_signatures.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": [], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_waf_attack_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_attack_signatures/)
- Property reference

## Direct properties

- [attack_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_attack_signatures/properties/attack_signatures/): complete subsection reference.

<a id="schema-known_version"></a>

### known_version property

Type: `"string"`. Optional.

Version of the attack signatures list that the client currently has, can be used for caching.

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
| `attack_signatures` | [attack_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_attack_signatures/properties/attack_signatures/#section) |
| `attack_signatures.accuracy` | [attack_signatures.accuracy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_attack_signatures/properties/attack_signatures/#schema-attack_signatures--accuracy) |
| `attack_signatures.applies_to` | [attack_signatures.applies_to](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_attack_signatures/properties/attack_signatures/#schema-attack_signatures--applies_to) |
| `attack_signatures.attack_type` | [attack_signatures.attack_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_attack_signatures/properties/attack_signatures/#schema-attack_signatures--attack_type) |
| `attack_signatures.description_spec` | [attack_signatures.description_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_attack_signatures/properties/attack_signatures/#schema-attack_signatures--description_spec) |
| `attack_signatures.id` | [attack_signatures.id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_attack_signatures/properties/attack_signatures/#schema-attack_signatures--id) |
| `attack_signatures.last_update` | [attack_signatures.last_update](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_attack_signatures/properties/attack_signatures/#schema-attack_signatures--last_update) |
| `attack_signatures.name` | [attack_signatures.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_attack_signatures/properties/attack_signatures/#schema-attack_signatures--name) |
| `attack_signatures.references` | [attack_signatures.references](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_attack_signatures/properties/attack_signatures/#schema-attack_signatures--references) |
| `attack_signatures.risk` | [attack_signatures.risk](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_attack_signatures/properties/attack_signatures/#schema-attack_signatures--risk) |
| `attack_signatures.systems` | [attack_signatures.systems](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_attack_signatures/properties/attack_signatures/#schema-attack_signatures--systems) |
| `known_version` | [known_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_attack_signatures/properties/#schema-known_version) |
| `not_modified` | [not_modified](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_attack_signatures/properties/#schema-not_modified) |
| `version` | [version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_attack_signatures/properties/#schema-version) |
