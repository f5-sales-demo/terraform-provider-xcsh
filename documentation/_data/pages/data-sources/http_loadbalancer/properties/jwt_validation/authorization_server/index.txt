---
page_title: "jwt_validation.authorization_server"
subcategory: "Load Balancing"
description: "Reference to Authorization Server object."
xcsh_docs: {"aliases": ["jwt validation authorization server"], "body_bytes": 1557, "body_sha256": "sha256:b9ca6f8c73f0c1fee4052b974dd6a1dd01265e2e5aa84055a2f87ba5e8fb11bd", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation:authorization_server:authorization_servers"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation:authorization_server", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation", "path": "documentation/data-sources/http_loadbalancer/properties/jwt_validation/authorization_server/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3323202302020011-1013203000333112-1211331220023232-3232211311013021-0331110330133333-1013322200202110-2010313023231030-3213031121002022", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-020.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["jwt_validation", "authorization_server"], "schema_version": 1, "sections": [{"aliases": ["authorization servers"], "anchor": "section", "description": "Authorization Servers are configured separately in the 'Shared Objects' section of the Web App & API Protection workspace and used to fetch JWKS for JWT validation.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation:authorization_server:authorization_servers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["jwt_validation", "authorization_server", "authorization_servers"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/jwt_validation/authorization_server/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Reference to Authorization Server object.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# jwt_validation.authorization_server

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [jwt_validation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/jwt_validation/)
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

- [authorization_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/jwt_validation/authorization_server/authorization_servers/): complete subsection reference.

## Next pages

- [jwt_validation.authorization_server.authorization_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/jwt_validation/authorization_server/authorization_servers/)
- [jwt_validation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/jwt_validation/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
