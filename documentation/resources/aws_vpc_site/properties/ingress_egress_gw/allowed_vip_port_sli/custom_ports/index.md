---
page_title: "ingress_egress_gw.allowed_vip_port_sli.custom_ports"
subcategory: "Infrastructure"
description: "List of Custom port."
xcsh_docs: {"aliases": ["ingress egress gw allowed vip port sli custom ports"], "body_bytes": 2964, "body_sha256": "sha256:1293c987ca1332edef9c9a8606b8ff06f21bce471c43d6e991843bf412fcc647", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:allowed_vip_port_sli:custom_ports", "parent_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:allowed_vip_port_sli", "path": "documentation/resources/aws_vpc_site/properties/ingress_egress_gw/allowed_vip_port_sli/custom_ports/index.md", "product": "distributed-cloud", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3122133222012102-0021200231323322-1120133023230031-3030110022211130-1301323331203130-3322220112203132-0100123003110132-2023111212111132", "registry_path": "docs/guides/resources--aws_vpc_site--reference--group-002.md", "relationships": [{"anchor": "schema-ingress_egress_gw--allowed_vip_port_sli--custom_ports--port_ranges", "enforcement": "provider-schema", "group": "ingress_egress_gw.allowed_vip_port_sli.custom_ports:RequiredObjectAttributes:port_ranges", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:allowed_vip_port_sli:custom_ports", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_egress_gw", "allowed_vip_port_sli", "custom_ports"], "schema_version": 1, "sections": [{"aliases": ["port ranges"], "anchor": "schema-ingress_egress_gw--allowed_vip_port_sli--custom_ports--port_ranges", "description": "Port Ranges.", "document_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:allowed_vip_port_sli:custom_ports", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "allowed_vip_port_sli", "custom_ports", "port_ranges"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_vpc_site/properties/ingress_egress_gw/allowed_vip_port_sli/custom_ports/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of Custom port.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.allowed_vip_port_sli.custom_ports

Breadcrumbs:

- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/)
- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/)
- [ingress_egress_gw.allowed_vip_port_sli](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/allowed_vip_port_sli/)
- ingress_egress_gw.allowed_vip_port_sli.custom_ports

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Custom Ports. List of Custom port.

Upstream description:

List of Custom port.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("port_ranges")}
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
custom_ports {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-ingress_egress_gw--allowed_vip_port_sli--custom_ports--port_ranges"></a>

### port_ranges property

Type: `"string"`. Optional.

Port Ranges. Port Ranges.

Upstream description:

Port Ranges.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 512),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
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
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range_list": "true"
  }
}
```

## Next pages

- [ingress_egress_gw.allowed_vip_port_sli](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/allowed_vip_port_sli/)
- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/)
