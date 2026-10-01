---
page_title: "jwt_validation"
subcategory: "Load Balancing"
description: "jwt_validation for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2935, "body_sha256": "sha256:16cfd1c425dc7ccdc178a0329d01d586d80c128c19eca88151509c2cfc8ff915", "canonical_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:jwt_validation", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:jwt_validation:action", "xcsh-docs:data-sources:cdn_loadbalancer:properties:jwt_validation:authorization_server", "xcsh-docs:data-sources:cdn_loadbalancer:properties:jwt_validation:jwks_config", "xcsh-docs:data-sources:cdn_loadbalancer:properties:jwt_validation:mandatory_claims", "xcsh-docs:data-sources:cdn_loadbalancer:properties:jwt_validation:reserved_claims", "xcsh-docs:data-sources:cdn_loadbalancer:properties:jwt_validation:target", "xcsh-docs:data-sources:cdn_loadbalancer:properties:jwt_validation:token_location"], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:jwt_validation", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:reference", "path": "docs/guides/data-sources--cdn_loadbalancer--properties--jwt_validation.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["jwt_validation"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/jwt_validation/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "jwt_validation for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# jwt_validation

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- jwt_validation

<a id="section"></a>

Type: `"single"`. Computed.

JWT Validation stops JWT replay attacks and JWT tampering by cryptographically verifying incoming
JWTs before they are passed to your API origin. JWT Validation will also stop requests with expired
tokens or tokens that are not yet valid.

Upstream description:

JWT Validation stops JWT replay attacks and JWT tampering by cryptographically verifying incoming
JWTs before they are passed to your API origin. JWT Validation will also stop requests with expired
tokens or tokens that are not yet valid.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-jwks_configuration": "[\"authorization_server\",\"jwks_config\"]"
}
```

## Direct properties

- [action](data-sources--cdn_loadbalancer--properties--jwt_validation--action.md): complete subsection reference.

- [authorization_server](data-sources--cdn_loadbalancer--properties--jwt_validation--authorization_server.md): complete subsection reference.

- [jwks_config](data-sources--cdn_loadbalancer--properties--jwt_validation--jwks_config.md): complete subsection reference.

- [mandatory_claims](data-sources--cdn_loadbalancer--properties--jwt_validation--mandatory_claims.md): complete subsection reference.

- [reserved_claims](data-sources--cdn_loadbalancer--properties--jwt_validation--reserved_claims.md): complete subsection reference.

- [target](data-sources--cdn_loadbalancer--properties--jwt_validation--target.md): complete subsection reference.

- [token_location](data-sources--cdn_loadbalancer--properties--jwt_validation--token_location.md): complete subsection reference.

## Next pages

- [jwt_validation.action](data-sources--cdn_loadbalancer--properties--jwt_validation--action.md)
- [jwt_validation.authorization_server](data-sources--cdn_loadbalancer--properties--jwt_validation--authorization_server.md)
- [jwt_validation.jwks_config](data-sources--cdn_loadbalancer--properties--jwt_validation--jwks_config.md)
- [jwt_validation.mandatory_claims](data-sources--cdn_loadbalancer--properties--jwt_validation--mandatory_claims.md)
- [jwt_validation.reserved_claims](data-sources--cdn_loadbalancer--properties--jwt_validation--reserved_claims.md)
- [jwt_validation.target](data-sources--cdn_loadbalancer--properties--jwt_validation--target.md)
- [jwt_validation.token_location](data-sources--cdn_loadbalancer--properties--jwt_validation--token_location.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
