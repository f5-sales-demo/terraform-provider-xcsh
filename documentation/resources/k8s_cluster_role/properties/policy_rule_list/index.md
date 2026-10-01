---
page_title: "policy_rule_list"
subcategory: "Container"
description: "policy_rule_list for xcsh_k8s_cluster_role."
xcsh_docs: {"aliases": [], "body_bytes": 1598, "body_sha256": "sha256:57103c20a32e38625d91148b0f5b4581314a2d8c251dd459bc53f92e064221b7", "child_ids": ["xcsh-docs:resources:k8s_cluster_role:properties:policy_rule_list:policy_rule"], "collection_id": "xcsh-docs:resources:k8s_cluster_role:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_cluster_role:properties:policy_rule_list", "parent_id": "xcsh-docs:resources:k8s_cluster_role:reference", "path": "documentation/resources/k8s_cluster_role/properties/policy_rule_list/index.md", "provider_name": "k8s_cluster_role", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["policy_rule_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_cluster_role/properties/policy_rule_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "policy_rule_list for xcsh_k8s_cluster_role.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_cluster_roleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# policy_rule_list

Breadcrumbs:

- [xcsh_k8s_cluster_role](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/properties/)
- policy_rule_list

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Policy Rule List. List of rules for role permissions.

Upstream description:

List of rules for role permissions.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("policy_rule")}
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
policy_rule_list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/properties/policy_rule_list/policy_rule/): complete subsection reference.

## Next pages

- [policy_rule_list.policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/properties/policy_rule_list/policy_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/properties/)
- [xcsh_k8s_cluster_role](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/)
