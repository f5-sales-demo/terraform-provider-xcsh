---
page_title: "single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos"
subcategory: "Load Balancing"
description: "single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2498, "body_sha256": "sha256:83dcca8fb30f026a6d90da0414d0826c079710758d2692e20c48b02ae41b4153", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_discovery_from_code_scan:code_base_integrations:selected_repos", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_discovery_from_code_scan:code_base_integrations:selected_repos", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_discovery_from_code_scan:code_base_integrations", "path": "docs/guides/resources--http_loadbalancer--properties--single_lb_app--enable_discovery--api_discovery_from_code_scan--code_base_integrations--selected_repos.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["single_lb_app", "enable_discovery", "api_discovery_from_code_scan", "code_base_integrations", "selected_repos"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_discovery_from_code_scan/code_base_integrations/selected_repos/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [single_lb_app](resources--http_loadbalancer--properties--single_lb_app.md)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--properties--single_lb_app--enable_discovery.md)
- [single_lb_app.enable_discovery.api_discovery_from_code_scan](resources--http_loadbalancer--properties--single_lb_app--enable_discovery--api_discovery_from_code_scan.md)
- [single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations](resources--http_loadbalancer--properties--single_lb_app--enable_discovery--api_discovery_from_code_scan--code_base_integrations.md)
- single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Select which API repositories represent the LB applications.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("api_code_repo")}
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
selected_repos {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-single_lb_app--enable_discovery--api_discovery_from_code_scan--code_base_integrations--selected_repos--api_code_repo"></a>

### api_code_repo property

Type: `["list", "string"]`. Optional.

Code repository which contain API endpoints.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

## Next pages

- [single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations](resources--http_loadbalancer--properties--single_lb_app--enable_discovery--api_discovery_from_code_scan--code_base_integrations.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
