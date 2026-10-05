---
page_title: "origin_pool"
subcategory: "Load Balancing"
description: "Origin Pool for the CDN distribution."
xcsh_docs: {"aliases": ["backend servers", "origin pool", "origin servers", "upstream servers"], "body_bytes": 4010, "body_sha256": "sha256:3efb586a6937b5a973512509399efa2f102498ee0bbbecb0645899c357b5c237", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:more_origin_options", "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:no_tls", "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:origin_servers", "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:public_name", "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:reference", "path": "documentation/resources/cdn_loadbalancer/properties/origin_pool/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-012.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "origin_pool:ConflictingObjectAttributes:no_tls,use_tls", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:no_tls", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_pool:ConflictingObjectAttributes:no_tls,use_tls", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_pool:RequiredObjectAttributes:origin_servers", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:origin_servers", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_pool"], "schema_version": 1, "sections": [{"aliases": ["backend servers", "origin pool more origin options", "origin servers", "upstream servers"], "anchor": "section", "description": "Configuration parameter for more origin options.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:more_origin_options", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_pool", "more_origin_options"], "syntax": "block", "type": "object"}, {"aliases": ["origin pool no tls"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:no_tls", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_pool", "no_tls"], "syntax": "attribute", "type": "object"}, {"aliases": ["backend servers", "duration", "origin pool origin request timeout", "origin servers", "upstream servers"], "anchor": "schema-origin_pool--origin_request_timeout", "description": "Configures the time after which a request to the origin will time out waiting for a response.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_pool", "origin_request_timeout"], "syntax": "attribute", "type": "string"}, {"aliases": ["backend servers", "origin pool origin servers", "origin servers", "upstream servers"], "anchor": "section", "description": "List of original servers.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:origin_servers", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "origin_pool.origin_servers:ConflictingListObjectAttributes:public_ip,public_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:origin_servers:public_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_pool.origin_servers:ConflictingListObjectAttributes:public_ip,public_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:origin_servers:public_name", "type": "conflicts"}], "schema_path": ["origin_pool", "origin_servers"], "syntax": "block", "type": "object"}, {"aliases": ["origin pool public name"], "anchor": "section", "description": "Specify origin server with public DNS name.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:public_name", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-origin_pool--public_name--dns_name", "enforcement": "provider-schema", "group": "origin_pool.public_name:RequiredObjectAttributes:dns_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:public_name", "type": "requires"}], "schema_path": ["origin_pool", "public_name"], "syntax": "block", "type": "object"}, {"aliases": ["origin pool use tls"], "anchor": "section", "description": "Upstream TLS Parameters.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-origin_pool--use_tls--max_session_keys", "enforcement": "provider-schema", "group": "origin_pool.use_tls:ConflictingObjectAttributes:default_session_key_caching,max_session_keys", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls", "type": "conflicts"}, {"anchor": "schema-origin_pool--use_tls--max_session_keys", "enforcement": "provider-schema", "group": "origin_pool.use_tls:ConflictingObjectAttributes:disable_session_key_caching,max_session_keys", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls", "type": "conflicts"}, {"anchor": "schema-origin_pool--use_tls--sni", "enforcement": "provider-schema", "group": "origin_pool.use_tls:ConflictingObjectAttributes:disable_sni,sni", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls", "type": "conflicts"}, {"anchor": "schema-origin_pool--use_tls--sni", "enforcement": "provider-schema", "group": "origin_pool.use_tls:ConflictingObjectAttributes:sni,use_host_header_as_sni", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_pool.use_tls:ConflictingObjectAttributes:default_session_key_caching,disable_session_key_caching", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:default_session_key_caching", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_pool.use_tls:ConflictingObjectAttributes:default_session_key_caching,max_session_keys", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:default_session_key_caching", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_pool.use_tls:ConflictingObjectAttributes:default_session_key_caching,disable_session_key_caching", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:disable_session_key_caching", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_pool.use_tls:ConflictingObjectAttributes:disable_session_key_caching,max_session_keys", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:disable_session_key_caching", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_pool.use_tls:ConflictingObjectAttributes:disable_sni,sni", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:disable_sni", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_pool.use_tls:ConflictingObjectAttributes:disable_sni,use_host_header_as_sni", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:disable_sni", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_pool.use_tls:ConflictingObjectAttributes:no_mtls,use_mtls", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:no_mtls", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_pool.use_tls:ConflictingObjectAttributes:no_mtls,use_mtls_obj", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:no_mtls", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_pool.use_tls:ConflictingObjectAttributes:skip_server_verification,use_server_verification", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:skip_server_verification", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_pool.use_tls:ConflictingObjectAttributes:skip_server_verification,volterra_trusted_ca", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:skip_server_verification", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_pool.use_tls:ConflictingObjectAttributes:disable_sni,use_host_header_as_sni", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_host_header_as_sni", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_pool.use_tls:ConflictingObjectAttributes:sni,use_host_header_as_sni", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_host_header_as_sni", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_pool.use_tls:ConflictingObjectAttributes:no_mtls,use_mtls", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_pool.use_tls:ConflictingObjectAttributes:use_mtls,use_mtls_obj", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_pool.use_tls:ConflictingObjectAttributes:no_mtls,use_mtls_obj", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls_obj", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_pool.use_tls:ConflictingObjectAttributes:use_mtls,use_mtls_obj", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls_obj", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_pool.use_tls:ConflictingObjectAttributes:skip_server_verification,use_server_verification", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_server_verification", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_pool.use_tls:ConflictingObjectAttributes:use_server_verification,volterra_trusted_ca", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_server_verification", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_pool.use_tls:ConflictingObjectAttributes:skip_server_verification,volterra_trusted_ca", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:volterra_trusted_ca", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_pool.use_tls:ConflictingObjectAttributes:use_server_verification,volterra_trusted_ca", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:volterra_trusted_ca", "type": "conflicts"}], "schema_path": ["origin_pool", "use_tls"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/origin_pool/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Origin Pool for the CDN distribution.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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

Upstream description:

Origin Pool for the CDN distribution.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("origin_servers"),
  validators.ConflictingObjectAttributes("no_tls",
    "use_tls")}
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

## Next pages

- [origin_pool.more_origin_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/origin_pool/more_origin_options/)
- [origin_pool.no_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/origin_pool/no_tls/)
- [origin_pool.origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/origin_pool/origin_servers/)
- [origin_pool.public_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/origin_pool/public_name/)
- [origin_pool.use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/origin_pool/use_tls/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
