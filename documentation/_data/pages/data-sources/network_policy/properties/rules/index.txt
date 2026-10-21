---
page_title: "rules"
subcategory: "Security"
description: "Shape of Rule Choice."
xcsh_docs: {"aliases": ["rules"], "body_bytes": 978, "body_sha256": "sha256:8248e0267ca7f6771e0dccfe29e748285ae081484fab13e0fcf9e936eb1a67a1", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:network_policy:properties:rules:egress_rules", "xcsh-docs:data-sources:network_policy:properties:rules:ingress_rules"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:network_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_policy:properties:rules", "parent_id": "xcsh-docs:data-sources:network_policy:reference", "path": "documentation/data-sources/network_policy/properties/rules/index.md", "product": "distributed-cloud", "provider_name": "network_policy", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-0330200303122200-2303303233120202-0133310221030110-3023123232222121-0221333023110201-3023220333211320-3001222113111213-0310132131330212", "registry_path": "docs/guides/data-sources--network_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rules"], "schema_version": 1, "sections": [{"aliases": ["rules egress rules"], "anchor": "section", "description": "Ordered list of rules applied to connections from policy endpoints.", "document_id": "xcsh-docs:data-sources:network_policy:properties:rules:egress_rules", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["rules", "egress_rules"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules ingress rules"], "anchor": "section", "description": "Ordered list of rules applied to connections to policy endpoints.", "document_id": "xcsh-docs:data-sources:network_policy:properties:rules:ingress_rules", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["rules", "ingress_rules"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_policy/properties/rules/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "Shape of Rule Choice.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["network_policyCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules

Breadcrumbs:

- [xcsh_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy/properties/)
- rules

<a id="section"></a>

Type: `"single"`. Computed.

Rule Choice. Shape of Rule Choice.

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

- [egress_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy/properties/rules/egress_rules/): complete subsection reference.

- [ingress_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy/properties/rules/ingress_rules/): complete subsection reference.
