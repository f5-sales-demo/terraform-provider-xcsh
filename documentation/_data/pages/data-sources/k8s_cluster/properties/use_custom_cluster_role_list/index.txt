---
page_title: "use_custom_cluster_role_list"
subcategory: ""
description: "List of active cluster role list for a K8s cluster."
xcsh_docs: {"aliases": ["use custom cluster role list"], "body_bytes": 1924, "body_sha256": "sha256:cefb93e55c0f86ef23875a09734832ca2c612421b4f62d439b8d80913e34671b", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:k8s_cluster:properties:use_custom_cluster_role_list:cluster_roles"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:k8s_cluster:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:k8s_cluster:properties:use_custom_cluster_role_list", "parent_id": "xcsh-docs:data-sources:k8s_cluster:reference", "path": "documentation/data-sources/k8s_cluster/properties/use_custom_cluster_role_list/index.md", "product": "distributed-cloud", "provider_name": "k8s_cluster", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2210312112133200-2120022110103011-3332312223021001-2132020112232033-3130120230202301-1002132302112011-3330122112333332-2312202300223020", "registry_path": "docs/guides/data-sources--k8s_cluster--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["use_custom_cluster_role_list"], "schema_version": 1, "sections": [{"aliases": ["cluster roles"], "anchor": "section", "description": "List of active cluster role list for a K8s cluster.", "document_id": "xcsh-docs:data-sources:k8s_cluster:properties:use_custom_cluster_role_list:cluster_roles", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["use_custom_cluster_role_list", "cluster_roles"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/k8s_cluster/properties/use_custom_cluster_role_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "List of active cluster role list for a K8s cluster.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# use_custom_cluster_role_list

Breadcrumbs:

- [xcsh_k8s_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/properties/)
- use_custom_cluster_role_list

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: use\_custom\_cluster\_role\_list, use\_default\_cluster\_roles; Default:
use\_default\_cluster\_roles\] List of active cluster role list for a K8s cluster.

Upstream description:

List of active cluster role list for a K8s cluster.

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

- [use_custom_cluster_role_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/properties/use_custom_cluster_role_list/#section)
- [use_default_cluster_roles](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/properties/use_default_cluster_roles/#section)

Select alternatives according to the provider validators above.

## Direct properties

- [cluster_roles](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/properties/use_custom_cluster_role_list/cluster_roles/): complete subsection reference.

## Next pages

- [use_custom_cluster_role_list.cluster_roles](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/properties/use_custom_cluster_role_list/cluster_roles/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/properties/)
- [xcsh_k8s_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/)
