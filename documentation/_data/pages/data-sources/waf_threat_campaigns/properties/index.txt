---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_waf_threat_campaigns."
xcsh_docs: {"aliases": ["waf threat campaigns"], "body_bytes": 3993, "body_sha256": "sha256:b451c714d21925a1d86dd207bc29a1556e09734a4fd108ad2e908dce38f6c45b", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:waf_threat_campaigns:properties:threat_campaigns"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:waf_threat_campaigns:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:waf_threat_campaigns:reference", "parent_id": "xcsh-docs:data-sources:waf_threat_campaigns:fundamentals", "path": "documentation/data-sources/waf_threat_campaigns/properties/index.md", "product": "distributed-cloud", "provider_name": "waf_threat_campaigns", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "data-sources", "registry_anchor": "canonical-2312320213131010-2313332320110100-1222123311130230-1202203132203210-2221111330220221-2123233132102221-2022012303331212-0311010310111022", "registry_path": "docs/guides/data-sources--waf_threat_campaigns--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["known version"], "anchor": "schema-known_version", "description": "Version of the threat campaigns list that the client currently has, can be used for caching.", "document_id": "xcsh-docs:data-sources:waf_threat_campaigns:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["known_version"], "syntax": "attribute", "type": "string"}, {"aliases": ["not modified"], "anchor": "schema-not_modified", "description": "Indicates if the attack signatures list has not been modified since the last version.", "document_id": "xcsh-docs:data-sources:waf_threat_campaigns:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["not_modified"], "syntax": "attribute", "type": "bool"}, {"aliases": ["threat campaigns"], "anchor": "section", "description": "Threat Campaigns. A list of all supported threat campaigns.", "document_id": "xcsh-docs:data-sources:waf_threat_campaigns:properties:threat_campaigns", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["threat_campaigns"], "syntax": "attribute", "type": "object"}, {"aliases": ["version"], "anchor": "schema-version", "description": "Version of the attack signatures list, can be used for caching.", "document_id": "xcsh-docs:data-sources:waf_threat_campaigns:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["version"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/waf_threat_campaigns/properties/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Property reference for xcsh_waf_threat_campaigns.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": [], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
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
