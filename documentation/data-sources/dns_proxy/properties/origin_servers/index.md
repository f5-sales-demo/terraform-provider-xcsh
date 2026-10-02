---
page_title: "origin_servers"
subcategory: ""
description: "List of origin Servers for the DNS proxy."
xcsh_docs: {"aliases": ["backend servers", "origin servers", "upstream servers"], "body_bytes": 1541, "body_sha256": "sha256:a24e3803cd8201a1e3928d0db541b08b4fd96d3c9e3293bfedde577140d307de", "capabilities": ["dns", "load-balancing.backend-servers"], "category": "dns", "child_ids": ["xcsh-docs:data-sources:dns_proxy:properties:origin_servers:health_checks", "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers", "parent_id": "xcsh-docs:data-sources:dns_proxy:reference", "path": "documentation/data-sources/dns_proxy/properties/origin_servers/index.md", "product": "distributed-cloud", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3102221313302120-2300133213200331-0320001332113013-1100133203331131-3231110233212130-1312212033030000-3302023133321213-2023021333201010", "registry_path": "docs/guides/data-sources--dns_proxy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_servers"], "schema_version": 1, "sections": [{"aliases": ["backend servers", "health checks", "origin servers", "upstream servers"], "anchor": "section", "description": "Origin Server Health Checks.", "document_id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:health_checks", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_servers", "health_checks"], "syntax": "attribute", "type": "object"}, {"aliases": ["backend servers", "origin servers", "upstream servers"], "anchor": "section", "description": "List of origin servers for Proxy.", "document_id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["origin_servers", "origin_servers"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_proxy/properties/origin_servers/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "List of origin Servers for the DNS proxy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_servers

Breadcrumbs:

- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/)
- origin_servers

<a id="section"></a>

Type: `"single"`. Computed.

List of origin Servers for the DNS proxy.

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

- [health_checks](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/health_checks/): complete subsection reference.

- [origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/): complete subsection reference.

## Next pages

- [origin_servers.health_checks](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/health_checks/)
- [origin_servers.origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/)
- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/)
