---
page_title: "use_custom_cluster_role_bindings"
subcategory: ""
description: "List of active cluster role binding list for a K8s cluster."
xcsh_docs: {"aliases": ["use custom cluster role bindings"], "body_bytes": 2038, "body_sha256": "sha256:a1e90f144feb176da62a1550a223ae1fdc27a8e76b3086cbc35c6ad8e58f2f77", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:k8s_cluster:properties:use_custom_cluster_role_bindings:cluster_role_bindings"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:k8s_cluster:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:k8s_cluster:properties:use_custom_cluster_role_bindings", "parent_id": "xcsh-docs:data-sources:k8s_cluster:reference", "path": "documentation/data-sources/k8s_cluster/properties/use_custom_cluster_role_bindings/index.md", "product": "distributed-cloud", "provider_name": "k8s_cluster", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3321211133132301-2331122022223321-2012231113301103-1020230201103020-0302233321331011-2321020130101031-2012122130121322-3131100020033332", "registry_path": "docs/guides/data-sources--k8s_cluster--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["use_custom_cluster_role_bindings"], "schema_version": 1, "sections": [{"aliases": ["cluster role bindings"], "anchor": "section", "description": "List of active cluster role binding list for a K8s cluster.", "document_id": "xcsh-docs:data-sources:k8s_cluster:properties:use_custom_cluster_role_bindings:cluster_role_bindings", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["use_custom_cluster_role_bindings", "cluster_role_bindings"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/k8s_cluster/properties/use_custom_cluster_role_bindings/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of active cluster role binding list for a K8s cluster.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["k8s_clusterCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# use_custom_cluster_role_bindings

Breadcrumbs:

- [xcsh_k8s_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/properties/)
- use_custom_cluster_role_bindings

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: use\_custom\_cluster\_role\_bindings, use\_default\_cluster\_role\_bindings; Default:
use\_default\_cluster\_role\_bindings\] List of active cluster role binding list for a K8s cluster.

Upstream description:

List of active cluster role binding list for a K8s cluster.

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

- [use_custom_cluster_role_bindings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/properties/use_custom_cluster_role_bindings/#section)
- [use_default_cluster_role_bindings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/properties/use_default_cluster_role_bindings/#section)

Select alternatives according to the provider validators above.

## Direct properties

- [cluster_role_bindings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/properties/use_custom_cluster_role_bindings/cluster_role_bindings/): complete subsection reference.

## Next pages

- [use_custom_cluster_role_bindings.cluster_role_bindings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/properties/use_custom_cluster_role_bindings/cluster_role_bindings/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/properties/)
- [xcsh_k8s_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/)
