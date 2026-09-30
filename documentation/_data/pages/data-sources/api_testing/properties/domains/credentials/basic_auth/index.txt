---
page_title: "domains.credentials.basic_auth"
subcategory: ""
description: "domains.credentials.basic_auth for xcsh_api_testing."
xcsh_docs: {"aliases": [], "body_bytes": 2558, "body_sha256": "sha256:a1f01f911c17ce887c7118bde330639cc52b5b98f2962aa7b20238e60fee0747", "child_ids": ["xcsh-docs:data-sources:api_testing:properties:domains:credentials:basic_auth:password"], "collection_id": "xcsh-docs:data-sources:api_testing:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:api_testing:properties:domains:credentials:basic_auth", "parent_id": "xcsh-docs:data-sources:api_testing:properties:domains:credentials", "path": "documentation/data-sources/api_testing/properties/domains/credentials/basic_auth/index.md", "provider_name": "api_testing", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["domains", "credentials", "basic_auth"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/api_testing/properties/domains/credentials/basic_auth/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "domains.credentials.basic_auth for xcsh_api_testing.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["api_testingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# domains.credentials.basic_auth

Breadcrumbs:

- [xcsh_api_testing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/properties/)
- [domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/properties/domains/)
- [domains.credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/properties/domains/credentials/)
- domains.credentials.basic_auth

<a id="section"></a>

Type: `"single"`. Computed.

Basic Authentication.

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

- [password](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/properties/domains/credentials/basic_auth/password/): complete subsection reference.

<a id="schema-domains--credentials--basic_auth--user"></a>

### user property

Type: `"string"`. Computed.

User. Configuration parameter for user

Upstream description:

Configuration parameter for user

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
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

## Next pages

- [domains.credentials.basic_auth.password](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/properties/domains/credentials/basic_auth/password/)
- [domains.credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/properties/domains/credentials/)
- [xcsh_api_testing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/)
