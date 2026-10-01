---
page_title: "custom_storage_config.storage_class_list.storage_classes.netapp_trident.selector"
subcategory: ""
description: "custom_storage_config.storage_class_list.storage_classes.netapp_trident.selector for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 1963, "body_sha256": "sha256:61d1471f0c415901cd9513fb67ac53c82230bb88b2ff6e58a13496a140c2f5ea", "canonical_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes:netapp_trident:selector", "child_ids": [], "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes:netapp_trident:selector", "parent_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes:netapp_trident", "path": "docs/guides/data-sources--voltstack_site--properties--custom_storage_config--storage_class_list--storage_classes--netapp_trident--selector.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_storage_config", "storage_class_list", "storage_classes", "netapp_trident", "selector"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/custom_storage_config/storage_class_list/storage_classes/netapp_trident/selector/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_storage_config.storage_class_list.storage_classes.netapp_trident.selector for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_storage_config.storage_class_list.storage_classes.netapp_trident.selector

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md)
- [Property reference](data-sources--voltstack_site--reference.md)
- [custom_storage_config](data-sources--voltstack_site--properties--custom_storage_config.md)
- [custom_storage_config.storage_class_list](data-sources--voltstack_site--properties--custom_storage_config--storage_class_list.md)
- [custom_storage_config.storage_class_list.storage_classes](data-sources--voltstack_site--properties--custom_storage_config--storage_class_list--storage_classes.md)
- [custom_storage_config.storage_class_list.storage_classes.netapp_trident](data-sources--voltstack_site--properties--custom_storage_config--storage_class_list--storage_classes--netapp_trident.md)
- custom_storage_config.storage_class_list.storage_classes.netapp_trident.selector

<a id="section"></a>

Type: `"single"`. Computed.

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

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [custom_storage_config.storage_class_list.storage_classes.netapp_trident](data-sources--voltstack_site--properties--custom_storage_config--storage_class_list--storage_classes--netapp_trident.md)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md)
