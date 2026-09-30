---
page_title: "jwt_validation.authorization_server"
subcategory: "Load Balancing"
description: "jwt_validation.authorization_server for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1103, "body_sha256": "sha256:a9ccea548ce22c402020e62c28b2554997d86a75fd2def944b5bdd5539b2f9a4", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation:authorization_server", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation:authorization_server:authorization_servers"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation:authorization_server", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation", "path": "docs/guides/data-sources--http_loadbalancer--properties--jwt_validation--authorization_server.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["jwt_validation", "authorization_server"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/jwt_validation/authorization_server/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "jwt_validation.authorization_server for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# jwt_validation.authorization_server

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [jwt_validation](data-sources--http_loadbalancer--properties--jwt_validation.md)
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

- [authorization_servers](data-sources--http_loadbalancer--properties--jwt_validation--authorization_server--authorization_servers.md): complete subsection reference.

## Next pages

- [jwt_validation.authorization_server.authorization_servers](data-sources--http_loadbalancer--properties--jwt_validation--authorization_server--authorization_servers.md)
- [jwt_validation](data-sources--http_loadbalancer--properties--jwt_validation.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
