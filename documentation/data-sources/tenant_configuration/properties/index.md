---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_tenant_configuration."
xcsh_docs: {"aliases": ["tenant configuration"], "body_bytes": 8753, "body_sha256": "sha256:8075bde026085c8f2de9b67c7447dfc7276cd2e6b9b44487765d837d4d0acf77", "capabilities": ["administration"], "category": "administration", "child_ids": ["xcsh-docs:data-sources:tenant_configuration:properties:brute_force_detection", "xcsh-docs:data-sources:tenant_configuration:properties:password_policy", "xcsh-docs:data-sources:tenant_configuration:properties:tenant_details", "xcsh-docs:data-sources:tenant_configuration:properties:user_session_expiration"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:tenant_configuration:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tenant_configuration:reference", "parent_id": "xcsh-docs:data-sources:tenant_configuration:fundamentals", "path": "documentation/data-sources/tenant_configuration/properties/index.md", "product": "distributed-cloud", "provider_name": "tenant_configuration", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-3122132313001203-3223033303101001-1300213201000103-1022102223020331-0330111321322330-2333230000221030-2133000302200313-2230013023120232", "registry_path": "docs/guides/data-sources--tenant_configuration--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations applied to this resource.", "document_id": "xcsh-docs:data-sources:tenant_configuration:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["brute force detection"], "anchor": "section", "description": "Configuration parameter for brute force detection.", "document_id": "xcsh-docs:data-sources:tenant_configuration:properties:brute_force_detection", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["brute_force_detection"], "syntax": "attribute", "type": "object"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Description of the TenantConfiguration.", "document_id": "xcsh-docs:data-sources:tenant_configuration:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:data-sources:tenant_configuration:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Labels applied to this resource.", "document_id": "xcsh-docs:data-sources:tenant_configuration:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "Name of the TenantConfiguration.", "document_id": "xcsh-docs:data-sources:tenant_configuration:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "Namespace where the TenantConfiguration exists.", "document_id": "xcsh-docs:data-sources:tenant_configuration:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["password policy"], "anchor": "section", "description": "Policy configuration for this feature.", "document_id": "xcsh-docs:data-sources:tenant_configuration:properties:password_policy", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["password_policy"], "syntax": "attribute", "type": "object"}, {"aliases": ["tenant details"], "anchor": "section", "description": "BasicConfiguration.", "document_id": "xcsh-docs:data-sources:tenant_configuration:properties:tenant_details", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["tenant_details"], "syntax": "attribute", "type": "object"}, {"aliases": ["duration", "user session expiration"], "anchor": "section", "description": "Defines all session-related expiration for user sessions within a tenant's environment. Relationship between session_expiry and cookie_expiry: - session_expiry defines the 'absolute maximum duration' of a session and enforces RE-authentication after this time. - cookie_expiry defines the 'inactivity timeout', which", "document_id": "xcsh-docs:data-sources:tenant_configuration:properties:user_session_expiration", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["user_session_expiration"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tenant_configuration/properties/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Property reference for xcsh_tenant_configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["tenant_configurationCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_tenant_configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

- [brute_force_detection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/brute_force_detection/): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description of the TenantConfiguration.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the TenantConfiguration.

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace where the TenantConfiguration exists.

- [password_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/password_policy/): complete subsection reference.

- [tenant_details](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/tenant_details/): complete subsection reference.

- [user_session_expiration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/user_session_expiration/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/#schema-annotations) |
| `brute_force_detection` | [brute_force_detection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/brute_force_detection/#section) |
| `brute_force_detection.max_login_failures` | [brute_force_detection.max_login_failures](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/brute_force_detection/#schema-brute_force_detection--max_login_failures) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/#schema-description) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/#schema-namespace) |
| `password_policy` | [password_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/password_policy/#section) |
| `password_policy.digits` | [password_policy.digits](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/password_policy/#schema-password_policy--digits) |
| `password_policy.expire_password` | [password_policy.expire_password](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/password_policy/#schema-password_policy--expire_password) |
| `password_policy.lowercase_characters` | [password_policy.lowercase_characters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/password_policy/#schema-password_policy--lowercase_characters) |
| `password_policy.minimum_length` | [password_policy.minimum_length](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/password_policy/#schema-password_policy--minimum_length) |
| `password_policy.not_recently_used` | [password_policy.not_recently_used](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/password_policy/#schema-password_policy--not_recently_used) |
| `password_policy.not_username` | [password_policy.not_username](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/password_policy/#schema-password_policy--not_username) |
| `password_policy.special_characters` | [password_policy.special_characters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/password_policy/#schema-password_policy--special_characters) |
| `password_policy.uppercase_characters` | [password_policy.uppercase_characters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/password_policy/#schema-password_policy--uppercase_characters) |
| `tenant_details` | [tenant_details](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/tenant_details/#section) |
| `tenant_details.display_name` | [tenant_details.display_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/tenant_details/#schema-tenant_details--display_name) |
| `user_session_expiration` | [user_session_expiration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/user_session_expiration/#section) |
| `user_session_expiration.absolute_timeout` | [user_session_expiration.absolute_timeout](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/user_session_expiration/absolute_timeout/#section) |
| `user_session_expiration.absolute_timeout.hours` | [user_session_expiration.absolute_timeout.hours](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/user_session_expiration/absolute_timeout/hours/#section) |
| `user_session_expiration.absolute_timeout.hours.duration` | [user_session_expiration.absolute_timeout.hours.duration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/user_session_expiration/absolute_timeout/hours/#schema-user_session_expiration--absolute_timeout--hours--duration) |
| `user_session_expiration.absolute_timeout.minutes` | [user_session_expiration.absolute_timeout.minutes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/user_session_expiration/absolute_timeout/minutes/#section) |
| `user_session_expiration.absolute_timeout.minutes.duration` | [user_session_expiration.absolute_timeout.minutes.duration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/user_session_expiration/absolute_timeout/minutes/#schema-user_session_expiration--absolute_timeout--minutes--duration) |
| `user_session_expiration.idle_timeout` | [user_session_expiration.idle_timeout](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/user_session_expiration/idle_timeout/#section) |
| `user_session_expiration.idle_timeout.hours` | [user_session_expiration.idle_timeout.hours](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/user_session_expiration/idle_timeout/hours/#section) |
| `user_session_expiration.idle_timeout.hours.duration` | [user_session_expiration.idle_timeout.hours.duration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/user_session_expiration/idle_timeout/hours/#schema-user_session_expiration--idle_timeout--hours--duration) |
| `user_session_expiration.idle_timeout.minutes` | [user_session_expiration.idle_timeout.minutes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/user_session_expiration/idle_timeout/minutes/#section) |
| `user_session_expiration.idle_timeout.minutes.duration` | [user_session_expiration.idle_timeout.minutes.duration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/user_session_expiration/idle_timeout/minutes/#schema-user_session_expiration--idle_timeout--minutes--duration) |
