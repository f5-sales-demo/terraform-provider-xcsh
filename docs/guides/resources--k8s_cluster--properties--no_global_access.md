---
page_title: "no_global_access"
subcategory: ""
description: "no_global_access for xcsh_k8s_cluster."
xcsh_docs: {"aliases": [], "body_bytes": 976, "body_sha256": "sha256:b9c751d62aa9c3e5ec97f3fb344f8ae4a754d9122c2da5e99a2e1cb847b2efe4", "canonical_id": "xcsh-docs:resources:k8s_cluster:properties:no_global_access", "child_ids": [], "collection_id": "xcsh-docs:resources:k8s_cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_cluster:properties:no_global_access", "parent_id": "xcsh-docs:resources:k8s_cluster:reference", "path": "docs/guides/resources--k8s_cluster--properties--no_global_access.md", "provider_name": "k8s_cluster", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["no_global_access"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_cluster/properties/no_global_access/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "no_global_access for xcsh_k8s_cluster.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# no_global_access

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md)
- [Property reference](resources--k8s_cluster--reference.md)
- no_global_access

<a id="section"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for no global access. Defaults to \`map\[\]\`. Server applies default when
omitted.

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
no_global_access = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--k8s_cluster--reference.md)
- [xcsh_k8s_cluster](../resources/k8s_cluster.md)
