---
page_title: "default_config"
subcategory: ""
description: "default_config for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 1608, "body_sha256": "sha256:13b26f41aa76479221a244aacb4f748530d3d0235527f0e743582548c9ef6a4e", "child_ids": [], "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:default_config", "parent_id": "xcsh-docs:resources:fleet:reference", "path": "documentation/resources/fleet/properties/default_config/index.md", "provider_name": "fleet", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["default_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/default_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "default_config for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_config

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- default_config

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: default\_config, device\_list, interface\_list; Default: default\_config\] Enable this
option

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

- [default_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/default_config/#section)
- [device_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/device_list/#section)
- [interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/interface_list/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
default_config = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
