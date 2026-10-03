---
page_title: "ingress_gw.local_subnet"
subcategory: "Infrastructure"
description: "This defines choice about GCP VPC network for a view."
xcsh_docs: {"aliases": ["ingress gw local subnet"], "body_bytes": 1914, "body_sha256": "sha256:45b19712138e571ccdecce663dca8f521336e36a151d6c7419fa239cde9d9130", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_gw:local_subnet:existing_subnet", "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_gw:local_subnet:new_subnet"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_gw:local_subnet", "parent_id": "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_gw", "path": "documentation/data-sources/gcp_vpc_site/properties/ingress_gw/local_subnet/index.md", "product": "distributed-cloud", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2312120310333230-0311320001331211-3031012013232011-2020203221300223-2333202232121112-3122102111000210-1320331010012303-2223100320103202", "registry_path": "docs/guides/data-sources--gcp_vpc_site--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_gw", "local_subnet"], "schema_version": 1, "sections": [{"aliases": ["ingress gw local subnet existing subnet"], "anchor": "section", "description": "Name of existing GCP subnet.", "document_id": "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_gw:local_subnet:existing_subnet", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_gw", "local_subnet", "existing_subnet"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress gw local subnet new subnet"], "anchor": "section", "description": "Parameters for GCP subnet.", "document_id": "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_gw:local_subnet:new_subnet", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_gw", "local_subnet", "new_subnet"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/gcp_vpc_site/properties/ingress_gw/local_subnet/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This defines choice about GCP VPC network for a view.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_gw.local_subnet

Breadcrumbs:

- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/)
- [ingress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/ingress_gw/)
- ingress_gw.local_subnet

<a id="section"></a>

Type: `"single"`. Computed.

Defines choice about GCP VPC network for a view.

Upstream description:

This defines choice about GCP VPC network for a view.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"existing_subnet\",\"new_subnet\"]"
}
```

## Direct properties

- [existing_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/ingress_gw/local_subnet/existing_subnet/): complete subsection reference.

- [new_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/ingress_gw/local_subnet/new_subnet/): complete subsection reference.

## Next pages

- [ingress_gw.local_subnet.existing_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/ingress_gw/local_subnet/existing_subnet/)
- [ingress_gw.local_subnet.new_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/ingress_gw/local_subnet/new_subnet/)
- [ingress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/ingress_gw/)
- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/)
