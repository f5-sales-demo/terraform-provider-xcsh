---
page_title: "vn_config.global_network_list.global_network_connections"
subcategory: ""
description: "Global network connections."
xcsh_docs: {"aliases": ["vn config global network list global network connections"], "body_bytes": 2502, "body_sha256": "sha256:803afde9faf8e56e913b34ea347713cd3e08a28a788cf8d6e9004d406f574638", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:aws_tgw_site:properties:vn_config:global_network_list:global_network_connections:sli_to_global_dr", "xcsh-docs:resources:aws_tgw_site:properties:vn_config:global_network_list:global_network_connections:slo_to_global_dr"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:global_network_list:global_network_connections", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:global_network_list", "path": "documentation/resources/aws_tgw_site/properties/vn_config/global_network_list/global_network_connections/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-3020111000002321-0233001200031331-3203100211320223-1010113310201221-3113233011301211-1321112232003100-2320003022121100-2021100212310232", "registry_path": "docs/guides/resources--aws_tgw_site--reference--group-003.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "vn_config.global_network_list.global_network_connections:ConflictingListObjectAttributes:sli_to_global_dr,slo_to_global_dr", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:global_network_list:global_network_connections:sli_to_global_dr", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config.global_network_list.global_network_connections:ConflictingListObjectAttributes:sli_to_global_dr,slo_to_global_dr", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:global_network_list:global_network_connections:slo_to_global_dr", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["vn_config", "global_network_list", "global_network_connections"], "schema_version": 1, "sections": [{"aliases": ["vn config global network list global network connections sli to global dr"], "anchor": "section", "description": "Global network reference for direct connection.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:global_network_list:global_network_connections:sli_to_global_dr", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["vn_config", "global_network_list", "global_network_connections", "sli_to_global_dr"], "syntax": "block", "type": "object"}, {"aliases": ["vn config global network list global network connections slo to global dr"], "anchor": "section", "description": "Global network reference for direct connection.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:global_network_list:global_network_connections:slo_to_global_dr", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["vn_config", "global_network_list", "global_network_connections", "slo_to_global_dr"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/vn_config/global_network_list/global_network_connections/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Global network connections.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vn_config.global_network_list.global_network_connections

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/)
- [vn_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/)
- [vn_config.global_network_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/global_network_list/)
- vn_config.global_network_list.global_network_connections

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Global Network Connections. Global network connections.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("sli_to_global_dr",
    "slo_to_global_dr")}
```

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

Terraform syntax:

```terraform
global_network_connections {
  # Configure direct properties listed below.
}
```

## Direct properties

- [sli_to_global_dr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/global_network_list/global_network_connections/sli_to_global_dr/): complete subsection reference.

- [slo_to_global_dr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/global_network_list/global_network_connections/slo_to_global_dr/): complete subsection reference.
