---
page_title: "jwt_claims.check_not_present"
subcategory: ""
description: "jwt_claims.check_not_present for xcsh_service_policy_rule."
xcsh_docs: {"aliases": [], "body_bytes": 1057, "body_sha256": "sha256:108c1e2a34c5d700d0d76c47a6757f85b0e772ea5012851c8347990ce510b8ee", "canonical_id": "xcsh-docs:resources:service_policy_rule:properties:jwt_claims:check_not_present", "child_ids": [], "collection_id": "xcsh-docs:resources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy_rule:properties:jwt_claims:check_not_present", "parent_id": "xcsh-docs:resources:service_policy_rule:properties:jwt_claims", "path": "docs/guides/resources--service_policy_rule--properties--jwt_claims--check_not_present.md", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["jwt_claims", "check_not_present"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy_rule/properties/jwt_claims/check_not_present/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "jwt_claims.check_not_present for xcsh_service_policy_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# jwt_claims.check_not_present

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md)
- [Property reference](resources--service_policy_rule--reference.md)
- [jwt_claims](resources--service_policy_rule--properties--jwt_claims.md)
- jwt_claims.check_not_present

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

Upstream description:

This can be used for messages where no values are needed.

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
check_not_present = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [jwt_claims](resources--service_policy_rule--properties--jwt_claims.md)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md)
