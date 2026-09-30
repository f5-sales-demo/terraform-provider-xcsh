---
page_title: "ingress_egress_gw_ar.hub.express_route_enabled.connections.metadata"
subcategory: "Infrastructure"
description: "ingress_egress_gw_ar.hub.express_route_enabled.connections.metadata for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 4387, "body_sha256": "sha256:df8e24d9bcae04899272faff9fd57fc443d08b12d9c37a34c08748f2d0691c80", "canonical_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:connections:metadata", "child_ids": [], "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:connections:metadata", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:connections", "path": "docs/guides/resources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--connections--metadata.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_egress_gw_ar", "hub", "express_route_enabled", "connections", "metadata"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/connections/metadata/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw_ar.hub.express_route_enabled.connections.metadata for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# ingress_egress_gw_ar.hub.express_route_enabled.connections.metadata

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
- [Property reference](resources--azure_vnet_site--reference.md)
- [ingress_egress_gw_ar](resources--azure_vnet_site--properties--ingress_egress_gw_ar.md)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--properties--ingress_egress_gw_ar--hub.md)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled.md)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections](resources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--connections.md)
- ingress_egress_gw_ar.hub.express_route_enabled.connections.metadata

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
metadata {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-ingress_egress_gw_ar--hub--express_route_enabled--connections--metadata--description_spec"></a>

### description_spec property

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="schema-ingress_egress_gw_ar--hub--express_route_enabled--connections--metadata--name"></a>

### name property

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

## Next pages

- [ingress_egress_gw_ar.hub.express_route_enabled.connections](resources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--connections.md)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
