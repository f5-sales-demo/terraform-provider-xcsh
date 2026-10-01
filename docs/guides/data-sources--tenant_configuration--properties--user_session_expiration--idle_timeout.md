---
page_title: "user_session_expiration.idle_timeout"
subcategory: ""
description: "user_session_expiration.idle_timeout for xcsh_tenant_configuration."
xcsh_docs: {"aliases": [], "body_bytes": 1562, "body_sha256": "sha256:5ed22a7516bcc03f565ec9b7414f19d56c7daa8660bf3427817e399bafa8dec1", "canonical_id": "xcsh-docs:data-sources:tenant_configuration:properties:user_session_expiration:idle_timeout", "child_ids": ["xcsh-docs:data-sources:tenant_configuration:properties:user_session_expiration:idle_timeout:hours", "xcsh-docs:data-sources:tenant_configuration:properties:user_session_expiration:idle_timeout:minutes"], "collection_id": "xcsh-docs:data-sources:tenant_configuration:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tenant_configuration:properties:user_session_expiration:idle_timeout", "parent_id": "xcsh-docs:data-sources:tenant_configuration:properties:user_session_expiration", "path": "docs/guides/data-sources--tenant_configuration--properties--user_session_expiration--idle_timeout.md", "provider_name": "tenant_configuration", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["user_session_expiration", "idle_timeout"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tenant_configuration/properties/user_session_expiration/idle_timeout/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "user_session_expiration.idle_timeout for xcsh_tenant_configuration.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tenant_configurationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# user_session_expiration.idle_timeout

Breadcrumbs:

- [xcsh_tenant_configuration](../data-sources/tenant_configuration.md)
- [Property reference](data-sources--tenant_configuration--reference.md)
- [user_session_expiration](data-sources--tenant_configuration--properties--user_session_expiration.md)
- user_session_expiration.idle_timeout

<a id="section"></a>

Type: `"single"`. Computed.

Represents the cookie expiration duration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-unit_of_time": "[\"hours\",\"minutes\"]"
}
```

## Direct properties

- [hours](data-sources--tenant_configuration--properties--user_session_expiration--idle_timeout--hours.md): complete subsection reference.

- [minutes](data-sources--tenant_configuration--properties--user_session_expiration--idle_timeout--minutes.md): complete subsection reference.

## Next pages

- [user_session_expiration.idle_timeout.hours](data-sources--tenant_configuration--properties--user_session_expiration--idle_timeout--hours.md)
- [user_session_expiration.idle_timeout.minutes](data-sources--tenant_configuration--properties--user_session_expiration--idle_timeout--minutes.md)
- [user_session_expiration](data-sources--tenant_configuration--properties--user_session_expiration.md)
- [xcsh_tenant_configuration](../data-sources/tenant_configuration.md)
