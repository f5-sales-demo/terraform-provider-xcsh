---
page_title: "use_default_psp"
subcategory: ""
description: "use_default_psp for xcsh_k8s_cluster."
xcsh_docs: {"aliases": [], "body_bytes": 916, "body_sha256": "sha256:17eb3ac528f5b66c2e9645a5402bf0404971f74b873b84b9ecbbc94ad8ad4ee2", "canonical_id": "xcsh-docs:data-sources:k8s_cluster:properties:use_default_psp", "child_ids": [], "collection_id": "xcsh-docs:data-sources:k8s_cluster:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:k8s_cluster:properties:use_default_psp", "parent_id": "xcsh-docs:data-sources:k8s_cluster:reference", "path": "docs/guides/data-sources--k8s_cluster--properties--use_default_psp.md", "provider_name": "k8s_cluster", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["use_default_psp"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/k8s_cluster/properties/use_default_psp/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "use_default_psp for xcsh_k8s_cluster.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# use_default_psp

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md)
- [Property reference](data-sources--k8s_cluster--reference.md)
- use_default_psp

<a id="section"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for use default psp. Defaults to \`map\[\]\`. Server applies default when
omitted.

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

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](data-sources--k8s_cluster--reference.md)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md)
