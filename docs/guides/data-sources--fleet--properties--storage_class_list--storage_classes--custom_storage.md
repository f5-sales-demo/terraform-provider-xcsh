---
page_title: "storage_class_list.storage_classes.custom_storage"
subcategory: ""
description: "storage_class_list.storage_classes.custom_storage for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 2152, "body_sha256": "sha256:1b97eacabeee4b5d75a89beda721c5cb8528c770f77e0e58f18bf63cec01bc0e", "canonical_id": "xcsh-docs:data-sources:fleet:properties:storage_class_list:storage_classes:custom_storage", "child_ids": [], "collection_id": "xcsh-docs:data-sources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fleet:properties:storage_class_list:storage_classes:custom_storage", "parent_id": "xcsh-docs:data-sources:fleet:properties:storage_class_list:storage_classes", "path": "docs/guides/data-sources--fleet--properties--storage_class_list--storage_classes--custom_storage.md", "provider_name": "fleet", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["storage_class_list", "storage_classes", "custom_storage"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fleet/properties/storage_class_list/storage_classes/custom_storage/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "storage_class_list.storage_classes.custom_storage for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# storage_class_list.storage_classes.custom_storage

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md)
- [Property reference](data-sources--fleet--reference.md)
- [storage_class_list](data-sources--fleet--properties--storage_class_list.md)
- [storage_class_list.storage_classes](data-sources--fleet--properties--storage_class_list--storage_classes.md)
- storage_class_list.storage_classes.custom_storage

<a id="section"></a>

Type: `"single"`. Computed.

Custom Storage Class allows to insert Kubernetes storageclass definition which will be applied into
given site.

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

<a id="schema-storage_class_list--storage_classes--custom_storage--yaml"></a>

### yaml property

Type: `"string"`. Computed.

Storage Class YAML. K8s YAML for StorageClass.

Upstream description:

K8s YAML for StorageClass.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 4096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "Valid parseable YAML",
    "maxLength": 4096,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "validation": {
      "customRule": "Must be valid YAML"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

## Next pages

- [storage_class_list.storage_classes](data-sources--fleet--properties--storage_class_list--storage_classes.md)
- [xcsh_fleet](../data-sources/fleet.md)
