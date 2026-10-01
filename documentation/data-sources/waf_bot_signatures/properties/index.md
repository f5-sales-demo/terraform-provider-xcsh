---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_waf_bot_signatures."
xcsh_docs: {"aliases": [], "body_bytes": 3509, "body_sha256": "sha256:8e656d850e961aba154af302ac722d588a86990aaae0b75e9dbf7611ec619d40", "child_ids": ["xcsh-docs:data-sources:waf_bot_signatures:properties:bot_signatures"], "collection_id": "xcsh-docs:data-sources:waf_bot_signatures:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:waf_bot_signatures:reference", "parent_id": "xcsh-docs:data-sources:waf_bot_signatures:fundamentals", "path": "documentation/data-sources/waf_bot_signatures/properties/index.md", "provider_name": "waf_bot_signatures", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/waf_bot_signatures/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_waf_bot_signatures.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

## Next pages

- [bot_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_bot_signatures/properties/bot_signatures/)
- [xcsh_waf_bot_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_bot_signatures/)
