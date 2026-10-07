---
page_title: "routes.route_destination.mirror_policy"
subcategory: ""
description: "MirrorPolicy is used for shadowing traffic from one cluster to another. The approach used is \"fire and forget\", meaning it will not wait for the shadow cluster to respond before returning the response from the primary cluster. All normal statistics are collected for the shadow cluster making this feature useful for"
xcsh_docs: {"aliases": ["routes route destination mirror policy"], "body_bytes": 1856, "body_sha256": "sha256:df159bb609f91cea97cd803acfd41b5c4a527d7c4487db53450d604aba3b0959", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:route:properties:routes:route_destination:mirror_policy:cluster", "xcsh-docs:data-sources:route:properties:routes:route_destination:mirror_policy:percent"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:route:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:route:properties:routes:route_destination:mirror_policy", "parent_id": "xcsh-docs:data-sources:route:properties:routes:route_destination", "path": "documentation/data-sources/route/properties/routes/route_destination/mirror_policy/index.md", "product": "distributed-cloud", "provider_name": "route", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-3102020312000110-0133212131123210-2001013313121121-0310130120100121-2021031311110033-1002330012133302-0123002120221110-1203302331112303", "registry_path": "docs/guides/data-sources--route--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "route_destination", "mirror_policy"], "schema_version": 1, "sections": [{"aliases": ["routes route destination mirror policy cluster"], "anchor": "section", "description": "Specifies the cluster to which the requests will be mirrored. The cluster object referred here must be present.", "document_id": "xcsh-docs:data-sources:route:properties:routes:route_destination:mirror_policy:cluster", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["routes", "route_destination", "mirror_policy", "cluster"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes route destination mirror policy percent"], "anchor": "section", "description": "Fraction used where sampling percentages are needed. Example sampled requests.", "document_id": "xcsh-docs:data-sources:route:properties:routes:route_destination:mirror_policy:percent", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "route_destination", "mirror_policy", "percent"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/route/properties/routes/route_destination/mirror_policy/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "MirrorPolicy is used for shadowing traffic from one cluster to another. The approach used is \"fire and forget\", meaning it will not wait for the shadow cluster to respond before returning the response from the primary cluster. All normal statistics are collected for the shadow cluster making this feature useful for", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["routeCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.route_destination.mirror_policy

Breadcrumbs:

- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/)
- [routes.route_destination](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/)
- routes.route_destination.mirror_policy

<a id="section"></a>

Type: `"single"`. Computed.

MirrorPolicy is used for shadowing traffic from one cluster to another. The approach used is 'fire
and forget', meaning it will not wait for the shadow cluster to respond before returning the
response from the primary cluster. All normal statistics are collected for the shadow cluster
making..

Additional upstream details:

The approach used is "fire and forget", meaning it will not wait for the shadow cluster to respond
before returning the response from the primary cluster. All normal statistics are collected for the
shadow cluster making this feature useful for testing and troubleshooting.

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

- [cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/mirror_policy/cluster/): complete subsection reference.

- [percent](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/mirror_policy/percent/): complete subsection reference.
