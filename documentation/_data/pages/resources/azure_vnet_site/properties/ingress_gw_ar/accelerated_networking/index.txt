---
page_title: "ingress_gw_ar.accelerated_networking"
subcategory: "Infrastructure"
description: "Accelerated Networking to reduce Latency, When Mode is toggled, traffic disruption will be seen."
xcsh_docs: {"aliases": ["ingress gw ar accelerated networking"], "body_bytes": 2402, "body_sha256": "sha256:a74620d0e46c35bcf1fb4ee8afd2b11e7090e028d43c0d3d26342b36c788e953", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:azure_vnet_site:properties:ingress_gw_ar:accelerated_networking:disable_spec", "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw_ar:accelerated_networking:enable"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw_ar:accelerated_networking", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw_ar", "path": "documentation/resources/azure_vnet_site/properties/ingress_gw_ar/accelerated_networking/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2200003022200020-3321101331013131-3231222121120031-0230130103103200-3213023223323122-0122232103232303-3023103103121130-2122203210113132", "registry_path": "docs/guides/resources--azure_vnet_site--reference--group-008.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw_ar.accelerated_networking:ConflictingObjectAttributes:disable_spec,enable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw_ar:accelerated_networking:disable_spec", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw_ar.accelerated_networking:ConflictingObjectAttributes:disable_spec,enable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw_ar:accelerated_networking:enable", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_gw_ar", "accelerated_networking"], "schema_version": 1, "sections": [{"aliases": ["disable spec"], "anchor": "section", "description": "Enable this option", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw_ar:accelerated_networking:disable_spec", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_gw_ar", "accelerated_networking", "disable_spec"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw_ar:accelerated_networking:enable", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_gw_ar", "accelerated_networking", "enable"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/ingress_gw_ar/accelerated_networking/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Accelerated Networking to reduce Latency, When Mode is toggled, traffic disruption will be seen.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_gw_ar.accelerated_networking

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/)
- [ingress_gw_ar](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_gw_ar/)
- ingress_gw_ar.accelerated_networking

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Accelerated Networking to reduce Latency, When Mode is toggled, traffic disruption will be seen.

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

- [disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_gw_ar/accelerated_networking/disable_spec/): complete subsection reference.

- [enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_gw_ar/accelerated_networking/enable/): complete subsection reference.

## Next pages

- [ingress_gw_ar.accelerated_networking.disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_gw_ar/accelerated_networking/disable_spec/)
- [ingress_gw_ar.accelerated_networking.enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_gw_ar/accelerated_networking/enable/)
- [ingress_gw_ar](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_gw_ar/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
