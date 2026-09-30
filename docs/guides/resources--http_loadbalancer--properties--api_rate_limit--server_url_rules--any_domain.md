---
page_title: "api_rate_limit.server_url_rules.any_domain"
subcategory: "Load Balancing"
description: "api_rate_limit.server_url_rules.any_domain for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1103, "body_sha256": "sha256:9e85cee5c91de75249a3bafab782e331b84d158f8047999881f44a72fdc904a2", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules:any_domain", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules:any_domain", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules", "path": "docs/guides/resources--http_loadbalancer--properties--api_rate_limit--server_url_rules--any_domain.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_rate_limit", "server_url_rules", "any_domain"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_rate_limit/server_url_rules/any_domain/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_rate_limit.server_url_rules.any_domain for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# api_rate_limit.server_url_rules.any_domain

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [api_rate_limit](resources--http_loadbalancer--properties--api_rate_limit.md)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--properties--api_rate_limit--server_url_rules.md)
- api_rate_limit.server_url_rules.any_domain

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
any_domain = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [api_rate_limit.server_url_rules](resources--http_loadbalancer--properties--api_rate_limit--server_url_rules.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
