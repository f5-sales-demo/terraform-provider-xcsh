---
page_title: "pod_security_admission_specs"
subcategory: ""
description: "Uniform Resource Identifier"
xcsh_docs: {"aliases": ["pod security admission specs"], "body_bytes": 3383, "body_sha256": "sha256:1569be7b175648ff47acb327426b189754a33c66b7e7b9aefc9abe5e71be0bf1", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:k8s_pod_security_admission:properties:pod_security_admission_specs:audit", "xcsh-docs:resources:k8s_pod_security_admission:properties:pod_security_admission_specs:baseline", "xcsh-docs:resources:k8s_pod_security_admission:properties:pod_security_admission_specs:enforce", "xcsh-docs:resources:k8s_pod_security_admission:properties:pod_security_admission_specs:privileged", "xcsh-docs:resources:k8s_pod_security_admission:properties:pod_security_admission_specs:restricted", "xcsh-docs:resources:k8s_pod_security_admission:properties:pod_security_admission_specs:warn"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:k8s_pod_security_admission:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_pod_security_admission:properties:pod_security_admission_specs", "parent_id": "xcsh-docs:resources:k8s_pod_security_admission:reference", "path": "documentation/resources/k8s_pod_security_admission/properties/pod_security_admission_specs/index.md", "product": "distributed-cloud", "provider_name": "k8s_pod_security_admission", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-1000030230112200-1231331301101110-1303212112010002-0022032303300133-0310013011110210-1012020000003201-2013031303111200-2003222011220032", "registry_path": "docs/guides/resources--k8s_pod_security_admission--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "pod_security_admission_specs:ConflictingListObjectAttributes:audit,enforce", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:k8s_pod_security_admission:properties:pod_security_admission_specs:audit", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "pod_security_admission_specs:ConflictingListObjectAttributes:audit,warn", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:k8s_pod_security_admission:properties:pod_security_admission_specs:audit", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "pod_security_admission_specs:ConflictingListObjectAttributes:baseline,privileged", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:k8s_pod_security_admission:properties:pod_security_admission_specs:baseline", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "pod_security_admission_specs:ConflictingListObjectAttributes:baseline,restricted", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:k8s_pod_security_admission:properties:pod_security_admission_specs:baseline", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "pod_security_admission_specs:ConflictingListObjectAttributes:audit,enforce", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:k8s_pod_security_admission:properties:pod_security_admission_specs:enforce", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "pod_security_admission_specs:ConflictingListObjectAttributes:enforce,warn", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:k8s_pod_security_admission:properties:pod_security_admission_specs:enforce", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "pod_security_admission_specs:ConflictingListObjectAttributes:baseline,privileged", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:k8s_pod_security_admission:properties:pod_security_admission_specs:privileged", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "pod_security_admission_specs:ConflictingListObjectAttributes:privileged,restricted", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:k8s_pod_security_admission:properties:pod_security_admission_specs:privileged", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "pod_security_admission_specs:ConflictingListObjectAttributes:baseline,restricted", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:k8s_pod_security_admission:properties:pod_security_admission_specs:restricted", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "pod_security_admission_specs:ConflictingListObjectAttributes:privileged,restricted", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:k8s_pod_security_admission:properties:pod_security_admission_specs:restricted", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "pod_security_admission_specs:ConflictingListObjectAttributes:audit,warn", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:k8s_pod_security_admission:properties:pod_security_admission_specs:warn", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "pod_security_admission_specs:ConflictingListObjectAttributes:enforce,warn", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:k8s_pod_security_admission:properties:pod_security_admission_specs:warn", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["pod_security_admission_specs"], "schema_version": 1, "sections": [{"aliases": ["pod security admission specs audit"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:k8s_pod_security_admission:properties:pod_security_admission_specs:audit", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["pod_security_admission_specs", "audit"], "syntax": "attribute", "type": "object"}, {"aliases": ["pod security admission specs baseline"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:k8s_pod_security_admission:properties:pod_security_admission_specs:baseline", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["pod_security_admission_specs", "baseline"], "syntax": "attribute", "type": "object"}, {"aliases": ["pod security admission specs enforce"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:k8s_pod_security_admission:properties:pod_security_admission_specs:enforce", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["pod_security_admission_specs", "enforce"], "syntax": "attribute", "type": "object"}, {"aliases": ["pod security admission specs privileged"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:k8s_pod_security_admission:properties:pod_security_admission_specs:privileged", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["pod_security_admission_specs", "privileged"], "syntax": "attribute", "type": "object"}, {"aliases": ["pod security admission specs restricted"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:k8s_pod_security_admission:properties:pod_security_admission_specs:restricted", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["pod_security_admission_specs", "restricted"], "syntax": "attribute", "type": "object"}, {"aliases": ["pod security admission specs warn"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:k8s_pod_security_admission:properties:pod_security_admission_specs:warn", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["pod_security_admission_specs", "warn"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_pod_security_admission/properties/pod_security_admission_specs/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Uniform Resource Identifier", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["k8s_pod_security_admissionCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
