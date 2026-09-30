---
page_title: "routes.route_destination.retract_cluster"
subcategory: ""
description: "routes.route_destination.retract_cluster for xcsh_route."
xcsh_docs: {"aliases": [], "body_bytes": 964, "body_sha256": "sha256:8fbc8e2a76d32a53044fe93fb78d2c782b6bbde99365eb13ebc0037c3227e32a", "canonical_id": "xcsh-docs:resources:route:properties:routes:route_destination:retract_cluster", "child_ids": [], "collection_id": "xcsh-docs:resources:route:collection", "completeness": "complete", "id": "xcsh-docs:resources:route:properties:routes:route_destination:retract_cluster", "parent_id": "xcsh-docs:resources:route:properties:routes:route_destination", "path": "docs/guides/resources--route--properties--routes--route_destination--retract_cluster.md", "provider_name": "route", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes", "route_destination", "retract_cluster"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/route/properties/routes/route_destination/retract_cluster/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.route_destination.retract_cluster for xcsh_route.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["routeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# routes.route_destination.retract_cluster

Breadcrumbs:

- [xcsh_route](../resources/route.md)
- [Property reference](resources--route--reference.md)
- [routes](resources--route--properties--routes.md)
- [routes.route_destination](resources--route--properties--routes--route_destination.md)
- routes.route_destination.retract_cluster

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
retract_cluster = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [routes.route_destination](resources--route--properties--routes--route_destination.md)
- [xcsh_route](../resources/route.md)
