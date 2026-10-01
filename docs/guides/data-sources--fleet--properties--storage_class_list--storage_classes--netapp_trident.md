---
page_title: "storage_class_list.storage_classes.netapp_trident"
subcategory: ""
description: "storage_class_list.storage_classes.netapp_trident for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 2169, "body_sha256": "sha256:8bf49b25a62a8c7b27e54a1e88306608d3d217aced5ddf29489d8d2330d1f787", "canonical_id": "xcsh-docs:data-sources:fleet:properties:storage_class_list:storage_classes:netapp_trident", "child_ids": ["xcsh-docs:data-sources:fleet:properties:storage_class_list:storage_classes:netapp_trident:selector"], "collection_id": "xcsh-docs:data-sources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fleet:properties:storage_class_list:storage_classes:netapp_trident", "parent_id": "xcsh-docs:data-sources:fleet:properties:storage_class_list:storage_classes", "path": "docs/guides/data-sources--fleet--properties--storage_class_list--storage_classes--netapp_trident.md", "provider_name": "fleet", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["storage_class_list", "storage_classes", "netapp_trident"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fleet/properties/storage_class_list/storage_classes/netapp_trident/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "storage_class_list.storage_classes.netapp_trident for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# storage_class_list.storage_classes.netapp_trident

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md)
- [Property reference](data-sources--fleet--reference.md)
- [storage_class_list](data-sources--fleet--properties--storage_class_list.md)
- [storage_class_list.storage_classes](data-sources--fleet--properties--storage_class_list--storage_classes.md)
- storage_class_list.storage_classes.netapp_trident

<a id="section"></a>

Type: `"single"`. Computed.

Storage class Device configuration for NetApp Trident.

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

- [selector](data-sources--fleet--properties--storage_class_list--storage_classes--netapp_trident--selector.md): complete subsection reference.

<a id="schema-storage_class_list--storage_classes--netapp_trident--storage_pools"></a>

### storage_pools property

Type: `"string"`. Computed.

The storagePools parameter is used to further restrict the set of pools that match any specified
attributes.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512"
  }
}
```

## Next pages

- [storage_class_list.storage_classes.netapp_trident.selector](data-sources--fleet--properties--storage_class_list--storage_classes--netapp_trident--selector.md)
- [storage_class_list.storage_classes](data-sources--fleet--properties--storage_class_list--storage_classes.md)
- [xcsh_fleet](../data-sources/fleet.md)
