---
page_title: "server_block_filters"
subcategory: ""
description: "Filters discovered server blocks based on server name, domain and ports. Atleast, one field should be populated for each filter. X-textBlockContent: If no filters are specified, all server blocks will be discovered by default."
xcsh_docs: {"aliases": ["server block filters"], "body_bytes": 4307, "body_sha256": "sha256:292870d40c48409e48f7614c5f1709d6e4f4458937f47442346f094f2547f49e", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:nginx_service_discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nginx_service_discovery:properties:server_block_filters", "parent_id": "xcsh-docs:data-sources:nginx_service_discovery:reference", "path": "documentation/data-sources/nginx_service_discovery/properties/server_block_filters/index.md", "product": "distributed-cloud", "provider_name": "nginx_service_discovery", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3301113020113113-0301200303022012-1223200003223313-2333133121110221-3213132022300322-3323132030003003-2202003121023320-1113212320113302", "registry_path": "docs/guides/data-sources--nginx_service_discovery--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["server_block_filters"], "schema_version": 1, "sections": [{"aliases": ["name regex"], "anchor": "schema-server_block_filters--name_regex", "description": "Regular expression to match the server name or domain that must be discovered.", "document_id": "xcsh-docs:data-sources:nginx_service_discovery:properties:server_block_filters", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["server_block_filters", "name_regex"], "syntax": "attribute", "type": "string"}, {"aliases": ["port ranges"], "anchor": "schema-server_block_filters--port_ranges", "description": "A string containing a comma separated list of individual service ports or port ranges. Each port range consists of a single port or two ports separated by \"-\". For example, 8000-8191. Maximum number of ports allowed is 1024.", "document_id": "xcsh-docs:data-sources:nginx_service_discovery:properties:server_block_filters", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["server_block_filters", "port_ranges"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nginx_service_discovery/properties/server_block_filters/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Filters discovered server blocks based on server name, domain and ports. Atleast, one field should be populated for each filter. X-textBlockContent: If no filters are specified, all server blocks will be discovered by default.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["nginx_service_discoveryCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# server_block_filters

Breadcrumbs:

- [xcsh_nginx_service_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_service_discovery/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_service_discovery/properties/)
- server_block_filters

<a id="section"></a>

Type: `"list"`. Computed.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

## Direct properties

<a id="schema-server_block_filters--name_regex"></a>

### name_regex property

Type: `"string"`. Computed.

Regular expression to match the server name or domain that must be discovered.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

Type: `"string"`. Computed.

String containing a comma separated list of individual service ports or port ranges. Each port range
consists of a single port or two ports separated by '-'. For example, 8000-8191.

Upstream description:

A string containing a comma separated list of individual service ports or port ranges. Each port
range consists of a single port or two ports separated by "-". For example, 8000-8191. Maximum
number of ports allowed is 1024.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_service_discovery/properties/)
- [xcsh_nginx_service_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_service_discovery/)
