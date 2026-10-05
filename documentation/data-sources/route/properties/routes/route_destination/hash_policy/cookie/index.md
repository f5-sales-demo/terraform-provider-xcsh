---
page_title: "routes.route_destination.hash_policy.cookie"
subcategory: ""
description: "Two types of cookie affinity: 1. Passive. Takes a cookie that's present in the cookies header and hashes on its value. 2. Generated. Generates and sets a cookie with an expiration (TTL) on the first request from the client in its response to the client, based on the endpoint the request gets sent to. The client then"
xcsh_docs: {"aliases": ["routes route destination hash policy cookie"], "body_bytes": 9240, "body_sha256": "sha256:370e8772f65e3d9992e21585f41419658bde96d6281894bb13652a3c9a01080d", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:route:properties:routes:route_destination:hash_policy:cookie:add_httponly", "xcsh-docs:data-sources:route:properties:routes:route_destination:hash_policy:cookie:add_secure", "xcsh-docs:data-sources:route:properties:routes:route_destination:hash_policy:cookie:ignore_httponly", "xcsh-docs:data-sources:route:properties:routes:route_destination:hash_policy:cookie:ignore_samesite", "xcsh-docs:data-sources:route:properties:routes:route_destination:hash_policy:cookie:ignore_secure", "xcsh-docs:data-sources:route:properties:routes:route_destination:hash_policy:cookie:samesite_lax", "xcsh-docs:data-sources:route:properties:routes:route_destination:hash_policy:cookie:samesite_none", "xcsh-docs:data-sources:route:properties:routes:route_destination:hash_policy:cookie:samesite_strict"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:route:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:route:properties:routes:route_destination:hash_policy:cookie", "parent_id": "xcsh-docs:data-sources:route:properties:routes:route_destination:hash_policy", "path": "documentation/data-sources/route/properties/routes/route_destination/hash_policy/cookie/index.md", "product": "distributed-cloud", "provider_name": "route", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-3300222200100202-2221333320332130-2220132213300033-3113023022203311-1103220320113033-3021320101300123-2103000333122303-1032032230202231", "registry_path": "docs/guides/data-sources--route--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "route_destination", "hash_policy", "cookie"], "schema_version": 1, "sections": [{"aliases": ["routes route destination hash policy cookie add httponly"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:route:properties:routes:route_destination:hash_policy:cookie:add_httponly", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "route_destination", "hash_policy", "cookie", "add_httponly"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes route destination hash policy cookie add secure"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:route:properties:routes:route_destination:hash_policy:cookie:add_secure", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "route_destination", "hash_policy", "cookie", "add_secure"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes route destination hash policy cookie ignore httponly"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:route:properties:routes:route_destination:hash_policy:cookie:ignore_httponly", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "route_destination", "hash_policy", "cookie", "ignore_httponly"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes route destination hash policy cookie ignore samesite"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:route:properties:routes:route_destination:hash_policy:cookie:ignore_samesite", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "route_destination", "hash_policy", "cookie", "ignore_samesite"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes route destination hash policy cookie ignore secure"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:route:properties:routes:route_destination:hash_policy:cookie:ignore_secure", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "route_destination", "hash_policy", "cookie", "ignore_secure"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes route destination hash policy cookie name"], "anchor": "schema-routes--route_destination--hash_policy--cookie--name", "description": "The name of the cookie that will be used to obtain the hash key. If the cookie is not present and TTL below is not set, no hash will be produced.", "document_id": "xcsh-docs:data-sources:route:properties:routes:route_destination:hash_policy:cookie", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "route_destination", "hash_policy", "cookie", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["routes route destination hash policy cookie path"], "anchor": "schema-routes--route_destination--hash_policy--cookie--path", "description": "The name of the path for the cookie. If no path is specified here, no path will be set for the cookie.", "document_id": "xcsh-docs:data-sources:route:properties:routes:route_destination:hash_policy:cookie", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "route_destination", "hash_policy", "cookie", "path"], "syntax": "attribute", "type": "string"}, {"aliases": ["routes route destination hash policy cookie samesite lax"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:route:properties:routes:route_destination:hash_policy:cookie:samesite_lax", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "route_destination", "hash_policy", "cookie", "samesite_lax"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes route destination hash policy cookie samesite none"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:route:properties:routes:route_destination:hash_policy:cookie:samesite_none", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "route_destination", "hash_policy", "cookie", "samesite_none"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes route destination hash policy cookie samesite strict"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:route:properties:routes:route_destination:hash_policy:cookie:samesite_strict", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "route_destination", "hash_policy", "cookie", "samesite_strict"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes route destination hash policy cookie ttl"], "anchor": "schema-routes--route_destination--hash_policy--cookie--ttl", "description": "If specified, a cookie with the TTL will be generated if the cookie is not present. If the TTL is present and zero, the generated cookie will be a session cookie. TTL value is in milliseconds.", "document_id": "xcsh-docs:data-sources:route:properties:routes:route_destination:hash_policy:cookie", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "route_destination", "hash_policy", "cookie", "ttl"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/route/properties/routes/route_destination/hash_policy/cookie/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Two types of cookie affinity: 1. Passive. Takes a cookie that's present in the cookies header and hashes on its value. 2. Generated. Generates and sets a cookie with an expiration (TTL) on the first request from the client in its response to the client, based on the endpoint the request gets sent to. The client then", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["routeCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.route_destination.hash_policy.cookie

Breadcrumbs:

- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/)
- [routes.route_destination](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/)
- [routes.route_destination.hash_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/hash_policy/)
- routes.route_destination.hash_policy.cookie

