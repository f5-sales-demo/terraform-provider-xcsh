---
page_title: "details"
subcategory: ""
description: "Peer Group Top Good Bots Details. Configuration parameter for details"
xcsh_docs: {"aliases": ["details"], "body_bytes": 1500, "body_sha256": "sha256:abd6fb1542bcfac2f90dc4f4c04fd824e3a9db6533e958116ec9858f8e325e6b", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_peer_top_reason_codes:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_peer_top_reason_codes:properties:details", "parent_id": "xcsh-docs:data-sources:bot_peer_top_reason_codes:reference", "path": "documentation/data-sources/bot_peer_top_reason_codes/properties/details/index.md", "product": "distributed-cloud", "provider_name": "bot_peer_top_reason_codes", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "data-sources", "registry_anchor": "canonical-0203221032032130-2023312202213021-3222012212023301-1230333201122110-0032211322322231-3231233000301113-2020320102221201-3210321330131002", "registry_path": "docs/guides/data-sources--bot_peer_top_reason_codes--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["details"], "schema_version": 1, "sections": [{"aliases": ["details name"], "anchor": "schema-details--name", "description": "Name. The Name of Item.", "document_id": "xcsh-docs:data-sources:bot_peer_top_reason_codes:properties:details", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["details", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["details peer count"], "anchor": "schema-details--peer_count", "description": "Peer Count. The count of Peer of the Item.", "document_id": "xcsh-docs:data-sources:bot_peer_top_reason_codes:properties:details", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["details", "peer_count"], "syntax": "attribute", "type": "string"}, {"aliases": ["details peer percentage"], "anchor": "schema-details--peer_percentage", "description": "Peer Percentage. The Peer Percentage of the Item.", "document_id": "xcsh-docs:data-sources:bot_peer_top_reason_codes:properties:details", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["details", "peer_percentage"], "syntax": "attribute", "type": "number"}, {"aliases": ["details self count"], "anchor": "schema-details--self_count", "description": "Self Count. The count of Self of the item.", "document_id": "xcsh-docs:data-sources:bot_peer_top_reason_codes:properties:details", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["details", "self_count"], "syntax": "attribute", "type": "string"}, {"aliases": ["details self percentage"], "anchor": "schema-details--self_percentage", "description": "Self Percentage. The Self Percentage of the Item.", "document_id": "xcsh-docs:data-sources:bot_peer_top_reason_codes:properties:details", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["details", "self_percentage"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_peer_top_reason_codes/properties/details/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "Peer Group Top Good Bots Details. Configuration parameter for details", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": [], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# details

Breadcrumbs:

- [xcsh_bot_peer_top_reason_codes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_top_reason_codes/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_top_reason_codes/properties/)
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
