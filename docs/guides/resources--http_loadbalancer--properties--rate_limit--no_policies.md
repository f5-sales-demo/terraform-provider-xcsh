---
page_title: "rate_limit.no_policies"
subcategory: "Load Balancing"
description: "rate_limit.no_policies for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1091, "body_sha256": "sha256:cfd4661f95543f280070c956805c5430c65ef2a932cc81ea469425107ea0c925", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:rate_limit:no_policies", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:rate_limit:no_policies", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:rate_limit", "path": "docs/guides/resources--http_loadbalancer--properties--rate_limit--no_policies.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rate_limit", "no_policies"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/rate_limit/no_policies/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rate_limit.no_policies for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rate_limit.no_policies

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [rate_limit](resources--http_loadbalancer--properties--rate_limit.md)
- rate_limit.no_policies

<a id="section"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for no policies. Defaults to \`map\[\]\`. Server applies default when
omitted.

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
no_policies = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [rate_limit](resources--http_loadbalancer--properties--rate_limit.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
