---
page_title: "upstream_conn_pool_reuse_type"
subcategory: ""
description: "upstream_conn_pool_reuse_type for xcsh_cluster."
xcsh_docs: {"aliases": [], "body_bytes": 1824, "body_sha256": "sha256:80584968190435c5e17a02dd8caa7f7c9cc44ce84ccad08ddfd12a95f04a398b", "canonical_id": "xcsh-docs:resources:cluster:properties:upstream_conn_pool_reuse_type", "child_ids": ["xcsh-docs:resources:cluster:properties:upstream_conn_pool_reuse_type:disable_conn_pool_reuse", "xcsh-docs:resources:cluster:properties:upstream_conn_pool_reuse_type:enable_conn_pool_reuse"], "collection_id": "xcsh-docs:resources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:cluster:properties:upstream_conn_pool_reuse_type", "parent_id": "xcsh-docs:resources:cluster:reference", "path": "docs/guides/resources--cluster--properties--upstream_conn_pool_reuse_type.md", "provider_name": "cluster", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["upstream_conn_pool_reuse_type"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cluster/properties/upstream_conn_pool_reuse_type/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "upstream_conn_pool_reuse_type for xcsh_cluster.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# upstream_conn_pool_reuse_type

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md)
- [Property reference](resources--cluster--reference.md)
- upstream_conn_pool_reuse_type

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Select upstream connection pool reuse state for every downstream connection. This configuration
choice is for HTTP(S) LB only.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_conn_pool_reuse",
    "enable_conn_pool_reuse")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-map_downstream_to_upstream_conn_pool_type": "[\"disable_conn_pool_reuse\",\"enable_conn_pool_reuse\"]"
}
```

Terraform syntax:

```terraform
upstream_conn_pool_reuse_type {
  # Configure direct properties listed below.
}
```

## Direct properties

- [disable_conn_pool_reuse](resources--cluster--properties--upstream_conn_pool_reuse_type--disable_conn_pool_reuse.md): complete subsection reference.

- [enable_conn_pool_reuse](resources--cluster--properties--upstream_conn_pool_reuse_type--enable_conn_pool_reuse.md): complete subsection reference.

## Next pages

- [upstream_conn_pool_reuse_type.disable_conn_pool_reuse](resources--cluster--properties--upstream_conn_pool_reuse_type--disable_conn_pool_reuse.md)
- [upstream_conn_pool_reuse_type.enable_conn_pool_reuse](resources--cluster--properties--upstream_conn_pool_reuse_type--enable_conn_pool_reuse.md)
- [Property reference](resources--cluster--reference.md)
- [xcsh_cluster](../resources/cluster.md)
