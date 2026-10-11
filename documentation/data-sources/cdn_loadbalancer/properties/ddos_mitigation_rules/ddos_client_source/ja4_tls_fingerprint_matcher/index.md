---
page_title: "ddos_mitigation_rules.ddos_client_source.ja4_tls_fingerprint_matcher"
subcategory: "Load Balancing"
description: "An extended version of JA3 that includes additional fields for more comprehensive fingerprinting of SSL/TLS clients and potentially has a different structure and length."
xcsh_docs: {"aliases": ["ddos mitigation rules ddos client source ja4 tls fingerprint matcher"], "body_bytes": 2372, "body_sha256": "sha256:399088bd350629bae4f5c9154aa682db1a1b6ef52b0f1719c08fd937011ed3d0", "capabilities": ["cdn"], "category": "cdn", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:ddos_mitigation_rules:ddos_client_source:ja4_tls_fingerprint_matcher", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:ddos_mitigation_rules:ddos_client_source", "path": "documentation/data-sources/cdn_loadbalancer/properties/ddos_mitigation_rules/ddos_client_source/ja4_tls_fingerprint_matcher/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "data-sources", "registry_anchor": "canonical-1010020312100303-2110330230121312-2302333202321113-3003032012103131-2011020302013313-2321131022313002-0021311001101200-3122032201311210", "registry_path": "docs/guides/data-sources--cdn_loadbalancer--reference--group-010.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ddos_mitigation_rules", "ddos_client_source", "ja4_tls_fingerprint_matcher"], "schema_version": 1, "sections": [{"aliases": ["ddos mitigation rules ddos client source ja4 tls fingerprint matcher exact values"], "anchor": "schema-ddos_mitigation_rules--ddos_client_source--ja4_tls_fingerprint_matcher--exact_values", "description": "A list of exact JA4 TLS fingerprint to match the input JA4 TLS fingerprint against.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:ddos_mitigation_rules:ddos_client_source:ja4_tls_fingerprint_matcher", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ddos_mitigation_rules", "ddos_client_source", "ja4_tls_fingerprint_matcher", "exact_values"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/ddos_mitigation_rules/ddos_client_source/ja4_tls_fingerprint_matcher/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "An extended version of JA3 that includes additional fields for more comprehensive fingerprinting of SSL/TLS clients and potentially has a different structure and length.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ddos_mitigation_rules.ddos_client_source.ja4_tls_fingerprint_matcher

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- [ddos_mitigation_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/ddos_mitigation_rules/)
- [ddos_mitigation_rules.ddos_client_source](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/ddos_mitigation_rules/ddos_client_source/)
- ddos_mitigation_rules.ddos_client_source.ja4_tls_fingerprint_matcher

<a id="section"></a>

Type: `"single"`. Computed.

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

## Direct properties

<a id="schema-ddos_mitigation_rules--ddos_client_source--ja4_tls_fingerprint_matcher--exact_values"></a>

### exact_values property

Type: `["list", "string"]`. Computed.

A list of exact JA4 TLS fingerprint to match the input JA4 TLS fingerprint against.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
