---
page_title: "policy_based_challenge.rule_list"
subcategory: "Load Balancing"
description: "policy_based_challenge.rule_list for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1627, "body_sha256": "sha256:c9f066936a0de3c4e8e63cac1dadb9b251dae05e0b78d00d263415a0bfcd2c16", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge", "path": "documentation/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/index.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["policy_based_challenge", "rule_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "policy_based_challenge.rule_list for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# policy_based_challenge.rule_list

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [policy_based_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/)
- policy_based_challenge.rule_list

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

List of challenge rules to be used in policy based challenge.

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
rule_list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/): complete subsection reference.

## Next pages

- [policy_based_challenge.rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/)
- [policy_based_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
