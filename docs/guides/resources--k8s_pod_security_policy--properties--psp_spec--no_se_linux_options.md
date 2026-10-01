---
page_title: "psp_spec.no_se_linux_options"
subcategory: ""
description: "psp_spec.no_se_linux_options for xcsh_k8s_pod_security_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1094, "body_sha256": "sha256:564d48d652a0aef79d238117f13d59b67d5635f69f2beb8f86799714a787e785", "canonical_id": "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:no_se_linux_options", "child_ids": [], "collection_id": "xcsh-docs:resources:k8s_pod_security_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:no_se_linux_options", "parent_id": "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec", "path": "docs/guides/resources--k8s_pod_security_policy--properties--psp_spec--no_se_linux_options.md", "provider_name": "k8s_pod_security_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["psp_spec", "no_se_linux_options"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_pod_security_policy/properties/psp_spec/no_se_linux_options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "psp_spec.no_se_linux_options for xcsh_k8s_pod_security_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_pod_security_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# psp_spec.no_se_linux_options

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md)
- [Property reference](resources--k8s_pod_security_policy--reference.md)
- [psp_spec](resources--k8s_pod_security_policy--properties--psp_spec.md)
- psp_spec.no_se_linux_options

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for no se linux options.

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
no_se_linux_options {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [psp_spec](resources--k8s_pod_security_policy--properties--psp_spec.md)
- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md)
