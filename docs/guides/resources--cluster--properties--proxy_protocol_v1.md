---
page_title: "proxy_protocol_v1"
subcategory: ""
description: "proxy_protocol_v1 for xcsh_cluster."
xcsh_docs: {"aliases": [], "body_bytes": 884, "body_sha256": "sha256:1cff6113f8ac4fae047cb194cc606e97ec9d47bf773e801c2d67b59d6dad62eb", "canonical_id": "xcsh-docs:resources:cluster:properties:proxy_protocol_v1", "child_ids": [], "collection_id": "xcsh-docs:resources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:cluster:properties:proxy_protocol_v1", "parent_id": "xcsh-docs:resources:cluster:reference", "path": "docs/guides/resources--cluster--properties--proxy_protocol_v1.md", "provider_name": "cluster", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["proxy_protocol_v1"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cluster/properties/proxy_protocol_v1/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "proxy_protocol_v1 for xcsh_cluster.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# proxy_protocol_v1

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md)
- [Property reference](resources--cluster--reference.md)
- proxy_protocol_v1

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for proxy protocol v1.

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
proxy_protocol_v1 = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--cluster--reference.md)
- [xcsh_cluster](../resources/cluster.md)
