---
page_title: "ingress_gw_ar.node.local_subnet.subnet"
subcategory: "Infrastructure"
description: "Parameters for Azure subnet."
xcsh_docs: {"aliases": ["ingress gw ar node local subnet subnet"], "body_bytes": 3923, "body_sha256": "sha256:cd6d2930ace507fcbb9d6e323a9a1eb0485561695623bb1a9025361ecb4cba4f", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw_ar:node:local_subnet:subnet:vnet_resource_group"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw_ar:node:local_subnet:subnet", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw_ar:node:local_subnet", "path": "documentation/data-sources/azure_vnet_site/properties/ingress_gw_ar/node/local_subnet/subnet/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0120112130012203-0011223102301103-0102312003333130-3013201100112010-2021322033200020-0233332221002121-0223112121022121-2201212233031210", "registry_path": "docs/guides/data-sources--azure_vnet_site--reference--group-008.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_gw_ar", "node", "local_subnet", "subnet"], "schema_version": 1, "sections": [{"aliases": ["subnet name"], "anchor": "schema-ingress_gw_ar--node--local_subnet--subnet--subnet_name", "description": "Name of existing subnet.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw_ar:node:local_subnet:subnet", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_gw_ar", "node", "local_subnet", "subnet", "subnet_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["subnet resource grp"], "anchor": "schema-ingress_gw_ar--node--local_subnet--subnet--subnet_resource_grp", "description": "Exclusive with Specify name of Resource Group.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw_ar:node:local_subnet:subnet", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_gw_ar", "node", "local_subnet", "subnet", "subnet_resource_grp"], "syntax": "attribute", "type": "string"}, {"aliases": ["vnet resource group"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw_ar:node:local_subnet:subnet:vnet_resource_group", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_gw_ar", "node", "local_subnet", "subnet", "vnet_resource_group"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/ingress_gw_ar/node/local_subnet/subnet/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Parameters for Azure subnet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_gw_ar.node.local_subnet.subnet

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/)
- [ingress_gw_ar](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw_ar/)
- [ingress_gw_ar.node](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw_ar/node/)
- [ingress_gw_ar.node.local_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw_ar/node/local_subnet/)
- ingress_gw_ar.node.local_subnet.subnet

<a id="section"></a>

Type: `"single"`. Computed.

Subnet specification for network segmentation.

Upstream description:

Parameters for Azure subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-resource_group_choice": "[\"subnet_resource_grp\",\"vnet_resource_group\"]"
}
```

## Direct properties

<a id="schema-ingress_gw_ar--node--local_subnet--subnet--subnet_name"></a>

### subnet_name property

Type: `"string"`. Computed.

Subnet Name. Name of existing subnet.

Upstream description:

Name of existing subnet.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="schema-ingress_gw_ar--node--local_subnet--subnet--subnet_resource_grp"></a>

### subnet_resource_grp property

Type: `"string"`. Computed.

Exclusive with \[vnet\_resource\_group\] Specify name of Resource Group.

Upstream description:

Exclusive with \[vnet\_resource\_group\] Specify name of Resource Group.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [vnet_resource_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw_ar/node/local_subnet/subnet/vnet_resource_group/): complete subsection reference.

## Next pages

- [ingress_gw_ar.node.local_subnet.subnet.vnet_resource_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw_ar/node/local_subnet/subnet/vnet_resource_group/)
- [ingress_gw_ar.node.local_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw_ar/node/local_subnet/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
