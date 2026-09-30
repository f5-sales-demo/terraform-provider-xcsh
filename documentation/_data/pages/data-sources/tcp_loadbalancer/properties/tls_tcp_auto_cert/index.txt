---
page_title: "tls_tcp_auto_cert"
subcategory: "Load Balancing"
description: "tls_tcp_auto_cert for xcsh_tcp_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1917, "body_sha256": "sha256:35cf6433248e5ebe42f9095dce6f740022daedd97e0970c3b48638a8046efadc", "child_ids": ["xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp_auto_cert:no_mtls", "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp_auto_cert:tls_config", "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp_auto_cert:use_mtls"], "collection_id": "xcsh-docs:data-sources:tcp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp_auto_cert", "parent_id": "xcsh-docs:data-sources:tcp_loadbalancer:reference", "path": "documentation/data-sources/tcp_loadbalancer/properties/tls_tcp_auto_cert/index.md", "provider_name": "tcp_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["tls_tcp_auto_cert"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tcp_loadbalancer/properties/tls_tcp_auto_cert/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tls_tcp_auto_cert for xcsh_tcp_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tcp_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# tls_tcp_auto_cert

Breadcrumbs:

- [xcsh_tcp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/)
- tls_tcp_auto_cert

<a id="section"></a>

Type: `"single"`. Computed.

Choice for selecting TLS over TCP proxy with automatic certificates.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

## Direct properties

- [no_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/tls_tcp_auto_cert/no_mtls/): complete subsection reference.

- [tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/tls_tcp_auto_cert/tls_config/): complete subsection reference.

- [use_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/tls_tcp_auto_cert/use_mtls/): complete subsection reference.

## Next pages

- [tls_tcp_auto_cert.no_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/tls_tcp_auto_cert/no_mtls/)
- [tls_tcp_auto_cert.tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/tls_tcp_auto_cert/tls_config/)
- [tls_tcp_auto_cert.use_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/tls_tcp_auto_cert/use_mtls/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/)
- [xcsh_tcp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/)
