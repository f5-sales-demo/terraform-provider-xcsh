---
page_title: "routes.route_destination.mirror_policy"
subcategory: ""
description: "routes.route_destination.mirror_policy for xcsh_route."
xcsh_docs: {"aliases": [], "body_bytes": 2283, "body_sha256": "sha256:71a9f7edc86df1d6247013821b5b76d393d69f85cfaf4a80a90b0fe43b0364e0", "canonical_id": "xcsh-docs:resources:route:properties:routes:route_destination:mirror_policy", "child_ids": ["xcsh-docs:resources:route:properties:routes:route_destination:mirror_policy:cluster", "xcsh-docs:resources:route:properties:routes:route_destination:mirror_policy:percent"], "collection_id": "xcsh-docs:resources:route:collection", "completeness": "complete", "id": "xcsh-docs:resources:route:properties:routes:route_destination:mirror_policy", "parent_id": "xcsh-docs:resources:route:properties:routes:route_destination", "path": "docs/guides/resources--route--properties--routes--route_destination--mirror_policy.md", "provider_name": "route", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes", "route_destination", "mirror_policy"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/route/properties/routes/route_destination/mirror_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.route_destination.mirror_policy for xcsh_route.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["routeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.route_destination.mirror_policy

Breadcrumbs:

- [xcsh_route](../resources/route.md)
- [Property reference](resources--route--reference.md)
- [routes](resources--route--properties--routes.md)
- [routes.route_destination](resources--route--properties--routes--route_destination.md)
- routes.route_destination.mirror_policy

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

MirrorPolicy is used for shadowing traffic from one cluster to another. The approach used is 'fire
and forget', meaning it will not wait for the shadow cluster to respond before returning the
response from the primary cluster. All normal statistics are collected for the shadow cluster
making..

Upstream description:

MirrorPolicy is used for shadowing traffic from one cluster to another. The approach used is "fire
and forget", meaning it will not wait for the shadow cluster to respond before returning the
response from the primary cluster. All normal statistics are collected for the shadow cluster making
this feature useful for testing and troubleshooting.

Provider validators and defaults (from schema source):

```go
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

- [cluster](resources--route--properties--routes--route_destination--mirror_policy--cluster.md): complete subsection reference.

- [percent](resources--route--properties--routes--route_destination--mirror_policy--percent.md): complete subsection reference.

## Next pages

- [routes.route_destination.mirror_policy.cluster](resources--route--properties--routes--route_destination--mirror_policy--cluster.md)
- [routes.route_destination.mirror_policy.percent](resources--route--properties--routes--route_destination--mirror_policy--percent.md)
- [routes.route_destination](resources--route--properties--routes--route_destination.md)
- [xcsh_route](../resources/route.md)
