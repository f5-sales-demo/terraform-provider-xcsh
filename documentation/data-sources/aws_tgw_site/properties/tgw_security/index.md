---
page_title: "tgw_security"
subcategory: ""
description: "Security Configuration for transit gateway."
xcsh_docs: {"aliases": ["tgw security"], "body_bytes": 4706, "body_sha256": "sha256:5d2a82329586af0afe0cc8b5c04cf41ce91cdaa9e92fca9f1a8659235b8495db", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:aws_tgw_site:properties:tgw_security:active_east_west_service_policies", "xcsh-docs:data-sources:aws_tgw_site:properties:tgw_security:active_enhanced_firewall_policies", "xcsh-docs:data-sources:aws_tgw_site:properties:tgw_security:active_forward_proxy_policies", "xcsh-docs:data-sources:aws_tgw_site:properties:tgw_security:active_network_policies", "xcsh-docs:data-sources:aws_tgw_site:properties:tgw_security:east_west_service_policy_allow_all", "xcsh-docs:data-sources:aws_tgw_site:properties:tgw_security:forward_proxy_allow_all", "xcsh-docs:data-sources:aws_tgw_site:properties:tgw_security:no_east_west_policy", "xcsh-docs:data-sources:aws_tgw_site:properties:tgw_security:no_forward_proxy", "xcsh-docs:data-sources:aws_tgw_site:properties:tgw_security:no_network_policy"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_tgw_site:properties:tgw_security", "parent_id": "xcsh-docs:data-sources:aws_tgw_site:reference", "path": "documentation/data-sources/aws_tgw_site/properties/tgw_security/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0232013222110121-3223210311032231-0332221012122031-2132020312222132-2303023010020032-2101212313223322-1203003033031332-1111320102110212", "registry_path": "docs/guides/data-sources--aws_tgw_site--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["tgw_security"], "schema_version": 1, "sections": [{"aliases": ["active east west service policies"], "anchor": "section", "description": "Active service policies for the east-west proxy.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:tgw_security:active_east_west_service_policies", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["tgw_security", "active_east_west_service_policies"], "syntax": "attribute", "type": "object"}, {"aliases": ["active enhanced firewall policies"], "anchor": "section", "description": "List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS available under firewall policies with an additional option for service insertion.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:tgw_security:active_enhanced_firewall_policies", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["tgw_security", "active_enhanced_firewall_policies"], "syntax": "attribute", "type": "object"}, {"aliases": ["active forward proxy policies"], "anchor": "section", "description": "Ordered List of Forward Proxy Policies active.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:tgw_security:active_forward_proxy_policies", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["tgw_security", "active_forward_proxy_policies"], "syntax": "attribute", "type": "object"}, {"aliases": ["active network policies"], "anchor": "section", "description": "List of firewall policy views.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:tgw_security:active_network_policies", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["tgw_security", "active_network_policies"], "syntax": "attribute", "type": "object"}, {"aliases": ["east west service policy allow all"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:tgw_security:east_west_service_policy_allow_all", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tgw_security", "east_west_service_policy_allow_all"], "syntax": "attribute", "type": "object"}, {"aliases": ["forward proxy allow all"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:tgw_security:forward_proxy_allow_all", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tgw_security", "forward_proxy_allow_all"], "syntax": "attribute", "type": "object"}, {"aliases": ["no east west policy"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:tgw_security:no_east_west_policy", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tgw_security", "no_east_west_policy"], "syntax": "attribute", "type": "object"}, {"aliases": ["no forward proxy"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:tgw_security:no_forward_proxy", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tgw_security", "no_forward_proxy"], "syntax": "attribute", "type": "object"}, {"aliases": ["no network policy"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:tgw_security:no_network_policy", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tgw_security", "no_network_policy"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_tgw_site/properties/tgw_security/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Security Configuration for transit gateway.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

## Next pages

- [tgw_security.active_east_west_service_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/tgw_security/active_east_west_service_policies/)
- [tgw_security.active_enhanced_firewall_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/tgw_security/active_enhanced_firewall_policies/)
- [tgw_security.active_forward_proxy_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/tgw_security/active_forward_proxy_policies/)
- [tgw_security.active_network_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/tgw_security/active_network_policies/)
- [tgw_security.east_west_service_policy_allow_all](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/tgw_security/east_west_service_policy_allow_all/)
- [tgw_security.forward_proxy_allow_all](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/tgw_security/forward_proxy_allow_all/)
- [tgw_security.no_east_west_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/tgw_security/no_east_west_policy/)
- [tgw_security.no_forward_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/tgw_security/no_forward_proxy/)
- [tgw_security.no_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/tgw_security/no_network_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/)
- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/)
