---
page_title: "access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list"
subcategory: ""
description: "Reference to Root CA Certificate."
xcsh_docs: {"aliases": ["access info tls config common params validation params trusted ca trusted ca list"], "body_bytes": 6502, "body_sha256": "sha256:e461525190f16c1e3a5e2036a9ddcec205360330fb08c47713a15a240a343554", "capabilities": ["identity"], "category": "identity", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:common_params:validation_params:trusted_ca:trusted_ca_list", "parent_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:common_params:validation_params:trusted_ca", "path": "documentation/data-sources/secret_management_access/properties/access_info/tls_config/common_params/validation_params/trusted_ca/trusted_ca_list/index.md", "product": "distributed-cloud", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-3213223132332031-0010001010211223-2310021322133310-0221230213332012-3023122303213312-2101331020231003-1311031113300131-2323011120232000", "registry_path": "docs/guides/data-sources--secret_management_access--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["access_info", "tls_config", "common_params", "validation_params", "trusted_ca", "trusted_ca_list"], "schema_version": 1, "sections": [{"aliases": ["access info tls config common params validation params trusted ca trusted ca list kind"], "anchor": "schema-access_info--tls_config--common_params--validation_params--trusted_ca--trusted_ca_list--kind", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. \"route\")", "document_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:common_params:validation_params:trusted_ca:trusted_ca_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "reference_identity": {"member": "kind", "scope_path": ["access_info", "tls_config", "common_params", "validation_params", "trusted_ca", "trusted_ca_list"], "source": "receipt-pinned-schema-identity", "upstream_message": "ves.io.schema.ObjectRefType", "version": 1}, "relationships": [], "schema_path": ["access_info", "tls_config", "common_params", "validation_params", "trusted_ca", "trusted_ca_list", "kind"], "syntax": "attribute", "type": "string"}, {"aliases": ["access info tls config common params validation params trusted ca trusted ca list name"], "anchor": "schema-access_info--tls_config--common_params--validation_params--trusted_ca--trusted_ca_list--name", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name.", "document_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:common_params:validation_params:trusted_ca:trusted_ca_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "reference_identity": {"member": "name", "scope_path": ["access_info", "tls_config", "common_params", "validation_params", "trusted_ca", "trusted_ca_list"], "source": "receipt-pinned-schema-identity", "upstream_message": "ves.io.schema.ObjectRefType", "version": 1}, "relationships": [], "schema_path": ["access_info", "tls_config", "common_params", "validation_params", "trusted_ca", "trusted_ca_list", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["access info tls config common params validation params trusted ca trusted ca list namespace"], "anchor": "schema-access_info--tls_config--common_params--validation_params--trusted_ca--trusted_ca_list--namespace", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace.", "document_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:common_params:validation_params:trusted_ca:trusted_ca_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "reference_identity": {"member": "namespace", "scope_path": ["access_info", "tls_config", "common_params", "validation_params", "trusted_ca", "trusted_ca_list"], "source": "receipt-pinned-schema-identity", "upstream_message": "ves.io.schema.ObjectRefType", "version": 1}, "relationships": [], "schema_path": ["access_info", "tls_config", "common_params", "validation_params", "trusted_ca", "trusted_ca_list", "namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["access info tls config common params validation params trusted ca trusted ca list tenant"], "anchor": "schema-access_info--tls_config--common_params--validation_params--trusted_ca--trusted_ca_list--tenant", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant.", "document_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:common_params:validation_params:trusted_ca:trusted_ca_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "reference_identity": {"member": "tenant", "scope_path": ["access_info", "tls_config", "common_params", "validation_params", "trusted_ca", "trusted_ca_list"], "source": "receipt-pinned-schema-identity", "upstream_message": "ves.io.schema.ObjectRefType", "version": 1}, "relationships": [], "schema_path": ["access_info", "tls_config", "common_params", "validation_params", "trusted_ca", "trusted_ca_list", "tenant"], "syntax": "attribute", "type": "string"}, {"aliases": ["access info tls config common params validation params trusted ca trusted ca list uid"], "anchor": "schema-access_info--tls_config--common_params--validation_params--trusted_ca--trusted_ca_list--uid", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. Route's) uid.", "document_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:common_params:validation_params:trusted_ca:trusted_ca_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "reference_identity": {"member": "uid", "scope_path": ["access_info", "tls_config", "common_params", "validation_params", "trusted_ca", "trusted_ca_list"], "source": "receipt-pinned-schema-identity", "upstream_message": "ves.io.schema.ObjectRefType", "version": 1}, "relationships": [], "schema_path": ["access_info", "tls_config", "common_params", "validation_params", "trusted_ca", "trusted_ca_list", "uid"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/secret_management_access/properties/access_info/tls_config/common_params/validation_params/trusted_ca/trusted_ca_list/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Reference to Root CA Certificate.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list

Breadcrumbs:

- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/)
- [access_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/)
- [access_info.tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/)
- [access_info.tls_config.common_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/common_params/)
- [access_info.tls_config.common_params.validation_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/common_params/validation_params/)
- [access_info.tls_config.common_params.validation_params.trusted_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/common_params/validation_params/trusted_ca/)
- access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list

<a id="section"></a>

Type: `"list"`. Computed.

Root CA Certificate Reference. Reference to Root CA Certificate.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

## Direct properties

<a id="schema-access_info--tls_config--common_params--validation_params--trusted_ca--trusted_ca_list--kind"></a>

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="schema-access_info--tls_config--common_params--validation_params--trusted_ca--trusted_ca_list--name"></a>

### name property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="schema-access_info--tls_config--common_params--validation_params--trusted_ca--trusted_ca_list--namespace"></a>

### namespace property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="schema-access_info--tls_config--common_params--validation_params--trusted_ca--trusted_ca_list--tenant"></a>

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="schema-access_info--tls_config--common_params--validation_params--trusted_ca--trusted_ca_list--uid"></a>

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
