---
page_title: "enable_api_discovery.custom_api_auth_discovery"
subcategory: "Load Balancing"
description: "API Discovery Advanced settings."
xcsh_docs: {"aliases": ["enable api discovery custom api auth discovery"], "body_bytes": 1254, "body_sha256": "sha256:c77571f45f9facdb79404c146ed3c84dbc02f057fd7e3c3dd58b2ee5d8669cb4", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:enable_api_discovery:custom_api_auth_discovery:api_discovery_ref"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:enable_api_discovery:custom_api_auth_discovery", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:enable_api_discovery", "path": "documentation/resources/http_loadbalancer/properties/enable_api_discovery/custom_api_auth_discovery/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-0000012321113332-1333302321200013-0013010310122302-2330031222212000-0110030002020101-3212323100133110-1003113113210123-3333201032123220", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-018.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["enable_api_discovery", "custom_api_auth_discovery"], "schema_version": 1, "sections": [{"aliases": ["enable api discovery custom api auth discovery api discovery ref"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:enable_api_discovery:custom_api_auth_discovery:api_discovery_ref", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-enable_api_discovery--custom_api_auth_discovery--api_discovery_ref--name", "enforcement": "provider-schema", "group": "enable_api_discovery.custom_api_auth_discovery.api_discovery_ref:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:enable_api_discovery:custom_api_auth_discovery:api_discovery_ref", "type": "requires"}], "schema_path": ["enable_api_discovery", "custom_api_auth_discovery", "api_discovery_ref"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/enable_api_discovery/custom_api_auth_discovery/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "API Discovery Advanced settings.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_api_discovery.custom_api_auth_discovery

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [enable_api_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/enable_api_discovery/)
- enable_api_discovery.custom_api_auth_discovery

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

- [api_discovery_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/enable_api_discovery/custom_api_auth_discovery/api_discovery_ref/): complete subsection reference.
