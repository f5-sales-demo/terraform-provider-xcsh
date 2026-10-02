---
page_title: "ingress_gw_ar.node.local_subnet.subnet_param"
subcategory: "Infrastructure"
description: "Parameters for creating a new cloud subnet."
xcsh_docs: {"aliases": ["ingress gw ar node local subnet subnet param"], "body_bytes": 2773, "body_sha256": "sha256:90c9affe47e17186368975f555be447d728dea54fd9034bff09b681a57e56a0b", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw_ar:node:local_subnet:subnet_param", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw_ar:node:local_subnet", "path": "documentation/resources/azure_vnet_site/properties/ingress_gw_ar/node/local_subnet/subnet_param/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1102132030323030-3111130302032110-2012332231330301-1223111211000233-2230123120200211-2221331030322201-2000022111322133-3010013110301311", "registry_path": "docs/guides/resources--azure_vnet_site--reference--group-008.md", "relationships": [{"anchor": "schema-ingress_gw_ar--node--local_subnet--subnet_param--ipv4", "enforcement": "provider-schema", "group": "ingress_gw_ar.node.local_subnet.subnet_param:RequiredObjectAttributes:ipv4", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw_ar:node:local_subnet:subnet_param", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_gw_ar", "node", "local_subnet", "subnet_param"], "schema_version": 1, "sections": [{"aliases": ["ipv4"], "anchor": "schema-ingress_gw_ar--node--local_subnet--subnet_param--ipv4", "description": "IPv4 subnet prefix for this subnet.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw_ar:node:local_subnet:subnet_param", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_gw_ar", "node", "local_subnet", "subnet_param", "ipv4"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/ingress_gw_ar/node/local_subnet/subnet_param/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Parameters for creating a new cloud subnet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_gw_ar.node.local_subnet.subnet_param

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/)
- [ingress_gw_ar](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_gw_ar/)
- [ingress_gw_ar.node](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_gw_ar/node/)
- [ingress_gw_ar.node.local_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_gw_ar/node/local_subnet/)
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

- [ingress_gw_ar.node.local_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_gw_ar/node/local_subnet/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
