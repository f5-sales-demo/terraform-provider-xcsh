---
page_title: "azure_vnet_site.vnet_attachments"
subcategory: ""
description: "Configuration parameter for vnet attachments."
xcsh_docs: {"aliases": ["azure vnet site vnet attachments"], "body_bytes": 1570, "body_sha256": "sha256:c3cba08b85220bb172012f38e3d3cf230f9ed542f99ad129c9bdc56ce9d508d0", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:cloud_connect:properties:azure_vnet_site:vnet_attachments:vnet_list"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cloud_connect:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_connect:properties:azure_vnet_site:vnet_attachments", "parent_id": "xcsh-docs:resources:cloud_connect:properties:azure_vnet_site", "path": "documentation/resources/cloud_connect/properties/azure_vnet_site/vnet_attachments/index.md", "product": "distributed-cloud", "provider_name": "cloud_connect", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3022033033322013-0223203001230023-1301122301311130-3323100132012112-1213320010021300-3231311210113021-1200002123210200-2332333310000100", "registry_path": "docs/guides/resources--cloud_connect--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["azure_vnet_site", "vnet_attachments"], "schema_version": 1, "sections": [{"aliases": ["azure vnet site vnet attachments vnet list"], "anchor": "section", "description": "Collection of items or values", "document_id": "xcsh-docs:resources:cloud_connect:properties:azure_vnet_site:vnet_attachments:vnet_list", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "azure_vnet_site.vnet_attachments.vnet_list:ConflictingListObjectAttributes:custom_routing,default_route", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cloud_connect:properties:azure_vnet_site:vnet_attachments:vnet_list:custom_routing", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "azure_vnet_site.vnet_attachments.vnet_list:ConflictingListObjectAttributes:custom_routing,manual_routing", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cloud_connect:properties:azure_vnet_site:vnet_attachments:vnet_list:custom_routing", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "azure_vnet_site.vnet_attachments.vnet_list:ConflictingListObjectAttributes:custom_routing,default_route", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cloud_connect:properties:azure_vnet_site:vnet_attachments:vnet_list:default_route", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "azure_vnet_site.vnet_attachments.vnet_list:ConflictingListObjectAttributes:default_route,manual_routing", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cloud_connect:properties:azure_vnet_site:vnet_attachments:vnet_list:default_route", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "azure_vnet_site.vnet_attachments.vnet_list:ConflictingListObjectAttributes:custom_routing,manual_routing", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cloud_connect:properties:azure_vnet_site:vnet_attachments:vnet_list:manual_routing", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "azure_vnet_site.vnet_attachments.vnet_list:ConflictingListObjectAttributes:default_route,manual_routing", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cloud_connect:properties:azure_vnet_site:vnet_attachments:vnet_list:manual_routing", "type": "conflicts"}, {"anchor": "schema-azure_vnet_site--vnet_attachments--vnet_list--subscription_id", "enforcement": "provider-schema", "group": "azure_vnet_site.vnet_attachments.vnet_list:RequiredListObjectAttributes:subscription_id,vnet_id", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:cloud_connect:properties:azure_vnet_site:vnet_attachments:vnet_list", "type": "requires"}, {"anchor": "schema-azure_vnet_site--vnet_attachments--vnet_list--vnet_id", "enforcement": "provider-schema", "group": "azure_vnet_site.vnet_attachments.vnet_list:RequiredListObjectAttributes:subscription_id,vnet_id", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:cloud_connect:properties:azure_vnet_site:vnet_attachments:vnet_list", "type": "requires"}], "schema_path": ["azure_vnet_site", "vnet_attachments", "vnet_list"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_connect/properties/azure_vnet_site/vnet_attachments/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configuration parameter for vnet attachments.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cloud_connectCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# azure_vnet_site.vnet_attachments

Breadcrumbs:

- [xcsh_cloud_connect](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/)
- [azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/azure_vnet_site/)
- azure_vnet_site.vnet_attachments

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
vnet_attachments {
  # Configure direct properties listed below.
}
```

## Direct properties

- [vnet_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/azure_vnet_site/vnet_attachments/vnet_list/): complete subsection reference.

## Next pages

- [azure_vnet_site.vnet_attachments.vnet_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/azure_vnet_site/vnet_attachments/vnet_list/)
- [azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/azure_vnet_site/)
- [xcsh_cloud_connect](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/)
