---
page_title: "jwt_validation.authorization_server"
subcategory: "Load Balancing"
description: "jwt_validation.authorization_server for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1461, "body_sha256": "sha256:6ece48c597428efd504258eefc9214905077fc3835a969f72d611c4b927d2067", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:authorization_server", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:authorization_server:authorization_servers"], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:authorization_server", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation", "path": "docs/guides/resources--cdn_loadbalancer--properties--jwt_validation--authorization_server.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["jwt_validation", "authorization_server"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/jwt_validation/authorization_server/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "jwt_validation.authorization_server for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# jwt_validation.authorization_server

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [jwt_validation](resources--cdn_loadbalancer--properties--jwt_validation.md)
- jwt_validation.authorization_server

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Reference to Authorization Server object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("authorization_servers")}
```

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

Terraform syntax:

```terraform
authorization_server {
  # Configure direct properties listed below.
}
```

## Direct properties

- [authorization_servers](resources--cdn_loadbalancer--properties--jwt_validation--authorization_server--authorization_servers.md): complete subsection reference.

## Next pages

- [jwt_validation.authorization_server.authorization_servers](resources--cdn_loadbalancer--properties--jwt_validation--authorization_server--authorization_servers.md)
- [jwt_validation](resources--cdn_loadbalancer--properties--jwt_validation.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
