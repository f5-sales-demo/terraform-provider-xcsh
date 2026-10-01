---
page_title: "mitigation_type.rules.threat_level"
subcategory: ""
description: "mitigation_type.rules.threat_level for xcsh_malicious_user_mitigation."
xcsh_docs: {"aliases": [], "body_bytes": 2380, "body_sha256": "sha256:70501ba6da5e7b1a99b1600a14699b7813adfca48175edd64ac37cbbbaa22630", "canonical_id": "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules:threat_level", "child_ids": ["xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules:threat_level:high", "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules:threat_level:low", "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules:threat_level:medium"], "collection_id": "xcsh-docs:resources:malicious_user_mitigation:collection", "completeness": "complete", "id": "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules:threat_level", "parent_id": "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules", "path": "docs/guides/resources--malicious_user_mitigation--properties--mitigation_type--rules--threat_level.md", "provider_name": "malicious_user_mitigation", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["mitigation_type", "rules", "threat_level"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/malicious_user_mitigation/properties/mitigation_type/rules/threat_level/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "mitigation_type.rules.threat_level for xcsh_malicious_user_mitigation.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["malicious_user_mitigationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# mitigation_type.rules.threat_level

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md)
- [Property reference](resources--malicious_user_mitigation--reference.md)
- [mitigation_type](resources--malicious_user_mitigation--properties--mitigation_type.md)
- [mitigation_type.rules](resources--malicious_user_mitigation--properties--mitigation_type--rules.md)
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

- [high](resources--malicious_user_mitigation--properties--mitigation_type--rules--threat_level--high.md): complete subsection reference.

- [low](resources--malicious_user_mitigation--properties--mitigation_type--rules--threat_level--low.md): complete subsection reference.

- [medium](resources--malicious_user_mitigation--properties--mitigation_type--rules--threat_level--medium.md): complete subsection reference.

## Next pages

- [mitigation_type.rules.threat_level.high](resources--malicious_user_mitigation--properties--mitigation_type--rules--threat_level--high.md)
- [mitigation_type.rules.threat_level.low](resources--malicious_user_mitigation--properties--mitigation_type--rules--threat_level--low.md)
- [mitigation_type.rules.threat_level.medium](resources--malicious_user_mitigation--properties--mitigation_type--rules--threat_level--medium.md)
- [mitigation_type.rules](resources--malicious_user_mitigation--properties--mitigation_type--rules.md)
- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md)
