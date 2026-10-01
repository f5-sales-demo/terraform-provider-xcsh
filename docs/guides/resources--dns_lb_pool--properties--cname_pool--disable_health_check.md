---
page_title: "cname_pool.disable_health_check"
subcategory: ""
description: "cname_pool.disable_health_check for xcsh_dns_lb_pool."
xcsh_docs: {"aliases": [], "body_bytes": 1013, "body_sha256": "sha256:d52bba9736a0c9e0476772f94fc4d3b36b9ae9990a2851cb593a8fe4786e064e", "canonical_id": "xcsh-docs:resources:dns_lb_pool:properties:cname_pool:disable_health_check", "child_ids": [], "collection_id": "xcsh-docs:resources:dns_lb_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_lb_pool:properties:cname_pool:disable_health_check", "parent_id": "xcsh-docs:resources:dns_lb_pool:properties:cname_pool", "path": "docs/guides/resources--dns_lb_pool--properties--cname_pool--disable_health_check.md", "provider_name": "dns_lb_pool", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cname_pool", "disable_health_check"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_lb_pool/properties/cname_pool/disable_health_check/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cname_pool.disable_health_check for xcsh_dns_lb_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_lb_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cname_pool.disable_health_check

Breadcrumbs:

- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md)
- [Property reference](resources--dns_lb_pool--reference.md)
- [cname_pool](resources--dns_lb_pool--properties--cname_pool.md)
- cname_pool.disable_health_check

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable health check.

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
disable_health_check = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [cname_pool](resources--dns_lb_pool--properties--cname_pool.md)
- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md)
