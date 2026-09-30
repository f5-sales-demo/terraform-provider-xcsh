---
page_title: "tenant_details"
subcategory: ""
description: "tenant_details for xcsh_tenant_configuration."
xcsh_docs: {"aliases": [], "body_bytes": 2207, "body_sha256": "sha256:1b7f5c71bdadf3033dcf5835545aa7c5736f87934c3a899b94b3cac0f010b729", "canonical_id": "xcsh-docs:resources:tenant_configuration:properties:tenant_details", "child_ids": [], "collection_id": "xcsh-docs:resources:tenant_configuration:collection", "completeness": "complete", "id": "xcsh-docs:resources:tenant_configuration:properties:tenant_details", "parent_id": "xcsh-docs:resources:tenant_configuration:reference", "path": "docs/guides/resources--tenant_configuration--properties--tenant_details.md", "provider_name": "tenant_configuration", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tenant_details"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tenant_configuration/properties/tenant_details/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tenant_details for xcsh_tenant_configuration.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tenant_configurationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# tenant_details

Breadcrumbs:

- [xcsh_tenant_configuration](../resources/tenant_configuration.md)
- [Property reference](resources--tenant_configuration--reference.md)
- tenant_details

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

BasicConfiguration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("display_name")}
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
tenant_details {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-tenant_details--display_name"></a>

### display_name property

Type: `"string"`. Optional.

Changes the tenant name displayed during login without affecting your company’s domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[^<>&\\\"]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.pattern": "^[^<>&\\\"]+$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.pattern": "^[^<>&\\\"]+$"
  }
}
```

## Next pages

- [Property reference](resources--tenant_configuration--reference.md)
- [xcsh_tenant_configuration](../resources/tenant_configuration.md)
