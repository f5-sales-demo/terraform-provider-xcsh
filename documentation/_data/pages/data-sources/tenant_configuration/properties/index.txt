---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_tenant_configuration."
xcsh_docs: {"aliases": [], "body_bytes": 9473, "body_sha256": "sha256:d82ddd2e6a9960e5d85999e6bc5ae2f209998f814a131f8cd6fa958cc2e77326", "child_ids": ["xcsh-docs:data-sources:tenant_configuration:properties:brute_force_detection", "xcsh-docs:data-sources:tenant_configuration:properties:password_policy", "xcsh-docs:data-sources:tenant_configuration:properties:tenant_details", "xcsh-docs:data-sources:tenant_configuration:properties:user_session_expiration"], "collection_id": "xcsh-docs:data-sources:tenant_configuration:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tenant_configuration:reference", "parent_id": "xcsh-docs:data-sources:tenant_configuration:fundamentals", "path": "documentation/data-sources/tenant_configuration/properties/index.md", "provider_name": "tenant_configuration", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tenant_configuration/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_tenant_configuration.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tenant_configurationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

## Next pages

- [brute_force_detection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/brute_force_detection/)
- [password_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/password_policy/)
- [tenant_details](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/tenant_details/)
- [user_session_expiration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/user_session_expiration/)
- [xcsh_tenant_configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/)
