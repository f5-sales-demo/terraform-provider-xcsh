---
page_title: "disable_pfs"
subcategory: ""
description: "disable_pfs for xcsh_ike2."
xcsh_docs: {"aliases": [], "body_bytes": 743, "body_sha256": "sha256:f47a1f91637cc6d79820d1c04791addbf86212e5cd3695777f59dbf71b397c8c", "canonical_id": "xcsh-docs:resources:ike2:properties:disable_pfs", "child_ids": [], "collection_id": "xcsh-docs:resources:ike2:collection", "completeness": "complete", "id": "xcsh-docs:resources:ike2:properties:disable_pfs", "parent_id": "xcsh-docs:resources:ike2:reference", "path": "docs/guides/resources--ike2--properties--disable_pfs.md", "provider_name": "ike2", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["disable_pfs"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/ike2/properties/disable_pfs/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "disable_pfs for xcsh_ike2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["ike2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# disable_pfs

Breadcrumbs:

- [xcsh_ike2](../resources/ike2.md)
- [Property reference](resources--ike2--reference.md)
- disable_pfs

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable pfs.

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
disable_pfs = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--ike2--reference.md)
- [xcsh_ike2](../resources/ike2.md)
