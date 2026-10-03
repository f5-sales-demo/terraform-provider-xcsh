---
page_title: "ingress_egress_gw.inside_subnet.existing_subnet"
subcategory: "Infrastructure"
description: "Name of existing GCP subnet."
xcsh_docs: {"aliases": ["ingress egress gw inside subnet existing subnet"], "body_bytes": 2874, "body_sha256": "sha256:b72eac175cda9c700b8e5a27de4e4f65c0415c443ad08a3884a7aaac96631495", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:inside_subnet:existing_subnet", "parent_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:inside_subnet", "path": "documentation/resources/gcp_vpc_site/properties/ingress_egress_gw/inside_subnet/existing_subnet/index.md", "product": "distributed-cloud", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1000322103211221-1112011201320010-0333201331033220-2231001323112111-0331321301030211-1012100102312313-1010230101033132-3030020121111113", "registry_path": "docs/guides/resources--gcp_vpc_site--reference--group-002.md", "relationships": [{"anchor": "schema-ingress_egress_gw--inside_subnet--existing_subnet--subnet_name", "enforcement": "provider-schema", "group": "ingress_egress_gw.inside_subnet.existing_subnet:RequiredObjectAttributes:subnet_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:inside_subnet:existing_subnet", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_egress_gw", "inside_subnet", "existing_subnet"], "schema_version": 1, "sections": [{"aliases": ["ingress egress gw inside subnet existing subnet subnet name"], "anchor": "schema-ingress_egress_gw--inside_subnet--existing_subnet--subnet_name", "description": "Name of your subnet in VPC network.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:inside_subnet:existing_subnet", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "inside_subnet", "existing_subnet", "subnet_name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/gcp_vpc_site/properties/ingress_egress_gw/inside_subnet/existing_subnet/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Name of existing GCP subnet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.inside_subnet.existing_subnet

Breadcrumbs:

- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/)
- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/)
- [ingress_egress_gw.inside_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/inside_subnet/)
- ingress_egress_gw.inside_subnet.existing_subnet

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for existing subnet.

Upstream description:

Name of existing GCP subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("subnet_name")}
```

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
existing_subnet {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-ingress_egress_gw--inside_subnet--existing_subnet--subnet_name"></a>

### subnet_name property

Type: `"string"`. Optional.

VPC Subnet Name. Name of your subnet in VPC network.

Upstream description:

Name of your subnet in VPC network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

## Next pages

- [ingress_egress_gw.inside_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/inside_subnet/)
- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/)
