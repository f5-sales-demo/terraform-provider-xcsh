---
page_title: "enable_forward_proxy"
subcategory: "Networking"
description: "Fine tune forward proxy behavior Few configurations allowed are White listed ports and IP prefixes: Forward proxy does application protocol detection and server name(SNI) detection by peeking into the traffic on the incoming downstream connection. Few protocols doesn't have client sending the first data. In such"
xcsh_docs: {"aliases": ["duration", "enable forward proxy", "operation timeout"], "body_bytes": 7035, "body_sha256": "sha256:c5b0f336787ac93f8374715ec16789580ca4d43236f7e7058514eec68ec25071", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:network_connector:properties:enable_forward_proxy:no_interception", "xcsh-docs:data-sources:network_connector:properties:enable_forward_proxy:tls_intercept"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:network_connector:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_connector:properties:enable_forward_proxy", "parent_id": "xcsh-docs:data-sources:network_connector:reference", "path": "documentation/data-sources/network_connector/properties/enable_forward_proxy/index.md", "product": "distributed-cloud", "provider_name": "network_connector", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2213111031121031-3003322303011211-0020021331302310-0213312231333222-0202212313032220-0002012131031003-0322101122200102-1301210120231201", "registry_path": "docs/guides/data-sources--network_connector--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["enable_forward_proxy"], "schema_version": 1, "sections": [{"aliases": ["connection timeout", "duration", "operation timeout"], "anchor": "schema-enable_forward_proxy--connection_timeout", "description": "The timeout for new network connections to upstream server. This is specified in milliseconds. The default value is 2000 (2 seconds)", "document_id": "xcsh-docs:data-sources:network_connector:properties:enable_forward_proxy", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_forward_proxy", "connection_timeout"], "syntax": "attribute", "type": "number"}, {"aliases": ["max connect attempts"], "anchor": "schema-enable_forward_proxy--max_connect_attempts", "description": "Specifies the allowed number of retries on connect failure to upstream server. Defaults to 1.", "document_id": "xcsh-docs:data-sources:network_connector:properties:enable_forward_proxy", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_forward_proxy", "max_connect_attempts"], "syntax": "attribute", "type": "number"}, {"aliases": ["no interception"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:network_connector:properties:enable_forward_proxy:no_interception", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_forward_proxy", "no_interception"], "syntax": "attribute", "type": "object"}, {"aliases": ["tls intercept"], "anchor": "section", "description": "Configuration to enable TLS interception.", "document_id": "xcsh-docs:data-sources:network_connector:properties:enable_forward_proxy:tls_intercept", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["enable_forward_proxy", "tls_intercept"], "syntax": "attribute", "type": "object"}, {"aliases": ["white listed ports"], "anchor": "schema-enable_forward_proxy--white_listed_ports", "description": "Traffic to these destination TCP ports is not subjected to protocol parsing Example \"tmate\" server port.", "document_id": "xcsh-docs:data-sources:network_connector:properties:enable_forward_proxy", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_forward_proxy", "white_listed_ports"], "syntax": "attribute", "type": "list"}, {"aliases": ["white listed prefixes"], "anchor": "schema-enable_forward_proxy--white_listed_prefixes", "description": "Traffic to these destination IP prefixes is not subjected to protocol parsing Example \"tmate\" server IP.", "document_id": "xcsh-docs:data-sources:network_connector:properties:enable_forward_proxy", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_forward_proxy", "white_listed_prefixes"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_connector/properties/enable_forward_proxy/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Fine tune forward proxy behavior Few configurations allowed are White listed ports and IP prefixes: Forward proxy does application protocol detection and server name(SNI) detection by peeking into the traffic on the incoming downstream connection. Few protocols doesn't have client sending the first data. In such", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["network_connectorCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_forward_proxy

Breadcrumbs:

- [xcsh_network_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/)
- enable_forward_proxy

<a id="section"></a>

Type: `"single"`. Computed.

Fine tune forward proxy behavior Few configurations allowed are White listed ports and IP prefixes:
Forward proxy does application protocol detection and server name(SNI) detection by peeking into the
traffic on the incoming downstream connection. Few protocols doesn't have client sending the..

Upstream description:

Fine tune forward proxy behavior

Few configurations allowed are

White listed ports and IP prefixes: Forward proxy does application protocol detection and server
name(SNI) detection by peeking into the traffic on the incoming downstream connection. Few protocols
doesn't have client sending the first data. In such cases, protocol and SNI detection fails. This
configuration allows, skipping protocol and SNI detection for whitelisted IP-prefix-list and ports
connection\_timeout: The timeout for new network connections to upstream server.
Max\_connect\_attempts: Maximum number of attempts made to make new network connection to upstream
server.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-tls_interception_choice": "[\"no_interception\",\"tls_intercept\"]"
}
```

## Direct properties

<a id="schema-enable_forward_proxy--connection_timeout"></a>

### connection_timeout property

Type: `"number"`. Computed.

The timeout for new network connections to upstream server. This is specified in milliseconds. The
(2 seconds). Defaults to \`2000\`.

Upstream description:

The timeout for new network connections to upstream server. This is specified in milliseconds. The
default value is 2000 (2 seconds)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="schema-enable_forward_proxy--max_connect_attempts"></a>

### max_connect_attempts property

Type: `"number"`. Computed.

Specifies the allowed number of retries on connect failure to upstream server. Defaults to \`1\`.

Upstream description:

Specifies the allowed number of retries on connect failure to upstream server. Defaults to 1.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "8"
  }
}
```

- [no_interception](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/enable_forward_proxy/no_interception/): complete subsection reference.

- [tls_intercept](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/enable_forward_proxy/tls_intercept/): complete subsection reference.

<a id="schema-enable_forward_proxy--white_listed_ports"></a>

### white_listed_ports property

Type: `["list", "number"]`. Computed.

Traffic to these destination TCP ports is not subjected to protocol parsing Example 'tmate' server
port.

Upstream description:

Traffic to these destination TCP ports is not subjected to protocol parsing Example "tmate" server
port.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.repeated.items.uint32.lte": "65535",
    "ves.io.schema.rules.repeated.max_items": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.uint32.lte": "65535",
    "ves.io.schema.rules.repeated.max_items": "64"
  }
}
```

<a id="schema-enable_forward_proxy--white_listed_prefixes"></a>

### white_listed_prefixes property

Type: `["list", "string"]`. Computed.

Traffic to these destination IP prefixes is not subjected to protocol parsing Example 'tmate' server
IP.

Upstream description:

Traffic to these destination IP prefixes is not subjected to protocol parsing Example "tmate" server
IP.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Next pages

- [enable_forward_proxy.no_interception](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/enable_forward_proxy/no_interception/)
- [enable_forward_proxy.tls_intercept](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/enable_forward_proxy/tls_intercept/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/)
- [xcsh_network_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/)
