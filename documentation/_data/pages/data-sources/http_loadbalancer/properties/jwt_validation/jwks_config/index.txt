---
page_title: "jwt_validation.jwks_config"
subcategory: "Load Balancing"
description: "The JSON Web Key Set (JWKS) is a set of keys used to verify JSON Web Token (JWT) issued by the Authorization Server. See RFC 7517 for more details."
xcsh_docs: {"aliases": ["jwt validation jwks config"], "body_bytes": 1885, "body_sha256": "sha256:a146609c8beff5266e5af0311899296d1e0506d716d630eac793363deb91095b", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation:jwks_config", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation", "path": "documentation/data-sources/http_loadbalancer/properties/jwt_validation/jwks_config/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1221233210212303-1303230013013133-2330013321313313-2331033212023133-0023202113210300-1000220111211222-1323110320030023-2001232301101300", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-020.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["jwt_validation", "jwks_config"], "schema_version": 1, "sections": [{"aliases": ["cleartext"], "anchor": "schema-jwt_validation--jwks_config--cleartext", "description": "The JSON Web Key Set (JWKS) is a set of keys used to verify JSON Web Token (JWT) issued by the Authorization Server. See RFC 7517 for more details.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation:jwks_config", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["jwt_validation", "jwks_config", "cleartext"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/jwt_validation/jwks_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "The JSON Web Key Set (JWKS) is a set of keys used to verify JSON Web Token (JWT) issued by the Authorization Server. See RFC 7517 for more details.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# jwt_validation.jwks_config

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [jwt_validation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/jwt_validation/)
- jwt_validation.jwks_config

<a id="section"></a>

Type: `"single"`. Computed.

The JSON Web Key Set (JWKS) is a set of keys used to verify JSON Web Token (JWT) issued by the
Authorization Server. See RFC 7517 for more details.

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

<a id="schema-jwt_validation--jwks_config--cleartext"></a>

### cleartext property

Type: `"string"`. Computed.

The JSON Web Key Set (JWKS) is a set of keys used to verify JSON Web Token (JWT) issued by the
Authorization Server. See RFC 7517 for more details.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Next pages

- [jwt_validation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/jwt_validation/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
