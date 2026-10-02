---
page_title: "sli_to_global_dr"
subcategory: "Networking"
description: "Global network reference for direct connection."
xcsh_docs: {"aliases": ["sli to global dr"], "body_bytes": 2059, "body_sha256": "sha256:2a406c1e832dc69a2ebde38c0c03198274c1ea88b230717b2b78b3ed000f42c8", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:network_connector:properties:sli_to_global_dr:global_vn"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:network_connector:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_connector:properties:sli_to_global_dr", "parent_id": "xcsh-docs:resources:network_connector:reference", "path": "documentation/resources/network_connector/properties/sli_to_global_dr/index.md", "product": "distributed-cloud", "provider_name": "network_connector", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1323220001110322-1200211010233233-0233223320201012-0231301213210211-1002133133033023-0030103033202111-0202301000230323-0310330231012320", "registry_path": "docs/guides/resources--network_connector--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["sli_to_global_dr"], "schema_version": 1, "sections": [{"aliases": ["global vn"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:network_connector:properties:sli_to_global_dr:global_vn", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-sli_to_global_dr--global_vn--name", "enforcement": "provider-schema", "group": "sli_to_global_dr.global_vn:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:network_connector:properties:sli_to_global_dr:global_vn", "type": "requires"}], "schema_path": ["sli_to_global_dr", "global_vn"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_connector/properties/sli_to_global_dr/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Global network reference for direct connection.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["network_connectorCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# sli_to_global_dr

Breadcrumbs:

- [xcsh_network_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/)
- sli_to_global_dr

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: sli\_to\_global\_dr, sli\_to\_slo\_snat, slo\_to\_global\_dr\] Global network reference for
direct connection.

Upstream description:

Global network reference for direct connection.

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

- [sli_to_global_dr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/sli_to_global_dr/#section)
- [sli_to_slo_snat](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/sli_to_slo_snat/#section)
- [slo_to_global_dr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/slo_to_global_dr/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
sli_to_global_dr {
  # Configure direct properties listed below.
}
```

## Direct properties

- [global_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/sli_to_global_dr/global_vn/): complete subsection reference.

## Next pages

- [sli_to_global_dr.global_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/sli_to_global_dr/global_vn/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/)
- [xcsh_network_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/)
