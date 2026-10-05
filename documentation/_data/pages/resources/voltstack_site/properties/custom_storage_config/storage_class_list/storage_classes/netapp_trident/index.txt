---
page_title: "custom_storage_config.storage_class_list.storage_classes.netapp_trident"
subcategory: ""
description: "Storage class Device configuration for NetApp Trident."
xcsh_docs: {"aliases": ["custom storage config storage class list storage classes netapp trident"], "body_bytes": 3305, "body_sha256": "sha256:feca2afe2c3938a186c7902da17012c5ce8a429c54f2de0ae299dd370b8bf380", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes:netapp_trident:selector"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes:netapp_trident", "parent_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes", "path": "documentation/resources/voltstack_site/properties/custom_storage_config/storage_class_list/storage_classes/netapp_trident/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-3010213222022331-2313320003203232-3130302120220302-0130222212302333-2103232323303223-2210112231333120-1311222210032110-3323312220122300", "registry_path": "docs/guides/resources--voltstack_site--reference--group-006.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_storage_config", "storage_class_list", "storage_classes", "netapp_trident"], "schema_version": 1, "sections": [{"aliases": ["custom storage config storage class list storage classes netapp trident selector"], "anchor": "section", "description": "Using the Selector field, each StorageClass calls out which virtual pool(s) may be used to host a volume. The volume will have the aspects defined in the chosen virtual pool.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes:netapp_trident:selector", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["custom_storage_config", "storage_class_list", "storage_classes", "netapp_trident", "selector"], "syntax": "block", "type": "object"}, {"aliases": ["custom storage config storage class list storage classes netapp trident storage pools"], "anchor": "schema-custom_storage_config--storage_class_list--storage_classes--netapp_trident--storage_pools", "description": "The storagePools parameter is used to further restrict the set of pools that match any specified attributes.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes:netapp_trident", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_class_list", "storage_classes", "netapp_trident", "storage_pools"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/custom_storage_config/storage_class_list/storage_classes/netapp_trident/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Storage class Device configuration for NetApp Trident.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
