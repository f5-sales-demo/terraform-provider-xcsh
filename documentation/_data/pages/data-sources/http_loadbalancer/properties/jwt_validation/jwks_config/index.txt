---
page_title: "jwt_validation.jwks_config"
subcategory: "Load Balancing"
description: "The JSON Web Key Set (JWKS) is a set of keys used to verify JSON Web Token (JWT) issued by the Authorization Server. See RFC 7517 for more details."
xcsh_docs: {"aliases": ["jwt validation jwks config"], "body_bytes": 1619, "body_sha256": "sha256:f1df70440d4fa82f700863a652355a8dbcf4b96c1ea344df2328c6ed8d7638a8", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation:jwks_config", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation", "path": "documentation/data-sources/http_loadbalancer/properties/jwt_validation/jwks_config/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-1221233210212303-1303230013013133-2330013321313313-2331033212023133-0023202113210300-1000220111211222-1323110320030023-2001232301101300", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-019.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["jwt_validation", "jwks_config"], "schema_version": 1, "sections": [{"aliases": ["jwt validation jwks config cleartext"], "anchor": "schema-jwt_validation--jwks_config--cleartext", "description": "The JSON Web Key Set (JWKS) is a set of keys used to verify JSON Web Token (JWT) issued by the Authorization Server. See RFC 7517 for more details.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation:jwks_config", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["jwt_validation", "jwks_config", "cleartext"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/jwt_validation/jwks_config/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "The JSON Web Key Set (JWKS) is a set of keys used to verify JSON Web Token (JWT) issued by the Authorization Server. See RFC 7517 for more details.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
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
      "validatedAt": "2026-10-20T00:10:01+00:00"
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
