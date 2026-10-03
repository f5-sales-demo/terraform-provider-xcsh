---
page_title: "allow_all_requests"
subcategory: "Security"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["allow all requests"], "body_bytes": 2000, "body_sha256": "sha256:1ebce4e4cec88334858cee11c0c88bb389a3613635dc0579c7809bb091626423", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:service_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy:properties:allow_all_requests", "parent_id": "xcsh-docs:resources:service_policy:reference", "path": "documentation/resources/service_policy/properties/allow_all_requests/index.md", "product": "distributed-cloud", "provider_name": "service_policy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0323122132121121-2032121222012200-2021300003103320-2200300021120320-3213122032133013-1012030312210210-1131212101302130-1131330232103211", "registry_path": "docs/guides/resources--service_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["allow_all_requests"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy/properties/allow_all_requests/index.txt", "spec_pin_digest": "sha256:62f71ec22260bc99f65753ef4581eb9e0dec1b65c506bb0d53099db73a05e19e", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.2", "schema_components": ["service_policyCreateRequest"], "target_commit": "1a0b5141f4589ffaf7bb696a4369a16ee74ae2ff"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# allow_all_requests

Breadcrumbs:

- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/)
- allow_all_requests

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: allow\_all\_requests, allow\_list, deny\_all\_requests, deny\_list, rule\_list\]
Configuration parameter for allow all requests.

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

- [allow_all_requests](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/allow_all_requests/#section)
- [allow_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/allow_list/#section)
- [deny_all_requests](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/deny_all_requests/#section)
- [deny_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/deny_list/#section)
- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
allow_all_requests = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/)
- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/)
