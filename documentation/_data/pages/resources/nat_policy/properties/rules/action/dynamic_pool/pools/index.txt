---
page_title: "rules.action.dynamic_pool.pools"
subcategory: ""
description: "List of IPv4 prefixes that represent an endpoint."
xcsh_docs: {"aliases": ["rules action dynamic pool pools"], "body_bytes": 1345, "body_sha256": "sha256:67c7f82bd1d37975cb21ab9c00d6eb53e7ae2c29d17c9d88e8711973b5445aa6", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:nat_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:nat_policy:properties:rules:action:dynamic_pool:pools", "parent_id": "xcsh-docs:resources:nat_policy:properties:rules:action:dynamic_pool", "path": "documentation/resources/nat_policy/properties/rules/action/dynamic_pool/pools/index.md", "product": "distributed-cloud", "provider_name": "nat_policy", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "resources", "registry_anchor": "canonical-1203231020021100-0311032330223312-0303233201201121-2330102332101102-2123101221333330-0231110331032012-1021232112103021-3130111102312233", "registry_path": "docs/guides/resources--nat_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "action", "dynamic_pool", "pools"], "schema_version": 1, "sections": [{"aliases": ["rules action dynamic pool pools prefixes"], "anchor": "schema-rules--action--dynamic_pool--pools--prefixes", "description": "List of IPv4 prefixes that represent an endpoint.", "document_id": "xcsh-docs:resources:nat_policy:properties:rules:action:dynamic_pool:pools", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "action", "dynamic_pool", "pools", "prefixes"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nat_policy/properties/rules/action/dynamic_pool/pools/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "List of IPv4 prefixes that represent an endpoint.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["nat_policyCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.action.dynamic_pool.pools

Breadcrumbs:

- [xcsh_nat_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/)
- [rules.action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/action/)
- [rules.action.dynamic_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/action/dynamic_pool/)
- rules.action.dynamic_pool.pools

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

List of IPv4 prefixes that represent an endpoint.

Terraform syntax:

```terraform
pools {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-rules--action--dynamic_pool--pools--prefixes"></a>

### prefixes property

Type: `["list", "string"]`. Optional.

List of IPv4 prefixes that represent an endpoint.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
```
