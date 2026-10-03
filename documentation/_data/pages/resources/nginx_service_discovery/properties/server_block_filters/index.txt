---
page_title: "server_block_filters"
subcategory: ""
description: "Filters discovered server blocks based on server name, domain and ports. Atleast, one field should be populated for each filter. X-textBlockContent: If no filters are specified, all server blocks will be discovered by default."
xcsh_docs: {"aliases": ["server block filters"], "body_bytes": 4700, "body_sha256": "sha256:47c0500aa786112abfe2e0c5f9e412516c2794b9fe6461c6071c45df0a1d8adc", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:nginx_service_discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:nginx_service_discovery:properties:server_block_filters", "parent_id": "xcsh-docs:resources:nginx_service_discovery:reference", "path": "documentation/resources/nginx_service_discovery/properties/server_block_filters/index.md", "product": "distributed-cloud", "provider_name": "nginx_service_discovery", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2222022131131122-3312320112111130-2002010102031033-1121330200131200-3323032021213333-3303212210100132-0101031000301200-2230133310332123", "registry_path": "docs/guides/resources--nginx_service_discovery--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["server_block_filters"], "schema_version": 1, "sections": [{"aliases": ["server block filters name regex"], "anchor": "schema-server_block_filters--name_regex", "description": "Regular expression to match the server name or domain that must be discovered.", "document_id": "xcsh-docs:resources:nginx_service_discovery:properties:server_block_filters", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["server_block_filters", "name_regex"], "syntax": "attribute", "type": "string"}, {"aliases": ["server block filters port ranges"], "anchor": "schema-server_block_filters--port_ranges", "description": "A string containing a comma separated list of individual service ports or port ranges. Each port range consists of a single port or two ports separated by \"-\". For example, 8000-8191. Maximum number of ports allowed is 1024.", "document_id": "xcsh-docs:resources:nginx_service_discovery:properties:server_block_filters", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["server_block_filters", "port_ranges"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nginx_service_discovery/properties/server_block_filters/index.txt", "spec_pin_digest": "sha256:0e278b4afb598ac8628ac74dca6a1b2831cd8607de1d26f56e72e3461a7440c7", "summary": "Filters discovered server blocks based on server name, domain and ports. Atleast, one field should be populated for each filter. X-textBlockContent: If no filters are specified, all server blocks will be discovered by default.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.0", "schema_components": ["nginx_service_discoveryCreateRequest"], "target_commit": "ac024ccbfb8b9f84412813f3e2ab9f5821937ef2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# server_block_filters

Breadcrumbs:

- [xcsh_nginx_service_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nginx_service_discovery/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nginx_service_discovery/properties/)
- server_block_filters

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Filters discovered server blocks based on server name, domain and ports. Atleast, one field should
be populated for each filter. X-textBlockContent: If no filters are specified, all server blocks
will be discovered by default.

Upstream description:

Filters discovered server blocks based on server name, domain and ports. Atleast, one field should
be populated for each filter.

X-textBlockContent: If no filters are specified, all server blocks will be discovered by default.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 10,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 10,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
server_block_filters {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-server_block_filters--name_regex"></a>

### name_regex property

Type: `"string"`. Optional.

Regular expression to match the server name or domain that must be discovered.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="schema-server_block_filters--port_ranges"></a>

### port_ranges property

Type: `"string"`. Optional.

String containing a comma separated list of individual service ports or port ranges. Each port range
consists of a single port or two ports separated by '-'. For example, 8000-8191.

Upstream description:

A string containing a comma separated list of individual service ports or port ranges. Each port
range consists of a single port or two ports separated by "-". For example, 8000-8191. Maximum
number of ports allowed is 1024.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(512),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "1024",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "1024",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  }
}
```

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nginx_service_discovery/properties/)
- [xcsh_nginx_service_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nginx_service_discovery/)
