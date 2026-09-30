---
page_title: "routes.route_destination.mirror_policy"
subcategory: ""
description: "routes.route_destination.mirror_policy for xcsh_route."
xcsh_docs: {"aliases": [], "body_bytes": 2682, "body_sha256": "sha256:9c1d7c47dd2f9bbab1b5cdd0c3d6f55f1d0d0acc6a209611d694aa47c24bda30", "child_ids": ["xcsh-docs:resources:route:properties:routes:route_destination:mirror_policy:cluster", "xcsh-docs:resources:route:properties:routes:route_destination:mirror_policy:percent"], "collection_id": "xcsh-docs:resources:route:collection", "completeness": "complete", "id": "xcsh-docs:resources:route:properties:routes:route_destination:mirror_policy", "parent_id": "xcsh-docs:resources:route:properties:routes:route_destination", "path": "documentation/resources/route/properties/routes/route_destination/mirror_policy/index.md", "provider_name": "route", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["routes", "route_destination", "mirror_policy"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/route/properties/routes/route_destination/mirror_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.route_destination.mirror_policy for xcsh_route.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["routeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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

- [cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/mirror_policy/cluster/): complete subsection reference.

- [percent](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/mirror_policy/percent/): complete subsection reference.

## Next pages

- [routes.route_destination.mirror_policy.cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/mirror_policy/cluster/)
- [routes.route_destination.mirror_policy.percent](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/mirror_policy/percent/)
- [routes.route_destination](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/)
- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/)
