---
page_title: "psp_spec.run_as_user.id_ranges"
subcategory: ""
description: "List of range of ID(s)"
xcsh_docs: {"aliases": ["psp spec run as user id ranges"], "body_bytes": 3861, "body_sha256": "sha256:383b319689aebbef6f69a50e513f5dd68e7f2bb5f31b88a3bfb1ac38a3d46f38", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:k8s_pod_security_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec:run_as_user:id_ranges", "parent_id": "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec:run_as_user", "path": "documentation/data-sources/k8s_pod_security_policy/properties/psp_spec/run_as_user/id_ranges/index.md", "product": "distributed-cloud", "provider_name": "k8s_pod_security_policy", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-3202000213223201-2101132131203101-0011330011003113-0001121103011211-0021022132131100-1203300131121230-1233013322010010-1012030021102103", "registry_path": "docs/guides/data-sources--k8s_pod_security_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["psp_spec", "run_as_user", "id_ranges"], "schema_version": 1, "sections": [{"aliases": ["psp spec run as user id ranges max id"], "anchor": "schema-psp_spec--run_as_user--id_ranges--max_id", "description": "Ending(maximum) ID for for ID range.", "document_id": "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec:run_as_user:id_ranges", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["psp_spec", "run_as_user", "id_ranges", "max_id"], "syntax": "attribute", "type": "number"}, {"aliases": ["psp spec run as user id ranges min id"], "anchor": "schema-psp_spec--run_as_user--id_ranges--min_id", "description": "Starting(minimum) ID for for ID range.", "document_id": "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec:run_as_user:id_ranges", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["psp_spec", "run_as_user", "id_ranges", "min_id"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/k8s_pod_security_policy/properties/psp_spec/run_as_user/id_ranges/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "List of range of ID(s)", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["k8s_pod_security_policyCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# psp_spec.run_as_user.id_ranges

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_policy/properties/)
- [psp_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_policy/properties/psp_spec/)
- [psp_spec.run_as_user](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_policy/properties/psp_spec/run_as_user/)
- psp_spec.run_as_user.id_ranges

<a id="section"></a>

Type: `"list"`. Computed.

ID Ranges. List of range of ID(s)

Upstream description:

List of range of ID(s)

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

## Direct properties

<a id="schema-psp_spec--run_as_user--id_ranges--max_id"></a>

### max_id property

Type: `"number"`. Computed.

Ending ID. Ending(maximum) ID for for ID range.

Upstream description:

Ending(maximum) ID for for ID range.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="schema-psp_spec--run_as_user--id_ranges--min_id"></a>

### min_id property

Type: `"number"`. Computed.

Starting ID. Starting(minimum) ID for for ID range.

Upstream description:

Starting(minimum) ID for for ID range.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [psp_spec.run_as_user](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_policy/properties/psp_spec/run_as_user/)
- [xcsh_k8s_pod_security_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_policy/)
