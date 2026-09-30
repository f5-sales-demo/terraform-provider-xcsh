---
page_title: "psp_spec.fs_group_strategy_options"
subcategory: ""
description: "psp_spec.fs_group_strategy_options for xcsh_k8s_pod_security_policy."
xcsh_docs: {"aliases": [], "body_bytes": 2877, "body_sha256": "sha256:2a1b1e3a8f980d78d02afe5316e967de39f36a1efa5537a8e67f6cef6d11bdb4", "child_ids": ["xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:fs_group_strategy_options:id_ranges"], "collection_id": "xcsh-docs:resources:k8s_pod_security_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:fs_group_strategy_options", "parent_id": "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec", "path": "documentation/resources/k8s_pod_security_policy/properties/psp_spec/fs_group_strategy_options/index.md", "provider_name": "k8s_pod_security_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["psp_spec", "fs_group_strategy_options"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_pod_security_policy/properties/psp_spec/fs_group_strategy_options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "psp_spec.fs_group_strategy_options for xcsh_k8s_pod_security_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_pod_security_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# psp_spec.fs_group_strategy_options

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/)
- [psp_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/)
- psp_spec.fs_group_strategy_options

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for fs group strategy options.

Upstream description:

ID ranges and rules.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("rule")}
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

Terraform syntax:

```terraform
fs_group_strategy_options {
  # Configure direct properties listed below.
}
```

## Direct properties

- [id_ranges](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/fs_group_strategy_options/id_ranges/): complete subsection reference.

<a id="schema-psp_spec--fs_group_strategy_options--rule"></a>

### rule property

Type: `"string"`. Optional.

Rule indicated how the FS group ID range is used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

- [psp_spec.fs_group_strategy_options.id_ranges](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/fs_group_strategy_options/id_ranges/)
- [psp_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/)
- [xcsh_k8s_pod_security_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/)
