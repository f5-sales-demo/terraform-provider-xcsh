---
page_title: "mitigation_type.rules.mitigation_action"
subcategory: ""
description: "Supported actions that can be taken to mitigate malicious activity from a user."
xcsh_docs: {"aliases": ["mitigation type rules mitigation action"], "body_bytes": 1874, "body_sha256": "sha256:356efc1db20d948474180fb7999f3dd10854b169b7524103bdbf513296847134", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:malicious_user_mitigation:properties:mitigation_type:rules:mitigation_action:block_temporarily", "xcsh-docs:data-sources:malicious_user_mitigation:properties:mitigation_type:rules:mitigation_action:captcha_challenge", "xcsh-docs:data-sources:malicious_user_mitigation:properties:mitigation_type:rules:mitigation_action:javascript_challenge"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:malicious_user_mitigation:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:malicious_user_mitigation:properties:mitigation_type:rules:mitigation_action", "parent_id": "xcsh-docs:data-sources:malicious_user_mitigation:properties:mitigation_type:rules", "path": "documentation/data-sources/malicious_user_mitigation/properties/mitigation_type/rules/mitigation_action/index.md", "product": "distributed-cloud", "provider_name": "malicious_user_mitigation", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-0030030303020100-0311203100121220-3331220233330221-0101300133002333-1121230110221131-1301220103000110-2210230100031301-1200132002320103", "registry_path": "docs/guides/data-sources--malicious_user_mitigation--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["mitigation_type", "rules", "mitigation_action"], "schema_version": 1, "sections": [{"aliases": ["mitigation type rules mitigation action block temporarily"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:malicious_user_mitigation:properties:mitigation_type:rules:mitigation_action:block_temporarily", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["mitigation_type", "rules", "mitigation_action", "block_temporarily"], "syntax": "attribute", "type": "object"}, {"aliases": ["mitigation type rules mitigation action captcha challenge"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:malicious_user_mitigation:properties:mitigation_type:rules:mitigation_action:captcha_challenge", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["mitigation_type", "rules", "mitigation_action", "captcha_challenge"], "syntax": "attribute", "type": "object"}, {"aliases": ["mitigation type rules mitigation action javascript challenge"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:malicious_user_mitigation:properties:mitigation_type:rules:mitigation_action:javascript_challenge", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["mitigation_type", "rules", "mitigation_action", "javascript_challenge"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/malicious_user_mitigation/properties/mitigation_type/rules/mitigation_action/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Supported actions that can be taken to mitigate malicious activity from a user.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["malicious_user_mitigationCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# mitigation_type.rules.mitigation_action

Breadcrumbs:

- [xcsh_malicious_user_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/malicious_user_mitigation/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/malicious_user_mitigation/properties/)
- [mitigation_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/malicious_user_mitigation/properties/mitigation_type/)
- [mitigation_type.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/malicious_user_mitigation/properties/mitigation_type/rules/)
- mitigation_type.rules.mitigation_action

<a id="section"></a>

Type: `"single"`. Computed.

Supported actions that can be taken to mitigate malicious activity from a user.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-mitigation_action": "[\"block_temporarily\",\"captcha_challenge\",\"javascript_challenge\"]"
}
```

## Direct properties

- [block_temporarily](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/malicious_user_mitigation/properties/mitigation_type/rules/mitigation_action/block_temporarily/): complete subsection reference.

- [captcha_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/malicious_user_mitigation/properties/mitigation_type/rules/mitigation_action/captcha_challenge/): complete subsection reference.

- [javascript_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/malicious_user_mitigation/properties/mitigation_type/rules/mitigation_action/javascript_challenge/): complete subsection reference.
