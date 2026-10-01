---
page_title: "api_rate_limit.server_url_rules.client_matcher.ip_matcher"
subcategory: "Load Balancing"
description: "api_rate_limit.server_url_rules.client_matcher.ip_matcher for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2510, "body_sha256": "sha256:c3fb4e84ae499e9f3d887351dfcd7a557dfbb366285c353b510163f82f2a872a", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules:client_matcher:ip_matcher", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules:client_matcher:ip_matcher:prefix_sets"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules:client_matcher:ip_matcher", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules:client_matcher", "path": "docs/guides/resources--http_loadbalancer--properties--api_rate_limit--server_url_rules--client_matcher--ip_matcher.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_rate_limit", "server_url_rules", "client_matcher", "ip_matcher"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_rate_limit/server_url_rules/client_matcher/ip_matcher/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_rate_limit.server_url_rules.client_matcher.ip_matcher for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_rate_limit.server_url_rules.client_matcher.ip_matcher

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [api_rate_limit](resources--http_loadbalancer--properties--api_rate_limit.md)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--properties--api_rate_limit--server_url_rules.md)
- [api_rate_limit.server_url_rules.client_matcher](resources--http_loadbalancer--properties--api_rate_limit--server_url_rules--client_matcher.md)
- api_rate_limit.server_url_rules.client_matcher.ip_matcher

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Upstream description:

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("prefix_sets")}
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
ip_matcher {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-api_rate_limit--server_url_rules--client_matcher--ip_matcher--invert_matcher"></a>

### invert_matcher property

Type: `"bool"`. Optional.

Invert IP Matcher. Invert the match result.

Upstream description:

Invert the match result.

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

- [prefix_sets](resources--http_loadbalancer--properties--api_rate_limit--server_url_rules--client_matcher--ip_matcher--prefix_sets.md): complete subsection reference.

## Next pages

- [api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets](resources--http_loadbalancer--properties--api_rate_limit--server_url_rules--client_matcher--ip_matcher--prefix_sets.md)
- [api_rate_limit.server_url_rules.client_matcher](resources--http_loadbalancer--properties--api_rate_limit--server_url_rules--client_matcher.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
