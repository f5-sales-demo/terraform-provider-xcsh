---
page_title: "https"
subcategory: "Load Balancing"
description: "https for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1653, "body_sha256": "sha256:4a06ac4ea36ddd1eeb4d127519a4758ca21c1705826c003c5eead1e0af694e49", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:https", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:https:tls_cert_options"], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:https", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:reference", "path": "docs/guides/resources--cdn_loadbalancer--properties--https.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["https"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/https/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "https for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# https

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
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

- [tls_cert_options](resources--cdn_loadbalancer--properties--https--tls_cert_options.md): complete subsection reference.

## Next pages

- [https.tls_cert_options](resources--cdn_loadbalancer--properties--https--tls_cert_options.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
