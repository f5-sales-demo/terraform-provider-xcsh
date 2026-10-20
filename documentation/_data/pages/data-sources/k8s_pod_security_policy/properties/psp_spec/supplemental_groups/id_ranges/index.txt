---
page_title: "psp_spec.supplemental_groups.id_ranges"
subcategory: ""
description: "List of range of ID(s)"
xcsh_docs: {"aliases": ["psp spec supplemental groups id ranges"], "body_bytes": 3442, "body_sha256": "sha256:b015b86c05c4b3d17c1abe70f9495318453c8037b8c71f44fe655d7736c6c852", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:k8s_pod_security_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec:supplemental_groups:id_ranges", "parent_id": "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec:supplemental_groups", "path": "documentation/data-sources/k8s_pod_security_policy/properties/psp_spec/supplemental_groups/id_ranges/index.md", "product": "distributed-cloud", "provider_name": "k8s_pod_security_policy", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-3022130322331311-0030123221213010-3003100132303010-2321033012030202-0122222323213221-1300130300013033-2212220331301110-2330112301333321", "registry_path": "docs/guides/data-sources--k8s_pod_security_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["psp_spec", "supplemental_groups", "id_ranges"], "schema_version": 1, "sections": [{"aliases": ["psp spec supplemental groups id ranges max id"], "anchor": "schema-psp_spec--supplemental_groups--id_ranges--max_id", "description": "Ending(maximum) ID for for ID range.", "document_id": "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec:supplemental_groups:id_ranges", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["psp_spec", "supplemental_groups", "id_ranges", "max_id"], "syntax": "attribute", "type": "number"}, {"aliases": ["psp spec supplemental groups id ranges min id"], "anchor": "schema-psp_spec--supplemental_groups--id_ranges--min_id", "description": "Starting(minimum) ID for for ID range.", "document_id": "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec:supplemental_groups:id_ranges", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["psp_spec", "supplemental_groups", "id_ranges", "min_id"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/k8s_pod_security_policy/properties/psp_spec/supplemental_groups/id_ranges/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "List of range of ID(s)", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["k8s_pod_security_policyCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# psp_spec.supplemental_groups.id_ranges

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_policy/properties/)
- [psp_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_policy/properties/psp_spec/)
- [psp_spec.supplemental_groups](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_policy/properties/psp_spec/supplemental_groups/)
- psp_spec.supplemental_groups.id_ranges

<a id="section"></a>

Type: `"list"`. Computed.

ID Ranges. List of range of ID(s)

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
      "validatedAt": "2026-10-20T00:10:01+00:00"
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

<a id="schema-psp_spec--supplemental_groups--id_ranges--max_id"></a>

### max_id property

Type: `"number"`. Computed.

Ending ID. Ending(maximum) ID for for ID range.

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
      "validatedAt": "2026-10-20T00:10:01+00:00"
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

<a id="schema-psp_spec--supplemental_groups--id_ranges--min_id"></a>

### min_id property

Type: `"number"`. Computed.

Starting ID. Starting(minimum) ID for for ID range.

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
      "validatedAt": "2026-10-20T00:10:01+00:00"
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
