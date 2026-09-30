---
page_title: "more_option.response_cookies_to_add.ignore_partitioned"
subcategory: "Load Balancing"
description: "more_option.response_cookies_to_add.ignore_partitioned for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1174, "body_sha256": "sha256:88d80db135db92bb54e6afc8603d0c9868e13eb9d1fc900b15d40949f6ac1f25", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:more_option:response_cookies_to_add:ignore_partitioned", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:more_option:response_cookies_to_add:ignore_partitioned", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:more_option:response_cookies_to_add", "path": "docs/guides/resources--http_loadbalancer--properties--more_option--response_cookies_to_add--ignore_partitioned.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["more_option", "response_cookies_to_add", "ignore_partitioned"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/more_option/response_cookies_to_add/ignore_partitioned/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "more_option.response_cookies_to_add.ignore_partitioned for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# more_option.response_cookies_to_add.ignore_partitioned

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [more_option](resources--http_loadbalancer--properties--more_option.md)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--properties--more_option--response_cookies_to_add.md)
- more_option.response_cookies_to_add.ignore_partitioned

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore partitioned.

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
ignore_partitioned = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [more_option.response_cookies_to_add](resources--http_loadbalancer--properties--more_option--response_cookies_to_add.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
