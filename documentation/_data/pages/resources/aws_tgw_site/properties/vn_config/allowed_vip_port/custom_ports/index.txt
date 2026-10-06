---
page_title: "vn_config.allowed_vip_port.custom_ports"
subcategory: ""
description: "List of Custom port."
xcsh_docs: {"aliases": ["vn config allowed vip port custom ports"], "body_bytes": 2572, "body_sha256": "sha256:f1e6707ed0302c4980f3ec578df183bf50ef5bb451f4e4b6d67e81e61ef4cd06", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port:custom_ports", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port", "path": "documentation/resources/aws_tgw_site/properties/vn_config/allowed_vip_port/custom_ports/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-3221131113321332-3102323112131130-3300330122232331-3230201332010100-3102320020231131-1000233201321320-3032230022020103-3112211302031212", "registry_path": "docs/guides/resources--aws_tgw_site--reference--group-003.md", "relationships": [{"anchor": "schema-vn_config--allowed_vip_port--custom_ports--port_ranges", "enforcement": "provider-schema", "group": "vn_config.allowed_vip_port.custom_ports:RequiredObjectAttributes:port_ranges", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port:custom_ports", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["vn_config", "allowed_vip_port", "custom_ports"], "schema_version": 1, "sections": [{"aliases": ["vn config allowed vip port custom ports port ranges"], "anchor": "schema-vn_config--allowed_vip_port--custom_ports--port_ranges", "description": "Port Ranges.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port:custom_ports", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vn_config", "allowed_vip_port", "custom_ports", "port_ranges"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/vn_config/allowed_vip_port/custom_ports/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "List of Custom port.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vn_config.allowed_vip_port.custom_ports

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/)
- [vn_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/)
- [vn_config.allowed_vip_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/allowed_vip_port/)
- vn_config.allowed_vip_port.custom_ports

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

<a id="schema-vn_config--allowed_vip_port--custom_ports--port_ranges"></a>

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
