---
page_title: "tls_parameters.disable_session_key_caching"
subcategory: ""
description: "tls_parameters.disable_session_key_caching for xcsh_cluster."
xcsh_docs: {"aliases": [], "body_bytes": 938, "body_sha256": "sha256:b31ba4591cffd3e9564694953b0623668d7d3f7b1c595c98b0fa2680b407acd0", "canonical_id": "xcsh-docs:resources:cluster:properties:tls_parameters:disable_session_key_caching", "child_ids": [], "collection_id": "xcsh-docs:resources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:cluster:properties:tls_parameters:disable_session_key_caching", "parent_id": "xcsh-docs:resources:cluster:properties:tls_parameters", "path": "docs/guides/resources--cluster--properties--tls_parameters--disable_session_key_caching.md", "provider_name": "cluster", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tls_parameters", "disable_session_key_caching"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cluster/properties/tls_parameters/disable_session_key_caching/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tls_parameters.disable_session_key_caching for xcsh_cluster.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# tls_parameters.disable_session_key_caching

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md)
- [Property reference](resources--cluster--reference.md)
- [tls_parameters](resources--cluster--properties--tls_parameters.md)
- tls_parameters.disable_session_key_caching

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable session key caching.

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
disable_session_key_caching = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [tls_parameters](resources--cluster--properties--tls_parameters.md)
- [xcsh_cluster](../resources/cluster.md)
