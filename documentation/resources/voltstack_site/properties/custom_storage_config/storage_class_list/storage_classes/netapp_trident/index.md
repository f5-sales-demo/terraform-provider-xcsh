---
page_title: "custom_storage_config.storage_class_list.storage_classes.netapp_trident"
subcategory: ""
description: "Storage class Device configuration for NetApp Trident."
xcsh_docs: {"aliases": ["custom storage config storage class list storage classes netapp trident"], "body_bytes": 3305, "body_sha256": "sha256:8fd5390ba5ee2c169475bb21479292398446ed97327a80fd4df3a4a7656efe91", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes:netapp_trident:selector"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes:netapp_trident", "parent_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes", "path": "documentation/resources/voltstack_site/properties/custom_storage_config/storage_class_list/storage_classes/netapp_trident/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3010213222022331-2313320003203232-3130302120220302-0130222212302333-2103232323303223-2210112231333120-1311222210032110-3323312220122300", "registry_path": "docs/guides/resources--voltstack_site--reference--group-006.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_storage_config", "storage_class_list", "storage_classes", "netapp_trident"], "schema_version": 1, "sections": [{"aliases": ["selector"], "anchor": "section", "description": "Using the Selector field, each StorageClass calls out which virtual pool(s) may be used to host a volume. The volume will have the aspects defined in the chosen virtual pool.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes:netapp_trident:selector", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["custom_storage_config", "storage_class_list", "storage_classes", "netapp_trident", "selector"], "syntax": "block", "type": "object"}, {"aliases": ["storage pools"], "anchor": "schema-custom_storage_config--storage_class_list--storage_classes--netapp_trident--storage_pools", "description": "The storagePools parameter is used to further restrict the set of pools that match any specified attributes.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes:netapp_trident", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_class_list", "storage_classes", "netapp_trident", "storage_pools"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/custom_storage_config/storage_class_list/storage_classes/netapp_trident/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Storage class Device configuration for NetApp Trident.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_storage_config.storage_class_list.storage_classes.netapp_trident

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/)
- [custom_storage_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/)
- [custom_storage_config.storage_class_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_class_list/)
- [custom_storage_config.storage_class_list.storage_classes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_class_list/storage_classes/)
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

- [selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_class_list/storage_classes/netapp_trident/selector/): complete subsection reference.

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

- [custom_storage_config.storage_class_list.storage_classes.netapp_trident.selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_class_list/storage_classes/netapp_trident/selector/)
- [custom_storage_config.storage_class_list.storage_classes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_class_list/storage_classes/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
