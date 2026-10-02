---
page_title: "enable_api_discovery.api_discovery_from_code_scan"
subcategory: "Load Balancing"
description: "Select Code Base and Repositories."
xcsh_docs: {"aliases": ["enable api discovery api discovery from code scan"], "body_bytes": 1648, "body_sha256": "sha256:b668de3e2bfbbc94167d86b5a30125dc0d4791d0cb1a7312f80aeae8e2a840f4", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:enable_api_discovery:api_discovery_from_code_scan:code_base_integrations"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:enable_api_discovery:api_discovery_from_code_scan", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:enable_api_discovery", "path": "documentation/data-sources/http_loadbalancer/properties/enable_api_discovery/api_discovery_from_code_scan/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2203231110333303-0020330220102233-1212031302200323-2010132310300200-2323303123313032-3023312023010120-1132113021000231-1003103022033213", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-017.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["enable_api_discovery", "api_discovery_from_code_scan"], "schema_version": 1, "sections": [{"aliases": ["code base integrations"], "anchor": "section", "description": "Configuration parameter for code base integrations", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:enable_api_discovery:api_discovery_from_code_scan:code_base_integrations", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["enable_api_discovery", "api_discovery_from_code_scan", "code_base_integrations"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/enable_api_discovery/api_discovery_from_code_scan/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Select Code Base and Repositories.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_api_discovery.api_discovery_from_code_scan

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [enable_api_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/enable_api_discovery/)
- enable_api_discovery.api_discovery_from_code_scan

<a id="section"></a>

Type: `"single"`. Computed.

Select Code Base and Repositories.

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

- [code_base_integrations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/enable_api_discovery/api_discovery_from_code_scan/code_base_integrations/): complete subsection reference.

## Next pages

- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/enable_api_discovery/api_discovery_from_code_scan/code_base_integrations/)
- [enable_api_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/enable_api_discovery/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
