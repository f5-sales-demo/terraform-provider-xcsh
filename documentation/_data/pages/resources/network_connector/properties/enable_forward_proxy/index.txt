---
page_title: "enable_forward_proxy"
subcategory: "Networking"
description: "Fine tune forward proxy behavior Few configurations allowed are White listed ports and IP prefixes: Forward proxy does application protocol detection and server name(SNI) detection by peeking into the traffic on the incoming downstream connection. Few protocols doesn't have client sending the first data. In such"
xcsh_docs: {"aliases": ["duration", "enable forward proxy"], "body_bytes": 6930, "body_sha256": "sha256:1fc35e6d59d772ff30e5bcf6e52dd14123c927cde78db02191336abb9cf08f4b", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:network_connector:properties:enable_forward_proxy:no_interception", "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:network_connector:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy", "parent_id": "xcsh-docs:resources:network_connector:reference", "path": "documentation/resources/network_connector/properties/enable_forward_proxy/index.md", "product": "distributed-cloud", "provider_name": "network_connector", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-1101110030313313-2102212303300111-0320130013022202-3311211123122310-2200103232211301-3223222300011312-2202123000112213-2220311201321020", "registry_path": "docs/guides/resources--network_connector--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "enable_forward_proxy:ConflictingObjectAttributes:no_interception,tls_intercept", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:no_interception", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "enable_forward_proxy:ConflictingObjectAttributes:no_interception,tls_intercept", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["enable_forward_proxy"], "schema_version": 1, "sections": [{"aliases": ["duration", "enable forward proxy connection timeout"], "anchor": "schema-enable_forward_proxy--connection_timeout", "description": "The timeout for new network connections to upstream server. This is specified in milliseconds. The default value is 2000 (2 seconds)", "document_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_forward_proxy", "connection_timeout"], "syntax": "attribute", "type": "number"}, {"aliases": ["enable forward proxy max connect attempts"], "anchor": "schema-enable_forward_proxy--max_connect_attempts", "description": "Specifies the allowed number of retries on connect failure to upstream server. Defaults to 1.", "document_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_forward_proxy", "max_connect_attempts"], "syntax": "attribute", "type": "number"}, {"aliases": ["enable forward proxy no interception"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:no_interception", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_forward_proxy", "no_interception"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable forward proxy tls intercept"], "anchor": "section", "description": "Configuration to enable TLS interception.", "document_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-enable_forward_proxy--tls_intercept--trusted_ca_url", "enforcement": "provider-schema", "group": "enable_forward_proxy.tls_intercept:ConflictingObjectAttributes:trusted_ca_url,volterra_trusted_ca", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "enable_forward_proxy.tls_intercept:ConflictingObjectAttributes:custom_certificate,volterra_certificate", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:custom_certificate", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "enable_forward_proxy.tls_intercept:ConflictingObjectAttributes:enable_for_all_domains,policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:enable_for_all_domains", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "enable_forward_proxy.tls_intercept:ConflictingObjectAttributes:enable_for_all_domains,policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:policy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "enable_forward_proxy.tls_intercept:ConflictingObjectAttributes:custom_certificate,volterra_certificate", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:volterra_certificate", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "enable_forward_proxy.tls_intercept:ConflictingObjectAttributes:trusted_ca_url,volterra_trusted_ca", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:volterra_trusted_ca", "type": "conflicts"}], "schema_path": ["enable_forward_proxy", "tls_intercept"], "syntax": "block", "type": "object"}, {"aliases": ["enable forward proxy white listed ports"], "anchor": "schema-enable_forward_proxy--white_listed_ports", "description": "Traffic to these destination TCP ports is not subjected to protocol parsing Example \"tmate\" server port.", "document_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_forward_proxy", "white_listed_ports"], "syntax": "attribute", "type": "list"}, {"aliases": ["enable forward proxy white listed prefixes"], "anchor": "schema-enable_forward_proxy--white_listed_prefixes", "description": "Traffic to these destination IP prefixes is not subjected to protocol parsing Example \"tmate\" server IP.", "document_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_forward_proxy", "white_listed_prefixes"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_connector/properties/enable_forward_proxy/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Fine tune forward proxy behavior Few configurations allowed are White listed ports and IP prefixes: Forward proxy does application protocol detection and server name(SNI) detection by peeking into the traffic on the incoming downstream connection. Few protocols doesn't have client sending the first data. In such", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["network_connectorCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_forward_proxy

Breadcrumbs:

- [xcsh_network_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/)
- enable_forward_proxy

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Fine tune forward proxy behavior

Few configurations allowed are

White listed ports and IP prefixes: Forward proxy does application protocol detection and server
name(SNI) detection by peeking into the traffic on the incoming downstream connection. Few protocols
doesn't have client sending the first data. In such cases, protocol and SNI detection fails. This
configuration allows, skipping protocol and SNI detection for whitelisted IP-prefix-list and ports
connection\_timeout: The timeout for new network connections to upstream server.
Max\_connect\_attempts: Maximum number of attempts made to make new network connection to upstream
server.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_interception",
    "tls_intercept")}
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
  "x-ves-oneof-field-tls_interception_choice": "[\"no_interception\",\"tls_intercept\"]"
}
```

Terraform syntax:

```terraform
enable_forward_proxy {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-enable_forward_proxy--connection_timeout"></a>

### connection_timeout property

Type: `"number"`. Optional.

The timeout for new network connections to upstream server. This is specified in milliseconds. The
(2 seconds). Defaults to \`2000\`.

Additional upstream details:

The default value is 2000 (2 seconds)

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.AtMost(600000),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

Type: `"number"`. Optional.

Specifies the allowed number of retries on connect failure to upstream server. Defaults to \`1\`.

Additional upstream details:

Defaults to 1.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.AtMost(8),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [no_interception](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/no_interception/): complete subsection reference.

- [tls_intercept](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/): complete subsection reference.

<a id="schema-enable_forward_proxy--white_listed_ports"></a>

### white_listed_ports property

Type: `["list", "number"]`. Optional.

Traffic to these destination TCP ports is not subjected to protocol parsing Example 'tmate' server
port.

Additional upstream details:

Traffic to these destination TCP ports is not subjected to protocol parsing Example "tmate" server
port.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

Type: `["list", "string"]`. Optional.

Traffic to these destination IP prefixes is not subjected to protocol parsing Example 'tmate' server
IP.

Additional upstream details:

Traffic to these destination IP prefixes is not subjected to protocol parsing Example "tmate" server
IP.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
