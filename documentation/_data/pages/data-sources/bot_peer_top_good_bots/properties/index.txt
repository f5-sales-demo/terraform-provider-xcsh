---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_bot_peer_top_good_bots."
xcsh_docs: {"aliases": ["bot peer top good bots"], "body_bytes": 4035, "body_sha256": "sha256:86f62fbd6b5c0da70a5868fa94c58b81e59b465ee5d1c6df4b2012f828097659", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:bot_peer_top_good_bots:properties:details"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_peer_top_good_bots:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_peer_top_good_bots:reference", "parent_id": "xcsh-docs:data-sources:bot_peer_top_good_bots:fundamentals", "path": "documentation/data-sources/bot_peer_top_good_bots/properties/index.md", "product": "distributed-cloud", "provider_name": "bot_peer_top_good_bots", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1222302012121320-3310003310000213-1123332130111312-0011233021231032-0101303313230300-0003312003231211-2302311133103331-3323312201112032", "registry_path": "docs/guides/data-sources--bot_peer_top_good_bots--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["details"], "anchor": "section", "description": "Peer Group Top Good Bots Details. Configuration parameter for details", "document_id": "xcsh-docs:data-sources:bot_peer_top_good_bots:properties:details", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["details"], "syntax": "attribute", "type": "object"}, {"aliases": ["end time"], "anchor": "schema-end_time", "description": "End Time. End time of the query period.", "document_id": "xcsh-docs:data-sources:bot_peer_top_good_bots:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["end_time"], "syntax": "attribute", "type": "string"}, {"aliases": ["limit"], "anchor": "schema-limit", "description": "Limits the number of transactions returned in the response Optional: If not specified (with default value 0), all transactions that match the query will be returned in the response.", "document_id": "xcsh-docs:data-sources:bot_peer_top_good_bots:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["limit"], "syntax": "attribute", "type": "number"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "Namespace. namespace is used to scope the query. Only virtual_host in given namespace will be considered.", "document_id": "xcsh-docs:data-sources:bot_peer_top_good_bots:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["rank by"], "anchor": "schema-rank_by", "description": "Or 'PEER' Key for ranking - SELF: Rank by self Use Self as key to query Use Peer as key to query. Possible values are `SELF`, `PEER`. Defaults to `SELF`.", "document_id": "xcsh-docs:data-sources:bot_peer_top_good_bots:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rank_by"], "syntax": "attribute", "type": "string"}, {"aliases": ["start time"], "anchor": "schema-start_time", "description": "Start Time. Start time of the query period.", "document_id": "xcsh-docs:data-sources:bot_peer_top_good_bots:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["start_time"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_peer_top_good_bots/properties/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Property reference for xcsh_bot_peer_top_good_bots.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_bot_peer_top_good_bots](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_top_good_bots/)
- Property reference

## Direct properties

- [details](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_top_good_bots/properties/details/): complete subsection reference.

<a id="schema-end_time"></a>

### end_time property

Type: `"string"`. Optional.

End Time. End time of the query period.

<a id="schema-limit"></a>

### limit property

Type: `"number"`. Optional.

Limits the number of transactions returned in the response Optional: If not specified (with default
value 0), all transactions that match the query will be returned in the response.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 500),
}
```

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace. namespace is used to scope the query. Only virtual\_host in given namespace will be
considered.

<a id="schema-rank_by"></a>

### rank_by property

Type: `"string"`. Optional.

\[Enum: SELF|PEER\] Or 'PEER' Key for ranking - SELF: Rank by self Use Self as key to query Use Peer
as key to query. Possible values are \`SELF\`, \`PEER\`. Defaults to \`SELF\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("SELF",
    "PEER"),
}
```

<a id="schema-start_time"></a>

### start_time property

Type: `"string"`. Optional.

Start Time. Start time of the query period.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `details` | [details](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_top_good_bots/properties/details/#section) |
| `details.name` | [details.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_top_good_bots/properties/details/#schema-details--name) |
| `details.peer_count` | [details.peer_count](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_top_good_bots/properties/details/#schema-details--peer_count) |
| `details.peer_percentage` | [details.peer_percentage](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_top_good_bots/properties/details/#schema-details--peer_percentage) |
| `details.self_count` | [details.self_count](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_top_good_bots/properties/details/#schema-details--self_count) |
| `details.self_percentage` | [details.self_percentage](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_top_good_bots/properties/details/#schema-details--self_percentage) |
| `end_time` | [end_time](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_top_good_bots/properties/#schema-end_time) |
| `limit` | [limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_top_good_bots/properties/#schema-limit) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_top_good_bots/properties/#schema-namespace) |
| `rank_by` | [rank_by](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_top_good_bots/properties/#schema-rank_by) |
| `start_time` | [start_time](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_top_good_bots/properties/#schema-start_time) |

## Next pages

- [details](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_top_good_bots/properties/details/)
- [xcsh_bot_peer_top_good_bots](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_top_good_bots/)
