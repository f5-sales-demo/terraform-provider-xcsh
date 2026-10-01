---
page_title: "pod_security_admission_specs"
subcategory: ""
description: "pod_security_admission_specs for xcsh_k8s_pod_security_admission."
xcsh_docs: {"aliases": [], "body_bytes": 3362, "body_sha256": "sha256:f7b8c700c58f75c84c9f2711a41c385a13cc1ade0726b5f35eed74b92beea237", "canonical_id": "xcsh-docs:data-sources:k8s_pod_security_admission:properties:pod_security_admission_specs", "child_ids": ["xcsh-docs:data-sources:k8s_pod_security_admission:properties:pod_security_admission_specs:audit", "xcsh-docs:data-sources:k8s_pod_security_admission:properties:pod_security_admission_specs:baseline", "xcsh-docs:data-sources:k8s_pod_security_admission:properties:pod_security_admission_specs:enforce", "xcsh-docs:data-sources:k8s_pod_security_admission:properties:pod_security_admission_specs:privileged", "xcsh-docs:data-sources:k8s_pod_security_admission:properties:pod_security_admission_specs:restricted", "xcsh-docs:data-sources:k8s_pod_security_admission:properties:pod_security_admission_specs:warn"], "collection_id": "xcsh-docs:data-sources:k8s_pod_security_admission:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:k8s_pod_security_admission:properties:pod_security_admission_specs", "parent_id": "xcsh-docs:data-sources:k8s_pod_security_admission:reference", "path": "docs/guides/data-sources--k8s_pod_security_admission--properties--pod_security_admission_specs.md", "provider_name": "k8s_pod_security_admission", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["pod_security_admission_specs"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/k8s_pod_security_admission/properties/pod_security_admission_specs/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "pod_security_admission_specs for xcsh_k8s_pod_security_admission.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_pod_security_admissionCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# pod_security_admission_specs

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](../data-sources/k8s_pod_security_admission.md)
- [Property reference](data-sources--k8s_pod_security_admission--reference.md)
- pod_security_admission_specs

<a id="section"></a>

Type: `"list"`. Computed.

K8s Pod Security Admission. Uniform Resource Identifier

Upstream description:

Uniform Resource Identifier

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

- [audit](data-sources--k8s_pod_security_admission--properties--pod_security_admission_specs--audit.md): complete subsection reference.

- [baseline](data-sources--k8s_pod_security_admission--properties--pod_security_admission_specs--baseline.md): complete subsection reference.

- [enforce](data-sources--k8s_pod_security_admission--properties--pod_security_admission_specs--enforce.md): complete subsection reference.

- [privileged](data-sources--k8s_pod_security_admission--properties--pod_security_admission_specs--privileged.md): complete subsection reference.

- [restricted](data-sources--k8s_pod_security_admission--properties--pod_security_admission_specs--restricted.md): complete subsection reference.

- [warn](data-sources--k8s_pod_security_admission--properties--pod_security_admission_specs--warn.md): complete subsection reference.

## Next pages

- [pod_security_admission_specs.audit](data-sources--k8s_pod_security_admission--properties--pod_security_admission_specs--audit.md)
- [pod_security_admission_specs.baseline](data-sources--k8s_pod_security_admission--properties--pod_security_admission_specs--baseline.md)
- [pod_security_admission_specs.enforce](data-sources--k8s_pod_security_admission--properties--pod_security_admission_specs--enforce.md)
- [pod_security_admission_specs.privileged](data-sources--k8s_pod_security_admission--properties--pod_security_admission_specs--privileged.md)
- [pod_security_admission_specs.restricted](data-sources--k8s_pod_security_admission--properties--pod_security_admission_specs--restricted.md)
- [pod_security_admission_specs.warn](data-sources--k8s_pod_security_admission--properties--pod_security_admission_specs--warn.md)
- [Property reference](data-sources--k8s_pod_security_admission--reference.md)
- [xcsh_k8s_pod_security_admission](../data-sources/k8s_pod_security_admission.md)
