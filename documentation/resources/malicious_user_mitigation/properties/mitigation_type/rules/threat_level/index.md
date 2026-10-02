---
page_title: "mitigation_type.rules.threat_level"
subcategory: ""
description: "Threat level estimated for each user based on the user's activity and reputation."
xcsh_docs: {"aliases": ["mitigation type rules threat level"], "body_bytes": 2974, "body_sha256": "sha256:a70da855e1fdf902454fd2f620f10670fc238109b99b7689a438b8d3d44fe995", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules:threat_level:high", "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules:threat_level:low", "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules:threat_level:medium"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:malicious_user_mitigation:collection", "completeness": "complete", "id": "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules:threat_level", "parent_id": "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules", "path": "documentation/resources/malicious_user_mitigation/properties/mitigation_type/rules/threat_level/index.md", "product": "distributed-cloud", "provider_name": "malicious_user_mitigation", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3331212021222023-0103201222303231-1032210101232212-3220023023001111-1221021033133130-3032030322013121-0302133310013322-3030121313211021", "registry_path": "docs/guides/resources--malicious_user_mitigation--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "mitigation_type.rules.threat_level:ConflictingObjectAttributes:high,low", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules:threat_level:high", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "mitigation_type.rules.threat_level:ConflictingObjectAttributes:high,medium", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules:threat_level:high", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "mitigation_type.rules.threat_level:ConflictingObjectAttributes:high,low", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules:threat_level:low", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "mitigation_type.rules.threat_level:ConflictingObjectAttributes:low,medium", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules:threat_level:low", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "mitigation_type.rules.threat_level:ConflictingObjectAttributes:high,medium", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules:threat_level:medium", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "mitigation_type.rules.threat_level:ConflictingObjectAttributes:low,medium", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules:threat_level:medium", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["mitigation_type", "rules", "threat_level"], "schema_version": 1, "sections": [{"aliases": ["high"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules:threat_level:high", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["mitigation_type", "rules", "threat_level", "high"], "syntax": "attribute", "type": "object"}, {"aliases": ["low"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules:threat_level:low", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["mitigation_type", "rules", "threat_level", "low"], "syntax": "attribute", "type": "object"}, {"aliases": ["medium"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules:threat_level:medium", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["mitigation_type", "rules", "threat_level", "medium"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/malicious_user_mitigation/properties/mitigation_type/rules/threat_level/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Threat level estimated for each user based on the user's activity and reputation.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["malicious_user_mitigationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# mitigation_type.rules.threat_level

Breadcrumbs:

- [xcsh_malicious_user_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/malicious_user_mitigation/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/malicious_user_mitigation/properties/)
- [mitigation_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/malicious_user_mitigation/properties/mitigation_type/)
- [mitigation_type.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/malicious_user_mitigation/properties/mitigation_type/rules/)
- mitigation_type.rules.threat_level

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Threat level estimated for each user based on the user's activity and reputation.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("high",
    "low"),
  validators.ConflictingObjectAttributes("high",
    "medium"),
  validators.ConflictingObjectAttributes("low",
    "medium")}
```

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

Terraform syntax:

```terraform
threat_level {
  # Configure direct properties listed below.
}
```

## Direct properties

- [high](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/malicious_user_mitigation/properties/mitigation_type/rules/threat_level/high/): complete subsection reference.

- [low](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/malicious_user_mitigation/properties/mitigation_type/rules/threat_level/low/): complete subsection reference.

- [medium](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/malicious_user_mitigation/properties/mitigation_type/rules/threat_level/medium/): complete subsection reference.

## Next pages

- [mitigation_type.rules.threat_level.high](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/malicious_user_mitigation/properties/mitigation_type/rules/threat_level/high/)
- [mitigation_type.rules.threat_level.low](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/malicious_user_mitigation/properties/mitigation_type/rules/threat_level/low/)
- [mitigation_type.rules.threat_level.medium](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/malicious_user_mitigation/properties/mitigation_type/rules/threat_level/medium/)
- [mitigation_type.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/malicious_user_mitigation/properties/mitigation_type/rules/)
- [xcsh_malicious_user_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/malicious_user_mitigation/)
