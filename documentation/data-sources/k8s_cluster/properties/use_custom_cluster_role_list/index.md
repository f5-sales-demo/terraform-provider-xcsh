---
page_title: "use_custom_cluster_role_list"
subcategory: ""
description: "List of active cluster role list for a K8s cluster."
xcsh_docs: {"aliases": ["use custom cluster role list"], "body_bytes": 1427, "body_sha256": "sha256:f79da430fe0c3bfcc0afc2aedf8180729addb935c5f6677164861bee7780bc59", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:k8s_cluster:properties:use_custom_cluster_role_list:cluster_roles"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:k8s_cluster:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:k8s_cluster:properties:use_custom_cluster_role_list", "parent_id": "xcsh-docs:data-sources:k8s_cluster:reference", "path": "documentation/data-sources/k8s_cluster/properties/use_custom_cluster_role_list/index.md", "product": "distributed-cloud", "provider_name": "k8s_cluster", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-2210312112133200-2120022110103011-3332312223021001-2132020112232033-3130120230202301-1002132302112011-3330122112333332-2312202300223020", "registry_path": "docs/guides/data-sources--k8s_cluster--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["use_custom_cluster_role_list"], "schema_version": 1, "sections": [{"aliases": ["use custom cluster role list cluster roles"], "anchor": "section", "description": "List of active cluster role list for a K8s cluster.", "document_id": "xcsh-docs:data-sources:k8s_cluster:properties:use_custom_cluster_role_list:cluster_roles", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["use_custom_cluster_role_list", "cluster_roles"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/k8s_cluster/properties/use_custom_cluster_role_list/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "List of active cluster role list for a K8s cluster.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["k8s_clusterCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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
