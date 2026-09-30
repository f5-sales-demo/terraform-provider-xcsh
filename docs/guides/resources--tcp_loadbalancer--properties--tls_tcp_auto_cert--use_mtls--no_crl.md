---
page_title: "tls_tcp_auto_cert.use_mtls.no_crl"
subcategory: "Load Balancing"
description: "tls_tcp_auto_cert.use_mtls.no_crl for xcsh_tcp_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1059, "body_sha256": "sha256:8ad26bb581c05358034518185c4f7afcb5f97e89e66c5bf379b5cc087525730e", "canonical_id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp_auto_cert:use_mtls:no_crl", "child_ids": [], "collection_id": "xcsh-docs:resources:tcp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp_auto_cert:use_mtls:no_crl", "parent_id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp_auto_cert:use_mtls", "path": "docs/guides/resources--tcp_loadbalancer--properties--tls_tcp_auto_cert--use_mtls--no_crl.md", "provider_name": "tcp_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tls_tcp_auto_cert", "use_mtls", "no_crl"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tcp_loadbalancer/properties/tls_tcp_auto_cert/use_mtls/no_crl/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tls_tcp_auto_cert.use_mtls.no_crl for xcsh_tcp_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tcp_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# tls_tcp_auto_cert.use_mtls.no_crl

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md)
- [Property reference](resources--tcp_loadbalancer--reference.md)
- [tls_tcp_auto_cert](resources--tcp_loadbalancer--properties--tls_tcp_auto_cert.md)
- [tls_tcp_auto_cert.use_mtls](resources--tcp_loadbalancer--properties--tls_tcp_auto_cert--use_mtls.md)
- tls_tcp_auto_cert.use_mtls.no_crl

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
no_crl = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [tls_tcp_auto_cert.use_mtls](resources--tcp_loadbalancer--properties--tls_tcp_auto_cert--use_mtls.md)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md)
