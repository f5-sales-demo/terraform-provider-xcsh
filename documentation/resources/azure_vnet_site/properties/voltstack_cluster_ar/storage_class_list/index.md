---
page_title: "voltstack_cluster_ar.storage_class_list"
subcategory: "Infrastructure"
description: "Add additional custom storage classes in Kubernetes for this site."
xcsh_docs: {"aliases": ["voltstack cluster ar storage class list"], "body_bytes": 1690, "body_sha256": "sha256:db955c0bd5ea27eb8107fc5c6ea5df612b815475c5f04966322642a5385c19f3", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar:storage_class_list:storage_classes"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar:storage_class_list", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar", "path": "documentation/resources/azure_vnet_site/properties/voltstack_cluster_ar/storage_class_list/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3333220321002101-1012013013201131-3112320212130103-2132312101213111-3011030103300022-0201231001332133-2223211331031221-1312311232113101", "registry_path": "docs/guides/resources--azure_vnet_site--reference--group-010.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["voltstack_cluster_ar", "storage_class_list"], "schema_version": 1, "sections": [{"aliases": ["voltstack cluster ar storage class list storage classes"], "anchor": "section", "description": "List of custom storage classes.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar:storage_class_list:storage_classes", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-voltstack_cluster_ar--storage_class_list--storage_classes--storage_class_name", "enforcement": "provider-schema", "group": "voltstack_cluster_ar.storage_class_list.storage_classes:RequiredListObjectAttributes:storage_class_name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar:storage_class_list:storage_classes", "type": "requires"}], "schema_path": ["voltstack_cluster_ar", "storage_class_list", "storage_classes"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/voltstack_cluster_ar/storage_class_list/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Add additional custom storage classes in Kubernetes for this site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# voltstack_cluster_ar.storage_class_list

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/)
- [voltstack_cluster_ar](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster_ar/)
- voltstack_cluster_ar.storage_class_list

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Add additional custom storage classes in Kubernetes for this site.

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
storage_class_list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [storage_classes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster_ar/storage_class_list/storage_classes/): complete subsection reference.

## Next pages

- [voltstack_cluster_ar.storage_class_list.storage_classes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster_ar/storage_class_list/storage_classes/)
- [voltstack_cluster_ar](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster_ar/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
