---
page_title: "details"
subcategory: ""
description: "details for xcsh_bot_peer_top_reason_codes."
xcsh_docs: {"aliases": [], "body_bytes": 1541, "body_sha256": "sha256:fefd18f39412c3423496801a4b13045e95d857306f9e83e16187189999670ed3", "canonical_id": "xcsh-docs:data-sources:bot_peer_top_reason_codes:properties:details", "child_ids": [], "collection_id": "xcsh-docs:data-sources:bot_peer_top_reason_codes:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_peer_top_reason_codes:properties:details", "parent_id": "xcsh-docs:data-sources:bot_peer_top_reason_codes:reference", "path": "docs/guides/data-sources--bot_peer_top_reason_codes--properties--details.md", "provider_name": "bot_peer_top_reason_codes", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["details"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_peer_top_reason_codes/properties/details/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "details for xcsh_bot_peer_top_reason_codes.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# details

Breadcrumbs:

- [xcsh_bot_peer_top_reason_codes](../data-sources/bot_peer_top_reason_codes.md)
- [Property reference](data-sources--bot_peer_top_reason_codes--reference.md)
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

- [Property reference](data-sources--bot_peer_top_reason_codes--reference.md)
- [xcsh_bot_peer_top_reason_codes](../data-sources/bot_peer_top_reason_codes.md)
