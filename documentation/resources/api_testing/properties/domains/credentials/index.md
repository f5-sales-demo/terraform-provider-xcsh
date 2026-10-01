---
page_title: "domains.credentials"
subcategory: ""
description: "domains.credentials for xcsh_api_testing."
xcsh_docs: {"aliases": [], "body_bytes": 5004, "body_sha256": "sha256:80bff73a2159acbdb5940e6f4260af9a081289f82d8ca48f370e1f2e1ee24601", "child_ids": ["xcsh-docs:resources:api_testing:properties:domains:credentials:admin", "xcsh-docs:resources:api_testing:properties:domains:credentials:api_key", "xcsh-docs:resources:api_testing:properties:domains:credentials:basic_auth", "xcsh-docs:resources:api_testing:properties:domains:credentials:bearer_token", "xcsh-docs:resources:api_testing:properties:domains:credentials:login_endpoint", "xcsh-docs:resources:api_testing:properties:domains:credentials:standard"], "collection_id": "xcsh-docs:resources:api_testing:collection", "completeness": "complete", "id": "xcsh-docs:resources:api_testing:properties:domains:credentials", "parent_id": "xcsh-docs:resources:api_testing:properties:domains", "path": "documentation/resources/api_testing/properties/domains/credentials/index.md", "provider_name": "api_testing", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["domains", "credentials"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_testing/properties/domains/credentials/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "domains.credentials for xcsh_api_testing.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["api_testingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# domains.credentials

Breadcrumbs:

- [xcsh_api_testing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/)
- [domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/)
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

- [admin](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/credentials/admin/): complete subsection reference.

- [api_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/credentials/api_key/): complete subsection reference.

- [basic_auth](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/credentials/basic_auth/): complete subsection reference.

- [bearer_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/credentials/bearer_token/): complete subsection reference.

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

- [login_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/credentials/login_endpoint/): complete subsection reference.

- [standard](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/credentials/standard/): complete subsection reference.

## Next pages

- [domains.credentials.admin](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/credentials/admin/)
- [domains.credentials.api_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/credentials/api_key/)
- [domains.credentials.basic_auth](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/credentials/basic_auth/)
- [domains.credentials.bearer_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/credentials/bearer_token/)
- [domains.credentials.login_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/credentials/login_endpoint/)
- [domains.credentials.standard](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/credentials/standard/)
- [domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/)
- [xcsh_api_testing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/)
