---
page_title: "ingress_egress_gw.hub.spoke_vnets.vnet"
subcategory: "Infrastructure"
description: "Resource group and name of existing Azure VNet."
xcsh_docs: {"aliases": ["ingress egress gw hub spoke vnets vnet"], "body_bytes": 5241, "body_sha256": "sha256:08b489ca9c5bc150238b387ba49f91f2c6f279bd430e266ba4217b0ffe4d8771", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:spoke_vnets:vnet:f5_orchestrated_routing", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:spoke_vnets:vnet:manual_routing"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:spoke_vnets:vnet", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:spoke_vnets", "path": "documentation/resources/azure_vnet_site/properties/ingress_egress_gw/hub/spoke_vnets/vnet/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "resources", "registry_anchor": "canonical-2302031320113300-0301333313323230-0333110232313210-0101232221120320-3231232112322012-2032000320331210-3121323331330321-0111222101233020", "registry_path": "docs/guides/resources--azure_vnet_site--reference--group-004.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw.hub.spoke_vnets.vnet:ConflictingObjectAttributes:f5_orchestrated_routing,manual_routing", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:spoke_vnets:vnet:f5_orchestrated_routing", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw.hub.spoke_vnets.vnet:ConflictingObjectAttributes:f5_orchestrated_routing,manual_routing", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:spoke_vnets:vnet:manual_routing", "type": "conflicts"}, {"anchor": "schema-ingress_egress_gw--hub--spoke_vnets--vnet--resource_group", "enforcement": "provider-schema", "group": "ingress_egress_gw.hub.spoke_vnets.vnet:RequiredObjectAttributes:resource_group,vnet_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:spoke_vnets:vnet", "type": "requires"}, {"anchor": "schema-ingress_egress_gw--hub--spoke_vnets--vnet--vnet_name", "enforcement": "provider-schema", "group": "ingress_egress_gw.hub.spoke_vnets.vnet:RequiredObjectAttributes:resource_group,vnet_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:spoke_vnets:vnet", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_egress_gw", "hub", "spoke_vnets", "vnet"], "schema_version": 1, "sections": [{"aliases": ["f5 orchestrated routing"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:spoke_vnets:vnet:f5_orchestrated_routing", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "hub", "spoke_vnets", "vnet", "f5_orchestrated_routing"], "syntax": "attribute", "type": "object"}, {"aliases": ["manual routing"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:spoke_vnets:vnet:manual_routing", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "hub", "spoke_vnets", "vnet", "manual_routing"], "syntax": "attribute", "type": "object"}, {"aliases": ["resource group"], "anchor": "schema-ingress_egress_gw--hub--spoke_vnets--vnet--resource_group", "description": "Resource group of existing VNet.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:spoke_vnets:vnet", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "hub", "spoke_vnets", "vnet", "resource_group"], "syntax": "attribute", "type": "string"}, {"aliases": ["vnet name"], "anchor": "schema-ingress_egress_gw--hub--spoke_vnets--vnet--vnet_name", "description": "Name of existing VNet.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:spoke_vnets:vnet", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "hub", "spoke_vnets", "vnet", "vnet_name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/ingress_egress_gw/hub/spoke_vnets/vnet/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Resource group and name of existing Azure VNet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.hub.spoke_vnets.vnet

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/)
- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/)
- [ingress_egress_gw.hub](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/hub/)
- [ingress_egress_gw.hub.spoke_vnets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/hub/spoke_vnets/)
- ingress_egress_gw.hub.spoke_vnets.vnet

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Resource group and name of existing Azure VNet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("resource_group",
    "vnet_name"),
  validators.ConflictingObjectAttributes("f5_orchestrated_routing",
    "manual_routing")}
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
  "x-ves-oneof-field-routing_type": "[\"f5_orchestrated_routing\",\"manual_routing\"]"
}
```

Terraform syntax:

```terraform
vnet {
  # Configure direct properties listed below.
}
```

## Direct properties

- [f5_orchestrated_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/hub/spoke_vnets/vnet/f5_orchestrated_routing/): complete subsection reference.

- [manual_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/hub/spoke_vnets/vnet/manual_routing/): complete subsection reference.

<a id="schema-ingress_egress_gw--hub--spoke_vnets--vnet--resource_group"></a>

### resource_group property

Type: `"string"`. Optional.

Existing VNet Resource Group. Resource group of existing VNet.

Upstream description:

Resource group of existing VNet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="schema-ingress_egress_gw--hub--spoke_vnets--vnet--vnet_name"></a>

### vnet_name property

Type: `"string"`. Optional.

Existing VNet Name. Name of existing VNet.

Upstream description:

Name of existing VNet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

## Next pages

- [ingress_egress_gw.hub.spoke_vnets.vnet.f5_orchestrated_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/hub/spoke_vnets/vnet/f5_orchestrated_routing/)
- [ingress_egress_gw.hub.spoke_vnets.vnet.manual_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/hub/spoke_vnets/vnet/manual_routing/)
- [ingress_egress_gw.hub.spoke_vnets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/hub/spoke_vnets/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
