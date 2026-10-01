---
page_title: "use_custom_cluster_role_list"
subcategory: ""
description: "use_custom_cluster_role_list for xcsh_k8s_cluster."
xcsh_docs: {"aliases": [], "body_bytes": 1779, "body_sha256": "sha256:c230a76142ec3a32e11aaf05061686074c7077a86e9c148f197aea16ca591f61", "canonical_id": "xcsh-docs:resources:k8s_cluster:properties:use_custom_cluster_role_list", "child_ids": ["xcsh-docs:resources:k8s_cluster:properties:use_custom_cluster_role_list:cluster_roles"], "collection_id": "xcsh-docs:resources:k8s_cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_cluster:properties:use_custom_cluster_role_list", "parent_id": "xcsh-docs:resources:k8s_cluster:reference", "path": "docs/guides/resources--k8s_cluster--properties--use_custom_cluster_role_list.md", "provider_name": "k8s_cluster", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["use_custom_cluster_role_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_cluster/properties/use_custom_cluster_role_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "use_custom_cluster_role_list for xcsh_k8s_cluster.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# use_custom_cluster_role_list

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md)
- [Property reference](resources--k8s_cluster--reference.md)
- use_custom_cluster_role_list

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: use\_custom\_cluster\_role\_list, use\_default\_cluster\_roles; Default:
use\_default\_cluster\_roles\] List of active cluster role list for a K8s cluster.

Upstream description:

List of active cluster role list for a K8s cluster.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cluster_roles")}
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

- [use_custom_cluster_role_list](resources--k8s_cluster--properties--use_custom_cluster_role_list.md#section)
- [use_default_cluster_roles](resources--k8s_cluster--properties--use_default_cluster_roles.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
use_custom_cluster_role_list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [cluster_roles](resources--k8s_cluster--properties--use_custom_cluster_role_list--cluster_roles.md): complete subsection reference.

## Next pages

- [use_custom_cluster_role_list.cluster_roles](resources--k8s_cluster--properties--use_custom_cluster_role_list--cluster_roles.md)
- [Property reference](resources--k8s_cluster--reference.md)
- [xcsh_k8s_cluster](../resources/k8s_cluster.md)
