---
page_title: "mitigation_type.rules"
subcategory: ""
description: "Define the threat levels and the corresponding mitigation actions to be taken."
xcsh_docs: {"aliases": ["mitigation type rules"], "body_bytes": 2115, "body_sha256": "sha256:79fa5de458b4975f37209593081b548a43012058b20c82f6a8cf673c5f1feea5", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:malicious_user_mitigation:properties:mitigation_type:rules:mitigation_action", "xcsh-docs:data-sources:malicious_user_mitigation:properties:mitigation_type:rules:threat_level"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:malicious_user_mitigation:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:malicious_user_mitigation:properties:mitigation_type:rules", "parent_id": "xcsh-docs:data-sources:malicious_user_mitigation:properties:mitigation_type", "path": "documentation/data-sources/malicious_user_mitigation/properties/mitigation_type/rules/index.md", "product": "distributed-cloud", "provider_name": "malicious_user_mitigation", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "data-sources", "registry_anchor": "canonical-0010202120012102-2133302201320223-0111010332312020-3232011323010010-3112032100230212-0103000312223210-3002211221233301-0302332311123121", "registry_path": "docs/guides/data-sources--malicious_user_mitigation--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["mitigation_type", "rules"], "schema_version": 1, "sections": [{"aliases": ["mitigation type rules mitigation action"], "anchor": "section", "description": "Supported actions that can be taken to mitigate malicious activity from a user.", "document_id": "xcsh-docs:data-sources:malicious_user_mitigation:properties:mitigation_type:rules:mitigation_action", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["mitigation_type", "rules", "mitigation_action"], "syntax": "attribute", "type": "object"}, {"aliases": ["mitigation type rules threat level"], "anchor": "section", "description": "Threat level estimated for each user based on the user's activity and reputation.", "document_id": "xcsh-docs:data-sources:malicious_user_mitigation:properties:mitigation_type:rules:threat_level", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["mitigation_type", "rules", "threat_level"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/malicious_user_mitigation/properties/mitigation_type/rules/index.txt", "spec_pin_digest": "sha256:66dcd7f1752e908b5a58a7638e63662f0daa121544d954167e4fab269122ff4a", "summary": "Define the threat levels and the corresponding mitigation actions to be taken.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.1", "schema_components": ["malicious_user_mitigationCreateRequest"], "target_commit": "bafefa99cbad3ad2e3c2c4fd9c0611a0cbcc36b3"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# mitigation_type.rules

Breadcrumbs:

- [xcsh_malicious_user_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/malicious_user_mitigation/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/malicious_user_mitigation/properties/)
- [mitigation_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/malicious_user_mitigation/properties/mitigation_type/)
- mitigation_type.rules

<a id="section"></a>

Type: `"list"`. Computed.

Define the threat levels and the corresponding mitigation actions to be taken.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 3,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-17T10:10:42+00:00"
    },
    "minItems": 0,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true",
    "ves.io.schema.rules.repeated.unique_threat_level": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true",
    "ves.io.schema.rules.repeated.unique_threat_level": "true"
  }
}
```

## Direct properties

- [mitigation_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/malicious_user_mitigation/properties/mitigation_type/rules/mitigation_action/): complete subsection reference.

- [threat_level](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/malicious_user_mitigation/properties/mitigation_type/rules/threat_level/): complete subsection reference.
