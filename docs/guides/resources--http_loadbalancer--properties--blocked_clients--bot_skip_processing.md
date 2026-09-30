---
page_title: "blocked_clients.bot_skip_processing"
subcategory: "Load Balancing"
description: "blocked_clients.bot_skip_processing for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 952, "body_sha256": "sha256:500aebde68c31fd0a8a09c6cb77a5f918be61aa451760ba491a8ee2bfad7c005", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:blocked_clients:bot_skip_processing", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:blocked_clients:bot_skip_processing", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:blocked_clients", "path": "docs/guides/resources--http_loadbalancer--properties--blocked_clients--bot_skip_processing.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["blocked_clients", "bot_skip_processing"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/blocked_clients/bot_skip_processing/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "blocked_clients.bot_skip_processing for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# blocked_clients.bot_skip_processing

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [blocked_clients](resources--http_loadbalancer--properties--blocked_clients.md)
- blocked_clients.bot_skip_processing

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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
bot_skip_processing = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [blocked_clients](resources--http_loadbalancer--properties--blocked_clients.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
