---
page_title: "ingress_egress_gw.global_network_list.global_network_connections"
subcategory: "Infrastructure"
description: "ingress_egress_gw.global_network_list.global_network_connections for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 2549, "body_sha256": "sha256:39279acd9489eed4332d8faccf58ffd35f4b4433e8a70cfd71c97a468d950e3f", "canonical_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:global_network_list:global_network_connections", "child_ids": ["xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:global_network_list:global_network_connections:sli_to_global_dr", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:global_network_list:global_network_connections:slo_to_global_dr"], "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:global_network_list:global_network_connections", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:global_network_list", "path": "docs/guides/data-sources--azure_vnet_site--properties--ingress_egress_gw--global_network_list--global_network_connections.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_egress_gw", "global_network_list", "global_network_connections"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/ingress_egress_gw/global_network_list/global_network_connections/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw.global_network_list.global_network_connections for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# ingress_egress_gw.global_network_list.global_network_connections

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md)
- [Property reference](data-sources--azure_vnet_site--reference.md)
- [ingress_egress_gw](data-sources--azure_vnet_site--properties--ingress_egress_gw.md)
- [ingress_egress_gw.global_network_list](data-sources--azure_vnet_site--properties--ingress_egress_gw--global_network_list.md)
- ingress_egress_gw.global_network_list.global_network_connections

<a id="section"></a>

Type: `"list"`. Computed.

Global Network Connections. Global network connections.

Upstream description:

Global network connections.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

## Direct properties

- [sli_to_global_dr](data-sources--azure_vnet_site--properties--ingress_egress_gw--global_network_list--global_network_connections--sli_to_global_dr.md): complete subsection reference.

- [slo_to_global_dr](data-sources--azure_vnet_site--properties--ingress_egress_gw--global_network_list--global_network_connections--slo_to_global_dr.md): complete subsection reference.

## Next pages

- [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr](data-sources--azure_vnet_site--properties--ingress_egress_gw--global_network_list--global_network_connections--sli_to_global_dr.md)
- [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr](data-sources--azure_vnet_site--properties--ingress_egress_gw--global_network_list--global_network_connections--slo_to_global_dr.md)
- [ingress_egress_gw.global_network_list](data-sources--azure_vnet_site--properties--ingress_egress_gw--global_network_list.md)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md)
