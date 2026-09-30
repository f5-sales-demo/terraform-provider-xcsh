---
page_title: "http2_options"
subcategory: ""
description: "http2_options for xcsh_cluster."
xcsh_docs: {"aliases": [], "body_bytes": 1006, "body_sha256": "sha256:63f79fa1398d4b1f7d536b11b8018639362e31e754affca5ed2061606d9dee5f", "canonical_id": "xcsh-docs:resources:cluster:properties:http2_options", "child_ids": [], "collection_id": "xcsh-docs:resources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:cluster:properties:http2_options", "parent_id": "xcsh-docs:resources:cluster:reference", "path": "docs/guides/resources--cluster--properties--http2_options.md", "provider_name": "cluster", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["http2_options"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cluster/properties/http2_options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "http2_options for xcsh_cluster.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# http2_options

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md)
- [Property reference](resources--cluster--reference.md)
- http2_options

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Http2 Protocol OPTIONS for upstream connections.

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
http2_options {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-http2_options--enabled"></a>

### enabled property

Type: `"bool"`. Optional.

Enable/disable HTTP2 Protocol for upstream connections.

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

## Next pages

- [Property reference](resources--cluster--reference.md)
- [xcsh_cluster](../resources/cluster.md)
