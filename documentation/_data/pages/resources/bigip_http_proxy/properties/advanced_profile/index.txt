---
page_title: "advanced_profile"
subcategory: ""
description: "This defines various advanced Profile OPTIONS for a Loadbalancer."
xcsh_docs: {"aliases": ["advanced profile"], "body_bytes": 1268, "body_sha256": "sha256:6c10c922f1c699450437954e62c781cf76ab2053f01439f92eb20d25dd5c9ddd", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:bigip_http_proxy:properties:advanced_profile:disable_spec", "xcsh-docs:resources:bigip_http_proxy:properties:advanced_profile:enable_default_profile"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bigip_http_proxy:properties:advanced_profile", "parent_id": "xcsh-docs:resources:bigip_http_proxy:reference", "path": "documentation/resources/bigip_http_proxy/properties/advanced_profile/index.md", "product": "distributed-cloud", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "resources", "registry_anchor": "canonical-2112133030313203-1202233221221323-0231033212122201-1200202122332313-2022030302010022-1221300001120300-3120211033311213-2323103320212230", "registry_path": "docs/guides/resources--bigip_http_proxy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["advanced_profile"], "schema_version": 1, "sections": [{"aliases": ["advanced profile disable spec"], "anchor": "section", "description": "Enable this option", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:advanced_profile:disable_spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["advanced_profile", "disable_spec"], "syntax": "attribute", "type": "object"}, {"aliases": ["advanced profile enable default profile"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:advanced_profile:enable_default_profile", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["advanced_profile", "enable_default_profile"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bigip_http_proxy/properties/advanced_profile/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "This defines various advanced Profile OPTIONS for a Loadbalancer.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# advanced_profile

Breadcrumbs:

- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/)
- advanced_profile

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
advanced_profile {
  # Configure direct properties listed below.
}
```

## Direct properties

- [disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/advanced_profile/disable_spec/): complete subsection reference.

- [enable_default_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/advanced_profile/enable_default_profile/): complete subsection reference.
