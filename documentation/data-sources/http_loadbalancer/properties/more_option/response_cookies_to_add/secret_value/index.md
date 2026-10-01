---
page_title: "more_option.response_cookies_to_add.secret_value"
subcategory: "Load Balancing"
description: "more_option.response_cookies_to_add.secret_value for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2422, "body_sha256": "sha256:958e9715982fb3172aec266e7845c7061a325279cceb7e387f6df617405dcfe0", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:more_option:response_cookies_to_add:secret_value:blindfold_secret_info", "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:response_cookies_to_add:secret_value:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:response_cookies_to_add:secret_value", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:response_cookies_to_add", "path": "documentation/data-sources/http_loadbalancer/properties/more_option/response_cookies_to_add/secret_value/index.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["more_option", "response_cookies_to_add", "secret_value"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/more_option/response_cookies_to_add/secret_value/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "more_option.response_cookies_to_add.secret_value for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# more_option.response_cookies_to_add.secret_value

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [more_option](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/)
- [more_option.response_cookies_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/response_cookies_to_add/)
- more_option.response_cookies_to_add.secret_value

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

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/response_cookies_to_add/secret_value/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/response_cookies_to_add/secret_value/clear_secret_info/): complete subsection reference.

## Next pages

- [more_option.response_cookies_to_add.secret_value.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/response_cookies_to_add/secret_value/blindfold_secret_info/)
- [more_option.response_cookies_to_add.secret_value.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/response_cookies_to_add/secret_value/clear_secret_info/)
- [more_option.response_cookies_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/response_cookies_to_add/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
