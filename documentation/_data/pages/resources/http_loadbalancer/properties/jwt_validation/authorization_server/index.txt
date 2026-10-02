---
page_title: "jwt_validation.authorization_server"
subcategory: "Load Balancing"
description: "Reference to Authorization Server object."
xcsh_docs: {"aliases": ["jwt validation authorization server"], "body_bytes": 1825, "body_sha256": "sha256:16ca1740bb374dc996cff039b8d4266890f0763233433475accbf9a3fdb5131a", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:jwt_validation:authorization_server:authorization_servers"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:jwt_validation:authorization_server", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:jwt_validation", "path": "documentation/resources/http_loadbalancer/properties/jwt_validation/authorization_server/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2022003113111113-0032002130102133-2220130121300000-0122202022230311-2300211112222121-1303223203111120-1213002002313001-1311023131121310", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-020.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "jwt_validation.authorization_server:RequiredObjectAttributes:authorization_servers", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:jwt_validation:authorization_server:authorization_servers", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["jwt_validation", "authorization_server"], "schema_version": 1, "sections": [{"aliases": ["authorization servers"], "anchor": "section", "description": "Authorization Servers are configured separately in the 'Shared Objects' section of the Web App & API Protection workspace and used to fetch JWKS for JWT validation.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:jwt_validation:authorization_server:authorization_servers", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-jwt_validation--authorization_server--authorization_servers--name", "enforcement": "provider-schema", "group": "jwt_validation.authorization_server.authorization_servers:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:jwt_validation:authorization_server:authorization_servers", "type": "requires"}], "schema_path": ["jwt_validation", "authorization_server", "authorization_servers"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/jwt_validation/authorization_server/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Reference to Authorization Server object.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# jwt_validation.authorization_server

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [jwt_validation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/jwt_validation/)
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

- [authorization_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/jwt_validation/authorization_server/authorization_servers/): complete subsection reference.

## Next pages

- [jwt_validation.authorization_server.authorization_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/jwt_validation/authorization_server/authorization_servers/)
- [jwt_validation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/jwt_validation/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
