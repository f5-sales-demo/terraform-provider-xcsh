---
page_title: "subjects"
subcategory: ""
description: "List of subjects (user, group or service account) to which this role is bound."
xcsh_docs: {"aliases": ["subjects"], "body_bytes": 5024, "body_sha256": "sha256:2007286117e9977da78aea342d4dba833a08229d38c896fde5b30f35b10ab322", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:k8s_cluster_role_binding:properties:subjects:service_account"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:k8s_cluster_role_binding:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_cluster_role_binding:properties:subjects", "parent_id": "xcsh-docs:resources:k8s_cluster_role_binding:reference", "path": "documentation/resources/k8s_cluster_role_binding/properties/subjects/index.md", "product": "distributed-cloud", "provider_name": "k8s_cluster_role_binding", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3202231323013321-2032223021010103-2021303133012001-3313313131013202-1232010101210021-2302132012221100-0313212310022231-0130230122331332", "registry_path": "docs/guides/resources--k8s_cluster_role_binding--reference--group-001.md", "relationships": [{"anchor": "schema-subjects--group", "enforcement": "provider-schema", "group": "subjects:ConflictingListObjectAttributes:group,service_account", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:k8s_cluster_role_binding:properties:subjects", "type": "conflicts"}, {"anchor": "schema-subjects--group", "enforcement": "provider-schema", "group": "subjects:ConflictingListObjectAttributes:group,user", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:k8s_cluster_role_binding:properties:subjects", "type": "conflicts"}, {"anchor": "schema-subjects--user", "enforcement": "provider-schema", "group": "subjects:ConflictingListObjectAttributes:group,user", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:k8s_cluster_role_binding:properties:subjects", "type": "conflicts"}, {"anchor": "schema-subjects--user", "enforcement": "provider-schema", "group": "subjects:ConflictingListObjectAttributes:service_account,user", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:k8s_cluster_role_binding:properties:subjects", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "subjects:ConflictingListObjectAttributes:group,service_account", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:k8s_cluster_role_binding:properties:subjects:service_account", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "subjects:ConflictingListObjectAttributes:service_account,user", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:k8s_cluster_role_binding:properties:subjects:service_account", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["subjects"], "schema_version": 1, "sections": [{"aliases": ["subjects group"], "anchor": "schema-subjects--group", "description": "Exclusive with Group ID of the user group.", "document_id": "xcsh-docs:resources:k8s_cluster_role_binding:properties:subjects", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["subjects", "group"], "syntax": "attribute", "type": "string"}, {"aliases": ["subjects service account"], "anchor": "section", "description": "ServiceAccountType.", "document_id": "xcsh-docs:resources:k8s_cluster_role_binding:properties:subjects:service_account", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-subjects--service_account--name", "enforcement": "provider-schema", "group": "subjects.service_account:RequiredObjectAttributes:name,namespace", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:k8s_cluster_role_binding:properties:subjects:service_account", "type": "requires"}, {"anchor": "schema-subjects--service_account--namespace", "enforcement": "provider-schema", "group": "subjects.service_account:RequiredObjectAttributes:name,namespace", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:k8s_cluster_role_binding:properties:subjects:service_account", "type": "requires"}], "schema_path": ["subjects", "service_account"], "syntax": "block", "type": "object"}, {"aliases": ["subjects user"], "anchor": "schema-subjects--user", "description": "Exclusive with User ID of the user.", "document_id": "xcsh-docs:resources:k8s_cluster_role_binding:properties:subjects", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["subjects", "user"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_cluster_role_binding/properties/subjects/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of subjects (user, group or service account) to which this role is bound.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["k8s_cluster_role_bindingCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# subjects

Breadcrumbs:

- [xcsh_k8s_cluster_role_binding](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role_binding/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role_binding/properties/)
- subjects

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

List of subjects (user, group or service account) to which this role is bound.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("group",
    "service_account"),
  validators.ConflictingListObjectAttributes("group",
    "user"),
  validators.ConflictingListObjectAttributes("service_account",
    "user")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
subjects {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-subjects--group"></a>

### group property

Type: `"string"`. Optional.

Exclusive with \[service\_account user\] Group ID of the user group.

Upstream description:

Exclusive with \[service\_account user\] Group ID of the user group.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [service_account](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role_binding/properties/subjects/service_account/): complete subsection reference.

<a id="schema-subjects--user"></a>

### user property

Type: `"string"`. Optional.

Exclusive with \[group service\_account\] User ID of the user.

Upstream description:

Exclusive with \[group service\_account\] User ID of the user.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

## Next pages

- [subjects.service_account](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role_binding/properties/subjects/service_account/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role_binding/properties/)
- [xcsh_k8s_cluster_role_binding](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role_binding/)
