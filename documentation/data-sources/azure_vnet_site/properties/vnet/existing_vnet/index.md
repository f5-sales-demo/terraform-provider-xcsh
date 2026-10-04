---
page_title: "vnet.existing_vnet"
subcategory: "Infrastructure"
description: "Resource group and name of existing Azure VNet."
xcsh_docs: {"aliases": ["vnet existing vnet"], "body_bytes": 4014, "body_sha256": "sha256:ab1720acb6c7c8476c800418c2ec6b2024aa120129ea3a2cd0e7e97d46fa8bd6", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:azure_vnet_site:properties:vnet:existing_vnet:f5_orchestrated_routing", "xcsh-docs:data-sources:azure_vnet_site:properties:vnet:existing_vnet:manual_routing"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:vnet:existing_vnet", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:properties:vnet", "path": "documentation/data-sources/azure_vnet_site/properties/vnet/existing_vnet/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0231221313132102-1320122330300311-2130101210302123-0012233023020101-1030232202100103-0311202112233011-2030330103111003-1012312012202322", "registry_path": "docs/guides/data-sources--azure_vnet_site--reference--group-008.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["vnet", "existing_vnet"], "schema_version": 1, "sections": [{"aliases": ["vnet existing vnet f5 orchestrated routing"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:vnet:existing_vnet:f5_orchestrated_routing", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vnet", "existing_vnet", "f5_orchestrated_routing"], "syntax": "attribute", "type": "object"}, {"aliases": ["vnet existing vnet manual routing"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:vnet:existing_vnet:manual_routing", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vnet", "existing_vnet", "manual_routing"], "syntax": "attribute", "type": "object"}, {"aliases": ["vnet existing vnet resource group"], "anchor": "schema-vnet--existing_vnet--resource_group", "description": "Resource group of existing VNet.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:vnet:existing_vnet", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vnet", "existing_vnet", "resource_group"], "syntax": "attribute", "type": "string"}, {"aliases": ["vnet existing vnet vnet name"], "anchor": "schema-vnet--existing_vnet--vnet_name", "description": "Name of existing VNet.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:vnet:existing_vnet", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vnet", "existing_vnet", "vnet_name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/vnet/existing_vnet/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Resource group and name of existing Azure VNet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vnet.existing_vnet

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/)
- [vnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/vnet/)
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

- [f5_orchestrated_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/vnet/existing_vnet/f5_orchestrated_routing/): complete subsection reference.

- [manual_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/vnet/existing_vnet/manual_routing/): complete subsection reference.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [vnet.existing_vnet.f5_orchestrated_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/vnet/existing_vnet/f5_orchestrated_routing/)
- [vnet.existing_vnet.manual_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/vnet/existing_vnet/manual_routing/)
- [vnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/vnet/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
