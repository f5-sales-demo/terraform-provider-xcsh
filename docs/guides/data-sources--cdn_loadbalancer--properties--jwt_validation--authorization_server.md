---
page_title: "jwt_validation.authorization_server"
subcategory: "Load Balancing"
description: "jwt_validation.authorization_server for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1094, "body_sha256": "sha256:f6a4657da2ee9b2ea24ee1a9f9c060aba5d448b703ea12bc4328ab63109ca3b5", "canonical_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:jwt_validation:authorization_server", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:jwt_validation:authorization_server:authorization_servers"], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:jwt_validation:authorization_server", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:jwt_validation", "path": "docs/guides/data-sources--cdn_loadbalancer--properties--jwt_validation--authorization_server.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["jwt_validation", "authorization_server"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/jwt_validation/authorization_server/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "jwt_validation.authorization_server for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# jwt_validation.authorization_server

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- [jwt_validation](data-sources--cdn_loadbalancer--properties--jwt_validation.md)
- jwt_validation.authorization_server

<a id="section"></a>

Type: `"single"`. Computed.

Reference to Authorization Server object.

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

## Direct properties

- [authorization_servers](data-sources--cdn_loadbalancer--properties--jwt_validation--authorization_server--authorization_servers.md): complete subsection reference.

## Next pages

- [jwt_validation.authorization_server.authorization_servers](data-sources--cdn_loadbalancer--properties--jwt_validation--authorization_server--authorization_servers.md)
- [jwt_validation](data-sources--cdn_loadbalancer--properties--jwt_validation.md)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
