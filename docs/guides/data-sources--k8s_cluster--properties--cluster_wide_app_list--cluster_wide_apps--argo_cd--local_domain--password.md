---
page_title: "cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password"
subcategory: ""
description: "cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password for xcsh_k8s_cluster."
xcsh_docs: {"aliases": [], "body_bytes": 2326, "body_sha256": "sha256:cfc4cf161325f40c63e84c0ecd89d3f27a5dfb58ee602cd6dafa7e0ab9f257df", "canonical_id": "xcsh-docs:data-sources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:argo_cd:local_domain:password", "child_ids": ["xcsh-docs:data-sources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:argo_cd:local_domain:password:blindfold_secret_info", "xcsh-docs:data-sources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:argo_cd:local_domain:password:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:k8s_cluster:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:argo_cd:local_domain:password", "parent_id": "xcsh-docs:data-sources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:argo_cd:local_domain", "path": "docs/guides/data-sources--k8s_cluster--properties--cluster_wide_app_list--cluster_wide_apps--argo_cd--local_domain--password.md", "provider_name": "k8s_cluster", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cluster_wide_app_list", "cluster_wide_apps", "argo_cd", "local_domain", "password"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/argo_cd/local_domain/password/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password for xcsh_k8s_cluster.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md)
- [Property reference](data-sources--k8s_cluster--reference.md)
- [cluster_wide_app_list](data-sources--k8s_cluster--properties--cluster_wide_app_list.md)
- [cluster_wide_app_list.cluster_wide_apps](data-sources--k8s_cluster--properties--cluster_wide_app_list--cluster_wide_apps.md)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd](data-sources--k8s_cluster--properties--cluster_wide_app_list--cluster_wide_apps--argo_cd.md)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain](data-sources--k8s_cluster--properties--cluster_wide_app_list--cluster_wide_apps--argo_cd--local_domain.md)
- cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password

<a id="section"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

## Direct properties

- [blindfold_secret_info](data-sources--k8s_cluster--properties--cluster_wide_app_list--cluster_wide_apps--argo_cd--local_domain--password--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](data-sources--k8s_cluster--properties--cluster_wide_app_list--cluster_wide_apps--argo_cd--local_domain--password--clear_secret_info.md): complete subsection reference.

## Next pages

- [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_secret_info](data-sources--k8s_cluster--properties--cluster_wide_app_list--cluster_wide_apps--argo_cd--local_domain--password--blindfold_secret_info.md)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.clear_secret_info](data-sources--k8s_cluster--properties--cluster_wide_app_list--cluster_wide_apps--argo_cd--local_domain--password--clear_secret_info.md)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain](data-sources--k8s_cluster--properties--cluster_wide_app_list--cluster_wide_apps--argo_cd--local_domain.md)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md)
