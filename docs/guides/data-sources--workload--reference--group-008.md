---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-403589b1465936cf5278dec49d1d397c14dd3c9843d842c54a686ec820135483"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 40d6e4750c2c / 7

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route](data-sources--workload--reference--group-007.md#canonical-d7e654f0d96f665293b19ab80b310bb70cf41bd6d32068cd65ffac82fb134936)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-6f1b39b7d16fd9d8a4737bb0c4fefab43db418fb409a93d67b085aae4cc30ed1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f8491aa6b229265f9d2d8b375e01ddac79f966237cbb24cfd2369c4343ed82d7"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 0bcc06f7ed60 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-fe3574b4f2af2fc6d8f8000c8ef6e8baa5bd8847f507f148a99a7c6940e27c62)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-d2f2640345dbdf8a6ca20b7b3a2ceb6c31007e44a4f25546827758c4a23df912)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-7aee7d3f01b8589d73b60b2b9ba278f9df11405cd3f693ab4300b3a9cfbf189c)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-007.md#canonical-834006c937eaad6c8f1af05c2ebf1bd39aa6ba171e7d68395eaafc99f2b23801)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-007.md#canonical-d355696c587815652ac9eb7f34afc67c59702436c2082b5011e8d692aedcf066)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route](data-sources--workload--reference--group-007.md#canonical-d7e654f0d96f665293b19ab80b310bb70cf41bd6d32068cd65ffac82fb134936)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response

<a id="canonical-8b2bc83f55ad339a56679b02d8c543eeda8dbb8d486370ee97ab296cd25d0030"></a>

Type: `"single"`. Computed.

Send this direct response in case of route match action is direct response.

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

<a id="canonical-86fa2c311c158f327147ffad84755657a35dd59162c5557c017f1164aef148d8"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 0bcc06f7ed60 / 3

<a id="canonical-c91b648d5e1a36d979a6b46b5580a9a8b1bfa373fd81dbbda2eb1e47c30e6394"></a>

<a id="canonical-3a4d0b93ac1603265e38f0133b78c999b5c37c8e341efbf23984297d2505ce90"></a>

## response_body_encoded property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 0bcc06f7ed60 / 4

Type: `"string"`. Computed.

Response body to send. Currently supported URL schemes is string:/// for which message should be
encoded in Base64 format. The message can be either plain text or HTML.

Upstream description:

Response body to send. Currently supported URL schemes is string:/// for which message should be
encoded in Base64 format. The message can be either plain text or HTML. E.g. "&lt;p&gt; Access
Denied &lt;/p&gt;". Base64 encoded string URL for this is
string:///PHA+IEFjY2VzcyBEZW5pZWQgPC9wPg==.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 65536,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 65536
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 65536,
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
    "ves.io.schema.rules.string.max_bytes": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-e1c4dd0e49a9ef8db8d46b81469568cb78cc3999557a9e90e25e581b2bb123b6"></a>

<a id="canonical-8f7320976870556613bd7834d8fb8134799d296241751cf462b7c0e20d6b4a8d"></a>

## response_code property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 0bcc06f7ed60 / 5

Type: `"number"`. Computed.

Response Code. Response code to send.

Upstream description:

Response code to send.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 599,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 100
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "100",
    "ves.io.schema.rules.uint32.lte": "599"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "100",
    "ves.io.schema.rules.uint32.lte": "599"
  }
}
```

<a id="canonical-f00e38b5461fcfc64947a2989622c45040d3baf590135e993d1339828dd61da0"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 0bcc06f7ed60 / 6

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route](data-sources--workload--reference--group-007.md#canonical-d7e654f0d96f665293b19ab80b310bb70cf41bd6d32068cd65ffac82fb134936)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-610b06066040c7fc9f728e970640a9895280a77086b1fc0104d83cf4cea3741e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d6055a9d1a219d8ea68d84047d11546fb1a7dd2f1bc40bbb7f344322ec6c3d31"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / acfa74d6b45a / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-fe3574b4f2af2fc6d8f8000c8ef6e8baa5bd8847f507f148a99a7c6940e27c62)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-d2f2640345dbdf8a6ca20b7b3a2ceb6c31007e44a4f25546827758c4a23df912)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-7aee7d3f01b8589d73b60b2b9ba278f9df11405cd3f693ab4300b3a9cfbf189c)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-007.md#canonical-834006c937eaad6c8f1af05c2ebf1bd39aa6ba171e7d68395eaafc99f2b23801)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-007.md#canonical-d355696c587815652ac9eb7f34afc67c59702436c2082b5011e8d692aedcf066)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route

<a id="canonical-c3ee2e582cee6ac8114ad0794e1ae099f5bce500e0a478d13012d6658b972657"></a>

Type: `"single"`. Computed.

Redirect route matches on path, incoming header, incoming port and/or HTTP method and redirects the
matching traffic to a different URL.

Upstream description:

A redirect route matches on path, incoming header, incoming port and/or HTTP method and redirects
the matching traffic to a different URL.

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

<a id="canonical-6de6ebea69efb6c4cf83bf02a7f71f1ef6ef1632e7e8fc0b971762426047b42d"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / acfa74d6b45a / 3

- [headers](data-sources--workload--reference--group-008.md#canonical-f34272fc6cda27c84136d6665616825ac4688590dbdb8a207285c17de9b2cedc): complete subsection reference.

<a id="canonical-086756341cc177edc8c768941132d0346f154c13f28e1d2fbf2f9d5d1fd327cb"></a>

<a id="canonical-6b49450d55be4d3ae05bd0e24391ffb9d91ea20fb83d280292a80055d921379b"></a>

## http_method property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / acfa74d6b45a / 4

Type: `"string"`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

Upstream description:

Specifies the HTTP method used to access a resource.

Any HTTP Method.

Receipt-pinned upstream constraints:

```json
{
  "default": "ANY",
  "enum": [
    "ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [incoming_port](data-sources--workload--reference--group-008.md#canonical-23b17183d9782bc281be1775fd6fd125241fe0daa4639d0db1345d0cfe8d3cb2): complete subsection reference.

- [path](data-sources--workload--reference--group-008.md#canonical-745d5c44a40fbd19b160e1d81d3bca451f90a07040f4c62c013b070ffeba186b): complete subsection reference.

- [route_redirect](data-sources--workload--reference--group-008.md#canonical-777fe1102aeb952d0dcfc9ea7ebaaacd51eb2198f903fc5979749fc4b8a6db18): complete subsection reference.

<a id="canonical-f46340cc2694c84d2917678e0eb46e4d50f58fb432235a0c172f5e6200c0a8b7"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / acfa74d6b45a / 5

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.headers](data-sources--workload--reference--group-008.md#canonical-f34272fc6cda27c84136d6665616825ac4688590dbdb8a207285c17de9b2cedc)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port](data-sources--workload--reference--group-008.md#canonical-23b17183d9782bc281be1775fd6fd125241fe0daa4639d0db1345d0cfe8d3cb2)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.path](data-sources--workload--reference--group-008.md#canonical-745d5c44a40fbd19b160e1d81d3bca451f90a07040f4c62c013b070ffeba186b)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](data-sources--workload--reference--group-008.md#canonical-777fe1102aeb952d0dcfc9ea7ebaaacd51eb2198f903fc5979749fc4b8a6db18)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-007.md#canonical-d355696c587815652ac9eb7f34afc67c59702436c2082b5011e8d692aedcf066)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-f34272fc6cda27c84136d6665616825ac4688590dbdb8a207285c17de9b2cedc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b3610c9b327468180b025fbe99042c512e07f5d6bdfdf9c63256631d72a6f49c"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.headers — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / d611735c1a34 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-fe3574b4f2af2fc6d8f8000c8ef6e8baa5bd8847f507f148a99a7c6940e27c62)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-d2f2640345dbdf8a6ca20b7b3a2ceb6c31007e44a4f25546827758c4a23df912)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-7aee7d3f01b8589d73b60b2b9ba278f9df11405cd3f693ab4300b3a9cfbf189c)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-007.md#canonical-834006c937eaad6c8f1af05c2ebf1bd39aa6ba171e7d68395eaafc99f2b23801)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-007.md#canonical-d355696c587815652ac9eb7f34afc67c59702436c2082b5011e8d692aedcf066)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-008.md#canonical-610b06066040c7fc9f728e970640a9895280a77086b1fc0104d83cf4cea3741e)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.headers

<a id="canonical-b06fbf5e289056ef2da72cc51f891d31bb0dcd4d02be09926819ae4c3402287c"></a>

Type: `"list"`. Computed.

Headers. List of (key, value) headers.

Upstream description:

List of (key, value) headers.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
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
    "minItems": 0,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2ad3a4cdddff6b8be2921b6d185da5cd851d9f0ce70f15ab6328a1fda0d76083"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / d611735c1a34 / 3

<a id="canonical-7f8c938b6951e89e2e38bf96df6039e6748f5dea6306c69fc796183775ebed74"></a>

<a id="canonical-5f3c09509d368490aa06adc952be5df5110d4f773d26b328890444d126c4f52a"></a>

## exact property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / d611735c1a34 / 4

Type: `"string"`. Computed.

Exclusive with \[presence regex\] Header value to match exactly.

Upstream description:

Exclusive with \[presence regex\] Header value to match exactly.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-acceb1cbfd2e8001eda6b3209c77dc442efeda8ecbac209d2c2217d1be5a33e8"></a>

<a id="canonical-f39a4fb8f726249c106b2ca426a1b244a11c7239dd9903887d960f9832331026"></a>

## invert_match property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / d611735c1a34 / 5

Type: `"bool"`. Computed.

Invert the result of the match to detect missing header or non-matching value.

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

<a id="canonical-8c5958c3e65b2bd1a67249af981c27ac6e106f44d74f5f9bba2a2531a533ad73"></a>

<a id="canonical-fa3297ad315539b1e124496fec406d3f733f65dd78bb2ebf3b7a5bdcdd7d6206"></a>

## name property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / d611735c1a34 / 6

Type: `"string"`. Computed.

Name. Name of the header.

Upstream description:

Name of the header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-15d29c1298e76c38412a2bc4c9f6dca93aa98af07af0f9cc4814c8a24e4c11a9"></a>

<a id="canonical-29c757dd019bfc5bcad6ea23c622deb33a112832727a8bd478121fe248a217d9"></a>

## presence property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / d611735c1a34 / 7

Type: `"bool"`. Computed.

Exclusive with \[exact regex\] If true, check for presence of header.

Upstream description:

Exclusive with \[exact regex\] If true, check for presence of header.

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

<a id="canonical-4683150186d38be049b87a6cd4e213f5181ed4f51f3378ed40286b550ea78998"></a>

<a id="canonical-67087e234485758cb4e7c2ae102d8ba2a764e8966ada3b8f43186f4771359aef"></a>

## regex property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / d611735c1a34 / 8

Type: `"string"`. Computed.

Exclusive with \[exact presence\] Regex match of the header value in re2 format.

Upstream description:

Exclusive with \[exact presence\] Regex match of the header value in re2 format.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-57f7bf132da03e503e755951e84b1200af5400c2302d41c0da489d709f5ec615"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / d611735c1a34 / 9

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-008.md#canonical-610b06066040c7fc9f728e970640a9895280a77086b1fc0104d83cf4cea3741e)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-23b17183d9782bc281be1775fd6fd125241fe0daa4639d0db1345d0cfe8d3cb2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-935d484c30b5d8038225e6f1d81af1d3690cabc7bcbf7b717acb144d1805bb79"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 9d32055100a0 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-fe3574b4f2af2fc6d8f8000c8ef6e8baa5bd8847f507f148a99a7c6940e27c62)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-d2f2640345dbdf8a6ca20b7b3a2ceb6c31007e44a4f25546827758c4a23df912)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-7aee7d3f01b8589d73b60b2b9ba278f9df11405cd3f693ab4300b3a9cfbf189c)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-007.md#canonical-834006c937eaad6c8f1af05c2ebf1bd39aa6ba171e7d68395eaafc99f2b23801)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-007.md#canonical-d355696c587815652ac9eb7f34afc67c59702436c2082b5011e8d692aedcf066)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-008.md#canonical-610b06066040c7fc9f728e970640a9895280a77086b1fc0104d83cf4cea3741e)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port

<a id="canonical-ce6511af99bf8b291d311f7228a37afd4d91cc41ef72aab654e64354bd45625b"></a>

Type: `"single"`. Computed.

Port match of the request can be a range or a specific port.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_match": "[\"no_port_match\",\"port\",\"port_ranges\"]"
}
```

<a id="canonical-72488afac7beea1c7c6d463045cf38e14d4c1b6e8a699d37723bf99b3b5f97ae"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 9d32055100a0 / 3

- [no_port_match](data-sources--workload--reference--group-008.md#canonical-a4bfdb21e2aad4c0d58e0f18401117ebc17b5bd8f0f07401c5672d6487113bd8): complete subsection reference.

<a id="canonical-e0a6c5a1456cc2e77da1e313398937069aaec6cb6f9f2b6af26a3ddeac0ed5a8"></a>

<a id="canonical-5b5594b47ea07b59a982058e3d188ec551cd34c8093565cdeaaf912c0abf38c0"></a>

## port property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 9d32055100a0 / 4

Type: `"number"`. Computed.

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

Upstream description:

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

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

<a id="canonical-08c974fcebe664467d950a55b7b6a58f030a70a86a8c8c9138b2eab416439886"></a>

<a id="canonical-7babff1a7f9bb4c0cc7158088c8be6aa2b7af8536ef373a16658a40b5d434c94"></a>

## port_ranges property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 9d32055100a0 / 5

Type: `"string"`. Computed.

Exclusive with \[no\_port\_match port\] Port range to match.

Upstream description:

Exclusive with \[no\_port\_match port\] Port range to match.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 32,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 32,
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
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  }
}
```

<a id="canonical-f2790aab54b3cf6db8945e0b4438145af89979f20bd5f35f4d61f6f9a410cd8f"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 9d32055100a0 / 6

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port.no_port_match](data-sources--workload--reference--group-008.md#canonical-a4bfdb21e2aad4c0d58e0f18401117ebc17b5bd8f0f07401c5672d6487113bd8)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-008.md#canonical-610b06066040c7fc9f728e970640a9895280a77086b1fc0104d83cf4cea3741e)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-a4bfdb21e2aad4c0d58e0f18401117ebc17b5bd8f0f07401c5672d6487113bd8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-64e5ac6691ceefff687d311ffc711e744067badffa74015318888ca5caa5be9a"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port.no_port_match — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / cadc7b0eedee / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-fe3574b4f2af2fc6d8f8000c8ef6e8baa5bd8847f507f148a99a7c6940e27c62)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-d2f2640345dbdf8a6ca20b7b3a2ceb6c31007e44a4f25546827758c4a23df912)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-7aee7d3f01b8589d73b60b2b9ba278f9df11405cd3f693ab4300b3a9cfbf189c)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-007.md#canonical-834006c937eaad6c8f1af05c2ebf1bd39aa6ba171e7d68395eaafc99f2b23801)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-007.md#canonical-d355696c587815652ac9eb7f34afc67c59702436c2082b5011e8d692aedcf066)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-008.md#canonical-610b06066040c7fc9f728e970640a9895280a77086b1fc0104d83cf4cea3741e)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port](data-sources--workload--reference--group-008.md#canonical-23b17183d9782bc281be1775fd6fd125241fe0daa4639d0db1345d0cfe8d3cb2)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port.no_port_match

<a id="canonical-2d2ed7538458231799f438bc70f96b86273d33de6e188059ca4ec872ee735d98"></a>

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

<a id="canonical-d76df789a17e4d9bfa04761e68dd459db3aacf49aad930159a5348c32dfc6cec"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / cadc7b0eedee / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9634f68e2f2ecb0fbaa62ed24019bba6358f4644dbfd11284b3db9c0d5d3166a"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / cadc7b0eedee / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port](data-sources--workload--reference--group-008.md#canonical-23b17183d9782bc281be1775fd6fd125241fe0daa4639d0db1345d0cfe8d3cb2)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-745d5c44a40fbd19b160e1d81d3bca451f90a07040f4c62c013b070ffeba186b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a4de5a735b5ed3d2eb05f6506fc0afc5d811aa787fc1a91021fb044a1d18b52d"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.path — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / c3e6a2922da2 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-fe3574b4f2af2fc6d8f8000c8ef6e8baa5bd8847f507f148a99a7c6940e27c62)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-d2f2640345dbdf8a6ca20b7b3a2ceb6c31007e44a4f25546827758c4a23df912)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-7aee7d3f01b8589d73b60b2b9ba278f9df11405cd3f693ab4300b3a9cfbf189c)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-007.md#canonical-834006c937eaad6c8f1af05c2ebf1bd39aa6ba171e7d68395eaafc99f2b23801)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-007.md#canonical-d355696c587815652ac9eb7f34afc67c59702436c2082b5011e8d692aedcf066)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-008.md#canonical-610b06066040c7fc9f728e970640a9895280a77086b1fc0104d83cf4cea3741e)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.path

<a id="canonical-ba1f0dac27e6138b7e0629f604ab5e76b10ba7069d27fff23a78d828f60f1fb5"></a>

Type: `"single"`. Computed.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

<a id="canonical-c741c159ab183cec256d98d04c0d8e6765c98bec42838fff4dac94bee7ce70c0"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / c3e6a2922da2 / 3

<a id="canonical-fc6a9b946c4d0bcd386201fd9c5a395d4790ceebaf9e77eae81f3729e9ea5d3d"></a>

<a id="canonical-0e38a6c74c068a0659a2363dd48367eaf43b074f90b4f8be51ab678b51fdbb84"></a>

## path property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / c3e6a2922da2 / 4

Type: `"string"`. Computed.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-dec4c84b865d0f9827ed0cee7653486cf4c834c173fd10829a0f3c7ed9331530"></a>

<a id="canonical-1577c3011b1a837ded6da677382ec9c9f6411da0a205eb0eb8536bfe86526a9d"></a>

## prefix property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / c3e6a2922da2 / 5

Type: `"string"`. Computed.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-e105736d950db001b6b79e6643116f0ba417501e2a4080b3429c6d55511c27d6"></a>

<a id="canonical-cf4c18f4f3671da2a2994f2a88041a8c7a1217ae64188cae7a38a0f1f8a2c981"></a>

## regex property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / c3e6a2922da2 / 6

Type: `"string"`. Computed.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-569379bbdaf647572ede4e8813a649230b751a00b7abd654d5499788683514e9"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / c3e6a2922da2 / 7

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-008.md#canonical-610b06066040c7fc9f728e970640a9895280a77086b1fc0104d83cf4cea3741e)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-777fe1102aeb952d0dcfc9ea7ebaaacd51eb2198f903fc5979749fc4b8a6db18"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bdf29537b51304e8873579beab61f65570cbdd8ae068d51c0ea250dcd0d5e9d5"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / c10e9c6b3dbc / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-fe3574b4f2af2fc6d8f8000c8ef6e8baa5bd8847f507f148a99a7c6940e27c62)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-d2f2640345dbdf8a6ca20b7b3a2ceb6c31007e44a4f25546827758c4a23df912)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-7aee7d3f01b8589d73b60b2b9ba278f9df11405cd3f693ab4300b3a9cfbf189c)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-007.md#canonical-834006c937eaad6c8f1af05c2ebf1bd39aa6ba171e7d68395eaafc99f2b23801)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-007.md#canonical-d355696c587815652ac9eb7f34afc67c59702436c2082b5011e8d692aedcf066)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-008.md#canonical-610b06066040c7fc9f728e970640a9895280a77086b1fc0104d83cf4cea3741e)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect

<a id="canonical-b1ef1d77bdb7238cb05acad24b55ebe50518e3c4a2fd567cb7086a246ef1c731"></a>

Type: `"single"`. Computed.

Route redirect parameters when match action is redirect.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-query_params": "[\"remove_all_params\",\"replace_params\",\"retain_all_params\"]",
  "x-ves-oneof-field-redirect_path_choice": "[\"path_redirect\",\"prefix_rewrite\"]"
}
```

<a id="canonical-92e645cd950498c2988170d831b46bb83d683d34d259e3a5f942042554f563cd"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / c10e9c6b3dbc / 3

<a id="canonical-87db6b85607a3f816a9360af707a9e372a83fab9bf376f270f887458916f6614"></a>

<a id="canonical-ed7bd4737a7e33d8d162039a4599457f998404bafc93bdc3650286c6a4351230"></a>

## host_redirect property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / c10e9c6b3dbc / 4

Type: `"string"`. Computed.

Swap host part of incoming URL in redirect URL.

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

<a id="canonical-ddaf4c1c8bf08d33ee5acca69bdf2c270b08aae3a3adaa1a583f0147bb37869c"></a>

<a id="canonical-54998d45a6ce7d96c3756df9342b51ce839b992aee390baf8d4daf66f923d14d"></a>

## path_redirect property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / c10e9c6b3dbc / 5

Type: `"string"`. Computed.

Exclusive with \[prefix\_rewrite\] swap path part of incoming URL in redirect URL.

Upstream description:

Exclusive with \[prefix\_rewrite\] swap path part of incoming URL in redirect URL.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-c1db08c0657d5bee29c22a9bb15ff5630d00359e5e6f06f94ce554d26ec72dd0"></a>

<a id="canonical-15d5a606653c4aa80b11b75b40c346b93bde0246a48cd180e2518a18b11c3fe5"></a>

## prefix_rewrite property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / c10e9c6b3dbc / 6

Type: `"string"`. Computed.

Exclusive with \[path\_redirect\] In Redirect response, the matched prefix (or path) should be
swapped with this value. This option allows redirect URLs be dynamically created based on the
request.

Upstream description:

Exclusive with \[path\_redirect\] In Redirect response, the matched prefix (or path) should be
swapped with this value. This option allows redirect URLs be dynamically created based on the
request.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-216d0523a757a5f4dc0921f7c3f20a9a0783f411140a806ee81d5e4f47007ca3"></a>

<a id="canonical-f13d87d3aac86198644d6b00831a782ace8bc28adc68980f42572c7985216d01"></a>

## proto_redirect property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / c10e9c6b3dbc / 7

Type: `"string"`. Computed.

\[Enum: incoming-proto|http|https\] Swap protocol part of incoming URL in redirect URL The protocol
can be swapped with either HTTP or HTTPS When incoming-proto option is specified, swapping of
protocol is not done. Possible values are \`incoming-proto\`, \`http\`, \`https\`.

Upstream description:

Swap protocol part of incoming URL in redirect URL The protocol can be swapped with either HTTP or
HTTPS When incoming-proto option is specified, swapping of protocol is not done.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "incoming-proto",
    "http",
    "https"
  ],
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
    "ves.io.schema.rules.string.in": "[\\\"incoming-proto\\\",\\\"http\\\",\\\"https\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"incoming-proto\\\",\\\"http\\\",\\\"https\\\"]"
  }
}
```

- [remove_all_params](data-sources--workload--reference--group-008.md#canonical-6dae1914b0d869e395968d41d5b5fc7edb6f27739cd7447b6aa3ca6512fae3b1): complete subsection reference.

<a id="canonical-6738e02d4fbd2d0380a2f7dc5442cd7ab71e51009e06533b9ae2325727c8a26a"></a>

<a id="canonical-71cd6c727c042d80ae8eeb8761bb86d913107d2d833ecca960f75131d2ce8fa5"></a>

## replace_params property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / c10e9c6b3dbc / 8

Type: `"string"`. Computed.

Exclusive with \[remove\_all\_params retain\_all\_params\].

Upstream description:

Exclusive with \[remove\_all\_params retain\_all\_params\]

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-69cb673ab6352ee905a95b8654e164787f4ebff81af4ee27287a3dc68ea27a45"></a>

<a id="canonical-2b6ed3368bc83cce93a2bbf645f47df4501f79cf91724881603e65f9796023b6"></a>

## response_code property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / c10e9c6b3dbc / 9

Type: `"number"`. Computed.

The HTTP status code to use in the redirect response.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 599,
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
    "ves.io.schema.rules.uint32.lte": "599"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "599"
  }
}
```

- [retain_all_params](data-sources--workload--reference--group-008.md#canonical-873e44ea780363304c858b0b6153feccfbd36c3e6588b57c66e33e8c9e29104c): complete subsection reference.

<a id="canonical-3a5e8cd612f1334bf54e54d803f2b68d8b8ba139746451a904cdc0be9a707798"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / c10e9c6b3dbc / 10

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.remove_all_params](data-sources--workload--reference--group-008.md#canonical-6dae1914b0d869e395968d41d5b5fc7edb6f27739cd7447b6aa3ca6512fae3b1)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.retain_all_params](data-sources--workload--reference--group-008.md#canonical-873e44ea780363304c858b0b6153feccfbd36c3e6588b57c66e33e8c9e29104c)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-008.md#canonical-610b06066040c7fc9f728e970640a9895280a77086b1fc0104d83cf4cea3741e)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-6dae1914b0d869e395968d41d5b5fc7edb6f27739cd7447b6aa3ca6512fae3b1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8e081fd57f5187555fac247b8a65a024b5d98f7c6bb66a3129199f5135baa62d"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.remove_all_params — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / b63a5a293dbf / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-fe3574b4f2af2fc6d8f8000c8ef6e8baa5bd8847f507f148a99a7c6940e27c62)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-d2f2640345dbdf8a6ca20b7b3a2ceb6c31007e44a4f25546827758c4a23df912)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-7aee7d3f01b8589d73b60b2b9ba278f9df11405cd3f693ab4300b3a9cfbf189c)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-007.md#canonical-834006c937eaad6c8f1af05c2ebf1bd39aa6ba171e7d68395eaafc99f2b23801)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-007.md#canonical-d355696c587815652ac9eb7f34afc67c59702436c2082b5011e8d692aedcf066)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-008.md#canonical-610b06066040c7fc9f728e970640a9895280a77086b1fc0104d83cf4cea3741e)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](data-sources--workload--reference--group-008.md#canonical-777fe1102aeb952d0dcfc9ea7ebaaacd51eb2198f903fc5979749fc4b8a6db18)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.remove_all_params

<a id="canonical-4e40b32f07183333f246c17bcd17cdb7fb0abbaa7dac37557b5ff1ee789ce834"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for remove all params.

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

<a id="canonical-09f70d0a9c8fed36d16016ad426d03712af9c45e92dcabbbea405fc89e351700"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / b63a5a293dbf / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e5edb0bf386dfd87936118e0025c9c72d0d533640e35f06dc42d1dea49e31f30"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / b63a5a293dbf / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](data-sources--workload--reference--group-008.md#canonical-777fe1102aeb952d0dcfc9ea7ebaaacd51eb2198f903fc5979749fc4b8a6db18)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-873e44ea780363304c858b0b6153feccfbd36c3e6588b57c66e33e8c9e29104c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-95672a4b958e12d132717b63df00d2000b8b80655ef15489cfce9fdff8055d36"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.retain_all_params — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 51512bab0dd3 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-fe3574b4f2af2fc6d8f8000c8ef6e8baa5bd8847f507f148a99a7c6940e27c62)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-d2f2640345dbdf8a6ca20b7b3a2ceb6c31007e44a4f25546827758c4a23df912)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-7aee7d3f01b8589d73b60b2b9ba278f9df11405cd3f693ab4300b3a9cfbf189c)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-007.md#canonical-834006c937eaad6c8f1af05c2ebf1bd39aa6ba171e7d68395eaafc99f2b23801)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-007.md#canonical-d355696c587815652ac9eb7f34afc67c59702436c2082b5011e8d692aedcf066)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-008.md#canonical-610b06066040c7fc9f728e970640a9895280a77086b1fc0104d83cf4cea3741e)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](data-sources--workload--reference--group-008.md#canonical-777fe1102aeb952d0dcfc9ea7ebaaacd51eb2198f903fc5979749fc4b8a6db18)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.retain_all_params

<a id="canonical-fecc5ed190cb4c7fde6d4a7e87442c6d35fc11639e57642a1508497247379511"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for retain all params.

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

<a id="canonical-85dc5f5790a33018ca35edc3daf863762c3b9528b1762f0c6de900534b9ce11a"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 51512bab0dd3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1bc5bb8d31930fa7a2b64bacd18d99c1621a4878965a604180868250931c7d9e"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 51512bab0dd3 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](data-sources--workload--reference--group-008.md#canonical-777fe1102aeb952d0dcfc9ea7ebaaacd51eb2198f903fc5979749fc4b8a6db18)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-3459a5e4944ffe06d04b9f59b34abcb30fe13bff696f6bb79e1cbd4f58384869"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-81486e3fdb32b6d6174eeb657df66a130907f3a6c3f13288e35a5763db295b01"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 9e60fa1980c0 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-fe3574b4f2af2fc6d8f8000c8ef6e8baa5bd8847f507f148a99a7c6940e27c62)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-d2f2640345dbdf8a6ca20b7b3a2ceb6c31007e44a4f25546827758c4a23df912)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-7aee7d3f01b8589d73b60b2b9ba278f9df11405cd3f693ab4300b3a9cfbf189c)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-007.md#canonical-834006c937eaad6c8f1af05c2ebf1bd39aa6ba171e7d68395eaafc99f2b23801)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-007.md#canonical-d355696c587815652ac9eb7f34afc67c59702436c2082b5011e8d692aedcf066)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route

<a id="canonical-3c84a5f83b7a847c8eda8edcb122f52db648e7cb780ef9615ecfc476ef16927a"></a>

Type: `"single"`. Computed.

Simple route matches on path and/or HTTP method and forwards the matching traffic to the default
origin pool specified outside.

Upstream description:

A simple route matches on path and/or HTTP method and forwards the matching traffic to the default
origin pool specified outside.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-host_rewrite_params": "[\"auto_host_rewrite\",\"disable_host_rewrite\",\"host_rewrite\"]"
}
```

