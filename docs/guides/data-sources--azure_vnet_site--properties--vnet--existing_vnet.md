---
page_title: "vnet.existing_vnet"
subcategory: "Infrastructure"
description: "vnet.existing_vnet for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 3561, "body_sha256": "sha256:58febe139c7c8a50fadc0dc9ebb772e91149e5fd34962a6c7c19e9ac60d6c533", "canonical_id": "xcsh-docs:data-sources:azure_vnet_site:properties:vnet:existing_vnet", "child_ids": ["xcsh-docs:data-sources:azure_vnet_site:properties:vnet:existing_vnet:f5_orchestrated_routing", "xcsh-docs:data-sources:azure_vnet_site:properties:vnet:existing_vnet:manual_routing"], "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:vnet:existing_vnet", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:properties:vnet", "path": "docs/guides/data-sources--azure_vnet_site--properties--vnet--existing_vnet.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["vnet", "existing_vnet"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/vnet/existing_vnet/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "vnet.existing_vnet for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vnet.existing_vnet

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md)
- [Property reference](data-sources--azure_vnet_site--reference.md)
- [vnet](data-sources--azure_vnet_site--properties--vnet.md)
- vnet.existing_vnet

<a id="section"></a>

Type: `"single"`. Computed.

Resource group and name of existing Azure VNet.

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

## Direct properties

- [f5_orchestrated_routing](data-sources--azure_vnet_site--properties--vnet--existing_vnet--f5_orchestrated_routing.md): complete subsection reference.

- [manual_routing](data-sources--azure_vnet_site--properties--vnet--existing_vnet--manual_routing.md): complete subsection reference.

<a id="schema-vnet--existing_vnet--resource_group"></a>

### resource_group property

Type: `"string"`. Computed.

Existing VNet Resource Group. Resource group of existing VNet.

Upstream description:

Resource group of existing VNet.

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

Type: `"string"`. Computed.

Existing VNet Name. Name of existing VNet.

Upstream description:

Name of existing VNet.

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

- [vnet.existing_vnet.f5_orchestrated_routing](data-sources--azure_vnet_site--properties--vnet--existing_vnet--f5_orchestrated_routing.md)
- [vnet.existing_vnet.manual_routing](data-sources--azure_vnet_site--properties--vnet--existing_vnet--manual_routing.md)
- [vnet](data-sources--azure_vnet_site--properties--vnet.md)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md)
