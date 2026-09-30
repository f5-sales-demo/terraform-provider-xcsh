---
page_title: "ingress_gw"
subcategory: "Infrastructure"
description: "ingress_gw for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 2604, "body_sha256": "sha256:ba3af03c8d1653127b13b2f70e655fcca74f93d26ddb94240a3feab86619573c", "canonical_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw", "child_ids": ["xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw:accelerated_networking", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw:az_nodes", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw:performance_enhancement_mode"], "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:reference", "path": "docs/guides/data-sources--azure_vnet_site--properties--ingress_gw.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_gw"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/ingress_gw/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_gw for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# ingress_gw

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md)
- [Property reference](data-sources--azure_vnet_site--reference.md)
- ingress_gw

<a id="section"></a>

Type: `"single"`. Computed.

Single interface Azure ingress site on on Recommended Region.

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

- [accelerated_networking](data-sources--azure_vnet_site--properties--ingress_gw--accelerated_networking.md): complete subsection reference.

- [az_nodes](data-sources--azure_vnet_site--properties--ingress_gw--az_nodes.md): complete subsection reference.

<a id="schema-ingress_gw--azure_certified_hw"></a>

### azure_certified_hw property

Type: `"string"`. Computed.

\[Enum: azure-byol-voltmesh\] Azure Certified Hardware. Name for Azure certified hardware. The only
possible value is \`azure-byol-voltmesh\`.

Upstream description:

Name for Azure certified hardware.

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

- [performance_enhancement_mode](data-sources--azure_vnet_site--properties--ingress_gw--performance_enhancement_mode.md): complete subsection reference.

## Next pages

- [ingress_gw.accelerated_networking](data-sources--azure_vnet_site--properties--ingress_gw--accelerated_networking.md)
- [ingress_gw.az_nodes](data-sources--azure_vnet_site--properties--ingress_gw--az_nodes.md)
- [ingress_gw.performance_enhancement_mode](data-sources--azure_vnet_site--properties--ingress_gw--performance_enhancement_mode.md)
- [Property reference](data-sources--azure_vnet_site--reference.md)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md)
