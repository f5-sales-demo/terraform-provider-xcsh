---
page_title: "tgw_security.active_east_west_service_policies"
subcategory: ""
description: "Active service policies for the east-west proxy."
xcsh_docs: {"aliases": ["tgw security active east west service policies"], "body_bytes": 1202, "body_sha256": "sha256:149cf8bdcdef33c3bf69049e8235304d9958c2d88f58a9c4659a797d98fa890c", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:aws_tgw_site:properties:tgw_security:active_east_west_service_policies:service_policies"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:tgw_security:active_east_west_service_policies", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:tgw_security", "path": "documentation/resources/aws_tgw_site/properties/tgw_security/active_east_west_service_policies/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-2233103232021301-3211110120311313-1122222320100212-3011331000331302-3311002321132211-2222000130113213-0321012021212213-3312331031310030", "registry_path": "docs/guides/resources--aws_tgw_site--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["tgw_security", "active_east_west_service_policies"], "schema_version": 1, "sections": [{"aliases": ["tgw security active east west service policies service policies"], "anchor": "section", "description": "A list of references to service_policy objects.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:tgw_security:active_east_west_service_policies:service_policies", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-tgw_security--active_east_west_service_policies--service_policies--name", "enforcement": "provider-schema", "group": "tgw_security.active_east_west_service_policies.service_policies:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:tgw_security:active_east_west_service_policies:service_policies", "type": "requires"}], "schema_path": ["tgw_security", "active_east_west_service_policies", "service_policies"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/tgw_security/active_east_west_service_policies/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Active service policies for the east-west proxy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tgw_security.active_east_west_service_policies

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/)
- [tgw_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/tgw_security/)
- tgw_security.active_east_west_service_policies

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Active service policies for the east-west proxy.

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
active_east_west_service_policies {
  # Configure direct properties listed below.
}
```

## Direct properties

- [service_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/tgw_security/active_east_west_service_policies/service_policies/): complete subsection reference.
