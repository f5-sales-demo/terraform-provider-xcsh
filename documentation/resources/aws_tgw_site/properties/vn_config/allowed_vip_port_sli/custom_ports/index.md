---
page_title: "vn_config.allowed_vip_port_sli.custom_ports"
subcategory: ""
description: "List of Custom port."
xcsh_docs: {"aliases": ["vn config allowed vip port sli custom ports"], "body_bytes": 2592, "body_sha256": "sha256:8167ceacc6afd2a9dc393efd8272ef6c22a187be2762295b6065ab24a6643dd5", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port_sli:custom_ports", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port_sli", "path": "documentation/resources/aws_tgw_site/properties/vn_config/allowed_vip_port_sli/custom_ports/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-3321310231011111-1012213322010001-2302023112230311-3201122333313213-0311112312230121-1121310022231133-0332012002322201-1131002032331331", "registry_path": "docs/guides/resources--aws_tgw_site--reference--group-003.md", "relationships": [{"anchor": "schema-vn_config--allowed_vip_port_sli--custom_ports--port_ranges", "enforcement": "provider-schema", "group": "vn_config.allowed_vip_port_sli.custom_ports:RequiredObjectAttributes:port_ranges", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port_sli:custom_ports", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["vn_config", "allowed_vip_port_sli", "custom_ports"], "schema_version": 1, "sections": [{"aliases": ["vn config allowed vip port sli custom ports port ranges"], "anchor": "schema-vn_config--allowed_vip_port_sli--custom_ports--port_ranges", "description": "Port Ranges.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port_sli:custom_ports", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vn_config", "allowed_vip_port_sli", "custom_ports", "port_ranges"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/vn_config/allowed_vip_port_sli/custom_ports/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "List of Custom port.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vn_config.allowed_vip_port_sli.custom_ports

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/)
- [vn_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/)
- [vn_config.allowed_vip_port_sli](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/allowed_vip_port_sli/)
- vn_config.allowed_vip_port_sli.custom_ports

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Custom Ports. List of Custom port.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="schema-vn_config--allowed_vip_port_sli--custom_ports--port_ranges"></a>

### port_ranges property

Type: `"string"`. Optional.

Port Ranges. Port Ranges.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
