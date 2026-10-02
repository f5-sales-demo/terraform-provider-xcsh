---
page_title: "azure_vnet_site"
subcategory: ""
description: "Cloud Connect Azure VNet Site Type."
xcsh_docs: {"aliases": ["azure vnet site"], "body_bytes": 1737, "body_sha256": "sha256:b2bfeb4d62953a411741a707d095d4182d4633c099cc299d90a5acc51ec59c93", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:cloud_connect:properties:azure_vnet_site:site", "xcsh-docs:resources:cloud_connect:properties:azure_vnet_site:vnet_attachments"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cloud_connect:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_connect:properties:azure_vnet_site", "parent_id": "xcsh-docs:resources:cloud_connect:reference", "path": "documentation/resources/cloud_connect/properties/azure_vnet_site/index.md", "product": "distributed-cloud", "provider_name": "cloud_connect", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3021103322130320-1200003010013121-0010020300322100-0312322031223102-3220323002313103-2002300323233103-0021222100212022-3300231133311003", "registry_path": "docs/guides/resources--cloud_connect--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["azure_vnet_site"], "schema_version": 1, "sections": [{"aliases": ["site"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:cloud_connect:properties:azure_vnet_site:site", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-azure_vnet_site--site--name", "enforcement": "provider-schema", "group": "azure_vnet_site.site:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_connect:properties:azure_vnet_site:site", "type": "requires"}], "schema_path": ["azure_vnet_site", "site"], "syntax": "block", "type": "object"}, {"aliases": ["vnet attachments"], "anchor": "section", "description": "Configuration parameter for vnet attachments.", "document_id": "xcsh-docs:resources:cloud_connect:properties:azure_vnet_site:vnet_attachments", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["azure_vnet_site", "vnet_attachments"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_connect/properties/azure_vnet_site/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Cloud Connect Azure VNet Site Type.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_connectCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# azure_vnet_site

Breadcrumbs:

- [xcsh_cloud_connect](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/)
- azure_vnet_site

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Azure VNet Site Type. Cloud Connect Azure VNet Site Type.

Upstream description:

Cloud Connect Azure VNet Site Type.

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
azure_vnet_site {
  # Configure direct properties listed below.
}
```

## Direct properties

- [site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/azure_vnet_site/site/): complete subsection reference.

- [vnet_attachments](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/azure_vnet_site/vnet_attachments/): complete subsection reference.

## Next pages

- [azure_vnet_site.site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/azure_vnet_site/site/)
- [azure_vnet_site.vnet_attachments](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/azure_vnet_site/vnet_attachments/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/)
- [xcsh_cloud_connect](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/)
