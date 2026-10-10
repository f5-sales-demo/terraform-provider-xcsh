---
page_title: "https"
subcategory: "Load Balancing"
description: "Choice for selecting CDN Distribution with bring your own certificates."
xcsh_docs: {"aliases": ["https"], "body_bytes": 1507, "body_sha256": "sha256:12c1f911da7475b864a9c4b093bc90c763da7a0bbc10dca06d0794cb72a0709c", "capabilities": ["cdn", "load-balancing.tls"], "category": "cdn", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:https:tls_cert_options"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:https", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:reference", "path": "documentation/data-sources/cdn_loadbalancer/properties/https/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-3230312322320110-2123202322212223-0312012332030102-0320301210203212-1331020322112130-2221022302312033-0120321223220102-3121132202012233", "registry_path": "docs/guides/data-sources--cdn_loadbalancer--reference--group-011.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["https"], "schema_version": 1, "sections": [{"aliases": ["https add hsts"], "anchor": "schema-https--add_hsts", "description": "Add HTTP Strict-Transport-Security response header.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:https", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["https", "add_hsts"], "syntax": "attribute", "type": "bool"}, {"aliases": ["https http redirect"], "anchor": "schema-https--http_redirect", "description": "Redirect HTTP traffic to HTTPS.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:https", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["https", "http_redirect"], "syntax": "attribute", "type": "bool"}, {"aliases": ["https tls cert options"], "anchor": "section", "description": "TLS Certificate OPTIONS.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:https:tls_cert_options", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["https", "tls_cert_options"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/https/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Choice for selecting CDN Distribution with bring your own certificates.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- https

<a id="section"></a>

Type: `"single"`. Computed.

Choice for selecting CDN Distribution with bring your own certificates.

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

<a id="schema-https--add_hsts"></a>

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

<a id="schema-https--http_redirect"></a>

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

- [tls_cert_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/https/tls_cert_options/): complete subsection reference.
