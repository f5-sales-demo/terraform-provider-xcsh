---
page_title: "trusted_clients.http_header"
subcategory: "Load Balancing"
description: "Request header name and value pairs."
xcsh_docs: {"aliases": ["trusted clients http header"], "body_bytes": 1083, "body_sha256": "sha256:b6edaf912435ac45394e48d9596b7c739378561b4a4bc22a0a8f3be5016c4128", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:trusted_clients:http_header:headers"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:trusted_clients:http_header", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:trusted_clients", "path": "documentation/data-sources/cdn_loadbalancer/properties/trusted_clients/http_header/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-3331313011113302-2133211203211001-3331033213021103-3121012200312032-0111302102020100-2110033001021302-3231323300322103-3311322122012121", "registry_path": "docs/guides/data-sources--cdn_loadbalancer--reference--group-015.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["trusted_clients", "http_header"], "schema_version": 1, "sections": [{"aliases": ["trusted clients http header headers"], "anchor": "section", "description": "List of HTTP header name and value pairs.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:trusted_clients:http_header:headers", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["trusted_clients", "http_header", "headers"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/trusted_clients/http_header/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "Request header name and value pairs.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# trusted_clients.http_header

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- [trusted_clients](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/trusted_clients/)
- trusted_clients.http_header

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for http header.

Additional upstream details:

Request header name and value pairs.

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

- [headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/trusted_clients/http_header/headers/): complete subsection reference.
