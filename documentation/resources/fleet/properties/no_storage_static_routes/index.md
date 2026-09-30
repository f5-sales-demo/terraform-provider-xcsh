---
page_title: "no_storage_static_routes"
subcategory: ""
description: "no_storage_static_routes for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 1505, "body_sha256": "sha256:71d3a1a84f7dc600e2c1e7b6bc64a6108e88ebb7824906f7fb4fbc0502cb7eb1", "child_ids": [], "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:no_storage_static_routes", "parent_id": "xcsh-docs:resources:fleet:reference", "path": "documentation/resources/fleet/properties/no_storage_static_routes/index.md", "provider_name": "fleet", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["no_storage_static_routes"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/no_storage_static_routes/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "no_storage_static_routes for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# no_storage_static_routes

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- no_storage_static_routes

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: no\_storage\_static\_routes, storage\_static\_routes; Default:
no\_storage\_static\_routes\] Configuration parameter for no storage static routes.

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

OneOf alternatives in this subsection:

- [no_storage_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/no_storage_static_routes/#section)
- [storage_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_static_routes/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
no_storage_static_routes = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