<a id="section"></a>

Type: `"single"`. Computed.

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

## Direct properties

- [add_httponly](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/hash_policy/cookie/add_httponly/): complete subsection reference.

- [add_secure](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/hash_policy/cookie/add_secure/): complete subsection reference.

- [ignore_httponly](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/hash_policy/cookie/ignore_httponly/): complete subsection reference.

- [ignore_samesite](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/hash_policy/cookie/ignore_samesite/): complete subsection reference.

- [ignore_secure](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/hash_policy/cookie/ignore_secure/): complete subsection reference.

<a id="schema-routes--route_destination--hash_policy--cookie--name"></a>

### name property

Type: `"string"`. Computed.

The name of the cookie that will be used to obtain the hash key. If the cookie is not present and
TTL below is not set, no hash will be produced.

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="schema-routes--route_destination--hash_policy--cookie--path"></a>

### path property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [samesite_lax](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/hash_policy/cookie/samesite_lax/): complete subsection reference.

- [samesite_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/hash_policy/cookie/samesite_none/): complete subsection reference.

- [samesite_strict](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/hash_policy/cookie/samesite_strict/): complete subsection reference.

<a id="schema-routes--route_destination--hash_policy--cookie--ttl"></a>

### ttl property

Type: `"number"`. Computed.

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

- [routes.route_destination.hash_policy.cookie.add_httponly](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/hash_policy/cookie/add_httponly/)
- [routes.route_destination.hash_policy.cookie.add_secure](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/hash_policy/cookie/add_secure/)
- [routes.route_destination.hash_policy.cookie.ignore_httponly](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/hash_policy/cookie/ignore_httponly/)
- [routes.route_destination.hash_policy.cookie.ignore_samesite](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/hash_policy/cookie/ignore_samesite/)
- [routes.route_destination.hash_policy.cookie.ignore_secure](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/hash_policy/cookie/ignore_secure/)
- [routes.route_destination.hash_policy.cookie.samesite_lax](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/hash_policy/cookie/samesite_lax/)
- [routes.route_destination.hash_policy.cookie.samesite_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/hash_policy/cookie/samesite_none/)
- [routes.route_destination.hash_policy.cookie.samesite_strict](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/hash_policy/cookie/samesite_strict/)
- [routes.route_destination.hash_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/hash_policy/)
- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/)
