---
page_title: "routes.simple_route.advanced_options.response_cookies_to_add"
subcategory: "Load Balancing"
description: "routes.simple_route.advanced_options.response_cookies_to_add for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 15548, "body_sha256": "sha256:565f720ebda4e166af0ffce6674bb588cd95bc95774bb06e429a9261b74bb652", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:response_cookies_to_add", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:response_cookies_to_add:add_httponly", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:response_cookies_to_add:add_partitioned", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:response_cookies_to_add:add_secure", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:response_cookies_to_add:ignore_domain", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:response_cookies_to_add:ignore_expiry", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:response_cookies_to_add:ignore_httponly", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:response_cookies_to_add:ignore_max_age", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:response_cookies_to_add:ignore_partitioned", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:response_cookies_to_add:ignore_path", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:response_cookies_to_add:ignore_samesite", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:response_cookies_to_add:ignore_secure", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:response_cookies_to_add:ignore_value", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:response_cookies_to_add:samesite_lax", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:response_cookies_to_add:samesite_none", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:response_cookies_to_add:samesite_strict", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:response_cookies_to_add:secret_value"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:response_cookies_to_add", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options", "path": "docs/guides/data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_cookies_to_add.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes", "simple_route", "advanced_options", "response_cookies_to_add"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/routes/simple_route/advanced_options/response_cookies_to_add/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.simple_route.advanced_options.response_cookies_to_add for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.simple_route.advanced_options.response_cookies_to_add

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [routes](data-sources--http_loadbalancer--properties--routes.md)
- [routes.simple_route](data-sources--http_loadbalancer--properties--routes--simple_route.md)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options.md)
- routes.simple_route.advanced_options.response_cookies_to_add

<a id="section"></a>

Type: `"list"`. Computed.

Cookies are name-value pairs along with optional attribute parameters to be added to HTTP response
being sent towards downstream. Cookies specified at this level are applied after cookies from
matched Route are applied.

Upstream description:

Cookies are name-value pairs along with optional attribute parameters to be added to HTTP response
being sent towards downstream. Cookies specified at this level are applied after cookies from
matched Route are applied.

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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

<a id="schema-routes--simple_route--advanced_options--response_cookies_to_add--add_domain"></a>

### add_domain property

Type: `"string"`. Computed.

Exclusive with \[ignore\_domain\] Add domain attribute.

Upstream description:

Exclusive with \[ignore\_domain\] Add domain attribute.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="schema-routes--simple_route--advanced_options--response_cookies_to_add--add_expiry"></a>

### add_expiry property

Type: `"string"`. Computed.

Exclusive with \[ignore\_expiry\] Add expiry attribute.

Upstream description:

Exclusive with \[ignore\_expiry\] Add expiry attribute.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [add_httponly](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_cookies_to_add--add_httponly.md): complete subsection reference.

- [add_partitioned](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_cookies_to_add--add_partitioned.md): complete subsection reference.

<a id="schema-routes--simple_route--advanced_options--response_cookies_to_add--add_path"></a>

### add_path property

Type: `"string"`. Computed.

Exclusive with \[ignore\_path\] Add path attribute.

Upstream description:

Exclusive with \[ignore\_path\] Add path attribute.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [add_secure](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_cookies_to_add--add_secure.md): complete subsection reference.

- [ignore_domain](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_cookies_to_add--ignore_domain.md): complete subsection reference.

- [ignore_expiry](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_cookies_to_add--ignore_expiry.md): complete subsection reference.

- [ignore_httponly](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_cookies_to_add--ignore_httponly.md): complete subsection reference.

- [ignore_max_age](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_cookies_to_add--ignore_max_age.md): complete subsection reference.

- [ignore_partitioned](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_cookies_to_add--ignore_partitioned.md): complete subsection reference.

- [ignore_path](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_cookies_to_add--ignore_path.md): complete subsection reference.

