---
page_title: "auto_http_config"
subcategory: ""
description: "auto_http_config for xcsh_cluster."
xcsh_docs: {"aliases": [], "body_bytes": 1252, "body_sha256": "sha256:3cce74553489499be209218c47c1c775211eac716850874c7a1097bcfb879366", "canonical_id": "xcsh-docs:resources:cluster:properties:auto_http_config", "child_ids": [], "collection_id": "xcsh-docs:resources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:cluster:properties:auto_http_config", "parent_id": "xcsh-docs:resources:cluster:reference", "path": "docs/guides/resources--cluster--properties--auto_http_config.md", "provider_name": "cluster", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["auto_http_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cluster/properties/auto_http_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "auto_http_config for xcsh_cluster.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# auto_http_config

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md)
- [Property reference](resources--cluster--reference.md)
- auto_http_config

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: auto\_http\_config, http1\_config, http2\_options\] Enable this option

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

OneOf alternatives in this subsection:

- [auto_http_config](resources--cluster--properties--auto_http_config.md#section)
- [http1_config](resources--cluster--properties--http1_config.md#section)
- [http2_options](resources--cluster--properties--http2_options.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
auto_http_config = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--cluster--reference.md)
- [xcsh_cluster](../resources/cluster.md)
