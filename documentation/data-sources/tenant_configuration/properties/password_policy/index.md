---
page_title: "password_policy"
subcategory: ""
description: "Policy configuration for this feature."
xcsh_docs: {"aliases": ["password policy"], "body_bytes": 6328, "body_sha256": "sha256:15e2378d41aecd0a67b9ac659461e8c0ba9945437fddd74b70b4cc6ef9ae19a8", "capabilities": ["administration"], "category": "administration", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:tenant_configuration:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tenant_configuration:properties:password_policy", "parent_id": "xcsh-docs:data-sources:tenant_configuration:reference", "path": "documentation/data-sources/tenant_configuration/properties/password_policy/index.md", "product": "distributed-cloud", "provider_name": "tenant_configuration", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-1001120333133230-3300303310031021-1113120310330110-0311101223100330-0011032012332230-0223330003003033-3212320310131032-0232103120223312", "registry_path": "docs/guides/data-sources--tenant_configuration--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["password_policy"], "schema_version": 1, "sections": [{"aliases": ["password policy digits"], "anchor": "schema-password_policy--digits", "description": "The number of digits required to be in the password string.", "document_id": "xcsh-docs:data-sources:tenant_configuration:properties:password_policy", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["password_policy", "digits"], "syntax": "attribute", "type": "number"}, {"aliases": ["password policy expire password"], "anchor": "schema-password_policy--expire_password", "description": "The number of days for which the password is valid. After the number of days has expired, the user is required to change their password.", "document_id": "xcsh-docs:data-sources:tenant_configuration:properties:password_policy", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["password_policy", "expire_password"], "syntax": "attribute", "type": "number"}, {"aliases": ["password policy lowercase characters"], "anchor": "schema-password_policy--lowercase_characters", "description": "The number of lower case letters required to be in the password string.", "document_id": "xcsh-docs:data-sources:tenant_configuration:properties:password_policy", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["password_policy", "lowercase_characters"], "syntax": "attribute", "type": "number"}, {"aliases": ["password policy minimum length"], "anchor": "schema-password_policy--minimum_length", "description": "Minimum length of password.", "document_id": "xcsh-docs:data-sources:tenant_configuration:properties:password_policy", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["password_policy", "minimum_length"], "syntax": "attribute", "type": "number"}, {"aliases": ["password policy not recently used"], "anchor": "schema-password_policy--not_recently_used", "description": "This policy is used to restrict user from using previously used passwords. Number that's set determines number of last passwords which user cannot use as new password.", "document_id": "xcsh-docs:data-sources:tenant_configuration:properties:password_policy", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["password_policy", "not_recently_used"], "syntax": "attribute", "type": "number"}, {"aliases": ["password policy not username"], "anchor": "schema-password_policy--not_username", "description": "When set, the password is not allowed to be the same as the username.", "document_id": "xcsh-docs:data-sources:tenant_configuration:properties:password_policy", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["password_policy", "not_username"], "syntax": "attribute", "type": "bool"}, {"aliases": ["password policy special characters"], "anchor": "schema-password_policy--special_characters", "description": "The number of special characters like '?!#%$' required to be in the password string.", "document_id": "xcsh-docs:data-sources:tenant_configuration:properties:password_policy", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["password_policy", "special_characters"], "syntax": "attribute", "type": "number"}, {"aliases": ["password policy uppercase characters"], "anchor": "schema-password_policy--uppercase_characters", "description": "The number of upper case letters required to be in the password string.", "document_id": "xcsh-docs:data-sources:tenant_configuration:properties:password_policy", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["password_policy", "uppercase_characters"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tenant_configuration/properties/password_policy/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Policy configuration for this feature.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["tenant_configurationCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# password_policy

Breadcrumbs:

- [xcsh_tenant_configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/)
- password_policy

<a id="section"></a>

Type: `"single"`. Computed.

Policy configuration for this feature.

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

<a id="schema-password_policy--digits"></a>

### digits property

Type: `"number"`. Computed.

The number of digits required to be in the password string.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

Type: `"number"`. Computed.

The number of days for which the password is valid. After the number of days has expired, the user
is required to change their password.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

Type: `"number"`. Computed.

The number of lower case letters required to be in the password string.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

Type: `"number"`. Computed.

Minimum Length. Minimum length of password.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

Type: `"number"`. Computed.

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

Type: `"bool"`. Computed.

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

Type: `"number"`. Computed.

The number of special characters like '?!\#%$' required to be in the password string.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

Type: `"number"`. Computed.

The number of upper case letters required to be in the password string.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