- [ignore_samesite](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_cookies_to_add--ignore_samesite.md): complete subsection reference.

- [ignore_secure](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_cookies_to_add--ignore_secure.md): complete subsection reference.

- [ignore_value](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_cookies_to_add--ignore_value.md): complete subsection reference.

<a id="schema-routes--simple_route--advanced_options--response_cookies_to_add--max_age_value"></a>

### max_age_value property

Type: `"number"`. Computed.

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Upstream description:

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 34560000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "34560000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "34560000"
  }
}
```

<a id="schema-routes--simple_route--advanced_options--response_cookies_to_add--name"></a>

### name property

Type: `"string"`. Computed.

Name. Name of the cookie in Cookie header.

Upstream description:

Name of the cookie in Cookie header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="schema-routes--simple_route--advanced_options--response_cookies_to_add--overwrite"></a>

### overwrite property

Type: `"bool"`. Computed.

Should the value be overwritten? If true, the value is overwritten to existing values. not
overwrite. Defaults to \`do\`.

Upstream description:

Should the value be overwritten? If true, the value is overwritten to existing values. Default value
is do not overwrite.

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

- [samesite_lax](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_cookies_to_add--samesite_lax.md): complete subsection reference.

- [samesite_none](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_cookies_to_add--samesite_none.md): complete subsection reference.

- [samesite_strict](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_cookies_to_add--samesite_strict.md): complete subsection reference.

- [secret_value](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_cookies_to_add--secret_value.md): complete subsection reference.

<a id="schema-routes--simple_route--advanced_options--response_cookies_to_add--value"></a>

### value property

Type: `"string"`. Computed.

Exclusive with \[ignore\_value secret\_value\] Value of the Cookie header.

Upstream description:

Exclusive with \[ignore\_value secret\_value\] Value of the Cookie header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

## Next pages

- [routes.simple_route.advanced_options.response_cookies_to_add.add_httponly](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_cookies_to_add--add_httponly.md)
- [routes.simple_route.advanced_options.response_cookies_to_add.add_partitioned](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_cookies_to_add--add_partitioned.md)
- [routes.simple_route.advanced_options.response_cookies_to_add.add_secure](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_cookies_to_add--add_secure.md)
- [routes.simple_route.advanced_options.response_cookies_to_add.ignore_domain](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_cookies_to_add--ignore_domain.md)
- [routes.simple_route.advanced_options.response_cookies_to_add.ignore_expiry](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_cookies_to_add--ignore_expiry.md)
- [routes.simple_route.advanced_options.response_cookies_to_add.ignore_httponly](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_cookies_to_add--ignore_httponly.md)
- [routes.simple_route.advanced_options.response_cookies_to_add.ignore_max_age](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_cookies_to_add--ignore_max_age.md)
- [routes.simple_route.advanced_options.response_cookies_to_add.ignore_partitioned](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_cookies_to_add--ignore_partitioned.md)
- [routes.simple_route.advanced_options.response_cookies_to_add.ignore_path](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_cookies_to_add--ignore_path.md)
- [routes.simple_route.advanced_options.response_cookies_to_add.ignore_samesite](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_cookies_to_add--ignore_samesite.md)
- [routes.simple_route.advanced_options.response_cookies_to_add.ignore_secure](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_cookies_to_add--ignore_secure.md)
- [routes.simple_route.advanced_options.response_cookies_to_add.ignore_value](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_cookies_to_add--ignore_value.md)
- [routes.simple_route.advanced_options.response_cookies_to_add.samesite_lax](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_cookies_to_add--samesite_lax.md)
- [routes.simple_route.advanced_options.response_cookies_to_add.samesite_none](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_cookies_to_add--samesite_none.md)
- [routes.simple_route.advanced_options.response_cookies_to_add.samesite_strict](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_cookies_to_add--samesite_strict.md)
- [routes.simple_route.advanced_options.response_cookies_to_add.secret_value](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_cookies_to_add--secret_value.md)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
