---
page_title: "routes.route_destination.hash_policy.cookie.ignore_httponly"
subcategory: ""
description: "routes.route_destination.hash_policy.cookie.ignore_httponly for xcsh_route."
xcsh_docs: {"aliases": [], "body_bytes": 1410, "body_sha256": "sha256:02864dcb1ff409aef5aeda8ab79dfed233faecffdc32080223ec7be5a3a525a7", "canonical_id": "xcsh-docs:resources:route:properties:routes:route_destination:hash_policy:cookie:ignore_httponly", "child_ids": [], "collection_id": "xcsh-docs:resources:route:collection", "completeness": "complete", "id": "xcsh-docs:resources:route:properties:routes:route_destination:hash_policy:cookie:ignore_httponly", "parent_id": "xcsh-docs:resources:route:properties:routes:route_destination:hash_policy:cookie", "path": "docs/guides/resources--route--properties--routes--route_destination--hash_policy--cookie--ignore_httponly.md", "provider_name": "route", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes", "route_destination", "hash_policy", "cookie", "ignore_httponly"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/route/properties/routes/route_destination/hash_policy/cookie/ignore_httponly/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.route_destination.hash_policy.cookie.ignore_httponly for xcsh_route.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["routeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.route_destination.hash_policy.cookie.ignore_httponly

Breadcrumbs:

- [xcsh_route](../resources/route.md)
- [Property reference](resources--route--reference.md)
- [routes](resources--route--properties--routes.md)
- [routes.route_destination](resources--route--properties--routes--route_destination.md)
- [routes.route_destination.hash_policy](resources--route--properties--routes--route_destination--hash_policy.md)
- [routes.route_destination.hash_policy.cookie](resources--route--properties--routes--route_destination--hash_policy--cookie.md)
- routes.route_destination.hash_policy.cookie.ignore_httponly

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore httponly.

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
ignore_httponly = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [routes.route_destination.hash_policy.cookie](resources--route--properties--routes--route_destination--hash_policy--cookie.md)
- [xcsh_route](../resources/route.md)
