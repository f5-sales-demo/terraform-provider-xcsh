---
page_title: "custom_anonymization.anonymization_config.http_header"
subcategory: "Security"
description: "custom_anonymization.anonymization_config.http_header for xcsh_app_firewall."
xcsh_docs: {"aliases": [], "body_bytes": 2425, "body_sha256": "sha256:7984e8d2437918e57d2289c94afcd5956414d265d466389fd1cf6d0240d1190c", "canonical_id": "xcsh-docs:data-sources:app_firewall:properties:custom_anonymization:anonymization_config:http_header", "child_ids": [], "collection_id": "xcsh-docs:data-sources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_firewall:properties:custom_anonymization:anonymization_config:http_header", "parent_id": "xcsh-docs:data-sources:app_firewall:properties:custom_anonymization:anonymization_config", "path": "docs/guides/data-sources--app_firewall--properties--custom_anonymization--anonymization_config--http_header.md", "provider_name": "app_firewall", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_anonymization", "anonymization_config", "http_header"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_firewall/properties/custom_anonymization/anonymization_config/http_header/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_anonymization.anonymization_config.http_header for xcsh_app_firewall.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_firewallCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_anonymization.anonymization_config.http_header

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md)
- [Property reference](data-sources--app_firewall--reference.md)
- [custom_anonymization](data-sources--app_firewall--properties--custom_anonymization.md)
- [custom_anonymization.anonymization_config](data-sources--app_firewall--properties--custom_anonymization--anonymization_config.md)
- custom_anonymization.anonymization_config.http_header

<a id="section"></a>

Type: `"single"`. Computed.

Configure anonymization for HTTP Headers.

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

<a id="schema-custom_anonymization--anonymization_config--http_header--header_name"></a>

### header_name property

Type: `"string"`. Computed.

Masks the HTTP header value. The setting does not mask the HTTP header name. Wildcard matching can
be used by prefixing or suffixing the HTTP header name with a wildcard asterisk (\*), or by using
only an asterisk to match any HTTP header name.

Upstream description:

Masks the HTTP header value. The setting does not mask the HTTP header name. Wildcard matching can
be used by prefixing or suffixing the HTTP header name with a wildcard asterisk (\*), or by using
only an asterisk to match any HTTP header name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
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
    "ves.io.schema.rules.string.http_header_field": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true"
  }
}
```

## Next pages

- [custom_anonymization.anonymization_config](data-sources--app_firewall--properties--custom_anonymization--anonymization_config.md)
- [xcsh_app_firewall](../data-sources/app_firewall.md)
