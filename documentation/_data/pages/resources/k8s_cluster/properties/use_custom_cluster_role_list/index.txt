---
page_title: "use_custom_cluster_role_list"
subcategory: ""
description: "List of active cluster role list for a K8s cluster."
xcsh_docs: {"aliases": ["use custom cluster role list"], "body_bytes": 2189, "body_sha256": "sha256:e69427e3177de4ed78eff84fde1b82388e2e4538ab27ba1c15f0ba86131e486d", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:k8s_cluster:properties:use_custom_cluster_role_list:cluster_roles"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:k8s_cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_cluster:properties:use_custom_cluster_role_list", "parent_id": "xcsh-docs:resources:k8s_cluster:reference", "path": "documentation/resources/k8s_cluster/properties/use_custom_cluster_role_list/index.md", "product": "distributed-cloud", "provider_name": "k8s_cluster", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2033312333111003-3123320312002221-3221201213321321-0213023322032121-3232230123201131-0222231311213030-3330211220123210-0223222330320322", "registry_path": "docs/guides/resources--k8s_cluster--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "use_custom_cluster_role_list:RequiredObjectAttributes:cluster_roles", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:k8s_cluster:properties:use_custom_cluster_role_list:cluster_roles", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["use_custom_cluster_role_list"], "schema_version": 1, "sections": [{"aliases": ["cluster roles"], "anchor": "section", "description": "List of active cluster role list for a K8s cluster.", "document_id": "xcsh-docs:resources:k8s_cluster:properties:use_custom_cluster_role_list:cluster_roles", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["use_custom_cluster_role_list", "cluster_roles"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_cluster/properties/use_custom_cluster_role_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "List of active cluster role list for a K8s cluster.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# use_custom_cluster_role_list

Breadcrumbs:

- [xcsh_k8s_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/)
- use_custom_cluster_role_list

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: use\_custom\_cluster\_role\_list, use\_default\_cluster\_roles; Default:
use\_default\_cluster\_roles\] List of active cluster role list for a K8s cluster.

Upstream description:

List of active cluster role list for a K8s cluster.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cluster_roles")}
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

OneOf alternatives in this subsection:

- [use_custom_cluster_role_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/use_custom_cluster_role_list/#section)
- [use_default_cluster_roles](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/use_default_cluster_roles/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
use_custom_cluster_role_list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [cluster_roles](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/use_custom_cluster_role_list/cluster_roles/): complete subsection reference.

## Next pages

- [use_custom_cluster_role_list.cluster_roles](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/use_custom_cluster_role_list/cluster_roles/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/)
- [xcsh_k8s_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/)
