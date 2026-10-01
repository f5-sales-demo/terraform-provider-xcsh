---
page_title: "custom_storage_config.storage_class_list.storage_classes.netapp_trident"
subcategory: ""
description: "custom_storage_config.storage_class_list.storage_classes.netapp_trident for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 2857, "body_sha256": "sha256:cce94e49f951b4e72cb29cf5839b2d035a8890ede08452196413965af631b5cc", "canonical_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes:netapp_trident", "child_ids": ["xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes:netapp_trident:selector"], "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes:netapp_trident", "parent_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes", "path": "docs/guides/resources--voltstack_site--properties--custom_storage_config--storage_class_list--storage_classes--netapp_trident.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_storage_config", "storage_class_list", "storage_classes", "netapp_trident"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/custom_storage_config/storage_class_list/storage_classes/netapp_trident/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_storage_config.storage_class_list.storage_classes.netapp_trident for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_storage_config.storage_class_list.storage_classes.netapp_trident

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md)
- [Property reference](resources--voltstack_site--reference.md)
- [custom_storage_config](resources--voltstack_site--properties--custom_storage_config.md)
- [custom_storage_config.storage_class_list](resources--voltstack_site--properties--custom_storage_config--storage_class_list.md)
- [custom_storage_config.storage_class_list.storage_classes](resources--voltstack_site--properties--custom_storage_config--storage_class_list--storage_classes.md)
- custom_storage_config.storage_class_list.storage_classes.netapp_trident

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
netapp_trident {
  # Configure direct properties listed below.
}
```

## Direct properties

- [selector](resources--voltstack_site--properties--custom_storage_config--storage_class_list--storage_classes--netapp_trident--selector.md): complete subsection reference.

<a id="schema-custom_storage_config--storage_class_list--storage_classes--netapp_trident--storage_pools"></a>

### storage_pools property

Type: `"string"`. Optional.

The storagePools parameter is used to further restrict the set of pools that match any specified
attributes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(512),
}
```

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

- [custom_storage_config.storage_class_list.storage_classes.netapp_trident.selector](resources--voltstack_site--properties--custom_storage_config--storage_class_list--storage_classes--netapp_trident--selector.md)
- [custom_storage_config.storage_class_list.storage_classes](resources--voltstack_site--properties--custom_storage_config--storage_class_list--storage_classes.md)
- [xcsh_voltstack_site](../resources/voltstack_site.md)
