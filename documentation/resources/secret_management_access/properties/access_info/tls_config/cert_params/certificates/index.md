---
page_title: "access_info.tls_config.cert_params.certificates"
subcategory: ""
description: "Client TLS Certificate required for mTLS authentication."
xcsh_docs: {"aliases": ["access info tls config cert params certificates", "cert", "certificate", "existing certificates", "tls certificates"], "body_bytes": 5800, "body_sha256": "sha256:8a34da3e4543575a18a2e48496a9e3d56f9b0879c6a69d40cbf109579c8f7ef4", "capabilities": ["identity"], "category": "identity", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:cert_params:certificates", "parent_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:cert_params", "path": "documentation/resources/secret_management_access/properties/access_info/tls_config/cert_params/certificates/index.md", "product": "distributed-cloud", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "resources", "registry_anchor": "canonical-3312332300331200-3331110220030031-3103112121301103-3231130233013113-2300320033021120-2331321033322311-1310230202031310-1211311013311133", "registry_path": "docs/guides/resources--secret_management_access--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["access_info", "tls_config", "cert_params", "certificates"], "schema_version": 1, "sections": [{"aliases": ["access info tls config cert params certificates kind"], "anchor": "schema-access_info--tls_config--cert_params--certificates--kind", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. \"route\")", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:cert_params:certificates", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "reference_identity": {"member": "kind", "scope_path": ["access_info", "tls_config", "cert_params", "certificates"], "source": "receipt-pinned-schema-identity", "upstream_message": "ves.io.schema.ObjectRefType", "version": 1}, "relationships": [], "schema_path": ["access_info", "tls_config", "cert_params", "certificates", "kind"], "syntax": "attribute", "type": "string"}, {"aliases": ["access info tls config cert params certificates name"], "anchor": "schema-access_info--tls_config--cert_params--certificates--name", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name.", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:cert_params:certificates", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "reference_identity": {"member": "name", "scope_path": ["access_info", "tls_config", "cert_params", "certificates"], "source": "receipt-pinned-schema-identity", "upstream_message": "ves.io.schema.ObjectRefType", "version": 1}, "relationships": [], "schema_path": ["access_info", "tls_config", "cert_params", "certificates", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["access info tls config cert params certificates namespace"], "anchor": "schema-access_info--tls_config--cert_params--certificates--namespace", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace.", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:cert_params:certificates", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "reference_identity": {"member": "namespace", "scope_path": ["access_info", "tls_config", "cert_params", "certificates"], "source": "receipt-pinned-schema-identity", "upstream_message": "ves.io.schema.ObjectRefType", "version": 1}, "relationships": [], "schema_path": ["access_info", "tls_config", "cert_params", "certificates", "namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["access info tls config cert params certificates tenant"], "anchor": "schema-access_info--tls_config--cert_params--certificates--tenant", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant.", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:cert_params:certificates", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "reference_identity": {"member": "tenant", "scope_path": ["access_info", "tls_config", "cert_params", "certificates"], "source": "receipt-pinned-schema-identity", "upstream_message": "ves.io.schema.ObjectRefType", "version": 1}, "relationships": [], "schema_path": ["access_info", "tls_config", "cert_params", "certificates", "tenant"], "syntax": "attribute", "type": "string"}, {"aliases": ["access info tls config cert params certificates uid"], "anchor": "schema-access_info--tls_config--cert_params--certificates--uid", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. Route's) uid.", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:cert_params:certificates", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "reference_identity": {"member": "uid", "scope_path": ["access_info", "tls_config", "cert_params", "certificates"], "source": "receipt-pinned-schema-identity", "upstream_message": "ves.io.schema.ObjectRefType", "version": 1}, "relationships": [], "schema_path": ["access_info", "tls_config", "cert_params", "certificates", "uid"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/secret_management_access/properties/access_info/tls_config/cert_params/certificates/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "Client TLS Certificate required for mTLS authentication.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# access_info.tls_config.cert_params.certificates

Breadcrumbs:

- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/)
- [access_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/)
- [access_info.tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/tls_config/)
- [access_info.tls_config.cert_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/tls_config/cert_params/)
- access_info.tls_config.cert_params.certificates

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Client TLS Certificate required for mTLS authentication.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

Terraform syntax:

```terraform
certificates {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-access_info--tls_config--cert_params--certificates--kind"></a>

### kind property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-16T10:31:02+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-access_info--tls_config--cert_params--certificates--name"></a>

### name property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-16T10:31:02+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-access_info--tls_config--cert_params--certificates--namespace"></a>

### namespace property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
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
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-10-16T10:31:02+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-access_info--tls_config--cert_params--certificates--tenant"></a>

### tenant property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-16T10:31:02+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-access_info--tls_config--cert_params--certificates--uid"></a>

### uid property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-16T10:31:02+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```
