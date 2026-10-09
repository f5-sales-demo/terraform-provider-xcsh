---
page_title: "single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains"
subcategory: "Load Balancing"
description: "Enter domains and their credentials to allow authenticated API crawling. You can only include domains you own that are associated with this Load Balancer."
xcsh_docs: {"aliases": ["single lb app enable discovery api crawler api crawler config domains"], "body_bytes": 3959, "body_sha256": "sha256:de2bd75beb1a001b79b21e13f4a1142fbabcd65d156ce55185b475bafe138016", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler:api_crawler_config:domains:simple_login"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler:api_crawler_config:domains", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler:api_crawler_config", "path": "documentation/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_crawler/api_crawler_config/domains/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-2123220030322000-1021313231133233-0031223003213010-0212112312101230-2002311230123033-3012302000131202-2123203120133100-3333100300210003", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-027.md", "relationships": [{"anchor": "schema-single_lb_app--enable_discovery--api_crawler--api_crawler_config--domains--domain", "enforcement": "provider-schema", "group": "single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains:RequiredListObjectAttributes:domain", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler:api_crawler_config:domains", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["single_lb_app", "enable_discovery", "api_crawler", "api_crawler_config", "domains"], "schema_version": 1, "sections": [{"aliases": ["single lb app enable discovery api crawler api crawler config domains domain"], "anchor": "schema-single_lb_app--enable_discovery--api_crawler--api_crawler_config--domains--domain", "description": "Select the domain to execute API Crawling with given credentials.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler:api_crawler_config:domains", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["single_lb_app", "enable_discovery", "api_crawler", "api_crawler_config", "domains", "domain"], "syntax": "attribute", "type": "string"}, {"aliases": ["login", "login result", "sign in", "single lb app enable discovery api crawler api crawler config domains simple login"], "anchor": "section", "description": "Configuration parameter for simple login.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler:api_crawler_config:domains:simple_login", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["single_lb_app", "enable_discovery", "api_crawler", "api_crawler_config", "domains", "simple_login"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_crawler/api_crawler_config/domains/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Enter domains and their credentials to allow authenticated API crawling. You can only include domains you own that are associated with this Load Balancer.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [single_lb_app](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/single_lb_app/)
- [single_lb_app.enable_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/)
- [single_lb_app.enable_discovery.api_crawler](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_crawler/)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_crawler/api_crawler_config/)
- single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Enter domains and their credentials to allow authenticated API crawling. You can only include
domains you own that are associated with this Load Balancer.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("domain")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

Terraform syntax:

```terraform
domains {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-single_lb_app--enable_discovery--api_crawler--api_crawler_config--domains--domain"></a>

### domain property

Type: `"string"`. Optional.

Select the domain to execute API Crawling with given credentials.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.vh_domain": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.vh_domain": "true"
  }
}
```

- [simple_login](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_crawler/api_crawler_config/domains/simple_login/): complete subsection reference.
