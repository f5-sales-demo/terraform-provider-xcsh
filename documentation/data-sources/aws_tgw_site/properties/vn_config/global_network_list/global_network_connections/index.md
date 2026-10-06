---
page_title: "vn_config.global_network_list.global_network_connections"
subcategory: ""
description: "Global network connections."
xcsh_docs: {"aliases": ["vn config global network list global network connections"], "body_bytes": 2171, "body_sha256": "sha256:01ec762cc74bdc3cfd3a33354667a10422f6af169b5a3e8befb8170b720cad8f", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:global_network_list:global_network_connections:sli_to_global_dr", "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:global_network_list:global_network_connections:slo_to_global_dr"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:global_network_list:global_network_connections", "parent_id": "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:global_network_list", "path": "documentation/data-sources/aws_tgw_site/properties/vn_config/global_network_list/global_network_connections/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-3232120200123113-1320230001313313-3102221313212003-1210311123333310-0112111101013103-0020001231321012-2110103023031113-2112210130220102", "registry_path": "docs/guides/data-sources--aws_tgw_site--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["vn_config", "global_network_list", "global_network_connections"], "schema_version": 1, "sections": [{"aliases": ["vn config global network list global network connections sli to global dr"], "anchor": "section", "description": "Global network reference for direct connection.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:global_network_list:global_network_connections:sli_to_global_dr", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["vn_config", "global_network_list", "global_network_connections", "sli_to_global_dr"], "syntax": "attribute", "type": "object"}, {"aliases": ["vn config global network list global network connections slo to global dr"], "anchor": "section", "description": "Global network reference for direct connection.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:global_network_list:global_network_connections:slo_to_global_dr", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["vn_config", "global_network_list", "global_network_connections", "slo_to_global_dr"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_tgw_site/properties/vn_config/global_network_list/global_network_connections/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Global network connections.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vn_config.global_network_list.global_network_connections

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/)
- [vn_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/)
- [vn_config.global_network_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/global_network_list/)
- vn_config.global_network_list.global_network_connections

<a id="section"></a>

Type: `"list"`. Computed.

Global Network Connections. Global network connections.

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

- [sli_to_global_dr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/global_network_list/global_network_connections/sli_to_global_dr/): complete subsection reference.

- [slo_to_global_dr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/global_network_list/global_network_connections/slo_to_global_dr/): complete subsection reference.
