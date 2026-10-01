---
page_title: "icmp_health_check"
subcategory: ""
description: "icmp_health_check for xcsh_dns_lb_health_check."
xcsh_docs: {"aliases": [], "body_bytes": 956, "body_sha256": "sha256:1ba91ccd2e7e84dc0edc771809fbb5632833e913eb026253f9ce755371c00159", "canonical_id": "xcsh-docs:resources:dns_lb_health_check:properties:icmp_health_check", "child_ids": [], "collection_id": "xcsh-docs:resources:dns_lb_health_check:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_lb_health_check:properties:icmp_health_check", "parent_id": "xcsh-docs:resources:dns_lb_health_check:reference", "path": "docs/guides/resources--dns_lb_health_check--properties--icmp_health_check.md", "provider_name": "dns_lb_health_check", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["icmp_health_check"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_lb_health_check/properties/icmp_health_check/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "icmp_health_check for xcsh_dns_lb_health_check.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_lb_health_checkCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# icmp_health_check

Breadcrumbs:

- [xcsh_dns_lb_health_check](../resources/dns_lb_health_check.md)
- [Property reference](resources--dns_lb_health_check--reference.md)
- icmp_health_check

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for icmp health check.

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
icmp_health_check = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--dns_lb_health_check--reference.md)
- [xcsh_dns_lb_health_check](../resources/dns_lb_health_check.md)
