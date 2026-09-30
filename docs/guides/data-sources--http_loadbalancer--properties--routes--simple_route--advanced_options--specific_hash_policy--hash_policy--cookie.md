---
page_title: "routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie"
subcategory: "Load Balancing"
description: "routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 9757, "body_sha256": "sha256:e75eaa3b8e503bacbf01899dbe1d0df3a73583ff59bd650a1154d3f438c00f96", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy:hash_policy:cookie", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy:hash_policy:cookie:add_httponly", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy:hash_policy:cookie:add_secure", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy:hash_policy:cookie:ignore_httponly", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy:hash_policy:cookie:ignore_samesite", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy:hash_policy:cookie:ignore_secure", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy:hash_policy:cookie:samesite_lax", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy:hash_policy:cookie:samesite_none", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy:hash_policy:cookie:samesite_strict"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy:hash_policy:cookie", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy:hash_policy", "path": "docs/guides/data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--specific_hash_policy--hash_policy--cookie.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes", "simple_route", "advanced_options", "specific_hash_policy", "hash_policy", "cookie"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/routes/simple_route/advanced_options/specific_hash_policy/hash_policy/cookie/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [routes](data-sources--http_loadbalancer--properties--routes.md)
- [routes.simple_route](data-sources--http_loadbalancer--properties--routes--simple_route.md)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options.md)
- [routes.simple_route.advanced_options.specific_hash_policy](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--specific_hash_policy.md)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--specific_hash_policy--hash_policy.md)
- routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie

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

- [add_httponly](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--specific_hash_policy--hash_policy--cookie--add_httponly.md): complete subsection reference.

- [add_secure](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--specific_hash_policy--hash_policy--cookie--add_secure.md): complete subsection reference.

- [ignore_httponly](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--specific_hash_policy--hash_policy--cookie--ignore_httponly.md): complete subsection reference.

- [ignore_samesite](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--specific_hash_policy--hash_policy--cookie--ignore_samesite.md): complete subsection reference.

- [ignore_secure](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--specific_hash_policy--hash_policy--cookie--ignore_secure.md): complete subsection reference.

<a id="schema-routes--simple_route--advanced_options--specific_hash_policy--hash_policy--cookie--name"></a>

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

<a id="schema-routes--simple_route--advanced_options--specific_hash_policy--hash_policy--cookie--path"></a>

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

- [samesite_lax](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--specific_hash_policy--hash_policy--cookie--samesite_lax.md): complete subsection reference.

- [samesite_none](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--specific_hash_policy--hash_policy--cookie--samesite_none.md): complete subsection reference.

- [samesite_strict](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--specific_hash_policy--hash_policy--cookie--samesite_strict.md): complete subsection reference.

<a id="schema-routes--simple_route--advanced_options--specific_hash_policy--hash_policy--cookie--ttl"></a>

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

- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.add_httponly](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--specific_hash_policy--hash_policy--cookie--add_httponly.md)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.add_secure](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--specific_hash_policy--hash_policy--cookie--add_secure.md)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.ignore_httponly](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--specific_hash_policy--hash_policy--cookie--ignore_httponly.md)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.ignore_samesite](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--specific_hash_policy--hash_policy--cookie--ignore_samesite.md)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.ignore_secure](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--specific_hash_policy--hash_policy--cookie--ignore_secure.md)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.samesite_lax](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--specific_hash_policy--hash_policy--cookie--samesite_lax.md)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.samesite_none](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--specific_hash_policy--hash_policy--cookie--samesite_none.md)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.samesite_strict](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--specific_hash_policy--hash_policy--cookie--samesite_strict.md)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--specific_hash_policy--hash_policy.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
