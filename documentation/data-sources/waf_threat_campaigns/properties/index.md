---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_waf_threat_campaigns."
xcsh_docs: {"aliases": [], "body_bytes": 4272, "body_sha256": "sha256:137e1524c8447300e1711fd72ec8ec3fb792ed55edb095552697449d9b173d8b", "child_ids": ["xcsh-docs:data-sources:waf_threat_campaigns:properties:threat_campaigns"], "collection_id": "xcsh-docs:data-sources:waf_threat_campaigns:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:waf_threat_campaigns:reference", "parent_id": "xcsh-docs:data-sources:waf_threat_campaigns:fundamentals", "path": "documentation/data-sources/waf_threat_campaigns/properties/index.md", "provider_name": "waf_threat_campaigns", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/waf_threat_campaigns/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_waf_threat_campaigns.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_waf_threat_campaigns](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_threat_campaigns/)
- Property reference

## Direct properties

<a id="schema-known_version"></a>

### known_version property

Type: `"string"`. Optional.

Version of the threat campaigns list that the client currently has, can be used for caching.

<a id="schema-not_modified"></a>

### not_modified property

Type: `"bool"`. Computed.

Indicates if the attack signatures list has not been modified since the last version.

- [threat_campaigns](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_threat_campaigns/properties/threat_campaigns/): complete subsection reference.

<a id="schema-version"></a>

### version property

Type: `"string"`. Computed.

Version of the attack signatures list, can be used for caching.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `known_version` | [known_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_threat_campaigns/properties/#schema-known_version) |
| `not_modified` | [not_modified](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_threat_campaigns/properties/#schema-not_modified) |
| `threat_campaigns` | [threat_campaigns](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_threat_campaigns/properties/threat_campaigns/#section) |
| `threat_campaigns.attack_type` | [threat_campaigns.attack_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_threat_campaigns/properties/threat_campaigns/#schema-threat_campaigns--attack_type) |
| `threat_campaigns.description_spec` | [threat_campaigns.description_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_threat_campaigns/properties/threat_campaigns/#schema-threat_campaigns--description_spec) |
| `threat_campaigns.id` | [threat_campaigns.id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_threat_campaigns/properties/threat_campaigns/#schema-threat_campaigns--id) |
| `threat_campaigns.intent` | [threat_campaigns.intent](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_threat_campaigns/properties/threat_campaigns/#schema-threat_campaigns--intent) |
| `threat_campaigns.last_update` | [threat_campaigns.last_update](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_threat_campaigns/properties/threat_campaigns/#schema-threat_campaigns--last_update) |
| `threat_campaigns.malwares` | [threat_campaigns.malwares](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_threat_campaigns/properties/threat_campaigns/#schema-threat_campaigns--malwares) |
| `threat_campaigns.name` | [threat_campaigns.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_threat_campaigns/properties/threat_campaigns/#schema-threat_campaigns--name) |
| `threat_campaigns.references` | [threat_campaigns.references](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_threat_campaigns/properties/threat_campaigns/#schema-threat_campaigns--references) |
| `threat_campaigns.risk` | [threat_campaigns.risk](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_threat_campaigns/properties/threat_campaigns/#schema-threat_campaigns--risk) |
| `threat_campaigns.systems` | [threat_campaigns.systems](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_threat_campaigns/properties/threat_campaigns/#schema-threat_campaigns--systems) |
| `version` | [version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_threat_campaigns/properties/#schema-version) |

## Next pages

- [threat_campaigns](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_threat_campaigns/properties/threat_campaigns/)
- [xcsh_waf_threat_campaigns](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_threat_campaigns/)
