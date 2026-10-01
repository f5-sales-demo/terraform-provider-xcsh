---
page_title: "policy_based_challenge.rule_list.rules.spec.enable_captcha_challenge"
subcategory: "Load Balancing"
description: "policy_based_challenge.rule_list.rules.spec.enable_captcha_challenge for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1618, "body_sha256": "sha256:f94e219eeb745e5593730dd49919e2b45667caabb4c74aa777e8b70518ad6de4", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:enable_captcha_challenge", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:enable_captcha_challenge", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec", "path": "docs/guides/resources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--enable_captcha_challenge.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "enable_captcha_challenge"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/enable_captcha_challenge/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "policy_based_challenge.rule_list.rules.spec.enable_captcha_challenge for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# policy_based_challenge.rule_list.rules.spec.enable_captcha_challenge

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [policy_based_challenge](resources--http_loadbalancer--properties--policy_based_challenge.md)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--properties--policy_based_challenge--rule_list.md)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules.md)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec.md)
- policy_based_challenge.rule_list.rules.spec.enable_captcha_challenge

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable captcha challenge.

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
enable_captcha_challenge = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
