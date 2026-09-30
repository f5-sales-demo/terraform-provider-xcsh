---
page_title: "user_session_expiration.idle_timeout"
subcategory: ""
description: "user_session_expiration.idle_timeout for xcsh_tenant_configuration."
xcsh_docs: {"aliases": [], "body_bytes": 1916, "body_sha256": "sha256:51211c3461c3b4c73c30af01303d7afd5b481ca797e82a6af9d32e7b2310f0c6", "child_ids": ["xcsh-docs:data-sources:tenant_configuration:properties:user_session_expiration:idle_timeout:hours", "xcsh-docs:data-sources:tenant_configuration:properties:user_session_expiration:idle_timeout:minutes"], "collection_id": "xcsh-docs:data-sources:tenant_configuration:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tenant_configuration:properties:user_session_expiration:idle_timeout", "parent_id": "xcsh-docs:data-sources:tenant_configuration:properties:user_session_expiration", "path": "documentation/data-sources/tenant_configuration/properties/user_session_expiration/idle_timeout/index.md", "provider_name": "tenant_configuration", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["user_session_expiration", "idle_timeout"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tenant_configuration/properties/user_session_expiration/idle_timeout/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "user_session_expiration.idle_timeout for xcsh_tenant_configuration.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tenant_configurationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# user_session_expiration.idle_timeout

Breadcrumbs:

- [xcsh_tenant_configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/)
- [user_session_expiration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/user_session_expiration/)
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

- [hours](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/user_session_expiration/idle_timeout/hours/): complete subsection reference.

- [minutes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/user_session_expiration/idle_timeout/minutes/): complete subsection reference.

## Next pages

- [user_session_expiration.idle_timeout.hours](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/user_session_expiration/idle_timeout/hours/)
- [user_session_expiration.idle_timeout.minutes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/user_session_expiration/idle_timeout/minutes/)
- [user_session_expiration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/user_session_expiration/)
- [xcsh_tenant_configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/)
