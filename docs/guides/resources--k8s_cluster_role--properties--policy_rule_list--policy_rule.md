---
page_title: "policy_rule_list.policy_rule"
subcategory: "Container"
description: "policy_rule_list.policy_rule for xcsh_k8s_cluster_role."
xcsh_docs: {"aliases": [], "body_bytes": 2515, "body_sha256": "sha256:0ef864a5d3ee11abc74bd79d1ae28462ec7940b0a32bf6622f0fb95b7c6e71d9", "canonical_id": "xcsh-docs:resources:k8s_cluster_role:properties:policy_rule_list:policy_rule", "child_ids": ["xcsh-docs:resources:k8s_cluster_role:properties:policy_rule_list:policy_rule:non_resource_url_list", "xcsh-docs:resources:k8s_cluster_role:properties:policy_rule_list:policy_rule:resource_list"], "collection_id": "xcsh-docs:resources:k8s_cluster_role:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_cluster_role:properties:policy_rule_list:policy_rule", "parent_id": "xcsh-docs:resources:k8s_cluster_role:properties:policy_rule_list", "path": "docs/guides/resources--k8s_cluster_role--properties--policy_rule_list--policy_rule.md", "provider_name": "k8s_cluster_role", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["policy_rule_list", "policy_rule"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_cluster_role/properties/policy_rule_list/policy_rule/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "policy_rule_list.policy_rule for xcsh_k8s_cluster_role.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_cluster_roleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# policy_rule_list.policy_rule

Breadcrumbs:

- [xcsh_k8s_cluster_role](../resources/k8s_cluster_role.md)
- [Property reference](resources--k8s_cluster_role--reference.md)
- [policy_rule_list](resources--k8s_cluster_role--properties--policy_rule_list.md)
- policy_rule_list.policy_rule

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Policy Rules. List of rules for role permissions.

Upstream description:

List of rules for role permissions.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("non_resource_url_list",
    "resource_list")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
policy_rule {
  # Configure direct properties listed below.
}
```

## Direct properties

- [non_resource_url_list](resources--k8s_cluster_role--properties--policy_rule_list--policy_rule--non_resource_url_list.md): complete subsection reference.

- [resource_list](resources--k8s_cluster_role--properties--policy_rule_list--policy_rule--resource_list.md): complete subsection reference.

## Next pages

- [policy_rule_list.policy_rule.non_resource_url_list](resources--k8s_cluster_role--properties--policy_rule_list--policy_rule--non_resource_url_list.md)
- [policy_rule_list.policy_rule.resource_list](resources--k8s_cluster_role--properties--policy_rule_list--policy_rule--resource_list.md)
- [policy_rule_list](resources--k8s_cluster_role--properties--policy_rule_list.md)
- [xcsh_k8s_cluster_role](../resources/k8s_cluster_role.md)
