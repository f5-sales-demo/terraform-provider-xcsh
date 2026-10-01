---
page_title: "ingress_egress_gw_ar.hub.spoke_vnets.vnet"
subcategory: "Infrastructure"
description: "ingress_egress_gw_ar.hub.spoke_vnets.vnet for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 4753, "body_sha256": "sha256:e06ff16a5f9cb7e2758c77b4a2b47657e282f2af2d1492416005c4bab0a469dc", "canonical_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:spoke_vnets:vnet", "child_ids": ["xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:spoke_vnets:vnet:f5_orchestrated_routing", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:spoke_vnets:vnet:manual_routing"], "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:spoke_vnets:vnet", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:spoke_vnets", "path": "docs/guides/resources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--spoke_vnets--vnet.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_egress_gw_ar", "hub", "spoke_vnets", "vnet"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/spoke_vnets/vnet/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw_ar.hub.spoke_vnets.vnet for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw_ar.hub.spoke_vnets.vnet

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
- [Property reference](resources--azure_vnet_site--reference.md)
- [ingress_egress_gw_ar](resources--azure_vnet_site--properties--ingress_egress_gw_ar.md)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--properties--ingress_egress_gw_ar--hub.md)
- [ingress_egress_gw_ar.hub.spoke_vnets](resources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--spoke_vnets.md)
- ingress_egress_gw_ar.hub.spoke_vnets.vnet

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

- [f5_orchestrated_routing](resources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--spoke_vnets--vnet--f5_orchestrated_routing.md): complete subsection reference.

- [manual_routing](resources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--spoke_vnets--vnet--manual_routing.md): complete subsection reference.

<a id="schema-ingress_egress_gw_ar--hub--spoke_vnets--vnet--resource_group"></a>

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

<a id="schema-ingress_egress_gw_ar--hub--spoke_vnets--vnet--vnet_name"></a>

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

- [ingress_egress_gw_ar.hub.spoke_vnets.vnet.f5_orchestrated_routing](resources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--spoke_vnets--vnet--f5_orchestrated_routing.md)
- [ingress_egress_gw_ar.hub.spoke_vnets.vnet.manual_routing](resources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--spoke_vnets--vnet--manual_routing.md)
- [ingress_egress_gw_ar.hub.spoke_vnets](resources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--spoke_vnets.md)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
