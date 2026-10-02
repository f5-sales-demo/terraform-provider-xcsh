---
page_title: "voltstack_cluster.storage_class_list.storage_classes"
subcategory: "Infrastructure"
description: "List of custom storage classes."
xcsh_docs: {"aliases": ["voltstack cluster storage class list storage classes"], "body_bytes": 3831, "body_sha256": "sha256:2e8f8fb3c4a2fdcb8da50444a2a28b4680c1dd9752e80073e2bc36c863d53617", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster:storage_class_list:storage_classes", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster:storage_class_list", "path": "documentation/resources/azure_vnet_site/properties/voltstack_cluster/storage_class_list/storage_classes/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0310003233021333-1233033012301212-3212331021002322-2300331102212320-1310113332033123-2010022000102222-1311331231220120-0033111013023030", "registry_path": "docs/guides/resources--azure_vnet_site--reference--group-009.md", "relationships": [{"anchor": "schema-voltstack_cluster--storage_class_list--storage_classes--storage_class_name", "enforcement": "provider-schema", "group": "voltstack_cluster.storage_class_list.storage_classes:RequiredListObjectAttributes:storage_class_name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster:storage_class_list:storage_classes", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["voltstack_cluster", "storage_class_list", "storage_classes"], "schema_version": 1, "sections": [{"aliases": ["default storage class"], "anchor": "schema-voltstack_cluster--storage_class_list--storage_classes--default_storage_class", "description": "Make this storage class default storage class for the K8s cluster.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster:storage_class_list:storage_classes", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["voltstack_cluster", "storage_class_list", "storage_classes", "default_storage_class"], "syntax": "attribute", "type": "bool"}, {"aliases": ["storage class name"], "anchor": "schema-voltstack_cluster--storage_class_list--storage_classes--storage_class_name", "description": "Name of the storage class as it will appear in K8s.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster:storage_class_list:storage_classes", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["voltstack_cluster", "storage_class_list", "storage_classes", "storage_class_name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/voltstack_cluster/storage_class_list/storage_classes/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of custom storage classes.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# voltstack_cluster.storage_class_list.storage_classes

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/)
- [voltstack_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster/)
- [voltstack_cluster.storage_class_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster/storage_class_list/)
- voltstack_cluster.storage_class_list.storage_classes

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

<a id="schema-voltstack_cluster--storage_class_list--storage_classes--default_storage_class"></a>

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

<a id="schema-voltstack_cluster--storage_class_list--storage_classes--storage_class_name"></a>

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

- [voltstack_cluster.storage_class_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster/storage_class_list/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
