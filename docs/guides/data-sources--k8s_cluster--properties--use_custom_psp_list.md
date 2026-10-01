---
page_title: "use_custom_psp_list"
subcategory: ""
description: "use_custom_psp_list for xcsh_k8s_cluster."
xcsh_docs: {"aliases": [], "body_bytes": 1439, "body_sha256": "sha256:30eaf5a8332ebfcc499ae32ef5cb9cdadd3fcba7ee86ea02f0ecb4aa29745b9f", "canonical_id": "xcsh-docs:data-sources:k8s_cluster:properties:use_custom_psp_list", "child_ids": ["xcsh-docs:data-sources:k8s_cluster:properties:use_custom_psp_list:pod_security_policies"], "collection_id": "xcsh-docs:data-sources:k8s_cluster:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:k8s_cluster:properties:use_custom_psp_list", "parent_id": "xcsh-docs:data-sources:k8s_cluster:reference", "path": "docs/guides/data-sources--k8s_cluster--properties--use_custom_psp_list.md", "provider_name": "k8s_cluster", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["use_custom_psp_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/k8s_cluster/properties/use_custom_psp_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "use_custom_psp_list for xcsh_k8s_cluster.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# use_custom_psp_list

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md)
- [Property reference](data-sources--k8s_cluster--reference.md)
- use_custom_psp_list

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: use\_custom\_psp\_list, use\_default\_psp; Default: use\_default\_psp\] List of active Pod
security policies for a K8s cluster.

Upstream description:

List of active Pod security policies for a K8s cluster.

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

- [use_custom_psp_list](data-sources--k8s_cluster--properties--use_custom_psp_list.md#section)
- [use_default_psp](data-sources--k8s_cluster--properties--use_default_psp.md#section)

Select alternatives according to the provider validators above.

## Direct properties

- [pod_security_policies](data-sources--k8s_cluster--properties--use_custom_psp_list--pod_security_policies.md): complete subsection reference.

## Next pages

- [use_custom_psp_list.pod_security_policies](data-sources--k8s_cluster--properties--use_custom_psp_list--pod_security_policies.md)
- [Property reference](data-sources--k8s_cluster--reference.md)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md)
