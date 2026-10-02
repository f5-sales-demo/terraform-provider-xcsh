---
page_title: "details"
subcategory: ""
description: "Peer Group Top Good Bots Details. Configuration parameter for details"
xcsh_docs: {"aliases": ["details"], "body_bytes": 1749, "body_sha256": "sha256:21aa4781e924694d7ddabdef1d778d3440631f89686d56369377e5bd30656f5b", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_peer_top_reason_codes:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_peer_top_reason_codes:properties:details", "parent_id": "xcsh-docs:data-sources:bot_peer_top_reason_codes:reference", "path": "documentation/data-sources/bot_peer_top_reason_codes/properties/details/index.md", "product": "distributed-cloud", "provider_name": "bot_peer_top_reason_codes", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0203221032032130-2023312202213021-3222012212023301-1230333201122110-0032211322322231-3231233000301113-2020320102221201-3210321330131002", "registry_path": "docs/guides/data-sources--bot_peer_top_reason_codes--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["details"], "schema_version": 1, "sections": [{"aliases": ["name"], "anchor": "schema-details--name", "description": "Name. The Name of Item.", "document_id": "xcsh-docs:data-sources:bot_peer_top_reason_codes:properties:details", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["details", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["peer count"], "anchor": "schema-details--peer_count", "description": "Peer Count. The count of Peer of the Item.", "document_id": "xcsh-docs:data-sources:bot_peer_top_reason_codes:properties:details", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["details", "peer_count"], "syntax": "attribute", "type": "string"}, {"aliases": ["peer percentage"], "anchor": "schema-details--peer_percentage", "description": "Peer Percentage. The Peer Percentage of the Item.", "document_id": "xcsh-docs:data-sources:bot_peer_top_reason_codes:properties:details", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["details", "peer_percentage"], "syntax": "attribute", "type": "number"}, {"aliases": ["self count"], "anchor": "schema-details--self_count", "description": "Self Count. The count of Self of the item.", "document_id": "xcsh-docs:data-sources:bot_peer_top_reason_codes:properties:details", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["details", "self_count"], "syntax": "attribute", "type": "string"}, {"aliases": ["self percentage"], "anchor": "schema-details--self_percentage", "description": "Self Percentage. The Self Percentage of the Item.", "document_id": "xcsh-docs:data-sources:bot_peer_top_reason_codes:properties:details", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["details", "self_percentage"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_peer_top_reason_codes/properties/details/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Peer Group Top Good Bots Details. Configuration parameter for details", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_top_reason_codes/properties/)
- [xcsh_bot_peer_top_reason_codes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_top_reason_codes/)
