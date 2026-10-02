---
page_title: "voltstack_cluster_ar.storage_class_list.storage_classes"
subcategory: "Infrastructure"
description: "List of custom storage classes."
xcsh_docs: {"aliases": ["voltstack cluster ar storage class list storage classes"], "body_bytes": 3861, "body_sha256": "sha256:22f10597a31ff1c2bbe4f204f9246bd139204a8379aa4fefc455abc6f1628863", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar:storage_class_list:storage_classes", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar:storage_class_list", "path": "documentation/resources/azure_vnet_site/properties/voltstack_cluster_ar/storage_class_list/storage_classes/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2202310203302333-2232312100213231-3132312310310130-0023223321002322-0332303231312011-0101020212132111-3323212201012201-0013102110003113", "registry_path": "docs/guides/resources--azure_vnet_site--reference--group-010.md", "relationships": [{"anchor": "schema-voltstack_cluster_ar--storage_class_list--storage_classes--storage_class_name", "enforcement": "provider-schema", "group": "voltstack_cluster_ar.storage_class_list.storage_classes:RequiredListObjectAttributes:storage_class_name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar:storage_class_list:storage_classes", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["voltstack_cluster_ar", "storage_class_list", "storage_classes"], "schema_version": 1, "sections": [{"aliases": ["default storage class"], "anchor": "schema-voltstack_cluster_ar--storage_class_list--storage_classes--default_storage_class", "description": "Make this storage class default storage class for the K8s cluster.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar:storage_class_list:storage_classes", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["voltstack_cluster_ar", "storage_class_list", "storage_classes", "default_storage_class"], "syntax": "attribute", "type": "bool"}, {"aliases": ["storage class name"], "anchor": "schema-voltstack_cluster_ar--storage_class_list--storage_classes--storage_class_name", "description": "Name of the storage class as it will appear in K8s.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar:storage_class_list:storage_classes", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["voltstack_cluster_ar", "storage_class_list", "storage_classes", "storage_class_name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/voltstack_cluster_ar/storage_class_list/storage_classes/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "List of custom storage classes.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# voltstack_cluster_ar.storage_class_list.storage_classes

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/)
- [voltstack_cluster_ar](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster_ar/)
- [voltstack_cluster_ar.storage_class_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster_ar/storage_class_list/)
- voltstack_cluster_ar.storage_class_list.storage_classes

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

List of Storage Classes. List of custom storage classes.

Upstream description:

List of custom storage classes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("storage_class_name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
storage_classes {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-voltstack_cluster_ar--storage_class_list--storage_classes--default_storage_class"></a>

### default_storage_class property

Type: `"bool"`. Optional.

Make this storage class default storage class for the K8s cluster.

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

<a id="schema-voltstack_cluster_ar--storage_class_list--storage_classes--storage_class_name"></a>

### storage_class_name property

Type: `"string"`. Optional.

Name of the storage class as it will appear in K8s.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

## Next pages

- [voltstack_cluster_ar.storage_class_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster_ar/storage_class_list/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
