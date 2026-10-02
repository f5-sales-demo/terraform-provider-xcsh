---
page_title: "ingress_egress_gw_ar.node.inside_subnet"
subcategory: "Infrastructure"
description: "Parameters for Azure subnet."
xcsh_docs: {"aliases": ["ingress egress gw ar node inside subnet"], "body_bytes": 2473, "body_sha256": "sha256:3c7465dd26ac6ea204e16823a7a80a890db37156136e1c70cddb8ce4880dfba5", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:node:inside_subnet:subnet", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:node:inside_subnet:subnet_param"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:node:inside_subnet", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:node", "path": "documentation/resources/azure_vnet_site/properties/ingress_egress_gw_ar/node/inside_subnet/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3123033001331333-0110023212132131-1211100322223012-1223013001233103-0310122321200010-1320000331033012-1033320103220030-1112123121221121", "registry_path": "docs/guides/resources--azure_vnet_site--reference--group-007.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw_ar.node.inside_subnet:ConflictingObjectAttributes:subnet,subnet_param", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:node:inside_subnet:subnet", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw_ar.node.inside_subnet:ConflictingObjectAttributes:subnet,subnet_param", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:node:inside_subnet:subnet_param", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_egress_gw_ar", "node", "inside_subnet"], "schema_version": 1, "sections": [{"aliases": ["subnet"], "anchor": "section", "description": "Parameters for Azure subnet.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:node:inside_subnet:subnet", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-ingress_egress_gw_ar--node--inside_subnet--subnet--subnet_resource_grp", "enforcement": "provider-schema", "group": "ingress_egress_gw_ar.node.inside_subnet.subnet:ConflictingObjectAttributes:subnet_resource_grp,vnet_resource_group", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:node:inside_subnet:subnet", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw_ar.node.inside_subnet.subnet:ConflictingObjectAttributes:subnet_resource_grp,vnet_resource_group", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:node:inside_subnet:subnet:vnet_resource_group", "type": "conflicts"}, {"anchor": "schema-ingress_egress_gw_ar--node--inside_subnet--subnet--subnet_name", "enforcement": "provider-schema", "group": "ingress_egress_gw_ar.node.inside_subnet.subnet:RequiredObjectAttributes:subnet_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:node:inside_subnet:subnet", "type": "requires"}], "schema_path": ["ingress_egress_gw_ar", "node", "inside_subnet", "subnet"], "syntax": "block", "type": "object"}, {"aliases": ["subnet param"], "anchor": "section", "description": "Parameters for creating a new cloud subnet.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:node:inside_subnet:subnet_param", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-ingress_egress_gw_ar--node--inside_subnet--subnet_param--ipv4", "enforcement": "provider-schema", "group": "ingress_egress_gw_ar.node.inside_subnet.subnet_param:RequiredObjectAttributes:ipv4", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:node:inside_subnet:subnet_param", "type": "requires"}], "schema_path": ["ingress_egress_gw_ar", "node", "inside_subnet", "subnet_param"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/ingress_egress_gw_ar/node/inside_subnet/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Parameters for Azure subnet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw_ar.node.inside_subnet

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/)
- [ingress_egress_gw_ar](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/)
- [ingress_egress_gw_ar.node](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/node/)
- ingress_egress_gw_ar.node.inside_subnet

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for inside subnet.

Upstream description:

Parameters for Azure subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("subnet",
    "subnet_param")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"subnet\",\"subnet_param\"]"
}
```

Terraform syntax:

```terraform
inside_subnet {
  # Configure direct properties listed below.
}
```

## Direct properties

- [subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/node/inside_subnet/subnet/): complete subsection reference.

- [subnet_param](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/node/inside_subnet/subnet_param/): complete subsection reference.

## Next pages

- [ingress_egress_gw_ar.node.inside_subnet.subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/node/inside_subnet/subnet/)
- [ingress_egress_gw_ar.node.inside_subnet.subnet_param](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/node/inside_subnet/subnet_param/)
- [ingress_egress_gw_ar.node](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/node/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
