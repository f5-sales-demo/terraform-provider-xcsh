---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-09ae151ee3fd673953923ba38f01cca4b9e8320c274835159d379a26b854c042"></a>

## Next pages — policy_based_challenge.always_enable_js_challenge / 92a91724a92b / 4

- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-7b0045090e8c5ef02ef0ee49e64dc70034bfde0d7833823e9b5744869cc7527f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dfe0de6056534f85a0e395f776d09f7342d4fde338aa7dfd9048cbc7af3afa4a"></a>

## policy_based_challenge.captcha_challenge_parameters — policy_based_challenge.captcha_challenge_parameters / 89db63f85f9e / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- policy_based_challenge.captcha_challenge_parameters

<a id="canonical-e47ea3e43339433168d057049199d5a34d098086bbc75280fc1aecbb7dbfbe45"></a>

Type: `"object"`. single nested block, Optional.

Enables loadbalancer to perform captcha challenge Captcha challenge will be based on Google
Recaptcha. With this feature enabled, only clients that pass the captcha challenge will be allowed
to complete the HTTP request. When loadbalancer is configured to do Captcha Challenge, it will
redirect..

Upstream description:

Enables loadbalancer to perform captcha challenge

Captcha challenge will be based on Google Recaptcha.

With this feature enabled, only clients that pass the captcha challenge will be allowed to complete
the HTTP request.

When loadbalancer is configured to do Captcha Challenge, it will redirect the browser to an HTML
page on every new HTTP request. This HTML page will have captcha challenge embedded in it. Client
will be allowed to make the request only if the captcha challenge is successful. Loadbalancer will
tag response header with a cookie to avoid Captcha challenge for subsequent requests.

CAPTCHA is mainly used as a security check to ensure only human users can pass through. Generally,
computers or bots are not capable of solving a captcha.

You can enable either Javascript challenge or Captcha challenge on a virtual host.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cookie_expiry")}
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
captcha_challenge_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-042018574273b56b777ce86fcf5e7b06c07f1f9f3bf7866e8799c6f951c10d9c"></a>

## Direct properties — policy_based_challenge.captcha_challenge_parameters / 89db63f85f9e / 3

<a id="canonical-ad9e90998480ba8dcd1a164b58e3983f06f7ef06cf62c2fd2baad93b05dea0da"></a>

<a id="canonical-437fd10dec31359b40a16b7a9f0e4f64b901fb508fe505582a1e00c70ffd2971"></a>

## cookie_expiry property — policy_based_challenge.captcha_challenge_parameters / 89db63f85f9e / 4

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

<a id="canonical-0702f7c60d65d9e55464370c3f6efbe163f2a9d55042a94ec90fc0e09b4a8372"></a>

<a id="canonical-f2bb5161e667f6bda49d20e2f3f1e84c8415871add6339f5a2e00cff36f98fe6"></a>

## custom_page property — policy_based_challenge.captcha_challenge_parameters / 89db63f85f9e / 5

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

<a id="canonical-461cfd6bebb5bf2315683381bf55447a26245a5835852b2c0fd63dfc904d4016"></a>

## Next pages — policy_based_challenge.captcha_challenge_parameters / 89db63f85f9e / 6

- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-d9f646b33f10838cb5c1fa84d31cba280e5e29994f263a1f481fb63fdca0c6a5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9f2f881fee6f7b4fb8c97a88e5387493ae985c2b57f3cafa68e487bc4c8f63a5"></a>

## policy_based_challenge.default_captcha_challenge_parameters — policy_based_challenge.default_captcha_challenge_parameters / bbf3cb25bb3e / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- policy_based_challenge.default_captcha_challenge_parameters

<a id="canonical-bf247701f376db5ca1edde592089c9cd341cc739898494101fab2f5d37e90cf2"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default captcha challenge parameters.

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
default_captcha_challenge_parameters = {}
```

<a id="canonical-6185a922c5ca28991366c292e626168f1500fbf01c3fe672714a8aa3e5e5ef6c"></a>

## Direct properties — policy_based_challenge.default_captcha_challenge_parameters / bbf3cb25bb3e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f3f44d9eb17e8852056096fd9d053a601aeb5075e7bcdd4341f8a451c17ad180"></a>

## Next pages — policy_based_challenge.default_captcha_challenge_parameters / bbf3cb25bb3e / 4

- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-692ec2ce0a9ecfaf67735387d16a2d6fe5151fe4cd40e59433e26c6985cb9611"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b1a6fe168679293f50ebb0258433068b81a3a157327b16f636d10f38706442df"></a>

## policy_based_challenge.default_js_challenge_parameters — policy_based_challenge.default_js_challenge_parameters / ff06ea388624 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- policy_based_challenge.default_js_challenge_parameters

<a id="canonical-1d5315ef290c29b2a8acb0696d1cd80c6e1fa5d7b93f6d7b24b95039e2451a76"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default js challenge parameters.

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
default_js_challenge_parameters = {}
```

<a id="canonical-46b42866f2041b964275eebdfb58d6143de8f7e4778512fd0c5a69b24dea941c"></a>

## Direct properties — policy_based_challenge.default_js_challenge_parameters / ff06ea388624 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-890db4aa0f6173d018caf9f6676ad53823df1333acd1bb8a2b38d8e614fd6729"></a>

## Next pages — policy_based_challenge.default_js_challenge_parameters / ff06ea388624 / 4

- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-18b41c01fa4dc61af170cf40418a88eda5822bfc3294987633cb23c646109550"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ff1efaea24a757625fb2f6131759285c3c577e3295daa35692715b9435bd5392"></a>

## policy_based_challenge.default_mitigation_settings — policy_based_challenge.default_mitigation_settings / 23a813024926 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- policy_based_challenge.default_mitigation_settings

<a id="canonical-0eac55676d979cbaca9b6a4faf3bd3cbecf2e7e3c6b3c770adc1a72715ac4f6e"></a>

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
default_mitigation_settings = {}
```

<a id="canonical-46312ebbe9f07a54f42d5533dfd23fae47011734c7c503c594c470c75058a0c3"></a>

## Direct properties — policy_based_challenge.default_mitigation_settings / 23a813024926 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-936ffdb898c063624e26742b66eaf823b135cd947eb61a3e4d3fce39f60b299f"></a>

## Next pages — policy_based_challenge.default_mitigation_settings / 23a813024926 / 4

- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-8af097ae92a51a14440a8a17fb2c5207137124945ac0711b1367170897ddc838"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3107e3ee363471f0308aaae0f9ce153158da93f9d3707f96da953c0889366463"></a>

## policy_based_challenge.default_temporary_blocking_parameters — policy_based_challenge.default_temporary_blocking_parameters / 782d4a9fd616 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- policy_based_challenge.default_temporary_blocking_parameters

<a id="canonical-ca54d14c379686acf148766b75c4958d88abdf180187739172b3286146b57ee0"></a>

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
default_temporary_blocking_parameters = {}
```

<a id="canonical-eef5af2c699dbcd1bcbbea9a0cccd6a27d17af5dd779146764749b8cdb1ec194"></a>

## Direct properties — policy_based_challenge.default_temporary_blocking_parameters / 782d4a9fd616 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2c55e1a3d24497920fb6f53ddbc25322eb9e6d2337e6e2d1d067258a190834ba"></a>

## Next pages — policy_based_challenge.default_temporary_blocking_parameters / 782d4a9fd616 / 4

- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-87630b046362400b830792b5610ed8697a78a4d5d3c8ad67ccefe199393f7342"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b5d0d3ca289a8f046a1fccf3b14008038dbe037ae8e7d78073b99eff0ff0352c"></a>

## policy_based_challenge.js_challenge_parameters — policy_based_challenge.js_challenge_parameters / 72e7aff531b2 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- policy_based_challenge.js_challenge_parameters

<a id="canonical-0f46f00600c478716cb1e5f356e040303d0fbf89842de9647ac70bbc3501f848"></a>

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
js_challenge_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-311378a43087b82e3952b75efa78eb3f0726fe111a42104e32af4996e63294e1"></a>

## Direct properties — policy_based_challenge.js_challenge_parameters / 72e7aff531b2 / 3

<a id="canonical-84246b3bfcaad3be7653c975edcc8c97750387c8941f0be084e86df9404cde35"></a>

<a id="canonical-ddc31455f00b58a879e096825b961e4ac3cba102eeddb215943326acf7eec031"></a>

## cookie_expiry property — policy_based_challenge.js_challenge_parameters / 72e7aff531b2 / 4

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

<a id="canonical-10bc557aae1d17797d12265126ff06aeb5b6908ef96e0ef036296fb13009368c"></a>

<a id="canonical-3e5f23ee99a871d275dd531d8eaa5339addfada2d59c7a7f2ab8c9a17d43fac9"></a>

## custom_page property — policy_based_challenge.js_challenge_parameters / 72e7aff531b2 / 5

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

<a id="canonical-efdd154fa961b7708e4efd3a4e732321cb8f20d5dfae41a7f9bacf2966ec6a05"></a>

<a id="canonical-c52ffd2c74150bad434f3ad7a812be7547256584524734b0c7b16c98e180e5e5"></a>

## js_script_delay property — policy_based_challenge.js_challenge_parameters / 72e7aff531b2 / 6

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

<a id="canonical-3c7f093120ba9ab5b38c60de46ffd959ab4330bfafda26eeb4124438305e3706"></a>

## Next pages — policy_based_challenge.js_challenge_parameters / 72e7aff531b2 / 7

- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-a73000a20bec1df010b97ca8b8f4f3989e4041fc954594295ea8e4bd63edbfa6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2eafd20ccf707904076f6a631bcfcfe4cb2ddb9ca1ea062975591e1de7d80930"></a>

## policy_based_challenge.malicious_user_mitigation — policy_based_challenge.malicious_user_mitigation / 35b25772e45b / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- policy_based_challenge.malicious_user_mitigation

