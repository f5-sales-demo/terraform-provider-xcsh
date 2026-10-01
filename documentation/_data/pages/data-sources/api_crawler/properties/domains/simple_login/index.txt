---
page_title: "domains.simple_login"
subcategory: ""
description: "domains.simple_login for xcsh_api_crawler."
xcsh_docs: {"aliases": [], "body_bytes": 2331, "body_sha256": "sha256:c5ca3bd8ff327925224ca1d204c3ad7d0b66cafd3057352870ce5e9776809a28", "child_ids": ["xcsh-docs:data-sources:api_crawler:properties:domains:simple_login:password"], "collection_id": "xcsh-docs:data-sources:api_crawler:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:api_crawler:properties:domains:simple_login", "parent_id": "xcsh-docs:data-sources:api_crawler:properties:domains", "path": "documentation/data-sources/api_crawler/properties/domains/simple_login/index.md", "provider_name": "api_crawler", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["domains", "simple_login"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/api_crawler/properties/domains/simple_login/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "domains.simple_login for xcsh_api_crawler.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["api_crawlerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# domains.simple_login

Breadcrumbs:

- [xcsh_api_crawler](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_crawler/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_crawler/properties/)
- [domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_crawler/properties/domains/)
- domains.simple_login

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for simple login.

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

- [password](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_crawler/properties/domains/simple_login/password/): complete subsection reference.

<a id="schema-domains--simple_login--user"></a>

### user property

Type: `"string"`. Computed.

Enter the username to assign credentials for the selected domain to crawl.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

## Next pages

- [domains.simple_login.password](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_crawler/properties/domains/simple_login/password/)
- [domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_crawler/properties/domains/)
- [xcsh_api_crawler](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_crawler/)
