---
page_title: "site_subnet_params"
subcategory: ""
description: "Configure subnet parameters per site."
xcsh_docs: {"aliases": ["site subnet params"], "body_bytes": 3187, "body_sha256": "sha256:379250dec0e5226dbe4b50b13caaee5fb09df0db6ba620442eee09cfa77da1d0", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:subnet:properties:site_subnet_params:dhcp", "xcsh-docs:resources:subnet:properties:site_subnet_params:site", "xcsh-docs:resources:subnet:properties:site_subnet_params:static_ip", "xcsh-docs:resources:subnet:properties:site_subnet_params:subnet_dhcp_server_params"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:subnet:collection", "completeness": "complete", "id": "xcsh-docs:resources:subnet:properties:site_subnet_params", "parent_id": "xcsh-docs:resources:subnet:reference", "path": "documentation/resources/subnet/properties/site_subnet_params/index.md", "product": "distributed-cloud", "provider_name": "subnet", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0311002121020233-0121010300123320-3332310122013110-1001032303211011-3120300131202223-1122230330101010-1133333103020011-0333012002330013", "registry_path": "docs/guides/resources--subnet--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "site_subnet_params:ConflictingListObjectAttributes:dhcp,static_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:subnet:properties:site_subnet_params:dhcp", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "site_subnet_params:ConflictingListObjectAttributes:dhcp,static_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:subnet:properties:site_subnet_params:static_ip", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["site_subnet_params"], "schema_version": 1, "sections": [{"aliases": ["site subnet params dhcp"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:subnet:properties:site_subnet_params:dhcp", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["site_subnet_params", "dhcp"], "syntax": "attribute", "type": "object"}, {"aliases": ["site subnet params site"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:subnet:properties:site_subnet_params:site", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-site_subnet_params--site--name", "enforcement": "provider-schema", "group": "site_subnet_params.site:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:subnet:properties:site_subnet_params:site", "type": "requires"}], "schema_path": ["site_subnet_params", "site"], "syntax": "block", "type": "object"}, {"aliases": ["site subnet params static ip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:subnet:properties:site_subnet_params:static_ip", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["site_subnet_params", "static_ip"], "syntax": "attribute", "type": "object"}, {"aliases": ["site subnet params subnet dhcp server params"], "anchor": "section", "description": "Subnet DHCP parameters will be a subset of network_interface.dhcpserverparameterstype as all features in network_interface.dhcpserverparameterstype may not be supported in a subnet.", "document_id": "xcsh-docs:resources:subnet:properties:site_subnet_params:subnet_dhcp_server_params", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["site_subnet_params", "subnet_dhcp_server_params"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/subnet/properties/site_subnet_params/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Configure subnet parameters per site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["subnetCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# site_subnet_params

Breadcrumbs:

- [xcsh_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/subnet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/subnet/properties/)
- site_subnet_params

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Site Subnet Parameters. Configure subnet parameters per site.

Upstream description:

Configure subnet parameters per site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("dhcp",
    "static_ip")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
site_subnet_params {
  # Configure direct properties listed below.
}
```

## Direct properties

- [dhcp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/subnet/properties/site_subnet_params/dhcp/): complete subsection reference.

- [site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/subnet/properties/site_subnet_params/site/): complete subsection reference.

- [static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/subnet/properties/site_subnet_params/static_ip/): complete subsection reference.

- [subnet_dhcp_server_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/subnet/properties/site_subnet_params/subnet_dhcp_server_params/): complete subsection reference.

## Next pages

- [site_subnet_params.dhcp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/subnet/properties/site_subnet_params/dhcp/)
- [site_subnet_params.site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/subnet/properties/site_subnet_params/site/)
- [site_subnet_params.static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/subnet/properties/site_subnet_params/static_ip/)
- [site_subnet_params.subnet_dhcp_server_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/subnet/properties/site_subnet_params/subnet_dhcp_server_params/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/subnet/properties/)
- [xcsh_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/subnet/)
