---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_bot_peer_traffic_overview."
xcsh_docs: {"aliases": ["bot peer traffic overview"], "body_bytes": 5316, "body_sha256": "sha256:4c53e4d66a01f96e24aa8fd240685b986bd261094e1a29fe0b872e7a948561d9", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:bot_peer_traffic_overview:properties:details"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_peer_traffic_overview:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_peer_traffic_overview:reference", "parent_id": "xcsh-docs:data-sources:bot_peer_traffic_overview:fundamentals", "path": "documentation/data-sources/bot_peer_traffic_overview/properties/index.md", "product": "distributed-cloud", "provider_name": "bot_peer_traffic_overview", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1103031213222001-1003122002030331-1121300031120000-1103232121210320-1003000301131102-3132201013033132-1001310031332103-2310031113130322", "registry_path": "docs/guides/data-sources--bot_peer_traffic_overview--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["details"], "anchor": "section", "description": "Peer Group Traffic Overview Details. Configuration parameter for details", "document_id": "xcsh-docs:data-sources:bot_peer_traffic_overview:properties:details", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["details"], "syntax": "attribute", "type": "object"}, {"aliases": ["end time"], "anchor": "schema-end_time", "description": "End Time. End time of the query period.", "document_id": "xcsh-docs:data-sources:bot_peer_traffic_overview:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["end_time"], "syntax": "attribute", "type": "string"}, {"aliases": ["limit"], "anchor": "schema-limit", "description": "Limits the number of transactions returned in the response Optional: If not specified (with default value 0), all transactions that match the query will be returned in the response.", "document_id": "xcsh-docs:data-sources:bot_peer_traffic_overview:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["limit"], "syntax": "attribute", "type": "number"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "Namespace. namespace is used to scope the query. Only virtual_host in given namespace will be considered.", "document_id": "xcsh-docs:data-sources:bot_peer_traffic_overview:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["peer percentage"], "anchor": "schema-peer_percentage", "description": "Peer Percentage. The Peer Percentage of the Item.", "document_id": "xcsh-docs:data-sources:bot_peer_traffic_overview:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["peer_percentage"], "syntax": "attribute", "type": "number"}, {"aliases": ["peer total"], "anchor": "schema-peer_total", "description": "Peer Total. The total number of Peer of the Item.", "document_id": "xcsh-docs:data-sources:bot_peer_traffic_overview:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["peer_total"], "syntax": "attribute", "type": "string"}, {"aliases": ["rank by"], "anchor": "schema-rank_by", "description": "Or 'PEER' Key for ranking - SELF: Rank by self Use Self as key to query Use Peer as key to query. Possible values are `SELF`, `PEER`. Defaults to `SELF`.", "document_id": "xcsh-docs:data-sources:bot_peer_traffic_overview:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rank_by"], "syntax": "attribute", "type": "string"}, {"aliases": ["self percentage"], "anchor": "schema-self_percentage", "description": "Self Percentage. The Self Percentage of the Item.", "document_id": "xcsh-docs:data-sources:bot_peer_traffic_overview:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["self_percentage"], "syntax": "attribute", "type": "number"}, {"aliases": ["self total"], "anchor": "schema-self_total", "description": "Self Total. The total number of Self of the item.", "document_id": "xcsh-docs:data-sources:bot_peer_traffic_overview:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["self_total"], "syntax": "attribute", "type": "string"}, {"aliases": ["start time"], "anchor": "schema-start_time", "description": "Start Time. Start time of the query period.", "document_id": "xcsh-docs:data-sources:bot_peer_traffic_overview:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["start_time"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_peer_traffic_overview/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_bot_peer_traffic_overview.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_bot_peer_traffic_overview](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_traffic_overview/)
- Property reference

## Direct properties

- [details](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_traffic_overview/properties/details/): complete subsection reference.

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

<a id="schema-peer_percentage"></a>

### peer_percentage property

Type: `"number"`. Computed.

Peer Percentage. The Peer Percentage of the Item.

<a id="schema-peer_total"></a>

### peer_total property

Type: `"string"`. Computed.

Peer Total. The total number of Peer of the Item.

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

<a id="schema-self_percentage"></a>

### self_percentage property

Type: `"number"`. Computed.

Self Percentage. The Self Percentage of the Item.

<a id="schema-self_total"></a>

### self_total property

Type: `"string"`. Computed.

Self Total. The total number of Self of the item.

<a id="schema-start_time"></a>

### start_time property

Type: `"string"`. Optional.

Start Time. Start time of the query period.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `details` | [details](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_traffic_overview/properties/details/#section) |
| `details.name` | [details.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_traffic_overview/properties/details/#schema-details--name) |
| `details.peer_count` | [details.peer_count](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_traffic_overview/properties/details/#schema-details--peer_count) |
| `details.peer_percentage` | [details.peer_percentage](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_traffic_overview/properties/details/#schema-details--peer_percentage) |
| `details.self_count` | [details.self_count](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_traffic_overview/properties/details/#schema-details--self_count) |
| `details.self_percentage` | [details.self_percentage](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_traffic_overview/properties/details/#schema-details--self_percentage) |
| `end_time` | [end_time](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_traffic_overview/properties/#schema-end_time) |
| `limit` | [limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_traffic_overview/properties/#schema-limit) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_traffic_overview/properties/#schema-namespace) |
| `peer_percentage` | [peer_percentage](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_traffic_overview/properties/#schema-peer_percentage) |
| `peer_total` | [peer_total](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_traffic_overview/properties/#schema-peer_total) |
| `rank_by` | [rank_by](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_traffic_overview/properties/#schema-rank_by) |
| `self_percentage` | [self_percentage](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_traffic_overview/properties/#schema-self_percentage) |
| `self_total` | [self_total](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_traffic_overview/properties/#schema-self_total) |
| `start_time` | [start_time](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_traffic_overview/properties/#schema-start_time) |

## Next pages

- [details](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_traffic_overview/properties/details/)
- [xcsh_bot_peer_traffic_overview](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_traffic_overview/)
