---
page_title: "ingress_egress_gw.allowed_vip_port_sli.custom_ports"
subcategory: "Infrastructure"
description: "List of Custom port."
xcsh_docs: {"aliases": ["ingress egress gw allowed vip port sli custom ports"], "body_bytes": 2569, "body_sha256": "sha256:53ebe306d67516842754ee5e5da28d3d3bf03f0e463d43400ad99245d7127a92", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_egress_gw:allowed_vip_port_sli:custom_ports", "parent_id": "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_egress_gw:allowed_vip_port_sli", "path": "documentation/data-sources/aws_vpc_site/properties/ingress_egress_gw/allowed_vip_port_sli/custom_ports/index.md", "product": "distributed-cloud", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1130112132310010-1220213112223231-0203331332210233-2221031031232113-0302212111012110-3110112121100211-3221131200003122-3330201020311113", "registry_path": "docs/guides/data-sources--aws_vpc_site--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_egress_gw", "allowed_vip_port_sli", "custom_ports"], "schema_version": 1, "sections": [{"aliases": ["port ranges"], "anchor": "schema-ingress_egress_gw--allowed_vip_port_sli--custom_ports--port_ranges", "description": "Port Ranges.", "document_id": "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_egress_gw:allowed_vip_port_sli:custom_ports", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "allowed_vip_port_sli", "custom_ports", "port_ranges"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_vpc_site/properties/ingress_egress_gw/allowed_vip_port_sli/custom_ports/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "List of Custom port.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.allowed_vip_port_sli.custom_ports

Breadcrumbs:

- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/)
- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/)
- [ingress_egress_gw.allowed_vip_port_sli](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/allowed_vip_port_sli/)
- ingress_egress_gw.allowed_vip_port_sli.custom_ports

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

<a id="schema-ingress_egress_gw--allowed_vip_port_sli--custom_ports--port_ranges"></a>

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

- [ingress_egress_gw.allowed_vip_port_sli](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/allowed_vip_port_sli/)
- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/)
