---
page_title: "psp_spec.allowed_capabilities"
subcategory: ""
description: "List of capabilities that docker container has."
xcsh_docs: {"aliases": ["psp spec allowed capabilities"], "body_bytes": 2501, "body_sha256": "sha256:0d24279e9fd479723fbf29886baf3181cc8dd3ebb8b114de2301ca3f4ba29496", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:k8s_pod_security_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:allowed_capabilities", "parent_id": "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec", "path": "documentation/resources/k8s_pod_security_policy/properties/psp_spec/allowed_capabilities/index.md", "product": "distributed-cloud", "provider_name": "k8s_pod_security_policy", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-2010232122200033-3131102120101000-2230012232013333-0030312320331011-0222312113203301-2111131111321223-3330321021211233-2330333310103110", "registry_path": "docs/guides/resources--k8s_pod_security_policy--reference--group-001.md", "relationships": [{"anchor": "schema-psp_spec--allowed_capabilities--capabilities", "enforcement": "provider-schema", "group": "psp_spec.allowed_capabilities:RequiredObjectAttributes:capabilities", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:allowed_capabilities", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["psp_spec", "allowed_capabilities"], "schema_version": 1, "sections": [{"aliases": ["psp spec allowed capabilities capabilities"], "anchor": "schema-psp_spec--allowed_capabilities--capabilities", "description": "List of capabilities that docker container has.", "document_id": "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:allowed_capabilities", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["psp_spec", "allowed_capabilities", "capabilities"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_pod_security_policy/properties/psp_spec/allowed_capabilities/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "List of capabilities that docker container has.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["k8s_pod_security_policyCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# psp_spec.allowed_capabilities

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/)
- [psp_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/)
- psp_spec.allowed_capabilities

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

List of capabilities that docker container has.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("capabilities")}
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
allowed_capabilities {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-psp_spec--allowed_capabilities--capabilities"></a>

### capabilities property

Type: `["list", "string"]`. Optional.

List of capabilities that docker container has.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```
