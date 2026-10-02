---
page_title: "ingress_gw.az_nodes.local_subnet.subnet_param"
subcategory: "Infrastructure"
description: "Parameters for creating a new cloud subnet."
xcsh_docs: {"aliases": ["ingress gw az nodes local subnet subnet param"], "body_bytes": 2506, "body_sha256": "sha256:85a6d69eac82d8616469057a165750bf1e464c2162c97d10be3cb86f67baf9dc", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_gw:az_nodes:local_subnet:subnet_param", "parent_id": "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_gw:az_nodes:local_subnet", "path": "documentation/data-sources/aws_vpc_site/properties/ingress_gw/az_nodes/local_subnet/subnet_param/index.md", "product": "distributed-cloud", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1203313031012233-0321002102032011-3230233312011210-1213221332232333-1110110220003022-1111013010021010-2312201232130133-2021331033221213", "registry_path": "docs/guides/data-sources--aws_vpc_site--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_gw", "az_nodes", "local_subnet", "subnet_param"], "schema_version": 1, "sections": [{"aliases": ["ipv4"], "anchor": "schema-ingress_gw--az_nodes--local_subnet--subnet_param--ipv4", "description": "IPv4 subnet prefix for this subnet.", "document_id": "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_gw:az_nodes:local_subnet:subnet_param", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_gw", "az_nodes", "local_subnet", "subnet_param", "ipv4"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_vpc_site/properties/ingress_gw/az_nodes/local_subnet/subnet_param/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Parameters for creating a new cloud subnet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_gw.az_nodes.local_subnet.subnet_param

Breadcrumbs:

- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/)
- [ingress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_gw/)
- [ingress_gw.az_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_gw/az_nodes/)
- [ingress_gw.az_nodes.local_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_gw/az_nodes/local_subnet/)
- ingress_gw.az_nodes.local_subnet.subnet_param

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

<a id="schema-ingress_gw--az_nodes--local_subnet--subnet_param--ipv4"></a>

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

- [ingress_gw.az_nodes.local_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_gw/az_nodes/local_subnet/)
- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/)
