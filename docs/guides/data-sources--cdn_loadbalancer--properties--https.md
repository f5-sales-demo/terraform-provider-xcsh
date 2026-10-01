---
page_title: "https"
subcategory: "Load Balancing"
description: "https for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1656, "body_sha256": "sha256:4a56d1578d69c609d217004a1baf5c7ba8c209c8e62790a89b66711e131b9043", "canonical_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:https", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:https:tls_cert_options"], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:https", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:reference", "path": "docs/guides/data-sources--cdn_loadbalancer--properties--https.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["https"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/https/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "https for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
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

- [tls_cert_options](data-sources--cdn_loadbalancer--properties--https--tls_cert_options.md): complete subsection reference.

## Next pages

- [https.tls_cert_options](data-sources--cdn_loadbalancer--properties--https--tls_cert_options.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
