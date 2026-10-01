---
page_title: "xcsh_bigip_http_proxy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bigip_http_proxy reference."
---

# xcsh_bigip_http_proxy reference

<a id="canonical-82d6325bbc7490554ab16b1f4abf3a164d8a808f94642753626bb210c535cc01"></a>

## domains property — proxy_config / 61979ee4e669 / 4

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

- [http](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0527c8a402207963cd1087c7d402ed9c01dd1c393607dc347fed1de79199cf85): complete subsection reference.

- [https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079): complete subsection reference.

- [https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-d31637bdd4b16cbe32ef66cceeb81a16f39e53c629c297d234c78d6c5b502e7c): complete subsection reference.

<a id="canonical-e08d46679817c3476bd8d448fde9bf2492031737ed403aa7971a0ecf66ea2f2f"></a>

## Next pages — proxy_config / 61979ee4e669 / 5

- [proxy_config.http](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0527c8a402207963cd1087c7d402ed9c01dd1c393607dc347fed1de79199cf85)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-d31637bdd4b16cbe32ef66cceeb81a16f39e53c629c297d234c78d6c5b502e7c)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-0527c8a402207963cd1087c7d402ed9c01dd1c393607dc347fed1de79199cf85"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ef40d6b9d500bff10b4cc77e2b2cfc742f22810b75016ccd6311c92d03cddde5"></a>

## proxy_config.http — proxy_config.http / ee8e2823b57b / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- proxy_config.http

<a id="canonical-21eb2dfc95fb160fb5553439ab04f01ab33308041b081c0c6d61bb97fd7dddc1"></a>

Type: `"single"`. Computed.

HTTP Choice. Choice for selecting HTTP proxy.

Upstream description:

Choice for selecting HTTP proxy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]"
}
```

<a id="canonical-4909a6fc63d80ab55dfd93ccbe5475fe8256e600a9e4864f53eab9c5a4cf05b8"></a>

## Direct properties — proxy_config.http / ee8e2823b57b / 3

<a id="canonical-648743fe43d44e0cdf5142b4029cc7a4220dba065d3c52e07f89f34c0a78a8b8"></a>

<a id="canonical-ce4ce37c051f917126c7f8723c225e5669b169175452422fdcda1b8e119ad392"></a>

## dns_volterra_managed property — proxy_config.http / ee8e2823b57b / 4

Type: `"bool"`. Computed.

DNS records for domains will be managed automatically by F5 Distributed Cloud. As a prerequisite,
the domain must be delegated to F5 Distributed Cloud using Delegated domain feature or a DNS CNAME
record should be created in your DNS provider's portal.

Upstream description:

DNS records for domains will be managed automatically by F5 Distributed Cloud. As a prerequisite,
the domain must be delegated to F5 Distributed Cloud using Delegated domain feature or a DNS CNAME
record should be created in your DNS provider's portal.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-4b039b90cf6df9350fc9d1b2fcd716e89f61ed0c2f3f7794a95dae36942813be"></a>

<a id="canonical-3eaeb97bd0d6b354d11571a21f0d3258ce768e925699fb4dcd96ff1604ebcdc0"></a>

## port property — proxy_config.http / ee8e2823b57b / 5

Type: `"number"`. Computed.

Exclusive with \[port\_ranges\] HTTP port to Listen.

Upstream description:

Exclusive with \[port\_ranges\] HTTP port to Listen.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-ee0cf97a35d039d283a1fc50d8f7cdb7443ee6dd25096a243501e312aa768962"></a>

<a id="canonical-83b131c03989adcce9f18da3edb42e86399eba6b4ee714ef027ffb2aa401b86b"></a>

## port_ranges property — proxy_config.http / ee8e2823b57b / 6

Type: `"string"`. Computed.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by "-".

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  }
}
```

<a id="canonical-85d84f5cc4aa052d737a7d6c426c409be03b4e24bc36346dff3444288559157e"></a>

## Next pages — proxy_config.http / ee8e2823b57b / 7

- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e30b782ba793582c0485f9992b34be0d9bb76cde5f1f34845740c0e0cf08f131"></a>

## proxy_config.https — proxy_config.https / ef6a9e4d00e4 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- proxy_config.https

<a id="canonical-dd901f55626c9e948b2f573c2c432c7ee4ba082fed149ee32bc7ad169d816664"></a>

Type: `"single"`. Computed.

Choice for selecting HTTP proxy with bring your own certificates.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-default_lb_choice": "[\"default_loadbalancer\",\"non_default_loadbalancer\"]",
  "x-ves-oneof-field-path_normalize_choice": "[\"disable_path_normalize\",\"enable_path_normalize\"]",
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]",
  "x-ves-oneof-field-server_header_choice": "[\"append_server_name\",\"default_header\",\"pass_through\",\"server_name\"]",
  "x-ves-oneof-field-tls_certificates_choice": "[\"tls_cert_params\",\"tls_parameters\"]"
}
```

<a id="canonical-76cec0efddbf351bacb2e48374765a931d9bcbb8eb3d26db6528edc0e70f160f"></a>

## Direct properties — proxy_config.https / ef6a9e4d00e4 / 3

<a id="canonical-4ff823f5346ae746c06df05a5f3dc9d846889c3457913b5ed8532a5f6c3e520e"></a>

<a id="canonical-4a725096ad03f3495ba52c11de86a4d170c2c97ba34472b271ca1961140cf51e"></a>

## add_hsts property — proxy_config.https / ef6a9e4d00e4 / 4

Type: `"bool"`. Computed.

Add HTTP Strict-Transport-Security response header.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2b8ed58e9211d74604aa818eb866052d21b76624f7da499f93bc7668ba9fd394"></a>

<a id="canonical-db0d597c42b3fb7a32c4a254937f059467779cc5bf97c4375c443f4ec03babf2"></a>

## append_server_name property — proxy_config.https / ef6a9e4d00e4 / 5

Type: `"string"`. Computed.

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

Upstream description:

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [coalescing_options](data-sources--bigip_http_proxy--reference--group-003.md#canonical-44a2e548065b20343b44d2d05d1784f3a79419b49b51b25f30748f6173aee9ba): complete subsection reference.

<a id="canonical-aff1d4b427548484d39a43d47b072bfeccc792915077d00bb1496d3e98d12812"></a>

<a id="canonical-2c51770ebd30f252df78ad9ff6a081f14a844565994b765c947e6b432bbc2457"></a>

## connection_idle_timeout property — proxy_config.https / ef6a9e4d00e4 / 6

Type: `"number"`. Computed.

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed.

Upstream description:

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed. Note
that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive. This is
specified in milliseconds. The default value is 2 minutes.

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

- [default_header](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3d6444196be1b482bd50b6e5b35b7aff9fd17c1a82505e5c6f219930418578f0): complete subsection reference.

- [default_loadbalancer](data-sources--bigip_http_proxy--reference--group-003.md#canonical-77c4cc3b353c2be550f7694552e03d90d0f88ed113d978a8ad8b64ebcea66d7e): complete subsection reference.

- [disable_path_normalize](data-sources--bigip_http_proxy--reference--group-003.md#canonical-c5b75ed1bb8c285180b01609b7de1e5e37313edf185e3ae52b83cbf431dab98a): complete subsection reference.

- [enable_path_normalize](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2a4ab511b50266604ad43b69f15132243adba43a4b1467917c18eefd5236d683): complete subsection reference.

- [http_protocol_options](data-sources--bigip_http_proxy--reference--group-003.md#canonical-671fa1cb101593849578d460c6f25aefaf1bef75e66fa627b13d8c68c2091504): complete subsection reference.

<a id="canonical-20908927e7d18777e9e355949006236eb52873ec25ec29a29d9b51a375efde8b"></a>

<a id="canonical-c8b62e4f4c17a48ea29c30baf78807e59a17eb5a19fa584be027ced65a0165e6"></a>

## http_redirect property — proxy_config.https / ef6a9e4d00e4 / 7

Type: `"bool"`. Computed.

HTTP Redirect to HTTPS. Redirect HTTP traffic to HTTPS.

Upstream description:

Redirect HTTP traffic to HTTPS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [non_default_loadbalancer](data-sources--bigip_http_proxy--reference--group-003.md#canonical-085303cfcdcea8d83f14c48e6b89e3a10995dee8912d5f866f76eb25e28fcca1): complete subsection reference.

- [pass_through](data-sources--bigip_http_proxy--reference--group-003.md#canonical-22efcd0f6f8d570d38450e7b024e23ab7ba8498a55ed48ee2e29e373568ffc6a): complete subsection reference.

<a id="canonical-51a1da62662fb9e82c7cf08c7c443c3fb43963456752979757317e4712116b08"></a>

<a id="canonical-2e9c302741f9f9144a960b9384e4df72189d50d1de5583e4df0c3b714f6e106c"></a>

## port property — proxy_config.https / ef6a9e4d00e4 / 8

Type: `"number"`. Computed.

Exclusive with \[port\_ranges\] HTTPS port to Listen.

Upstream description:

Exclusive with \[port\_ranges\] HTTPS port to Listen.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-4a4551e56df180e7a67eb1032cd755a159ee0bc3b2155b13191238c931691a75"></a>

<a id="canonical-0e4de7cddcff2bc1e479874e7fc2fe1c6b9c983be8e5816f20349b4a5e680668"></a>

## port_ranges property — proxy_config.https / ef6a9e4d00e4 / 9

Type: `"string"`. Computed.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by "-".

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  }
}
```

<a id="canonical-7ea2faea4d0ed8df2108c81b38269d95fdf0b1a33ed3ede4857b6278a4a1656a"></a>

<a id="canonical-1f3fd18a7976395e114ccba8dc6c143a8389b1d62123f74b54d080689a4e8038"></a>

## server_name property — proxy_config.https / ef6a9e4d00e4 / 10

Type: `"string"`. Computed.

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

