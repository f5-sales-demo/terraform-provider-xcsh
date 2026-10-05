---
page_title: "http_proxy.more_option.response_headers_to_add.secret_value"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["http proxy more option response headers to add secret value"], "body_bytes": 2546, "body_sha256": "sha256:557415e59552f5e135cf55e7a777751c8c06a925f7cd5d01aec9390ac000039e", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:proxy:properties:http_proxy:more_option:response_headers_to_add:secret_value:blindfold_secret_info", "xcsh-docs:data-sources:proxy:properties:http_proxy:more_option:response_headers_to_add:secret_value:clear_secret_info"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:proxy:properties:http_proxy:more_option:response_headers_to_add:secret_value", "parent_id": "xcsh-docs:data-sources:proxy:properties:http_proxy:more_option:response_headers_to_add", "path": "documentation/data-sources/proxy/properties/http_proxy/more_option/response_headers_to_add/secret_value/index.md", "product": "distributed-cloud", "provider_name": "proxy", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-0023313330203210-1020302000012220-1300212120112013-3310321023331120-0120032003002310-2233233121032232-1123021133232123-3210303003030011", "registry_path": "docs/guides/data-sources--proxy--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["http_proxy", "more_option", "response_headers_to_add", "secret_value"], "schema_version": 1, "sections": [{"aliases": ["http proxy more option response headers to add secret value blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:data-sources:proxy:properties:http_proxy:more_option:response_headers_to_add:secret_value:blindfold_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["http_proxy", "more_option", "response_headers_to_add", "secret_value", "blindfold_secret_info"], "syntax": "attribute", "type": "object"}, {"aliases": ["http proxy more option response headers to add secret value clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:data-sources:proxy:properties:http_proxy:more_option:response_headers_to_add:secret_value:clear_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["http_proxy", "more_option", "response_headers_to_add", "secret_value", "clear_secret_info"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/proxy/properties/http_proxy/more_option/response_headers_to_add/secret_value/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http_proxy.more_option.response_headers_to_add.secret_value

Breadcrumbs:

- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/)
- [http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/http_proxy/)
- [http_proxy.more_option](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/http_proxy/more_option/)
- [http_proxy.more_option.response_headers_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/http_proxy/more_option/response_headers_to_add/)
- http_proxy.more_option.response_headers_to_add.secret_value

<a id="section"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

## Direct properties

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/http_proxy/more_option/response_headers_to_add/secret_value/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/http_proxy/more_option/response_headers_to_add/secret_value/clear_secret_info/): complete subsection reference.

## Next pages

- [http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/http_proxy/more_option/response_headers_to_add/secret_value/blindfold_secret_info/)
- [http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/http_proxy/more_option/response_headers_to_add/secret_value/clear_secret_info/)
- [http_proxy.more_option.response_headers_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/http_proxy/more_option/response_headers_to_add/)
- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/)
