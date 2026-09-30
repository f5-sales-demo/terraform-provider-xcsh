---
page_title: "psp_spec.fs_group_strategy_options.id_ranges"
subcategory: ""
description: "psp_spec.fs_group_strategy_options.id_ranges for xcsh_k8s_pod_security_policy."
xcsh_docs: {"aliases": [], "body_bytes": 4101, "body_sha256": "sha256:085a1a2becfbf843c672f2437908a1d7b0a031705c53cc07bcc31e326949860b", "canonical_id": "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:fs_group_strategy_options:id_ranges", "child_ids": [], "collection_id": "xcsh-docs:resources:k8s_pod_security_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:fs_group_strategy_options:id_ranges", "parent_id": "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:fs_group_strategy_options", "path": "docs/guides/resources--k8s_pod_security_policy--properties--psp_spec--fs_group_strategy_options--id_ranges.md", "provider_name": "k8s_pod_security_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["psp_spec", "fs_group_strategy_options", "id_ranges"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_pod_security_policy/properties/psp_spec/fs_group_strategy_options/id_ranges/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "psp_spec.fs_group_strategy_options.id_ranges for xcsh_k8s_pod_security_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_pod_security_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# psp_spec.fs_group_strategy_options.id_ranges

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md)
- [Property reference](resources--k8s_pod_security_policy--reference.md)
- [psp_spec](resources--k8s_pod_security_policy--properties--psp_spec.md)
- [psp_spec.fs_group_strategy_options](resources--k8s_pod_security_policy--properties--psp_spec--fs_group_strategy_options.md)
- psp_spec.fs_group_strategy_options.id_ranges

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

ID Ranges. List of range of ID(s)

Upstream description:

List of range of ID(s)

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("max_id",
    "min_id")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
id_ranges {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-psp_spec--fs_group_strategy_options--id_ranges--max_id"></a>

### max_id property

Type: `"number"`. Optional.

Ending ID. Ending(maximum) ID for for ID range.

Upstream description:

Ending(maximum) ID for for ID range.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="schema-psp_spec--fs_group_strategy_options--id_ranges--min_id"></a>

### min_id property

Type: `"number"`. Optional.

Starting ID. Starting(minimum) ID for for ID range.

Upstream description:

Starting(minimum) ID for for ID range.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

## Next pages

- [psp_spec.fs_group_strategy_options](resources--k8s_pod_security_policy--properties--psp_spec--fs_group_strategy_options.md)
- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md)
