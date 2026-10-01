---
page_title: "rule_list.rules.spec.jwt_claims.check_present"
subcategory: "Security"
description: "rule_list.rules.spec.jwt_claims.check_present for xcsh_service_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1376, "body_sha256": "sha256:31f3f2fd313c17f4142d5fec0331aef97603f8250a7aba8eacd515872e71d261", "canonical_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:jwt_claims:check_present", "child_ids": [], "collection_id": "xcsh-docs:resources:service_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:jwt_claims:check_present", "parent_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:jwt_claims", "path": "docs/guides/resources--service_policy--properties--rule_list--rules--spec--jwt_claims--check_present.md", "provider_name": "service_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rule_list", "rules", "spec", "jwt_claims", "check_present"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy/properties/rule_list/rules/spec/jwt_claims/check_present/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rule_list.rules.spec.jwt_claims.check_present for xcsh_service_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules.spec.jwt_claims.check_present

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md)
- [Property reference](resources--service_policy--reference.md)
- [rule_list](resources--service_policy--properties--rule_list.md)
- [rule_list.rules](resources--service_policy--properties--rule_list--rules.md)
- [rule_list.rules.spec](resources--service_policy--properties--rule_list--rules--spec.md)
- [rule_list.rules.spec.jwt_claims](resources--service_policy--properties--rule_list--rules--spec--jwt_claims.md)
- rule_list.rules.spec.jwt_claims.check_present

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

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
check_present = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [rule_list.rules.spec.jwt_claims](resources--service_policy--properties--rule_list--rules--spec--jwt_claims.md)
- [xcsh_service_policy](../resources/service_policy.md)
