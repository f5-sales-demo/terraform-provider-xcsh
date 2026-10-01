---
page_title: "origin_pool"
subcategory: "Load Balancing"
description: "origin_pool for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 3302, "body_sha256": "sha256:fca48021083a4e574446031a74767356676d769b73c871873312bf73e13eda10", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:more_origin_options", "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:no_tls", "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:origin_servers", "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:public_name", "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls"], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:reference", "path": "docs/guides/resources--cdn_loadbalancer--properties--origin_pool.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["origin_pool"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/origin_pool/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "origin_pool for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_pool

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- origin_pool

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for origin pool.

Upstream description:

Origin Pool for the CDN distribution.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("origin_servers"),
  validators.ConflictingObjectAttributes("no_tls",
    "use_tls")}
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
  "x-ves-oneof-field-tls_choice": "[\"no_tls\",\"use_tls\"]"
}
```

Terraform syntax:

```terraform
origin_pool {
  # Configure direct properties listed below.
}
```

## Direct properties

- [more_origin_options](resources--cdn_loadbalancer--properties--origin_pool--more_origin_options.md): complete subsection reference.

- [no_tls](resources--cdn_loadbalancer--properties--origin_pool--no_tls.md): complete subsection reference.

<a id="schema-origin_pool--origin_request_timeout"></a>

### origin_request_timeout property

Type: `"string"`. Optional.

Configures the time after which a request to the origin will time out waiting for a response.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_time_interval": "10m",
    "ves.io.schema.rules.string.min_time_interval": "10s",
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_time_interval": "10m",
    "ves.io.schema.rules.string.min_time_interval": "10s",
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

- [origin_servers](resources--cdn_loadbalancer--properties--origin_pool--origin_servers.md): complete subsection reference.

- [public_name](resources--cdn_loadbalancer--properties--origin_pool--public_name.md): complete subsection reference.

- [use_tls](resources--cdn_loadbalancer--properties--origin_pool--use_tls.md): complete subsection reference.

## Next pages

- [origin_pool.more_origin_options](resources--cdn_loadbalancer--properties--origin_pool--more_origin_options.md)
- [origin_pool.no_tls](resources--cdn_loadbalancer--properties--origin_pool--no_tls.md)
- [origin_pool.origin_servers](resources--cdn_loadbalancer--properties--origin_pool--origin_servers.md)
- [origin_pool.public_name](resources--cdn_loadbalancer--properties--origin_pool--public_name.md)
- [origin_pool.use_tls](resources--cdn_loadbalancer--properties--origin_pool--use_tls.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
