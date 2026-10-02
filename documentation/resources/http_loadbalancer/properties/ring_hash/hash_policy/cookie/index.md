---
page_title: "ring_hash.hash_policy.cookie"
subcategory: "Load Balancing"
description: "Two types of cookie affinity: 1. Passive. Takes a cookie that's present in the cookies header and hashes on its value. 2. Generated. Generates and sets a cookie with an expiration (TTL) on the first request from the client in its response to the client, based on the endpoint the request gets sent to. The client then"
xcsh_docs: {"aliases": ["ring hash hash policy cookie"], "body_bytes": 9875, "body_sha256": "sha256:535aa642f152ebed966d226ad93f0425f78a4896c88312253c0f80602fb25e83", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy:cookie:add_httponly", "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy:cookie:add_secure", "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy:cookie:ignore_httponly", "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy:cookie:ignore_samesite", "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy:cookie:ignore_secure", "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy:cookie:samesite_lax", "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy:cookie:samesite_none", "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy:cookie:samesite_strict"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy:cookie", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy", "path": "documentation/resources/http_loadbalancer/properties/ring_hash/hash_policy/cookie/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3131013300110333-2102200210330133-2102221222232012-3000213212101123-0202022013221110-1322110330301102-3003222330332021-2332003223213022", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-024.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "ring_hash.hash_policy.cookie:ConflictingObjectAttributes:add_httponly,ignore_httponly", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy:cookie:add_httponly", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ring_hash.hash_policy.cookie:ConflictingObjectAttributes:add_secure,ignore_secure", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy:cookie:add_secure", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ring_hash.hash_policy.cookie:ConflictingObjectAttributes:add_httponly,ignore_httponly", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy:cookie:ignore_httponly", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ring_hash.hash_policy.cookie:ConflictingObjectAttributes:ignore_samesite,samesite_lax", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy:cookie:ignore_samesite", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ring_hash.hash_policy.cookie:ConflictingObjectAttributes:ignore_samesite,samesite_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy:cookie:ignore_samesite", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ring_hash.hash_policy.cookie:ConflictingObjectAttributes:ignore_samesite,samesite_strict", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy:cookie:ignore_samesite", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ring_hash.hash_policy.cookie:ConflictingObjectAttributes:add_secure,ignore_secure", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy:cookie:ignore_secure", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ring_hash.hash_policy.cookie:ConflictingObjectAttributes:ignore_samesite,samesite_lax", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy:cookie:samesite_lax", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ring_hash.hash_policy.cookie:ConflictingObjectAttributes:samesite_lax,samesite_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy:cookie:samesite_lax", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ring_hash.hash_policy.cookie:ConflictingObjectAttributes:samesite_lax,samesite_strict", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy:cookie:samesite_lax", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ring_hash.hash_policy.cookie:ConflictingObjectAttributes:ignore_samesite,samesite_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy:cookie:samesite_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ring_hash.hash_policy.cookie:ConflictingObjectAttributes:samesite_lax,samesite_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy:cookie:samesite_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ring_hash.hash_policy.cookie:ConflictingObjectAttributes:samesite_none,samesite_strict", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy:cookie:samesite_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ring_hash.hash_policy.cookie:ConflictingObjectAttributes:ignore_samesite,samesite_strict", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy:cookie:samesite_strict", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ring_hash.hash_policy.cookie:ConflictingObjectAttributes:samesite_lax,samesite_strict", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy:cookie:samesite_strict", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ring_hash.hash_policy.cookie:ConflictingObjectAttributes:samesite_none,samesite_strict", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy:cookie:samesite_strict", "type": "conflicts"}, {"anchor": "schema-ring_hash--hash_policy--cookie--name", "enforcement": "provider-schema", "group": "ring_hash.hash_policy.cookie:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy:cookie", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["ring_hash", "hash_policy", "cookie"], "schema_version": 1, "sections": [{"aliases": ["add httponly"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy:cookie:add_httponly", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ring_hash", "hash_policy", "cookie", "add_httponly"], "syntax": "attribute", "type": "object"}, {"aliases": ["add secure"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy:cookie:add_secure", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ring_hash", "hash_policy", "cookie", "add_secure"], "syntax": "attribute", "type": "object"}, {"aliases": ["ignore httponly"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy:cookie:ignore_httponly", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ring_hash", "hash_policy", "cookie", "ignore_httponly"], "syntax": "attribute", "type": "object"}, {"aliases": ["ignore samesite"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy:cookie:ignore_samesite", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ring_hash", "hash_policy", "cookie", "ignore_samesite"], "syntax": "attribute", "type": "object"}, {"aliases": ["ignore secure"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy:cookie:ignore_secure", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ring_hash", "hash_policy", "cookie", "ignore_secure"], "syntax": "attribute", "type": "object"}, {"aliases": ["name"], "anchor": "schema-ring_hash--hash_policy--cookie--name", "description": "The name of the cookie that will be used to obtain the hash key. If the cookie is not present and TTL below is not set, no hash will be produced.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy:cookie", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ring_hash", "hash_policy", "cookie", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["path"], "anchor": "schema-ring_hash--hash_policy--cookie--path", "description": "The name of the path for the cookie. If no path is specified here, no path will be set for the cookie.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy:cookie", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ring_hash", "hash_policy", "cookie", "path"], "syntax": "attribute", "type": "string"}, {"aliases": ["samesite lax"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy:cookie:samesite_lax", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ring_hash", "hash_policy", "cookie", "samesite_lax"], "syntax": "attribute", "type": "object"}, {"aliases": ["samesite none"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy:cookie:samesite_none", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ring_hash", "hash_policy", "cookie", "samesite_none"], "syntax": "attribute", "type": "object"}, {"aliases": ["samesite strict"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy:cookie:samesite_strict", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ring_hash", "hash_policy", "cookie", "samesite_strict"], "syntax": "attribute", "type": "object"}, {"aliases": ["ttl"], "anchor": "schema-ring_hash--hash_policy--cookie--ttl", "description": "If specified, a cookie with the TTL will be generated if the cookie is not present. If the TTL is present and zero, the generated cookie will be a session cookie. TTL value is in milliseconds.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy:cookie", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ring_hash", "hash_policy", "cookie", "ttl"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/ring_hash/hash_policy/cookie/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Two types of cookie affinity: 1. Passive. Takes a cookie that's present in the cookies header and hashes on its value. 2. Generated. Generates and sets a cookie with an expiration (TTL) on the first request from the client in its response to the client, based on the endpoint the request gets sent to. The client then", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ring_hash.hash_policy.cookie

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [ring_hash](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/ring_hash/)
- [ring_hash.hash_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/ring_hash/hash_policy/)
- ring_hash.hash_policy.cookie

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Two types of cookie affinity: 1. Passive. Takes a cookie that's present in the cookies header and
hashes on its value. 2. Generated. Generates and sets a cookie with an expiration (TTL) on the first
request from the client in its response to the client, based on the endpoint the request gets..

Upstream description:

Two types of cookie affinity:

&#8203;1. Passive. Takes a cookie that's present in the cookies header and hashes on its value.

&#8203;2. Generated. Generates and sets a cookie with an expiration (TTL) on the first request from
the client in its response to the client, based on the endpoint the request gets sent to. The client
then presents this on the next and all subsequent requests. The hash of this is sufficient to ensure
these requests GET sent to the same endpoint. The cookie is generated by hashing the source and
destination ports and addresses so that multiple independent HTTP2 streams on the same connection
will independently receive the same cookie, even if they arrive simultaneously.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name"),
  validators.ConflictingObjectAttributes("add_httponly",
    "ignore_httponly"),
  validators.ConflictingObjectAttributes("add_secure",
    "ignore_secure"),
  validators.ConflictingObjectAttributes("ignore_samesite",
    "samesite_lax"),
  validators.ConflictingObjectAttributes("ignore_samesite",
    "samesite_none"),
  validators.ConflictingObjectAttributes("ignore_samesite",
    "samesite_strict"),
  validators.ConflictingObjectAttributes("samesite_lax",
    "samesite_none"),
  validators.ConflictingObjectAttributes("samesite_lax",
    "samesite_strict"),
  validators.ConflictingObjectAttributes("samesite_none",
    "samesite_strict")}
