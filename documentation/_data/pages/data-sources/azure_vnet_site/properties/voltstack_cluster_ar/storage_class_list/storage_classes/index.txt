---
page_title: "voltstack_cluster_ar.storage_class_list.storage_classes"
subcategory: "Infrastructure"
description: "List of custom storage classes."
xcsh_docs: {"aliases": ["voltstack cluster ar storage class list storage classes"], "body_bytes": 3457, "body_sha256": "sha256:6106cfc08f14779ac735b854a465dc7f51daefea4d291505a797a35ac0f41a30", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar:storage_class_list:storage_classes", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar:storage_class_list", "path": "documentation/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/storage_class_list/storage_classes/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0313330212232110-2111131032323020-0102103030211000-2100323330001331-0322212011200230-0203322323303112-1221103123130110-3111022132022100", "registry_path": "docs/guides/data-sources--azure_vnet_site--reference--group-010.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["voltstack_cluster_ar", "storage_class_list", "storage_classes"], "schema_version": 1, "sections": [{"aliases": ["default storage class"], "anchor": "schema-voltstack_cluster_ar--storage_class_list--storage_classes--default_storage_class", "description": "Make this storage class default storage class for the K8s cluster.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar:storage_class_list:storage_classes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["voltstack_cluster_ar", "storage_class_list", "storage_classes", "default_storage_class"], "syntax": "attribute", "type": "bool"}, {"aliases": ["storage class name"], "anchor": "schema-voltstack_cluster_ar--storage_class_list--storage_classes--storage_class_name", "description": "Name of the storage class as it will appear in K8s.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar:storage_class_list:storage_classes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["voltstack_cluster_ar", "storage_class_list", "storage_classes", "storage_class_name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/storage_class_list/storage_classes/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of custom storage classes.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# voltstack_cluster_ar.storage_class_list.storage_classes

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/)
- [voltstack_cluster_ar](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/)
- [voltstack_cluster_ar.storage_class_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/storage_class_list/)
- voltstack_cluster_ar.storage_class_list.storage_classes

<a id="section"></a>

Type: `"list"`. Computed.

List of Storage Classes. List of custom storage classes.

Upstream description:

List of custom storage classes.

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

## Direct properties

<a id="schema-voltstack_cluster_ar--storage_class_list--storage_classes--default_storage_class"></a>

### default_storage_class property

Type: `"bool"`. Computed.

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

Type: `"string"`. Computed.

Name of the storage class as it will appear in K8s.

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

- [voltstack_cluster_ar.storage_class_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/storage_class_list/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
