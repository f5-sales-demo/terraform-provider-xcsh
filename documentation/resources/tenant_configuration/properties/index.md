---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_tenant_configuration."
xcsh_docs: {"aliases": ["tenant configuration"], "body_bytes": 11234, "body_sha256": "sha256:6bb5a989863bd32d1b0dde64baede0d19581ac86a0731c4be08c159e84e8a77a", "capabilities": ["administration"], "category": "administration", "child_ids": ["xcsh-docs:resources:tenant_configuration:properties:brute_force_detection", "xcsh-docs:resources:tenant_configuration:properties:password_policy", "xcsh-docs:resources:tenant_configuration:properties:tenant_details", "xcsh-docs:resources:tenant_configuration:properties:timeouts", "xcsh-docs:resources:tenant_configuration:properties:user_session_expiration"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:tenant_configuration:collection", "completeness": "complete", "id": "xcsh-docs:resources:tenant_configuration:reference", "parent_id": "xcsh-docs:resources:tenant_configuration:fundamentals", "path": "documentation/resources/tenant_configuration/properties/index.md", "product": "distributed-cloud", "provider_name": "tenant_configuration", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0112010312211210-2330131301100232-0100013221012023-0100320210300013-2333113013332121-3020301131332100-2310323331122100-1302302023032202", "registry_path": "docs/guides/resources--tenant_configuration--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.", "document_id": "xcsh-docs:resources:tenant_configuration:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["brute force detection"], "anchor": "section", "description": "Configuration parameter for brute force detection.", "document_id": "xcsh-docs:resources:tenant_configuration:properties:brute_force_detection", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["brute_force_detection"], "syntax": "block", "type": "object"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:resources:tenant_configuration:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["disable"], "anchor": "schema-disable", "description": "A value of true administratively disables the object.", "document_id": "xcsh-docs:resources:tenant_configuration:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["disable"], "syntax": "attribute", "type": "bool"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:resources:tenant_configuration:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Labels is a user defined key value map that can be attached to resources for organization and filtering.", "document_id": "xcsh-docs:resources:tenant_configuration:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "Name of the Tenant Configuration. Must be unique within the namespace.", "document_id": "xcsh-docs:resources:tenant_configuration:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "Namespace where the Tenant Configuration is created.", "document_id": "xcsh-docs:resources:tenant_configuration:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["password policy"], "anchor": "section", "description": "Policy configuration for this feature.", "document_id": "xcsh-docs:resources:tenant_configuration:properties:password_policy", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-password_policy--minimum_length", "enforcement": "provider-schema", "group": "password_policy:RequiredObjectAttributes:minimum_length", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:tenant_configuration:properties:password_policy", "type": "requires"}], "schema_path": ["password_policy"], "syntax": "block", "type": "object"}, {"aliases": ["tenant details"], "anchor": "section", "description": "BasicConfiguration.", "document_id": "xcsh-docs:resources:tenant_configuration:properties:tenant_details", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-tenant_details--display_name", "enforcement": "provider-schema", "group": "tenant_details:RequiredObjectAttributes:display_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:tenant_configuration:properties:tenant_details", "type": "requires"}], "schema_path": ["tenant_details"], "syntax": "block", "type": "object"}, {"aliases": ["duration", "operation timeout", "timeouts"], "anchor": "section", "description": "timeouts", "document_id": "xcsh-docs:resources:tenant_configuration:properties:timeouts", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["timeouts"], "syntax": "block", "type": "object"}, {"aliases": ["authentication", "credential setup", "credentials", "duration", "operation timeout", "user session expiration"], "anchor": "section", "description": "Defines all session-related expiration for user sessions within a tenant's environment. Relationship between session_expiry and cookie_expiry: - session_expiry defines the 'absolute maximum duration' of a session and enforces RE-authentication after this time. - cookie_expiry defines the 'inactivity timeout', which", "document_id": "xcsh-docs:resources:tenant_configuration:properties:user_session_expiration", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["user_session_expiration"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tenant_configuration/properties/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Property reference for xcsh_tenant_configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["tenant_configurationCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_tenant_configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata.

- [brute_force_detection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/brute_force_detection/): complete subsection reference.

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

- [password_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/password_policy/): complete subsection reference.

- [tenant_details](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/tenant_details/): complete subsection reference.

- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/timeouts/): complete subsection reference.

