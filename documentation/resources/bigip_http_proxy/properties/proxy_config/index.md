---
page_title: "proxy_config"
subcategory: ""
description: "HTTP/HTTPS Load balancer."
xcsh_docs: {"aliases": ["proxy config"], "body_bytes": 4993, "body_sha256": "sha256:b3c5c317ac86195c9857868e4e8f120036899c48ff8029f82642fea87ebbb91a", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:http", "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https", "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https_auto_cert"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config", "parent_id": "xcsh-docs:resources:bigip_http_proxy:reference", "path": "documentation/resources/bigip_http_proxy/properties/proxy_config/index.md", "product": "distributed-cloud", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033", "registry_path": "docs/guides/resources--bigip_http_proxy--reference--group-003.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config:ConflictingObjectAttributes:http,https", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:http", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config:ConflictingObjectAttributes:http,https_auto_cert", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:http", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config:ConflictingObjectAttributes:http,https", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config:ConflictingObjectAttributes:https,https_auto_cert", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config:ConflictingObjectAttributes:http,https_auto_cert", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https_auto_cert", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config:ConflictingObjectAttributes:https,https_auto_cert", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https_auto_cert", "type": "conflicts"}, {"anchor": "schema-proxy_config--domains", "enforcement": "provider-schema", "group": "proxy_config:RequiredObjectAttributes:domains", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["proxy_config"], "schema_version": 1, "sections": [{"aliases": ["domains"], "anchor": "schema-proxy_config--domains", "description": "A list of domains (host/authority header) that will be matched to loadbalancer. Wildcard hosts are supported in the suffix or prefix form Domain search order: 1. Exact domain names: ``www.example.com``. 2. Prefix domain wildcards: ``*.example.com`` or ``*-bar.example.com``. 3. Special wildcard ``*`` matching any domain", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_config", "domains"], "syntax": "attribute", "type": "list"}, {"aliases": ["http"], "anchor": "section", "description": "Choice for selecting HTTP proxy.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:http", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-proxy_config--http--port", "enforcement": "provider-schema", "group": "proxy_config.http:ConflictingObjectAttributes:port,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:http", "type": "conflicts"}, {"anchor": "schema-proxy_config--http--port_ranges", "enforcement": "provider-schema", "group": "proxy_config.http:ConflictingObjectAttributes:port,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:http", "type": "conflicts"}], "schema_path": ["proxy_config", "http"], "syntax": "block", "type": "object"}, {"aliases": ["cert", "certificate", "existing certificates", "https", "tls certificates"], "anchor": "section", "description": "Choice for selecting HTTP proxy with bring your own certificates.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-proxy_config--https--append_server_name", "enforcement": "provider-schema", "group": "proxy_config.https:ConflictingObjectAttributes:append_server_name,default_header", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https", "type": "conflicts"}, {"anchor": "schema-proxy_config--https--append_server_name", "enforcement": "provider-schema", "group": "proxy_config.https:ConflictingObjectAttributes:append_server_name,pass_through", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https", "type": "conflicts"}, {"anchor": "schema-proxy_config--https--append_server_name", "enforcement": "provider-schema", "group": "proxy_config.https:ConflictingObjectAttributes:append_server_name,server_name", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https", "type": "conflicts"}, {"anchor": "schema-proxy_config--https--port", "enforcement": "provider-schema", "group": "proxy_config.https:ConflictingObjectAttributes:port,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https", "type": "conflicts"}, {"anchor": "schema-proxy_config--https--port_ranges", "enforcement": "provider-schema", "group": "proxy_config.https:ConflictingObjectAttributes:port,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https", "type": "conflicts"}, {"anchor": "schema-proxy_config--https--server_name", "enforcement": "provider-schema", "group": "proxy_config.https:ConflictingObjectAttributes:append_server_name,server_name", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https", "type": "conflicts"}, {"anchor": "schema-proxy_config--https--server_name", "enforcement": "provider-schema", "group": "proxy_config.https:ConflictingObjectAttributes:default_header,server_name", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https", "type": "conflicts"}, {"anchor": "schema-proxy_config--https--server_name", "enforcement": "provider-schema", "group": "proxy_config.https:ConflictingObjectAttributes:pass_through,server_name", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https:ConflictingObjectAttributes:append_server_name,default_header", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:default_header", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https:ConflictingObjectAttributes:default_header,pass_through", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:default_header", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https:ConflictingObjectAttributes:default_header,server_name", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:default_header", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https:ConflictingObjectAttributes:default_loadbalancer,non_default_loadbalancer", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:default_loadbalancer", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https:ConflictingObjectAttributes:disable_path_normalize,enable_path_normalize", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:disable_path_normalize", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https:ConflictingObjectAttributes:disable_path_normalize,enable_path_normalize", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:enable_path_normalize", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https:ConflictingObjectAttributes:default_loadbalancer,non_default_loadbalancer", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:non_default_loadbalancer", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https:ConflictingObjectAttributes:append_server_name,pass_through", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:pass_through", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https:ConflictingObjectAttributes:default_header,pass_through", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:pass_through", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https:ConflictingObjectAttributes:pass_through,server_name", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:pass_through", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https:ConflictingObjectAttributes:tls_cert_params,tls_parameters", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https:ConflictingObjectAttributes:tls_cert_params,tls_parameters", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_parameters", "type": "conflicts"}], "schema_path": ["proxy_config", "https"], "syntax": "block", "type": "object"}, {"aliases": ["cert", "certificate", "existing certificates", "https auto cert", "tls certificates"], "anchor": "section", "description": "Choice for selecting HTTP proxy with bring your own certificates.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https_auto_cert", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-proxy_config--https_auto_cert--append_server_name", "enforcement": "provider-schema", "group": "proxy_config.https_auto_cert:ConflictingObjectAttributes:append_server_name,default_header", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https_auto_cert", "type": "conflicts"}, {"anchor": "schema-proxy_config--https_auto_cert--append_server_name", "enforcement": "provider-schema", "group": "proxy_config.https_auto_cert:ConflictingObjectAttributes:append_server_name,pass_through", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https_auto_cert", "type": "conflicts"}, {"anchor": "schema-proxy_config--https_auto_cert--append_server_name", "enforcement": "provider-schema", "group": "proxy_config.https_auto_cert:ConflictingObjectAttributes:append_server_name,server_name", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https_auto_cert", "type": "conflicts"}, {"anchor": "schema-proxy_config--https_auto_cert--port", "enforcement": "provider-schema", "group": "proxy_config.https_auto_cert:ConflictingObjectAttributes:port,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https_auto_cert", "type": "conflicts"}, {"anchor": "schema-proxy_config--https_auto_cert--port_ranges", "enforcement": "provider-schema", "group": "proxy_config.https_auto_cert:ConflictingObjectAttributes:port,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https_auto_cert", "type": "conflicts"}, {"anchor": "schema-proxy_config--https_auto_cert--server_name", "enforcement": "provider-schema", "group": "proxy_config.https_auto_cert:ConflictingObjectAttributes:append_server_name,server_name", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https_auto_cert", "type": "conflicts"}, {"anchor": "schema-proxy_config--https_auto_cert--server_name", "enforcement": "provider-schema", "group": "proxy_config.https_auto_cert:ConflictingObjectAttributes:default_header,server_name", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https_auto_cert", "type": "conflicts"}, {"anchor": "schema-proxy_config--https_auto_cert--server_name", "enforcement": "provider-schema", "group": "proxy_config.https_auto_cert:ConflictingObjectAttributes:pass_through,server_name", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https_auto_cert", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https_auto_cert:ConflictingObjectAttributes:append_server_name,default_header", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https_auto_cert:default_header", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https_auto_cert:ConflictingObjectAttributes:default_header,pass_through", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https_auto_cert:default_header", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https_auto_cert:ConflictingObjectAttributes:default_header,server_name", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https_auto_cert:default_header", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https_auto_cert:ConflictingObjectAttributes:default_loadbalancer,non_default_loadbalancer", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https_auto_cert:default_loadbalancer", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https_auto_cert:ConflictingObjectAttributes:disable_path_normalize,enable_path_normalize", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https_auto_cert:disable_path_normalize", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https_auto_cert:ConflictingObjectAttributes:disable_path_normalize,enable_path_normalize", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https_auto_cert:enable_path_normalize", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https_auto_cert:ConflictingObjectAttributes:no_mtls,use_mtls", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https_auto_cert:no_mtls", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https_auto_cert:ConflictingObjectAttributes:default_loadbalancer,non_default_loadbalancer", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https_auto_cert:non_default_loadbalancer", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https_auto_cert:ConflictingObjectAttributes:append_server_name,pass_through", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https_auto_cert:pass_through", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https_auto_cert:ConflictingObjectAttributes:default_header,pass_through", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https_auto_cert:pass_through", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https_auto_cert:ConflictingObjectAttributes:pass_through,server_name", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https_auto_cert:pass_through", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https_auto_cert:ConflictingObjectAttributes:no_mtls,use_mtls", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https_auto_cert:use_mtls", "type": "conflicts"}], "schema_path": ["proxy_config", "https_auto_cert"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bigip_http_proxy/properties/proxy_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "HTTP/HTTPS Load balancer.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# proxy_config

Breadcrumbs:

- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/)
- proxy_config

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

