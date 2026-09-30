---
page_title: "proxy_config"
subcategory: ""
description: "proxy_config for xcsh_bigip_http_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 4894, "body_sha256": "sha256:947ce38266982e49dc7696a2be3d974180c94b33fc8c4da611247ffafdcef320", "child_ids": ["xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:http", "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https", "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https_auto_cert"], "collection_id": "xcsh-docs:resources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config", "parent_id": "xcsh-docs:resources:bigip_http_proxy:reference", "path": "documentation/resources/bigip_http_proxy/properties/proxy_config/index.md", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["proxy_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bigip_http_proxy/properties/proxy_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "proxy_config for xcsh_bigip_http_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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
