---
page_title: "ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key"
subcategory: "Infrastructure"
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["ingress egress gw ar hub express route enabled connections other subscription authorized key"], "body_bytes": 3776, "body_sha256": "sha256:117f737149ff4f0b2f0c20bc6c1b0a3f783519a456f9dfbc11b1365e23e3b0ac", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:connections:other_subscription:authorized_key:blindfold_secret_info", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:connections:other_subscription:authorized_key:clear_secret_info"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:connections:other_subscription:authorized_key", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:connections:other_subscription", "path": "documentation/resources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/connections/other_subscription/authorized_key/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3232003323122302-1230000103301323-3301112210011332-2310131213033333-3311200031221221-3202031103221200-3013310203220232-1213132032323203", "registry_path": "docs/guides/resources--azure_vnet_site--reference--group-006.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:connections:other_subscription:authorized_key:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:connections:other_subscription:authorized_key:clear_secret_info", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_egress_gw_ar", "hub", "express_route_enabled", "connections", "other_subscription", "authorized_key"], "schema_version": 1, "sections": [{"aliases": ["blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:connections:other_subscription:authorized_key:blindfold_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-ingress_egress_gw_ar--hub--express_route_enabled--connections--other_subscription--authorized_key--blindfold_secret_info--location", "enforcement": "provider-schema", "group": "ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info:RequiredObjectAttributes:location", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:connections:other_subscription:authorized_key:blindfold_secret_info", "type": "requires"}], "schema_path": ["ingress_egress_gw_ar", "hub", "express_route_enabled", "connections", "other_subscription", "authorized_key", "blindfold_secret_info"], "syntax": "block", "type": "object"}, {"aliases": ["clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:connections:other_subscription:authorized_key:clear_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-ingress_egress_gw_ar--hub--express_route_enabled--connections--other_subscription--authorized_key--clear_secret_info--url", "enforcement": "provider-schema", "group": "ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.clear_secret_info:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:connections:other_subscription:authorized_key:clear_secret_info", "type": "requires"}], "schema_path": ["ingress_egress_gw_ar", "hub", "express_route_enabled", "connections", "other_subscription", "authorized_key", "clear_secret_info"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/connections/other_subscription/authorized_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/)
- [ingress_egress_gw_ar](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/)
- [ingress_egress_gw_ar.hub](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/)
- [ingress_egress_gw_ar.hub.express_route_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/connections/)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/connections/other_subscription/)
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

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/connections/other_subscription/authorized_key/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/connections/other_subscription/authorized_key/clear_secret_info/): complete subsection reference.

## Next pages

- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/connections/other_subscription/authorized_key/blindfold_secret_info/)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/connections/other_subscription/authorized_key/clear_secret_info/)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/connections/other_subscription/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
