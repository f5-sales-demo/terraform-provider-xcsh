---
page_title: "routes.response_cookies_to_add"
subcategory: ""
description: "routes.response_cookies_to_add for xcsh_route."
xcsh_docs: {"aliases": [], "body_bytes": 14999, "body_sha256": "sha256:13d8c72828b1a8ad80482c06927c60bd2a996ed7ce156f222cef4186cfb74227", "canonical_id": "xcsh-docs:resources:route:properties:routes:response_cookies_to_add", "child_ids": ["xcsh-docs:resources:route:properties:routes:response_cookies_to_add:add_httponly", "xcsh-docs:resources:route:properties:routes:response_cookies_to_add:add_partitioned", "xcsh-docs:resources:route:properties:routes:response_cookies_to_add:add_secure", "xcsh-docs:resources:route:properties:routes:response_cookies_to_add:ignore_domain", "xcsh-docs:resources:route:properties:routes:response_cookies_to_add:ignore_expiry", "xcsh-docs:resources:route:properties:routes:response_cookies_to_add:ignore_httponly", "xcsh-docs:resources:route:properties:routes:response_cookies_to_add:ignore_max_age", "xcsh-docs:resources:route:properties:routes:response_cookies_to_add:ignore_partitioned", "xcsh-docs:resources:route:properties:routes:response_cookies_to_add:ignore_path", "xcsh-docs:resources:route:properties:routes:response_cookies_to_add:ignore_samesite", "xcsh-docs:resources:route:properties:routes:response_cookies_to_add:ignore_secure", "xcsh-docs:resources:route:properties:routes:response_cookies_to_add:ignore_value", "xcsh-docs:resources:route:properties:routes:response_cookies_to_add:samesite_lax", "xcsh-docs:resources:route:properties:routes:response_cookies_to_add:samesite_none", "xcsh-docs:resources:route:properties:routes:response_cookies_to_add:samesite_strict", "xcsh-docs:resources:route:properties:routes:response_cookies_to_add:secret_value"], "collection_id": "xcsh-docs:resources:route:collection", "completeness": "complete", "id": "xcsh-docs:resources:route:properties:routes:response_cookies_to_add", "parent_id": "xcsh-docs:resources:route:properties:routes", "path": "docs/guides/resources--route--properties--routes--response_cookies_to_add.md", "provider_name": "route", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes", "response_cookies_to_add"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/route/properties/routes/response_cookies_to_add/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.response_cookies_to_add for xcsh_route.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["routeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.response_cookies_to_add

Breadcrumbs:

- [xcsh_route](../resources/route.md)
- [Property reference](resources--route--reference.md)
- [routes](resources--route--properties--routes.md)
- routes.response_cookies_to_add

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Cookies are name-value pairs along with optional attribute parameters to be added to HTTP response
being sent towards downstream.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("add_domain",
    "ignore_domain"),
  validators.ConflictingListObjectAttributes("add_expiry",
    "ignore_expiry"),
  validators.ConflictingListObjectAttributes("add_httponly",
    "ignore_httponly"),
  validators.ConflictingListObjectAttributes("add_partitioned",
    "ignore_partitioned"),
  validators.ConflictingListObjectAttributes("add_path",
    "ignore_path"),
  validators.ConflictingListObjectAttributes("add_secure",
    "ignore_secure"),
  validators.ConflictingListObjectAttributes("ignore_max_age",
    "max_age_value"),
  validators.ConflictingListObjectAttributes("ignore_samesite",
    "samesite_lax"),
  validators.ConflictingListObjectAttributes("ignore_samesite",
    "samesite_none"),
  validators.ConflictingListObjectAttributes("ignore_samesite",
    "samesite_strict"),
  validators.ConflictingListObjectAttributes("ignore_value",
    "secret_value"),
  validators.ConflictingListObjectAttributes("ignore_value",
    "value"),
  validators.ConflictingListObjectAttributes("samesite_lax",
    "samesite_none"),
  validators.ConflictingListObjectAttributes("samesite_lax",
    "samesite_strict"),
  validators.ConflictingListObjectAttributes("samesite_none",
    "samesite_strict"),
  validators.ConflictingListObjectAttributes("secret_value",
    "value")}
