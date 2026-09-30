---
page_title: "routes.route_destination.query_params.remove_all_params"
subcategory: ""
description: "routes.route_destination.query_params.remove_all_params for xcsh_route."
xcsh_docs: {"aliases": [], "body_bytes": 1167, "body_sha256": "sha256:1bb2d494dd120009d80340af9f3a4f55d361d433e7db1d2aa8035b3390af82dc", "canonical_id": "xcsh-docs:resources:route:properties:routes:route_destination:query_params:remove_all_params", "child_ids": [], "collection_id": "xcsh-docs:resources:route:collection", "completeness": "complete", "id": "xcsh-docs:resources:route:properties:routes:route_destination:query_params:remove_all_params", "parent_id": "xcsh-docs:resources:route:properties:routes:route_destination:query_params", "path": "docs/guides/resources--route--properties--routes--route_destination--query_params--remove_all_params.md", "provider_name": "route", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes", "route_destination", "query_params", "remove_all_params"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/route/properties/routes/route_destination/query_params/remove_all_params/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.route_destination.query_params.remove_all_params for xcsh_route.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["routeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# routes.route_destination.query_params.remove_all_params

Breadcrumbs:

- [xcsh_route](../resources/route.md)
- [Property reference](resources--route--reference.md)
- [routes](resources--route--properties--routes.md)
- [routes.route_destination](resources--route--properties--routes--route_destination.md)
- [routes.route_destination.query_params](resources--route--properties--routes--route_destination--query_params.md)
- routes.route_destination.query_params.remove_all_params

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for remove all params.

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
remove_all_params = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [routes.route_destination.query_params](resources--route--properties--routes--route_destination--query_params.md)
- [xcsh_route](../resources/route.md)
