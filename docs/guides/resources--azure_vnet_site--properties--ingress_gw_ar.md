---
page_title: "ingress_gw_ar"
subcategory: "Infrastructure"
description: "ingress_gw_ar for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 3199, "body_sha256": "sha256:185d382284d7249dca87a2e1496af907446a54ede08d602a217d99768f21df7b", "canonical_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw_ar", "child_ids": ["xcsh-docs:resources:azure_vnet_site:properties:ingress_gw_ar:accelerated_networking", "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw_ar:node", "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw_ar:performance_enhancement_mode"], "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw_ar", "parent_id": "xcsh-docs:resources:azure_vnet_site:reference", "path": "docs/guides/resources--azure_vnet_site--properties--ingress_gw_ar.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_gw_ar"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/ingress_gw_ar/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_gw_ar for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_gw_ar

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
- [Property reference](resources--azure_vnet_site--reference.md)
- ingress_gw_ar

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for ingress gw ar.

Upstream description:

Single interface Azure ingress site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("azure_certified_hw")}
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
ingress_gw_ar {
  # Configure direct properties listed below.
}
```

## Direct properties

- [accelerated_networking](resources--azure_vnet_site--properties--ingress_gw_ar--accelerated_networking.md): complete subsection reference.

<a id="schema-ingress_gw_ar--azure_certified_hw"></a>

### azure_certified_hw property

Type: `"string"`. Optional.

\[Enum: azure-byol-voltmesh\] Azure Certified Hardware. Name for Azure certified hardware. The only
possible value is \`azure-byol-voltmesh\`.

Upstream description:

Name for Azure certified hardware.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
  stringvalidator.OneOf("azure-byol-voltmesh"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "azure-byol-voltmesh"
  ],
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.in": "[\\\"azure-byol-voltmesh\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"azure-byol-voltmesh\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [node](resources--azure_vnet_site--properties--ingress_gw_ar--node.md): complete subsection reference.

- [performance_enhancement_mode](resources--azure_vnet_site--properties--ingress_gw_ar--performance_enhancement_mode.md): complete subsection reference.

## Next pages

- [ingress_gw_ar.accelerated_networking](resources--azure_vnet_site--properties--ingress_gw_ar--accelerated_networking.md)
- [ingress_gw_ar.node](resources--azure_vnet_site--properties--ingress_gw_ar--node.md)
- [ingress_gw_ar.performance_enhancement_mode](resources--azure_vnet_site--properties--ingress_gw_ar--performance_enhancement_mode.md)
- [Property reference](resources--azure_vnet_site--reference.md)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
