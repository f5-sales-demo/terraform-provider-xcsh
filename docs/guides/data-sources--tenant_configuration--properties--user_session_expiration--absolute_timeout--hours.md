---
page_title: "user_session_expiration.absolute_timeout.hours"
subcategory: ""
description: "user_session_expiration.absolute_timeout.hours for xcsh_tenant_configuration."
xcsh_docs: {"aliases": [], "body_bytes": 2160, "body_sha256": "sha256:4b72afe34f35e4795a27061362b09e32b281a5019e37cd05d79757380104fe2d", "canonical_id": "xcsh-docs:data-sources:tenant_configuration:properties:user_session_expiration:absolute_timeout:hours", "child_ids": [], "collection_id": "xcsh-docs:data-sources:tenant_configuration:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tenant_configuration:properties:user_session_expiration:absolute_timeout:hours", "parent_id": "xcsh-docs:data-sources:tenant_configuration:properties:user_session_expiration:absolute_timeout", "path": "docs/guides/data-sources--tenant_configuration--properties--user_session_expiration--absolute_timeout--hours.md", "provider_name": "tenant_configuration", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["user_session_expiration", "absolute_timeout", "hours"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tenant_configuration/properties/user_session_expiration/absolute_timeout/hours/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "user_session_expiration.absolute_timeout.hours for xcsh_tenant_configuration.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tenant_configurationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# user_session_expiration.absolute_timeout.hours

Breadcrumbs:

- [xcsh_tenant_configuration](../data-sources/tenant_configuration.md)
- [Property reference](data-sources--tenant_configuration--reference.md)
- [user_session_expiration](data-sources--tenant_configuration--properties--user_session_expiration.md)
- [user_session_expiration.absolute_timeout](data-sources--tenant_configuration--properties--user_session_expiration--absolute_timeout.md)
- user_session_expiration.absolute_timeout.hours

<a id="section"></a>

Type: `"single"`. Computed.

Represents the session duration in hours.

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

<a id="schema-user_session_expiration--absolute_timeout--hours--duration"></a>

### duration property

Type: `"number"`. Computed.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 720,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "720"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "720"
  }
}
```

## Next pages

- [user_session_expiration.absolute_timeout](data-sources--tenant_configuration--properties--user_session_expiration--absolute_timeout.md)
- [xcsh_tenant_configuration](../data-sources/tenant_configuration.md)
