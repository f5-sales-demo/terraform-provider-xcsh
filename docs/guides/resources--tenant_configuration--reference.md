---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_tenant_configuration."
xcsh_docs: {"aliases": [], "body_bytes": 8790, "body_sha256": "sha256:c3d9b9cb5a0e673ce2d028fcb2aa214a5aed82e7f63859d0ac0dbeafb3bfc37f", "canonical_id": "xcsh-docs:resources:tenant_configuration:reference", "child_ids": ["xcsh-docs:resources:tenant_configuration:properties:brute_force_detection", "xcsh-docs:resources:tenant_configuration:properties:password_policy", "xcsh-docs:resources:tenant_configuration:properties:tenant_details", "xcsh-docs:resources:tenant_configuration:properties:timeouts", "xcsh-docs:resources:tenant_configuration:properties:user_session_expiration"], "collection_id": "xcsh-docs:resources:tenant_configuration:collection", "completeness": "complete", "id": "xcsh-docs:resources:tenant_configuration:reference", "parent_id": "xcsh-docs:resources:tenant_configuration:fundamentals", "path": "docs/guides/resources--tenant_configuration--reference.md", "provider_name": "tenant_configuration", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tenant_configuration/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_tenant_configuration.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tenant_configurationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_tenant_configuration](../resources/tenant_configuration.md)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata.

- [brute_force_detection](resources--tenant_configuration--properties--brute_force_detection.md): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Optional.

Human readable description for the object.

<a id="schema-disable"></a>

### disable property

Type: `"bool"`. Optional.

A value of true administratively disables the object.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Optional.

Labels is a user defined key value map that can be attached to resources for organization and
filtering.

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the Tenant Configuration. Must be unique within the namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NameValidator(),
}
```

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace where the Tenant Configuration is created.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NamespaceValidator(),
}
```

- [password_policy](resources--tenant_configuration--properties--password_policy.md): complete subsection reference.

- [tenant_details](resources--tenant_configuration--properties--tenant_details.md): complete subsection reference.

- [timeouts](resources--tenant_configuration--properties--timeouts.md): complete subsection reference.

