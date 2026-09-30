---
page_title: "disable_gpu"
subcategory: ""
description: "disable_gpu for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 1498, "body_sha256": "sha256:f51636b04f73b74df26a1f03952cc9daebfd32b4a8894ec9264a32c3d67defb3", "child_ids": [], "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:disable_gpu", "parent_id": "xcsh-docs:resources:fleet:reference", "path": "documentation/resources/fleet/properties/disable_gpu/index.md", "provider_name": "fleet", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["disable_gpu"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/disable_gpu/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "disable_gpu for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# disable_gpu

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- disable_gpu

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable\_gpu, enable\_gpu, enable\_vgpu; Default: disable\_gpu\] Configuration parameter
for disable gpu.

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

- [disable_gpu](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/disable_gpu/#section)
- [enable_gpu](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/enable_gpu/#section)
- [enable_vgpu](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/enable_vgpu/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_gpu = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
