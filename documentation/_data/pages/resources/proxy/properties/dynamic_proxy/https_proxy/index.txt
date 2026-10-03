---
page_title: "dynamic_proxy.https_proxy"
subcategory: ""
description: "Parameters for dynamic HTTPS proxy."
xcsh_docs: {"aliases": ["dynamic proxy https proxy"], "body_bytes": 1846, "body_sha256": "sha256:9e4f9c6f17d3cb2109f14e558b43c13f2a4307daa892ddaa7c131d224b9eae9a", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:more_option", "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy", "parent_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy", "path": "documentation/resources/proxy/properties/dynamic_proxy/https_proxy/index.md", "product": "distributed-cloud", "provider_name": "proxy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203", "registry_path": "docs/guides/resources--proxy--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["dynamic_proxy", "https_proxy"], "schema_version": 1, "sections": [{"aliases": ["dynamic proxy https proxy more option"], "anchor": "section", "description": "This defines various OPTIONS to define a route.", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:more_option", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-dynamic_proxy--https_proxy--more_option--max_requests_per_connection", "enforcement": "provider-schema", "group": "dynamic_proxy.https_proxy.more_option:ConflictingObjectAttributes:max_requests_per_connection,no_request_limit_per_connection", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:more_option", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dynamic_proxy.https_proxy.more_option:ConflictingObjectAttributes:disable_path_normalize,enable_path_normalize", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:more_option:disable_path_normalize", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dynamic_proxy.https_proxy.more_option:ConflictingObjectAttributes:disable_path_normalize,enable_path_normalize", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:more_option:enable_path_normalize", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dynamic_proxy.https_proxy.more_option:ConflictingObjectAttributes:max_requests_per_connection,no_request_limit_per_connection", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:more_option:no_request_limit_per_connection", "type": "conflicts"}], "schema_path": ["dynamic_proxy", "https_proxy", "more_option"], "syntax": "block", "type": "object"}, {"aliases": ["dynamic proxy https proxy tls params"], "anchor": "section", "description": "Inline TLS parameters.", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "dynamic_proxy.https_proxy.tls_params:ConflictingObjectAttributes:no_mtls,use_mtls", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:no_mtls", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dynamic_proxy.https_proxy.tls_params:ConflictingObjectAttributes:no_mtls,use_mtls", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:use_mtls", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dynamic_proxy.https_proxy.tls_params:RequiredObjectAttributes:tls_certificates", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_certificates", "type": "requires"}], "schema_path": ["dynamic_proxy", "https_proxy", "tls_params"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/dynamic_proxy/https_proxy/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Parameters for dynamic HTTPS proxy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["proxyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dynamic_proxy.https_proxy

Breadcrumbs:

- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/)
- [dynamic_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/)
- dynamic_proxy.https_proxy

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for https proxy.

Upstream description:

Parameters for dynamic HTTPS proxy.

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
https_proxy {
  # Configure direct properties listed below.
}
```

## Direct properties

- [more_option](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/https_proxy/more_option/): complete subsection reference.

- [tls_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/https_proxy/tls_params/): complete subsection reference.

## Next pages

- [dynamic_proxy.https_proxy.more_option](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/https_proxy/more_option/)
- [dynamic_proxy.https_proxy.tls_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/https_proxy/tls_params/)
- [dynamic_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/)
- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/)
