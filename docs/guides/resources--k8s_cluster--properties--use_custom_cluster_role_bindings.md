---
page_title: "use_custom_cluster_role_bindings"
subcategory: ""
description: "use_custom_cluster_role_bindings for xcsh_k8s_cluster."
xcsh_docs: {"aliases": [], "body_bytes": 1905, "body_sha256": "sha256:1f803d6f5d04c8e3509341817eb8dc4e4197362015f526b16a30f6b21c3ff9a4", "canonical_id": "xcsh-docs:resources:k8s_cluster:properties:use_custom_cluster_role_bindings", "child_ids": ["xcsh-docs:resources:k8s_cluster:properties:use_custom_cluster_role_bindings:cluster_role_bindings"], "collection_id": "xcsh-docs:resources:k8s_cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_cluster:properties:use_custom_cluster_role_bindings", "parent_id": "xcsh-docs:resources:k8s_cluster:reference", "path": "docs/guides/resources--k8s_cluster--properties--use_custom_cluster_role_bindings.md", "provider_name": "k8s_cluster", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["use_custom_cluster_role_bindings"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_cluster/properties/use_custom_cluster_role_bindings/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "use_custom_cluster_role_bindings for xcsh_k8s_cluster.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# use_custom_cluster_role_bindings

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md)
- [Property reference](resources--k8s_cluster--reference.md)
- use_custom_cluster_role_bindings

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: use\_custom\_cluster\_role\_bindings, use\_default\_cluster\_role\_bindings; Default:
use\_default\_cluster\_role\_bindings\] List of active cluster role binding list for a K8s cluster.

Upstream description:

List of active cluster role binding list for a K8s cluster.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cluster_role_bindings")}
```

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

- [use_custom_cluster_role_bindings](resources--k8s_cluster--properties--use_custom_cluster_role_bindings.md#section)
- [use_default_cluster_role_bindings](resources--k8s_cluster--properties--use_default_cluster_role_bindings.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
use_custom_cluster_role_bindings {
  # Configure direct properties listed below.
}
```

## Direct properties

- [cluster_role_bindings](resources--k8s_cluster--properties--use_custom_cluster_role_bindings--cluster_role_bindings.md): complete subsection reference.

## Next pages

- [use_custom_cluster_role_bindings.cluster_role_bindings](resources--k8s_cluster--properties--use_custom_cluster_role_bindings--cluster_role_bindings.md)
- [Property reference](resources--k8s_cluster--reference.md)
- [xcsh_k8s_cluster](../resources/k8s_cluster.md)
