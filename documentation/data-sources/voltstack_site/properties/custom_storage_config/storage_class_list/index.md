---
page_title: "custom_storage_config.storage_class_list"
subcategory: ""
description: "Add additional custom storage classes in Kubernetes for this fleet."
xcsh_docs: {"aliases": ["custom storage config storage class list"], "body_bytes": 1585, "body_sha256": "sha256:1a3d6e72f594e175c194bd82c60a17d098989ee703743163ee84431def931ece", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_class_list", "parent_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config", "path": "documentation/data-sources/voltstack_site/properties/custom_storage_config/storage_class_list/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0313302010302212-0330323323030201-2223132001112311-1033010323032321-2200130111100323-2131003310231011-1023031012213200-2333011201031333", "registry_path": "docs/guides/data-sources--voltstack_site--reference--group-005.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_storage_config", "storage_class_list"], "schema_version": 1, "sections": [{"aliases": ["custom storage config storage class list storage classes"], "anchor": "section", "description": "List of custom storage classes.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["custom_storage_config", "storage_class_list", "storage_classes"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/custom_storage_config/storage_class_list/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Add additional custom storage classes in Kubernetes for this fleet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_storage_config.storage_class_list

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/)
- [custom_storage_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/)
- custom_storage_config.storage_class_list

<a id="section"></a>

Type: `"single"`. Computed.

Add additional custom storage classes in Kubernetes for this fleet.

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

- [storage_classes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/storage_class_list/storage_classes/): complete subsection reference.

## Next pages

- [custom_storage_config.storage_class_list.storage_classes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/storage_class_list/storage_classes/)
- [custom_storage_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
