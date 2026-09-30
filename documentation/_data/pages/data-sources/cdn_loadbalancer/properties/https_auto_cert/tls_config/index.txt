---
page_title: "https_auto_cert.tls_config"
subcategory: "Load Balancing"
description: "https_auto_cert.tls_config for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1924, "body_sha256": "sha256:f969b59d91fa00352a60553c7f80d9303403339fd4948f5058d02cc0ab80df73", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:https_auto_cert:tls_config:tls_11_plus", "xcsh-docs:data-sources:cdn_loadbalancer:properties:https_auto_cert:tls_config:tls_12_plus"], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:https_auto_cert:tls_config", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:https_auto_cert", "path": "documentation/data-sources/cdn_loadbalancer/properties/https_auto_cert/tls_config/index.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["https_auto_cert", "tls_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/https_auto_cert/tls_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "https_auto_cert.tls_config for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# https_auto_cert.tls_config

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- [https_auto_cert](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/https_auto_cert/)
- https_auto_cert.tls_config

<a id="section"></a>

Type: `"single"`. Computed.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"tls_11_plus\",\"tls_12_plus\"]"
}
```

## Direct properties

- [tls_11_plus](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/https_auto_cert/tls_config/tls_11_plus/): complete subsection reference.

- [tls_12_plus](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/https_auto_cert/tls_config/tls_12_plus/): complete subsection reference.

## Next pages

- [https_auto_cert.tls_config.tls_11_plus](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/https_auto_cert/tls_config/tls_11_plus/)
- [https_auto_cert.tls_config.tls_12_plus](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/https_auto_cert/tls_config/tls_12_plus/)
- [https_auto_cert](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/https_auto_cert/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
