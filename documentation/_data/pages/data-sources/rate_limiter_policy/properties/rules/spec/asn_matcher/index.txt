---
page_title: "rules.spec.asn_matcher"
subcategory: "Security"
description: "Match any AS number contained in the list of bgp_asn_sets."
xcsh_docs: {"aliases": ["rules spec asn matcher"], "body_bytes": 1579, "body_sha256": "sha256:5c8107b89788aafec8214305acb7f37d216b0800b127c3069797783711e8123c", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:asn_matcher:asn_sets"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:rate_limiter_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:asn_matcher", "parent_id": "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec", "path": "documentation/data-sources/rate_limiter_policy/properties/rules/spec/asn_matcher/index.md", "product": "distributed-cloud", "provider_name": "rate_limiter_policy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3002221121111303-2331202203130210-1332331020200211-3010112211110012-0032023322323303-2113013110212101-2120023130220301-0322213133032000", "registry_path": "docs/guides/data-sources--rate_limiter_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "spec", "asn_matcher"], "schema_version": 1, "sections": [{"aliases": ["rules spec asn matcher asn sets"], "anchor": "section", "description": "A list of references to bgp_asn_set objects.", "document_id": "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:asn_matcher:asn_sets", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["rules", "spec", "asn_matcher", "asn_sets"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/rate_limiter_policy/properties/rules/spec/asn_matcher/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Match any AS number contained in the list of bgp_asn_sets.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["rate_limiter_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.spec.asn_matcher

Breadcrumbs:

- [xcsh_rate_limiter_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/)
- [rules.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/)
- rules.spec.asn_matcher

<a id="section"></a>

Type: `"single"`. Computed.

Match any AS number contained in the list of bgp\_asn\_sets.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Direct properties

- [asn_sets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/asn_matcher/asn_sets/): complete subsection reference.

## Next pages

- [rules.spec.asn_matcher.asn_sets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/asn_matcher/asn_sets/)
- [rules.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/)
- [xcsh_rate_limiter_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/)
