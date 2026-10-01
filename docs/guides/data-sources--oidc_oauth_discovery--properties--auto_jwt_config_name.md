---
page_title: "auto_jwt_config_name"
subcategory: ""
description: "auto_jwt_config_name for xcsh_oidc_oauth_discovery."
xcsh_docs: {"aliases": [], "body_bytes": 1840, "body_sha256": "sha256:3e43c87259f57bdef99e07e425d0b281ff30c14e9bb93b3f56f4bf9a01d7222d", "canonical_id": "xcsh-docs:data-sources:oidc_oauth_discovery:properties:auto_jwt_config_name", "child_ids": [], "collection_id": "xcsh-docs:data-sources:oidc_oauth_discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:oidc_oauth_discovery:properties:auto_jwt_config_name", "parent_id": "xcsh-docs:data-sources:oidc_oauth_discovery:reference", "path": "docs/guides/data-sources--oidc_oauth_discovery--properties--auto_jwt_config_name.md", "provider_name": "oidc_oauth_discovery", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["auto_jwt_config_name"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/oidc_oauth_discovery/properties/auto_jwt_config_name/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "auto_jwt_config_name for xcsh_oidc_oauth_discovery.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# auto_jwt_config_name

Breadcrumbs:

- [xcsh_oidc_oauth_discovery](../data-sources/oidc_oauth_discovery.md)
- [Property reference](data-sources--oidc_oauth_discovery--reference.md)
- auto_jwt_config_name

<a id="section"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

## Direct properties

<a id="schema-auto_jwt_config_name--name"></a>

### name property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

<a id="schema-auto_jwt_config_name--namespace"></a>

### namespace property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

<a id="schema-auto_jwt_config_name--tenant"></a>

### tenant property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

## Next pages

- [Property reference](data-sources--oidc_oauth_discovery--reference.md)
- [xcsh_oidc_oauth_discovery](../data-sources/oidc_oauth_discovery.md)
