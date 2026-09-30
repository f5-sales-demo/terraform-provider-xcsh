---
page_title: "enable_forward_proxy"
subcategory: "Networking"
description: "enable_forward_proxy for xcsh_network_connector."
xcsh_docs: {"aliases": [], "body_bytes": 6936, "body_sha256": "sha256:3224b15e509cf45057d047892ccce65cfd33d9c1350f94f896320fc3d7279fe9", "child_ids": ["xcsh-docs:data-sources:network_connector:properties:enable_forward_proxy:no_interception", "xcsh-docs:data-sources:network_connector:properties:enable_forward_proxy:tls_intercept"], "collection_id": "xcsh-docs:data-sources:network_connector:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_connector:properties:enable_forward_proxy", "parent_id": "xcsh-docs:data-sources:network_connector:reference", "path": "documentation/data-sources/network_connector/properties/enable_forward_proxy/index.md", "provider_name": "network_connector", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["enable_forward_proxy"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_connector/properties/enable_forward_proxy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "enable_forward_proxy for xcsh_network_connector.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_connectorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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
