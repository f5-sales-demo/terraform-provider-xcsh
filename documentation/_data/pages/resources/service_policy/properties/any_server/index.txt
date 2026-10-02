---
page_title: "any_server"
subcategory: "Security"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["any server"], "body_bytes": 1883, "body_sha256": "sha256:4f167ab99538493ffe2f79e0cad6c92dca9602fdcdc5af6db994948f50b4a8e3", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:service_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy:properties:any_server", "parent_id": "xcsh-docs:resources:service_policy:reference", "path": "documentation/resources/service_policy/properties/any_server/index.md", "product": "distributed-cloud", "provider_name": "service_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2312320122102201-0133030001231211-3021113302202221-0231313332122131-1331023031332202-2100201311220333-3033202032033122-2311001313122213", "registry_path": "docs/guides/resources--service_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["any_server"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy/properties/any_server/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["service_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# any_server

Breadcrumbs:

- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/)
- any_server

<a id="section"></a>

Type: `["object", {}]`. Optional, Computed.

\[OneOf: any\_server, server\_name, server\_name\_matcher, server\_selector\] Enable this option.
Defaults to \`map\[\]\`. Server applies default when omitted.

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

OneOf alternatives in this subsection:

- [any_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/any_server/#section)
- [server_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/#schema-server_name)
- [server_name_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/server_name_matcher/#section)
- [server_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/server_selector/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
any_server = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/)
- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/)
