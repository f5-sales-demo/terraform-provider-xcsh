---
page_title: "ingress_egress_gw_ar.hub"
subcategory: "Infrastructure"
description: "ingress_egress_gw_ar.hub for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 2023, "body_sha256": "sha256:c4116eaa4c160bf28cea5aa26ea568697bfe75e64aba32a4709aa724e4f6d0a4", "canonical_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:hub", "child_ids": ["xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_disabled", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:spoke_vnets"], "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:hub", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar", "path": "docs/guides/resources--azure_vnet_site--properties--ingress_egress_gw_ar--hub.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_egress_gw_ar", "hub"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw_ar.hub for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# ingress_egress_gw_ar.hub

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
- [Property reference](resources--azure_vnet_site--reference.md)
- [ingress_egress_gw_ar](resources--azure_vnet_site--properties--ingress_egress_gw_ar.md)
- ingress_egress_gw_ar.hub

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Hub VNet type. Hub VNet type.

Upstream description:

Hub VNet type.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("express_route_disabled",
    "express_route_enabled")}
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
  "x-ves-oneof-field-express_route_choice": "[\"express_route_disabled\",\"express_route_enabled\"]"
}
```

Terraform syntax:

```terraform
hub {
  # Configure direct properties listed below.
}
```

## Direct properties

- [express_route_disabled](resources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_disabled.md): complete subsection reference.

- [express_route_enabled](resources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled.md): complete subsection reference.

- [spoke_vnets](resources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--spoke_vnets.md): complete subsection reference.

## Next pages

- [ingress_egress_gw_ar.hub.express_route_disabled](resources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_disabled.md)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled.md)
- [ingress_egress_gw_ar.hub.spoke_vnets](resources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--spoke_vnets.md)
- [ingress_egress_gw_ar](resources--azure_vnet_site--properties--ingress_egress_gw_ar.md)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
