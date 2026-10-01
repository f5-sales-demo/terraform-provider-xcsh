---
page_title: "default_pool.origin_servers.cbip_service"
subcategory: "Load Balancing"
description: "default_pool.origin_servers.cbip_service for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2168, "body_sha256": "sha256:cc5b6eac396bf86c9e1e89c081057d50777d8e52486ccb9a97f62fbb32ba5f00", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:cbip_service", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:cbip_service", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers", "path": "docs/guides/resources--http_loadbalancer--properties--default_pool--origin_servers--cbip_service.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["default_pool", "origin_servers", "cbip_service"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/default_pool/origin_servers/cbip_service/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "default_pool.origin_servers.cbip_service for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool.origin_servers.cbip_service

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [default_pool](resources--http_loadbalancer--properties--default_pool.md)
- [default_pool.origin_servers](resources--http_loadbalancer--properties--default_pool--origin_servers.md)
- default_pool.origin_servers.cbip_service

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

<a id="schema-default_pool--origin_servers--cbip_service--service_name"></a>

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

- [default_pool.origin_servers](resources--http_loadbalancer--properties--default_pool--origin_servers.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
