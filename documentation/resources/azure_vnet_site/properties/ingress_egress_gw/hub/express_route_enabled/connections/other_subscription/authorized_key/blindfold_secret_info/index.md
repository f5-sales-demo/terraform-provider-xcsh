---
page_title: "ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info"
subcategory: "Infrastructure"
description: "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management."
xcsh_docs: {"aliases": ["ingress egress gw hub express route enabled connections other subscription authorized key blindfold secret info"], "body_bytes": 6268, "body_sha256": "sha256:4f584b3d50168720c86a498e06eadf3e6e34619155401e64db59844cc97a7e98", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:express_route_enabled:connections:other_subscription:authorized_key:blindfold_secret_info", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:express_route_enabled:connections:other_subscription:authorized_key", "path": "documentation/resources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/connections/other_subscription/authorized_key/blindfold_secret_info/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3322121333022000-3303002311021101-2220330011123100-3121333203123000-1121112212332003-1103232203011123-0212323223231213-3123010200020203", "registry_path": "docs/guides/resources--azure_vnet_site--reference--group-004.md", "relationships": [{"anchor": "schema-ingress_egress_gw--hub--express_route_enabled--connections--other_subscription--authorized_key--blindfold_secret_info--location", "enforcement": "provider-schema", "group": "ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info:RequiredObjectAttributes:location", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:express_route_enabled:connections:other_subscription:authorized_key:blindfold_secret_info", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_egress_gw", "hub", "express_route_enabled", "connections", "other_subscription", "authorized_key", "blindfold_secret_info"], "schema_version": 1, "sections": [{"aliases": ["ingress egress gw hub express route enabled connections other subscription authorized key blindfold secret info decryption provider"], "anchor": "schema-ingress_egress_gw--hub--express_route_enabled--connections--other_subscription--authorized_key--blindfold_secret_info--decryption_provider", "description": "Name of the Secret Management Access object that contains information about the backend Secret Management service.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:express_route_enabled:connections:other_subscription:authorized_key:blindfold_secret_info", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "hub", "express_route_enabled", "connections", "other_subscription", "authorized_key", "blindfold_secret_info", "decryption_provider"], "syntax": "attribute", "type": "string"}, {"aliases": ["ingress egress gw hub express route enabled connections other subscription authorized key blindfold secret info location"], "anchor": "schema-ingress_egress_gw--hub--express_route_enabled--connections--other_subscription--authorized_key--blindfold_secret_info--location", "description": "Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:express_route_enabled:connections:other_subscription:authorized_key:blindfold_secret_info", "flags": ["optional", "sensitive"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "hub", "express_route_enabled", "connections", "other_subscription", "authorized_key", "blindfold_secret_info", "location"], "syntax": "attribute", "type": "string"}, {"aliases": ["ingress egress gw hub express route enabled connections other subscription authorized key blindfold secret info store provider"], "anchor": "schema-ingress_egress_gw--hub--express_route_enabled--connections--other_subscription--authorized_key--blindfold_secret_info--store_provider", "description": "Name of the Secret Management Access object that contains information about the store to GET encrypted bytes This field needs to be provided only if the URL scheme is not string:///.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:express_route_enabled:connections:other_subscription:authorized_key:blindfold_secret_info", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "hub", "express_route_enabled", "connections", "other_subscription", "authorized_key", "blindfold_secret_info", "store_provider"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/connections/other_subscription/authorized_key/blindfold_secret_info/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/)
- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/)
- [ingress_egress_gw.hub](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/hub/)
- [ingress_egress_gw.hub.express_route_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/)
- [ingress_egress_gw.hub.express_route_enabled.connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/connections/)
- [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/connections/other_subscription/)
- [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/connections/other_subscription/authorized_key/)
- ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
```

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
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-ingress_egress_gw--hub--express_route_enabled--connections--other_subscription--authorized_key--blindfold_secret_info--decryption_provider"></a>

### decryption_provider property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-ingress_egress_gw--hub--express_route_enabled--connections--other_subscription--authorized_key--blindfold_secret_info--location"></a>

### location property

Type: `"string"`. Optional, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="schema-ingress_egress_gw--hub--express_route_enabled--connections--other_subscription--authorized_key--blindfold_secret_info--store_provider"></a>

### store_provider property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Next pages

- [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/connections/other_subscription/authorized_key/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
