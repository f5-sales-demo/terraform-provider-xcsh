---
page_title: "slow_ddos_mitigation.disable_request_timeout"
subcategory: "Load Balancing"
description: "slow_ddos_mitigation.disable_request_timeout for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1384, "body_sha256": "sha256:6b07edd087ce37b2b934638b833479af77ad3f778b3cb657e5037156a37a9ff6", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:slow_ddos_mitigation:disable_request_timeout", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:slow_ddos_mitigation", "path": "documentation/resources/http_loadbalancer/properties/slow_ddos_mitigation/disable_request_timeout/index.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["slow_ddos_mitigation", "disable_request_timeout"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/slow_ddos_mitigation/disable_request_timeout/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "slow_ddos_mitigation.disable_request_timeout for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# slow_ddos_mitigation.disable_request_timeout

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [slow_ddos_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/slow_ddos_mitigation/)
- slow_ddos_mitigation.disable_request_timeout

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable request timeout.

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
disable_request_timeout = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [slow_ddos_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/slow_ddos_mitigation/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
