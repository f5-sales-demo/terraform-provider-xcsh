---
page_title: "mitigation_type.rules"
subcategory: ""
description: "mitigation_type.rules for xcsh_malicious_user_mitigation."
xcsh_docs: {"aliases": [], "body_bytes": 2229, "body_sha256": "sha256:ca10967affecf56d574ff88b9776b37f4170f2fcdacbec2c1aebd405393b4a88", "canonical_id": "xcsh-docs:data-sources:malicious_user_mitigation:properties:mitigation_type:rules", "child_ids": ["xcsh-docs:data-sources:malicious_user_mitigation:properties:mitigation_type:rules:mitigation_action", "xcsh-docs:data-sources:malicious_user_mitigation:properties:mitigation_type:rules:threat_level"], "collection_id": "xcsh-docs:data-sources:malicious_user_mitigation:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:malicious_user_mitigation:properties:mitigation_type:rules", "parent_id": "xcsh-docs:data-sources:malicious_user_mitigation:properties:mitigation_type", "path": "docs/guides/data-sources--malicious_user_mitigation--properties--mitigation_type--rules.md", "provider_name": "malicious_user_mitigation", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["mitigation_type", "rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/malicious_user_mitigation/properties/mitigation_type/rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "mitigation_type.rules for xcsh_malicious_user_mitigation.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["malicious_user_mitigationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# mitigation_type.rules

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../data-sources/malicious_user_mitigation.md)
- [Property reference](data-sources--malicious_user_mitigation--reference.md)
- [mitigation_type](data-sources--malicious_user_mitigation--properties--mitigation_type.md)
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

- [mitigation_action](data-sources--malicious_user_mitigation--properties--mitigation_type--rules--mitigation_action.md): complete subsection reference.

- [threat_level](data-sources--malicious_user_mitigation--properties--mitigation_type--rules--threat_level.md): complete subsection reference.

## Next pages

- [mitigation_type.rules.mitigation_action](data-sources--malicious_user_mitigation--properties--mitigation_type--rules--mitigation_action.md)
- [mitigation_type.rules.threat_level](data-sources--malicious_user_mitigation--properties--mitigation_type--rules--threat_level.md)
- [mitigation_type](data-sources--malicious_user_mitigation--properties--mitigation_type.md)
- [xcsh_malicious_user_mitigation](../data-sources/malicious_user_mitigation.md)
