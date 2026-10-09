---
page_title: "https_auto_cert"
subcategory: "Load Balancing"
description: "Choice for selecting HTTPS CDN distribution with bring your own certificates."
xcsh_docs: {"aliases": ["automatic certificate management", "automatic certificates", "https auto cert", "managed TLS certificates"], "body_bytes": 1551, "body_sha256": "sha256:0cd813fa889738cfa752380c6576bf09e7b82708240c470ab68a63d02ac289bd", "capabilities": ["cdn", "load-balancing.tls"], "category": "cdn", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:https_auto_cert:tls_config"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:https_auto_cert", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:reference", "path": "documentation/data-sources/cdn_loadbalancer/properties/https_auto_cert/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-1022110223113311-1122120102013303-2121023120231002-0103032202213002-0201010020313211-3100130013113312-0132133320111300-3120303011211300", "registry_path": "docs/guides/data-sources--cdn_loadbalancer--reference--group-012.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["https_auto_cert"], "schema_version": 1, "sections": [{"aliases": ["https auto cert add hsts"], "anchor": "schema-https_auto_cert--add_hsts", "description": "Add HTTP Strict-Transport-Security response header.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:https_auto_cert", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["https_auto_cert", "add_hsts"], "syntax": "attribute", "type": "bool"}, {"aliases": ["https auto cert http redirect"], "anchor": "schema-https_auto_cert--http_redirect", "description": "Redirect HTTP traffic to HTTPS.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:https_auto_cert", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["https_auto_cert", "http_redirect"], "syntax": "attribute", "type": "bool"}, {"aliases": ["https auto cert tls config"], "anchor": "section", "description": "This defines various OPTIONS to configure TLS configuration parameters.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:https_auto_cert:tls_config", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["https_auto_cert", "tls_config"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/https_auto_cert/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Choice for selecting HTTPS CDN distribution with bring your own certificates.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https_auto_cert

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- https_auto_cert

<a id="section"></a>

Type: `"single"`. Computed.

Choice for selecting HTTPS CDN distribution with bring your own certificates.

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

<a id="schema-https_auto_cert--add_hsts"></a>

### add_hsts property

Type: `"bool"`. Computed.

Add HTTP Strict-Transport-Security response header.

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

<a id="schema-https_auto_cert--http_redirect"></a>

### http_redirect property

Type: `"bool"`. Computed.

HTTP Redirect to HTTPS. Redirect HTTP traffic to HTTPS.

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

- [tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/https_auto_cert/tls_config/): complete subsection reference.
