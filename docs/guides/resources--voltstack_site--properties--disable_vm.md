---
page_title: "disable_vm"
subcategory: ""
description: "disable_vm for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 1192, "body_sha256": "sha256:ba9a3893afca0f28c1ffdd922afeed8d25553a356c90ad2f8570690d8dc4d3da", "canonical_id": "xcsh-docs:resources:voltstack_site:properties:disable_vm", "child_ids": [], "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:disable_vm", "parent_id": "xcsh-docs:resources:voltstack_site:reference", "path": "docs/guides/resources--voltstack_site--properties--disable_vm.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["disable_vm"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/disable_vm/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "disable_vm for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# disable_vm

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md)
- [Property reference](resources--voltstack_site--reference.md)
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

- [disable_vm](resources--voltstack_site--properties--disable_vm.md#section)
- [enable_vm](resources--voltstack_site--properties--enable_vm.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_vm = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--voltstack_site--reference.md)
- [xcsh_voltstack_site](../resources/voltstack_site.md)
