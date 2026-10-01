---
page_title: "policy_based_challenge.rule_list.rules.spec.arg_matchers.check_not_present"
subcategory: "Load Balancing"
description: "policy_based_challenge.rule_list.rules.spec.arg_matchers.check_not_present for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1811, "body_sha256": "sha256:12b4468c206c86e7e8c5e545ee20463e14f28dfdfec85324026024b47a4274a0", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:arg_matchers:check_not_present", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:arg_matchers:check_not_present", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:arg_matchers", "path": "docs/guides/resources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--arg_matchers--check_not_present.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "arg_matchers", "check_not_present"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/arg_matchers/check_not_present/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "policy_based_challenge.rule_list.rules.spec.arg_matchers.check_not_present for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# policy_based_challenge.rule_list.rules.spec.arg_matchers.check_not_present

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [policy_based_challenge](resources--http_loadbalancer--properties--policy_based_challenge.md)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--properties--policy_based_challenge--rule_list.md)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules.md)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec.md)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers](resources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--arg_matchers.md)
- policy_based_challenge.rule_list.rules.spec.arg_matchers.check_not_present

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

- [policy_based_challenge.rule_list.rules.spec.arg_matchers](resources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--arg_matchers.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
