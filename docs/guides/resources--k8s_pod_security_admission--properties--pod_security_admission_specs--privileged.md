---
page_title: "pod_security_admission_specs.privileged"
subcategory: ""
description: "pod_security_admission_specs.privileged for xcsh_k8s_pod_security_admission."
xcsh_docs: {"aliases": [], "body_bytes": 1165, "body_sha256": "sha256:d03eb31171d8f14d750c593f548483a1707f85de0be449478fc9bd654613ea8c", "canonical_id": "xcsh-docs:resources:k8s_pod_security_admission:properties:pod_security_admission_specs:privileged", "child_ids": [], "collection_id": "xcsh-docs:resources:k8s_pod_security_admission:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_pod_security_admission:properties:pod_security_admission_specs:privileged", "parent_id": "xcsh-docs:resources:k8s_pod_security_admission:properties:pod_security_admission_specs", "path": "docs/guides/resources--k8s_pod_security_admission--properties--pod_security_admission_specs--privileged.md", "provider_name": "k8s_pod_security_admission", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["pod_security_admission_specs", "privileged"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_pod_security_admission/properties/pod_security_admission_specs/privileged/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "pod_security_admission_specs.privileged for xcsh_k8s_pod_security_admission.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_pod_security_admissionCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# pod_security_admission_specs.privileged

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](../resources/k8s_pod_security_admission.md)
- [Property reference](resources--k8s_pod_security_admission--reference.md)
- [pod_security_admission_specs](resources--k8s_pod_security_admission--properties--pod_security_admission_specs.md)
- pod_security_admission_specs.privileged

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
privileged = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [pod_security_admission_specs](resources--k8s_pod_security_admission--properties--pod_security_admission_specs.md)
- [xcsh_k8s_pod_security_admission](../resources/k8s_pod_security_admission.md)
