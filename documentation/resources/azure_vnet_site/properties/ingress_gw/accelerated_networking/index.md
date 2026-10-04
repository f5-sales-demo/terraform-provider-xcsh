---
page_title: "ingress_gw.accelerated_networking"
subcategory: "Infrastructure"
description: "Accelerated Networking to reduce Latency, When Mode is toggled, traffic disruption will be seen."
xcsh_docs: {"aliases": ["ingress gw accelerated networking"], "body_bytes": 2403, "body_sha256": "sha256:41989b683c4a8585d1603e5ccbdeb63dd5517f64c9cabe0d611ba784eea909d4", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:azure_vnet_site:properties:ingress_gw:accelerated_networking:disable_spec", "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw:accelerated_networking:enable"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw:accelerated_networking", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw", "path": "documentation/resources/azure_vnet_site/properties/ingress_gw/accelerated_networking/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0131311002000031-1200100123330122-2330213012000120-3122110310221331-3232121203223210-2021122211223001-1001213201330210-1211013212203320", "registry_path": "docs/guides/resources--azure_vnet_site--reference--group-007.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw.accelerated_networking:ConflictingObjectAttributes:disable_spec,enable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw:accelerated_networking:disable_spec", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw.accelerated_networking:ConflictingObjectAttributes:disable_spec,enable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw:accelerated_networking:enable", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_gw", "accelerated_networking"], "schema_version": 1, "sections": [{"aliases": ["ingress gw accelerated networking disable spec"], "anchor": "section", "description": "Enable this option", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw:accelerated_networking:disable_spec", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_gw", "accelerated_networking", "disable_spec"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress gw accelerated networking enable"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw:accelerated_networking:enable", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_gw", "accelerated_networking", "enable"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/ingress_gw/accelerated_networking/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Accelerated Networking to reduce Latency, When Mode is toggled, traffic disruption will be seen.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_gw.accelerated_networking

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/)
- [ingress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_gw/)
- ingress_gw.accelerated_networking

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Accelerated Networking to reduce Latency, When Mode is toggled, traffic disruption will be seen.
Server applies default when omitted.

Upstream description:

Accelerated Networking to reduce Latency, When Mode is toggled, traffic disruption will be seen.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_spec",
    "enable")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-accelerated_networking": "[\"disable\",\"enable\"]"
}
```

Terraform syntax:

```terraform
accelerated_networking {
  # Configure direct properties listed below.
}
```

## Direct properties

- [disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_gw/accelerated_networking/disable_spec/): complete subsection reference.

- [enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_gw/accelerated_networking/enable/): complete subsection reference.

## Next pages

- [ingress_gw.accelerated_networking.disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_gw/accelerated_networking/disable_spec/)
- [ingress_gw.accelerated_networking.enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_gw/accelerated_networking/enable/)
- [ingress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_gw/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
