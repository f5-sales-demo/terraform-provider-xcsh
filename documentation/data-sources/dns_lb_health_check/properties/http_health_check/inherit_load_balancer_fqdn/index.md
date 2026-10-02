---
page_title: "http_health_check.inherit_load_balancer_fqdn"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["http health check inherit load balancer fqdn"], "body_bytes": 1338, "body_sha256": "sha256:e0f0df6113ae0bac674b54feff7984227c75f1425290b804f4c6260d81a60f66", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:dns_lb_health_check:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_lb_health_check:properties:http_health_check:inherit_load_balancer_fqdn", "parent_id": "xcsh-docs:data-sources:dns_lb_health_check:properties:http_health_check", "path": "documentation/data-sources/dns_lb_health_check/properties/http_health_check/inherit_load_balancer_fqdn/index.md", "product": "distributed-cloud", "provider_name": "dns_lb_health_check", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0330332110210230-2100032203010033-0123232220110122-0102333322301132-1210323100210233-2002210103320320-0001331031312211-1010102203123131", "registry_path": "docs/guides/data-sources--dns_lb_health_check--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["http_health_check", "inherit_load_balancer_fqdn"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_lb_health_check/properties/http_health_check/inherit_load_balancer_fqdn/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_lb_health_checkCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http_health_check.inherit_load_balancer_fqdn

Breadcrumbs:

- [xcsh_dns_lb_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/)
- [http_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/http_health_check/)
- http_health_check.inherit_load_balancer_fqdn

<a id="section"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for inherit load balancer fqdn.

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

- [http_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/http_health_check/)
- [xcsh_dns_lb_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/)
