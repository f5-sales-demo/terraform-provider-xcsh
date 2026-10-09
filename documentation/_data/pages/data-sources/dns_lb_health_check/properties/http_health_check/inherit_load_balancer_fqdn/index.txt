---
page_title: "http_health_check.inherit_load_balancer_fqdn"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["http health check inherit load balancer fqdn"], "body_bytes": 1045, "body_sha256": "sha256:d5a8c3108adca60c29feee63f1b597b209659209295cfc106b4631ae7b03aadc", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:dns_lb_health_check:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_lb_health_check:properties:http_health_check:inherit_load_balancer_fqdn", "parent_id": "xcsh-docs:data-sources:dns_lb_health_check:properties:http_health_check", "path": "documentation/data-sources/dns_lb_health_check/properties/http_health_check/inherit_load_balancer_fqdn/index.md", "product": "distributed-cloud", "provider_name": "dns_lb_health_check", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-0330332110210230-2100032203010033-0123232220110122-0102333322301132-1210323100210233-2002210103320320-0001331031312211-1010102203123131", "registry_path": "docs/guides/data-sources--dns_lb_health_check--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["http_health_check", "inherit_load_balancer_fqdn"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_lb_health_check/properties/http_health_check/inherit_load_balancer_fqdn/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["dns_lb_health_checkCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
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

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.
