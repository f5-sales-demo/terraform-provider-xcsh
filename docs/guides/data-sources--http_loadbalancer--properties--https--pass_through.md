---
page_title: "https.pass_through"
subcategory: "Load Balancing"
description: "https.pass_through for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 854, "body_sha256": "sha256:d16e6ebe9e5b1b39a0c3839ada4ffdd0d70243610637bbed30d81721902324ef", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:https:pass_through", "child_ids": [], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:https:pass_through", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:https", "path": "docs/guides/data-sources--http_loadbalancer--properties--https--pass_through.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["https", "pass_through"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/https/pass_through/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "https.pass_through for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# https.pass_through

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [https](data-sources--http_loadbalancer--properties--https.md)
- https.pass_through

<a id="section"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for pass through.

Upstream description:

This can be used for messages where no values are needed.

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

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [https](data-sources--http_loadbalancer--properties--https.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
