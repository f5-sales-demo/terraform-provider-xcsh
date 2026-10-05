---
page_title: "origin_pool.origin_servers"
subcategory: "Load Balancing"
description: "List of original servers."
xcsh_docs: {"aliases": ["backend servers", "origin pool origin servers", "origin servers", "upstream servers"], "body_bytes": 3896, "body_sha256": "sha256:15fb3730a0900fc43f38feb2c63e07212f18edf7f16619cec3501b284833c754", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool:origin_servers:public_ip", "xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool:origin_servers:public_name"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool:origin_servers", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool", "path": "documentation/data-sources/cdn_loadbalancer/properties/origin_pool/origin_servers/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-1011000010211201-1230313033201303-0203320122033323-0202110232110203-0003303302302131-3023122211323223-2231102320110130-3001132020010021", "registry_path": "docs/guides/data-sources--cdn_loadbalancer--reference--group-012.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_pool", "origin_servers"], "schema_version": 1, "sections": [{"aliases": ["origin pool origin servers port"], "anchor": "schema-origin_pool--origin_servers--port", "description": "Port the workload can be reached on Enter a custom port only if your origin server uses a non-default port. Leave the value as 0 to automatically use 443 (TLS) or 80 (non-TLS).", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool:origin_servers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_pool", "origin_servers", "port"], "syntax": "attribute", "type": "number"}, {"aliases": ["origin pool origin servers public ip"], "anchor": "section", "description": "Specify origin server with public IP address.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool:origin_servers:public_ip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_pool", "origin_servers", "public_ip"], "syntax": "attribute", "type": "object"}, {"aliases": ["origin pool origin servers public name"], "anchor": "section", "description": "Specify origin server with public DNS name.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool:origin_servers:public_name", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_pool", "origin_servers", "public_name"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/origin_pool/origin_servers/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "List of original servers.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_pool.origin_servers

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- [origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/origin_pool/)
- origin_pool.origin_servers

<a id="section"></a>

Type: `"list"`. Computed.

List Of Origin Servers. List of original servers.

Upstream description:

List of original servers.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

<a id="schema-origin_pool--origin_servers--port"></a>

### port property

Type: `"number"`. Computed.

Port the workload can be reached on Enter a custom port only if your origin server uses a
non-default port. Leave the value as 0 to automatically use 443 (TLS) or 80 (non-TLS).

Upstream description:

Port the workload can be reached on Enter a custom port only if your origin server uses a
non-default port. Leave the value as 0 to automatically use 443 (TLS) or 80 (non-TLS).

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

- [public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/origin_pool/origin_servers/public_ip/): complete subsection reference.

- [public_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/origin_pool/origin_servers/public_name/): complete subsection reference.

## Next pages

- [origin_pool.origin_servers.public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/origin_pool/origin_servers/public_ip/)
- [origin_pool.origin_servers.public_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/origin_pool/origin_servers/public_name/)
- [origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/origin_pool/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
