---
page_title: "ingress_egress_gw.az_nodes"
subcategory: "Infrastructure"
description: "ingress_egress_gw.az_nodes for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 2831, "body_sha256": "sha256:74657dd2064dfc95ed2e4960d384d00c36336a55cfde90e282b69ad5e5c7dc15", "canonical_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:az_nodes", "child_ids": ["xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:az_nodes:inside_subnet", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:az_nodes:outside_subnet"], "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:az_nodes", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw", "path": "docs/guides/data-sources--azure_vnet_site--properties--ingress_egress_gw--az_nodes.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_egress_gw", "az_nodes"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/ingress_egress_gw/az_nodes/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw.az_nodes for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.az_nodes

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md)
- [Property reference](data-sources--azure_vnet_site--reference.md)
- [ingress_egress_gw](data-sources--azure_vnet_site--properties--ingress_egress_gw.md)
- ingress_egress_gw.az_nodes

<a id="section"></a>

Type: `"list"`. Computed.

Only Single AZ or Three AZ(s) nodes are supported currently.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.num_items": "1,3"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.num_items": "1,3"
  }
}
```

## Direct properties

<a id="schema-ingress_egress_gw--az_nodes--azure_az"></a>

### azure_az property

Type: `"string"`. Computed.

\[Enum: 1|2|3\] Zone depicting a grouping of datacenters within an Azure region. Expecting numeric
input. Possible values are \`1\`, \`2\`, \`3\`.

Upstream description:

A zone depicting a grouping of datacenters within an Azure region. Expecting numeric input.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "1",
    "2",
    "3"
  ],
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
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
    "ves.io.schema.rules.string.in": "[\\\"1\\\",\\\"2\\\",\\\"3\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"1\\\",\\\"2\\\",\\\"3\\\"]"
  }
}
```

- [inside_subnet](data-sources--azure_vnet_site--properties--ingress_egress_gw--az_nodes--inside_subnet.md): complete subsection reference.

- [outside_subnet](data-sources--azure_vnet_site--properties--ingress_egress_gw--az_nodes--outside_subnet.md): complete subsection reference.

## Next pages

- [ingress_egress_gw.az_nodes.inside_subnet](data-sources--azure_vnet_site--properties--ingress_egress_gw--az_nodes--inside_subnet.md)
- [ingress_egress_gw.az_nodes.outside_subnet](data-sources--azure_vnet_site--properties--ingress_egress_gw--az_nodes--outside_subnet.md)
- [ingress_egress_gw](data-sources--azure_vnet_site--properties--ingress_egress_gw.md)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md)
