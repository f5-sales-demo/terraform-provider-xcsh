---
page_title: "mitigation_type.rules"
subcategory: ""
description: "mitigation_type.rules for xcsh_malicious_user_mitigation."
xcsh_docs: {"aliases": [], "body_bytes": 2415, "body_sha256": "sha256:bf878a760b57e45fb2aaaf0548d5e713ef2826c0306fad8d9560adb670947be9", "canonical_id": "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules", "child_ids": ["xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules:mitigation_action", "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules:threat_level"], "collection_id": "xcsh-docs:resources:malicious_user_mitigation:collection", "completeness": "complete", "id": "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules", "parent_id": "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type", "path": "docs/guides/resources--malicious_user_mitigation--properties--mitigation_type--rules.md", "provider_name": "malicious_user_mitigation", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["mitigation_type", "rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/malicious_user_mitigation/properties/mitigation_type/rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "mitigation_type.rules for xcsh_malicious_user_mitigation.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["malicious_user_mitigationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# mitigation_type.rules

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md)
- [Property reference](resources--malicious_user_mitigation--reference.md)
- [mitigation_type](resources--malicious_user_mitigation--properties--mitigation_type.md)
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

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

## Direct properties

- [mitigation_action](resources--malicious_user_mitigation--properties--mitigation_type--rules--mitigation_action.md): complete subsection reference.

- [threat_level](resources--malicious_user_mitigation--properties--mitigation_type--rules--threat_level.md): complete subsection reference.

## Next pages

- [mitigation_type.rules.mitigation_action](resources--malicious_user_mitigation--properties--mitigation_type--rules--mitigation_action.md)
- [mitigation_type.rules.threat_level](resources--malicious_user_mitigation--properties--mitigation_type--rules--threat_level.md)
- [mitigation_type](resources--malicious_user_mitigation--properties--mitigation_type.md)
- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md)
