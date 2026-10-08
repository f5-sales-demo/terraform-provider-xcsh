---
page_title: "routes.route_destination.mirror_policy"
subcategory: ""
description: "MirrorPolicy is used for shadowing traffic from one cluster to another. The approach used is \"fire and forget\", meaning it will not wait for the shadow cluster to respond before returning the response from the primary cluster. All normal statistics are collected for the shadow cluster making this feature useful for"
xcsh_docs: {"aliases": ["routes route destination mirror policy"], "body_bytes": 2136, "body_sha256": "sha256:04b2db213076ed6bee700b8ec8679175c96d2710865e413cfe17624083f24fa8", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:route:properties:routes:route_destination:mirror_policy:cluster", "xcsh-docs:resources:route:properties:routes:route_destination:mirror_policy:percent"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:route:collection", "completeness": "complete", "id": "xcsh-docs:resources:route:properties:routes:route_destination:mirror_policy", "parent_id": "xcsh-docs:resources:route:properties:routes:route_destination", "path": "documentation/resources/route/properties/routes/route_destination/mirror_policy/index.md", "product": "distributed-cloud", "provider_name": "route", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-1312110103201222-3211323013120213-1220232003210122-1333011320200321-0302110311030231-3120313321230312-0300310313300210-1231213001201021", "registry_path": "docs/guides/resources--route--reference--group-003.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "routes.route_destination.mirror_policy:RequiredObjectAttributes:cluster", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:mirror_policy:cluster", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "route_destination", "mirror_policy"], "schema_version": 1, "sections": [{"aliases": ["routes route destination mirror policy cluster"], "anchor": "section", "description": "Specifies the cluster to which the requests will be mirrored. The cluster object referred here must be present.", "document_id": "xcsh-docs:resources:route:properties:routes:route_destination:mirror_policy:cluster", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["routes", "route_destination", "mirror_policy", "cluster"], "syntax": "block", "type": "object"}, {"aliases": ["routes route destination mirror policy percent"], "anchor": "section", "description": "Fraction used where sampling percentages are needed. Example sampled requests.", "document_id": "xcsh-docs:resources:route:properties:routes:route_destination:mirror_policy:percent", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-routes--route_destination--mirror_policy--percent--numerator", "enforcement": "provider-schema", "group": "routes.route_destination.mirror_policy.percent:RequiredObjectAttributes:numerator", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:mirror_policy:percent", "type": "requires"}], "schema_path": ["routes", "route_destination", "mirror_policy", "percent"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/route/properties/routes/route_destination/mirror_policy/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "MirrorPolicy is used for shadowing traffic from one cluster to another. The approach used is \"fire and forget\", meaning it will not wait for the shadow cluster to respond before returning the response from the primary cluster. All normal statistics are collected for the shadow cluster making this feature useful for", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["routeCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.route_destination.mirror_policy

Breadcrumbs:

- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/)
- [routes.route_destination](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/)
- routes.route_destination.mirror_policy

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

MirrorPolicy is used for shadowing traffic from one cluster to another. The approach used is 'fire
and forget', meaning it will not wait for the shadow cluster to respond before returning the
response from the primary cluster. All normal statistics are collected for the shadow cluster
making..

Additional upstream details:

The approach used is "fire and forget", meaning it will not wait for the shadow cluster to respond
before returning the response from the primary cluster. All normal statistics are collected for the
shadow cluster making this feature useful for testing and troubleshooting.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("cluster")}
```

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
mirror_policy {
  # Configure direct properties listed below.
}
```

## Direct properties

- [cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/mirror_policy/cluster/): complete subsection reference.

- [percent](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/mirror_policy/percent/): complete subsection reference.
