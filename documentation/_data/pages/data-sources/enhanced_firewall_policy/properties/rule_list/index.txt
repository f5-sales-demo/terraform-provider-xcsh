---
page_title: "rule_list"
subcategory: ""
description: "Custom Enhanced Firewall Policy Rules."
xcsh_docs: {"aliases": ["rule list"], "body_bytes": 891, "body_sha256": "sha256:405ceba6c2f07af2bde45e0ef382ebc2c2a60889798b47c7eb79190bbd02d2df", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:enhanced_firewall_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list", "parent_id": "xcsh-docs:data-sources:enhanced_firewall_policy:reference", "path": "documentation/data-sources/enhanced_firewall_policy/properties/rule_list/index.md", "product": "distributed-cloud", "provider_name": "enhanced_firewall_policy", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-2222221232113201-2332332231203200-0002112022102311-1000133032211001-1011223300133010-3300003012031231-1130121231201200-1013230320331012", "registry_path": "docs/guides/data-sources--enhanced_firewall_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rule_list"], "schema_version": 1, "sections": [{"aliases": ["rule list rules"], "anchor": "section", "description": "Ordered List of Enhanced Firewall Policy Rules.", "document_id": "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["rule_list", "rules"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/enhanced_firewall_policy/properties/rule_list/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "Custom Enhanced Firewall Policy Rules.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["enhanced_firewall_policyCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/)
- rule_list

<a id="section"></a>

Type: `"single"`. Computed.

Custom Enhanced Firewall Policy Rules. Custom Enhanced Firewall Policy Rules.

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

- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/): complete subsection reference.
