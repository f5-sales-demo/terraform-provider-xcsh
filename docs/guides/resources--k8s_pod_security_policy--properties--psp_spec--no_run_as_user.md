---
page_title: "psp_spec.no_run_as_user"
subcategory: ""
description: "psp_spec.no_run_as_user for xcsh_k8s_pod_security_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1061, "body_sha256": "sha256:d14c01c5b68a15d3cc87cc1c06cfcee6412acf47b21d4f2d0136d5f47779124e", "canonical_id": "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:no_run_as_user", "child_ids": [], "collection_id": "xcsh-docs:resources:k8s_pod_security_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:no_run_as_user", "parent_id": "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec", "path": "docs/guides/resources--k8s_pod_security_policy--properties--psp_spec--no_run_as_user.md", "provider_name": "k8s_pod_security_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["psp_spec", "no_run_as_user"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_pod_security_policy/properties/psp_spec/no_run_as_user/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "psp_spec.no_run_as_user for xcsh_k8s_pod_security_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_pod_security_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# psp_spec.no_run_as_user

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md)
- [Property reference](resources--k8s_pod_security_policy--reference.md)
- [psp_spec](resources--k8s_pod_security_policy--properties--psp_spec.md)
- psp_spec.no_run_as_user

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no run as user.

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
no_run_as_user = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [psp_spec](resources--k8s_pod_security_policy--properties--psp_spec.md)
- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md)
