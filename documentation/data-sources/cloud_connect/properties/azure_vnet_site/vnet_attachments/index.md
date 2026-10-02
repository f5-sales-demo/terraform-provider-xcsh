---
page_title: "azure_vnet_site.vnet_attachments"
subcategory: ""
description: "Configuration parameter for vnet attachments."
xcsh_docs: {"aliases": ["azure vnet site vnet attachments"], "body_bytes": 1466, "body_sha256": "sha256:8842c50334a73cd1b29519d9d5f5e6c0bba7970077a8271847b80517d93d6e18", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:cloud_connect:properties:azure_vnet_site:vnet_attachments:vnet_list"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cloud_connect:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_connect:properties:azure_vnet_site:vnet_attachments", "parent_id": "xcsh-docs:data-sources:cloud_connect:properties:azure_vnet_site", "path": "documentation/data-sources/cloud_connect/properties/azure_vnet_site/vnet_attachments/index.md", "product": "distributed-cloud", "provider_name": "cloud_connect", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0230320312132323-2200131002110130-0233330032303133-0331032222032113-2033113223103310-2310201323303211-0102210022233122-2210201210032231", "registry_path": "docs/guides/data-sources--cloud_connect--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["azure_vnet_site", "vnet_attachments"], "schema_version": 1, "sections": [{"aliases": ["vnet list"], "anchor": "section", "description": "Collection of items or values", "document_id": "xcsh-docs:data-sources:cloud_connect:properties:azure_vnet_site:vnet_attachments:vnet_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["azure_vnet_site", "vnet_attachments", "vnet_list"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_connect/properties/azure_vnet_site/vnet_attachments/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configuration parameter for vnet attachments.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cloud_connectCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# azure_vnet_site.vnet_attachments

Breadcrumbs:

- [xcsh_cloud_connect](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_connect/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_connect/properties/)
- [azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_connect/properties/azure_vnet_site/)
- azure_vnet_site.vnet_attachments

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for vnet attachments.

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

- [vnet_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_connect/properties/azure_vnet_site/vnet_attachments/vnet_list/): complete subsection reference.

## Next pages

- [azure_vnet_site.vnet_attachments.vnet_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_connect/properties/azure_vnet_site/vnet_attachments/vnet_list/)
- [azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_connect/properties/azure_vnet_site/)
- [xcsh_cloud_connect](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_connect/)
