---
page_title: "a_pool.disable_health_check"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["a pool disable health check"], "body_bytes": 1246, "body_sha256": "sha256:3ba3ac4c2e957dad20190de9ddf0c8d29369ac29a8bf20df91f1c1acda471217", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:dns_lb_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_lb_pool:properties:a_pool:disable_health_check", "parent_id": "xcsh-docs:resources:dns_lb_pool:properties:a_pool", "path": "documentation/resources/dns_lb_pool/properties/a_pool/disable_health_check/index.md", "product": "distributed-cloud", "provider_name": "dns_lb_pool", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0210213002213000-2202232100103103-0110323232101021-3003200012313322-1310021301300000-1313202210221010-1332112122100030-3031332233202111", "registry_path": "docs/guides/resources--dns_lb_pool--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["a_pool", "disable_health_check"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_lb_pool/properties/a_pool/disable_health_check/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_lb_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# a_pool.disable_health_check

Breadcrumbs:

- [xcsh_dns_lb_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/)
- [a_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/a_pool/)
- a_pool.disable_health_check

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

- [a_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/a_pool/)
- [xcsh_dns_lb_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/)
