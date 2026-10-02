---
page_title: "ddos_mitigation_rules.ddos_client_source.ja4_tls_fingerprint_matcher"
subcategory: "Load Balancing"
description: "An extended version of JA3 that includes additional fields for more comprehensive fingerprinting of SSL/TLS clients and potentially has a different structure and length."
xcsh_docs: {"aliases": ["ddos mitigation rules ddos client source ja4 tls fingerprint matcher"], "body_bytes": 3241, "body_sha256": "sha256:dd38af2e347d90f636b349b697c9b2991ba85ca0bffc6b98a81e1857cda8bab2", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:ddos_mitigation_rules:ddos_client_source:ja4_tls_fingerprint_matcher", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:ddos_mitigation_rules:ddos_client_source", "path": "documentation/resources/http_loadbalancer/properties/ddos_mitigation_rules/ddos_client_source/ja4_tls_fingerprint_matcher/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2210313030223112-0221132223013102-2133320023222302-2223212323022331-0111110010032032-2231320102021233-2132022013210322-2231020122012333", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-015.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ddos_mitigation_rules", "ddos_client_source", "ja4_tls_fingerprint_matcher"], "schema_version": 1, "sections": [{"aliases": ["exact values"], "anchor": "schema-ddos_mitigation_rules--ddos_client_source--ja4_tls_fingerprint_matcher--exact_values", "description": "A list of exact JA4 TLS fingerprint to match the input JA4 TLS fingerprint against.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:ddos_mitigation_rules:ddos_client_source:ja4_tls_fingerprint_matcher", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ddos_mitigation_rules", "ddos_client_source", "ja4_tls_fingerprint_matcher", "exact_values"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/ddos_mitigation_rules/ddos_client_source/ja4_tls_fingerprint_matcher/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "An extended version of JA3 that includes additional fields for more comprehensive fingerprinting of SSL/TLS clients and potentially has a different structure and length.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ddos_mitigation_rules.ddos_client_source.ja4_tls_fingerprint_matcher

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [ddos_mitigation_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/ddos_mitigation_rules/)
- [ddos_mitigation_rules.ddos_client_source](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/ddos_mitigation_rules/ddos_client_source/)
- ddos_mitigation_rules.ddos_client_source.ja4_tls_fingerprint_matcher

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Extended version of JA3 that includes additional fields for more comprehensive fingerprinting of
SSL/TLS clients and potentially has a different structure and length.

Upstream description:

An extended version of JA3 that includes additional fields for more comprehensive fingerprinting of
SSL/TLS clients and potentially has a different structure and length.

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
ja4_tls_fingerprint_matcher {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-ddos_mitigation_rules--ddos_client_source--ja4_tls_fingerprint_matcher--exact_values"></a>

### exact_values property

Type: `["list", "string"]`. Optional.

List of exact JA4 TLS fingerprint to match the input JA4 TLS fingerprint against.

Upstream description:

A list of exact JA4 TLS fingerprint to match the input JA4 TLS fingerprint against.

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
    "ves.io.schema.rules.repeated.items.string.len": "36",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "36",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Next pages

- [ddos_mitigation_rules.ddos_client_source](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/ddos_mitigation_rules/ddos_client_source/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