<a id="canonical-b5cec6cbdef0a2167f89c37a4a60c361e30fc417d5f92cfb26fc38228ec659c2"></a>

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
malicious_user_mitigation {
  # Configure direct properties listed below.
}
```

<a id="canonical-8bb3c244e642b768c982ec9c7b830158167d3cca22dfc73952bc6258a64002ca"></a>

## Direct properties — policy_based_challenge.malicious_user_mitigation / 35b25772e45b / 3

<a id="canonical-248e8edb35652adf5076d7e6ea11425ff7b043b3adea2b0944366868e966bd0b"></a>

<a id="canonical-fd66af3c2c9467ebc766ca5299d010965913b41171d62309d17550bea1b6c455"></a>

## name property — policy_based_challenge.malicious_user_mitigation / 35b25772e45b / 4

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

<a id="canonical-c7ee6a7da4e924d1bcebaec92b47be4dad3a4a27a0f997574f883bca98379bdd"></a>

<a id="canonical-a3d1da8d46f816e1a645c4dbd8eee6d8057bbb3e2584cf432a059a4d84c0b645"></a>

## namespace property — policy_based_challenge.malicious_user_mitigation / 35b25772e45b / 5

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

<a id="canonical-93e40f459dc9912273a3d27ca7431ac728ed86fe918152f608b9e81b277f333c"></a>

<a id="canonical-6d60e4f442c399ff5a5d1ff1dd7245341ef9810c446c2034f598c11dc910cdf3"></a>

## tenant property — policy_based_challenge.malicious_user_mitigation / 35b25772e45b / 6

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

<a id="canonical-ac2b73439f71432e0c26779bf7a7c4531b592d91fb50bc544f54e626f588e9b5"></a>

## Next pages — policy_based_challenge.malicious_user_mitigation / 35b25772e45b / 7

- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-e781909cfb9de5bc972cf81b5d973b59e0fb9cbd5403f19e92a44af189803430"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b48ea46280beb1f392be67d936b98d08aab89fb86a5f531a2c7e6edfbada8ae2"></a>

## policy_based_challenge.no_challenge — policy_based_challenge.no_challenge / c4faa17b530e / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- policy_based_challenge.no_challenge

<a id="canonical-238b8756b2e29925e67783983855a11e3c54097e6451fbfba15dc4cc8d29269d"></a>

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

<a id="canonical-dc9e458f4dd4290f72a61fd255a24adb64d19f9e597296db05811cbf1cc28a3e"></a>

## Direct properties — policy_based_challenge.no_challenge / c4faa17b530e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1142a02cc234531f81cdffcb19d6c03051a91da8423674f0c2ec4c497ad97e0a"></a>

## Next pages — policy_based_challenge.no_challenge / c4faa17b530e / 4

- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-e7e050a0106b5c243929bc99c717fe27390190239f5f0da6543c11a425e13c71"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4cb2a15f9f0ccfa5b46080b0b00009d27037deb90b3b821b7267de5733a8c045"></a>

## policy_based_challenge.rule_list — policy_based_challenge.rule_list / ed1f3dec47ad / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- policy_based_challenge.rule_list

<a id="canonical-9babc10f232a9666970a0bbee5986aaf154437de8d0cd505efce4bc92563f4de"></a>

Type: `"object"`. single nested block, Optional.

List of challenge rules to be used in policy based challenge.

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
rule_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-52fdf52e82a29c7deea2da8bb25c973f5cb01069e1de8deec5470aaf106a093e"></a>

## Direct properties — policy_based_challenge.rule_list / ed1f3dec47ad / 3

- [rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-6b6b9bb417be0328d783de67330d570206c65ca27165054e0fe0e20523560695): complete subsection reference.

<a id="canonical-fdf491c69c969abced5800256a5bd4a781dcb67f3fcdcd808d8a91bec3e2b207"></a>

## Next pages — policy_based_challenge.rule_list / ed1f3dec47ad / 4

- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-6b6b9bb417be0328d783de67330d570206c65ca27165054e0fe0e20523560695)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-6b6b9bb417be0328d783de67330d570206c65ca27165054e0fe0e20523560695"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b17517ac1465f3237bdfd2974dbdcccfb7267f7d58fb6e59d56624705f3f7170"></a>

## policy_based_challenge.rule_list.rules — policy_based_challenge.rule_list.rules / 04a20dc66403 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-e7e050a0106b5c243929bc99c717fe27390190239f5f0da6543c11a425e13c71)
- policy_based_challenge.rule_list.rules

<a id="canonical-512652050ce52c93c303dd141b3db94c290ee2c5ac2e407bb3efacf38620747d"></a>

Type: `"object"`. list nested block, Optional.

Rules that specify the match conditions and challenge type to be launched. When a challenge type is
selected to be always enabled, these rules can be used to disable challenge or launch a different
challenge for requests that match the specified conditions.

Upstream description:

Rules that specify the match conditions and challenge type to be launched. When a challenge type is
selected to be always enabled, these rules can be used to disable challenge or launch a different
challenge for requests that match the specified conditions.

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
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-8fafec32c9427bbf64c1a5562db8055ae7bb2f23770d895ea9d0c4af32c70ad4"></a>

## Direct properties — policy_based_challenge.rule_list.rules / 04a20dc66403 / 3

- [metadata](resources--cdn_loadbalancer--reference--group-013.md#canonical-f6296636bc69e4410bd599071809303ea2409b8cdf9386f635ca20723a755306): complete subsection reference.

- [spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-c1694f8f786e00f1b8d062766cb92e3219566178dd23899b40bc7c9d82c517a1): complete subsection reference.

<a id="canonical-0bd85c5d802e9b9db3929624c6b0d0da94be19fac43d04228f96bda47847fce1"></a>

## Next pages — policy_based_challenge.rule_list.rules / 04a20dc66403 / 4

- [policy_based_challenge.rule_list.rules.metadata](resources--cdn_loadbalancer--reference--group-013.md#canonical-f6296636bc69e4410bd599071809303ea2409b8cdf9386f635ca20723a755306)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-c1694f8f786e00f1b8d062766cb92e3219566178dd23899b40bc7c9d82c517a1)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-e7e050a0106b5c243929bc99c717fe27390190239f5f0da6543c11a425e13c71)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-f6296636bc69e4410bd599071809303ea2409b8cdf9386f635ca20723a755306"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3e7102a30dd63485e8a00ea8671c51f43ba165f6b343b0dcdc26dd59c8acf609"></a>

## policy_based_challenge.rule_list.rules.metadata — policy_based_challenge.rule_list.rules.metadata / 3b76698c565b / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-e7e050a0106b5c243929bc99c717fe27390190239f5f0da6543c11a425e13c71)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-6b6b9bb417be0328d783de67330d570206c65ca27165054e0fe0e20523560695)
- policy_based_challenge.rule_list.rules.metadata

<a id="canonical-11fb60e827c6a75e713271f87584eca29706f74dd927ff92f1de73bdeaa20fb7"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-47b96ef14fd0081fd963193fe2d2c5ccbb63654a3f451fc0bf5034fc47681191"></a>

## Direct properties — policy_based_challenge.rule_list.rules.metadata / 3b76698c565b / 3

<a id="canonical-7f24fe29d849b9bbc36936a6680ef1c2513941ced339d4c20cce2952ee95761c"></a>

<a id="canonical-31b9f4dfb07c8de660d00586de4b97dff4938716312d942cde7b00d755df6601"></a>

## description_spec property — policy_based_challenge.rule_list.rules.metadata / 3b76698c565b / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-bfc2534a6ac3445ac05c0ee7e6a3275a1d33527f0f3ac5d81d79d2b3f778c4d5"></a>

<a id="canonical-06a01db84e8d9f35d879b2edfd9fd2ed08fbd13673f72d57551043008a524eb9"></a>

## name property — policy_based_challenge.rule_list.rules.metadata / 3b76698c565b / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
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
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-37d3e354b4a2ad763308181cc0b2db6e3587838f36385d87fc6b7cb6cea8510f"></a>

## Next pages — policy_based_challenge.rule_list.rules.metadata / 3b76698c565b / 6

- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-6b6b9bb417be0328d783de67330d570206c65ca27165054e0fe0e20523560695)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-c1694f8f786e00f1b8d062766cb92e3219566178dd23899b40bc7c9d82c517a1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-44ec89bce00471be782f11db7dc32227fd031a87a54ecc960b158b4e79403408"></a>

## policy_based_challenge.rule_list.rules.spec — policy_based_challenge.rule_list.rules.spec / d1831211a9e9 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-e7e050a0106b5c243929bc99c717fe27390190239f5f0da6543c11a425e13c71)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-6b6b9bb417be0328d783de67330d570206c65ca27165054e0fe0e20523560695)
- policy_based_challenge.rule_list.rules.spec

<a id="canonical-902b63b28eeea034f823becc35d6cc83231f3c5f730f10cc2fabcdc04620c873"></a>

Type: `"object"`. single nested block, Optional.

Challenge Rule consists of an unordered list of predicates and an action. The predicates are
evaluated against a set of input fields that are extracted from or derived from an L7 request API. A
request API is considered to match the rule if all predicates in the rule evaluate to true for
that..

Upstream description:

A Challenge Rule consists of an unordered list of predicates and an action. The predicates are
evaluated against a set of input fields that are extracted from or derived from an L7 request API. A
request API is considered to match the rule if all predicates in the rule evaluate to true for that
request. Any predicates that are not specified in a rule are implicitly considered to be true. If a
request API matches a challenge rule, the configured challenge is enforced.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("any_asn",
    "asn_list"),
  validators.ConflictingObjectAttributes("any_asn",
    "asn_matcher"),
  validators.ConflictingObjectAttributes("any_client",
    "client_selector"),
  validators.ConflictingObjectAttributes("any_ip",
    "ip_matcher"),
  validators.ConflictingObjectAttributes("any_ip",
    "ip_prefix_list"),
  validators.ConflictingObjectAttributes("asn_list",
    "asn_matcher"),
  validators.ConflictingObjectAttributes("disable_challenge",
    "enable_captcha_challenge"),
  validators.ConflictingObjectAttributes("disable_challenge",
    "enable_javascript_challenge"),
  validators.ConflictingObjectAttributes("enable_captcha_challenge",
    "enable_javascript_challenge"),
  validators.ConflictingObjectAttributes("ip_matcher",
    "ip_prefix_list")}
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
  "x-ves-oneof-field-asn_choice": "[\"any_asn\",\"asn_list\",\"asn_matcher\"]",
  "x-ves-oneof-field-challenge_action": "[\"disable_challenge\",\"enable_captcha_challenge\",\"enable_javascript_challenge\"]",
  "x-ves-oneof-field-client_choice": "[\"any_client\",\"client_selector\"]",
  "x-ves-oneof-field-ip_choice": "[\"any_ip\",\"ip_matcher\",\"ip_prefix_list\"]",
  "x-ves-oneof-field-tls_fingerprint_choice": "[\"tls_fingerprint_matcher\"]"
}
```

Terraform syntax:

```terraform
spec {
  # Configure direct properties listed below.
}
```

<a id="canonical-0cb1dd7188a71d33e6ffad2230ac2f6c8e8da054d22b5d9c6427921500040dee"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec / d1831211a9e9 / 3

