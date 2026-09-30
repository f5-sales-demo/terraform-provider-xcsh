---
page_title: "routes.inherited_waf_exclusion"
subcategory: ""
description: "routes.inherited_waf_exclusion for xcsh_route."
xcsh_docs: {"aliases": [], "body_bytes": 860, "body_sha256": "sha256:70e73b2313362ca8db68166ab78ee034fa0df87a05179d546abd9325227402d8", "canonical_id": "xcsh-docs:resources:route:properties:routes:inherited_waf_exclusion", "child_ids": [], "collection_id": "xcsh-docs:resources:route:collection", "completeness": "complete", "id": "xcsh-docs:resources:route:properties:routes:inherited_waf_exclusion", "parent_id": "xcsh-docs:resources:route:properties:routes", "path": "docs/guides/resources--route--properties--routes--inherited_waf_exclusion.md", "provider_name": "route", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes", "inherited_waf_exclusion"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/route/properties/routes/inherited_waf_exclusion/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.inherited_waf_exclusion for xcsh_route.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["routeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# routes.inherited_waf_exclusion

Breadcrumbs:

- [xcsh_route](../resources/route.md)
- [Property reference](resources--route--reference.md)
- [routes](resources--route--properties--routes.md)
- routes.inherited_waf_exclusion

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for inherited waf exclusion.

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
inherited_waf_exclusion = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [routes](resources--route--properties--routes.md)
- [xcsh_route](../resources/route.md)
