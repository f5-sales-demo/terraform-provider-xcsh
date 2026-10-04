---
page_title: "https"
subcategory: "Load Balancing"
description: "Choice for selecting CDN Distribution with bring your own certificates."
xcsh_docs: {"aliases": ["https"], "body_bytes": 2060, "body_sha256": "sha256:d0bf276e9a25fc13dfde4f2b66a7217d97c8e321b003145ff72f3ecb9edf4eb0", "capabilities": ["cdn", "load-balancing.tls"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:https:tls_cert_options"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:https", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:reference", "path": "documentation/resources/cdn_loadbalancer/properties/https/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2101233301010003-1200221022012203-0031300003110020-0322211021320210-2003201303102013-1022203021201030-0102222103131230-0232323033112111", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-011.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["https"], "schema_version": 1, "sections": [{"aliases": ["https add hsts"], "anchor": "schema-https--add_hsts", "description": "Add HTTP Strict-Transport-Security response header.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:https", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["https", "add_hsts"], "syntax": "attribute", "type": "bool"}, {"aliases": ["https http redirect"], "anchor": "schema-https--http_redirect", "description": "Redirect HTTP traffic to HTTPS.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:https", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["https", "http_redirect"], "syntax": "attribute", "type": "bool"}, {"aliases": ["https tls cert options"], "anchor": "section", "description": "TLS Certificate OPTIONS.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:https:tls_cert_options", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "https.tls_cert_options:ConflictingObjectAttributes:tls_cert_params,tls_inline_params", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:https:tls_cert_options:tls_cert_params", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "https.tls_cert_options:ConflictingObjectAttributes:tls_cert_params,tls_inline_params", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:https:tls_cert_options:tls_inline_params", "type": "conflicts"}], "schema_path": ["https", "tls_cert_options"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/https/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Choice for selecting CDN Distribution with bring your own certificates.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- https

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
https {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-https--add_hsts"></a>

### add_hsts property

Type: `"bool"`. Optional.

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

Type: `"bool"`. Optional.

HTTP Redirect to HTTPS. Redirect HTTP traffic to HTTPS.

Upstream description:

Redirect HTTP traffic to HTTPS.

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

- [tls_cert_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/https/tls_cert_options/): complete subsection reference.

## Next pages

- [https.tls_cert_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/https/tls_cert_options/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
