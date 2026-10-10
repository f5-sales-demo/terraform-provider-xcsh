---
page_title: "single_lb_app.enable_discovery.custom_api_auth_discovery"
subcategory: "Load Balancing"
description: "API Discovery Advanced settings."
xcsh_docs: {"aliases": ["single lb app enable discovery custom api auth discovery"], "body_bytes": 1433, "body_sha256": "sha256:b142bace77c2d5cf92232b7c240ccaec2bf49ed99ed14baeda605d71fa02c2ca", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:custom_api_auth_discovery:api_discovery_ref"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:custom_api_auth_discovery", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery", "path": "documentation/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/custom_api_auth_discovery/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-0033020300201021-2112213210103002-3302223002101203-1012232012320110-3312220020100332-0310003001132021-2230020302333111-0112132130323101", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-028.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["single_lb_app", "enable_discovery", "custom_api_auth_discovery"], "schema_version": 1, "sections": [{"aliases": ["single lb app enable discovery custom api auth discovery api discovery ref"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:custom_api_auth_discovery:api_discovery_ref", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["single_lb_app", "enable_discovery", "custom_api_auth_discovery", "api_discovery_ref"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/custom_api_auth_discovery/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "API Discovery Advanced settings.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# single_lb_app.enable_discovery.custom_api_auth_discovery

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [single_lb_app](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/single_lb_app/)
- [single_lb_app.enable_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/)
- single_lb_app.enable_discovery.custom_api_auth_discovery

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

API Discovery Advanced Settings. API Discovery Advanced settings.

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
custom_api_auth_discovery {
  # Configure direct properties listed below.
}
```

## Direct properties

- [api_discovery_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/custom_api_auth_discovery/api_discovery_ref/): complete subsection reference.
