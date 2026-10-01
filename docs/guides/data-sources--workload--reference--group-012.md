---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-3653edbe040688307af078d565c4f798dcea8980bf7babb3083e32b97a080585"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / e4d23bd87a42 / 7

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route](data-sources--workload--reference--group-011.md#canonical-9ca52f5a46e1a8d2d087fbf29e79b67bc2ab87fc779ec2a4d25e1bfa72fce73d)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-3cb34f707d2dbec6123bc41dddc1fc07d5517460301bdffdd7cdfb2fbda20733"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4b2a19809096648831504646f925c8a626da81054319d76770961cd1d7ed1ac2"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.port — service.advertise_options.advertise_on_public.multi_ports.ports.port / c0ddfdacf5db / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-dd7ee64823cb466d5d5ec98c954f713c9db16ae3b5d676ba8ba400b91a109286)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-097d167f64e2b428df3091a67123e0cba76128c80bc50f5aa467387925f7ffd8)
- service.advertise_options.advertise_on_public.multi_ports.ports.port

<a id="canonical-2cd405316e36c7cd75ff44430a9cba2669a8768e2c6b76d93535922d7164b53e"></a>

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

<a id="canonical-4d93512fe08f40e3f09f48f948460426a3f1492f696d40c425cfe85f13185e04"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.port / c0ddfdacf5db / 3

- [info](data-sources--workload--reference--group-012.md#canonical-aa198edb01662ff2ef72b24a70f73b515e61669e539de8b0861fdd35322d9f31): complete subsection reference.

<a id="canonical-18e8e21486a3a392527fa7854a791a71bd1a74fadba243870c9037da875d1fc6"></a>

<a id="canonical-e9e7d1544916573a2f528058a429825d626ecb345a3b3e09d2d25d39b9f3368b"></a>

## name property — service.advertise_options.advertise_on_public.multi_ports.ports.port / c0ddfdacf5db / 4

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

<a id="canonical-5815e56374b05d6f15b3c655cde485acd9b3654fd6935308c486d8053e966c26"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.port / c0ddfdacf5db / 5

- [service.advertise_options.advertise_on_public.multi_ports.ports.port.info](data-sources--workload--reference--group-012.md#canonical-aa198edb01662ff2ef72b24a70f73b515e61669e539de8b0861fdd35322d9f31)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-097d167f64e2b428df3091a67123e0cba76128c80bc50f5aa467387925f7ffd8)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-aa198edb01662ff2ef72b24a70f73b515e61669e539de8b0861fdd35322d9f31"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4e061cdc0295fc86f924b8cd6f2634d746f19e143f00120e71cb514dee9d063f"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.port.info — service.advertise_options.advertise_on_public.multi_ports.ports.port.info / c3f81baa72d3 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-dd7ee64823cb466d5d5ec98c954f713c9db16ae3b5d676ba8ba400b91a109286)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-097d167f64e2b428df3091a67123e0cba76128c80bc50f5aa467387925f7ffd8)
- [service.advertise_options.advertise_on_public.multi_ports.ports.port](data-sources--workload--reference--group-012.md#canonical-3cb34f707d2dbec6123bc41dddc1fc07d5517460301bdffdd7cdfb2fbda20733)
- service.advertise_options.advertise_on_public.multi_ports.ports.port.info

<a id="canonical-fa18e3e31a4dac42a4b1140aeccbbca8a22b54bf85918b33ddc2463d1416770b"></a>

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

<a id="canonical-cf328c9701d9589e52274bf6c5017c0c09aa9fe19b3232d8fa46beaf2b174b1c"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.port.info / c3f81baa72d3 / 3

<a id="canonical-73a28279ca19a35572189657d3a7ef953b96a119faf03292ac242e38d853197e"></a>

<a id="canonical-08dc3aa3ec29bee225f9710f03c97313c472cd783efa2418c6cc279494802fcc"></a>

## port property — service.advertise_options.advertise_on_public.multi_ports.ports.port.info / c3f81baa72d3 / 4

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

<a id="canonical-69e716639276eb51fe91bfac89d50efe22819cc7c0f3a6b21da3d56f314d4701"></a>

<a id="canonical-7ede8f13bd1048ecbbadea598216d9884b119ad0f51dccf157ae045055a5eb2d"></a>

## protocol property — service.advertise_options.advertise_on_public.multi_ports.ports.port.info / c3f81baa72d3 / 5

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

- [same_as_port](data-sources--workload--reference--group-012.md#canonical-a58df315fe8b9488fb360ad9f0bd34c51e60f0990f6f9d30974c92dcb410df0e): complete subsection reference.

<a id="canonical-0a626e414c67186f0f3250c122cb6cd257c875b098b93a2152cc3facb4d9003c"></a>

<a id="canonical-c09a109c59b5715a2c7c01758fe8c32613a4e1ea0caa188deebb76aeb53d88db"></a>

## target_port property — service.advertise_options.advertise_on_public.multi_ports.ports.port.info / c3f81baa72d3 / 6

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

<a id="canonical-0b43f8a2ad7c186cd859b5aabe71b5752f450dddde4358f571a9ad2eaf61c679"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.port.info / c3f81baa72d3 / 7

- [service.advertise_options.advertise_on_public.multi_ports.ports.port.info.same_as_port](data-sources--workload--reference--group-012.md#canonical-a58df315fe8b9488fb360ad9f0bd34c51e60f0990f6f9d30974c92dcb410df0e)
- [service.advertise_options.advertise_on_public.multi_ports.ports.port](data-sources--workload--reference--group-012.md#canonical-3cb34f707d2dbec6123bc41dddc1fc07d5517460301bdffdd7cdfb2fbda20733)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-a58df315fe8b9488fb360ad9f0bd34c51e60f0990f6f9d30974c92dcb410df0e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a4bd88026da099806a4de70ecbf01fa460bec13b29347832ca6c88eb97d519ba"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.port.info.same_as_port — service.advertise_options.advertise_on_public.multi_ports.ports.port.info.same_a / e7817d5180aa / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-dd7ee64823cb466d5d5ec98c954f713c9db16ae3b5d676ba8ba400b91a109286)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-097d167f64e2b428df3091a67123e0cba76128c80bc50f5aa467387925f7ffd8)
- [service.advertise_options.advertise_on_public.multi_ports.ports.port](data-sources--workload--reference--group-012.md#canonical-3cb34f707d2dbec6123bc41dddc1fc07d5517460301bdffdd7cdfb2fbda20733)
- [service.advertise_options.advertise_on_public.multi_ports.ports.port.info](data-sources--workload--reference--group-012.md#canonical-aa198edb01662ff2ef72b24a70f73b515e61669e539de8b0861fdd35322d9f31)
- service.advertise_options.advertise_on_public.multi_ports.ports.port.info.same_as_port

<a id="canonical-7cb9308f10cdf217273c2413536a12dd91fdd189cb87a5cf4519a3f484a3de17"></a>

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

<a id="canonical-83f65d9bbfbb534754f3ed9ef83d9e0e88c6c370a9da9092d60574f680b59f3d"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.port.info.same_a / e7817d5180aa / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0ec35e665c645f6fb230cecd92929147cad52a9da6052129863100a9515b7735"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.port.info.same_a / e7817d5180aa / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.port.info](data-sources--workload--reference--group-012.md#canonical-aa198edb01662ff2ef72b24a70f73b515e61669e539de8b0861fdd35322d9f31)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-c35ed97785ea91f129e6fb702179520fb87869491603605ee0f31354741b57cc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dac825146bd2af6cde717dbe9f7541bef4b142a7118ee6d428691bfb28605942"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.tcp_loadbalancer — service.advertise_options.advertise_on_public.multi_ports.ports.tcp_loadbalancer / d86c78e2e171 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-dd7ee64823cb466d5d5ec98c954f713c9db16ae3b5d676ba8ba400b91a109286)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-097d167f64e2b428df3091a67123e0cba76128c80bc50f5aa467387925f7ffd8)
- service.advertise_options.advertise_on_public.multi_ports.ports.tcp_loadbalancer

<a id="canonical-63be8994b50ecc5f2e9ef73351be34a8d65d2003ef3e9f981d23b6abf3bfa5ed"></a>

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

<a id="canonical-dbd54a0e785bca77ce0d45cd78cbbbe12efc4f48ffbc8d2e96a036e469bbe717"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.tcp_loadbalancer / d86c78e2e171 / 3

<a id="canonical-9bc3a1990bbedbe1f28163a08304d0f896b2e8eae6a0e5c4ce08af554c34557a"></a>

<a id="canonical-275194db765718e30f58e6d0478bbad15908f8519a7291773b7a14ad0de7f1c2"></a>

## domains property — service.advertise_options.advertise_on_public.multi_ports.ports.tcp_loadbalancer / d86c78e2e171 / 4

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

<a id="canonical-3ba246aa3fd00bea10d1e2213422a8bc848a7a9eeab4a5717de69d08d0816250"></a>

<a id="canonical-980c25f9f049c4b470ad0896b39de57b96958a73776541e8a01faf1dc473a5e8"></a>

## with_sni property — service.advertise_options.advertise_on_public.multi_ports.ports.tcp_loadbalancer / d86c78e2e171 / 5

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

<a id="canonical-921b0a03344e2c76517ab4c8a2927fccd41279fbdcfc2d4841c5dc13ab390579"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.tcp_loadbalancer / d86c78e2e171 / 6

- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-097d167f64e2b428df3091a67123e0cba76128c80bc50f5aa467387925f7ffd8)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-a12bf02272720e68e73b00969912e4531958e55ac3c63866089f14cf774f7a3e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-98e229a7b08bafe2c33448d78d2d4c8d2eed1f53abfdece447a7b9f49b629d91"></a>

## service.advertise_options.advertise_on_public.port — service.advertise_options.advertise_on_public.port / f49577ee84ce / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- service.advertise_options.advertise_on_public.port

<a id="canonical-530ab5e4453df5c9044843d5d776d9d91d6834eb9dc2b1f6352215b61c09c6fc"></a>

Type: `"single"`. Computed.

Advertise Port. Advertise single port.

Upstream description:

Advertise single port.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-advertise_choice": "[\"http_loadbalancer\",\"tcp_loadbalancer\"]"
}
```

<a id="canonical-fcbdf67d38df8d0bc75406a59de2dfe7109f659aca6104e39068e526586a0eca"></a>

## Direct properties — service.advertise_options.advertise_on_public.port / f49577ee84ce / 3

- [http_loadbalancer](data-sources--workload--reference--group-012.md#canonical-6111d7521abb2c61754e5c7b0229218264c875772bcdbb8788031967782f42e0): complete subsection reference.

- [port](data-sources--workload--reference--group-015.md#canonical-17050b4cdfcb149ca58593a78653468f7c97221f865c86f032f71bb4f5f3cc2b): complete subsection reference.

- [tcp_loadbalancer](data-sources--workload--reference--group-015.md#canonical-0e62a645135c3b8ccf8743b299e24a9aa87050bc8ebdb2324540646b774daebe): complete subsection reference.

<a id="canonical-7e8d72439839cb679764f0e84bd483a6927feb8702761adbc29ddefb1ad3537a"></a>

## Next pages — service.advertise_options.advertise_on_public.port / f49577ee84ce / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-012.md#canonical-6111d7521abb2c61754e5c7b0229218264c875772bcdbb8788031967782f42e0)
- [service.advertise_options.advertise_on_public.port.port](data-sources--workload--reference--group-015.md#canonical-17050b4cdfcb149ca58593a78653468f7c97221f865c86f032f71bb4f5f3cc2b)
- [service.advertise_options.advertise_on_public.port.tcp_loadbalancer](data-sources--workload--reference--group-015.md#canonical-0e62a645135c3b8ccf8743b299e24a9aa87050bc8ebdb2324540646b774daebe)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-6111d7521abb2c61754e5c7b0229218264c875772bcdbb8788031967782f42e0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a7590b3897182ffe3ba2de9033dfb4fd10cfc9b8d0165fe6d34e2bc20d4f89ca"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer — service.advertise_options.advertise_on_public.port.http_loadbalancer / 23c448ba7c4a / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-a12bf02272720e68e73b00969912e4531958e55ac3c63866089f14cf774f7a3e)
- service.advertise_options.advertise_on_public.port.http_loadbalancer

<a id="canonical-d06169525fd293350fad2358d98eef172ad53843b4459eefd75f4efd72834588"></a>

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

<a id="canonical-c6d3df06513e4a24443bab34dfb646d130b1faefe986187cb39d077b5e6ac85c"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer / 23c448ba7c4a / 3

- [default_route](data-sources--workload--reference--group-012.md#canonical-855935f5092837615cba4e9cc4d147df6a7e0d4d23a588782e46f55f1c021223): complete subsection reference.

<a id="canonical-90594ff768a2bd11cdb8dae3fa86de5816e7cc1b18da1d4e988b3ca5917be8a6"></a>

<a id="canonical-d3fa33444b90bf71d654d29b652e76f3fd6030f4b12e41f787838bf7a3471609"></a>

## domains property — service.advertise_options.advertise_on_public.port.http_loadbalancer / 23c448ba7c4a / 4

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

- [http](data-sources--workload--reference--group-012.md#canonical-66ff8e7d0d9a67d815de9abb1c72199e8a35bfcb587e3e1f170c6f1cfddeed8f): complete subsection reference.

- [https](data-sources--workload--reference--group-012.md#canonical-538e9c665601a7ca0a6bc1a8fd5966adb4add9f6730278f137de37cbf55287da): complete subsection reference.

- [https_auto_cert](data-sources--workload--reference--group-013.md#canonical-e9f8e1bd4658fdad60095f468116bb700679427d9e5d73b9e33307e3f2beedd0): complete subsection reference.

- [specific_routes](data-sources--workload--reference--group-014.md#canonical-4c9aba5983dc47b04849ca6948ed0f4e2e69986678e90302904e7047d8f5d5bf): complete subsection reference.

<a id="canonical-6fe833538fc785345980aee92d06eee5f7da8aa15e0de67df033243c37559463"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer / 23c448ba7c4a / 5

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route](data-sources--workload--reference--group-012.md#canonical-855935f5092837615cba4e9cc4d147df6a7e0d4d23a588782e46f55f1c021223)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.http](data-sources--workload--reference--group-012.md#canonical-66ff8e7d0d9a67d815de9abb1c72199e8a35bfcb587e3e1f170c6f1cfddeed8f)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-012.md#canonical-538e9c665601a7ca0a6bc1a8fd5966adb4add9f6730278f137de37cbf55287da)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-013.md#canonical-e9f8e1bd4658fdad60095f468116bb700679427d9e5d73b9e33307e3f2beedd0)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](data-sources--workload--reference--group-014.md#canonical-4c9aba5983dc47b04849ca6948ed0f4e2e69986678e90302904e7047d8f5d5bf)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-a12bf02272720e68e73b00969912e4531958e55ac3c63866089f14cf774f7a3e)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-855935f5092837615cba4e9cc4d147df6a7e0d4d23a588782e46f55f1c021223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9a59171dee65536d0d3d7fc41a309fc25b589efabd933cef1df401f154de7175"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route — service.advertise_options.advertise_on_public.port.http_loadbalancer.default_rou / 618a1f326a3c / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-a12bf02272720e68e73b00969912e4531958e55ac3c63866089f14cf774f7a3e)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-012.md#canonical-6111d7521abb2c61754e5c7b0229218264c875772bcdbb8788031967782f42e0)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route

<a id="canonical-e4a44281e6f7bf69eda08efb41ad2d1f066efcec4c3634b6698deba9714be47a"></a>

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

<a id="canonical-3004c85c14ad9351772edfc69a9f44bd0c59953a1f941d15e40669867c4f2c0d"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.default_rou / 618a1f326a3c / 3

- [auto_host_rewrite](data-sources--workload--reference--group-012.md#canonical-de73f6a388ba56765a80894f1dc171ae6629e1c5b3df75f9640732950be13fa3): complete subsection reference.

- [disable_host_rewrite](data-sources--workload--reference--group-012.md#canonical-6b8844ae83ade29b9959fa155fb435ff3ed01906d239c9d420f958f964bc01b8): complete subsection reference.

<a id="canonical-ee0e43c4aa61fcbcbc79b859ccd5a695ca4ff076cd2a3b2f4d7e11b1b570d878"></a>

<a id="canonical-4341c740f6d0863b86b15951775e5faecc78e2c299c9cb0913f08d54006cff08"></a>

## host_rewrite property — service.advertise_options.advertise_on_public.port.http_loadbalancer.default_rou / 618a1f326a3c / 4

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

<a id="canonical-149e0918f254e4130d2e69ff4180c893021c624a449cbed64766ddc345118843"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.default_rou / 618a1f326a3c / 5

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route.auto_host_rewrite](data-sources--workload--reference--group-012.md#canonical-de73f6a388ba56765a80894f1dc171ae6629e1c5b3df75f9640732950be13fa3)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route.disable_host_rewrite](data-sources--workload--reference--group-012.md#canonical-6b8844ae83ade29b9959fa155fb435ff3ed01906d239c9d420f958f964bc01b8)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-012.md#canonical-6111d7521abb2c61754e5c7b0229218264c875772bcdbb8788031967782f42e0)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-de73f6a388ba56765a80894f1dc171ae6629e1c5b3df75f9640732950be13fa3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7b83a371ebf292af10b186ca5f4d7ae118eb35441421ee253990f13512b5d5a8"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route.auto_host_rewrite — service.advertise_options.advertise_on_public.port.http_loadbalancer.default_rou / 3e037ec36404 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-a12bf02272720e68e73b00969912e4531958e55ac3c63866089f14cf774f7a3e)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-012.md#canonical-6111d7521abb2c61754e5c7b0229218264c875772bcdbb8788031967782f42e0)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route](data-sources--workload--reference--group-012.md#canonical-855935f5092837615cba4e9cc4d147df6a7e0d4d23a588782e46f55f1c021223)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route.auto_host_rewrite

<a id="canonical-8e643e452a60d40f697086b305bc9c66d3640c5b0503b82f7be877200d4bd671"></a>

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

<a id="canonical-4d8f65c125775e41bf8bf2852544e0bab838de978a50e018a22fbab5070ec56a"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.default_rou / 3e037ec36404 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6ad8f3d6a05922e907b09949ffc4d3c0fdaf5b2889cbe5532daeff3598cf77c7"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.default_rou / 3e037ec36404 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route](data-sources--workload--reference--group-012.md#canonical-855935f5092837615cba4e9cc4d147df6a7e0d4d23a588782e46f55f1c021223)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-6b8844ae83ade29b9959fa155fb435ff3ed01906d239c9d420f958f964bc01b8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-539b990856b188e7e5fc6869fb33dd51a303364bb31249abe25494dd618a7d71"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route.disable_host_rewrite — service.advertise_options.advertise_on_public.port.http_loadbalancer.default_rou / 3c9b0f94e073 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-a12bf02272720e68e73b00969912e4531958e55ac3c63866089f14cf774f7a3e)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-012.md#canonical-6111d7521abb2c61754e5c7b0229218264c875772bcdbb8788031967782f42e0)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route](data-sources--workload--reference--group-012.md#canonical-855935f5092837615cba4e9cc4d147df6a7e0d4d23a588782e46f55f1c021223)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route.disable_host_rewrite

<a id="canonical-6beca10af5bb31195869fc04dd59efe12f110b7ee778f607cbe2fe2f99e9823d"></a>

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

<a id="canonical-6a85a3da10de73408eabc678304eb41da57873a2fed90f1d2eb9eb5cc1f2be22"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.default_rou / 3c9b0f94e073 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-84a6837badad9da75a5b3a913e71d7e0f71b9f4eb718f27b6a777c96c48115fa"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.default_rou / 3c9b0f94e073 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route](data-sources--workload--reference--group-012.md#canonical-855935f5092837615cba4e9cc4d147df6a7e0d4d23a588782e46f55f1c021223)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-66ff8e7d0d9a67d815de9abb1c72199e8a35bfcb587e3e1f170c6f1cfddeed8f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-84314fc08150f3c5262b35a864e6164858846793ac48699586e61defb1e1b173"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.http — service.advertise_options.advertise_on_public.port.http_loadbalancer.http / 05343a4a56a4 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-a12bf02272720e68e73b00969912e4531958e55ac3c63866089f14cf774f7a3e)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-012.md#canonical-6111d7521abb2c61754e5c7b0229218264c875772bcdbb8788031967782f42e0)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.http

<a id="canonical-6410ebc0001730309291aa5770168e9e8208e89dcde3855234e1417a6831a9ad"></a>

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

<a id="canonical-51b502d5175ff5efce5e1396b5935e9842439dc86288d25b574ba8aead63d015"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.http / 05343a4a56a4 / 3

<a id="canonical-adb3445c53dd59d0522ae79f71370e1771f230365823b7ee8c48e066aa163a7e"></a>

<a id="canonical-cbde91f0ae943ebff5fd8c30658813f373d82db4e8097b21530f0cf7bbd60a67"></a>

## dns_volterra_managed property — service.advertise_options.advertise_on_public.port.http_loadbalancer.http / 05343a4a56a4 / 4

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

<a id="canonical-4426a4af2aadb82ef884de3c6632a3155e40a982b831703b9ca305350ce6f5ab"></a>

<a id="canonical-469754cd6ae48eba2eefb7cc5103bf0fa2d8f9c36a4dbe613b1844baba4a5f0d"></a>

## port property — service.advertise_options.advertise_on_public.port.http_loadbalancer.http / 05343a4a56a4 / 5

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

<a id="canonical-a1411dc4a9cfa055069e9b9455ffee3808fdcbc35e9a39013f4fdc7c11c6c68a"></a>

<a id="canonical-b138e91c37e39c3811be06a74d2f1630a0fd7fbf72735beb054145055a1384c0"></a>

## port_ranges property — service.advertise_options.advertise_on_public.port.http_loadbalancer.http / 05343a4a56a4 / 6

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

<a id="canonical-621dd16240e558a6a286042fd3e8f92f4bb83d7547e8d34e751daa917e1cc625"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.http / 05343a4a56a4 / 7

- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-012.md#canonical-6111d7521abb2c61754e5c7b0229218264c875772bcdbb8788031967782f42e0)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-538e9c665601a7ca0a6bc1a8fd5966adb4add9f6730278f137de37cbf55287da"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-548e8f3bf82af54d4ba3941e3c956f3e83a944f8fd3a4ee3992c24e57c159b0d"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https — service.advertise_options.advertise_on_public.port.http_loadbalancer.https / 691616a8c91e / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-a12bf02272720e68e73b00969912e4531958e55ac3c63866089f14cf774f7a3e)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-012.md#canonical-6111d7521abb2c61754e5c7b0229218264c875772bcdbb8788031967782f42e0)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https

<a id="canonical-086ea2e903674d704e75e100b99d7011306a9fdb26511ffd4f2870af772feb70"></a>

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

<a id="canonical-099b2f2d6a11e019a07febb5b0bab174289997b5880ac0c50f5d278eac4d5b48"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https / 691616a8c91e / 3

<a id="canonical-592950f01af12774d8b31b435fbc826bf1da3696235ac0f9e6374a2c04b6f4c0"></a>

<a id="canonical-8e1f18123a82f0d3608fcf00fdd885591d76d8f5557fdcee64c17d1485825578"></a>

## add_hsts property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https / 691616a8c91e / 4

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

<a id="canonical-11f7c82fddcaa8e66a25d06237a5e49fa763ad6c9a1bcd14836b82b4a286ba54"></a>

<a id="canonical-441188b2f9654dcbbe920e06130f3840196f6dc5fb5329e1382864b31c26720b"></a>

## append_server_name property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https / 691616a8c91e / 5

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

- [coalescing_options](data-sources--workload--reference--group-012.md#canonical-2c1f32493d8e0d7f7bb26e6f98418fcb1e7cc3847dc28e3df328f2b5d685c579): complete subsection reference.

<a id="canonical-e8823db4e6dcdca8ff63d004e0a33f99195c50e02f36020112ca7efb4dae26f5"></a>

<a id="canonical-8f8795ff2872145a5ed989407ae98391b1fed2482d96da3f64f0b22136d33aa3"></a>

## connection_idle_timeout property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https / 691616a8c91e / 6

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

- [default_header](data-sources--workload--reference--group-012.md#canonical-f31cfa2d6feb9a9ed07cac79f20ffa5b6127557f514e4e6279bc6ebfdf21a84f): complete subsection reference.

- [default_loadbalancer](data-sources--workload--reference--group-012.md#canonical-a81a90dc9dfec22ce6e0e53dd8bd4c2faaba383ab78282acf6e470daaeb4466d): complete subsection reference.

- [disable_path_normalize](data-sources--workload--reference--group-012.md#canonical-f4991df325bb34083ec99501cc370ea3a6dcda7e4932d337220b2207e53b73c0): complete subsection reference.

- [enable_path_normalize](data-sources--workload--reference--group-012.md#canonical-855af0fc8cfa54583b9aa9d997d4116e9c22a3ef92d1237356aeeec5dd6f31a1): complete subsection reference.

- [http_protocol_options](data-sources--workload--reference--group-012.md#canonical-14dd7b5cf80aa90d068a3e4908d7a7c78128ed611b49b32800fe0256b8a723d0): complete subsection reference.

<a id="canonical-599d55906a5e0e41c9c45cd2f356f85ed41ffc3635c7acac3d485819018ee3fe"></a>

<a id="canonical-d2216e85bd9877371f816335d9ccf153beaa9e6bf0f1a80066305d06a12eba71"></a>

## http_redirect property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https / 691616a8c91e / 7

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

- [non_default_loadbalancer](data-sources--workload--reference--group-012.md#canonical-63ba93699f1452d5436cf4af39984b2243c57a34d6aa5af1869767b4a21660b8): complete subsection reference.

- [pass_through](data-sources--workload--reference--group-012.md#canonical-3b792759b2c1a209f680b735a949de95299f9c70635baab041f73ce0785a4248): complete subsection reference.

<a id="canonical-7b400d52fcbfd33c4d68d43bedafa9c54446514478564cdc1b0d6111e01d7d0e"></a>

<a id="canonical-fd9f602ffb1fdf9df7a957f3aba236f94743e4c1af95a0e92ed46b01a5512e0d"></a>

## port property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https / 691616a8c91e / 8

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

<a id="canonical-386a6f7289824b1e506f8fb6846e958dd64b0562ddffcca8983c05f47f43ca1f"></a>

<a id="canonical-f49393191de2207eceff46d4a508014925ec3df61d21a3aa25d270581b387e67"></a>

## port_ranges property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https / 691616a8c91e / 9

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

<a id="canonical-b7c1dca02c6ebb057b7a41c5c742e8931ca2be1ac552c03066486efa5aed286f"></a>

<a id="canonical-e2bc8cd4e4161390d6127fc403d6c4f10332aa75c6498577ef7b1d028f908f6a"></a>

## server_name property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https / 691616a8c91e / 10

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

- [tls_cert_params](data-sources--workload--reference--group-012.md#canonical-20b77c6c48600ccbb49e7368e5d17c8287ba89c0f87cdb8749b4811900aa8228): complete subsection reference.

- [tls_parameters](data-sources--workload--reference--group-013.md#canonical-53de6e33bf34524b090a82db2f7b6b11d8299cee04d082990dccaf4636b5c9a5): complete subsection reference.

<a id="canonical-243b4cf4eda145ad850ba88eb1f28598eadd278bb9aa39dd435fab01ca4bad09"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https / 691616a8c91e / 11

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options](data-sources--workload--reference--group-012.md#canonical-2c1f32493d8e0d7f7bb26e6f98418fcb1e7cc3847dc28e3df328f2b5d685c579)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.default_header](data-sources--workload--reference--group-012.md#canonical-f31cfa2d6feb9a9ed07cac79f20ffa5b6127557f514e4e6279bc6ebfdf21a84f)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.default_loadbalancer](data-sources--workload--reference--group-012.md#canonical-a81a90dc9dfec22ce6e0e53dd8bd4c2faaba383ab78282acf6e470daaeb4466d)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.disable_path_normalize](data-sources--workload--reference--group-012.md#canonical-f4991df325bb34083ec99501cc370ea3a6dcda7e4932d337220b2207e53b73c0)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.enable_path_normalize](data-sources--workload--reference--group-012.md#canonical-855af0fc8cfa54583b9aa9d997d4116e9c22a3ef92d1237356aeeec5dd6f31a1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-012.md#canonical-14dd7b5cf80aa90d068a3e4908d7a7c78128ed611b49b32800fe0256b8a723d0)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.non_default_loadbalancer](data-sources--workload--reference--group-012.md#canonical-63ba93699f1452d5436cf4af39984b2243c57a34d6aa5af1869767b4a21660b8)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.pass_through](data-sources--workload--reference--group-012.md#canonical-3b792759b2c1a209f680b735a949de95299f9c70635baab041f73ce0785a4248)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-012.md#canonical-20b77c6c48600ccbb49e7368e5d17c8287ba89c0f87cdb8749b4811900aa8228)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-013.md#canonical-53de6e33bf34524b090a82db2f7b6b11d8299cee04d082990dccaf4636b5c9a5)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-012.md#canonical-6111d7521abb2c61754e5c7b0229218264c875772bcdbb8788031967782f42e0)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-2c1f32493d8e0d7f7bb26e6f98418fcb1e7cc3847dc28e3df328f2b5d685c579"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f9e07070857e92e6f50195878cee8412fd4602098a0e811a4eb7a9e112244588"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coale / e4233b2d2c28 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-a12bf02272720e68e73b00969912e4531958e55ac3c63866089f14cf774f7a3e)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-012.md#canonical-6111d7521abb2c61754e5c7b0229218264c875772bcdbb8788031967782f42e0)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-012.md#canonical-538e9c665601a7ca0a6bc1a8fd5966adb4add9f6730278f137de37cbf55287da)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options

<a id="canonical-147e2992806a7953a1c8407a594200696bb95159d5cbd6a016ec7a38b45ab34d"></a>

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

<a id="canonical-ec7c0a49245b8605de881b8e6f78271cebb1656d27d3fec1f5fe87cc70b999c7"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coale / e4233b2d2c28 / 3

- [default_coalescing](data-sources--workload--reference--group-012.md#canonical-5e99045a02a68a7e60d99e1a46ea5b426a545def07ff7a6a85574d4fa9db1796): complete subsection reference.

- [strict_coalescing](data-sources--workload--reference--group-012.md#canonical-d0c4de7eca865a8b53838405455186b559c15007cb7f0aa3fad5b782a8066a35): complete subsection reference.

<a id="canonical-26ecd1749b8eaf4d86ed897f5900c60bce6d92fec5c69b86f046c9dcad20a81c"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coale / e4233b2d2c28 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options.default_coalescing](data-sources--workload--reference--group-012.md#canonical-5e99045a02a68a7e60d99e1a46ea5b426a545def07ff7a6a85574d4fa9db1796)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options.strict_coalescing](data-sources--workload--reference--group-012.md#canonical-d0c4de7eca865a8b53838405455186b559c15007cb7f0aa3fad5b782a8066a35)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-012.md#canonical-538e9c665601a7ca0a6bc1a8fd5966adb4add9f6730278f137de37cbf55287da)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-5e99045a02a68a7e60d99e1a46ea5b426a545def07ff7a6a85574d4fa9db1796"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-43e43b850003ecccbaa2a04b2464a46d071f254e3f4ca8ced48944f3c1e4e864"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options.default_coalescing — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coale / 843368e39496 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-a12bf02272720e68e73b00969912e4531958e55ac3c63866089f14cf774f7a3e)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-012.md#canonical-6111d7521abb2c61754e5c7b0229218264c875772bcdbb8788031967782f42e0)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-012.md#canonical-538e9c665601a7ca0a6bc1a8fd5966adb4add9f6730278f137de37cbf55287da)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options](data-sources--workload--reference--group-012.md#canonical-2c1f32493d8e0d7f7bb26e6f98418fcb1e7cc3847dc28e3df328f2b5d685c579)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options.default_coalescing

<a id="canonical-1a91c08ece35d8950c44918681f2893f2e48988e8b0205f2f6610f1cb60865bd"></a>

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

<a id="canonical-6fc53b2deeb941c06421d61b3afa692c10970ff4326ab103c2144b1bd7591763"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coale / 843368e39496 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2ee27a11745d6af8c9191db699af3e9bef6435a0222638a391b0f788ede0090b"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coale / 843368e39496 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options](data-sources--workload--reference--group-012.md#canonical-2c1f32493d8e0d7f7bb26e6f98418fcb1e7cc3847dc28e3df328f2b5d685c579)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-d0c4de7eca865a8b53838405455186b559c15007cb7f0aa3fad5b782a8066a35"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-48dcf76f94e5691b92218ab1d59405869acd05838513e05e8349f4cbf31c3f8d"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options.strict_coalescing — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coale / 97e24d1a979e / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-a12bf02272720e68e73b00969912e4531958e55ac3c63866089f14cf774f7a3e)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-012.md#canonical-6111d7521abb2c61754e5c7b0229218264c875772bcdbb8788031967782f42e0)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-012.md#canonical-538e9c665601a7ca0a6bc1a8fd5966adb4add9f6730278f137de37cbf55287da)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options](data-sources--workload--reference--group-012.md#canonical-2c1f32493d8e0d7f7bb26e6f98418fcb1e7cc3847dc28e3df328f2b5d685c579)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options.strict_coalescing

<a id="canonical-4a1d193300a01c6ec92552b67f39753acdec4d776d4d903a7f34afc0eff67d9c"></a>

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

<a id="canonical-4cfcc9fa0c7a984d3da1ad963aa863d82c1a3a5a87ad7d8a422e581a7c8fbeb6"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coale / 97e24d1a979e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f93ba4f96edf5e5c2caa4663ba02feb3537247ae970c3ea7a2e414280ce579c6"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coale / 97e24d1a979e / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options](data-sources--workload--reference--group-012.md#canonical-2c1f32493d8e0d7f7bb26e6f98418fcb1e7cc3847dc28e3df328f2b5d685c579)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-f31cfa2d6feb9a9ed07cac79f20ffa5b6127557f514e4e6279bc6ebfdf21a84f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-41edbafdd7a7dfd8764a2042568ca747334aa09f3a4cdfe5adc82ed37ba2e43c"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.default_header — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.defau / a3070281f325 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-a12bf02272720e68e73b00969912e4531958e55ac3c63866089f14cf774f7a3e)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-012.md#canonical-6111d7521abb2c61754e5c7b0229218264c875772bcdbb8788031967782f42e0)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-012.md#canonical-538e9c665601a7ca0a6bc1a8fd5966adb4add9f6730278f137de37cbf55287da)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.default_header

<a id="canonical-6a006268fa000502723fb50216ab2fe2a542b1d13c4f6b3d409bde5865727967"></a>

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

<a id="canonical-62d4f63ade4e5dce8e50896f232093f35fc1864281c550953e4bf153a1eabaf0"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.defau / a3070281f325 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3e57623b9cb5e6065aefc5f56d520f5f5d39fac7bd9df7c8e760a9cf77389789"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.defau / a3070281f325 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-012.md#canonical-538e9c665601a7ca0a6bc1a8fd5966adb4add9f6730278f137de37cbf55287da)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-a81a90dc9dfec22ce6e0e53dd8bd4c2faaba383ab78282acf6e470daaeb4466d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7c4896be576882509a2904278f741ef6fbce8af0cec11384b28c4f65de857fa5"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.default_loadbalancer — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.defau / 131c235a49b6 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-a12bf02272720e68e73b00969912e4531958e55ac3c63866089f14cf774f7a3e)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-012.md#canonical-6111d7521abb2c61754e5c7b0229218264c875772bcdbb8788031967782f42e0)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-012.md#canonical-538e9c665601a7ca0a6bc1a8fd5966adb4add9f6730278f137de37cbf55287da)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.default_loadbalancer

<a id="canonical-43f6f5ae427e439764c58b57c35b3ff49470874528c80b70cb1058b139003e89"></a>

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

<a id="canonical-eeb93b3291ad16f67e15257d7d062ef63df7dabcb51ec7a785e93b2f6d058369"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.defau / 131c235a49b6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a58b307e572c39a37206fccd5f250563053ce717c186c375b54800c33c176f03"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.defau / 131c235a49b6 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-012.md#canonical-538e9c665601a7ca0a6bc1a8fd5966adb4add9f6730278f137de37cbf55287da)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-f4991df325bb34083ec99501cc370ea3a6dcda7e4932d337220b2207e53b73c0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ddee6927f8de9857cbf686f4d0f413f0546fd4b38e96c9486f2b5be9c0697886"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.disable_path_normalize — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.disab / 69731a2308ac / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-a12bf02272720e68e73b00969912e4531958e55ac3c63866089f14cf774f7a3e)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-012.md#canonical-6111d7521abb2c61754e5c7b0229218264c875772bcdbb8788031967782f42e0)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-012.md#canonical-538e9c665601a7ca0a6bc1a8fd5966adb4add9f6730278f137de37cbf55287da)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.disable_path_normalize

<a id="canonical-e7526d1f424960fd50d39229861412306439309d92989195e35968cf4cb36dc0"></a>

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

<a id="canonical-55882aedae6f4ee9a3389a5714498fb5356cdbe95b8739383736410a22a47046"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.disab / 69731a2308ac / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-723566b68cb83e6e45ea33a23d6bd1ee2a29e97304c0398e07aa89175eb9b1d7"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.disab / 69731a2308ac / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-012.md#canonical-538e9c665601a7ca0a6bc1a8fd5966adb4add9f6730278f137de37cbf55287da)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-855af0fc8cfa54583b9aa9d997d4116e9c22a3ef92d1237356aeeec5dd6f31a1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-44bd014bbfd4052016e68e550713b1ebbbed0b94a73cc9d069d5e5d656dc07ec"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.enable_path_normalize — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.enabl / 64a4005c77f6 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-a12bf02272720e68e73b00969912e4531958e55ac3c63866089f14cf774f7a3e)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-012.md#canonical-6111d7521abb2c61754e5c7b0229218264c875772bcdbb8788031967782f42e0)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-012.md#canonical-538e9c665601a7ca0a6bc1a8fd5966adb4add9f6730278f137de37cbf55287da)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.enable_path_normalize

<a id="canonical-3dc8f8e80aecb4ba672c7acdf6f63261aae56a8a7e4dfb8a2a1c0beb9b4c0dfd"></a>

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

<a id="canonical-452f5bf13f78f25108eb2cad26d1ffdde4fae26fc8fc329cf6b55c6d1973c670"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.enabl / 64a4005c77f6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0cbd32d6b119305d59ec43acc3003f5c7036c334362b30eb3b7f77813759c3bf"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.enabl / 64a4005c77f6 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-012.md#canonical-538e9c665601a7ca0a6bc1a8fd5966adb4add9f6730278f137de37cbf55287da)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-14dd7b5cf80aa90d068a3e4908d7a7c78128ed611b49b32800fe0256b8a723d0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-350c927aecd25b88e1ef9ce6c9dae13bf3cdff5e079b3af24838d5c8f124ae00"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_ / 4793fd9849ac / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-a12bf02272720e68e73b00969912e4531958e55ac3c63866089f14cf774f7a3e)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-012.md#canonical-6111d7521abb2c61754e5c7b0229218264c875772bcdbb8788031967782f42e0)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-012.md#canonical-538e9c665601a7ca0a6bc1a8fd5966adb4add9f6730278f137de37cbf55287da)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options

<a id="canonical-96dcb6f973d080ffecb4496cbbb1ac20590a798b428c876e0a1f4a159f965b41"></a>

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

<a id="canonical-ad46dca4cb745f022026cb9c44731892a01825b16c41fb157ecabb6c20da8b2c"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_ / 4793fd9849ac / 3

- [http_protocol_enable_v1_only](data-sources--workload--reference--group-012.md#canonical-e5b54d4cbd06d6156d6ca0e81b92108d3cc188d35c47ef43a8b1a940edcf5bd5): complete subsection reference.

- [http_protocol_enable_v1_v2](data-sources--workload--reference--group-012.md#canonical-76d1b99b0d968f21c6325330f28d1b36692dd252d2768169b045993fd01eac2f): complete subsection reference.

- [http_protocol_enable_v2_only](data-sources--workload--reference--group-012.md#canonical-71056508ced834ee6f59f599b6bbe79752ea8efdad4ea7887f4d7ec4ebd074ae): complete subsection reference.

<a id="canonical-7be9eb1b4a9b038f7d28fc7741ed8a60792bf5240d4860f0b8e700f863bb7d0f"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_ / 4793fd9849ac / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-012.md#canonical-e5b54d4cbd06d6156d6ca0e81b92108d3cc188d35c47ef43a8b1a940edcf5bd5)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2](data-sources--workload--reference--group-012.md#canonical-76d1b99b0d968f21c6325330f28d1b36692dd252d2768169b045993fd01eac2f)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only](data-sources--workload--reference--group-012.md#canonical-71056508ced834ee6f59f599b6bbe79752ea8efdad4ea7887f4d7ec4ebd074ae)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-012.md#canonical-538e9c665601a7ca0a6bc1a8fd5966adb4add9f6730278f137de37cbf55287da)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-e5b54d4cbd06d6156d6ca0e81b92108d3cc188d35c47ef43a8b1a940edcf5bd5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-072c3715b023bd9e526d61ab2809da9ea472fb637981508d755c02e1b49a862b"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_ / c3e328670a9a / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-a12bf02272720e68e73b00969912e4531958e55ac3c63866089f14cf774f7a3e)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-012.md#canonical-6111d7521abb2c61754e5c7b0229218264c875772bcdbb8788031967782f42e0)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-012.md#canonical-538e9c665601a7ca0a6bc1a8fd5966adb4add9f6730278f137de37cbf55287da)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-012.md#canonical-14dd7b5cf80aa90d068a3e4908d7a7c78128ed611b49b32800fe0256b8a723d0)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-01a180b9d07e5cc71d634762d8a6c34712e1aa1523105f0b566d8c7b4ef0e2d4"></a>

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

<a id="canonical-cf462bdc6914f326191c3604bb84021954428e13dfc5ecaa20cfad98b23e5f45"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_ / c3e328670a9a / 3

- [header_transformation](data-sources--workload--reference--group-012.md#canonical-27b94c84a18f8dc158b7d7e0c2d01ef9d7b90e8869703ffbcdb2fa364d78e083): complete subsection reference.

<a id="canonical-e191d879a597d211d2401f5a2efd1fc463598bc7855b46b8658e8527b09c52fc"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_ / c3e328670a9a / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-012.md#canonical-27b94c84a18f8dc158b7d7e0c2d01ef9d7b90e8869703ffbcdb2fa364d78e083)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-012.md#canonical-14dd7b5cf80aa90d068a3e4908d7a7c78128ed611b49b32800fe0256b8a723d0)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-27b94c84a18f8dc158b7d7e0c2d01ef9d7b90e8869703ffbcdb2fa364d78e083"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5a6e3abefc4204416687671a68790e26afbadb5384c1fb68d81d0958e4f5cc36"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_ / 022e8cde9a0e / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-a12bf02272720e68e73b00969912e4531958e55ac3c63866089f14cf774f7a3e)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-012.md#canonical-6111d7521abb2c61754e5c7b0229218264c875772bcdbb8788031967782f42e0)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-012.md#canonical-538e9c665601a7ca0a6bc1a8fd5966adb4add9f6730278f137de37cbf55287da)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-012.md#canonical-14dd7b5cf80aa90d068a3e4908d7a7c78128ed611b49b32800fe0256b8a723d0)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-012.md#canonical-e5b54d4cbd06d6156d6ca0e81b92108d3cc188d35c47ef43a8b1a940edcf5bd5)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-6d5b3c9fab4115b1d311bdf444a5bd56ebed7812d85592c4221ef44c707d7204"></a>

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

<a id="canonical-d591719ca8c3ddb9ddbe862209408ea15811d5b48207b3356faf3cd4908a15ed"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_ / 022e8cde9a0e / 3

- [default_header_transformation](data-sources--workload--reference--group-012.md#canonical-7a6c6c5d782dc59be161fdc2abe3204b57b93a16f01bb206b464292869097e05): complete subsection reference.

- [preserve_case_header_transformation](data-sources--workload--reference--group-012.md#canonical-a461607e02b79fb7cc8069f54782bcb928f403ac3a86477b43668fb6fcc7850c): complete subsection reference.

- [proper_case_header_transformation](data-sources--workload--reference--group-012.md#canonical-2496d9fae82937bfa69298980db602040ff34bc7f628626e566b3bf04447faa0): complete subsection reference.

<a id="canonical-5be5a2875c85a14e372aa25b198cd27ac93350d170b93d5f0ae36c5caca47a11"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_ / 022e8cde9a0e / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](data-sources--workload--reference--group-012.md#canonical-7a6c6c5d782dc59be161fdc2abe3204b57b93a16f01bb206b464292869097e05)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](data-sources--workload--reference--group-012.md#canonical-a461607e02b79fb7cc8069f54782bcb928f403ac3a86477b43668fb6fcc7850c)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](data-sources--workload--reference--group-012.md#canonical-2496d9fae82937bfa69298980db602040ff34bc7f628626e566b3bf04447faa0)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-012.md#canonical-e5b54d4cbd06d6156d6ca0e81b92108d3cc188d35c47ef43a8b1a940edcf5bd5)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-7a6c6c5d782dc59be161fdc2abe3204b57b93a16f01bb206b464292869097e05"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e58b55719764c5de808ff73cd070719b946d9963783c764b91ceb651894a5837"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_ / b8127010e70c / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-a12bf02272720e68e73b00969912e4531958e55ac3c63866089f14cf774f7a3e)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-012.md#canonical-6111d7521abb2c61754e5c7b0229218264c875772bcdbb8788031967782f42e0)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-012.md#canonical-538e9c665601a7ca0a6bc1a8fd5966adb4add9f6730278f137de37cbf55287da)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-012.md#canonical-14dd7b5cf80aa90d068a3e4908d7a7c78128ed611b49b32800fe0256b8a723d0)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-012.md#canonical-e5b54d4cbd06d6156d6ca0e81b92108d3cc188d35c47ef43a8b1a940edcf5bd5)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-012.md#canonical-27b94c84a18f8dc158b7d7e0c2d01ef9d7b90e8869703ffbcdb2fa364d78e083)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-38cbb24a87645c8a1c1bd4ae0719871cc22596ae17cc60f1f9d0e4c766d85f93"></a>

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

<a id="canonical-9e598422f59de29357df4509999e8ed20d832172b57f09769b6401b3b46e6c90"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_ / b8127010e70c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-122cf00137e6ac78df54c63447ffde86738bb0c33cdce7550a5e7d1285f4de7e"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_ / b8127010e70c / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-012.md#canonical-27b94c84a18f8dc158b7d7e0c2d01ef9d7b90e8869703ffbcdb2fa364d78e083)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-a461607e02b79fb7cc8069f54782bcb928f403ac3a86477b43668fb6fcc7850c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9556311e4cbc03b04e55b299bd3b132bc1caca8023230a69f9815f00d9f84c73"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_ / c2bd2abf1ec3 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-a12bf02272720e68e73b00969912e4531958e55ac3c63866089f14cf774f7a3e)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-012.md#canonical-6111d7521abb2c61754e5c7b0229218264c875772bcdbb8788031967782f42e0)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-012.md#canonical-538e9c665601a7ca0a6bc1a8fd5966adb4add9f6730278f137de37cbf55287da)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-012.md#canonical-14dd7b5cf80aa90d068a3e4908d7a7c78128ed611b49b32800fe0256b8a723d0)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-012.md#canonical-e5b54d4cbd06d6156d6ca0e81b92108d3cc188d35c47ef43a8b1a940edcf5bd5)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-012.md#canonical-27b94c84a18f8dc158b7d7e0c2d01ef9d7b90e8869703ffbcdb2fa364d78e083)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-e92783135ec293697a8b6df81f44bf7363f1d51493811d0f26e02466853e80c1"></a>

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

<a id="canonical-9711761723887657b3ce1fe760179fe3b95bcdd911761a67861940bd5d2d9535"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_ / c2bd2abf1ec3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2d2a09287493b94c19d652cbe967d99fc8976f0413a42c1fb9f240ff389f3204"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_ / c2bd2abf1ec3 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-012.md#canonical-27b94c84a18f8dc158b7d7e0c2d01ef9d7b90e8869703ffbcdb2fa364d78e083)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-2496d9fae82937bfa69298980db602040ff34bc7f628626e566b3bf04447faa0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-23b0712eafb8cb2c23fb9827e6afb087d26aa6f4eed456359a5ba547f5c38c73"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_ / c64cd359f27a / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-a12bf02272720e68e73b00969912e4531958e55ac3c63866089f14cf774f7a3e)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-012.md#canonical-6111d7521abb2c61754e5c7b0229218264c875772bcdbb8788031967782f42e0)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-012.md#canonical-538e9c665601a7ca0a6bc1a8fd5966adb4add9f6730278f137de37cbf55287da)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-012.md#canonical-14dd7b5cf80aa90d068a3e4908d7a7c78128ed611b49b32800fe0256b8a723d0)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-012.md#canonical-e5b54d4cbd06d6156d6ca0e81b92108d3cc188d35c47ef43a8b1a940edcf5bd5)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-012.md#canonical-27b94c84a18f8dc158b7d7e0c2d01ef9d7b90e8869703ffbcdb2fa364d78e083)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-3dc7efb6f69d373efa6f02217e29bee0a2d6284de8c1434839e1c0a3d6474e4e"></a>

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

<a id="canonical-52b1a1911838e3339b8681fb154c774ee964bd1ed47de1c59e22520233e51b8e"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_ / c64cd359f27a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0da45b3c226740d83f7ac9983d1d05f4bad26b17b84b9e4a15f852521b935593"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_ / c64cd359f27a / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-012.md#canonical-27b94c84a18f8dc158b7d7e0c2d01ef9d7b90e8869703ffbcdb2fa364d78e083)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-76d1b99b0d968f21c6325330f28d1b36692dd252d2768169b045993fd01eac2f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5739fa975a0efdfd5b99824e3d6b45bab671ec5df9dd18b5542a28ddec716468"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2 — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_ / d5a1b7421424 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-a12bf02272720e68e73b00969912e4531958e55ac3c63866089f14cf774f7a3e)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-012.md#canonical-6111d7521abb2c61754e5c7b0229218264c875772bcdbb8788031967782f42e0)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-012.md#canonical-538e9c665601a7ca0a6bc1a8fd5966adb4add9f6730278f137de37cbf55287da)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-012.md#canonical-14dd7b5cf80aa90d068a3e4908d7a7c78128ed611b49b32800fe0256b8a723d0)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-973a1e3514299104dd15429fbfbe3658e63dc692571105e458c0975813fcea90"></a>

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

<a id="canonical-382f64bd056bb047663d985dd1bbe375f441f2a0d3f6870cd5f52016c6890922"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_ / d5a1b7421424 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ffd78bf903d979216491242ff4bca8f629a91ffd7738ddbb212637b064252e6d"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_ / d5a1b7421424 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-012.md#canonical-14dd7b5cf80aa90d068a3e4908d7a7c78128ed611b49b32800fe0256b8a723d0)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-71056508ced834ee6f59f599b6bbe79752ea8efdad4ea7887f4d7ec4ebd074ae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7331deb06421bfb394312b09724ae58dbd400817d193080d0d265b8911a59363"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_ / 9cd83aad8395 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-a12bf02272720e68e73b00969912e4531958e55ac3c63866089f14cf774f7a3e)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-012.md#canonical-6111d7521abb2c61754e5c7b0229218264c875772bcdbb8788031967782f42e0)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-012.md#canonical-538e9c665601a7ca0a6bc1a8fd5966adb4add9f6730278f137de37cbf55287da)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-012.md#canonical-14dd7b5cf80aa90d068a3e4908d7a7c78128ed611b49b32800fe0256b8a723d0)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-4d3807b8126480b1c57c84e43c0931ecb66fa37b1f9ddc533b89e0b52a23418a"></a>

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

<a id="canonical-2c12014239ccbe6ebeef37b4addf45e2b8df2687f57adbae10401a10bb6c238f"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_ / 9cd83aad8395 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3205e7e08b6c9418360409b25a86ce8d55b61fce945654a30089d150270c524e"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_ / 9cd83aad8395 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-012.md#canonical-14dd7b5cf80aa90d068a3e4908d7a7c78128ed611b49b32800fe0256b8a723d0)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-63ba93699f1452d5436cf4af39984b2243c57a34d6aa5af1869767b4a21660b8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4a16f057f4e7527a6910c278027d189e8a6d7baf053745ed8b947271aa7a277b"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.non_default_loadbalancer — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.non_d / face903a0097 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-a12bf02272720e68e73b00969912e4531958e55ac3c63866089f14cf774f7a3e)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-012.md#canonical-6111d7521abb2c61754e5c7b0229218264c875772bcdbb8788031967782f42e0)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-012.md#canonical-538e9c665601a7ca0a6bc1a8fd5966adb4add9f6730278f137de37cbf55287da)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.non_default_loadbalancer

<a id="canonical-e054b3a9dff84df4d583903e12466d897f8bd54c3593f0f14336328feb3a62fe"></a>

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

<a id="canonical-6b1580dc22af2f6436d656f8514c09e21101ea11f0a4383599298cb5e0273493"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.non_d / face903a0097 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3d02a4372d89153fd1f91c1980bd0b041ae9343b48f132cd517194c804a93888"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.non_d / face903a0097 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-012.md#canonical-538e9c665601a7ca0a6bc1a8fd5966adb4add9f6730278f137de37cbf55287da)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-3b792759b2c1a209f680b735a949de95299f9c70635baab041f73ce0785a4248"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2bf7ad336fbb1f0d21965957681f1c51b2583f51eaac05587739b5ce607e84e2"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.pass_through — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.pass_ / 8e2678114a39 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-a12bf02272720e68e73b00969912e4531958e55ac3c63866089f14cf774f7a3e)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-012.md#canonical-6111d7521abb2c61754e5c7b0229218264c875772bcdbb8788031967782f42e0)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-012.md#canonical-538e9c665601a7ca0a6bc1a8fd5966adb4add9f6730278f137de37cbf55287da)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.pass_through

<a id="canonical-de36c95eeb8825eaa4948c89cdf5206be70bbe7434f1ee29534ae59d1889236e"></a>

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

<a id="canonical-41fadb96cb9ca8490bc594989099a0115e0e2eb222c30dbe91f552c5c4822d35"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.pass_ / 8e2678114a39 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d106763d2b50f35f91a45b77963706dea38616da1fd6d514715cff61d3a0df2a"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.pass_ / 8e2678114a39 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-012.md#canonical-538e9c665601a7ca0a6bc1a8fd5966adb4add9f6730278f137de37cbf55287da)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-20b77c6c48600ccbb49e7368e5d17c8287ba89c0f87cdb8749b4811900aa8228"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dc37856a8e98980384be2597f15efaa374e2c5e6969ca7d4b072a9dd4dfb47ca"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 1c5dbfcb93c2 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-a12bf02272720e68e73b00969912e4531958e55ac3c63866089f14cf774f7a3e)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-012.md#canonical-6111d7521abb2c61754e5c7b0229218264c875772bcdbb8788031967782f42e0)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-012.md#canonical-538e9c665601a7ca0a6bc1a8fd5966adb4add9f6730278f137de37cbf55287da)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params

<a id="canonical-60c482d646fee581cb3699a7b1cc1390f1879857a5a013a1200a625ce8b5c4de"></a>

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

<a id="canonical-9a81bf1c88761cbd49846f6968cb1e4e860238b35695ae4c844fd575764fd13c"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 1c5dbfcb93c2 / 3

- [certificates](data-sources--workload--reference--group-012.md#canonical-eba3bc7a5b4b6cb14245d4336df9c99f81ae9fede647823fc9b8054bd22140db): complete subsection reference.

- [no_mtls](data-sources--workload--reference--group-012.md#canonical-e3593076e050c970739f4d07886e5d832ff3e67f2356d6ddca80ba47e6ea48e4): complete subsection reference.

- [tls_config](data-sources--workload--reference--group-012.md#canonical-2be72677650c8f510d95b9b7489ec207da2041468c8a3cffbd0c6cf0bae94cb7): complete subsection reference.

- [use_mtls](data-sources--workload--reference--group-012.md#canonical-607e34bf98d0ec368a202370433a07376fcc4882a42012a7abb7af44118c70c9): complete subsection reference.

<a id="canonical-5cad73ffc81622476223b4bb5429b96857f0026306f0d17805a97b310c4f60dd"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 1c5dbfcb93c2 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.certificates](data-sources--workload--reference--group-012.md#canonical-eba3bc7a5b4b6cb14245d4336df9c99f81ae9fede647823fc9b8054bd22140db)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.no_mtls](data-sources--workload--reference--group-012.md#canonical-e3593076e050c970739f4d07886e5d832ff3e67f2356d6ddca80ba47e6ea48e4)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-012.md#canonical-2be72677650c8f510d95b9b7489ec207da2041468c8a3cffbd0c6cf0bae94cb7)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-012.md#canonical-607e34bf98d0ec368a202370433a07376fcc4882a42012a7abb7af44118c70c9)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-012.md#canonical-538e9c665601a7ca0a6bc1a8fd5966adb4add9f6730278f137de37cbf55287da)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-eba3bc7a5b4b6cb14245d4336df9c99f81ae9fede647823fc9b8054bd22140db"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-27d77917e62648beb473c7d533e6d482e9f3f53777efff282bb196b9f517780d"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.certificates — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / c3bd1584eb04 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-a12bf02272720e68e73b00969912e4531958e55ac3c63866089f14cf774f7a3e)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-012.md#canonical-6111d7521abb2c61754e5c7b0229218264c875772bcdbb8788031967782f42e0)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-012.md#canonical-538e9c665601a7ca0a6bc1a8fd5966adb4add9f6730278f137de37cbf55287da)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-012.md#canonical-20b77c6c48600ccbb49e7368e5d17c8287ba89c0f87cdb8749b4811900aa8228)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.certificates

<a id="canonical-1c0c48086ce2a58adec94a04e4b6b9767b2db80076a4584971737b83ced8ed92"></a>

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

<a id="canonical-1ba68aa65f5acc783909581721f15b048031dd703082ad324f57e6d7e237c374"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / c3bd1584eb04 / 3

<a id="canonical-bd7206be0d985b0b14ea5427be23d7bf486136d826ec3b3e9f948d94373b3501"></a>

<a id="canonical-d7b63a700df2982a0b43c723bdf1cdd7c472a07351c4981162662f72ca8025c0"></a>

## name property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / c3bd1584eb04 / 4

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

<a id="canonical-259116fe16b580710667c4586020e6d6b9b83c0b768dad4f65ba9ff1996b2823"></a>

<a id="canonical-0e170470d83a0e947e013ce8fa3ff3c240eea39e467eec6916916563b39e5952"></a>

## namespace property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / c3bd1584eb04 / 5

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

<a id="canonical-8f3eb751bc748290a0ce57522e4ab0092e233824ef2cf35ad600a054b4d10f2b"></a>

<a id="canonical-5262741ce79bb7d1448559cfa7a1416543abb9bd0a8a4efc0c95fef1c9be8316"></a>

## tenant property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / c3bd1584eb04 / 6

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

<a id="canonical-0a25b0c52fc6be17e046b1f745c459c0547e244a34dbcbbed0190002f72e9a6b"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / c3bd1584eb04 / 7

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-012.md#canonical-20b77c6c48600ccbb49e7368e5d17c8287ba89c0f87cdb8749b4811900aa8228)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-e3593076e050c970739f4d07886e5d832ff3e67f2356d6ddca80ba47e6ea48e4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aa852142c1502ff94407924c233ef78729c25fa0176ee6c43212e624fcc21a67"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.no_mtls — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 96f27de10605 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-a12bf02272720e68e73b00969912e4531958e55ac3c63866089f14cf774f7a3e)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-012.md#canonical-6111d7521abb2c61754e5c7b0229218264c875772bcdbb8788031967782f42e0)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-012.md#canonical-538e9c665601a7ca0a6bc1a8fd5966adb4add9f6730278f137de37cbf55287da)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-012.md#canonical-20b77c6c48600ccbb49e7368e5d17c8287ba89c0f87cdb8749b4811900aa8228)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.no_mtls

<a id="canonical-41c21de1896c29ceb9cc947a73652c1d3ac02f25baf6e22688d731a85365e1e5"></a>

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

<a id="canonical-f189a21b4b92040a3c6bbb0331f26b3a9afd97bf5d3305b7df9c15851d1d6286"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 96f27de10605 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b51796c84d48fa1a32f0947f68bd804609e9add392f71162f21d457f17d829c2"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 96f27de10605 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-012.md#canonical-20b77c6c48600ccbb49e7368e5d17c8287ba89c0f87cdb8749b4811900aa8228)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-2be72677650c8f510d95b9b7489ec207da2041468c8a3cffbd0c6cf0bae94cb7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-11b3dde72658415cef3d7dc8ee4f7c1ce13ffae1afe6dc2a73ced3860b3b9802"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 587a33d4dc09 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-a12bf02272720e68e73b00969912e4531958e55ac3c63866089f14cf774f7a3e)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-012.md#canonical-6111d7521abb2c61754e5c7b0229218264c875772bcdbb8788031967782f42e0)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-012.md#canonical-538e9c665601a7ca0a6bc1a8fd5966adb4add9f6730278f137de37cbf55287da)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-012.md#canonical-20b77c6c48600ccbb49e7368e5d17c8287ba89c0f87cdb8749b4811900aa8228)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config

<a id="canonical-8eb0c79506e024ea908baf5b6221bb3eb0c3cbc0aab9e810f3a3363327f8b9f5"></a>

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

<a id="canonical-89babbcdf5cbbfa4315430a98e9db136668a3185a1b36b43c0fe56cd06543143"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 587a33d4dc09 / 3

- [custom_security](data-sources--workload--reference--group-012.md#canonical-74ca55c7e042bbf517df6a4712371eb387aa41aa9a9bdd6c1fbb4d2008091b9f): complete subsection reference.

- [default_security](data-sources--workload--reference--group-012.md#canonical-340c999fae28beed28c1a4af28767c7b5ca33fb2fd29a6a37d5e68ffb326d0b6): complete subsection reference.

- [low_security](data-sources--workload--reference--group-012.md#canonical-94d50f563933bbf2167c76f35d90d01d24ad6659a28b6b6caae18cc64ff708fb): complete subsection reference.

- [medium_security](data-sources--workload--reference--group-012.md#canonical-e287f8b249ca6faaed37b520e62356cbe03e53dd043fa6977397fd7378c467fe): complete subsection reference.

<a id="canonical-d90de5918d59a7991236caa336576f80419f44ccd644c0583a88948c2cd390c8"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 587a33d4dc09 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.custom_security](data-sources--workload--reference--group-012.md#canonical-74ca55c7e042bbf517df6a4712371eb387aa41aa9a9bdd6c1fbb4d2008091b9f)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.default_security](data-sources--workload--reference--group-012.md#canonical-340c999fae28beed28c1a4af28767c7b5ca33fb2fd29a6a37d5e68ffb326d0b6)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.low_security](data-sources--workload--reference--group-012.md#canonical-94d50f563933bbf2167c76f35d90d01d24ad6659a28b6b6caae18cc64ff708fb)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.medium_security](data-sources--workload--reference--group-012.md#canonical-e287f8b249ca6faaed37b520e62356cbe03e53dd043fa6977397fd7378c467fe)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-012.md#canonical-20b77c6c48600ccbb49e7368e5d17c8287ba89c0f87cdb8749b4811900aa8228)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-74ca55c7e042bbf517df6a4712371eb387aa41aa9a9bdd6c1fbb4d2008091b9f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e1eca51e484c3e57a01d29fb03314191556723f66981413e7d97c5c221a7b420"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.custom_security — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 1642a959b2ce / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-a12bf02272720e68e73b00969912e4531958e55ac3c63866089f14cf774f7a3e)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-012.md#canonical-6111d7521abb2c61754e5c7b0229218264c875772bcdbb8788031967782f42e0)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-012.md#canonical-538e9c665601a7ca0a6bc1a8fd5966adb4add9f6730278f137de37cbf55287da)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-012.md#canonical-20b77c6c48600ccbb49e7368e5d17c8287ba89c0f87cdb8749b4811900aa8228)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-012.md#canonical-2be72677650c8f510d95b9b7489ec207da2041468c8a3cffbd0c6cf0bae94cb7)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.custom_security

<a id="canonical-305d33819a189e8e57cc8d08d8d948bfca0b820d62756658dfa60516046b2b8b"></a>

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

<a id="canonical-e4c6012fea0bc316d6c3165173fd0e29045490eb2849e0a89dee5ddced02cb3a"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 1642a959b2ce / 3

<a id="canonical-ee5389980692c52920e0467c8141f11d27e6648084a59390d84ea0b217f3187f"></a>

<a id="canonical-edda8d54d33abb9e6c9c789bb2867207252271ce4e2ab2d06dc2509d96670e13"></a>

## cipher_suites property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 1642a959b2ce / 4

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

<a id="canonical-7269d32acb8ffa3b1fb36803f5793968d694dc39688f3395e59f87fa1aca87a3"></a>

<a id="canonical-06ae022851e5061649d4d67c541b708d8c6890c50525f06e9b962102de7334c2"></a>

## max_version property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 1642a959b2ce / 5

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

<a id="canonical-1442e84487c1ec800eac31cfe7972c0d7b2c3903e8817cb969fc648ce711e14c"></a>

<a id="canonical-bd36dfa358ca64abe9d3cc721598eec2146d26d7ad39133cbba42209a64a9729"></a>

## min_version property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 1642a959b2ce / 6

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

<a id="canonical-bd4fb23579b133aff5937d86c46f2c0533f5bf6dc386c8247d36889a959b517e"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 1642a959b2ce / 7

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-012.md#canonical-2be72677650c8f510d95b9b7489ec207da2041468c8a3cffbd0c6cf0bae94cb7)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-340c999fae28beed28c1a4af28767c7b5ca33fb2fd29a6a37d5e68ffb326d0b6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ca3d4700cba3747e511636cb7d0d7c72a51127da73b8ea84f051df75df6ba59c"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.default_security — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / a9c852e1b282 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-a12bf02272720e68e73b00969912e4531958e55ac3c63866089f14cf774f7a3e)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-012.md#canonical-6111d7521abb2c61754e5c7b0229218264c875772bcdbb8788031967782f42e0)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-012.md#canonical-538e9c665601a7ca0a6bc1a8fd5966adb4add9f6730278f137de37cbf55287da)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-012.md#canonical-20b77c6c48600ccbb49e7368e5d17c8287ba89c0f87cdb8749b4811900aa8228)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-012.md#canonical-2be72677650c8f510d95b9b7489ec207da2041468c8a3cffbd0c6cf0bae94cb7)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.default_security

<a id="canonical-d06abecea752ed37016e0fec6bf4685e26777d04dbb4f208c275441364f91a0e"></a>

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

<a id="canonical-1520eda2d1715dbf91d37665ca4d31e21a8cafe83279865817bdca27d6f1a813"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / a9c852e1b282 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-da66d7e5a1cd7c5d5b1f83295e0943e9febc4d2401a83a565f9ae69d2969704f"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / a9c852e1b282 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-012.md#canonical-2be72677650c8f510d95b9b7489ec207da2041468c8a3cffbd0c6cf0bae94cb7)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-94d50f563933bbf2167c76f35d90d01d24ad6659a28b6b6caae18cc64ff708fb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5557bbc9b73472fd1c6b42d27edcc7897feb368323598ea811b6733ffc5fb28b"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.low_security — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 279ef14c67f4 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-a12bf02272720e68e73b00969912e4531958e55ac3c63866089f14cf774f7a3e)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-012.md#canonical-6111d7521abb2c61754e5c7b0229218264c875772bcdbb8788031967782f42e0)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-012.md#canonical-538e9c665601a7ca0a6bc1a8fd5966adb4add9f6730278f137de37cbf55287da)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-012.md#canonical-20b77c6c48600ccbb49e7368e5d17c8287ba89c0f87cdb8749b4811900aa8228)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-012.md#canonical-2be72677650c8f510d95b9b7489ec207da2041468c8a3cffbd0c6cf0bae94cb7)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.low_security

<a id="canonical-9701300abe5aea9d5ab555174363c848579500869bedc520d74bf0e08d38ae9a"></a>

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

<a id="canonical-20ff9db54978b710e7f31981d8cef777cd64ac6be5a4b013e2df72d2999bd4b7"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 279ef14c67f4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-08adc67654487c603a5750112ba326dd651132fba10b2180e6908ac3815ef4fd"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 279ef14c67f4 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-012.md#canonical-2be72677650c8f510d95b9b7489ec207da2041468c8a3cffbd0c6cf0bae94cb7)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-e287f8b249ca6faaed37b520e62356cbe03e53dd043fa6977397fd7378c467fe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d6bc18625343770a3ae8c9e9c75719dc58da1a93a6b50d4147dec7751d2f9abc"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.medium_security — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 30d490601cdd / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-a12bf02272720e68e73b00969912e4531958e55ac3c63866089f14cf774f7a3e)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-012.md#canonical-6111d7521abb2c61754e5c7b0229218264c875772bcdbb8788031967782f42e0)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-012.md#canonical-538e9c665601a7ca0a6bc1a8fd5966adb4add9f6730278f137de37cbf55287da)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-012.md#canonical-20b77c6c48600ccbb49e7368e5d17c8287ba89c0f87cdb8749b4811900aa8228)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-012.md#canonical-2be72677650c8f510d95b9b7489ec207da2041468c8a3cffbd0c6cf0bae94cb7)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.medium_security

<a id="canonical-700105ff99263a7a5ec1d247ef76c5005256478374612a17add2e95cd2c2abc6"></a>

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

<a id="canonical-5030a134307947e17aa9717fb3a71e6fd65f58ab13a651943167e3c8ffc84cf5"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 30d490601cdd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f0609ad760df1a7dac9a618277e4b07cdd31aa10e653711ab643e4d3a4a0bf50"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 30d490601cdd / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-012.md#canonical-2be72677650c8f510d95b9b7489ec207da2041468c8a3cffbd0c6cf0bae94cb7)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-607e34bf98d0ec368a202370433a07376fcc4882a42012a7abb7af44118c70c9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4b08e817da77bede5edd53020508a516909bf09ca50b1a677a5ed45a5700135a"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 67e932a1a382 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-a12bf02272720e68e73b00969912e4531958e55ac3c63866089f14cf774f7a3e)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-012.md#canonical-6111d7521abb2c61754e5c7b0229218264c875772bcdbb8788031967782f42e0)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-012.md#canonical-538e9c665601a7ca0a6bc1a8fd5966adb4add9f6730278f137de37cbf55287da)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-012.md#canonical-20b77c6c48600ccbb49e7368e5d17c8287ba89c0f87cdb8749b4811900aa8228)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls

<a id="canonical-effa6b140d945cb1ae842b7c76656ccb211d9f4cd2f09851f6dcc614f3504b7e"></a>

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

<a id="canonical-696e1522328ece55403cdc1573952c6987353d07311570316c46feb74c31bc61"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 67e932a1a382 / 3

<a id="canonical-2ae0e05404bcf3e5068193fa3710e1ecae5b4b0b6bd2c913edd29dd8c28701f7"></a>

<a id="canonical-0c9027253dee8f8ea4c43e2144dc24022c1bc8093eba12d707b3cd49d79b8dd3"></a>

## client_certificate_optional property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 67e932a1a382 / 4

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

- [crl](data-sources--workload--reference--group-012.md#canonical-8a94da610bc634882c923adf442e69eb085d0c744e86baf6886b242c9d0afeed): complete subsection reference.

- [no_crl](data-sources--workload--reference--group-012.md#canonical-101d5053a84cb71b58f6f58363132ef3f87abfb843273898fd30af75fd14ce84): complete subsection reference.

- [trusted_ca](data-sources--workload--reference--group-013.md#canonical-22b89bc1c7f9e9506d9e2f0c79b2a92e44914b66ed2581d67869b61aa935e064): complete subsection reference.

<a id="canonical-c15be0d3bd073834c2ca08b5fcdead4a6dcbf220eec09691534223a4f2479e26"></a>

<a id="canonical-cd70b9cc2d295da20a9d81e5df6e000773586c568f4d5565bc6a48ff383c53f5"></a>

## trusted_ca_url property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 67e932a1a382 / 5

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

- [xfcc_disabled](data-sources--workload--reference--group-013.md#canonical-f552800fc652080a6bf38f5b636c484de7e2afd366d03df31eb290e0b2cf9d39): complete subsection reference.

- [xfcc_options](data-sources--workload--reference--group-013.md#canonical-5a2f46dc50403934a8804766e458aa5c871ff0c38ef819359f027cbf3f2f27f1): complete subsection reference.

<a id="canonical-353ea7febcc65a4d9ec0558400897840263a34ddbc6a8f3913771eaccdab768e"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 67e932a1a382 / 6

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.crl](data-sources--workload--reference--group-012.md#canonical-8a94da610bc634882c923adf442e69eb085d0c744e86baf6886b242c9d0afeed)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.no_crl](data-sources--workload--reference--group-012.md#canonical-101d5053a84cb71b58f6f58363132ef3f87abfb843273898fd30af75fd14ce84)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca](data-sources--workload--reference--group-013.md#canonical-22b89bc1c7f9e9506d9e2f0c79b2a92e44914b66ed2581d67869b61aa935e064)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_disabled](data-sources--workload--reference--group-013.md#canonical-f552800fc652080a6bf38f5b636c484de7e2afd366d03df31eb290e0b2cf9d39)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options](data-sources--workload--reference--group-013.md#canonical-5a2f46dc50403934a8804766e458aa5c871ff0c38ef819359f027cbf3f2f27f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-012.md#canonical-20b77c6c48600ccbb49e7368e5d17c8287ba89c0f87cdb8749b4811900aa8228)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-8a94da610bc634882c923adf442e69eb085d0c744e86baf6886b242c9d0afeed"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0977eaca7e333aa266ea272dce4c48c4a6ffb6f71e0f491aa40a237ef41fada3"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.crl — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / fb0a5a213d52 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-a12bf02272720e68e73b00969912e4531958e55ac3c63866089f14cf774f7a3e)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-012.md#canonical-6111d7521abb2c61754e5c7b0229218264c875772bcdbb8788031967782f42e0)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-012.md#canonical-538e9c665601a7ca0a6bc1a8fd5966adb4add9f6730278f137de37cbf55287da)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-012.md#canonical-20b77c6c48600ccbb49e7368e5d17c8287ba89c0f87cdb8749b4811900aa8228)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-012.md#canonical-607e34bf98d0ec368a202370433a07376fcc4882a42012a7abb7af44118c70c9)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.crl

<a id="canonical-026331cd72c62327ade00a5cb7dd55433ca18619b312c0867a514127ed7a1e34"></a>

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

<a id="canonical-cd95a6f0616cdec1c423edc020870d51b016209c4e29f81fba54e0619f55b2d7"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / fb0a5a213d52 / 3

<a id="canonical-fc65dfeff54f5f78b506e63947c2f70fb9e0a4a18872be9df1c344d5261aaaf8"></a>

<a id="canonical-e580053e181acf7e7f09d716e8b0c2bcdb787cc00c924f8fc986f52596a460d6"></a>

## name property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / fb0a5a213d52 / 4

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

<a id="canonical-cfce77528fc645f3c66fff21410ec48e552ffcfaa9f0fe5bdb0639adc0f203a8"></a>

<a id="canonical-b3aad200197895d3540a72e22356022789cfec60014db734ccdfd7916e16f046"></a>

## namespace property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / fb0a5a213d52 / 5

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

<a id="canonical-f468ba455c7e489c9c62a6ab1d1b9f686b1ca95c31c0a5a8fdb17a76d7090dda"></a>

<a id="canonical-1b172c69eee0c2c8297b548ddde9ee67af574554782c52e9e8b9e9ccb0b79d49"></a>

## tenant property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / fb0a5a213d52 / 6

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

<a id="canonical-ebec49f38b0673c0a2282da4d347394ac4f21614e55214e96311ed4dfd1b9c60"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / fb0a5a213d52 / 7

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-012.md#canonical-607e34bf98d0ec368a202370433a07376fcc4882a42012a7abb7af44118c70c9)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-101d5053a84cb71b58f6f58363132ef3f87abfb843273898fd30af75fd14ce84"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
