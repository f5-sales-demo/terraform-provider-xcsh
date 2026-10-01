---
page_title: "details"
subcategory: ""
description: "details for xcsh_bot_peer_traffic_overview."
xcsh_docs: {"aliases": [], "body_bytes": 1752, "body_sha256": "sha256:07c1fd4e70eba12f031153df1cd2b6b0ebfe0efd06ce6d569f0f61adf75ab43b", "child_ids": [], "collection_id": "xcsh-docs:data-sources:bot_peer_traffic_overview:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_peer_traffic_overview:properties:details", "parent_id": "xcsh-docs:data-sources:bot_peer_traffic_overview:reference", "path": "documentation/data-sources/bot_peer_traffic_overview/properties/details/index.md", "provider_name": "bot_peer_traffic_overview", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["details"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_peer_traffic_overview/properties/details/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "details for xcsh_bot_peer_traffic_overview.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
