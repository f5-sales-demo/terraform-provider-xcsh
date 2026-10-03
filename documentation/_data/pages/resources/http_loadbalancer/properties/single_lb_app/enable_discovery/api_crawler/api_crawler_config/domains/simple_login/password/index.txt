---
page_title: "single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password"
subcategory: "Load Balancing"
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["login", "login result", "sign in", "single lb app enable discovery api crawler api crawler config domains simple login password"], "body_bytes": 4049, "body_sha256": "sha256:7d9bb89131a8e50b5c745ca411860daea03d3c2ff713a8d09ee872acd74447aa", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler:api_crawler_config:domains:simple_login:password:blindfold_secret_info", "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler:api_crawler_config:domains:simple_login:password:clear_secret_info"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler:api_crawler_config:domains:simple_login:password", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler:api_crawler_config:domains:simple_login", "path": "documentation/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_crawler/api_crawler_config/domains/simple_login/password/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3212023010311330-0110001320133130-3222031233120311-1132010133120332-0223233200303331-0131332200020212-0110002020231331-2203200001011110", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-027.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler:api_crawler_config:domains:simple_login:password:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler:api_crawler_config:domains:simple_login:password:clear_secret_info", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["single_lb_app", "enable_discovery", "api_crawler", "api_crawler_config", "domains", "simple_login", "password"], "schema_version": 1, "sections": [{"aliases": ["login", "login result", "sign in", "single lb app enable discovery api crawler api crawler config domains simple login password blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler:api_crawler_config:domains:simple_login:password:blindfold_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-single_lb_app--enable_discovery--api_crawler--api_crawler_config--domains--simple_login--password--blindfold_secret_info--location", "enforcement": "provider-schema", "group": "single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info:RequiredObjectAttributes:location", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler:api_crawler_config:domains:simple_login:password:blindfold_secret_info", "type": "requires"}], "schema_path": ["single_lb_app", "enable_discovery", "api_crawler", "api_crawler_config", "domains", "simple_login", "password", "blindfold_secret_info"], "syntax": "block", "type": "object"}, {"aliases": ["login", "login result", "sign in", "single lb app enable discovery api crawler api crawler config domains simple login password clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler:api_crawler_config:domains:simple_login:password:clear_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-single_lb_app--enable_discovery--api_crawler--api_crawler_config--domains--simple_login--password--clear_secret_info--url", "enforcement": "provider-schema", "group": "single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler:api_crawler_config:domains:simple_login:password:clear_secret_info", "type": "requires"}], "schema_path": ["single_lb_app", "enable_discovery", "api_crawler", "api_crawler_config", "domains", "simple_login", "password", "clear_secret_info"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_crawler/api_crawler_config/domains/simple_login/password/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [single_lb_app](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/single_lb_app/)
- [single_lb_app.enable_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/)
- [single_lb_app.enable_discovery.api_crawler](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_crawler/)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_crawler/api_crawler_config/)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_crawler/api_crawler_config/domains/)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_crawler/api_crawler_config/domains/simple_login/)
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

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_crawler/api_crawler_config/domains/simple_login/password/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_crawler/api_crawler_config/domains/simple_login/password/clear_secret_info/): complete subsection reference.

## Next pages

- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_crawler/api_crawler_config/domains/simple_login/password/blindfold_secret_info/)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_crawler/api_crawler_config/domains/simple_login/password/clear_secret_info/)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_crawler/api_crawler_config/domains/simple_login/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
