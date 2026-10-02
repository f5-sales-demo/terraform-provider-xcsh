---
page_title: "proxy_config"
subcategory: ""
description: "HTTP/HTTPS Load balancer."
xcsh_docs: {"aliases": ["proxy config"], "body_bytes": 4411, "body_sha256": "sha256:1d047e4c4eecc0c30b7e85cd6ddd121675551453be3719684973a548c25f979f", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:http", "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https", "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https_auto_cert"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config", "parent_id": "xcsh-docs:data-sources:bigip_http_proxy:reference", "path": "documentation/data-sources/bigip_http_proxy/properties/proxy_config/index.md", "product": "distributed-cloud", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322", "registry_path": "docs/guides/data-sources--bigip_http_proxy--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["proxy_config"], "schema_version": 1, "sections": [{"aliases": ["domains"], "anchor": "schema-proxy_config--domains", "description": "A list of domains (host/authority header) that will be matched to loadbalancer. Wildcard hosts are supported in the suffix or prefix form Domain search order: 1. Exact domain names: ``www.example.com``. 2. Prefix domain wildcards: ``*.example.com`` or ``*-bar.example.com``. 3. Special wildcard ``*`` matching any", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_config", "domains"], "syntax": "attribute", "type": "list"}, {"aliases": ["http"], "anchor": "section", "description": "Choice for selecting HTTP proxy.", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:http", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["proxy_config", "http"], "syntax": "attribute", "type": "object"}, {"aliases": ["cert", "certificate", "existing certificates", "https", "tls certificates"], "anchor": "section", "description": "Choice for selecting HTTP proxy with bring your own certificates.", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["proxy_config", "https"], "syntax": "attribute", "type": "object"}, {"aliases": ["cert", "certificate", "existing certificates", "https auto cert", "tls certificates"], "anchor": "section", "description": "Choice for selecting HTTP proxy with bring your own certificates.", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https_auto_cert", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["proxy_config", "https_auto_cert"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bigip_http_proxy/properties/proxy_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "HTTP/HTTPS Load balancer.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

Upstream description:

HTTP/HTTPS Load balancer.

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

- [http](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/http/): complete subsection reference.

- [https](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/): complete subsection reference.

- [https_auto_cert](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/): complete subsection reference.

## Next pages

- [proxy_config.http](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/http/)
- [proxy_config.https](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/)
- [proxy_config.https_auto_cert](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/)
- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/)
