---
page_title: "cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain"
subcategory: ""
description: "cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain for xcsh_k8s_cluster."
xcsh_docs: {"aliases": [], "body_bytes": 5112, "body_sha256": "sha256:236f6a46f16101b5dfdac4a7ef330d8c86c803155d584b688b68dee15d0a2ca2", "child_ids": ["xcsh-docs:data-sources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:argo_cd:local_domain:default_port", "xcsh-docs:data-sources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:argo_cd:local_domain:password"], "collection_id": "xcsh-docs:data-sources:k8s_cluster:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:argo_cd:local_domain", "parent_id": "xcsh-docs:data-sources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:argo_cd", "path": "documentation/data-sources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/argo_cd/local_domain/index.md", "provider_name": "k8s_cluster", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["cluster_wide_app_list", "cluster_wide_apps", "argo_cd", "local_domain"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/argo_cd/local_domain/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain for xcsh_k8s_cluster.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain

Breadcrumbs:

- [xcsh_k8s_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/properties/)
- [cluster_wide_app_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/properties/cluster_wide_app_list/)
- [cluster_wide_app_list.cluster_wide_apps](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/argo_cd/)
- cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain

<a id="section"></a>

Type: `"single"`. Computed.

Parameters required to enable local access.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_choice": "[\"default_port\",\"port\"]"
}
```

## Direct properties

- [default_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/argo_cd/local_domain/default_port/): complete subsection reference.

<a id="schema-cluster_wide_app_list--cluster_wide_apps--argo_cd--local_domain--local_domain"></a>

### local_domain property

Type: `"string"`. Computed.

ArgoCD will be accessible at &lt;site name&gt;.&lt;local domain&gt;.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 192,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 192,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "192",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "192",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [password](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/argo_cd/local_domain/password/): complete subsection reference.

<a id="schema-cluster_wide_app_list--cluster_wide_apps--argo_cd--local_domain--port"></a>

### port property

Type: `"number"`. Computed.

Exclusive with \[default\_port\] Use custom ArgoCD port. Available port range is less than 65000
except reserved ports.

Upstream description:

Exclusive with \[default\_port\] Use custom ArgoCD port. Available port range is less than 65000
except reserved ports.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "networking",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.not_in_ranges": "0,6443,8005-8007,8443-8444,8505-8507,9005-9007,9090,9505-9507,9100,9115,9999,20914,23802,30805,30855,30905,30955,32222,18091-18095,65000-65334"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.not_in_ranges": "0,6443,8005-8007,8443-8444,8505-8507,9005-9007,9090,9505-9507,9100,9115,9999,20914,23802,30805,30855,30905,30955,32222,18091-18095,65000-65334"
  }
}
```

## Next pages

- [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.default_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/argo_cd/local_domain/default_port/)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/argo_cd/local_domain/password/)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/argo_cd/)
- [xcsh_k8s_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/)
