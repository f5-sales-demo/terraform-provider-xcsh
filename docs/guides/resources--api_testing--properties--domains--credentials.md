---
page_title: "domains.credentials"
subcategory: ""
description: "domains.credentials for xcsh_api_testing."
xcsh_docs: {"aliases": [], "body_bytes": 4060, "body_sha256": "sha256:6a9b0eb2e6a1687600f1dd4787841dab5eeb76867aebac61ce1f4effd08c5bc6", "canonical_id": "xcsh-docs:resources:api_testing:properties:domains:credentials", "child_ids": ["xcsh-docs:resources:api_testing:properties:domains:credentials:admin", "xcsh-docs:resources:api_testing:properties:domains:credentials:api_key", "xcsh-docs:resources:api_testing:properties:domains:credentials:basic_auth", "xcsh-docs:resources:api_testing:properties:domains:credentials:bearer_token", "xcsh-docs:resources:api_testing:properties:domains:credentials:login_endpoint", "xcsh-docs:resources:api_testing:properties:domains:credentials:standard"], "collection_id": "xcsh-docs:resources:api_testing:collection", "completeness": "complete", "id": "xcsh-docs:resources:api_testing:properties:domains:credentials", "parent_id": "xcsh-docs:resources:api_testing:properties:domains", "path": "docs/guides/resources--api_testing--properties--domains--credentials.md", "provider_name": "api_testing", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["domains", "credentials"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_testing/properties/domains/credentials/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "domains.credentials for xcsh_api_testing.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["api_testingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# domains.credentials

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md)
- [Property reference](resources--api_testing--reference.md)
- [domains](resources--api_testing--properties--domains.md)
- domains.credentials

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Add credentials for API testing to use in the selected environment.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("credential_name"),
  validators.ConflictingListObjectAttributes("admin",
    "standard"),
  validators.ConflictingListObjectAttributes("api_key",
    "basic_auth"),
  validators.ConflictingListObjectAttributes("api_key",
    "bearer_token"),
  validators.ConflictingListObjectAttributes("api_key",
    "login_endpoint"),
  validators.ConflictingListObjectAttributes("basic_auth",
    "bearer_token"),
  validators.ConflictingListObjectAttributes("basic_auth",
    "login_endpoint"),
  validators.ConflictingListObjectAttributes("bearer_token",
    "login_endpoint")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

Terraform syntax:

```terraform
credentials {
  # Configure direct properties listed below.
}
```

## Direct properties

- [admin](resources--api_testing--properties--domains--credentials--admin.md): complete subsection reference.

- [api_key](resources--api_testing--properties--domains--credentials--api_key.md): complete subsection reference.

- [basic_auth](resources--api_testing--properties--domains--credentials--basic_auth.md): complete subsection reference.

- [bearer_token](resources--api_testing--properties--domains--credentials--bearer_token.md): complete subsection reference.

<a id="schema-domains--credentials--credential_name"></a>

### credential_name property

Type: `"string"`. Optional.

Enter a unique name for the credentials used in API testing.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [login_endpoint](resources--api_testing--properties--domains--credentials--login_endpoint.md): complete subsection reference.

- [standard](resources--api_testing--properties--domains--credentials--standard.md): complete subsection reference.

## Next pages

- [domains.credentials.admin](resources--api_testing--properties--domains--credentials--admin.md)
- [domains.credentials.api_key](resources--api_testing--properties--domains--credentials--api_key.md)
- [domains.credentials.basic_auth](resources--api_testing--properties--domains--credentials--basic_auth.md)
- [domains.credentials.bearer_token](resources--api_testing--properties--domains--credentials--bearer_token.md)
- [domains.credentials.login_endpoint](resources--api_testing--properties--domains--credentials--login_endpoint.md)
- [domains.credentials.standard](resources--api_testing--properties--domains--credentials--standard.md)
- [domains](resources--api_testing--properties--domains.md)
- [xcsh_api_testing](../resources/api_testing.md)
