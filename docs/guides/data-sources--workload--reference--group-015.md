---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-0117213e6e26b21735b2c336656e85484dd9364e844e33ff58f7a65dc457ad98"></a>

## response_code property — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / adf45e400ffc / 9

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

- [retain_all_params](data-sources--workload--reference--group-015.md#canonical-126770edf9cdf8a43f1163cbaaffe0410fc9417bad279c4362a7b21fdf7c6422): complete subsection reference.

<a id="canonical-bc12392ad46518ef0e583271f190ff8578418ffd707547ba5d2a73d4a5e5da1a"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / adf45e400ffc / 10

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.remove_all_params](data-sources--workload--reference--group-015.md#canonical-a2dfc1a19b36db6253a2891c821ef0e9c507c2d79deaf3d185d97951aa47da3c)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.retain_all_params](data-sources--workload--reference--group-015.md#canonical-126770edf9cdf8a43f1163cbaaffe0410fc9417bad279c4362a7b21fdf7c6422)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-014.md#canonical-5a62b2f21ab553b586718957280996d37baede8377e98a41757fb8897fcaa499)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-a2dfc1a19b36db6253a2891c821ef0e9c507c2d79deaf3d185d97951aa47da3c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-db6c61ffa07c20b271d08053fc57cc0d4d416f57b62273f7b134c6f9453821fe"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.remove_all_params — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 6b7fb770e92b / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-a12bf02272720e68e73b00969912e4531958e55ac3c63866089f14cf774f7a3e)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-012.md#canonical-6111d7521abb2c61754e5c7b0229218264c875772bcdbb8788031967782f42e0)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](data-sources--workload--reference--group-014.md#canonical-4c9aba5983dc47b04849ca6948ed0f4e2e69986678e90302904e7047d8f5d5bf)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-014.md#canonical-d5c49233094262eaeca10374cad6895c6f07f30ff553958aeecc38b276faaa92)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-014.md#canonical-5a62b2f21ab553b586718957280996d37baede8377e98a41757fb8897fcaa499)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](data-sources--workload--reference--group-014.md#canonical-56e197682ddaa1a5f23b42e6ef332edd513fda8432ae1836ccb1084a44dc289e)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.remove_all_params

<a id="canonical-ead73b9ea16f4e74ab75c89102c5a00100cf41baac68a545b57f26be97129314"></a>

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

<a id="canonical-4b04bd39dfb45fe99c709656c48701d2ba7837cdaaf52575a7cdb262b16dd369"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 6b7fb770e92b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e0a6aa4e29ae1fd8377f6cab56261436838fa85038fe86ce8b48355fafa152a8"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 6b7fb770e92b / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](data-sources--workload--reference--group-014.md#canonical-56e197682ddaa1a5f23b42e6ef332edd513fda8432ae1836ccb1084a44dc289e)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-126770edf9cdf8a43f1163cbaaffe0410fc9417bad279c4362a7b21fdf7c6422"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4fed9ca196219cc20fb8826c28ae1834f58abb04da3005138888899bc395a1c2"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.retain_all_params — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 1ad3c64fd976 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-a12bf02272720e68e73b00969912e4531958e55ac3c63866089f14cf774f7a3e)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-012.md#canonical-6111d7521abb2c61754e5c7b0229218264c875772bcdbb8788031967782f42e0)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](data-sources--workload--reference--group-014.md#canonical-4c9aba5983dc47b04849ca6948ed0f4e2e69986678e90302904e7047d8f5d5bf)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-014.md#canonical-d5c49233094262eaeca10374cad6895c6f07f30ff553958aeecc38b276faaa92)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-014.md#canonical-5a62b2f21ab553b586718957280996d37baede8377e98a41757fb8897fcaa499)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](data-sources--workload--reference--group-014.md#canonical-56e197682ddaa1a5f23b42e6ef332edd513fda8432ae1836ccb1084a44dc289e)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.retain_all_params

<a id="canonical-bc86a2cb86663a301bfd94cda48729cee1e200301ecacad91ed8366b1e725723"></a>

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

<a id="canonical-924893b486874fc452d365bbfe6285b3176f164b52033422e1f244ea62fa2081"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 1ad3c64fd976 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5714c61b398724dbc885bd4b386d1719db8b7aaaad2e93381af07087afc3d0be"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 1ad3c64fd976 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](data-sources--workload--reference--group-014.md#canonical-56e197682ddaa1a5f23b42e6ef332edd513fda8432ae1836ccb1084a44dc289e)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-ef1fb0fd8f470ec5704a672b2b6c3859905a5b7fbe7ef23ecba2688aaea2fdfa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6016b686e135933b7c3a3f10a1250323a79e8ac21195750fb1f7fec8a210575e"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 16a26b98c601 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-a12bf02272720e68e73b00969912e4531958e55ac3c63866089f14cf774f7a3e)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-012.md#canonical-6111d7521abb2c61754e5c7b0229218264c875772bcdbb8788031967782f42e0)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](data-sources--workload--reference--group-014.md#canonical-4c9aba5983dc47b04849ca6948ed0f4e2e69986678e90302904e7047d8f5d5bf)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-014.md#canonical-d5c49233094262eaeca10374cad6895c6f07f30ff553958aeecc38b276faaa92)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route

<a id="canonical-2ca2ea77a517ffcbfe353d59f028b82a2c0fec22693b58309c3d3e288bf6fc01"></a>

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

<a id="canonical-7b9a771d9a6cf659f1a04d7064605d4d752e3c295633756189e7727cd9584d23"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 16a26b98c601 / 3