- [user_session_expiration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/user_session_expiration/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/#schema-annotations) |
| `brute_force_detection` | [brute_force_detection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/brute_force_detection/#section) |
| `brute_force_detection.max_login_failures` | [brute_force_detection.max_login_failures](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/brute_force_detection/#schema-brute_force_detection--max_login_failures) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/#schema-description) |
| `disable` | [disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/#schema-disable) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/#schema-namespace) |
| `password_policy` | [password_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/password_policy/#section) |
| `password_policy.digits` | [password_policy.digits](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/password_policy/#schema-password_policy--digits) |
| `password_policy.expire_password` | [password_policy.expire_password](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/password_policy/#schema-password_policy--expire_password) |
| `password_policy.lowercase_characters` | [password_policy.lowercase_characters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/password_policy/#schema-password_policy--lowercase_characters) |
| `password_policy.minimum_length` | [password_policy.minimum_length](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/password_policy/#schema-password_policy--minimum_length) |
| `password_policy.not_recently_used` | [password_policy.not_recently_used](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/password_policy/#schema-password_policy--not_recently_used) |
| `password_policy.not_username` | [password_policy.not_username](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/password_policy/#schema-password_policy--not_username) |
| `password_policy.special_characters` | [password_policy.special_characters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/password_policy/#schema-password_policy--special_characters) |
| `password_policy.uppercase_characters` | [password_policy.uppercase_characters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/password_policy/#schema-password_policy--uppercase_characters) |
| `tenant_details` | [tenant_details](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/tenant_details/#section) |
| `tenant_details.display_name` | [tenant_details.display_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/tenant_details/#schema-tenant_details--display_name) |
| `timeouts` | [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/timeouts/#section) |
| `timeouts.create` | [timeouts.create](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/timeouts/#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/timeouts/#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/timeouts/#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/timeouts/#schema-timeouts--update) |
| `user_session_expiration` | [user_session_expiration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/user_session_expiration/#section) |
| `user_session_expiration.absolute_timeout` | [user_session_expiration.absolute_timeout](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/user_session_expiration/absolute_timeout/#section) |
| `user_session_expiration.absolute_timeout.hours` | [user_session_expiration.absolute_timeout.hours](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/user_session_expiration/absolute_timeout/hours/#section) |
| `user_session_expiration.absolute_timeout.hours.duration` | [user_session_expiration.absolute_timeout.hours.duration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/user_session_expiration/absolute_timeout/hours/#schema-user_session_expiration--absolute_timeout--hours--duration) |
| `user_session_expiration.absolute_timeout.minutes` | [user_session_expiration.absolute_timeout.minutes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/user_session_expiration/absolute_timeout/minutes/#section) |
| `user_session_expiration.absolute_timeout.minutes.duration` | [user_session_expiration.absolute_timeout.minutes.duration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/user_session_expiration/absolute_timeout/minutes/#schema-user_session_expiration--absolute_timeout--minutes--duration) |
| `user_session_expiration.idle_timeout` | [user_session_expiration.idle_timeout](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/user_session_expiration/idle_timeout/#section) |
| `user_session_expiration.idle_timeout.hours` | [user_session_expiration.idle_timeout.hours](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/user_session_expiration/idle_timeout/hours/#section) |
| `user_session_expiration.idle_timeout.hours.duration` | [user_session_expiration.idle_timeout.hours.duration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/user_session_expiration/idle_timeout/hours/#schema-user_session_expiration--idle_timeout--hours--duration) |
| `user_session_expiration.idle_timeout.minutes` | [user_session_expiration.idle_timeout.minutes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/user_session_expiration/idle_timeout/minutes/#section) |
| `user_session_expiration.idle_timeout.minutes.duration` | [user_session_expiration.idle_timeout.minutes.duration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/user_session_expiration/idle_timeout/minutes/#schema-user_session_expiration--idle_timeout--minutes--duration) |

## Next pages

- [brute_force_detection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/brute_force_detection/)
- [password_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/password_policy/)
- [tenant_details](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/tenant_details/)
- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/timeouts/)
- [user_session_expiration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/user_session_expiration/)
- [xcsh_tenant_configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/)
