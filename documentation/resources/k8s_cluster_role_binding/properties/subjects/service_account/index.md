---
page_title: "subjects.service_account"
subcategory: ""
description: "ServiceAccountType."
xcsh_docs: {"aliases": ["subjects service account"], "body_bytes": 4499, "body_sha256": "sha256:2904816fb7ab9d2675b433fae1f27333a7247ed1120377a3052f2d181d0caca4", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:k8s_cluster_role_binding:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_cluster_role_binding:properties:subjects:service_account", "parent_id": "xcsh-docs:resources:k8s_cluster_role_binding:properties:subjects", "path": "documentation/resources/k8s_cluster_role_binding/properties/subjects/service_account/index.md", "product": "distributed-cloud", "provider_name": "k8s_cluster_role_binding", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-1133023320132221-1202231111122312-1013102311103130-2131231013113132-1001313031131130-3201231112123021-1002133103000223-1000120230001210", "registry_path": "docs/guides/resources--k8s_cluster_role_binding--reference--group-001.md", "relationships": [{"anchor": "schema-subjects--service_account--name", "enforcement": "provider-schema", "group": "subjects.service_account:RequiredObjectAttributes:name,namespace", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:k8s_cluster_role_binding:properties:subjects:service_account", "type": "requires"}, {"anchor": "schema-subjects--service_account--namespace", "enforcement": "provider-schema", "group": "subjects.service_account:RequiredObjectAttributes:name,namespace", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:k8s_cluster_role_binding:properties:subjects:service_account", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["subjects", "service_account"], "schema_version": 1, "sections": [{"aliases": ["subjects service account name"], "anchor": "schema-subjects--service_account--name", "description": "Name of the service account.", "document_id": "xcsh-docs:resources:k8s_cluster_role_binding:properties:subjects:service_account", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["subjects", "service_account", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["subjects service account namespace"], "anchor": "schema-subjects--service_account--namespace", "description": "Namespace of the service account.", "document_id": "xcsh-docs:resources:k8s_cluster_role_binding:properties:subjects:service_account", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["subjects", "service_account", "namespace"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_cluster_role_binding/properties/subjects/service_account/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "ServiceAccountType.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["k8s_cluster_role_bindingCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# subjects.service_account

Breadcrumbs:

- [xcsh_k8s_cluster_role_binding](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role_binding/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role_binding/properties/)
- [subjects](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role_binding/properties/subjects/)
- subjects.service_account

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

ServiceAccountType.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("name",
    "namespace")}
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
service_account {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-subjects--service_account--name"></a>

### name property

Type: `"string"`. Optional.

Name. Name of the service account.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-20T00:10:01+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="schema-subjects--service_account--namespace"></a>

### namespace property

Type: `"string"`. Optional, Computed.

Namespace. Namespace of the service account.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-20T00:10:01+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```