HTTP/HTTPS Load Balancer. HTTP/HTTPS Load balancer.

Upstream description:

HTTP/HTTPS Load balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("domains"),
  validators.ConflictingObjectAttributes("http",
    "https"),
  validators.ConflictingObjectAttributes("http",
    "https_auto_cert"),
  validators.ConflictingObjectAttributes("https",
    "https_auto_cert")}
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
  "x-ves-oneof-field-loadbalancer_type": "[\"http\",\"https\",\"https_auto_cert\"]"
}
```

Terraform syntax:

```terraform
proxy_config {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-proxy_config--domains"></a>

### domains property

Type: `["list", "string"]`. Optional.

List of domains (host/authority header) that will be matched to loadbalancer. Wildcard hosts are
supported in the suffix or prefix form Domain search order: 1. Exact domain names: \`\`.

Upstream description:

A list of domains (host/authority header) that will be matched to loadbalancer. Wildcard hosts are
supported in the suffix or prefix form

Domain search order: &#8203;1. Exact domain names: \`\`www&#46;example.com\`\`. &#8203;2. Prefix
domain wildcards: \`\`\*.example.com\`\` or \`\`\*-bar.example.com\`\`. &#8203;3. Special wildcard
\`\`\*\`\` matching any domain.

Wildcard will not match empty string. E.g. \`\`\*-bar.example.com\`\` will match
\`\`baz-bar.example.com\`\` but not \`\`-bar.example.com\`\`. The longest wildcards match first.

Domains are also used for SNI matching if the loadbalancer type is HTTPS Domains also indicate the
list of names for which DNS resolution will be done by VER.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [http](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/http/): complete subsection reference.

- [https](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/): complete subsection reference.

- [https_auto_cert](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https_auto_cert/): complete subsection reference.

## Next pages

- [proxy_config.http](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/http/)
- [proxy_config.https](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/)
- [proxy_config.https_auto_cert](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https_auto_cert/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/)
- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/)
