---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_waf_attack_signatures."
xcsh_docs: {"aliases": ["waf attack signatures"], "body_bytes": 4352, "body_sha256": "sha256:6e4969572d250ab8d19b249710b1b51a51fa900f4b788bf3ee307262b115981a", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:waf_attack_signatures:properties:attack_signatures"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:waf_attack_signatures:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:waf_attack_signatures:reference", "parent_id": "xcsh-docs:data-sources:waf_attack_signatures:fundamentals", "path": "documentation/data-sources/waf_attack_signatures/properties/index.md", "product": "distributed-cloud", "provider_name": "waf_attack_signatures", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3120200000301112-2000001222010231-1030000101003120-0002131010003101-1121111122231111-2121223333310231-1312321220022310-2131323003202332", "registry_path": "docs/guides/data-sources--waf_attack_signatures--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["attack signatures"], "anchor": "section", "description": "List of all supported attack signatures.", "document_id": "xcsh-docs:data-sources:waf_attack_signatures:properties:attack_signatures", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["attack_signatures"], "syntax": "attribute", "type": "object"}, {"aliases": ["known version"], "anchor": "schema-known_version", "description": "Version of the attack signatures list that the client currently has, can be used for caching.", "document_id": "xcsh-docs:data-sources:waf_attack_signatures:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["known_version"], "syntax": "attribute", "type": "string"}, {"aliases": ["not modified"], "anchor": "schema-not_modified", "description": "Indicates if the attack signatures list has not been modified since the last version.", "document_id": "xcsh-docs:data-sources:waf_attack_signatures:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["not_modified"], "syntax": "attribute", "type": "bool"}, {"aliases": ["version"], "anchor": "schema-version", "description": "Version of the attack signatures list, can be used for caching.", "document_id": "xcsh-docs:data-sources:waf_attack_signatures:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["version"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/waf_attack_signatures/properties/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Property reference for xcsh_waf_attack_signatures.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Next pages

- [attack_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_attack_signatures/properties/attack_signatures/)
- [xcsh_waf_attack_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_attack_signatures/)
