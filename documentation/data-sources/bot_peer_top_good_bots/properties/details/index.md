---
page_title: "details"
subcategory: ""
description: "Peer Group Top Good Bots Details. Configuration parameter for details"
xcsh_docs: {"aliases": ["details"], "body_bytes": 1491, "body_sha256": "sha256:9abe86e07aea0f8cb7e19155351e7177fd97b0e7591d911f1df05f0e00f3c0c7", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_peer_top_good_bots:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_peer_top_good_bots:properties:details", "parent_id": "xcsh-docs:data-sources:bot_peer_top_good_bots:reference", "path": "documentation/data-sources/bot_peer_top_good_bots/properties/details/index.md", "product": "distributed-cloud", "provider_name": "bot_peer_top_good_bots", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-0120211233303301-3121333012213312-3123002231120022-2000122121311231-0211210020111230-2210210002222200-3020330210302102-3310311130100121", "registry_path": "docs/guides/data-sources--bot_peer_top_good_bots--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["details"], "schema_version": 1, "sections": [{"aliases": ["details name"], "anchor": "schema-details--name", "description": "Name. The Name of Item.", "document_id": "xcsh-docs:data-sources:bot_peer_top_good_bots:properties:details", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["details", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["details peer count"], "anchor": "schema-details--peer_count", "description": "Peer Count. The count of Peer of the Item.", "document_id": "xcsh-docs:data-sources:bot_peer_top_good_bots:properties:details", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["details", "peer_count"], "syntax": "attribute", "type": "string"}, {"aliases": ["details peer percentage"], "anchor": "schema-details--peer_percentage", "description": "Peer Percentage. The Peer Percentage of the Item.", "document_id": "xcsh-docs:data-sources:bot_peer_top_good_bots:properties:details", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["details", "peer_percentage"], "syntax": "attribute", "type": "number"}, {"aliases": ["details self count"], "anchor": "schema-details--self_count", "description": "Self Count. The count of Self of the item.", "document_id": "xcsh-docs:data-sources:bot_peer_top_good_bots:properties:details", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["details", "self_count"], "syntax": "attribute", "type": "string"}, {"aliases": ["details self percentage"], "anchor": "schema-details--self_percentage", "description": "Self Percentage. The Self Percentage of the Item.", "document_id": "xcsh-docs:data-sources:bot_peer_top_good_bots:properties:details", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["details", "self_percentage"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_peer_top_good_bots/properties/details/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Peer Group Top Good Bots Details. Configuration parameter for details", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": [], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# details

Breadcrumbs:

- [xcsh_bot_peer_top_good_bots](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_top_good_bots/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_top_good_bots/properties/)
- details

<a id="section"></a>

Type: `"list"`. Computed.

Peer Group Top Good Bots Details. Configuration parameter for details

## Direct properties

<a id="schema-details--name"></a>

### name property

Type: `"string"`. Computed.

Name. The Name of Item.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

<a id="schema-details--peer_count"></a>

### peer_count property

Type: `"string"`. Computed.

Peer Count. The count of Peer of the Item.

<a id="schema-details--peer_percentage"></a>

### peer_percentage property

Type: `"number"`. Computed.

Peer Percentage. The Peer Percentage of the Item.

<a id="schema-details--self_count"></a>

### self_count property

Type: `"string"`. Computed.

Self Count. The count of Self of the item.

<a id="schema-details--self_percentage"></a>

### self_percentage property

Type: `"number"`. Computed.

Self Percentage. The Self Percentage of the Item.
