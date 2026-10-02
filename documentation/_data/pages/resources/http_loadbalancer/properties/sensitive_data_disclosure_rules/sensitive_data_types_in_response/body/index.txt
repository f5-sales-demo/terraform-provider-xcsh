---
page_title: "sensitive_data_disclosure_rules.sensitive_data_types_in_response.body"
subcategory: "Load Balancing"
description: "OPTIONS for HTTP Body Masking."
xcsh_docs: {"aliases": ["sensitive data disclosure rules sensitive data types in response body"], "body_bytes": 4187, "body_sha256": "sha256:16b1fcf7ea334da0bcacaabf9a4d230a7b44c6ff20f82f98aab65db45f16d33b", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:sensitive_data_disclosure_rules:sensitive_data_types_in_response:body", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:sensitive_data_disclosure_rules:sensitive_data_types_in_response", "path": "documentation/resources/http_loadbalancer/properties/sensitive_data_disclosure_rules/sensitive_data_types_in_response/body/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2323121012223210-0232113003332230-0221101000212202-0011210103332001-1122131032233232-1201112100203231-0302310310112000-1231223202202011", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-026.md", "relationships": [{"anchor": "schema-sensitive_data_disclosure_rules--sensitive_data_types_in_response--body--fields", "enforcement": "provider-schema", "group": "sensitive_data_disclosure_rules.sensitive_data_types_in_response.body:RequiredObjectAttributes:fields", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:sensitive_data_disclosure_rules:sensitive_data_types_in_response:body", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["sensitive_data_disclosure_rules", "sensitive_data_types_in_response", "body"], "schema_version": 1, "sections": [{"aliases": ["fields"], "anchor": "schema-sensitive_data_disclosure_rules--sensitive_data_types_in_response--body--fields", "description": "List of JSON Path field values. Use square brackets with an underscore to indicate array elements (e.g., person.emails). To reference JSON keys that contain spaces, enclose the entire path in double quotes. For example: \"person.first name\".", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:sensitive_data_disclosure_rules:sensitive_data_types_in_response:body", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["sensitive_data_disclosure_rules", "sensitive_data_types_in_response", "body", "fields"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/sensitive_data_disclosure_rules/sensitive_data_types_in_response/body/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "OPTIONS for HTTP Body Masking.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# sensitive_data_disclosure_rules.sensitive_data_types_in_response.body

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [sensitive_data_disclosure_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/sensitive_data_disclosure_rules/)
- [sensitive_data_disclosure_rules.sensitive_data_types_in_response](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/sensitive_data_disclosure_rules/sensitive_data_types_in_response/)
- sensitive_data_disclosure_rules.sensitive_data_types_in_response.body

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Body Section Masking OPTIONS. OPTIONS for HTTP Body Masking.

Upstream description:

OPTIONS for HTTP Body Masking.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("fields")}
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
body {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-sensitive_data_disclosure_rules--sensitive_data_types_in_response--body--fields"></a>

### fields property

Type: `["list", "string"]`. Optional.

List of JSON Path field values. Use square brackets with an underscore \[\_\] to indicate array
elements (e.g., person.emails\[\_\]). To reference JSON keys that contain spaces, enclose the entire
path in double quotes.

Upstream description:

List of JSON Path field values. Use square brackets with an underscore \[\_\] to indicate array
elements (e.g., person.emails\[\_\]). To reference JSON keys that contain spaces, enclose the entire
path in double quotes. For example: "person.first name".

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.items.string.json_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.json_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Next pages

- [sensitive_data_disclosure_rules.sensitive_data_types_in_response](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/sensitive_data_disclosure_rules/sensitive_data_types_in_response/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
