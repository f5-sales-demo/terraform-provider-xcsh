---
page_title: "mitigation_type"
subcategory: ""
description: "mitigation_type for xcsh_malicious_user_mitigation."
xcsh_docs: {"aliases": [], "body_bytes": 2237, "body_sha256": "sha256:b601d1897b0c9eb4848d957c9a163e89ef9a4a544092045ce705eed53d8902d8", "child_ids": ["xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules"], "collection_id": "xcsh-docs:resources:malicious_user_mitigation:collection", "completeness": "complete", "id": "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type", "parent_id": "xcsh-docs:resources:malicious_user_mitigation:reference", "path": "documentation/resources/malicious_user_mitigation/properties/mitigation_type/index.md", "provider_name": "malicious_user_mitigation", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["mitigation_type"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/malicious_user_mitigation/properties/mitigation_type/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "mitigation_type for xcsh_malicious_user_mitigation.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["malicious_user_mitigationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# mitigation_type

Breadcrumbs:

- [xcsh_malicious_user_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/malicious_user_mitigation/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/malicious_user_mitigation/properties/)
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

- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/malicious_user_mitigation/properties/mitigation_type/rules/): complete subsection reference.

## Next pages

- [mitigation_type.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/malicious_user_mitigation/properties/mitigation_type/rules/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/malicious_user_mitigation/properties/)
- [xcsh_malicious_user_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/malicious_user_mitigation/)
