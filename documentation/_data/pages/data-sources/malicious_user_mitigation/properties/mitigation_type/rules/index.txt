---
page_title: "mitigation_type.rules"
subcategory: ""
description: "Define the threat levels and the corresponding mitigation actions to be taken."
xcsh_docs: {"aliases": ["mitigation type rules"], "body_bytes": 2781, "body_sha256": "sha256:77ca02c03eadc76e2da8fbf20706ba5b4dae650d2404c0a21b4f576a84290b23", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:malicious_user_mitigation:properties:mitigation_type:rules:mitigation_action", "xcsh-docs:data-sources:malicious_user_mitigation:properties:mitigation_type:rules:threat_level"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:malicious_user_mitigation:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:malicious_user_mitigation:properties:mitigation_type:rules", "parent_id": "xcsh-docs:data-sources:malicious_user_mitigation:properties:mitigation_type", "path": "documentation/data-sources/malicious_user_mitigation/properties/mitigation_type/rules/index.md", "product": "distributed-cloud", "provider_name": "malicious_user_mitigation", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0010202120012102-2133302201320223-0111010332312020-3232011323010010-3112032100230212-0103000312223210-3002211221233301-0302332311123121", "registry_path": "docs/guides/data-sources--malicious_user_mitigation--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["mitigation_type", "rules"], "schema_version": 1, "sections": [{"aliases": ["mitigation action"], "anchor": "section", "description": "Supported actions that can be taken to mitigate malicious activity from a user.", "document_id": "xcsh-docs:data-sources:malicious_user_mitigation:properties:mitigation_type:rules:mitigation_action", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["mitigation_type", "rules", "mitigation_action"], "syntax": "attribute", "type": "object"}, {"aliases": ["threat level"], "anchor": "section", "description": "Threat level estimated for each user based on the user's activity and reputation.", "document_id": "xcsh-docs:data-sources:malicious_user_mitigation:properties:mitigation_type:rules:threat_level", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["mitigation_type", "rules", "threat_level"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/malicious_user_mitigation/properties/mitigation_type/rules/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Define the threat levels and the corresponding mitigation actions to be taken.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["malicious_user_mitigationCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

## Next pages

- [mitigation_type.rules.mitigation_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/malicious_user_mitigation/properties/mitigation_type/rules/mitigation_action/)
- [mitigation_type.rules.threat_level](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/malicious_user_mitigation/properties/mitigation_type/rules/threat_level/)
- [mitigation_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/malicious_user_mitigation/properties/mitigation_type/)
- [xcsh_malicious_user_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/malicious_user_mitigation/)
