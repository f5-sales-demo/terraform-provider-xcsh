---
page_title: "ingress_egress_gw.outside_subnet"
subcategory: "Infrastructure"
description: "This defines choice about GCP VPC network for a view."
xcsh_docs: {"aliases": ["ingress egress gw outside subnet"], "body_bytes": 2285, "body_sha256": "sha256:78689afe4cca1f376455a21d85a79f0d68944a73e98b067261f649afd02d5d63", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:outside_subnet:existing_subnet", "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:outside_subnet:new_subnet"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:outside_subnet", "parent_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw", "path": "documentation/resources/gcp_vpc_site/properties/ingress_egress_gw/outside_subnet/index.md", "product": "distributed-cloud", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3331120001113100-3203200201122222-1030313310121000-2333133202033022-0330313230303203-2301303322122303-1233232211210020-1221010302223301", "registry_path": "docs/guides/resources--gcp_vpc_site--reference--group-003.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw.outside_subnet:ConflictingObjectAttributes:existing_subnet,new_subnet", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:outside_subnet:existing_subnet", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw.outside_subnet:ConflictingObjectAttributes:existing_subnet,new_subnet", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:outside_subnet:new_subnet", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_egress_gw", "outside_subnet"], "schema_version": 1, "sections": [{"aliases": ["existing subnet"], "anchor": "section", "description": "Name of existing GCP subnet.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:outside_subnet:existing_subnet", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-ingress_egress_gw--outside_subnet--existing_subnet--subnet_name", "enforcement": "provider-schema", "group": "ingress_egress_gw.outside_subnet.existing_subnet:RequiredObjectAttributes:subnet_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:outside_subnet:existing_subnet", "type": "requires"}], "schema_path": ["ingress_egress_gw", "outside_subnet", "existing_subnet"], "syntax": "block", "type": "object"}, {"aliases": ["new subnet"], "anchor": "section", "description": "Parameters for GCP subnet.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:outside_subnet:new_subnet", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-ingress_egress_gw--outside_subnet--new_subnet--primary_ipv4", "enforcement": "provider-schema", "group": "ingress_egress_gw.outside_subnet.new_subnet:RequiredObjectAttributes:primary_ipv4", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:outside_subnet:new_subnet", "type": "requires"}], "schema_path": ["ingress_egress_gw", "outside_subnet", "new_subnet"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/gcp_vpc_site/properties/ingress_egress_gw/outside_subnet/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This defines choice about GCP VPC network for a view.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.outside_subnet

Breadcrumbs:

- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/)
- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/)
- ingress_egress_gw.outside_subnet

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Defines choice about GCP VPC network for a view.

Upstream description:

This defines choice about GCP VPC network for a view.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("existing_subnet",
    "new_subnet")}
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
  "x-ves-oneof-field-choice": "[\"existing_subnet\",\"new_subnet\"]"
}
```

Terraform syntax:

```terraform
outside_subnet {
  # Configure direct properties listed below.
}
```

## Direct properties

- [existing_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/outside_subnet/existing_subnet/): complete subsection reference.

- [new_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/outside_subnet/new_subnet/): complete subsection reference.

## Next pages

- [ingress_egress_gw.outside_subnet.existing_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/outside_subnet/existing_subnet/)
- [ingress_egress_gw.outside_subnet.new_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/outside_subnet/new_subnet/)
- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/)
- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/)
