---
page_title: "enable_trust_client_ip_headers"
subcategory: "Load Balancing"
description: "List of Client IP Headers."
xcsh_docs: {"aliases": ["enable trust client ip headers"], "body_bytes": 2995, "body_sha256": "sha256:a37ad2329e2c50e57d915e7795ff12140441f283e7d2db54dd682f914b548a83", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:enable_trust_client_ip_headers", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:reference", "path": "documentation/data-sources/http_loadbalancer/properties/enable_trust_client_ip_headers/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "data-sources", "registry_anchor": "canonical-0223302231131311-1212022331012012-1201212100221020-3222300100201333-2302211010002001-3333320223120100-2001323010222333-0112102133220100", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-018.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["enable_trust_client_ip_headers"], "schema_version": 1, "sections": [{"aliases": ["enable trust client ip headers client ip headers"], "anchor": "schema-enable_trust_client_ip_headers--client_ip_headers", "description": "Define the list of one or more Client IP Headers. Headers will be used in order from top to bottom, meaning if the first header is not present in the request, the system will proceed to check for the second header, and so on, until one of the listed headers is found. If none of the defined headers exist, or the value", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:enable_trust_client_ip_headers", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_trust_client_ip_headers", "client_ip_headers"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/enable_trust_client_ip_headers/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "List of Client IP Headers.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_trust_client_ip_headers

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- enable_trust_client_ip_headers

<a id="section"></a>

Type: `"single"`. Computed.

Trust Client IP Headers List. List of Client IP Headers.

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

<a id="schema-enable_trust_client_ip_headers--client_ip_headers"></a>

### client_ip_headers property

Type: `["list", "string"]`. Computed.

Define the list of one or more Client IP Headers. Headers will be used in order from top to bottom,
meaning if the first header is not present in the request, the system will proceed to check for the
second header, and so on, until one of the listed headers is found. If none of the defined headers
exist, or the value is not an IP address, then the system will use the source IP of the packet. If
multiple defined headers with different names are present in the request, the value of the first
header name in the configuration will be used. If multiple defined headers with the same name are
present in the request, values of all those headers will be combined. The system will read the
right-most IP address from header, if there are multiple IP addresses in the header value. For
X-Forwarded-For header, the system will read the IP address(rightmost - 1), as the client IP.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```
