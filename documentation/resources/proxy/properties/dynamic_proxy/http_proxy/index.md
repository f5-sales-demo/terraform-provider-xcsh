---
page_title: "dynamic_proxy.http_proxy"
subcategory: ""
description: "Parameters for dynamic HTTP proxy."
xcsh_docs: {"aliases": ["dynamic proxy http proxy"], "body_bytes": 1081, "body_sha256": "sha256:14a8b7831c3f0a51a6c86b45f1ad4bf081a94139fa07e36614d5a25845ae1162", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy", "parent_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy", "path": "documentation/resources/proxy/properties/dynamic_proxy/http_proxy/index.md", "product": "distributed-cloud", "provider_name": "proxy", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-3010030230300233-0022012030133013-0221201031310302-0303120102300030-2123221020311232-1011333012001321-3003022113003223-0232330232302103", "registry_path": "docs/guides/resources--proxy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["dynamic_proxy", "http_proxy"], "schema_version": 1, "sections": [{"aliases": ["dynamic proxy http proxy more option"], "anchor": "section", "description": "This defines various OPTIONS to define a route.", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["dynamic_proxy", "http_proxy", "more_option"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/dynamic_proxy/http_proxy/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "Parameters for dynamic HTTP proxy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["proxyCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dynamic_proxy.http_proxy

Breadcrumbs:

- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/)
- [dynamic_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/)
- dynamic_proxy.http_proxy

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Dynamic HTTP Proxy Type. Parameters for dynamic HTTP proxy.

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
http_proxy {
  # Configure direct properties listed below.
}
```

## Direct properties

- [more_option](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/http_proxy/more_option/): complete subsection reference.
