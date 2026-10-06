---
page_title: "site_acl.fast_acl_rules.action"
subcategory: ""
description: "FastAclRuleAction specifies possible action to be applied on traffic, possible action include dropping, forwarding or ratelimiting the traffic."
xcsh_docs: {"aliases": ["site acl fast acl rules action"], "body_bytes": 2089, "body_sha256": "sha256:f3f46f6d02ac3d83fc61348cf172e4aa84e48b716ec4eea1a252ce3dd1b535bc", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:fast_acl:properties:site_acl:fast_acl_rules:action:policer_action", "xcsh-docs:data-sources:fast_acl:properties:site_acl:fast_acl_rules:action:protocol_policer_action"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:fast_acl:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fast_acl:properties:site_acl:fast_acl_rules:action", "parent_id": "xcsh-docs:data-sources:fast_acl:properties:site_acl:fast_acl_rules", "path": "documentation/data-sources/fast_acl/properties/site_acl/fast_acl_rules/action/index.md", "product": "distributed-cloud", "provider_name": "fast_acl", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-0021121121010230-2321333120320031-2221223110112032-3330203231200002-2132020133232032-1030223130311213-3012301213203022-1222210222230332", "registry_path": "docs/guides/data-sources--fast_acl--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["site_acl", "fast_acl_rules", "action"], "schema_version": 1, "sections": [{"aliases": ["site acl fast acl rules action policer action"], "anchor": "section", "description": "Reference to policer object.", "document_id": "xcsh-docs:data-sources:fast_acl:properties:site_acl:fast_acl_rules:action:policer_action", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["site_acl", "fast_acl_rules", "action", "policer_action"], "syntax": "attribute", "type": "object"}, {"aliases": ["site acl fast acl rules action protocol policer action"], "anchor": "section", "description": "Reference to policer object.", "document_id": "xcsh-docs:data-sources:fast_acl:properties:site_acl:fast_acl_rules:action:protocol_policer_action", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["site_acl", "fast_acl_rules", "action", "protocol_policer_action"], "syntax": "attribute", "type": "object"}, {"aliases": ["site acl fast acl rules action simple action"], "anchor": "schema-site_acl--fast_acl_rules--action--simple_action", "description": "FastAclRuleSimpleAction specifies simple action like PASS or DENY Drop the traffic Forward the traffic.", "document_id": "xcsh-docs:data-sources:fast_acl:properties:site_acl:fast_acl_rules:action", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["site_acl", "fast_acl_rules", "action", "simple_action"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fast_acl/properties/site_acl/fast_acl_rules/action/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "FastAclRuleAction specifies possible action to be applied on traffic, possible action include dropping, forwarding or ratelimiting the traffic.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["fast_aclCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# site_acl.fast_acl_rules.action

Breadcrumbs:

- [xcsh_fast_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/)
- [site_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/site_acl/)
- [site_acl.fast_acl_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/site_acl/fast_acl_rules/)
- site_acl.fast_acl_rules.action

<a id="section"></a>

Type: `"single"`. Computed.

FastAclRuleAction specifies possible action to be applied on traffic, possible action include
dropping, forwarding or ratelimiting the traffic.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-action": "[\"policer_action\",\"protocol_policer_action\",\"simple_action\"]"
}
```

## Direct properties

- [policer_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/site_acl/fast_acl_rules/action/policer_action/): complete subsection reference.

- [protocol_policer_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/site_acl/fast_acl_rules/action/protocol_policer_action/): complete subsection reference.

<a id="schema-site_acl--fast_acl_rules--action--simple_action"></a>

### simple_action property

Type: `"string"`. Computed.

\[Enum: DENY|ALLOW\] FastAclRuleSimpleAction specifies simple action like PASS or DENY Drop the
traffic Forward the traffic. Possible values are \`DENY\`, \`ALLOW\`. Defaults to \`DENY\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "DENY",
  "enum": [
    "DENY",
    "ALLOW"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```
