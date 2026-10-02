---
page_title: "psp_spec.supplemental_groups"
subcategory: ""
description: "ID ranges and rules."
xcsh_docs: {"aliases": ["psp spec supplemental groups"], "body_bytes": 2933, "body_sha256": "sha256:e8b7f86758a74e64617b27a30d04b655a87166fda6d3db12c682e51bbcc04a7a", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:supplemental_groups:id_ranges"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:k8s_pod_security_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:supplemental_groups", "parent_id": "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec", "path": "documentation/resources/k8s_pod_security_policy/properties/psp_spec/supplemental_groups/index.md", "product": "distributed-cloud", "provider_name": "k8s_pod_security_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3212320333313232-2233312031023110-2223231010003002-2023103303022112-2302230232033013-1011011211221023-1010000011313231-2102321103331011", "registry_path": "docs/guides/resources--k8s_pod_security_policy--reference--group-001.md", "relationships": [{"anchor": "schema-psp_spec--supplemental_groups--rule", "enforcement": "provider-schema", "group": "psp_spec.supplemental_groups:RequiredObjectAttributes:rule", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:supplemental_groups", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["psp_spec", "supplemental_groups"], "schema_version": 1, "sections": [{"aliases": ["id ranges"], "anchor": "section", "description": "List of range of ID(s)", "document_id": "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:supplemental_groups:id_ranges", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-psp_spec--supplemental_groups--id_ranges--max_id", "enforcement": "provider-schema", "group": "psp_spec.supplemental_groups.id_ranges:RequiredListObjectAttributes:max_id,min_id", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:supplemental_groups:id_ranges", "type": "requires"}, {"anchor": "schema-psp_spec--supplemental_groups--id_ranges--min_id", "enforcement": "provider-schema", "group": "psp_spec.supplemental_groups.id_ranges:RequiredListObjectAttributes:max_id,min_id", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:supplemental_groups:id_ranges", "type": "requires"}], "schema_path": ["psp_spec", "supplemental_groups", "id_ranges"], "syntax": "block", "type": "object"}, {"aliases": ["rule"], "anchor": "schema-psp_spec--supplemental_groups--rule", "description": "Rule indicated how the FS group ID range is used.", "document_id": "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:supplemental_groups", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["psp_spec", "supplemental_groups", "rule"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_pod_security_policy/properties/psp_spec/supplemental_groups/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "ID ranges and rules.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["k8s_pod_security_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

- [psp_spec.supplemental_groups.id_ranges](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/supplemental_groups/id_ranges/)
- [psp_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/)
- [xcsh_k8s_pod_security_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/)
