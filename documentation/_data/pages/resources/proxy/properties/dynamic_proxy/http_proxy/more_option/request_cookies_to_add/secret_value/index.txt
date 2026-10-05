---
page_title: "dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["dynamic proxy http proxy more option request cookies to add secret value"], "body_bytes": 3151, "body_sha256": "sha256:e2655110858beb4357ca1ba07e66775102500b535c117851c615fa6ac0ae1a11", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:request_cookies_to_add:secret_value:blindfold_secret_info", "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:request_cookies_to_add:secret_value:clear_secret_info"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:request_cookies_to_add:secret_value", "parent_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:request_cookies_to_add", "path": "documentation/resources/proxy/properties/dynamic_proxy/http_proxy/more_option/request_cookies_to_add/secret_value/index.md", "product": "distributed-cloud", "provider_name": "proxy", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-2221121133321303-1301031031112023-1111231110021110-2223211301300001-1113002112011222-1211111333212102-1212331102122331-2003210210030301", "registry_path": "docs/guides/resources--proxy--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:request_cookies_to_add:secret_value:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:request_cookies_to_add:secret_value:clear_secret_info", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["dynamic_proxy", "http_proxy", "more_option", "request_cookies_to_add", "secret_value"], "schema_version": 1, "sections": [{"aliases": ["dynamic proxy http proxy more option request cookies to add secret value blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:request_cookies_to_add:secret_value:blindfold_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-dynamic_proxy--http_proxy--more_option--request_cookies_to_add--secret_value--blindfold_secret_info--location", "enforcement": "provider-schema", "group": "dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info:RequiredObjectAttributes:location", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:request_cookies_to_add:secret_value:blindfold_secret_info", "type": "requires"}], "schema_path": ["dynamic_proxy", "http_proxy", "more_option", "request_cookies_to_add", "secret_value", "blindfold_secret_info"], "syntax": "block", "type": "object"}, {"aliases": ["dynamic proxy http proxy more option request cookies to add secret value clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:request_cookies_to_add:secret_value:clear_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-dynamic_proxy--http_proxy--more_option--request_cookies_to_add--secret_value--clear_secret_info--url", "enforcement": "provider-schema", "group": "dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:request_cookies_to_add:secret_value:clear_secret_info", "type": "requires"}], "schema_path": ["dynamic_proxy", "http_proxy", "more_option", "request_cookies_to_add", "secret_value", "clear_secret_info"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/dynamic_proxy/http_proxy/more_option/request_cookies_to_add/secret_value/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value

Breadcrumbs:

- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/)
- [dynamic_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/)
- [dynamic_proxy.http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/http_proxy/)
- [dynamic_proxy.http_proxy.more_option](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/http_proxy/more_option/)
- [dynamic_proxy.http_proxy.more_option.request_cookies_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/http_proxy/more_option/request_cookies_to_add/)
- dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
```

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

Terraform syntax:

```terraform
secret_value {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/http_proxy/more_option/request_cookies_to_add/secret_value/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/http_proxy/more_option/request_cookies_to_add/secret_value/clear_secret_info/): complete subsection reference.

## Next pages

- [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/http_proxy/more_option/request_cookies_to_add/secret_value/blindfold_secret_info/)
- [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/http_proxy/more_option/request_cookies_to_add/secret_value/clear_secret_info/)
- [dynamic_proxy.http_proxy.more_option.request_cookies_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/http_proxy/more_option/request_cookies_to_add/)
- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/)
