---
page_title: "routes.simple_route.advanced_options.mirror_policy"
subcategory: "Load Balancing"
description: "MirrorPolicy is used for shadowing traffic from one origin pool to another. The approach used is \"fire and forget\", meaning it will not wait for the shadow origin pool to respond before returning the response from the primary origin pool. All normal statistics are collected for the shadow origin pool making this"
xcsh_docs: {"aliases": ["routes simple route advanced options mirror policy"], "body_bytes": 2180, "body_sha256": "sha256:827a563cc54ec355842d4c17755773a46aaf270f27cd0f7a1a7e637bf91d9a04", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:mirror_policy:origin_pool", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:mirror_policy:percent"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:mirror_policy", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options", "path": "documentation/data-sources/http_loadbalancer/properties/routes/simple_route/advanced_options/mirror_policy/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-0133131221300311-2302332121101120-1332302111032210-0113112131101132-3210212212011232-1001132300233230-3310112120032213-0122101112323311", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-025.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "simple_route", "advanced_options", "mirror_policy"], "schema_version": 1, "sections": [{"aliases": ["backend servers", "origin servers", "routes simple route advanced options mirror policy origin pool", "upstream servers"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:mirror_policy:origin_pool", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "simple_route", "advanced_options", "mirror_policy", "origin_pool"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes simple route advanced options mirror policy percent"], "anchor": "section", "description": "Fraction used where sampling percentages are needed. Example sampled requests.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:mirror_policy:percent", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "simple_route", "advanced_options", "mirror_policy", "percent"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/routes/simple_route/advanced_options/mirror_policy/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "MirrorPolicy is used for shadowing traffic from one origin pool to another. The approach used is \"fire and forget\", meaning it will not wait for the shadow origin pool to respond before returning the response from the primary origin pool. All normal statistics are collected for the shadow origin pool making this", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.simple_route.advanced_options.mirror_policy

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/)
- [routes.simple_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/simple_route/)
- [routes.simple_route.advanced_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/simple_route/advanced_options/)
- routes.simple_route.advanced_options.mirror_policy

<a id="section"></a>

Type: `"single"`. Computed.

MirrorPolicy is used for shadowing traffic from one origin pool to another. The approach used is
'fire and forget', meaning it will not wait for the shadow origin pool to respond before returning
the response from the primary origin pool. All normal statistics are collected for the shadow
origin..

Additional upstream details:

The approach used is "fire and forget", meaning it will not wait for the shadow origin pool to
respond before returning the response from the primary origin pool. All normal statistics are
collected for the shadow origin pool making this feature useful for testing and troubleshooting.

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

- [origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/simple_route/advanced_options/mirror_policy/origin_pool/): complete subsection reference.

- [percent](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/simple_route/advanced_options/mirror_policy/percent/): complete subsection reference.
