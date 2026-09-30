---
page_title: "jwt_validation.target.base_paths"
subcategory: "Load Balancing"
description: "jwt_validation.target.base_paths for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2185, "body_sha256": "sha256:7caedb79ad709dbd8fe6ad3227bb731c7ad5eef2df4a86489587ae357f9b5fc8", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation:target:base_paths", "child_ids": [], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation:target:base_paths", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation:target", "path": "docs/guides/data-sources--http_loadbalancer--properties--jwt_validation--target--base_paths.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["jwt_validation", "target", "base_paths"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/jwt_validation/target/base_paths/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "jwt_validation.target.base_paths for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# jwt_validation.target.base_paths

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [jwt_validation](data-sources--http_loadbalancer--properties--jwt_validation.md)
- [jwt_validation.target](data-sources--http_loadbalancer--properties--jwt_validation--target.md)
- jwt_validation.target.base_paths

<a id="section"></a>

Type: `"single"`. Computed.

Base Paths.

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

<a id="schema-jwt_validation--target--base_paths--base_paths"></a>

### base_paths property

Type: `["list", "string"]`. Computed.

Prefix Values. File system or URL path

Upstream description:

File system or URL path

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
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
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Next pages

- [jwt_validation.target](data-sources--http_loadbalancer--properties--jwt_validation--target.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
