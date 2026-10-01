---
page_title: "disable_vm"
subcategory: ""
description: "disable_vm for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 1430, "body_sha256": "sha256:a4484c0d0a0162f87056e59544bef91c9635c723fcd9bb939829fa1353c03340", "child_ids": [], "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:disable_vm", "parent_id": "xcsh-docs:resources:fleet:reference", "path": "documentation/resources/fleet/properties/disable_vm/index.md", "provider_name": "fleet", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["disable_vm"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/disable_vm/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "disable_vm for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# disable_vm

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- disable_vm

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable\_vm, enable\_vm; Default: disable\_vm\] Enable this option

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

- [disable_vm](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/disable_vm/#section)
- [enable_vm](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/enable_vm/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_vm = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
