---
page_title: "ingress_egress_gw.accelerated_networking"
subcategory: "Infrastructure"
description: "Accelerated Networking to reduce Latency, When Mode is toggled, traffic disruption will be seen."
xcsh_docs: {"aliases": ["ingress egress gw accelerated networking"], "body_bytes": 2178, "body_sha256": "sha256:74e65b64f393e5d0775f922a2d844b6bfa448f80e30354a04c74a819c2949da7", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:accelerated_networking:disable_spec", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:accelerated_networking:enable"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:accelerated_networking", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw", "path": "documentation/data-sources/azure_vnet_site/properties/ingress_egress_gw/accelerated_networking/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2201000223103010-1013230101201333-3001201223022133-1320232020230222-1203030001031101-1230210211021200-3203320130022310-0002211203203220", "registry_path": "docs/guides/data-sources--azure_vnet_site--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_egress_gw", "accelerated_networking"], "schema_version": 1, "sections": [{"aliases": ["ingress egress gw accelerated networking disable spec"], "anchor": "section", "description": "Enable this option", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:accelerated_networking:disable_spec", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "accelerated_networking", "disable_spec"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress egress gw accelerated networking enable"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:accelerated_networking:enable", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "accelerated_networking", "enable"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/ingress_egress_gw/accelerated_networking/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Accelerated Networking to reduce Latency, When Mode is toggled, traffic disruption will be seen.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.accelerated_networking

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/)
- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/)
- ingress_egress_gw.accelerated_networking

<a id="section"></a>

Type: `"single"`. Computed.

Accelerated Networking to reduce Latency, When Mode is toggled, traffic disruption will be seen.

Upstream description:

Accelerated Networking to reduce Latency, When Mode is toggled, traffic disruption will be seen.

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

## Direct properties

- [disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/accelerated_networking/disable_spec/): complete subsection reference.

- [enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/accelerated_networking/enable/): complete subsection reference.

## Next pages

- [ingress_egress_gw.accelerated_networking.disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/accelerated_networking/disable_spec/)
- [ingress_egress_gw.accelerated_networking.enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/accelerated_networking/enable/)
- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
