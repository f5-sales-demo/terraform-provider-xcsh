---
page_title: "more_option.response_cookies_to_add.samesite_strict"
subcategory: "Load Balancing"
description: "more_option.response_cookies_to_add.samesite_strict for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1235, "body_sha256": "sha256:6482c59aa86a926bf92546475a4a2dceb56bade2eaef714c8bc96d70cc573d61", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:more_option:response_cookies_to_add:samesite_strict", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:more_option:response_cookies_to_add:samesite_strict", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:more_option:response_cookies_to_add", "path": "docs/guides/resources--http_loadbalancer--properties--more_option--response_cookies_to_add--samesite_strict.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["more_option", "response_cookies_to_add", "samesite_strict"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/more_option/response_cookies_to_add/samesite_strict/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "more_option.response_cookies_to_add.samesite_strict for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# more_option.response_cookies_to_add.samesite_strict

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [more_option](resources--http_loadbalancer--properties--more_option.md)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--properties--more_option--response_cookies_to_add.md)
- more_option.response_cookies_to_add.samesite_strict

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
samesite_strict = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [more_option.response_cookies_to_add](resources--http_loadbalancer--properties--more_option--response_cookies_to_add.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
