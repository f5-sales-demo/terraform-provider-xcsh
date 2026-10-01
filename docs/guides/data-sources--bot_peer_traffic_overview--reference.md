---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_bot_peer_traffic_overview."
xcsh_docs: {"aliases": [], "body_bytes": 4329, "body_sha256": "sha256:e53bc570060e2c07b1057bbceb900b578edd2c5dcf8dc7d5bf27834423d21145", "canonical_id": "xcsh-docs:data-sources:bot_peer_traffic_overview:reference", "child_ids": ["xcsh-docs:data-sources:bot_peer_traffic_overview:properties:details"], "collection_id": "xcsh-docs:data-sources:bot_peer_traffic_overview:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_peer_traffic_overview:reference", "parent_id": "xcsh-docs:data-sources:bot_peer_traffic_overview:fundamentals", "path": "docs/guides/data-sources--bot_peer_traffic_overview--reference.md", "provider_name": "bot_peer_traffic_overview", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_peer_traffic_overview/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_bot_peer_traffic_overview.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_bot_peer_traffic_overview](../data-sources/bot_peer_traffic_overview.md)
- Property reference

## Direct properties

- [details](data-sources--bot_peer_traffic_overview--properties--details.md): complete subsection reference.

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
| `details` | [details](data-sources--bot_peer_traffic_overview--properties--details.md#section) |
| `details.name` | [details.name](data-sources--bot_peer_traffic_overview--properties--details.md#schema-details--name) |
| `details.peer_count` | [details.peer_count](data-sources--bot_peer_traffic_overview--properties--details.md#schema-details--peer_count) |
| `details.peer_percentage` | [details.peer_percentage](data-sources--bot_peer_traffic_overview--properties--details.md#schema-details--peer_percentage) |
| `details.self_count` | [details.self_count](data-sources--bot_peer_traffic_overview--properties--details.md#schema-details--self_count) |
| `details.self_percentage` | [details.self_percentage](data-sources--bot_peer_traffic_overview--properties--details.md#schema-details--self_percentage) |
| `end_time` | [end_time](data-sources--bot_peer_traffic_overview--reference.md#schema-end_time) |
| `limit` | [limit](data-sources--bot_peer_traffic_overview--reference.md#schema-limit) |
| `namespace` | [namespace](data-sources--bot_peer_traffic_overview--reference.md#schema-namespace) |
| `peer_percentage` | [peer_percentage](data-sources--bot_peer_traffic_overview--reference.md#schema-peer_percentage) |
| `peer_total` | [peer_total](data-sources--bot_peer_traffic_overview--reference.md#schema-peer_total) |
| `rank_by` | [rank_by](data-sources--bot_peer_traffic_overview--reference.md#schema-rank_by) |
| `self_percentage` | [self_percentage](data-sources--bot_peer_traffic_overview--reference.md#schema-self_percentage) |
| `self_total` | [self_total](data-sources--bot_peer_traffic_overview--reference.md#schema-self_total) |
| `start_time` | [start_time](data-sources--bot_peer_traffic_overview--reference.md#schema-start_time) |

## Next pages

- [details](data-sources--bot_peer_traffic_overview--properties--details.md)
- [xcsh_bot_peer_traffic_overview](../data-sources/bot_peer_traffic_overview.md)
