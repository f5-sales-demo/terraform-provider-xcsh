---
page_title: "use_tls.use_mtls.tls_certificates.disable_ocsp_stapling"
subcategory: "Load Balancing"
description: "use_tls.use_mtls.tls_certificates.disable_ocsp_stapling for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 1553, "body_sha256": "sha256:6f80e9a4d021ccca7cf973cac97c7aa621c33f36b8beaa0fa646f43aca3cb495", "child_ids": [], "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:origin_pool:properties:use_tls:use_mtls:tls_certificates:disable_ocsp_stapling", "parent_id": "xcsh-docs:resources:origin_pool:properties:use_tls:use_mtls:tls_certificates", "path": "documentation/resources/origin_pool/properties/use_tls/use_mtls/tls_certificates/disable_ocsp_stapling/index.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["use_tls", "use_mtls", "tls_certificates", "disable_ocsp_stapling"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/properties/use_tls/use_mtls/tls_certificates/disable_ocsp_stapling/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "use_tls.use_mtls.tls_certificates.disable_ocsp_stapling for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# use_tls.use_mtls.tls_certificates.disable_ocsp_stapling

Breadcrumbs:

- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/)
- [use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/)
- [use_tls.use_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/use_mtls/)
- [use_tls.use_mtls.tls_certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/use_mtls/tls_certificates/)
- use_tls.use_mtls.tls_certificates.disable_ocsp_stapling

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable ocsp stapling.

Upstream description:

This can be used for messages where no values are needed.

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
disable_ocsp_stapling = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [use_tls.use_mtls.tls_certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/use_mtls/tls_certificates/)
- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/)
