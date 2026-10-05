---
page_title: "discovery_target"
subcategory: ""
description: "Configuration parameter for discovery target."
xcsh_docs: {"aliases": ["discovery target"], "body_bytes": 1793, "body_sha256": "sha256:a5b9c0e15561ceac755f77ba28757300eda80bf5a38d22e5b5b8ac92b1af644d", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:nginx_service_discovery:properties:discovery_target:config_sync_group", "xcsh-docs:data-sources:nginx_service_discovery:properties:discovery_target:nginx_instance"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:nginx_service_discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nginx_service_discovery:properties:discovery_target", "parent_id": "xcsh-docs:data-sources:nginx_service_discovery:reference", "path": "documentation/data-sources/nginx_service_discovery/properties/discovery_target/index.md", "product": "distributed-cloud", "provider_name": "nginx_service_discovery", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-0220312301013130-1323112333310301-1023311301010200-1111320222132102-2112102102200100-3200110003311031-3031212331233013-0020102033031211", "registry_path": "docs/guides/data-sources--nginx_service_discovery--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["discovery_target"], "schema_version": 1, "sections": [{"aliases": ["discovery target config sync group"], "anchor": "section", "description": "Select new ConfigSyncGroup.", "document_id": "xcsh-docs:data-sources:nginx_service_discovery:properties:discovery_target:config_sync_group", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["discovery_target", "config_sync_group"], "syntax": "attribute", "type": "object"}, {"aliases": ["discovery target nginx instance"], "anchor": "section", "description": "Select new NGINX Instance.", "document_id": "xcsh-docs:data-sources:nginx_service_discovery:properties:discovery_target:nginx_instance", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["discovery_target", "nginx_instance"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nginx_service_discovery/properties/discovery_target/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Configuration parameter for discovery target.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["nginx_service_discoveryCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
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

## Next pages

- [discovery_target.config_sync_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_service_discovery/properties/discovery_target/config_sync_group/)
- [discovery_target.nginx_instance](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_service_discovery/properties/discovery_target/nginx_instance/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_service_discovery/properties/)
- [xcsh_nginx_service_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_service_discovery/)
