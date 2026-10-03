---
page_title: "ring_hash.hash_policy"
subcategory: "Load Balancing"
description: "Specifies a list of hash policies to use for ring hash load balancing. Each hash policy is evaluated individually and the combined result is used to route the request."
xcsh_docs: {"aliases": ["ring hash hash policy"], "body_bytes": 4361, "body_sha256": "sha256:80d0e4b0d9c887bcf6c6a5232cdf68f19df1aea7865b07b67607f4ab1443e76e", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:ring_hash:hash_policy:cookie"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:ring_hash:hash_policy", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:ring_hash", "path": "documentation/data-sources/http_loadbalancer/properties/ring_hash/hash_policy/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3100323310012210-0321302221233031-2333102213101122-0102010120300121-1121012300222112-2310201300213122-1023032330011310-3332013312113120", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-023.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ring_hash", "hash_policy"], "schema_version": 1, "sections": [{"aliases": ["ring hash hash policy cookie"], "anchor": "section", "description": "Two types of cookie affinity: 1. Passive. Takes a cookie that's present in the cookies header and hashes on its value. 2. Generated. Generates and sets a cookie with an expiration (TTL) on the first request from the client in its response to the client, based on the endpoint the request gets sent to. The client then", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:ring_hash:hash_policy:cookie", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ring_hash", "hash_policy", "cookie"], "syntax": "attribute", "type": "object"}, {"aliases": ["ring hash hash policy header name"], "anchor": "schema-ring_hash--hash_policy--header_name", "description": "Exclusive with The name or key of the request header that will be used to obtain the hash key.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:ring_hash:hash_policy", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ring_hash", "hash_policy", "header_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["ring hash hash policy source ip"], "anchor": "schema-ring_hash--hash_policy--source_ip", "description": "Exclusive with Hash based on source IP address.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:ring_hash:hash_policy", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ring_hash", "hash_policy", "source_ip"], "syntax": "attribute", "type": "bool"}, {"aliases": ["ring hash hash policy terminal"], "anchor": "schema-ring_hash--hash_policy--terminal", "description": "Specify if its a terminal policy.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:ring_hash:hash_policy", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ring_hash", "hash_policy", "terminal"], "syntax": "attribute", "type": "bool"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/ring_hash/hash_policy/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Specifies a list of hash policies to use for ring hash load balancing. Each hash policy is evaluated individually and the combined result is used to route the request.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ring_hash.hash_policy

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [ring_hash](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/ring_hash/)
- ring_hash.hash_policy

<a id="section"></a>

Type: `"list"`. Computed.

Specifies a list of hash policies to use for ring hash load balancing. Each hash policy is evaluated
individually and the combined result is used to route the request.

Upstream description:

Specifies a list of hash policies to use for ring hash load balancing. Each hash policy is evaluated
individually and the combined result is used to route the request.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

- [cookie](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/ring_hash/hash_policy/cookie/): complete subsection reference.

<a id="schema-ring_hash--hash_policy--header_name"></a>

### header_name property

Type: `"string"`. Computed.

Exclusive with \[cookie source\_ip\] The name or key of the request header that will be used to
obtain the hash key.

Upstream description:

Exclusive with \[cookie source\_ip\] The name or key of the request header that will be used to
obtain the hash key.

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
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="schema-ring_hash--hash_policy--source_ip"></a>

### source_ip property

Type: `"bool"`. Computed.

Exclusive with \[cookie header\_name\] Hash based on source IP address.

Upstream description:

Exclusive with \[cookie header\_name\] Hash based on source IP address.

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

<a id="schema-ring_hash--hash_policy--terminal"></a>

### terminal property

Type: `"bool"`. Computed.

Terminal. Specify if its a terminal policy.

Upstream description:

Specify if its a terminal policy.

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

## Next pages

- [ring_hash.hash_policy.cookie](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/ring_hash/hash_policy/cookie/)
- [ring_hash](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/ring_hash/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
