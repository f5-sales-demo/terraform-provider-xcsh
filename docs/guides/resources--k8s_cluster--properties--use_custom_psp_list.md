---
page_title: "use_custom_psp_list"
subcategory: ""
description: "use_custom_psp_list for xcsh_k8s_cluster."
xcsh_docs: {"aliases": [], "body_bytes": 1703, "body_sha256": "sha256:e329453bbbf6ddd3c1b7f42a36acdbbe523e18e0030228cd81c0cf8e1ebda4e5", "canonical_id": "xcsh-docs:resources:k8s_cluster:properties:use_custom_psp_list", "child_ids": ["xcsh-docs:resources:k8s_cluster:properties:use_custom_psp_list:pod_security_policies"], "collection_id": "xcsh-docs:resources:k8s_cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_cluster:properties:use_custom_psp_list", "parent_id": "xcsh-docs:resources:k8s_cluster:reference", "path": "docs/guides/resources--k8s_cluster--properties--use_custom_psp_list.md", "provider_name": "k8s_cluster", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["use_custom_psp_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_cluster/properties/use_custom_psp_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "use_custom_psp_list for xcsh_k8s_cluster.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# use_custom_psp_list

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md)
- [Property reference](resources--k8s_cluster--reference.md)
- use_custom_psp_list

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: use\_custom\_psp\_list, use\_default\_psp; Default: use\_default\_psp\] List of active Pod
security policies for a K8s cluster.

Upstream description:

List of active Pod security policies for a K8s cluster.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("pod_security_policies")}
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

- [use_custom_psp_list](resources--k8s_cluster--properties--use_custom_psp_list.md#section)
- [use_default_psp](resources--k8s_cluster--properties--use_default_psp.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
use_custom_psp_list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [pod_security_policies](resources--k8s_cluster--properties--use_custom_psp_list--pod_security_policies.md): complete subsection reference.

## Next pages

- [use_custom_psp_list.pod_security_policies](resources--k8s_cluster--properties--use_custom_psp_list--pod_security_policies.md)
- [Property reference](resources--k8s_cluster--reference.md)
- [xcsh_k8s_cluster](../resources/k8s_cluster.md)
