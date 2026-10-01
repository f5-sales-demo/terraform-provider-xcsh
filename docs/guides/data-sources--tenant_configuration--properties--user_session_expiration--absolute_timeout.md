---
page_title: "user_session_expiration.absolute_timeout"
subcategory: ""
description: "user_session_expiration.absolute_timeout for xcsh_tenant_configuration."
xcsh_docs: {"aliases": [], "body_bytes": 1595, "body_sha256": "sha256:3693651d119ed064f3ae54cb724b88883dde61203f96bfda2538b23f87b2f8ba", "canonical_id": "xcsh-docs:data-sources:tenant_configuration:properties:user_session_expiration:absolute_timeout", "child_ids": ["xcsh-docs:data-sources:tenant_configuration:properties:user_session_expiration:absolute_timeout:hours", "xcsh-docs:data-sources:tenant_configuration:properties:user_session_expiration:absolute_timeout:minutes"], "collection_id": "xcsh-docs:data-sources:tenant_configuration:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tenant_configuration:properties:user_session_expiration:absolute_timeout", "parent_id": "xcsh-docs:data-sources:tenant_configuration:properties:user_session_expiration", "path": "docs/guides/data-sources--tenant_configuration--properties--user_session_expiration--absolute_timeout.md", "provider_name": "tenant_configuration", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["user_session_expiration", "absolute_timeout"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tenant_configuration/properties/user_session_expiration/absolute_timeout/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "user_session_expiration.absolute_timeout for xcsh_tenant_configuration.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tenant_configurationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# user_session_expiration.absolute_timeout

Breadcrumbs:

- [xcsh_tenant_configuration](../data-sources/tenant_configuration.md)
- [Property reference](data-sources--tenant_configuration--reference.md)
- [user_session_expiration](data-sources--tenant_configuration--properties--user_session_expiration.md)
- user_session_expiration.absolute_timeout

<a id="section"></a>

Type: `"single"`. Computed.

Represents the session expiration duration.

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

- [hours](data-sources--tenant_configuration--properties--user_session_expiration--absolute_timeout--hours.md): complete subsection reference.

- [minutes](data-sources--tenant_configuration--properties--user_session_expiration--absolute_timeout--minutes.md): complete subsection reference.

## Next pages

- [user_session_expiration.absolute_timeout.hours](data-sources--tenant_configuration--properties--user_session_expiration--absolute_timeout--hours.md)
- [user_session_expiration.absolute_timeout.minutes](data-sources--tenant_configuration--properties--user_session_expiration--absolute_timeout--minutes.md)
- [user_session_expiration](data-sources--tenant_configuration--properties--user_session_expiration.md)
- [xcsh_tenant_configuration](../data-sources/tenant_configuration.md)
