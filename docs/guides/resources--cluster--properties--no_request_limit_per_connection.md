---
page_title: "no_request_limit_per_connection"
subcategory: ""
description: "no_request_limit_per_connection for xcsh_cluster."
xcsh_docs: {"aliases": [], "body_bytes": 940, "body_sha256": "sha256:4a6c87fdadb813f9ede8f35c6bdaaa3e665140463570575d1d6c49b901c8062c", "canonical_id": "xcsh-docs:resources:cluster:properties:no_request_limit_per_connection", "child_ids": [], "collection_id": "xcsh-docs:resources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:cluster:properties:no_request_limit_per_connection", "parent_id": "xcsh-docs:resources:cluster:reference", "path": "docs/guides/resources--cluster--properties--no_request_limit_per_connection.md", "provider_name": "cluster", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["no_request_limit_per_connection"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cluster/properties/no_request_limit_per_connection/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "no_request_limit_per_connection for xcsh_cluster.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# no_request_limit_per_connection

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md)
- [Property reference](resources--cluster--reference.md)
- no_request_limit_per_connection

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no request limit per connection.

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
no_request_limit_per_connection = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--cluster--reference.md)
- [xcsh_cluster](../resources/cluster.md)
