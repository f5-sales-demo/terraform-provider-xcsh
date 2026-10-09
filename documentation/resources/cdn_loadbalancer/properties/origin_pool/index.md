---
page_title: "origin_pool"
subcategory: "Load Balancing"
description: "Origin Pool for the CDN distribution."
xcsh_docs: {"aliases": ["backend servers", "origin pool", "origin servers", "upstream servers"], "body_bytes": 2807, "body_sha256": "sha256:f1b9aeeea9b56db2c66d7cd7db508584aaa1ed319374c476f2b6b2199f4e78d6", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:more_origin_options", "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:no_tls", "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:origin_servers", "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:public_name", "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:reference", "path": "documentation/resources/cdn_loadbalancer/properties/origin_pool/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-012.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_pool"], "schema_version": 1, "sections": [{"aliases": ["backend servers", "origin pool more origin options", "origin servers", "upstream servers"], "anchor": "section", "description": "Configuration parameter for more origin options.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:more_origin_options", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_pool", "more_origin_options"], "syntax": "block", "type": "object"}, {"aliases": ["origin pool no tls"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:no_tls", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_pool", "no_tls"], "syntax": "attribute", "type": "object"}, {"aliases": ["backend servers", "duration", "origin pool origin request timeout", "origin servers", "upstream servers"], "anchor": "schema-origin_pool--origin_request_timeout", "description": "Configures the time after which a request to the origin will time out waiting for a response.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_pool", "origin_request_timeout"], "syntax": "attribute", "type": "string"}, {"aliases": ["backend servers", "origin pool origin servers", "origin servers", "upstream servers"], "anchor": "section", "description": "List of original servers.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:origin_servers", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["origin_pool", "origin_servers"], "syntax": "block", "type": "object"}, {"aliases": ["origin pool public name"], "anchor": "section", "description": "Specify origin server with public DNS name.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:public_name", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_pool", "public_name"], "syntax": "block", "type": "object"}, {"aliases": ["origin pool use tls"], "anchor": "section", "description": "Upstream TLS Parameters.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_pool", "use_tls"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/origin_pool/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Origin Pool for the CDN distribution.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_pool

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- origin_pool

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for origin pool.

Additional upstream details:

Origin Pool for the CDN distribution.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-tls_choice": "[\"no_tls\",\"use_tls\"]"
}
```

Terraform syntax:

```terraform
origin_pool {
  # Configure direct properties listed below.
}
```

## Direct properties

- [more_origin_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/origin_pool/more_origin_options/): complete subsection reference.

- [no_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/origin_pool/no_tls/): complete subsection reference.

<a id="schema-origin_pool--origin_request_timeout"></a>

### origin_request_timeout property

Type: `"string"`. Optional.

Configures the time after which a request to the origin will time out waiting for a response.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_time_interval": "10m",
    "ves.io.schema.rules.string.min_time_interval": "10s",
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_time_interval": "10m",
    "ves.io.schema.rules.string.min_time_interval": "10s",
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

- [origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/origin_pool/origin_servers/): complete subsection reference.

- [public_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/origin_pool/public_name/): complete subsection reference.

- [use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/origin_pool/use_tls/): complete subsection reference.
