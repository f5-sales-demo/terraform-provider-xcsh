---
page_title: "http_proxy.more_option.no_request_limit_per_connection"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["http proxy more option no request limit per connection"], "body_bytes": 1455, "body_sha256": "sha256:9f1346bed4f5a2df586876fb5caa519ea2762338178fcfb333d7521913630dd0", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:http_proxy:more_option:no_request_limit_per_connection", "parent_id": "xcsh-docs:resources:proxy:properties:http_proxy:more_option", "path": "documentation/resources/proxy/properties/http_proxy/more_option/no_request_limit_per_connection/index.md", "product": "distributed-cloud", "provider_name": "proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1003122322301321-1333110123011112-3221222213203212-1023220213321033-2023002003031031-3010120133222233-1001030333011203-3301303002212231", "registry_path": "docs/guides/resources--proxy--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["http_proxy", "more_option", "no_request_limit_per_connection"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/http_proxy/more_option/no_request_limit_per_connection/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http_proxy.more_option.no_request_limit_per_connection

Breadcrumbs:

- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/)
- [http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/http_proxy/)
- [http_proxy.more_option](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/http_proxy/more_option/)
- http_proxy.more_option.no_request_limit_per_connection

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no request limit per connection.

Upstream description:

This can be used for messages where no values are needed.

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
no_request_limit_per_connection = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [http_proxy.more_option](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/http_proxy/more_option/)
- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/)
