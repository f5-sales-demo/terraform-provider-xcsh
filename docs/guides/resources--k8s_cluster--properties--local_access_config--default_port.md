---
page_title: "local_access_config.default_port"
subcategory: ""
description: "local_access_config.default_port for xcsh_k8s_cluster."
xcsh_docs: {"aliases": [], "body_bytes": 913, "body_sha256": "sha256:20c2c1d151850e24a9f45d845b82328dd0403298a89a9b551c3ce9cbe879895c", "canonical_id": "xcsh-docs:resources:k8s_cluster:properties:local_access_config:default_port", "child_ids": [], "collection_id": "xcsh-docs:resources:k8s_cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_cluster:properties:local_access_config:default_port", "parent_id": "xcsh-docs:resources:k8s_cluster:properties:local_access_config", "path": "docs/guides/resources--k8s_cluster--properties--local_access_config--default_port.md", "provider_name": "k8s_cluster", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["local_access_config", "default_port"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_cluster/properties/local_access_config/default_port/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "local_access_config.default_port for xcsh_k8s_cluster.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# local_access_config.default_port

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md)
- [Property reference](resources--k8s_cluster--reference.md)
- [local_access_config](resources--k8s_cluster--properties--local_access_config.md)
- local_access_config.default_port

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
default_port = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [local_access_config](resources--k8s_cluster--properties--local_access_config.md)
- [xcsh_k8s_cluster](../resources/k8s_cluster.md)
