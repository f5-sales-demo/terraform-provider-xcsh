---
page_title: "mitigation_type.rules.mitigation_action"
subcategory: ""
description: "Supported actions that can be taken to mitigate malicious activity from a user."
xcsh_docs: {"aliases": ["mitigation type rules mitigation action"], "body_bytes": 2868, "body_sha256": "sha256:cd4365bba0631d612b2b0b6b3d004bd2f3a6bdc4ad77f4f4357262fb67746101", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:malicious_user_mitigation:properties:mitigation_type:rules:mitigation_action:block_temporarily", "xcsh-docs:data-sources:malicious_user_mitigation:properties:mitigation_type:rules:mitigation_action:captcha_challenge", "xcsh-docs:data-sources:malicious_user_mitigation:properties:mitigation_type:rules:mitigation_action:javascript_challenge"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:malicious_user_mitigation:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:malicious_user_mitigation:properties:mitigation_type:rules:mitigation_action", "parent_id": "xcsh-docs:data-sources:malicious_user_mitigation:properties:mitigation_type:rules", "path": "documentation/data-sources/malicious_user_mitigation/properties/mitigation_type/rules/mitigation_action/index.md", "product": "distributed-cloud", "provider_name": "malicious_user_mitigation", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0030030303020100-0311203100121220-3331220233330221-0101300133002333-1121230110221131-1301220103000110-2210230100031301-1200132002320103", "registry_path": "docs/guides/data-sources--malicious_user_mitigation--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["mitigation_type", "rules", "mitigation_action"], "schema_version": 1, "sections": [{"aliases": ["block temporarily"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:malicious_user_mitigation:properties:mitigation_type:rules:mitigation_action:block_temporarily", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["mitigation_type", "rules", "mitigation_action", "block_temporarily"], "syntax": "attribute", "type": "object"}, {"aliases": ["captcha challenge"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:malicious_user_mitigation:properties:mitigation_type:rules:mitigation_action:captcha_challenge", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["mitigation_type", "rules", "mitigation_action", "captcha_challenge"], "syntax": "attribute", "type": "object"}, {"aliases": ["javascript challenge"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:malicious_user_mitigation:properties:mitigation_type:rules:mitigation_action:javascript_challenge", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["mitigation_type", "rules", "mitigation_action", "javascript_challenge"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/malicious_user_mitigation/properties/mitigation_type/rules/mitigation_action/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Supported actions that can be taken to mitigate malicious activity from a user.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["malicious_user_mitigationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

## Next pages

- [mitigation_type.rules.mitigation_action.block_temporarily](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/malicious_user_mitigation/properties/mitigation_type/rules/mitigation_action/block_temporarily/)
- [mitigation_type.rules.mitigation_action.captcha_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/malicious_user_mitigation/properties/mitigation_type/rules/mitigation_action/captcha_challenge/)
- [mitigation_type.rules.mitigation_action.javascript_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/malicious_user_mitigation/properties/mitigation_type/rules/mitigation_action/javascript_challenge/)
- [mitigation_type.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/malicious_user_mitigation/properties/mitigation_type/rules/)
- [xcsh_malicious_user_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/malicious_user_mitigation/)
