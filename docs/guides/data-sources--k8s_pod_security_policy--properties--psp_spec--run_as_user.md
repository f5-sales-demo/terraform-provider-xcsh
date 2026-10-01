---
page_title: "psp_spec.run_as_user"
subcategory: ""
description: "psp_spec.run_as_user for xcsh_k8s_pod_security_policy."
xcsh_docs: {"aliases": [], "body_bytes": 2125, "body_sha256": "sha256:486e282f4b8e2d0f9705c50b2740c096f843ea13f77b74b0eca69097d1fa57d0", "canonical_id": "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec:run_as_user", "child_ids": ["xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec:run_as_user:id_ranges"], "collection_id": "xcsh-docs:data-sources:k8s_pod_security_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec:run_as_user", "parent_id": "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec", "path": "docs/guides/data-sources--k8s_pod_security_policy--properties--psp_spec--run_as_user.md", "provider_name": "k8s_pod_security_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["psp_spec", "run_as_user"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/k8s_pod_security_policy/properties/psp_spec/run_as_user/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "psp_spec.run_as_user for xcsh_k8s_pod_security_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_pod_security_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# psp_spec.run_as_user

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md)
- [Property reference](data-sources--k8s_pod_security_policy--reference.md)
- [psp_spec](data-sources--k8s_pod_security_policy--properties--psp_spec.md)
- psp_spec.run_as_user

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for run as user.

Upstream description:

ID ranges and rules.

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

## Direct properties

- [id_ranges](data-sources--k8s_pod_security_policy--properties--psp_spec--run_as_user--id_ranges.md): complete subsection reference.

<a id="schema-psp_spec--run_as_user--rule"></a>

### rule property

Type: `"string"`. Computed.

Rule indicated how the FS group ID range is used.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

## Next pages

- [psp_spec.run_as_user.id_ranges](data-sources--k8s_pod_security_policy--properties--psp_spec--run_as_user--id_ranges.md)
- [psp_spec](data-sources--k8s_pod_security_policy--properties--psp_spec.md)
- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md)
