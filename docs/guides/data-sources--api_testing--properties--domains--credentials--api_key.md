---
page_title: "domains.credentials.api_key"
subcategory: ""
description: "domains.credentials.api_key for xcsh_api_testing."
xcsh_docs: {"aliases": [], "body_bytes": 2012, "body_sha256": "sha256:ad51f34eb1caa9d87f573a6ea88b4636dd58fc3035400fde3cf8a50401fcf404", "canonical_id": "xcsh-docs:data-sources:api_testing:properties:domains:credentials:api_key", "child_ids": ["xcsh-docs:data-sources:api_testing:properties:domains:credentials:api_key:value"], "collection_id": "xcsh-docs:data-sources:api_testing:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:api_testing:properties:domains:credentials:api_key", "parent_id": "xcsh-docs:data-sources:api_testing:properties:domains:credentials", "path": "docs/guides/data-sources--api_testing--properties--domains--credentials--api_key.md", "provider_name": "api_testing", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["domains", "credentials", "api_key"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/api_testing/properties/domains/credentials/api_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "domains.credentials.api_key for xcsh_api_testing.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["api_testingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# domains.credentials.api_key

Breadcrumbs:

- [xcsh_api_testing](../data-sources/api_testing.md)
- [Property reference](data-sources--api_testing--reference.md)
- [domains](data-sources--api_testing--properties--domains.md)
- [domains.credentials](data-sources--api_testing--properties--domains--credentials.md)
- domains.credentials.api_key

<a id="section"></a>

Type: `"single"`. Computed.

API Key

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

<a id="schema-domains--credentials--api_key--key"></a>

### key property

Type: `"string"`. Computed.

Key. Cryptographic key material

Upstream description:

Cryptographic key material

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [value](data-sources--api_testing--properties--domains--credentials--api_key--value.md): complete subsection reference.

## Next pages

- [domains.credentials.api_key.value](data-sources--api_testing--properties--domains--credentials--api_key--value.md)
- [domains.credentials](data-sources--api_testing--properties--domains--credentials.md)
- [xcsh_api_testing](../data-sources/api_testing.md)
