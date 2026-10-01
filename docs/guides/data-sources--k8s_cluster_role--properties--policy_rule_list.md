---
page_title: "policy_rule_list"
subcategory: "Container"
description: "policy_rule_list for xcsh_k8s_cluster_role."
xcsh_docs: {"aliases": [], "body_bytes": 1033, "body_sha256": "sha256:ff8e5dbec93220e3edeb943761460b9f76ed19be5c9012d5aea55ee04f168d65", "canonical_id": "xcsh-docs:data-sources:k8s_cluster_role:properties:policy_rule_list", "child_ids": ["xcsh-docs:data-sources:k8s_cluster_role:properties:policy_rule_list:policy_rule"], "collection_id": "xcsh-docs:data-sources:k8s_cluster_role:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:k8s_cluster_role:properties:policy_rule_list", "parent_id": "xcsh-docs:data-sources:k8s_cluster_role:reference", "path": "docs/guides/data-sources--k8s_cluster_role--properties--policy_rule_list.md", "provider_name": "k8s_cluster_role", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["policy_rule_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/k8s_cluster_role/properties/policy_rule_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "policy_rule_list for xcsh_k8s_cluster_role.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_cluster_roleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# policy_rule_list

Breadcrumbs:

- [xcsh_k8s_cluster_role](../data-sources/k8s_cluster_role.md)
- [Property reference](data-sources--k8s_cluster_role--reference.md)
- policy_rule_list

<a id="section"></a>

Type: `"single"`. Computed.

Policy Rule List. List of rules for role permissions.

Upstream description:

List of rules for role permissions.

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

## Direct properties

- [policy_rule](data-sources--k8s_cluster_role--properties--policy_rule_list--policy_rule.md): complete subsection reference.

## Next pages

- [policy_rule_list.policy_rule](data-sources--k8s_cluster_role--properties--policy_rule_list--policy_rule.md)
- [Property reference](data-sources--k8s_cluster_role--reference.md)
- [xcsh_k8s_cluster_role](../data-sources/k8s_cluster_role.md)
