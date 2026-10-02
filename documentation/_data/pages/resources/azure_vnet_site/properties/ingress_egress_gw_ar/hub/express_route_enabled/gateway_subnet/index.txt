---
page_title: "ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet"
subcategory: "Infrastructure"
description: "Parameters for Azure subnet."
xcsh_docs: {"aliases": ["ingress egress gw ar hub express route enabled gateway subnet"], "body_bytes": 3464, "body_sha256": "sha256:2ffd89e38eb23fd19d7101517b441154aef6174a0ec66752940b64d4b0e234d0", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:gateway_subnet:auto", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:gateway_subnet:subnet", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:gateway_subnet:subnet_param"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:gateway_subnet", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled", "path": "documentation/resources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/gateway_subnet/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3213100012220112-3132130321000300-2200303301211013-0231210003133130-3320101132223312-3113123112213233-0120032210312300-3200212200020121", "registry_path": "docs/guides/resources--azure_vnet_site--reference--group-006.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet:ConflictingObjectAttributes:auto,subnet", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:gateway_subnet:auto", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet:ConflictingObjectAttributes:auto,subnet_param", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:gateway_subnet:auto", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet:ConflictingObjectAttributes:auto,subnet", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:gateway_subnet:subnet", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet:ConflictingObjectAttributes:subnet,subnet_param", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:gateway_subnet:subnet", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet:ConflictingObjectAttributes:auto,subnet_param", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:gateway_subnet:subnet_param", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet:ConflictingObjectAttributes:subnet,subnet_param", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:gateway_subnet:subnet_param", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_egress_gw_ar", "hub", "express_route_enabled", "gateway_subnet"], "schema_version": 1, "sections": [{"aliases": ["auto"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:gateway_subnet:auto", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw_ar", "hub", "express_route_enabled", "gateway_subnet", "auto"], "syntax": "attribute", "type": "object"}, {"aliases": ["subnet"], "anchor": "section", "description": "Parameters for Azure special subnet which name is reserved. (i.e GatewaySubnet or RouteServerSubnet)", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:gateway_subnet:subnet", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-ingress_egress_gw_ar--hub--express_route_enabled--gateway_subnet--subnet--subnet_resource_grp", "enforcement": "provider-schema", "group": "ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet:ConflictingObjectAttributes:subnet_resource_grp,vnet_resource_group", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:gateway_subnet:subnet", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet:ConflictingObjectAttributes:subnet_resource_grp,vnet_resource_group", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:gateway_subnet:subnet:vnet_resource_group", "type": "conflicts"}], "schema_path": ["ingress_egress_gw_ar", "hub", "express_route_enabled", "gateway_subnet", "subnet"], "syntax": "block", "type": "object"}, {"aliases": ["subnet param"], "anchor": "section", "description": "Parameters for creating a new cloud subnet.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:gateway_subnet:subnet_param", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-ingress_egress_gw_ar--hub--express_route_enabled--gateway_subnet--subnet_param--ipv4", "enforcement": "provider-schema", "group": "ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet_param:RequiredObjectAttributes:ipv4", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:gateway_subnet:subnet_param", "type": "requires"}], "schema_path": ["ingress_egress_gw_ar", "hub", "express_route_enabled", "gateway_subnet", "subnet_param"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/gateway_subnet/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Parameters for Azure subnet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/)
- [ingress_egress_gw_ar](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/)
- [ingress_egress_gw_ar.hub](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/)
- [ingress_egress_gw_ar.hub.express_route_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/)
- ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for gateway subnet.

Upstream description:

Parameters for Azure subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("auto",
    "subnet"),
  validators.ConflictingObjectAttributes("auto",
    "subnet_param"),
  validators.ConflictingObjectAttributes("subnet",
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
  "x-ves-oneof-field-choice": "[\"auto\",\"subnet\",\"subnet_param\"]"
}
```

Terraform syntax:

```terraform
gateway_subnet {
  # Configure direct properties listed below.
}
```

## Direct properties

- [auto](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/gateway_subnet/auto/): complete subsection reference.

- [subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/gateway_subnet/subnet/): complete subsection reference.

- [subnet_param](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/gateway_subnet/subnet_param/): complete subsection reference.

## Next pages

- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.auto](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/gateway_subnet/auto/)
- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/gateway_subnet/subnet/)
- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet_param](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/gateway_subnet/subnet_param/)
- [ingress_egress_gw_ar.hub.express_route_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
