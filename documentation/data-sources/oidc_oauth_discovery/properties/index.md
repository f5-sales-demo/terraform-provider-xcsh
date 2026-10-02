---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_oidc_oauth_discovery."
xcsh_docs: {"aliases": ["oidc oauth discovery"], "body_bytes": 5235, "body_sha256": "sha256:862f032c2089995750ee02040080a5b97a0f23920a5f0a5405105125dea74925", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:data-sources:oidc_oauth_discovery:properties:auto_jwt_config_name", "xcsh-docs:data-sources:oidc_oauth_discovery:properties:trusted_ca"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:oidc_oauth_discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:oidc_oauth_discovery:reference", "parent_id": "xcsh-docs:data-sources:oidc_oauth_discovery:fundamentals", "path": "documentation/data-sources/oidc_oauth_discovery/properties/index.md", "product": "distributed-cloud", "provider_name": "oidc_oauth_discovery", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2000200023102030-2300333232322203-1113221101011222-0232203102332013-0201033112023021-3021031133133122-2122002130231000-3102322223332022", "registry_path": "docs/guides/data-sources--oidc_oauth_discovery--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["authentication", "authentication uri", "credential setup", "credentials"], "anchor": "schema-authentication_uri", "description": "Authentication URI. OAuth 2.0 authorization endpoint URI.", "document_id": "xcsh-docs:data-sources:oidc_oauth_discovery:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["authentication_uri"], "syntax": "attribute", "type": "string"}, {"aliases": ["auto jwt config name"], "anchor": "section", "description": "Type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:oidc_oauth_discovery:properties:auto_jwt_config_name", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["auto_jwt_config_name"], "syntax": "attribute", "type": "object"}, {"aliases": ["introspect uri"], "anchor": "schema-introspect_uri", "description": "OAuth 2.0 token introspection endpoint URI.", "document_id": "xcsh-docs:data-sources:oidc_oauth_discovery:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["introspect_uri"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "Namespace Namespace to scope the request.", "document_id": "xcsh-docs:data-sources:oidc_oauth_discovery:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["openid cfg uri"], "anchor": "schema-openid_cfg_uri", "description": "OpenID Configuration URI. OpenID provider metadata URI.", "document_id": "xcsh-docs:data-sources:oidc_oauth_discovery:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["openid_cfg_uri"], "syntax": "attribute", "type": "string"}, {"aliases": ["token uri"], "anchor": "schema-token_uri", "description": "Token URI. OAuth 2.0 token endpoint URI.", "document_id": "xcsh-docs:data-sources:oidc_oauth_discovery:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["token_uri"], "syntax": "attribute", "type": "string"}, {"aliases": ["token validation scope uri"], "anchor": "schema-token_validation_scope_uri", "description": "Scope used when validating tokens against the provider.", "document_id": "xcsh-docs:data-sources:oidc_oauth_discovery:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["token_validation_scope_uri"], "syntax": "attribute", "type": "string"}, {"aliases": ["trusted ca"], "anchor": "section", "description": "Type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:oidc_oauth_discovery:properties:trusted_ca", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["trusted_ca"], "syntax": "attribute", "type": "object"}, {"aliases": ["userinfo request uri"], "anchor": "schema-userinfo_request_uri", "description": "UserInfo Request URI. OpenID Connect UserInfo endpoint URI.", "document_id": "xcsh-docs:data-sources:oidc_oauth_discovery:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["userinfo_request_uri"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/oidc_oauth_discovery/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_oidc_oauth_discovery.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_oidc_oauth_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/oidc_oauth_discovery/)
- Property reference

## Direct properties

<a id="schema-authentication_uri"></a>

### authentication_uri property

Type: `"string"`. Computed.

Authentication URI. OAuth 2.0 authorization endpoint URI.

- [auto_jwt_config_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/oidc_oauth_discovery/properties/auto_jwt_config_name/): complete subsection reference.

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

- [trusted_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/oidc_oauth_discovery/properties/trusted_ca/): complete subsection reference.

<a id="schema-userinfo_request_uri"></a>

### userinfo_request_uri property

Type: `"string"`. Computed.

UserInfo Request URI. OpenID Connect UserInfo endpoint URI.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `authentication_uri` | [authentication_uri](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/oidc_oauth_discovery/properties/#schema-authentication_uri) |
| `auto_jwt_config_name` | [auto_jwt_config_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/oidc_oauth_discovery/properties/auto_jwt_config_name/#section) |
| `auto_jwt_config_name.name` | [auto_jwt_config_name.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/oidc_oauth_discovery/properties/auto_jwt_config_name/#schema-auto_jwt_config_name--name) |
| `auto_jwt_config_name.namespace` | [auto_jwt_config_name.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/oidc_oauth_discovery/properties/auto_jwt_config_name/#schema-auto_jwt_config_name--namespace) |
| `auto_jwt_config_name.tenant` | [auto_jwt_config_name.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/oidc_oauth_discovery/properties/auto_jwt_config_name/#schema-auto_jwt_config_name--tenant) |
| `introspect_uri` | [introspect_uri](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/oidc_oauth_discovery/properties/#schema-introspect_uri) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/oidc_oauth_discovery/properties/#schema-namespace) |
| `openid_cfg_uri` | [openid_cfg_uri](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/oidc_oauth_discovery/properties/#schema-openid_cfg_uri) |
| `token_uri` | [token_uri](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/oidc_oauth_discovery/properties/#schema-token_uri) |
| `token_validation_scope_uri` | [token_validation_scope_uri](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/oidc_oauth_discovery/properties/#schema-token_validation_scope_uri) |
| `trusted_ca` | [trusted_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/oidc_oauth_discovery/properties/trusted_ca/#section) |
| `trusted_ca.name` | [trusted_ca.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/oidc_oauth_discovery/properties/trusted_ca/#schema-trusted_ca--name) |
| `trusted_ca.namespace` | [trusted_ca.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/oidc_oauth_discovery/properties/trusted_ca/#schema-trusted_ca--namespace) |
| `trusted_ca.tenant` | [trusted_ca.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/oidc_oauth_discovery/properties/trusted_ca/#schema-trusted_ca--tenant) |
| `userinfo_request_uri` | [userinfo_request_uri](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/oidc_oauth_discovery/properties/#schema-userinfo_request_uri) |

## Next pages

- [auto_jwt_config_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/oidc_oauth_discovery/properties/auto_jwt_config_name/)
- [trusted_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/oidc_oauth_discovery/properties/trusted_ca/)
- [xcsh_oidc_oauth_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/oidc_oauth_discovery/)
