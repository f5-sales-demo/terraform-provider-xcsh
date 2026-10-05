---
page_title: "discovery_target.nginx_instance"
subcategory: ""
description: "Select new NGINX Instance."
xcsh_docs: {"aliases": ["discovery target nginx instance"], "body_bytes": 1887, "body_sha256": "sha256:7a48044d5639166bd653b408380a23e7c08f37004c737c387fc6ea6a4953be27", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:nginx_service_discovery:properties:discovery_target:nginx_instance:nginx_instance"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:nginx_service_discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:nginx_service_discovery:properties:discovery_target:nginx_instance", "parent_id": "xcsh-docs:resources:nginx_service_discovery:properties:discovery_target", "path": "documentation/resources/nginx_service_discovery/properties/discovery_target/nginx_instance/index.md", "product": "distributed-cloud", "provider_name": "nginx_service_discovery", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-3002103300302203-3233002012210030-2131002122302221-0203011202221231-3100111103220331-1220330203313221-0322022212312332-0230302020002133", "registry_path": "docs/guides/resources--nginx_service_discovery--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "discovery_target.nginx_instance:RequiredObjectAttributes:nginx_instance", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:nginx_service_discovery:properties:discovery_target:nginx_instance:nginx_instance", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["discovery_target", "nginx_instance"], "schema_version": 1, "sections": [{"aliases": ["discovery target nginx instance nginx instance"], "anchor": "section", "description": "Select new NGINX Instance.", "document_id": "xcsh-docs:resources:nginx_service_discovery:properties:discovery_target:nginx_instance:nginx_instance", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["discovery_target", "nginx_instance", "nginx_instance"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nginx_service_discovery/properties/discovery_target/nginx_instance/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Select new NGINX Instance.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["nginx_service_discoveryCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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

Upstream description:

Select new NGINX Instance.

Provider validators and defaults (from schema source):

```go
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

## Next pages

- [discovery_target.nginx_instance.nginx_instance](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nginx_service_discovery/properties/discovery_target/nginx_instance/nginx_instance/)
- [discovery_target](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nginx_service_discovery/properties/discovery_target/)
- [xcsh_nginx_service_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nginx_service_discovery/)
