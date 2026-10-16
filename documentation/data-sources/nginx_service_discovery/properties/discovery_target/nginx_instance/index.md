---
page_title: "discovery_target.nginx_instance"
subcategory: ""
description: "Select new NGINX Instance."
xcsh_docs: {"aliases": ["discovery target nginx instance"], "body_bytes": 1089, "body_sha256": "sha256:0fc68bc52ef05bec038c92ce981685058da31c89eab652e19eac1cd0e8e2710e", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:nginx_service_discovery:properties:discovery_target:nginx_instance:nginx_instance"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:nginx_service_discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nginx_service_discovery:properties:discovery_target:nginx_instance", "parent_id": "xcsh-docs:data-sources:nginx_service_discovery:properties:discovery_target", "path": "documentation/data-sources/nginx_service_discovery/properties/discovery_target/nginx_instance/index.md", "product": "distributed-cloud", "provider_name": "nginx_service_discovery", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "data-sources", "registry_anchor": "canonical-0300232101130030-0001232300211210-0303010113013301-3322131213033201-2233001322220221-2202230200021122-1232030300100130-1303110002110131", "registry_path": "docs/guides/data-sources--nginx_service_discovery--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["discovery_target", "nginx_instance"], "schema_version": 1, "sections": [{"aliases": ["discovery target nginx instance nginx instance"], "anchor": "section", "description": "Select new NGINX Instance.", "document_id": "xcsh-docs:data-sources:nginx_service_discovery:properties:discovery_target:nginx_instance:nginx_instance", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["discovery_target", "nginx_instance", "nginx_instance"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nginx_service_discovery/properties/discovery_target/nginx_instance/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "Select new NGINX Instance.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["nginx_service_discoveryCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# discovery_target.nginx_instance

Breadcrumbs:

- [xcsh_nginx_service_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_service_discovery/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_service_discovery/properties/)
- [discovery_target](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_service_discovery/properties/discovery_target/)
- discovery_target.nginx_instance

<a id="section"></a>

Type: `"single"`. Computed.

NGINXInstance Reference. Select new NGINX Instance.

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

## Direct properties

- [nginx_instance](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_service_discovery/properties/discovery_target/nginx_instance/nginx_instance/): complete subsection reference.
