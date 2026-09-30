---
page_title: "rule_list.rules.spec.query_params.check_present"
subcategory: "Security"
description: "rule_list.rules.spec.query_params.check_present for xcsh_service_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1289, "body_sha256": "sha256:c257662598861b08595e7d0f9f039773512aa468cbce58d7cb3c23f24a104cb6", "canonical_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:query_params:check_present", "child_ids": [], "collection_id": "xcsh-docs:resources:service_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:query_params:check_present", "parent_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:query_params", "path": "docs/guides/resources--service_policy--properties--rule_list--rules--spec--query_params--check_present.md", "provider_name": "service_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rule_list", "rules", "spec", "query_params", "check_present"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy/properties/rule_list/rules/spec/query_params/check_present/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rule_list.rules.spec.query_params.check_present for xcsh_service_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# rule_list.rules.spec.query_params.check_present

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md)
- [Property reference](resources--service_policy--reference.md)
- [rule_list](resources--service_policy--properties--rule_list.md)
- [rule_list.rules](resources--service_policy--properties--rule_list--rules.md)
- [rule_list.rules.spec](resources--service_policy--properties--rule_list--rules--spec.md)
- [rule_list.rules.spec.query_params](resources--service_policy--properties--rule_list--rules--spec--query_params.md)
- rule_list.rules.spec.query_params.check_present

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

- [rule_list.rules.spec.query_params](resources--service_policy--properties--rule_list--rules--spec--query_params.md)
- [xcsh_service_policy](../resources/service_policy.md)
