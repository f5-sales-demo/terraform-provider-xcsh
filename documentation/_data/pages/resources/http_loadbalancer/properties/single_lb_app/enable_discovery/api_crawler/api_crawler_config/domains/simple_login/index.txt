---
page_title: "single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login"
subcategory: "Load Balancing"
description: "Configuration parameter for simple login."
xcsh_docs: {"aliases": ["login", "login result", "sign in", "single lb app enable discovery api crawler api crawler config domains simple login"], "body_bytes": 3954, "body_sha256": "sha256:3abf59183f770b80fd19eda0bc60e060099a1371ce22f592427d9e1b497a1705", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler:api_crawler_config:domains:simple_login:password"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler:api_crawler_config:domains:simple_login", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler:api_crawler_config:domains", "path": "documentation/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_crawler/api_crawler_config/domains/simple_login/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2101003003311231-2012222323232001-0131232323112100-0031121001321121-2321110231310231-1102100200223121-0301332121121202-0311301230102310", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-027.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["single_lb_app", "enable_discovery", "api_crawler", "api_crawler_config", "domains", "simple_login"], "schema_version": 1, "sections": [{"aliases": ["password"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler:api_crawler_config:domains:simple_login:password", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler:api_crawler_config:domains:simple_login:password:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler:api_crawler_config:domains:simple_login:password:clear_secret_info", "type": "conflicts"}], "schema_path": ["single_lb_app", "enable_discovery", "api_crawler", "api_crawler_config", "domains", "simple_login", "password"], "syntax": "block", "type": "object"}, {"aliases": ["authentication", "credential setup", "credentials", "user"], "anchor": "schema-single_lb_app--enable_discovery--api_crawler--api_crawler_config--domains--simple_login--user", "description": "Enter the username to assign credentials for the selected domain to crawl.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler:api_crawler_config:domains:simple_login", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["single_lb_app", "enable_discovery", "api_crawler", "api_crawler_config", "domains", "simple_login", "user"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_crawler/api_crawler_config/domains/simple_login/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Configuration parameter for simple login.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [single_lb_app](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/single_lb_app/)
- [single_lb_app.enable_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/)
- [single_lb_app.enable_discovery.api_crawler](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_crawler/)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_crawler/api_crawler_config/)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_crawler/api_crawler_config/domains/)
- single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for simple login.

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
simple_login {
  # Configure direct properties listed below.
}
```

## Direct properties

- [password](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_crawler/api_crawler_config/domains/simple_login/password/): complete subsection reference.

<a id="schema-single_lb_app--enable_discovery--api_crawler--api_crawler_config--domains--simple_login--user"></a>

### user property

Type: `"string"`. Optional.

Enter the username to assign credentials for the selected domain to crawl.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

## Next pages

- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_crawler/api_crawler_config/domains/simple_login/password/)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_crawler/api_crawler_config/domains/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
