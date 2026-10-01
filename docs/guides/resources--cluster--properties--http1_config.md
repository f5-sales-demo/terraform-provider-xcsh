---
page_title: "http1_config"
subcategory: ""
description: "http1_config for xcsh_cluster."
xcsh_docs: {"aliases": [], "body_bytes": 1022, "body_sha256": "sha256:5dfdd62bf38a56c3ca6481a1785bd310e4359de0c238677bb015976cd5498d69", "canonical_id": "xcsh-docs:resources:cluster:properties:http1_config", "child_ids": ["xcsh-docs:resources:cluster:properties:http1_config:header_transformation"], "collection_id": "xcsh-docs:resources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:cluster:properties:http1_config", "parent_id": "xcsh-docs:resources:cluster:reference", "path": "docs/guides/resources--cluster--properties--http1_config.md", "provider_name": "cluster", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["http1_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cluster/properties/http1_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "http1_config for xcsh_cluster.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http1_config

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md)
- [Property reference](resources--cluster--reference.md)
- http1_config

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

HTTP/1.1 Protocol OPTIONS for upstream connections.

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
http1_config {
  # Configure direct properties listed below.
}
```

## Direct properties

- [header_transformation](resources--cluster--properties--http1_config--header_transformation.md): complete subsection reference.

## Next pages

- [http1_config.header_transformation](resources--cluster--properties--http1_config--header_transformation.md)
- [Property reference](resources--cluster--reference.md)
- [xcsh_cluster](../resources/cluster.md)