- [auto_host_rewrite](data-sources--workload--reference--group-015.md#canonical-8b46ef8cd9066881d3077c8d652f000738ce78424118b659082500f1d66674c5): complete subsection reference.

- [disable_host_rewrite](data-sources--workload--reference--group-015.md#canonical-fe2048a7a95f3759694162e0f5faf870752d429b20204dda07bd2ca096f56390): complete subsection reference.

<a id="canonical-0420750d424c0b6949bc7a378b23c5b22618d2fd965d15a3da6318c5e784bb2e"></a>

<a id="canonical-0be06f759ba58db5a28719d9e7d36f01c4a6764bce86a1772666c7bf871ac111"></a>

## host_rewrite property — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 16a26b98c601 / 4

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

<a id="canonical-fd6b1bb3a3644c244839362fd5b9c0df465fb95ef5d2253700cbaa4cc43255bf"></a>

<a id="canonical-222b4e535363fb44a652bcac700fc23290e90b00a033f95aafc20710c52be481"></a>

## http_method property — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 16a26b98c601 / 5

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

- [path](data-sources--workload--reference--group-015.md#canonical-b667f7f29fa0b13b065fcec28651f24bef937fabe53c4a7fcc113fd20d8e8f88): complete subsection reference.

<a id="canonical-96ac246368b4d5ddcab3baaadbdd662d05307a6e510a85798e42a9c863ddb27a"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 16a26b98c601 / 6

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.auto_host_rewrite](data-sources--workload--reference--group-015.md#canonical-8b46ef8cd9066881d3077c8d652f000738ce78424118b659082500f1d66674c5)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite](data-sources--workload--reference--group-015.md#canonical-fe2048a7a95f3759694162e0f5faf870752d429b20204dda07bd2ca096f56390)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.path](data-sources--workload--reference--group-015.md#canonical-b667f7f29fa0b13b065fcec28651f24bef937fabe53c4a7fcc113fd20d8e8f88)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-014.md#canonical-d5c49233094262eaeca10374cad6895c6f07f30ff553958aeecc38b276faaa92)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-8b46ef8cd9066881d3077c8d652f000738ce78424118b659082500f1d66674c5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5caf3f377071d5185d6b4033b55ae7d868dbb81685994da64be32cfd8edc5932"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.auto_host_rewrite — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 97fb18f5fe3b / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-a12bf02272720e68e73b00969912e4531958e55ac3c63866089f14cf774f7a3e)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-012.md#canonical-6111d7521abb2c61754e5c7b0229218264c875772bcdbb8788031967782f42e0)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](data-sources--workload--reference--group-014.md#canonical-4c9aba5983dc47b04849ca6948ed0f4e2e69986678e90302904e7047d8f5d5bf)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-014.md#canonical-d5c49233094262eaeca10374cad6895c6f07f30ff553958aeecc38b276faaa92)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route](data-sources--workload--reference--group-015.md#canonical-ef1fb0fd8f470ec5704a672b2b6c3859905a5b7fbe7ef23ecba2688aaea2fdfa)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.auto_host_rewrite

<a id="canonical-5dd2a242ab290116a0f786b088f141cc6f913d99085e6f2ebb086173118d12d2"></a>

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

<a id="canonical-49b301af53bb9e169b5802df92241a977d52b5dd96a2557a66025a75fbabb65e"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 97fb18f5fe3b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-aa670f0e77122f4673db408c204f23451498216414d0da8b9122711deb02a95f"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 97fb18f5fe3b / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route](data-sources--workload--reference--group-015.md#canonical-ef1fb0fd8f470ec5704a672b2b6c3859905a5b7fbe7ef23ecba2688aaea2fdfa)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-fe2048a7a95f3759694162e0f5faf870752d429b20204dda07bd2ca096f56390"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cc57eb3e2dcccab401bb41eebaeab5aee6ee5bec06f8a049a21f8ae8756671a8"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / a1d72a635652 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-a12bf02272720e68e73b00969912e4531958e55ac3c63866089f14cf774f7a3e)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-012.md#canonical-6111d7521abb2c61754e5c7b0229218264c875772bcdbb8788031967782f42e0)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](data-sources--workload--reference--group-014.md#canonical-4c9aba5983dc47b04849ca6948ed0f4e2e69986678e90302904e7047d8f5d5bf)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-014.md#canonical-d5c49233094262eaeca10374cad6895c6f07f30ff553958aeecc38b276faaa92)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route](data-sources--workload--reference--group-015.md#canonical-ef1fb0fd8f470ec5704a672b2b6c3859905a5b7fbe7ef23ecba2688aaea2fdfa)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite

<a id="canonical-202617bce747c7a17e0088ed94aa78778a73a8c830359f032063ce5a4a8625b1"></a>

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

<a id="canonical-fd6a820356c8f1c5e7e6899f5e30abb7a896f605857739db0a533ff8fb0ec9df"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / a1d72a635652 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-064c3d1f2e6e2e95f38ea99eb7871d209ab260c2be8d5e8b9fd61fd560d63656"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / a1d72a635652 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route](data-sources--workload--reference--group-015.md#canonical-ef1fb0fd8f470ec5704a672b2b6c3859905a5b7fbe7ef23ecba2688aaea2fdfa)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-b667f7f29fa0b13b065fcec28651f24bef937fabe53c4a7fcc113fd20d8e8f88"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-375ba300cbdb8ed3f5495b59bd23855abc52d5bbc856fa524840c74d7a06a1a4"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.path — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 9fb824411faf / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-a12bf02272720e68e73b00969912e4531958e55ac3c63866089f14cf774f7a3e)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-012.md#canonical-6111d7521abb2c61754e5c7b0229218264c875772bcdbb8788031967782f42e0)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](data-sources--workload--reference--group-014.md#canonical-4c9aba5983dc47b04849ca6948ed0f4e2e69986678e90302904e7047d8f5d5bf)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-014.md#canonical-d5c49233094262eaeca10374cad6895c6f07f30ff553958aeecc38b276faaa92)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route](data-sources--workload--reference--group-015.md#canonical-ef1fb0fd8f470ec5704a672b2b6c3859905a5b7fbe7ef23ecba2688aaea2fdfa)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.path

<a id="canonical-6041b70ca5a3b3c00716102489cbb20931bad247fa9259df3950ef9f6b2ac755"></a>

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

<a id="canonical-4b3bc4b2f289acee1c452f6a480c639e334861df3dc8e9692ae876eca5d70b12"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 9fb824411faf / 3

<a id="canonical-19a8af8a4bc8b145169e6d41d149160b9e8e34894d5c99908f036eb3100244b5"></a>

<a id="canonical-fe22df6d2d6b41c28cef571e24f35605ef5a7b6bbe570e40ac84d2554747c64b"></a>

## path property — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 9fb824411faf / 4

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

<a id="canonical-4366b506f1d7a63806dec3690ce84600ab674a4aacf1d382bd37652d935bce96"></a>

<a id="canonical-44c486a9f0674b9f11c2dc0750d95e5b83145882d777435e260429e238139191"></a>

## prefix property — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 9fb824411faf / 5

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

<a id="canonical-345a462152862f0b4a2e2b068dbf5150d5ccb037420e35bd672b833a4008b40b"></a>

<a id="canonical-3f991d669822eec353991a227f55d705cda4ebfdc05ea502ec645b49114bc692"></a>

## regex property — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 9fb824411faf / 6

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

<a id="canonical-00a53e5c42eeb109641326613eac94efc5a2c8be71d9e930b23647ee7517b6fa"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 9fb824411faf / 7

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route](data-sources--workload--reference--group-015.md#canonical-ef1fb0fd8f470ec5704a672b2b6c3859905a5b7fbe7ef23ecba2688aaea2fdfa)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-17050b4cdfcb149ca58593a78653468f7c97221f865c86f032f71bb4f5f3cc2b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dacd2c588e3ebcb498d53336ffd464b8146b606bc8217e8f20fcdb7c2f3ff55b"></a>

## service.advertise_options.advertise_on_public.port.port — service.advertise_options.advertise_on_public.port.port / 3f493b224767 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-a12bf02272720e68e73b00969912e4531958e55ac3c63866089f14cf774f7a3e)
- service.advertise_options.advertise_on_public.port.port

<a id="canonical-f678d439b6bb14e581b829d37b0a79697a90c424c297cc51164faae881072e6e"></a>

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

<a id="canonical-401db2790f88a086056663d315e083386eec83099a7fdd9a5d9f5701326cf6da"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.port / 3f493b224767 / 3

- [info](data-sources--workload--reference--group-015.md#canonical-c0eab234681276b5b679fec6a1889c6499c45b2573ac84626157885203a83704): complete subsection reference.

<a id="canonical-2b1d80eeaf7d754daa7092efc838d52ef037c71b7965052423f440e6c9549be2"></a>

## Next pages — service.advertise_options.advertise_on_public.port.port / 3f493b224767 / 4

- [service.advertise_options.advertise_on_public.port.port.info](data-sources--workload--reference--group-015.md#canonical-c0eab234681276b5b679fec6a1889c6499c45b2573ac84626157885203a83704)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-a12bf02272720e68e73b00969912e4531958e55ac3c63866089f14cf774f7a3e)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-c0eab234681276b5b679fec6a1889c6499c45b2573ac84626157885203a83704"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b0df09433123b0a6965117102958844807df3e3f15ee69ba16b2f3b34ec026b3"></a>

## service.advertise_options.advertise_on_public.port.port.info — service.advertise_options.advertise_on_public.port.port.info / 9f254dc23011 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-a12bf02272720e68e73b00969912e4531958e55ac3c63866089f14cf774f7a3e)
- [service.advertise_options.advertise_on_public.port.port](data-sources--workload--reference--group-015.md#canonical-17050b4cdfcb149ca58593a78653468f7c97221f865c86f032f71bb4f5f3cc2b)
- service.advertise_options.advertise_on_public.port.port.info

<a id="canonical-accdcf80e3c836f9ef978aac34f2675257c2c33bb5bda1390ac6e251b17f90ae"></a>

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

<a id="canonical-11530581a4a057bd22f2484ccdb4cd8828202e19d3b62af0cf3d3ec5dae2a5ac"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.port.info / 9f254dc23011 / 3

<a id="canonical-e6a4ddfab72cc57fc0165a4fb9fda3c3e618bd29cfee2bbcc0e89190857ed347"></a>

<a id="canonical-45141a1345423f7f1085121ee4e9c047c3b4250438d428ac79b4594ae8cd62c9"></a>

## port property — service.advertise_options.advertise_on_public.port.port.info / 9f254dc23011 / 4

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

<a id="canonical-7076dee562fa330e6004c73f6e7201f992aff33b68d0b14ddb8ddb39b789186e"></a>

<a id="canonical-2c7b3120ef9f28104b98b2c6d81ffc8907eb342ecfb476c7cf9c1473c647a499"></a>

## protocol property — service.advertise_options.advertise_on_public.port.port.info / 9f254dc23011 / 5

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

- [same_as_port](data-sources--workload--reference--group-015.md#canonical-15fcf3801e1f14953b431c97d23bb081262b037d665ff3b183ac0781d5c189e1): complete subsection reference.

<a id="canonical-35a666eb43814e65e8209c5e042e5ffaf0a083a55a56dc4232528edd3a4e906b"></a>

<a id="canonical-7d1d1f5f3c16889587f2506070a93dc43af5927d7be5e355458b3335687950f0"></a>

## target_port property — service.advertise_options.advertise_on_public.port.port.info / 9f254dc23011 / 6

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

<a id="canonical-61ea9b1796a4fac2e561a4fa7d30cd35fc3b6e87071a0035461500c29d554779"></a>

## Next pages — service.advertise_options.advertise_on_public.port.port.info / 9f254dc23011 / 7

- [service.advertise_options.advertise_on_public.port.port.info.same_as_port](data-sources--workload--reference--group-015.md#canonical-15fcf3801e1f14953b431c97d23bb081262b037d665ff3b183ac0781d5c189e1)
- [service.advertise_options.advertise_on_public.port.port](data-sources--workload--reference--group-015.md#canonical-17050b4cdfcb149ca58593a78653468f7c97221f865c86f032f71bb4f5f3cc2b)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-15fcf3801e1f14953b431c97d23bb081262b037d665ff3b183ac0781d5c189e1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-975ddf5445b1aa4ec930102e7222437d796942714bd92526f30b5c0d30f5d13a"></a>

## service.advertise_options.advertise_on_public.port.port.info.same_as_port — service.advertise_options.advertise_on_public.port.port.info.same_as_port / 5abfd1c96088 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-a12bf02272720e68e73b00969912e4531958e55ac3c63866089f14cf774f7a3e)
- [service.advertise_options.advertise_on_public.port.port](data-sources--workload--reference--group-015.md#canonical-17050b4cdfcb149ca58593a78653468f7c97221f865c86f032f71bb4f5f3cc2b)
- [service.advertise_options.advertise_on_public.port.port.info](data-sources--workload--reference--group-015.md#canonical-c0eab234681276b5b679fec6a1889c6499c45b2573ac84626157885203a83704)
- service.advertise_options.advertise_on_public.port.port.info.same_as_port

<a id="canonical-6e50cbdc253e074513ed8b63f1b986732f3ddeb1e04cac1547246e386ee3a4df"></a>

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

<a id="canonical-c2260881012afaf93b731c55d5729ba33e1a71d30986a5f936cc10aeaa855604"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.port.info.same_as_port / 5abfd1c96088 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-097984955a120e13b2372e639b7f112f542e7f7be0b62bb1cb28129dc9f16e4a"></a>

## Next pages — service.advertise_options.advertise_on_public.port.port.info.same_as_port / 5abfd1c96088 / 4

- [service.advertise_options.advertise_on_public.port.port.info](data-sources--workload--reference--group-015.md#canonical-c0eab234681276b5b679fec6a1889c6499c45b2573ac84626157885203a83704)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-0e62a645135c3b8ccf8743b299e24a9aa87050bc8ebdb2324540646b774daebe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-194ab8c75da7670890ad70c142360c048ba699b1b3a9373f72366f642eac21a8"></a>

## service.advertise_options.advertise_on_public.port.tcp_loadbalancer — service.advertise_options.advertise_on_public.port.tcp_loadbalancer / 819547eeaeb8 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-a12bf02272720e68e73b00969912e4531958e55ac3c63866089f14cf774f7a3e)
- service.advertise_options.advertise_on_public.port.tcp_loadbalancer

<a id="canonical-9e88b4d2119c588d6feb7b750a6e942d66e996c5d3799435f4469d28635061e7"></a>

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

<a id="canonical-36cfe31defee9f62d79f769b341b4a88f4a7e78cce709f1a44bd65e80c93d69f"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.tcp_loadbalancer / 819547eeaeb8 / 3

<a id="canonical-d15937447a36f66821440a54d05b65024a057869178775dcd2206636129c8f98"></a>

<a id="canonical-e378c99e4e05f17e0d95c5fe5975afeadc0366c8f9d16eb1272408f0ac43df89"></a>

## domains property — service.advertise_options.advertise_on_public.port.tcp_loadbalancer / 819547eeaeb8 / 4

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

<a id="canonical-89aaa3c2aef72c599c3643f3fb0e554cf63a5d44011ea120feb76e0768c73629"></a>

<a id="canonical-08a1b010a509e8edbbbe03fb52c86e4d400649e8ce6f0eec0e2e6351fde73d24"></a>

## with_sni property — service.advertise_options.advertise_on_public.port.tcp_loadbalancer / 819547eeaeb8 / 5

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

<a id="canonical-2ded209c9e78341eae45adad13d788153b33f49d435e7328cc3f5e59dfc86c21"></a>

## Next pages — service.advertise_options.advertise_on_public.port.tcp_loadbalancer / 819547eeaeb8 / 6

- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-a12bf02272720e68e73b00969912e4531958e55ac3c63866089f14cf774f7a3e)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-f3f8ccca96d163616e3cb5e7b42b4d261cfcc700a0e926fbe177a14cd8f2c3c0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-40d54dfb8ccb0f60e825a4de4b0553a29cabe863f04d5b694d2bfcdd4d42b9d3"></a>

## service.advertise_options.do_not_advertise — service.advertise_options.do_not_advertise / 229d66aac83b / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- service.advertise_options.do_not_advertise

<a id="canonical-511465f8bf87ab14c988a625468223e0657ddc52b930638d0fccf3702cd5b0a9"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for do not advertise.

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

<a id="canonical-220ae19efdbbe76aeb781f156e97fc1c37af0b579e27af8fe0e853d684194524"></a>

## Direct properties — service.advertise_options.do_not_advertise / 229d66aac83b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b263eb37e1879f373adc5b1895069973458e64b2c9cf2e704bd414b96a65216d"></a>

## Next pages — service.advertise_options.do_not_advertise / 229d66aac83b / 4

- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-db082f6c21bd52187fa7fe03d4ad630583fc4dd7bf94198dff5ab8dd3a742d24"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5d799e3eb7b627e448470c21c9b29384d65e4a69ed5701a610566d3ed4a5fbc9"></a>

## service.configuration — service.configuration / daf0e87e5330 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- service.configuration

<a id="canonical-7f9a500cbcc79767eb1bac4ea00ecaa5544301e3e2b1756783ce157a78564b4d"></a>

Type: `"single"`. Computed.

Configuration parameters of the workload.

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

<a id="canonical-b38f14e33141ec6180703b456ead9b9aa0c8fc404d18fab7f49ca37874a0724f"></a>

## Direct properties — service.configuration / daf0e87e5330 / 3

- [parameters](data-sources--workload--reference--group-015.md#canonical-061784d7cfc961779438fb52580bf0762c7daecaba80550074e7b13e08be3a6f): complete subsection reference.

<a id="canonical-91e3f9409064df52e28f24ce9e8e8a7767c8f84dca2b6f6d244c04365c079e13"></a>

## Next pages — service.configuration / daf0e87e5330 / 4

- [service.configuration.parameters](data-sources--workload--reference--group-015.md#canonical-061784d7cfc961779438fb52580bf0762c7daecaba80550074e7b13e08be3a6f)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-061784d7cfc961779438fb52580bf0762c7daecaba80550074e7b13e08be3a6f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2865a92b169b38c40af4d3229a6527a6bed90b8f04599fd051e8489e1359eaec"></a>

## service.configuration.parameters — service.configuration.parameters / f007a6582c75 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.configuration](data-sources--workload--reference--group-015.md#canonical-db082f6c21bd52187fa7fe03d4ad630583fc4dd7bf94198dff5ab8dd3a742d24)
- service.configuration.parameters

<a id="canonical-5bcf161a6b33704ce52c1190225fe3576f1719efa2217b56facf12a1e0b8df6d"></a>

Type: `"list"`. Computed.

Parameters. Parameters for the workload.

Upstream description:

Parameters for the workload.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-bb46428d2d71f40d1a2586955df6c591b7dcd3d4c743c3c2442071c93afc864b"></a>

## Direct properties — service.configuration.parameters / f007a6582c75 / 3

- [env_var](data-sources--workload--reference--group-015.md#canonical-0d0961a00c38e4a2f4e40eedda0ffd0b4f35cd3cbe1b6ba84372736a965a1ef7): complete subsection reference.

- [file](data-sources--workload--reference--group-015.md#canonical-fef138edae9e5e07cd09d3c62dc15e4f204703ad1235fd3a62e66985f30e0cb5): complete subsection reference.

<a id="canonical-203952f7103e83994455ac329664a24a841a67d60253a493ad68769cfd6820d9"></a>

## Next pages — service.configuration.parameters / f007a6582c75 / 4

- [service.configuration.parameters.env_var](data-sources--workload--reference--group-015.md#canonical-0d0961a00c38e4a2f4e40eedda0ffd0b4f35cd3cbe1b6ba84372736a965a1ef7)
- [service.configuration.parameters.file](data-sources--workload--reference--group-015.md#canonical-fef138edae9e5e07cd09d3c62dc15e4f204703ad1235fd3a62e66985f30e0cb5)
- [service.configuration](data-sources--workload--reference--group-015.md#canonical-db082f6c21bd52187fa7fe03d4ad630583fc4dd7bf94198dff5ab8dd3a742d24)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-0d0961a00c38e4a2f4e40eedda0ffd0b4f35cd3cbe1b6ba84372736a965a1ef7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d6016f1ec80f341e1d96f83506c32a8727798635dd8745b1b2564d1aba985165"></a>

## service.configuration.parameters.env_var — service.configuration.parameters.env_var / 932552701e0e / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.configuration](data-sources--workload--reference--group-015.md#canonical-db082f6c21bd52187fa7fe03d4ad630583fc4dd7bf94198dff5ab8dd3a742d24)
- [service.configuration.parameters](data-sources--workload--reference--group-015.md#canonical-061784d7cfc961779438fb52580bf0762c7daecaba80550074e7b13e08be3a6f)
- service.configuration.parameters.env_var

<a id="canonical-0fe25ab637a5f21a4dcd20e9bede84107cbc75096f519d093d19ca6b102c38dc"></a>

Type: `"single"`. Computed.

Environment Variable. Environment Variable.

Upstream description:

Environment Variable.

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

<a id="canonical-30155eadbb25f4c069ae793c4f572b3b41e1195cc2407232f569da347aac6356"></a>

## Direct properties — service.configuration.parameters.env_var / 932552701e0e / 3

<a id="canonical-f3c8bea369a6aa7a3a7bf97f909d5457c8e72169aecbef01c131e90b00c42c3a"></a>

<a id="canonical-4ddfcc7920943f82bbb1c5fdb2a994956d570ee80aff685e69f41bd1219296c9"></a>

## name property — service.configuration.parameters.env_var / 932552701e0e / 4

Type: `"string"`. Computed.

Name. Name of Environment Variable.

Upstream description:

Name of Environment Variable.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
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
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-9d29206a1d2cf49d174481df3fa2f5a30229a7effd961d3de8802b4dfe39d165"></a>

<a id="canonical-276e91cfa4536738d9bd4ab621acb3a74d9afeafc820606d0356b009c8afbcab"></a>

## value property — service.configuration.parameters.env_var / 932552701e0e / 5

Type: `"string"`. Computed.

Value. Value of Environment Variable.

Upstream description:

Value of Environment Variable.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-08c4375bba03b8db04027cb39617ae7f4b3ed4ff68c4c3f873fd3bfe9e81cf81"></a>

## Next pages — service.configuration.parameters.env_var / 932552701e0e / 6

- [service.configuration.parameters](data-sources--workload--reference--group-015.md#canonical-061784d7cfc961779438fb52580bf0762c7daecaba80550074e7b13e08be3a6f)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-fef138edae9e5e07cd09d3c62dc15e4f204703ad1235fd3a62e66985f30e0cb5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-53527075c721ace7f5e3f89e50bd7ec0758fd9763dbd8ec9c7342cb47ffb4346"></a>

## service.configuration.parameters.file — service.configuration.parameters.file / 8b4a3abe5396 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.configuration](data-sources--workload--reference--group-015.md#canonical-db082f6c21bd52187fa7fe03d4ad630583fc4dd7bf94198dff5ab8dd3a742d24)
- [service.configuration.parameters](data-sources--workload--reference--group-015.md#canonical-061784d7cfc961779438fb52580bf0762c7daecaba80550074e7b13e08be3a6f)
- service.configuration.parameters.file

<a id="canonical-e2589e1b0bf7a33beb15fc4f52d242811730e3ef937e43a32a250ec116232f49"></a>

Type: `"single"`. Computed.

Configuration File. Configuration File for the workload.

Upstream description:

Configuration File for the workload.

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

<a id="canonical-9b83607993f1d1c2038803ff5322147d79b3c89bc5654fb668b9c60679b66a3e"></a>

## Direct properties — service.configuration.parameters.file / 8b4a3abe5396 / 3

<a id="canonical-5329a2123463d43cd88ba51878133e04c0643c0b583e59367cf14af5fbe9f094"></a>

<a id="canonical-1a8bbdf1fe3d66c294ce695938cc7411c51d418d7b0e3f39f0e2ea67fd5fc865"></a>

## data property — service.configuration.parameters.file / 8b4a3abe5396 / 4

Type: `"string"`. Computed.

Data. File data

Upstream description:

File data

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 16384,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 16384,
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
    "ves.io.schema.rules.string.max_len": "16384",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "16384",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [mount](data-sources--workload--reference--group-015.md#canonical-6c99ae8299232afa552427f289c61c44462a3f18a5f1565ce1f3363f5b2a53ec): complete subsection reference.

<a id="canonical-21e1dc8a6c36d31081b67510a0655ac46c45320caaa9056b94c5d64517920128"></a>

<a id="canonical-6d6d247d52177cd00376da3f3308e025db74880e57c5648b9f206ea5de77fdbc"></a>

## name property — service.configuration.parameters.file / 8b4a3abe5396 / 5

Type: `"string"`. Computed.

Name. Name of the file.

Upstream description:

Name of the file.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
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
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-44f41d2f950b89ac1a4c08954d3fe1d8dada2a82b697ab02531028a05d4867bb"></a>

<a id="canonical-aab6cab2b9f0e716472da614d82fb5752ee3a9a676a4e9147d2911dba494ac29"></a>

## volume_name property — service.configuration.parameters.file / 8b4a3abe5396 / 6

Type: `"string"`. Computed.

Volume Name. Name of the Volume.

Upstream description:

Name of the Volume.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-842a6f936dc7cad7439cf7192f87ab206e98dffcfbf47c51965fb20bf5403cc5"></a>

## Next pages — service.configuration.parameters.file / 8b4a3abe5396 / 7

- [service.configuration.parameters.file.mount](data-sources--workload--reference--group-015.md#canonical-6c99ae8299232afa552427f289c61c44462a3f18a5f1565ce1f3363f5b2a53ec)
- [service.configuration.parameters](data-sources--workload--reference--group-015.md#canonical-061784d7cfc961779438fb52580bf0762c7daecaba80550074e7b13e08be3a6f)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-6c99ae8299232afa552427f289c61c44462a3f18a5f1565ce1f3363f5b2a53ec"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8ae1b276f8897218d113c08d74d455c535036373b79bd05044b1b9f2785a6839"></a>

## service.configuration.parameters.file.mount — service.configuration.parameters.file.mount / 73e0c06f91af / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.configuration](data-sources--workload--reference--group-015.md#canonical-db082f6c21bd52187fa7fe03d4ad630583fc4dd7bf94198dff5ab8dd3a742d24)
- [service.configuration.parameters](data-sources--workload--reference--group-015.md#canonical-061784d7cfc961779438fb52580bf0762c7daecaba80550074e7b13e08be3a6f)
- [service.configuration.parameters.file](data-sources--workload--reference--group-015.md#canonical-fef138edae9e5e07cd09d3c62dc15e4f204703ad1235fd3a62e66985f30e0cb5)
- service.configuration.parameters.file.mount

<a id="canonical-fd6a6ca3d0faa21e4c2f91f304d9aff2379b3d23cbb768560a9d8c1c432b0be2"></a>

Type: `"single"`. Computed.

Volume mount describes how volume is mounted inside a workload.

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

<a id="canonical-8a44f0d2ada81b6aa928a2703b5cd5f69d4b53cc8587d3e39035aa59580ece04"></a>

## Direct properties — service.configuration.parameters.file.mount / 73e0c06f91af / 3

<a id="canonical-adeb6aa208375ebfaaae142011e516b6134aa7ee1612a62adee7961a708c8351"></a>

<a id="canonical-e15dffc4904563ceb2731a40f79ebdacfccbc22ffcbe159fec45bd6089ff5308"></a>

## mode property — service.configuration.parameters.file.mount / 73e0c06f91af / 4

Type: `"string"`. Computed.

\[Enum: VOLUME\_MOUNT\_READ\_ONLY|VOLUME\_MOUNT\_READ\_WRITE\] Mode in which the volume should be
mounted to the workload - VOLUME\_MOUNT\_READ\_ONLY: ReadOnly Mount the volume in read-only mode -
VOLUME\_MOUNT\_READ\_WRITE: Read Write Mount the volume in read-write mode. Possible values are
\`VOLUME\_MOUNT\_READ\_ONLY\`, \`VOLUME\_MOUNT\_READ\_WRITE\`. Defaults to
\`VOLUME\_MOUNT\_READ\_ONLY\`.

Upstream description:

Mode in which the volume should be mounted to the workload

&#8203;- VOLUME\_MOUNT\_READ\_ONLY: ReadOnly

Mount the volume in read-only mode &#8203;- VOLUME\_MOUNT\_READ\_WRITE: Read Write

Mount the volume in read-write mode.

Receipt-pinned upstream constraints:

```json
{
  "default": "VOLUME_MOUNT_READ_ONLY",
  "enum": [
    "VOLUME_MOUNT_READ_ONLY",
    "VOLUME_MOUNT_READ_WRITE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-c5ffc3b662e51fca7cb9fb1065b3b0a8da8db51afecaf9b745d19170eaae38d8"></a>

<a id="canonical-681f7c8d97013360c58cf026b9832c62cdfaf61d19738e7cc810ad649da7a1f5"></a>

## mount_path property — service.configuration.parameters.file.mount / 73e0c06f91af / 5

Type: `"string"`. Computed.

Path within the workload container at which the volume should be mounted. Must not contain ':'.

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
    "pattern": "^[^:]*$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.pattern": "^[^:]*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.pattern": "^[^:]*$"
  }
}
```

<a id="canonical-e9daa9d81461a5f8965f2549bb8623d24fbbc97fc66726f63867c3725b1f526b"></a>

<a id="canonical-053ce4d0e73e1bdf6cf5388f9279bf92c45cb08c1f98e3b864f9e10498852d65"></a>

## sub_path property — service.configuration.parameters.file.mount / 73e0c06f91af / 6

Type: `"string"`. Computed.

Path within the volume from which the workload's volume should be mounted. Defaults to '' (volume's
root).

Upstream description:

Path within the volume from which the workload's volume should be mounted. Defaults to "" (volume's
root).

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-5bacdc4c7b4843ae23f6cb33d0f8f63cd50aaf06c4fdb8dad8db9bd6e6c755a0"></a>

## Next pages — service.configuration.parameters.file.mount / 73e0c06f91af / 7

- [service.configuration.parameters.file](data-sources--workload--reference--group-015.md#canonical-fef138edae9e5e07cd09d3c62dc15e4f204703ad1235fd3a62e66985f30e0cb5)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-f3e5d811f1c619d3bb1afde6648794fd06112ea88c6a7458c78327031b9486cf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6efe2ca92f9485105c7f2448b6e2b96308a0d2417708f5fe824132f75bdeadac"></a>

## service.containers — service.containers / ec021bcd7142 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- service.containers

<a id="canonical-4c4f7595e3cb771951795878b2ba46d560a6ed8a3cd642e5f1f101a26b0c1a4d"></a>

Type: `"list"`. Computed.

Containers. Containers to use for service.

Upstream description:

Containers to use for service.

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

<a id="canonical-18d914c1da4b86c149e319ba88f7ab33ca702ef0e5b342b0df6e8d60aff82a88"></a>

## Direct properties — service.containers / ec021bcd7142 / 3

<a id="canonical-697b47f2bae3f94fa4f55603abc51f560640486b2341a72692724186fa1ea9f2"></a>

<a id="canonical-1c4c18901d1122c52bdd79b89f1b98663410289a01cb591d6c17891180a4bc03"></a>

## args property — service.containers / ec021bcd7142 / 4

Type: `["list", "string"]`. Computed.

Arguments to the entrypoint. Overrides the docker image's CMD.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

<a id="canonical-8977ed5e2cfafc70b7c86ea4e11950dd0c347eb8917091824b980c3eeec784dc"></a>

<a id="canonical-8127b2f625de4b53a056a26331acc035cf56086aa60f5634f00b1c8dba80651a"></a>

## command property — service.containers / ec021bcd7142 / 5

Type: `["list", "string"]`. Computed.

Command to execute. Overrides the docker image's ENTRYPOINT.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

- [custom_flavor](data-sources--workload--reference--group-015.md#canonical-28abd6179a6d01293463d9f81b7d0df3616cc0ac18ba4d5d642e02b17be3452c): complete subsection reference.

- [default_flavor](data-sources--workload--reference--group-015.md#canonical-af11a4e2315a0604cbd7c70811da8372af27fd7f4598381bfd426b1742df1ded): complete subsection reference.

<a id="canonical-860e43884b01a2a9dc98ee264ff6ba764ed30def518f2f99893c8c3802bb2f9f"></a>

<a id="canonical-597ea473e22fab97059a46ec947a4c52a75da64e3e037bb6bb3fa72626a109d3"></a>

## flavor property — service.containers / ec021bcd7142 / 6

Type: `"string"`. Computed.

\[Enum:
CONTAINER\_FLAVOR\_TYPE\_TINY|CONTAINER\_FLAVOR\_TYPE\_MEDIUM|CONTAINER\_FLAVOR\_TYPE\_LARGE\]
Container Flavor type - CONTAINER\_FLAVOR\_TYPE\_TINY: Tiny Tiny containers have limit of 0.1 vCPU
and 256 MiB (mebibyte) memory - CONTAINER\_FLAVOR\_TYPE\_MEDIUM: Medium Medium containers have limit
of 0.25 vCPU and 512 MiB (mebibyte) memory - CONTAINER\_FLAVOR\_TYPE\_LARGE: Large Large containers
have.. Possible values are \`CONTAINER\_FLAVOR\_TYPE\_TINY\`, \`CONTAINER\_FLAVOR\_TYPE\_MEDIUM\`,
\`CONTAINER\_FLAVOR\_TYPE\_LARGE\`. Defaults to \`CONTAINER\_FLAVOR\_TYPE\_TINY\`.

Upstream description:

Container Flavor type

&#8203;- CONTAINER\_FLAVOR\_TYPE\_TINY: Tiny

Tiny containers have limit of 0.1 vCPU and 256 MiB (mebibyte) memory &#8203;-
CONTAINER\_FLAVOR\_TYPE\_MEDIUM: Medium

Medium containers have limit of 0.25 vCPU and 512 MiB (mebibyte) memory &#8203;-
CONTAINER\_FLAVOR\_TYPE\_LARGE: Large

Large containers have limit of 1 vCPU and 2048 MiB (mebibyte) memory.

Receipt-pinned upstream constraints:

```json
{
  "default": "CONTAINER_FLAVOR_TYPE_TINY",
  "enum": [
    "CONTAINER_FLAVOR_TYPE_TINY",
    "CONTAINER_FLAVOR_TYPE_MEDIUM",
    "CONTAINER_FLAVOR_TYPE_LARGE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [image](data-sources--workload--reference--group-015.md#canonical-8deafb0d6e4d8202fe93e341fde40d23121e50940e399bfbec467b76cccbc67f): complete subsection reference.

<a id="canonical-387ce125548841f44b1647e3812355809b500cfeca633973b913b735c2343d4b"></a>

<a id="canonical-fca96dcefc49ba0b7b1fa14dfeb6f70e163f0f755e5bc96dc0d39a6d14d53dab"></a>

## init_container property — service.containers / ec021bcd7142 / 7

Type: `"bool"`. Computed.

Specialized container that runs before application container and runs to completion.

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

- [liveness_check](data-sources--workload--reference--group-015.md#canonical-9359bc49c0401d6ea98bd188449578db7855267e305be7d84d0d7b8e1137b1d8): complete subsection reference.

<a id="canonical-ff7ec9fb46528d0c60ee77e127197f495e84773bfc80bd0dd7e7bef7cd2fe031"></a>

<a id="canonical-6877507c7bd00dd47e55cb7974e3b8233305f97ecb02ffd0351356e239ebbca4"></a>

## name property — service.containers / ec021bcd7142 / 8

Type: `"string"`. Computed.

Name. Name of the container.

Upstream description:

Name of the container.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
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
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [readiness_check](data-sources--workload--reference--group-015.md#canonical-59a7c74fa8e3b4ee945cdef67d7014c8eb4b7fd15e9d23ac48180839bbbcd306): complete subsection reference.

<a id="canonical-631474bea1ae9fc6cb2b2931a76fc0a274cadd1b385991a4084dbaf45f979610"></a>

## Next pages — service.containers / ec021bcd7142 / 9

- [service.containers.custom_flavor](data-sources--workload--reference--group-015.md#canonical-28abd6179a6d01293463d9f81b7d0df3616cc0ac18ba4d5d642e02b17be3452c)
- [service.containers.default_flavor](data-sources--workload--reference--group-015.md#canonical-af11a4e2315a0604cbd7c70811da8372af27fd7f4598381bfd426b1742df1ded)
- [service.containers.image](data-sources--workload--reference--group-015.md#canonical-8deafb0d6e4d8202fe93e341fde40d23121e50940e399bfbec467b76cccbc67f)
- [service.containers.liveness_check](data-sources--workload--reference--group-015.md#canonical-9359bc49c0401d6ea98bd188449578db7855267e305be7d84d0d7b8e1137b1d8)
- [service.containers.readiness_check](data-sources--workload--reference--group-015.md#canonical-59a7c74fa8e3b4ee945cdef67d7014c8eb4b7fd15e9d23ac48180839bbbcd306)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-28abd6179a6d01293463d9f81b7d0df3616cc0ac18ba4d5d642e02b17be3452c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-270f6372723ae497232ca734892491b4df341f941a3fcb4972c844696c2b6bfe"></a>

## service.containers.custom_flavor — service.containers.custom_flavor / 6048dacb421d / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.containers](data-sources--workload--reference--group-015.md#canonical-f3e5d811f1c619d3bb1afde6648794fd06112ea88c6a7458c78327031b9486cf)
- service.containers.custom_flavor

<a id="canonical-4bd9f4db5e4193baec17233b28890dc58e3155b0b0ec3ae66e1c5f129f85fc29"></a>

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

<a id="canonical-373496355ac207e08cd811a9109a2a2d670f6870ca249b6189db6dfe33fbbbc7"></a>

## Direct properties — service.containers.custom_flavor / 6048dacb421d / 3

<a id="canonical-82fc4fafb8d195062c181b580716abdb581cf7c63e3a782a8bcae3e29a3bc40e"></a>

<a id="canonical-9e85e78adee5a858f4a703df608d3e422197109d1ab4c5085a4447404c10e70c"></a>

## name property — service.containers.custom_flavor / 6048dacb421d / 4

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

<a id="canonical-8ed2cec179c54205b5160ddbf726967618f9f4b7ecad0a19fb02e66017c5368c"></a>

<a id="canonical-76158ffa437c8ddba32043c5a086389cb87e030889525ee0475fe27608ebf0b6"></a>

## namespace property — service.containers.custom_flavor / 6048dacb421d / 5

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

<a id="canonical-b1f0727ffffd460a9df712927aa582f85618bd3bed97c5071ad89e31264fb967"></a>

<a id="canonical-1a6c8c8331f1b0616cbac4309fc3b7271385aa33b7e69b992b00e44880051ddb"></a>

## tenant property — service.containers.custom_flavor / 6048dacb421d / 6

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

<a id="canonical-7e90813b65953240b64cc991f52473ddc7f5d5aef8d562418b651f29cb2b7abd"></a>

## Next pages — service.containers.custom_flavor / 6048dacb421d / 7

- [service.containers](data-sources--workload--reference--group-015.md#canonical-f3e5d811f1c619d3bb1afde6648794fd06112ea88c6a7458c78327031b9486cf)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-af11a4e2315a0604cbd7c70811da8372af27fd7f4598381bfd426b1742df1ded"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-14e81a322b888f7fd892b1f84d49e99a187bc054fd8cf32bf90fe8e1d1dd206a"></a>

## service.containers.default_flavor — service.containers.default_flavor / 2bde63711a2f / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.containers](data-sources--workload--reference--group-015.md#canonical-f3e5d811f1c619d3bb1afde6648794fd06112ea88c6a7458c78327031b9486cf)
- service.containers.default_flavor

<a id="canonical-52768005cc15afbdc396841353fc07b7ae476196f5836a900d3d48e89f388e30"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default flavor.

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

<a id="canonical-d32ffd72d3899e00ccf1675160b8a6b9fdbe308df2fffd42b4cdb9510b51d86c"></a>

## Direct properties — service.containers.default_flavor / 2bde63711a2f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-902dbfa53ab4b128374044623941598914f7e69ce2099cad4c87968a202381a8"></a>

## Next pages — service.containers.default_flavor / 2bde63711a2f / 4

- [service.containers](data-sources--workload--reference--group-015.md#canonical-f3e5d811f1c619d3bb1afde6648794fd06112ea88c6a7458c78327031b9486cf)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-8deafb0d6e4d8202fe93e341fde40d23121e50940e399bfbec467b76cccbc67f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-efad28a16e89f665ad15ed01104a3eb5bdd0722df63a10d276a54aaf628f4d44"></a>

## service.containers.image — service.containers.image / 03242f25f3dc / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.containers](data-sources--workload--reference--group-015.md#canonical-f3e5d811f1c619d3bb1afde6648794fd06112ea88c6a7458c78327031b9486cf)
- service.containers.image

<a id="canonical-8dcd7abb437d2d285b7fc8c6c1c5e533800cdb42272ab9fc40bbbafa2ae6345f"></a>

Type: `"single"`. Computed.

ImageType configures the image to use, how to pull the image, and the associated secrets to use if
any.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-registry_choice": "[\"container_registry\",\"public\"]"
}
```

<a id="canonical-4c33b80725ee688478e8931269261a001552d131df06b1839d9d5d9be581b033"></a>

## Direct properties — service.containers.image / 03242f25f3dc / 3

- [container_registry](data-sources--workload--reference--group-015.md#canonical-20bae4d7083ed3019603d04058da356d9da6b3c1162b6cab6ec8e2dbd0241eb0): complete subsection reference.

<a id="canonical-a4fbb7c1275ade6423520f5ac314dfe8b28f9b5eecd877db6aaf8e60cb21be83"></a>

<a id="canonical-b11d5fde648d3cce888b14ac453d9862bd3ddaed5880b2a325a835bc73839743"></a>

## name property — service.containers.image / 03242f25f3dc / 4

Type: `"string"`. Computed.

Name is a container image which are usually given a name such as alpine, ubuntu, or
quay.I/O/etcd:0.13. The format is registry/image:tag or registry/image@image-digest. If registry is
not specified, the Docker public registry is assumed.

Upstream description:

Name is a container image which are usually given a name such as alpine, ubuntu, or
quay.I/O/etcd:0.13. The format is registry/image:tag or registry/image@image-digest. If registry is
not specified, the Docker public registry is assumed. If tag is not specified, latest is assumed.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
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
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [public](data-sources--workload--reference--group-015.md#canonical-47092a70d8d29f378e6f5ec1d2a20c29fbe8ce7c0440c9a0a8f0b73878ad60e0): complete subsection reference.

<a id="canonical-e816852d39bba31dc4a3e3e340eecb374e39c6016678775f14e6904732e64cd2"></a>

<a id="canonical-5b6e0ccc2911f2228acae7e176ad5237b69f2a4314837374ce83a01b10bd3c29"></a>

## pull_policy property — service.containers.image / 03242f25f3dc / 5

Type: `"string"`. Computed.

\[Enum:
IMAGE\_PULL\_POLICY\_DEFAULT|IMAGE\_PULL\_POLICY\_IF\_NOT\_PRESENT|IMAGE\_PULL\_POLICY\_ALWAYS|IMAGE\_PULL\_POLICY\_NEVER\]
Image pull policy type enumerates the policy choices to use for pulling the image prior to starting
the workload - IMAGE\_PULL\_POLICY\_DEFAULT: Default Default will always pull image if :latest tag
is specified in image name. If :latest tag is not specified in image name, it will pull image only..
Possible values are \`IMAGE\_PULL\_POLICY\_DEFAULT\`, \`IMAGE\_PULL\_POLICY\_IF\_NOT\_PRESENT\`,
\`IMAGE\_PULL\_POLICY\_ALWAYS\`, \`IMAGE\_PULL\_POLICY\_NEVER\`. Defaults to
\`IMAGE\_PULL\_POLICY\_DEFAULT\`.

Upstream description:

Image pull policy type enumerates the policy choices to use for pulling the image prior to starting
the workload

&#8203;- IMAGE\_PULL\_POLICY\_DEFAULT: Default

Default will always pull image if :latest tag is specified in image name. If :latest tag is not
specified in image name, it will pull image only if it does not already exist on the node &#8203;-
IMAGE\_PULL\_POLICY\_IF\_NOT\_PRESENT: IfNotPresent

Only pull the image if it does not already exist on the node &#8203;- IMAGE\_PULL\_POLICY\_ALWAYS:
Always

Always pull the image &#8203;- IMAGE\_PULL\_POLICY\_NEVER: Never

Never pull the image.

Receipt-pinned upstream constraints:

```json
{
  "default": "IMAGE_PULL_POLICY_DEFAULT",
  "enum": [
    "IMAGE_PULL_POLICY_DEFAULT",
    "IMAGE_PULL_POLICY_IF_NOT_PRESENT",
    "IMAGE_PULL_POLICY_ALWAYS",
    "IMAGE_PULL_POLICY_NEVER"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-188084626f54a8ad4bc1b10d9eb5f80cb3300d006e5e02c967a2c8a19000064b"></a>

## Next pages — service.containers.image / 03242f25f3dc / 6

- [service.containers.image.container_registry](data-sources--workload--reference--group-015.md#canonical-20bae4d7083ed3019603d04058da356d9da6b3c1162b6cab6ec8e2dbd0241eb0)
- [service.containers.image.public](data-sources--workload--reference--group-015.md#canonical-47092a70d8d29f378e6f5ec1d2a20c29fbe8ce7c0440c9a0a8f0b73878ad60e0)
- [service.containers](data-sources--workload--reference--group-015.md#canonical-f3e5d811f1c619d3bb1afde6648794fd06112ea88c6a7458c78327031b9486cf)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-20bae4d7083ed3019603d04058da356d9da6b3c1162b6cab6ec8e2dbd0241eb0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b92b186fc9ac7965e9f17ea3066b90363e3ae3c05b31c9b77c9e5fdd41663cf1"></a>

## service.containers.image.container_registry — service.containers.image.container_registry / 1911c357cc6e / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.containers](data-sources--workload--reference--group-015.md#canonical-f3e5d811f1c619d3bb1afde6648794fd06112ea88c6a7458c78327031b9486cf)
- [service.containers.image](data-sources--workload--reference--group-015.md#canonical-8deafb0d6e4d8202fe93e341fde40d23121e50940e399bfbec467b76cccbc67f)
- service.containers.image.container_registry

<a id="canonical-918d380450ac4b0c0be2a4b90189cb1503ac1f094ca7062146db1505d4bc6ab3"></a>

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

<a id="canonical-cd8fd7d65590f1521e252ceff7968a4500cab43126edb6c30181a74859668e35"></a>

## Direct properties — service.containers.image.container_registry / 1911c357cc6e / 3

<a id="canonical-62744c31264ab7dcc9213008623be2fcbe247224d68f29902b0ef800079c3034"></a>

<a id="canonical-8930669a316be0d19947df13637b8174c7172df9da341932290624549dd74dfc"></a>

## name property — service.containers.image.container_registry / 1911c357cc6e / 4

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

<a id="canonical-4e55869a2a08ca0e1b51c763217ccd2af2be456dddbb6cbde5bb2aba42118fab"></a>

<a id="canonical-9b82caf691d6e2302b6316e8dc087972fdad83321877e08fc012b82a97eb9b6b"></a>

## namespace property — service.containers.image.container_registry / 1911c357cc6e / 5

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

<a id="canonical-f9ad33640515467ef60329d9a9db4b549f3b5bd83a288d7d7d24404a3a9b688b"></a>

<a id="canonical-5423221629eaf68acf8d1350d8ce3121be11d41735f77e7f3f150a3f0e5f885a"></a>

## tenant property — service.containers.image.container_registry / 1911c357cc6e / 6

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

<a id="canonical-afdb3e4b1313b3e8c63e3d992c961da04ce79af22bc64020c1d94cf5849c6bff"></a>

## Next pages — service.containers.image.container_registry / 1911c357cc6e / 7

- [service.containers.image](data-sources--workload--reference--group-015.md#canonical-8deafb0d6e4d8202fe93e341fde40d23121e50940e399bfbec467b76cccbc67f)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-47092a70d8d29f378e6f5ec1d2a20c29fbe8ce7c0440c9a0a8f0b73878ad60e0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b08d4dd808987f7bfdc37a015a171c3da2aea8726102fbb8722bed4d06968876"></a>

## service.containers.image.public — service.containers.image.public / 34209e27fac3 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.containers](data-sources--workload--reference--group-015.md#canonical-f3e5d811f1c619d3bb1afde6648794fd06112ea88c6a7458c78327031b9486cf)
- [service.containers.image](data-sources--workload--reference--group-015.md#canonical-8deafb0d6e4d8202fe93e341fde40d23121e50940e399bfbec467b76cccbc67f)
- service.containers.image.public

<a id="canonical-e591b80d5ab64d95db6d1dc0e82979165e923c4f04a1fbf3aed734cf93535889"></a>

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

<a id="canonical-99242bcdbec29c7e7ff0e0131ef332cde3edd1d54e9a25977ddd98178a068e88"></a>

## Direct properties — service.containers.image.public / 34209e27fac3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7aee1bddaa6caf325a8e49f8e9959c61e7dabbd7786f74558980cd873faf85ff"></a>

## Next pages — service.containers.image.public / 34209e27fac3 / 4

- [service.containers.image](data-sources--workload--reference--group-015.md#canonical-8deafb0d6e4d8202fe93e341fde40d23121e50940e399bfbec467b76cccbc67f)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-9359bc49c0401d6ea98bd188449578db7855267e305be7d84d0d7b8e1137b1d8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f65fbf6ba971c0ee25a40681fe5fece1a11b3f291509aa599625d6c0b01526e4"></a>

## service.containers.liveness_check — service.containers.liveness_check / 83f099b5f198 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.containers](data-sources--workload--reference--group-015.md#canonical-f3e5d811f1c619d3bb1afde6648794fd06112ea88c6a7458c78327031b9486cf)
- service.containers.liveness_check

<a id="canonical-c22a7d636b840d908774187e5256b7bec75f7dde5cdc011d22091736ba091122"></a>

Type: `"single"`. Computed.

HealthCheckType describes a health check to be performed against a container to determine whether it
has started up or is alive or ready to receive traffic.

Upstream description:

HealthCheckType describes a health check to be performed against a container to determine whether it
has started up or is alive or ready to receive traffic.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-health_check_choice": "[\"exec_health_check\",\"http_health_check\",\"tcp_health_check\"]"
}
```

<a id="canonical-ffeda275ebfa2a7af7220c29fb9c42d5cf4f1c7c9f3de76a2c7b8e836b6cdcb5"></a>

## Direct properties — service.containers.liveness_check / 83f099b5f198 / 3

- [exec_health_check](data-sources--workload--reference--group-015.md#canonical-7eb98db615c0805aad960a3584a395c8b76c86cd2e36c6c93526d0cb63f6301a): complete subsection reference.

<a id="canonical-2ab33d1392ac7a9d1d36c63b04aed6d0b653ca841a6d78e46a2ac491b9e2e9eb"></a>

<a id="canonical-78dbb0f30a83b894608cb26d442b814f5f06f608960e4cb93215636aba42b282"></a>

## healthy_threshold property — service.containers.liveness_check / 83f099b5f198 / 4

Type: `"number"`. Computed.

Number of consecutive successful responses after having failed before declaring healthy. In other
words, this is the number of healthy health checks required before marking healthy. Note that during
startup and liveliness, only a single successful health check is required to mark a container..

Upstream description:

Number of consecutive successful responses after having failed before declaring healthy. In other
words, this is the number of healthy health checks required before marking healthy. Note that during
startup and liveliness, only a single successful health check is required to mark a container
healthy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

- [http_health_check](data-sources--workload--reference--group-015.md#canonical-5bb7a7f1b5475f8f80e3df328c843ecc0b5893208f40e5ebdf3c37d9d827bde0): complete subsection reference.

<a id="canonical-19f7c15905165733825308e418fdf09b8f2bbfd9ed5ba97a3ea5a34feae70e8b"></a>

<a id="canonical-dd26f84578fc708587a039f145b0ff8890bf98019472f858e3de76cf6ec0481e"></a>

## initial_delay property — service.containers.liveness_check / 83f099b5f198 / 5

Type: `"number"`. Computed.

Number of seconds after the container has started before health checks are initiated.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
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
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

<a id="canonical-3d90217e2149af5616c4e601d87675aa2db0115aa8b18e58935e7847dc3da9f4"></a>

<a id="canonical-6e042d309a525dd0414578da09c059109cdc6d96ad2df772bde25be4c123790b"></a>

## interval property — service.containers.liveness_check / 83f099b5f198 / 6

Type: `"number"`. Computed.

Time interval in seconds between two health check requests.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

- [tcp_health_check](data-sources--workload--reference--group-015.md#canonical-af2794efee684bc26c4e0d0e0a56bdfd087f12eaca7cc8db982a848a0188c811): complete subsection reference.

<a id="canonical-244e967cfb091dc2fcdb6606712a1c13c4be02073e2b7d237304044d85dae3d6"></a>

<a id="canonical-58f20681685941b1ab7cf36068f070a18ddb9906906b8b076cd5d544ece2ac73"></a>

## timeout property — service.containers.liveness_check / 83f099b5f198 / 7

Type: `"number"`. Computed.

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure.

Upstream description:

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

<a id="canonical-3792d64ae0e2a92a51acd5e6cd304f46b9bbbfbb56da8c0780db02f6f7b0f459"></a>

<a id="canonical-b3cb9125764ab29fff40970f781c1dfb9d15ea2af86cc0c11a876316e23eb180"></a>

## unhealthy_threshold property — service.containers.liveness_check / 83f099b5f198 / 8

Type: `"number"`. Computed.

Number of consecutive failed responses before declaring unhealthy. In other words, this is the
number of unhealthy health checks required before a container is marked unhealthy.

Upstream description:

Number of consecutive failed responses before declaring unhealthy. In other words, this is the
number of unhealthy health checks required before a container is marked unhealthy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

<a id="canonical-703443afea7c77d63590eaa7089dbbe8210baf128675981d78f2747530991446"></a>

## Next pages — service.containers.liveness_check / 83f099b5f198 / 9

- [service.containers.liveness_check.exec_health_check](data-sources--workload--reference--group-015.md#canonical-7eb98db615c0805aad960a3584a395c8b76c86cd2e36c6c93526d0cb63f6301a)
- [service.containers.liveness_check.http_health_check](data-sources--workload--reference--group-015.md#canonical-5bb7a7f1b5475f8f80e3df328c843ecc0b5893208f40e5ebdf3c37d9d827bde0)
- [service.containers.liveness_check.tcp_health_check](data-sources--workload--reference--group-015.md#canonical-af2794efee684bc26c4e0d0e0a56bdfd087f12eaca7cc8db982a848a0188c811)
- [service.containers](data-sources--workload--reference--group-015.md#canonical-f3e5d811f1c619d3bb1afde6648794fd06112ea88c6a7458c78327031b9486cf)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-7eb98db615c0805aad960a3584a395c8b76c86cd2e36c6c93526d0cb63f6301a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-49763dd1530baef42d7e18045817441234109e4fa0962b7492e68b5ab5e8780b"></a>

## service.containers.liveness_check.exec_health_check — service.containers.liveness_check.exec_health_check / 984b0135e475 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.containers](data-sources--workload--reference--group-015.md#canonical-f3e5d811f1c619d3bb1afde6648794fd06112ea88c6a7458c78327031b9486cf)
- [service.containers.liveness_check](data-sources--workload--reference--group-015.md#canonical-9359bc49c0401d6ea98bd188449578db7855267e305be7d84d0d7b8e1137b1d8)
- service.containers.liveness_check.exec_health_check

<a id="canonical-2906b1bfe1ce98d184e52f514feeb6d599784e0583227a2c950d3f63e0658d28"></a>

Type: `"single"`. Computed.

ExecHealthCheckType describes a health check based on 'run in container' action. Exit status of 0 is
treated as live/healthy and non-zero is unhealthy.

Upstream description:

ExecHealthCheckType describes a health check based on "run in container" action. Exit status of 0 is
treated as live/healthy and non-zero is unhealthy.

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

<a id="canonical-30ec3572a6d0ec5c5215e34900665d0a7de780eb548b552e5401f0c0135947b7"></a>

## Direct properties — service.containers.liveness_check.exec_health_check / 984b0135e475 / 3

<a id="canonical-7238756554dd23b91e853d4e42784adc54d72552391fb88516a67cc577a75c24"></a>

<a id="canonical-702c6e5d53644c9555d8834a6c822e29bb21ddabf27f55ccdd14793b8be8d2e3"></a>

## command property — service.containers.liveness_check.exec_health_check / 984b0135e475 / 4

Type: `["list", "string"]`. Computed.

Command is the command line to execute inside the container, the working directory for the command
is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a
shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to..

Upstream description:

Command is the command line to execute inside the container, the working directory for the command
is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a
shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to
explicitly call out to that shell.

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
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-211f62ebd9544533889b1b53f3cd07eba874a4788d5ca774a231d742ec793bf2"></a>

## Next pages — service.containers.liveness_check.exec_health_check / 984b0135e475 / 5

- [service.containers.liveness_check](data-sources--workload--reference--group-015.md#canonical-9359bc49c0401d6ea98bd188449578db7855267e305be7d84d0d7b8e1137b1d8)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-5bb7a7f1b5475f8f80e3df328c843ecc0b5893208f40e5ebdf3c37d9d827bde0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-45d9c3c62e5051a87df9ba716f1f0ddfec7d02e5d1c00aedc0bb35b306f2db95"></a>

## service.containers.liveness_check.http_health_check — service.containers.liveness_check.http_health_check / 189f03021cbb / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.containers](data-sources--workload--reference--group-015.md#canonical-f3e5d811f1c619d3bb1afde6648794fd06112ea88c6a7458c78327031b9486cf)
- [service.containers.liveness_check](data-sources--workload--reference--group-015.md#canonical-9359bc49c0401d6ea98bd188449578db7855267e305be7d84d0d7b8e1137b1d8)
- service.containers.liveness_check.http_health_check

<a id="canonical-023b5a48bf8c090d5dae14ba4f1c069490d4180d33a5ff2abf40c570d416e86e"></a>

Type: `"single"`. Computed.

HTTPHealthCheckType describes a health check based on HTTP GET requests.

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

<a id="canonical-4b84e9a21ca12f81ad906d2eff03e52d5323070d8330af82bdcf66c187bcddb2"></a>

## Direct properties — service.containers.liveness_check.http_health_check / 189f03021cbb / 3

<a id="canonical-e5245d2d2dff9ba76ba161e7684c27832f1839af335df9120fec0bcc12f2bbd2"></a>

<a id="canonical-7a31cf845eba812a1193959f23aa3c0d3625a3289f00865ec3878f306c81a70d"></a>

## headers property — service.containers.liveness_check.http_health_check / 189f03021cbb / 4

Type: `["map", "string"]`. Computed.

Specifies a list of HTTP headers that should be added to each request that is sent to the health
checked container. This is a list of key-value pairs.

Upstream description:

Specifies a list of HTTP headers that should be added to each request that is sent to the health
checked container. This is a list of key-value pairs.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "2048",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "2048",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-e9a5cc78c504ece891e74b1aa9245582f6a4e1f0f97dc4c7ea6d54d6d3a838ae"></a>

<a id="canonical-f964b23fc7822e18ab5819411afc788e1bd1c44b19ea3e12d242cb2802ecbc0c"></a>

## host_header property — service.containers.liveness_check.http_health_check / 189f03021cbb / 5

Type: `"string"`. Computed.

The value of the host header in the HTTP health check request.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 262,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 262,
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
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  }
}
```

<a id="canonical-24908e0c562fbab7809d248ae7d0d98c291ed7f9c32ba9717914982f0645131f"></a>

<a id="canonical-26bf61746876e6493dc98494f6b7d88ad3889a534538a2200fbc37e0cdc9e612"></a>

## path property — service.containers.liveness_check.http_health_check / 189f03021cbb / 6

Type: `"string"`. Computed.

Path. Path to access on the HTTP server.

Upstream description:

Path to access on the HTTP server.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

- [port](data-sources--workload--reference--group-015.md#canonical-bbfb53c7127bfd3768cf5b8cbb7868e52379dfbbca471f7e7ffde7a30a545189): complete subsection reference.

<a id="canonical-81717315a9bb8efc9e32fe138fe9199fbf5b2ea3fb1068405b650aafcf66e9a3"></a>

## Next pages — service.containers.liveness_check.http_health_check / 189f03021cbb / 7

- [service.containers.liveness_check.http_health_check.port](data-sources--workload--reference--group-015.md#canonical-bbfb53c7127bfd3768cf5b8cbb7868e52379dfbbca471f7e7ffde7a30a545189)
- [service.containers.liveness_check](data-sources--workload--reference--group-015.md#canonical-9359bc49c0401d6ea98bd188449578db7855267e305be7d84d0d7b8e1137b1d8)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-bbfb53c7127bfd3768cf5b8cbb7868e52379dfbbca471f7e7ffde7a30a545189"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a28aa5bf8d59683846e302632e724c26f3ee1b894b4c55b0c73b9b36bae573e8"></a>

## service.containers.liveness_check.http_health_check.port — service.containers.liveness_check.http_health_check.port / ef894ef0ce71 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.containers](data-sources--workload--reference--group-015.md#canonical-f3e5d811f1c619d3bb1afde6648794fd06112ea88c6a7458c78327031b9486cf)
- [service.containers.liveness_check](data-sources--workload--reference--group-015.md#canonical-9359bc49c0401d6ea98bd188449578db7855267e305be7d84d0d7b8e1137b1d8)
- [service.containers.liveness_check.http_health_check](data-sources--workload--reference--group-015.md#canonical-5bb7a7f1b5475f8f80e3df328c843ecc0b5893208f40e5ebdf3c37d9d827bde0)
- service.containers.liveness_check.http_health_check.port

<a id="canonical-4ad19f14d184cdaabac90244c3589614c48834b3275582e03610bd348221aee4"></a>

Type: `"single"`. Computed.

Port. Port

Upstream description:

Port

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_choice": "[\"name\",\"num\"]"
}
```

<a id="canonical-607525f114d4344a9031b02798e05c36c141edac25bccdf544c408a7b921d8be"></a>

## Direct properties — service.containers.liveness_check.http_health_check.port / ef894ef0ce71 / 3

<a id="canonical-7dd1ce892f718f28fad1b004f96a79ae9232cad1a1468beeecedd5d0bdcaa78c"></a>

<a id="canonical-2799c3eb1d67d8eb3ceea71fd4c8ab02a19a56b90ce52850b0d7cb832a8cbbed"></a>

## name property — service.containers.liveness_check.http_health_check.port / ef894ef0ce71 / 4

Type: `"string"`. Computed.

Port Name. Exclusive with \[num\] Port Name.

Upstream description:

Exclusive with \[num\] Port Name.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.iana_svc_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.iana_svc_name": "true"
  }
}
```

<a id="canonical-8e9131c331652a7c072983c439068dc738a94a49a9befaf0133333dc3d5abadf"></a>

<a id="canonical-54232a16f9f301fcef51fe1ca8d9dccd52d63eee08e344655798cf6e583a70d5"></a>

## num property — service.containers.liveness_check.http_health_check.port / ef894ef0ce71 / 5

Type: `"number"`. Computed.

Port Number. Exclusive with \[name\] Port number.

Upstream description:

Exclusive with \[name\] Port number.

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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-eaa83774c49902483d22ec2d39969e477144c1132dcb5ce4fc3a5450756ff869"></a>

## Next pages — service.containers.liveness_check.http_health_check.port / ef894ef0ce71 / 6

- [service.containers.liveness_check.http_health_check](data-sources--workload--reference--group-015.md#canonical-5bb7a7f1b5475f8f80e3df328c843ecc0b5893208f40e5ebdf3c37d9d827bde0)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-af2794efee684bc26c4e0d0e0a56bdfd087f12eaca7cc8db982a848a0188c811"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d523a27e44a540c427c11b1d6a80f334c874e49d2bae1a7eaec68d5bf3752d31"></a>

## service.containers.liveness_check.tcp_health_check — service.containers.liveness_check.tcp_health_check / 692bf0fff7c5 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.containers](data-sources--workload--reference--group-015.md#canonical-f3e5d811f1c619d3bb1afde6648794fd06112ea88c6a7458c78327031b9486cf)
- [service.containers.liveness_check](data-sources--workload--reference--group-015.md#canonical-9359bc49c0401d6ea98bd188449578db7855267e305be7d84d0d7b8e1137b1d8)
- service.containers.liveness_check.tcp_health_check

<a id="canonical-dc51b781e86c0baf0907ae5470a3bad255b6c2784e14732523602493d179abf3"></a>

Type: `"single"`. Computed.

TCPHealthCheckType describes a health check based on opening a TCP connection.

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

<a id="canonical-2ef9708818b6fac41477ef1e291057a09fe2a50e3212733246f9fc605f34c24b"></a>

## Direct properties — service.containers.liveness_check.tcp_health_check / 692bf0fff7c5 / 3

- [port](data-sources--workload--reference--group-015.md#canonical-a8dbf12307700631f4175ef3373f9c3e5b377f46699b0df8dc97451400c4dc4b): complete subsection reference.

<a id="canonical-29927bc7d56646cd2367ffb256c790945bfb82d3b64274bbfbae382781bf10c8"></a>

## Next pages — service.containers.liveness_check.tcp_health_check / 692bf0fff7c5 / 4

- [service.containers.liveness_check.tcp_health_check.port](data-sources--workload--reference--group-015.md#canonical-a8dbf12307700631f4175ef3373f9c3e5b377f46699b0df8dc97451400c4dc4b)
- [service.containers.liveness_check](data-sources--workload--reference--group-015.md#canonical-9359bc49c0401d6ea98bd188449578db7855267e305be7d84d0d7b8e1137b1d8)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-a8dbf12307700631f4175ef3373f9c3e5b377f46699b0df8dc97451400c4dc4b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3354a211a65cc66035add14b72797a85b1d934fc3c9465a00efeae598770378c"></a>

## service.containers.liveness_check.tcp_health_check.port — service.containers.liveness_check.tcp_health_check.port / 30644738299c / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.containers](data-sources--workload--reference--group-015.md#canonical-f3e5d811f1c619d3bb1afde6648794fd06112ea88c6a7458c78327031b9486cf)
- [service.containers.liveness_check](data-sources--workload--reference--group-015.md#canonical-9359bc49c0401d6ea98bd188449578db7855267e305be7d84d0d7b8e1137b1d8)
- [service.containers.liveness_check.tcp_health_check](data-sources--workload--reference--group-015.md#canonical-af2794efee684bc26c4e0d0e0a56bdfd087f12eaca7cc8db982a848a0188c811)
- service.containers.liveness_check.tcp_health_check.port

<a id="canonical-79c60a5644db992d1e5f06be28c3d8e3cddb92d693506537b787381409b45a80"></a>

Type: `"single"`. Computed.

Port. Port

Upstream description:

Port

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_choice": "[\"name\",\"num\"]"
}
```

<a id="canonical-04d4d09517276de8b4936907e08cd87190bb49aff480d976e0e0b5c1a81d00e9"></a>

## Direct properties — service.containers.liveness_check.tcp_health_check.port / 30644738299c / 3

<a id="canonical-e8b645879e46432786d3f5d3e6e42b99667c3df53980b59796e9fc741fe303d0"></a>

<a id="canonical-f767e19df3bceffbb6195662a9e0fddcbc91ad14b45f728ea168e82e3e429e19"></a>

## name property — service.containers.liveness_check.tcp_health_check.port / 30644738299c / 4

Type: `"string"`. Computed.

Port Name. Exclusive with \[num\] Port Name.

Upstream description:

Exclusive with \[num\] Port Name.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.iana_svc_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.iana_svc_name": "true"
  }
}
```

<a id="canonical-83b1086858fa02a2f10df4995fe038ae6e07041c652c031e05344143ca6bf1bc"></a>

<a id="canonical-38a43ab467e9494f39068787fe3a74c0388687efd5321d2cf41e2e458ab34596"></a>

## num property — service.containers.liveness_check.tcp_health_check.port / 30644738299c / 5

Type: `"number"`. Computed.

Port Number. Exclusive with \[name\] Port number.

Upstream description:

Exclusive with \[name\] Port number.

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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-7a0d687680cab817f1048b7639323939229707836e547eb38e6d874c9f9720c6"></a>

## Next pages — service.containers.liveness_check.tcp_health_check.port / 30644738299c / 6

- [service.containers.liveness_check.tcp_health_check](data-sources--workload--reference--group-015.md#canonical-af2794efee684bc26c4e0d0e0a56bdfd087f12eaca7cc8db982a848a0188c811)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-59a7c74fa8e3b4ee945cdef67d7014c8eb4b7fd15e9d23ac48180839bbbcd306"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-80fa952d7c63662a20aa2b52957c620f495e973f2d04cbb8fc5767338eaec856"></a>

## service.containers.readiness_check — service.containers.readiness_check / ad31914e3da8 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.containers](data-sources--workload--reference--group-015.md#canonical-f3e5d811f1c619d3bb1afde6648794fd06112ea88c6a7458c78327031b9486cf)
- service.containers.readiness_check

<a id="canonical-9bcb811ace5f02c5a0f0a0b4d70b80f9d08e3ecab575643a04c4175ab0695aeb"></a>

Type: `"single"`. Computed.

HealthCheckType describes a health check to be performed against a container to determine whether it
has started up or is alive or ready to receive traffic.

Upstream description:

HealthCheckType describes a health check to be performed against a container to determine whether it
has started up or is alive or ready to receive traffic.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-health_check_choice": "[\"exec_health_check\",\"http_health_check\",\"tcp_health_check\"]"
}
```

<a id="canonical-589da97c4aa4eed749a3558dad791345fd6a1533d4557ab5ef83c825967bee99"></a>

## Direct properties — service.containers.readiness_check / ad31914e3da8 / 3

- [exec_health_check](data-sources--workload--reference--group-015.md#canonical-e36ee4ad2342d3a657c2918e82346bb4ffdc577168b6ddabe985b61193ecc260): complete subsection reference.

<a id="canonical-e81eadbe01d5e6636fae5075442dcb71e890aa63bb638bd64098c86beac038fe"></a>

<a id="canonical-9fc7668b8bdaaae8a7d779f6029fb6538b6c2a9d49989389736cb2e1937c6f24"></a>

## healthy_threshold property — service.containers.readiness_check / ad31914e3da8 / 4

Type: `"number"`. Computed.

Number of consecutive successful responses after having failed before declaring healthy. In other
words, this is the number of healthy health checks required before marking healthy. Note that during
startup and liveliness, only a single successful health check is required to mark a container..

Upstream description:

Number of consecutive successful responses after having failed before declaring healthy. In other
words, this is the number of healthy health checks required before marking healthy. Note that during
startup and liveliness, only a single successful health check is required to mark a container
healthy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

- [http_health_check](data-sources--workload--reference--group-015.md#canonical-6befa4e193cb7b5813d380a3bfab6163480f62900154096555d9558cdc924a9c): complete subsection reference.

<a id="canonical-95a84f0e9999bc7fa36ef61f846af94dd8ab5028c43746cd21d3bfec58b0ff9a"></a>

<a id="canonical-f5ca804968c9542f652de9020ef0c1c59aa706ab5423c5b9f295798a67426a95"></a>

## initial_delay property — service.containers.readiness_check / ad31914e3da8 / 5

Type: `"number"`. Computed.

Number of seconds after the container has started before health checks are initiated.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
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
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

<a id="canonical-3f0ded464bc54ce29806c9eafd4b8166a40c225c22280cfe94c8f70c163cd470"></a>

<a id="canonical-7d905fb9869fca6762c4839c3a074de926d8a726055fa9251f088fed4291661e"></a>

## interval property — service.containers.readiness_check / ad31914e3da8 / 6

Type: `"number"`. Computed.

Time interval in seconds between two health check requests.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

- [tcp_health_check](data-sources--workload--reference--group-015.md#canonical-66544e90429c84a69ccf75312fc51fa9b9ccf6f1c63b409e83b9b20330a42d9c): complete subsection reference.

<a id="canonical-35657a6ae5f07a3e6628f332971d044bdafe17cecd80de9b045f6d8321d765cd"></a>

<a id="canonical-e9e7a428224591e667671b55bb9e0ea7dc838eeeac899611dd02a0bce3680ed5"></a>

## timeout property — service.containers.readiness_check / ad31914e3da8 / 7

Type: `"number"`. Computed.

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure.

Upstream description:

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

<a id="canonical-7291e826797d28536ec5d462cc85c98273a2b3ebda546540280e9120c371cb71"></a>

<a id="canonical-601dfed40530aa829bba3508f8d6678cc285ec5a5a607484c5d3a82f237bd8ed"></a>

## unhealthy_threshold property — service.containers.readiness_check / ad31914e3da8 / 8

Type: `"number"`. Computed.

Number of consecutive failed responses before declaring unhealthy. In other words, this is the
number of unhealthy health checks required before a container is marked unhealthy.

Upstream description:

Number of consecutive failed responses before declaring unhealthy. In other words, this is the
number of unhealthy health checks required before a container is marked unhealthy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

<a id="canonical-6fdeae0752214349108f5cc6739c680b22081d76a0fa499cb68a11461f916485"></a>

## Next pages — service.containers.readiness_check / ad31914e3da8 / 9

- [service.containers.readiness_check.exec_health_check](data-sources--workload--reference--group-015.md#canonical-e36ee4ad2342d3a657c2918e82346bb4ffdc577168b6ddabe985b61193ecc260)
- [service.containers.readiness_check.http_health_check](data-sources--workload--reference--group-015.md#canonical-6befa4e193cb7b5813d380a3bfab6163480f62900154096555d9558cdc924a9c)
- [service.containers.readiness_check.tcp_health_check](data-sources--workload--reference--group-015.md#canonical-66544e90429c84a69ccf75312fc51fa9b9ccf6f1c63b409e83b9b20330a42d9c)
- [service.containers](data-sources--workload--reference--group-015.md#canonical-f3e5d811f1c619d3bb1afde6648794fd06112ea88c6a7458c78327031b9486cf)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-e36ee4ad2342d3a657c2918e82346bb4ffdc577168b6ddabe985b61193ecc260"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-18ca1d144edba0fda1814cf1e3e6d56ae50c7252ee58b205e68af9dc924655ee"></a>

## service.containers.readiness_check.exec_health_check — service.containers.readiness_check.exec_health_check / a3cc2fa504b4 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.containers](data-sources--workload--reference--group-015.md#canonical-f3e5d811f1c619d3bb1afde6648794fd06112ea88c6a7458c78327031b9486cf)
- [service.containers.readiness_check](data-sources--workload--reference--group-015.md#canonical-59a7c74fa8e3b4ee945cdef67d7014c8eb4b7fd15e9d23ac48180839bbbcd306)
- service.containers.readiness_check.exec_health_check

<a id="canonical-c9d63db7b061ab02dc1ae2d507104ff7604d6d1c7eb5a6864f825d117aab0327"></a>

Type: `"single"`. Computed.

ExecHealthCheckType describes a health check based on 'run in container' action. Exit status of 0 is
treated as live/healthy and non-zero is unhealthy.

Upstream description:

ExecHealthCheckType describes a health check based on "run in container" action. Exit status of 0 is
treated as live/healthy and non-zero is unhealthy.

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

<a id="canonical-e4d1cf917d205a1c85bd3552bad2a26186acfb028c2c3e9c716991c9a0d87fdf"></a>

## Direct properties — service.containers.readiness_check.exec_health_check / a3cc2fa504b4 / 3

<a id="canonical-c89449f9743cc9cacd6427b000a0e629ec412a75a65a412ea0e73268d2a30400"></a>

<a id="canonical-a4a960364a7e9755d6da92ac93e2d909bd424f5a9163f27f2f830870da56a1ea"></a>

## command property — service.containers.readiness_check.exec_health_check / a3cc2fa504b4 / 4

Type: `["list", "string"]`. Computed.

Command is the command line to execute inside the container, the working directory for the command
is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a
shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to..

Upstream description:

Command is the command line to execute inside the container, the working directory for the command
is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a
shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to
explicitly call out to that shell.

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
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-51330cdc331605e65a87db8d7aeb193f17019c90654026150ee1de1934ff5f37"></a>

## Next pages — service.containers.readiness_check.exec_health_check / a3cc2fa504b4 / 5

- [service.containers.readiness_check](data-sources--workload--reference--group-015.md#canonical-59a7c74fa8e3b4ee945cdef67d7014c8eb4b7fd15e9d23ac48180839bbbcd306)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-6befa4e193cb7b5813d380a3bfab6163480f62900154096555d9558cdc924a9c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7021fca4a103003f77779098ab78003904a9436ae0ad6dd374e4d73c9faa102f"></a>

## service.containers.readiness_check.http_health_check — service.containers.readiness_check.http_health_check / f00cb0cdef76 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.containers](data-sources--workload--reference--group-015.md#canonical-f3e5d811f1c619d3bb1afde6648794fd06112ea88c6a7458c78327031b9486cf)
- [service.containers.readiness_check](data-sources--workload--reference--group-015.md#canonical-59a7c74fa8e3b4ee945cdef67d7014c8eb4b7fd15e9d23ac48180839bbbcd306)
- service.containers.readiness_check.http_health_check

<a id="canonical-90cfca05f9d706e88022411c88a695fe0cf93cd6353969a3b699e59dc6c92f3f"></a>

Type: `"single"`. Computed.

HTTPHealthCheckType describes a health check based on HTTP GET requests.

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

<a id="canonical-94c4cb526b36f86384ad0177f5194f8793c2ebe7a23fe81fed164b23d1d03c4b"></a>

## Direct properties — service.containers.readiness_check.http_health_check / f00cb0cdef76 / 3

<a id="canonical-2b41f10f3e4fcb94e631d79f8b5eed582084c32d1f79b7d1442772c853392f6b"></a>

<a id="canonical-8f6d2675dc19d74711221cb632fd64d96f6d4ba4d9aba411ec816906acf53e69"></a>

## headers property — service.containers.readiness_check.http_health_check / f00cb0cdef76 / 4

Type: `["map", "string"]`. Computed.

Specifies a list of HTTP headers that should be added to each request that is sent to the health
checked container. This is a list of key-value pairs.

Upstream description:

Specifies a list of HTTP headers that should be added to each request that is sent to the health
checked container. This is a list of key-value pairs.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "2048",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "2048",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-0bfb1abb0ffb3548a6c3de693ed178770ceacab786faa78b932d745511be960f"></a>

<a id="canonical-9855fe414d86423925ac265a880019795d86e103d0f493657213b6cf7314f772"></a>

## host_header property — service.containers.readiness_check.http_health_check / f00cb0cdef76 / 5

Type: `"string"`. Computed.

The value of the host header in the HTTP health check request.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 262,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 262,
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
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  }
}
```

<a id="canonical-3693818234fd4cee5bf158fa3e5b9c6571b386293fca8a2158f58a44590859ea"></a>

<a id="canonical-15bf6d1cfec2d63476f4f51035857f7d20b41fb43620b1005610df27f831b754"></a>

## path property — service.containers.readiness_check.http_health_check / f00cb0cdef76 / 6

Type: `"string"`. Computed.

Path. Path to access on the HTTP server.

Upstream description:

Path to access on the HTTP server.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

- [port](data-sources--workload--reference--group-015.md#canonical-bdc72ba4eb1ab50d41e956cbecce03398fa35aedb956272f83cf34396c3d879b): complete subsection reference.

<a id="canonical-788c7a14995cf16470911ba397ac1c7dab9cfd957b5b44f2fdd9a423ae7452be"></a>

## Next pages — service.containers.readiness_check.http_health_check / f00cb0cdef76 / 7

- [service.containers.readiness_check.http_health_check.port](data-sources--workload--reference--group-015.md#canonical-bdc72ba4eb1ab50d41e956cbecce03398fa35aedb956272f83cf34396c3d879b)
- [service.containers.readiness_check](data-sources--workload--reference--group-015.md#canonical-59a7c74fa8e3b4ee945cdef67d7014c8eb4b7fd15e9d23ac48180839bbbcd306)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-bdc72ba4eb1ab50d41e956cbecce03398fa35aedb956272f83cf34396c3d879b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7718153d5c2b5f97dc26571f48f052c7d220c7d93a9b274de34e6d997deef45f"></a>

## service.containers.readiness_check.http_health_check.port — service.containers.readiness_check.http_health_check.port / ab7fcf47819e / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.containers](data-sources--workload--reference--group-015.md#canonical-f3e5d811f1c619d3bb1afde6648794fd06112ea88c6a7458c78327031b9486cf)
- [service.containers.readiness_check](data-sources--workload--reference--group-015.md#canonical-59a7c74fa8e3b4ee945cdef67d7014c8eb4b7fd15e9d23ac48180839bbbcd306)
- [service.containers.readiness_check.http_health_check](data-sources--workload--reference--group-015.md#canonical-6befa4e193cb7b5813d380a3bfab6163480f62900154096555d9558cdc924a9c)
- service.containers.readiness_check.http_health_check.port

<a id="canonical-54e88d7b92c94ec74fa62b336929106f902772dbbc33f4938c4c1e059ee6242c"></a>

Type: `"single"`. Computed.

Port. Port

Upstream description:

Port

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_choice": "[\"name\",\"num\"]"
}
```

<a id="canonical-9fa314e309c0c9eea103183ed534b5723fa296a86d3e6e55ad3c673e2f471314"></a>

## Direct properties — service.containers.readiness_check.http_health_check.port / ab7fcf47819e / 3

<a id="canonical-a7ed771af8c5af2b73eb581a86e84005e70df61d40ef0d08129a2badbcb2b1e7"></a>

<a id="canonical-42a2ed52091d08d2b7be56ddef9671e7baa418ff249991b9b01996cf86a6573f"></a>

## name property — service.containers.readiness_check.http_health_check.port / ab7fcf47819e / 4

Type: `"string"`. Computed.

Port Name. Exclusive with \[num\] Port Name.

Upstream description:

Exclusive with \[num\] Port Name.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.iana_svc_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.iana_svc_name": "true"
  }
}
```

<a id="canonical-ed1b2830a6ca7dc4860cdca3eeb7b36940438c1e0bed51b0d7624d719da86fcd"></a>

<a id="canonical-9508e8216b47e0e4fe9cdc9ff928d3f533b0c7c84cdeced5647cfbe543cb3734"></a>

## num property — service.containers.readiness_check.http_health_check.port / ab7fcf47819e / 5

Type: `"number"`. Computed.

Port Number. Exclusive with \[name\] Port number.

Upstream description:

Exclusive with \[name\] Port number.

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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-5bc69e9580e5740c50a71ed4193c091680d82dcef912b96ace6ca0b5cde87f74"></a>

## Next pages — service.containers.readiness_check.http_health_check.port / ab7fcf47819e / 6

- [service.containers.readiness_check.http_health_check](data-sources--workload--reference--group-015.md#canonical-6befa4e193cb7b5813d380a3bfab6163480f62900154096555d9558cdc924a9c)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-66544e90429c84a69ccf75312fc51fa9b9ccf6f1c63b409e83b9b20330a42d9c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-14b731db1502d87f5b7ff9cb51e017e00c090682c4710d13e943a6c83f563037"></a>

## service.containers.readiness_check.tcp_health_check — service.containers.readiness_check.tcp_health_check / a820cdad3b6c / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.containers](data-sources--workload--reference--group-015.md#canonical-f3e5d811f1c619d3bb1afde6648794fd06112ea88c6a7458c78327031b9486cf)
- [service.containers.readiness_check](data-sources--workload--reference--group-015.md#canonical-59a7c74fa8e3b4ee945cdef67d7014c8eb4b7fd15e9d23ac48180839bbbcd306)
- service.containers.readiness_check.tcp_health_check

<a id="canonical-2686257cf9b487c529a3caf32a066c441170d3b029530062a8c2a415fff5f4e9"></a>

Type: `"single"`. Computed.

TCPHealthCheckType describes a health check based on opening a TCP connection.

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

<a id="canonical-0c72ce26dded03c86b3ba6458f69d58c9863780f5d5fe7920e4b1df3124e016e"></a>

## Direct properties — service.containers.readiness_check.tcp_health_check / a820cdad3b6c / 3

- [port](data-sources--workload--reference--group-015.md#canonical-c2a8f68052b9f6e5f387d6b7c2fd46ca111ce6283a29f077482fb56cead5e1d9): complete subsection reference.

<a id="canonical-de128c1855106408c6727a3b1060064a6a3ef134ee90f4b17c7669c3129c0a85"></a>

## Next pages — service.containers.readiness_check.tcp_health_check / a820cdad3b6c / 4

- [service.containers.readiness_check.tcp_health_check.port](data-sources--workload--reference--group-015.md#canonical-c2a8f68052b9f6e5f387d6b7c2fd46ca111ce6283a29f077482fb56cead5e1d9)
- [service.containers.readiness_check](data-sources--workload--reference--group-015.md#canonical-59a7c74fa8e3b4ee945cdef67d7014c8eb4b7fd15e9d23ac48180839bbbcd306)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-c2a8f68052b9f6e5f387d6b7c2fd46ca111ce6283a29f077482fb56cead5e1d9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-70ee0ba5eee39a43bd7777f78e79103ffbd9852e8170885de1c5e84dfec65cde"></a>

## service.containers.readiness_check.tcp_health_check.port — service.containers.readiness_check.tcp_health_check.port / f34d734680fc / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.containers](data-sources--workload--reference--group-015.md#canonical-f3e5d811f1c619d3bb1afde6648794fd06112ea88c6a7458c78327031b9486cf)
- [service.containers.readiness_check](data-sources--workload--reference--group-015.md#canonical-59a7c74fa8e3b4ee945cdef67d7014c8eb4b7fd15e9d23ac48180839bbbcd306)
- [service.containers.readiness_check.tcp_health_check](data-sources--workload--reference--group-015.md#canonical-66544e90429c84a69ccf75312fc51fa9b9ccf6f1c63b409e83b9b20330a42d9c)
- service.containers.readiness_check.tcp_health_check.port

<a id="canonical-c6c371334b823d03458663cb3c660f759126f828f6c5925ae0ed3e9e8cab735a"></a>

Type: `"single"`. Computed.

Port. Port

Upstream description:

Port

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_choice": "[\"name\",\"num\"]"
}
```

<a id="canonical-2b793ec68a4bccb7042a0ce05777e59f77acf2354c1b58c3f456461a26126f1e"></a>

## Direct properties — service.containers.readiness_check.tcp_health_check.port / f34d734680fc / 3

<a id="canonical-1967a283d31c744f66ffa6a09d046e56c3205127ad486a1d73f043eeb729407d"></a>

<a id="canonical-978f1492b6601d0b54f04ac4d208e6cbbfc99c8a753f125c0b2d9610ccfdadcc"></a>

## name property — service.containers.readiness_check.tcp_health_check.port / f34d734680fc / 4

Type: `"string"`. Computed.

Port Name. Exclusive with \[num\] Port Name.

Upstream description:

Exclusive with \[num\] Port Name.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.iana_svc_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.iana_svc_name": "true"
  }
}
```

<a id="canonical-26938b9da38e741fda3e44fb007e787683ce5888611e591a2e3237c3ec2749d3"></a>

<a id="canonical-eec66b86ed5154a0b2a42fcf99d65c23e8412f9acee190afb32314d8d020d36e"></a>

## num property — service.containers.readiness_check.tcp_health_check.port / f34d734680fc / 5

Type: `"number"`. Computed.

Port Number. Exclusive with \[name\] Port number.

Upstream description:

Exclusive with \[name\] Port number.

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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-c427e1d32cf3395526bac1c0fd6e91d98f4f2195574938e02c152a32b807c8a0"></a>

## Next pages — service.containers.readiness_check.tcp_health_check.port / f34d734680fc / 6

- [service.containers.readiness_check.tcp_health_check](data-sources--workload--reference--group-015.md#canonical-66544e90429c84a69ccf75312fc51fa9b9ccf6f1c63b409e83b9b20330a42d9c)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-1938114db94be122e9e6f32ed52bcf7ce88fba44423daeb9b9b3cb3545bbc9b8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-30a4b356faea6fdaaf92ab30c922db9709278ad2a96a895c55ee4a159c90a49f"></a>

## service.deploy_options — service.deploy_options / 487dbdd633a7 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- service.deploy_options

<a id="canonical-2dc16d57454b0eef6a7b7f74569fa39dd885d382cb2a9a0db3201e1a007ed4fb"></a>

Type: `"single"`. Computed.

Deploy OPTIONS are used to configure the workload deployment OPTIONS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-deploy_choice": "[\"all_res\",\"default_virtual_sites\",\"deploy_ce_sites\",\"deploy_ce_virtual_sites\",\"deploy_re_sites\",\"deploy_re_virtual_sites\"]"
}
```

<a id="canonical-c973dcaad34c0a293c49ed007cb0932c332f68ce2d4460b8e190df8f21fc896f"></a>

## Direct properties — service.deploy_options / 487dbdd633a7 / 3

- [all_res](data-sources--workload--reference--group-015.md#canonical-2203a0f123a6b3b9ea934ff5299daadfbd8c8059ace2823399c8eca623e4601e): complete subsection reference.

- [default_virtual_sites](data-sources--workload--reference--group-015.md#canonical-16febc382a742a7127c55a35de9fa9eadce0359070a0723c54aa620a6a6fad1e): complete subsection reference.

- [deploy_ce_sites](data-sources--workload--reference--group-015.md#canonical-68fed3b3476f5ab1ded0fadb89e7ac5ef9100971649b10ac6dafef7958fed7e3): complete subsection reference.

- [deploy_ce_virtual_sites](data-sources--workload--reference--group-015.md#canonical-cef9ccb04cabafe0e69397e4e9dd11108b7fc7b1665d41a2a8e9df90c333c703): complete subsection reference.

- [deploy_re_sites](data-sources--workload--reference--group-015.md#canonical-73b8a1786044b7eb659373d38c873cfa2d98caa0d71ed74406f1a394a16f97ae): complete subsection reference.

- [deploy_re_virtual_sites](data-sources--workload--reference--group-015.md#canonical-4efe9b660641250eaef1573c02be5b28c5b6f1c1bdf439c1e9acf505e56b3874): complete subsection reference.

<a id="canonical-ba70f515a14d7833356e055dcefe0ad48797081cc45808a2af84c11f94d8d8c1"></a>

## Next pages — service.deploy_options / 487dbdd633a7 / 4

- [service.deploy_options.all_res](data-sources--workload--reference--group-015.md#canonical-2203a0f123a6b3b9ea934ff5299daadfbd8c8059ace2823399c8eca623e4601e)
- [service.deploy_options.default_virtual_sites](data-sources--workload--reference--group-015.md#canonical-16febc382a742a7127c55a35de9fa9eadce0359070a0723c54aa620a6a6fad1e)
- [service.deploy_options.deploy_ce_sites](data-sources--workload--reference--group-015.md#canonical-68fed3b3476f5ab1ded0fadb89e7ac5ef9100971649b10ac6dafef7958fed7e3)
- [service.deploy_options.deploy_ce_virtual_sites](data-sources--workload--reference--group-015.md#canonical-cef9ccb04cabafe0e69397e4e9dd11108b7fc7b1665d41a2a8e9df90c333c703)
- [service.deploy_options.deploy_re_sites](data-sources--workload--reference--group-015.md#canonical-73b8a1786044b7eb659373d38c873cfa2d98caa0d71ed74406f1a394a16f97ae)
- [service.deploy_options.deploy_re_virtual_sites](data-sources--workload--reference--group-015.md#canonical-4efe9b660641250eaef1573c02be5b28c5b6f1c1bdf439c1e9acf505e56b3874)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-2203a0f123a6b3b9ea934ff5299daadfbd8c8059ace2823399c8eca623e4601e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dfdf81ccdf66820a1a2a178477372a9c1c6eb976563d21503b2f97c200e318c0"></a>

## service.deploy_options.all_res — service.deploy_options.all_res / 3181f2f880f6 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.deploy_options](data-sources--workload--reference--group-015.md#canonical-1938114db94be122e9e6f32ed52bcf7ce88fba44423daeb9b9b3cb3545bbc9b8)
- service.deploy_options.all_res

<a id="canonical-a1944ba484e6e79692752515c0a263a62e4121e7595aaf99224b560f0e80d27b"></a>

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

<a id="canonical-260c5b8d37fe27e1de3f15df1460067bf4eb6fd64188619a5d9497ce5f9ddd75"></a>

## Direct properties — service.deploy_options.all_res / 3181f2f880f6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2f2d1efac75387c88c4a78246779a87dcb2bf941b01cd6398117a5b56d3e2159"></a>

## Next pages — service.deploy_options.all_res / 3181f2f880f6 / 4

- [service.deploy_options](data-sources--workload--reference--group-015.md#canonical-1938114db94be122e9e6f32ed52bcf7ce88fba44423daeb9b9b3cb3545bbc9b8)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-16febc382a742a7127c55a35de9fa9eadce0359070a0723c54aa620a6a6fad1e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d705968eedb86f2b7d8b02648902837451a93d2b1fb0696757b71af134715a3d"></a>

## service.deploy_options.default_virtual_sites — service.deploy_options.default_virtual_sites / 230f23b1a47e / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.deploy_options](data-sources--workload--reference--group-015.md#canonical-1938114db94be122e9e6f32ed52bcf7ce88fba44423daeb9b9b3cb3545bbc9b8)
- service.deploy_options.default_virtual_sites

<a id="canonical-e8cbcd75dc310a64f21e0e0b5574c59bd9cbc3356cd33ad12c4b2ec0cd4d0577"></a>

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

<a id="canonical-d846bad15ced90cfa2b7ce02e572c322932f7c2584653d3b78f99201e7d24d9e"></a>

## Direct properties — service.deploy_options.default_virtual_sites / 230f23b1a47e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7e11b84febd4d7705c8aefb30e861e77fb3d643e1ef5f2354457cdb6099e292f"></a>

## Next pages — service.deploy_options.default_virtual_sites / 230f23b1a47e / 4

- [service.deploy_options](data-sources--workload--reference--group-015.md#canonical-1938114db94be122e9e6f32ed52bcf7ce88fba44423daeb9b9b3cb3545bbc9b8)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-68fed3b3476f5ab1ded0fadb89e7ac5ef9100971649b10ac6dafef7958fed7e3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-399904df9f8b70fda4ba31af2f6f67d137b5fe32719269f3def6fabb24fe81d4"></a>

## service.deploy_options.deploy_ce_sites — service.deploy_options.deploy_ce_sites / 7ddc08f39fe6 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.deploy_options](data-sources--workload--reference--group-015.md#canonical-1938114db94be122e9e6f32ed52bcf7ce88fba44423daeb9b9b3cb3545bbc9b8)
- service.deploy_options.deploy_ce_sites

<a id="canonical-b0bd4fdfea63ea60b2e39421195be9a42b85081a952a338e94abd3090523cc2a"></a>

Type: `"single"`. Computed.

Defines a way to deploy a workload on specific Customer sites.

Upstream description:

This defines a way to deploy a workload on specific Customer sites.

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

<a id="canonical-9a8845a1597d0fde2417e22e11749a682518664fe194a349008b58793d50e260"></a>

## Direct properties — service.deploy_options.deploy_ce_sites / 7ddc08f39fe6 / 3

- [site](data-sources--workload--reference--group-015.md#canonical-713597d99ba586ad9152cb7ee7064123d01a1ffa62f817f63a7e5d5a8db2efd3): complete subsection reference.

<a id="canonical-1b843d870cdc0e4fe343155859214a8c46920cafad5f9228cc430d5510ae0725"></a>

## Next pages — service.deploy_options.deploy_ce_sites / 7ddc08f39fe6 / 4

- [service.deploy_options.deploy_ce_sites.site](data-sources--workload--reference--group-015.md#canonical-713597d99ba586ad9152cb7ee7064123d01a1ffa62f817f63a7e5d5a8db2efd3)
- [service.deploy_options](data-sources--workload--reference--group-015.md#canonical-1938114db94be122e9e6f32ed52bcf7ce88fba44423daeb9b9b3cb3545bbc9b8)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-713597d99ba586ad9152cb7ee7064123d01a1ffa62f817f63a7e5d5a8db2efd3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-00913e73a90e02154b6339be58106284ca563e20978558baa10f628f4bdaaca1"></a>

## service.deploy_options.deploy_ce_sites.site — service.deploy_options.deploy_ce_sites.site / 461f245ef79a / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.deploy_options](data-sources--workload--reference--group-015.md#canonical-1938114db94be122e9e6f32ed52bcf7ce88fba44423daeb9b9b3cb3545bbc9b8)
- [service.deploy_options.deploy_ce_sites](data-sources--workload--reference--group-015.md#canonical-68fed3b3476f5ab1ded0fadb89e7ac5ef9100971649b10ac6dafef7958fed7e3)
- service.deploy_options.deploy_ce_sites.site

<a id="canonical-c5180d556fce4ba60e13ddbc5848f05e7e14c50c0293d4bef5ede78890adb1d4"></a>

Type: `"list"`. Computed.

Which customer sites should this workload be deployed.

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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-8d60d2e79c068a172ba1ed064f78bcc4655d7501633b21bd6c547831f9f52700"></a>

## Direct properties — service.deploy_options.deploy_ce_sites.site / 461f245ef79a / 3

<a id="canonical-f200df21a64b022d47fe0d05ef84315a6b52cd61e878f49e9ed393ce98854590"></a>

<a id="canonical-01e41298443d6111c5d4595f08aedc1e0502dc6565fec7f216c6259192f0ae23"></a>

## name property — service.deploy_options.deploy_ce_sites.site / 461f245ef79a / 4

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

<a id="canonical-1f2280e95e6913413c66da9d6771eabc8dcff7392ccb0f62b7e2c09512f93ec8"></a>

<a id="canonical-105a912cb51e8ebea46eb5b133e301383d7157f6b7014d26d5216111042a0c83"></a>

## namespace property — service.deploy_options.deploy_ce_sites.site / 461f245ef79a / 5

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

<a id="canonical-03572a2ce8f50bdeb16149f2f204c8dc7718a7f339f2118dc5b6c31635a0b19d"></a>

<a id="canonical-89e0b210d088caba029bb452e6964fea5f3be5135dc67fcf061c3f016de49c6f"></a>

## tenant property — service.deploy_options.deploy_ce_sites.site / 461f245ef79a / 6

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

<a id="canonical-29afd163f4978ed131fe57245960489bb4b1393027827b62164db5c627efbd82"></a>

## Next pages — service.deploy_options.deploy_ce_sites.site / 461f245ef79a / 7

- [service.deploy_options.deploy_ce_sites](data-sources--workload--reference--group-015.md#canonical-68fed3b3476f5ab1ded0fadb89e7ac5ef9100971649b10ac6dafef7958fed7e3)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-cef9ccb04cabafe0e69397e4e9dd11108b7fc7b1665d41a2a8e9df90c333c703"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-814392b3602e796aaca722837e29083ec346887bef068a419010b3843d2a9434"></a>

## service.deploy_options.deploy_ce_virtual_sites — service.deploy_options.deploy_ce_virtual_sites / 6516d11e06df / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.deploy_options](data-sources--workload--reference--group-015.md#canonical-1938114db94be122e9e6f32ed52bcf7ce88fba44423daeb9b9b3cb3545bbc9b8)
- service.deploy_options.deploy_ce_virtual_sites

<a id="canonical-7d04d9a56a72141cee3361d25c1e0bb30e07a3195bef4e9456decd985dbe3cfc"></a>

Type: `"single"`. Computed.

Defines a way to deploy a workload on specific Customer virtual sites.

Upstream description:

This defines a way to deploy a workload on specific Customer virtual sites.

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

<a id="canonical-d5aae2ac12fb17b878947ab7981161e09462a00a8f36d64323d9016500e6fc2a"></a>

## Direct properties — service.deploy_options.deploy_ce_virtual_sites / 6516d11e06df / 3

- [virtual_site](data-sources--workload--reference--group-015.md#canonical-8208b0077d5f5a3506f1d9e1b3540fbb289d4a976af532cafc1ba160261273b0): complete subsection reference.

<a id="canonical-c174449ebfe956499694c24e979f2eff5a393fe3b16d18cb152c72c4ef5574b3"></a>

## Next pages — service.deploy_options.deploy_ce_virtual_sites / 6516d11e06df / 4

- [service.deploy_options.deploy_ce_virtual_sites.virtual_site](data-sources--workload--reference--group-015.md#canonical-8208b0077d5f5a3506f1d9e1b3540fbb289d4a976af532cafc1ba160261273b0)
- [service.deploy_options](data-sources--workload--reference--group-015.md#canonical-1938114db94be122e9e6f32ed52bcf7ce88fba44423daeb9b9b3cb3545bbc9b8)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-8208b0077d5f5a3506f1d9e1b3540fbb289d4a976af532cafc1ba160261273b0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-159733b04a27cdbfb71034905fb361a7a13490c856c35af791cfe8c3ed6012dd"></a>

## service.deploy_options.deploy_ce_virtual_sites.virtual_site — service.deploy_options.deploy_ce_virtual_sites.virtual_site / 810af92d19e0 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.deploy_options](data-sources--workload--reference--group-015.md#canonical-1938114db94be122e9e6f32ed52bcf7ce88fba44423daeb9b9b3cb3545bbc9b8)
- [service.deploy_options.deploy_ce_virtual_sites](data-sources--workload--reference--group-015.md#canonical-cef9ccb04cabafe0e69397e4e9dd11108b7fc7b1665d41a2a8e9df90c333c703)
- service.deploy_options.deploy_ce_virtual_sites.virtual_site

<a id="canonical-d3a050349b4b847a311bb7922655f16170bff647a21e879c1a11e5882cba421a"></a>

Type: `"list"`. Computed.

Which customer virtual sites should this workload be deployed.

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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-63e477bd2edbbc2d3d8fe4f3d0ff0b09599535b7aa82ed5c170e4107c7dde562"></a>

## Direct properties — service.deploy_options.deploy_ce_virtual_sites.virtual_site / 810af92d19e0 / 3

<a id="canonical-f1b5b4b79a96bed6f3811d2427875504cb598eb37ae6643bc580a2349df14065"></a>

<a id="canonical-e58d903284f2c06dd4e74494938dae4068a680890f0c6953b39ff87b947cee91"></a>

## name property — service.deploy_options.deploy_ce_virtual_sites.virtual_site / 810af92d19e0 / 4

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

<a id="canonical-f24d7d42b02aacae1287c4c8c6fd053d67501cd5916c61bd01fd4121ece243ad"></a>

<a id="canonical-9f1224b33716857ddad9f047cb2b648c48beb3e78086e7ed9ae032f886ad2147"></a>

## namespace property — service.deploy_options.deploy_ce_virtual_sites.virtual_site / 810af92d19e0 / 5

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

<a id="canonical-37fe9d200963ec16ec5d6133a5b56220b062948e3ebf8450c5cb834331f415f0"></a>

<a id="canonical-649e438ffb21d49962a4672eed8984a638c8ba8be11f06ae9066a2da8abfc399"></a>

## tenant property — service.deploy_options.deploy_ce_virtual_sites.virtual_site / 810af92d19e0 / 6

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

<a id="canonical-6ccbacae598b769ad5f41891ed2357caf90dfde021b45b40c9faa20b845da812"></a>

## Next pages — service.deploy_options.deploy_ce_virtual_sites.virtual_site / 810af92d19e0 / 7

- [service.deploy_options.deploy_ce_virtual_sites](data-sources--workload--reference--group-015.md#canonical-cef9ccb04cabafe0e69397e4e9dd11108b7fc7b1665d41a2a8e9df90c333c703)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-73b8a1786044b7eb659373d38c873cfa2d98caa0d71ed74406f1a394a16f97ae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fb31709967387c57ffc42a7aefd14b1537dbe00e2f9d6a02e416d386ae232d9f"></a>

## service.deploy_options.deploy_re_sites — service.deploy_options.deploy_re_sites / a72f338450bd / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.deploy_options](data-sources--workload--reference--group-015.md#canonical-1938114db94be122e9e6f32ed52bcf7ce88fba44423daeb9b9b3cb3545bbc9b8)
- service.deploy_options.deploy_re_sites

<a id="canonical-a3628a231bca29fad1f73fc0a02e22588a264bb1859543b60b3961c299a1399b"></a>

Type: `"single"`. Computed.

Defines a way to deploy a workload on specific Regional Edge sites.

Upstream description:

This defines a way to deploy a workload on specific Regional Edge sites.

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

<a id="canonical-01bc98788b4f431899067988ddee61b11b08f52da9a959db9c7df0e49536d949"></a>

## Direct properties — service.deploy_options.deploy_re_sites / a72f338450bd / 3

- [site](data-sources--workload--reference--group-015.md#canonical-5586fe78174464eb1d8992935fda4cb887886bff9f33f11e5e4c30e7d5855a77): complete subsection reference.

<a id="canonical-bdf79098ddfae7076f282b476eacfeeab48051f8616672be310d6f3563ad8447"></a>

## Next pages — service.deploy_options.deploy_re_sites / a72f338450bd / 4

- [service.deploy_options.deploy_re_sites.site](data-sources--workload--reference--group-015.md#canonical-5586fe78174464eb1d8992935fda4cb887886bff9f33f11e5e4c30e7d5855a77)
- [service.deploy_options](data-sources--workload--reference--group-015.md#canonical-1938114db94be122e9e6f32ed52bcf7ce88fba44423daeb9b9b3cb3545bbc9b8)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-5586fe78174464eb1d8992935fda4cb887886bff9f33f11e5e4c30e7d5855a77"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-69fcbd81ecb9a38b25ec653850c0255023ed1ee21ec45ac600bb77adeb2fe3a5"></a>

## service.deploy_options.deploy_re_sites.site — service.deploy_options.deploy_re_sites.site / 32f9cfbb03a7 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.deploy_options](data-sources--workload--reference--group-015.md#canonical-1938114db94be122e9e6f32ed52bcf7ce88fba44423daeb9b9b3cb3545bbc9b8)
- [service.deploy_options.deploy_re_sites](data-sources--workload--reference--group-015.md#canonical-73b8a1786044b7eb659373d38c873cfa2d98caa0d71ed74406f1a394a16f97ae)
- service.deploy_options.deploy_re_sites.site

<a id="canonical-e4555233eb895fe2dd72cf3a17b33b735615efbb9c7476c21f503bdc1aec55ed"></a>

Type: `"list"`. Computed.

Which regional edge sites should this workload be deployed.

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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-126818e51df87c0b27ed38b88f871cdfc73674b7c17a3f468e6bb93964ad533a"></a>

## Direct properties — service.deploy_options.deploy_re_sites.site / 32f9cfbb03a7 / 3

<a id="canonical-0edd9421c1d3e20bc80ca3542fe67408437c876b6e22cd7613afb99210154c6a"></a>

<a id="canonical-260b87702969d42be925a4c311f3c7bec199b92401f2c379c6e7abfd9e582583"></a>

## name property — service.deploy_options.deploy_re_sites.site / 32f9cfbb03a7 / 4

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

<a id="canonical-3a457cba353269c96ebd96824721e7d1c90855b021589feb5fb27df51f5568c3"></a>

<a id="canonical-ad7000ef84833cfff5e65510a6998e78ebd6b586658698e612f241f1f3793ac1"></a>

## namespace property — service.deploy_options.deploy_re_sites.site / 32f9cfbb03a7 / 5

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

<a id="canonical-11701f6576245e232585c828b24e9a1ad41d14ef8331570c37c20290b0ea0867"></a>

<a id="canonical-4f78cc4ae1f21971a306ad6396ae0552fab7648ca2f7e71b5aa87a7d32f498d6"></a>

## tenant property — service.deploy_options.deploy_re_sites.site / 32f9cfbb03a7 / 6

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

<a id="canonical-4c617674357e58659e339bf6e6faac050677b56991b7a283d8701b0beb8c8080"></a>

## Next pages — service.deploy_options.deploy_re_sites.site / 32f9cfbb03a7 / 7

- [service.deploy_options.deploy_re_sites](data-sources--workload--reference--group-015.md#canonical-73b8a1786044b7eb659373d38c873cfa2d98caa0d71ed74406f1a394a16f97ae)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-4efe9b660641250eaef1573c02be5b28c5b6f1c1bdf439c1e9acf505e56b3874"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-572970648da117e81356ad35ab18c815669412f9023b05bc95eb8ed36f64707a"></a>

## service.deploy_options.deploy_re_virtual_sites — service.deploy_options.deploy_re_virtual_sites / 8e0c78583684 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.deploy_options](data-sources--workload--reference--group-015.md#canonical-1938114db94be122e9e6f32ed52bcf7ce88fba44423daeb9b9b3cb3545bbc9b8)
- service.deploy_options.deploy_re_virtual_sites

<a id="canonical-61d893d5b6b3723a2ee2c5e03d587c460a9babac37cb14f9119f974fad651422"></a>

Type: `"single"`. Computed.

Defines a way to deploy a workload on specific Regional Edge virtual sites.

Upstream description:

This defines a way to deploy a workload on specific Regional Edge virtual sites.

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

<a id="canonical-061fdadc254257f15210c18de5c25367e239c028d0da247c3bee9ae8a83209ac"></a>

## Direct properties — service.deploy_options.deploy_re_virtual_sites / 8e0c78583684 / 3

- [virtual_site](data-sources--workload--reference--group-015.md#canonical-6eea28a118c53149a19b053f9221a75e3b496fdb0e5661fa9af1c1d1579c9d9e): complete subsection reference.

<a id="canonical-3aee3771156bbb44e7882711a0a364da6001e5ded15cecb0759b143840e54a7d"></a>

## Next pages — service.deploy_options.deploy_re_virtual_sites / 8e0c78583684 / 4

- [service.deploy_options.deploy_re_virtual_sites.virtual_site](data-sources--workload--reference--group-015.md#canonical-6eea28a118c53149a19b053f9221a75e3b496fdb0e5661fa9af1c1d1579c9d9e)
- [service.deploy_options](data-sources--workload--reference--group-015.md#canonical-1938114db94be122e9e6f32ed52bcf7ce88fba44423daeb9b9b3cb3545bbc9b8)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-6eea28a118c53149a19b053f9221a75e3b496fdb0e5661fa9af1c1d1579c9d9e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c0c3b9338496df26e9a9c2cdd62860c1a52bf248cefec3ff188b5486b5cbe0a5"></a>

## service.deploy_options.deploy_re_virtual_sites.virtual_site — service.deploy_options.deploy_re_virtual_sites.virtual_site / 022810bd998b / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.deploy_options](data-sources--workload--reference--group-015.md#canonical-1938114db94be122e9e6f32ed52bcf7ce88fba44423daeb9b9b3cb3545bbc9b8)
- [service.deploy_options.deploy_re_virtual_sites](data-sources--workload--reference--group-015.md#canonical-4efe9b660641250eaef1573c02be5b28c5b6f1c1bdf439c1e9acf505e56b3874)
- service.deploy_options.deploy_re_virtual_sites.virtual_site

<a id="canonical-df9c0e1b28e2cbebff04823794de14d66b851c21d3e23b0ede894e762e975c7f"></a>

Type: `"list"`. Computed.

Which regional edge virtual sites should this workload be deployed.

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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-30b0fabf693e64757c38bacecde03d8cd8013a0b21b8b45b17be2e8c4e8a24fc"></a>

## Direct properties — service.deploy_options.deploy_re_virtual_sites.virtual_site / 022810bd998b / 3

<a id="canonical-837039a89a6f88c1848fd550775f4d088b8bc2f1d97815ef23082c73aca6cfe2"></a>

<a id="canonical-8a5e3770825f4c92503ccc039c6351492219a86c41291ec7d5102022132647dd"></a>

## name property — service.deploy_options.deploy_re_virtual_sites.virtual_site / 022810bd998b / 4

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

<a id="canonical-3b3d33e8843c9c2b307174a7778e8a4b5bebd36ba3a33872afcdb5bfb7ee6213"></a>

<a id="canonical-cdb59a88ad275a17d313955ae47ad502e7b110e580c138b6da6755edf7e3f434"></a>

## namespace property — service.deploy_options.deploy_re_virtual_sites.virtual_site / 022810bd998b / 5

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

<a id="canonical-6bb670126e3f7bc5d20624646b2338ad85ad81aa1f4beee134164dd31cc8f9a2"></a>

<a id="canonical-4a87c203acbf51c07276a79ebb42072d8e97ecdc03db5652d99c479995699959"></a>

## tenant property — service.deploy_options.deploy_re_virtual_sites.virtual_site / 022810bd998b / 6

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

<a id="canonical-34c9d7ca573c2da28d80fe39f65565144a855cc608678f7c2b84a137fb2c24da"></a>

## Next pages — service.deploy_options.deploy_re_virtual_sites.virtual_site / 022810bd998b / 7

- [service.deploy_options.deploy_re_virtual_sites](data-sources--workload--reference--group-015.md#canonical-4efe9b660641250eaef1573c02be5b28c5b6f1c1bdf439c1e9acf505e56b3874)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-6ebab7e2e2f069d9b9109fa69dd5907985757d4d36cc6e56ff9fd00e40411269"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2d5816aa5e5023519ea30ad1fa984d4d88f780a78fde23f2014450fc9199e6aa"></a>

## service.scale_to_zero — service.scale_to_zero / 9d880d13eb21 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- service.scale_to_zero

<a id="canonical-4a92133dc95772c553f9b5cb65de9ec69ea21ba6ae64df732b94a697ff7e4eb6"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for scale to zero.

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

<a id="canonical-26f5e260101f8a94bffb99e89ffbc9d4fe2b5e2da8bda749e98e2993907591ee"></a>

## Direct properties — service.scale_to_zero / 9d880d13eb21 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-797096f8d8ccfca8c1f9eb51c6f14a9a0b28eb5b2506e31f116a15e34ec1417a"></a>

## Next pages — service.scale_to_zero / 9d880d13eb21 / 4

- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-8650f561bb6addd9f137c4e4b67ba52ffb9efd5338ba2a38ed06a7df7fb0e7b4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