- [any_asn](resources--cdn_loadbalancer--reference--group-013.md#canonical-d9a973ab70a1795ad2cc596d01451640661db33454ca4eae9682a3aee47222c0): complete subsection reference.

- [any_client](resources--cdn_loadbalancer--reference--group-013.md#canonical-7206887c7dbab891a8c0f52f7cbfc42b6d5d61bc0ec9a2c1620e16371bb57130): complete subsection reference.

- [any_ip](resources--cdn_loadbalancer--reference--group-013.md#canonical-d4a4a7eeb03bf122911fb075de8bfe6a31f2ca809a67a17166210b326f8fbdd1): complete subsection reference.

- [arg_matchers](resources--cdn_loadbalancer--reference--group-013.md#canonical-91547eacfa98c58e1043609f628832ee32fc5235cd79582a09576709331a83d0): complete subsection reference.

- [asn_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-03bebc21360864c8730a328b95bedfd640ed0706789c184f62fcb03247eaf576): complete subsection reference.

- [asn_matcher](resources--cdn_loadbalancer--reference--group-013.md#canonical-0349e2c4ed40628b09167941f918ef8b690da59403f4f4194e0884b053f26411): complete subsection reference.

- [body_matcher](resources--cdn_loadbalancer--reference--group-013.md#canonical-08d06452c0b0e58e1e5141f6483bde76637e6f5c1be0a0c898a76638e99823c4): complete subsection reference.

- [client_selector](resources--cdn_loadbalancer--reference--group-013.md#canonical-ec7c39dfc583277454d8fc3ff5a931696d8107535dae7ea3f2297140818d1b9e): complete subsection reference.

- [cookie_matchers](resources--cdn_loadbalancer--reference--group-013.md#canonical-f7820fa160919665e80f0206825016d7ef69fa7426efa4747f1bd99a82c18fdc): complete subsection reference.

- [disable_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-0ef8295b621b356b70726ee6ddbf291d8ddb571073af277b924b4c59840e39e9): complete subsection reference.

- [domain_matcher](resources--cdn_loadbalancer--reference--group-013.md#canonical-4df3dac25f5add58043c4a45e0bb68c779e67f54e56313602e09305a6f14e67f): complete subsection reference.

- [enable_captcha_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-d3d73cf1d86a379d97a04d2ce491647779e67beb9c7eca6beb3fc6cc6213b0f8): complete subsection reference.

- [enable_javascript_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-51e7028886650667705be61ff103bbd51fc27e835c8c9be24db09884020dbf44): complete subsection reference.

<a id="canonical-9a819e1de160ba466577119807392eba9b15b72715aef0c97b0a9af76a20ed42"></a>

<a id="canonical-e31c5aedfaba45e6990a2d2499e17a62618b12ae7b51318f1d490d28aee0577a"></a>

## expiration_timestamp property — policy_based_challenge.rule_list.rules.spec / d1831211a9e9 / 4

Type: `"string"`. Optional.

Specifies expiration\_timestamp the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Upstream description:

The expiration\_timestamp is the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "date-time",
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

- [headers](resources--cdn_loadbalancer--reference--group-013.md#canonical-fc8041ae07d56ba34ef24020aa6a5accfa48b0ec10c1362557b6eeff2a4f3d88): complete subsection reference.

- [http_method](resources--cdn_loadbalancer--reference--group-013.md#canonical-b6776a702c1500c90976697d1bffab0f64d0ecfa783f880ca9faa078bcd35243): complete subsection reference.

- [ip_matcher](resources--cdn_loadbalancer--reference--group-013.md#canonical-698f05ca9efe080743409822581d5ad76e0d8fd235fb254c80eccfbcca391788): complete subsection reference.

- [ip_prefix_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-463f810aaab48ed3b9009874cd2f5f4a0608b779ac170ed0dcd87ab8a766b824): complete subsection reference.

- [path](resources--cdn_loadbalancer--reference--group-013.md#canonical-7b35f428c37707746358a0356c25eec89e66574da2d70f66d5b38214c515df95): complete subsection reference.

- [query_params](resources--cdn_loadbalancer--reference--group-013.md#canonical-511830efa2a971e782cd6ba252499d4cf40f63d59a6550161b659d22def20550): complete subsection reference.

- [tls_fingerprint_matcher](resources--cdn_loadbalancer--reference--group-014.md#canonical-ee283ca70e76f514eda54bce44c1789446a7a62c7a862a69afffcfe0cce80cfd): complete subsection reference.

<a id="canonical-47ad5da4f108e0ac31f3da0af5ac63827e8daba3da215f83b6f7d3d002427683"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec / d1831211a9e9 / 5

- [policy_based_challenge.rule_list.rules.spec.any_asn](resources--cdn_loadbalancer--reference--group-013.md#canonical-d9a973ab70a1795ad2cc596d01451640661db33454ca4eae9682a3aee47222c0)
- [policy_based_challenge.rule_list.rules.spec.any_client](resources--cdn_loadbalancer--reference--group-013.md#canonical-7206887c7dbab891a8c0f52f7cbfc42b6d5d61bc0ec9a2c1620e16371bb57130)
- [policy_based_challenge.rule_list.rules.spec.any_ip](resources--cdn_loadbalancer--reference--group-013.md#canonical-d4a4a7eeb03bf122911fb075de8bfe6a31f2ca809a67a17166210b326f8fbdd1)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers](resources--cdn_loadbalancer--reference--group-013.md#canonical-91547eacfa98c58e1043609f628832ee32fc5235cd79582a09576709331a83d0)
- [policy_based_challenge.rule_list.rules.spec.asn_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-03bebc21360864c8730a328b95bedfd640ed0706789c184f62fcb03247eaf576)
- [policy_based_challenge.rule_list.rules.spec.asn_matcher](resources--cdn_loadbalancer--reference--group-013.md#canonical-0349e2c4ed40628b09167941f918ef8b690da59403f4f4194e0884b053f26411)
- [policy_based_challenge.rule_list.rules.spec.body_matcher](resources--cdn_loadbalancer--reference--group-013.md#canonical-08d06452c0b0e58e1e5141f6483bde76637e6f5c1be0a0c898a76638e99823c4)
- [policy_based_challenge.rule_list.rules.spec.client_selector](resources--cdn_loadbalancer--reference--group-013.md#canonical-ec7c39dfc583277454d8fc3ff5a931696d8107535dae7ea3f2297140818d1b9e)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](resources--cdn_loadbalancer--reference--group-013.md#canonical-f7820fa160919665e80f0206825016d7ef69fa7426efa4747f1bd99a82c18fdc)
- [policy_based_challenge.rule_list.rules.spec.disable_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-0ef8295b621b356b70726ee6ddbf291d8ddb571073af277b924b4c59840e39e9)
- [policy_based_challenge.rule_list.rules.spec.domain_matcher](resources--cdn_loadbalancer--reference--group-013.md#canonical-4df3dac25f5add58043c4a45e0bb68c779e67f54e56313602e09305a6f14e67f)
- [policy_based_challenge.rule_list.rules.spec.enable_captcha_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-d3d73cf1d86a379d97a04d2ce491647779e67beb9c7eca6beb3fc6cc6213b0f8)
- [policy_based_challenge.rule_list.rules.spec.enable_javascript_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-51e7028886650667705be61ff103bbd51fc27e835c8c9be24db09884020dbf44)
- [policy_based_challenge.rule_list.rules.spec.headers](resources--cdn_loadbalancer--reference--group-013.md#canonical-fc8041ae07d56ba34ef24020aa6a5accfa48b0ec10c1362557b6eeff2a4f3d88)
- [policy_based_challenge.rule_list.rules.spec.http_method](resources--cdn_loadbalancer--reference--group-013.md#canonical-b6776a702c1500c90976697d1bffab0f64d0ecfa783f880ca9faa078bcd35243)
- [policy_based_challenge.rule_list.rules.spec.ip_matcher](resources--cdn_loadbalancer--reference--group-013.md#canonical-698f05ca9efe080743409822581d5ad76e0d8fd235fb254c80eccfbcca391788)
- [policy_based_challenge.rule_list.rules.spec.ip_prefix_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-463f810aaab48ed3b9009874cd2f5f4a0608b779ac170ed0dcd87ab8a766b824)
- [policy_based_challenge.rule_list.rules.spec.path](resources--cdn_loadbalancer--reference--group-013.md#canonical-7b35f428c37707746358a0356c25eec89e66574da2d70f66d5b38214c515df95)
- [policy_based_challenge.rule_list.rules.spec.query_params](resources--cdn_loadbalancer--reference--group-013.md#canonical-511830efa2a971e782cd6ba252499d4cf40f63d59a6550161b659d22def20550)
- [policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher](resources--cdn_loadbalancer--reference--group-014.md#canonical-ee283ca70e76f514eda54bce44c1789446a7a62c7a862a69afffcfe0cce80cfd)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-6b6b9bb417be0328d783de67330d570206c65ca27165054e0fe0e20523560695)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-d9a973ab70a1795ad2cc596d01451640661db33454ca4eae9682a3aee47222c0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2397559f178e632613b2359f88b74fc10413ec372e98b1d3b8c8b53b09e64ff3"></a>

## policy_based_challenge.rule_list.rules.spec.any_asn — policy_based_challenge.rule_list.rules.spec.any_asn / 8f27fb7c478c / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-e7e050a0106b5c243929bc99c717fe27390190239f5f0da6543c11a425e13c71)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-6b6b9bb417be0328d783de67330d570206c65ca27165054e0fe0e20523560695)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-c1694f8f786e00f1b8d062766cb92e3219566178dd23899b40bc7c9d82c517a1)
- policy_based_challenge.rule_list.rules.spec.any_asn

<a id="canonical-d6b7ded9f82a97c7dda441f0a5e30b1587039ff85578f040e1c63d60afc01088"></a>

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
any_asn = {}
```

<a id="canonical-f26f8671302f605dd0dd63f160b7f4a0c85afd22f78d6170cb514aaf680e2e00"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.any_asn / 8f27fb7c478c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b7cd9cf2d7362072bacdb9936db79dcd30652ac7d0507372680b9706c55487c0"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.any_asn / 8f27fb7c478c / 4

- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-c1694f8f786e00f1b8d062766cb92e3219566178dd23899b40bc7c9d82c517a1)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-7206887c7dbab891a8c0f52f7cbfc42b6d5d61bc0ec9a2c1620e16371bb57130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-920bd6a138438c59847cfd2ebcd8c0e021a04779b2cdf4fdd3861ba941a3e7f8"></a>

## policy_based_challenge.rule_list.rules.spec.any_client — policy_based_challenge.rule_list.rules.spec.any_client / 36228a21d03a / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-e7e050a0106b5c243929bc99c717fe27390190239f5f0da6543c11a425e13c71)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-6b6b9bb417be0328d783de67330d570206c65ca27165054e0fe0e20523560695)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-c1694f8f786e00f1b8d062766cb92e3219566178dd23899b40bc7c9d82c517a1)
- policy_based_challenge.rule_list.rules.spec.any_client

<a id="canonical-16995b7d9d1331f28c52541be71ecf301fd70fc0c788f860e61cd493b99eb8a6"></a>

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
any_client = {}
```

<a id="canonical-ba33ca50ab6d478025f2a9c402418ca53463d6ec86bcb97ad1010b3f2ba3450c"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.any_client / 36228a21d03a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2ae7d5d64b204db50a3e8f79b2bf181a9fa444e30f4bf8a853a5601d50b90d79"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.any_client / 36228a21d03a / 4

- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-c1694f8f786e00f1b8d062766cb92e3219566178dd23899b40bc7c9d82c517a1)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-d4a4a7eeb03bf122911fb075de8bfe6a31f2ca809a67a17166210b326f8fbdd1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1962207b07899bf417f1a0b0a9a0d50eccc2b094df13b0c3c8b380a165133fd2"></a>

## policy_based_challenge.rule_list.rules.spec.any_ip — policy_based_challenge.rule_list.rules.spec.any_ip / 2b8e5d62df23 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-e7e050a0106b5c243929bc99c717fe27390190239f5f0da6543c11a425e13c71)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-6b6b9bb417be0328d783de67330d570206c65ca27165054e0fe0e20523560695)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-c1694f8f786e00f1b8d062766cb92e3219566178dd23899b40bc7c9d82c517a1)
- policy_based_challenge.rule_list.rules.spec.any_ip

<a id="canonical-d4b37ce8ce4ecdbcad1c31c38ffa309734832a1b206ee979d8f6c6a4c0b935c9"></a>

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
any_ip = {}
```

<a id="canonical-7e390a685e717455dcf08722ceb8e9f6cf51e93f65aa9d9b87212f959f829864"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.any_ip / 2b8e5d62df23 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-aa62aec10f1b657d80418cadeedcdca8f7673290c057a92617337ea98c5c96a8"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.any_ip / 2b8e5d62df23 / 4

- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-c1694f8f786e00f1b8d062766cb92e3219566178dd23899b40bc7c9d82c517a1)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-91547eacfa98c58e1043609f628832ee32fc5235cd79582a09576709331a83d0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-318bdecf2e79a88c97830098f1a06ce0299a27c7ad391ec02f772bc0c52db670"></a>

## policy_based_challenge.rule_list.rules.spec.arg_matchers — policy_based_challenge.rule_list.rules.spec.arg_matchers / 540eb1344b07 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-e7e050a0106b5c243929bc99c717fe27390190239f5f0da6543c11a425e13c71)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-6b6b9bb417be0328d783de67330d570206c65ca27165054e0fe0e20523560695)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-c1694f8f786e00f1b8d062766cb92e3219566178dd23899b40bc7c9d82c517a1)
- policy_based_challenge.rule_list.rules.spec.arg_matchers

<a id="canonical-62f7f88b3df278768eee21d93cd7783c61fc71883f5a752579aff048e6061d00"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for all POST args that need to be matched. The criteria for matching each arg are
described in individual instances of ArgMatcherType. The actual arg values are extracted from the
request API as a list of strings for each arg selector name.

Upstream description:

A list of predicates for all POST args that need to be matched. The criteria for matching each arg
are described in individual instances of ArgMatcherType. The actual arg values are extracted from
the request API as a list of strings for each arg selector name. Note that all specified arg matcher
predicates must evaluate to true. A request body greater than 64KB will not be evaluated.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
```

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
arg_matchers {
  # Configure direct properties listed below.
}
```

<a id="canonical-1ed9a1413917af2f516da4406f8e77f03b2ed2d14493c3f1ed2ede470654ea7a"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.arg_matchers / 540eb1344b07 / 3

- [check_not_present](resources--cdn_loadbalancer--reference--group-013.md#canonical-47eee7c17f2e60904a7c8afd97ca823cef2c12400ae5895375e0653689f8935d): complete subsection reference.

- [check_present](resources--cdn_loadbalancer--reference--group-013.md#canonical-6465f3918f538ced096cd89948af76f7c7b87e835937872314ab693f499dc87f): complete subsection reference.

<a id="canonical-4830d51706f352c5c57e839a87ed6a2aa76be5dd44abf1a53bc080074c62ef45"></a>

<a id="canonical-fdd58d6500b722eb025e697c34023367ba752a9c2e14f7d9c1be327928b4ff0b"></a>

## invert_matcher property — policy_based_challenge.rule_list.rules.spec.arg_matchers / 540eb1344b07 / 4

Type: `"bool"`. Optional.

Invert Matcher. Invert Match of the expression defined.

Upstream description:

Invert Match of the expression defined.

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

- [item](resources--cdn_loadbalancer--reference--group-013.md#canonical-1cad7de35f24c06bd5433703dc43433a634de49991b126cf4a6b783097deded7): complete subsection reference.

<a id="canonical-94f469d7e69decf7217e7017866a3f97f3735c2ed5b0a737bb76623e30a9d2b6"></a>

<a id="canonical-bd024f576ce820dea9fa9e259265bed61de07c437b0ecd52736242a68b3e0c2d"></a>

## name property — policy_based_challenge.rule_list.rules.spec.arg_matchers / 540eb1344b07 / 5

Type: `"string"`. Optional.

Case-sensitive JSON path in the HTTP request body.

Upstream description:

A case-sensitive JSON path in the HTTP request body.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
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
    "ves.io.schema.rules.string.json_path": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.json_path": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-0f2342b61aec29c32c4981fe7a4ff6996308889d461237ed83fc130703f86885"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.arg_matchers / 540eb1344b07 / 6

- [policy_based_challenge.rule_list.rules.spec.arg_matchers.check_not_present](resources--cdn_loadbalancer--reference--group-013.md#canonical-47eee7c17f2e60904a7c8afd97ca823cef2c12400ae5895375e0653689f8935d)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers.check_present](resources--cdn_loadbalancer--reference--group-013.md#canonical-6465f3918f538ced096cd89948af76f7c7b87e835937872314ab693f499dc87f)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers.item](resources--cdn_loadbalancer--reference--group-013.md#canonical-1cad7de35f24c06bd5433703dc43433a634de49991b126cf4a6b783097deded7)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-c1694f8f786e00f1b8d062766cb92e3219566178dd23899b40bc7c9d82c517a1)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-47eee7c17f2e60904a7c8afd97ca823cef2c12400ae5895375e0653689f8935d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0c7993a1acb078a7ad92c15269d7a4a8318c55dfdfe9b6287c7ea94848b0ecab"></a>

## policy_based_challenge.rule_list.rules.spec.arg_matchers.check_not_present — policy_based_challenge.rule_list.rules.spec.arg_matchers.check_not_present / 3058fcbbb485 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-e7e050a0106b5c243929bc99c717fe27390190239f5f0da6543c11a425e13c71)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-6b6b9bb417be0328d783de67330d570206c65ca27165054e0fe0e20523560695)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-c1694f8f786e00f1b8d062766cb92e3219566178dd23899b40bc7c9d82c517a1)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers](resources--cdn_loadbalancer--reference--group-013.md#canonical-91547eacfa98c58e1043609f628832ee32fc5235cd79582a09576709331a83d0)
- policy_based_challenge.rule_list.rules.spec.arg_matchers.check_not_present

<a id="canonical-87cc531498517aad44fb35b7aa4ac56838105173d3426a48410b5201c933f015"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

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
check_not_present = {}
```

<a id="canonical-c83db09d3a5f05022157e7da31be83ec45656242bae51edd8e0046dbe0feb4bc"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.arg_matchers.check_not_present / 3058fcbbb485 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e4f4d44866a6c70e6ef3ccbbf03d285481f77ad3d2019fb2f51be1ecdf8489bb"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.arg_matchers.check_not_present / 3058fcbbb485 / 4

- [policy_based_challenge.rule_list.rules.spec.arg_matchers](resources--cdn_loadbalancer--reference--group-013.md#canonical-91547eacfa98c58e1043609f628832ee32fc5235cd79582a09576709331a83d0)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-6465f3918f538ced096cd89948af76f7c7b87e835937872314ab693f499dc87f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aefa56275f3a13a1181ec82057467a0f0aedf6bdbb6e8ff67b29e87f2469f15c"></a>

## policy_based_challenge.rule_list.rules.spec.arg_matchers.check_present — policy_based_challenge.rule_list.rules.spec.arg_matchers.check_present / 5a9b930dfa9d / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-e7e050a0106b5c243929bc99c717fe27390190239f5f0da6543c11a425e13c71)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-6b6b9bb417be0328d783de67330d570206c65ca27165054e0fe0e20523560695)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-c1694f8f786e00f1b8d062766cb92e3219566178dd23899b40bc7c9d82c517a1)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers](resources--cdn_loadbalancer--reference--group-013.md#canonical-91547eacfa98c58e1043609f628832ee32fc5235cd79582a09576709331a83d0)
- policy_based_challenge.rule_list.rules.spec.arg_matchers.check_present

<a id="canonical-4e2547b57bc923d50a4717312fd889c7fbfa6092799e0e471053f062d1cd4633"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

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
check_present = {}
```

<a id="canonical-3aea06bd72a117c9493133bbce9c5a778cce1d1e6a39780c30a6591732a6ab69"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.arg_matchers.check_present / 5a9b930dfa9d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-711a74265ea7aec007d1198a8be3a8fb09b387c8b1074e0a2e0b55e93d1f8a06"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.arg_matchers.check_present / 5a9b930dfa9d / 4

- [policy_based_challenge.rule_list.rules.spec.arg_matchers](resources--cdn_loadbalancer--reference--group-013.md#canonical-91547eacfa98c58e1043609f628832ee32fc5235cd79582a09576709331a83d0)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-1cad7de35f24c06bd5433703dc43433a634de49991b126cf4a6b783097deded7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d23dc362046e25a72bbf356e3ee92fe15f33cf2f1c40ad60d1fa84754548cb21"></a>

## policy_based_challenge.rule_list.rules.spec.arg_matchers.item — policy_based_challenge.rule_list.rules.spec.arg_matchers.item / 980fc6b3e990 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-e7e050a0106b5c243929bc99c717fe27390190239f5f0da6543c11a425e13c71)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-6b6b9bb417be0328d783de67330d570206c65ca27165054e0fe0e20523560695)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-c1694f8f786e00f1b8d062766cb92e3219566178dd23899b40bc7c9d82c517a1)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers](resources--cdn_loadbalancer--reference--group-013.md#canonical-91547eacfa98c58e1043609f628832ee32fc5235cd79582a09576709331a83d0)
- policy_based_challenge.rule_list.rules.spec.arg_matchers.item

<a id="canonical-b91dc41fd09651bb7c9e7d2ee587bd049ed4179e13dcf4b04f612abe28f0eb40"></a>

Type: `"object"`. single nested block, Optional.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-57122153eeb0714c2f6f80d160f95e5019b88fd81b2102279db6e4e597c5a279"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.arg_matchers.item / 980fc6b3e990 / 3

<a id="canonical-a73616564c9cb89a3fc7031083b1d5f3a76896652af0b52e40f1a695559cc48f"></a>

<a id="canonical-cbd613e3aace9107f02eae3bcd21f8dcf6b57e73ac2a9931664708e900f58110"></a>

## exact_values property — policy_based_challenge.rule_list.rules.spec.arg_matchers.item / 980fc6b3e990 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-7a716825bb3b0f16cf1f6a4bc7c240f8446962291146829012a4178886e7b096"></a>

<a id="canonical-023a385df9ce09513888eeb2989dcb48973d78131c921840c3a1a913c3e84d10"></a>

## regex_values property — policy_based_challenge.rule_list.rules.spec.arg_matchers.item / 980fc6b3e990 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-95381afe7d3d2ec18af61e5e7dff8d2351cd3809c53078464b343f4d79b5e1e7"></a>

<a id="canonical-e0548a72d08771c229374dc07eac66a74467da1629b358b5bf8730b326d9fdec"></a>

## transformers property — policy_based_challenge.rule_list.rules.spec.arg_matchers.item / 980fc6b3e990 / 6

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-a8bf6aa21fba94a86e05e53a3caa31b51be576ab971eb97b1dfe9295b787b04e"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.arg_matchers.item / 980fc6b3e990 / 7

- [policy_based_challenge.rule_list.rules.spec.arg_matchers](resources--cdn_loadbalancer--reference--group-013.md#canonical-91547eacfa98c58e1043609f628832ee32fc5235cd79582a09576709331a83d0)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-03bebc21360864c8730a328b95bedfd640ed0706789c184f62fcb03247eaf576"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-22f99676d5043fec3afc64ea9a3b016e404374d59c3eb0bd7a653e7f3adad89a"></a>

## policy_based_challenge.rule_list.rules.spec.asn_list — policy_based_challenge.rule_list.rules.spec.asn_list / 0ca2c7188e3c / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-e7e050a0106b5c243929bc99c717fe27390190239f5f0da6543c11a425e13c71)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-6b6b9bb417be0328d783de67330d570206c65ca27165054e0fe0e20523560695)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-c1694f8f786e00f1b8d062766cb92e3219566178dd23899b40bc7c9d82c517a1)
- policy_based_challenge.rule_list.rules.spec.asn_list

<a id="canonical-e308992a234f2050373eca092e18a63fd3d1c8210e4b127e4caaf3a89fb64568"></a>

Type: `"object"`. single nested block, Optional.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("as_numbers")}
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
asn_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-05337eb6001a35c814cb6ecf918cd9000b7a5337991eccf8fb225fd627b5751c"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.asn_list / 0ca2c7188e3c / 3

<a id="canonical-7bb8bc86b0e1822e9343db16beaeb93b0e344b92b2abf8e367ce5059d88645eb"></a>

<a id="canonical-637083a0358a5050d9eb511fb7bfb225af17c21159563bcef81011e45444d4db"></a>

## as_numbers property — policy_based_challenge.rule_list.rules.spec.asn_list / 0ca2c7188e3c / 4

Type: `["list", "number"]`. Optional.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 16),
}
```

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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-57e5613019e84015af48261cf1ebfc3feb13ab309826c4ade9f40b1e7d3ea5bf"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.asn_list / 0ca2c7188e3c / 5

- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-c1694f8f786e00f1b8d062766cb92e3219566178dd23899b40bc7c9d82c517a1)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-0349e2c4ed40628b09167941f918ef8b690da59403f4f4194e0884b053f26411"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-012085d8c32d2b7813f0b7c1e8fe0ae53408b6184c2d68fa651fddfc89ff9678"></a>

## policy_based_challenge.rule_list.rules.spec.asn_matcher — policy_based_challenge.rule_list.rules.spec.asn_matcher / eda4d9ab3224 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-e7e050a0106b5c243929bc99c717fe27390190239f5f0da6543c11a425e13c71)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-6b6b9bb417be0328d783de67330d570206c65ca27165054e0fe0e20523560695)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-c1694f8f786e00f1b8d062766cb92e3219566178dd23899b40bc7c9d82c517a1)
- policy_based_challenge.rule_list.rules.spec.asn_matcher

<a id="canonical-edf643ad9bd997b63b9c1f9f05d98c0ba556fe32724ec11eb7771f6a30ee3dac"></a>

Type: `"object"`. single nested block, Optional.

Match any AS number contained in the list of bgp\_asn\_sets.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("asn_sets")}
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
asn_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-97c8b6f468f5e13d2ed595443ad1cc62c65d8e7df3a1f17af1cb8a47035a65f3"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.asn_matcher / eda4d9ab3224 / 3

- [asn_sets](resources--cdn_loadbalancer--reference--group-013.md#canonical-0330773cbd5de417df1ce551bc56293aa6a8eaa2aba985626879acc4265b3a9c): complete subsection reference.

<a id="canonical-73fe787bccb46062ad91439bbec450ef997524a253c770d0a9c09346e33b85a6"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.asn_matcher / eda4d9ab3224 / 4

- [policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets](resources--cdn_loadbalancer--reference--group-013.md#canonical-0330773cbd5de417df1ce551bc56293aa6a8eaa2aba985626879acc4265b3a9c)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-c1694f8f786e00f1b8d062766cb92e3219566178dd23899b40bc7c9d82c517a1)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-0330773cbd5de417df1ce551bc56293aa6a8eaa2aba985626879acc4265b3a9c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-742a7f77856abd32ae6b1e91137152ebd030e14229c5493362aa3c97b6dcf9db"></a>

## policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets — policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets / 5fbcc93b3365 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-e7e050a0106b5c243929bc99c717fe27390190239f5f0da6543c11a425e13c71)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-6b6b9bb417be0328d783de67330d570206c65ca27165054e0fe0e20523560695)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-c1694f8f786e00f1b8d062766cb92e3219566178dd23899b40bc7c9d82c517a1)
- [policy_based_challenge.rule_list.rules.spec.asn_matcher](resources--cdn_loadbalancer--reference--group-013.md#canonical-0349e2c4ed40628b09167941f918ef8b690da59403f4f4194e0884b053f26411)
- policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets

<a id="canonical-591a56802ff826670238501d0082b795786416f04a6554cadae34ec1ac850dd2"></a>

Type: `"object"`. list nested block, Optional.

List of references to bgp\_asn\_set objects.

Upstream description:

A list of references to bgp\_asn\_set objects.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

Terraform syntax:

```terraform
asn_sets {
  # Configure direct properties listed below.
}
```

<a id="canonical-aeb1b93b486888c7b97a08d6d83445d7c81676df174a735418df144585e7bd97"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets / 5fbcc93b3365 / 3

<a id="canonical-1c3e30fdd5385b3eddf09ae30f28127ff37e698ca43a874aac00967cbc159d83"></a>

<a id="canonical-5dd9057757d2053de32f4683117cdac7a124b02723924dcba5f7e3c7a8c341e3"></a>

## kind property — policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets / 5fbcc93b3365 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-d549a925aea644b5406587e503f4396c560df8dd2ce1e7adaaf5b84804997c61"></a>

<a id="canonical-159c7748b731ed9ddb9ae5c03caf854c6c5873b5a842114791d779089dcfc541"></a>

## name property — policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets / 5fbcc93b3365 / 5

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-292ac99a268a3a46378cd1ca0cb918b7972045f1d5e502f89a8c6821afba7e8b"></a>

<a id="canonical-07323562d000c7f047f9841a08e3a4c074afe7eaab1cc0a1a3a3115b52e5111c"></a>

## namespace property — policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets / 5fbcc93b3365 / 6

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
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
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
  }
}
```

<a id="canonical-f3116458746676a2168d303c33a58add7be93a336865d6d694c513a7f5d53062"></a>

<a id="canonical-c8c5d7d4789c51a71854dedbe928af62451ab7a4457232b3da3c58e208406ded"></a>

## tenant property — policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets / 5fbcc93b3365 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-3754d28bdfc11e414b9f245b95b0560183c96914b06d7f244951f72ca84a3b62"></a>

<a id="canonical-38cefaf2ae7833c933f0b93cdf48bd6993be3a4c792123686cc389cb9a441e8c"></a>

## uid property — policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets / 5fbcc93b3365 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-e630122a32a82aa967ea094e78a7e6b6124b68874dbe164256caeebcfa20a40a"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets / 5fbcc93b3365 / 9

- [policy_based_challenge.rule_list.rules.spec.asn_matcher](resources--cdn_loadbalancer--reference--group-013.md#canonical-0349e2c4ed40628b09167941f918ef8b690da59403f4f4194e0884b053f26411)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-08d06452c0b0e58e1e5141f6483bde76637e6f5c1be0a0c898a76638e99823c4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cf184a141229ef8074c6327664632d6f668953813d3e33cb789dc3370a0b4c2f"></a>

## policy_based_challenge.rule_list.rules.spec.body_matcher — policy_based_challenge.rule_list.rules.spec.body_matcher / 914440cf6750 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-e7e050a0106b5c243929bc99c717fe27390190239f5f0da6543c11a425e13c71)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-6b6b9bb417be0328d783de67330d570206c65ca27165054e0fe0e20523560695)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-c1694f8f786e00f1b8d062766cb92e3219566178dd23899b40bc7c9d82c517a1)
- policy_based_challenge.rule_list.rules.spec.body_matcher

<a id="canonical-4d33d14652e872b138f0faceea0ab599543886c7824fb4d2595e878de668ea0e"></a>

Type: `"object"`. single nested block, Optional.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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
body_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-804d380d705dd54ef554f18293678d61a2cc4df01ed52bc834ea5ef300d9aae6"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.body_matcher / 914440cf6750 / 3

<a id="canonical-ceacedd50389ac064eef6b34af3b8f8047fc6158bf8513bcf5cb83da31b0b82c"></a>

<a id="canonical-5ee82176b97d73d7b8bc020075b84ec26bfa04e09416f5502d856d8fd5105a6e"></a>

## exact_values property — policy_based_challenge.rule_list.rules.spec.body_matcher / 914440cf6750 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2ec24399f7f3ef748bb972c1084e4bf40961e275bb42f2f01005f8a4059192cf"></a>

<a id="canonical-a1258ddbcf13eefe9215976ae390eec550755d43ac6b46a420e0f6e4ec4988a6"></a>

## regex_values property — policy_based_challenge.rule_list.rules.spec.body_matcher / 914440cf6750 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-48dc5a5ac166f9959b7c7b88a13f64bb090fc81f698077a89df7ff58dee8c41b"></a>

<a id="canonical-4c68e1cf436adb2ec055285eb02333ca6163e44b2e2a764f023e1d777b7e3f9b"></a>

## transformers property — policy_based_challenge.rule_list.rules.spec.body_matcher / 914440cf6750 / 6

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-e09a95f763e71f8f1d862a9460de96bedebf62caae25f876fc9d0b10619d7a65"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.body_matcher / 914440cf6750 / 7

- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-c1694f8f786e00f1b8d062766cb92e3219566178dd23899b40bc7c9d82c517a1)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-ec7c39dfc583277454d8fc3ff5a931696d8107535dae7ea3f2297140818d1b9e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bdcb4bdeeff0c3446af7ec74598e90a05f729e5408c7e335298e6661a581c1dc"></a>

## policy_based_challenge.rule_list.rules.spec.client_selector — policy_based_challenge.rule_list.rules.spec.client_selector / 9db9ec7e7fbb / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-e7e050a0106b5c243929bc99c717fe27390190239f5f0da6543c11a425e13c71)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-6b6b9bb417be0328d783de67330d570206c65ca27165054e0fe0e20523560695)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-c1694f8f786e00f1b8d062766cb92e3219566178dd23899b40bc7c9d82c517a1)
- policy_based_challenge.rule_list.rules.spec.client_selector

<a id="canonical-9a1c6bb15807f48c56dda3d88ee7a9a4495faa6bf5cfbe02b738f2976fe5e978"></a>

Type: `"object"`. single nested block, Optional.

Type can be used to establish a 'selector reference' from one object(called selector) to a set of
other objects(called selectees) based on the value of expressions. A label selector is a label query
over a set of resources. An empty label selector matches all objects.

Upstream description:

This type can be used to establish a 'selector reference' from one object(called selector) to a set
of other objects(called selectees) based on the value of expressions. A label selector is a label
query over a set of resources. An empty label selector matches all objects. A null label selector
matches no objects. Label selector is immutable. Expressions is a list of strings of label selection
expression. Each string has "," separated values which are "AND" and all strings are logically "OR".
BNF for expression string &lt;selector-syntax&gt; ::= &lt;requirement&gt; | &lt;requirement&gt; ","
&lt;selector-syntax&gt; &lt;requirement&gt; ::= \[!\] KEY \[ &lt;set-based-restriction&gt; |
&lt;exact-match-restriction&gt; \] &lt;set-based-restriction&gt; ::= "" |
&lt;inclusion-exclusion&gt; &lt;value-set&gt; &lt;inclusion-exclusion&gt; ::= &lt;inclusion&gt; |
&lt;exclusion&gt; &lt;exclusion&gt; ::= "n&#111;tin" &lt;inclusion&gt; ::= "in" &lt;value-set&gt;
::= "(" &lt;values&gt; ")" &lt;values&gt; ::= VALUE | VALUE "," &lt;values&gt;
&lt;exact-match-restriction&gt; ::= \["="|"=="|"!="\] VALUE.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("expressions")}
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
client_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-47aa2119f215f7cc22097ee181d159c2f34163769b068646a97e6b0589e56255"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.client_selector / 9db9ec7e7fbb / 3

<a id="canonical-302744962fdab9fd0eaedcdb0cfa3d9e49b632481ff17e3f443cdb0eee7736b6"></a>

<a id="canonical-582ab6433f88744be00eed72f6f2663a80ae75d30a83cf4bacde7dd222a29d29"></a>

## expressions property — policy_based_challenge.rule_list.rules.spec.client_selector / 9db9ec7e7fbb / 4

Type: `["list", "string"]`. Optional.

Expressions contains the Kubernetes style label expression for selections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(1),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-f49ae6e96d78410b4ebfd0e027c7abdebbb54422498bb3b04722da964b15c5a9"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.client_selector / 9db9ec7e7fbb / 5

- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-c1694f8f786e00f1b8d062766cb92e3219566178dd23899b40bc7c9d82c517a1)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-f7820fa160919665e80f0206825016d7ef69fa7426efa4747f1bd99a82c18fdc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3d69c3940032eb82858496e300f3d175a02b94f5863675bf461b63e7904346d7"></a>

## policy_based_challenge.rule_list.rules.spec.cookie_matchers — policy_based_challenge.rule_list.rules.spec.cookie_matchers / e26bce5d2ab6 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-e7e050a0106b5c243929bc99c717fe27390190239f5f0da6543c11a425e13c71)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-6b6b9bb417be0328d783de67330d570206c65ca27165054e0fe0e20523560695)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-c1694f8f786e00f1b8d062766cb92e3219566178dd23899b40bc7c9d82c517a1)
- policy_based_challenge.rule_list.rules.spec.cookie_matchers

<a id="canonical-daefd02e3ed10cd9e2bcbcb673e14ccb5dbb5c8ffedefefe8c9f807080488396"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for all cookies that need to be matched. The criteria for matching each cookie is
described in individual instances of CookieMatcherType. The actual cookie values are extracted from
the request API as a list of strings for each cookie name.

Upstream description:

A list of predicates for all cookies that need to be matched. The criteria for matching each cookie
is described in individual instances of CookieMatcherType. The actual cookie values are extracted
from the request API as a list of strings for each cookie name. Note that all specified cookie
matcher predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
```

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
cookie_matchers {
  # Configure direct properties listed below.
}
```

<a id="canonical-ac143b3af8a18489cb0a0b22ce5fab1e778ebd456c83c78a1d51a0f70cb44c9c"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.cookie_matchers / e26bce5d2ab6 / 3

- [check_not_present](resources--cdn_loadbalancer--reference--group-013.md#canonical-20d5ddc07c5f0fb721f61e6ef999028209176e22f3dcf1b79dadffcec4645a31): complete subsection reference.

- [check_present](resources--cdn_loadbalancer--reference--group-013.md#canonical-1868a324d80ee4844003204112911338e5a514fcc45673ef13389cf39680948c): complete subsection reference.

<a id="canonical-7bee9e29b69e0b1cfe4ff9b89fc380109903a67a9c7d6b5a93449fe94bce76f8"></a>

<a id="canonical-c8e15cbb57da6c971e5f9e8f5a5570ff095f77df7176552bd706bb69c30118a8"></a>

## invert_matcher property — policy_based_challenge.rule_list.rules.spec.cookie_matchers / e26bce5d2ab6 / 4

Type: `"bool"`. Optional.

Invert Matcher. Invert Match of the expression defined.

Upstream description:

Invert Match of the expression defined.

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

- [item](resources--cdn_loadbalancer--reference--group-013.md#canonical-130c5ef0a0e206245f9d47f6d25661eb3230f1e9c596aa254b317edd9fc70bb2): complete subsection reference.

<a id="canonical-195c372f32cc8f22f5eb8e2549f8881d1e3d42b10d349c6bc4669ef0a98dcd4b"></a>

<a id="canonical-16e297b0a05daf20d384cdab9a205344b308f42a2241ec16f0ca5941fe8983fa"></a>

## name property — policy_based_challenge.rule_list.rules.spec.cookie_matchers / e26bce5d2ab6 / 5

Type: `"string"`. Optional.

Cookie Name. A case-sensitive cookie name.

Upstream description:

A case-sensitive cookie name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-e6d5741889ef7d0d3496fce4020818a1b853a870789038ac7ce1ad16b7c7ef44"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.cookie_matchers / e26bce5d2ab6 / 6

- [policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_not_present](resources--cdn_loadbalancer--reference--group-013.md#canonical-20d5ddc07c5f0fb721f61e6ef999028209176e22f3dcf1b79dadffcec4645a31)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_present](resources--cdn_loadbalancer--reference--group-013.md#canonical-1868a324d80ee4844003204112911338e5a514fcc45673ef13389cf39680948c)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers.item](resources--cdn_loadbalancer--reference--group-013.md#canonical-130c5ef0a0e206245f9d47f6d25661eb3230f1e9c596aa254b317edd9fc70bb2)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-c1694f8f786e00f1b8d062766cb92e3219566178dd23899b40bc7c9d82c517a1)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-20d5ddc07c5f0fb721f61e6ef999028209176e22f3dcf1b79dadffcec4645a31"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3b5245f7fc85f8bd8756c5072aaa974fdac2ceaf27c546be0b69747f607445db"></a>

## policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_not_present — policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_not_present / e46253d7113f / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-e7e050a0106b5c243929bc99c717fe27390190239f5f0da6543c11a425e13c71)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-6b6b9bb417be0328d783de67330d570206c65ca27165054e0fe0e20523560695)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-c1694f8f786e00f1b8d062766cb92e3219566178dd23899b40bc7c9d82c517a1)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](resources--cdn_loadbalancer--reference--group-013.md#canonical-f7820fa160919665e80f0206825016d7ef69fa7426efa4747f1bd99a82c18fdc)
- policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_not_present

<a id="canonical-787d3ddf5d6e09f9718b7cbbc57196faa3f12de2b3644fe0c0c1cd7d599e0275"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

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
check_not_present = {}
```

<a id="canonical-c0d0b0c86bdb3ea2785e9c9a319b46ace385a9eed43ef1bc6313dceb647e85de"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_not_present / e46253d7113f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a82c3733a7f821868a9b80a99d16c92cb5b624f61d97a75f7f6f405cec552e8c"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_not_present / e46253d7113f / 4

- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](resources--cdn_loadbalancer--reference--group-013.md#canonical-f7820fa160919665e80f0206825016d7ef69fa7426efa4747f1bd99a82c18fdc)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-1868a324d80ee4844003204112911338e5a514fcc45673ef13389cf39680948c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-63b6accc1056a5729bdbc2d00577186a13e81aa64c95e13f26c1b7056028f0b6"></a>

## policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_present — policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_present / dd3d74b9757e / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-e7e050a0106b5c243929bc99c717fe27390190239f5f0da6543c11a425e13c71)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-6b6b9bb417be0328d783de67330d570206c65ca27165054e0fe0e20523560695)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-c1694f8f786e00f1b8d062766cb92e3219566178dd23899b40bc7c9d82c517a1)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](resources--cdn_loadbalancer--reference--group-013.md#canonical-f7820fa160919665e80f0206825016d7ef69fa7426efa4747f1bd99a82c18fdc)
- policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_present

<a id="canonical-61303d0d44ba6f1b2139ffc96bcbd9fc2bacbc9e622ad9ab1c22575c71790ea6"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

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
check_present = {}
```

<a id="canonical-d5a0df80e7fbd3868105bbfd229ff618f4793e0c0e79583cb8b321e40b5a1e66"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_present / dd3d74b9757e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-214c70be9d081aff347db93f3ac5ca942ccc57d53ab344ea11983a05ae92b60c"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_present / dd3d74b9757e / 4

- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](resources--cdn_loadbalancer--reference--group-013.md#canonical-f7820fa160919665e80f0206825016d7ef69fa7426efa4747f1bd99a82c18fdc)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-130c5ef0a0e206245f9d47f6d25661eb3230f1e9c596aa254b317edd9fc70bb2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5f64c6596f5c049c6c543e5803ddadfb6dc32c7fb683354a50cf02f302caa7d9"></a>

## policy_based_challenge.rule_list.rules.spec.cookie_matchers.item — policy_based_challenge.rule_list.rules.spec.cookie_matchers.item / b3bc7ddd3d93 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-e7e050a0106b5c243929bc99c717fe27390190239f5f0da6543c11a425e13c71)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-6b6b9bb417be0328d783de67330d570206c65ca27165054e0fe0e20523560695)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-c1694f8f786e00f1b8d062766cb92e3219566178dd23899b40bc7c9d82c517a1)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](resources--cdn_loadbalancer--reference--group-013.md#canonical-f7820fa160919665e80f0206825016d7ef69fa7426efa4747f1bd99a82c18fdc)
- policy_based_challenge.rule_list.rules.spec.cookie_matchers.item

<a id="canonical-2742e5a2c500f1f6e913e5c9b7053c5528c260b6a829d576a1d7f343959e4428"></a>

Type: `"object"`. single nested block, Optional.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-1b1c38c3f58c30e6cc6a1da8ed51887e7f5d6d8ade82abffe152da242e17ba54"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.cookie_matchers.item / b3bc7ddd3d93 / 3

<a id="canonical-f174bb436f1eadbfb2f2edef2a0e9a61b60fa3c692b903f34619473c74b51915"></a>

<a id="canonical-a48fdd20586b506693bb50350ebbb61dbc5c70f9f3348b2035575db563d2f422"></a>

## exact_values property — policy_based_challenge.rule_list.rules.spec.cookie_matchers.item / b3bc7ddd3d93 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-39b276785837f81801196e765dc74350e9a5b6d2f19366cf44faf0d4e5d8a929"></a>

<a id="canonical-8b08d6fae311c61ef10c64e8162921d4921965ad2ae52619f6737d926914b30e"></a>

## regex_values property — policy_based_challenge.rule_list.rules.spec.cookie_matchers.item / b3bc7ddd3d93 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-771270e0402adde7dae64ca9fbc5b69a70abd50387a837f2b8640b579ba14326"></a>

<a id="canonical-2a123fc7b069e35b537cc21c6f9d31587f4927c544705de9dde981be6e162b9d"></a>

## transformers property — policy_based_challenge.rule_list.rules.spec.cookie_matchers.item / b3bc7ddd3d93 / 6

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-c2aa11874abef0ba80b810af7edc539e9de53d83938a9a652b1929697d1821d9"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.cookie_matchers.item / b3bc7ddd3d93 / 7

- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](resources--cdn_loadbalancer--reference--group-013.md#canonical-f7820fa160919665e80f0206825016d7ef69fa7426efa4747f1bd99a82c18fdc)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-0ef8295b621b356b70726ee6ddbf291d8ddb571073af277b924b4c59840e39e9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3e9175a00b2c9529ace8cc561bbc13c6bc7532d09c56c24baa5e75324d1e6709"></a>

## policy_based_challenge.rule_list.rules.spec.disable_challenge — policy_based_challenge.rule_list.rules.spec.disable_challenge / 57495618cdcc / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-e7e050a0106b5c243929bc99c717fe27390190239f5f0da6543c11a425e13c71)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-6b6b9bb417be0328d783de67330d570206c65ca27165054e0fe0e20523560695)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-c1694f8f786e00f1b8d062766cb92e3219566178dd23899b40bc7c9d82c517a1)
- policy_based_challenge.rule_list.rules.spec.disable_challenge

<a id="canonical-3303edf9f89bdfae21688c18e56c64c9ace0b0dac6144d61629ed71b55b2412f"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable challenge.

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
disable_challenge = {}
```

<a id="canonical-37bccbf82a51f24c1c05966040b4c45d93b78bdad3fa06aa839797a82182b990"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.disable_challenge / 57495618cdcc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-db4bdf68652407fa498381aa05314a911680d68541ac68c01e16d941f9ac5414"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.disable_challenge / 57495618cdcc / 4

- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-c1694f8f786e00f1b8d062766cb92e3219566178dd23899b40bc7c9d82c517a1)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-4df3dac25f5add58043c4a45e0bb68c779e67f54e56313602e09305a6f14e67f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4b03cbac21df2960a61c4b8f5546beeb9ee0066a7d10c962c4dfa41deb2d3668"></a>

## policy_based_challenge.rule_list.rules.spec.domain_matcher — policy_based_challenge.rule_list.rules.spec.domain_matcher / 9fbb7ea4cc32 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-e7e050a0106b5c243929bc99c717fe27390190239f5f0da6543c11a425e13c71)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-6b6b9bb417be0328d783de67330d570206c65ca27165054e0fe0e20523560695)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-c1694f8f786e00f1b8d062766cb92e3219566178dd23899b40bc7c9d82c517a1)
- policy_based_challenge.rule_list.rules.spec.domain_matcher

<a id="canonical-51af80a11941870e0a66ba8b0b5fb7ed03187b340001565243453678fd6989e3"></a>

Type: `"object"`. single nested block, Optional.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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
domain_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-42d19f828ba82b9bc1eaefc14f1183e7366430cfaae21a882695d5575951588b"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.domain_matcher / 9fbb7ea4cc32 / 3

<a id="canonical-55910c57e8dca6b798aeb8ded6f10db8fcd63c940af8f6554c3fdafe388937b0"></a>

<a id="canonical-f1490623bff857350b70911df555f3e263277b37b7e66ced1d0586a5358e1e03"></a>

## exact_values property — policy_based_challenge.rule_list.rules.spec.domain_matcher / 9fbb7ea4cc32 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3b30c9aeeac08f55717f4489a4657887cd9fa1d0b217e933d9c534add99ced79"></a>

<a id="canonical-4348455f7ee35f6fffe7ecb4988c11ccf0f58ece2c45b3ea21ae8de5d59545a1"></a>

## regex_values property — policy_based_challenge.rule_list.rules.spec.domain_matcher / 9fbb7ea4cc32 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-67d47cc7f75f2e65d30ebc43608f72730dc8284deed94e3b1f6d7a84fe4d634e"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.domain_matcher / 9fbb7ea4cc32 / 6

- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-c1694f8f786e00f1b8d062766cb92e3219566178dd23899b40bc7c9d82c517a1)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-d3d73cf1d86a379d97a04d2ce491647779e67beb9c7eca6beb3fc6cc6213b0f8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d6e6cac28e2bfc3a6d13961464d6426b5f70282e1cf71da7d0e34a27d1f8848c"></a>

## policy_based_challenge.rule_list.rules.spec.enable_captcha_challenge — policy_based_challenge.rule_list.rules.spec.enable_captcha_challenge / f9b436cc13e9 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-e7e050a0106b5c243929bc99c717fe27390190239f5f0da6543c11a425e13c71)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-6b6b9bb417be0328d783de67330d570206c65ca27165054e0fe0e20523560695)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-c1694f8f786e00f1b8d062766cb92e3219566178dd23899b40bc7c9d82c517a1)
- policy_based_challenge.rule_list.rules.spec.enable_captcha_challenge

<a id="canonical-ee169ad6750c5409cbd5c618a7e8cc7e665c967187b3a5a40455de448b5a3a04"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable captcha challenge.

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
enable_captcha_challenge = {}
```

<a id="canonical-ff1aebaed8619373ac2209a8fd793e27c02764add47e33ef2270d8f52e92bf1d"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.enable_captcha_challenge / f9b436cc13e9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9a2edffae0760cc88eff9a75024dc1e0b4cb648933cf3ed45dc517c76da847ca"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.enable_captcha_challenge / f9b436cc13e9 / 4

- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-c1694f8f786e00f1b8d062766cb92e3219566178dd23899b40bc7c9d82c517a1)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-51e7028886650667705be61ff103bbd51fc27e835c8c9be24db09884020dbf44"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dec62b9a8879d7fc08d7208727ca4e88fa021b308836621df78f8c081d5c7462"></a>

## policy_based_challenge.rule_list.rules.spec.enable_javascript_challenge — policy_based_challenge.rule_list.rules.spec.enable_javascript_challenge / 42c0acc108a7 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-e7e050a0106b5c243929bc99c717fe27390190239f5f0da6543c11a425e13c71)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-6b6b9bb417be0328d783de67330d570206c65ca27165054e0fe0e20523560695)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-c1694f8f786e00f1b8d062766cb92e3219566178dd23899b40bc7c9d82c517a1)
- policy_based_challenge.rule_list.rules.spec.enable_javascript_challenge

<a id="canonical-dc757821c6e32cdaa05d1094f7798b7f09b88f7341a81f665e9937271ba316cf"></a>

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
enable_javascript_challenge = {}
```

<a id="canonical-99fd45ccfaa4b4e995c9b61a25d17a28e7831ac824c1d021ed0e7a9e3adc1937"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.enable_javascript_challenge / 42c0acc108a7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-19ea9897983f0c473fda50c51d58dd1565c3716ec7d0d6db62e425fc9fbdb91f"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.enable_javascript_challenge / 42c0acc108a7 / 4

- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-c1694f8f786e00f1b8d062766cb92e3219566178dd23899b40bc7c9d82c517a1)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-fc8041ae07d56ba34ef24020aa6a5accfa48b0ec10c1362557b6eeff2a4f3d88"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4b3274060ff08f044959d07768f330a853ebdc1ddbdfbb5fe5cf584e25cb0e05"></a>

## policy_based_challenge.rule_list.rules.spec.headers — policy_based_challenge.rule_list.rules.spec.headers / 7c240a3ae24a / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-e7e050a0106b5c243929bc99c717fe27390190239f5f0da6543c11a425e13c71)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-6b6b9bb417be0328d783de67330d570206c65ca27165054e0fe0e20523560695)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-c1694f8f786e00f1b8d062766cb92e3219566178dd23899b40bc7c9d82c517a1)
- policy_based_challenge.rule_list.rules.spec.headers

<a id="canonical-33ae45be083c4b5d8b7003069d2c9ed0e441c6f1d19a0834553889b140df83c1"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for various HTTP headers that need to match. The criteria for matching each HTTP
header are described in individual HeaderMatcherType instances. The actual HTTP header values are
extracted from the request API as a list of strings for each HTTP header type.

Upstream description:

A list of predicates for various HTTP headers that need to match. The criteria for matching each
HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values
are extracted from the request API as a list of strings for each HTTP header type. Note that all
specified header predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
```

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
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
headers {
  # Configure direct properties listed below.
}
```

<a id="canonical-61c88db05d403ddaea4e172389acc4dc06a59ecb24a1416cadaec862fab85f2e"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.headers / 7c240a3ae24a / 3

- [check_not_present](resources--cdn_loadbalancer--reference--group-013.md#canonical-9fbace9e870cbb378ed85ffc3b3669b762adf0b21ccb081c5dc8acb1ea7b19f4): complete subsection reference.

- [check_present](resources--cdn_loadbalancer--reference--group-013.md#canonical-bbbc6ed8518c5abc3ca44f5883305065afef0acad4860e2f37e9fba4f8aae860): complete subsection reference.

<a id="canonical-9b9ce168e8e1ff737447c8595ed2d799837df6d6cf1dc25f829efc0249f97064"></a>

<a id="canonical-461fe291d25cb58e31f7fbfda34060a32e5554141d9b5d4542017bbef5f5e28a"></a>

## invert_matcher property — policy_based_challenge.rule_list.rules.spec.headers / 7c240a3ae24a / 4

Type: `"bool"`. Optional.

Invert Header Matcher. Invert the match result.

Upstream description:

Invert the match result.

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

- [item](resources--cdn_loadbalancer--reference--group-013.md#canonical-08127dedac9c6413e149634f18018fbe890c8ecd6aeb35ee5a9b44474d62a4ca): complete subsection reference.

<a id="canonical-190d6f68564e1cadc4fa49d5e741c4cfec8ac4feba5ddcbd3c6b718ccd783d1d"></a>

<a id="canonical-f7afcbb44c5a26a94000eab0cb824fd19e2297b70eaf45f9cfc11b5dbdf686a5"></a>

## name property — policy_based_challenge.rule_list.rules.spec.headers / 7c240a3ae24a / 5

Type: `"string"`. Optional.

Header Name. A case-insensitive HTTP header name.

Upstream description:

A case-insensitive HTTP header name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-a032cdd24845dbce832612dfac865e3450783d83623faf55fbe0dd1f63e3a337"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.headers / 7c240a3ae24a / 6

- [policy_based_challenge.rule_list.rules.spec.headers.check_not_present](resources--cdn_loadbalancer--reference--group-013.md#canonical-9fbace9e870cbb378ed85ffc3b3669b762adf0b21ccb081c5dc8acb1ea7b19f4)
- [policy_based_challenge.rule_list.rules.spec.headers.check_present](resources--cdn_loadbalancer--reference--group-013.md#canonical-bbbc6ed8518c5abc3ca44f5883305065afef0acad4860e2f37e9fba4f8aae860)
- [policy_based_challenge.rule_list.rules.spec.headers.item](resources--cdn_loadbalancer--reference--group-013.md#canonical-08127dedac9c6413e149634f18018fbe890c8ecd6aeb35ee5a9b44474d62a4ca)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-c1694f8f786e00f1b8d062766cb92e3219566178dd23899b40bc7c9d82c517a1)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-9fbace9e870cbb378ed85ffc3b3669b762adf0b21ccb081c5dc8acb1ea7b19f4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f1b0c283306cb1c14ba4a4d7f67970ec11bf4327ced8af0b887bdd604aa03969"></a>

## policy_based_challenge.rule_list.rules.spec.headers.check_not_present — policy_based_challenge.rule_list.rules.spec.headers.check_not_present / 6d476d607b7f / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-e7e050a0106b5c243929bc99c717fe27390190239f5f0da6543c11a425e13c71)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-6b6b9bb417be0328d783de67330d570206c65ca27165054e0fe0e20523560695)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-c1694f8f786e00f1b8d062766cb92e3219566178dd23899b40bc7c9d82c517a1)
- [policy_based_challenge.rule_list.rules.spec.headers](resources--cdn_loadbalancer--reference--group-013.md#canonical-fc8041ae07d56ba34ef24020aa6a5accfa48b0ec10c1362557b6eeff2a4f3d88)
- policy_based_challenge.rule_list.rules.spec.headers.check_not_present

<a id="canonical-6309c93827032a2484f6ea6f72aa2a8449d9f54cc46cb531015e355e3a39b9f8"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

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
check_not_present = {}
```

<a id="canonical-9be5a270ebd91666def49ea3bed9b68e9dc269ae0c429e0366282fa5d6c6b15c"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.headers.check_not_present / 6d476d607b7f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b79939c0ec25e317de1bba0f62affd1fcd031310a9bbfa35cb85a51d7bf3b813"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.headers.check_not_present / 6d476d607b7f / 4

- [policy_based_challenge.rule_list.rules.spec.headers](resources--cdn_loadbalancer--reference--group-013.md#canonical-fc8041ae07d56ba34ef24020aa6a5accfa48b0ec10c1362557b6eeff2a4f3d88)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-bbbc6ed8518c5abc3ca44f5883305065afef0acad4860e2f37e9fba4f8aae860"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b03dd613aa315d8ed0347ea896b8923974c67e54151f548350bd15cfaf2949c5"></a>

## policy_based_challenge.rule_list.rules.spec.headers.check_present — policy_based_challenge.rule_list.rules.spec.headers.check_present / 871704c5aa73 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-e7e050a0106b5c243929bc99c717fe27390190239f5f0da6543c11a425e13c71)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-6b6b9bb417be0328d783de67330d570206c65ca27165054e0fe0e20523560695)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-c1694f8f786e00f1b8d062766cb92e3219566178dd23899b40bc7c9d82c517a1)
- [policy_based_challenge.rule_list.rules.spec.headers](resources--cdn_loadbalancer--reference--group-013.md#canonical-fc8041ae07d56ba34ef24020aa6a5accfa48b0ec10c1362557b6eeff2a4f3d88)
- policy_based_challenge.rule_list.rules.spec.headers.check_present

<a id="canonical-e88650bc05a322ecabaaa68ed646b86fe7ef9b805778c3603ff84098763450ec"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

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
check_present = {}
```

<a id="canonical-1f23e783de07421df8f9785a5519377b833ad7e1631334716bb478aba5b57782"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.headers.check_present / 871704c5aa73 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2d337aef79f8dbe0f492cc5766fa13c1e3e9287384e66ec4dd02c804b1cf9cbe"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.headers.check_present / 871704c5aa73 / 4

- [policy_based_challenge.rule_list.rules.spec.headers](resources--cdn_loadbalancer--reference--group-013.md#canonical-fc8041ae07d56ba34ef24020aa6a5accfa48b0ec10c1362557b6eeff2a4f3d88)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-08127dedac9c6413e149634f18018fbe890c8ecd6aeb35ee5a9b44474d62a4ca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f705a7b577a4c49d7d25cd954070384096a165d1c87fa480f6e624836929d139"></a>

## policy_based_challenge.rule_list.rules.spec.headers.item — policy_based_challenge.rule_list.rules.spec.headers.item / b14082afeb12 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-e7e050a0106b5c243929bc99c717fe27390190239f5f0da6543c11a425e13c71)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-6b6b9bb417be0328d783de67330d570206c65ca27165054e0fe0e20523560695)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-c1694f8f786e00f1b8d062766cb92e3219566178dd23899b40bc7c9d82c517a1)
- [policy_based_challenge.rule_list.rules.spec.headers](resources--cdn_loadbalancer--reference--group-013.md#canonical-fc8041ae07d56ba34ef24020aa6a5accfa48b0ec10c1362557b6eeff2a4f3d88)
- policy_based_challenge.rule_list.rules.spec.headers.item

<a id="canonical-4166cf828060c4ccd6199a949b338fd74467f7f80ccd9c47f6e8ab8b57e01154"></a>

Type: `"object"`. single nested block, Optional.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-9e17f154045dcf4a39a0970440405125a3d3691e6bb1be275b4879bb05edfeee"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.headers.item / b14082afeb12 / 3

<a id="canonical-5c9e98af33e75a644405abc65c650e21a3b3bd78008c6d43ea45fbc8d81f7182"></a>

<a id="canonical-9899c0b59d34a2d625d6bd615ac514524c028fb8985a1fa99d73ecd01663e7e4"></a>

## exact_values property — policy_based_challenge.rule_list.rules.spec.headers.item / b14082afeb12 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-c07db77f357eee6b7aad237a904f2305cce0491e8a3a59bc9953a0165190dc45"></a>

<a id="canonical-e1cbc53be893015cbaa29a13a421b9864f91d8d5d985f70d2b733f0111c269ed"></a>

## regex_values property — policy_based_challenge.rule_list.rules.spec.headers.item / b14082afeb12 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2c44fdb5a41004d85085425acf6eecd6e461526073922a27760422fc048dc430"></a>

<a id="canonical-c04042a716e7264eaf62fdc643633810523a6eda53d5667d549ac7e0ebe363ed"></a>

## transformers property — policy_based_challenge.rule_list.rules.spec.headers.item / b14082afeb12 / 6

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-7015b0d5b1c3a02ba20a341988cbcf7d42b9c826b720a44a95010aa7c17c7ef6"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.headers.item / b14082afeb12 / 7

- [policy_based_challenge.rule_list.rules.spec.headers](resources--cdn_loadbalancer--reference--group-013.md#canonical-fc8041ae07d56ba34ef24020aa6a5accfa48b0ec10c1362557b6eeff2a4f3d88)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-b6776a702c1500c90976697d1bffab0f64d0ecfa783f880ca9faa078bcd35243"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-80175a34c41106a0a7fb4693c4376d8590753a6d7e5dc4b188fab08915e725ac"></a>

## policy_based_challenge.rule_list.rules.spec.http_method — policy_based_challenge.rule_list.rules.spec.http_method / 8f0d653ac304 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-e7e050a0106b5c243929bc99c717fe27390190239f5f0da6543c11a425e13c71)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-6b6b9bb417be0328d783de67330d570206c65ca27165054e0fe0e20523560695)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-c1694f8f786e00f1b8d062766cb92e3219566178dd23899b40bc7c9d82c517a1)
- policy_based_challenge.rule_list.rules.spec.http_method

<a id="canonical-f89994b7ac353a14a5843de67892b17b0118222ea0565da51e8de9d03bac16ae"></a>

Type: `"object"`. single nested block, Optional.

HTTP method matcher specifies a list of methods to match an input HTTP method. The match is
considered successful if the input method is a member of the list. The result of the match based on
the method list is inverted if invert\_matcher is true.

Upstream description:

A HTTP method matcher specifies a list of methods to match an input HTTP method. The match is
considered successful if the input method is a member of the list. The result of the match based on
the method list is inverted if invert\_matcher is true.

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
http_method {
  # Configure direct properties listed below.
}
```

<a id="canonical-92de7b54b74ed6d7c7147ef1b914e19ce4c0646229e969bed7484261294fa5f0"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.http_method / 8f0d653ac304 / 3

<a id="canonical-a3bae577699806c4fb8172fde0e4ea39a1510873aac5b6193a880ccf67345116"></a>

<a id="canonical-c17f44630779e0ee65ad10fe6036f188cb4ff651e61e7aa459eea68ce0f146e6"></a>

## invert_matcher property — policy_based_challenge.rule_list.rules.spec.http_method / 8f0d653ac304 / 4

Type: `"bool"`. Optional.

Invert Method Matcher. Invert the match result.

Upstream description:

Invert the match result.

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

<a id="canonical-044bf24a76da552fb1b40ff18fd8209187cdcc9f17a0ac6245ab89094d925e50"></a>

<a id="canonical-e1aa45e8bca9bd91f61c39c89d4b074517da428f7dfb4053dc75be41bcb25935"></a>

## methods property — policy_based_challenge.rule_list.rules.spec.http_method / 8f0d653ac304 / 5

Type: `["list", "string"]`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] List of methods values to
match against. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`,
\`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Upstream description:

List of methods values to match against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-5eb079d60b51bacb4511ac576d7c895ab7e07dd51795c9a3659f7cdd3ad9c94d"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.http_method / 8f0d653ac304 / 6

- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-c1694f8f786e00f1b8d062766cb92e3219566178dd23899b40bc7c9d82c517a1)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-698f05ca9efe080743409822581d5ad76e0d8fd235fb254c80eccfbcca391788"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-49499fc9fc63657c9c8b75c7f538edce30b6cfa6254f78b79a46ec5de1ed0819"></a>

## policy_based_challenge.rule_list.rules.spec.ip_matcher — policy_based_challenge.rule_list.rules.spec.ip_matcher / 7749a906799e / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-e7e050a0106b5c243929bc99c717fe27390190239f5f0da6543c11a425e13c71)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-6b6b9bb417be0328d783de67330d570206c65ca27165054e0fe0e20523560695)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-c1694f8f786e00f1b8d062766cb92e3219566178dd23899b40bc7c9d82c517a1)
- policy_based_challenge.rule_list.rules.spec.ip_matcher

<a id="canonical-bafdea486eba89e86ea4b33465d494ef3c13f194e5e02c3b3e0b3b749e6c96d4"></a>

Type: `"object"`. single nested block, Optional.

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Upstream description:

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("prefix_sets")}
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
ip_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-500c37b39038b48881e00b3451b6fd2ca731441c3dbad19b32b75ffb17287987"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.ip_matcher / 7749a906799e / 3

<a id="canonical-8458e1177c30f78c70da63d8ee498efb4cb9d18cee7f2081eb6be4cf496f0c77"></a>

<a id="canonical-98354048c95cf876c1b123bcd461a41c962915df1857546cd0b32c1b28feb442"></a>

## invert_matcher property — policy_based_challenge.rule_list.rules.spec.ip_matcher / 7749a906799e / 4

Type: `"bool"`. Optional.

Invert IP Matcher. Invert the match result.

Upstream description:

Invert the match result.

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

- [prefix_sets](resources--cdn_loadbalancer--reference--group-013.md#canonical-b51882fe21f10d4cef58edbd6ac815085436859dd6303697463b258e36a5ca5e): complete subsection reference.

<a id="canonical-76e07762e9d734a91e76ac71571abca293aeec7140b15b4f96ac570102f0edd9"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.ip_matcher / 7749a906799e / 5

- [policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets](resources--cdn_loadbalancer--reference--group-013.md#canonical-b51882fe21f10d4cef58edbd6ac815085436859dd6303697463b258e36a5ca5e)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-c1694f8f786e00f1b8d062766cb92e3219566178dd23899b40bc7c9d82c517a1)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-b51882fe21f10d4cef58edbd6ac815085436859dd6303697463b258e36a5ca5e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2d9c7409fc2fe26563910e48b1b69ad270a89740093ab893aa42f628af261e5e"></a>

## policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets — policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets / c48cc35d34b1 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-e7e050a0106b5c243929bc99c717fe27390190239f5f0da6543c11a425e13c71)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-6b6b9bb417be0328d783de67330d570206c65ca27165054e0fe0e20523560695)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-c1694f8f786e00f1b8d062766cb92e3219566178dd23899b40bc7c9d82c517a1)
- [policy_based_challenge.rule_list.rules.spec.ip_matcher](resources--cdn_loadbalancer--reference--group-013.md#canonical-698f05ca9efe080743409822581d5ad76e0d8fd235fb254c80eccfbcca391788)
- policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets

<a id="canonical-9b8416eb6101863d4c25ced550de0cd6ee90ec51b07fac4319ab397bd4f286a0"></a>

Type: `"object"`. list nested block, Optional.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

Terraform syntax:

```terraform
prefix_sets {
  # Configure direct properties listed below.
}
```

<a id="canonical-ccbde7c9c92e8a23c5ba5a84fb81e91f9de8645a5d38b811a2235b9c55e5ff41"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets / c48cc35d34b1 / 3

<a id="canonical-2ee15a8c728c3f35bd6e246cd67b4724eeb621126c6a4b303f93f11588467992"></a>

<a id="canonical-44fca43b96d053f4340c726dae05404b8a5e7685e18ec0f9e82e7770bbdb6a0a"></a>

## kind property — policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets / c48cc35d34b1 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-d978743c6868da83ae9c1fcde7e27cbde18e098312c798bfd3d3a7f951bd6031"></a>

<a id="canonical-758171da4ba1a545b8e919772c5d2c77c28bcac30daea0e42bb8447651b67326"></a>

## name property — policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets / c48cc35d34b1 / 5

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-051a2bf2e8f5902718bc3c2479826eaec52e936f57184bc9d1e6ea8cb0f676f7"></a>

<a id="canonical-c5909221d2a63596b12ab7bb81702172f28fd9453a76c4144941fa99d4e69864"></a>

## namespace property — policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets / c48cc35d34b1 / 6

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
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
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
  }
}
```

<a id="canonical-3603b9ed1518e880cf94f91fc7b1d660e71096366a29ec77c78b2dcb4274f180"></a>

<a id="canonical-665aeb12f3abddd4f755d7b25df4b293f74a5ea335cee39f20dcc1b1357cf286"></a>

## tenant property — policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets / c48cc35d34b1 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-b06a0e7df0cd59ff568be85fd42bdb8963a208b1498f7c1ed2e43457f4096f90"></a>

<a id="canonical-53cfe94fff6ca1f596f55cf51385f7c1d8a3ccc96d0333bbfe7a6c0c87994179"></a>

## uid property — policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets / c48cc35d34b1 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-2168a706041c8b13a51c1adace8fad28e7cbd8a70b04c8218282ea04f8bd37ce"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets / c48cc35d34b1 / 9

- [policy_based_challenge.rule_list.rules.spec.ip_matcher](resources--cdn_loadbalancer--reference--group-013.md#canonical-698f05ca9efe080743409822581d5ad76e0d8fd235fb254c80eccfbcca391788)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-463f810aaab48ed3b9009874cd2f5f4a0608b779ac170ed0dcd87ab8a766b824"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0a9c6d2ee426aa6f3cf9efb9ab40d5d1ab5d45dbc749827f9502fa6ae23dd250"></a>

## policy_based_challenge.rule_list.rules.spec.ip_prefix_list — policy_based_challenge.rule_list.rules.spec.ip_prefix_list / f68f640e2db3 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-e7e050a0106b5c243929bc99c717fe27390190239f5f0da6543c11a425e13c71)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-6b6b9bb417be0328d783de67330d570206c65ca27165054e0fe0e20523560695)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-c1694f8f786e00f1b8d062766cb92e3219566178dd23899b40bc7c9d82c517a1)
- policy_based_challenge.rule_list.rules.spec.ip_prefix_list

<a id="canonical-5b0080e15ab2a1273100d1f1f819682a6cb920e394165676ee83fb6417cf1933"></a>

Type: `"object"`. single nested block, Optional.

List of IP Prefix strings to match against.

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
ip_prefix_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-6fc77ccb211463bfc957d53fb022314c2178e3578007609bb4999f27e9203ee1"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.ip_prefix_list / f68f640e2db3 / 3

<a id="canonical-b6b833d7b2382a42ce80648bcb69f0abc508dc04904557b5bfe0f2fa42eca982"></a>

<a id="canonical-0642689d9ca49a6876750ea2743b388a216a5accea75b91ccba77f03368473f8"></a>

## invert_match property — policy_based_challenge.rule_list.rules.spec.ip_prefix_list / f68f640e2db3 / 4

Type: `"bool"`. Optional.

Invert Match Result. Invert the match result.

Upstream description:

Invert the match result.

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

<a id="canonical-3cb540423eabaad89065a2e4f8fb4c84ef617a5b2e121227e46b5af00050d775"></a>

<a id="canonical-4fbb06e45bd4cdf8fc46bb3805fb980a8051d97b4af21ded4c7f7248937ab89b"></a>

## ip_prefixes property — policy_based_challenge.rule_list.rules.spec.ip_prefix_list / f68f640e2db3 / 5

Type: `["list", "string"]`. Optional.

IPv4 Prefix List. List of IPv4 prefix strings.

Upstream description:

List of IPv4 prefix strings.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
```

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
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-57e2a8f62d162e740233255cbb963d2ca82c45779e0f5d979ff9e426ad61b07b"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.ip_prefix_list / f68f640e2db3 / 6

- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-c1694f8f786e00f1b8d062766cb92e3219566178dd23899b40bc7c9d82c517a1)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-7b35f428c37707746358a0356c25eec89e66574da2d70f66d5b38214c515df95"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9775e8e568e5d3c8ec4ce0c1db3757281102304bf8d3a770d55afd2fdc7aaab5"></a>

## policy_based_challenge.rule_list.rules.spec.path — policy_based_challenge.rule_list.rules.spec.path / 4e0de6495400 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-e7e050a0106b5c243929bc99c717fe27390190239f5f0da6543c11a425e13c71)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-6b6b9bb417be0328d783de67330d570206c65ca27165054e0fe0e20523560695)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-c1694f8f786e00f1b8d062766cb92e3219566178dd23899b40bc7c9d82c517a1)
- policy_based_challenge.rule_list.rules.spec.path

<a id="canonical-4021514ce3d29e8d38c95829e07f9abab558633771d2a36e5096deaccda82737"></a>

Type: `"object"`. single nested block, Optional.

Path matcher specifies multiple criteria for matching an HTTP path string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of path prefixes, a list of exact path values and a list of regular expressions.

Upstream description:

A path matcher specifies multiple criteria for matching an HTTP path string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of path prefixes, a list of exact path values and a list of regular expressions.

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
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-aebe64c3efd397f3c7c7aa6bc205711465fbc859d7d0945e070df63a79d08792"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.path / 4e0de6495400 / 3

<a id="canonical-f433ee91d290a8e168e12a2f0706e15762bb8650e38a5f737f7e3e73810e3140"></a>

<a id="canonical-3d90eb48606f8e6b10d293a03ab0dc9378d640aaebba2b6e8e0686cf4265053b"></a>

## encoded_path_matcher property — policy_based_challenge.rule_list.rules.spec.path / 4e0de6495400 / 4

Type: `"bool"`. Optional.

Match against the encoded, escaped path.

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

<a id="canonical-5aacf8c1c0d5a4a156e62d1835945d31a35cd0a785768e12d01e95f7c0ad39a1"></a>

<a id="canonical-854719a38530473940389727c0998671fdb01f99e6275ecdc17dc8b32fc11377"></a>

## exact_values property — policy_based_challenge.rule_list.rules.spec.path / 4e0de6495400 / 5

Type: `["list", "string"]`. Optional.

List of exact path values to match the input HTTP path against.

Upstream description:

A list of exact path values to match the input HTTP path against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-4b53dcd6122f9a97270d81d071ef0c68562f4e4a0c5d98671e6db94f6d136ae2"></a>

<a id="canonical-d7ca0f2ad9525d374ddfcc1dcc78b5218af5cfcb91f4cfe1a9fdcb7e830db126"></a>

## invert_matcher property — policy_based_challenge.rule_list.rules.spec.path / 4e0de6495400 / 6

Type: `"bool"`. Optional.

Invert Path Matcher. Invert the match result.

Upstream description:

Invert the match result.

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

<a id="canonical-35e2dace988bb72aa34b2062e86f1dac2b0bd9a76eda1f29ed4eb4049d93fb63"></a>

<a id="canonical-cb4ff7ff632a00918ec590511759b760431203d5b72700f92596df972d4981df"></a>

## prefix_values property — policy_based_challenge.rule_list.rules.spec.path / 4e0de6495400 / 7

Type: `["list", "string"]`. Optional.

List of path prefix values to match the input HTTP path against.

Upstream description:

A list of path prefix values to match the input HTTP path against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-7f760e682425f203746834daf6e253275c45f8c3c1e7a9750ff0b1db0da0168e"></a>

<a id="canonical-8aa4ad1705a69c2dea643f1af7ecfe8678e4c39ea02caea7d1dc56c749fe02a7"></a>

## regex_values property — policy_based_challenge.rule_list.rules.spec.path / 4e0de6495400 / 8

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input HTTP path against.

Upstream description:

A list of regular expressions to match the input HTTP path against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-327fb86838b4c55668f46d6899d44c8a663c0c57c85aa9b2e06c51c588819753"></a>

<a id="canonical-0beb8a89dfd3dd4a27449b17f9aec1e821272238c2e7a76b1202ec57c01967f3"></a>

## suffix_values property — policy_based_challenge.rule_list.rules.spec.path / 4e0de6495400 / 9

Type: `["list", "string"]`. Optional.

List of path suffix values to match the input HTTP path against.

Upstream description:

A list of path suffix values to match the input HTTP path against.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "64",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "64",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-d1ee139bdc5498719ef0587441671555cc59f0c920130e4f2213586ae89e01f6"></a>

<a id="canonical-f009f553f9ef6a7c22d4fa10e049e0367d5eb0697ff9bd476bae857094bea016"></a>

## transformers property — policy_based_challenge.rule_list.rules.spec.path / 4e0de6495400 / 10

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-88d209eb18ad4404f5704996e3ff9ccfa9e67fcdfe147fb3aa5fa8bb2c1a88c5"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.path / 4e0de6495400 / 11

- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-c1694f8f786e00f1b8d062766cb92e3219566178dd23899b40bc7c9d82c517a1)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-511830efa2a971e782cd6ba252499d4cf40f63d59a6550161b659d22def20550"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5e92d73542335e93ec265616f9239f4fdecebb6d3af322e43f5e78f4333e1190"></a>

## policy_based_challenge.rule_list.rules.spec.query_params — policy_based_challenge.rule_list.rules.spec.query_params / 009b4697b931 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-e7e050a0106b5c243929bc99c717fe27390190239f5f0da6543c11a425e13c71)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-6b6b9bb417be0328d783de67330d570206c65ca27165054e0fe0e20523560695)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-c1694f8f786e00f1b8d062766cb92e3219566178dd23899b40bc7c9d82c517a1)
- policy_based_challenge.rule_list.rules.spec.query_params

<a id="canonical-3c1d6acc2bbd316d96364dd7bed0a6366e3e0431852225c907371050b2382820"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for all query parameters that need to be matched. The criteria for matching each
query parameter are described in individual instances of QueryParameterMatcherType. The actual query
parameter values are extracted from the request API as a list of strings for each query..

Upstream description:

A list of predicates for all query parameters that need to be matched. The criteria for matching
each query parameter are described in individual instances of QueryParameterMatcherType. The actual
query parameter values are extracted from the request API as a list of strings for each query
parameter name. Note that all specified query parameter predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("key"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
```

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
query_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-aa1932b50cea8f165263ed5c092faee2c225a6993ec9e2beb37f1cd9bd32df52"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.query_params / 009b4697b931 / 3

- [check_not_present](resources--cdn_loadbalancer--reference--group-013.md#canonical-76c3c8a22cf0e473bd9e82d7f6b36c7cfd408ad7737402028dae2c344d70590a): complete subsection reference.

- [check_present](resources--cdn_loadbalancer--reference--group-013.md#canonical-05827f52f304ac421a5a95c55139e3d6ba110e156327463d0895f7a59fb1b5a1): complete subsection reference.

<a id="canonical-d211685eea5f5926522a12472641fee3b27d724d89d714e82bd7381136e4c990"></a>

<a id="canonical-1f30ef766f61f032e7c184379b2d490d7ec12824db25a836bd252c9b30dd891e"></a>

## invert_matcher property — policy_based_challenge.rule_list.rules.spec.query_params / 009b4697b931 / 4

Type: `"bool"`. Optional.

Invert Query Parameter Matcher. Invert the match result.

Upstream description:

Invert the match result.

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

- [item](resources--cdn_loadbalancer--reference--group-013.md#canonical-30798cc5aa1f9e99534033d31c0052dcd28e1cfeedce619c21e3025633eed62f): complete subsection reference.

<a id="canonical-194fb6c83a3a0c244401be2ea7a050b82141f2a2b7494d784599a47de8acf2ea"></a>

<a id="canonical-6381e789d483825a9bad032d986816b848c88f8f76d342d5d2f46950653507fb"></a>

## key property — policy_based_challenge.rule_list.rules.spec.query_params / 009b4697b931 / 5

Type: `"string"`. Optional.

Case-sensitive HTTP query parameter name.

Upstream description:

A case-sensitive HTTP query parameter name.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-fdd9074fd00a8dba182f8ba11d63a2872eaefb961c9bdc6e4b7c280f0f0f89ec"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.query_params / 009b4697b931 / 6

- [policy_based_challenge.rule_list.rules.spec.query_params.check_not_present](resources--cdn_loadbalancer--reference--group-013.md#canonical-76c3c8a22cf0e473bd9e82d7f6b36c7cfd408ad7737402028dae2c344d70590a)
- [policy_based_challenge.rule_list.rules.spec.query_params.check_present](resources--cdn_loadbalancer--reference--group-013.md#canonical-05827f52f304ac421a5a95c55139e3d6ba110e156327463d0895f7a59fb1b5a1)
- [policy_based_challenge.rule_list.rules.spec.query_params.item](resources--cdn_loadbalancer--reference--group-013.md#canonical-30798cc5aa1f9e99534033d31c0052dcd28e1cfeedce619c21e3025633eed62f)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-c1694f8f786e00f1b8d062766cb92e3219566178dd23899b40bc7c9d82c517a1)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-76c3c8a22cf0e473bd9e82d7f6b36c7cfd408ad7737402028dae2c344d70590a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c27d8e47c6c73646faad81d0ae9f56071ab39948cd9d3733d96ac9cb41c2f3ea"></a>

## policy_based_challenge.rule_list.rules.spec.query_params.check_not_present — policy_based_challenge.rule_list.rules.spec.query_params.check_not_present / 2596f28e3966 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-e7e050a0106b5c243929bc99c717fe27390190239f5f0da6543c11a425e13c71)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-6b6b9bb417be0328d783de67330d570206c65ca27165054e0fe0e20523560695)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-c1694f8f786e00f1b8d062766cb92e3219566178dd23899b40bc7c9d82c517a1)
- [policy_based_challenge.rule_list.rules.spec.query_params](resources--cdn_loadbalancer--reference--group-013.md#canonical-511830efa2a971e782cd6ba252499d4cf40f63d59a6550161b659d22def20550)
- policy_based_challenge.rule_list.rules.spec.query_params.check_not_present

<a id="canonical-fbcf949c776e1810c0155a7e3fae4ee8e5be0fa2cfbf180277d81bd578bbca17"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

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
check_not_present = {}
```

<a id="canonical-d43ec8b87f3a60276169826e2630bcec96a7a55000748c3d7d941f721f2cce8a"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.query_params.check_not_present / 2596f28e3966 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f494068c8ab25e3e27e453ed538a9685e69d244ae4230c16d5b7512682209bc1"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.query_params.check_not_present / 2596f28e3966 / 4

- [policy_based_challenge.rule_list.rules.spec.query_params](resources--cdn_loadbalancer--reference--group-013.md#canonical-511830efa2a971e782cd6ba252499d4cf40f63d59a6550161b659d22def20550)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-05827f52f304ac421a5a95c55139e3d6ba110e156327463d0895f7a59fb1b5a1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d274894f805e8ecdb0bddf90e2e9b801ead8b7b7ccb6ed68be383dae8a277f4a"></a>

## policy_based_challenge.rule_list.rules.spec.query_params.check_present — policy_based_challenge.rule_list.rules.spec.query_params.check_present / d27e7c7749e6 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-e7e050a0106b5c243929bc99c717fe27390190239f5f0da6543c11a425e13c71)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-6b6b9bb417be0328d783de67330d570206c65ca27165054e0fe0e20523560695)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-c1694f8f786e00f1b8d062766cb92e3219566178dd23899b40bc7c9d82c517a1)
- [policy_based_challenge.rule_list.rules.spec.query_params](resources--cdn_loadbalancer--reference--group-013.md#canonical-511830efa2a971e782cd6ba252499d4cf40f63d59a6550161b659d22def20550)
- policy_based_challenge.rule_list.rules.spec.query_params.check_present

<a id="canonical-68e2961f0fe5794863da60dd7706e5d0ef88a8e2b7e63338df7d3084ff4c2276"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

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
check_present = {}
```

<a id="canonical-2f973c3fc4d624277fe260c7e8f24911a7cd583ca897eac8c7a11fd66c848b97"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.query_params.check_present / d27e7c7749e6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e7fae099304773b77884e1d5922e3c837fb113059e83c696d27c580902874b68"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.query_params.check_present / d27e7c7749e6 / 4

- [policy_based_challenge.rule_list.rules.spec.query_params](resources--cdn_loadbalancer--reference--group-013.md#canonical-511830efa2a971e782cd6ba252499d4cf40f63d59a6550161b659d22def20550)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-30798cc5aa1f9e99534033d31c0052dcd28e1cfeedce619c21e3025633eed62f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-38cb6525a798a26d02065a48677ed390d78126099678fb0754ab4a5324dfbcc4"></a>

## policy_based_challenge.rule_list.rules.spec.query_params.item — policy_based_challenge.rule_list.rules.spec.query_params.item / 284418b93dcf / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-e7e050a0106b5c243929bc99c717fe27390190239f5f0da6543c11a425e13c71)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-6b6b9bb417be0328d783de67330d570206c65ca27165054e0fe0e20523560695)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-c1694f8f786e00f1b8d062766cb92e3219566178dd23899b40bc7c9d82c517a1)
- [policy_based_challenge.rule_list.rules.spec.query_params](resources--cdn_loadbalancer--reference--group-013.md#canonical-511830efa2a971e782cd6ba252499d4cf40f63d59a6550161b659d22def20550)
- policy_based_challenge.rule_list.rules.spec.query_params.item

<a id="canonical-123c0971e295da3d80ff8d566e86cb5bd44a03a65fe4f5ace2d94be07ba3399c"></a>

Type: `"object"`. single nested block, Optional.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-ac5f87b5d6f63e21a21c47da7ddc6c0450f1532117b0f8ccc3e2749e5e41f040"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.query_params.item / 284418b93dcf / 3

<a id="canonical-b22f2205d45fd67abbd5963ac2794c2731a7fb2fe26b37458d9cfa0e8b5a1dfe"></a>
