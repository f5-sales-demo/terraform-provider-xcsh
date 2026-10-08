---
page_title: "user_session_expiration"
subcategory: ""
description: "Defines all session-related expiration for user sessions within a tenant's environment. Relationship between session_expiry and cookie_expiry: - session_expiry defines the 'absolute maximum duration' of a session and enforces RE-authentication after this time. - cookie_expiry defines the 'inactivity timeout', which"
xcsh_docs: {"aliases": ["duration", "user session expiration"], "body_bytes": 1586, "body_sha256": "sha256:3f413f32b2e2f8ca3e84b96ad0f6ae1c14537f4687090dbdb40dae61ecf5642a", "capabilities": ["administration"], "category": "administration", "child_ids": ["xcsh-docs:data-sources:tenant_configuration:properties:user_session_expiration:absolute_timeout", "xcsh-docs:data-sources:tenant_configuration:properties:user_session_expiration:idle_timeout"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:tenant_configuration:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tenant_configuration:properties:user_session_expiration", "parent_id": "xcsh-docs:data-sources:tenant_configuration:reference", "path": "documentation/data-sources/tenant_configuration/properties/user_session_expiration/index.md", "product": "distributed-cloud", "provider_name": "tenant_configuration", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-0330320322333220-2110200200023002-3313000311210332-1022022021222222-0022333101311100-0311330013121030-0213312112223101-2110012011110223", "registry_path": "docs/guides/data-sources--tenant_configuration--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["user_session_expiration"], "schema_version": 1, "sections": [{"aliases": ["duration", "user session expiration absolute timeout"], "anchor": "section", "description": "Represents the session expiration duration.", "document_id": "xcsh-docs:data-sources:tenant_configuration:properties:user_session_expiration:absolute_timeout", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["user_session_expiration", "absolute_timeout"], "syntax": "attribute", "type": "object"}, {"aliases": ["duration", "user session expiration idle timeout"], "anchor": "section", "description": "Represents the cookie expiration duration.", "document_id": "xcsh-docs:data-sources:tenant_configuration:properties:user_session_expiration:idle_timeout", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["user_session_expiration", "idle_timeout"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tenant_configuration/properties/user_session_expiration/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Defines all session-related expiration for user sessions within a tenant's environment. Relationship between session_expiry and cookie_expiry: - session_expiry defines the 'absolute maximum duration' of a session and enforces RE-authentication after this time. - cookie_expiry defines the 'inactivity timeout', which", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["tenant_configurationCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# user_session_expiration

Breadcrumbs:

- [xcsh_tenant_configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/)
- user_session_expiration

<a id="section"></a>

Type: `"single"`. Computed.

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

## Direct properties

- [absolute_timeout](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/user_session_expiration/absolute_timeout/): complete subsection reference.

- [idle_timeout](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/user_session_expiration/idle_timeout/): complete subsection reference.
