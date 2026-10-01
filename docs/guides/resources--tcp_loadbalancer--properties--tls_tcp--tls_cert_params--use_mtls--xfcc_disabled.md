---
page_title: "tls_tcp.tls_cert_params.use_mtls.xfcc_disabled"
subcategory: "Load Balancing"
description: "tls_tcp.tls_cert_params.use_mtls.xfcc_disabled for xcsh_tcp_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1295, "body_sha256": "sha256:70f62397a7e3c6ba949c8ed7d029ef80c64acd8de5ae926f7bcacbc06dcd7f78", "canonical_id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_cert_params:use_mtls:xfcc_disabled", "child_ids": [], "collection_id": "xcsh-docs:resources:tcp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_cert_params:use_mtls:xfcc_disabled", "parent_id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_cert_params:use_mtls", "path": "docs/guides/resources--tcp_loadbalancer--properties--tls_tcp--tls_cert_params--use_mtls--xfcc_disabled.md", "provider_name": "tcp_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tls_tcp", "tls_cert_params", "use_mtls", "xfcc_disabled"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tcp_loadbalancer/properties/tls_tcp/tls_cert_params/use_mtls/xfcc_disabled/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tls_tcp.tls_cert_params.use_mtls.xfcc_disabled for xcsh_tcp_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tcp_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_tcp.tls_cert_params.use_mtls.xfcc_disabled

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md)
- [Property reference](resources--tcp_loadbalancer--reference.md)
- [tls_tcp](resources--tcp_loadbalancer--properties--tls_tcp.md)
- [tls_tcp.tls_cert_params](resources--tcp_loadbalancer--properties--tls_tcp--tls_cert_params.md)
- [tls_tcp.tls_cert_params.use_mtls](resources--tcp_loadbalancer--properties--tls_tcp--tls_cert_params--use_mtls.md)
- tls_tcp.tls_cert_params.use_mtls.xfcc_disabled

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
xfcc_disabled = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [tls_tcp.tls_cert_params.use_mtls](resources--tcp_loadbalancer--properties--tls_tcp--tls_cert_params--use_mtls.md)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md)
