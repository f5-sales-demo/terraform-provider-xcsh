---
page_title: "discovery_target"
subcategory: ""
description: "Configuration parameter for discovery target."
xcsh_docs: {"aliases": ["discovery target"], "body_bytes": 1166, "body_sha256": "sha256:b1222a5425b241af038bc7df28c61721814fef9a3db0f6cf8d0780bb618825b8", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:nginx_service_discovery:properties:discovery_target:config_sync_group", "xcsh-docs:data-sources:nginx_service_discovery:properties:discovery_target:nginx_instance"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:nginx_service_discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nginx_service_discovery:properties:discovery_target", "parent_id": "xcsh-docs:data-sources:nginx_service_discovery:reference", "path": "documentation/data-sources/nginx_service_discovery/properties/discovery_target/index.md", "product": "distributed-cloud", "provider_name": "nginx_service_discovery", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-0220312301013130-1323112333310301-1023311301010200-1111320222132102-2112102102200100-3200110003311031-3031212331233013-0020102033031211", "registry_path": "docs/guides/data-sources--nginx_service_discovery--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["discovery_target"], "schema_version": 1, "sections": [{"aliases": ["discovery target config sync group"], "anchor": "section", "description": "Select new ConfigSyncGroup.", "document_id": "xcsh-docs:data-sources:nginx_service_discovery:properties:discovery_target:config_sync_group", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["discovery_target", "config_sync_group"], "syntax": "attribute", "type": "object"}, {"aliases": ["discovery target nginx instance"], "anchor": "section", "description": "Select new NGINX Instance.", "document_id": "xcsh-docs:data-sources:nginx_service_discovery:properties:discovery_target:nginx_instance", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["discovery_target", "nginx_instance"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nginx_service_discovery/properties/discovery_target/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "Configuration parameter for discovery target.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["nginx_service_discoveryCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# discovery_target

Breadcrumbs:

- [xcsh_nginx_service_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_service_discovery/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_service_discovery/properties/)
- discovery_target

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for discovery target.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-target": "[\"config_sync_group\",\"nginx_instance\"]"
}
```

## Direct properties

- [config_sync_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_service_discovery/properties/discovery_target/config_sync_group/): complete subsection reference.

- [nginx_instance](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_service_discovery/properties/discovery_target/nginx_instance/): complete subsection reference.
