---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_oidc_oauth_discovery."
xcsh_docs: {"aliases": [], "body_bytes": 4051, "body_sha256": "sha256:43fe79ceb83a6a9768b5819922093518225c9ae8eb9ef238e9ca706875fa9b79", "canonical_id": "xcsh-docs:data-sources:oidc_oauth_discovery:reference", "child_ids": ["xcsh-docs:data-sources:oidc_oauth_discovery:properties:auto_jwt_config_name", "xcsh-docs:data-sources:oidc_oauth_discovery:properties:trusted_ca"], "collection_id": "xcsh-docs:data-sources:oidc_oauth_discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:oidc_oauth_discovery:reference", "parent_id": "xcsh-docs:data-sources:oidc_oauth_discovery:fundamentals", "path": "docs/guides/data-sources--oidc_oauth_discovery--reference.md", "provider_name": "oidc_oauth_discovery", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/oidc_oauth_discovery/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_oidc_oauth_discovery.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_oidc_oauth_discovery](../data-sources/oidc_oauth_discovery.md)
- Property reference

## Direct properties

<a id="schema-authentication_uri"></a>

### authentication_uri property

Type: `"string"`. Computed.

Authentication URI. OAuth 2.0 authorization endpoint URI.

- [auto_jwt_config_name](data-sources--oidc_oauth_discovery--properties--auto_jwt_config_name.md): complete subsection reference.

<a id="schema-introspect_uri"></a>

### introspect_uri property

Type: `"string"`. Computed.

OAuth 2.0 token introspection endpoint URI.

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace Namespace to scope the request.

<a id="schema-openid_cfg_uri"></a>

### openid_cfg_uri property

Type: `"string"`. Optional.

OpenID Configuration URI. OpenID provider metadata URI.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

<a id="schema-token_uri"></a>

### token_uri property

Type: `"string"`. Computed.

Token URI. OAuth 2.0 token endpoint URI.

<a id="schema-token_validation_scope_uri"></a>

### token_validation_scope_uri property

Type: `"string"`. Computed.

Scope used when validating tokens against the provider.

- [trusted_ca](data-sources--oidc_oauth_discovery--properties--trusted_ca.md): complete subsection reference.

<a id="schema-userinfo_request_uri"></a>

### userinfo_request_uri property

Type: `"string"`. Computed.

UserInfo Request URI. OpenID Connect UserInfo endpoint URI.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `authentication_uri` | [authentication_uri](data-sources--oidc_oauth_discovery--reference.md#schema-authentication_uri) |
| `auto_jwt_config_name` | [auto_jwt_config_name](data-sources--oidc_oauth_discovery--properties--auto_jwt_config_name.md#section) |
| `auto_jwt_config_name.name` | [auto_jwt_config_name.name](data-sources--oidc_oauth_discovery--properties--auto_jwt_config_name.md#schema-auto_jwt_config_name--name) |
| `auto_jwt_config_name.namespace` | [auto_jwt_config_name.namespace](data-sources--oidc_oauth_discovery--properties--auto_jwt_config_name.md#schema-auto_jwt_config_name--namespace) |
| `auto_jwt_config_name.tenant` | [auto_jwt_config_name.tenant](data-sources--oidc_oauth_discovery--properties--auto_jwt_config_name.md#schema-auto_jwt_config_name--tenant) |
| `introspect_uri` | [introspect_uri](data-sources--oidc_oauth_discovery--reference.md#schema-introspect_uri) |
| `namespace` | [namespace](data-sources--oidc_oauth_discovery--reference.md#schema-namespace) |
| `openid_cfg_uri` | [openid_cfg_uri](data-sources--oidc_oauth_discovery--reference.md#schema-openid_cfg_uri) |
| `token_uri` | [token_uri](data-sources--oidc_oauth_discovery--reference.md#schema-token_uri) |
| `token_validation_scope_uri` | [token_validation_scope_uri](data-sources--oidc_oauth_discovery--reference.md#schema-token_validation_scope_uri) |
| `trusted_ca` | [trusted_ca](data-sources--oidc_oauth_discovery--properties--trusted_ca.md#section) |
| `trusted_ca.name` | [trusted_ca.name](data-sources--oidc_oauth_discovery--properties--trusted_ca.md#schema-trusted_ca--name) |
| `trusted_ca.namespace` | [trusted_ca.namespace](data-sources--oidc_oauth_discovery--properties--trusted_ca.md#schema-trusted_ca--namespace) |
| `trusted_ca.tenant` | [trusted_ca.tenant](data-sources--oidc_oauth_discovery--properties--trusted_ca.md#schema-trusted_ca--tenant) |
| `userinfo_request_uri` | [userinfo_request_uri](data-sources--oidc_oauth_discovery--reference.md#schema-userinfo_request_uri) |

## Next pages

- [auto_jwt_config_name](data-sources--oidc_oauth_discovery--properties--auto_jwt_config_name.md)
- [trusted_ca](data-sources--oidc_oauth_discovery--properties--trusted_ca.md)
- [xcsh_oidc_oauth_discovery](../data-sources/oidc_oauth_discovery.md)
