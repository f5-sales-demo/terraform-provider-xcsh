---
page_title: "request_constraints.max_header_count_none"
subcategory: ""
description: "request_constraints.max_header_count_none for xcsh_service_policy_rule."
xcsh_docs: {"aliases": [], "body_bytes": 1127, "body_sha256": "sha256:994ab519fbcb491dc1f565babb96f745c1f6e1733743a8af0a11bfc551a2ecdf", "canonical_id": "xcsh-docs:resources:service_policy_rule:properties:request_constraints:max_header_count_none", "child_ids": [], "collection_id": "xcsh-docs:resources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy_rule:properties:request_constraints:max_header_count_none", "parent_id": "xcsh-docs:resources:service_policy_rule:properties:request_constraints", "path": "docs/guides/resources--service_policy_rule--properties--request_constraints--max_header_count_none.md", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["request_constraints", "max_header_count_none"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy_rule/properties/request_constraints/max_header_count_none/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "request_constraints.max_header_count_none for xcsh_service_policy_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# request_constraints.max_header_count_none

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md)
- [Property reference](resources--service_policy_rule--reference.md)
- [request_constraints](resources--service_policy_rule--properties--request_constraints.md)
- request_constraints.max_header_count_none

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max header count none.

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
max_header_count_none = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [request_constraints](resources--service_policy_rule--properties--request_constraints.md)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md)
