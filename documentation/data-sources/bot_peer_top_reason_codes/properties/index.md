---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_bot_peer_top_reason_codes."
xcsh_docs: {"aliases": [], "body_bytes": 4086, "body_sha256": "sha256:b1657ea4a10ad609862d879d6f393de34cd242765758ce2b471be0e899ad6889", "child_ids": ["xcsh-docs:data-sources:bot_peer_top_reason_codes:properties:details"], "collection_id": "xcsh-docs:data-sources:bot_peer_top_reason_codes:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_peer_top_reason_codes:reference", "parent_id": "xcsh-docs:data-sources:bot_peer_top_reason_codes:fundamentals", "path": "documentation/data-sources/bot_peer_top_reason_codes/properties/index.md", "provider_name": "bot_peer_top_reason_codes", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_peer_top_reason_codes/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_bot_peer_top_reason_codes.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_bot_peer_top_reason_codes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_top_reason_codes/)
- Property reference

## Direct properties

- [details](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_top_reason_codes/properties/details/): complete subsection reference.

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
| `details` | [details](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_top_reason_codes/properties/details/#section) |
| `details.name` | [details.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_top_reason_codes/properties/details/#schema-details--name) |
| `details.peer_count` | [details.peer_count](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_top_reason_codes/properties/details/#schema-details--peer_count) |
| `details.peer_percentage` | [details.peer_percentage](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_top_reason_codes/properties/details/#schema-details--peer_percentage) |
| `details.self_count` | [details.self_count](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_top_reason_codes/properties/details/#schema-details--self_count) |
| `details.self_percentage` | [details.self_percentage](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_top_reason_codes/properties/details/#schema-details--self_percentage) |
| `end_time` | [end_time](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_top_reason_codes/properties/#schema-end_time) |
| `limit` | [limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_top_reason_codes/properties/#schema-limit) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_top_reason_codes/properties/#schema-namespace) |
| `rank_by` | [rank_by](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_top_reason_codes/properties/#schema-rank_by) |
| `start_time` | [start_time](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_top_reason_codes/properties/#schema-start_time) |

## Next pages

- [details](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_top_reason_codes/properties/details/)
- [xcsh_bot_peer_top_reason_codes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_top_reason_codes/)
