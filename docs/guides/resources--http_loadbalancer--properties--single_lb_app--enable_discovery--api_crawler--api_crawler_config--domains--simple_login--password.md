---
page_title: "single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password"
subcategory: "Load Balancing"
description: "single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 3381, "body_sha256": "sha256:f4509a3fb4a6c4d49a26ab38f14bc45a6b97e39293600f85721ae0f66e1d5484", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler:api_crawler_config:domains:simple_login:password", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler:api_crawler_config:domains:simple_login:password:blindfold_secret_info", "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler:api_crawler_config:domains:simple_login:password:clear_secret_info"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler:api_crawler_config:domains:simple_login:password", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler:api_crawler_config:domains:simple_login", "path": "docs/guides/resources--http_loadbalancer--properties--single_lb_app--enable_discovery--api_crawler--api_crawler_config--domains--simple_login--password.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["single_lb_app", "enable_discovery", "api_crawler", "api_crawler_config", "domains", "simple_login", "password"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_crawler/api_crawler_config/domains/simple_login/password/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [single_lb_app](resources--http_loadbalancer--properties--single_lb_app.md)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--properties--single_lb_app--enable_discovery.md)
- [single_lb_app.enable_discovery.api_crawler](resources--http_loadbalancer--properties--single_lb_app--enable_discovery--api_crawler.md)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config](resources--http_loadbalancer--properties--single_lb_app--enable_discovery--api_crawler--api_crawler_config.md)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains](resources--http_loadbalancer--properties--single_lb_app--enable_discovery--api_crawler--api_crawler_config--domains.md)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login](resources--http_loadbalancer--properties--single_lb_app--enable_discovery--api_crawler--api_crawler_config--domains--simple_login.md)
- single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
password {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](resources--http_loadbalancer--properties--single_lb_app--enable_discovery--api_crawler--api_crawler_config--domains--simple_login--password--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](resources--http_loadbalancer--properties--single_lb_app--enable_discovery--api_crawler--api_crawler_config--domains--simple_login--password--clear_secret_info.md): complete subsection reference.

## Next pages

- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info](resources--http_loadbalancer--properties--single_lb_app--enable_discovery--api_crawler--api_crawler_config--domains--simple_login--password--blindfold_secret_info.md)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info](resources--http_loadbalancer--properties--single_lb_app--enable_discovery--api_crawler--api_crawler_config--domains--simple_login--password--clear_secret_info.md)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login](resources--http_loadbalancer--properties--single_lb_app--enable_discovery--api_crawler--api_crawler_config--domains--simple_login.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
