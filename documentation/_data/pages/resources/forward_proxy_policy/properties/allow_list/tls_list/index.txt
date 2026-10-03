---
page_title: "allow_list.tls_list"
subcategory: "Security"
description: "Domains in SNI for TLS connections."
xcsh_docs: {"aliases": ["allow list tls list"], "body_bytes": 6187, "body_sha256": "sha256:c4e1932ebbbc867b5dad977f61ca16c0e1f6161f3e456d7e8368c150cc039195", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:forward_proxy_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:forward_proxy_policy:properties:allow_list:tls_list", "parent_id": "xcsh-docs:resources:forward_proxy_policy:properties:allow_list", "path": "documentation/resources/forward_proxy_policy/properties/allow_list/tls_list/index.md", "product": "distributed-cloud", "provider_name": "forward_proxy_policy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2210220233212330-2121101130022310-2101231231222001-0223201021123312-3112220102330021-2120202030302111-2222203232332103-0133113122113112", "registry_path": "docs/guides/resources--forward_proxy_policy--reference--group-001.md", "relationships": [{"anchor": "schema-allow_list--tls_list--exact_value", "enforcement": "provider-schema", "group": "allow_list.tls_list:ConflictingListObjectAttributes:exact_value,regex_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:forward_proxy_policy:properties:allow_list:tls_list", "type": "conflicts"}, {"anchor": "schema-allow_list--tls_list--exact_value", "enforcement": "provider-schema", "group": "allow_list.tls_list:ConflictingListObjectAttributes:exact_value,suffix_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:forward_proxy_policy:properties:allow_list:tls_list", "type": "conflicts"}, {"anchor": "schema-allow_list--tls_list--regex_value", "enforcement": "provider-schema", "group": "allow_list.tls_list:ConflictingListObjectAttributes:exact_value,regex_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:forward_proxy_policy:properties:allow_list:tls_list", "type": "conflicts"}, {"anchor": "schema-allow_list--tls_list--regex_value", "enforcement": "provider-schema", "group": "allow_list.tls_list:ConflictingListObjectAttributes:regex_value,suffix_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:forward_proxy_policy:properties:allow_list:tls_list", "type": "conflicts"}, {"anchor": "schema-allow_list--tls_list--suffix_value", "enforcement": "provider-schema", "group": "allow_list.tls_list:ConflictingListObjectAttributes:exact_value,suffix_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:forward_proxy_policy:properties:allow_list:tls_list", "type": "conflicts"}, {"anchor": "schema-allow_list--tls_list--suffix_value", "enforcement": "provider-schema", "group": "allow_list.tls_list:ConflictingListObjectAttributes:regex_value,suffix_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:forward_proxy_policy:properties:allow_list:tls_list", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["allow_list", "tls_list"], "schema_version": 1, "sections": [{"aliases": ["allow list tls list exact value"], "anchor": "schema-allow_list--tls_list--exact_value", "description": "Exclusive with Exact domain name.", "document_id": "xcsh-docs:resources:forward_proxy_policy:properties:allow_list:tls_list", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["allow_list", "tls_list", "exact_value"], "syntax": "attribute", "type": "string"}, {"aliases": ["allow list tls list regex value"], "anchor": "schema-allow_list--tls_list--regex_value", "description": "Exclusive with Regular Expression value for the domain name.", "document_id": "xcsh-docs:resources:forward_proxy_policy:properties:allow_list:tls_list", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["allow_list", "tls_list", "regex_value"], "syntax": "attribute", "type": "string"}, {"aliases": ["allow list tls list suffix value"], "anchor": "schema-allow_list--tls_list--suffix_value", "description": "Exclusive with Suffix of domain name e.g \"xyz.com\" will match \"*.xyz.com\" and \"xyz.com\"", "document_id": "xcsh-docs:resources:forward_proxy_policy:properties:allow_list:tls_list", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["allow_list", "tls_list", "suffix_value"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/forward_proxy_policy/properties/allow_list/tls_list/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Domains in SNI for TLS connections.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["forward_proxy_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# allow_list.tls_list

Breadcrumbs:

- [xcsh_forward_proxy_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/)
- [allow_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/allow_list/)
- allow_list.tls_list

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

TLS Domains. Domains in SNI for TLS connections.

Upstream description:

Domains in SNI for TLS connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingListObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingListObjectAttributes("regex_value",
    "suffix_value")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
tls_list {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-allow_list--tls_list--exact_value"></a>

### exact_value property

Type: `"string"`. Optional.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

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
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="schema-allow_list--tls_list--regex_value"></a>

### regex_value property

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

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
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="schema-allow_list--tls_list--suffix_value"></a>

### suffix_value property

Type: `"string"`. Optional.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

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
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

## Next pages

- [allow_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/allow_list/)
- [xcsh_forward_proxy_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/)
