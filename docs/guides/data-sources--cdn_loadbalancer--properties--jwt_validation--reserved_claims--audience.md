---
page_title: "jwt_validation.reserved_claims.audience"
subcategory: "Load Balancing"
description: "jwt_validation.reserved_claims.audience for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2220, "body_sha256": "sha256:f6dc408dbc609b60c852aec16e8f4f64b3d5c920663810f43d73445a4b066d30", "canonical_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:jwt_validation:reserved_claims:audience", "child_ids": [], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:jwt_validation:reserved_claims:audience", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:jwt_validation:reserved_claims", "path": "docs/guides/data-sources--cdn_loadbalancer--properties--jwt_validation--reserved_claims--audience.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["jwt_validation", "reserved_claims", "audience"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/jwt_validation/reserved_claims/audience/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "jwt_validation.reserved_claims.audience for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# jwt_validation.reserved_claims.audience

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- [jwt_validation](data-sources--cdn_loadbalancer--properties--jwt_validation.md)
- [jwt_validation.reserved_claims](data-sources--cdn_loadbalancer--properties--jwt_validation--reserved_claims.md)
- jwt_validation.reserved_claims.audience

<a id="section"></a>

Type: `"single"`. Computed.

Audiences

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

<a id="schema-jwt_validation--reserved_claims--audience--audiences"></a>

### audiences property

Type: `["list", "string"]`. Computed.

Values. Configuration parameter for audiences

Upstream description:

Configuration parameter for audiences

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Next pages

- [jwt_validation.reserved_claims](data-sources--cdn_loadbalancer--properties--jwt_validation--reserved_claims.md)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
