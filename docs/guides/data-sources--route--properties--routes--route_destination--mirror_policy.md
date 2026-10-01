---
page_title: "routes.route_destination.mirror_policy"
subcategory: ""
description: "routes.route_destination.mirror_policy for xcsh_route."
xcsh_docs: {"aliases": [], "body_bytes": 2045, "body_sha256": "sha256:aa61294320c7eeb7b4df946952b5c99ba0ff62bcba7ce4816a4ff9b09e79dacd", "canonical_id": "xcsh-docs:data-sources:route:properties:routes:route_destination:mirror_policy", "child_ids": ["xcsh-docs:data-sources:route:properties:routes:route_destination:mirror_policy:cluster", "xcsh-docs:data-sources:route:properties:routes:route_destination:mirror_policy:percent"], "collection_id": "xcsh-docs:data-sources:route:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:route:properties:routes:route_destination:mirror_policy", "parent_id": "xcsh-docs:data-sources:route:properties:routes:route_destination", "path": "docs/guides/data-sources--route--properties--routes--route_destination--mirror_policy.md", "provider_name": "route", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes", "route_destination", "mirror_policy"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/route/properties/routes/route_destination/mirror_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.route_destination.mirror_policy for xcsh_route.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["routeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.route_destination.mirror_policy

Breadcrumbs:

- [xcsh_route](../data-sources/route.md)
- [Property reference](data-sources--route--reference.md)
- [routes](data-sources--route--properties--routes.md)
- [routes.route_destination](data-sources--route--properties--routes--route_destination.md)
- routes.route_destination.mirror_policy

<a id="section"></a>

Type: `"single"`. Computed.

MirrorPolicy is used for shadowing traffic from one cluster to another. The approach used is 'fire
and forget', meaning it will not wait for the shadow cluster to respond before returning the
response from the primary cluster. All normal statistics are collected for the shadow cluster
making..

Upstream description:

MirrorPolicy is used for shadowing traffic from one cluster to another. The approach used is "fire
and forget", meaning it will not wait for the shadow cluster to respond before returning the
response from the primary cluster. All normal statistics are collected for the shadow cluster making
this feature useful for testing and troubleshooting.

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

- [cluster](data-sources--route--properties--routes--route_destination--mirror_policy--cluster.md): complete subsection reference.

- [percent](data-sources--route--properties--routes--route_destination--mirror_policy--percent.md): complete subsection reference.

## Next pages

- [routes.route_destination.mirror_policy.cluster](data-sources--route--properties--routes--route_destination--mirror_policy--cluster.md)
- [routes.route_destination.mirror_policy.percent](data-sources--route--properties--routes--route_destination--mirror_policy--percent.md)
- [routes.route_destination](data-sources--route--properties--routes--route_destination.md)
- [xcsh_route](../data-sources/route.md)