<a id="canonical-5e19d9c31705edd7e754ac6ab65a3cb1d57d7f893a359d002f4b82f34af03020"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 9e60fa1980c0 / 3

- [auto_host_rewrite](data-sources--workload--reference--group-008.md#canonical-d558e55d212e2169c910d96dcd194abaab38184dfd5277ce62f9e5a6f76b739a): complete subsection reference.

- [disable_host_rewrite](data-sources--workload--reference--group-008.md#canonical-04ddc259d7c90f22f09c6fb75fa17f3b70e8c3faa3a5e5a7a0c9baa648eeb1cd): complete subsection reference.

<a id="canonical-c34d7c3bb2cf7a9264a913a7013d59b8e0f5116beb1390bb5ee8a04860f41736"></a>

<a id="canonical-c13ba23486941025bd0bbb4c1920610e9d6448638596bb4c3aee69aba09a7ee7"></a>

## host_rewrite property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 9e60fa1980c0 / 4

Type: `"string"`. Computed.

Exclusive with \[auto\_host\_rewrite disable\_host\_rewrite\] Host header will be swapped with this
value.

Upstream description:

Exclusive with \[auto\_host\_rewrite disable\_host\_rewrite\] Host header will be swapped with this
value.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-14999a4d91fb7cd9cf1d82c79bb9a5b3511f500da4774ded16cd26c1a8606fc3"></a>

<a id="canonical-05448b0eeba5f156f3729fa57eef07465842cc678bf0d5a06836ee61c98b6ba1"></a>

## http_method property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 9e60fa1980c0 / 5

Type: `"string"`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

Upstream description:

Specifies the HTTP method used to access a resource.

Any HTTP Method.

Receipt-pinned upstream constraints:

```json
{
  "default": "ANY",
  "enum": [
    "ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [path](data-sources--workload--reference--group-008.md#canonical-53fe9ff0fee3228a9c160e6b8a1288167767fe31ec1b395e0951f77cbbad08cc): complete subsection reference.

<a id="canonical-91c77e3af87d914f7b8c4b55aa8c2ace987d027ab919215ba2d5f624d7fa7f1f"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 9e60fa1980c0 / 6

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.auto_host_rewrite](data-sources--workload--reference--group-008.md#canonical-d558e55d212e2169c910d96dcd194abaab38184dfd5277ce62f9e5a6f76b739a)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite](data-sources--workload--reference--group-008.md#canonical-04ddc259d7c90f22f09c6fb75fa17f3b70e8c3faa3a5e5a7a0c9baa648eeb1cd)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.path](data-sources--workload--reference--group-008.md#canonical-53fe9ff0fee3228a9c160e6b8a1288167767fe31ec1b395e0951f77cbbad08cc)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-007.md#canonical-d355696c587815652ac9eb7f34afc67c59702436c2082b5011e8d692aedcf066)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-d558e55d212e2169c910d96dcd194abaab38184dfd5277ce62f9e5a6f76b739a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1fe8f79c35c2c90aa8ef95ead81987b3421bbb7c7b081b85a5b558ee2512ca47"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.auto_host_rewrite — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / cba75057cc37 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-fe3574b4f2af2fc6d8f8000c8ef6e8baa5bd8847f507f148a99a7c6940e27c62)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-d2f2640345dbdf8a6ca20b7b3a2ceb6c31007e44a4f25546827758c4a23df912)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-7aee7d3f01b8589d73b60b2b9ba278f9df11405cd3f693ab4300b3a9cfbf189c)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-007.md#canonical-834006c937eaad6c8f1af05c2ebf1bd39aa6ba171e7d68395eaafc99f2b23801)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-007.md#canonical-d355696c587815652ac9eb7f34afc67c59702436c2082b5011e8d692aedcf066)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route](data-sources--workload--reference--group-008.md#canonical-3459a5e4944ffe06d04b9f59b34abcb30fe13bff696f6bb79e1cbd4f58384869)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.auto_host_rewrite

<a id="canonical-9061c7218b6e9094fc05d26f6c26e0e3fcf34da1136448e2242e0526f85f146b"></a>

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

<a id="canonical-f7837b3e6899edccab07de3a4a3b0536652542fc129dea4707228945aaa4afdb"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / cba75057cc37 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a191d63a06a7a9b08af545915e18bad712662612c7b22e5f645cde82eca0f036"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / cba75057cc37 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route](data-sources--workload--reference--group-008.md#canonical-3459a5e4944ffe06d04b9f59b34abcb30fe13bff696f6bb79e1cbd4f58384869)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-04ddc259d7c90f22f09c6fb75fa17f3b70e8c3faa3a5e5a7a0c9baa648eeb1cd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bef20ded10d881de6d890f0471956e9bf04d28425c10819937267892bfe23b5f"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 088b611b7084 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-fe3574b4f2af2fc6d8f8000c8ef6e8baa5bd8847f507f148a99a7c6940e27c62)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-d2f2640345dbdf8a6ca20b7b3a2ceb6c31007e44a4f25546827758c4a23df912)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-7aee7d3f01b8589d73b60b2b9ba278f9df11405cd3f693ab4300b3a9cfbf189c)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-007.md#canonical-834006c937eaad6c8f1af05c2ebf1bd39aa6ba171e7d68395eaafc99f2b23801)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-007.md#canonical-d355696c587815652ac9eb7f34afc67c59702436c2082b5011e8d692aedcf066)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route](data-sources--workload--reference--group-008.md#canonical-3459a5e4944ffe06d04b9f59b34abcb30fe13bff696f6bb79e1cbd4f58384869)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite

<a id="canonical-035ff484cc36258dbb045a6b889e6525464cb9926a94901f48c9c718087d77e3"></a>

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

<a id="canonical-d5e76f4d8427fddebefc84a92f8d8942affa1d0567cd4b6da2da821721fbaaa9"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 088b611b7084 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-afd502c8f5e932b3ff2ad979c2a2865821f3d52030d267a6bbf033d076e759c9"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 088b611b7084 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route](data-sources--workload--reference--group-008.md#canonical-3459a5e4944ffe06d04b9f59b34abcb30fe13bff696f6bb79e1cbd4f58384869)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-53fe9ff0fee3228a9c160e6b8a1288167767fe31ec1b395e0951f77cbbad08cc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-99df83fee97511309f33fe2b1bfadf8adbdfe9baeae587ad93497cbc46d49c6f"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.path — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 7a732bdb6046 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-fe3574b4f2af2fc6d8f8000c8ef6e8baa5bd8847f507f148a99a7c6940e27c62)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-d2f2640345dbdf8a6ca20b7b3a2ceb6c31007e44a4f25546827758c4a23df912)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-7aee7d3f01b8589d73b60b2b9ba278f9df11405cd3f693ab4300b3a9cfbf189c)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-007.md#canonical-834006c937eaad6c8f1af05c2ebf1bd39aa6ba171e7d68395eaafc99f2b23801)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-007.md#canonical-d355696c587815652ac9eb7f34afc67c59702436c2082b5011e8d692aedcf066)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route](data-sources--workload--reference--group-008.md#canonical-3459a5e4944ffe06d04b9f59b34abcb30fe13bff696f6bb79e1cbd4f58384869)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.path

<a id="canonical-905773bdd39294157cb07495b78df292871e23a7e8d9388bd2515527e7d5dc22"></a>

Type: `"single"`. Computed.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

<a id="canonical-73af67274a922bec86694dab8bbe487f904c819028f6176f6c515b54ec7f4e7b"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 7a732bdb6046 / 3

<a id="canonical-efc3d25556e970a6565d33c523b0e946512431365031cf47f4a5b7a17d02ec74"></a>

<a id="canonical-f2ab27c533f874967693ec780c91afa147c6c4ef44cf43fd22b16343a24501a2"></a>

## path property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 7a732bdb6046 / 4

Type: `"string"`. Computed.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-df9e781939791f2062b15164c39c0b52c5701fe52cd0da4251743e6df08a267b"></a>

<a id="canonical-21617b7291b6d92c241fe3a03aa60682272ef69ff10f0c86817ec3a9edc144b4"></a>

## prefix property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 7a732bdb6046 / 5

Type: `"string"`. Computed.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3b9f070257b3ae34abc2600b7ba95459203b1717f59404578a6735e135775fb2"></a>

<a id="canonical-3a98bc65aee4c3ed8ba707de1afdec5d72474e8ab382176c6c0794f19d9512dc"></a>

## regex property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 7a732bdb6046 / 6

Type: `"string"`. Computed.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-e57e6ab338f34bb0e304e74c9978c091be7cc00acd77da23d601c4039e9a6bb5"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 7a732bdb6046 / 7

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route](data-sources--workload--reference--group-008.md#canonical-3459a5e4944ffe06d04b9f59b34abcb30fe13bff696f6bb79e1cbd4f58384869)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-7c613d4efb08afb76c10e5423143ef3447c3a5c0c851a6306c101db9750216aa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-628f7addcc0116d4bd6658ff6d2d079a99c3fcfc91b9081dd8b3482fbe352272"></a>

## service.advertise_options.advertise_custom.ports.port — service.advertise_options.advertise_custom.ports.port / 34da10fb04c4 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-fe3574b4f2af2fc6d8f8000c8ef6e8baa5bd8847f507f148a99a7c6940e27c62)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-d2f2640345dbdf8a6ca20b7b3a2ceb6c31007e44a4f25546827758c4a23df912)
- service.advertise_options.advertise_custom.ports.port

<a id="canonical-2649a479112b4632f8a45916668304dbb50d2482e027b729070fad5be07dcba1"></a>

Type: `"single"`. Computed.

Port. Port of the workload.

Upstream description:

Port of the workload.

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

<a id="canonical-45ccef440b444ad17594cd1db50888379099476c8d7ef8b088a673a4d27616b9"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.port / 34da10fb04c4 / 3

- [info](data-sources--workload--reference--group-008.md#canonical-1469fbc807dfc79c795510f925d081e814a20e0543978f334f37577fa55b62c0): complete subsection reference.

<a id="canonical-7a85698b3e022d210ebb5cfa84c1a8456b637ee90408d59b5fb042863d1910e8"></a>

<a id="canonical-3b78d30b6853a91fad442aad889ee93bb63dec067836f28ea671f5cf4b2e2aa5"></a>

## name property — service.advertise_options.advertise_custom.ports.port / 34da10fb04c4 / 4

Type: `"string"`. Computed.

Name. Name of the Port.

Upstream description:

Name of the Port.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.iana_svc_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.iana_svc_name": "true"
  }
}
```

<a id="canonical-70107504f047aaac910d9ede88b93cc216b29d9b27b4b9baa3d46caec39bf76c"></a>

## Next pages — service.advertise_options.advertise_custom.ports.port / 34da10fb04c4 / 5

- [service.advertise_options.advertise_custom.ports.port.info](data-sources--workload--reference--group-008.md#canonical-1469fbc807dfc79c795510f925d081e814a20e0543978f334f37577fa55b62c0)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-d2f2640345dbdf8a6ca20b7b3a2ceb6c31007e44a4f25546827758c4a23df912)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-1469fbc807dfc79c795510f925d081e814a20e0543978f334f37577fa55b62c0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cf841c7e550d52bb7229b0fab5cc8fe722e6ab364059c2b759078f6c78f576e3"></a>

## service.advertise_options.advertise_custom.ports.port.info — service.advertise_options.advertise_custom.ports.port.info / ba846a5cdcdf / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-fe3574b4f2af2fc6d8f8000c8ef6e8baa5bd8847f507f148a99a7c6940e27c62)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-d2f2640345dbdf8a6ca20b7b3a2ceb6c31007e44a4f25546827758c4a23df912)
- [service.advertise_options.advertise_custom.ports.port](data-sources--workload--reference--group-008.md#canonical-7c613d4efb08afb76c10e5423143ef3447c3a5c0c851a6306c101db9750216aa)
- service.advertise_options.advertise_custom.ports.port.info

<a id="canonical-a1c4464445c08b648d0c19fcc064d7e413b2b14c443f8ea4196838d53d660cbc"></a>

Type: `"single"`. Computed.

Port Information. Port information.

Upstream description:

Port information.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-target_port_choice": "[\"same_as_port\",\"target_port\"]"
}
```

<a id="canonical-0cb5e4fa4f94b3b16ed70e5b615281865d423f3beccfdfc47711f41267c0bf3b"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.port.info / ba846a5cdcdf / 3

<a id="canonical-f28781fa8cf6143025c315395dd8c435996381abb4a419e9865ebd96f8f6d18d"></a>

<a id="canonical-76798a3a6953d004e5ad0ebf21624f407660bc8fed59774349727a9ec731ab55"></a>

## port property — service.advertise_options.advertise_custom.ports.port.info / ba846a5cdcdf / 4

Type: `"number"`. Computed.

Port. Port the workload can be reached on.

Upstream description:

Port the workload can be reached on.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-cbdea9a2bc415d5e12cb67b62d715d572bbc2687d711ae783c451b1dfcafed6d"></a>

<a id="canonical-cd15d5aac97020ef347bff3d95f85e66c682da615906bc185a107af7fdd169af"></a>

## protocol property — service.advertise_options.advertise_custom.ports.port.info / ba846a5cdcdf / 5

Type: `"string"`. Computed.

\[Enum: PROTOCOL\_TCP|PROTOCOL\_HTTP|PROTOCOL\_HTTP2|PROTOCOL\_TLS\_WITH\_SNI|PROTOCOL\_UDP\] Type
of protocol - PROTOCOL\_TCP: TCP TCP - PROTOCOL\_HTTP: HTTP HTTP - PROTOCOL\_HTTP2: HTTP2 HTTP2 -
PROTOCOL\_TLS\_WITH\_SNI: TLS with SNI TLS with SNI - PROTOCOL\_UDP: UDP UDP. Possible values are
\`PROTOCOL\_TCP\`, \`PROTOCOL\_HTTP\`, \`PROTOCOL\_HTTP2\`, \`PROTOCOL\_TLS\_WITH\_SNI\`,
\`PROTOCOL\_UDP\`. Defaults to \`PROTOCOL\_TCP\`.

Upstream description:

Type of protocol

&#8203;- PROTOCOL\_TCP: TCP

TCP &#8203;- PROTOCOL\_HTTP: HTTP

HTTP &#8203;- PROTOCOL\_HTTP2: HTTP2

HTTP2 &#8203;- PROTOCOL\_TLS\_WITH\_SNI: TLS with SNI

TLS with SNI &#8203;- PROTOCOL\_UDP: UDP

UDP.

Receipt-pinned upstream constraints:

```json
{
  "default": "PROTOCOL_TCP",
  "enum": [
    "PROTOCOL_TCP",
    "PROTOCOL_HTTP",
    "PROTOCOL_HTTP2",
    "PROTOCOL_TLS_WITH_SNI",
    "PROTOCOL_UDP"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [same_as_port](data-sources--workload--reference--group-008.md#canonical-6416dad6164e60438be49e73aee1113d07f90d583e1f45ba587b447098356974): complete subsection reference.

<a id="canonical-fbf8ee012642f0ac53f7b39c33858549427da50ce16efc5c528e2b83a2fb37ac"></a>

<a id="canonical-2ad79fdc93ff867eb6b23ecd83ecb4348c71987dc4a71e7d5cbe40befea07d60"></a>

## target_port property — service.advertise_options.advertise_custom.ports.port.info / ba846a5cdcdf / 6

Type: `"number"`. Computed.

Exclusive with \[same\_as\_port\] Port the workload is listening on.

Upstream description:

Exclusive with \[same\_as\_port\] Port the workload is listening on.

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
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-6c3b8523ec5f523ff54ad16eee61972651432c5414d28f93069beaaf8b4f08fd"></a>

## Next pages — service.advertise_options.advertise_custom.ports.port.info / ba846a5cdcdf / 7

- [service.advertise_options.advertise_custom.ports.port.info.same_as_port](data-sources--workload--reference--group-008.md#canonical-6416dad6164e60438be49e73aee1113d07f90d583e1f45ba587b447098356974)
- [service.advertise_options.advertise_custom.ports.port](data-sources--workload--reference--group-008.md#canonical-7c613d4efb08afb76c10e5423143ef3447c3a5c0c851a6306c101db9750216aa)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-6416dad6164e60438be49e73aee1113d07f90d583e1f45ba587b447098356974"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2e647fa1c8b7065f4e84ca5c191646088e65e4ca254f172ebf7fa5f5c21f7bfc"></a>

## service.advertise_options.advertise_custom.ports.port.info.same_as_port — service.advertise_options.advertise_custom.ports.port.info.same_as_port / 65ddb1b48de7 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-fe3574b4f2af2fc6d8f8000c8ef6e8baa5bd8847f507f148a99a7c6940e27c62)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-d2f2640345dbdf8a6ca20b7b3a2ceb6c31007e44a4f25546827758c4a23df912)
- [service.advertise_options.advertise_custom.ports.port](data-sources--workload--reference--group-008.md#canonical-7c613d4efb08afb76c10e5423143ef3447c3a5c0c851a6306c101db9750216aa)
- [service.advertise_options.advertise_custom.ports.port.info](data-sources--workload--reference--group-008.md#canonical-1469fbc807dfc79c795510f925d081e814a20e0543978f334f37577fa55b62c0)
- service.advertise_options.advertise_custom.ports.port.info.same_as_port

<a id="canonical-a4a78636f9289feb3b5fe4fa6f01ac8199c28d09028dbfc17b83f9245d8632e6"></a>

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

<a id="canonical-0cdb9377ff6802d9b6b05d8af5866036a4813390523d35ff519cbd5d53124448"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.port.info.same_as_port / 65ddb1b48de7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-06e6ff3377f633a2a6c943dd287a8e194aa8a78406b04841891a722b0acc3b3f"></a>

## Next pages — service.advertise_options.advertise_custom.ports.port.info.same_as_port / 65ddb1b48de7 / 4

- [service.advertise_options.advertise_custom.ports.port.info](data-sources--workload--reference--group-008.md#canonical-1469fbc807dfc79c795510f925d081e814a20e0543978f334f37577fa55b62c0)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-671edfd9f0410a2c5d55e0e8ca6a82232079d713e8ffed70a89baf527d62d65d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3915f8335b9e42c27ac809b4e02418d29d8021304edfa5df390e3c77de0a5d27"></a>

## service.advertise_options.advertise_custom.ports.tcp_loadbalancer — service.advertise_options.advertise_custom.ports.tcp_loadbalancer / 7f4acf3841ea / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-fe3574b4f2af2fc6d8f8000c8ef6e8baa5bd8847f507f148a99a7c6940e27c62)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-d2f2640345dbdf8a6ca20b7b3a2ceb6c31007e44a4f25546827758c4a23df912)
- service.advertise_options.advertise_custom.ports.tcp_loadbalancer

<a id="canonical-a57c6741a39ab3371d49dca3c4ebcb0d5296e4da5d5019b59bc3503979ae7409"></a>

Type: `"single"`. Computed.

Configuration parameter for tcp loadbalancer.

Upstream description:

TCP loadbalancer.

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

<a id="canonical-8160fa8a2b4ed910d3f0a8e8eefe1affdd5d79e435061e237cbba2b083de2099"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.tcp_loadbalancer / 7f4acf3841ea / 3

<a id="canonical-2ef1cdff325e4dbe120df1b612d20916b3989eea2e3c805c9b7b793b0fa009e1"></a>

<a id="canonical-b5bff061155a88dcfad1006ad706805ead5dbf79efb305b0279c6ebbaf043e5e"></a>

## domains property — service.advertise_options.advertise_custom.ports.tcp_loadbalancer / 7f4acf3841ea / 4

Type: `["list", "string"]`. Computed.

List of additional domains (host/authority header) that will be matched to this loadbalancer.
Domains are also used for SNI matching if the is true Domains also indicate the list of names for
which DNS resolution will be done by VER.

Upstream description:

A list of additional domains (host/authority header) that will be matched to this loadbalancer.

Domains are also used for SNI matching if the \`with\_sni\` is true Domains also indicate the list
of names for which DNS resolution will be done by VER.

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
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-cf61bdb8892d8353c6ae58136bbc1bd202723bb48d75eae1f8837dd761bae626"></a>

<a id="canonical-9c91a169e0da64d7422fdff03c7af10595068e3401ac45c52ecb5ccc2948847f"></a>

## with_sni property — service.advertise_options.advertise_custom.ports.tcp_loadbalancer / 7f4acf3841ea / 5

Type: `"bool"`. Computed.

Set to true to enable TCP loadbalancer with SNI.

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

<a id="canonical-67a66724ec8349ea68cab1e7368350ea462fa299893faa65ca3215969fcf1902"></a>

## Next pages — service.advertise_options.advertise_custom.ports.tcp_loadbalancer / 7f4acf3841ea / 6

- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-d2f2640345dbdf8a6ca20b7b3a2ceb6c31007e44a4f25546827758c4a23df912)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-9d22595127cd5d0e2ed99473331e76f716d729f28680fb396db73663194f15db"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-db9b3c6b3a7586242bccea602096d6fc3f435e13e9bf2c9e614ac5b2a130d174"></a>

## service.advertise_options.advertise_in_cluster — service.advertise_options.advertise_in_cluster / e4b6f881d60e / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- service.advertise_options.advertise_in_cluster

<a id="canonical-48829fd45f7cfd0af5b4afe3b00ad19cae5f94e529cfa0e6fa96361224328954"></a>

Type: `"single"`. Computed.

Advertise the workload locally in-cluster.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_choice": "[\"multi_ports\",\"port\"]"
}
```

<a id="canonical-fb07e372ca31f900ed2bcd76ba37b3b940cebbbf35dfcb2975d6e2670fe7a644"></a>

## Direct properties — service.advertise_options.advertise_in_cluster / e4b6f881d60e / 3

- [multi_ports](data-sources--workload--reference--group-008.md#canonical-3b7228e326eea6861f6be0cd35dbf0bc149a0346cce737edba94e0a003445cb5): complete subsection reference.

- [port](data-sources--workload--reference--group-008.md#canonical-ab753dfa0757cc6eb78321f6b0da20bf777027d579bb9021434710affa18439b): complete subsection reference.

<a id="canonical-b8158aa0367b9acbf20587eae8c73321f200437a00a026f59944875a49144444"></a>

## Next pages — service.advertise_options.advertise_in_cluster / e4b6f881d60e / 4

- [service.advertise_options.advertise_in_cluster.multi_ports](data-sources--workload--reference--group-008.md#canonical-3b7228e326eea6861f6be0cd35dbf0bc149a0346cce737edba94e0a003445cb5)
- [service.advertise_options.advertise_in_cluster.port](data-sources--workload--reference--group-008.md#canonical-ab753dfa0757cc6eb78321f6b0da20bf777027d579bb9021434710affa18439b)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-3b7228e326eea6861f6be0cd35dbf0bc149a0346cce737edba94e0a003445cb5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-504bbe30346e67ccb00e4cdcd61ef5c04a3129e16289874e1cfb31fffd265b28"></a>

## service.advertise_options.advertise_in_cluster.multi_ports — service.advertise_options.advertise_in_cluster.multi_ports / fe97b46eecca / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_in_cluster](data-sources--workload--reference--group-008.md#canonical-9d22595127cd5d0e2ed99473331e76f716d729f28680fb396db73663194f15db)
- service.advertise_options.advertise_in_cluster.multi_ports

<a id="canonical-99a7e5160b419caa5f44ff41873354fdd87edcf449ea279c6360b21e45415f55"></a>

Type: `"single"`. Computed.

Multiple Ports. Multiple ports.

Upstream description:

Multiple ports.

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

<a id="canonical-7af470bcd785f88114beb3488deefcbeb52c16fb818db8d734f79aa3c1e0dd30"></a>

## Direct properties — service.advertise_options.advertise_in_cluster.multi_ports / fe97b46eecca / 3

- [ports](data-sources--workload--reference--group-008.md#canonical-606cc9d2f6a469bc9a9898eed59d27edf06fd060703f716f1762dc7de45cbdfb): complete subsection reference.

<a id="canonical-04ebeffa1f36defe6102a95ef3b3c391cbdaeadc86fe92737e571f69780428d0"></a>

## Next pages — service.advertise_options.advertise_in_cluster.multi_ports / fe97b46eecca / 4

- [service.advertise_options.advertise_in_cluster.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-606cc9d2f6a469bc9a9898eed59d27edf06fd060703f716f1762dc7de45cbdfb)
- [service.advertise_options.advertise_in_cluster](data-sources--workload--reference--group-008.md#canonical-9d22595127cd5d0e2ed99473331e76f716d729f28680fb396db73663194f15db)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-606cc9d2f6a469bc9a9898eed59d27edf06fd060703f716f1762dc7de45cbdfb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8527738e5e82d1ea02e6edfb52a4dfb9291bf56c7b041b0f827f72a9b0e0c0b4"></a>

## service.advertise_options.advertise_in_cluster.multi_ports.ports — service.advertise_options.advertise_in_cluster.multi_ports.ports / 2e708d2b587a / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_in_cluster](data-sources--workload--reference--group-008.md#canonical-9d22595127cd5d0e2ed99473331e76f716d729f28680fb396db73663194f15db)
- [service.advertise_options.advertise_in_cluster.multi_ports](data-sources--workload--reference--group-008.md#canonical-3b7228e326eea6861f6be0cd35dbf0bc149a0346cce737edba94e0a003445cb5)
- service.advertise_options.advertise_in_cluster.multi_ports.ports

<a id="canonical-dcc9d5e56f77acb7e675aac7f1a369c6b54751917cfc7bb450915a0c879d907b"></a>

Type: `"list"`. Computed.

Ports. Ports to advertise.

Upstream description:

Ports to advertise.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-bc2a3f0458a0d684001d2b7b291bc0dbe5a8f9be6e8763315420a1233ee12159"></a>

## Direct properties — service.advertise_options.advertise_in_cluster.multi_ports.ports / 2e708d2b587a / 3

- [info](data-sources--workload--reference--group-008.md#canonical-751c853afedaab2c95f58599722f9954d9912782fdc282230e7bb093874ecbb5): complete subsection reference.

<a id="canonical-95aabf619818441aea8ea6c9ff75724ff1c5c871e0dde6755c083a569d87ffe9"></a>

<a id="canonical-9002ac4eda3251a12b9e8e64b83c57a11c04cb5c16e12dbaf6de458dfcaf4582"></a>

## name property — service.advertise_options.advertise_in_cluster.multi_ports.ports / 2e708d2b587a / 4

Type: `"string"`. Computed.

Name. Name of the Port.

Upstream description:

Name of the Port.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.iana_svc_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.iana_svc_name": "true"
  }
}
```

<a id="canonical-eb1e551fe13824452676337d052ebbbf05d5ce48d0977386eed1e0fac6179c54"></a>

## Next pages — service.advertise_options.advertise_in_cluster.multi_ports.ports / 2e708d2b587a / 5

- [service.advertise_options.advertise_in_cluster.multi_ports.ports.info](data-sources--workload--reference--group-008.md#canonical-751c853afedaab2c95f58599722f9954d9912782fdc282230e7bb093874ecbb5)
- [service.advertise_options.advertise_in_cluster.multi_ports](data-sources--workload--reference--group-008.md#canonical-3b7228e326eea6861f6be0cd35dbf0bc149a0346cce737edba94e0a003445cb5)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-751c853afedaab2c95f58599722f9954d9912782fdc282230e7bb093874ecbb5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-209feb0d58367853eb3542dac64d93260ef87d9ccf65d25c5f5ed21d01cd6cbb"></a>

## service.advertise_options.advertise_in_cluster.multi_ports.ports.info — service.advertise_options.advertise_in_cluster.multi_ports.ports.info / 2fd4b7086eb0 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_in_cluster](data-sources--workload--reference--group-008.md#canonical-9d22595127cd5d0e2ed99473331e76f716d729f28680fb396db73663194f15db)
- [service.advertise_options.advertise_in_cluster.multi_ports](data-sources--workload--reference--group-008.md#canonical-3b7228e326eea6861f6be0cd35dbf0bc149a0346cce737edba94e0a003445cb5)
- [service.advertise_options.advertise_in_cluster.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-606cc9d2f6a469bc9a9898eed59d27edf06fd060703f716f1762dc7de45cbdfb)
- service.advertise_options.advertise_in_cluster.multi_ports.ports.info

<a id="canonical-5bb2c54e3087524fd9ae01ace93ad5fcc1e92bd7cc39e2a956f30fe1dc1dd10b"></a>

Type: `"single"`. Computed.

Port Information. Port information.

Upstream description:

Port information.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-target_port_choice": "[\"same_as_port\",\"target_port\"]"
}
```

<a id="canonical-8dfaa1c602098296c9fd5109d7579a6eea89aefd6113db4db75a529b35e87203"></a>

## Direct properties — service.advertise_options.advertise_in_cluster.multi_ports.ports.info / 2fd4b7086eb0 / 3

<a id="canonical-c6f9db44b439773c8832ed4d0734453ff31fb89fba187d928707d2611aa08e17"></a>

<a id="canonical-b3d856c88771bbad9c2ed234fd3181af2303ff8273ad5fdf655aa0d6b02c3f37"></a>

## port property — service.advertise_options.advertise_in_cluster.multi_ports.ports.info / 2fd4b7086eb0 / 4

Type: `"number"`. Computed.

Port. Port the workload can be reached on.

Upstream description:

Port the workload can be reached on.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-6473a6f0cf3b9ca41613307c7ee73a3a3554ab67bc7c2780b8f0a1f3e9517bf5"></a>

<a id="canonical-8813f5b2cbe345b47fa515e20bbb4809986f5b9845f2c81013586c9bf33494c7"></a>

## protocol property — service.advertise_options.advertise_in_cluster.multi_ports.ports.info / 2fd4b7086eb0 / 5

Type: `"string"`. Computed.

\[Enum: PROTOCOL\_TCP|PROTOCOL\_HTTP|PROTOCOL\_HTTP2|PROTOCOL\_TLS\_WITH\_SNI|PROTOCOL\_UDP\] Type
of protocol - PROTOCOL\_TCP: TCP TCP - PROTOCOL\_HTTP: HTTP HTTP - PROTOCOL\_HTTP2: HTTP2 HTTP2 -
PROTOCOL\_TLS\_WITH\_SNI: TLS with SNI TLS with SNI - PROTOCOL\_UDP: UDP UDP. Possible values are
\`PROTOCOL\_TCP\`, \`PROTOCOL\_HTTP\`, \`PROTOCOL\_HTTP2\`, \`PROTOCOL\_TLS\_WITH\_SNI\`,
\`PROTOCOL\_UDP\`. Defaults to \`PROTOCOL\_TCP\`.

Upstream description:

Type of protocol

&#8203;- PROTOCOL\_TCP: TCP

TCP &#8203;- PROTOCOL\_HTTP: HTTP

HTTP &#8203;- PROTOCOL\_HTTP2: HTTP2

HTTP2 &#8203;- PROTOCOL\_TLS\_WITH\_SNI: TLS with SNI

TLS with SNI &#8203;- PROTOCOL\_UDP: UDP

UDP.

Receipt-pinned upstream constraints:

```json
{
  "default": "PROTOCOL_TCP",
  "enum": [
    "PROTOCOL_TCP",
    "PROTOCOL_HTTP",
    "PROTOCOL_HTTP2",
    "PROTOCOL_TLS_WITH_SNI",
    "PROTOCOL_UDP"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [same_as_port](data-sources--workload--reference--group-008.md#canonical-f6b39c4ae1d4fa0d2b93580fd4aa69c7f634cfeb35d32b1c6865155e4321b1f7): complete subsection reference.

<a id="canonical-39b7555cebe0aac0bf11f469f3eeca226f490b4c1f8646e6b666ac83ac27661a"></a>

<a id="canonical-9fcc412ff0ed2f47b52ec7ff524dac25c0ca137e2d5c1eb0a1dd91aed4b0b639"></a>

## target_port property — service.advertise_options.advertise_in_cluster.multi_ports.ports.info / 2fd4b7086eb0 / 6

Type: `"number"`. Computed.

Exclusive with \[same\_as\_port\] Port the workload is listening on.

Upstream description:

Exclusive with \[same\_as\_port\] Port the workload is listening on.

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
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-eb213cf5ab50525da27d49060cc26d9ae70c048f865af4a69908fd4454dfe311"></a>

## Next pages — service.advertise_options.advertise_in_cluster.multi_ports.ports.info / 2fd4b7086eb0 / 7

- [service.advertise_options.advertise_in_cluster.multi_ports.ports.info.same_as_port](data-sources--workload--reference--group-008.md#canonical-f6b39c4ae1d4fa0d2b93580fd4aa69c7f634cfeb35d32b1c6865155e4321b1f7)
- [service.advertise_options.advertise_in_cluster.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-606cc9d2f6a469bc9a9898eed59d27edf06fd060703f716f1762dc7de45cbdfb)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-f6b39c4ae1d4fa0d2b93580fd4aa69c7f634cfeb35d32b1c6865155e4321b1f7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-86d332cc0c54b4e5bca6f84916b848ac511bf93f1a67a71c299eda9990f516dd"></a>

## service.advertise_options.advertise_in_cluster.multi_ports.ports.info.same_as_port — service.advertise_options.advertise_in_cluster.multi_ports.ports.info.same_as_po / 9aa510fdc31e / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_in_cluster](data-sources--workload--reference--group-008.md#canonical-9d22595127cd5d0e2ed99473331e76f716d729f28680fb396db73663194f15db)
- [service.advertise_options.advertise_in_cluster.multi_ports](data-sources--workload--reference--group-008.md#canonical-3b7228e326eea6861f6be0cd35dbf0bc149a0346cce737edba94e0a003445cb5)
- [service.advertise_options.advertise_in_cluster.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-606cc9d2f6a469bc9a9898eed59d27edf06fd060703f716f1762dc7de45cbdfb)
- [service.advertise_options.advertise_in_cluster.multi_ports.ports.info](data-sources--workload--reference--group-008.md#canonical-751c853afedaab2c95f58599722f9954d9912782fdc282230e7bb093874ecbb5)
- service.advertise_options.advertise_in_cluster.multi_ports.ports.info.same_as_port

<a id="canonical-6ac7622572109b132d4cfbf442ba405eeb07b7e25dfd1cd9f59d2f17b9c19763"></a>

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

<a id="canonical-74a27f055bdf4cf1416aba8128ffe7c7eafb07a2766581ea3ef55f0236e43252"></a>

## Direct properties — service.advertise_options.advertise_in_cluster.multi_ports.ports.info.same_as_po / 9aa510fdc31e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cb4f66c8225eb6ed9d69c1a40de90e3609ac3278b17bfce9e068f8656741c3bc"></a>

## Next pages — service.advertise_options.advertise_in_cluster.multi_ports.ports.info.same_as_po / 9aa510fdc31e / 4

- [service.advertise_options.advertise_in_cluster.multi_ports.ports.info](data-sources--workload--reference--group-008.md#canonical-751c853afedaab2c95f58599722f9954d9912782fdc282230e7bb093874ecbb5)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-ab753dfa0757cc6eb78321f6b0da20bf777027d579bb9021434710affa18439b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-443d7b82cf74ff43aee318a5ddd93c7e5f1eadb7fdc03c74766de5189d08ece4"></a>

## service.advertise_options.advertise_in_cluster.port — service.advertise_options.advertise_in_cluster.port / 227b93901b3e / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_in_cluster](data-sources--workload--reference--group-008.md#canonical-9d22595127cd5d0e2ed99473331e76f716d729f28680fb396db73663194f15db)
- service.advertise_options.advertise_in_cluster.port

<a id="canonical-3a2b54b9ed01165760290e2e57ab525e1a50bdbec38f6a57d67325f425878e0b"></a>

Type: `"single"`. Computed.

Port. Single port.

Upstream description:

Single port.

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

<a id="canonical-af297c2c7c148a163c11dbf9b7b3d23647f4ab249144f6e9f378a7791fcb1a3a"></a>

## Direct properties — service.advertise_options.advertise_in_cluster.port / 227b93901b3e / 3

- [info](data-sources--workload--reference--group-008.md#canonical-508dfbd6f6a5491264f85e66edd2cc3c275652694bd85829f55bc03b2c7e4ea0): complete subsection reference.

<a id="canonical-d06f153257e7ba2ea6ba1a024975caf580fabca219dfe1f7483ff9a818041444"></a>

## Next pages — service.advertise_options.advertise_in_cluster.port / 227b93901b3e / 4

- [service.advertise_options.advertise_in_cluster.port.info](data-sources--workload--reference--group-008.md#canonical-508dfbd6f6a5491264f85e66edd2cc3c275652694bd85829f55bc03b2c7e4ea0)
- [service.advertise_options.advertise_in_cluster](data-sources--workload--reference--group-008.md#canonical-9d22595127cd5d0e2ed99473331e76f716d729f28680fb396db73663194f15db)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-508dfbd6f6a5491264f85e66edd2cc3c275652694bd85829f55bc03b2c7e4ea0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3a9289ee74397788ff39854d3d64f6db089176aa42ea2f13c8fa4e9b7760d8f6"></a>

## service.advertise_options.advertise_in_cluster.port.info — service.advertise_options.advertise_in_cluster.port.info / 04f10b7b6b29 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_in_cluster](data-sources--workload--reference--group-008.md#canonical-9d22595127cd5d0e2ed99473331e76f716d729f28680fb396db73663194f15db)
- [service.advertise_options.advertise_in_cluster.port](data-sources--workload--reference--group-008.md#canonical-ab753dfa0757cc6eb78321f6b0da20bf777027d579bb9021434710affa18439b)
- service.advertise_options.advertise_in_cluster.port.info

<a id="canonical-671470e2a26be39b7269b05a819c6e721599df283c9f25e1c5a8a9798d520843"></a>

Type: `"single"`. Computed.

Port Information. Port information.

Upstream description:

Port information.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-target_port_choice": "[\"same_as_port\",\"target_port\"]"
}
```

<a id="canonical-6e9010534b969df57f611d9a4a8138e445e7d2e7e7969e7fed2d3847536a70b0"></a>

## Direct properties — service.advertise_options.advertise_in_cluster.port.info / 04f10b7b6b29 / 3

<a id="canonical-2e1577e526632a9739f41610d6278a4118a60b8ea5fa036e66df9b034e23bf4f"></a>

<a id="canonical-784e4f0b522fb5380b1738bf107e8ab3f06d3c4f020431c595f436e6d93bfb42"></a>

## port property — service.advertise_options.advertise_in_cluster.port.info / 04f10b7b6b29 / 4

Type: `"number"`. Computed.

Port. Port the workload can be reached on.

Upstream description:

Port the workload can be reached on.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-d81148bbecbc398c7bd8a0b722cea941f03de0728c3b50be9fddfcaf36ac6f4a"></a>

<a id="canonical-7f7987265024ba27751e25b2bbea572164cb0d8e5ff9473d47a69760b2c4f38d"></a>

## protocol property — service.advertise_options.advertise_in_cluster.port.info / 04f10b7b6b29 / 5

Type: `"string"`. Computed.

\[Enum: PROTOCOL\_TCP|PROTOCOL\_HTTP|PROTOCOL\_HTTP2|PROTOCOL\_TLS\_WITH\_SNI|PROTOCOL\_UDP\] Type
of protocol - PROTOCOL\_TCP: TCP TCP - PROTOCOL\_HTTP: HTTP HTTP - PROTOCOL\_HTTP2: HTTP2 HTTP2 -
PROTOCOL\_TLS\_WITH\_SNI: TLS with SNI TLS with SNI - PROTOCOL\_UDP: UDP UDP. Possible values are
\`PROTOCOL\_TCP\`, \`PROTOCOL\_HTTP\`, \`PROTOCOL\_HTTP2\`, \`PROTOCOL\_TLS\_WITH\_SNI\`,
\`PROTOCOL\_UDP\`. Defaults to \`PROTOCOL\_TCP\`.

Upstream description:

Type of protocol

&#8203;- PROTOCOL\_TCP: TCP

TCP &#8203;- PROTOCOL\_HTTP: HTTP

HTTP &#8203;- PROTOCOL\_HTTP2: HTTP2

HTTP2 &#8203;- PROTOCOL\_TLS\_WITH\_SNI: TLS with SNI

TLS with SNI &#8203;- PROTOCOL\_UDP: UDP

UDP.

Receipt-pinned upstream constraints:

```json
{
  "default": "PROTOCOL_TCP",
  "enum": [
    "PROTOCOL_TCP",
    "PROTOCOL_HTTP",
    "PROTOCOL_HTTP2",
    "PROTOCOL_TLS_WITH_SNI",
    "PROTOCOL_UDP"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [same_as_port](data-sources--workload--reference--group-008.md#canonical-44df07f856f561260870cf5a27187354d97551cc940c733baafccb7ae8eeaf40): complete subsection reference.

<a id="canonical-64a72817101a713fda93d7bc5020531ae882fe2d0cb3d32c581caa82e74d8ca3"></a>

<a id="canonical-f05fc62e929714118f6752542589c905ba4635e20d44d1d48c0162ad5d47fd5e"></a>

## target_port property — service.advertise_options.advertise_in_cluster.port.info / 04f10b7b6b29 / 6

Type: `"number"`. Computed.

Exclusive with \[same\_as\_port\] Port the workload is listening on.

Upstream description:

Exclusive with \[same\_as\_port\] Port the workload is listening on.

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
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-9e8cbac9a0148e8d9d49566c88ec7166fbb74d1d3c34e544196541dd6961ca7e"></a>

## Next pages — service.advertise_options.advertise_in_cluster.port.info / 04f10b7b6b29 / 7

- [service.advertise_options.advertise_in_cluster.port.info.same_as_port](data-sources--workload--reference--group-008.md#canonical-44df07f856f561260870cf5a27187354d97551cc940c733baafccb7ae8eeaf40)
- [service.advertise_options.advertise_in_cluster.port](data-sources--workload--reference--group-008.md#canonical-ab753dfa0757cc6eb78321f6b0da20bf777027d579bb9021434710affa18439b)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-44df07f856f561260870cf5a27187354d97551cc940c733baafccb7ae8eeaf40"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6485dcafd87dfdd0f59c42a0cb5faa44c19f2cd0b1fd26bbd8786d9a754efe9e"></a>

## service.advertise_options.advertise_in_cluster.port.info.same_as_port — service.advertise_options.advertise_in_cluster.port.info.same_as_port / 1e620c2dfd7f / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_in_cluster](data-sources--workload--reference--group-008.md#canonical-9d22595127cd5d0e2ed99473331e76f716d729f28680fb396db73663194f15db)
- [service.advertise_options.advertise_in_cluster.port](data-sources--workload--reference--group-008.md#canonical-ab753dfa0757cc6eb78321f6b0da20bf777027d579bb9021434710affa18439b)
- [service.advertise_options.advertise_in_cluster.port.info](data-sources--workload--reference--group-008.md#canonical-508dfbd6f6a5491264f85e66edd2cc3c275652694bd85829f55bc03b2c7e4ea0)
- service.advertise_options.advertise_in_cluster.port.info.same_as_port

<a id="canonical-acffff78840cd3764ae4c60aeab46e39b7f7893a22afdc07c0019826499e8b4d"></a>

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

<a id="canonical-18a36cb2b93255a0aa11992208544549e24db6a4c22edbd0e3226a2be980ce6b"></a>

## Direct properties — service.advertise_options.advertise_in_cluster.port.info.same_as_port / 1e620c2dfd7f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3e5fecda20a4a21d391e6e84b283c559a9e1e32344bb1a27b2a5e6d606275354"></a>

## Next pages — service.advertise_options.advertise_in_cluster.port.info.same_as_port / 1e620c2dfd7f / 4

- [service.advertise_options.advertise_in_cluster.port.info](data-sources--workload--reference--group-008.md#canonical-508dfbd6f6a5491264f85e66edd2cc3c275652694bd85829f55bc03b2c7e4ea0)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-751ed9accfd600225b5c3a2437517cfdc18c1ddc5f2d7ba062b60b10acfbaa27"></a>

## service.advertise_options.advertise_on_public — service.advertise_options.advertise_on_public / b4dc7bb18fd7 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- service.advertise_options.advertise_on_public

<a id="canonical-02d105c9585ee30e59977a406ed9b959e398dd9967b4fe5aba40aaf5da13441c"></a>

Type: `"single"`. Computed.

Advertise this workload via loadbalancer on Internet with default VIP.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-advertise_choice": "[\"multi_ports\",\"port\"]"
}
```

<a id="canonical-dfc46ed8f0b3b9feaddb7a079c80da39b0f6c6d87d9d8ee8e16326394f7345e4"></a>

## Direct properties — service.advertise_options.advertise_on_public / b4dc7bb18fd7 / 3

- [multi_ports](data-sources--workload--reference--group-008.md#canonical-dd7ee64823cb466d5d5ec98c954f713c9db16ae3b5d676ba8ba400b91a109286): complete subsection reference.

- [port](data-sources--workload--reference--group-012.md#canonical-a12bf02272720e68e73b00969912e4531958e55ac3c63866089f14cf774f7a3e): complete subsection reference.

<a id="canonical-e68155207b7258126d3f38718379067a8621ba7ac7a614a7c881f7a7e38945e0"></a>

## Next pages — service.advertise_options.advertise_on_public / b4dc7bb18fd7 / 4

- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-dd7ee64823cb466d5d5ec98c954f713c9db16ae3b5d676ba8ba400b91a109286)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-a12bf02272720e68e73b00969912e4531958e55ac3c63866089f14cf774f7a3e)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-dd7ee64823cb466d5d5ec98c954f713c9db16ae3b5d676ba8ba400b91a109286"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b7ec56436ecf61317224d3101abfaaa13676f7bf3c6ffabe6fa07f6802d75e4d"></a>

## service.advertise_options.advertise_on_public.multi_ports — service.advertise_options.advertise_on_public.multi_ports / 9fac50a9b80a / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- service.advertise_options.advertise_on_public.multi_ports

<a id="canonical-5eac8c381b1970cf830b43bdfb6fe8efa3c114b4ec804730cc56c5d33fc2a663"></a>

Type: `"single"`. Computed.

Advertise Multiple Ports. Advertise multiple ports.

Upstream description:

Advertise multiple ports.

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

<a id="canonical-c75c5a01cc78469afc37732b1f5e525b7d522813714c13383959c355817ed6b4"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports / 9fac50a9b80a / 3

- [ports](data-sources--workload--reference--group-008.md#canonical-097d167f64e2b428df3091a67123e0cba76128c80bc50f5aa467387925f7ffd8): complete subsection reference.

<a id="canonical-a93860307222136c96bc263b8acfaccf389c715fd559ed46aeff0fdb4572c06e"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports / 9fac50a9b80a / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-097d167f64e2b428df3091a67123e0cba76128c80bc50f5aa467387925f7ffd8)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-097d167f64e2b428df3091a67123e0cba76128c80bc50f5aa467387925f7ffd8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b0aef087f98cd19b0b613cda56e31c2773c77a261b44e84bec3bbb58c19682a6"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports — service.advertise_options.advertise_on_public.multi_ports.ports / d5cc1c3dbf98 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-dd7ee64823cb466d5d5ec98c954f713c9db16ae3b5d676ba8ba400b91a109286)
- service.advertise_options.advertise_on_public.multi_ports.ports

<a id="canonical-f69694509e36bbb03ff27de047e690dde0ed050808dfe977266760645aa01b0f"></a>

Type: `"list"`. Computed.

Ports. Ports to advertise.

Upstream description:

Ports to advertise.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-2813b3900c273c23bd1fc19f97d9e4eb1bf9d4fc098c52ad40f5dade90b87f31"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports / d5cc1c3dbf98 / 3

- [http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-cb558922372d639b8f935627c0dc3027a3586d18d6ce937b77b1623f0a87a4b1): complete subsection reference.

- [port](data-sources--workload--reference--group-012.md#canonical-3cb34f707d2dbec6123bc41dddc1fc07d5517460301bdffdd7cdfb2fbda20733): complete subsection reference.

- [tcp_loadbalancer](data-sources--workload--reference--group-012.md#canonical-c35ed97785ea91f129e6fb702179520fb87869491603605ee0f31354741b57cc): complete subsection reference.

<a id="canonical-fcd2eecf7a403d42c06f8d0f4f485a5dd0332969641edafead4231916bb46059"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports / d5cc1c3dbf98 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-cb558922372d639b8f935627c0dc3027a3586d18d6ce937b77b1623f0a87a4b1)
- [service.advertise_options.advertise_on_public.multi_ports.ports.port](data-sources--workload--reference--group-012.md#canonical-3cb34f707d2dbec6123bc41dddc1fc07d5517460301bdffdd7cdfb2fbda20733)
- [service.advertise_options.advertise_on_public.multi_ports.ports.tcp_loadbalancer](data-sources--workload--reference--group-012.md#canonical-c35ed97785ea91f129e6fb702179520fb87869491603605ee0f31354741b57cc)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-dd7ee64823cb466d5d5ec98c954f713c9db16ae3b5d676ba8ba400b91a109286)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-cb558922372d639b8f935627c0dc3027a3586d18d6ce937b77b1623f0a87a4b1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-08f31f569d631f43bab394c507a01289081a34a264dacfac5e6670e6d8d100fa"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / cb724c9e1355 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-dd7ee64823cb466d5d5ec98c954f713c9db16ae3b5d676ba8ba400b91a109286)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-097d167f64e2b428df3091a67123e0cba76128c80bc50f5aa467387925f7ffd8)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer

<a id="canonical-fe27fff3463248c00dbfb66fda6b5266b7ba6abff4eaf0f2bd2e912ad42426aa"></a>

Type: `"single"`. Computed.

Configuration parameter for http loadbalancer.

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
  "x-ves-oneof-field-loadbalancer_type": "[\"http\",\"https\",\"https_auto_cert\"]",
  "x-ves-oneof-field-route_choice": "[\"default_route\",\"specific_routes\"]"
}
```

<a id="canonical-bf9222fd1b6e01eb7ab3a7c5a1bd50dc4ae9e402558ec3c2deaf45184b463518"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / cb724c9e1355 / 3

- [default_route](data-sources--workload--reference--group-008.md#canonical-8a013f734db838b60bb754f8b4400742c1fa325ad9136f756530d2ce762cca50): complete subsection reference.

<a id="canonical-8e2e41c68ac0122aee883245b715c537e9eba2bf44e74455173b5e432f2841ce"></a>

<a id="canonical-aabd5cd573058bcb6020a5c7db93002fcd066243f9ca017b31f9fe24afa5f1d2"></a>

## domains property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / cb724c9e1355 / 4

Type: `["list", "string"]`. Computed.

List of domains (host/authority header) that will be matched to loadbalancer. Wildcard hosts are
supported in the suffix or prefix form Domain search order: 1. Exact domain names: \`\` is invalid
Domains are also used for SNI matching if the loadbalancer type is HTTPS Domains also indicate the..

Upstream description:

A list of domains (host/authority header) that will be matched to loadbalancer. Wildcard hosts are
supported in the suffix or prefix form

Domain search order: &#8203;1. Exact domain names: \`\`www&#46;example.com\`\`. &#8203;2. Prefix
domain wildcards: \`\`\*.example.com\`\` or \`\`\*.bar.example.com\`\`. &#8203;3. Special wildcard
\`\`\*\`\` matching any domain.

Wildcard will not match empty string. E.g. \`\`\*.example.com\`\` will match \`\`bar.example.com\`\`
and \`\`baz-bar.example.com\`\` but not \`\`.example.com\`\`. The longest wildcards match first.
Wildcards must match a whole DNS label. E.g. \`\`\*.example.com\`\` and \*.bar.example.com are
valid, however \`\`\*bar.example.com\`\` or \`\`\*-bar.example.com\`\` is invalid

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

- [http](data-sources--workload--reference--group-008.md#canonical-45600eb7238b372a39221db2939141c5b76227188be6c31728d84de9918dc6ba): complete subsection reference.

- [https](data-sources--workload--reference--group-008.md#canonical-6ae7f2b71d6a5a98e32ded4d2b787278e735d1f3c72f9f2d0a5affc2d4f91758): complete subsection reference.

- [https_auto_cert](data-sources--workload--reference--group-010.md#canonical-3c6fe7acbb3e7d43101bb33001feeb507501240dfc6045bf42be34bbd5c362a4): complete subsection reference.

- [specific_routes](data-sources--workload--reference--group-011.md#canonical-30327ed3011b3335416b21fd86305c3cd5b23a9eb80a4d5c2ccc09979303187b): complete subsection reference.

<a id="canonical-51dd44c05a36d4ad0c7fc44357e377020adecd2a7cc9ceef1317c00168d69fff"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / cb724c9e1355 / 5

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route](data-sources--workload--reference--group-008.md#canonical-8a013f734db838b60bb754f8b4400742c1fa325ad9136f756530d2ce762cca50)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.http](data-sources--workload--reference--group-008.md#canonical-45600eb7238b372a39221db2939141c5b76227188be6c31728d84de9918dc6ba)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-6ae7f2b71d6a5a98e32ded4d2b787278e735d1f3c72f9f2d0a5affc2d4f91758)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-010.md#canonical-3c6fe7acbb3e7d43101bb33001feeb507501240dfc6045bf42be34bbd5c362a4)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-011.md#canonical-30327ed3011b3335416b21fd86305c3cd5b23a9eb80a4d5c2ccc09979303187b)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-097d167f64e2b428df3091a67123e0cba76128c80bc50f5aa467387925f7ffd8)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-8a013f734db838b60bb754f8b4400742c1fa325ad9136f756530d2ce762cca50"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-37861fb069dd4cb531f82a0b3f2caff25bbd78b405e147395321210610077c22"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 0066c2970710 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-dd7ee64823cb466d5d5ec98c954f713c9db16ae3b5d676ba8ba400b91a109286)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-097d167f64e2b428df3091a67123e0cba76128c80bc50f5aa467387925f7ffd8)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-cb558922372d639b8f935627c0dc3027a3586d18d6ce937b77b1623f0a87a4b1)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route

<a id="canonical-48cea4e4effa466c15e59dbf69150e3ff6973e623db965442e9620efbf5ab567"></a>

Type: `"single"`. Computed.

Configuration parameter for default route.

Upstream description:

Default route matching all APIs.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-host_rewrite_params": "[\"auto_host_rewrite\",\"disable_host_rewrite\",\"host_rewrite\"]"
}
```

<a id="canonical-6b835b1f94e2da0999e977bada383f8fb460c2bc9e45bc1d771cdae3eacc58f5"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 0066c2970710 / 3

- [auto_host_rewrite](data-sources--workload--reference--group-008.md#canonical-94244114990cf4af0f96c9013717cc86e14cf95fff7f1000cd40dbf21fcd2321): complete subsection reference.

- [disable_host_rewrite](data-sources--workload--reference--group-008.md#canonical-155a0f81c5480c9f1c280323d2e7447e24622935faf8c1b4c7cd14e3cd6139c7): complete subsection reference.

<a id="canonical-2da9ace63a3e5cef7dddd6a0c527945695edbaf6069982947251c6db6a0835d7"></a>

<a id="canonical-fb6254e64367c064fbdf79839c41003738984c7ac71d280ac3f2198d2313eae3"></a>

## host_rewrite property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 0066c2970710 / 4

Type: `"string"`. Computed.

Exclusive with \[auto\_host\_rewrite disable\_host\_rewrite\] Host header will be swapped with this
value.

Upstream description:

Exclusive with \[auto\_host\_rewrite disable\_host\_rewrite\] Host header will be swapped with this
value.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-3406006fc2f8f85a8d8d1c42a586faacd42d87c2c8608a570f85b3a42b9e0053"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 0066c2970710 / 5

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.auto_host_rewrite](data-sources--workload--reference--group-008.md#canonical-94244114990cf4af0f96c9013717cc86e14cf95fff7f1000cd40dbf21fcd2321)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.disable_host_rewrite](data-sources--workload--reference--group-008.md#canonical-155a0f81c5480c9f1c280323d2e7447e24622935faf8c1b4c7cd14e3cd6139c7)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-cb558922372d639b8f935627c0dc3027a3586d18d6ce937b77b1623f0a87a4b1)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-94244114990cf4af0f96c9013717cc86e14cf95fff7f1000cd40dbf21fcd2321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-62bc95e0b468ea64d11276b598f0c768fd081ccc224ea248594b7e41122afe62"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.auto_host_rewrite — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / c2c0b4fb483b / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-dd7ee64823cb466d5d5ec98c954f713c9db16ae3b5d676ba8ba400b91a109286)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-097d167f64e2b428df3091a67123e0cba76128c80bc50f5aa467387925f7ffd8)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-cb558922372d639b8f935627c0dc3027a3586d18d6ce937b77b1623f0a87a4b1)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route](data-sources--workload--reference--group-008.md#canonical-8a013f734db838b60bb754f8b4400742c1fa325ad9136f756530d2ce762cca50)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.auto_host_rewrite

<a id="canonical-241961ff54ffe3463f614443016f1ae5e0a4fffbe6cc3c8d9540d0b8d51dcd3d"></a>

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

<a id="canonical-f34b398bbc839311f0bd5e8209933dcf8859be548c897ca2278dcdb793d5e7d3"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / c2c0b4fb483b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-41bff2f1d19035466390e927e8afbf639cf21eb848543cf8269d075154bd3560"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / c2c0b4fb483b / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route](data-sources--workload--reference--group-008.md#canonical-8a013f734db838b60bb754f8b4400742c1fa325ad9136f756530d2ce762cca50)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-155a0f81c5480c9f1c280323d2e7447e24622935faf8c1b4c7cd14e3cd6139c7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-88e9c3b21af8cc3a24c2192244717bf5cd2d9d13aa3a70113bc4eb15c6ecf2ac"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.disable_host_rewrite — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / e83b3b31634e / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-dd7ee64823cb466d5d5ec98c954f713c9db16ae3b5d676ba8ba400b91a109286)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-097d167f64e2b428df3091a67123e0cba76128c80bc50f5aa467387925f7ffd8)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-cb558922372d639b8f935627c0dc3027a3586d18d6ce937b77b1623f0a87a4b1)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route](data-sources--workload--reference--group-008.md#canonical-8a013f734db838b60bb754f8b4400742c1fa325ad9136f756530d2ce762cca50)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.disable_host_rewrite

<a id="canonical-226b655e5bdce79798e985cdfc3e8da3d7e59fc68df4460aba34f445db8d9cc4"></a>

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

<a id="canonical-2093dd1728f5f3a0d3c877502330de94d4bfb1a53579c442ab0ea6a930394425"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / e83b3b31634e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f4139f4d0c00fc70bc43be2c604247ad405fcbea09b7874f960089b249542ef9"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / e83b3b31634e / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route](data-sources--workload--reference--group-008.md#canonical-8a013f734db838b60bb754f8b4400742c1fa325ad9136f756530d2ce762cca50)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-45600eb7238b372a39221db2939141c5b76227188be6c31728d84de9918dc6ba"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8f4c04cbc58d8d93e4e74d63fc881e30ef177da4bb39a76567f7a461bb22788a"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.http — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 96ce45f89017 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-dd7ee64823cb466d5d5ec98c954f713c9db16ae3b5d676ba8ba400b91a109286)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-097d167f64e2b428df3091a67123e0cba76128c80bc50f5aa467387925f7ffd8)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-cb558922372d639b8f935627c0dc3027a3586d18d6ce937b77b1623f0a87a4b1)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.http

<a id="canonical-c4434f7b858b37be3cd49ce8aac3d92eb08c6c34c4c15115a42963bea86b8e07"></a>

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

<a id="canonical-2d7c20126bda5180f972040c0172a28ce4509128b5a516e41ed7c4c447df513d"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 96ce45f89017 / 3

<a id="canonical-b232fc4e3c7ce5376e1b4e587385ff1c78f41b8051976e2edc23a1ec89c66f0c"></a>

<a id="canonical-3774ff9f985bc5debefb54fa72838bcdf4f76bcb4725686392f6c3172dfc3bd6"></a>

## dns_volterra_managed property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 96ce45f89017 / 4

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

<a id="canonical-469caf3d93d992e4dd24e14125a73616dfd386b689c3717c391e611669d84ced"></a>

<a id="canonical-9ae9eafd7840d1e6a472fc27078610f1112aeef6a8ff3b76304129be1407f19a"></a>

## port property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 96ce45f89017 / 5

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

<a id="canonical-0d463038f71231113aa63b884455a20ae83816424b4e6d1a79be3b693582077d"></a>

<a id="canonical-a203de81cea638bd8c2f2c3260bae948a212e2f9b9f6925353968b0b4438f802"></a>

## port_ranges property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 96ce45f89017 / 6

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

<a id="canonical-b869196877af18f57c3c48cff300a2bdc3f3bf20481067192c5fbe906af99269"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 96ce45f89017 / 7

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-cb558922372d639b8f935627c0dc3027a3586d18d6ce937b77b1623f0a87a4b1)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-6ae7f2b71d6a5a98e32ded4d2b787278e735d1f3c72f9f2d0a5affc2d4f91758"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-efffe083382b39d890468181287b0ec27b072bbc9d05aa4d84389316772661c3"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 2b318df32d67 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-dd7ee64823cb466d5d5ec98c954f713c9db16ae3b5d676ba8ba400b91a109286)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-097d167f64e2b428df3091a67123e0cba76128c80bc50f5aa467387925f7ffd8)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-cb558922372d639b8f935627c0dc3027a3586d18d6ce937b77b1623f0a87a4b1)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https

<a id="canonical-6aa5736598525f3b8c3e4ba5a46da71cd5664c07045106b111bdfe1e9f95243c"></a>

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

<a id="canonical-50a6846da4cc8c4971e53d425450d4a39550690f1002b5f6c83885d22090a029"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 2b318df32d67 / 3

<a id="canonical-2919c0100ef6afbea06be4be046d855e19b04de5826dac665f86fa9a9cad7062"></a>

<a id="canonical-784e82c10d8a0e64d16751c33383c724dcc5cda704a00dd6a329e3985a31fcf1"></a>

## add_hsts property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 2b318df32d67 / 4

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

<a id="canonical-a320f058d16820cdb84e3607c5144be79f71c68fd9a28fe12079d84e6ab5e9a5"></a>

<a id="canonical-1bb751237fb163001fed3c8252c980ec459ae389653b469f44ec3665009bae3d"></a>

## append_server_name property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 2b318df32d67 / 5

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

- [coalescing_options](data-sources--workload--reference--group-008.md#canonical-493f98289a6ea117fc1214ef9322833c89818d34723ce2ca794f3400e622cd87): complete subsection reference.

<a id="canonical-0b952290a511ac1d89176ea96fb374f45e1267489973cad12ee72fde7dd27fdd"></a>

<a id="canonical-4ae0cff30cdd2ee38b9dfebd41512e774da3c55a88d968a20bb0e37670c1f9b3"></a>

## connection_idle_timeout property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 2b318df32d67 / 6

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

- [default_header](data-sources--workload--reference--group-008.md#canonical-7457085cf0da0ee561b8ff93ea1d126b21698ce974ee4acc6fc30185432857d8): complete subsection reference.

- [default_loadbalancer](data-sources--workload--reference--group-009.md#canonical-c17febba60e59e71e0c97c21de0dbad4e533bab108b791d79bacb74bcba01eec): complete subsection reference.

- [disable_path_normalize](data-sources--workload--reference--group-009.md#canonical-e24f56068cfbc79e78760d6b5d18999dcdaffb8812860363cdee839f858f145e): complete subsection reference.

- [enable_path_normalize](data-sources--workload--reference--group-009.md#canonical-32a942bf61be073e001e64217411263bf6b6e632da520c55e627d358303d51c8): complete subsection reference.

- [http_protocol_options](data-sources--workload--reference--group-009.md#canonical-6800266df44556232a0ae5fb5fb8ae633a0e9c70f834cb3e8ac413369e6ae809): complete subsection reference.

<a id="canonical-f1b37a50d02dbf53532f9a099602b9923b0da18eae3c00c88607d128c5395b4c"></a>

<a id="canonical-3d0f7be5832c0c8c46965aff96c41e1b42b202eb38a12748d98ff03f2c26dd05"></a>

## http_redirect property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 2b318df32d67 / 7

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

- [non_default_loadbalancer](data-sources--workload--reference--group-009.md#canonical-47ba5e756987d3ffa1322289780f5bdf28b35d3daeea7da8721b009cf49d8c2f): complete subsection reference.

- [pass_through](data-sources--workload--reference--group-009.md#canonical-afd0c2ea1dc2dbd7ccb915a076e791d2d67777db7a8e37bfa4f695051f467a6c): complete subsection reference.

<a id="canonical-46f0e3ee529424a0da2386de27ad501f8c5e294d7527bcdecdd07d04682856eb"></a>

<a id="canonical-1f9064cd19f32c04c6873670e2d66c168d74e1f3f98d261fbb712bac0c892585"></a>

## port property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 2b318df32d67 / 8

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

<a id="canonical-bd0be61a0a05a3056670d5f9d649fca701f3a3dcba234072de5b8c5aa54a08fe"></a>

<a id="canonical-dc23e0a3c2146294b47ab05a9e1732f4cef70581047acdf6c290bff2e7f371bf"></a>

## port_ranges property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 2b318df32d67 / 9

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

<a id="canonical-ef6ad8215d85b81e99c9cdb3b96b9543e719402b2d6f443b112d6db5f4a6ff29"></a>

<a id="canonical-f8b416ed9eed2f91bd7017d0af1ec26e060673ff794029f8768d9074d7fe16d0"></a>

## server_name property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 2b318df32d67 / 10

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

- [tls_cert_params](data-sources--workload--reference--group-009.md#canonical-fa55d26d81706f6376fab9a26b13baf7022243e71883ee06dc2c093196e44083): complete subsection reference.

- [tls_parameters](data-sources--workload--reference--group-009.md#canonical-ccb6358bafbe4fd33ea49ce251d3f0473136052ee3906207ff885267e7a59217): complete subsection reference.

<a id="canonical-75e56f7737e1cc1972b53c4291e4f55c90ce13990db19614e103cd145eb1007f"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 2b318df32d67 / 11

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options](data-sources--workload--reference--group-008.md#canonical-493f98289a6ea117fc1214ef9322833c89818d34723ce2ca794f3400e622cd87)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.default_header](data-sources--workload--reference--group-008.md#canonical-7457085cf0da0ee561b8ff93ea1d126b21698ce974ee4acc6fc30185432857d8)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.default_loadbalancer](data-sources--workload--reference--group-009.md#canonical-c17febba60e59e71e0c97c21de0dbad4e533bab108b791d79bacb74bcba01eec)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.disable_path_normalize](data-sources--workload--reference--group-009.md#canonical-e24f56068cfbc79e78760d6b5d18999dcdaffb8812860363cdee839f858f145e)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.enable_path_normalize](data-sources--workload--reference--group-009.md#canonical-32a942bf61be073e001e64217411263bf6b6e632da520c55e627d358303d51c8)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-009.md#canonical-6800266df44556232a0ae5fb5fb8ae633a0e9c70f834cb3e8ac413369e6ae809)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.non_default_loadbalancer](data-sources--workload--reference--group-009.md#canonical-47ba5e756987d3ffa1322289780f5bdf28b35d3daeea7da8721b009cf49d8c2f)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.pass_through](data-sources--workload--reference--group-009.md#canonical-afd0c2ea1dc2dbd7ccb915a076e791d2d67777db7a8e37bfa4f695051f467a6c)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-009.md#canonical-fa55d26d81706f6376fab9a26b13baf7022243e71883ee06dc2c093196e44083)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-009.md#canonical-ccb6358bafbe4fd33ea49ce251d3f0473136052ee3906207ff885267e7a59217)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-cb558922372d639b8f935627c0dc3027a3586d18d6ce937b77b1623f0a87a4b1)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-493f98289a6ea117fc1214ef9322833c89818d34723ce2ca794f3400e622cd87"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-40101d2b7ae928ee4d4ee5a19c93a0fc10f3d89dc99d6cea659f04dbc8bc0489"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / aabee42da0f4 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-dd7ee64823cb466d5d5ec98c954f713c9db16ae3b5d676ba8ba400b91a109286)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-097d167f64e2b428df3091a67123e0cba76128c80bc50f5aa467387925f7ffd8)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-cb558922372d639b8f935627c0dc3027a3586d18d6ce937b77b1623f0a87a4b1)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-6ae7f2b71d6a5a98e32ded4d2b787278e735d1f3c72f9f2d0a5affc2d4f91758)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options

<a id="canonical-4677a0e27e2be322ce3d0e768ae7d7e27a36336ec98dddf7ecb0f436937ce7dc"></a>

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

<a id="canonical-a414bf8d4548e15d497efe6c492aa5cfba8150f2e14203a9c3cf1da189db8c1e"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / aabee42da0f4 / 3

- [default_coalescing](data-sources--workload--reference--group-008.md#canonical-e06e2d999c96a9c18c3c5016ff86b80bc415364d3bb56b81c8828fd6fcf6cd40): complete subsection reference.

- [strict_coalescing](data-sources--workload--reference--group-008.md#canonical-5a5b3ee3cc158302f6e2204cc1ae98f16a7fa4b09d7316a2d2721e6f665a7c24): complete subsection reference.

<a id="canonical-5889c6c64e96796cbdd5fbcd1f8d696b9fde403cb8530a7589b53567894fd2e9"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / aabee42da0f4 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options.default_coalescing](data-sources--workload--reference--group-008.md#canonical-e06e2d999c96a9c18c3c5016ff86b80bc415364d3bb56b81c8828fd6fcf6cd40)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options.strict_coalescing](data-sources--workload--reference--group-008.md#canonical-5a5b3ee3cc158302f6e2204cc1ae98f16a7fa4b09d7316a2d2721e6f665a7c24)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-6ae7f2b71d6a5a98e32ded4d2b787278e735d1f3c72f9f2d0a5affc2d4f91758)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-e06e2d999c96a9c18c3c5016ff86b80bc415364d3bb56b81c8828fd6fcf6cd40"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9538345fd57ed8c805ad97a0e7ad97102795763fe5f5a166dc71694db6151707"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options.default_coalescing — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 300164db9a0d / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-dd7ee64823cb466d5d5ec98c954f713c9db16ae3b5d676ba8ba400b91a109286)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-097d167f64e2b428df3091a67123e0cba76128c80bc50f5aa467387925f7ffd8)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-cb558922372d639b8f935627c0dc3027a3586d18d6ce937b77b1623f0a87a4b1)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-6ae7f2b71d6a5a98e32ded4d2b787278e735d1f3c72f9f2d0a5affc2d4f91758)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options](data-sources--workload--reference--group-008.md#canonical-493f98289a6ea117fc1214ef9322833c89818d34723ce2ca794f3400e622cd87)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options.default_coalescing

<a id="canonical-a27e3c9baef8b4e0e1d873a7bedef0a3c35278a1a63ee1738566916a76594611"></a>

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

<a id="canonical-07e79454f1793dd09000b76ad73331ba4b31c977c598aad29804dd44372f4061"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 300164db9a0d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-82aff7c4d614db2541768d3f8c9d706397190d1c089f76388d5e520630d49fd4"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 300164db9a0d / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options](data-sources--workload--reference--group-008.md#canonical-493f98289a6ea117fc1214ef9322833c89818d34723ce2ca794f3400e622cd87)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-5a5b3ee3cc158302f6e2204cc1ae98f16a7fa4b09d7316a2d2721e6f665a7c24"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6bf76b2397e7d65786c8390abe35486ffcfb0fa073627247cd87642b8b1c455b"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options.strict_coalescing — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 5ae822841cd4 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-dd7ee64823cb466d5d5ec98c954f713c9db16ae3b5d676ba8ba400b91a109286)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-097d167f64e2b428df3091a67123e0cba76128c80bc50f5aa467387925f7ffd8)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-cb558922372d639b8f935627c0dc3027a3586d18d6ce937b77b1623f0a87a4b1)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-6ae7f2b71d6a5a98e32ded4d2b787278e735d1f3c72f9f2d0a5affc2d4f91758)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options](data-sources--workload--reference--group-008.md#canonical-493f98289a6ea117fc1214ef9322833c89818d34723ce2ca794f3400e622cd87)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options.strict_coalescing

<a id="canonical-885f9d89e9a35939306fdcb6aca890859e78c48c20328a0e45af9fa4d45b07af"></a>

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

<a id="canonical-8b5881885fd49fe1fb2bf2bd4d252060b32f84cf5b4b467e10b9cb7636b51e4b"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 5ae822841cd4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-39e7e2ae21e32171b3e699f016c9a18ad5cdb77af34bb1c3aa58d8156494f0a5"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 5ae822841cd4 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options](data-sources--workload--reference--group-008.md#canonical-493f98289a6ea117fc1214ef9322833c89818d34723ce2ca794f3400e622cd87)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-7457085cf0da0ee561b8ff93ea1d126b21698ce974ee4acc6fc30185432857d8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
