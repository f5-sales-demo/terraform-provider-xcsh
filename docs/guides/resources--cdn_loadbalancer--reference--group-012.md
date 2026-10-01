---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-dd11ceec830979f93bf9e12d68b4ad0a20ccc2e807a1a41831c299bad48045f7"></a>

## Next pages — jwt_validation.token_location / 9809b0becced / 4

- [jwt_validation.token_location.bearer_token](resources--cdn_loadbalancer--reference--group-012.md#canonical-b772b8cbe6dcb9788c6cee31a05ff241c10c084208939551659cfd5657f93b7c)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-011.md#canonical-ed941b0598c35b0b9845ee273333bbd7ddfe716ab488784720e0406aca406b8d)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-b772b8cbe6dcb9788c6cee31a05ff241c10c084208939551659cfd5657f93b7c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fbb435920529336a1fa90b73802f0975d6fe28989ea0db28e686922da0abb34e"></a>

## jwt_validation.token_location.bearer_token — jwt_validation.token_location.bearer_token / 9653ad36daf1 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-011.md#canonical-ed941b0598c35b0b9845ee273333bbd7ddfe716ab488784720e0406aca406b8d)
- [jwt_validation.token_location](resources--cdn_loadbalancer--reference--group-011.md#canonical-b203809ffe9d8acf4752853f2fe13f88799f51ab1c589aaa0bc8e392624c98c2)
- jwt_validation.token_location.bearer_token

<a id="canonical-5a07a523d4c6ffd1a3d95f7bddbc5979c45f95e2ccb8709ad846419f5efdc46b"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for bearer token.

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

Terraform syntax:

```terraform
bearer_token {}
```

<a id="canonical-19a3c25711bc26ab0a0edcff50bfc3754cc4ee965121c8b5c360eae603404c30"></a>

## Direct properties — jwt_validation.token_location.bearer_token / 9653ad36daf1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4c2b8ddc101e57fd048dc5fc8be512ccde4a4bcbbbf45ecbd8358ac805eaca48"></a>

## Next pages — jwt_validation.token_location.bearer_token / 9653ad36daf1 / 4

- [jwt_validation.token_location](resources--cdn_loadbalancer--reference--group-011.md#canonical-b203809ffe9d8acf4752853f2fe13f88799f51ab1c589aaa0bc8e392624c98c2)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-722c0532d8572d916e96bd684dfab040b2fc03c7412dd033ee3fd232b8eca199"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-efa3063b5d3514af3d82456e9642742c25715f3bd8a7c387692c0c68ecb8442f"></a>

## l7_ddos_action_block — l7_ddos_action_block / 53ac58700781 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- l7_ddos_action_block

<a id="canonical-108416d0bb48ac93a1212a43300016a083327d9704cdde2b857d023b1255fe54"></a>

Type: `["object", {}]`. Optional.

\[OneOf: l7\_ddos\_action\_block, l7\_ddos\_action\_default, l7\_ddos\_action\_js\_challenge;
Default: l7\_ddos\_action\_default\] Enable this option

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

OneOf alternatives in this subsection:

- [l7_ddos_action_block](resources--cdn_loadbalancer--reference--group-012.md#canonical-108416d0bb48ac93a1212a43300016a083327d9704cdde2b857d023b1255fe54)
- [l7_ddos_action_default](resources--cdn_loadbalancer--reference--group-012.md#canonical-cd309cf6f96d36f96139852b71ecb594707781ec54de79dfbb076b8beff6030d)
- [l7_ddos_action_js_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-8bdbd2b078351491aece4982f44aee0ae7664b3f75ad5c7ba9668339bd9c585d)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
l7_ddos_action_block = {}
```

<a id="canonical-d8745bd5297428cb72383e12c5a2279c7d79df5b718dba75233185dad1106961"></a>

## Direct properties — l7_ddos_action_block / 53ac58700781 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-173dc7c60a7ca962ad6bbeab17f061fb62d11608ca08bf740bda24dd1815040c"></a>

## Next pages — l7_ddos_action_block / 53ac58700781 / 4

- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-acf231ce5ffcd5d0095980e96c8d1e272272f831c0f9d1a7e3ea58cd5daaf982"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-30b82018f5e8e0a1dc40d88fc86200dec94770f4b11765965f1b3ad7cbd3edb6"></a>

## l7_ddos_action_default — l7_ddos_action_default / b7ab48459779 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- l7_ddos_action_default

<a id="canonical-cd309cf6f96d36f96139852b71ecb594707781ec54de79dfbb076b8beff6030d"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
l7_ddos_action_default = {}
```

<a id="canonical-4c80fd7bee452006e9fbafd4fba38d2c960c247c8e77c100a9887014fcaef814"></a>

## Direct properties — l7_ddos_action_default / b7ab48459779 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-dd91953cc506f43a3ea8b09b5a7de5a4778e5eae31f59966a588ca4efa9627c3"></a>

## Next pages — l7_ddos_action_default / b7ab48459779 / 4

- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-ec16815ce3b83fcc24968fa670a2b071b426e9de1613736dc0f5ac839f891325"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f0f37b0483d62645559704bcb92735391f9c412c4846df932250f3a3b903db10"></a>

## l7_ddos_action_js_challenge — l7_ddos_action_js_challenge / e6dc10857a33 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- l7_ddos_action_js_challenge

<a id="canonical-8bdbd2b078351491aece4982f44aee0ae7664b3f75ad5c7ba9668339bd9c585d"></a>

Type: `"object"`. single nested block, Optional.

Enables loadbalancer to perform client browser compatibility test by redirecting to a page with
Javascript. With this feature enabled, only clients that are capable of executing Javascript(mostly
browsers) will be allowed to complete the HTTP request. When loadbalancer is configured to do..

Upstream description:

Enables loadbalancer to perform client browser compatibility test by redirecting to a page with
Javascript.

With this feature enabled, only clients that are capable of executing Javascript(mostly browsers)
will be allowed to complete the HTTP request.

When loadbalancer is configured to do Javascript Challenge, it will redirect the browser to an HTML
page on every new HTTP request. This HTML page will have Javascript embedded in it. Loadbalancer
chooses a set of random numbers for every new client and sends these numbers along with an encrypted
answer with the request such that it embed these numbers as input in the Javascript. Javascript will
run on the requester browser and perform a complex Math operation. Script will submit the answer to
loadbalancer. Loadbalancer will validate the answer by comparing the calculated answer with the
decrypted answer (which was encrypted when it was sent back as reply) and allow the request to the
upstream server only if the answer is correct. Loadbalancer will tag response header with a cookie
to avoid Javascript challenge for subsequent requests.

Javascript challenge serves following purposes \* Validate that the request is coming via a browser
that is capable for running Javascript \* Force the browser to run a complex operation, f(X), that
requires it to spend a large number of CPU cycles. This is to slow down a potential DoS attacker by
making it difficult to launch a large request flood without having to spend even larger CPU cost at
their end.

You can enable either Javascript challenge or Captcha challenge on a virtual host.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cookie_expiry",
    "js_script_delay")}
```

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

Terraform syntax:

```terraform
l7_ddos_action_js_challenge {
  # Configure direct properties listed below.
}
```

<a id="canonical-c189b405e125e0f46fa7a819d9db4d9b0c125109cc2dadae0b46d344ea026748"></a>

## Direct properties — l7_ddos_action_js_challenge / e6dc10857a33 / 3

<a id="canonical-b13a186ef180c0657aa44911c7a281d45ed9ef1c5541c09e83703f46617ff7ba"></a>

<a id="canonical-fbcc54f0b20c37f792f11e90b446e69a536392c60a63349dd705cfbb60862d28"></a>

## cookie_expiry property — l7_ddos_action_js_challenge / e6dc10857a33 / 4

Type: `"number"`. Optional.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Upstream description:

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 86400),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 86400,
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

<a id="canonical-37d77a0347d26e48a317390f275051a10c2a208313db6525dcc7b5ece722ed7e"></a>

<a id="canonical-dc1900b6da56ed4e47bfd989fe5ded399dc947b6228d106bfc131b83d03fe331"></a>

## custom_page property — l7_ddos_action_js_challenge / e6dc10857a33 / 5

Type: `"string"`. Optional.

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format.

Upstream description:

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". Base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(65536),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 65536,
  "x-f5xc-constraints": {
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
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-2274bcf7934097317b64740f60993decba6a8fc7c3618e4c10a33bca418ccab1"></a>

<a id="canonical-cd606806f4a9caab114b7ea5502091ec27abcafafd6fe6d011f81135c4c3ca10"></a>

## js_script_delay property — l7_ddos_action_js_challenge / e6dc10857a33 / 6

Type: `"number"`. Optional.

Delay introduced by Javascript, in milliseconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1000, 60000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1000
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1000",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1000",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

<a id="canonical-a040d649eb25811552c06a8902ce5b51bfbafbe37e91e28adac05f18b4f37805"></a>

## Next pages — l7_ddos_action_js_challenge / e6dc10857a33 / 7

- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-f6005363adc29c9893da440ca4e14b77a49e5969955121dcce8465038b52940d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-63dfc29854fdc502115c8732dc7fd2f4a318992e96d6575fc8b25eeac382aeed"></a>

## no_challenge — no_challenge / ea54d28898dc / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- no_challenge

<a id="canonical-daead701143ebd9525dd2c685a8cb14eb0f1a03fb72c0f48599432246cfec8f5"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no challenge.

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

Terraform syntax:

```terraform
no_challenge = {}
```

<a id="canonical-bf73de029e5bd4df5e226700fcf301de864454faf08f09237c95c508952981cd"></a>

## Direct properties — no_challenge / ea54d28898dc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-df08991283977a0a2c4e40977d4cbd022754a568925b03bb1ab59c637773d3d6"></a>

## Next pages — no_challenge / ea54d28898dc / 4

- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-db5bd1441ea30ff501c1ef1d5ef4672bf5bb2236a075287daff00db31ea5a53b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-592b8df1a1fd5069425b986e9035c7934efceca8c447769f1fa194842fa533f5"></a>

## no_service_policies — no_service_policies / 9dfe76d810f6 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- no_service_policies

<a id="canonical-47d30b47b99bf8c43e28d8dbd60a7c59c4978d1cef8bf16b5d55004cb441d60b"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no service policies.

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

Terraform syntax:

```terraform
no_service_policies = {}
```

<a id="canonical-d321e0fffd0e422612ed38a8eedfac34e726823d7491e61b2ae95e1d3d602df2"></a>

## Direct properties — no_service_policies / 9dfe76d810f6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8cd05012a61063a19187f32560f93fe218c214c57bc7dc8fe036cfd94445bdde"></a>

## Next pages — no_service_policies / 9dfe76d810f6 / 4

- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-603a6ee1a88486640208f36e01e4e00f68db830979076e4ad137090fd9373676"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6743739abe4a06967653463d19f76b4574dee4d5f4c9e9bd503f73dff5ad5e03"></a>

## origin_pool — origin_pool / 270d7af1f39d / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- origin_pool

<a id="canonical-7f2f0baec8604ad5a43bfef2e2f9b42fa2045324d8bcbf4c059c3ed59c280744"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for origin pool.

Upstream description:

Origin Pool for the CDN distribution.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("origin_servers"),
  validators.ConflictingObjectAttributes("no_tls",
    "use_tls")}
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
  "x-ves-oneof-field-tls_choice": "[\"no_tls\",\"use_tls\"]"
}
```

Terraform syntax:

```terraform
origin_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-b01575044f1b242907b123e1f562d850201485c5b19e877d5d12df4e8cc66313"></a>

## Direct properties — origin_pool / 270d7af1f39d / 3

- [more_origin_options](resources--cdn_loadbalancer--reference--group-012.md#canonical-27811f6d981d465d64ccb7c54c7f75cee5c05f55e77052d889e1fc3ead6bf292): complete subsection reference.

- [no_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0de2b216fa16de709f9ccc84c3cec67086617673d4ec7225820eedf3eba8e387): complete subsection reference.

<a id="canonical-d9e0afad3c4ca655d3e1c47579e2ef7eb375e1e1a112897ea9f95e61a01b79cf"></a>

<a id="canonical-60707245b6f749a5e5756a9135cad22cb823cfe09a4c13feea49e58da4951ea4"></a>

## origin_request_timeout property — origin_pool / 270d7af1f39d / 4

Type: `"string"`. Optional.

Configures the time after which a request to the origin will time out waiting for a response.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_time_interval": "10m",
    "ves.io.schema.rules.string.min_time_interval": "10s",
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_time_interval": "10m",
    "ves.io.schema.rules.string.min_time_interval": "10s",
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

- [origin_servers](resources--cdn_loadbalancer--reference--group-012.md#canonical-2aa032604efbb54188b2b1598036ae5f8093dc714dafbae721e8bcc09281010f): complete subsection reference.

- [public_name](resources--cdn_loadbalancer--reference--group-012.md#canonical-4242ea1ebb2a5bfd6878cdb2aaff691de32bfd6875a61fb269a7f51c48907eae): complete subsection reference.

- [use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-213416fcbc6044c18fab1b3315ab81a7eed3b03909873e5492aeff7c88e527e0): complete subsection reference.

<a id="canonical-85632fc5954b5414fa7ad5f4872066759acecfa0858ad6ee34a0b8ee0df54347"></a>

## Next pages — origin_pool / 270d7af1f39d / 5

- [origin_pool.more_origin_options](resources--cdn_loadbalancer--reference--group-012.md#canonical-27811f6d981d465d64ccb7c54c7f75cee5c05f55e77052d889e1fc3ead6bf292)
- [origin_pool.no_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0de2b216fa16de709f9ccc84c3cec67086617673d4ec7225820eedf3eba8e387)
- [origin_pool.origin_servers](resources--cdn_loadbalancer--reference--group-012.md#canonical-2aa032604efbb54188b2b1598036ae5f8093dc714dafbae721e8bcc09281010f)
- [origin_pool.public_name](resources--cdn_loadbalancer--reference--group-012.md#canonical-4242ea1ebb2a5bfd6878cdb2aaff691de32bfd6875a61fb269a7f51c48907eae)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-213416fcbc6044c18fab1b3315ab81a7eed3b03909873e5492aeff7c88e527e0)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-27811f6d981d465d64ccb7c54c7f75cee5c05f55e77052d889e1fc3ead6bf292"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-103ae6783a8a9db2675d4b63202189519c7257c48db337872983def468480029"></a>

## origin_pool.more_origin_options — origin_pool.more_origin_options / fafa8b99dd68 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-603a6ee1a88486640208f36e01e4e00f68db830979076e4ad137090fd9373676)
- origin_pool.more_origin_options

<a id="canonical-f28b28d34f2719dcf2b17a9d0faf184a7305fde784c3132139d0aa6d132fa15f"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for more origin options.

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

Terraform syntax:

```terraform
more_origin_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-2694605deb6b4a0f6062e2824ab3dff117581c6a17a6d8c973b49b3050579d17"></a>

## Direct properties — origin_pool.more_origin_options / fafa8b99dd68 / 3

<a id="canonical-a903069c994d60e969f82a677e00b54e32f435c71965e81642e3561b0030dd9e"></a>

<a id="canonical-7178750e720ecc7d16711afcf414bd345e6e811076a8abb5387da395babb1211"></a>

## enable_byte_range_request property — origin_pool.more_origin_options / fafa8b99dd68 / 4

Type: `"bool"`. Optional.

Choice to enable/disable byte range requests towards origin.

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

<a id="canonical-76a2c5cf133be50886afb7f8a1365acb1256826989765482dd541ea1f66fcf2f"></a>

<a id="canonical-933ef7207a1b22b8bc747e1a258c51318b9b412820f3ee4ec2a5b57a3895ba4c"></a>

## websocket_proxy property — origin_pool.more_origin_options / fafa8b99dd68 / 5

Type: `"bool"`. Optional.

Option to enable proxying of websocket connections to the origin server.

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

<a id="canonical-b5e44e978ceefed133104c0452508cba332e006e79721dc6e1997c599c0de0c3"></a>

## Next pages — origin_pool.more_origin_options / fafa8b99dd68 / 6

- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-603a6ee1a88486640208f36e01e4e00f68db830979076e4ad137090fd9373676)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-0de2b216fa16de709f9ccc84c3cec67086617673d4ec7225820eedf3eba8e387"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c62264d767c1c5c1ddb35f5cfadf16d3c3f53e4a797bcc5ab88cf27576efb85f"></a>

## origin_pool.no_tls — origin_pool.no_tls / 01a5ba9e7a36 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-603a6ee1a88486640208f36e01e4e00f68db830979076e4ad137090fd9373676)
- origin_pool.no_tls

<a id="canonical-63054df8e80f8977362ad08ca0e95547dfe692eaca0a3836031773f308f2c4ea"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
no_tls = {}
```

<a id="canonical-6d5bd908bc2baf33ee01a606e965814b48d52a0a28c6c2d0f5378215aab056f9"></a>

## Direct properties — origin_pool.no_tls / 01a5ba9e7a36 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-291806ab3822aea57c8573c5a6bbb663ff0926c6dce1b017552b45ae849f7627"></a>

## Next pages — origin_pool.no_tls / 01a5ba9e7a36 / 4

- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-603a6ee1a88486640208f36e01e4e00f68db830979076e4ad137090fd9373676)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-2aa032604efbb54188b2b1598036ae5f8093dc714dafbae721e8bcc09281010f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-55d072f83959c978eff49543305759a51de43f89c221cec898f789c60a4be7b2"></a>

## origin_pool.origin_servers — origin_pool.origin_servers / c965b57eacf8 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-603a6ee1a88486640208f36e01e4e00f68db830979076e4ad137090fd9373676)
- origin_pool.origin_servers

<a id="canonical-a56250f0b905da61d17575c52f596303cd144b56d87dbba748d0a002a788f55c"></a>

Type: `"object"`. list nested block, Optional.

List Of Origin Servers. List of original servers.

Upstream description:

List of original servers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("public_ip",
    "public_name")}
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

Terraform syntax:

```terraform
origin_servers {
  # Configure direct properties listed below.
}
```

<a id="canonical-2e883f430589fca31d207629d6d8e7e2a969720cec783f8fc16de30ac297ffb6"></a>

## Direct properties — origin_pool.origin_servers / c965b57eacf8 / 3

<a id="canonical-b92a48ff74980f53f7b6032b6de92b7c23c6f4e9859f0826afb4804772504040"></a>

<a id="canonical-6344ceea0558c74d177e4612848087504e7b57556f4947b77c792def04a9d22e"></a>

## port property — origin_pool.origin_servers / c965b57eacf8 / 4

Type: `"number"`. Optional.

Port the workload can be reached on Enter a custom port only if your origin server uses a
non-default port. Leave the value as 0 to automatically use 443 (TLS) or 80 (non-TLS).

Upstream description:

Port the workload can be reached on Enter a custom port only if your origin server uses a
non-default port. Leave the value as 0 to automatically use 443 (TLS) or 80 (non-TLS).

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
}
```

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
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

- [public_ip](resources--cdn_loadbalancer--reference--group-012.md#canonical-fdad4bd94271dc4de7197c3f54894a75ed1fb0621cc4cf7219062168c29e4373): complete subsection reference.

- [public_name](resources--cdn_loadbalancer--reference--group-012.md#canonical-c095575d097dd8caa5807a2ad19c33fbfd185eb0f867ac8fe7db7c8b377ac6e7): complete subsection reference.

<a id="canonical-2383243655c34d07add945a8659e418dfd6d89724d8074a21fbad800cae45a2e"></a>

## Next pages — origin_pool.origin_servers / c965b57eacf8 / 5

- [origin_pool.origin_servers.public_ip](resources--cdn_loadbalancer--reference--group-012.md#canonical-fdad4bd94271dc4de7197c3f54894a75ed1fb0621cc4cf7219062168c29e4373)
- [origin_pool.origin_servers.public_name](resources--cdn_loadbalancer--reference--group-012.md#canonical-c095575d097dd8caa5807a2ad19c33fbfd185eb0f867ac8fe7db7c8b377ac6e7)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-603a6ee1a88486640208f36e01e4e00f68db830979076e4ad137090fd9373676)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-fdad4bd94271dc4de7197c3f54894a75ed1fb0621cc4cf7219062168c29e4373"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c679f59100e31c8539f5a6b29f1be9a91376222f9e4575804840e4113fa06b44"></a>

## origin_pool.origin_servers.public_ip — origin_pool.origin_servers.public_ip / c6bcd2c7a753 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-603a6ee1a88486640208f36e01e4e00f68db830979076e4ad137090fd9373676)
- [origin_pool.origin_servers](resources--cdn_loadbalancer--reference--group-012.md#canonical-2aa032604efbb54188b2b1598036ae5f8093dc714dafbae721e8bcc09281010f)
- origin_pool.origin_servers.public_ip

<a id="canonical-e07837d4453b0e27335a864c161e15b274d09962c42d478d0e2f7d1dc3bb5868"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with public IP address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-public_ip_choice": "[\"ip\"]"
}
```

Terraform syntax:

```terraform
public_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-bad9d961dbc3e7196620ab2ff6c58ab269c62f65c42c05aaa1023cec0548ac8a"></a>

## Direct properties — origin_pool.origin_servers.public_ip / c6bcd2c7a753 / 3

<a id="canonical-af3d51b4b34e7c2ecc1c3a2a502e125c6da6d0b5d3934340b1088c303f33e4bf"></a>

<a id="canonical-84e03e534894690ff40f4b3139b16d6b7eb019931d8a7318d101be126baf740c"></a>

## ip property — origin_pool.origin_servers.public_ip / c6bcd2c7a753 / 4

Type: `"string"`. Optional.

Public IPv4. Exclusive with \[\] Public IPv4 address.

Upstream description:

Exclusive with \[\] Public IPv4 address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-c41e27b226e349d2e31b88ef4586d3b064160a25f45832a9a3e1f0a66dc47600"></a>

## Next pages — origin_pool.origin_servers.public_ip / c6bcd2c7a753 / 5

- [origin_pool.origin_servers](resources--cdn_loadbalancer--reference--group-012.md#canonical-2aa032604efbb54188b2b1598036ae5f8093dc714dafbae721e8bcc09281010f)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-c095575d097dd8caa5807a2ad19c33fbfd185eb0f867ac8fe7db7c8b377ac6e7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-46776233e0302c724c4a47e284d6e3c54a751f7bf7f4089957aed48d48227fd7"></a>

## origin_pool.origin_servers.public_name — origin_pool.origin_servers.public_name / 1f1b0e4387f7 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-603a6ee1a88486640208f36e01e4e00f68db830979076e4ad137090fd9373676)
- [origin_pool.origin_servers](resources--cdn_loadbalancer--reference--group-012.md#canonical-2aa032604efbb54188b2b1598036ae5f8093dc714dafbae721e8bcc09281010f)
- origin_pool.origin_servers.public_name

<a id="canonical-4f1bb11e081e0c2640bdad4a5afb44bfd51bab26b973b34b69ed10e2d967c483"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with public DNS name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("dns_name")}
```

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

Terraform syntax:

```terraform
public_name {
  # Configure direct properties listed below.
}
```

<a id="canonical-fe8ae72deb7160f5bb86e1e97419ad55d743fcb2255a666685e435870c7bc1f4"></a>

## Direct properties — origin_pool.origin_servers.public_name / 1f1b0e4387f7 / 3

<a id="canonical-ffc3ccf07827b1119d8357c32d7acaf1807c4ea21a51306c43ab565c1bc001d5"></a>

<a id="canonical-84164f890ade137cc84bbe27d7c1e741e89995e37606c1ba00cbc613657d0981"></a>

## dns_name property — origin_pool.origin_servers.public_name / 1f1b0e4387f7 / 4

Type: `"string"`. Optional.

DNS Name. DNS Name

Upstream description:

DNS Name

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9.-]",
      "description": "Dot-separated DNS labels"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "formatDescription": "RFC 1123 FQDN: lowercase, dot-separated labels, max 253 chars total",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-dec4521f8b34b59b15d1a074198018685c497ee8f7c563507af0fdaf5a394140"></a>

<a id="canonical-6f973fbfc13f07a1c1a1ebc7235b6cc4f58c706da874b88532f49a1249cca41b"></a>

## refresh_interval property — origin_pool.origin_servers.public_name / 1f1b0e4387f7 / 5

Type: `"number"`. Optional.

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Upstream description:

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  validators.Int64RangeSetValidator(
    validators.Int64Range{Minimum: 0, Maximum: 0},
    validators.Int64Range{Minimum: 10, Maximum: 604800},
  ),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 604800,
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
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  }
}
```

<a id="canonical-6859daf1a1cc75fbc55d3aadef1593311ae6320a13f61f6a3d82d6af26fb1ce9"></a>

## Next pages — origin_pool.origin_servers.public_name / 1f1b0e4387f7 / 6

- [origin_pool.origin_servers](resources--cdn_loadbalancer--reference--group-012.md#canonical-2aa032604efbb54188b2b1598036ae5f8093dc714dafbae721e8bcc09281010f)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-4242ea1ebb2a5bfd6878cdb2aaff691de32bfd6875a61fb269a7f51c48907eae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ebd3e3f0bc355f1f96be13dc0748c60c2236babf56c25707b660296e1274851d"></a>

## origin_pool.public_name — origin_pool.public_name / 90ec7b0f585c / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-603a6ee1a88486640208f36e01e4e00f68db830979076e4ad137090fd9373676)
- origin_pool.public_name

<a id="canonical-3388831623d6cb0154f1ed442f88075984042fe808e34b90bf5bf0079fd2c027"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with public DNS name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("dns_name")}
```

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

Terraform syntax:

```terraform
public_name {
  # Configure direct properties listed below.
}
```

<a id="canonical-61686707b4c7ee5616d08cdaab4ef8229cb77115667922c137a4519eda3cb322"></a>

## Direct properties — origin_pool.public_name / 90ec7b0f585c / 3

<a id="canonical-2bc30ce48bdbf61e435fd4aeea67cc41ef9a92a1b22e1d4da0c19abb42ea9e4a"></a>

<a id="canonical-1cbc2766148d258d125727bc0de6a2e4e7250823ce0a5531933fa96ea5525b26"></a>

## dns_name property — origin_pool.public_name / 90ec7b0f585c / 4

Type: `"string"`. Optional.

DNS Name. DNS Name

Upstream description:

DNS Name

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9.-]",
      "description": "Dot-separated DNS labels"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "formatDescription": "RFC 1123 FQDN: lowercase, dot-separated labels, max 253 chars total",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-642710e2d110ae25113e3e9fe9dfe8258c593e7781405e8d4bd1f2a87153a364"></a>

<a id="canonical-b2ef8c4081b4816b1f1fab23b3495f6364c066d138c7720efcc67e4a11d87cc7"></a>

## refresh_interval property — origin_pool.public_name / 90ec7b0f585c / 5

Type: `"number"`. Optional.

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Upstream description:

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  validators.Int64RangeSetValidator(
    validators.Int64Range{Minimum: 0, Maximum: 0},
    validators.Int64Range{Minimum: 10, Maximum: 604800},
  ),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 604800,
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
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  }
}
```

<a id="canonical-a79388a9868b50066faf65474f42b5c291811e07231902e93e81e799ba083426"></a>

## Next pages — origin_pool.public_name / 90ec7b0f585c / 6

- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-603a6ee1a88486640208f36e01e4e00f68db830979076e4ad137090fd9373676)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-213416fcbc6044c18fab1b3315ab81a7eed3b03909873e5492aeff7c88e527e0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7ed361b610d2fa896509a768ab2bbeb5fd3b6ab5fe8802f90c35711df6ea1ae0"></a>

## origin_pool.use_tls — origin_pool.use_tls / e292ed730189 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-603a6ee1a88486640208f36e01e4e00f68db830979076e4ad137090fd9373676)
- origin_pool.use_tls

<a id="canonical-4dc86c45cb57e6d7736662647e4b71e65cfc4514a4aadb65bef3b90c202f4ef8"></a>

Type: `"object"`. single nested block, Optional.

TLS Parameters for Origin Servers. Upstream TLS Parameters.

Upstream description:

Upstream TLS Parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_session_key_caching",
    "disable_session_key_caching"),
  validators.ConflictingObjectAttributes("default_session_key_caching",
    "max_session_keys"),
  validators.ConflictingObjectAttributes("disable_session_key_caching",
    "max_session_keys"),
  validators.ConflictingObjectAttributes("disable_sni",
    "sni"),
  validators.ConflictingObjectAttributes("disable_sni",
    "use_host_header_as_sni"),
  validators.ConflictingObjectAttributes("no_mtls",
    "use_mtls"),
  validators.ConflictingObjectAttributes("no_mtls",
    "use_mtls_obj"),
  validators.ConflictingObjectAttributes("skip_server_verification",
    "use_server_verification"),
  validators.ConflictingObjectAttributes("skip_server_verification",
    "volterra_trusted_ca"),
  validators.ConflictingObjectAttributes("sni",
    "use_host_header_as_sni"),
  validators.ConflictingObjectAttributes("use_mtls",
    "use_mtls_obj"),
  validators.ConflictingObjectAttributes("use_server_verification",
    "volterra_trusted_ca")}
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
  "x-ves-oneof-field-max_session_keys_type": "[\"default_session_key_caching\",\"disable_session_key_caching\",\"max_session_keys\"]",
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\",\"use_mtls_obj\"]",
  "x-ves-oneof-field-server_validation_choice": "[\"skip_server_verification\",\"use_server_verification\",\"volterra_trusted_ca\"]",
  "x-ves-oneof-field-sni_choice": "[\"disable_sni\",\"sni\",\"use_host_header_as_sni\"]"
}
```

Terraform syntax:

```terraform
use_tls {
  # Configure direct properties listed below.
}
```

<a id="canonical-db385596f37445cf7be7cbe1a3612385264e15d66b262b36372f6ee6537251d0"></a>

## Direct properties — origin_pool.use_tls / e292ed730189 / 3

- [default_session_key_caching](resources--cdn_loadbalancer--reference--group-012.md#canonical-32e3c7998095ab542f0be4b722bd81e5edcd880ff638f6037856328382e4bc85): complete subsection reference.

- [disable_session_key_caching](resources--cdn_loadbalancer--reference--group-012.md#canonical-5aa19f59cd5a7a6929f6e49f1a43c33d9fa1ae4c72e2a6fedb6de4c2e3e0b306): complete subsection reference.

- [disable_sni](resources--cdn_loadbalancer--reference--group-012.md#canonical-42c5343cb92e1c0338886a55aa3fe1311b2c1614bf23a12525a7fb33b78eaad3): complete subsection reference.

<a id="canonical-9a00a6e27956e0d8c49e6702bdb79dfaa66a69742caffc0c5360f5ba8b8b22c9"></a>

<a id="canonical-55cabaade892f0e09762a51bcb9c09b998f72286a3e8d87b0937912e6295db23"></a>

## max_session_keys property — origin_pool.use_tls / e292ed730189 / 4

Type: `"number"`. Optional.

Exclusive with \[default\_session\_key\_caching disable\_session\_key\_caching\] Number of session
keys that are cached.

Upstream description:

Exclusive with \[default\_session\_key\_caching disable\_session\_key\_caching\]

Number of session keys that are cached.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(2, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 2
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "64"
  }
}
```

- [no_mtls](resources--cdn_loadbalancer--reference--group-012.md#canonical-33570878ec06f2cbfd6f47241a85d143808bf503b0f405584f85624b88137713): complete subsection reference.

- [skip_server_verification](resources--cdn_loadbalancer--reference--group-012.md#canonical-2d325c2640c2e19b93cee7471584783d207e9b5f35198936e0c3014bee75de78): complete subsection reference.

<a id="canonical-6164bd85aa8eb8d2dcaa900fb55d41df4c0d3236828113eeaa34d7949c91dcfb"></a>

<a id="canonical-9c425d55c0ef3a6a88924cea734226cb6ab3de9ad05a8781065a64b964b58c02"></a>

## sni property — origin_pool.use_tls / e292ed730189 / 5

Type: `"string"`. Optional.

Exclusive with \[disable\_sni use\_host\_header\_as\_sni\] SNI value to be used.

Upstream description:

Exclusive with \[disable\_sni use\_host\_header\_as\_sni\] SNI value to be used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [tls_config](resources--cdn_loadbalancer--reference--group-012.md#canonical-5fbf27b546aabdc4f2af2cbc3fd167a146805e1786c315777301990e3296462c): complete subsection reference.

- [use_host_header_as_sni](resources--cdn_loadbalancer--reference--group-012.md#canonical-dbcf2cdbe2fd6076d6aa3893cf10d388b09145424d5ea691535e20b2fddac59d): complete subsection reference.

- [use_mtls](resources--cdn_loadbalancer--reference--group-012.md#canonical-10135e1d0e39aea2354081a9383060b17612baabf9b8d1e55f79954afb3713e8): complete subsection reference.

- [use_mtls_obj](resources--cdn_loadbalancer--reference--group-012.md#canonical-8d22410b61f69f6fcc934ed0d400e6cab770df64afe0a892a58ce3b15297e50d): complete subsection reference.

- [use_server_verification](resources--cdn_loadbalancer--reference--group-012.md#canonical-1087b1f966ea83fbe26326ef2fa2e8a742d53896d3dd13c74807c9234554e48b): complete subsection reference.

- [volterra_trusted_ca](resources--cdn_loadbalancer--reference--group-012.md#canonical-deec9d5dd45c2417e305be03115967af8c43414a8f7d65c9a8c6706dba4427cd): complete subsection reference.

<a id="canonical-6185cbb0736c40ec1fb1b0d4590c604a3f428a66c0bc25d4d2ad83ad0c7157a2"></a>

## Next pages — origin_pool.use_tls / e292ed730189 / 6

- [origin_pool.use_tls.default_session_key_caching](resources--cdn_loadbalancer--reference--group-012.md#canonical-32e3c7998095ab542f0be4b722bd81e5edcd880ff638f6037856328382e4bc85)
- [origin_pool.use_tls.disable_session_key_caching](resources--cdn_loadbalancer--reference--group-012.md#canonical-5aa19f59cd5a7a6929f6e49f1a43c33d9fa1ae4c72e2a6fedb6de4c2e3e0b306)
- [origin_pool.use_tls.disable_sni](resources--cdn_loadbalancer--reference--group-012.md#canonical-42c5343cb92e1c0338886a55aa3fe1311b2c1614bf23a12525a7fb33b78eaad3)
- [origin_pool.use_tls.no_mtls](resources--cdn_loadbalancer--reference--group-012.md#canonical-33570878ec06f2cbfd6f47241a85d143808bf503b0f405584f85624b88137713)
- [origin_pool.use_tls.skip_server_verification](resources--cdn_loadbalancer--reference--group-012.md#canonical-2d325c2640c2e19b93cee7471584783d207e9b5f35198936e0c3014bee75de78)
- [origin_pool.use_tls.tls_config](resources--cdn_loadbalancer--reference--group-012.md#canonical-5fbf27b546aabdc4f2af2cbc3fd167a146805e1786c315777301990e3296462c)
- [origin_pool.use_tls.use_host_header_as_sni](resources--cdn_loadbalancer--reference--group-012.md#canonical-dbcf2cdbe2fd6076d6aa3893cf10d388b09145424d5ea691535e20b2fddac59d)
- [origin_pool.use_tls.use_mtls](resources--cdn_loadbalancer--reference--group-012.md#canonical-10135e1d0e39aea2354081a9383060b17612baabf9b8d1e55f79954afb3713e8)
- [origin_pool.use_tls.use_mtls_obj](resources--cdn_loadbalancer--reference--group-012.md#canonical-8d22410b61f69f6fcc934ed0d400e6cab770df64afe0a892a58ce3b15297e50d)
- [origin_pool.use_tls.use_server_verification](resources--cdn_loadbalancer--reference--group-012.md#canonical-1087b1f966ea83fbe26326ef2fa2e8a742d53896d3dd13c74807c9234554e48b)
- [origin_pool.use_tls.volterra_trusted_ca](resources--cdn_loadbalancer--reference--group-012.md#canonical-deec9d5dd45c2417e305be03115967af8c43414a8f7d65c9a8c6706dba4427cd)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-603a6ee1a88486640208f36e01e4e00f68db830979076e4ad137090fd9373676)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-32e3c7998095ab542f0be4b722bd81e5edcd880ff638f6037856328382e4bc85"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3043124c96150b92ee4bb0467aefb1752588d36686ce78745f223fc277d3dd78"></a>

## origin_pool.use_tls.default_session_key_caching — origin_pool.use_tls.default_session_key_caching / d4ad6ad46a8c / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-603a6ee1a88486640208f36e01e4e00f68db830979076e4ad137090fd9373676)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-213416fcbc6044c18fab1b3315ab81a7eed3b03909873e5492aeff7c88e527e0)
- origin_pool.use_tls.default_session_key_caching

<a id="canonical-cb87bfaaeda5c1d3f7d8017c6f2d281f63b9b8342f8e2e8f572ffaa15cf2dc7f"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for default session key caching. Defaults to \`map\[\]\`. Server applies
default when omitted.

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

Terraform syntax:

```terraform
default_session_key_caching = {}
```

<a id="canonical-fe595f9642a49089ec12d498ee5ade97958a9caf28a34988e2077d11c1c8b06a"></a>

## Direct properties — origin_pool.use_tls.default_session_key_caching / d4ad6ad46a8c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-01d11f9360c40d5612e09e606d4baf8408b0b84025ee1cea08da929a87456e2e"></a>

## Next pages — origin_pool.use_tls.default_session_key_caching / d4ad6ad46a8c / 4

- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-213416fcbc6044c18fab1b3315ab81a7eed3b03909873e5492aeff7c88e527e0)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-5aa19f59cd5a7a6929f6e49f1a43c33d9fa1ae4c72e2a6fedb6de4c2e3e0b306"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3d1ca5ff44bf89e4b85d084f5ea79b6ace12e3a5731486eee8b681cbc7e13b2f"></a>

## origin_pool.use_tls.disable_session_key_caching — origin_pool.use_tls.disable_session_key_caching / 9e35e447b804 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-603a6ee1a88486640208f36e01e4e00f68db830979076e4ad137090fd9373676)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-213416fcbc6044c18fab1b3315ab81a7eed3b03909873e5492aeff7c88e527e0)
- origin_pool.use_tls.disable_session_key_caching

<a id="canonical-5405f17cd0a48311052e5b0c1911969becd941965edd9375d465a315f614a46a"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable session key caching.

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

Terraform syntax:

```terraform
disable_session_key_caching = {}
```

<a id="canonical-b490cd32770f3c891ce94d4026b2c48eda752f8a1d1e8cc495750313f784f5d8"></a>

## Direct properties — origin_pool.use_tls.disable_session_key_caching / 9e35e447b804 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-78c074a28bd3fefe7d3e059eeb1906f026135715651172feb7be2a8380e8ee5e"></a>

## Next pages — origin_pool.use_tls.disable_session_key_caching / 9e35e447b804 / 4

- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-213416fcbc6044c18fab1b3315ab81a7eed3b03909873e5492aeff7c88e527e0)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-42c5343cb92e1c0338886a55aa3fe1311b2c1614bf23a12525a7fb33b78eaad3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-be003ef695b9737d494f59b979c3574039a68cb927be47ce59f64e54bf573011"></a>

## origin_pool.use_tls.disable_sni — origin_pool.use_tls.disable_sni / e1782fec519d / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-603a6ee1a88486640208f36e01e4e00f68db830979076e4ad137090fd9373676)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-213416fcbc6044c18fab1b3315ab81a7eed3b03909873e5492aeff7c88e527e0)
- origin_pool.use_tls.disable_sni

<a id="canonical-2a45ddc0c6cb8e5761a52c8caa2801f28725e1e82f9cf0bac3a5aa632a0a971d"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable sni.

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

Terraform syntax:

```terraform
disable_sni = {}
```

<a id="canonical-17597fb868bd43568809c111ac77343846f3d17091eac101a697075711caf447"></a>

## Direct properties — origin_pool.use_tls.disable_sni / e1782fec519d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8e0d1d0f9857620d431eaab873d86415891a17ed01f939c93556923bdacb0047"></a>

## Next pages — origin_pool.use_tls.disable_sni / e1782fec519d / 4

- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-213416fcbc6044c18fab1b3315ab81a7eed3b03909873e5492aeff7c88e527e0)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-33570878ec06f2cbfd6f47241a85d143808bf503b0f405584f85624b88137713"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7ec924d631388676b4239192a4ec3da99283a59f2340fdfbc6c24c422ac9d0ea"></a>

## origin_pool.use_tls.no_mtls — origin_pool.use_tls.no_mtls / 8471dd9e04e7 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-603a6ee1a88486640208f36e01e4e00f68db830979076e4ad137090fd9373676)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-213416fcbc6044c18fab1b3315ab81a7eed3b03909873e5492aeff7c88e527e0)
- origin_pool.use_tls.no_mtls

<a id="canonical-2ab919a01d179ff5253243fc994d16adaeb530c7d5e1f245f9457f1ab801ef6c"></a>

Type: `["object", {}]`. Optional, Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

Terraform syntax:

```terraform
no_mtls = {}
```

<a id="canonical-c1cf17f012d1bda3c0da4de6a8dbe38189be110983283bd15672506cd77ea3fa"></a>

## Direct properties — origin_pool.use_tls.no_mtls / 8471dd9e04e7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8d4b59188571cd10556aa75c83dff402add714bb2a9dd4cfceb58808857d8a60"></a>

## Next pages — origin_pool.use_tls.no_mtls / 8471dd9e04e7 / 4

- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-213416fcbc6044c18fab1b3315ab81a7eed3b03909873e5492aeff7c88e527e0)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-2d325c2640c2e19b93cee7471584783d207e9b5f35198936e0c3014bee75de78"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-304797208af393127d1867dc1a1c406f2dbf98f001909c3ce485024b65ab5f80"></a>

## origin_pool.use_tls.skip_server_verification — origin_pool.use_tls.skip_server_verification / 8d2a4be22d1d / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-603a6ee1a88486640208f36e01e4e00f68db830979076e4ad137090fd9373676)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-213416fcbc6044c18fab1b3315ab81a7eed3b03909873e5492aeff7c88e527e0)
- origin_pool.use_tls.skip_server_verification

<a id="canonical-061be30216ff7313b04aa1bb7dc01b9098c7066632af6f9c7752f6bc5a8e0ac9"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
skip_server_verification = {}
```

<a id="canonical-1809a0b4c98dae5d600cfc84de5075fca31733159f2ff582ee6389dade2fef01"></a>

## Direct properties — origin_pool.use_tls.skip_server_verification / 8d2a4be22d1d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e79613f51e569243ef479c2852fca6b9b5b57b6c201390e732c84e216ee302de"></a>

## Next pages — origin_pool.use_tls.skip_server_verification / 8d2a4be22d1d / 4

- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-213416fcbc6044c18fab1b3315ab81a7eed3b03909873e5492aeff7c88e527e0)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-5fbf27b546aabdc4f2af2cbc3fd167a146805e1786c315777301990e3296462c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-23c402fd6fd059300857ba8f77070a4183be8b57ffb275f4f0c374d4752d5023"></a>

## origin_pool.use_tls.tls_config — origin_pool.use_tls.tls_config / 64629b4255c8 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-603a6ee1a88486640208f36e01e4e00f68db830979076e4ad137090fd9373676)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-213416fcbc6044c18fab1b3315ab81a7eed3b03909873e5492aeff7c88e527e0)
- origin_pool.use_tls.tls_config

<a id="canonical-84343809c0f97ee24036fea5979abcdefea8a3e03ae158ea3d68c88ed43111b5"></a>

Type: `"object"`. single nested block, Optional.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_security",
    "default_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "low_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("default_security",
    "low_security"),
  validators.ConflictingObjectAttributes("default_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("low_security",
    "medium_security")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "tls",
    "constraintType": "object",
    "deterministic": true,
    "metadata": {
      "category": "tls",
      "confidence": 0.99,
      "note": "Required when use_tls is selected — API returns 400 if nil",
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
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

Terraform syntax:

```terraform
tls_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-331efb8aecd1e77087729f1400e10aa5a1b9f4161d3b6c8799d2e2856eb3d6b1"></a>

## Direct properties — origin_pool.use_tls.tls_config / 64629b4255c8 / 3

- [custom_security](resources--cdn_loadbalancer--reference--group-012.md#canonical-2d586a03ba26fbe6aeeae1d4c7c25fa52d01577791520227587fdbb1f6a5471a): complete subsection reference.

- [default_security](resources--cdn_loadbalancer--reference--group-012.md#canonical-9dd9bd81d6182c6f7a71237eae20b2a1e41ed92d1e2c459ee67d8516f211fb4b): complete subsection reference.

- [low_security](resources--cdn_loadbalancer--reference--group-012.md#canonical-c8c20066a8af2f46911573f1c236efed0107177e33c15974fccf3347e2df78e1): complete subsection reference.

- [medium_security](resources--cdn_loadbalancer--reference--group-012.md#canonical-3f5c02b42ea4132ac80a06e381dd159ee3cd410c7ac3de77c5272ff81246ad95): complete subsection reference.

<a id="canonical-124c9a3b1b0e9410f59e8be86caeb662782c6acbea140cdd0d3a7cda74d7de4a"></a>

## Next pages — origin_pool.use_tls.tls_config / 64629b4255c8 / 4

- [origin_pool.use_tls.tls_config.custom_security](resources--cdn_loadbalancer--reference--group-012.md#canonical-2d586a03ba26fbe6aeeae1d4c7c25fa52d01577791520227587fdbb1f6a5471a)
- [origin_pool.use_tls.tls_config.default_security](resources--cdn_loadbalancer--reference--group-012.md#canonical-9dd9bd81d6182c6f7a71237eae20b2a1e41ed92d1e2c459ee67d8516f211fb4b)
- [origin_pool.use_tls.tls_config.low_security](resources--cdn_loadbalancer--reference--group-012.md#canonical-c8c20066a8af2f46911573f1c236efed0107177e33c15974fccf3347e2df78e1)
- [origin_pool.use_tls.tls_config.medium_security](resources--cdn_loadbalancer--reference--group-012.md#canonical-3f5c02b42ea4132ac80a06e381dd159ee3cd410c7ac3de77c5272ff81246ad95)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-213416fcbc6044c18fab1b3315ab81a7eed3b03909873e5492aeff7c88e527e0)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-2d586a03ba26fbe6aeeae1d4c7c25fa52d01577791520227587fdbb1f6a5471a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-df53d4fe038a69041329cefd8471f5861f28089db56033959513f35fafd02dbb"></a>

## origin_pool.use_tls.tls_config.custom_security — origin_pool.use_tls.tls_config.custom_security / ee18aa8b6a91 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-603a6ee1a88486640208f36e01e4e00f68db830979076e4ad137090fd9373676)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-213416fcbc6044c18fab1b3315ab81a7eed3b03909873e5492aeff7c88e527e0)
- [origin_pool.use_tls.tls_config](resources--cdn_loadbalancer--reference--group-012.md#canonical-5fbf27b546aabdc4f2af2cbc3fd167a146805e1786c315777301990e3296462c)
- origin_pool.use_tls.tls_config.custom_security

<a id="canonical-6020ce137c66de74b3ee34c7d164d584c037e2a33de8bc9ad727b1fb2f8e3780"></a>

Type: `"object"`. single nested block, Optional.

Defines TLS protocol config including min/max versions and allowed ciphers.

Upstream description:

This defines TLS protocol config including min/max versions and allowed ciphers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cipher_suites")}
```

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

Terraform syntax:

```terraform
custom_security {
  # Configure direct properties listed below.
}
```

<a id="canonical-e2c19f496e5682439d1c489157824c410b14a2ba964ad956ddea5d40bd47eeb9"></a>

## Direct properties — origin_pool.use_tls.tls_config.custom_security / ee18aa8b6a91 / 3

<a id="canonical-03064ebc58192eb741aba0378ffe282e429245569b45d76a120996c411eb18b1"></a>

<a id="canonical-84684bc892819dee243b14c43bfc1bb4284f67569607d8865457a2c5a55f0607"></a>

## cipher_suites property — origin_pool.use_tls.tls_config.custom_security / ee18aa8b6a91 / 4

Type: `["list", "string"]`. Optional.

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

<a id="canonical-821bf8420d07b5787dd156085cfe05b2cf9facb576e74fec9bc7c5a5251c181c"></a>

<a id="canonical-89dcf65c2e25b5dedfc5920092b9a06e96e81fe8774175aa744fb41d30dbd8f5"></a>

## max_version property — origin_pool.use_tls.tls_config.custom_security / ee18aa8b6a91 / 5

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

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

<a id="canonical-5af5ecc6f0ab3b9f2bfd9c26675683282b2f82493dfc396bd8730307c622ade6"></a>

<a id="canonical-ecee803ef785070803d1fd5b4871b82b42d0cb0b1be585af04b6505260dc6a46"></a>

## min_version property — origin_pool.use_tls.tls_config.custom_security / ee18aa8b6a91 / 6

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

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

<a id="canonical-49feb7af8c899add50620c490dae3973d8e74df2d10823890de1d1c970076452"></a>

## Next pages — origin_pool.use_tls.tls_config.custom_security / ee18aa8b6a91 / 7

- [origin_pool.use_tls.tls_config](resources--cdn_loadbalancer--reference--group-012.md#canonical-5fbf27b546aabdc4f2af2cbc3fd167a146805e1786c315777301990e3296462c)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-9dd9bd81d6182c6f7a71237eae20b2a1e41ed92d1e2c459ee67d8516f211fb4b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-039a7027ad12c29452ec814c8b95061be30e84ed5960715cba53469c7d58ba7a"></a>

## origin_pool.use_tls.tls_config.default_security — origin_pool.use_tls.tls_config.default_security / 24be33bfada8 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-603a6ee1a88486640208f36e01e4e00f68db830979076e4ad137090fd9373676)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-213416fcbc6044c18fab1b3315ab81a7eed3b03909873e5492aeff7c88e527e0)
- [origin_pool.use_tls.tls_config](resources--cdn_loadbalancer--reference--group-012.md#canonical-5fbf27b546aabdc4f2af2cbc3fd167a146805e1786c315777301990e3296462c)
- origin_pool.use_tls.tls_config.default_security

<a id="canonical-1f4ac167ded080222032d5a574fbd202a90bbeeed6ed03487d321f558f91ec14"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default_security = {}
```

<a id="canonical-4649374241a066d7963a29b1e536fdead8442d09c06f981e3a00e68e2eb60be1"></a>

## Direct properties — origin_pool.use_tls.tls_config.default_security / 24be33bfada8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-57a2a0a161788d94064275b6617b8415cc0813f0918a9384730394904160293d"></a>

## Next pages — origin_pool.use_tls.tls_config.default_security / 24be33bfada8 / 4

- [origin_pool.use_tls.tls_config](resources--cdn_loadbalancer--reference--group-012.md#canonical-5fbf27b546aabdc4f2af2cbc3fd167a146805e1786c315777301990e3296462c)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-c8c20066a8af2f46911573f1c236efed0107177e33c15974fccf3347e2df78e1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ba1c479240b0d19efeb728a8510c16e18762ae9f5ffe6cc3c5eaf0d2093ba419"></a>

## origin_pool.use_tls.tls_config.low_security — origin_pool.use_tls.tls_config.low_security / 6c47c56e4be2 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-603a6ee1a88486640208f36e01e4e00f68db830979076e4ad137090fd9373676)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-213416fcbc6044c18fab1b3315ab81a7eed3b03909873e5492aeff7c88e527e0)
- [origin_pool.use_tls.tls_config](resources--cdn_loadbalancer--reference--group-012.md#canonical-5fbf27b546aabdc4f2af2cbc3fd167a146805e1786c315777301990e3296462c)
- origin_pool.use_tls.tls_config.low_security

<a id="canonical-dc04ee52c127a31c5e7a8bb01b58bbd30ae8b0285b11c437a883ac5cdc95c1e4"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
low_security = {}
```

<a id="canonical-b40f528c6922d40a3fe4493a3bb5d4ce8bba205c5e66b07fa018228f393207bc"></a>

## Direct properties — origin_pool.use_tls.tls_config.low_security / 6c47c56e4be2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e40c9bfe65792d22c5e5811dedd530b39f92f62afb8f9c8242c30732588b4576"></a>

## Next pages — origin_pool.use_tls.tls_config.low_security / 6c47c56e4be2 / 4

- [origin_pool.use_tls.tls_config](resources--cdn_loadbalancer--reference--group-012.md#canonical-5fbf27b546aabdc4f2af2cbc3fd167a146805e1786c315777301990e3296462c)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-3f5c02b42ea4132ac80a06e381dd159ee3cd410c7ac3de77c5272ff81246ad95"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3fe1eb6faf0db8e6226b2183174cef4ea840db8811c1573650f41fb56ffdb3bc"></a>

## origin_pool.use_tls.tls_config.medium_security — origin_pool.use_tls.tls_config.medium_security / 8368eb8b30e2 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-603a6ee1a88486640208f36e01e4e00f68db830979076e4ad137090fd9373676)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-213416fcbc6044c18fab1b3315ab81a7eed3b03909873e5492aeff7c88e527e0)
- [origin_pool.use_tls.tls_config](resources--cdn_loadbalancer--reference--group-012.md#canonical-5fbf27b546aabdc4f2af2cbc3fd167a146805e1786c315777301990e3296462c)
- origin_pool.use_tls.tls_config.medium_security

<a id="canonical-0e08a425c520aa4d7797a1c330bd8ac603b95dce4e51cf44c1866680d5f59944"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
medium_security = {}
```

<a id="canonical-ffc95514d5f8f4bae034fcf45664be1b81510718f0e25bd7cc9391842bbe479a"></a>

## Direct properties — origin_pool.use_tls.tls_config.medium_security / 8368eb8b30e2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-93c6ce433f99fe44003a955dac0d448a441a2a057bfb9d4111d7a74d87f9f866"></a>

## Next pages — origin_pool.use_tls.tls_config.medium_security / 8368eb8b30e2 / 4

- [origin_pool.use_tls.tls_config](resources--cdn_loadbalancer--reference--group-012.md#canonical-5fbf27b546aabdc4f2af2cbc3fd167a146805e1786c315777301990e3296462c)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-dbcf2cdbe2fd6076d6aa3893cf10d388b09145424d5ea691535e20b2fddac59d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c1f921f10e369e9b95b7b6a2712f325ca61cd1b8052824b779abaac8339031b1"></a>

## origin_pool.use_tls.use_host_header_as_sni — origin_pool.use_tls.use_host_header_as_sni / 21fad6838fde / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-603a6ee1a88486640208f36e01e4e00f68db830979076e4ad137090fd9373676)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-213416fcbc6044c18fab1b3315ab81a7eed3b03909873e5492aeff7c88e527e0)
- origin_pool.use_tls.use_host_header_as_sni

<a id="canonical-ff4e4b111ddf26607ff2390a1e16a4014912e1150d940c77917dadd081dcd61b"></a>

Type: `["object", {}]`. Optional, Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

Terraform syntax:

```terraform
use_host_header_as_sni = {}
```

<a id="canonical-4d60315cf30d3c9eb51bef8d691d380f383fdc09b8c9cdc63d3499f361ad8cb8"></a>

## Direct properties — origin_pool.use_tls.use_host_header_as_sni / 21fad6838fde / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a32fe48b3c011b26a4d07c04ad88effa75046c70cb2372622e6993a98778173f"></a>

## Next pages — origin_pool.use_tls.use_host_header_as_sni / 21fad6838fde / 4

- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-213416fcbc6044c18fab1b3315ab81a7eed3b03909873e5492aeff7c88e527e0)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-10135e1d0e39aea2354081a9383060b17612baabf9b8d1e55f79954afb3713e8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-67c629de1f5fa9553e15a9040f272e9c7a3f908881525423445d564c54ccebfd"></a>

## origin_pool.use_tls.use_mtls — origin_pool.use_tls.use_mtls / 6b7ab7f85999 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-603a6ee1a88486640208f36e01e4e00f68db830979076e4ad137090fd9373676)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-213416fcbc6044c18fab1b3315ab81a7eed3b03909873e5492aeff7c88e527e0)
- origin_pool.use_tls.use_mtls

<a id="canonical-3537c556440ab569af08c0543f0eefe3a3c8cc117a8dcc599680ea9fcd2dddb3"></a>

Type: `"object"`. single nested block, Optional.

MTLS Certificate. MTLS Client Certificate.

Upstream description:

MTLS Client Certificate.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("tls_certificates")}
```

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

Terraform syntax:

```terraform
use_mtls {
  # Configure direct properties listed below.
}
```

<a id="canonical-b5439e182e51e534350c56792f43a271633eb138bb83ac8a91b1c27314286a2f"></a>

## Direct properties — origin_pool.use_tls.use_mtls / 6b7ab7f85999 / 3

- [tls_certificates](resources--cdn_loadbalancer--reference--group-012.md#canonical-82ac64a572854798e6b0a645aceb2e852c3086186e6c948b917fa8ba325e00c8): complete subsection reference.

<a id="canonical-25ffbe2ec3241bb7b8fa9b552f450ec7bda205f7682d7afc6acef203c7df97bd"></a>

## Next pages — origin_pool.use_tls.use_mtls / 6b7ab7f85999 / 4

- [origin_pool.use_tls.use_mtls.tls_certificates](resources--cdn_loadbalancer--reference--group-012.md#canonical-82ac64a572854798e6b0a645aceb2e852c3086186e6c948b917fa8ba325e00c8)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-213416fcbc6044c18fab1b3315ab81a7eed3b03909873e5492aeff7c88e527e0)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-82ac64a572854798e6b0a645aceb2e852c3086186e6c948b917fa8ba325e00c8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-62f66f9659c9549876a8435add4b395b1e70db882676882bc9257398c73d44e6"></a>

## origin_pool.use_tls.use_mtls.tls_certificates — origin_pool.use_tls.use_mtls.tls_certificates / 359010854a63 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-603a6ee1a88486640208f36e01e4e00f68db830979076e4ad137090fd9373676)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-213416fcbc6044c18fab1b3315ab81a7eed3b03909873e5492aeff7c88e527e0)
- [origin_pool.use_tls.use_mtls](resources--cdn_loadbalancer--reference--group-012.md#canonical-10135e1d0e39aea2354081a9383060b17612baabf9b8d1e55f79954afb3713e8)
- origin_pool.use_tls.use_mtls.tls_certificates

<a id="canonical-7930ae90aa9c964857e1302b89d9e7f14693010dc0b1037899d6bd030eb7b747"></a>

Type: `"object"`. list nested block, Optional.

MTLS Client Certificate. MTLS Client Certificate.

Upstream description:

MTLS Client Certificate.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("certificate_url"),
  validators.ConflictingListObjectAttributes("custom_hash_algorithms",
    "disable_ocsp_stapling"),
  validators.ConflictingListObjectAttributes("custom_hash_algorithms",
    "use_system_defaults"),
  validators.ConflictingListObjectAttributes("disable_ocsp_stapling",
    "use_system_defaults")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
tls_certificates {
  # Configure direct properties listed below.
}
```

<a id="canonical-2343d91de10b2e6b4cd72146a4938910263c25684114024e54f4118ab7537b09"></a>

## Direct properties — origin_pool.use_tls.use_mtls.tls_certificates / 359010854a63 / 3

<a id="canonical-dd6ecc12b3d1b26d2a1e69fdd106296b208bfebd32d65a3db8b39f71cec7fe52"></a>

<a id="canonical-ddd38461d72f01980f547a5c5e3f4756e3e035b4853583565fc71711d0b9d0f7"></a>

## certificate_url property — origin_pool.use_tls.use_mtls.tls_certificates / 359010854a63 / 4

Type: `"string"`. Optional.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Upstream description:

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

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

- [custom_hash_algorithms](resources--cdn_loadbalancer--reference--group-012.md#canonical-86d386be06ca14594ed95de857a351f710de9d194db03280a02e2f1cac761c96): complete subsection reference.

<a id="canonical-82084ec3fb9f5b968f23c0f7eb2757556f6650f838a056873fd88d23c0499bf0"></a>

<a id="canonical-948585d42c4b1392ae347e90c11518c0878ccd7b62cfb1fb812bc649d5f22985"></a>

## description_spec property — origin_pool.use_tls.use_mtls.tls_certificates / 359010854a63 / 5

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--cdn_loadbalancer--reference--group-012.md#canonical-0c0069b0631af29ee549d4a8f6c48b185cc765b09bec83755134a0fadd765342): complete subsection reference.

- [private_key](resources--cdn_loadbalancer--reference--group-012.md#canonical-967320c45c30b7c8351ef143f6f2b61f70b60ab46dcf659e6b7cea8baa1a2235): complete subsection reference.

- [use_system_defaults](resources--cdn_loadbalancer--reference--group-012.md#canonical-399419a9a1d6bd7c2360b417d22db396de0cf3ce77b6ba6b35e253a8529038ad): complete subsection reference.

<a id="canonical-838a07464aad2868316d93c3386fc78b8fbb6404f02b48bc820cb1fbee2278d6"></a>

## Next pages — origin_pool.use_tls.use_mtls.tls_certificates / 359010854a63 / 6

- [origin_pool.use_tls.use_mtls.tls_certificates.custom_hash_algorithms](resources--cdn_loadbalancer--reference--group-012.md#canonical-86d386be06ca14594ed95de857a351f710de9d194db03280a02e2f1cac761c96)
- [origin_pool.use_tls.use_mtls.tls_certificates.disable_ocsp_stapling](resources--cdn_loadbalancer--reference--group-012.md#canonical-0c0069b0631af29ee549d4a8f6c48b185cc765b09bec83755134a0fadd765342)
- [origin_pool.use_tls.use_mtls.tls_certificates.private_key](resources--cdn_loadbalancer--reference--group-012.md#canonical-967320c45c30b7c8351ef143f6f2b61f70b60ab46dcf659e6b7cea8baa1a2235)
- [origin_pool.use_tls.use_mtls.tls_certificates.use_system_defaults](resources--cdn_loadbalancer--reference--group-012.md#canonical-399419a9a1d6bd7c2360b417d22db396de0cf3ce77b6ba6b35e253a8529038ad)
- [origin_pool.use_tls.use_mtls](resources--cdn_loadbalancer--reference--group-012.md#canonical-10135e1d0e39aea2354081a9383060b17612baabf9b8d1e55f79954afb3713e8)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-86d386be06ca14594ed95de857a351f710de9d194db03280a02e2f1cac761c96"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b9e699d74898546d181ceb495947412afb35bd9475a639d3457056dcf0ca3590"></a>

## origin_pool.use_tls.use_mtls.tls_certificates.custom_hash_algorithms — origin_pool.use_tls.use_mtls.tls_certificates.custom_hash_algorithms / 0066da8e2a06 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-603a6ee1a88486640208f36e01e4e00f68db830979076e4ad137090fd9373676)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-213416fcbc6044c18fab1b3315ab81a7eed3b03909873e5492aeff7c88e527e0)
- [origin_pool.use_tls.use_mtls](resources--cdn_loadbalancer--reference--group-012.md#canonical-10135e1d0e39aea2354081a9383060b17612baabf9b8d1e55f79954afb3713e8)
- [origin_pool.use_tls.use_mtls.tls_certificates](resources--cdn_loadbalancer--reference--group-012.md#canonical-82ac64a572854798e6b0a645aceb2e852c3086186e6c948b917fa8ba325e00c8)
- origin_pool.use_tls.use_mtls.tls_certificates.custom_hash_algorithms

<a id="canonical-fb8ff43e7afaa8625dc8d02714ff27fbc643eb98d564b355b69e5304ab000cdb"></a>

Type: `"object"`. single nested block, Optional.

Specifies the hash algorithms to be used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("hash_algorithms")}
```

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

Terraform syntax:

```terraform
custom_hash_algorithms {
  # Configure direct properties listed below.
}
```

<a id="canonical-d31335948bd29ca47d19949cebeeb7f583bdf6f52ab3ccf2bed2810260abac28"></a>

## Direct properties — origin_pool.use_tls.use_mtls.tls_certificates.custom_hash_algorithms / 0066da8e2a06 / 3

<a id="canonical-84a25a31bed9a5b3eb6003f6890f008e0656a09e54c44175f3f89d853eba6773"></a>

<a id="canonical-55bc07dac1ff42758438a2e26dc0406eaff6ed821c76ef1166ed676aa5e03b1d"></a>

## hash_algorithms property — origin_pool.use_tls.use_mtls.tls_certificates.custom_hash_algorithms / 0066da8e2a06 / 4

Type: `["list", "string"]`. Optional.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

Upstream description:

Ordered list of hash algorithms to be used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 4),
}
```

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

<a id="canonical-24ba4c539e0447b071424d3688162c943dfe0ab6a88d633be8557c9a645383b9"></a>

## Next pages — origin_pool.use_tls.use_mtls.tls_certificates.custom_hash_algorithms / 0066da8e2a06 / 5

- [origin_pool.use_tls.use_mtls.tls_certificates](resources--cdn_loadbalancer--reference--group-012.md#canonical-82ac64a572854798e6b0a645aceb2e852c3086186e6c948b917fa8ba325e00c8)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-0c0069b0631af29ee549d4a8f6c48b185cc765b09bec83755134a0fadd765342"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fb29273d0cd3d88972e9504ad74c23755d127c10c4e5766598d3a86263a64a2d"></a>

## origin_pool.use_tls.use_mtls.tls_certificates.disable_ocsp_stapling — origin_pool.use_tls.use_mtls.tls_certificates.disable_ocsp_stapling / 0e8c662eee68 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-603a6ee1a88486640208f36e01e4e00f68db830979076e4ad137090fd9373676)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-213416fcbc6044c18fab1b3315ab81a7eed3b03909873e5492aeff7c88e527e0)
- [origin_pool.use_tls.use_mtls](resources--cdn_loadbalancer--reference--group-012.md#canonical-10135e1d0e39aea2354081a9383060b17612baabf9b8d1e55f79954afb3713e8)
- [origin_pool.use_tls.use_mtls.tls_certificates](resources--cdn_loadbalancer--reference--group-012.md#canonical-82ac64a572854798e6b0a645aceb2e852c3086186e6c948b917fa8ba325e00c8)
- origin_pool.use_tls.use_mtls.tls_certificates.disable_ocsp_stapling

<a id="canonical-40e6be2a80dbd5ee3697e3c95440110f4155f0736619b1c1127b0124a0ab27b3"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
disable_ocsp_stapling = {}
```

<a id="canonical-97f1d771af6a740b36f222254ea8444fe0de3385ce50ad1a36ea3bd6d4feb086"></a>

## Direct properties — origin_pool.use_tls.use_mtls.tls_certificates.disable_ocsp_stapling / 0e8c662eee68 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-65fe85f5aac750bbe4ff8feed6696a51f176a4e208b5a7173cfb2d8c96d628d6"></a>

## Next pages — origin_pool.use_tls.use_mtls.tls_certificates.disable_ocsp_stapling / 0e8c662eee68 / 4

- [origin_pool.use_tls.use_mtls.tls_certificates](resources--cdn_loadbalancer--reference--group-012.md#canonical-82ac64a572854798e6b0a645aceb2e852c3086186e6c948b917fa8ba325e00c8)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-967320c45c30b7c8351ef143f6f2b61f70b60ab46dcf659e6b7cea8baa1a2235"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-742983d1a09fc2236cc7bec0e7abe535052e6a0e5bd5d3f5faba5d3b5837416b"></a>

## origin_pool.use_tls.use_mtls.tls_certificates.private_key — origin_pool.use_tls.use_mtls.tls_certificates.private_key / 8c15ccfd6ef2 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-603a6ee1a88486640208f36e01e4e00f68db830979076e4ad137090fd9373676)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-213416fcbc6044c18fab1b3315ab81a7eed3b03909873e5492aeff7c88e527e0)
- [origin_pool.use_tls.use_mtls](resources--cdn_loadbalancer--reference--group-012.md#canonical-10135e1d0e39aea2354081a9383060b17612baabf9b8d1e55f79954afb3713e8)
- [origin_pool.use_tls.use_mtls.tls_certificates](resources--cdn_loadbalancer--reference--group-012.md#canonical-82ac64a572854798e6b0a645aceb2e852c3086186e6c948b917fa8ba325e00c8)
- origin_pool.use_tls.use_mtls.tls_certificates.private_key

<a id="canonical-ff748ed014ce10843f29eef15afe221d5bd79ac727691bf151e361928abe8706"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
private_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-3c88f74ebb1bebf0a9968e4af502901111686b0d21eaa4c0b6047e597564fc21"></a>

## Direct properties — origin_pool.use_tls.use_mtls.tls_certificates.private_key / 8c15ccfd6ef2 / 3

- [blindfold_secret_info](resources--cdn_loadbalancer--reference--group-012.md#canonical-e6ac1f06e943e7fbe4387e73630604a5e698766637a52bbcae0904796e1ea952): complete subsection reference.

- [clear_secret_info](resources--cdn_loadbalancer--reference--group-012.md#canonical-c446bccdb8b2af6550e24b22f950c364d29450c1b3f71d918afa660f5a9806dc): complete subsection reference.

<a id="canonical-421662e2e9c8e6986428a93980ed9b2990ecadb7e443183e53bf07ebbb98e7b9"></a>

## Next pages — origin_pool.use_tls.use_mtls.tls_certificates.private_key / 8c15ccfd6ef2 / 4

- [origin_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info](resources--cdn_loadbalancer--reference--group-012.md#canonical-e6ac1f06e943e7fbe4387e73630604a5e698766637a52bbcae0904796e1ea952)
- [origin_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info](resources--cdn_loadbalancer--reference--group-012.md#canonical-c446bccdb8b2af6550e24b22f950c364d29450c1b3f71d918afa660f5a9806dc)
- [origin_pool.use_tls.use_mtls.tls_certificates](resources--cdn_loadbalancer--reference--group-012.md#canonical-82ac64a572854798e6b0a645aceb2e852c3086186e6c948b917fa8ba325e00c8)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-e6ac1f06e943e7fbe4387e73630604a5e698766637a52bbcae0904796e1ea952"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2ec3e752befcf376f515c425b41cb100ff870665afdb164c4f9df666cdc5f88f"></a>

## origin_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info — origin_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info / 186cbe96fcb2 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-603a6ee1a88486640208f36e01e4e00f68db830979076e4ad137090fd9373676)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-213416fcbc6044c18fab1b3315ab81a7eed3b03909873e5492aeff7c88e527e0)
- [origin_pool.use_tls.use_mtls](resources--cdn_loadbalancer--reference--group-012.md#canonical-10135e1d0e39aea2354081a9383060b17612baabf9b8d1e55f79954afb3713e8)
- [origin_pool.use_tls.use_mtls.tls_certificates](resources--cdn_loadbalancer--reference--group-012.md#canonical-82ac64a572854798e6b0a645aceb2e852c3086186e6c948b917fa8ba325e00c8)
- [origin_pool.use_tls.use_mtls.tls_certificates.private_key](resources--cdn_loadbalancer--reference--group-012.md#canonical-967320c45c30b7c8351ef143f6f2b61f70b60ab46dcf659e6b7cea8baa1a2235)
- origin_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-bd773461ae1808cab7da434ab3303cc635306c6d5fbcace4ee0d8664ecccbea0"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
```

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

Terraform syntax:

```terraform
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-56bbe65eabf42a01bb950e7c76842a56e3a12cb0126a5dbe3e9dfeb3e9c9f45f"></a>

## Direct properties — origin_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info / 186cbe96fcb2 / 3

<a id="canonical-02bc3f4830eedfddf9b669a2d02d204200be19394f7b29cf528abd683d53130e"></a>

<a id="canonical-219f7bbcaf9bc06e9e91513808b1a88a603f75f63c8921ab5672d7163592eee5"></a>

## decryption_provider property — origin_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info / 186cbe96fcb2 / 4

Type: `"string"`. Optional.

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

<a id="canonical-066ba6675343868a75bcb965280ee4ffe9d335d13e7e88bc743c65c9e18725cf"></a>

<a id="canonical-6e5dea57b7124dccd9ce4d9b10f7a25cfb254403cdd810dadd54b6193cadf7fa"></a>

## location property — origin_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info / 186cbe96fcb2 / 5

Type: `"string"`. Optional, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

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

<a id="canonical-b14254cb3474a543293f49257f0fe5074ad86c12254ca76a27ea2282cfd2bdcd"></a>

<a id="canonical-9aa4fb99c066b519a214286ef2d031fc5b353c1e219ce271c97ec691597f677c"></a>

## store_provider property — origin_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info / 186cbe96fcb2 / 6

Type: `"string"`. Optional.

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

<a id="canonical-72daf9b861421a83d3d8c03f565f7e0e7a02b1181c86e40c8ad7044e5e89d19d"></a>

## Next pages — origin_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info / 186cbe96fcb2 / 7

- [origin_pool.use_tls.use_mtls.tls_certificates.private_key](resources--cdn_loadbalancer--reference--group-012.md#canonical-967320c45c30b7c8351ef143f6f2b61f70b60ab46dcf659e6b7cea8baa1a2235)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-c446bccdb8b2af6550e24b22f950c364d29450c1b3f71d918afa660f5a9806dc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1fcae6dc791ba6f404507767ebe5ba25cf0da9bdde7ecb8cc7bf9aec8b78cb6b"></a>

## origin_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info — origin_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info / db3d3df81e41 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-603a6ee1a88486640208f36e01e4e00f68db830979076e4ad137090fd9373676)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-213416fcbc6044c18fab1b3315ab81a7eed3b03909873e5492aeff7c88e527e0)
- [origin_pool.use_tls.use_mtls](resources--cdn_loadbalancer--reference--group-012.md#canonical-10135e1d0e39aea2354081a9383060b17612baabf9b8d1e55f79954afb3713e8)
- [origin_pool.use_tls.use_mtls.tls_certificates](resources--cdn_loadbalancer--reference--group-012.md#canonical-82ac64a572854798e6b0a645aceb2e852c3086186e6c948b917fa8ba325e00c8)
- [origin_pool.use_tls.use_mtls.tls_certificates.private_key](resources--cdn_loadbalancer--reference--group-012.md#canonical-967320c45c30b7c8351ef143f6f2b61f70b60ab46dcf659e6b7cea8baa1a2235)
- origin_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info

<a id="canonical-2bb41600a88fefa906027cce0a2cc8f407e4ae150d55c08bcfc9dfe270bf8f28"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
```

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

Terraform syntax:

```terraform
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-f7ae8ef07093cc53e6c03b69b83b2cae9a8a7dd26abf4bf7e9ea2a6cedec6348"></a>

## Direct properties — origin_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info / db3d3df81e41 / 3

<a id="canonical-8e56b3b115f37135282db73b5c526b44c25d567478e7f17af966d6e215c918e5"></a>

<a id="canonical-448fabc65fe2a8befb9c572eebbd26591239f0a7ca4af0d8396bcc46bc006049"></a>

## provider_ref property — origin_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info / db3d3df81e41 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-4a95d18680be2eedc0a4716192f9c611ce65950b9fe3ddb317138638e79e8bc9"></a>

<a id="canonical-aa352d9f623de26ebca02ff9ad340488ccd84030d60912cc3a7ba71fd847ee0b"></a>

## url property — origin_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info / db3d3df81e41 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

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

<a id="canonical-4b7d3b0ce02aa127eaff034d1937b242458f9a7d1b897453e47ed5c660c0e764"></a>

## Next pages — origin_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info / db3d3df81e41 / 6

- [origin_pool.use_tls.use_mtls.tls_certificates.private_key](resources--cdn_loadbalancer--reference--group-012.md#canonical-967320c45c30b7c8351ef143f6f2b61f70b60ab46dcf659e6b7cea8baa1a2235)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-399419a9a1d6bd7c2360b417d22db396de0cf3ce77b6ba6b35e253a8529038ad"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-726df1954b4d59243cca609b0454bd00d4d99e7ac8b12425a57bfb476e7afeff"></a>

## origin_pool.use_tls.use_mtls.tls_certificates.use_system_defaults — origin_pool.use_tls.use_mtls.tls_certificates.use_system_defaults / b8b567cb35d8 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-603a6ee1a88486640208f36e01e4e00f68db830979076e4ad137090fd9373676)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-213416fcbc6044c18fab1b3315ab81a7eed3b03909873e5492aeff7c88e527e0)
- [origin_pool.use_tls.use_mtls](resources--cdn_loadbalancer--reference--group-012.md#canonical-10135e1d0e39aea2354081a9383060b17612baabf9b8d1e55f79954afb3713e8)
- [origin_pool.use_tls.use_mtls.tls_certificates](resources--cdn_loadbalancer--reference--group-012.md#canonical-82ac64a572854798e6b0a645aceb2e852c3086186e6c948b917fa8ba325e00c8)
- origin_pool.use_tls.use_mtls.tls_certificates.use_system_defaults

<a id="canonical-65c0bf023d5965cd03911edc6abbbd3f14cf20539df6f88c13157f70484a6e86"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
use_system_defaults = {}
```

<a id="canonical-88346ad55bf0bd7bc46d9754984f510c164d98293009a511f8501c3fa821eb58"></a>

## Direct properties — origin_pool.use_tls.use_mtls.tls_certificates.use_system_defaults / b8b567cb35d8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ea82978d4b86b2fe30dd206aa122f26b708cb698d81ca45e26590f0abf97dde8"></a>

## Next pages — origin_pool.use_tls.use_mtls.tls_certificates.use_system_defaults / b8b567cb35d8 / 4

- [origin_pool.use_tls.use_mtls.tls_certificates](resources--cdn_loadbalancer--reference--group-012.md#canonical-82ac64a572854798e6b0a645aceb2e852c3086186e6c948b917fa8ba325e00c8)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-8d22410b61f69f6fcc934ed0d400e6cab770df64afe0a892a58ce3b15297e50d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-77d80e7d97f42018ae94fc669050dfc3e0fdf7c990f40984eda64166b2c13baf"></a>

## origin_pool.use_tls.use_mtls_obj — origin_pool.use_tls.use_mtls_obj / 1b6e91babbe5 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-603a6ee1a88486640208f36e01e4e00f68db830979076e4ad137090fd9373676)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-213416fcbc6044c18fab1b3315ab81a7eed3b03909873e5492aeff7c88e527e0)
- origin_pool.use_tls.use_mtls_obj

<a id="canonical-64f4b9e604086c5907c77f38c61672b8bdeb9dbd2b58a7fa0b1b2ab0f67bd8e7"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
```

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

Terraform syntax:

```terraform
use_mtls_obj {
  # Configure direct properties listed below.
}
```

<a id="canonical-40cd418f42941f21aaa90dced65fd4edc4e6e32da150bbf299f9bfc283730c7c"></a>

## Direct properties — origin_pool.use_tls.use_mtls_obj / 1b6e91babbe5 / 3

<a id="canonical-5f19b5c799ecb44947d11aa0e8245f646c71d95d53f401031b9e2073aa712cae"></a>

<a id="canonical-9ab602e9a1dcc912be1d7200eca548c43c415c95f58ed58dbb499daa4599f8b5"></a>

## name property — origin_pool.use_tls.use_mtls_obj / 1b6e91babbe5 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-1ee773cc1597ec0af8cce0a6baf67681cb9cde68aaf539e599bf3faeab7008df"></a>

<a id="canonical-d8ae8935c24d843548c2f1cd2ae51fec37b218a20796918065575bac414b85de"></a>

## namespace property — origin_pool.use_tls.use_mtls_obj / 1b6e91babbe5 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-61f6a170b5959ecf934794493d04561fbdfaab1fcc32560e2343922520330f69"></a>

<a id="canonical-a9225d6fdd59c247cc981fc6b428410e7abcc049ccccf83e23fb09c199478b60"></a>

## tenant property — origin_pool.use_tls.use_mtls_obj / 1b6e91babbe5 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-feda9eca397d84518b35d8bb58c5a6a7c0d47ecac80eb76b1a07761f7d7e243d"></a>

## Next pages — origin_pool.use_tls.use_mtls_obj / 1b6e91babbe5 / 7

- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-213416fcbc6044c18fab1b3315ab81a7eed3b03909873e5492aeff7c88e527e0)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-1087b1f966ea83fbe26326ef2fa2e8a742d53896d3dd13c74807c9234554e48b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f4b66e97286b13726a7447a175721cb8633bcd8805f3c62f2fccd8015ded760e"></a>

## origin_pool.use_tls.use_server_verification — origin_pool.use_tls.use_server_verification / 809582ff1870 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-603a6ee1a88486640208f36e01e4e00f68db830979076e4ad137090fd9373676)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-213416fcbc6044c18fab1b3315ab81a7eed3b03909873e5492aeff7c88e527e0)
- origin_pool.use_tls.use_server_verification

<a id="canonical-a870f7a7052742fc52123c389b5402c8ce7058d9129343c6bf39e7e91543b793"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for use server verification.

Upstream description:

Upstream TLS Validation Context.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("trusted_ca",
    "trusted_ca_url")}
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
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]"
}
```

Terraform syntax:

```terraform
use_server_verification {
  # Configure direct properties listed below.
}
```

<a id="canonical-b283604b646baf5ff395e3035289579737a87ed017b409b6802b9b3dcf0fa107"></a>

## Direct properties — origin_pool.use_tls.use_server_verification / 809582ff1870 / 3

- [trusted_ca](resources--cdn_loadbalancer--reference--group-012.md#canonical-c29cb27a1f14d6d864dc25a9a2f0cf6f90aaef434b736c9c0fc35661c929be70): complete subsection reference.

<a id="canonical-c67c3b14da3a9dba6134911b62af33fb4609dc35afa21193eeb22eaf34c32fcb"></a>

<a id="canonical-7cef43e1bfc08a25264d1271cb544e1a855bee2208d6642be7af99197640d622"></a>

## trusted_ca_url property — origin_pool.use_tls.use_server_verification / 809582ff1870 / 4

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Origin Pool for
verification of server's certificate.

Upstream description:

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Origin Pool for
verification of server's certificate.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

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

<a id="canonical-7e6c75813d25f21456857b7517bb7ed708443078421a9d0f410739111305e11d"></a>

## Next pages — origin_pool.use_tls.use_server_verification / 809582ff1870 / 5

- [origin_pool.use_tls.use_server_verification.trusted_ca](resources--cdn_loadbalancer--reference--group-012.md#canonical-c29cb27a1f14d6d864dc25a9a2f0cf6f90aaef434b736c9c0fc35661c929be70)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-213416fcbc6044c18fab1b3315ab81a7eed3b03909873e5492aeff7c88e527e0)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-c29cb27a1f14d6d864dc25a9a2f0cf6f90aaef434b736c9c0fc35661c929be70"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a498b0adb990258b90f335cbbd3a0ed0396182e79fcef0a1c61c0d8d3b530cc7"></a>

## origin_pool.use_tls.use_server_verification.trusted_ca — origin_pool.use_tls.use_server_verification.trusted_ca / 1579f45680e9 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-603a6ee1a88486640208f36e01e4e00f68db830979076e4ad137090fd9373676)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-213416fcbc6044c18fab1b3315ab81a7eed3b03909873e5492aeff7c88e527e0)
- [origin_pool.use_tls.use_server_verification](resources--cdn_loadbalancer--reference--group-012.md#canonical-1087b1f966ea83fbe26326ef2fa2e8a742d53896d3dd13c74807c9234554e48b)
- origin_pool.use_tls.use_server_verification.trusted_ca

<a id="canonical-b3bca070b52bdb24d5042faffc5f38d127b5a8c846f8221f164800a86b129d52"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
```

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

Terraform syntax:

```terraform
trusted_ca {
  # Configure direct properties listed below.
}
```

<a id="canonical-9bc2bcec3ab09882e3350808caf12fe33ba64458094271b442459c4e4dd98ff9"></a>

## Direct properties — origin_pool.use_tls.use_server_verification.trusted_ca / 1579f45680e9 / 3

<a id="canonical-0d0085309e8019af8f783fa1496776edd6724f7674d2307171151d650a123060"></a>

<a id="canonical-a0d5da7c5167ffe14b23c07ee9ec247f19cc5c035e9d93abf836a2ff974bb5ad"></a>

## name property — origin_pool.use_tls.use_server_verification.trusted_ca / 1579f45680e9 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-b77c9e5eab6bd938f55072073ba70a1a9d41d2605d48284a6d99abc15607f5b1"></a>

<a id="canonical-64fedb0aba9cb08846b5946b5222d43a9cbe03c08dd4ffb588090d9cbb83e946"></a>

## namespace property — origin_pool.use_tls.use_server_verification.trusted_ca / 1579f45680e9 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-a7f08b4496303cffb707835e52f68b9b370a73b89aa114e49dcc2d70d6cf5a86"></a>

<a id="canonical-d682589ed8e72bc78b4ff7ce0719b439da42813458357fc903cd245df04819b7"></a>

## tenant property — origin_pool.use_tls.use_server_verification.trusted_ca / 1579f45680e9 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-9ca03d1598af79c5963c6e4bd41f34d92d9887015bea0b2a42017f0949bba78e"></a>

## Next pages — origin_pool.use_tls.use_server_verification.trusted_ca / 1579f45680e9 / 7

- [origin_pool.use_tls.use_server_verification](resources--cdn_loadbalancer--reference--group-012.md#canonical-1087b1f966ea83fbe26326ef2fa2e8a742d53896d3dd13c74807c9234554e48b)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-deec9d5dd45c2417e305be03115967af8c43414a8f7d65c9a8c6706dba4427cd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-63423db7c403fc3c0bf673698279c1305e84bd9b2eb79277db973f7dff4ca3d3"></a>

## origin_pool.use_tls.volterra_trusted_ca — origin_pool.use_tls.volterra_trusted_ca / e65761e1cf94 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-603a6ee1a88486640208f36e01e4e00f68db830979076e4ad137090fd9373676)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-213416fcbc6044c18fab1b3315ab81a7eed3b03909873e5492aeff7c88e527e0)
- origin_pool.use_tls.volterra_trusted_ca

<a id="canonical-957b240914792340572af5057323a33956f1e1758f24a6f8b176dac0fa443a4f"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for volterra trusted ca. Defaults to \`map\[\]\`. Server applies default
when omitted.

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

Terraform syntax:

```terraform
volterra_trusted_ca = {}
```

<a id="canonical-86a9d39f3165631011a01a8c176f3e8f132cdd70cd80948266f430ec6896ee4a"></a>

## Direct properties — origin_pool.use_tls.volterra_trusted_ca / e65761e1cf94 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b544901759160e85089ad5d587762b420af59af678cb35d8d77b67d789918f42"></a>

## Next pages — origin_pool.use_tls.volterra_trusted_ca / e65761e1cf94 / 4

- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-213416fcbc6044c18fab1b3315ab81a7eed3b03909873e5492aeff7c88e527e0)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-6237c1513ed4eb47b747c88f4987c86621425202c792c064ab4bfdfacbad388a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e75cefdafac2c1ff1070cf0294f7464d77353890e0e9399266912e404c4c80f7"></a>

## other_settings — other_settings / 279ed62fa2f8 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- other_settings

<a id="canonical-6af79f052160012a636489a8afd3b118b96a007fa1bead9e720f91c75cbda072"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for other settings.

Upstream description:

Other Settings.

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

Terraform syntax:

```terraform
other_settings {
  # Configure direct properties listed below.
}
```

<a id="canonical-d7e965600b9760ef0cc7fa7c12be30c9f618a4caaf395b9436e28e3a41ef2a29"></a>

## Direct properties — other_settings / 279ed62fa2f8 / 3

<a id="canonical-c14aa084bc1e9c7d71a3b32bf702714a3cebead5b8f97583098985618b3be385"></a>

<a id="canonical-5d2dd1e186061088aded12eb1ae261ca8c8a8e6c86af2fbd8a7977b8a588a86b"></a>

## add_location property — other_settings / 279ed62fa2f8 / 4

Type: `"bool"`. Optional.

Add Location. X-example: true Appends header x-F5 Distributed Cloud-location = &lt;RE-site-name&gt;
in responses.

Upstream description:

X-example: true Appends header x-F5 Distributed Cloud-location = &lt;RE-site-name&gt; in responses.

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

- [header_options](resources--cdn_loadbalancer--reference--group-012.md#canonical-2d180d161dffb067f7db1e0720d11e345c4d3005e1d9ea84233978dd5f6e4cdb): complete subsection reference.

- [logging_options](resources--cdn_loadbalancer--reference--group-012.md#canonical-9fd681a2d1eb0f615230fa310f857bb9bca5668fad19c13074eec63c56995bdc): complete subsection reference.

<a id="canonical-19ec48ed386904a74ace1c75a26c89d451a6d3c5f6dbba49187943b0f082c6ec"></a>

## Next pages — other_settings / 279ed62fa2f8 / 5

- [other_settings.header_options](resources--cdn_loadbalancer--reference--group-012.md#canonical-2d180d161dffb067f7db1e0720d11e345c4d3005e1d9ea84233978dd5f6e4cdb)
- [other_settings.logging_options](resources--cdn_loadbalancer--reference--group-012.md#canonical-9fd681a2d1eb0f615230fa310f857bb9bca5668fad19c13074eec63c56995bdc)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-2d180d161dffb067f7db1e0720d11e345c4d3005e1d9ea84233978dd5f6e4cdb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-af7909b11613bdc5091c65fdd6d7baec4a0a83c3b547e88869b4fbd8deb20dc8"></a>

## other_settings.header_options — other_settings.header_options / 5d4e69652685 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [other_settings](resources--cdn_loadbalancer--reference--group-012.md#canonical-6237c1513ed4eb47b747c88f4987c86621425202c792c064ab4bfdfacbad388a)
- other_settings.header_options

<a id="canonical-181a11c86b4513c3b8b323bca46678f8901cb62ff8b380e5ffe6f4d86ea45f0f"></a>

Type: `"object"`. single nested block, Optional.

Defines various OPTIONS related to request/response headers.

Upstream description:

This defines various OPTIONS related to request/response headers.

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

Terraform syntax:

```terraform
header_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-f2eac06389c5a4e9095cf009a7b1226b6e0b34dd62d2166e495c2ed3ef703b3a"></a>

## Direct properties — other_settings.header_options / 5d4e69652685 / 3

- [request_headers_to_add](resources--cdn_loadbalancer--reference--group-012.md#canonical-d4798b10c5cf6cce2b217dade31c80cc63e30f112061600bc31514aef034f162): complete subsection reference.

<a id="canonical-f7b2e7f395c47587ac7454fec15ed3238ccc24e4f5ff881f34f551c9d52c3e72"></a>

<a id="canonical-5b5ed22ac6c24bfaebb9909aa2a4a0b89d95d529543161140a8bc5a4b10fa59d"></a>

## request_headers_to_remove property — other_settings.header_options / 5d4e69652685 / 4

Type: `["list", "string"]`. Optional.

List of keys of Headers to be removed from the HTTP request being sent towards upstream.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.items.string.pattern": "^[0-9A-Za-z_\\\\-\\\\.]+$",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.items.string.pattern": "^[0-9A-Za-z_\\\\-\\\\.]+$",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [response_headers_to_add](resources--cdn_loadbalancer--reference--group-012.md#canonical-37a2e73252f0c5f2e400fdd23fa6bacd2cb864a19cb5b14313dac520aca15eba): complete subsection reference.

<a id="canonical-b0a1924c4cd5d434e79cd73d68c329ee58d28b9490d93db22d765cad2864cbcd"></a>

<a id="canonical-38e5189d1e5f5194cb55a1440efa912213ea5a904c47230662798f896a3aff6a"></a>

## response_headers_to_remove property — other_settings.header_options / 5d4e69652685 / 5

Type: `["list", "string"]`. Optional.

List of keys of Headers to be removed from the HTTP response being sent towards downstream.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

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
    "ves.io.schema.rules.repeated.items.string.pattern": "^[0-9A-Za-z_\\\\-\\\\.]+$",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.pattern": "^[0-9A-Za-z_\\\\-\\\\.]+$",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3f7c732ab2a5d209b3d7552f2daf5f5c36ac2a4e863ca8563dee383f164cbc3b"></a>

## Next pages — other_settings.header_options / 5d4e69652685 / 6

- [other_settings.header_options.request_headers_to_add](resources--cdn_loadbalancer--reference--group-012.md#canonical-d4798b10c5cf6cce2b217dade31c80cc63e30f112061600bc31514aef034f162)
- [other_settings.header_options.response_headers_to_add](resources--cdn_loadbalancer--reference--group-012.md#canonical-37a2e73252f0c5f2e400fdd23fa6bacd2cb864a19cb5b14313dac520aca15eba)
- [other_settings](resources--cdn_loadbalancer--reference--group-012.md#canonical-6237c1513ed4eb47b747c88f4987c86621425202c792c064ab4bfdfacbad388a)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-d4798b10c5cf6cce2b217dade31c80cc63e30f112061600bc31514aef034f162"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-53d4ea91683046a47ed03cd1d621bf3fdac2bdf3079bb99fc2609ef904fe8b1a"></a>

## other_settings.header_options.request_headers_to_add — other_settings.header_options.request_headers_to_add / 6fa94a0a3a72 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [other_settings](resources--cdn_loadbalancer--reference--group-012.md#canonical-6237c1513ed4eb47b747c88f4987c86621425202c792c064ab4bfdfacbad388a)
- [other_settings.header_options](resources--cdn_loadbalancer--reference--group-012.md#canonical-2d180d161dffb067f7db1e0720d11e345c4d3005e1d9ea84233978dd5f6e4cdb)
- other_settings.header_options.request_headers_to_add

<a id="canonical-a7c04b80144eef0c366bc7bf1afe167742031765e43e1d3648555d1a736d6288"></a>

Type: `"object"`. list nested block, Optional.

Headers are key-value pairs to be added to HTTP request being routed towards upstream. Headers
specified at this level are applied after headers from matched Route are applied.

Upstream description:

Headers are key-value pairs to be added to HTTP request being routed towards upstream. Headers
specified at this level are applied after headers from matched Route are applied.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("secret_value",
    "value")}
```

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

Terraform syntax:

```terraform
request_headers_to_add {
  # Configure direct properties listed below.
}
```

<a id="canonical-eb20e43b08818aee7bab029d3a578d7e5d2dd3694ad52e10e44df6d4b91b36c7"></a>

## Direct properties — other_settings.header_options.request_headers_to_add / 6fa94a0a3a72 / 3

<a id="canonical-d3798a7d5b05d9668ea6e1e0764145e4173f5c21ddd28494f13414e0bf9890e9"></a>

<a id="canonical-545dc7292ced8aa5fb9673053689d39a7be078039b605f44e95646bd0cb5aa8b"></a>

## append property — other_settings.header_options.request_headers_to_add / 6fa94a0a3a72 / 4

Type: `"bool"`. Optional.

Should the value be appended? If true, the value is appended to existing values. not append.
Defaults to \`do\`.

Upstream description:

Should the value be appended? If true, the value is appended to existing values. Default value is do
not append.

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

<a id="canonical-2ebf1611a69185561f07e66ee93d43e73a407cba5a9a68d7baaaf734286ff3cd"></a>

<a id="canonical-a127b429586b5a44b908f272c88e9644d0871d988fe86b1abb4bc3b0c53bd304"></a>

## name property — other_settings.header_options.request_headers_to_add / 6fa94a0a3a72 / 5

Type: `"string"`. Optional.

Name. Name of the HTTP header.

Upstream description:

Name of the HTTP header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

- [secret_value](resources--cdn_loadbalancer--reference--group-012.md#canonical-54655a9afee4324aa1176effe02a8013a23c0588ec540bd531574a16fd718327): complete subsection reference.

<a id="canonical-ca163ec7710b15d17ef30a7f2a884ff466a8758c406b3113bee494e784a1df3b"></a>

<a id="canonical-3d5e5250fba92d89e7be67be46b1772f643752593a06399ae47c05a9841cb533"></a>

## value property — other_settings.header_options.request_headers_to_add / 6fa94a0a3a72 / 6

Type: `"string"`. Optional.

Exclusive with \[secret\_value\] Value of the HTTP header.

Upstream description:

Exclusive with \[secret\_value\] Value of the HTTP header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8096),
}
```

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

<a id="canonical-421c7540ce634b138047a185f813a810663683e82bb9c4b964ada0eba11f8f2d"></a>

## Next pages — other_settings.header_options.request_headers_to_add / 6fa94a0a3a72 / 7

- [other_settings.header_options.request_headers_to_add.secret_value](resources--cdn_loadbalancer--reference--group-012.md#canonical-54655a9afee4324aa1176effe02a8013a23c0588ec540bd531574a16fd718327)
- [other_settings.header_options](resources--cdn_loadbalancer--reference--group-012.md#canonical-2d180d161dffb067f7db1e0720d11e345c4d3005e1d9ea84233978dd5f6e4cdb)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-54655a9afee4324aa1176effe02a8013a23c0588ec540bd531574a16fd718327"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7b92c3f2a5a65fa8d4a5633016a5af96869dca599262c0bb5c874db880276498"></a>

## other_settings.header_options.request_headers_to_add.secret_value — other_settings.header_options.request_headers_to_add.secret_value / 6bf4cd0a439d / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [other_settings](resources--cdn_loadbalancer--reference--group-012.md#canonical-6237c1513ed4eb47b747c88f4987c86621425202c792c064ab4bfdfacbad388a)
- [other_settings.header_options](resources--cdn_loadbalancer--reference--group-012.md#canonical-2d180d161dffb067f7db1e0720d11e345c4d3005e1d9ea84233978dd5f6e4cdb)
- [other_settings.header_options.request_headers_to_add](resources--cdn_loadbalancer--reference--group-012.md#canonical-d4798b10c5cf6cce2b217dade31c80cc63e30f112061600bc31514aef034f162)
- other_settings.header_options.request_headers_to_add.secret_value

<a id="canonical-fd115aaa901e81789a478cb5149fe97555e91a911ebe6d8fd5da7a1894d8eac0"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
secret_value {
  # Configure direct properties listed below.
}
```

<a id="canonical-1099a69d5a5fd7a66529e4d7901e5b2b985504cbd2d67bcee0073e203fac457b"></a>

## Direct properties — other_settings.header_options.request_headers_to_add.secret_value / 6bf4cd0a439d / 3

- [blindfold_secret_info](resources--cdn_loadbalancer--reference--group-012.md#canonical-bdac05043bbdb8962dfc64ac7cda26256cc4e8607da1f32cd3a09b1f1f0e9ae9): complete subsection reference.

- [clear_secret_info](resources--cdn_loadbalancer--reference--group-012.md#canonical-f72e28b2f8990fce5059c8e8599c2698fabfdfd9c8f11f2591869b53104c2f95): complete subsection reference.

<a id="canonical-28647e8ab6b39b650741e320bc0e61e6212452eaed6132faf0c39bc4194790a8"></a>

## Next pages — other_settings.header_options.request_headers_to_add.secret_value / 6bf4cd0a439d / 4

- [other_settings.header_options.request_headers_to_add.secret_value.blindfold_secret_info](resources--cdn_loadbalancer--reference--group-012.md#canonical-bdac05043bbdb8962dfc64ac7cda26256cc4e8607da1f32cd3a09b1f1f0e9ae9)
- [other_settings.header_options.request_headers_to_add.secret_value.clear_secret_info](resources--cdn_loadbalancer--reference--group-012.md#canonical-f72e28b2f8990fce5059c8e8599c2698fabfdfd9c8f11f2591869b53104c2f95)
- [other_settings.header_options.request_headers_to_add](resources--cdn_loadbalancer--reference--group-012.md#canonical-d4798b10c5cf6cce2b217dade31c80cc63e30f112061600bc31514aef034f162)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-bdac05043bbdb8962dfc64ac7cda26256cc4e8607da1f32cd3a09b1f1f0e9ae9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fe71da41a77a0ce2694ace795d41fe98ae0120367104697c9de2d0727727b76e"></a>

## other_settings.header_options.request_headers_to_add.secret_value.blindfold_secret_info — other_settings.header_options.request_headers_to_add.secret_value.blindfold_secr / 4fb6adb57551 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [other_settings](resources--cdn_loadbalancer--reference--group-012.md#canonical-6237c1513ed4eb47b747c88f4987c86621425202c792c064ab4bfdfacbad388a)
- [other_settings.header_options](resources--cdn_loadbalancer--reference--group-012.md#canonical-2d180d161dffb067f7db1e0720d11e345c4d3005e1d9ea84233978dd5f6e4cdb)
- [other_settings.header_options.request_headers_to_add](resources--cdn_loadbalancer--reference--group-012.md#canonical-d4798b10c5cf6cce2b217dade31c80cc63e30f112061600bc31514aef034f162)
- [other_settings.header_options.request_headers_to_add.secret_value](resources--cdn_loadbalancer--reference--group-012.md#canonical-54655a9afee4324aa1176effe02a8013a23c0588ec540bd531574a16fd718327)
- other_settings.header_options.request_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-5755b6ae677c9cb29ffd8b761ec68e4ec79af9f7d5e9417dd9fc2b94dddc0c05"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
```

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

Terraform syntax:

```terraform
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-be6335334286d16a810f15f43582a610d23b5af8054b280ead8820e140fd6bae"></a>

## Direct properties — other_settings.header_options.request_headers_to_add.secret_value.blindfold_secr / 4fb6adb57551 / 3

<a id="canonical-94a2b7ec0589bcc0b531af593ba8eaaf040cd53fe44b5901d5b6f532d4ec6407"></a>

<a id="canonical-8c3b5c5032e51cc513c577e409706640fe355520d208731c18e86f4a34536bfa"></a>

## decryption_provider property — other_settings.header_options.request_headers_to_add.secret_value.blindfold_secr / 4fb6adb57551 / 4

Type: `"string"`. Optional.

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

<a id="canonical-bdb8385a092c8613256e9df5125764b18444a72567f3fc13f0fd2a79086af056"></a>

<a id="canonical-d21d21cc23cbd1b00d544ed0324570ec1db8dcb0b636ca4060438067b37862f2"></a>

## location property — other_settings.header_options.request_headers_to_add.secret_value.blindfold_secr / 4fb6adb57551 / 5

Type: `"string"`. Optional, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

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

<a id="canonical-91c656ee371c71bbc33bdc4eb6fad0beecd2a477ab8c405770ba4ea6d8152521"></a>

<a id="canonical-0d3517ed4b76e7e7d3452310e991f9193cb24c3913afb0182d2b0f7256ecc898"></a>

## store_provider property — other_settings.header_options.request_headers_to_add.secret_value.blindfold_secr / 4fb6adb57551 / 6

Type: `"string"`. Optional.

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

<a id="canonical-ac6407eebf6c0e98d9c33f1738958c2225e0ef34548d10e31bee616d46b1f521"></a>

## Next pages — other_settings.header_options.request_headers_to_add.secret_value.blindfold_secr / 4fb6adb57551 / 7

- [other_settings.header_options.request_headers_to_add.secret_value](resources--cdn_loadbalancer--reference--group-012.md#canonical-54655a9afee4324aa1176effe02a8013a23c0588ec540bd531574a16fd718327)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-f72e28b2f8990fce5059c8e8599c2698fabfdfd9c8f11f2591869b53104c2f95"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eaac9e05ed7481ae0b7d85e4014ee6897fb963033cef029eae185a47929d8cde"></a>

## other_settings.header_options.request_headers_to_add.secret_value.clear_secret_info — other_settings.header_options.request_headers_to_add.secret_value.clear_secret_i / 64777a4dde95 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [other_settings](resources--cdn_loadbalancer--reference--group-012.md#canonical-6237c1513ed4eb47b747c88f4987c86621425202c792c064ab4bfdfacbad388a)
- [other_settings.header_options](resources--cdn_loadbalancer--reference--group-012.md#canonical-2d180d161dffb067f7db1e0720d11e345c4d3005e1d9ea84233978dd5f6e4cdb)
- [other_settings.header_options.request_headers_to_add](resources--cdn_loadbalancer--reference--group-012.md#canonical-d4798b10c5cf6cce2b217dade31c80cc63e30f112061600bc31514aef034f162)
- [other_settings.header_options.request_headers_to_add.secret_value](resources--cdn_loadbalancer--reference--group-012.md#canonical-54655a9afee4324aa1176effe02a8013a23c0588ec540bd531574a16fd718327)
- other_settings.header_options.request_headers_to_add.secret_value.clear_secret_info

<a id="canonical-c5bce739939727f759fc4ba03d01bc0ece063a2188343ae05906af5c0f157a18"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
```

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

Terraform syntax:

```terraform
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-a0fdb933c69edcd97884d540098666d1d3f0a33de6d9745a1a2e5fc2e2c1a922"></a>

## Direct properties — other_settings.header_options.request_headers_to_add.secret_value.clear_secret_i / 64777a4dde95 / 3

<a id="canonical-66f6feae705d1e484b493972073703642c15defb77f65dc70730da598d897e1c"></a>

<a id="canonical-f63413ad60db9da9f112fd67a44fb009af7ffaade6a416d1b714c9d1740a0963"></a>

## provider_ref property — other_settings.header_options.request_headers_to_add.secret_value.clear_secret_i / 64777a4dde95 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-c13b44bfbbd32181333b21b5f7973090b7207bda858653a584f94dab81cd3493"></a>

<a id="canonical-01b6156d74d62616937d6676809e7f99f1721c24da286f2372c9ab39e1a28ee3"></a>

## url property — other_settings.header_options.request_headers_to_add.secret_value.clear_secret_i / 64777a4dde95 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

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

<a id="canonical-16d33c533cad9a2ac6b85db70b13dc6bd77708e719d8b5bd84005016aba46e68"></a>

## Next pages — other_settings.header_options.request_headers_to_add.secret_value.clear_secret_i / 64777a4dde95 / 6

- [other_settings.header_options.request_headers_to_add.secret_value](resources--cdn_loadbalancer--reference--group-012.md#canonical-54655a9afee4324aa1176effe02a8013a23c0588ec540bd531574a16fd718327)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-37a2e73252f0c5f2e400fdd23fa6bacd2cb864a19cb5b14313dac520aca15eba"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-63f83625b27663ce8bc1f7f9e5f63675c51909d6da9ad31275c85c03e6b2d89d"></a>

## other_settings.header_options.response_headers_to_add — other_settings.header_options.response_headers_to_add / 10d98e884593 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [other_settings](resources--cdn_loadbalancer--reference--group-012.md#canonical-6237c1513ed4eb47b747c88f4987c86621425202c792c064ab4bfdfacbad388a)
- [other_settings.header_options](resources--cdn_loadbalancer--reference--group-012.md#canonical-2d180d161dffb067f7db1e0720d11e345c4d3005e1d9ea84233978dd5f6e4cdb)
- other_settings.header_options.response_headers_to_add

<a id="canonical-907c7c99863e82f22a0addf7aa776389298c32c317ffce9b0e7ec43de6b68328"></a>

Type: `"object"`. list nested block, Optional.

Headers are key-value pairs to be added to HTTP response being sent towards downstream. Headers
specified at this level are applied after headers from matched Route are applied.

Upstream description:

Headers are key-value pairs to be added to HTTP response being sent towards downstream. Headers
specified at this level are applied after headers from matched Route are applied.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("secret_value",
    "value")}
```

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

Terraform syntax:

```terraform
response_headers_to_add {
  # Configure direct properties listed below.
}
```

<a id="canonical-9d295b496d71274134879c41c33ce5b5dbb1f53dd78589fef724387e712639cb"></a>

## Direct properties — other_settings.header_options.response_headers_to_add / 10d98e884593 / 3

<a id="canonical-188b4c110366c84b697c02177d0014eafc4d35376d681adf1b43d17f1f7eccce"></a>

<a id="canonical-474695168050729abe823730f070df3ae2a2f9bd5e40b9f511dc4f0f537e4f5e"></a>

## append property — other_settings.header_options.response_headers_to_add / 10d98e884593 / 4

Type: `"bool"`. Optional.

Should the value be appended? If true, the value is appended to existing values. not append.
Defaults to \`do\`.

Upstream description:

Should the value be appended? If true, the value is appended to existing values. Default value is do
not append.

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

<a id="canonical-dba253337f1faee6aca3485c846fb309d9915a8ab3a31a57b55193244d601f04"></a>

<a id="canonical-e12d97251e3789eb36f22c849da5902714622059413b9f71464d47f423ef3ad6"></a>

## name property — other_settings.header_options.response_headers_to_add / 10d98e884593 / 5

Type: `"string"`. Optional.

Name. Name of the HTTP header.

Upstream description:

Name of the HTTP header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

- [secret_value](resources--cdn_loadbalancer--reference--group-012.md#canonical-7166f49d195fc76d8a96e71552ed6d550a5f9e24a9550c0401d47305f27cf1da): complete subsection reference.

<a id="canonical-70d6e363897268bfe1f1b46d69087b77a3b0dfd78a03d602afe6dc1db3710cf7"></a>

<a id="canonical-f48bacb9e52173630cc5c7e7cdefb77b25e1a7e0ecf1d192ff73c88e950c464a"></a>

## value property — other_settings.header_options.response_headers_to_add / 10d98e884593 / 6

Type: `"string"`. Optional.

Exclusive with \[secret\_value\] Value of the HTTP header.

Upstream description:

Exclusive with \[secret\_value\] Value of the HTTP header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8096),
}
```

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

<a id="canonical-752292a18b735bfb91dd381599e0884b47783f67eb988b5bd9de818d68277b2f"></a>

## Next pages — other_settings.header_options.response_headers_to_add / 10d98e884593 / 7

- [other_settings.header_options.response_headers_to_add.secret_value](resources--cdn_loadbalancer--reference--group-012.md#canonical-7166f49d195fc76d8a96e71552ed6d550a5f9e24a9550c0401d47305f27cf1da)
- [other_settings.header_options](resources--cdn_loadbalancer--reference--group-012.md#canonical-2d180d161dffb067f7db1e0720d11e345c4d3005e1d9ea84233978dd5f6e4cdb)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-7166f49d195fc76d8a96e71552ed6d550a5f9e24a9550c0401d47305f27cf1da"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c999904047f89451fa883152e9980f7b9ac6e4d24afd856a30804cd111ee273b"></a>

## other_settings.header_options.response_headers_to_add.secret_value — other_settings.header_options.response_headers_to_add.secret_value / 6943f8fd920a / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [other_settings](resources--cdn_loadbalancer--reference--group-012.md#canonical-6237c1513ed4eb47b747c88f4987c86621425202c792c064ab4bfdfacbad388a)
- [other_settings.header_options](resources--cdn_loadbalancer--reference--group-012.md#canonical-2d180d161dffb067f7db1e0720d11e345c4d3005e1d9ea84233978dd5f6e4cdb)
- [other_settings.header_options.response_headers_to_add](resources--cdn_loadbalancer--reference--group-012.md#canonical-37a2e73252f0c5f2e400fdd23fa6bacd2cb864a19cb5b14313dac520aca15eba)
- other_settings.header_options.response_headers_to_add.secret_value

<a id="canonical-944c3131e7c138c70ba8618ca93c1d6fe22667ef005ee9f20fc27113823b15d9"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
secret_value {
  # Configure direct properties listed below.
}
```

<a id="canonical-7b6916dc5730045136dd88d563bb61994272d631dcb258f317f52b1d3ce81fb2"></a>

## Direct properties — other_settings.header_options.response_headers_to_add.secret_value / 6943f8fd920a / 3

- [blindfold_secret_info](resources--cdn_loadbalancer--reference--group-012.md#canonical-1465ea8630c4bfffcc565837c56ce1db371352f56b99572fe42a122b8a853eb8): complete subsection reference.

- [clear_secret_info](resources--cdn_loadbalancer--reference--group-012.md#canonical-138f5e3f06a69afb99ea214f0ed440f7e0ed88171eb43710d0d005e3ad5058ee): complete subsection reference.

<a id="canonical-229967460a60ef54e443ac2f526d7caeb75d3598e0865c02749b9b54f3c1a2c3"></a>

## Next pages — other_settings.header_options.response_headers_to_add.secret_value / 6943f8fd920a / 4

- [other_settings.header_options.response_headers_to_add.secret_value.blindfold_secret_info](resources--cdn_loadbalancer--reference--group-012.md#canonical-1465ea8630c4bfffcc565837c56ce1db371352f56b99572fe42a122b8a853eb8)
- [other_settings.header_options.response_headers_to_add.secret_value.clear_secret_info](resources--cdn_loadbalancer--reference--group-012.md#canonical-138f5e3f06a69afb99ea214f0ed440f7e0ed88171eb43710d0d005e3ad5058ee)
- [other_settings.header_options.response_headers_to_add](resources--cdn_loadbalancer--reference--group-012.md#canonical-37a2e73252f0c5f2e400fdd23fa6bacd2cb864a19cb5b14313dac520aca15eba)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-1465ea8630c4bfffcc565837c56ce1db371352f56b99572fe42a122b8a853eb8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7a29ce1021e0d1c164bb1c6735830550528384b0bd188bb2c24550dd6314aa14"></a>

## other_settings.header_options.response_headers_to_add.secret_value.blindfold_secret_info — other_settings.header_options.response_headers_to_add.secret_value.blindfold_sec / 2904f44054d7 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [other_settings](resources--cdn_loadbalancer--reference--group-012.md#canonical-6237c1513ed4eb47b747c88f4987c86621425202c792c064ab4bfdfacbad388a)
- [other_settings.header_options](resources--cdn_loadbalancer--reference--group-012.md#canonical-2d180d161dffb067f7db1e0720d11e345c4d3005e1d9ea84233978dd5f6e4cdb)
- [other_settings.header_options.response_headers_to_add](resources--cdn_loadbalancer--reference--group-012.md#canonical-37a2e73252f0c5f2e400fdd23fa6bacd2cb864a19cb5b14313dac520aca15eba)
- [other_settings.header_options.response_headers_to_add.secret_value](resources--cdn_loadbalancer--reference--group-012.md#canonical-7166f49d195fc76d8a96e71552ed6d550a5f9e24a9550c0401d47305f27cf1da)
- other_settings.header_options.response_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-16f828048bd54ca54ba054bd6e2abe76c093a178e34b7f1aca497dfb678e15ae"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
```

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

Terraform syntax:

```terraform
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-5baa3b1469f6c286f8c02a3e22f08873aac8828e9c1dfd98ba90599aefb44983"></a>

## Direct properties — other_settings.header_options.response_headers_to_add.secret_value.blindfold_sec / 2904f44054d7 / 3

<a id="canonical-62d9aec25c61d48dec8324ae1ac1d0b3b21a2b4984d41ffe6af88b9bcd80de05"></a>

<a id="canonical-e29ffd10e5e47d325fb7fb78621c9a51e810e7de047046f91e986b2c50570679"></a>

## decryption_provider property — other_settings.header_options.response_headers_to_add.secret_value.blindfold_sec / 2904f44054d7 / 4

Type: `"string"`. Optional.

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

<a id="canonical-29edf4774e18bcf1a10b254a561a59d821540f431e68758ce0bf58a793c1f263"></a>

<a id="canonical-03aa7353d78bc4c14717ab97e63bb23aa2a70b570fb61f470f62c485f03ebb4a"></a>

## location property — other_settings.header_options.response_headers_to_add.secret_value.blindfold_sec / 2904f44054d7 / 5

Type: `"string"`. Optional, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

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

<a id="canonical-6f33c2385552577581a1e07bdb64135d194637e8d6acf7ff7ee794ab70856980"></a>

<a id="canonical-d79f744336296d4f3f35f0f7ac0552a95e27109b712cfad30fedbed899635796"></a>

## store_provider property — other_settings.header_options.response_headers_to_add.secret_value.blindfold_sec / 2904f44054d7 / 6

Type: `"string"`. Optional.

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

<a id="canonical-054154247e63925a90cb60e93c06650b62f7336de5314bbcc725d0c7ad61e5dd"></a>

## Next pages — other_settings.header_options.response_headers_to_add.secret_value.blindfold_sec / 2904f44054d7 / 7

- [other_settings.header_options.response_headers_to_add.secret_value](resources--cdn_loadbalancer--reference--group-012.md#canonical-7166f49d195fc76d8a96e71552ed6d550a5f9e24a9550c0401d47305f27cf1da)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-138f5e3f06a69afb99ea214f0ed440f7e0ed88171eb43710d0d005e3ad5058ee"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e407ba09430e6ac31d990c78f6a4522168bbe26e89f0b6a00070befa95ac5ac9"></a>

## other_settings.header_options.response_headers_to_add.secret_value.clear_secret_info — other_settings.header_options.response_headers_to_add.secret_value.clear_secret_ / 52516ec5d154 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [other_settings](resources--cdn_loadbalancer--reference--group-012.md#canonical-6237c1513ed4eb47b747c88f4987c86621425202c792c064ab4bfdfacbad388a)
- [other_settings.header_options](resources--cdn_loadbalancer--reference--group-012.md#canonical-2d180d161dffb067f7db1e0720d11e345c4d3005e1d9ea84233978dd5f6e4cdb)
- [other_settings.header_options.response_headers_to_add](resources--cdn_loadbalancer--reference--group-012.md#canonical-37a2e73252f0c5f2e400fdd23fa6bacd2cb864a19cb5b14313dac520aca15eba)
- [other_settings.header_options.response_headers_to_add.secret_value](resources--cdn_loadbalancer--reference--group-012.md#canonical-7166f49d195fc76d8a96e71552ed6d550a5f9e24a9550c0401d47305f27cf1da)
- other_settings.header_options.response_headers_to_add.secret_value.clear_secret_info

<a id="canonical-b57c9e0b9f06dd0aba771babd2f7ea710fea76efcf51eb3a973bb3af579a7f5e"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
```

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

Terraform syntax:

```terraform
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-d09c1c5c628beb6d8fa2a03fa5153156185f142bb7d67c4ca2d9d011bf9572fc"></a>

## Direct properties — other_settings.header_options.response_headers_to_add.secret_value.clear_secret_ / 52516ec5d154 / 3

<a id="canonical-71f7c0421d50a6c297e203c2f32b38be76fc54dda2915a4333a16ae30d51c2ad"></a>

<a id="canonical-b477ea830b5c6b109301e7a394d4a596410a7174eee5c91c7b251ec869495daa"></a>

## provider_ref property — other_settings.header_options.response_headers_to_add.secret_value.clear_secret_ / 52516ec5d154 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-57782513bc6b3876bbdc17591b3a751a831f618b543098cf1fbcb19856c4259f"></a>

<a id="canonical-dd2a40b1e18604c9fe85073f0003decda96382409cef8ce144941a55d973095b"></a>

## url property — other_settings.header_options.response_headers_to_add.secret_value.clear_secret_ / 52516ec5d154 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

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

<a id="canonical-fef5a5c9db47a5359a7b6f3fd6c2912d66fdf229b2c923d9811ee38ac24753e1"></a>

## Next pages — other_settings.header_options.response_headers_to_add.secret_value.clear_secret_ / 52516ec5d154 / 6

- [other_settings.header_options.response_headers_to_add.secret_value](resources--cdn_loadbalancer--reference--group-012.md#canonical-7166f49d195fc76d8a96e71552ed6d550a5f9e24a9550c0401d47305f27cf1da)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-9fd681a2d1eb0f615230fa310f857bb9bca5668fad19c13074eec63c56995bdc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2fd5c557a6beb355cd3a5acf2a4d182ae1474df3271231f2982ea471929796e2"></a>

## other_settings.logging_options — other_settings.logging_options / 802e3bf5f12e / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [other_settings](resources--cdn_loadbalancer--reference--group-012.md#canonical-6237c1513ed4eb47b747c88f4987c86621425202c792c064ab4bfdfacbad388a)
- other_settings.logging_options

<a id="canonical-8f3f5d0fd6be80267a20ab539d0e929335598fa49dc1cfb89ccfe6472f30a6c1"></a>

Type: `"object"`. single nested block, Optional.

Defines various OPTIONS related to logging.

Upstream description:

This defines various OPTIONS related to logging.

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

Terraform syntax:

```terraform
logging_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-2e816150e62b24970f4814ac25e21def5c2ab192f11f2a1238b3fabc586f6513"></a>

## Direct properties — other_settings.logging_options / 802e3bf5f12e / 3

- [client_log_options](resources--cdn_loadbalancer--reference--group-012.md#canonical-d0a335080c54d06f04900eb765d2e1762f138d71804977a47875212bdb512964): complete subsection reference.

- [origin_log_options](resources--cdn_loadbalancer--reference--group-012.md#canonical-51e87ef1ab501018986ef459ea0d5bf73f1a40ab4eaf89255cbd333b605cd727): complete subsection reference.

<a id="canonical-3c63b8b9a7fc03cd5af819d9451cd4e0d4edfa49e60ad23ad7be23b054cf9fbe"></a>

## Next pages — other_settings.logging_options / 802e3bf5f12e / 4

- [other_settings.logging_options.client_log_options](resources--cdn_loadbalancer--reference--group-012.md#canonical-d0a335080c54d06f04900eb765d2e1762f138d71804977a47875212bdb512964)
- [other_settings.logging_options.origin_log_options](resources--cdn_loadbalancer--reference--group-012.md#canonical-51e87ef1ab501018986ef459ea0d5bf73f1a40ab4eaf89255cbd333b605cd727)
- [other_settings](resources--cdn_loadbalancer--reference--group-012.md#canonical-6237c1513ed4eb47b747c88f4987c86621425202c792c064ab4bfdfacbad388a)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-d0a335080c54d06f04900eb765d2e1762f138d71804977a47875212bdb512964"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-706bb2ff87aa6fcbe95d49646b9ca8202dcbb066641239a583a51079317bc32e"></a>

## other_settings.logging_options.client_log_options — other_settings.logging_options.client_log_options / 5e1f030f0123 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [other_settings](resources--cdn_loadbalancer--reference--group-012.md#canonical-6237c1513ed4eb47b747c88f4987c86621425202c792c064ab4bfdfacbad388a)
- [other_settings.logging_options](resources--cdn_loadbalancer--reference--group-012.md#canonical-9fd681a2d1eb0f615230fa310f857bb9bca5668fad19c13074eec63c56995bdc)
- other_settings.logging_options.client_log_options

<a id="canonical-c8bcd381e15057b2a1c0423d4b89a51bc049b49fd82032ab1832f9d50d3fad90"></a>

Type: `"object"`. single nested block, Optional.

Headers to Log. List of headers to Log.

Upstream description:

List of headers to Log.

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

Terraform syntax:

```terraform
client_log_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-0bf13543952fd5bcf36825ac7a87e655b106bec43ac7fc8064ba6ef5a5cdd516"></a>

## Direct properties — other_settings.logging_options.client_log_options / 5e1f030f0123 / 3

<a id="canonical-6f18c7224b492b4a2aa815084c1b4f90afac88d3baf789c7eb1f112446e6e3e7"></a>

<a id="canonical-3a15082c233beaff480ff14fd70d31b7d24ffab66b8e0235e09cf9999018049a"></a>

## header_list property — other_settings.logging_options.client_log_options / 5e1f030f0123 / 4

Type: `["list", "string"]`. Optional.

Headers. List of headers.

Upstream description:

List of headers.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-fe4f05882eb0c280f466dbf77f61f0732f9823b5d1e350fb74ebf8fb3a4b46b1"></a>

## Next pages — other_settings.logging_options.client_log_options / 5e1f030f0123 / 5

- [other_settings.logging_options](resources--cdn_loadbalancer--reference--group-012.md#canonical-9fd681a2d1eb0f615230fa310f857bb9bca5668fad19c13074eec63c56995bdc)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-51e87ef1ab501018986ef459ea0d5bf73f1a40ab4eaf89255cbd333b605cd727"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c0440dbe5913e2892feebb087457fa1d103624e71ce09ac178a1c088b2d88708"></a>

## other_settings.logging_options.origin_log_options — other_settings.logging_options.origin_log_options / bc3ea79aee2f / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [other_settings](resources--cdn_loadbalancer--reference--group-012.md#canonical-6237c1513ed4eb47b747c88f4987c86621425202c792c064ab4bfdfacbad388a)
- [other_settings.logging_options](resources--cdn_loadbalancer--reference--group-012.md#canonical-9fd681a2d1eb0f615230fa310f857bb9bca5668fad19c13074eec63c56995bdc)
- other_settings.logging_options.origin_log_options

<a id="canonical-d68dd8ada87ee4df2936eebf397a50896947d167640cf6da67fcdc72bb949341"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for origin log options.

Upstream description:

List of headers to Log.

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

Terraform syntax:

```terraform
origin_log_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-35c1b96849fa1e0a60f2a1a46f62fd9c34d9039ba150d0cc3e95466a05c88efa"></a>

## Direct properties — other_settings.logging_options.origin_log_options / bc3ea79aee2f / 3

<a id="canonical-77d3121f5c26678c23b9a267661f0046fb62a0f8eb171398a163acb0010f0eb3"></a>

<a id="canonical-e974336c69d50cce599a7e5a8623a29bc576a2a019d9f1000910cbb6d906acb0"></a>

## header_list property — other_settings.logging_options.origin_log_options / bc3ea79aee2f / 4

Type: `["list", "string"]`. Optional.

Headers. List of headers.

Upstream description:

List of headers.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-33b7cf65c418141fb0a337d2db7a96041b3ffd5d161b69fe34443e2e9cfedee1"></a>

## Next pages — other_settings.logging_options.origin_log_options / bc3ea79aee2f / 5

- [other_settings.logging_options](resources--cdn_loadbalancer--reference--group-012.md#canonical-9fd681a2d1eb0f615230fa310f857bb9bca5668fad19c13074eec63c56995bdc)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9fa535951579453c3ddc91cd31534c008bad613961bc46489bd427e59ac6a756"></a>

## policy_based_challenge — policy_based_challenge / 961a19cc189c / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- policy_based_challenge

<a id="canonical-d4be8fe0e687fee1885e9b29faa62a129ea348a6ad5c2bd7c73688b1e3a5fa88"></a>

Type: `"object"`. single nested block, Optional.

Specifies the settings for policy rule based challenge.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("always_enable_captcha_challenge",
    "always_enable_js_challenge"),
  validators.ConflictingObjectAttributes("always_enable_captcha_challenge",
    "no_challenge"),
  validators.ConflictingObjectAttributes("always_enable_js_challenge",
    "no_challenge"),
  validators.ConflictingObjectAttributes("captcha_challenge_parameters",
    "default_captcha_challenge_parameters"),
  validators.ConflictingObjectAttributes("default_js_challenge_parameters",
    "js_challenge_parameters"),
  validators.ConflictingObjectAttributes("default_mitigation_settings",
    "malicious_user_mitigation"),
  validators.ConflictingObjectAttributes("default_temporary_blocking_parameters",
    "temporary_user_blocking")}
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
  "x-ves-oneof-field-captcha_challenge_parameters_choice": "[\"captcha_challenge_parameters\",\"default_captcha_challenge_parameters\"]",
  "x-ves-oneof-field-challenge_choice": "[\"always_enable_captcha_challenge\",\"always_enable_js_challenge\",\"no_challenge\"]",
  "x-ves-oneof-field-js_challenge_parameters_choice": "[\"default_js_challenge_parameters\",\"js_challenge_parameters\"]",
  "x-ves-oneof-field-malicious_user_mitigation_choice": "[\"default_mitigation_settings\",\"malicious_user_mitigation\"]",
  "x-ves-oneof-field-temporary_blocking_parameters_choice": "[\"default_temporary_blocking_parameters\",\"temporary_user_blocking\"]"
}
```

Terraform syntax:

```terraform
policy_based_challenge {
  # Configure direct properties listed below.
}
```

<a id="canonical-a155620768b796961fa7e21ef0087b88df363c24537b2c8fd373ba18ad8d806a"></a>

## Direct properties — policy_based_challenge / 961a19cc189c / 3

- [always_enable_captcha_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-ed41e8b257dba61d5faa50328ed39449fb350f60099ec38f09a1745133c68a4d): complete subsection reference.

- [always_enable_js_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-e8de5bd4e333592379aa9c51751add264265d3277e414e5b77568d04d6357119): complete subsection reference.

- [captcha_challenge_parameters](resources--cdn_loadbalancer--reference--group-013.md#canonical-7b0045090e8c5ef02ef0ee49e64dc70034bfde0d7833823e9b5744869cc7527f): complete subsection reference.

- [default_captcha_challenge_parameters](resources--cdn_loadbalancer--reference--group-013.md#canonical-d9f646b33f10838cb5c1fa84d31cba280e5e29994f263a1f481fb63fdca0c6a5): complete subsection reference.

- [default_js_challenge_parameters](resources--cdn_loadbalancer--reference--group-013.md#canonical-692ec2ce0a9ecfaf67735387d16a2d6fe5151fe4cd40e59433e26c6985cb9611): complete subsection reference.

- [default_mitigation_settings](resources--cdn_loadbalancer--reference--group-013.md#canonical-18b41c01fa4dc61af170cf40418a88eda5822bfc3294987633cb23c646109550): complete subsection reference.

- [default_temporary_blocking_parameters](resources--cdn_loadbalancer--reference--group-013.md#canonical-8af097ae92a51a14440a8a17fb2c5207137124945ac0711b1367170897ddc838): complete subsection reference.

- [js_challenge_parameters](resources--cdn_loadbalancer--reference--group-013.md#canonical-87630b046362400b830792b5610ed8697a78a4d5d3c8ad67ccefe199393f7342): complete subsection reference.

- [malicious_user_mitigation](resources--cdn_loadbalancer--reference--group-013.md#canonical-a73000a20bec1df010b97ca8b8f4f3989e4041fc954594295ea8e4bd63edbfa6): complete subsection reference.

- [no_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-e781909cfb9de5bc972cf81b5d973b59e0fb9cbd5403f19e92a44af189803430): complete subsection reference.

- [rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-e7e050a0106b5c243929bc99c717fe27390190239f5f0da6543c11a425e13c71): complete subsection reference.

- [temporary_user_blocking](resources--cdn_loadbalancer--reference--group-014.md#canonical-885e305a1d0df444d63d64da6c7a4b0a54459b03c5eb86b881c6ad7b4cb37d06): complete subsection reference.

<a id="canonical-cc67555e9fc9ba5f89608a697191250b1bb53119c74a2c51c0f8763fcfad8237"></a>

## Next pages — policy_based_challenge / 961a19cc189c / 4

- [policy_based_challenge.always_enable_captcha_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-ed41e8b257dba61d5faa50328ed39449fb350f60099ec38f09a1745133c68a4d)
- [policy_based_challenge.always_enable_js_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-e8de5bd4e333592379aa9c51751add264265d3277e414e5b77568d04d6357119)
- [policy_based_challenge.captcha_challenge_parameters](resources--cdn_loadbalancer--reference--group-013.md#canonical-7b0045090e8c5ef02ef0ee49e64dc70034bfde0d7833823e9b5744869cc7527f)
- [policy_based_challenge.default_captcha_challenge_parameters](resources--cdn_loadbalancer--reference--group-013.md#canonical-d9f646b33f10838cb5c1fa84d31cba280e5e29994f263a1f481fb63fdca0c6a5)
- [policy_based_challenge.default_js_challenge_parameters](resources--cdn_loadbalancer--reference--group-013.md#canonical-692ec2ce0a9ecfaf67735387d16a2d6fe5151fe4cd40e59433e26c6985cb9611)
- [policy_based_challenge.default_mitigation_settings](resources--cdn_loadbalancer--reference--group-013.md#canonical-18b41c01fa4dc61af170cf40418a88eda5822bfc3294987633cb23c646109550)
- [policy_based_challenge.default_temporary_blocking_parameters](resources--cdn_loadbalancer--reference--group-013.md#canonical-8af097ae92a51a14440a8a17fb2c5207137124945ac0711b1367170897ddc838)
- [policy_based_challenge.js_challenge_parameters](resources--cdn_loadbalancer--reference--group-013.md#canonical-87630b046362400b830792b5610ed8697a78a4d5d3c8ad67ccefe199393f7342)
- [policy_based_challenge.malicious_user_mitigation](resources--cdn_loadbalancer--reference--group-013.md#canonical-a73000a20bec1df010b97ca8b8f4f3989e4041fc954594295ea8e4bd63edbfa6)
- [policy_based_challenge.no_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-e781909cfb9de5bc972cf81b5d973b59e0fb9cbd5403f19e92a44af189803430)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-e7e050a0106b5c243929bc99c717fe27390190239f5f0da6543c11a425e13c71)
- [policy_based_challenge.temporary_user_blocking](resources--cdn_loadbalancer--reference--group-014.md#canonical-885e305a1d0df444d63d64da6c7a4b0a54459b03c5eb86b881c6ad7b4cb37d06)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-ed41e8b257dba61d5faa50328ed39449fb350f60099ec38f09a1745133c68a4d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-16bc9fc66e7f17eb96fe350e1b33e6a4f8fdc3d89a3d8d687d2c020f0ebf1c87"></a>

## policy_based_challenge.always_enable_captcha_challenge — policy_based_challenge.always_enable_captcha_challenge / cc169eafc1d0 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- policy_based_challenge.always_enable_captcha_challenge

<a id="canonical-5d61433c99285626d91149a3311003940ee29c340cc72b4dc68a45bfccca1cd1"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for always enable captcha challenge.

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

Terraform syntax:

```terraform
always_enable_captcha_challenge = {}
```

<a id="canonical-3fd4f9345c9d9de3ae325447a965a66eebdcff17d90c30193f312e08b062b29d"></a>

## Direct properties — policy_based_challenge.always_enable_captcha_challenge / cc169eafc1d0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9a895e48e5bf0ce24101d12dfa97b5bbab8923a2e4a5c85d779f1767000d8a75"></a>

## Next pages — policy_based_challenge.always_enable_captcha_challenge / cc169eafc1d0 / 4

- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-e8de5bd4e333592379aa9c51751add264265d3277e414e5b77568d04d6357119"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-38ac3cc411e15e859b91fbb791c3928f8bc1c4c3f037af7158fb37e5ad5d2d84"></a>

## policy_based_challenge.always_enable_js_challenge — policy_based_challenge.always_enable_js_challenge / 92a91724a92b / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- policy_based_challenge.always_enable_js_challenge

<a id="canonical-198ed767d87febf8c79702b9bc1f1b9d72e2913483f229dbf035a2d424221e30"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for always enable js challenge.

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

Terraform syntax:

```terraform
always_enable_js_challenge = {}
```

<a id="canonical-cab766b1920dadd24c3e9467c304796c9ec5cf024bfc10e09789e98156fd542a"></a>

## Direct properties — policy_based_challenge.always_enable_js_challenge / 92a91724a92b / 3

This is an empty object or choice marker. It has no direct properties.
