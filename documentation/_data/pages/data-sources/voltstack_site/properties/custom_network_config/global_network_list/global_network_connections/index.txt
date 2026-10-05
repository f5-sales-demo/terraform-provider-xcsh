---
page_title: "custom_network_config.global_network_list.global_network_connections"
subcategory: ""
description: "Global network connections."
xcsh_docs: {"aliases": ["custom network config global network list global network connections"], "body_bytes": 3190, "body_sha256": "sha256:304f05f66738b30b83140018e5b4b87301ccaa01edb86fbba344f070a63cb78a", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:global_network_list:global_network_connections:sli_to_global_dr", "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:global_network_list:global_network_connections:slo_to_global_dr"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:global_network_list:global_network_connections", "parent_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:global_network_list", "path": "documentation/data-sources/voltstack_site/properties/custom_network_config/global_network_list/global_network_connections/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-0002012313111030-3102003230303312-0031131103322323-3100313310133013-3001110203120122-2101231322103233-0012331010121233-2320002223103232", "registry_path": "docs/guides/data-sources--voltstack_site--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_network_config", "global_network_list", "global_network_connections"], "schema_version": 1, "sections": [{"aliases": ["custom network config global network list global network connections sli to global dr"], "anchor": "section", "description": "Global network reference for direct connection.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:global_network_list:global_network_connections:sli_to_global_dr", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["custom_network_config", "global_network_list", "global_network_connections", "sli_to_global_dr"], "syntax": "attribute", "type": "object"}, {"aliases": ["custom network config global network list global network connections slo to global dr"], "anchor": "section", "description": "Global network reference for direct connection.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:global_network_list:global_network_connections:slo_to_global_dr", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["custom_network_config", "global_network_list", "global_network_connections", "slo_to_global_dr"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/custom_network_config/global_network_list/global_network_connections/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Global network connections.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_network_config.global_network_list.global_network_connections

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/)
- [custom_network_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/)
- [custom_network_config.global_network_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/global_network_list/)
- custom_network_config.global_network_list.global_network_connections

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [sli_to_global_dr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/global_network_list/global_network_connections/sli_to_global_dr/): complete subsection reference.

- [slo_to_global_dr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/global_network_list/global_network_connections/slo_to_global_dr/): complete subsection reference.

## Next pages

- [custom_network_config.global_network_list.global_network_connections.sli_to_global_dr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/global_network_list/global_network_connections/sli_to_global_dr/)
- [custom_network_config.global_network_list.global_network_connections.slo_to_global_dr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/global_network_list/global_network_connections/slo_to_global_dr/)
- [custom_network_config.global_network_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/global_network_list/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
