---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_waf_attack_signatures."
xcsh_docs: {"aliases": [], "body_bytes": 3428, "body_sha256": "sha256:a2ee6603f7fa6a7b54384b4768703b5d2bc593de0422040889a7ebfd3fc770f7", "canonical_id": "xcsh-docs:data-sources:waf_attack_signatures:reference", "child_ids": ["xcsh-docs:data-sources:waf_attack_signatures:properties:attack_signatures"], "collection_id": "xcsh-docs:data-sources:waf_attack_signatures:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:waf_attack_signatures:reference", "parent_id": "xcsh-docs:data-sources:waf_attack_signatures:fundamentals", "path": "docs/guides/data-sources--waf_attack_signatures--reference.md", "provider_name": "waf_attack_signatures", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/waf_attack_signatures/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_waf_attack_signatures.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_waf_attack_signatures](../data-sources/waf_attack_signatures.md)
- Property reference

## Direct properties

- [attack_signatures](data-sources--waf_attack_signatures--properties--attack_signatures.md): complete subsection reference.

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
| `attack_signatures` | [attack_signatures](data-sources--waf_attack_signatures--properties--attack_signatures.md#section) |
| `attack_signatures.accuracy` | [attack_signatures.accuracy](data-sources--waf_attack_signatures--properties--attack_signatures.md#schema-attack_signatures--accuracy) |
| `attack_signatures.applies_to` | [attack_signatures.applies_to](data-sources--waf_attack_signatures--properties--attack_signatures.md#schema-attack_signatures--applies_to) |
| `attack_signatures.attack_type` | [attack_signatures.attack_type](data-sources--waf_attack_signatures--properties--attack_signatures.md#schema-attack_signatures--attack_type) |
| `attack_signatures.description_spec` | [attack_signatures.description_spec](data-sources--waf_attack_signatures--properties--attack_signatures.md#schema-attack_signatures--description_spec) |
| `attack_signatures.id` | [attack_signatures.id](data-sources--waf_attack_signatures--properties--attack_signatures.md#schema-attack_signatures--id) |
| `attack_signatures.last_update` | [attack_signatures.last_update](data-sources--waf_attack_signatures--properties--attack_signatures.md#schema-attack_signatures--last_update) |
| `attack_signatures.name` | [attack_signatures.name](data-sources--waf_attack_signatures--properties--attack_signatures.md#schema-attack_signatures--name) |
| `attack_signatures.references` | [attack_signatures.references](data-sources--waf_attack_signatures--properties--attack_signatures.md#schema-attack_signatures--references) |
| `attack_signatures.risk` | [attack_signatures.risk](data-sources--waf_attack_signatures--properties--attack_signatures.md#schema-attack_signatures--risk) |
| `attack_signatures.systems` | [attack_signatures.systems](data-sources--waf_attack_signatures--properties--attack_signatures.md#schema-attack_signatures--systems) |
| `known_version` | [known_version](data-sources--waf_attack_signatures--reference.md#schema-known_version) |
| `not_modified` | [not_modified](data-sources--waf_attack_signatures--reference.md#schema-not_modified) |
| `version` | [version](data-sources--waf_attack_signatures--reference.md#schema-version) |

## Next pages

- [attack_signatures](data-sources--waf_attack_signatures--properties--attack_signatures.md)
- [xcsh_waf_attack_signatures](../data-sources/waf_attack_signatures.md)
