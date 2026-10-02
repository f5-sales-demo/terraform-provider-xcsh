---
page_title: "mitigation_type.rules.threat_level"
subcategory: ""
description: "Threat level estimated for each user based on the user's activity and reputation."
xcsh_docs: {"aliases": ["mitigation type rules threat level"], "body_bytes": 2974, "body_sha256": "sha256:a70da855e1fdf902454fd2f620f10670fc238109b99b7689a438b8d3d44fe995", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules:threat_level:high", "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules:threat_level:low", "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules:threat_level:medium"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:malicious_user_mitigation:collection", "completeness": "complete", "id": "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules:threat_level", "parent_id": "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules", "path": "documentation/resources/malicious_user_mitigation/properties/mitigation_type/rules/threat_level/index.md", "product": "distributed-cloud", "provider_name": "malicious_user_mitigation", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3331212021222023-0103201222303231-1032210101232212-3220023023001111-1221021033133130-3032030322013121-0302133310013322-3030121313211021", "registry_path": "docs/guides/resources--malicious_user_mitigation--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "mitigation_type.rules.threat_level:ConflictingObjectAttributes:high,low", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules:threat_level:high", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "mitigation_type.rules.threat_level:ConflictingObjectAttributes:high,medium", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules:threat_level:high", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "mitigation_type.rules.threat_level:ConflictingObjectAttributes:high,low", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules:threat_level:low", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "mitigation_type.rules.threat_level:ConflictingObjectAttributes:low,medium", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules:threat_level:low", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "mitigation_type.rules.threat_level:ConflictingObjectAttributes:high,medium", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules:threat_level:medium", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "mitigation_type.rules.threat_level:ConflictingObjectAttributes:low,medium", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules:threat_level:medium", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["mitigation_type", "rules", "threat_level"], "schema_version": 1, "sections": [{"aliases": ["high"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules:threat_level:high", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["mitigation_type", "rules", "threat_level", "high"], "syntax": "attribute", "type": "object"}, {"aliases": ["low"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules:threat_level:low", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["mitigation_type", "rules", "threat_level", "low"], "syntax": "attribute", "type": "object"}, {"aliases": ["medium"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules:threat_level:medium", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["mitigation_type", "rules", "threat_level", "medium"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/malicious_user_mitigation/properties/mitigation_type/rules/threat_level/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Threat level estimated for each user based on the user's activity and reputation.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["malicious_user_mitigationCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
