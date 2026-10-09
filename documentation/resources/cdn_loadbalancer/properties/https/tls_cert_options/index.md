---
page_title: "https.tls_cert_options"
subcategory: "Load Balancing"
description: "TLS Certificate OPTIONS."
xcsh_docs: {"aliases": ["https tls cert options"], "body_bytes": 1456, "body_sha256": "sha256:7208d6d069dbb8573a496e4b1e1fc2d59df6f22da003cd511056f0870eddb350", "capabilities": ["cdn", "load-balancing.tls"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:https:tls_cert_options:tls_cert_params", "xcsh-docs:resources:cdn_loadbalancer:properties:https:tls_cert_options:tls_inline_params"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:https:tls_cert_options", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:https", "path": "documentation/resources/cdn_loadbalancer/properties/https/tls_cert_options/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-0333210302232102-3132301301101323-1211333333001031-0100031110002301-1231020320311323-2013321101333113-2121323320223113-3103033031222132", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-011.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["https", "tls_cert_options"], "schema_version": 1, "sections": [{"aliases": ["https tls cert options tls cert params"], "anchor": "section", "description": "Select TLS Parameters and Certificates.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:https:tls_cert_options:tls_cert_params", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["https", "tls_cert_options", "tls_cert_params"], "syntax": "block", "type": "object"}, {"aliases": ["https tls cert options tls inline params"], "anchor": "section", "description": "Inline TLS parameters.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:https:tls_cert_options:tls_inline_params", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["https", "tls_cert_options", "tls_inline_params"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/https/tls_cert_options/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "TLS Certificate OPTIONS.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https.tls_cert_options

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [https](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/https/)
- https.tls_cert_options

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for tls cert options.

Additional upstream details:

TLS Certificate OPTIONS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-tls_certificates_choice": "[\"tls_cert_params\",\"tls_inline_params\"]"
}
```

Terraform syntax:

```terraform
tls_cert_options {
  # Configure direct properties listed below.
}
```

## Direct properties

- [tls_cert_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/https/tls_cert_options/tls_cert_params/): complete subsection reference.

- [tls_inline_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/https/tls_cert_options/tls_inline_params/): complete subsection reference.
