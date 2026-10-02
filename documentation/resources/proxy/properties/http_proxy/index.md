---
page_title: "http_proxy"
subcategory: ""
description: "Parameters for HTTP Connect Proxy."
xcsh_docs: {"aliases": ["http proxy"], "body_bytes": 1678, "body_sha256": "sha256:0ec4b1ea8b5230d0bbeefae698a2d326d3fa0af4bac0bdf727549ef8e00d968e", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:proxy:properties:http_proxy:enable_http", "xcsh-docs:resources:proxy:properties:http_proxy:more_option"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:http_proxy", "parent_id": "xcsh-docs:resources:proxy:reference", "path": "documentation/resources/proxy/properties/http_proxy/index.md", "product": "distributed-cloud", "provider_name": "proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331", "registry_path": "docs/guides/resources--proxy--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["http_proxy"], "schema_version": 1, "sections": [{"aliases": ["enable http"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:proxy:properties:http_proxy:enable_http", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["http_proxy", "enable_http"], "syntax": "block", "type": "object"}, {"aliases": ["more option"], "anchor": "section", "description": "This defines various OPTIONS to define a route.", "document_id": "xcsh-docs:resources:proxy:properties:http_proxy:more_option", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-http_proxy--more_option--max_requests_per_connection", "enforcement": "provider-schema", "group": "http_proxy.more_option:ConflictingObjectAttributes:max_requests_per_connection,no_request_limit_per_connection", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:http_proxy:more_option", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http_proxy.more_option:ConflictingObjectAttributes:disable_path_normalize,enable_path_normalize", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:http_proxy:more_option:disable_path_normalize", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http_proxy.more_option:ConflictingObjectAttributes:disable_path_normalize,enable_path_normalize", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:http_proxy:more_option:enable_path_normalize", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http_proxy.more_option:ConflictingObjectAttributes:max_requests_per_connection,no_request_limit_per_connection", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:http_proxy:more_option:no_request_limit_per_connection", "type": "conflicts"}], "schema_path": ["http_proxy", "more_option"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/http_proxy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Parameters for HTTP Connect Proxy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http_proxy

Breadcrumbs:

- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/)
- http_proxy

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

HTTP Connect Proxy. Parameters for HTTP Connect Proxy.

Upstream description:

Parameters for HTTP Connect Proxy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-http_https_choice": "[\"enable_http\"]"
}
```

Terraform syntax:

```terraform
http_proxy {
  # Configure direct properties listed below.
}
```

## Direct properties

- [enable_http](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/http_proxy/enable_http/): complete subsection reference.

- [more_option](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/http_proxy/more_option/): complete subsection reference.

## Next pages

- [http_proxy.enable_http](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/http_proxy/enable_http/)
- [http_proxy.more_option](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/http_proxy/more_option/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/)
- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/)
