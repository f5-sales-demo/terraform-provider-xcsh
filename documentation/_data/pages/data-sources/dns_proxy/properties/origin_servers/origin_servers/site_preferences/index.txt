---
page_title: "origin_servers.origin_servers.site_preferences"
subcategory: ""
description: "Carries the references to one or more sites."
xcsh_docs: {"aliases": ["origin servers origin servers site preferences"], "body_bytes": 1661, "body_sha256": "sha256:02920d0e77ef146f2cdbc6f0b7e77bc7d92d633ac16457e3a836906d34d3fead", "capabilities": ["dns", "load-balancing.backend-servers"], "category": "dns", "child_ids": ["xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:site_preferences:refs"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:site_preferences", "parent_id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers", "path": "documentation/data-sources/dns_proxy/properties/origin_servers/origin_servers/site_preferences/index.md", "product": "distributed-cloud", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-2030003201020323-1323020011011112-2000103202210133-0303032330110302-0332002331302301-3130130232111302-3222033101030313-1331120330031102", "registry_path": "docs/guides/data-sources--dns_proxy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_servers", "origin_servers", "site_preferences"], "schema_version": 1, "sections": [{"aliases": ["origin servers origin servers site preferences refs"], "anchor": "section", "description": "Reference to one or more sites.", "document_id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:site_preferences:refs", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["origin_servers", "origin_servers", "site_preferences", "refs"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_proxy/properties/origin_servers/origin_servers/site_preferences/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Carries the references to one or more sites.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_servers.origin_servers.site_preferences

Breadcrumbs:

- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/)
- [origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/)
- [origin_servers.origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/)
- origin_servers.origin_servers.site_preferences

<a id="section"></a>

Type: `"single"`. Computed.

Carries the references to one or more sites.

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

- [refs](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/site_preferences/refs/): complete subsection reference.

## Next pages

- [origin_servers.origin_servers.site_preferences.refs](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/site_preferences/refs/)
- [origin_servers.origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/)
- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/)
