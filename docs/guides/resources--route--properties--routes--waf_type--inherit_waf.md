---
page_title: "routes.waf_type.inherit_waf"
subcategory: ""
description: "routes.waf_type.inherit_waf for xcsh_route."
xcsh_docs: {"aliases": [], "body_bytes": 1019, "body_sha256": "sha256:c50f3f0e6ce0f4b30c58e215e8051468776e276e8498c352733f056aa71cd6c4", "canonical_id": "xcsh-docs:resources:route:properties:routes:waf_type:inherit_waf", "child_ids": [], "collection_id": "xcsh-docs:resources:route:collection", "completeness": "complete", "id": "xcsh-docs:resources:route:properties:routes:waf_type:inherit_waf", "parent_id": "xcsh-docs:resources:route:properties:routes:waf_type", "path": "docs/guides/resources--route--properties--routes--waf_type--inherit_waf.md", "provider_name": "route", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes", "waf_type", "inherit_waf"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/route/properties/routes/waf_type/inherit_waf/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.waf_type.inherit_waf for xcsh_route.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["routeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.waf_type.inherit_waf

Breadcrumbs:

- [xcsh_route](../resources/route.md)
- [Property reference](resources--route--reference.md)
- [routes](resources--route--properties--routes.md)
- [routes.waf_type](resources--route--properties--routes--waf_type.md)
- routes.waf_type.inherit_waf

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for inherit waf.

Upstream description:

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
inherit_waf = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [routes.waf_type](resources--route--properties--routes--waf_type.md)
- [xcsh_route](../resources/route.md)
