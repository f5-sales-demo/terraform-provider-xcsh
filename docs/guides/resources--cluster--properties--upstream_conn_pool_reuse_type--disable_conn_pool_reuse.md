---
page_title: "upstream_conn_pool_reuse_type.disable_conn_pool_reuse"
subcategory: ""
description: "upstream_conn_pool_reuse_type.disable_conn_pool_reuse for xcsh_cluster."
xcsh_docs: {"aliases": [], "body_bytes": 1012, "body_sha256": "sha256:c35f0fdff2471f21b3e5807cdc27d98280f6bd7ec9b363d9504f6f138cebc192", "canonical_id": "xcsh-docs:resources:cluster:properties:upstream_conn_pool_reuse_type:disable_conn_pool_reuse", "child_ids": [], "collection_id": "xcsh-docs:resources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:cluster:properties:upstream_conn_pool_reuse_type:disable_conn_pool_reuse", "parent_id": "xcsh-docs:resources:cluster:properties:upstream_conn_pool_reuse_type", "path": "docs/guides/resources--cluster--properties--upstream_conn_pool_reuse_type--disable_conn_pool_reuse.md", "provider_name": "cluster", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["upstream_conn_pool_reuse_type", "disable_conn_pool_reuse"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cluster/properties/upstream_conn_pool_reuse_type/disable_conn_pool_reuse/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "upstream_conn_pool_reuse_type.disable_conn_pool_reuse for xcsh_cluster.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# upstream_conn_pool_reuse_type.disable_conn_pool_reuse

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md)
- [Property reference](resources--cluster--reference.md)
- [upstream_conn_pool_reuse_type](resources--cluster--properties--upstream_conn_pool_reuse_type.md)
- upstream_conn_pool_reuse_type.disable_conn_pool_reuse

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable conn pool reuse.

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

Terraform syntax:

```terraform
disable_conn_pool_reuse = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [upstream_conn_pool_reuse_type](resources--cluster--properties--upstream_conn_pool_reuse_type.md)
- [xcsh_cluster](../resources/cluster.md)
