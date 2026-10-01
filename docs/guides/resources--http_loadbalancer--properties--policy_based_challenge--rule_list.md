---
page_title: "policy_based_challenge.rule_list"
subcategory: "Load Balancing"
description: "policy_based_challenge.rule_list for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1272, "body_sha256": "sha256:4ef14670a00e3180fdb41d8b0a6b5126f080288bd8cc786d9a859f888d283fb1", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge", "path": "docs/guides/resources--http_loadbalancer--properties--policy_based_challenge--rule_list.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["policy_based_challenge", "rule_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "policy_based_challenge.rule_list for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# policy_based_challenge.rule_list

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [policy_based_challenge](resources--http_loadbalancer--properties--policy_based_challenge.md)
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

- [rules](resources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules.md): complete subsection reference.

## Next pages

- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules.md)
- [policy_based_challenge](resources--http_loadbalancer--properties--policy_based_challenge.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
