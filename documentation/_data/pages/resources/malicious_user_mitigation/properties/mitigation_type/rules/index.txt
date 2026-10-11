---
page_title: "mitigation_type.rules"
subcategory: ""
description: "Define the threat levels and the corresponding mitigation actions to be taken."
xcsh_docs: {"aliases": ["mitigation type rules"], "body_bytes": 2214, "body_sha256": "sha256:69d58663e49814ef773848dab9a9cb181a99d3f223b2538f96be8f5b67547bc3", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules:mitigation_action", "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules:threat_level"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:malicious_user_mitigation:collection", "completeness": "complete", "id": "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules", "parent_id": "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type", "path": "documentation/resources/malicious_user_mitigation/properties/mitigation_type/rules/index.md", "product": "distributed-cloud", "provider_name": "malicious_user_mitigation", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "resources", "registry_anchor": "canonical-2203232003033103-0203333122231231-1322022310231230-0112032113103131-2230132122210032-0000332021330012-2222332212022102-1333332133331303", "registry_path": "docs/guides/resources--malicious_user_mitigation--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["mitigation_type", "rules"], "schema_version": 1, "sections": [{"aliases": ["mitigation type rules mitigation action"], "anchor": "section", "description": "Supported actions that can be taken to mitigate malicious activity from a user.", "document_id": "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules:mitigation_action", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "mitigation_type.rules.mitigation_action:ConflictingObjectAttributes:block_temporarily,captcha_challenge", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules:mitigation_action:block_temporarily", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "mitigation_type.rules.mitigation_action:ConflictingObjectAttributes:block_temporarily,javascript_challenge", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules:mitigation_action:block_temporarily", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "mitigation_type.rules.mitigation_action:ConflictingObjectAttributes:block_temporarily,captcha_challenge", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules:mitigation_action:captcha_challenge", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "mitigation_type.rules.mitigation_action:ConflictingObjectAttributes:captcha_challenge,javascript_challenge", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules:mitigation_action:captcha_challenge", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "mitigation_type.rules.mitigation_action:ConflictingObjectAttributes:block_temporarily,javascript_challenge", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules:mitigation_action:javascript_challenge", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "mitigation_type.rules.mitigation_action:ConflictingObjectAttributes:captcha_challenge,javascript_challenge", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules:mitigation_action:javascript_challenge", "type": "conflicts"}], "schema_path": ["mitigation_type", "rules", "mitigation_action"], "syntax": "block", "type": "object"}, {"aliases": ["mitigation type rules threat level"], "anchor": "section", "description": "Threat level estimated for each user based on the user's activity and reputation.", "document_id": "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules:threat_level", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "mitigation_type.rules.threat_level:ConflictingObjectAttributes:high,low", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules:threat_level:high", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "mitigation_type.rules.threat_level:ConflictingObjectAttributes:high,medium", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules:threat_level:high", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "mitigation_type.rules.threat_level:ConflictingObjectAttributes:high,low", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules:threat_level:low", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "mitigation_type.rules.threat_level:ConflictingObjectAttributes:low,medium", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules:threat_level:low", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "mitigation_type.rules.threat_level:ConflictingObjectAttributes:high,medium", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules:threat_level:medium", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "mitigation_type.rules.threat_level:ConflictingObjectAttributes:low,medium", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules:threat_level:medium", "type": "conflicts"}], "schema_path": ["mitigation_type", "rules", "threat_level"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/malicious_user_mitigation/properties/mitigation_type/rules/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Define the threat levels and the corresponding mitigation actions to be taken.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["malicious_user_mitigationCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# mitigation_type.rules

Breadcrumbs:

- [xcsh_malicious_user_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/malicious_user_mitigation/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/malicious_user_mitigation/properties/)
- [mitigation_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/malicious_user_mitigation/properties/mitigation_type/)
- mitigation_type.rules

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

## Direct properties

- [mitigation_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/malicious_user_mitigation/properties/mitigation_type/rules/mitigation_action/): complete subsection reference.

- [threat_level](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/malicious_user_mitigation/properties/mitigation_type/rules/threat_level/): complete subsection reference.
