---
page_title: "caching_policy.custom_cache_rule"
subcategory: "Load Balancing"
description: "Caching policies for CDN."
xcsh_docs: {"aliases": ["caching policy custom cache rule"], "body_bytes": 1168, "body_sha256": "sha256:53ef68ebbcf8848f4ce496d70d7ca65b221d5a2d0d2cde7577b649f7bd8e1e8a", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:caching_policy:custom_cache_rule:cdn_cache_rules"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:caching_policy:custom_cache_rule", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:caching_policy", "path": "documentation/resources/http_loadbalancer/properties/caching_policy/custom_cache_rule/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-1112323230330001-3210133301302130-1001103132120102-1113111200333202-2010211133110132-3203313202133021-1030000033033020-1212102233212101", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-013.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["caching_policy", "custom_cache_rule"], "schema_version": 1, "sections": [{"aliases": ["caching policy custom cache rule cdn cache rules"], "anchor": "section", "description": "Reference to CDN Cache Rule configuration object.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:caching_policy:custom_cache_rule:cdn_cache_rules", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-caching_policy--custom_cache_rule--cdn_cache_rules--name", "enforcement": "provider-schema", "group": "caching_policy.custom_cache_rule.cdn_cache_rules:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:caching_policy:custom_cache_rule:cdn_cache_rules", "type": "requires"}], "schema_path": ["caching_policy", "custom_cache_rule", "cdn_cache_rules"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/caching_policy/custom_cache_rule/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Caching policies for CDN.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# caching_policy.custom_cache_rule

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [caching_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/caching_policy/)
- caching_policy.custom_cache_rule

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
custom_cache_rule {
  # Configure direct properties listed below.
}
```

## Direct properties

- [cdn_cache_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/caching_policy/custom_cache_rule/cdn_cache_rules/): complete subsection reference.
