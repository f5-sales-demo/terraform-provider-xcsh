---
page_title: "ingress_gw_ar.node.local_subnet.subnet_param"
subcategory: "Infrastructure"
description: "ingress_gw_ar.node.local_subnet.subnet_param for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 2419, "body_sha256": "sha256:6c8e8c2c464fa731eb44ed79c594010917434fccc955dd3d5ca42d2811c8f452", "canonical_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw_ar:node:local_subnet:subnet_param", "child_ids": [], "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw_ar:node:local_subnet:subnet_param", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw_ar:node:local_subnet", "path": "docs/guides/resources--azure_vnet_site--properties--ingress_gw_ar--node--local_subnet--subnet_param.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_gw_ar", "node", "local_subnet", "subnet_param"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/ingress_gw_ar/node/local_subnet/subnet_param/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_gw_ar.node.local_subnet.subnet_param for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_gw_ar.node.local_subnet.subnet_param

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
- [Property reference](resources--azure_vnet_site--reference.md)
- [ingress_gw_ar](resources--azure_vnet_site--properties--ingress_gw_ar.md)
- [ingress_gw_ar.node](resources--azure_vnet_site--properties--ingress_gw_ar--node.md)
- [ingress_gw_ar.node.local_subnet](resources--azure_vnet_site--properties--ingress_gw_ar--node--local_subnet.md)
- ingress_gw_ar.node.local_subnet.subnet_param

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

<a id="schema-ingress_gw_ar--node--local_subnet--subnet_param--ipv4"></a>

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

- [ingress_gw_ar.node.local_subnet](resources--azure_vnet_site--properties--ingress_gw_ar--node--local_subnet.md)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
