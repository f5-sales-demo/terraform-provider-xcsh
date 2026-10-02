---
page_title: "allowed_response_codes"
subcategory: "Security"
description: "List of HTTP response status codes that are allowed."
xcsh_docs: {"aliases": ["allowed response codes"], "body_bytes": 2750, "body_sha256": "sha256:46159ac939bdec7f4e7950638d844a88e665107d15067d48031003125f4b3c78", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_firewall:properties:allowed_response_codes", "parent_id": "xcsh-docs:resources:app_firewall:reference", "path": "documentation/resources/app_firewall/properties/allowed_response_codes/index.md", "product": "distributed-cloud", "provider_name": "app_firewall", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0120311133313213-1113131100301122-3201330201312210-0023331230030310-2032013032030031-0122113221102121-2233331112123232-0130001130213312", "registry_path": "docs/guides/resources--app_firewall--reference--group-001.md", "relationships": [{"anchor": "schema-allowed_response_codes--response_code", "enforcement": "provider-schema", "group": "allowed_response_codes:RequiredObjectAttributes:response_code", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:app_firewall:properties:allowed_response_codes", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["allowed_response_codes"], "schema_version": 1, "sections": [{"aliases": ["response code"], "anchor": "schema-allowed_response_codes--response_code", "description": "List of HTTP response status codes that are allowed.", "document_id": "xcsh-docs:resources:app_firewall:properties:allowed_response_codes", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["allowed_response_codes", "response_code"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_firewall/properties/allowed_response_codes/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of HTTP response status codes that are allowed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["app_firewallCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# allowed_response_codes

Breadcrumbs:

- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/)
- allowed_response_codes

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

List of HTTP response status codes that are allowed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("response_code")}
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
allowed_response_codes {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-allowed_response_codes--response_code"></a>

### response_code property

Type: `["list", "number"]`. Optional.

List of HTTP response status codes that are allowed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 48),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 48,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 48,
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
    "ves.io.schema.rules.repeated.items.uint32.gte": "100",
    "ves.io.schema.rules.repeated.items.uint32.lte": "999",
    "ves.io.schema.rules.repeated.max_items": "48",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.uint32.gte": "100",
    "ves.io.schema.rules.repeated.items.uint32.lte": "999",
    "ves.io.schema.rules.repeated.max_items": "48",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/)
- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/)
