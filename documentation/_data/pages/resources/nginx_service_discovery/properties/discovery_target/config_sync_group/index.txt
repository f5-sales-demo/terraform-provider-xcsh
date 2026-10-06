---
page_title: "discovery_target.config_sync_group"
subcategory: ""
description: "Select new ConfigSyncGroup."
xcsh_docs: {"aliases": ["discovery target config sync group"], "body_bytes": 1428, "body_sha256": "sha256:c78f7e49a474ea0267e47507cab318c5cfda7839ea7fe4c6fc10c23b83791912", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:nginx_service_discovery:properties:discovery_target:config_sync_group:config_sync_group"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:nginx_service_discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:nginx_service_discovery:properties:discovery_target:config_sync_group", "parent_id": "xcsh-docs:resources:nginx_service_discovery:properties:discovery_target", "path": "documentation/resources/nginx_service_discovery/properties/discovery_target/config_sync_group/index.md", "product": "distributed-cloud", "provider_name": "nginx_service_discovery", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-2011212030103323-1211101011023030-1210222010122301-0322112032010033-1113232200111120-0030013221311330-3123122103230032-1221201330220122", "registry_path": "docs/guides/resources--nginx_service_discovery--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "discovery_target.config_sync_group:RequiredObjectAttributes:config_sync_group", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:nginx_service_discovery:properties:discovery_target:config_sync_group:config_sync_group", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["discovery_target", "config_sync_group"], "schema_version": 1, "sections": [{"aliases": ["discovery target config sync group config sync group"], "anchor": "section", "description": "Select new ConfigSyncGroup.", "document_id": "xcsh-docs:resources:nginx_service_discovery:properties:discovery_target:config_sync_group:config_sync_group", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["discovery_target", "config_sync_group", "config_sync_group"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nginx_service_discovery/properties/discovery_target/config_sync_group/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Select new ConfigSyncGroup.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["nginx_service_discoveryCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# discovery_target.config_sync_group

Breadcrumbs:

- [xcsh_nginx_service_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nginx_service_discovery/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nginx_service_discovery/properties/)
- [discovery_target](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nginx_service_discovery/properties/discovery_target/)
- discovery_target.config_sync_group

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for config sync group.

Additional upstream details:

Select new ConfigSyncGroup.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("config_sync_group")}
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
config_sync_group {
  # Configure direct properties listed below.
}
```

## Direct properties

- [config_sync_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nginx_service_discovery/properties/discovery_target/config_sync_group/config_sync_group/): complete subsection reference.
