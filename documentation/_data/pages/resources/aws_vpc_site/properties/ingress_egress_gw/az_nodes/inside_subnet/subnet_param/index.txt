---
page_title: "ingress_egress_gw.az_nodes.inside_subnet.subnet_param"
subcategory: "Infrastructure"
description: "Parameters for creating a new cloud subnet."
xcsh_docs: {"aliases": ["ingress egress gw az nodes inside subnet subnet param"], "body_bytes": 2833, "body_sha256": "sha256:44cc99ff2bc73eb2926aa1c6de5aa7d38b8c8e387fd31fed82ab17e157e63f71", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:az_nodes:inside_subnet:subnet_param", "parent_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:az_nodes:inside_subnet", "path": "documentation/resources/aws_vpc_site/properties/ingress_egress_gw/az_nodes/inside_subnet/subnet_param/index.md", "product": "distributed-cloud", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0022223021120122-1132023300112000-0302312311211020-3313203131202122-1233201202001102-0323020203103133-1332012212301021-0110232130101313", "registry_path": "docs/guides/resources--aws_vpc_site--reference--group-002.md", "relationships": [{"anchor": "schema-ingress_egress_gw--az_nodes--inside_subnet--subnet_param--ipv4", "enforcement": "provider-schema", "group": "ingress_egress_gw.az_nodes.inside_subnet.subnet_param:RequiredObjectAttributes:ipv4", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:az_nodes:inside_subnet:subnet_param", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_egress_gw", "az_nodes", "inside_subnet", "subnet_param"], "schema_version": 1, "sections": [{"aliases": ["ipv4"], "anchor": "schema-ingress_egress_gw--az_nodes--inside_subnet--subnet_param--ipv4", "description": "IPv4 subnet prefix for this subnet.", "document_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:az_nodes:inside_subnet:subnet_param", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "az_nodes", "inside_subnet", "subnet_param", "ipv4"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_vpc_site/properties/ingress_egress_gw/az_nodes/inside_subnet/subnet_param/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Parameters for creating a new cloud subnet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.az_nodes.inside_subnet.subnet_param

Breadcrumbs:

- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/)
- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/)
- [ingress_egress_gw.az_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/az_nodes/)
- [ingress_egress_gw.az_nodes.inside_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/az_nodes/inside_subnet/)
- ingress_egress_gw.az_nodes.inside_subnet.subnet_param

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Parameters for creating a new cloud subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ipv4")}
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
subnet_param {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-ingress_egress_gw--az_nodes--inside_subnet--subnet_param--ipv4"></a>

### ipv4 property

Type: `"string"`. Optional.

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

- [ingress_egress_gw.az_nodes.inside_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/az_nodes/inside_subnet/)
- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/)
