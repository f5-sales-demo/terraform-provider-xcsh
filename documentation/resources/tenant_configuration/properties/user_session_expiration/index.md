---
page_title: "user_session_expiration"
subcategory: ""
description: "user_session_expiration for xcsh_tenant_configuration."
xcsh_docs: {"aliases": [], "body_bytes": 2548, "body_sha256": "sha256:92ef9ba51b9b27d7bc644575fe24bfcc35a2135d9835665057fc29d757a5d390", "child_ids": ["xcsh-docs:resources:tenant_configuration:properties:user_session_expiration:absolute_timeout", "xcsh-docs:resources:tenant_configuration:properties:user_session_expiration:idle_timeout"], "collection_id": "xcsh-docs:resources:tenant_configuration:collection", "completeness": "complete", "id": "xcsh-docs:resources:tenant_configuration:properties:user_session_expiration", "parent_id": "xcsh-docs:resources:tenant_configuration:reference", "path": "documentation/resources/tenant_configuration/properties/user_session_expiration/index.md", "provider_name": "tenant_configuration", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["user_session_expiration"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tenant_configuration/properties/user_session_expiration/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "user_session_expiration for xcsh_tenant_configuration.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tenant_configurationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# user_session_expiration

Breadcrumbs:

- [xcsh_tenant_configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/)
- user_session_expiration

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Defines all session-related expiration for user sessions within a tenant's environment. Relationship
between session\_expiry and cookie\_expiry: - session\_expiry defines the 'absolute maximum
duration' of a session and enforces RE-authentication after this time. - cookie\_expiry defines
the..

Upstream description:

Defines all session-related expiration for user sessions within a tenant's environment. Relationship
between session\_expiry and cookie\_expiry: &#8203;- session\_expiry defines the 'absolute maximum
duration' of a session and enforces RE-authentication after this time. &#8203;- cookie\_expiry
defines the 'inactivity timeout', which resets on user activity and only logs out users after idle
periods. Together, these ensure the user is logged out when either the session reaches its maximum
age or the user is inactive for too long.

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

Terraform syntax:

```terraform
user_session_expiration {
  # Configure direct properties listed below.
}
```

## Direct properties

- [absolute_timeout](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/user_session_expiration/absolute_timeout/): complete subsection reference.

- [idle_timeout](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/user_session_expiration/idle_timeout/): complete subsection reference.

## Next pages

- [user_session_expiration.absolute_timeout](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/user_session_expiration/absolute_timeout/)
- [user_session_expiration.idle_timeout](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/user_session_expiration/idle_timeout/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/)
- [xcsh_tenant_configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/)
