---
page_title: "disable_gpu"
subcategory: ""
description: "disable_gpu for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 1579, "body_sha256": "sha256:9c056bb5836abf509f64daa470f1cbe2fb1426d1d8141ac95a52ef6d01507281", "child_ids": [], "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:disable_gpu", "parent_id": "xcsh-docs:resources:voltstack_site:reference", "path": "documentation/resources/voltstack_site/properties/disable_gpu/index.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["disable_gpu"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/disable_gpu/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "disable_gpu for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# disable_gpu

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/)
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

- [disable_gpu](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/disable_gpu/#section)
- [enable_gpu](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/enable_gpu/#section)
- [enable_vgpu](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/enable_vgpu/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_gpu = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
