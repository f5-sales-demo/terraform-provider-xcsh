---
page_title: "origin_pool.origin_servers"
subcategory: "Load Balancing"
description: "List of original servers."
xcsh_docs: {"aliases": ["backend servers", "origin pool origin servers", "origin servers", "upstream servers"], "body_bytes": 4300, "body_sha256": "sha256:00524bde007dcd0c34c7c9fea5e4dac5deba9d99da5fee5765a74dc9d383e4bf", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:origin_servers:public_ip", "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:origin_servers:public_name"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:origin_servers", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool", "path": "documentation/resources/cdn_loadbalancer/properties/origin_pool/origin_servers/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0222220003021200-1032332323111001-2020230223011121-2000031222321133-2000210331301301-1031223323223213-0201322023303000-2102200100010033", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-012.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "origin_pool.origin_servers:ConflictingListObjectAttributes:public_ip,public_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:origin_servers:public_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_pool.origin_servers:ConflictingListObjectAttributes:public_ip,public_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:origin_servers:public_name", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_pool", "origin_servers"], "schema_version": 1, "sections": [{"aliases": ["backend servers", "origin servers", "port", "upstream servers"], "anchor": "schema-origin_pool--origin_servers--port", "description": "Port the workload can be reached on Enter a custom port only if your origin server uses a non-default port. Leave the value as 0 to automatically use 443 (TLS) or 80 (non-TLS).", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:origin_servers", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_pool", "origin_servers", "port"], "syntax": "attribute", "type": "number"}, {"aliases": ["backend servers", "origin servers", "public ip", "upstream servers"], "anchor": "section", "description": "Specify origin server with public IP address.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:origin_servers:public_ip", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_pool", "origin_servers", "public_ip"], "syntax": "block", "type": "object"}, {"aliases": ["backend servers", "origin servers", "public name", "upstream servers"], "anchor": "section", "description": "Specify origin server with public DNS name.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:origin_servers:public_name", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-origin_pool--origin_servers--public_name--dns_name", "enforcement": "provider-schema", "group": "origin_pool.origin_servers.public_name:RequiredObjectAttributes:dns_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:origin_servers:public_name", "type": "requires"}], "schema_path": ["origin_pool", "origin_servers", "public_name"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/origin_pool/origin_servers/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "List of original servers.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_pool.origin_servers

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/origin_pool/)
- origin_pool.origin_servers

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

List Of Origin Servers. List of original servers.

Upstream description:

List of original servers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("public_ip",
    "public_name")}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

Terraform syntax:

```terraform
origin_servers {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-origin_pool--origin_servers--port"></a>

### port property

Type: `"number"`. Optional.

Port the workload can be reached on Enter a custom port only if your origin server uses a
non-default port. Leave the value as 0 to automatically use 443 (TLS) or 80 (non-TLS).

Upstream description:

Port the workload can be reached on Enter a custom port only if your origin server uses a
non-default port. Leave the value as 0 to automatically use 443 (TLS) or 80 (non-TLS).

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/origin_pool/origin_servers/public_ip/): complete subsection reference.

- [public_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/origin_pool/origin_servers/public_name/): complete subsection reference.

## Next pages

- [origin_pool.origin_servers.public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/origin_pool/origin_servers/public_ip/)
- [origin_pool.origin_servers.public_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/origin_pool/origin_servers/public_name/)
- [origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/origin_pool/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
