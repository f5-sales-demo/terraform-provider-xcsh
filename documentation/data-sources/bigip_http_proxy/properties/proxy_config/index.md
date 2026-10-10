---
page_title: "proxy_config"
subcategory: ""
description: "HTTP/HTTPS Load balancer."
xcsh_docs: {"aliases": ["proxy config"], "body_bytes": 3458, "body_sha256": "sha256:1ac12b1b67a5c1d0692e5486b7b9952acbbb5e99cb1a655627a4c4dbbcf22933", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:http", "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https", "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https_auto_cert"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config", "parent_id": "xcsh-docs:data-sources:bigip_http_proxy:reference", "path": "documentation/data-sources/bigip_http_proxy/properties/proxy_config/index.md", "product": "distributed-cloud", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322", "registry_path": "docs/guides/data-sources--bigip_http_proxy--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["proxy_config"], "schema_version": 1, "sections": [{"aliases": ["proxy config domains"], "anchor": "schema-proxy_config--domains", "description": "A list of domains (host/authority header) that will be matched to loadbalancer. Wildcard hosts are supported in the suffix or prefix form Domain search order: 1. Exact domain names: ``www.example.com``. 2. Prefix domain wildcards: ``*.example.com`` or ``*-bar.example.com``. 3. Special wildcard ``*`` matching any", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_config", "domains"], "syntax": "attribute", "type": "list"}, {"aliases": ["proxy config http"], "anchor": "section", "description": "Choice for selecting HTTP proxy.", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:http", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["proxy_config", "http"], "syntax": "attribute", "type": "object"}, {"aliases": ["proxy config https"], "anchor": "section", "description": "Choice for selecting HTTP proxy with bring your own certificates.", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["proxy_config", "https"], "syntax": "attribute", "type": "object"}, {"aliases": ["automatic certificate management", "automatic certificates", "managed TLS certificates", "proxy config https auto cert"], "anchor": "section", "description": "Choice for selecting HTTP proxy with bring your own certificates.", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https_auto_cert", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["proxy_config", "https_auto_cert"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bigip_http_proxy/properties/proxy_config/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "HTTP/HTTPS Load balancer.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# proxy_config

Breadcrumbs:

- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/)
- proxy_config

<a id="section"></a>

Type: `"single"`. Computed.

HTTP/HTTPS Load Balancer. HTTP/HTTPS Load balancer.

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

## Direct properties

<a id="schema-proxy_config--domains"></a>

### domains property

Type: `["list", "string"]`. Computed.

A list of domains (host/authority header) that will be matched to loadbalancer. Wildcard hosts are
supported in the suffix or prefix form

Domain search order: &#8203;1. Exact domain names: \`\`www&#46;example.com\`\`. &#8203;2. Prefix
domain wildcards: \`\`\*.example.com\`\` or \`\`\*-bar.example.com\`\`. &#8203;3. Special wildcard
\`\`\*\`\` matching any domain.

Wildcard will not match empty string. E.g. \`\`\*-bar.example.com\`\` will match
\`\`baz-bar.example.com\`\` but not \`\`-bar.example.com\`\`. The longest wildcards match first.

Domains are also used for SNI matching if the loadbalancer type is HTTPS Domains also indicate the
list of names for which DNS resolution will be done by VER.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [http](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/http/): complete subsection reference.

- [https](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/): complete subsection reference.

- [https_auto_cert](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/): complete subsection reference.
