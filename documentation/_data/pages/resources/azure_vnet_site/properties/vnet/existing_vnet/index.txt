---
page_title: "vnet.existing_vnet"
subcategory: "Infrastructure"
description: "vnet.existing_vnet for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 4652, "body_sha256": "sha256:6fcee296e5030e1c6ff49c5340ca29c4b01dcb8726991aa9ba12a94f56f2b4e6", "child_ids": ["xcsh-docs:resources:azure_vnet_site:properties:vnet:existing_vnet:f5_orchestrated_routing", "xcsh-docs:resources:azure_vnet_site:properties:vnet:existing_vnet:manual_routing"], "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:vnet:existing_vnet", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:vnet", "path": "documentation/resources/azure_vnet_site/properties/vnet/existing_vnet/index.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["vnet", "existing_vnet"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/vnet/existing_vnet/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "vnet.existing_vnet for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vnet.existing_vnet

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/)
- [vnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/vnet/)
- vnet.existing_vnet

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
existing_vnet {
  # Configure direct properties listed below.
}
```

## Direct properties

- [f5_orchestrated_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/vnet/existing_vnet/f5_orchestrated_routing/): complete subsection reference.

- [manual_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/vnet/existing_vnet/manual_routing/): complete subsection reference.

<a id="schema-vnet--existing_vnet--resource_group"></a>

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

<a id="schema-vnet--existing_vnet--vnet_name"></a>

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

- [vnet.existing_vnet.f5_orchestrated_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/vnet/existing_vnet/f5_orchestrated_routing/)
- [vnet.existing_vnet.manual_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/vnet/existing_vnet/manual_routing/)
- [vnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/vnet/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
