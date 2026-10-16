---
page_title: "enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login"
subcategory: "Load Balancing"
description: "Configuration parameter for simple login."
xcsh_docs: {"aliases": ["enable api discovery api crawler api crawler config domains simple login", "login", "login result", "sign in"], "body_bytes": 2800, "body_sha256": "sha256:b9a9218bd269a8bee05276696b2a6c353e3a9fcd0c66fccb1e401bffae87c0c9", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:enable_api_discovery:api_crawler:api_crawler_config:domains:simple_login:password"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:enable_api_discovery:api_crawler:api_crawler_config:domains:simple_login", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:enable_api_discovery:api_crawler:api_crawler_config:domains", "path": "documentation/data-sources/cdn_loadbalancer/properties/enable_api_discovery/api_crawler/api_crawler_config/domains/simple_login/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "data-sources", "registry_anchor": "canonical-1100331302331013-1303333312221230-3201221100230223-1133113103013222-1032303012011110-1331210020320201-1302222122102332-1312321323212231", "registry_path": "docs/guides/data-sources--cdn_loadbalancer--reference--group-011.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["enable_api_discovery", "api_crawler", "api_crawler_config", "domains", "simple_login"], "schema_version": 1, "sections": [{"aliases": ["enable api discovery api crawler api crawler config domains simple login password", "login", "login result", "sign in"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:enable_api_discovery:api_crawler:api_crawler_config:domains:simple_login:password", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["enable_api_discovery", "api_crawler", "api_crawler_config", "domains", "simple_login", "password"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable api discovery api crawler api crawler config domains simple login user", "login", "login result", "sign in"], "anchor": "schema-enable_api_discovery--api_crawler--api_crawler_config--domains--simple_login--user", "description": "Enter the username to assign credentials for the selected domain to crawl.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:enable_api_discovery:api_crawler:api_crawler_config:domains:simple_login", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_api_discovery", "api_crawler", "api_crawler_config", "domains", "simple_login", "user"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/enable_api_discovery/api_crawler/api_crawler_config/domains/simple_login/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "Configuration parameter for simple login.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- [enable_api_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/enable_api_discovery/)
- [enable_api_discovery.api_crawler](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/enable_api_discovery/api_crawler/)
- [enable_api_discovery.api_crawler.api_crawler_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/enable_api_discovery/api_crawler/api_crawler_config/)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/enable_api_discovery/api_crawler/api_crawler_config/domains/)
- enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login

<a id="section"></a>

Type: `"single"`. Computed.

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

## Direct properties

- [password](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/enable_api_discovery/api_crawler/api_crawler_config/domains/simple_login/password/): complete subsection reference.

<a id="schema-enable_api_discovery--api_crawler--api_crawler_config--domains--simple_login--user"></a>

### user property

Type: `"string"`. Computed.

Enter the username to assign credentials for the selected domain to crawl.

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
      "validatedAt": "2026-10-15T10:49:30+00:00"
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