```

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

Terraform syntax:

```terraform
cookie {
  # Configure direct properties listed below.
}
```

## Direct properties

- [add_httponly](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/ring_hash/hash_policy/cookie/add_httponly/): complete subsection reference.

- [add_secure](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/ring_hash/hash_policy/cookie/add_secure/): complete subsection reference.

- [ignore_httponly](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/ring_hash/hash_policy/cookie/ignore_httponly/): complete subsection reference.

- [ignore_samesite](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/ring_hash/hash_policy/cookie/ignore_samesite/): complete subsection reference.

- [ignore_secure](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/ring_hash/hash_policy/cookie/ignore_secure/): complete subsection reference.

<a id="schema-ring_hash--hash_policy--cookie--name"></a>

### name property

Type: `"string"`. Optional.

The name of the cookie that will be used to obtain the hash key. If the cookie is not present and
TTL below is not set, no hash will be produced.

Upstream description:

The name of the cookie that will be used to obtain the hash key. If the cookie is not present and
TTL below is not set, no hash will be produced.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="schema-ring_hash--hash_policy--cookie--path"></a>

### path property

Type: `"string"`. Optional.

The name of the path for the cookie. If no path is specified here, no path will be set for the
cookie.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [samesite_lax](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/ring_hash/hash_policy/cookie/samesite_lax/): complete subsection reference.

- [samesite_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/ring_hash/hash_policy/cookie/samesite_none/): complete subsection reference.

- [samesite_strict](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/ring_hash/hash_policy/cookie/samesite_strict/): complete subsection reference.

<a id="schema-ring_hash--hash_policy--cookie--ttl"></a>

### ttl property

Type: `"number"`. Optional.

If specified, a cookie with the TTL will be generated if the cookie is not present. If the TTL is
present and zero, the generated cookie will be a session cookie. TTL value is in milliseconds.

Upstream description:

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

## Next pages

- [ring_hash.hash_policy.cookie.add_httponly](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/ring_hash/hash_policy/cookie/add_httponly/)
- [ring_hash.hash_policy.cookie.add_secure](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/ring_hash/hash_policy/cookie/add_secure/)
- [ring_hash.hash_policy.cookie.ignore_httponly](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/ring_hash/hash_policy/cookie/ignore_httponly/)
- [ring_hash.hash_policy.cookie.ignore_samesite](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/ring_hash/hash_policy/cookie/ignore_samesite/)
- [ring_hash.hash_policy.cookie.ignore_secure](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/ring_hash/hash_policy/cookie/ignore_secure/)
- [ring_hash.hash_policy.cookie.samesite_lax](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/ring_hash/hash_policy/cookie/samesite_lax/)
- [ring_hash.hash_policy.cookie.samesite_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/ring_hash/hash_policy/cookie/samesite_none/)
- [ring_hash.hash_policy.cookie.samesite_strict](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/ring_hash/hash_policy/cookie/samesite_strict/)
- [ring_hash.hash_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/ring_hash/hash_policy/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
