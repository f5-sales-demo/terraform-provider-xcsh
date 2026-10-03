---
page_title: "http_proxy.more_option.response_cookies_to_add.secret_value"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["http proxy more option response cookies to add secret value"], "body_bytes": 2822, "body_sha256": "sha256:a51a05b3322c1b04ac96c9302fa0adeb8cc06f102480e662aa36bd6403bd8d77", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:proxy:properties:http_proxy:more_option:response_cookies_to_add:secret_value:blindfold_secret_info", "xcsh-docs:resources:proxy:properties:http_proxy:more_option:response_cookies_to_add:secret_value:clear_secret_info"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:http_proxy:more_option:response_cookies_to_add:secret_value", "parent_id": "xcsh-docs:resources:proxy:properties:http_proxy:more_option:response_cookies_to_add", "path": "documentation/resources/proxy/properties/http_proxy/more_option/response_cookies_to_add/secret_value/index.md", "product": "distributed-cloud", "provider_name": "proxy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0032223220331321-2122103333322322-0112302321320233-1222120311123120-3122320100100323-2000101331022231-0320013223311032-0110033301203222", "registry_path": "docs/guides/resources--proxy--reference--group-004.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "http_proxy.more_option.response_cookies_to_add.secret_value:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:http_proxy:more_option:response_cookies_to_add:secret_value:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http_proxy.more_option.response_cookies_to_add.secret_value:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:http_proxy:more_option:response_cookies_to_add:secret_value:clear_secret_info", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["http_proxy", "more_option", "response_cookies_to_add", "secret_value"], "schema_version": 1, "sections": [{"aliases": ["http proxy more option response cookies to add secret value blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:resources:proxy:properties:http_proxy:more_option:response_cookies_to_add:secret_value:blindfold_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-http_proxy--more_option--response_cookies_to_add--secret_value--blindfold_secret_info--location", "enforcement": "provider-schema", "group": "http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info:RequiredObjectAttributes:location", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:http_proxy:more_option:response_cookies_to_add:secret_value:blindfold_secret_info", "type": "requires"}], "schema_path": ["http_proxy", "more_option", "response_cookies_to_add", "secret_value", "blindfold_secret_info"], "syntax": "block", "type": "object"}, {"aliases": ["http proxy more option response cookies to add secret value clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:resources:proxy:properties:http_proxy:more_option:response_cookies_to_add:secret_value:clear_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-http_proxy--more_option--response_cookies_to_add--secret_value--clear_secret_info--url", "enforcement": "provider-schema", "group": "http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:http_proxy:more_option:response_cookies_to_add:secret_value:clear_secret_info", "type": "requires"}], "schema_path": ["http_proxy", "more_option", "response_cookies_to_add", "secret_value", "clear_secret_info"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/http_proxy/more_option/response_cookies_to_add/secret_value/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["proxyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http_proxy.more_option.response_cookies_to_add.secret_value

Breadcrumbs:

- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/)
- [http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/http_proxy/)
- [http_proxy.more_option](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/http_proxy/more_option/)
- [http_proxy.more_option.response_cookies_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/http_proxy/more_option/response_cookies_to_add/)
- http_proxy.more_option.response_cookies_to_add.secret_value

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

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/http_proxy/more_option/response_cookies_to_add/secret_value/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/http_proxy/more_option/response_cookies_to_add/secret_value/clear_secret_info/): complete subsection reference.

## Next pages

- [http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/http_proxy/more_option/response_cookies_to_add/secret_value/blindfold_secret_info/)
- [http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/http_proxy/more_option/response_cookies_to_add/secret_value/clear_secret_info/)
- [http_proxy.more_option.response_cookies_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/http_proxy/more_option/response_cookies_to_add/)
- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/)
