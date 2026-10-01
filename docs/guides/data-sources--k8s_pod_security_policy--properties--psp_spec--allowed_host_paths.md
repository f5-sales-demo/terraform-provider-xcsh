---
page_title: "psp_spec.allowed_host_paths"
subcategory: ""
description: "psp_spec.allowed_host_paths for xcsh_k8s_pod_security_policy."
xcsh_docs: {"aliases": [], "body_bytes": 2936, "body_sha256": "sha256:1b6096c01a483b5b6696b44d9b364fb9360b2b9e9454bafceb9c546674d170b7", "canonical_id": "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec:allowed_host_paths", "child_ids": [], "collection_id": "xcsh-docs:data-sources:k8s_pod_security_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec:allowed_host_paths", "parent_id": "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec", "path": "docs/guides/data-sources--k8s_pod_security_policy--properties--psp_spec--allowed_host_paths.md", "provider_name": "k8s_pod_security_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["psp_spec", "allowed_host_paths"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/k8s_pod_security_policy/properties/psp_spec/allowed_host_paths/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "psp_spec.allowed_host_paths for xcsh_k8s_pod_security_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_pod_security_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# psp_spec.allowed_host_paths

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md)
- [Property reference](data-sources--k8s_pod_security_policy--reference.md)
- [psp_spec](data-sources--k8s_pod_security_policy--properties--psp_spec.md)
- psp_spec.allowed_host_paths

<a id="section"></a>

Type: `"list"`. Computed.

Restrict list of host paths, default all host paths are allowed.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

<a id="schema-psp_spec--allowed_host_paths--path_prefix"></a>

### path_prefix property

Type: `"string"`. Computed.

Host path prefix is the path prefix that the host volume must match. It does not support \*.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="schema-psp_spec--allowed_host_paths--read_only"></a>

### read_only property

Type: `"bool"`. Computed.

Volume will be allowed to mount read only.

Upstream description:

This volume will be allowed to mount read only.

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

## Next pages

- [psp_spec](data-sources--k8s_pod_security_policy--properties--psp_spec.md)
- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md)
