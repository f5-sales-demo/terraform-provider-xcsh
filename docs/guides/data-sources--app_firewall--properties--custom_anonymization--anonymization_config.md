---
page_title: "custom_anonymization.anonymization_config"
subcategory: "Security"
description: "custom_anonymization.anonymization_config for xcsh_app_firewall."
xcsh_docs: {"aliases": [], "body_bytes": 2498, "body_sha256": "sha256:447e653052a517b9179ba87c3426df25750f2a3d5867ebc9af685930279610f8", "canonical_id": "xcsh-docs:data-sources:app_firewall:properties:custom_anonymization:anonymization_config", "child_ids": ["xcsh-docs:data-sources:app_firewall:properties:custom_anonymization:anonymization_config:cookie", "xcsh-docs:data-sources:app_firewall:properties:custom_anonymization:anonymization_config:http_header", "xcsh-docs:data-sources:app_firewall:properties:custom_anonymization:anonymization_config:query_parameter"], "collection_id": "xcsh-docs:data-sources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_firewall:properties:custom_anonymization:anonymization_config", "parent_id": "xcsh-docs:data-sources:app_firewall:properties:custom_anonymization", "path": "docs/guides/data-sources--app_firewall--properties--custom_anonymization--anonymization_config.md", "provider_name": "app_firewall", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_anonymization", "anonymization_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_firewall/properties/custom_anonymization/anonymization_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_anonymization.anonymization_config for xcsh_app_firewall.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_firewallCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_anonymization.anonymization_config

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md)
- [Property reference](data-sources--app_firewall--reference.md)
- [custom_anonymization](data-sources--app_firewall--properties--custom_anonymization.md)
- custom_anonymization.anonymization_config

<a id="section"></a>

Type: `"list"`. Computed.

List of HTTP headers, cookies and query parameters whose values will be masked.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

- [cookie](data-sources--app_firewall--properties--custom_anonymization--anonymization_config--cookie.md): complete subsection reference.

- [http_header](data-sources--app_firewall--properties--custom_anonymization--anonymization_config--http_header.md): complete subsection reference.

- [query_parameter](data-sources--app_firewall--properties--custom_anonymization--anonymization_config--query_parameter.md): complete subsection reference.

## Next pages

- [custom_anonymization.anonymization_config.cookie](data-sources--app_firewall--properties--custom_anonymization--anonymization_config--cookie.md)
- [custom_anonymization.anonymization_config.http_header](data-sources--app_firewall--properties--custom_anonymization--anonymization_config--http_header.md)
- [custom_anonymization.anonymization_config.query_parameter](data-sources--app_firewall--properties--custom_anonymization--anonymization_config--query_parameter.md)
- [custom_anonymization](data-sources--app_firewall--properties--custom_anonymization.md)
- [xcsh_app_firewall](../data-sources/app_firewall.md)
