---
page_title: "request_headers_to_add"
subcategory: ""
description: "Headers are key-value pairs to be added to HTTP request being routed towards upstream. Headers specified at this level are applied after headers from matched Route are applied."
xcsh_docs: {"aliases": ["request headers to add"], "body_bytes": 5668, "body_sha256": "sha256:0c1de986553a45e1a309dc8bd778ae683be5854e2c8ba67bdb3a4a76e591ce8f", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:virtual_host:properties:request_headers_to_add:secret_value"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:request_headers_to_add", "parent_id": "xcsh-docs:resources:virtual_host:reference", "path": "documentation/resources/virtual_host/properties/request_headers_to_add/index.md", "product": "distributed-cloud", "provider_name": "virtual_host", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3023100112011132-2302030100320013-2030100013010220-0112333330230033-0110121200123102-2012130203132311-0221233210230001-2212303220211221", "registry_path": "docs/guides/resources--virtual_host--reference--group-002.md", "relationships": [{"anchor": "schema-request_headers_to_add--value", "enforcement": "provider-schema", "group": "request_headers_to_add:ConflictingListObjectAttributes:secret_value,value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:request_headers_to_add", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "request_headers_to_add:ConflictingListObjectAttributes:secret_value,value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:request_headers_to_add:secret_value", "type": "conflicts"}, {"anchor": "schema-request_headers_to_add--name", "enforcement": "provider-schema", "group": "request_headers_to_add:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:request_headers_to_add", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["request_headers_to_add"], "schema_version": 1, "sections": [{"aliases": ["append"], "anchor": "schema-request_headers_to_add--append", "description": "Should the value be appended? If true, the value is appended to existing values. Default value is do not append.", "document_id": "xcsh-docs:resources:virtual_host:properties:request_headers_to_add", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["request_headers_to_add", "append"], "syntax": "attribute", "type": "bool"}, {"aliases": ["name"], "anchor": "schema-request_headers_to_add--name", "description": "Name of the HTTP header.", "document_id": "xcsh-docs:resources:virtual_host:properties:request_headers_to_add", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["request_headers_to_add", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["secret value"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:virtual_host:properties:request_headers_to_add:secret_value", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "request_headers_to_add.secret_value:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:request_headers_to_add:secret_value:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "request_headers_to_add.secret_value:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:request_headers_to_add:secret_value:clear_secret_info", "type": "conflicts"}], "schema_path": ["request_headers_to_add", "secret_value"], "syntax": "block", "type": "object"}, {"aliases": ["value"], "anchor": "schema-request_headers_to_add--value", "description": "Exclusive with Value of the HTTP header.", "document_id": "xcsh-docs:resources:virtual_host:properties:request_headers_to_add", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["request_headers_to_add", "value"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/request_headers_to_add/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Headers are key-value pairs to be added to HTTP request being routed towards upstream. Headers specified at this level are applied after headers from matched Route are applied.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# request_headers_to_add

Breadcrumbs:

- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/)
- request_headers_to_add

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Headers are key-value pairs to be added to HTTP request being routed towards upstream. Headers
specified at this level are applied after headers from matched Route are applied.

Upstream description:

Headers are key-value pairs to be added to HTTP request being routed towards upstream. Headers
specified at this level are applied after headers from matched Route are applied.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("secret_value",
    "value")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
request_headers_to_add {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-request_headers_to_add--append"></a>

### append property

Type: `"bool"`. Optional.

Should the value be appended? If true, the value is appended to existing values. not append.
Defaults to \`do\`.

Upstream description:

Should the value be appended? If true, the value is appended to existing values. Default value is do
not append.

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

<a id="schema-request_headers_to_add--name"></a>

### name property

Type: `"string"`. Optional.

Name. Name of the HTTP header.

Upstream description:

Name of the HTTP header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [secret_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/request_headers_to_add/secret_value/): complete subsection reference.

<a id="schema-request_headers_to_add--value"></a>

### value property

Type: `"string"`. Optional.

Exclusive with \[secret\_value\] Value of the HTTP header.

Upstream description:

Exclusive with \[secret\_value\] Value of the HTTP header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8096),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

## Next pages

- [request_headers_to_add.secret_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/request_headers_to_add/secret_value/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/)
- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/)
