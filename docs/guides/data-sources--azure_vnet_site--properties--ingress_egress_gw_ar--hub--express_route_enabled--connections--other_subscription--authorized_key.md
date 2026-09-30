---
page_title: "ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key"
subcategory: "Infrastructure"
description: "ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 2778, "body_sha256": "sha256:9a8020cc2367987e47a609c707b7bb49c26e84494d90cdcc084a6970d05f7a6d", "canonical_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:connections:other_subscription:authorized_key", "child_ids": ["xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:connections:other_subscription:authorized_key:blindfold_secret_info", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:connections:other_subscription:authorized_key:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:connections:other_subscription:authorized_key", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:connections:other_subscription", "path": "docs/guides/data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--connections--other_subscription--authorized_key.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_egress_gw_ar", "hub", "express_route_enabled", "connections", "other_subscription", "authorized_key"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/connections/other_subscription/authorized_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md)
- [Property reference](data-sources--azure_vnet_site--reference.md)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar.md)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub.md)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled.md)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--connections.md)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--connections--other_subscription.md)
- ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key

<a id="section"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

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

## Direct properties

- [blindfold_secret_info](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--connections--other_subscription--authorized_key--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--connections--other_subscription--authorized_key--clear_secret_info.md): complete subsection reference.

## Next pages

- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--connections--other_subscription--authorized_key--blindfold_secret_info.md)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.clear_secret_info](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--connections--other_subscription--authorized_key--clear_secret_info.md)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--connections--other_subscription.md)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md)
