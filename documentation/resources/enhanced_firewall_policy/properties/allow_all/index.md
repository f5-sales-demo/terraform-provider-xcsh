---
page_title: "allow_all"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["allow all"], "body_bytes": 2460, "body_sha256": "sha256:9c467c2d0d20af4be1d33c86eb94695e57932c5985a4fd385e944dce244e11f7", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:enhanced_firewall_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:enhanced_firewall_policy:properties:allow_all", "parent_id": "xcsh-docs:resources:enhanced_firewall_policy:reference", "path": "documentation/resources/enhanced_firewall_policy/properties/allow_all/index.md", "product": "distributed-cloud", "provider_name": "enhanced_firewall_policy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1102131223103231-1221212033311132-2212320133110100-2123233220030223-0320001123233230-0220122332332022-1031011132232112-2000130111010012", "registry_path": "docs/guides/resources--enhanced_firewall_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["allow_all"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/enhanced_firewall_policy/properties/allow_all/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["enhanced_firewall_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# allow_all

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/)
- allow_all

<a id="section"></a>

Type: `["object", {}]`. Optional, Computed.

\[OneOf: allow\_all, allowed\_destinations, allowed\_sources, denied\_destinations, denied\_sources,
deny\_all, rule\_list\] Enable this option. Defaults to \`map\[\]\`. Server applies default when
omitted.

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

- [allow_all](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/allow_all/#section)
- [allowed_destinations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/allowed_destinations/#section)
- [allowed_sources](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/allowed_sources/#section)
- [denied_destinations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/denied_destinations/#section)
- [denied_sources](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/denied_sources/#section)
- [deny_all](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/deny_all/#section)
- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
allow_all = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/)
- [xcsh_enhanced_firewall_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/)
