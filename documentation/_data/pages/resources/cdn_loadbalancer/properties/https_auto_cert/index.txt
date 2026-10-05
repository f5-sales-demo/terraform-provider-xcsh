---
page_title: "https_auto_cert"
subcategory: "Load Balancing"
description: "Choice for selecting HTTPS CDN distribution with bring your own certificates."
xcsh_docs: {"aliases": ["automatic certificate management", "automatic certificates", "https auto cert", "managed TLS certificates"], "body_bytes": 2122, "body_sha256": "sha256:912e6993afc15f476d092d948a3cb6b3504c58587eadc6a1c71ba088181187e9", "capabilities": ["cdn", "load-balancing.tls"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:https_auto_cert:tls_config"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:https_auto_cert", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:reference", "path": "documentation/resources/cdn_loadbalancer/properties/https_auto_cert/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-2333321023320231-1103212001100120-1332321202002133-1230013202111101-3203220322212103-3223230223121121-1120131300100331-2001211211221323", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-011.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["https_auto_cert"], "schema_version": 1, "sections": [{"aliases": ["https auto cert add hsts"], "anchor": "schema-https_auto_cert--add_hsts", "description": "Add HTTP Strict-Transport-Security response header.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:https_auto_cert", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["https_auto_cert", "add_hsts"], "syntax": "attribute", "type": "bool"}, {"aliases": ["https auto cert http redirect"], "anchor": "schema-https_auto_cert--http_redirect", "description": "Redirect HTTP traffic to HTTPS.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:https_auto_cert", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["https_auto_cert", "http_redirect"], "syntax": "attribute", "type": "bool"}, {"aliases": ["https auto cert tls config"], "anchor": "section", "description": "This defines various OPTIONS to configure TLS configuration parameters.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:https_auto_cert:tls_config", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "https_auto_cert.tls_config:ConflictingObjectAttributes:tls_11_plus,tls_12_plus", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:https_auto_cert:tls_config:tls_11_plus", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "https_auto_cert.tls_config:ConflictingObjectAttributes:tls_11_plus,tls_12_plus", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:https_auto_cert:tls_config:tls_12_plus", "type": "conflicts"}], "schema_path": ["https_auto_cert", "tls_config"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/https_auto_cert/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Choice for selecting HTTPS CDN distribution with bring your own certificates.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https_auto_cert

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- https_auto_cert

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
https_auto_cert {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-https_auto_cert--add_hsts"></a>

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

<a id="schema-https_auto_cert--http_redirect"></a>

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

- [tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/https_auto_cert/tls_config/): complete subsection reference.

## Next pages

- [https_auto_cert.tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/https_auto_cert/tls_config/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
