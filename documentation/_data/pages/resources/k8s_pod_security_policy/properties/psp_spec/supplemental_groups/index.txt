---
page_title: "psp_spec.supplemental_groups"
subcategory: ""
description: "ID ranges and rules."
xcsh_docs: {"aliases": ["psp spec supplemental groups"], "body_bytes": 2496, "body_sha256": "sha256:50be6867be14c22a0b1717ae9db9fcbb7aa3fc8dbc41f44db4472a463ee6581c", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:supplemental_groups:id_ranges"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:k8s_pod_security_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:supplemental_groups", "parent_id": "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec", "path": "documentation/resources/k8s_pod_security_policy/properties/psp_spec/supplemental_groups/index.md", "product": "distributed-cloud", "provider_name": "k8s_pod_security_policy", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-3212320333313232-2233312031023110-2223231010003002-2023103303022112-2302230232033013-1011011211221023-1010000011313231-2102321103331011", "registry_path": "docs/guides/resources--k8s_pod_security_policy--reference--group-001.md", "relationships": [{"anchor": "schema-psp_spec--supplemental_groups--rule", "enforcement": "provider-schema", "group": "psp_spec.supplemental_groups:RequiredObjectAttributes:rule", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:supplemental_groups", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["psp_spec", "supplemental_groups"], "schema_version": 1, "sections": [{"aliases": ["psp spec supplemental groups id ranges"], "anchor": "section", "description": "List of range of ID(s)", "document_id": "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:supplemental_groups:id_ranges", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-psp_spec--supplemental_groups--id_ranges--max_id", "enforcement": "provider-schema", "group": "psp_spec.supplemental_groups.id_ranges:RequiredListObjectAttributes:max_id,min_id", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:supplemental_groups:id_ranges", "type": "requires"}, {"anchor": "schema-psp_spec--supplemental_groups--id_ranges--min_id", "enforcement": "provider-schema", "group": "psp_spec.supplemental_groups.id_ranges:RequiredListObjectAttributes:max_id,min_id", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:supplemental_groups:id_ranges", "type": "requires"}], "schema_path": ["psp_spec", "supplemental_groups", "id_ranges"], "syntax": "block", "type": "object"}, {"aliases": ["psp spec supplemental groups rule"], "anchor": "schema-psp_spec--supplemental_groups--rule", "description": "Rule indicated how the FS group ID range is used.", "document_id": "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:supplemental_groups", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["psp_spec", "supplemental_groups", "rule"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_pod_security_policy/properties/psp_spec/supplemental_groups/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "ID ranges and rules.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["k8s_pod_security_policyCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# psp_spec.supplemental_groups

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/)
- [psp_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/)
- psp_spec.supplemental_groups

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

ID(User,Group,FSGroup) Strategy. ID ranges and rules.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
supplemental_groups {
  # Configure direct properties listed below.
}
```

## Direct properties

- [id_ranges](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/supplemental_groups/id_ranges/): complete subsection reference.

<a id="schema-psp_spec--supplemental_groups--rule"></a>

### rule property

Type: `"string"`. Optional.

Rule indicated how the FS group ID range is used.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
