---
page_title: "ingress_egress_gw.hub.express_route_enabled.connections.other_subscription"
subcategory: "Infrastructure"
description: "Express Route Circuit Config From Other Subscription."
xcsh_docs: {"aliases": ["ingress egress gw hub express route enabled connections other subscription"], "body_bytes": 3444, "body_sha256": "sha256:7c33d04abacd8b949ad624f780eb1896a947903e088668f92d8651db4a6750ad", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:express_route_enabled:connections:other_subscription:authorized_key"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:express_route_enabled:connections:other_subscription", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:express_route_enabled:connections", "path": "documentation/resources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/connections/other_subscription/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2301322201002220-0001232112300130-3231132310203032-1113000310011321-0111332203232001-0113013221022201-2201323012212310-3321022033023033", "registry_path": "docs/guides/resources--azure_vnet_site--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_egress_gw", "hub", "express_route_enabled", "connections", "other_subscription"], "schema_version": 1, "sections": [{"aliases": ["ingress egress gw hub express route enabled connections other subscription authorized key"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:express_route_enabled:connections:other_subscription:authorized_key", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:express_route_enabled:connections:other_subscription:authorized_key:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:express_route_enabled:connections:other_subscription:authorized_key:clear_secret_info", "type": "conflicts"}], "schema_path": ["ingress_egress_gw", "hub", "express_route_enabled", "connections", "other_subscription", "authorized_key"], "syntax": "block", "type": "object"}, {"aliases": ["ingress egress gw hub express route enabled connections other subscription circuit id"], "anchor": "schema-ingress_egress_gw--hub--express_route_enabled--connections--other_subscription--circuit_id", "description": "Circuit ID.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:express_route_enabled:connections:other_subscription", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "hub", "express_route_enabled", "connections", "other_subscription", "circuit_id"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/connections/other_subscription/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Express Route Circuit Config From Other Subscription.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.hub.express_route_enabled.connections.other_subscription

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/)
- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/)
- [ingress_egress_gw.hub](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/hub/)
- [ingress_egress_gw.hub.express_route_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/)
- [ingress_egress_gw.hub.express_route_enabled.connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/connections/)
- ingress_egress_gw.hub.express_route_enabled.connections.other_subscription

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
other_subscription {
  # Configure direct properties listed below.
}
```

## Direct properties

- [authorized_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/connections/other_subscription/authorized_key/): complete subsection reference.

<a id="schema-ingress_egress_gw--hub--express_route_enabled--connections--other_subscription--circuit_id"></a>

### circuit_id property

Type: `"string"`. Optional.

Circuit ID. Circuit ID.

Upstream description:

Circuit ID.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(512),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/connections/other_subscription/authorized_key/)
- [ingress_egress_gw.hub.express_route_enabled.connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/connections/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
