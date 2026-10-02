---
page_title: "https.tls_cert_options.tls_cert_params"
subcategory: "Load Balancing"
description: "Select TLS Parameters and Certificates."
xcsh_docs: {"aliases": ["cert", "certificate", "existing certificates", "https tls cert options tls cert params", "tls certificates"], "body_bytes": 2986, "body_sha256": "sha256:c605231677afb1c9800c57efe9540b031f7f6121823fcfc6eb38cfec89d6ce13", "capabilities": ["cdn", "load-balancing.tls"], "category": "cdn", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:https:tls_cert_options:tls_cert_params:certificates", "xcsh-docs:data-sources:cdn_loadbalancer:properties:https:tls_cert_options:tls_cert_params:no_mtls", "xcsh-docs:data-sources:cdn_loadbalancer:properties:https:tls_cert_options:tls_cert_params:tls_config", "xcsh-docs:data-sources:cdn_loadbalancer:properties:https:tls_cert_options:tls_cert_params:use_mtls"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:https:tls_cert_options:tls_cert_params", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:https:tls_cert_options", "path": "documentation/data-sources/cdn_loadbalancer/properties/https/tls_cert_options/tls_cert_params/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2102213030323200-0011230110200231-0212313130013000-0223212220103300-2133100002311113-2123312123122300-2130112200001323-1022310011132032", "registry_path": "docs/guides/data-sources--cdn_loadbalancer--reference--group-010.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["https", "tls_cert_options", "tls_cert_params"], "schema_version": 1, "sections": [{"aliases": ["cert", "certificate", "certificates", "existing certificates", "tls certificates"], "anchor": "section", "description": "Select one or more certificates with any domain names.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:https:tls_cert_options:tls_cert_params:certificates", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["https", "tls_cert_options", "tls_cert_params", "certificates"], "syntax": "attribute", "type": "object"}, {"aliases": ["no mtls"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:https:tls_cert_options:tls_cert_params:no_mtls", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["https", "tls_cert_options", "tls_cert_params", "no_mtls"], "syntax": "attribute", "type": "object"}, {"aliases": ["tls config"], "anchor": "section", "description": "This defines various OPTIONS to configure TLS configuration parameters.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:https:tls_cert_options:tls_cert_params:tls_config", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["https", "tls_cert_options", "tls_cert_params", "tls_config"], "syntax": "attribute", "type": "object"}, {"aliases": ["use mtls"], "anchor": "section", "description": "Validation context for downstream client TLS connections.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:https:tls_cert_options:tls_cert_params:use_mtls", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["https", "tls_cert_options", "tls_cert_params", "use_mtls"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/https/tls_cert_options/tls_cert_params/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Select TLS Parameters and Certificates.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https.tls_cert_options.tls_cert_params

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- [https](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/https/)
- [https.tls_cert_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/https/tls_cert_options/)
- https.tls_cert_options.tls_cert_params

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for tls cert params.

Upstream description:

Select TLS Parameters and Certificates.

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

- [certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/https/tls_cert_options/tls_cert_params/certificates/): complete subsection reference.

- [no_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/https/tls_cert_options/tls_cert_params/no_mtls/): complete subsection reference.

- [tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/https/tls_cert_options/tls_cert_params/tls_config/): complete subsection reference.

- [use_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/https/tls_cert_options/tls_cert_params/use_mtls/): complete subsection reference.

## Next pages

- [https.tls_cert_options.tls_cert_params.certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/https/tls_cert_options/tls_cert_params/certificates/)
- [https.tls_cert_options.tls_cert_params.no_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/https/tls_cert_options/tls_cert_params/no_mtls/)
- [https.tls_cert_options.tls_cert_params.tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/https/tls_cert_options/tls_cert_params/tls_config/)
- [https.tls_cert_options.tls_cert_params.use_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/https/tls_cert_options/tls_cert_params/use_mtls/)
- [https.tls_cert_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/https/tls_cert_options/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
