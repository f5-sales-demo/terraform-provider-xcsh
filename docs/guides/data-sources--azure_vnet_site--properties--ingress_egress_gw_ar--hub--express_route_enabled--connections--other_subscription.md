---
page_title: "ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription"
subcategory: "Infrastructure"
description: "ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 2665, "body_sha256": "sha256:837fd10f85dcce9813c876c9e481795d3cac103fafe1b88e62207d7cc7f3802b", "canonical_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:connections:other_subscription", "child_ids": ["xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:connections:other_subscription:authorized_key"], "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:connections:other_subscription", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:connections", "path": "docs/guides/data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--connections--other_subscription.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_egress_gw_ar", "hub", "express_route_enabled", "connections", "other_subscription"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/connections/other_subscription/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md)
- [Property reference](data-sources--azure_vnet_site--reference.md)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar.md)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub.md)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled.md)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--connections.md)
- ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription

<a id="section"></a>

Type: `"single"`. Computed.

Express Route Circuit Config From Other Subscription.

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

## Direct properties

- [authorized_key](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--connections--other_subscription--authorized_key.md): complete subsection reference.

<a id="schema-ingress_egress_gw_ar--hub--express_route_enabled--connections--other_subscription--circuit_id"></a>

### circuit_id property

Type: `"string"`. Computed.

Circuit ID. Circuit ID.

Upstream description:

Circuit ID.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
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
    "ves.io.schema.rules.string.max_len": "512"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512"
  }
}
```

## Next pages

- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--connections--other_subscription--authorized_key.md)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--connections.md)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md)
