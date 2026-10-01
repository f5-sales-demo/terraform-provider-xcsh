---
page_title: "mitigation_type"
subcategory: ""
description: "mitigation_type for xcsh_malicious_user_mitigation."
xcsh_docs: {"aliases": [], "body_bytes": 1929, "body_sha256": "sha256:b87afb5404ab5695de76f42c00fbe05d5770ec61216d6cb35823f6b0b6a8f5e7", "canonical_id": "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type", "child_ids": ["xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules"], "collection_id": "xcsh-docs:resources:malicious_user_mitigation:collection", "completeness": "complete", "id": "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type", "parent_id": "xcsh-docs:resources:malicious_user_mitigation:reference", "path": "docs/guides/resources--malicious_user_mitigation--properties--mitigation_type.md", "provider_name": "malicious_user_mitigation", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["mitigation_type"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/malicious_user_mitigation/properties/mitigation_type/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "mitigation_type for xcsh_malicious_user_mitigation.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["malicious_user_mitigationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# mitigation_type

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md)
- [Property reference](resources--malicious_user_mitigation--reference.md)
- mitigation_type

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Settings that specify the actions to be taken when malicious users are determined to be at different
threat levels. User's activity is monitored and continuously analyzed for malicious behavior. From
this analysis, a threat-level is assigned to each user. Server applies default when omitted.

Upstream description:

Settings that specify the actions to be taken when malicious users are determined to be at different
threat levels. User's activity is monitored and continuously analyzed for malicious behavior. From
this analysis, a threat-level is assigned to each user. The settings defined in malicious user
mitigation specify what mitigation actions to take for user determined to be at different threat
levels.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("rules")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
mitigation_type {
  # Configure direct properties listed below.
}
```

## Direct properties

- [rules](resources--malicious_user_mitigation--properties--mitigation_type--rules.md): complete subsection reference.

## Next pages

- [mitigation_type.rules](resources--malicious_user_mitigation--properties--mitigation_type--rules.md)
- [Property reference](resources--malicious_user_mitigation--reference.md)
- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md)