- [user_session_expiration](resources--tenant_configuration--properties--user_session_expiration.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--tenant_configuration--reference.md#schema-annotations) |
| `brute_force_detection` | [brute_force_detection](resources--tenant_configuration--properties--brute_force_detection.md#section) |
| `brute_force_detection.max_login_failures` | [brute_force_detection.max_login_failures](resources--tenant_configuration--properties--brute_force_detection.md#schema-brute_force_detection--max_login_failures) |
| `description` | [description](resources--tenant_configuration--reference.md#schema-description) |
| `disable` | [disable](resources--tenant_configuration--reference.md#schema-disable) |
| `id` | [id](resources--tenant_configuration--reference.md#schema-id) |
| `labels` | [labels](resources--tenant_configuration--reference.md#schema-labels) |
| `name` | [name](resources--tenant_configuration--reference.md#schema-name) |
| `namespace` | [namespace](resources--tenant_configuration--reference.md#schema-namespace) |
| `password_policy` | [password_policy](resources--tenant_configuration--properties--password_policy.md#section) |
| `password_policy.digits` | [password_policy.digits](resources--tenant_configuration--properties--password_policy.md#schema-password_policy--digits) |
| `password_policy.expire_password` | [password_policy.expire_password](resources--tenant_configuration--properties--password_policy.md#schema-password_policy--expire_password) |
| `password_policy.lowercase_characters` | [password_policy.lowercase_characters](resources--tenant_configuration--properties--password_policy.md#schema-password_policy--lowercase_characters) |
| `password_policy.minimum_length` | [password_policy.minimum_length](resources--tenant_configuration--properties--password_policy.md#schema-password_policy--minimum_length) |
| `password_policy.not_recently_used` | [password_policy.not_recently_used](resources--tenant_configuration--properties--password_policy.md#schema-password_policy--not_recently_used) |
| `password_policy.not_username` | [password_policy.not_username](resources--tenant_configuration--properties--password_policy.md#schema-password_policy--not_username) |
| `password_policy.special_characters` | [password_policy.special_characters](resources--tenant_configuration--properties--password_policy.md#schema-password_policy--special_characters) |
| `password_policy.uppercase_characters` | [password_policy.uppercase_characters](resources--tenant_configuration--properties--password_policy.md#schema-password_policy--uppercase_characters) |
| `tenant_details` | [tenant_details](resources--tenant_configuration--properties--tenant_details.md#section) |
| `tenant_details.display_name` | [tenant_details.display_name](resources--tenant_configuration--properties--tenant_details.md#schema-tenant_details--display_name) |
| `timeouts` | [timeouts](resources--tenant_configuration--properties--timeouts.md#section) |
| `timeouts.create` | [timeouts.create](resources--tenant_configuration--properties--timeouts.md#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](resources--tenant_configuration--properties--timeouts.md#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](resources--tenant_configuration--properties--timeouts.md#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](resources--tenant_configuration--properties--timeouts.md#schema-timeouts--update) |
| `user_session_expiration` | [user_session_expiration](resources--tenant_configuration--properties--user_session_expiration.md#section) |
| `user_session_expiration.absolute_timeout` | [user_session_expiration.absolute_timeout](resources--tenant_configuration--properties--user_session_expiration--absolute_timeout.md#section) |
| `user_session_expiration.absolute_timeout.hours` | [user_session_expiration.absolute_timeout.hours](resources--tenant_configuration--properties--user_session_expiration--absolute_timeout--hours.md#section) |
| `user_session_expiration.absolute_timeout.hours.duration` | [user_session_expiration.absolute_timeout.hours.duration](resources--tenant_configuration--properties--user_session_expiration--absolute_timeout--hours.md#schema-user_session_expiration--absolute_timeout--hours--duration) |
| `user_session_expiration.absolute_timeout.minutes` | [user_session_expiration.absolute_timeout.minutes](resources--tenant_configuration--properties--user_session_expiration--absolute_timeout--minutes.md#section) |
| `user_session_expiration.absolute_timeout.minutes.duration` | [user_session_expiration.absolute_timeout.minutes.duration](resources--tenant_configuration--properties--user_session_expiration--absolute_timeout--minutes.md#schema-user_session_expiration--absolute_timeout--minutes--duration) |
| `user_session_expiration.idle_timeout` | [user_session_expiration.idle_timeout](resources--tenant_configuration--properties--user_session_expiration--idle_timeout.md#section) |
| `user_session_expiration.idle_timeout.hours` | [user_session_expiration.idle_timeout.hours](resources--tenant_configuration--properties--user_session_expiration--idle_timeout--hours.md#section) |
| `user_session_expiration.idle_timeout.hours.duration` | [user_session_expiration.idle_timeout.hours.duration](resources--tenant_configuration--properties--user_session_expiration--idle_timeout--hours.md#schema-user_session_expiration--idle_timeout--hours--duration) |
| `user_session_expiration.idle_timeout.minutes` | [user_session_expiration.idle_timeout.minutes](resources--tenant_configuration--properties--user_session_expiration--idle_timeout--minutes.md#section) |
| `user_session_expiration.idle_timeout.minutes.duration` | [user_session_expiration.idle_timeout.minutes.duration](resources--tenant_configuration--properties--user_session_expiration--idle_timeout--minutes.md#schema-user_session_expiration--idle_timeout--minutes--duration) |

## Next pages

- [brute_force_detection](resources--tenant_configuration--properties--brute_force_detection.md)
- [password_policy](resources--tenant_configuration--properties--password_policy.md)
- [tenant_details](resources--tenant_configuration--properties--tenant_details.md)
- [timeouts](resources--tenant_configuration--properties--timeouts.md)
- [user_session_expiration](resources--tenant_configuration--properties--user_session_expiration.md)
- [xcsh_tenant_configuration](../resources/tenant_configuration.md)
