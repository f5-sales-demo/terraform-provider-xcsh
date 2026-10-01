---
page_title: "routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie"
subcategory: "Load Balancing"
description: "routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 10828, "body_sha256": "sha256:df2719fd0beed3201addcb0f1ce64fca3ce8f811c32ca5fd41bd0b5c40f0406c", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy:hash_policy:cookie", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy:hash_policy:cookie:add_httponly", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy:hash_policy:cookie:add_secure", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy:hash_policy:cookie:ignore_httponly", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy:hash_policy:cookie:ignore_samesite", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy:hash_policy:cookie:ignore_secure", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy:hash_policy:cookie:samesite_lax", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy:hash_policy:cookie:samesite_none", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy:hash_policy:cookie:samesite_strict"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy:hash_policy:cookie", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy:hash_policy", "path": "docs/guides/resources--http_loadbalancer--properties--routes--simple_route--advanced_options--specific_hash_policy--hash_policy--cookie.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes", "simple_route", "advanced_options", "specific_hash_policy", "hash_policy", "cookie"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/specific_hash_policy/hash_policy/cookie/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [routes](resources--http_loadbalancer--properties--routes.md)
- [routes.simple_route](resources--http_loadbalancer--properties--routes--simple_route.md)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--properties--routes--simple_route--advanced_options.md)
- [routes.simple_route.advanced_options.specific_hash_policy](resources--http_loadbalancer--properties--routes--simple_route--advanced_options--specific_hash_policy.md)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy](resources--http_loadbalancer--properties--routes--simple_route--advanced_options--specific_hash_policy--hash_policy.md)
- routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie

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

- [add_httponly](resources--http_loadbalancer--properties--routes--simple_route--advanced_options--specific_hash_policy--hash_policy--cookie--add_httponly.md): complete subsection reference.

- [add_secure](resources--http_loadbalancer--properties--routes--simple_route--advanced_options--specific_hash_policy--hash_policy--cookie--add_secure.md): complete subsection reference.

- [ignore_httponly](resources--http_loadbalancer--properties--routes--simple_route--advanced_options--specific_hash_policy--hash_policy--cookie--ignore_httponly.md): complete subsection reference.

- [ignore_samesite](resources--http_loadbalancer--properties--routes--simple_route--advanced_options--specific_hash_policy--hash_policy--cookie--ignore_samesite.md): complete subsection reference.

- [ignore_secure](resources--http_loadbalancer--properties--routes--simple_route--advanced_options--specific_hash_policy--hash_policy--cookie--ignore_secure.md): complete subsection reference.

<a id="schema-routes--simple_route--advanced_options--specific_hash_policy--hash_policy--cookie--name"></a>

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

<a id="schema-routes--simple_route--advanced_options--specific_hash_policy--hash_policy--cookie--path"></a>

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

- [samesite_lax](resources--http_loadbalancer--properties--routes--simple_route--advanced_options--specific_hash_policy--hash_policy--cookie--samesite_lax.md): complete subsection reference.

- [samesite_none](resources--http_loadbalancer--properties--routes--simple_route--advanced_options--specific_hash_policy--hash_policy--cookie--samesite_none.md): complete subsection reference.

- [samesite_strict](resources--http_loadbalancer--properties--routes--simple_route--advanced_options--specific_hash_policy--hash_policy--cookie--samesite_strict.md): complete subsection reference.

<a id="schema-routes--simple_route--advanced_options--specific_hash_policy--hash_policy--cookie--ttl"></a>

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

- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.add_httponly](resources--http_loadbalancer--properties--routes--simple_route--advanced_options--specific_hash_policy--hash_policy--cookie--add_httponly.md)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.add_secure](resources--http_loadbalancer--properties--routes--simple_route--advanced_options--specific_hash_policy--hash_policy--cookie--add_secure.md)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.ignore_httponly](resources--http_loadbalancer--properties--routes--simple_route--advanced_options--specific_hash_policy--hash_policy--cookie--ignore_httponly.md)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.ignore_samesite](resources--http_loadbalancer--properties--routes--simple_route--advanced_options--specific_hash_policy--hash_policy--cookie--ignore_samesite.md)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.ignore_secure](resources--http_loadbalancer--properties--routes--simple_route--advanced_options--specific_hash_policy--hash_policy--cookie--ignore_secure.md)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.samesite_lax](resources--http_loadbalancer--properties--routes--simple_route--advanced_options--specific_hash_policy--hash_policy--cookie--samesite_lax.md)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.samesite_none](resources--http_loadbalancer--properties--routes--simple_route--advanced_options--specific_hash_policy--hash_policy--cookie--samesite_none.md)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.samesite_strict](resources--http_loadbalancer--properties--routes--simple_route--advanced_options--specific_hash_policy--hash_policy--cookie--samesite_strict.md)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy](resources--http_loadbalancer--properties--routes--simple_route--advanced_options--specific_hash_policy--hash_policy.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
