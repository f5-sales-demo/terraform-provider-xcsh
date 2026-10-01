---
page_title: "custom_storage_config.no_storage_device"
subcategory: ""
description: "custom_storage_config.no_storage_device for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 1088, "body_sha256": "sha256:8c3c9495622e9297ac42964db1f5eab62eb265e5816d502f1a67b13836bb689b", "canonical_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:no_storage_device", "child_ids": [], "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:no_storage_device", "parent_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config", "path": "docs/guides/resources--voltstack_site--properties--custom_storage_config--no_storage_device.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_storage_config", "no_storage_device"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/custom_storage_config/no_storage_device/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_storage_config.no_storage_device for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_storage_config.no_storage_device

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md)
- [Property reference](resources--voltstack_site--reference.md)
- [custom_storage_config](resources--voltstack_site--properties--custom_storage_config.md)
- custom_storage_config.no_storage_device

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no storage device.

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
no_storage_device = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [custom_storage_config](resources--voltstack_site--properties--custom_storage_config.md)
- [xcsh_voltstack_site](../resources/voltstack_site.md)
