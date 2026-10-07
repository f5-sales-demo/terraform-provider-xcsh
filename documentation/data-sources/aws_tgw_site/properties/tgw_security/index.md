---
page_title: "tgw_security"
subcategory: ""
description: "Security Configuration for transit gateway."
xcsh_docs: {"aliases": ["tgw security"], "body_bytes": 2869, "body_sha256": "sha256:02b1f635f0a0d4aa7cd2b61de96dec402c778970e07aa5aeaf1690cbd67e4a96", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:aws_tgw_site:properties:tgw_security:active_east_west_service_policies", "xcsh-docs:data-sources:aws_tgw_site:properties:tgw_security:active_enhanced_firewall_policies", "xcsh-docs:data-sources:aws_tgw_site:properties:tgw_security:active_forward_proxy_policies", "xcsh-docs:data-sources:aws_tgw_site:properties:tgw_security:active_network_policies", "xcsh-docs:data-sources:aws_tgw_site:properties:tgw_security:east_west_service_policy_allow_all", "xcsh-docs:data-sources:aws_tgw_site:properties:tgw_security:forward_proxy_allow_all", "xcsh-docs:data-sources:aws_tgw_site:properties:tgw_security:no_east_west_policy", "xcsh-docs:data-sources:aws_tgw_site:properties:tgw_security:no_forward_proxy", "xcsh-docs:data-sources:aws_tgw_site:properties:tgw_security:no_network_policy"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_tgw_site:properties:tgw_security", "parent_id": "xcsh-docs:data-sources:aws_tgw_site:reference", "path": "documentation/data-sources/aws_tgw_site/properties/tgw_security/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-0232013222110121-3223210311032231-0332221012122031-2132020312222132-2303023010020032-2101212313223322-1203003033031332-1111320102110212", "registry_path": "docs/guides/data-sources--aws_tgw_site--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["tgw_security"], "schema_version": 1, "sections": [{"aliases": ["tgw security active east west service policies"], "anchor": "section", "description": "Active service policies for the east-west proxy.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:tgw_security:active_east_west_service_policies", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["tgw_security", "active_east_west_service_policies"], "syntax": "attribute", "type": "object"}, {"aliases": ["tgw security active enhanced firewall policies"], "anchor": "section", "description": "List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS available under firewall policies with an additional option for service insertion.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:tgw_security:active_enhanced_firewall_policies", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["tgw_security", "active_enhanced_firewall_policies"], "syntax": "attribute", "type": "object"}, {"aliases": ["tgw security active forward proxy policies"], "anchor": "section", "description": "Ordered List of Forward Proxy Policies active.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:tgw_security:active_forward_proxy_policies", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["tgw_security", "active_forward_proxy_policies"], "syntax": "attribute", "type": "object"}, {"aliases": ["tgw security active network policies"], "anchor": "section", "description": "List of firewall policy views.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:tgw_security:active_network_policies", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["tgw_security", "active_network_policies"], "syntax": "attribute", "type": "object"}, {"aliases": ["tgw security east west service policy allow all"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:tgw_security:east_west_service_policy_allow_all", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tgw_security", "east_west_service_policy_allow_all"], "syntax": "attribute", "type": "object"}, {"aliases": ["tgw security forward proxy allow all"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:tgw_security:forward_proxy_allow_all", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tgw_security", "forward_proxy_allow_all"], "syntax": "attribute", "type": "object"}, {"aliases": ["tgw security no east west policy"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:tgw_security:no_east_west_policy", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tgw_security", "no_east_west_policy"], "syntax": "attribute", "type": "object"}, {"aliases": ["tgw security no forward proxy"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:tgw_security:no_forward_proxy", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tgw_security", "no_forward_proxy"], "syntax": "attribute", "type": "object"}, {"aliases": ["tgw security no network policy"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:tgw_security:no_network_policy", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tgw_security", "no_network_policy"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_tgw_site/properties/tgw_security/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Security Configuration for transit gateway.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tgw_security

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/)
- tgw_security

<a id="section"></a>

Type: `"single"`. Computed.

Security Configuration for transit gateway.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-east_west_service_policy_choice": "[\"active_east_west_service_policies\",\"east_west_service_policy_allow_all\",\"no_east_west_policy\"]",
  "x-ves-oneof-field-forward_proxy_choice": "[\"active_forward_proxy_policies\",\"forward_proxy_allow_all\",\"no_forward_proxy\"]",
  "x-ves-oneof-field-network_policy_choice": "[\"active_enhanced_firewall_policies\",\"active_network_policies\",\"no_network_policy\"]"
}
```

## Direct properties

- [active_east_west_service_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/tgw_security/active_east_west_service_policies/): complete subsection reference.

- [active_enhanced_firewall_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/tgw_security/active_enhanced_firewall_policies/): complete subsection reference.

- [active_forward_proxy_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/tgw_security/active_forward_proxy_policies/): complete subsection reference.

- [active_network_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/tgw_security/active_network_policies/): complete subsection reference.

- [east_west_service_policy_allow_all](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/tgw_security/east_west_service_policy_allow_all/): complete subsection reference.

- [forward_proxy_allow_all](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/tgw_security/forward_proxy_allow_all/): complete subsection reference.

- [no_east_west_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/tgw_security/no_east_west_policy/): complete subsection reference.

- [no_forward_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/tgw_security/no_forward_proxy/): complete subsection reference.

- [no_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/tgw_security/no_network_policy/): complete subsection reference.
