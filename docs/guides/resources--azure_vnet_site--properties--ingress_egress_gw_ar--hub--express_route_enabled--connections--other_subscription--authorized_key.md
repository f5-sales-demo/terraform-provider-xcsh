---
page_title: "ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key"
subcategory: "Infrastructure"
description: "ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 3149, "body_sha256": "sha256:7d95f0aaa3d273df0dc70a02852933df02cf69f8c18214a426baa076c4f93249", "canonical_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:connections:other_subscription:authorized_key", "child_ids": ["xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:connections:other_subscription:authorized_key:blindfold_secret_info", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:connections:other_subscription:authorized_key:clear_secret_info"], "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:connections:other_subscription:authorized_key", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:connections:other_subscription", "path": "docs/guides/resources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--connections--other_subscription--authorized_key.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_egress_gw_ar", "hub", "express_route_enabled", "connections", "other_subscription", "authorized_key"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/connections/other_subscription/authorized_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
- [Property reference](resources--azure_vnet_site--reference.md)
- [ingress_egress_gw_ar](resources--azure_vnet_site--properties--ingress_egress_gw_ar.md)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--properties--ingress_egress_gw_ar--hub.md)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled.md)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections](resources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--connections.md)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription](resources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--connections--other_subscription.md)
- ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
authorized_key {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](resources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--connections--other_subscription--authorized_key--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](resources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--connections--other_subscription--authorized_key--clear_secret_info.md): complete subsection reference.

## Next pages

- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info](resources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--connections--other_subscription--authorized_key--blindfold_secret_info.md)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.clear_secret_info](resources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--connections--other_subscription--authorized_key--clear_secret_info.md)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription](resources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--connections--other_subscription.md)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
