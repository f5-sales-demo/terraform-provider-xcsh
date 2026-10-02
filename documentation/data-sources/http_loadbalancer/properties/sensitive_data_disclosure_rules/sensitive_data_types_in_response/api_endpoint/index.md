---
page_title: "sensitive_data_disclosure_rules.sensitive_data_types_in_response.api_endpoint"
subcategory: "Load Balancing"
description: "This defines API endpoint."
xcsh_docs: {"aliases": ["sensitive data disclosure rules sensitive data types in response api endpoint"], "body_bytes": 4132, "body_sha256": "sha256:9c09ab650083599f4136d50e7cc26b4ab5c9a73e6127bda613087a867f011526", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:sensitive_data_disclosure_rules:sensitive_data_types_in_response:api_endpoint", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:sensitive_data_disclosure_rules:sensitive_data_types_in_response", "path": "documentation/data-sources/http_loadbalancer/properties/sensitive_data_disclosure_rules/sensitive_data_types_in_response/api_endpoint/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0223211112203132-3213101003131222-1230332300321100-1333011021323120-0103233313332121-2333321230102002-0320003223022021-3313002203010002", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-026.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["sensitive_data_disclosure_rules", "sensitive_data_types_in_response", "api_endpoint"], "schema_version": 1, "sections": [{"aliases": ["methods"], "anchor": "schema-sensitive_data_disclosure_rules--sensitive_data_types_in_response--api_endpoint--methods", "description": "Methods to be matched.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:sensitive_data_disclosure_rules:sensitive_data_types_in_response:api_endpoint", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["sensitive_data_disclosure_rules", "sensitive_data_types_in_response", "api_endpoint", "methods"], "syntax": "attribute", "type": "list"}, {"aliases": ["path"], "anchor": "schema-sensitive_data_disclosure_rules--sensitive_data_types_in_response--api_endpoint--path", "description": "Path to be matched.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:sensitive_data_disclosure_rules:sensitive_data_types_in_response:api_endpoint", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["sensitive_data_disclosure_rules", "sensitive_data_types_in_response", "api_endpoint", "path"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/sensitive_data_disclosure_rules/sensitive_data_types_in_response/api_endpoint/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This defines API endpoint.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# sensitive_data_disclosure_rules.sensitive_data_types_in_response.api_endpoint

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [sensitive_data_disclosure_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/sensitive_data_disclosure_rules/)
- [sensitive_data_disclosure_rules.sensitive_data_types_in_response](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/sensitive_data_disclosure_rules/sensitive_data_types_in_response/)
- sensitive_data_disclosure_rules.sensitive_data_types_in_response.api_endpoint

<a id="section"></a>

Type: `"single"`. Computed.

API Endpoint. This defines API endpoint.

Upstream description:

This defines API endpoint.

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

## Direct properties

<a id="schema-sensitive_data_disclosure_rules--sensitive_data_types_in_response--api_endpoint--methods"></a>

### methods property

Type: `["list", "string"]`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Methods. Methods to be
matched. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`,
\`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Upstream description:

Methods to be matched.

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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="schema-sensitive_data_disclosure_rules--sensitive_data_types_in_response--api_endpoint--path"></a>

### path property

Type: `"string"`. Computed.

Path. Path to be matched.

Upstream description:

Path to be matched.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  }
}
```

## Next pages

- [sensitive_data_disclosure_rules.sensitive_data_types_in_response](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/sensitive_data_disclosure_rules/sensitive_data_types_in_response/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
