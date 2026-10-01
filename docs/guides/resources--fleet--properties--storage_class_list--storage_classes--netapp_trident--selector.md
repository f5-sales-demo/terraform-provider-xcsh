---
page_title: "storage_class_list.storage_classes.netapp_trident.selector"
subcategory: ""
description: "storage_class_list.storage_classes.netapp_trident.selector for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 1613, "body_sha256": "sha256:fad5d1252647e93ba1f0561bc8c28aafcc72537c6933ddb7069c395bb62c7672", "canonical_id": "xcsh-docs:resources:fleet:properties:storage_class_list:storage_classes:netapp_trident:selector", "child_ids": [], "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:storage_class_list:storage_classes:netapp_trident:selector", "parent_id": "xcsh-docs:resources:fleet:properties:storage_class_list:storage_classes:netapp_trident", "path": "docs/guides/resources--fleet--properties--storage_class_list--storage_classes--netapp_trident--selector.md", "provider_name": "fleet", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["storage_class_list", "storage_classes", "netapp_trident", "selector"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/storage_class_list/storage_classes/netapp_trident/selector/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "storage_class_list.storage_classes.netapp_trident.selector for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# storage_class_list.storage_classes.netapp_trident.selector

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md)
- [Property reference](resources--fleet--reference.md)
- [storage_class_list](resources--fleet--properties--storage_class_list.md)
- [storage_class_list.storage_classes](resources--fleet--properties--storage_class_list--storage_classes.md)
- [storage_class_list.storage_classes.netapp_trident](resources--fleet--properties--storage_class_list--storage_classes--netapp_trident.md)
- storage_class_list.storage_classes.netapp_trident.selector

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Using the Selector field, each StorageClass calls out which virtual pool(s) may be used to host a
volume. The volume will have the aspects defined in the chosen virtual pool.

Upstream description:

Using the Selector field, each StorageClass calls out which virtual pool(s) may be used to host a
volume. The volume will have the aspects defined in the chosen virtual pool.

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
selector {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [storage_class_list.storage_classes.netapp_trident](resources--fleet--properties--storage_class_list--storage_classes--netapp_trident.md)
- [xcsh_fleet](../resources/fleet.md)
