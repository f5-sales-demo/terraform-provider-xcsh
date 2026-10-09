---
page_title: "proxy_config.https.http_protocol_options"
subcategory: ""
description: "HTTP protocol configuration OPTIONS for downstream connections."
xcsh_docs: {"aliases": ["proxy config https http protocol options"], "body_bytes": 1975, "body_sha256": "sha256:247555f814f60ded1e2f66aa3a114d0bbc79d2ded4bc11fe52a17b58e2b00320", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:http_protocol_options:http_protocol_enable_v1_only", "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:http_protocol_options:http_protocol_enable_v1_v2", "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:http_protocol_options:http_protocol_enable_v2_only"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:http_protocol_options", "parent_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https", "path": "documentation/resources/bigip_http_proxy/properties/proxy_config/https/http_protocol_options/index.md", "product": "distributed-cloud", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-1200103230101333-2113021032322120-3323321233020213-3302012133221032-2222302001101103-1013231202302013-3031233132323311-3300212010030113", "registry_path": "docs/guides/resources--bigip_http_proxy--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["proxy_config", "https", "http_protocol_options"], "schema_version": 1, "sections": [{"aliases": ["proxy config https http protocol options http protocol enable v1 only"], "anchor": "section", "description": "HTTP/1.1 Protocol OPTIONS for downstream connections.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:http_protocol_options:http_protocol_enable_v1_only", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["proxy_config", "https", "http_protocol_options", "http_protocol_enable_v1_only"], "syntax": "block", "type": "object"}, {"aliases": ["proxy config https http protocol options http protocol enable v1 v2"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:http_protocol_options:http_protocol_enable_v1_v2", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_config", "https", "http_protocol_options", "http_protocol_enable_v1_v2"], "syntax": "attribute", "type": "object"}, {"aliases": ["proxy config https http protocol options http protocol enable v2 only"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:http_protocol_options:http_protocol_enable_v2_only", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_config", "https", "http_protocol_options", "http_protocol_enable_v2_only"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bigip_http_proxy/properties/proxy_config/https/http_protocol_options/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "HTTP protocol configuration OPTIONS for downstream connections.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# proxy_config.https.http_protocol_options

Breadcrumbs:

- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/)
- [proxy_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/)
- [proxy_config.https](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/)
- proxy_config.https.http_protocol_options

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

HTTP protocol configuration OPTIONS for downstream connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-http_protocol_choice": "[\"http_protocol_enable_v1_only\",\"http_protocol_enable_v1_v2\",\"http_protocol_enable_v2_only\"]"
}
```

Terraform syntax:

```terraform
http_protocol_options {
  # Configure direct properties listed below.
}
```

## Direct properties

- [http_protocol_enable_v1_only](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/http_protocol_options/http_protocol_enable_v1_only/): complete subsection reference.

- [http_protocol_enable_v1_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/http_protocol_options/http_protocol_enable_v1_v2/): complete subsection reference.

- [http_protocol_enable_v2_only](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/http_protocol_options/http_protocol_enable_v2_only/): complete subsection reference.
