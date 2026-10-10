---
page_title: "cookie_stickiness"
subcategory: "Load Balancing"
description: "Two types of cookie affinity: 1. Passive. Takes a cookie that's present in the cookies header and hashes on its value. 2. Generated. Generates and sets a cookie with an expiration (TTL) on the first request from the client in its response to the client, based on the endpoint the request gets sent to. The client then"
xcsh_docs: {"aliases": ["cookie stickiness"], "body_bytes": 7200, "body_sha256": "sha256:4dcb0c83c4c62934462f317c7e9ea2ce880fd9a04de568ea26a1a2708aae4567", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:cookie_stickiness:add_httponly", "xcsh-docs:resources:http_loadbalancer:properties:cookie_stickiness:add_secure", "xcsh-docs:resources:http_loadbalancer:properties:cookie_stickiness:ignore_httponly", "xcsh-docs:resources:http_loadbalancer:properties:cookie_stickiness:ignore_samesite", "xcsh-docs:resources:http_loadbalancer:properties:cookie_stickiness:ignore_secure", "xcsh-docs:resources:http_loadbalancer:properties:cookie_stickiness:samesite_lax", "xcsh-docs:resources:http_loadbalancer:properties:cookie_stickiness:samesite_none", "xcsh-docs:resources:http_loadbalancer:properties:cookie_stickiness:samesite_strict"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:cookie_stickiness", "parent_id": "xcsh-docs:resources:http_loadbalancer:reference", "path": "documentation/resources/http_loadbalancer/properties/cookie_stickiness/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-3322222220120200-1330331312120131-2323231032320212-0121111210230332-1011313101330022-0002302120301202-0110200313122233-1010303103322111", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-015.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cookie_stickiness"], "schema_version": 1, "sections": [{"aliases": ["cookie stickiness add httponly"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:cookie_stickiness:add_httponly", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cookie_stickiness", "add_httponly"], "syntax": "attribute", "type": "object"}, {"aliases": ["cookie stickiness add secure"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:cookie_stickiness:add_secure", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cookie_stickiness", "add_secure"], "syntax": "attribute", "type": "object"}, {"aliases": ["cookie stickiness ignore httponly"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:cookie_stickiness:ignore_httponly", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cookie_stickiness", "ignore_httponly"], "syntax": "attribute", "type": "object"}, {"aliases": ["cookie stickiness ignore samesite"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:cookie_stickiness:ignore_samesite", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cookie_stickiness", "ignore_samesite"], "syntax": "attribute", "type": "object"}, {"aliases": ["cookie stickiness ignore secure"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:cookie_stickiness:ignore_secure", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cookie_stickiness", "ignore_secure"], "syntax": "attribute", "type": "object"}, {"aliases": ["cookie stickiness name"], "anchor": "schema-cookie_stickiness--name", "description": "The name of the cookie that will be used to obtain the hash key. If the cookie is not present and TTL below is not set, no hash will be produced.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:cookie_stickiness", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cookie_stickiness", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["cookie stickiness path"], "anchor": "schema-cookie_stickiness--path", "description": "The name of the path for the cookie. If no path is specified here, no path will be set for the cookie.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:cookie_stickiness", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cookie_stickiness", "path"], "syntax": "attribute", "type": "string"}, {"aliases": ["cookie stickiness samesite lax"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:cookie_stickiness:samesite_lax", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cookie_stickiness", "samesite_lax"], "syntax": "attribute", "type": "object"}, {"aliases": ["cookie stickiness samesite none"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:cookie_stickiness:samesite_none", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cookie_stickiness", "samesite_none"], "syntax": "attribute", "type": "object"}, {"aliases": ["cookie stickiness samesite strict"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:cookie_stickiness:samesite_strict", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cookie_stickiness", "samesite_strict"], "syntax": "attribute", "type": "object"}, {"aliases": ["cookie stickiness ttl"], "anchor": "schema-cookie_stickiness--ttl", "description": "If specified, a cookie with the TTL will be generated if the cookie is not present. If the TTL is present and zero, the generated cookie will be a session cookie. TTL value is in milliseconds.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:cookie_stickiness", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cookie_stickiness", "ttl"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/cookie_stickiness/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Two types of cookie affinity: 1. Passive. Takes a cookie that's present in the cookies header and hashes on its value. 2. Generated. Generates and sets a cookie with an expiration (TTL) on the first request from the client in its response to the client, based on the endpoint the request gets sent to. The client then", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cookie_stickiness

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- cookie_stickiness

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: cookie\_stickiness, least\_active, random, ring\_hash, round\_robin,
source\_ip\_stickiness; Default: round\_robin\] Two types of cookie affinity: 1. Passive. Takes a
cookie that's present in the cookies header and hashes on its value. 2. Generated. Generates and
sets a cookie with an expiration (TTL) on the first request from the client in its response to the
client, based on the endpoint the request gets..

Additional upstream details:

Two types of cookie affinity:

&#8203;1. Generates and sets a cookie with an expiration (TTL) on the first request from the client
in its response to the client, based on the endpoint the request gets sent to. The client then
presents this on the next and all subsequent requests. The hash of this is sufficient to ensure
these requests GET sent to the same endpoint. The cookie is generated by hashing the source and
destination ports and addresses so that multiple independent HTTP2 streams on the same connection
will independently receive the same cookie, even if they arrive simultaneously.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-httponly": "[\"add_httponly\",\"ignore_httponly\"]",
  "x-ves-oneof-field-samesite": "[\"ignore_samesite\",\"samesite_lax\",\"samesite_none\",\"samesite_strict\"]",
  "x-ves-oneof-field-secure": "[\"add_secure\",\"ignore_secure\"]"
}
```

OneOf alternatives in this subsection:

- [cookie_stickiness](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/cookie_stickiness/#section)
- [least_active](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/least_active/#section)
- [random](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/random/#section)
- [ring_hash](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/ring_hash/#section)
- [round_robin](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/round_robin/#section)
- [source_ip_stickiness](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/source_ip_stickiness/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
cookie_stickiness {
  # Configure direct properties listed below.
}
```

## Direct properties

- [add_httponly](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/cookie_stickiness/add_httponly/): complete subsection reference.

- [add_secure](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/cookie_stickiness/add_secure/): complete subsection reference.

- [ignore_httponly](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/cookie_stickiness/ignore_httponly/): complete subsection reference.

- [ignore_samesite](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/cookie_stickiness/ignore_samesite/): complete subsection reference.

- [ignore_secure](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/cookie_stickiness/ignore_secure/): complete subsection reference.

<a id="schema-cookie_stickiness--name"></a>

### name property

Type: `"string"`. Optional.

The name of the cookie that will be used to obtain the hash key. If the cookie is not present and
TTL below is not set, no hash will be produced.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="schema-cookie_stickiness--path"></a>

### path property

Type: `"string"`. Optional.

The name of the path for the cookie. If no path is specified here, no path will be set for the
cookie.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.8,
      "source": "inferred",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [samesite_lax](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/cookie_stickiness/samesite_lax/): complete subsection reference.

- [samesite_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/cookie_stickiness/samesite_none/): complete subsection reference.

- [samesite_strict](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/cookie_stickiness/samesite_strict/): complete subsection reference.

<a id="schema-cookie_stickiness--ttl"></a>

### ttl property

Type: `"number"`. Optional.

If specified, a cookie with the TTL will be generated if the cookie is not present. If the TTL is
present and zero, the generated cookie will be a session cookie. TTL value is in milliseconds.

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
