---
page_title: "origin_pool.use_tls.use_mtls"
subcategory: "Load Balancing"
description: "origin_pool.use_tls.use_mtls for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1437, "body_sha256": "sha256:f77e3075843d6edebf08341e141655ac1fe0aafecebb2c0e48f32a477d63107f", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates"], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls", "path": "docs/guides/resources--cdn_loadbalancer--properties--origin_pool--use_tls--use_mtls.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["origin_pool", "use_tls", "use_mtls"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/origin_pool/use_tls/use_mtls/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "origin_pool.use_tls.use_mtls for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# origin_pool.use_tls.use_mtls

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [origin_pool](resources--cdn_loadbalancer--properties--origin_pool.md)
- [origin_pool.use_tls](resources--cdn_loadbalancer--properties--origin_pool--use_tls.md)
- origin_pool.use_tls.use_mtls

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

MTLS Certificate. MTLS Client Certificate.

Upstream description:

MTLS Client Certificate.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("tls_certificates")}
```

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
use_mtls {
  # Configure direct properties listed below.
}
```

## Direct properties

- [tls_certificates](resources--cdn_loadbalancer--properties--origin_pool--use_tls--use_mtls--tls_certificates.md): complete subsection reference.

## Next pages

- [origin_pool.use_tls.use_mtls.tls_certificates](resources--cdn_loadbalancer--properties--origin_pool--use_tls--use_mtls--tls_certificates.md)
- [origin_pool.use_tls](resources--cdn_loadbalancer--properties--origin_pool--use_tls.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
