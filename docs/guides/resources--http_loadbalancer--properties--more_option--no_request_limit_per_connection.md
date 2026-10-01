---
page_title: "more_option.no_request_limit_per_connection"
subcategory: "Load Balancing"
description: "more_option.no_request_limit_per_connection for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1105, "body_sha256": "sha256:22645c43034b712ad8834cbd355da70cdbbf51e729c33bdb747829b972ac3d8b", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:more_option:no_request_limit_per_connection", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:more_option:no_request_limit_per_connection", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:more_option", "path": "docs/guides/resources--http_loadbalancer--properties--more_option--no_request_limit_per_connection.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["more_option", "no_request_limit_per_connection"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/more_option/no_request_limit_per_connection/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "more_option.no_request_limit_per_connection for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# more_option.no_request_limit_per_connection

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [more_option](resources--http_loadbalancer--properties--more_option.md)
- more_option.no_request_limit_per_connection

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no request limit per connection.

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

Terraform syntax:

```terraform
no_request_limit_per_connection = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [more_option](resources--http_loadbalancer--properties--more_option.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
