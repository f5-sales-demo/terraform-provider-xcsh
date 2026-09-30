---
page_title: "vn_config.allowed_vip_port_sli.custom_ports"
subcategory: ""
description: "vn_config.allowed_vip_port_sli.custom_ports for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": [], "body_bytes": 2398, "body_sha256": "sha256:d173ac852c9ff6cec41766db133228ef1a4c902ecf3b4912c0372940d2bbbda7", "child_ids": [], "collection_id": "xcsh-docs:data-sources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:allowed_vip_port_sli:custom_ports", "parent_id": "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:allowed_vip_port_sli", "path": "documentation/data-sources/aws_tgw_site/properties/vn_config/allowed_vip_port_sli/custom_ports/index.md", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["vn_config", "allowed_vip_port_sli", "custom_ports"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_tgw_site/properties/vn_config/allowed_vip_port_sli/custom_ports/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "vn_config.allowed_vip_port_sli.custom_ports for xcsh_aws_tgw_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# vn_config.allowed_vip_port_sli.custom_ports

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/)
- [vn_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/)
- [vn_config.allowed_vip_port_sli](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/allowed_vip_port_sli/)
- vn_config.allowed_vip_port_sli.custom_ports

<a id="section"></a>

Type: `"single"`. Computed.

Custom Ports. List of Custom port.

Upstream description:

List of Custom port.

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

## Direct properties

<a id="schema-vn_config--allowed_vip_port_sli--custom_ports--port_ranges"></a>

### port_ranges property

Type: `"string"`. Computed.

Port Ranges. Port Ranges.

Upstream description:

Port Ranges.

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

- [vn_config.allowed_vip_port_sli](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/allowed_vip_port_sli/)
- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/)
