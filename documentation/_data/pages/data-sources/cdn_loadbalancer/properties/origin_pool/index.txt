---
page_title: "origin_pool"
subcategory: "Load Balancing"
description: "Origin Pool for the CDN distribution."
xcsh_docs: {"aliases": ["backend servers", "origin pool", "origin servers", "upstream servers"], "body_bytes": 3712, "body_sha256": "sha256:e08bfc308b4a7cc3ec1d04bd892e5bcedbe877787510066f58c8e562c49f9992", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool:more_origin_options", "xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool:no_tls", "xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool:origin_servers", "xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool:public_name", "xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool:use_tls"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:reference", "path": "documentation/data-sources/cdn_loadbalancer/properties/origin_pool/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0223231232013221-1031133321030110-3120202101012130-1330201133233301-2333202311002001-1210313131130123-2211221120033030-3320230211121110", "registry_path": "docs/guides/data-sources--cdn_loadbalancer--reference--group-012.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_pool"], "schema_version": 1, "sections": [{"aliases": ["backend servers", "more origin options", "origin servers", "upstream servers"], "anchor": "section", "description": "Configuration parameter for more origin options.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool:more_origin_options", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_pool", "more_origin_options"], "syntax": "attribute", "type": "object"}, {"aliases": ["no tls"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool:no_tls", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_pool", "no_tls"], "syntax": "attribute", "type": "object"}, {"aliases": ["backend servers", "duration", "operation timeout", "origin request timeout", "origin servers", "upstream servers"], "anchor": "schema-origin_pool--origin_request_timeout", "description": "Configures the time after which a request to the origin will time out waiting for a response.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_pool", "origin_request_timeout"], "syntax": "attribute", "type": "string"}, {"aliases": ["backend servers", "origin servers", "upstream servers"], "anchor": "section", "description": "List of original servers.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool:origin_servers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["origin_pool", "origin_servers"], "syntax": "attribute", "type": "object"}, {"aliases": ["backend servers", "origin servers", "public name", "upstream servers"], "anchor": "section", "description": "Specify origin server with public DNS name.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool:public_name", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_pool", "public_name"], "syntax": "attribute", "type": "object"}, {"aliases": ["use tls"], "anchor": "section", "description": "Upstream TLS Parameters.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool:use_tls", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_pool", "use_tls"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/origin_pool/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Origin Pool for the CDN distribution.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_pool

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- origin_pool

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for origin pool.

Upstream description:

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

## Direct properties

- [more_origin_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/origin_pool/more_origin_options/): complete subsection reference.

- [no_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/origin_pool/no_tls/): complete subsection reference.

<a id="schema-origin_pool--origin_request_timeout"></a>

### origin_request_timeout property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/origin_pool/origin_servers/): complete subsection reference.

- [public_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/origin_pool/public_name/): complete subsection reference.

- [use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/origin_pool/use_tls/): complete subsection reference.

## Next pages

- [origin_pool.more_origin_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/origin_pool/more_origin_options/)
- [origin_pool.no_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/origin_pool/no_tls/)
- [origin_pool.origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/origin_pool/origin_servers/)
- [origin_pool.public_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/origin_pool/public_name/)
- [origin_pool.use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/origin_pool/use_tls/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
