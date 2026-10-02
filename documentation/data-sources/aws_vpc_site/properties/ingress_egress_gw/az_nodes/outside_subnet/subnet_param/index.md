---
page_title: "ingress_egress_gw.az_nodes.outside_subnet.subnet_param"
subcategory: "Infrastructure"
description: "Parameters for creating a new cloud subnet."
xcsh_docs: {"aliases": ["ingress egress gw az nodes outside subnet subnet param"], "body_bytes": 2597, "body_sha256": "sha256:9273e48e327de94a6dcb836dcf3d7511fec1ec0d60d21a43263edf4ddbbee36e", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_egress_gw:az_nodes:outside_subnet:subnet_param", "parent_id": "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_egress_gw:az_nodes:outside_subnet", "path": "documentation/data-sources/aws_vpc_site/properties/ingress_egress_gw/az_nodes/outside_subnet/subnet_param/index.md", "product": "distributed-cloud", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3132211011121030-2333230030220010-0211313003321033-2230300122032223-3010000301332002-1301323302202111-1032113111223002-3110021013313012", "registry_path": "docs/guides/data-sources--aws_vpc_site--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_egress_gw", "az_nodes", "outside_subnet", "subnet_param"], "schema_version": 1, "sections": [{"aliases": ["ipv4"], "anchor": "schema-ingress_egress_gw--az_nodes--outside_subnet--subnet_param--ipv4", "description": "IPv4 subnet prefix for this subnet.", "document_id": "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_egress_gw:az_nodes:outside_subnet:subnet_param", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "az_nodes", "outside_subnet", "subnet_param", "ipv4"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_vpc_site/properties/ingress_egress_gw/az_nodes/outside_subnet/subnet_param/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Parameters for creating a new cloud subnet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.az_nodes.outside_subnet.subnet_param

Breadcrumbs:

- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/)
- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/)
- [ingress_egress_gw.az_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/az_nodes/)
- [ingress_egress_gw.az_nodes.outside_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/az_nodes/outside_subnet/)
- ingress_egress_gw.az_nodes.outside_subnet.subnet_param

<a id="section"></a>

Type: `"single"`. Computed.

Parameters for creating a new cloud subnet.

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

<a id="schema-ingress_egress_gw--az_nodes--outside_subnet--subnet_param--ipv4"></a>

### ipv4 property

Type: `"string"`. Computed.

IPv4 Subnet. IPv4 subnet prefix for this subnet.

Upstream description:

IPv4 subnet prefix for this subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  }
}
```

## Next pages

- [ingress_egress_gw.az_nodes.outside_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/az_nodes/outside_subnet/)
- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/)
