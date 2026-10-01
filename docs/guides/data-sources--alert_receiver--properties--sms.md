---
page_title: "sms"
subcategory: ""
description: "sms for xcsh_alert_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1596, "body_sha256": "sha256:83f63cb728d9e673ec6967b3d71ebffc12f425f0009c92a37c75992d54374683", "canonical_id": "xcsh-docs:data-sources:alert_receiver:properties:sms", "child_ids": [], "collection_id": "xcsh-docs:data-sources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:alert_receiver:properties:sms", "parent_id": "xcsh-docs:data-sources:alert_receiver:reference", "path": "docs/guides/data-sources--alert_receiver--properties--sms.md", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["sms"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/alert_receiver/properties/sms/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "sms for xcsh_alert_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# sms

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md)
- [Property reference](data-sources--alert_receiver--reference.md)
- sms

<a id="section"></a>

Type: `"single"`. Computed.

SMS Configuration.

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

<a id="schema-sms--contact_number"></a>

### contact_number property

Type: `"string"`. Computed.

Contact number of the user in ITU E.164 format \[+\]\[country code\]\[subscriber number including
area code\].

Upstream description:

Contact number of the user in ITU E.164 format \[+\]\[country code\]\[subscriber number including
area code\]

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.phone_number": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.phone_number": "true"
  }
}
```

## Next pages

- [Property reference](data-sources--alert_receiver--reference.md)
- [xcsh_alert_receiver](../data-sources/alert_receiver.md)
