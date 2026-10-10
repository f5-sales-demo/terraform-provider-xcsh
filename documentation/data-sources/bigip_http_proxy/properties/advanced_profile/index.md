---
page_title: "advanced_profile"
subcategory: ""
description: "This defines various advanced Profile OPTIONS for a Loadbalancer."
xcsh_docs: {"aliases": ["advanced profile"], "body_bytes": 1155, "body_sha256": "sha256:e5d44b1d27b86757bd9f9f06914b3de04d0fbb643f5a541d573ba1e3629f9aaa", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:bigip_http_proxy:properties:advanced_profile:disable_spec", "xcsh-docs:data-sources:bigip_http_proxy:properties:advanced_profile:enable_default_profile"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bigip_http_proxy:properties:advanced_profile", "parent_id": "xcsh-docs:data-sources:bigip_http_proxy:reference", "path": "documentation/data-sources/bigip_http_proxy/properties/advanced_profile/index.md", "product": "distributed-cloud", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-0211003021332023-0013000312333003-0331301233132013-2000113303313022-0231112311002001-1300321212312100-1303030031223022-3332202130110223", "registry_path": "docs/guides/data-sources--bigip_http_proxy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["advanced_profile"], "schema_version": 1, "sections": [{"aliases": ["advanced profile disable spec"], "anchor": "section", "description": "Enable this option", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:advanced_profile:disable_spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["advanced_profile", "disable_spec"], "syntax": "attribute", "type": "object"}, {"aliases": ["advanced profile enable default profile"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:advanced_profile:enable_default_profile", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["advanced_profile", "enable_default_profile"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bigip_http_proxy/properties/advanced_profile/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "This defines various advanced Profile OPTIONS for a Loadbalancer.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# advanced_profile

Breadcrumbs:

- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/)
- advanced_profile

<a id="section"></a>

Type: `"single"`. Computed.

This defines various advanced Profile OPTIONS for a Loadbalancer.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"disable\",\"enable_default_profile\"]"
}
```

## Direct properties

- [disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/advanced_profile/disable_spec/): complete subsection reference.

- [enable_default_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/advanced_profile/enable_default_profile/): complete subsection reference.
