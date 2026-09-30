---
page_title: "use_tls.use_mtls"
subcategory: "Load Balancing"
description: "use_tls.use_mtls for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 1207, "body_sha256": "sha256:a7609dd43a8b4134fafd6a6c8cd462f3ae387a1b6ea25bc04f5dc4a42b256a33", "canonical_id": "xcsh-docs:resources:origin_pool:properties:use_tls:use_mtls", "child_ids": ["xcsh-docs:resources:origin_pool:properties:use_tls:use_mtls:tls_certificates"], "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:origin_pool:properties:use_tls:use_mtls", "parent_id": "xcsh-docs:resources:origin_pool:properties:use_tls", "path": "docs/guides/resources--origin_pool--properties--use_tls--use_mtls.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["use_tls", "use_mtls"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/properties/use_tls/use_mtls/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "use_tls.use_mtls for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# use_tls.use_mtls

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md)
- [Property reference](resources--origin_pool--reference.md)
- [use_tls](resources--origin_pool--properties--use_tls.md)
- use_tls.use_mtls

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

- [tls_certificates](resources--origin_pool--properties--use_tls--use_mtls--tls_certificates.md): complete subsection reference.

## Next pages

- [use_tls.use_mtls.tls_certificates](resources--origin_pool--properties--use_tls--use_mtls--tls_certificates.md)
- [use_tls](resources--origin_pool--properties--use_tls.md)
- [xcsh_origin_pool](../resources/origin_pool.md)
