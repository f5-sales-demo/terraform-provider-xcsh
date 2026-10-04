---
page_title: "psp_spec.supplemental_groups.id_ranges"
subcategory: ""
description: "List of range of ID(s)"
xcsh_docs: {"aliases": ["psp spec supplemental groups id ranges"], "body_bytes": 4458, "body_sha256": "sha256:3ce23c4025400deb66e7aeee9ffa006a8cb8c2c071e8fb1570162f43a92caec1", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:k8s_pod_security_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:supplemental_groups:id_ranges", "parent_id": "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:supplemental_groups", "path": "documentation/resources/k8s_pod_security_policy/properties/psp_spec/supplemental_groups/id_ranges/index.md", "product": "distributed-cloud", "provider_name": "k8s_pod_security_policy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3201111331020301-0321203222222222-3030211031010121-0131023303302002-1003333330320321-1131220111023023-2321122030120121-0033230301231332", "registry_path": "docs/guides/resources--k8s_pod_security_policy--reference--group-001.md", "relationships": [{"anchor": "schema-psp_spec--supplemental_groups--id_ranges--max_id", "enforcement": "provider-schema", "group": "psp_spec.supplemental_groups.id_ranges:RequiredListObjectAttributes:max_id,min_id", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:supplemental_groups:id_ranges", "type": "requires"}, {"anchor": "schema-psp_spec--supplemental_groups--id_ranges--min_id", "enforcement": "provider-schema", "group": "psp_spec.supplemental_groups.id_ranges:RequiredListObjectAttributes:max_id,min_id", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:supplemental_groups:id_ranges", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["psp_spec", "supplemental_groups", "id_ranges"], "schema_version": 1, "sections": [{"aliases": ["psp spec supplemental groups id ranges max id"], "anchor": "schema-psp_spec--supplemental_groups--id_ranges--max_id", "description": "Ending(maximum) ID for for ID range.", "document_id": "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:supplemental_groups:id_ranges", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["psp_spec", "supplemental_groups", "id_ranges", "max_id"], "syntax": "attribute", "type": "number"}, {"aliases": ["psp spec supplemental groups id ranges min id"], "anchor": "schema-psp_spec--supplemental_groups--id_ranges--min_id", "description": "Starting(minimum) ID for for ID range.", "document_id": "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:supplemental_groups:id_ranges", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["psp_spec", "supplemental_groups", "id_ranges", "min_id"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_pod_security_policy/properties/psp_spec/supplemental_groups/id_ranges/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "List of range of ID(s)", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["k8s_pod_security_policyCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# psp_spec.supplemental_groups.id_ranges

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/)
- [psp_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/)
- [psp_spec.supplemental_groups](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/supplemental_groups/)
- psp_spec.supplemental_groups.id_ranges

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

Terraform syntax:

```terraform
id_ranges {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-psp_spec--supplemental_groups--id_ranges--max_id"></a>

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

<a id="schema-psp_spec--supplemental_groups--id_ranges--min_id"></a>

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

- [psp_spec.supplemental_groups](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/supplemental_groups/)
- [xcsh_k8s_pod_security_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/)
