---
page_title: "discovery_target"
subcategory: ""
description: "Configuration parameter for discovery target."
xcsh_docs: {"aliases": ["discovery target"], "body_bytes": 2075, "body_sha256": "sha256:e6e5a92614083e20050f3db1183e30b72f6a3dd6d30ad03fa8115e58b8e80373", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:nginx_service_discovery:properties:discovery_target:config_sync_group", "xcsh-docs:resources:nginx_service_discovery:properties:discovery_target:nginx_instance"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:nginx_service_discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:nginx_service_discovery:properties:discovery_target", "parent_id": "xcsh-docs:resources:nginx_service_discovery:reference", "path": "documentation/resources/nginx_service_discovery/properties/discovery_target/index.md", "product": "distributed-cloud", "provider_name": "nginx_service_discovery", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "resources", "registry_anchor": "canonical-0200111022000313-3021311210312113-2123133211211313-0300212021112210-0021332122331022-2330302200002331-0210301332130032-1132030132233312", "registry_path": "docs/guides/resources--nginx_service_discovery--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "discovery_target:ConflictingObjectAttributes:config_sync_group,nginx_instance", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nginx_service_discovery:properties:discovery_target:config_sync_group", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "discovery_target:ConflictingObjectAttributes:config_sync_group,nginx_instance", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nginx_service_discovery:properties:discovery_target:nginx_instance", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["discovery_target"], "schema_version": 1, "sections": [{"aliases": ["config sync group"], "anchor": "section", "description": "Select new ConfigSyncGroup.", "document_id": "xcsh-docs:resources:nginx_service_discovery:properties:discovery_target:config_sync_group", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "discovery_target.config_sync_group:RequiredObjectAttributes:config_sync_group", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:nginx_service_discovery:properties:discovery_target:config_sync_group:config_sync_group", "type": "requires"}], "schema_path": ["discovery_target", "config_sync_group"], "syntax": "block", "type": "object"}, {"aliases": ["nginx instance"], "anchor": "section", "description": "Select new NGINX Instance.", "document_id": "xcsh-docs:resources:nginx_service_discovery:properties:discovery_target:nginx_instance", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "discovery_target.nginx_instance:RequiredObjectAttributes:nginx_instance", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:nginx_service_discovery:properties:discovery_target:nginx_instance:nginx_instance", "type": "requires"}], "schema_path": ["discovery_target", "nginx_instance"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nginx_service_discovery/properties/discovery_target/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configuration parameter for discovery target.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["nginx_service_discoveryCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# discovery_target

Breadcrumbs:

- [xcsh_nginx_service_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nginx_service_discovery/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nginx_service_discovery/properties/)
- discovery_target

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for discovery target.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("config_sync_group",
    "nginx_instance")}
```

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

Terraform syntax:

```terraform
discovery_target {
  # Configure direct properties listed below.
}
```

## Direct properties

- [config_sync_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nginx_service_discovery/properties/discovery_target/config_sync_group/): complete subsection reference.

- [nginx_instance](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nginx_service_discovery/properties/discovery_target/nginx_instance/): complete subsection reference.

## Next pages

- [discovery_target.config_sync_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nginx_service_discovery/properties/discovery_target/config_sync_group/)
- [discovery_target.nginx_instance](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nginx_service_discovery/properties/discovery_target/nginx_instance/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nginx_service_discovery/properties/)
- [xcsh_nginx_service_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nginx_service_discovery/)
