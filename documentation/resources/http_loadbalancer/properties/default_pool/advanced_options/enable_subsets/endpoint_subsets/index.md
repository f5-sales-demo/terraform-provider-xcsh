---
page_title: "default_pool.advanced_options.enable_subsets.endpoint_subsets"
subcategory: "Load Balancing"
description: "List of subset class. Subsets class is defined using list of keys. Every unique combination of values of these keys form a subset within the class."
xcsh_docs: {"aliases": ["default pool advanced options enable subsets endpoint subsets"], "body_bytes": 3687, "body_sha256": "sha256:2ab6ff923ce682d89c8f4856212f53bba5f33be73326c9a39d1bfa28bcff03a2", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:enable_subsets:endpoint_subsets", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:enable_subsets", "path": "documentation/resources/http_loadbalancer/properties/default_pool/advanced_options/enable_subsets/endpoint_subsets/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1323012021010023-2303102102321011-0331330121102101-3230020120302033-1222230310221220-3003022122203230-0012010311033122-2233222113222233", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-015.md", "relationships": [{"anchor": "schema-default_pool--advanced_options--enable_subsets--endpoint_subsets--keys", "enforcement": "provider-schema", "group": "default_pool.advanced_options.enable_subsets.endpoint_subsets:RequiredListObjectAttributes:keys", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:enable_subsets:endpoint_subsets", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["default_pool", "advanced_options", "enable_subsets", "endpoint_subsets"], "schema_version": 1, "sections": [{"aliases": ["default pool advanced options enable subsets endpoint subsets keys"], "anchor": "schema-default_pool--advanced_options--enable_subsets--endpoint_subsets--keys", "description": "List of keys that define a cluster subset class.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:enable_subsets:endpoint_subsets", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_pool", "advanced_options", "enable_subsets", "endpoint_subsets", "keys"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/default_pool/advanced_options/enable_subsets/endpoint_subsets/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of subset class. Subsets class is defined using list of keys. Every unique combination of values of these keys form a subset within the class.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool.advanced_options.enable_subsets.endpoint_subsets

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [default_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/)
- [default_pool.advanced_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/advanced_options/)
- [default_pool.advanced_options.enable_subsets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/advanced_options/enable_subsets/)
- default_pool.advanced_options.enable_subsets.endpoint_subsets

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

List of subset class. Subsets class is defined using list of keys. Every unique combination of
values of these keys form a subset within the class.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("keys")}
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
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

Terraform syntax:

```terraform
endpoint_subsets {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-default_pool--advanced_options--enable_subsets--endpoint_subsets--keys"></a>

### keys property

Type: `["list", "string"]`. Optional.

List of keys that define a cluster subset class.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
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
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

## Next pages

- [default_pool.advanced_options.enable_subsets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/advanced_options/enable_subsets/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
