---
page_title: "pod_security_admission_specs"
subcategory: ""
description: "Uniform Resource Identifier"
xcsh_docs: {"aliases": ["pod security admission specs"], "body_bytes": 2710, "body_sha256": "sha256:3799ae360f306d1db76c42082099ba77ae83a443f09c495f7afa165d41fef454", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:k8s_pod_security_admission:properties:pod_security_admission_specs:audit", "xcsh-docs:data-sources:k8s_pod_security_admission:properties:pod_security_admission_specs:baseline", "xcsh-docs:data-sources:k8s_pod_security_admission:properties:pod_security_admission_specs:enforce", "xcsh-docs:data-sources:k8s_pod_security_admission:properties:pod_security_admission_specs:privileged", "xcsh-docs:data-sources:k8s_pod_security_admission:properties:pod_security_admission_specs:restricted", "xcsh-docs:data-sources:k8s_pod_security_admission:properties:pod_security_admission_specs:warn"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:k8s_pod_security_admission:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:k8s_pod_security_admission:properties:pod_security_admission_specs", "parent_id": "xcsh-docs:data-sources:k8s_pod_security_admission:reference", "path": "documentation/data-sources/k8s_pod_security_admission/properties/pod_security_admission_specs/index.md", "product": "distributed-cloud", "provider_name": "k8s_pod_security_admission", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-1211112021002321-2221300323023323-3200102331112201-0013023312200103-1031010012002313-3231212323301230-1011020323221331-3120332333030331", "registry_path": "docs/guides/data-sources--k8s_pod_security_admission--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["pod_security_admission_specs"], "schema_version": 1, "sections": [{"aliases": ["pod security admission specs audit"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:k8s_pod_security_admission:properties:pod_security_admission_specs:audit", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["pod_security_admission_specs", "audit"], "syntax": "attribute", "type": "object"}, {"aliases": ["pod security admission specs baseline"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:k8s_pod_security_admission:properties:pod_security_admission_specs:baseline", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["pod_security_admission_specs", "baseline"], "syntax": "attribute", "type": "object"}, {"aliases": ["pod security admission specs enforce"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:k8s_pod_security_admission:properties:pod_security_admission_specs:enforce", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["pod_security_admission_specs", "enforce"], "syntax": "attribute", "type": "object"}, {"aliases": ["pod security admission specs privileged"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:k8s_pod_security_admission:properties:pod_security_admission_specs:privileged", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["pod_security_admission_specs", "privileged"], "syntax": "attribute", "type": "object"}, {"aliases": ["pod security admission specs restricted"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:k8s_pod_security_admission:properties:pod_security_admission_specs:restricted", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["pod_security_admission_specs", "restricted"], "syntax": "attribute", "type": "object"}, {"aliases": ["pod security admission specs warn"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:k8s_pod_security_admission:properties:pod_security_admission_specs:warn", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["pod_security_admission_specs", "warn"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/k8s_pod_security_admission/properties/pod_security_admission_specs/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Uniform Resource Identifier", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["k8s_pod_security_admissionCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# pod_security_admission_specs

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_admission/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_admission/properties/)
- pod_security_admission_specs

<a id="section"></a>

Type: `"list"`. Computed.

K8s Pod Security Admission. Uniform Resource Identifier

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

- [audit](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_admission/properties/pod_security_admission_specs/audit/): complete subsection reference.

- [baseline](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_admission/properties/pod_security_admission_specs/baseline/): complete subsection reference.

- [enforce](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_admission/properties/pod_security_admission_specs/enforce/): complete subsection reference.

- [privileged](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_admission/properties/pod_security_admission_specs/privileged/): complete subsection reference.

- [restricted](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_admission/properties/pod_security_admission_specs/restricted/): complete subsection reference.

- [warn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_admission/properties/pod_security_admission_specs/warn/): complete subsection reference.
