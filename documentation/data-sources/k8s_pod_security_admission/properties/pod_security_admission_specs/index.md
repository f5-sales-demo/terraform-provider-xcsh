---
page_title: "pod_security_admission_specs"
subcategory: ""
description: "Uniform Resource Identifier"
xcsh_docs: {"aliases": ["pod security admission specs"], "body_bytes": 4170, "body_sha256": "sha256:3c3238498ea1a941d2ef7fc11e1bb80ca99528798d78c1e7c8faa55adea89212", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:k8s_pod_security_admission:properties:pod_security_admission_specs:audit", "xcsh-docs:data-sources:k8s_pod_security_admission:properties:pod_security_admission_specs:baseline", "xcsh-docs:data-sources:k8s_pod_security_admission:properties:pod_security_admission_specs:enforce", "xcsh-docs:data-sources:k8s_pod_security_admission:properties:pod_security_admission_specs:privileged", "xcsh-docs:data-sources:k8s_pod_security_admission:properties:pod_security_admission_specs:restricted", "xcsh-docs:data-sources:k8s_pod_security_admission:properties:pod_security_admission_specs:warn"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:k8s_pod_security_admission:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:k8s_pod_security_admission:properties:pod_security_admission_specs", "parent_id": "xcsh-docs:data-sources:k8s_pod_security_admission:reference", "path": "documentation/data-sources/k8s_pod_security_admission/properties/pod_security_admission_specs/index.md", "product": "distributed-cloud", "provider_name": "k8s_pod_security_admission", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1211112021002321-2221300323023323-3200102331112201-0013023312200103-1031010012002313-3231212323301230-1011020323221331-3120332333030331", "registry_path": "docs/guides/data-sources--k8s_pod_security_admission--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["pod_security_admission_specs"], "schema_version": 1, "sections": [{"aliases": ["audit"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:k8s_pod_security_admission:properties:pod_security_admission_specs:audit", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["pod_security_admission_specs", "audit"], "syntax": "attribute", "type": "object"}, {"aliases": ["baseline"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:k8s_pod_security_admission:properties:pod_security_admission_specs:baseline", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["pod_security_admission_specs", "baseline"], "syntax": "attribute", "type": "object"}, {"aliases": ["enforce"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:k8s_pod_security_admission:properties:pod_security_admission_specs:enforce", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["pod_security_admission_specs", "enforce"], "syntax": "attribute", "type": "object"}, {"aliases": ["privileged"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:k8s_pod_security_admission:properties:pod_security_admission_specs:privileged", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["pod_security_admission_specs", "privileged"], "syntax": "attribute", "type": "object"}, {"aliases": ["restricted"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:k8s_pod_security_admission:properties:pod_security_admission_specs:restricted", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["pod_security_admission_specs", "restricted"], "syntax": "attribute", "type": "object"}, {"aliases": ["warn"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:k8s_pod_security_admission:properties:pod_security_admission_specs:warn", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["pod_security_admission_specs", "warn"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/k8s_pod_security_admission/properties/pod_security_admission_specs/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Uniform Resource Identifier", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["k8s_pod_security_admissionCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

Upstream description:

Uniform Resource Identifier

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

## Next pages

- [pod_security_admission_specs.audit](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_admission/properties/pod_security_admission_specs/audit/)
- [pod_security_admission_specs.baseline](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_admission/properties/pod_security_admission_specs/baseline/)
- [pod_security_admission_specs.enforce](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_admission/properties/pod_security_admission_specs/enforce/)
- [pod_security_admission_specs.privileged](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_admission/properties/pod_security_admission_specs/privileged/)
- [pod_security_admission_specs.restricted](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_admission/properties/pod_security_admission_specs/restricted/)
- [pod_security_admission_specs.warn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_admission/properties/pod_security_admission_specs/warn/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_admission/properties/)
- [xcsh_k8s_pod_security_admission](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_admission/)
