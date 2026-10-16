---
page_title: "routes.route_destination.hash_policy.cookie.samesite_lax"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["routes route destination hash policy cookie samesite lax"], "body_bytes": 1479, "body_sha256": "sha256:311563bed6b45939290517c8f984dfc42348fd7862ff6d813d03862d0a052a79", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:route:collection", "completeness": "complete", "id": "xcsh-docs:resources:route:properties:routes:route_destination:hash_policy:cookie:samesite_lax", "parent_id": "xcsh-docs:resources:route:properties:routes:route_destination:hash_policy:cookie", "path": "documentation/resources/route/properties/routes/route_destination/hash_policy/cookie/samesite_lax/index.md", "product": "distributed-cloud", "provider_name": "route", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "resources", "registry_anchor": "canonical-0330022013103032-3312123211210022-0220100001032323-1232302301033311-3322102231220223-0003033002123112-0130212200111323-0011313112122003", "registry_path": "docs/guides/resources--route--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "route_destination", "hash_policy", "cookie", "samesite_lax"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/route/properties/routes/route_destination/hash_policy/cookie/samesite_lax/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["routeCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.route_destination.hash_policy.cookie.samesite_lax

Breadcrumbs:

- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/)
- [routes.route_destination](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/)
- [routes.route_destination.hash_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/hash_policy/)
- [routes.route_destination.hash_policy.cookie](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/hash_policy/cookie/)
- routes.route_destination.hash_policy.cookie.samesite_lax

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

This can be used for messages where no values are needed.

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
samesite_lax = {}
```

This is an empty object or choice marker. It has no direct properties.
