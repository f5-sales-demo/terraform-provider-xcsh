---
page_title: "origin_servers.cbip_service"
subcategory: "Load Balancing"
description: "Specify origin server with Classic BIG-IP Service (Virtual Server)"
xcsh_docs: {"aliases": ["origin servers cbip service"], "body_bytes": 1910, "body_sha256": "sha256:922eb14d64609e9e9d26ae6e00495a5ca8fa98be4bfb87683400be30f4a080b8", "capabilities": ["load-balancing", "load-balancing.backend-servers"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:origin_pool:properties:origin_servers:cbip_service", "parent_id": "xcsh-docs:resources:origin_pool:properties:origin_servers", "path": "documentation/resources/origin_pool/properties/origin_servers/cbip_service/index.md", "product": "distributed-cloud", "provider_name": "origin_pool", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-1010331102312323-2231313310123113-1310231000100331-3110100312012302-3331111003312110-2102012030220303-1220320003131210-3032300320110001", "registry_path": "docs/guides/resources--origin_pool--reference--group-002.md", "relationships": [{"anchor": "schema-origin_servers--cbip_service--service_name", "enforcement": "provider-schema", "group": "origin_servers.cbip_service:RequiredObjectAttributes:service_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:cbip_service", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_servers", "cbip_service"], "schema_version": 1, "sections": [{"aliases": ["origin servers cbip service service name"], "anchor": "schema-origin_servers--cbip_service--service_name", "description": "Name of the discovered Classic BIG-IP virtual server to be used as origin.", "document_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:cbip_service", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_servers", "cbip_service", "service_name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/properties/origin_servers/cbip_service/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Specify origin server with Classic BIG-IP Service (Virtual Server)", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["origin_poolCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_servers.cbip_service

Breadcrumbs:

- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/)
- [origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/)
- origin_servers.cbip_service

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with Classic BIG-IP Service (Virtual Server).

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("service_name")}
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
cbip_service {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-origin_servers--cbip_service--service_name"></a>

### service_name property

Type: `"string"`. Optional.

Name of the discovered Classic BIG-IP virtual server to be used as origin.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```
