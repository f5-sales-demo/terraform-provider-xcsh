---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_tenant_configuration."
xcsh_docs: {"aliases": [], "body_bytes": 7439, "body_sha256": "sha256:d157133191863160b3e5fab07931d234f42609976747563fdaa30679fbea1806", "canonical_id": "xcsh-docs:data-sources:tenant_configuration:reference", "child_ids": ["xcsh-docs:data-sources:tenant_configuration:properties:brute_force_detection", "xcsh-docs:data-sources:tenant_configuration:properties:password_policy", "xcsh-docs:data-sources:tenant_configuration:properties:tenant_details", "xcsh-docs:data-sources:tenant_configuration:properties:user_session_expiration"], "collection_id": "xcsh-docs:data-sources:tenant_configuration:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tenant_configuration:reference", "parent_id": "xcsh-docs:data-sources:tenant_configuration:fundamentals", "path": "docs/guides/data-sources--tenant_configuration--reference.md", "provider_name": "tenant_configuration", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tenant_configuration/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_tenant_configuration.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tenant_configurationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_tenant_configuration](../data-sources/tenant_configuration.md)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

- [brute_force_detection](data-sources--tenant_configuration--properties--brute_force_detection.md): complete subsection reference.

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

- [password_policy](data-sources--tenant_configuration--properties--password_policy.md): complete subsection reference.

- [tenant_details](data-sources--tenant_configuration--properties--tenant_details.md): complete subsection reference.

- [user_session_expiration](data-sources--tenant_configuration--properties--user_session_expiration.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--tenant_configuration--reference.md#schema-annotations) |
| `brute_force_detection` | [brute_force_detection](data-sources--tenant_configuration--properties--brute_force_detection.md#section) |
| `brute_force_detection.max_login_failures` | [brute_force_detection.max_login_failures](data-sources--tenant_configuration--properties--brute_force_detection.md#schema-brute_force_detection--max_login_failures) |
| `description` | [description](data-sources--tenant_configuration--reference.md#schema-description) |
| `id` | [id](data-sources--tenant_configuration--reference.md#schema-id) |
| `labels` | [labels](data-sources--tenant_configuration--reference.md#schema-labels) |
| `name` | [name](data-sources--tenant_configuration--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--tenant_configuration--reference.md#schema-namespace) |
| `password_policy` | [password_policy](data-sources--tenant_configuration--properties--password_policy.md#section) |
| `password_policy.digits` | [password_policy.digits](data-sources--tenant_configuration--properties--password_policy.md#schema-password_policy--digits) |
| `password_policy.expire_password` | [password_policy.expire_password](data-sources--tenant_configuration--properties--password_policy.md#schema-password_policy--expire_password) |
| `password_policy.lowercase_characters` | [password_policy.lowercase_characters](data-sources--tenant_configuration--properties--password_policy.md#schema-password_policy--lowercase_characters) |
| `password_policy.minimum_length` | [password_policy.minimum_length](data-sources--tenant_configuration--properties--password_policy.md#schema-password_policy--minimum_length) |
| `password_policy.not_recently_used` | [password_policy.not_recently_used](data-sources--tenant_configuration--properties--password_policy.md#schema-password_policy--not_recently_used) |
| `password_policy.not_username` | [password_policy.not_username](data-sources--tenant_configuration--properties--password_policy.md#schema-password_policy--not_username) |
| `password_policy.special_characters` | [password_policy.special_characters](data-sources--tenant_configuration--properties--password_policy.md#schema-password_policy--special_characters) |
| `password_policy.uppercase_characters` | [password_policy.uppercase_characters](data-sources--tenant_configuration--properties--password_policy.md#schema-password_policy--uppercase_characters) |
| `tenant_details` | [tenant_details](data-sources--tenant_configuration--properties--tenant_details.md#section) |
| `tenant_details.display_name` | [tenant_details.display_name](data-sources--tenant_configuration--properties--tenant_details.md#schema-tenant_details--display_name) |
| `user_session_expiration` | [user_session_expiration](data-sources--tenant_configuration--properties--user_session_expiration.md#section) |
| `user_session_expiration.absolute_timeout` | [user_session_expiration.absolute_timeout](data-sources--tenant_configuration--properties--user_session_expiration--absolute_timeout.md#section) |
| `user_session_expiration.absolute_timeout.hours` | [user_session_expiration.absolute_timeout.hours](data-sources--tenant_configuration--properties--user_session_expiration--absolute_timeout--hours.md#section) |
| `user_session_expiration.absolute_timeout.hours.duration` | [user_session_expiration.absolute_timeout.hours.duration](data-sources--tenant_configuration--properties--user_session_expiration--absolute_timeout--hours.md#schema-user_session_expiration--absolute_timeout--hours--duration) |
| `user_session_expiration.absolute_timeout.minutes` | [user_session_expiration.absolute_timeout.minutes](data-sources--tenant_configuration--properties--user_session_expiration--absolute_timeout--minutes.md#section) |
| `user_session_expiration.absolute_timeout.minutes.duration` | [user_session_expiration.absolute_timeout.minutes.duration](data-sources--tenant_configuration--properties--user_session_expiration--absolute_timeout--minutes.md#schema-user_session_expiration--absolute_timeout--minutes--duration) |
| `user_session_expiration.idle_timeout` | [user_session_expiration.idle_timeout](data-sources--tenant_configuration--properties--user_session_expiration--idle_timeout.md#section) |
| `user_session_expiration.idle_timeout.hours` | [user_session_expiration.idle_timeout.hours](data-sources--tenant_configuration--properties--user_session_expiration--idle_timeout--hours.md#section) |
| `user_session_expiration.idle_timeout.hours.duration` | [user_session_expiration.idle_timeout.hours.duration](data-sources--tenant_configuration--properties--user_session_expiration--idle_timeout--hours.md#schema-user_session_expiration--idle_timeout--hours--duration) |
| `user_session_expiration.idle_timeout.minutes` | [user_session_expiration.idle_timeout.minutes](data-sources--tenant_configuration--properties--user_session_expiration--idle_timeout--minutes.md#section) |
| `user_session_expiration.idle_timeout.minutes.duration` | [user_session_expiration.idle_timeout.minutes.duration](data-sources--tenant_configuration--properties--user_session_expiration--idle_timeout--minutes.md#schema-user_session_expiration--idle_timeout--minutes--duration) |

## Next pages

- [brute_force_detection](data-sources--tenant_configuration--properties--brute_force_detection.md)
- [password_policy](data-sources--tenant_configuration--properties--password_policy.md)
- [tenant_details](data-sources--tenant_configuration--properties--tenant_details.md)
- [user_session_expiration](data-sources--tenant_configuration--properties--user_session_expiration.md)
- [xcsh_tenant_configuration](../data-sources/tenant_configuration.md)
