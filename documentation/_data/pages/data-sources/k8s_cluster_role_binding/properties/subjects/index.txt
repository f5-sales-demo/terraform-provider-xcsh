---
page_title: "subjects"
subcategory: ""
description: "subjects for xcsh_k8s_cluster_role_binding."
xcsh_docs: {"aliases": [], "body_bytes": 4327, "body_sha256": "sha256:a0e4ea07abac3674d64a9c8cd207292533dbeb3c2c18a4b91beafa4d36089591", "child_ids": ["xcsh-docs:data-sources:k8s_cluster_role_binding:properties:subjects:service_account"], "collection_id": "xcsh-docs:data-sources:k8s_cluster_role_binding:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:k8s_cluster_role_binding:properties:subjects", "parent_id": "xcsh-docs:data-sources:k8s_cluster_role_binding:reference", "path": "documentation/data-sources/k8s_cluster_role_binding/properties/subjects/index.md", "provider_name": "k8s_cluster_role_binding", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["subjects"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/k8s_cluster_role_binding/properties/subjects/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "subjects for xcsh_k8s_cluster_role_binding.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_cluster_role_bindingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# subjects

Breadcrumbs:

- [xcsh_k8s_cluster_role_binding](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster_role_binding/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster_role_binding/properties/)
- subjects

<a id="section"></a>

Type: `"list"`. Computed.

List of subjects (user, group or service account) to which this role is bound.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

<a id="schema-subjects--group"></a>

### group property

Type: `"string"`. Computed.

Exclusive with \[service\_account user\] Group ID of the user group.

Upstream description:

Exclusive with \[service\_account user\] Group ID of the user group.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [service_account](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster_role_binding/properties/subjects/service_account/): complete subsection reference.

<a id="schema-subjects--user"></a>

### user property

Type: `"string"`. Computed.

Exclusive with \[group service\_account\] User ID of the user.

Upstream description:

Exclusive with \[group service\_account\] User ID of the user.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

## Next pages

- [subjects.service_account](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster_role_binding/properties/subjects/service_account/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster_role_binding/properties/)
- [xcsh_k8s_cluster_role_binding](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster_role_binding/)
