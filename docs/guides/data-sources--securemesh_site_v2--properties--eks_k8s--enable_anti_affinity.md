---
page_title: "eks_k8s.enable_anti_affinity"
subcategory: ""
description: "eks_k8s.enable_anti_affinity for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 1313, "body_sha256": "sha256:515118aaa41e310243543f3454a9ec312dcd5b955e35d73f7725d4c1eb3622e4", "canonical_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:eks_k8s:enable_anti_affinity", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:eks_k8s:enable_anti_affinity:rules"], "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:eks_k8s:enable_anti_affinity", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:eks_k8s", "path": "docs/guides/data-sources--securemesh_site_v2--properties--eks_k8s--enable_anti_affinity.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["eks_k8s", "enable_anti_affinity"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/eks_k8s/enable_anti_affinity/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "eks_k8s.enable_anti_affinity for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# eks_k8s.enable_anti_affinity

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
- [Property reference](data-sources--securemesh_site_v2--reference.md)
- [eks_k8s](data-sources--securemesh_site_v2--properties--eks_k8s.md)
- eks_k8s.enable_anti_affinity

<a id="section"></a>

Type: `"single"`. Computed.

Configuration for pod anti-affinity scheduling rules. Define multiple rules to control how different
applications/components are distributed across your Kubernetes cluster.

Upstream description:

Configuration for pod anti-affinity scheduling rules. Define multiple rules to control how different
applications/components are distributed across your Kubernetes cluster.

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

- [rules](data-sources--securemesh_site_v2--properties--eks_k8s--enable_anti_affinity--rules.md): complete subsection reference.

## Next pages

- [eks_k8s.enable_anti_affinity.rules](data-sources--securemesh_site_v2--properties--eks_k8s--enable_anti_affinity--rules.md)
- [eks_k8s](data-sources--securemesh_site_v2--properties--eks_k8s.md)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
