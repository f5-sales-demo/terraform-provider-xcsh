---
page_title: "mitigation_type.rules.threat_level"
subcategory: ""
description: "Threat level estimated for each user based on the user's activity and reputation."
xcsh_docs: {"aliases": ["mitigation type rules threat level"], "body_bytes": 2605, "body_sha256": "sha256:4bf0e53690c825c82eeca33f02284f3abdff44da37c93a60ae1ce3f0cc12af62", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:malicious_user_mitigation:properties:mitigation_type:rules:threat_level:high", "xcsh-docs:data-sources:malicious_user_mitigation:properties:mitigation_type:rules:threat_level:low", "xcsh-docs:data-sources:malicious_user_mitigation:properties:mitigation_type:rules:threat_level:medium"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:malicious_user_mitigation:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:malicious_user_mitigation:properties:mitigation_type:rules:threat_level", "parent_id": "xcsh-docs:data-sources:malicious_user_mitigation:properties:mitigation_type:rules", "path": "documentation/data-sources/malicious_user_mitigation/properties/mitigation_type/rules/threat_level/index.md", "product": "distributed-cloud", "provider_name": "malicious_user_mitigation", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-2101121211102013-0231111031232020-2230030313301212-2333132212311123-2022133110130301-3110310310213022-3113323001223321-0022020313302022", "registry_path": "docs/guides/data-sources--malicious_user_mitigation--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["mitigation_type", "rules", "threat_level"], "schema_version": 1, "sections": [{"aliases": ["mitigation type rules threat level high"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:malicious_user_mitigation:properties:mitigation_type:rules:threat_level:high", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["mitigation_type", "rules", "threat_level", "high"], "syntax": "attribute", "type": "object"}, {"aliases": ["mitigation type rules threat level low"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:malicious_user_mitigation:properties:mitigation_type:rules:threat_level:low", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["mitigation_type", "rules", "threat_level", "low"], "syntax": "attribute", "type": "object"}, {"aliases": ["mitigation type rules threat level medium"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:malicious_user_mitigation:properties:mitigation_type:rules:threat_level:medium", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["mitigation_type", "rules", "threat_level", "medium"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/malicious_user_mitigation/properties/mitigation_type/rules/threat_level/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Threat level estimated for each user based on the user's activity and reputation.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["malicious_user_mitigationCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# mitigation_type.rules.threat_level

Breadcrumbs:

- [xcsh_malicious_user_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/malicious_user_mitigation/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/malicious_user_mitigation/properties/)
- [mitigation_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/malicious_user_mitigation/properties/mitigation_type/)
- [mitigation_type.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/malicious_user_mitigation/properties/mitigation_type/rules/)
- mitigation_type.rules.threat_level

<a id="section"></a>

Type: `"single"`. Computed.

Threat level estimated for each user based on the user's activity and reputation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-threat_level": "[\"high\",\"low\",\"medium\"]"
}
```

## Direct properties

- [high](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/malicious_user_mitigation/properties/mitigation_type/rules/threat_level/high/): complete subsection reference.

- [low](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/malicious_user_mitigation/properties/mitigation_type/rules/threat_level/low/): complete subsection reference.

- [medium](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/malicious_user_mitigation/properties/mitigation_type/rules/threat_level/medium/): complete subsection reference.

## Next pages

- [mitigation_type.rules.threat_level.high](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/malicious_user_mitigation/properties/mitigation_type/rules/threat_level/high/)
- [mitigation_type.rules.threat_level.low](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/malicious_user_mitigation/properties/mitigation_type/rules/threat_level/low/)
- [mitigation_type.rules.threat_level.medium](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/malicious_user_mitigation/properties/mitigation_type/rules/threat_level/medium/)
- [mitigation_type.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/malicious_user_mitigation/properties/mitigation_type/rules/)
- [xcsh_malicious_user_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/malicious_user_mitigation/)
