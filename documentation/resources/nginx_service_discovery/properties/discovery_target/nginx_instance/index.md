---
page_title: "discovery_target.nginx_instance"
subcategory: ""
description: "Select new NGINX Instance."
xcsh_docs: {"aliases": ["discovery target nginx instance"], "body_bytes": 1383, "body_sha256": "sha256:c647c5ee02e831f8d84a82ec29e970760247d65a6640708228d2c43300b8e63b", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:nginx_service_discovery:properties:discovery_target:nginx_instance:nginx_instance"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:nginx_service_discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:nginx_service_discovery:properties:discovery_target:nginx_instance", "parent_id": "xcsh-docs:resources:nginx_service_discovery:properties:discovery_target", "path": "documentation/resources/nginx_service_discovery/properties/discovery_target/nginx_instance/index.md", "product": "distributed-cloud", "provider_name": "nginx_service_discovery", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-3002103300302203-3233002012210030-2131002122302221-0203011202221231-3100111103220331-1220330203313221-0322022212312332-0230302020002133", "registry_path": "docs/guides/resources--nginx_service_discovery--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "discovery_target.nginx_instance:RequiredObjectAttributes:nginx_instance", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:nginx_service_discovery:properties:discovery_target:nginx_instance:nginx_instance", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["discovery_target", "nginx_instance"], "schema_version": 1, "sections": [{"aliases": ["discovery target nginx instance nginx instance"], "anchor": "section", "description": "Select new NGINX Instance.", "document_id": "xcsh-docs:resources:nginx_service_discovery:properties:discovery_target:nginx_instance:nginx_instance", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["discovery_target", "nginx_instance", "nginx_instance"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nginx_service_discovery/properties/discovery_target/nginx_instance/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "Select new NGINX Instance.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["nginx_service_discoveryCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# discovery_target.nginx_instance

Breadcrumbs:

- [xcsh_nginx_service_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nginx_service_discovery/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nginx_service_discovery/properties/)
- [discovery_target](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nginx_service_discovery/properties/discovery_target/)
- discovery_target.nginx_instance

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

NGINXInstance Reference. Select new NGINX Instance.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("nginx_instance")}
```

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
nginx_instance {
  # Configure direct properties listed below.
}
```

## Direct properties

- [nginx_instance](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nginx_service_discovery/properties/discovery_target/nginx_instance/nginx_instance/): complete subsection reference.
