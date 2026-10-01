---
page_title: "policy_based_challenge.rule_list.rules.spec.query_params.check_present"
subcategory: "Load Balancing"
description: "policy_based_challenge.rule_list.rules.spec.query_params.check_present for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1784, "body_sha256": "sha256:0f469e982cf2c4c3344bf8338297fe0a807aa85b4b7534adc8705723210aa9df", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:query_params:check_present", "child_ids": [], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:query_params:check_present", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:query_params", "path": "docs/guides/resources--cdn_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--query_params--check_present.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "query_params", "check_present"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/query_params/check_present/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "policy_based_challenge.rule_list.rules.spec.query_params.check_present for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# policy_based_challenge.rule_list.rules.spec.query_params.check_present

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [policy_based_challenge](resources--cdn_loadbalancer--properties--policy_based_challenge.md)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--properties--policy_based_challenge--rule_list.md)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--properties--policy_based_challenge--rule_list--rules.md)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec.md)
- [policy_based_challenge.rule_list.rules.spec.query_params](resources--cdn_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--query_params.md)
- policy_based_challenge.rule_list.rules.spec.query_params.check_present

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

- [policy_based_challenge.rule_list.rules.spec.query_params](resources--cdn_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--query_params.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
