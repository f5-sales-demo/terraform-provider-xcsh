---
page_title: "advanced_profile.enable_default_profile"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["advanced profile enable default profile"], "body_bytes": 1349, "body_sha256": "sha256:9ffcc0020652954e8857f7b013395c6731c18b9f1a405c89da81af0233214496", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bigip_http_proxy:properties:advanced_profile:enable_default_profile", "parent_id": "xcsh-docs:resources:bigip_http_proxy:properties:advanced_profile", "path": "documentation/resources/bigip_http_proxy/properties/advanced_profile/enable_default_profile/index.md", "product": "distributed-cloud", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3133021223323122-3320012200231021-3301031021103333-1013210113222003-1132233333333102-3313303211213212-1000321311333320-1011120110101032", "registry_path": "docs/guides/resources--bigip_http_proxy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["advanced_profile", "enable_default_profile"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bigip_http_proxy/properties/advanced_profile/enable_default_profile/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# advanced_profile.enable_default_profile

Breadcrumbs:

- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/)
- [advanced_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/advanced_profile/)
- advanced_profile.enable_default_profile

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable default profile.

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
enable_default_profile = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [advanced_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/advanced_profile/)
- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/)
