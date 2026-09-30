---
page_title: "ingress_egress_gw_ar.node.outside_subnet.subnet_param"
subcategory: "Infrastructure"
description: "ingress_egress_gw_ar.node.outside_subnet.subnet_param for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 2765, "body_sha256": "sha256:e086a91317a70143050b56b312f1ac1236d1c3d521f7e4b3b14bec176a7c15cc", "child_ids": [], "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:node:outside_subnet:subnet_param", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:node:outside_subnet", "path": "documentation/resources/azure_vnet_site/properties/ingress_egress_gw_ar/node/outside_subnet/subnet_param/index.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["ingress_egress_gw_ar", "node", "outside_subnet", "subnet_param"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/ingress_egress_gw_ar/node/outside_subnet/subnet_param/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw_ar.node.outside_subnet.subnet_param for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# ingress_egress_gw_ar.node.outside_subnet.subnet_param

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/)
- [ingress_egress_gw_ar](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/)
- [ingress_egress_gw_ar.node](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/node/)
- [ingress_egress_gw_ar.node.outside_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/node/outside_subnet/)
- ingress_egress_gw_ar.node.outside_subnet.subnet_param

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

<a id="schema-ingress_egress_gw_ar--node--outside_subnet--subnet_param--ipv4"></a>

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

- [ingress_egress_gw_ar.node.outside_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/node/outside_subnet/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
