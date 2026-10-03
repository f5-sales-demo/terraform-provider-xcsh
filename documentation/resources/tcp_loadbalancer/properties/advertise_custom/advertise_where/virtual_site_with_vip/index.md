---
page_title: "advertise_custom.advertise_where.virtual_site_with_vip"
subcategory: "Load Balancing"
description: "This defines a reference to a customer site virtual site along with network type and IP where a load balancer could be advertised."
xcsh_docs: {"aliases": ["advertise custom advertise where virtual site with vip"], "body_bytes": 4296, "body_sha256": "sha256:de16b60b0241a4c69ea5cdf8c272e5b1d84540f92c70f7948b136fe138054a46", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:virtual_site_with_vip:virtual_site"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:tcp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:virtual_site_with_vip", "parent_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where", "path": "documentation/resources/tcp_loadbalancer/properties/advertise_custom/advertise_where/virtual_site_with_vip/index.md", "product": "distributed-cloud", "provider_name": "tcp_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1121201333003001-3211103001030232-1120020013233002-1201212112111322-2313032201232300-2313132333210321-3020122003000202-2203220100301301", "registry_path": "docs/guides/resources--tcp_loadbalancer--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["advertise_custom", "advertise_where", "virtual_site_with_vip"], "schema_version": 1, "sections": [{"aliases": ["advertise custom advertise where virtual site with vip ip"], "anchor": "schema-advertise_custom--advertise_where--virtual_site_with_vip--ip", "description": "Use given IP address as VIP on the site.", "document_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:virtual_site_with_vip", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["advertise_custom", "advertise_where", "virtual_site_with_vip", "ip"], "syntax": "attribute", "type": "string"}, {"aliases": ["advertise custom advertise where virtual site with vip network"], "anchor": "schema-advertise_custom--advertise_where--virtual_site_with_vip--network", "description": "This defines network types to be used on virtual-site with specified VIP All outside networks. All inside networks.", "document_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:virtual_site_with_vip", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["advertise_custom", "advertise_where", "virtual_site_with_vip", "network"], "syntax": "attribute", "type": "string"}, {"aliases": ["advertise custom advertise where virtual site with vip virtual site"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:virtual_site_with_vip:virtual_site", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-advertise_custom--advertise_where--virtual_site_with_vip--virtual_site--name", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where.virtual_site_with_vip.virtual_site:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:virtual_site_with_vip:virtual_site", "type": "requires"}], "schema_path": ["advertise_custom", "advertise_where", "virtual_site_with_vip", "virtual_site"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tcp_loadbalancer/properties/advertise_custom/advertise_where/virtual_site_with_vip/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This defines a reference to a customer site virtual site along with network type and IP where a load balancer could be advertised.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["tcp_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# advertise_custom.advertise_where.virtual_site_with_vip

Breadcrumbs:

- [xcsh_tcp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/)
- [advertise_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/advertise_custom/)
- [advertise_custom.advertise_where](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/advertise_custom/advertise_where/)
- advertise_custom.advertise_where.virtual_site_with_vip

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Defines a reference to a customer site virtual site along with network type and IP where a load
balancer could be advertised.

Upstream description:

This defines a reference to a customer site virtual site along with network type and IP where a load
balancer could be advertised.

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
virtual_site_with_vip {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-advertise_custom--advertise_where--virtual_site_with_vip--ip"></a>

### ip property

Type: `"string"`. Optional.

Use given IP address as VIP on the site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="schema-advertise_custom--advertise_where--virtual_site_with_vip--network"></a>

### network property

Type: `"string"`. Optional.

\[Enum: SITE\_NETWORK\_SPECIFIED\_VIP\_OUTSIDE|SITE\_NETWORK\_SPECIFIED\_VIP\_INSIDE\] Defines
network types to be used on virtual-site with specified VIP All outside networks. All inside
networks. Possible values are \`SITE\_NETWORK\_SPECIFIED\_VIP\_OUTSIDE\`,
\`SITE\_NETWORK\_SPECIFIED\_VIP\_INSIDE\`. Defaults to \`SITE\_NETWORK\_SPECIFIED\_VIP\_OUTSIDE\`.

Upstream description:

This defines network types to be used on virtual-site with specified VIP

All outside networks. All inside networks.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("SITE_NETWORK_SPECIFIED_VIP_OUTSIDE",
    "SITE_NETWORK_SPECIFIED_VIP_INSIDE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "SITE_NETWORK_SPECIFIED_VIP_OUTSIDE",
  "enum": [
    "SITE_NETWORK_SPECIFIED_VIP_OUTSIDE",
    "SITE_NETWORK_SPECIFIED_VIP_INSIDE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/advertise_custom/advertise_where/virtual_site_with_vip/virtual_site/): complete subsection reference.

## Next pages

- [advertise_custom.advertise_where.virtual_site_with_vip.virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/advertise_custom/advertise_where/virtual_site_with_vip/virtual_site/)
- [advertise_custom.advertise_where](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/advertise_custom/advertise_where/)
- [xcsh_tcp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/)
