---
page_title: "cluster_scoped_access_deny"
subcategory: ""
description: "cluster_scoped_access_deny for xcsh_k8s_cluster."
xcsh_docs: {"aliases": [], "body_bytes": 1376, "body_sha256": "sha256:16ea174ddd281daf2f313707553be8341cbdea07e731bc7b81fa1d49ed4b8b94", "canonical_id": "xcsh-docs:resources:k8s_cluster:properties:cluster_scoped_access_deny", "child_ids": [], "collection_id": "xcsh-docs:resources:k8s_cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_cluster:properties:cluster_scoped_access_deny", "parent_id": "xcsh-docs:resources:k8s_cluster:reference", "path": "docs/guides/resources--k8s_cluster--properties--cluster_scoped_access_deny.md", "provider_name": "k8s_cluster", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cluster_scoped_access_deny"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_cluster/properties/cluster_scoped_access_deny/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cluster_scoped_access_deny for xcsh_k8s_cluster.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cluster_scoped_access_deny

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md)
- [Property reference](resources--k8s_cluster--reference.md)
- cluster_scoped_access_deny

<a id="section"></a>

Type: `["object", {}]`. Optional, Computed.

\[OneOf: cluster\_scoped\_access\_deny, cluster\_scoped\_access\_permit\] Enable this option.
Defaults to \`map\[\]\`. Server applies default when omitted.

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

- [cluster_scoped_access_deny](resources--k8s_cluster--properties--cluster_scoped_access_deny.md#section)
- [cluster_scoped_access_permit](resources--k8s_cluster--properties--cluster_scoped_access_permit.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
cluster_scoped_access_deny = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--k8s_cluster--reference.md)
- [xcsh_k8s_cluster](../resources/k8s_cluster.md)
