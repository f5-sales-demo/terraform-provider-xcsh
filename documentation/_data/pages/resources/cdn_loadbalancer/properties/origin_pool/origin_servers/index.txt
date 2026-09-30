---
page_title: "origin_pool.origin_servers"
subcategory: "Load Balancing"
description: "origin_pool.origin_servers for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 4201, "body_sha256": "sha256:e309a69c3e249f8f7844fb60ddb24f2ef5dd0ac94b82db384352ef05e42a9b8a", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:origin_servers:public_ip", "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:origin_servers:public_name"], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:origin_servers", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool", "path": "documentation/resources/cdn_loadbalancer/properties/origin_pool/origin_servers/index.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["origin_pool", "origin_servers"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/origin_pool/origin_servers/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "origin_pool.origin_servers for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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