Upstream description:

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [tls_cert_params](data-sources--bigip_http_proxy--reference--group-003.md#canonical-9c4118e7d008c6f588dc371e243ef7160412df5c78621103437520f1bc1cc038): complete subsection reference.

- [tls_parameters](data-sources--bigip_http_proxy--reference--group-003.md#canonical-12b1668564b57fed383102e0655e5c2fd223f895bd3f947a1500d20e88bd4b90): complete subsection reference.

<a id="canonical-0987d7e16d43956fed51d3b856c65aace0c864250eef95909c28eb728b7e4e87"></a>

## Next pages — proxy_config.https / ef6a9e4d00e4 / 11

- [proxy_config.https.coalescing_options](data-sources--bigip_http_proxy--reference--group-003.md#canonical-44a2e548065b20343b44d2d05d1784f3a79419b49b51b25f30748f6173aee9ba)
- [proxy_config.https.default_header](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3d6444196be1b482bd50b6e5b35b7aff9fd17c1a82505e5c6f219930418578f0)
- [proxy_config.https.default_loadbalancer](data-sources--bigip_http_proxy--reference--group-003.md#canonical-77c4cc3b353c2be550f7694552e03d90d0f88ed113d978a8ad8b64ebcea66d7e)
- [proxy_config.https.disable_path_normalize](data-sources--bigip_http_proxy--reference--group-003.md#canonical-c5b75ed1bb8c285180b01609b7de1e5e37313edf185e3ae52b83cbf431dab98a)
- [proxy_config.https.enable_path_normalize](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2a4ab511b50266604ad43b69f15132243adba43a4b1467917c18eefd5236d683)
- [proxy_config.https.http_protocol_options](data-sources--bigip_http_proxy--reference--group-003.md#canonical-671fa1cb101593849578d460c6f25aefaf1bef75e66fa627b13d8c68c2091504)
- [proxy_config.https.non_default_loadbalancer](data-sources--bigip_http_proxy--reference--group-003.md#canonical-085303cfcdcea8d83f14c48e6b89e3a10995dee8912d5f866f76eb25e28fcca1)
- [proxy_config.https.pass_through](data-sources--bigip_http_proxy--reference--group-003.md#canonical-22efcd0f6f8d570d38450e7b024e23ab7ba8498a55ed48ee2e29e373568ffc6a)
- [proxy_config.https.tls_cert_params](data-sources--bigip_http_proxy--reference--group-003.md#canonical-9c4118e7d008c6f588dc371e243ef7160412df5c78621103437520f1bc1cc038)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-003.md#canonical-12b1668564b57fed383102e0655e5c2fd223f895bd3f947a1500d20e88bd4b90)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-44a2e548065b20343b44d2d05d1784f3a79419b49b51b25f30748f6173aee9ba"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ca91ae5a8102e3f9571dc86fd1301b02cf5c19fe681e82da272d88b182e4d0d9"></a>

## proxy_config.https.coalescing_options — proxy_config.https.coalescing_options / fdb5a064105d / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- proxy_config.https.coalescing_options

<a id="canonical-2bb45ed6c0b9beea3ebf4634c1330f2dff816e6abf44ef6bef18f0174751843a"></a>

Type: `"single"`. Computed.

TLS connection coalescing configuration (not compatible with mTLS).

Upstream description:

TLS connection coalescing configuration (not compatible with mTLS)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-coalescing_choice": "[\"default_coalescing\",\"strict_coalescing\"]"
}
```

<a id="canonical-2afd92d103a8be9a496bce01d7d719007c5f045c1b2e928bfb5283b9ca260842"></a>

## Direct properties — proxy_config.https.coalescing_options / fdb5a064105d / 3

- [default_coalescing](data-sources--bigip_http_proxy--reference--group-003.md#canonical-21c50d716ca2d01baf90db9953e185d3c08d985ecc72f1aeb443d088f913a660): complete subsection reference.

- [strict_coalescing](data-sources--bigip_http_proxy--reference--group-003.md#canonical-c0443001841f2684c435dd9b41939aef86627ae2e826715a8e77fb2b285ad23c): complete subsection reference.

<a id="canonical-f490263cb66dcc4a70e2de215e099e46b05f3e5f7b2e51362093218da7e601bb"></a>

## Next pages — proxy_config.https.coalescing_options / fdb5a064105d / 4

- [proxy_config.https.coalescing_options.default_coalescing](data-sources--bigip_http_proxy--reference--group-003.md#canonical-21c50d716ca2d01baf90db9953e185d3c08d985ecc72f1aeb443d088f913a660)
- [proxy_config.https.coalescing_options.strict_coalescing](data-sources--bigip_http_proxy--reference--group-003.md#canonical-c0443001841f2684c435dd9b41939aef86627ae2e826715a8e77fb2b285ad23c)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-21c50d716ca2d01baf90db9953e185d3c08d985ecc72f1aeb443d088f913a660"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a4db09cb3857c8728a2ab30e3b84acfcdccf386e1ac90156af063ec862d97646"></a>

## proxy_config.https.coalescing_options.default_coalescing — proxy_config.https.coalescing_options.default_coalescing / b2383c0e07ed / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- [proxy_config.https.coalescing_options](data-sources--bigip_http_proxy--reference--group-003.md#canonical-44a2e548065b20343b44d2d05d1784f3a79419b49b51b25f30748f6173aee9ba)
- proxy_config.https.coalescing_options.default_coalescing

<a id="canonical-de71922070fc1906b598ae6ffae880cd7c2df79126ddc940dd58a25694ddc202"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default coalescing.

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-37b84782932548c591557ed966b4b6a6fb05fac346327c84233680fbbfadc6eb"></a>

## Direct properties — proxy_config.https.coalescing_options.default_coalescing / b2383c0e07ed / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f20b61c36a17f75879c01f1ec0734eaead536378f5d0e54a91b935ba38ff9747"></a>

## Next pages — proxy_config.https.coalescing_options.default_coalescing / b2383c0e07ed / 4

- [proxy_config.https.coalescing_options](data-sources--bigip_http_proxy--reference--group-003.md#canonical-44a2e548065b20343b44d2d05d1784f3a79419b49b51b25f30748f6173aee9ba)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-c0443001841f2684c435dd9b41939aef86627ae2e826715a8e77fb2b285ad23c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c0e044fd9c5e5a1d1d54463e40fb88245139d509b49e372a7c1fc090e719ac32"></a>

## proxy_config.https.coalescing_options.strict_coalescing — proxy_config.https.coalescing_options.strict_coalescing / 5d89c17984d5 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- [proxy_config.https.coalescing_options](data-sources--bigip_http_proxy--reference--group-003.md#canonical-44a2e548065b20343b44d2d05d1784f3a79419b49b51b25f30748f6173aee9ba)
- proxy_config.https.coalescing_options.strict_coalescing

<a id="canonical-d2ff70092c3e5b95d821b2fcd78fb7cb75c711d0e3cb5b3230a78d633a033cd3"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for strict coalescing.

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-64a28b07c1324996dbee1368c9b75c4ecfe64f99f35281ae30f4971ea7c0af6a"></a>

## Direct properties — proxy_config.https.coalescing_options.strict_coalescing / 5d89c17984d5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c9258a3734de4806e7c103de4c5c9dd6f3185b7ac58cac0c9cbe051c44dd8743"></a>

## Next pages — proxy_config.https.coalescing_options.strict_coalescing / 5d89c17984d5 / 4

- [proxy_config.https.coalescing_options](data-sources--bigip_http_proxy--reference--group-003.md#canonical-44a2e548065b20343b44d2d05d1784f3a79419b49b51b25f30748f6173aee9ba)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-3d6444196be1b482bd50b6e5b35b7aff9fd17c1a82505e5c6f219930418578f0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a082378d23b1408f7da28ac5357a50b57ae6618dfaa5ef9ce65e59a7122de054"></a>

## proxy_config.https.default_header — proxy_config.https.default_header / 439a671bd8a7 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- proxy_config.https.default_header

<a id="canonical-f528313c3821e2f19438a3951bc0e57cd13b323e3c51e0bfdcd8ccbaf3a8ff0c"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default header.

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-4a28f9d3dc1fe95ff097343c9038b3ce7ca336924c30b547d2c0fdf6fe540336"></a>

## Direct properties — proxy_config.https.default_header / 439a671bd8a7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bdd4fa2569230c36a6b23679dc84cf61289072ba42b62e837f91586b5864864b"></a>

## Next pages — proxy_config.https.default_header / 439a671bd8a7 / 4

- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-77c4cc3b353c2be550f7694552e03d90d0f88ed113d978a8ad8b64ebcea66d7e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5f3608947cafae2606ed2c859b26e26b63493d02d075e44ae2a15bafa43f7589"></a>

## proxy_config.https.default_loadbalancer — proxy_config.https.default_loadbalancer / ceaf0cbd12fa / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- proxy_config.https.default_loadbalancer

<a id="canonical-b3b45d0999990021ad93fce611771658f08b3c728ca86748c38b13a415083477"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default loadbalancer.

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0fc75bde74766c3b98923e022c56231b81678141eea78d8d2d8e7b039a7bd893"></a>

## Direct properties — proxy_config.https.default_loadbalancer / ceaf0cbd12fa / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d2e56666b31a72c9312e0ea92ce4222e7232184154a7c71dcea70ad91d6cb40a"></a>

## Next pages — proxy_config.https.default_loadbalancer / ceaf0cbd12fa / 4

- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-c5b75ed1bb8c285180b01609b7de1e5e37313edf185e3ae52b83cbf431dab98a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-729ef4f605b483f05b81f98737b2cf4a888d69411f4ce7d5ad282ff5e7800194"></a>

## proxy_config.https.disable_path_normalize — proxy_config.https.disable_path_normalize / b1f367e0519c / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- proxy_config.https.disable_path_normalize

<a id="canonical-553211d3ccfa833e73bb06216caf83be870b78874807f9e4ff110b64816f0fd5"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-96affe054ff8d7444c49920ea06dbe42d0f8371d68437123b6824d295bedb637"></a>

## Direct properties — proxy_config.https.disable_path_normalize / b1f367e0519c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7757eb8607e654dea4ef23ab6185a9101e90fcb4bd0a03a3cefb0e0db1f29fda"></a>

## Next pages — proxy_config.https.disable_path_normalize / b1f367e0519c / 4

- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-2a4ab511b50266604ad43b69f15132243adba43a4b1467917c18eefd5236d683"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-39efb6815b370aef8218462926408e187b5e3b337627bd0ebb645a96558e74e4"></a>

## proxy_config.https.enable_path_normalize — proxy_config.https.enable_path_normalize / 65e532de6232 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- proxy_config.https.enable_path_normalize

<a id="canonical-adb7715efa237152995753e3637fb6abb6319af6f01852d3506a9569fd6547af"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-4713dd9b988a334e1e9437a9e8ec321150a32ea4a98d382fc91bfd12a7f9f30b"></a>

## Direct properties — proxy_config.https.enable_path_normalize / 65e532de6232 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9592b13743ecc5bcb096e583544fb215ed17f77e7ba56acc9d96384263d30bd3"></a>

## Next pages — proxy_config.https.enable_path_normalize / 65e532de6232 / 4

- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-671fa1cb101593849578d460c6f25aefaf1bef75e66fa627b13d8c68c2091504"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c65fc031e27a255758eefba73ed8603456711101a4414e5060d298167ec0992b"></a>

## proxy_config.https.http_protocol_options — proxy_config.https.http_protocol_options / ba9f79975c9a / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- proxy_config.https.http_protocol_options

<a id="canonical-88b8f8c01afc110747bc76859b953d5c3dcca88db9208418ce8cda417edde962"></a>

Type: `"single"`. Computed.

HTTP protocol configuration OPTIONS for downstream connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-http_protocol_choice": "[\"http_protocol_enable_v1_only\",\"http_protocol_enable_v1_v2\",\"http_protocol_enable_v2_only\"]"
}
```

<a id="canonical-430bd4cf192f8bcaa051d5f74b0dbe72c3256736fc18b6dd4c5a5679cd475388"></a>

## Direct properties — proxy_config.https.http_protocol_options / ba9f79975c9a / 3

- [http_protocol_enable_v1_only](data-sources--bigip_http_proxy--reference--group-003.md#canonical-ef2d68e74507b7af684ac53aa4c3201dc54f65891796b4b1aacc3ca11035e461): complete subsection reference.

- [http_protocol_enable_v1_v2](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2ab2428060c9fc2844cebf1b20ec47abfff4fcf27d14b7446674e831afcd6e77): complete subsection reference.

- [http_protocol_enable_v2_only](data-sources--bigip_http_proxy--reference--group-003.md#canonical-ed2039da05bb43a28d41024a8be556ab36dfc93090c5276bc731baa914fe7cea): complete subsection reference.

<a id="canonical-b92e1ebbc969d5a58fbfffa1cc9b928dd0480ade90dae33415c6a89f1998cf7a"></a>

## Next pages — proxy_config.https.http_protocol_options / ba9f79975c9a / 4

- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--bigip_http_proxy--reference--group-003.md#canonical-ef2d68e74507b7af684ac53aa4c3201dc54f65891796b4b1aacc3ca11035e461)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_v2](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2ab2428060c9fc2844cebf1b20ec47abfff4fcf27d14b7446674e831afcd6e77)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v2_only](data-sources--bigip_http_proxy--reference--group-003.md#canonical-ed2039da05bb43a28d41024a8be556ab36dfc93090c5276bc731baa914fe7cea)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-ef2d68e74507b7af684ac53aa4c3201dc54f65891796b4b1aacc3ca11035e461"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-053f1c9b62ed1934124d452f0b14c211aa84e2a78c91a5d1be9c9c98be62a929"></a>

## proxy_config.https.http_protocol_options.http_protocol_enable_v1_only — proxy_config.https.http_protocol_options.http_protocol_enable_v1_only / 18a5875b6f93 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- [proxy_config.https.http_protocol_options](data-sources--bigip_http_proxy--reference--group-003.md#canonical-671fa1cb101593849578d460c6f25aefaf1bef75e66fa627b13d8c68c2091504)
- proxy_config.https.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-36f0474f9725c1fff5ab944bf147e3462b63ce4207ab974a43cf07f854312ae8"></a>

Type: `"single"`. Computed.

HTTP/1.1 Protocol OPTIONS for downstream connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-fee0d721a35863d8c31812064393783911b53aed9f9efbd18719c7b6adbd5cf6"></a>

## Direct properties — proxy_config.https.http_protocol_options.http_protocol_enable_v1_only / 18a5875b6f93 / 3

- [header_transformation](data-sources--bigip_http_proxy--reference--group-003.md#canonical-7d4012e82ca07d8a322c69d96c29091ea164157db0df7b69c7fafe9826113421): complete subsection reference.

<a id="canonical-1ed816c36dac85d333c0b86642d15b4945307d323fd42282c29464fd748b4f60"></a>

## Next pages — proxy_config.https.http_protocol_options.http_protocol_enable_v1_only / 18a5875b6f93 / 4

- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--bigip_http_proxy--reference--group-003.md#canonical-7d4012e82ca07d8a322c69d96c29091ea164157db0df7b69c7fafe9826113421)
- [proxy_config.https.http_protocol_options](data-sources--bigip_http_proxy--reference--group-003.md#canonical-671fa1cb101593849578d460c6f25aefaf1bef75e66fa627b13d8c68c2091504)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-7d4012e82ca07d8a322c69d96c29091ea164157db0df7b69c7fafe9826113421"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8a626573ec192c9140c4e9b02d8398cc71f71408273f9fd6da0d9ee49b08684f"></a>

## proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation — proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_tra / 0d0aa9af3886 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- [proxy_config.https.http_protocol_options](data-sources--bigip_http_proxy--reference--group-003.md#canonical-671fa1cb101593849578d460c6f25aefaf1bef75e66fa627b13d8c68c2091504)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--bigip_http_proxy--reference--group-003.md#canonical-ef2d68e74507b7af684ac53aa4c3201dc54f65891796b4b1aacc3ca11035e461)
- proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-743bfceeb01bbbd69e2135921a9d8201fd5a7172135258b71021ba464d3593f8"></a>

Type: `"single"`. Computed.

Header Transformation OPTIONS for HTTP/1.1 request/response headers.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-header_transformation_choice": "[\"default_header_transformation\",\"preserve_case_header_transformation\",\"proper_case_header_transformation\"]"
}
```

<a id="canonical-e196e7ccfd87e9898d006a8c0a20dbec2b54854af221b485a47ac16252e65e72"></a>

## Direct properties — proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_tra / 0d0aa9af3886 / 3

- [default_header_transformation](data-sources--bigip_http_proxy--reference--group-003.md#canonical-00d1bbc189b84817ae57919cd6eddeefda9042e58d4a35c34b86b2929f5e5a08): complete subsection reference.

- [preserve_case_header_transformation](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2b886a83e4b1416206c88156afdc6b6b4937a7657bc2d39039bcc7fc7bcf7f46): complete subsection reference.

- [proper_case_header_transformation](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1bb6109a75bd879ccaa141c0959b5752819e2a85ac858c795170a72024659c1c): complete subsection reference.

<a id="canonical-68fe3a9c2e6b8dd2330b2ba530aa788382b28d42ce2d848c012d88f005011015"></a>

## Next pages — proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_tra / 0d0aa9af3886 / 4

- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](data-sources--bigip_http_proxy--reference--group-003.md#canonical-00d1bbc189b84817ae57919cd6eddeefda9042e58d4a35c34b86b2929f5e5a08)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2b886a83e4b1416206c88156afdc6b6b4937a7657bc2d39039bcc7fc7bcf7f46)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1bb6109a75bd879ccaa141c0959b5752819e2a85ac858c795170a72024659c1c)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--bigip_http_proxy--reference--group-003.md#canonical-ef2d68e74507b7af684ac53aa4c3201dc54f65891796b4b1aacc3ca11035e461)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-00d1bbc189b84817ae57919cd6eddeefda9042e58d4a35c34b86b2929f5e5a08"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aaa9539d348bb4ae39d748c72cc03c69463d0bb7ac33a832ed18b6260902a968"></a>

## proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation — proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_tra / cc8b15a9a2a8 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- [proxy_config.https.http_protocol_options](data-sources--bigip_http_proxy--reference--group-003.md#canonical-671fa1cb101593849578d460c6f25aefaf1bef75e66fa627b13d8c68c2091504)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--bigip_http_proxy--reference--group-003.md#canonical-ef2d68e74507b7af684ac53aa4c3201dc54f65891796b4b1aacc3ca11035e461)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--bigip_http_proxy--reference--group-003.md#canonical-7d4012e82ca07d8a322c69d96c29091ea164157db0df7b69c7fafe9826113421)
- proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-c28a8c7d4bb1aaf35bf4b122b5f9e6e6318b0645b6e3b923b66b4177c05ba008"></a>

Type: `["object", {}]`. Computed.

Use the platform's current default HTTP header transformation behavior.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-e730581d1d5b39bf7528ccff58d7a468472c04b578bfd637786740dce08878f7"></a>

## Direct properties — proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_tra / cc8b15a9a2a8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f36a813e0a24e15c6de0557f25c67388ccf3c5fb5f67566d56ea41c3fb3442cb"></a>

## Next pages — proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_tra / cc8b15a9a2a8 / 4

- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--bigip_http_proxy--reference--group-003.md#canonical-7d4012e82ca07d8a322c69d96c29091ea164157db0df7b69c7fafe9826113421)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-2b886a83e4b1416206c88156afdc6b6b4937a7657bc2d39039bcc7fc7bcf7f46"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-02a572a8282bd75977649cad05125abc876b602c7493729fa0c9f4269459e9bb"></a>

## proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation — proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_tra / 5791887b48b6 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- [proxy_config.https.http_protocol_options](data-sources--bigip_http_proxy--reference--group-003.md#canonical-671fa1cb101593849578d460c6f25aefaf1bef75e66fa627b13d8c68c2091504)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--bigip_http_proxy--reference--group-003.md#canonical-ef2d68e74507b7af684ac53aa4c3201dc54f65891796b4b1aacc3ca11035e461)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--bigip_http_proxy--reference--group-003.md#canonical-7d4012e82ca07d8a322c69d96c29091ea164157db0df7b69c7fafe9826113421)
- proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-9ed3e912acfd1c8a71e5744533de1b75ae134373b19afeb0d69731c6cdfe527f"></a>

Type: `["object", {}]`. Computed.

Preserve HTTP header-name case when upstream case must remain unchanged.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-93fc7d4774a3dfc472021aaa9fa53348a80a5746ea7d3d0d3376e044d6d6b0de"></a>

## Direct properties — proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_tra / 5791887b48b6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f6710546dcf33964ef815269cb62e14e7ba763a10b22eac222ab6f4fda273011"></a>

## Next pages — proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_tra / 5791887b48b6 / 4

- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--bigip_http_proxy--reference--group-003.md#canonical-7d4012e82ca07d8a322c69d96c29091ea164157db0df7b69c7fafe9826113421)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-1bb6109a75bd879ccaa141c0959b5752819e2a85ac858c795170a72024659c1c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-baca5ab4e9152cac0c0280fb9fc85e2977fe90ba1eaae1ee378992c58b596f08"></a>

## proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation — proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_tra / 12898e890298 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- [proxy_config.https.http_protocol_options](data-sources--bigip_http_proxy--reference--group-003.md#canonical-671fa1cb101593849578d460c6f25aefaf1bef75e66fa627b13d8c68c2091504)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--bigip_http_proxy--reference--group-003.md#canonical-ef2d68e74507b7af684ac53aa4c3201dc54f65891796b4b1aacc3ca11035e461)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--bigip_http_proxy--reference--group-003.md#canonical-7d4012e82ca07d8a322c69d96c29091ea164157db0df7b69c7fafe9826113421)
- proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-0abd2ef9a2719279ca5e96069f430c99f6c10ec8bdc458d6a45d3431fda352a6"></a>

Type: `["object", {}]`. Computed.

Transform HTTP header names to proper case when explicit transformation is required.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-46f6ca868c85ffc43bde58b0d610af2ed45cf5c86aad9dd69319616ca736bb37"></a>

## Direct properties — proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_tra / 12898e890298 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5ab3ca78481619ead27275e23040f1c00288acb9a125cb613a937162548d1e32"></a>

## Next pages — proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_tra / 12898e890298 / 4

- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--bigip_http_proxy--reference--group-003.md#canonical-7d4012e82ca07d8a322c69d96c29091ea164157db0df7b69c7fafe9826113421)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-2ab2428060c9fc2844cebf1b20ec47abfff4fcf27d14b7446674e831afcd6e77"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5f0f24a841271e903fcf807939b43bdc1f22e2599962db1d4fb2a8ab7f29cfa0"></a>

## proxy_config.https.http_protocol_options.http_protocol_enable_v1_v2 — proxy_config.https.http_protocol_options.http_protocol_enable_v1_v2 / 435d95a98dde / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- [proxy_config.https.http_protocol_options](data-sources--bigip_http_proxy--reference--group-003.md#canonical-671fa1cb101593849578d460c6f25aefaf1bef75e66fa627b13d8c68c2091504)
- proxy_config.https.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-a7746050171947b18e44db6f304addc90fda2e120e51c7a838fc1425b6a79e37"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for http protocol enable v1 v2.

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-19451fd611fd497032d415e9fecce3492afc8e524b8b8fb8cf83546d153598b8"></a>

## Direct properties — proxy_config.https.http_protocol_options.http_protocol_enable_v1_v2 / 435d95a98dde / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3ecd1310afddfc9a9b4b92c2f348ea381ab31a0042ddbef6e263b79cac776124"></a>

## Next pages — proxy_config.https.http_protocol_options.http_protocol_enable_v1_v2 / 435d95a98dde / 4

- [proxy_config.https.http_protocol_options](data-sources--bigip_http_proxy--reference--group-003.md#canonical-671fa1cb101593849578d460c6f25aefaf1bef75e66fa627b13d8c68c2091504)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-ed2039da05bb43a28d41024a8be556ab36dfc93090c5276bc731baa914fe7cea"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-62b243ea96c822a674ae047367eb71fc84280846aed4230f23c608e96c6099ed"></a>

## proxy_config.https.http_protocol_options.http_protocol_enable_v2_only — proxy_config.https.http_protocol_options.http_protocol_enable_v2_only / 00db46f7ab33 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- [proxy_config.https.http_protocol_options](data-sources--bigip_http_proxy--reference--group-003.md#canonical-671fa1cb101593849578d460c6f25aefaf1bef75e66fa627b13d8c68c2091504)
- proxy_config.https.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-ae7b3dd620008e85b67443ac5fa10cbbb78ed79d50d95ab7881e9a60e6bea7eb"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for http protocol enable v2 only.

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-91bc3cefc2239ed03aeec501cd0ce50845a651320068aa3db8a748165b03aa60"></a>

## Direct properties — proxy_config.https.http_protocol_options.http_protocol_enable_v2_only / 00db46f7ab33 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6fa85187943b2e8b28fe04694a56a7a0e30064ee3257a1b69a4246d54ad6ff65"></a>

## Next pages — proxy_config.https.http_protocol_options.http_protocol_enable_v2_only / 00db46f7ab33 / 4

- [proxy_config.https.http_protocol_options](data-sources--bigip_http_proxy--reference--group-003.md#canonical-671fa1cb101593849578d460c6f25aefaf1bef75e66fa627b13d8c68c2091504)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-085303cfcdcea8d83f14c48e6b89e3a10995dee8912d5f866f76eb25e28fcca1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-250d36e27210b83b4df606d86528735da3ea23adfa1bba36d8de01b788e4291f"></a>

## proxy_config.https.non_default_loadbalancer — proxy_config.https.non_default_loadbalancer / 3d4b9d11f4ea / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- proxy_config.https.non_default_loadbalancer

<a id="canonical-24aef35840aaafdeae16289b7b8b67a4e8a64e5499bea8d0422145f7f8afb136"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for non default loadbalancer.

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-c4d5f88db2ac85c1ed58c14b837a0a520c77c453d4876b3fcbcb396d803dd97b"></a>

## Direct properties — proxy_config.https.non_default_loadbalancer / 3d4b9d11f4ea / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c241afee7a852906cdbf1af94351807789f3a014a41d1e1c3b466e8994df3a7e"></a>

## Next pages — proxy_config.https.non_default_loadbalancer / 3d4b9d11f4ea / 4

- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-22efcd0f6f8d570d38450e7b024e23ab7ba8498a55ed48ee2e29e373568ffc6a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d113ac63329795384126ea88e55b415ba007fb397f1db4ef4c3e37fda0837872"></a>

## proxy_config.https.pass_through — proxy_config.https.pass_through / 608c7c129d4b / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- proxy_config.https.pass_through

<a id="canonical-3e2a806e783dd1d3d4ff68e9e156cf9348d83fa690c080b6be0b9614f61067b9"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for pass through.

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2899bcaa705bc8128a441d83f69f87fe53300d346de3469df19eab4b1254f24d"></a>

## Direct properties — proxy_config.https.pass_through / 608c7c129d4b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f106db5bb6964614c5661006270058c81e33b79140e10bf702f41eeb5e9a308c"></a>

## Next pages — proxy_config.https.pass_through / 608c7c129d4b / 4

- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-9c4118e7d008c6f588dc371e243ef7160412df5c78621103437520f1bc1cc038"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e3e5f32f9e112a8416f532ea958fc7e470cb9cc482864c58d0ed0a9a8f9091df"></a>

## proxy_config.https.tls_cert_params — proxy_config.https.tls_cert_params / 42e2a81c7d0b / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- proxy_config.https.tls_cert_params

<a id="canonical-dad39f9ccbfc1c7a0ab0fc441dc66ffc5ac7eb3e6b0866318b196507ac076f6f"></a>

Type: `"single"`. Computed.

Configuration parameter for tls cert params.

Upstream description:

Select TLS Parameters and Certificates.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

<a id="canonical-f7e883b6c938e3eccdc99d54bc9cd14dfd223e8ecc3cd305b892a499ca091ab8"></a>

## Direct properties — proxy_config.https.tls_cert_params / 42e2a81c7d0b / 3

- [certificates](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0aeeefe346a9b0a474fed3f5269e9ff3589f9c77029bc79c1f033d25ee7ff8e0): complete subsection reference.

- [no_mtls](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1801b48b94d3f41c7b779531d6ae5c7b5795025ba0540f044791aa8310473ac5): complete subsection reference.

- [tls_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-ae2556bad9d17e008ae3b5ddd221b4299262c0dbe1d204357303e644fd15660b): complete subsection reference.

- [use_mtls](data-sources--bigip_http_proxy--reference--group-003.md#canonical-7769bad9d339b8f2d09e15b5f21aec04f7c8375af5555e2851e4dd948ea8cf73): complete subsection reference.

<a id="canonical-16923964fe9eab3ed46788ec4fa2fb85342f56ff7475f4072418355b1f82ca4b"></a>

## Next pages — proxy_config.https.tls_cert_params / 42e2a81c7d0b / 4

- [proxy_config.https.tls_cert_params.certificates](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0aeeefe346a9b0a474fed3f5269e9ff3589f9c77029bc79c1f033d25ee7ff8e0)
- [proxy_config.https.tls_cert_params.no_mtls](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1801b48b94d3f41c7b779531d6ae5c7b5795025ba0540f044791aa8310473ac5)
- [proxy_config.https.tls_cert_params.tls_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-ae2556bad9d17e008ae3b5ddd221b4299262c0dbe1d204357303e644fd15660b)
- [proxy_config.https.tls_cert_params.use_mtls](data-sources--bigip_http_proxy--reference--group-003.md#canonical-7769bad9d339b8f2d09e15b5f21aec04f7c8375af5555e2851e4dd948ea8cf73)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-0aeeefe346a9b0a474fed3f5269e9ff3589f9c77029bc79c1f033d25ee7ff8e0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-63265fca8b135207cc03ae0c4f1a763a050db7bc2ddfdf187553bfe96949732c"></a>

## proxy_config.https.tls_cert_params.certificates — proxy_config.https.tls_cert_params.certificates / 22a526e05e5b / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- [proxy_config.https.tls_cert_params](data-sources--bigip_http_proxy--reference--group-003.md#canonical-9c4118e7d008c6f588dc371e243ef7160412df5c78621103437520f1bc1cc038)
- proxy_config.https.tls_cert_params.certificates

<a id="canonical-0611780e6bd773be16d67af6f7b5fd6d452d0734ae716ca368d0ae69c777cc49"></a>

Type: `"list"`. Computed.

Select one or more certificates with any domain names.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-b63e780b3b6e5320ff33895dffe807d2b98da5c24784b42217cfe923b4c75a6e"></a>

## Direct properties — proxy_config.https.tls_cert_params.certificates / 22a526e05e5b / 3

<a id="canonical-d798df0c0f749383d960edfedc3ec8916ab68d9a86f7c2cb4364d674467de128"></a>

<a id="canonical-22d2654b9a53ef71310263e64d12bc9f914f9e007f79ffcdb9b8cc181676a638"></a>

## name property — proxy_config.https.tls_cert_params.certificates / 22a526e05e5b / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-b08a4ed4b901aa71bd430af88e1112cf5316f291bf75650972ffa71d55f8c6f3"></a>

<a id="canonical-80fee8e69b6a7e467c56aa0b36136261844a434e092bab4b528c4b6d06ecb6eb"></a>

## namespace property — proxy_config.https.tls_cert_params.certificates / 22a526e05e5b / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-24e2b2d2f62c13644cad7ea7a02bfaa062cb3482aa67b697508ec48fbb61e25f"></a>

<a id="canonical-05875cc0248e333f7ddce1339fee951bc0ef42fbf8d96629132ce7f2bbaa3e2d"></a>

## tenant property — proxy_config.https.tls_cert_params.certificates / 22a526e05e5b / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-f5f72373c15ab11057930e587891fe69ff79e07430e0165cfce262b3d1480baa"></a>

## Next pages — proxy_config.https.tls_cert_params.certificates / 22a526e05e5b / 7

- [proxy_config.https.tls_cert_params](data-sources--bigip_http_proxy--reference--group-003.md#canonical-9c4118e7d008c6f588dc371e243ef7160412df5c78621103437520f1bc1cc038)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-1801b48b94d3f41c7b779531d6ae5c7b5795025ba0540f044791aa8310473ac5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-de32989a5cfe49623c3998c9dc23f46d4faddd2988ec86205cb95a4ccb80b8be"></a>

## proxy_config.https.tls_cert_params.no_mtls — proxy_config.https.tls_cert_params.no_mtls / b8971100f1eb / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- [proxy_config.https.tls_cert_params](data-sources--bigip_http_proxy--reference--group-003.md#canonical-9c4118e7d008c6f588dc371e243ef7160412df5c78621103437520f1bc1cc038)
- proxy_config.https.tls_cert_params.no_mtls

<a id="canonical-92cfa3e06a3aa555ea2769be636086ea2361e0631f38094a469a1287b2fa8626"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-dd2cc0142d76c11ce2b8058cfb48e7491682316f58f4dffef6efe8b4656a79c6"></a>

## Direct properties — proxy_config.https.tls_cert_params.no_mtls / b8971100f1eb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a13b2ab2be8466bbc7c81f95d2b2b42e1e601111c2e99097e011ffaec408551a"></a>

## Next pages — proxy_config.https.tls_cert_params.no_mtls / b8971100f1eb / 4

- [proxy_config.https.tls_cert_params](data-sources--bigip_http_proxy--reference--group-003.md#canonical-9c4118e7d008c6f588dc371e243ef7160412df5c78621103437520f1bc1cc038)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-ae2556bad9d17e008ae3b5ddd221b4299262c0dbe1d204357303e644fd15660b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-271def63d4a9fcb0c6bac16965f7ffbe45ccbd8dbe9c571cdfc9b8ebed285db0"></a>

## proxy_config.https.tls_cert_params.tls_config — proxy_config.https.tls_cert_params.tls_config / 037ca2eeaf7b / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- [proxy_config.https.tls_cert_params](data-sources--bigip_http_proxy--reference--group-003.md#canonical-9c4118e7d008c6f588dc371e243ef7160412df5c78621103437520f1bc1cc038)
- proxy_config.https.tls_cert_params.tls_config

<a id="canonical-fdd7c41ef9ab58c73a7b5493d547f7ceca5998a15ffaa87cb4c24cc8a43d36c9"></a>

Type: `"single"`. Computed.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

<a id="canonical-67917d192f05fc08a88b0efb66efcb6210c06309954eae0db4cb95930dbfb1a6"></a>

## Direct properties — proxy_config.https.tls_cert_params.tls_config / 037ca2eeaf7b / 3

- [custom_security](data-sources--bigip_http_proxy--reference--group-003.md#canonical-208178bf67352ee26faeb200a523db947874150d5f61d9af3a3e0d19cddcfee4): complete subsection reference.

- [default_security](data-sources--bigip_http_proxy--reference--group-003.md#canonical-aabee3f656a945d0c07a107ee8a7d7a0e129fab9298f0826d6eae40c2d2b97c6): complete subsection reference.

- [low_security](data-sources--bigip_http_proxy--reference--group-003.md#canonical-62c2951d7398cbe207d4317d0ae913869ba677fd2c5ed037f755449c773c35f1): complete subsection reference.

- [medium_security](data-sources--bigip_http_proxy--reference--group-003.md#canonical-5d9fcbb921a0f12c232cd06e1b7567459545e788f297a365b80e74229fdb6205): complete subsection reference.

<a id="canonical-fda0a7e6573c80827bf04b731c3c50fdb20ff366a98b8cbf5abe80c8f4ffa351"></a>

## Next pages — proxy_config.https.tls_cert_params.tls_config / 037ca2eeaf7b / 4

- [proxy_config.https.tls_cert_params.tls_config.custom_security](data-sources--bigip_http_proxy--reference--group-003.md#canonical-208178bf67352ee26faeb200a523db947874150d5f61d9af3a3e0d19cddcfee4)
- [proxy_config.https.tls_cert_params.tls_config.default_security](data-sources--bigip_http_proxy--reference--group-003.md#canonical-aabee3f656a945d0c07a107ee8a7d7a0e129fab9298f0826d6eae40c2d2b97c6)
- [proxy_config.https.tls_cert_params.tls_config.low_security](data-sources--bigip_http_proxy--reference--group-003.md#canonical-62c2951d7398cbe207d4317d0ae913869ba677fd2c5ed037f755449c773c35f1)
- [proxy_config.https.tls_cert_params.tls_config.medium_security](data-sources--bigip_http_proxy--reference--group-003.md#canonical-5d9fcbb921a0f12c232cd06e1b7567459545e788f297a365b80e74229fdb6205)
- [proxy_config.https.tls_cert_params](data-sources--bigip_http_proxy--reference--group-003.md#canonical-9c4118e7d008c6f588dc371e243ef7160412df5c78621103437520f1bc1cc038)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-208178bf67352ee26faeb200a523db947874150d5f61d9af3a3e0d19cddcfee4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1465e76fb9ebd4a6eadc5a2c26f3a93f65be64f8afcccaca6c2aa4244d9d14e3"></a>

## proxy_config.https.tls_cert_params.tls_config.custom_security — proxy_config.https.tls_cert_params.tls_config.custom_security / a4e0ae6cc2f1 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- [proxy_config.https.tls_cert_params](data-sources--bigip_http_proxy--reference--group-003.md#canonical-9c4118e7d008c6f588dc371e243ef7160412df5c78621103437520f1bc1cc038)
- [proxy_config.https.tls_cert_params.tls_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-ae2556bad9d17e008ae3b5ddd221b4299262c0dbe1d204357303e644fd15660b)
- proxy_config.https.tls_cert_params.tls_config.custom_security

<a id="canonical-22cdf5f4c0084cfe07da48c299aba7699ae129420aa9ebccfad1797f736131da"></a>

Type: `"single"`. Computed.

Defines TLS protocol config including min/max versions and allowed ciphers.

Upstream description:

This defines TLS protocol config including min/max versions and allowed ciphers.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-532f4c50b498157d2d149dab8a298aa986f303470031444ab6667cdf0e1dabbc"></a>

## Direct properties — proxy_config.https.tls_cert_params.tls_config.custom_security / a4e0ae6cc2f1 / 3

<a id="canonical-93eaef405c1541dea463240928b4a414bd15cad0d25f0f0276a5d3e954b18b81"></a>

<a id="canonical-e020c28caf4a0ec946c79c2c8818a80a7b9922d168bd1b298de0af1997f18e7b"></a>

## cipher_suites property — proxy_config.https.tls_cert_params.tls_config.custom_security / a4e0ae6cc2f1 / 4

Type: `["list", "string"]`. Computed.

The TLS listener will only support the specified cipher list.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-6c1b6c9fcfc2b617affecfd8e598c9795d135c150e9f6a158f3f03f5ec5d4ef1"></a>

<a id="canonical-0852bdf3e0536f912b4ffe031449ce178e4fdfd2342f49e6f23a4d9d529737d6"></a>

## max_version property — proxy_config.https.tls_cert_params.tls_config.custom_security / a4e0ae6cc2f1 / 5

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-ef278476fa0becb653fb48a0300aec1216bbf77cb2ff67db4e846c54fe83a2c5"></a>

<a id="canonical-cd689d88f723b74f8a56a1f30123388b72a3a4915322d61b6b84dcf855f6ad77"></a>

## min_version property — proxy_config.https.tls_cert_params.tls_config.custom_security / a4e0ae6cc2f1 / 6

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1df6d7076631381bd9c906e9eef29029793cd1350c1cbf33cf8d28bea74fe080"></a>

## Next pages — proxy_config.https.tls_cert_params.tls_config.custom_security / a4e0ae6cc2f1 / 7

- [proxy_config.https.tls_cert_params.tls_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-ae2556bad9d17e008ae3b5ddd221b4299262c0dbe1d204357303e644fd15660b)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-aabee3f656a945d0c07a107ee8a7d7a0e129fab9298f0826d6eae40c2d2b97c6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-50ee2e48e48096c22c552c83cab8377d16a01df5a5af3c367203c772bc9630df"></a>

## proxy_config.https.tls_cert_params.tls_config.default_security — proxy_config.https.tls_cert_params.tls_config.default_security / cc2fe614d992 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- [proxy_config.https.tls_cert_params](data-sources--bigip_http_proxy--reference--group-003.md#canonical-9c4118e7d008c6f588dc371e243ef7160412df5c78621103437520f1bc1cc038)
- [proxy_config.https.tls_cert_params.tls_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-ae2556bad9d17e008ae3b5ddd221b4299262c0dbe1d204357303e644fd15660b)
- proxy_config.https.tls_cert_params.tls_config.default_security

<a id="canonical-85748ced1bee1f3e10a12fa39991786e38b3f29333d0a4e143eeb1180a50f931"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-4b35ad61856975a945497378d506b5c254bb879e186568879ab3eced5983f4c5"></a>

## Direct properties — proxy_config.https.tls_cert_params.tls_config.default_security / cc2fe614d992 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d40e62ee96de6bb4fcb10d5a2b09ecc23a76eaf0289f4fd34d52322c688a945a"></a>

## Next pages — proxy_config.https.tls_cert_params.tls_config.default_security / cc2fe614d992 / 4

- [proxy_config.https.tls_cert_params.tls_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-ae2556bad9d17e008ae3b5ddd221b4299262c0dbe1d204357303e644fd15660b)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-62c2951d7398cbe207d4317d0ae913869ba677fd2c5ed037f755449c773c35f1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e70482560d3a27f83b026f91ff67ab5cdf755b17fedabfccfa59843fff8cf877"></a>

## proxy_config.https.tls_cert_params.tls_config.low_security — proxy_config.https.tls_cert_params.tls_config.low_security / b7e97dbaf3c0 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- [proxy_config.https.tls_cert_params](data-sources--bigip_http_proxy--reference--group-003.md#canonical-9c4118e7d008c6f588dc371e243ef7160412df5c78621103437520f1bc1cc038)
- [proxy_config.https.tls_cert_params.tls_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-ae2556bad9d17e008ae3b5ddd221b4299262c0dbe1d204357303e644fd15660b)
- proxy_config.https.tls_cert_params.tls_config.low_security

<a id="canonical-1086531268895e0a9c9d2a576c780c2aad56121641ac58def9287e55a720cf71"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-7928781486014fb366866946f4a4feaa15be51ff4bdf6d390b851a54990fa046"></a>

## Direct properties — proxy_config.https.tls_cert_params.tls_config.low_security / b7e97dbaf3c0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ff1638969683b730285100cb930ce6e6907b143789343b0dff3236928307bd58"></a>

## Next pages — proxy_config.https.tls_cert_params.tls_config.low_security / b7e97dbaf3c0 / 4

- [proxy_config.https.tls_cert_params.tls_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-ae2556bad9d17e008ae3b5ddd221b4299262c0dbe1d204357303e644fd15660b)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-5d9fcbb921a0f12c232cd06e1b7567459545e788f297a365b80e74229fdb6205"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a1e4e11b7d5f591423596925905878a968e5d4076b5182500eba7bf068b2aac5"></a>

## proxy_config.https.tls_cert_params.tls_config.medium_security — proxy_config.https.tls_cert_params.tls_config.medium_security / bfb36062027c / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- [proxy_config.https.tls_cert_params](data-sources--bigip_http_proxy--reference--group-003.md#canonical-9c4118e7d008c6f588dc371e243ef7160412df5c78621103437520f1bc1cc038)
- [proxy_config.https.tls_cert_params.tls_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-ae2556bad9d17e008ae3b5ddd221b4299262c0dbe1d204357303e644fd15660b)
- proxy_config.https.tls_cert_params.tls_config.medium_security

<a id="canonical-a48a826a048394e4d09accca26bef4b4a526fcec111578db67cb0a42b05f6b70"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-8c057748bb2ba65b427383e3eefe2ffd237c390b3d7df1d0e2e87e8d234abb5c"></a>

## Direct properties — proxy_config.https.tls_cert_params.tls_config.medium_security / bfb36062027c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3cbd1b38bf0ab26631bc1073bebf28f61538e9bbc0763cf0948db291922aaf1f"></a>

## Next pages — proxy_config.https.tls_cert_params.tls_config.medium_security / bfb36062027c / 4

- [proxy_config.https.tls_cert_params.tls_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-ae2556bad9d17e008ae3b5ddd221b4299262c0dbe1d204357303e644fd15660b)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-7769bad9d339b8f2d09e15b5f21aec04f7c8375af5555e2851e4dd948ea8cf73"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8f356f385a54f49a59f71b1bb51603a6b52f89bdc2da75bcea3dbcb56d924609"></a>

## proxy_config.https.tls_cert_params.use_mtls — proxy_config.https.tls_cert_params.use_mtls / b24d46a12fcf / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- [proxy_config.https.tls_cert_params](data-sources--bigip_http_proxy--reference--group-003.md#canonical-9c4118e7d008c6f588dc371e243ef7160412df5c78621103437520f1bc1cc038)
- proxy_config.https.tls_cert_params.use_mtls

<a id="canonical-fbffd65c11f0b287819064008a0106603475ad221134487763047f31df6108c2"></a>

Type: `"single"`. Computed.

Validation context for downstream client TLS connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-crl_choice": "[\"crl\",\"no_crl\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-xfcc_header": "[\"xfcc_disabled\",\"xfcc_options\"]"
}
```

<a id="canonical-e48d5214c160bec22c0393e2b95a7a734d3015a5a2233d6852555ff7d7591a61"></a>

## Direct properties — proxy_config.https.tls_cert_params.use_mtls / b24d46a12fcf / 3

<a id="canonical-2b2611e990ab9465263f1bc8f04b2b41c805f46fa01d8ce626d22dc2759c3081"></a>

<a id="canonical-fc1018704a4a6489032ec69b0bd6ff152dd2393f3b31542c15e0eed4663c68a7"></a>

## client_certificate_optional property — proxy_config.https.tls_cert_params.use_mtls / b24d46a12fcf / 4

Type: `"bool"`. Computed.

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated.

Upstream description:

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated. If the client
does not provide a certificate, the connection will be accepted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [crl](data-sources--bigip_http_proxy--reference--group-003.md#canonical-b4a772f9188044f2071f0de09017487527a27027a63e68e87d9132c13a004c7c): complete subsection reference.

- [no_crl](data-sources--bigip_http_proxy--reference--group-003.md#canonical-06a9ef51b6939e0a42cd19c2b86cc97ffe2562f3a373e7c550f9e0b7be5407ba): complete subsection reference.

- [trusted_ca](data-sources--bigip_http_proxy--reference--group-003.md#canonical-7b1d6ff3a8e80810774e0902512d3ac11c34f5105ec434a3eae52564b3ea6187): complete subsection reference.

<a id="canonical-d4074ee8dca765f904a93be3a97ceb7d8723f6778f72e2dfa34bba172554aefa"></a>

<a id="canonical-8493d6358c586887bc13397e2e704faff60743b76ac6a9265aaa6d0679db2242"></a>

## trusted_ca_url property — proxy_config.https.tls_cert_params.use_mtls / b24d46a12fcf / 5

Type: `"string"`. Computed.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Upstream description:

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

- [xfcc_disabled](data-sources--bigip_http_proxy--reference--group-003.md#canonical-5979cad9538974b7b62f997bdcf47a8b9f5490e25016127660506041133e5744): complete subsection reference.

- [xfcc_options](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2858b3d235d0df46698fbb4ca2b340782ea2a603ef613d03f2c4b5cb23b96664): complete subsection reference.

<a id="canonical-74b30899040fce264b245af08aafb3a9b8fdad48f6728d9cbd86ae76b4b7e884"></a>

## Next pages — proxy_config.https.tls_cert_params.use_mtls / b24d46a12fcf / 6

- [proxy_config.https.tls_cert_params.use_mtls.crl](data-sources--bigip_http_proxy--reference--group-003.md#canonical-b4a772f9188044f2071f0de09017487527a27027a63e68e87d9132c13a004c7c)
- [proxy_config.https.tls_cert_params.use_mtls.no_crl](data-sources--bigip_http_proxy--reference--group-003.md#canonical-06a9ef51b6939e0a42cd19c2b86cc97ffe2562f3a373e7c550f9e0b7be5407ba)
- [proxy_config.https.tls_cert_params.use_mtls.trusted_ca](data-sources--bigip_http_proxy--reference--group-003.md#canonical-7b1d6ff3a8e80810774e0902512d3ac11c34f5105ec434a3eae52564b3ea6187)
- [proxy_config.https.tls_cert_params.use_mtls.xfcc_disabled](data-sources--bigip_http_proxy--reference--group-003.md#canonical-5979cad9538974b7b62f997bdcf47a8b9f5490e25016127660506041133e5744)
- [proxy_config.https.tls_cert_params.use_mtls.xfcc_options](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2858b3d235d0df46698fbb4ca2b340782ea2a603ef613d03f2c4b5cb23b96664)
- [proxy_config.https.tls_cert_params](data-sources--bigip_http_proxy--reference--group-003.md#canonical-9c4118e7d008c6f588dc371e243ef7160412df5c78621103437520f1bc1cc038)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-b4a772f9188044f2071f0de09017487527a27027a63e68e87d9132c13a004c7c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-66ca292fc55cd54d28d278cf34cb01ae6bc6a8b25b3d72ffc464a54e726f1267"></a>

## proxy_config.https.tls_cert_params.use_mtls.crl — proxy_config.https.tls_cert_params.use_mtls.crl / 39fdb6e2b790 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- [proxy_config.https.tls_cert_params](data-sources--bigip_http_proxy--reference--group-003.md#canonical-9c4118e7d008c6f588dc371e243ef7160412df5c78621103437520f1bc1cc038)
- [proxy_config.https.tls_cert_params.use_mtls](data-sources--bigip_http_proxy--reference--group-003.md#canonical-7769bad9d339b8f2d09e15b5f21aec04f7c8375af5555e2851e4dd948ea8cf73)
- proxy_config.https.tls_cert_params.use_mtls.crl

<a id="canonical-87b1cb7dd59e861a77f18359da5f94e5ce9fe992a51571a785f0b26c6dd53e47"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-a8913fa43f43eef7ae5b4cbbd3e3342ad77d215bfc3ab2fe16c922470fb3a9ba"></a>

## Direct properties — proxy_config.https.tls_cert_params.use_mtls.crl / 39fdb6e2b790 / 3

<a id="canonical-40a0880643acbe78132bed6a8e8ae243ea32715b08227613d40bc324720b500b"></a>

<a id="canonical-5250825754bf9a6769c9d6adb4eba5368d44e122a5ef9c1b3d3214ea78b0a5d7"></a>

## name property — proxy_config.https.tls_cert_params.use_mtls.crl / 39fdb6e2b790 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-8d41cc7df037d64fa0a04759daeea6d873864fa7d98a13a1b083130f09a76441"></a>

<a id="canonical-07a27257d7df20e8c7c9f87b28a015f54c7bfc7207357cdbe42ff78cefeff889"></a>

## namespace property — proxy_config.https.tls_cert_params.use_mtls.crl / 39fdb6e2b790 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-f39fb51beeb1c29f3a0c452367f37b88991d0e36f76b19534d0ddd6e43427a01"></a>

<a id="canonical-ef84dacb43147ba779721fed677343200f33823a280d7b116cd0604eef286b88"></a>

## tenant property — proxy_config.https.tls_cert_params.use_mtls.crl / 39fdb6e2b790 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-476df77ad71aad691260ad86df8677e95b2e274ecb804b996a1b3129bd66a66f"></a>

## Next pages — proxy_config.https.tls_cert_params.use_mtls.crl / 39fdb6e2b790 / 7

- [proxy_config.https.tls_cert_params.use_mtls](data-sources--bigip_http_proxy--reference--group-003.md#canonical-7769bad9d339b8f2d09e15b5f21aec04f7c8375af5555e2851e4dd948ea8cf73)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-06a9ef51b6939e0a42cd19c2b86cc97ffe2562f3a373e7c550f9e0b7be5407ba"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7db961b40b80eaeca292958da49c9d88e7c6039146d0dfe7b9fc0b19b04de307"></a>

## proxy_config.https.tls_cert_params.use_mtls.no_crl — proxy_config.https.tls_cert_params.use_mtls.no_crl / 35e09bb31dce / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- [proxy_config.https.tls_cert_params](data-sources--bigip_http_proxy--reference--group-003.md#canonical-9c4118e7d008c6f588dc371e243ef7160412df5c78621103437520f1bc1cc038)
- [proxy_config.https.tls_cert_params.use_mtls](data-sources--bigip_http_proxy--reference--group-003.md#canonical-7769bad9d339b8f2d09e15b5f21aec04f7c8375af5555e2851e4dd948ea8cf73)
- proxy_config.https.tls_cert_params.use_mtls.no_crl

<a id="canonical-e8b1f086cb7beab7e8412f07b062bbb2087b673d47bfada377c46a8d89a051cd"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2e22fa1ef8a62b0362f6e819892c80e7173c37ec2425a12a9006b06b9ef9d5b3"></a>

## Direct properties — proxy_config.https.tls_cert_params.use_mtls.no_crl / 35e09bb31dce / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bf7c770e7da9e5f8b61cd437a249b76b5cafb4a253af601c30b6ad33016ccfa2"></a>

## Next pages — proxy_config.https.tls_cert_params.use_mtls.no_crl / 35e09bb31dce / 4

- [proxy_config.https.tls_cert_params.use_mtls](data-sources--bigip_http_proxy--reference--group-003.md#canonical-7769bad9d339b8f2d09e15b5f21aec04f7c8375af5555e2851e4dd948ea8cf73)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-7b1d6ff3a8e80810774e0902512d3ac11c34f5105ec434a3eae52564b3ea6187"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e6d9da8be0912684626e3e64709a1c139d5e87f53d408ea886073e35c397dff5"></a>

## proxy_config.https.tls_cert_params.use_mtls.trusted_ca — proxy_config.https.tls_cert_params.use_mtls.trusted_ca / ddb06af1539b / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- [proxy_config.https.tls_cert_params](data-sources--bigip_http_proxy--reference--group-003.md#canonical-9c4118e7d008c6f588dc371e243ef7160412df5c78621103437520f1bc1cc038)
- [proxy_config.https.tls_cert_params.use_mtls](data-sources--bigip_http_proxy--reference--group-003.md#canonical-7769bad9d339b8f2d09e15b5f21aec04f7c8375af5555e2851e4dd948ea8cf73)
- proxy_config.https.tls_cert_params.use_mtls.trusted_ca

<a id="canonical-78e86b53479cc7ba11186fa9d23584b10f0d9c3649cb1f7bb5caee09b81621ed"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-bb045e77d9a093063fbacb74f34345a3899f3ecec5bbf29e8d6855cc2eb3dee8"></a>

## Direct properties — proxy_config.https.tls_cert_params.use_mtls.trusted_ca / ddb06af1539b / 3

<a id="canonical-f9d926d359d22daece2ca478f04f62a5d6ede7790b813592d95cccf95846cca9"></a>

<a id="canonical-0d0aebcc348f20bffd88064673bdbbb458c3f89659db10079bcc3452a99a00cc"></a>

## name property — proxy_config.https.tls_cert_params.use_mtls.trusted_ca / ddb06af1539b / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-1f86d16f54fcf2f6295371cb0384538956b8459b76a62d17e72274023d84c45e"></a>

<a id="canonical-a392b4b7f7f509adac4722e46350d291e23bd962fbcc20d096e97a3676c1e523"></a>

## namespace property — proxy_config.https.tls_cert_params.use_mtls.trusted_ca / ddb06af1539b / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-54316ad634518660fc3110085cd67c6327ee70475d99959b4eb64e324b56eece"></a>

<a id="canonical-8c7b57c3d927031e65a9660cdcab511b923f55fd1cd86534d438f990aa8f0c7f"></a>

## tenant property — proxy_config.https.tls_cert_params.use_mtls.trusted_ca / ddb06af1539b / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-8205b60378b47b990682895f1a56e9c598f88b2127aefd6984fd448038204a81"></a>

## Next pages — proxy_config.https.tls_cert_params.use_mtls.trusted_ca / ddb06af1539b / 7

- [proxy_config.https.tls_cert_params.use_mtls](data-sources--bigip_http_proxy--reference--group-003.md#canonical-7769bad9d339b8f2d09e15b5f21aec04f7c8375af5555e2851e4dd948ea8cf73)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-5979cad9538974b7b62f997bdcf47a8b9f5490e25016127660506041133e5744"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7b11b801fd4eb953c03e34aeda4116c9e2018190ac9714c48a16c1f966c37f37"></a>

## proxy_config.https.tls_cert_params.use_mtls.xfcc_disabled — proxy_config.https.tls_cert_params.use_mtls.xfcc_disabled / 26efcc918b44 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- [proxy_config.https.tls_cert_params](data-sources--bigip_http_proxy--reference--group-003.md#canonical-9c4118e7d008c6f588dc371e243ef7160412df5c78621103437520f1bc1cc038)
- [proxy_config.https.tls_cert_params.use_mtls](data-sources--bigip_http_proxy--reference--group-003.md#canonical-7769bad9d339b8f2d09e15b5f21aec04f7c8375af5555e2851e4dd948ea8cf73)
- proxy_config.https.tls_cert_params.use_mtls.xfcc_disabled

<a id="canonical-bc5e959778ef8730f77b73b2f60ff901a8072c85c329a9bf997da73e2484138d"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-631f164660c32173325544dc2dd50ce6fd1ba3233c7d4780ea93da7e82f1384e"></a>

## Direct properties — proxy_config.https.tls_cert_params.use_mtls.xfcc_disabled / 26efcc918b44 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2cd69e8cc3e5567b565255a5a3ec8dff25cbee5863bf56437670b307c3a0a405"></a>

## Next pages — proxy_config.https.tls_cert_params.use_mtls.xfcc_disabled / 26efcc918b44 / 4

- [proxy_config.https.tls_cert_params.use_mtls](data-sources--bigip_http_proxy--reference--group-003.md#canonical-7769bad9d339b8f2d09e15b5f21aec04f7c8375af5555e2851e4dd948ea8cf73)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-2858b3d235d0df46698fbb4ca2b340782ea2a603ef613d03f2c4b5cb23b96664"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4df726727fb2e6c05663d9c1816998efd97891a6bdf3deef6d7a29a6bd1aca2a"></a>

## proxy_config.https.tls_cert_params.use_mtls.xfcc_options — proxy_config.https.tls_cert_params.use_mtls.xfcc_options / 346f17c0ff5f / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- [proxy_config.https.tls_cert_params](data-sources--bigip_http_proxy--reference--group-003.md#canonical-9c4118e7d008c6f588dc371e243ef7160412df5c78621103437520f1bc1cc038)
- [proxy_config.https.tls_cert_params.use_mtls](data-sources--bigip_http_proxy--reference--group-003.md#canonical-7769bad9d339b8f2d09e15b5f21aec04f7c8375af5555e2851e4dd948ea8cf73)
- proxy_config.https.tls_cert_params.use_mtls.xfcc_options

<a id="canonical-3f046c1ebf3638a0858c7d9078d0845ce11fc433356a76d7d1faf81babcb456a"></a>

Type: `"single"`. Computed.

X-Forwarded-Client-Cert header elements to be added to requests.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1f29aab2bf2fe7ebdab708d1c53c87159488cd25932c398ed94e2e8a02299618"></a>

## Direct properties — proxy_config.https.tls_cert_params.use_mtls.xfcc_options / 346f17c0ff5f / 3

<a id="canonical-1de6bccd1f65d658ace2452fa2599725f4ed1a192d67564120c5ffc81ad38b93"></a>

<a id="canonical-8cbd30964e4eef3e0d524c24cc0ab0e8c7999385d9433afbe127178d798be664"></a>

## xfcc_header_elements property — proxy_config.https.tls_cert_params.use_mtls.xfcc_options / 346f17c0ff5f / 4

Type: `["list", "string"]`. Computed.

\[Enum: XFCC\_NONE|XFCC\_CERT|XFCC\_CHAIN|XFCC\_SUBJECT|XFCC\_URI|XFCC\_DNS\]
X-Forwarded-Client-Cert header elements to be added to requests. Possible values are \`XFCC\_NONE\`,
\`XFCC\_CERT\`, \`XFCC\_CHAIN\`, \`XFCC\_SUBJECT\`, \`XFCC\_URI\`, \`XFCC\_DNS\`. Defaults to
\`XFCC\_NONE\`.

Upstream description:

X-Forwarded-Client-Cert header elements to be added to requests.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  }
}
```

<a id="canonical-5dd3aba3154596d8d5fcafc793278d6a05d04e678837923c4fa4d7bb25c977e5"></a>

## Next pages — proxy_config.https.tls_cert_params.use_mtls.xfcc_options / 346f17c0ff5f / 5

- [proxy_config.https.tls_cert_params.use_mtls](data-sources--bigip_http_proxy--reference--group-003.md#canonical-7769bad9d339b8f2d09e15b5f21aec04f7c8375af5555e2851e4dd948ea8cf73)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-12b1668564b57fed383102e0655e5c2fd223f895bd3f947a1500d20e88bd4b90"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4548ad2982fb5962f501c5ee75aa171684ebc032b7e5d79ec16a2b26bd985c33"></a>

## proxy_config.https.tls_parameters — proxy_config.https.tls_parameters / a05f3e670ce6 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- proxy_config.https.tls_parameters

<a id="canonical-616c3576f8ddcbe0b9f675d0038e8d7a78d45d2b4f1a7e956eb5e69f0c969554"></a>

Type: `"single"`. Computed.

Configuration parameter for tls parameters.

Upstream description:

Inline TLS parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

<a id="canonical-92d96e0fc355bc09a7801d90badbac91747861e24551ee98c675731a0accfefa"></a>

## Direct properties — proxy_config.https.tls_parameters / a05f3e670ce6 / 3

- [no_mtls](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d395d0bd821c1a0ab8d2111c440ebc28d34079e2c7e34203414b6bdeee8e4782): complete subsection reference.

- [tls_certificates](data-sources--bigip_http_proxy--reference--group-003.md#canonical-ecf9c85358445cb80e77e18ae57762f3136159ff52ec7e36742fdd88bf7c5a6d): complete subsection reference.

- [tls_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-f85d66899e619671943d5da274f98e215365a6a9cd9991d9e0733009dd9f6485): complete subsection reference.

- [use_mtls](data-sources--bigip_http_proxy--reference--group-003.md#canonical-86372ca7fbc2e20f711797628e4eeb717449261b6371d6fc347c5aba93edf190): complete subsection reference.

<a id="canonical-5e8bccf7ab27230be8a0386bdac537f37933716b3eaccfe58f8c706b40e15d12"></a>

## Next pages — proxy_config.https.tls_parameters / a05f3e670ce6 / 4

- [proxy_config.https.tls_parameters.no_mtls](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d395d0bd821c1a0ab8d2111c440ebc28d34079e2c7e34203414b6bdeee8e4782)
- [proxy_config.https.tls_parameters.tls_certificates](data-sources--bigip_http_proxy--reference--group-003.md#canonical-ecf9c85358445cb80e77e18ae57762f3136159ff52ec7e36742fdd88bf7c5a6d)
- [proxy_config.https.tls_parameters.tls_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-f85d66899e619671943d5da274f98e215365a6a9cd9991d9e0733009dd9f6485)
- [proxy_config.https.tls_parameters.use_mtls](data-sources--bigip_http_proxy--reference--group-003.md#canonical-86372ca7fbc2e20f711797628e4eeb717449261b6371d6fc347c5aba93edf190)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-d395d0bd821c1a0ab8d2111c440ebc28d34079e2c7e34203414b6bdeee8e4782"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ce0b3d1a2890f83c24643d7bdd6752c4fabf06eaa11f1b4a3ba2f15f6f992c91"></a>

## proxy_config.https.tls_parameters.no_mtls — proxy_config.https.tls_parameters.no_mtls / 0169eb5c2e1d / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-003.md#canonical-12b1668564b57fed383102e0655e5c2fd223f895bd3f947a1500d20e88bd4b90)
- proxy_config.https.tls_parameters.no_mtls

<a id="canonical-501cfa27afd123dc948b98a32678c2f0c2905ee94c98b1c06f1049eeca513649"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-7f1b8401c4e63ee7cffc6c4e646e4cbeb794eafe687214cdaffdddfd85ec29f4"></a>

## Direct properties — proxy_config.https.tls_parameters.no_mtls / 0169eb5c2e1d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-70fb66fb77309b4a9cb5603d6da0a4090fe3c94d67ba9ae07f9c6f1876567651"></a>

## Next pages — proxy_config.https.tls_parameters.no_mtls / 0169eb5c2e1d / 4

- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-003.md#canonical-12b1668564b57fed383102e0655e5c2fd223f895bd3f947a1500d20e88bd4b90)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-ecf9c85358445cb80e77e18ae57762f3136159ff52ec7e36742fdd88bf7c5a6d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-22263046636306b6f2bede793149cf55d7969e9ce7dac51314e6f824defd1033"></a>

## proxy_config.https.tls_parameters.tls_certificates — proxy_config.https.tls_parameters.tls_certificates / dfba273b9db8 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-003.md#canonical-12b1668564b57fed383102e0655e5c2fd223f895bd3f947a1500d20e88bd4b90)
- proxy_config.https.tls_parameters.tls_certificates

<a id="canonical-f8ed05f6a8d9c1731b317ee50da934274e9dc7746aab3209e98db6870df473d3"></a>

Type: `"list"`. Computed.

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

Upstream description:

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-d475edd72acf43577fcb77053e2dc9e52efc1a23fb0e6f35d3137cec09b82bac"></a>

## Direct properties — proxy_config.https.tls_parameters.tls_certificates / dfba273b9db8 / 3

<a id="canonical-3333b7597f01e01855df0ce34f29fdcaf09bf17f3b703da54b8949af1f88101e"></a>

<a id="canonical-c9e3038115d088cf7934317a06fdc45034a17802545a09662c0a7661417c9f67"></a>

## certificate_url property — proxy_config.https.tls_parameters.tls_certificates / dfba273b9db8 / 4

Type: `"string"`. Computed.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Upstream description:

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [custom_hash_algorithms](data-sources--bigip_http_proxy--reference--group-003.md#canonical-858f4b9096e837eceaddb15427bb3ba7c451624d96bc4c9daa24c1316d56cfd2): complete subsection reference.

<a id="canonical-10b04008c44f768cddb228fc5d481ce77824a7d3ae64732e1e99db66900feef0"></a>

<a id="canonical-7ee1d335f9c297c084a22cedf5c88850fcd5f1d88b02e815e362cfaf36b65d6c"></a>

## description_spec property — proxy_config.https.tls_parameters.tls_certificates / dfba273b9db8 / 5

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](data-sources--bigip_http_proxy--reference--group-003.md#canonical-7c148f8d682655c10ee62a5173dff8f06f8b2d37a26cc52a2261b73cecd1a900): complete subsection reference.

- [private_key](data-sources--bigip_http_proxy--reference--group-003.md#canonical-00a1e0956e529103c7fc45661d2e029b7ad9403715c25747f601e92c8b10856e): complete subsection reference.

- [use_system_defaults](data-sources--bigip_http_proxy--reference--group-003.md#canonical-de34422009cdc27964f3478475f61c3795911b8be75ea959f3b738317a692381): complete subsection reference.

<a id="canonical-5b2d0c94e56c1f586402db66e33a0a2b236588d67a0f37721aec9f7a19937eb7"></a>

## Next pages — proxy_config.https.tls_parameters.tls_certificates / dfba273b9db8 / 6

- [proxy_config.https.tls_parameters.tls_certificates.custom_hash_algorithms](data-sources--bigip_http_proxy--reference--group-003.md#canonical-858f4b9096e837eceaddb15427bb3ba7c451624d96bc4c9daa24c1316d56cfd2)
- [proxy_config.https.tls_parameters.tls_certificates.disable_ocsp_stapling](data-sources--bigip_http_proxy--reference--group-003.md#canonical-7c148f8d682655c10ee62a5173dff8f06f8b2d37a26cc52a2261b73cecd1a900)
- [proxy_config.https.tls_parameters.tls_certificates.private_key](data-sources--bigip_http_proxy--reference--group-003.md#canonical-00a1e0956e529103c7fc45661d2e029b7ad9403715c25747f601e92c8b10856e)
- [proxy_config.https.tls_parameters.tls_certificates.use_system_defaults](data-sources--bigip_http_proxy--reference--group-003.md#canonical-de34422009cdc27964f3478475f61c3795911b8be75ea959f3b738317a692381)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-003.md#canonical-12b1668564b57fed383102e0655e5c2fd223f895bd3f947a1500d20e88bd4b90)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-858f4b9096e837eceaddb15427bb3ba7c451624d96bc4c9daa24c1316d56cfd2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d388733c71b4759c7059e89ab6a021be427a12848abb401172deab0ab8db4bf8"></a>

## proxy_config.https.tls_parameters.tls_certificates.custom_hash_algorithms — proxy_config.https.tls_parameters.tls_certificates.custom_hash_algorithms / f0ddfc31ef8b / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-003.md#canonical-12b1668564b57fed383102e0655e5c2fd223f895bd3f947a1500d20e88bd4b90)
- [proxy_config.https.tls_parameters.tls_certificates](data-sources--bigip_http_proxy--reference--group-003.md#canonical-ecf9c85358445cb80e77e18ae57762f3136159ff52ec7e36742fdd88bf7c5a6d)
- proxy_config.https.tls_parameters.tls_certificates.custom_hash_algorithms

<a id="canonical-59159874ca0574997727ab420b1cd5665a1726c55dc9cbceaad58295870cb466"></a>

Type: `"single"`. Computed.

Specifies the hash algorithms to be used.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2096c8d627e1c87f4d6d8dad01ea401b3ba98b8124542f5f8538c2df51ec9003"></a>

## Direct properties — proxy_config.https.tls_parameters.tls_certificates.custom_hash_algorithms / f0ddfc31ef8b / 3

<a id="canonical-a02167508df3c2834ccc7bc86abcf344aecef6445ea121d7cbcd6f0bad17c6f9"></a>

<a id="canonical-8581585879be037a67cca050b5b2c14efc91ac37f7832b437d91766025dd6cd8"></a>

## hash_algorithms property — proxy_config.https.tls_parameters.tls_certificates.custom_hash_algorithms / f0ddfc31ef8b / 4

Type: `["list", "string"]`. Computed.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

Upstream description:

Ordered list of hash algorithms to be used.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-6cc8908467b2b9f973e271aace30f5a30ec2b15d3acc92c4ccf64064d40a40b9"></a>

## Next pages — proxy_config.https.tls_parameters.tls_certificates.custom_hash_algorithms / f0ddfc31ef8b / 5

- [proxy_config.https.tls_parameters.tls_certificates](data-sources--bigip_http_proxy--reference--group-003.md#canonical-ecf9c85358445cb80e77e18ae57762f3136159ff52ec7e36742fdd88bf7c5a6d)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-7c148f8d682655c10ee62a5173dff8f06f8b2d37a26cc52a2261b73cecd1a900"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-80520a9362b514af58fc2a157b1ee5a8f0c87db33a445d3f2d640f132bce8756"></a>

## proxy_config.https.tls_parameters.tls_certificates.disable_ocsp_stapling — proxy_config.https.tls_parameters.tls_certificates.disable_ocsp_stapling / c03204307239 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-003.md#canonical-12b1668564b57fed383102e0655e5c2fd223f895bd3f947a1500d20e88bd4b90)
- [proxy_config.https.tls_parameters.tls_certificates](data-sources--bigip_http_proxy--reference--group-003.md#canonical-ecf9c85358445cb80e77e18ae57762f3136159ff52ec7e36742fdd88bf7c5a6d)
- proxy_config.https.tls_parameters.tls_certificates.disable_ocsp_stapling

<a id="canonical-d615d3b515d9481ea20f9dd59191dbe3673f703f9e3851305d789c2c57b3efa1"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable ocsp stapling.

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-f8c756e69f66ed9da3bd11e50c038464de484c7362768f3cf59fc1814d337635"></a>

## Direct properties — proxy_config.https.tls_parameters.tls_certificates.disable_ocsp_stapling / c03204307239 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6c95f26e7ff0eefde3621f704da6a5dc4812c3de4b8cbc197004ce20cc433931"></a>

## Next pages — proxy_config.https.tls_parameters.tls_certificates.disable_ocsp_stapling / c03204307239 / 4

- [proxy_config.https.tls_parameters.tls_certificates](data-sources--bigip_http_proxy--reference--group-003.md#canonical-ecf9c85358445cb80e77e18ae57762f3136159ff52ec7e36742fdd88bf7c5a6d)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-00a1e0956e529103c7fc45661d2e029b7ad9403715c25747f601e92c8b10856e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-92832503ad3c3b30b6f8a0f15cb83e3ede76bbd2db838457d7d087b4ef9c15b8"></a>

## proxy_config.https.tls_parameters.tls_certificates.private_key — proxy_config.https.tls_parameters.tls_certificates.private_key / c5a4d87574a5 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-003.md#canonical-12b1668564b57fed383102e0655e5c2fd223f895bd3f947a1500d20e88bd4b90)
- [proxy_config.https.tls_parameters.tls_certificates](data-sources--bigip_http_proxy--reference--group-003.md#canonical-ecf9c85358445cb80e77e18ae57762f3136159ff52ec7e36742fdd88bf7c5a6d)
- proxy_config.https.tls_parameters.tls_certificates.private_key

<a id="canonical-0eafb96d97e189c015b2728668122db5376f55c8c72536cf55a7037bee33dbf1"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

<a id="canonical-5b1c3b76c53a63a9b14403cde6c2fac93f5b74f19f287e9854a10772804779c9"></a>

## Direct properties — proxy_config.https.tls_parameters.tls_certificates.private_key / c5a4d87574a5 / 3

- [blindfold_secret_info](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1d078bbbfcd4cf5f8db338d396907afeb0f3a3859487f5c9dde3a8e844a9e22b): complete subsection reference.

- [clear_secret_info](data-sources--bigip_http_proxy--reference--group-003.md#canonical-76fded2013924c0c6af1adab1e1a053b5d496ac39c84f68c58f67cef77afccb5): complete subsection reference.

<a id="canonical-2721053ebde4f57a9655588216676293135309d7f8e1673cd4aae46c95eb9f72"></a>

## Next pages — proxy_config.https.tls_parameters.tls_certificates.private_key / c5a4d87574a5 / 4

- [proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1d078bbbfcd4cf5f8db338d396907afeb0f3a3859487f5c9dde3a8e844a9e22b)
- [proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info](data-sources--bigip_http_proxy--reference--group-003.md#canonical-76fded2013924c0c6af1adab1e1a053b5d496ac39c84f68c58f67cef77afccb5)
- [proxy_config.https.tls_parameters.tls_certificates](data-sources--bigip_http_proxy--reference--group-003.md#canonical-ecf9c85358445cb80e77e18ae57762f3136159ff52ec7e36742fdd88bf7c5a6d)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-1d078bbbfcd4cf5f8db338d396907afeb0f3a3859487f5c9dde3a8e844a9e22b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-00d21ba6eb480a68cbc692aa01b026bd3d112902aec524ec284e6bdc12acdb4a"></a>

## proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info — proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_ / f3cb48bac04f / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-003.md#canonical-12b1668564b57fed383102e0655e5c2fd223f895bd3f947a1500d20e88bd4b90)
- [proxy_config.https.tls_parameters.tls_certificates](data-sources--bigip_http_proxy--reference--group-003.md#canonical-ecf9c85358445cb80e77e18ae57762f3136159ff52ec7e36742fdd88bf7c5a6d)
- [proxy_config.https.tls_parameters.tls_certificates.private_key](data-sources--bigip_http_proxy--reference--group-003.md#canonical-00a1e0956e529103c7fc45661d2e029b7ad9403715c25747f601e92c8b10856e)
- proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-a02f905df5a36a4855936b53857a2a35d89d76b788d84465aba0a27d70d82728"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-dafbc22e3a38ffbc10e056576df440501f1eb8590cfd834ad34ff90ebca66466"></a>

## Direct properties — proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_ / f3cb48bac04f / 3

<a id="canonical-e390e67c2c6f60b4dee55afd9e8f1feb00e9206ebf43755d94e1cdbd2cf53646"></a>

<a id="canonical-759eae77d8879fefa841182d6f51a4abf42c8388c39313fcccbce6712beb14d2"></a>

## decryption_provider property — proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_ / f3cb48bac04f / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

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
  }
}
```

<a id="canonical-df8f08a784509358b45532877bf0e53fa614260bcefd75447daf95c5507bcb05"></a>

<a id="canonical-cd956387fd39fb1e72da0622964ccff2c2fb6b525886cbbbebeececb3ea18746"></a>

## location property — proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_ / f3cb48bac04f / 5

Type: `"string"`. Computed, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-033f72565abea0539b2cd9b4b665c992166bb42cd024f4468af6c618bbe16c4a"></a>

<a id="canonical-6ce10cf32380914f15a047c6813d336ebaf5acef8d85e6c2334d6aaf7d6ac2ef"></a>

## store_provider property — proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_ / f3cb48bac04f / 6

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

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
  }
}
```

<a id="canonical-c3c57574d85574f3b28db93f621459500b9b1b53e11f7370acfd00f13d3adb95"></a>

## Next pages — proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_ / f3cb48bac04f / 7

- [proxy_config.https.tls_parameters.tls_certificates.private_key](data-sources--bigip_http_proxy--reference--group-003.md#canonical-00a1e0956e529103c7fc45661d2e029b7ad9403715c25747f601e92c8b10856e)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-76fded2013924c0c6af1adab1e1a053b5d496ac39c84f68c58f67cef77afccb5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-44ff3d1cd17378e73c9d0c31ba358f6228a5b9df0ae009e0fbf3fc51132cc277"></a>

## proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info — proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info / 968884c233db / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-003.md#canonical-12b1668564b57fed383102e0655e5c2fd223f895bd3f947a1500d20e88bd4b90)
- [proxy_config.https.tls_parameters.tls_certificates](data-sources--bigip_http_proxy--reference--group-003.md#canonical-ecf9c85358445cb80e77e18ae57762f3136159ff52ec7e36742fdd88bf7c5a6d)
- [proxy_config.https.tls_parameters.tls_certificates.private_key](data-sources--bigip_http_proxy--reference--group-003.md#canonical-00a1e0956e529103c7fc45661d2e029b7ad9403715c25747f601e92c8b10856e)
- proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info

<a id="canonical-a32c34a570f339a45a0c285cd9021b625c1b3e4337c5c974ae712be4cb8f22da"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1dc1df4e82b4b6cbc56e6a3234862053bc4ae02a9f7d605d305b6f4908ae714c"></a>

## Direct properties — proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info / 968884c233db / 3

<a id="canonical-260eb859e7a8a9efd1bcbd41ec9cf71eb095220ddc02dcdc618bd6ed5d7f33c6"></a>

<a id="canonical-dc1bc1cd500517f24dd77d69669165dc327289bf081892b8f1a235356deb94ea"></a>

## provider_ref property — proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info / 968884c233db / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-9fb601a012444488e9320052bb5130184440d4423e1037096814a7f40880bda4"></a>

<a id="canonical-86a2774894afd041970ed49d16afc89846e12a31fdfaea6875d4913aa5e81d17"></a>

## url property — proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info / 968884c233db / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-340b5633c557e5adf5c54745c7afd38b1a2e476fc9f11a8fdfae7dff0a7f5e14"></a>

## Next pages — proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info / 968884c233db / 6

- [proxy_config.https.tls_parameters.tls_certificates.private_key](data-sources--bigip_http_proxy--reference--group-003.md#canonical-00a1e0956e529103c7fc45661d2e029b7ad9403715c25747f601e92c8b10856e)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-de34422009cdc27964f3478475f61c3795911b8be75ea959f3b738317a692381"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0025435fa41fbe99f6ce81061c679ef1e8df039175cec061be7482e692fb3a2f"></a>

## proxy_config.https.tls_parameters.tls_certificates.use_system_defaults — proxy_config.https.tls_parameters.tls_certificates.use_system_defaults / c1aab8be806a / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-003.md#canonical-12b1668564b57fed383102e0655e5c2fd223f895bd3f947a1500d20e88bd4b90)
- [proxy_config.https.tls_parameters.tls_certificates](data-sources--bigip_http_proxy--reference--group-003.md#canonical-ecf9c85358445cb80e77e18ae57762f3136159ff52ec7e36742fdd88bf7c5a6d)
- proxy_config.https.tls_parameters.tls_certificates.use_system_defaults

<a id="canonical-9d186f08a180eb5faa48197d1d092092cf2767ee870fd050cf038463cc0f6a4a"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for use system defaults.

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1ee54eccc5111be8248cbe32553fd093fe4c6c05332314aff441a27564658f56"></a>

## Direct properties — proxy_config.https.tls_parameters.tls_certificates.use_system_defaults / c1aab8be806a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-64bf07548cc704d2ebe3e45f033c1f048bd3fa561ffd1db783e3785f8cd3fb63"></a>

## Next pages — proxy_config.https.tls_parameters.tls_certificates.use_system_defaults / c1aab8be806a / 4

- [proxy_config.https.tls_parameters.tls_certificates](data-sources--bigip_http_proxy--reference--group-003.md#canonical-ecf9c85358445cb80e77e18ae57762f3136159ff52ec7e36742fdd88bf7c5a6d)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-f85d66899e619671943d5da274f98e215365a6a9cd9991d9e0733009dd9f6485"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-521e378cf64e089bb5db1c8ea5e119b2b3d9b817291e97e4d2f416740e8973a4"></a>

## proxy_config.https.tls_parameters.tls_config — proxy_config.https.tls_parameters.tls_config / c1112765b0d7 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-003.md#canonical-12b1668564b57fed383102e0655e5c2fd223f895bd3f947a1500d20e88bd4b90)
- proxy_config.https.tls_parameters.tls_config

<a id="canonical-99b1a5e02623cd1d34fc531e98203bd0a768414cd8cda669cb2dbb066f9427eb"></a>

Type: `"single"`. Computed.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

<a id="canonical-269321008600fc9b376542516871aa24adbd527a08723d613302bec1502f3a14"></a>

## Direct properties — proxy_config.https.tls_parameters.tls_config / c1112765b0d7 / 3

- [custom_security](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3c60f6da9694218fb3051479cd4ae6839e1d1e2c4e673b23b837ae1e2e175268): complete subsection reference.

- [default_security](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3ed63bd061fdac7cc226c995f36938f661a1100ea20bc57de9c83df695259d1c): complete subsection reference.

- [low_security](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3424d97af32411d3972e52f2b950d15f862787bbb9b921d012827fe960d7956d): complete subsection reference.

- [medium_security](data-sources--bigip_http_proxy--reference--group-003.md#canonical-b530df3792dc26c5f21a992373bbdf02920e2c3f93d4f4248ef21840e3a9a4d8): complete subsection reference.

<a id="canonical-e0979873a591007f080f4869b60b20a6e32c4ad509279e893fdb59f8077cb191"></a>

## Next pages — proxy_config.https.tls_parameters.tls_config / c1112765b0d7 / 4

- [proxy_config.https.tls_parameters.tls_config.custom_security](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3c60f6da9694218fb3051479cd4ae6839e1d1e2c4e673b23b837ae1e2e175268)
- [proxy_config.https.tls_parameters.tls_config.default_security](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3ed63bd061fdac7cc226c995f36938f661a1100ea20bc57de9c83df695259d1c)
- [proxy_config.https.tls_parameters.tls_config.low_security](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3424d97af32411d3972e52f2b950d15f862787bbb9b921d012827fe960d7956d)
- [proxy_config.https.tls_parameters.tls_config.medium_security](data-sources--bigip_http_proxy--reference--group-003.md#canonical-b530df3792dc26c5f21a992373bbdf02920e2c3f93d4f4248ef21840e3a9a4d8)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-003.md#canonical-12b1668564b57fed383102e0655e5c2fd223f895bd3f947a1500d20e88bd4b90)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-3c60f6da9694218fb3051479cd4ae6839e1d1e2c4e673b23b837ae1e2e175268"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b1d9d3586b16e831a1693f6678b9ff4e3cebd04f63736d1b3a0bda71b572fd7c"></a>

## proxy_config.https.tls_parameters.tls_config.custom_security — proxy_config.https.tls_parameters.tls_config.custom_security / 51f89503f057 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-003.md#canonical-12b1668564b57fed383102e0655e5c2fd223f895bd3f947a1500d20e88bd4b90)
- [proxy_config.https.tls_parameters.tls_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-f85d66899e619671943d5da274f98e215365a6a9cd9991d9e0733009dd9f6485)
- proxy_config.https.tls_parameters.tls_config.custom_security

<a id="canonical-dab74c818c4054c0c4d8faf2505db7b64529942469a949e9c8c7ebd21c6a363d"></a>

Type: `"single"`. Computed.

Defines TLS protocol config including min/max versions and allowed ciphers.

Upstream description:

This defines TLS protocol config including min/max versions and allowed ciphers.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0f4587eaee0b3eda1bcc5021254d61ee0a00270ac8469088218418907c149ee1"></a>

## Direct properties — proxy_config.https.tls_parameters.tls_config.custom_security / 51f89503f057 / 3

<a id="canonical-cd066ff0879df58c3e76146a7ffe5a18d9124cdf03a0054818e875f9b26e9b61"></a>

<a id="canonical-cae529befa3bc0b5cc8c6bb0680be533e0fd67c6beddf7d8ce25db4162798eb4"></a>

## cipher_suites property — proxy_config.https.tls_parameters.tls_config.custom_security / 51f89503f057 / 4

Type: `["list", "string"]`. Computed.

The TLS listener will only support the specified cipher list.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-08982dc8b3eec74b7d8de02401e83f777844bb9bdc3f8c48becfaccad15b0ca8"></a>

<a id="canonical-54b53e036162cf6c3f01914e3588d40595c173dff51ed85a256c0a86f1f38ac8"></a>

## max_version property — proxy_config.https.tls_parameters.tls_config.custom_security / 51f89503f057 / 5

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-f9823dfece0c40f23af44eece687cb5d1db21ac1f8257b06624493749a7ae486"></a>

<a id="canonical-07832f480e4bb49eb4f298d2ed89c06694a9d3e529c5749ad97379f8a3fd50ed"></a>

## min_version property — proxy_config.https.tls_parameters.tls_config.custom_security / 51f89503f057 / 6

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2fc51eea3ec1651bbe849353548055a24d215d8dd48a43a0a04c7e444e7124d2"></a>

## Next pages — proxy_config.https.tls_parameters.tls_config.custom_security / 51f89503f057 / 7

- [proxy_config.https.tls_parameters.tls_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-f85d66899e619671943d5da274f98e215365a6a9cd9991d9e0733009dd9f6485)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-3ed63bd061fdac7cc226c995f36938f661a1100ea20bc57de9c83df695259d1c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f058cdfc36904a115fa044da942b1c586c99ad9caeb2c566cac8ddd6b81cdab4"></a>

## proxy_config.https.tls_parameters.tls_config.default_security — proxy_config.https.tls_parameters.tls_config.default_security / f46a10d19e83 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-003.md#canonical-12b1668564b57fed383102e0655e5c2fd223f895bd3f947a1500d20e88bd4b90)
- [proxy_config.https.tls_parameters.tls_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-f85d66899e619671943d5da274f98e215365a6a9cd9991d9e0733009dd9f6485)
- proxy_config.https.tls_parameters.tls_config.default_security

<a id="canonical-8628d1e0dde400350dea5d17041c91fcfde3a91d386ed85ff663907db1b2f8b8"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-266b4254c9bf9d51f1b5007e882f03e0491f65794da7bc6c7161d6c0c165233d"></a>

## Direct properties — proxy_config.https.tls_parameters.tls_config.default_security / f46a10d19e83 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8c1564b406ca402bb1a2d964eece9a8e32b5b687b7a2964e6614766e2b3bf1d6"></a>

## Next pages — proxy_config.https.tls_parameters.tls_config.default_security / f46a10d19e83 / 4

- [proxy_config.https.tls_parameters.tls_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-f85d66899e619671943d5da274f98e215365a6a9cd9991d9e0733009dd9f6485)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-3424d97af32411d3972e52f2b950d15f862787bbb9b921d012827fe960d7956d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6e26708f24283ebf2851883f66c64befb32eb8b03c4d7730f68a49b359ca4df0"></a>

## proxy_config.https.tls_parameters.tls_config.low_security — proxy_config.https.tls_parameters.tls_config.low_security / 4128fdd20d7e / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-003.md#canonical-12b1668564b57fed383102e0655e5c2fd223f895bd3f947a1500d20e88bd4b90)
- [proxy_config.https.tls_parameters.tls_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-f85d66899e619671943d5da274f98e215365a6a9cd9991d9e0733009dd9f6485)
- proxy_config.https.tls_parameters.tls_config.low_security

<a id="canonical-68b5605ff70509746a3aa996744a2f4f16b665e6268c85fc228aeb92d4d6b5ba"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2e8125f8bd54db51ca40787612228d9bae79f307b9e38c7f18c9e0bc29e6b3cd"></a>

## Direct properties — proxy_config.https.tls_parameters.tls_config.low_security / 4128fdd20d7e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c3e28d625aa7f39fae0a74ad441ec12dac2373ed3caeb46fc637f4a6fdff6088"></a>

## Next pages — proxy_config.https.tls_parameters.tls_config.low_security / 4128fdd20d7e / 4

- [proxy_config.https.tls_parameters.tls_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-f85d66899e619671943d5da274f98e215365a6a9cd9991d9e0733009dd9f6485)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-b530df3792dc26c5f21a992373bbdf02920e2c3f93d4f4248ef21840e3a9a4d8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7e99bde3da48808d34d1bdc567e40a2637578a584d014cc4bbaf3a3f8e7452ab"></a>

## proxy_config.https.tls_parameters.tls_config.medium_security — proxy_config.https.tls_parameters.tls_config.medium_security / 3fd1407427ea / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-003.md#canonical-12b1668564b57fed383102e0655e5c2fd223f895bd3f947a1500d20e88bd4b90)
- [proxy_config.https.tls_parameters.tls_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-f85d66899e619671943d5da274f98e215365a6a9cd9991d9e0733009dd9f6485)
- proxy_config.https.tls_parameters.tls_config.medium_security

<a id="canonical-c59a90728296f4bee9e7132a3890c2ffe3dc3c7695f3f38718b347e12cf45da9"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-cb9d8c966abc3703ffb34ca743818d544a7bdf864078c39017d9448be59a5eb6"></a>

## Direct properties — proxy_config.https.tls_parameters.tls_config.medium_security / 3fd1407427ea / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8e2a4f2e1b493407e024f5ed1db26337bd7cd00d22b0a802ad1f01661eed4964"></a>

## Next pages — proxy_config.https.tls_parameters.tls_config.medium_security / 3fd1407427ea / 4

- [proxy_config.https.tls_parameters.tls_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-f85d66899e619671943d5da274f98e215365a6a9cd9991d9e0733009dd9f6485)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-86372ca7fbc2e20f711797628e4eeb717449261b6371d6fc347c5aba93edf190"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f6c33dad058b905e1733d7d03a7566d07b372981079049f1683f4876bbdbbf57"></a>

## proxy_config.https.tls_parameters.use_mtls — proxy_config.https.tls_parameters.use_mtls / c688c6caa028 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-003.md#canonical-12b1668564b57fed383102e0655e5c2fd223f895bd3f947a1500d20e88bd4b90)
- proxy_config.https.tls_parameters.use_mtls

<a id="canonical-646ac82f0848ff79909c59a0584afaa46ed95ca6deb78715afca511ccc27ff62"></a>

Type: `"single"`. Computed.

Validation context for downstream client TLS connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-crl_choice": "[\"crl\",\"no_crl\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-xfcc_header": "[\"xfcc_disabled\",\"xfcc_options\"]"
}
```

<a id="canonical-fca2160d0c2bb3583a4720e4c5ab7f590fbcb278e7165f3e277621b00e0e400c"></a>

## Direct properties — proxy_config.https.tls_parameters.use_mtls / c688c6caa028 / 3

<a id="canonical-6aa523923112c9e3b317960a51a578529eb155f301923472e4a90242f866d56b"></a>

<a id="canonical-4e6d8fba0120b9997c5f1bf30ddfd7b0b2e89603ec6ba52a35d3dfff2e4bd08c"></a>

## client_certificate_optional property — proxy_config.https.tls_parameters.use_mtls / c688c6caa028 / 4

Type: `"bool"`. Computed.

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated.

Upstream description:

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated. If the client
does not provide a certificate, the connection will be accepted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [crl](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d32fbc7b152cf520302d9994a2ee54f62d806b43ea3652ec7f5487803bac7890): complete subsection reference.

- [no_crl](data-sources--bigip_http_proxy--reference--group-003.md#canonical-e72a9b293fa71c4ebedda269dced3c5eee63c931aec016e43856df1ab53ce6f5): complete subsection reference.

- [trusted_ca](data-sources--bigip_http_proxy--reference--group-003.md#canonical-57eec0b230ec014d0539ea4c314f2de844694950150b686c754390fc04113036): complete subsection reference.

<a id="canonical-23f63b3bb5124a716a10c3a548bd5e6ed44cb4070c8c44589ef422826ed2ae31"></a>

<a id="canonical-5bc03e223b0ba99105fdf1a1ba383c123386ea6b44762fd19ab513465ab4dca2"></a>

## trusted_ca_url property — proxy_config.https.tls_parameters.use_mtls / c688c6caa028 / 5

Type: `"string"`. Computed.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Upstream description:

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

- [xfcc_disabled](data-sources--bigip_http_proxy--reference--group-004.md#canonical-99a7b6ea2820551a1c2746a1ee6bb38e89719aa6fb15655b46cd67a766a646ad): complete subsection reference.

- [xfcc_options](data-sources--bigip_http_proxy--reference--group-004.md#canonical-917867436879ba25f4b16d939daa8aae74dc830bffd147d9ff4cccd7d94d1236): complete subsection reference.

<a id="canonical-ef28311ec9c3ed8c4a1443124a2c5b8771ddc5d75e18d446d583d770edca6c9b"></a>

## Next pages — proxy_config.https.tls_parameters.use_mtls / c688c6caa028 / 6

- [proxy_config.https.tls_parameters.use_mtls.crl](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d32fbc7b152cf520302d9994a2ee54f62d806b43ea3652ec7f5487803bac7890)
- [proxy_config.https.tls_parameters.use_mtls.no_crl](data-sources--bigip_http_proxy--reference--group-003.md#canonical-e72a9b293fa71c4ebedda269dced3c5eee63c931aec016e43856df1ab53ce6f5)
- [proxy_config.https.tls_parameters.use_mtls.trusted_ca](data-sources--bigip_http_proxy--reference--group-003.md#canonical-57eec0b230ec014d0539ea4c314f2de844694950150b686c754390fc04113036)
- [proxy_config.https.tls_parameters.use_mtls.xfcc_disabled](data-sources--bigip_http_proxy--reference--group-004.md#canonical-99a7b6ea2820551a1c2746a1ee6bb38e89719aa6fb15655b46cd67a766a646ad)
- [proxy_config.https.tls_parameters.use_mtls.xfcc_options](data-sources--bigip_http_proxy--reference--group-004.md#canonical-917867436879ba25f4b16d939daa8aae74dc830bffd147d9ff4cccd7d94d1236)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-003.md#canonical-12b1668564b57fed383102e0655e5c2fd223f895bd3f947a1500d20e88bd4b90)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-d32fbc7b152cf520302d9994a2ee54f62d806b43ea3652ec7f5487803bac7890"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-212daeda50eb8bbbcc2c8f3a55ca6a0c0c9e60b58fea2ec37be264a16662abd7"></a>

## proxy_config.https.tls_parameters.use_mtls.crl — proxy_config.https.tls_parameters.use_mtls.crl / a9c56799bba2 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-003.md#canonical-12b1668564b57fed383102e0655e5c2fd223f895bd3f947a1500d20e88bd4b90)
- [proxy_config.https.tls_parameters.use_mtls](data-sources--bigip_http_proxy--reference--group-003.md#canonical-86372ca7fbc2e20f711797628e4eeb717449261b6371d6fc347c5aba93edf190)
- proxy_config.https.tls_parameters.use_mtls.crl

<a id="canonical-ac282277f0bfac738e807036bc78f94d0fa4f31c48f83f6e2c059f052da40b3e"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-7b24c135fa12663b5d8f811d13f865cc7cdcb8a2acfc0f75a77f0e7758151516"></a>

## Direct properties — proxy_config.https.tls_parameters.use_mtls.crl / a9c56799bba2 / 3

<a id="canonical-088cb962e8083098b42993fc7769fe21401a492f276ac5efe7e627b8bc9f1180"></a>

<a id="canonical-65c1344c850fb35a7aaf682a0fe938d0b23862b1b59fff28a34a04712b8c314f"></a>

## name property — proxy_config.https.tls_parameters.use_mtls.crl / a9c56799bba2 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-a193b3fd0778488c6d732b5e8e814998215c68ec4156075cad3cd91301cc9886"></a>

<a id="canonical-15a25cb6fe2287b923750e8cd0f64d1ba1e4c3fb9e95e1a613bb01633be13ef8"></a>

## namespace property — proxy_config.https.tls_parameters.use_mtls.crl / a9c56799bba2 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-acae7262f19349826969a4e4b975cc320376b3d7e5946de7bf2c7cf6c123ae71"></a>

<a id="canonical-131439dd4dc945be00003c3881bba1d8c7774b22ed970d022fd4aff0b0cfee72"></a>

## tenant property — proxy_config.https.tls_parameters.use_mtls.crl / a9c56799bba2 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-e19e3afc0bb888a2488cdbfeb13912390693fd6f17854cb97386b3faf84b722e"></a>

## Next pages — proxy_config.https.tls_parameters.use_mtls.crl / a9c56799bba2 / 7

- [proxy_config.https.tls_parameters.use_mtls](data-sources--bigip_http_proxy--reference--group-003.md#canonical-86372ca7fbc2e20f711797628e4eeb717449261b6371d6fc347c5aba93edf190)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-e72a9b293fa71c4ebedda269dced3c5eee63c931aec016e43856df1ab53ce6f5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-602ca193f9ab89a9e347f8bdb1332aa6a309310946127a59fc5012a160242877"></a>

## proxy_config.https.tls_parameters.use_mtls.no_crl — proxy_config.https.tls_parameters.use_mtls.no_crl / 68a208042d33 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-003.md#canonical-12b1668564b57fed383102e0655e5c2fd223f895bd3f947a1500d20e88bd4b90)
- [proxy_config.https.tls_parameters.use_mtls](data-sources--bigip_http_proxy--reference--group-003.md#canonical-86372ca7fbc2e20f711797628e4eeb717449261b6371d6fc347c5aba93edf190)
- proxy_config.https.tls_parameters.use_mtls.no_crl

<a id="canonical-6d651873c3744c970e0cf679680fd08e0ec6cbf3864dd22005fe74f5d8aab46b"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-5aec0f908765565ee1a29ebf145e4caff260033abc0f9cad2f9a8828817a5ec9"></a>

## Direct properties — proxy_config.https.tls_parameters.use_mtls.no_crl / 68a208042d33 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d6a53716dcb660c0e114679eefecf186e598a0cd2d8ca8dd75810631673df490"></a>

## Next pages — proxy_config.https.tls_parameters.use_mtls.no_crl / 68a208042d33 / 4

- [proxy_config.https.tls_parameters.use_mtls](data-sources--bigip_http_proxy--reference--group-003.md#canonical-86372ca7fbc2e20f711797628e4eeb717449261b6371d6fc347c5aba93edf190)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-57eec0b230ec014d0539ea4c314f2de844694950150b686c754390fc04113036"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
