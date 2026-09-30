---
page_title: "multi_lb_app"
subcategory: "Load Balancing"
description: "multi_lb_app for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1107, "body_sha256": "sha256:e326b156b41247c9b5d8190dc3be845deff63c86b50cfa88a2a3e3415c6dd9ab", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:multi_lb_app", "child_ids": [], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:multi_lb_app", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:reference", "path": "docs/guides/data-sources--http_loadbalancer--properties--multi_lb_app.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["multi_lb_app"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/multi_lb_app/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "multi_lb_app for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# multi_lb_app

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- multi_lb_app

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: multi\_lb\_app, single\_lb\_app\] Configuration parameter for multi lb app.

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

OneOf alternatives in this subsection:

- [multi_lb_app](data-sources--http_loadbalancer--properties--multi_lb_app.md#section)
- [single_lb_app](data-sources--http_loadbalancer--properties--single_lb_app.md#section)

Select alternatives according to the provider validators above.

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](data-sources--http_loadbalancer--reference.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
