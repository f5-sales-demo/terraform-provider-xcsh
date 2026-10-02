---
page_title: "details"
subcategory: ""
description: "Peer Group Top Good Bots Details. Configuration parameter for details"
xcsh_docs: {"aliases": ["details"], "body_bytes": 1731, "body_sha256": "sha256:132138824771d7ba0a192cf12bb00bac562ed5e4a83c041b1c749a442af54b6d", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_peer_top_good_bots:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_peer_top_good_bots:properties:details", "parent_id": "xcsh-docs:data-sources:bot_peer_top_good_bots:reference", "path": "documentation/data-sources/bot_peer_top_good_bots/properties/details/index.md", "product": "distributed-cloud", "provider_name": "bot_peer_top_good_bots", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0120211233303301-3121333012213312-3123002231120022-2000122121311231-0211210020111230-2210210002222200-3020330210302102-3310311130100121", "registry_path": "docs/guides/data-sources--bot_peer_top_good_bots--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["details"], "schema_version": 1, "sections": [{"aliases": ["name"], "anchor": "schema-details--name", "description": "Name. The Name of Item.", "document_id": "xcsh-docs:data-sources:bot_peer_top_good_bots:properties:details", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["details", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["peer count"], "anchor": "schema-details--peer_count", "description": "Peer Count. The count of Peer of the Item.", "document_id": "xcsh-docs:data-sources:bot_peer_top_good_bots:properties:details", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["details", "peer_count"], "syntax": "attribute", "type": "string"}, {"aliases": ["peer percentage"], "anchor": "schema-details--peer_percentage", "description": "Peer Percentage. The Peer Percentage of the Item.", "document_id": "xcsh-docs:data-sources:bot_peer_top_good_bots:properties:details", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["details", "peer_percentage"], "syntax": "attribute", "type": "number"}, {"aliases": ["self count"], "anchor": "schema-details--self_count", "description": "Self Count. The count of Self of the item.", "document_id": "xcsh-docs:data-sources:bot_peer_top_good_bots:properties:details", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["details", "self_count"], "syntax": "attribute", "type": "string"}, {"aliases": ["self percentage"], "anchor": "schema-details--self_percentage", "description": "Self Percentage. The Self Percentage of the Item.", "document_id": "xcsh-docs:data-sources:bot_peer_top_good_bots:properties:details", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["details", "self_percentage"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_peer_top_good_bots/properties/details/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Peer Group Top Good Bots Details. Configuration parameter for details", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_top_good_bots/properties/)
- [xcsh_bot_peer_top_good_bots](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_top_good_bots/)
