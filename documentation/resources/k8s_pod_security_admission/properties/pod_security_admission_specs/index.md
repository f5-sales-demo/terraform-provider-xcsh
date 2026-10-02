---
page_title: "pod_security_admission_specs"
subcategory: ""
description: "Uniform Resource Identifier"
xcsh_docs: {"aliases": ["pod security admission specs"], "body_bytes": 4789, "body_sha256": "sha256:d9a7fb07d2631469e364df49118260cee047164cd8b298c58a9acb3180a6b6cb", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:k8s_pod_security_admission:properties:pod_security_admission_specs:audit", "xcsh-docs:resources:k8s_pod_security_admission:properties:pod_security_admission_specs:baseline", "xcsh-docs:resources:k8s_pod_security_admission:properties:pod_security_admission_specs:enforce", "xcsh-docs:resources:k8s_pod_security_admission:properties:pod_security_admission_specs:privileged", "xcsh-docs:resources:k8s_pod_security_admission:properties:pod_security_admission_specs:restricted", "xcsh-docs:resources:k8s_pod_security_admission:properties:pod_security_admission_specs:warn"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:k8s_pod_security_admission:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_pod_security_admission:properties:pod_security_admission_specs", "parent_id": "xcsh-docs:resources:k8s_pod_security_admission:reference", "path": "documentation/resources/k8s_pod_security_admission/properties/pod_security_admission_specs/index.md", "product": "distributed-cloud", "provider_name": "k8s_pod_security_admission", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1000030230112200-1231331301101110-1303212112010002-0022032303300133-0310013011110210-1012020000003201-2013031303111200-2003222011220032", "registry_path": "docs/guides/resources--k8s_pod_security_admission--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["pod_security_admission_specs"], "schema_version": 1, "sections": [{"aliases": ["audit"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:k8s_pod_security_admission:properties:pod_security_admission_specs:audit", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["pod_security_admission_specs", "audit"], "syntax": "attribute", "type": "object"}, {"aliases": ["baseline"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:k8s_pod_security_admission:properties:pod_security_admission_specs:baseline", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["pod_security_admission_specs", "baseline"], "syntax": "attribute", "type": "object"}, {"aliases": ["enforce"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:k8s_pod_security_admission:properties:pod_security_admission_specs:enforce", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["pod_security_admission_specs", "enforce"], "syntax": "attribute", "type": "object"}, {"aliases": ["privileged"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:k8s_pod_security_admission:properties:pod_security_admission_specs:privileged", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["pod_security_admission_specs", "privileged"], "syntax": "attribute", "type": "object"}, {"aliases": ["restricted"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:k8s_pod_security_admission:properties:pod_security_admission_specs:restricted", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["pod_security_admission_specs", "restricted"], "syntax": "attribute", "type": "object"}, {"aliases": ["warn"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:k8s_pod_security_admission:properties:pod_security_admission_specs:warn", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["pod_security_admission_specs", "warn"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_pod_security_admission/properties/pod_security_admission_specs/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Uniform Resource Identifier", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_pod_security_admissionCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# pod_security_admission_specs

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_admission/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_admission/properties/)
- pod_security_admission_specs

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

K8s Pod Security Admission. Uniform Resource Identifier

Upstream description:

Uniform Resource Identifier

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("audit",
    "enforce"),
  validators.ConflictingListObjectAttributes("audit",
    "warn"),
  validators.ConflictingListObjectAttributes("baseline",
    "privileged"),
  validators.ConflictingListObjectAttributes("baseline",
    "restricted"),
  validators.ConflictingListObjectAttributes("enforce",
    "warn"),
  validators.ConflictingListObjectAttributes("privileged",
    "restricted")}
```

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

Terraform syntax:

```terraform
pod_security_admission_specs {
  # Configure direct properties listed below.
}
```

## Direct properties

- [audit](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_admission/properties/pod_security_admission_specs/audit/): complete subsection reference.

- [baseline](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_admission/properties/pod_security_admission_specs/baseline/): complete subsection reference.

- [enforce](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_admission/properties/pod_security_admission_specs/enforce/): complete subsection reference.

- [privileged](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_admission/properties/pod_security_admission_specs/privileged/): complete subsection reference.

- [restricted](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_admission/properties/pod_security_admission_specs/restricted/): complete subsection reference.

- [warn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_admission/properties/pod_security_admission_specs/warn/): complete subsection reference.

## Next pages

- [pod_security_admission_specs.audit](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_admission/properties/pod_security_admission_specs/audit/)
- [pod_security_admission_specs.baseline](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_admission/properties/pod_security_admission_specs/baseline/)
- [pod_security_admission_specs.enforce](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_admission/properties/pod_security_admission_specs/enforce/)
- [pod_security_admission_specs.privileged](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_admission/properties/pod_security_admission_specs/privileged/)
- [pod_security_admission_specs.restricted](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_admission/properties/pod_security_admission_specs/restricted/)
- [pod_security_admission_specs.warn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_admission/properties/pod_security_admission_specs/warn/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_admission/properties/)
- [xcsh_k8s_pod_security_admission](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_admission/)
