---
page_title: "password_policy"
subcategory: ""
description: "Policy configuration for this feature."
xcsh_docs: {"aliases": ["password policy"], "body_bytes": 8033, "body_sha256": "sha256:ea748296a989a0a1005e0ff775f4b8f4031689c9dc65788ff8f40181b12fb8d6", "capabilities": ["administration"], "category": "administration", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:tenant_configuration:collection", "completeness": "complete", "id": "xcsh-docs:resources:tenant_configuration:properties:password_policy", "parent_id": "xcsh-docs:resources:tenant_configuration:reference", "path": "documentation/resources/tenant_configuration/properties/password_policy/index.md", "product": "distributed-cloud", "provider_name": "tenant_configuration", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-3102022230023020-1021013032122012-3020123221123311-3313322012121231-3031033300031300-3312331012000020-2031003212112110-2203300222010022", "registry_path": "docs/guides/resources--tenant_configuration--reference--group-001.md", "relationships": [{"anchor": "schema-password_policy--minimum_length", "enforcement": "provider-schema", "group": "password_policy:RequiredObjectAttributes:minimum_length", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:tenant_configuration:properties:password_policy", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["password_policy"], "schema_version": 1, "sections": [{"aliases": ["password policy digits"], "anchor": "schema-password_policy--digits", "description": "The number of digits required to be in the password string.", "document_id": "xcsh-docs:resources:tenant_configuration:properties:password_policy", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["password_policy", "digits"], "syntax": "attribute", "type": "number"}, {"aliases": ["password policy expire password"], "anchor": "schema-password_policy--expire_password", "description": "The number of days for which the password is valid. After the number of days has expired, the user is required to change their password.", "document_id": "xcsh-docs:resources:tenant_configuration:properties:password_policy", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["password_policy", "expire_password"], "syntax": "attribute", "type": "number"}, {"aliases": ["password policy lowercase characters"], "anchor": "schema-password_policy--lowercase_characters", "description": "The number of lower case letters required to be in the password string.", "document_id": "xcsh-docs:resources:tenant_configuration:properties:password_policy", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["password_policy", "lowercase_characters"], "syntax": "attribute", "type": "number"}, {"aliases": ["password policy minimum length"], "anchor": "schema-password_policy--minimum_length", "description": "Minimum length of password.", "document_id": "xcsh-docs:resources:tenant_configuration:properties:password_policy", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["password_policy", "minimum_length"], "syntax": "attribute", "type": "number"}, {"aliases": ["password policy not recently used"], "anchor": "schema-password_policy--not_recently_used", "description": "This policy is used to restrict user from using previously used passwords. Number that's set determines number of last passwords which user cannot use as new password.", "document_id": "xcsh-docs:resources:tenant_configuration:properties:password_policy", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["password_policy", "not_recently_used"], "syntax": "attribute", "type": "number"}, {"aliases": ["password policy not username"], "anchor": "schema-password_policy--not_username", "description": "When set, the password is not allowed to be the same as the username.", "document_id": "xcsh-docs:resources:tenant_configuration:properties:password_policy", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["password_policy", "not_username"], "syntax": "attribute", "type": "bool"}, {"aliases": ["password policy special characters"], "anchor": "schema-password_policy--special_characters", "description": "The number of special characters like '?!#%$' required to be in the password string.", "document_id": "xcsh-docs:resources:tenant_configuration:properties:password_policy", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["password_policy", "special_characters"], "syntax": "attribute", "type": "number"}, {"aliases": ["password policy uppercase characters"], "anchor": "schema-password_policy--uppercase_characters", "description": "The number of upper case letters required to be in the password string.", "document_id": "xcsh-docs:resources:tenant_configuration:properties:password_policy", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["password_policy", "uppercase_characters"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tenant_configuration/properties/password_policy/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Policy configuration for this feature.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["tenant_configurationCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# password_policy

Breadcrumbs:

- [xcsh_tenant_configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/)
- password_policy

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Policy configuration for this feature.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("minimum_length")}
```

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

Terraform syntax:

```terraform
password_policy {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-password_policy--digits"></a>

### digits property

Type: `"number"`. Optional.

The number of digits required to be in the password string.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

<a id="schema-password_policy--expire_password"></a>

### expire_password property

Type: `"number"`. Optional.

The number of days for which the password is valid. After the number of days has expired, the user
is required to change their password.

Upstream description:

The number of days for which the password is valid. After the number of days has expired, the user
is required to change their password.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(1080),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1080,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "1080"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "1080"
  }
}
```

<a id="schema-password_policy--lowercase_characters"></a>

### lowercase_characters property

Type: `"number"`. Optional.

The number of lower case letters required to be in the password string.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

<a id="schema-password_policy--minimum_length"></a>

### minimum_length property

Type: `"number"`. Optional.

Minimum Length. Minimum length of password.

Upstream description:

Minimum length of password.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtLeast(7),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 7
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "7"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "7"
  }
}
```

<a id="schema-password_policy--not_recently_used"></a>

### not_recently_used property

Type: `"number"`. Optional.

Policy is used to restrict user from using previously used passwords. Number that's set determines
number of last passwords which user cannot use as new password.

Upstream description:

This policy is used to restrict user from using previously used passwords. Number that's set
determines number of last passwords which user cannot use as new password.

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

<a id="schema-password_policy--not_username"></a>

### not_username property

Type: `"bool"`. Optional.

When set, the password is not allowed to be the same as the username.

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

<a id="schema-password_policy--special_characters"></a>

### special_characters property

Type: `"number"`. Optional.

The number of special characters like '?!\#%$' required to be in the password string.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

<a id="schema-password_policy--uppercase_characters"></a>

### uppercase_characters property

Type: `"number"`. Optional.

The number of upper case letters required to be in the password string.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/)
- [xcsh_tenant_configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/)
