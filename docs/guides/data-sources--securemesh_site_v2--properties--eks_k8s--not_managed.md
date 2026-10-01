---
page_title: "eks_k8s.not_managed"
subcategory: ""
description: "eks_k8s.not_managed for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 1318, "body_sha256": "sha256:621902896398bf6bac02cce194956069173520ca6b9ad92d77cb999e4eb0c512", "canonical_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:eks_k8s:not_managed", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:eks_k8s:not_managed:node_list"], "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:eks_k8s:not_managed", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:eks_k8s", "path": "docs/guides/data-sources--securemesh_site_v2--properties--eks_k8s--not_managed.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["eks_k8s", "not_managed"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/eks_k8s/not_managed/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "eks_k8s.not_managed for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# eks_k8s.not_managed

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
- [Property reference](data-sources--securemesh_site_v2--reference.md)
- [eks_k8s](data-sources--securemesh_site_v2--properties--eks_k8s.md)
- eks_k8s.not_managed

<a id="section"></a>

Type: `"single"`. Computed.

Section will show nodes associated with this site.

Upstream description:

This section will show nodes associated with this site. Note: For sites that are not orchestrated by
F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it
will be shown in this section.

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

- [node_list](data-sources--securemesh_site_v2--properties--eks_k8s--not_managed--node_list.md): complete subsection reference.

## Next pages

- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--properties--eks_k8s--not_managed--node_list.md)
- [eks_k8s](data-sources--securemesh_site_v2--properties--eks_k8s.md)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