```

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

Terraform syntax:

```terraform
response_cookies_to_add {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-routes--response_cookies_to_add--add_domain"></a>

### add_domain property

Type: `"string"`. Optional.

Exclusive with \[ignore\_domain\] Add domain attribute.

Upstream description:

Exclusive with \[ignore\_domain\] Add domain attribute.

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

<a id="schema-routes--response_cookies_to_add--add_expiry"></a>

### add_expiry property

Type: `"string"`. Optional.

Exclusive with \[ignore\_expiry\] Add expiry attribute.

Upstream description:

Exclusive with \[ignore\_expiry\] Add expiry attribute.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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

- [add_httponly](resources--route--properties--routes--response_cookies_to_add--add_httponly.md): complete subsection reference.

- [add_partitioned](resources--route--properties--routes--response_cookies_to_add--add_partitioned.md): complete subsection reference.

<a id="schema-routes--response_cookies_to_add--add_path"></a>

### add_path property

Type: `"string"`. Optional.

Exclusive with \[ignore\_path\] Add path attribute.

Upstream description:

Exclusive with \[ignore\_path\] Add path attribute.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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

- [add_secure](resources--route--properties--routes--response_cookies_to_add--add_secure.md): complete subsection reference.

- [ignore_domain](resources--route--properties--routes--response_cookies_to_add--ignore_domain.md): complete subsection reference.

- [ignore_expiry](resources--route--properties--routes--response_cookies_to_add--ignore_expiry.md): complete subsection reference.

- [ignore_httponly](resources--route--properties--routes--response_cookies_to_add--ignore_httponly.md): complete subsection reference.

- [ignore_max_age](resources--route--properties--routes--response_cookies_to_add--ignore_max_age.md): complete subsection reference.

- [ignore_partitioned](resources--route--properties--routes--response_cookies_to_add--ignore_partitioned.md): complete subsection reference.

- [ignore_path](resources--route--properties--routes--response_cookies_to_add--ignore_path.md): complete subsection reference.

- [ignore_samesite](resources--route--properties--routes--response_cookies_to_add--ignore_samesite.md): complete subsection reference.

- [ignore_secure](resources--route--properties--routes--response_cookies_to_add--ignore_secure.md): complete subsection reference.

- [ignore_value](resources--route--properties--routes--response_cookies_to_add--ignore_value.md): complete subsection reference.

<a id="schema-routes--response_cookies_to_add--max_age_value"></a>

### max_age_value property

Type: `"number"`. Optional.

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Upstream description:

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(34560000),
}
```

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

<a id="schema-routes--response_cookies_to_add--name"></a>

### name property

Type: `"string"`. Optional.

Name. Name of the cookie in Cookie header.

Upstream description:

Name of the cookie in Cookie header.

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

<a id="schema-routes--response_cookies_to_add--overwrite"></a>

### overwrite property

Type: `"bool"`. Optional.

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

- [samesite_lax](resources--route--properties--routes--response_cookies_to_add--samesite_lax.md): complete subsection reference.

- [samesite_none](resources--route--properties--routes--response_cookies_to_add--samesite_none.md): complete subsection reference.

- [samesite_strict](resources--route--properties--routes--response_cookies_to_add--samesite_strict.md): complete subsection reference.

- [secret_value](resources--route--properties--routes--response_cookies_to_add--secret_value.md): complete subsection reference.

<a id="schema-routes--response_cookies_to_add--value"></a>

### value property

Type: `"string"`. Optional.

Exclusive with \[ignore\_value secret\_value\] Value of the Cookie header.

Upstream description:

Exclusive with \[ignore\_value secret\_value\] Value of the Cookie header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8096),
}
```

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

- [routes.response_cookies_to_add.add_httponly](resources--route--properties--routes--response_cookies_to_add--add_httponly.md)
- [routes.response_cookies_to_add.add_partitioned](resources--route--properties--routes--response_cookies_to_add--add_partitioned.md)
- [routes.response_cookies_to_add.add_secure](resources--route--properties--routes--response_cookies_to_add--add_secure.md)
- [routes.response_cookies_to_add.ignore_domain](resources--route--properties--routes--response_cookies_to_add--ignore_domain.md)
- [routes.response_cookies_to_add.ignore_expiry](resources--route--properties--routes--response_cookies_to_add--ignore_expiry.md)
- [routes.response_cookies_to_add.ignore_httponly](resources--route--properties--routes--response_cookies_to_add--ignore_httponly.md)
- [routes.response_cookies_to_add.ignore_max_age](resources--route--properties--routes--response_cookies_to_add--ignore_max_age.md)
- [routes.response_cookies_to_add.ignore_partitioned](resources--route--properties--routes--response_cookies_to_add--ignore_partitioned.md)
- [routes.response_cookies_to_add.ignore_path](resources--route--properties--routes--response_cookies_to_add--ignore_path.md)
- [routes.response_cookies_to_add.ignore_samesite](resources--route--properties--routes--response_cookies_to_add--ignore_samesite.md)
- [routes.response_cookies_to_add.ignore_secure](resources--route--properties--routes--response_cookies_to_add--ignore_secure.md)
- [routes.response_cookies_to_add.ignore_value](resources--route--properties--routes--response_cookies_to_add--ignore_value.md)
- [routes.response_cookies_to_add.samesite_lax](resources--route--properties--routes--response_cookies_to_add--samesite_lax.md)
- [routes.response_cookies_to_add.samesite_none](resources--route--properties--routes--response_cookies_to_add--samesite_none.md)
- [routes.response_cookies_to_add.samesite_strict](resources--route--properties--routes--response_cookies_to_add--samesite_strict.md)
- [routes.response_cookies_to_add.secret_value](resources--route--properties--routes--response_cookies_to_add--secret_value.md)
- [routes](resources--route--properties--routes.md)
- [xcsh_route](../resources/route.md)
