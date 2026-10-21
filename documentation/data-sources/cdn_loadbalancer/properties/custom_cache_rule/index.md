---
page_title: "custom_cache_rule"
subcategory: "Load Balancing"
description: "Caching policies for CDN."
xcsh_docs: {"aliases": ["custom cache rule"], "body_bytes": 871, "body_sha256": "sha256:009ac07cda64c6df092b95200058cd827502e39f262158424c563396185d6753", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:custom_cache_rule:cdn_cache_rules"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:custom_cache_rule", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:reference", "path": "documentation/data-sources/cdn_loadbalancer/properties/custom_cache_rule/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-0012310100301121-0211013102110231-1333303221223123-0323131031032113-2122320213230021-2211113110211001-3231311331000132-0313231022331320", "registry_path": "docs/guides/data-sources--cdn_loadbalancer--reference--group-010.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_cache_rule"], "schema_version": 1, "sections": [{"aliases": ["custom cache rule cdn cache rules"], "anchor": "section", "description": "Reference to CDN Cache Rule configuration object.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:custom_cache_rule:cdn_cache_rules", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["custom_cache_rule", "cdn_cache_rules"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/custom_cache_rule/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "Caching policies for CDN.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_cache_rule

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- custom_cache_rule

<a id="section"></a>

Type: `"single"`. Computed.

Custom Cache Rules. Caching policies for CDN.

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

- [cdn_cache_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/custom_cache_rule/cdn_cache_rules/): complete subsection reference.
