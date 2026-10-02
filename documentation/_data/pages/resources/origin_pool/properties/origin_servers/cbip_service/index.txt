---
page_title: "origin_servers.cbip_service"
subcategory: "Load Balancing"
description: "Specify origin server with Classic BIG-IP Service (Virtual Server)"
xcsh_docs: {"aliases": ["origin servers cbip service"], "body_bytes": 2213, "body_sha256": "sha256:ba1ee89ac4d8bac9052c01e3f4639e052a6539f6eb98c8cc1ac0957b122a47fe", "capabilities": ["load-balancing", "load-balancing.backend-servers"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:origin_pool:properties:origin_servers:cbip_service", "parent_id": "xcsh-docs:resources:origin_pool:properties:origin_servers", "path": "documentation/resources/origin_pool/properties/origin_servers/cbip_service/index.md", "product": "distributed-cloud", "provider_name": "origin_pool", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1010331102312323-2231313310123113-1310231000100331-3110100312012302-3331111003312110-2102012030220303-1220320003131210-3032300320110001", "registry_path": "docs/guides/resources--origin_pool--reference--group-002.md", "relationships": [{"anchor": "schema-origin_servers--cbip_service--service_name", "enforcement": "provider-schema", "group": "origin_servers.cbip_service:RequiredObjectAttributes:service_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:cbip_service", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_servers", "cbip_service"], "schema_version": 1, "sections": [{"aliases": ["service name"], "anchor": "schema-origin_servers--cbip_service--service_name", "description": "Name of the discovered Classic BIG-IP virtual server to be used as origin.", "document_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:cbip_service", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_servers", "cbip_service", "service_name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/properties/origin_servers/cbip_service/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Specify origin server with Classic BIG-IP Service (Virtual Server)", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["origin_poolCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

Upstream description:

Specify origin server with Classic BIG-IP Service (Virtual Server)

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

## Next pages

- [origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/)
- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/)
