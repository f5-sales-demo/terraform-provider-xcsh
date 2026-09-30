---
page_title: "no_request_limit_per_connection"
subcategory: ""
description: "no_request_limit_per_connection for xcsh_cluster."
xcsh_docs: {"aliases": [], "body_bytes": 841, "body_sha256": "sha256:5c72d63889a5aa554df0ab594726b66fce05c477b1237b42fe216f7f120f7fef", "canonical_id": "xcsh-docs:resources:cluster:properties:no_request_limit_per_connection", "child_ids": [], "collection_id": "xcsh-docs:resources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:cluster:properties:no_request_limit_per_connection", "parent_id": "xcsh-docs:resources:cluster:reference", "path": "docs/guides/resources--cluster--properties--no_request_limit_per_connection.md", "provider_name": "cluster", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["no_request_limit_per_connection"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cluster/properties/no_request_limit_per_connection/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "no_request_limit_per_connection for xcsh_cluster.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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
