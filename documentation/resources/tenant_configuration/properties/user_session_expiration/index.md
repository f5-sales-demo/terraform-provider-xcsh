---
page_title: "user_session_expiration"
subcategory: ""
description: "Defines all session-related expiration for user sessions within a tenant's environment. Relationship between session_expiry and cookie_expiry: - session_expiry defines the 'absolute maximum duration' of a session and enforces RE-authentication after this time. - cookie_expiry defines the 'inactivity timeout', which"
xcsh_docs: {"aliases": ["duration", "user session expiration"], "body_bytes": 2647, "body_sha256": "sha256:fdc644056c8682d048b652b0d489907748cf1e3807dfcbbf067f919db938c53f", "capabilities": ["administration"], "category": "administration", "child_ids": ["xcsh-docs:resources:tenant_configuration:properties:user_session_expiration:absolute_timeout", "xcsh-docs:resources:tenant_configuration:properties:user_session_expiration:idle_timeout"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:tenant_configuration:collection", "completeness": "complete", "id": "xcsh-docs:resources:tenant_configuration:properties:user_session_expiration", "parent_id": "xcsh-docs:resources:tenant_configuration:reference", "path": "documentation/resources/tenant_configuration/properties/user_session_expiration/index.md", "product": "distributed-cloud", "provider_name": "tenant_configuration", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-3002032020221331-0023102232020311-2221003000102010-1231312030013330-1231033101320300-2312123103202330-1322112233100103-2133031101202322", "registry_path": "docs/guides/resources--tenant_configuration--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["user_session_expiration"], "schema_version": 1, "sections": [{"aliases": ["duration", "user session expiration absolute timeout"], "anchor": "section", "description": "Represents the session expiration duration.", "document_id": "xcsh-docs:resources:tenant_configuration:properties:user_session_expiration:absolute_timeout", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "user_session_expiration.absolute_timeout:ConflictingObjectAttributes:hours,minutes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tenant_configuration:properties:user_session_expiration:absolute_timeout:hours", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "user_session_expiration.absolute_timeout:ConflictingObjectAttributes:hours,minutes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tenant_configuration:properties:user_session_expiration:absolute_timeout:minutes", "type": "conflicts"}], "schema_path": ["user_session_expiration", "absolute_timeout"], "syntax": "block", "type": "object"}, {"aliases": ["duration", "user session expiration idle timeout"], "anchor": "section", "description": "Represents the cookie expiration duration.", "document_id": "xcsh-docs:resources:tenant_configuration:properties:user_session_expiration:idle_timeout", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "user_session_expiration.idle_timeout:ConflictingObjectAttributes:hours,minutes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tenant_configuration:properties:user_session_expiration:idle_timeout:hours", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "user_session_expiration.idle_timeout:ConflictingObjectAttributes:hours,minutes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tenant_configuration:properties:user_session_expiration:idle_timeout:minutes", "type": "conflicts"}], "schema_path": ["user_session_expiration", "idle_timeout"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tenant_configuration/properties/user_session_expiration/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Defines all session-related expiration for user sessions within a tenant's environment. Relationship between session_expiry and cookie_expiry: - session_expiry defines the 'absolute maximum duration' of a session and enforces RE-authentication after this time. - cookie_expiry defines the 'inactivity timeout', which", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["tenant_configurationCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

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
