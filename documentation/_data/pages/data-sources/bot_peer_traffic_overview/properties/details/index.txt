---
page_title: "details"
subcategory: ""
description: "Peer Group Traffic Overview Details. Configuration parameter for details"
xcsh_docs: {"aliases": ["details"], "body_bytes": 1752, "body_sha256": "sha256:07c1fd4e70eba12f031153df1cd2b6b0ebfe0efd06ce6d569f0f61adf75ab43b", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_peer_traffic_overview:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_peer_traffic_overview:properties:details", "parent_id": "xcsh-docs:data-sources:bot_peer_traffic_overview:reference", "path": "documentation/data-sources/bot_peer_traffic_overview/properties/details/index.md", "product": "distributed-cloud", "provider_name": "bot_peer_traffic_overview", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "data-sources", "registry_anchor": "canonical-0333331230020000-0300110232113000-3213121023333031-3020213011130113-1010110023023203-0313001113312222-1212322132313021-0110023331001001", "registry_path": "docs/guides/data-sources--bot_peer_traffic_overview--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["details"], "schema_version": 1, "sections": [{"aliases": ["name"], "anchor": "schema-details--name", "description": "Name. The Name of Item.", "document_id": "xcsh-docs:data-sources:bot_peer_traffic_overview:properties:details", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["details", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["peer count"], "anchor": "schema-details--peer_count", "description": "Peer Count. The count of Peer of the Item.", "document_id": "xcsh-docs:data-sources:bot_peer_traffic_overview:properties:details", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["details", "peer_count"], "syntax": "attribute", "type": "string"}, {"aliases": ["peer percentage"], "anchor": "schema-details--peer_percentage", "description": "Peer Percentage. The Peer Percentage of the Item.", "document_id": "xcsh-docs:data-sources:bot_peer_traffic_overview:properties:details", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["details", "peer_percentage"], "syntax": "attribute", "type": "number"}, {"aliases": ["self count"], "anchor": "schema-details--self_count", "description": "Self Count. The count of Self of the item.", "document_id": "xcsh-docs:data-sources:bot_peer_traffic_overview:properties:details", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["details", "self_count"], "syntax": "attribute", "type": "string"}, {"aliases": ["self percentage"], "anchor": "schema-details--self_percentage", "description": "Self Percentage. The Self Percentage of the Item.", "document_id": "xcsh-docs:data-sources:bot_peer_traffic_overview:properties:details", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["details", "self_percentage"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_peer_traffic_overview/properties/details/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Peer Group Traffic Overview Details. Configuration parameter for details", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# details

Breadcrumbs:

- [xcsh_bot_peer_traffic_overview](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_traffic_overview/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_traffic_overview/properties/)
- details

<a id="section"></a>

Type: `"list"`. Computed.

Peer Group Traffic Overview Details. Configuration parameter for details

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

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_traffic_overview/properties/)
- [xcsh_bot_peer_traffic_overview](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_traffic_overview/)
