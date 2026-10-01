---
page_title: "domains"
subcategory: ""
description: "domains for xcsh_api_crawler."
xcsh_docs: {"aliases": [], "body_bytes": 2533, "body_sha256": "sha256:0300388377244b14435686f507ef45c0b2be28a2b6653632fba5d786267a52a5", "canonical_id": "xcsh-docs:data-sources:api_crawler:properties:domains", "child_ids": ["xcsh-docs:data-sources:api_crawler:properties:domains:simple_login"], "collection_id": "xcsh-docs:data-sources:api_crawler:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:api_crawler:properties:domains", "parent_id": "xcsh-docs:data-sources:api_crawler:reference", "path": "docs/guides/data-sources--api_crawler--properties--domains.md", "provider_name": "api_crawler", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["domains"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/api_crawler/properties/domains/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "domains for xcsh_api_crawler.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["api_crawlerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# domains

Breadcrumbs:

- [xcsh_api_crawler](../data-sources/api_crawler.md)
- [Property reference](data-sources--api_crawler--reference.md)
- domains

<a id="section"></a>

Type: `"list"`. Computed.

API Crawler. API Crawler Configuration.

Upstream description:

API Crawler Configuration.

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
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

## Direct properties

<a id="schema-domains--domain"></a>

### domain property

Type: `"string"`. Computed.

Select the domain to execute API Crawling with given credentials.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.vh_domain": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.vh_domain": "true"
  }
}
```

- [simple_login](data-sources--api_crawler--properties--domains--simple_login.md): complete subsection reference.

## Next pages

- [domains.simple_login](data-sources--api_crawler--properties--domains--simple_login.md)
- [Property reference](data-sources--api_crawler--reference.md)
- [xcsh_api_crawler](../data-sources/api_crawler.md)
