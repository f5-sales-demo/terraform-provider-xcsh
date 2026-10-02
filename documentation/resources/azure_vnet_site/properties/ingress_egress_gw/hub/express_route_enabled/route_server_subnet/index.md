---
page_title: "ingress_egress_gw.hub.express_route_enabled.route_server_subnet"
subcategory: "Infrastructure"
description: "Parameters for Azure subnet."
xcsh_docs: {"aliases": ["ingress egress gw hub express route enabled route server subnet"], "body_bytes": 3472, "body_sha256": "sha256:3415797b2fdc004f19a11d8fccc514af50816827b9e43fdff8fa54ee0ee1e7e4", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:express_route_enabled:route_server_subnet:auto", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:express_route_enabled:route_server_subnet:subnet", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:express_route_enabled:route_server_subnet:subnet_param"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:express_route_enabled:route_server_subnet", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:express_route_enabled", "path": "documentation/resources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/route_server_subnet/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2103203211303022-1313020201021223-1000201211311301-0322301320330323-0002101320001203-0103230211020221-1203100233321101-2321133213031022", "registry_path": "docs/guides/resources--azure_vnet_site--reference--group-004.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw.hub.express_route_enabled.route_server_subnet:ConflictingObjectAttributes:auto,subnet", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:express_route_enabled:route_server_subnet:auto", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw.hub.express_route_enabled.route_server_subnet:ConflictingObjectAttributes:auto,subnet_param", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:express_route_enabled:route_server_subnet:auto", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw.hub.express_route_enabled.route_server_subnet:ConflictingObjectAttributes:auto,subnet", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:express_route_enabled:route_server_subnet:subnet", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw.hub.express_route_enabled.route_server_subnet:ConflictingObjectAttributes:subnet,subnet_param", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:express_route_enabled:route_server_subnet:subnet", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw.hub.express_route_enabled.route_server_subnet:ConflictingObjectAttributes:auto,subnet_param", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:express_route_enabled:route_server_subnet:subnet_param", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw.hub.express_route_enabled.route_server_subnet:ConflictingObjectAttributes:subnet,subnet_param", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:express_route_enabled:route_server_subnet:subnet_param", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_egress_gw", "hub", "express_route_enabled", "route_server_subnet"], "schema_version": 1, "sections": [{"aliases": ["auto"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:express_route_enabled:route_server_subnet:auto", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "hub", "express_route_enabled", "route_server_subnet", "auto"], "syntax": "attribute", "type": "object"}, {"aliases": ["subnet"], "anchor": "section", "description": "Parameters for Azure special subnet which name is reserved. (i.e GatewaySubnet or RouteServerSubnet)", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:express_route_enabled:route_server_subnet:subnet", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-ingress_egress_gw--hub--express_route_enabled--route_server_subnet--subnet--subnet_resource_grp", "enforcement": "provider-schema", "group": "ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet:ConflictingObjectAttributes:subnet_resource_grp,vnet_resource_group", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:express_route_enabled:route_server_subnet:subnet", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet:ConflictingObjectAttributes:subnet_resource_grp,vnet_resource_group", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:express_route_enabled:route_server_subnet:subnet:vnet_resource_group", "type": "conflicts"}], "schema_path": ["ingress_egress_gw", "hub", "express_route_enabled", "route_server_subnet", "subnet"], "syntax": "block", "type": "object"}, {"aliases": ["subnet param"], "anchor": "section", "description": "Parameters for creating a new cloud subnet.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:express_route_enabled:route_server_subnet:subnet_param", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-ingress_egress_gw--hub--express_route_enabled--route_server_subnet--subnet_param--ipv4", "enforcement": "provider-schema", "group": "ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet_param:RequiredObjectAttributes:ipv4", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:express_route_enabled:route_server_subnet:subnet_param", "type": "requires"}], "schema_path": ["ingress_egress_gw", "hub", "express_route_enabled", "route_server_subnet", "subnet_param"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/route_server_subnet/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Parameters for Azure subnet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.hub.express_route_enabled.route_server_subnet

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/)
- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/)
- [ingress_egress_gw.hub](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/hub/)
- [ingress_egress_gw.hub.express_route_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/)
- ingress_egress_gw.hub.express_route_enabled.route_server_subnet

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for route server subnet.

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
route_server_subnet {
  # Configure direct properties listed below.
}
```

## Direct properties

- [auto](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/route_server_subnet/auto/): complete subsection reference.

- [subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/route_server_subnet/subnet/): complete subsection reference.

- [subnet_param](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/route_server_subnet/subnet_param/): complete subsection reference.

## Next pages

- [ingress_egress_gw.hub.express_route_enabled.route_server_subnet.auto](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/route_server_subnet/auto/)
- [ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/route_server_subnet/subnet/)
- [ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet_param](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/route_server_subnet/subnet_param/)
- [ingress_egress_gw.hub.express_route_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
