---
page_title: "dynamic_proxy.https_proxy"
subcategory: ""
description: "Parameters for dynamic HTTPS proxy."
xcsh_docs: {"aliases": ["dynamic proxy https proxy"], "body_bytes": 1303, "body_sha256": "sha256:81a1edae8f1dcb3f7d4f3723642c88935a4315d81bd8e38a6700cab54f4927f6", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:more_option", "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy", "parent_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy", "path": "documentation/resources/proxy/properties/dynamic_proxy/https_proxy/index.md", "product": "distributed-cloud", "provider_name": "proxy", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203", "registry_path": "docs/guides/resources--proxy--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["dynamic_proxy", "https_proxy"], "schema_version": 1, "sections": [{"aliases": ["dynamic proxy https proxy more option"], "anchor": "section", "description": "This defines various OPTIONS to define a route.", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:more_option", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-dynamic_proxy--https_proxy--more_option--max_requests_per_connection", "enforcement": "provider-schema", "group": "dynamic_proxy.https_proxy.more_option:ConflictingObjectAttributes:max_requests_per_connection,no_request_limit_per_connection", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:more_option", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dynamic_proxy.https_proxy.more_option:ConflictingObjectAttributes:disable_path_normalize,enable_path_normalize", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:more_option:disable_path_normalize", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dynamic_proxy.https_proxy.more_option:ConflictingObjectAttributes:disable_path_normalize,enable_path_normalize", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:more_option:enable_path_normalize", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dynamic_proxy.https_proxy.more_option:ConflictingObjectAttributes:max_requests_per_connection,no_request_limit_per_connection", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:more_option:no_request_limit_per_connection", "type": "conflicts"}], "schema_path": ["dynamic_proxy", "https_proxy", "more_option"], "syntax": "block", "type": "object"}, {"aliases": ["dynamic proxy https proxy tls params"], "anchor": "section", "description": "Inline TLS parameters.", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "dynamic_proxy.https_proxy.tls_params:ConflictingObjectAttributes:no_mtls,use_mtls", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:no_mtls", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dynamic_proxy.https_proxy.tls_params:ConflictingObjectAttributes:no_mtls,use_mtls", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:use_mtls", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dynamic_proxy.https_proxy.tls_params:RequiredObjectAttributes:tls_certificates", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_certificates", "type": "requires"}], "schema_path": ["dynamic_proxy", "https_proxy", "tls_params"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/dynamic_proxy/https_proxy/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Parameters for dynamic HTTPS proxy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["proxyCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
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

Additional upstream details:

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
