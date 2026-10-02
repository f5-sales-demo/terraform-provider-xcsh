---
page_title: "https.tls_cert_options.tls_inline_params.tls_certificates.use_system_defaults"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["cert", "certificate", "existing certificates", "https tls cert options tls inline params tls certificates use system defaults", "tls certificates"], "body_bytes": 2023, "body_sha256": "sha256:a111634ccc9dfe98a8c27431e8eba4ae07f197cc51151d9b14cad32c68c10e60", "capabilities": ["cdn", "load-balancing.tls"], "category": "cdn", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:https:tls_cert_options:tls_inline_params:tls_certificates:use_system_defaults", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:https:tls_cert_options:tls_inline_params:tls_certificates", "path": "documentation/resources/cdn_loadbalancer/properties/https/tls_cert_options/tls_inline_params/tls_certificates/use_system_defaults/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0020003122102012-0323213211021111-3103300212000331-2133023113001032-3323011311000002-1300013002231113-3133302101120331-3121200113232223", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-011.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["https", "tls_cert_options", "tls_inline_params", "tls_certificates", "use_system_defaults"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/https/tls_cert_options/tls_inline_params/tls_certificates/use_system_defaults/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https.tls_cert_options.tls_inline_params.tls_certificates.use_system_defaults

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [https](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/https/)
- [https.tls_cert_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/https/tls_cert_options/)
- [https.tls_cert_options.tls_inline_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/https/tls_cert_options/tls_inline_params/)
- [https.tls_cert_options.tls_inline_params.tls_certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/https/tls_cert_options/tls_inline_params/tls_certificates/)
- https.tls_cert_options.tls_inline_params.tls_certificates.use_system_defaults

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for use system defaults.

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
use_system_defaults = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [https.tls_cert_options.tls_inline_params.tls_certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/https/tls_cert_options/tls_inline_params/tls_certificates/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
