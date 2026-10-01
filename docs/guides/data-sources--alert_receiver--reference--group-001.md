---
page_title: "xcsh_alert_receiver reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_alert_receiver reference."
---

# xcsh_alert_receiver reference

<a id="canonical-9e77ba4988d239ee5fcf37742214d6c3d97c59c318f7ff0d06d4ec6a09a48451"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bfb9661a7822faadd7ca20ca639f804ffa7c46bd72d20367462b280aefe574ec"></a>

## Property reference — Property reference / 4cbdf4076975 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)
- Property reference

<a id="canonical-2d86e2b358ce13738e695668dabbbaa040c96846af305f7732c6e1dc85b1431c"></a>

## Direct properties — Property reference / 4cbdf4076975 / 3

<a id="canonical-421eea1da04d79da081b9243d0fd89fbe326b4fa9ce5255e29774d0d1f8cd905"></a>

<a id="canonical-c67f7c0831e4237958718501138eee5e25632a8ec34e687c3b54eb8a92a9bc55"></a>

## annotations property — Property reference / 4cbdf4076975 / 4

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Upstream description:

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

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
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-4700e520e48dc7592155d8684cf4a1abb50926d1a28542bf9e78e5cc2182a4ea"></a>

<a id="canonical-baa5d8323ac6997793bcbed1fe74c94d2ff81e641d660bbc6b988c1dc7713215"></a>

## description property — Property reference / 4cbdf4076975 / 5

Type: `"string"`. Computed.

Description of the AlertReceiver.

Upstream description:

Human readable description for the object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1200,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1200
    },
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  }
}
```

- [email](data-sources--alert_receiver--reference--group-001.md#canonical-21a18a279214d46252b2df55a8691c76cca00b0caf61a1e7ee172c4192b4b500): complete subsection reference.

<a id="canonical-1d92dfe43c33e908d3c8c49e80caceaf2b3ec640fefbbe399f9b8a93f9a566ac"></a>

<a id="canonical-798085346fd813f8a39818a59e096a043c058cbbcc9f3580d7f27408c5a38004"></a>

## id property — Property reference / 4cbdf4076975 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-cfc5e46ab67b75ccf6545a65bd32fb15a303db48a9db5e71177da76868d9169c"></a>

<a id="canonical-93e78ccebe1769b908ff21cf0ddf7ee980c4175166dd0156edf84c2de94c9806"></a>

## labels property — Property reference / 4cbdf4076975 / 7

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

Upstream description:

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

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

<a id="canonical-b1ec5620420d90275faf76b278cca7375c1b3450f0d46cc3fb56a2878a5bc395"></a>

<a id="canonical-fa30ae5b4d8a64e879e967ebbeefd3b6aa127f18141a34b0c50f64c082ee33f2"></a>

## name property — Property reference / 4cbdf4076975 / 8

Type: `"string"`. Required.

Name of the AlertReceiver.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-f6b624e28670581791a33d7c46575852664962b7febbb890c68304849f9f97ae"></a>

<a id="canonical-97bd7d83fc090815d2dcefed5df521baa0965cf2358e790a4299b4393cc7809d"></a>

## namespace property — Property reference / 4cbdf4076975 / 9

Type: `"string"`. Required.

Namespace where the AlertReceiver exists.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

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

- [opsgenie](data-sources--alert_receiver--reference--group-001.md#canonical-65e9783467a272c8aef1f2fa069a34e95b84a3a47427010b1b359d211ecd801b): complete subsection reference.

- [pagerduty](data-sources--alert_receiver--reference--group-001.md#canonical-b4e0c565cef2aed583f5d66d89ae5511e98d68efca125edb833b6cf215d8c1e5): complete subsection reference.

- [slack](data-sources--alert_receiver--reference--group-001.md#canonical-cff83f266f4e4fce9d31bd4c15d3a18c9aac2fc1bb545e7a39e47b9c6c38343b): complete subsection reference.

- [sms](data-sources--alert_receiver--reference--group-001.md#canonical-5223bcadfc376ad1be835710290f970f343fcd54aaa3472835c3c119da43455c): complete subsection reference.

- [webhook](data-sources--alert_receiver--reference--group-001.md#canonical-30aebc9511b4f9fb141ec5e484c794819b7452f6fd388fd46ab47d5adfc81033): complete subsection reference.

<a id="canonical-1aedce1ba1b7743674b197092a084f5c0509c097b59a9277f81148e8654064bc"></a>

## All schema paths — Property reference / 4cbdf4076975 / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--alert_receiver--reference--group-001.md#canonical-421eea1da04d79da081b9243d0fd89fbe326b4fa9ce5255e29774d0d1f8cd905) |
| `description` | [description](data-sources--alert_receiver--reference--group-001.md#canonical-4700e520e48dc7592155d8684cf4a1abb50926d1a28542bf9e78e5cc2182a4ea) |
| `email` | [email](data-sources--alert_receiver--reference--group-001.md#canonical-34a38c9f017713444b2b336cbcb45ed19a4a47ba911fc8b6a9854e205d67e2ab) |
| `email.email` | [email.email](data-sources--alert_receiver--reference--group-001.md#canonical-cbb29e463d1e4544b00e21e79b42b87b1ae442a5a7e217f9179506a621fbf05b) |
| `id` | [id](data-sources--alert_receiver--reference--group-001.md#canonical-1d92dfe43c33e908d3c8c49e80caceaf2b3ec640fefbbe399f9b8a93f9a566ac) |
| `labels` | [labels](data-sources--alert_receiver--reference--group-001.md#canonical-cfc5e46ab67b75ccf6545a65bd32fb15a303db48a9db5e71177da76868d9169c) |
| `name` | [name](data-sources--alert_receiver--reference--group-001.md#canonical-b1ec5620420d90275faf76b278cca7375c1b3450f0d46cc3fb56a2878a5bc395) |
| `namespace` | [namespace](data-sources--alert_receiver--reference--group-001.md#canonical-f6b624e28670581791a33d7c46575852664962b7febbb890c68304849f9f97ae) |
| `opsgenie` | [opsgenie](data-sources--alert_receiver--reference--group-001.md#canonical-9188a7fa450f59adbd4c7bc3efa56887c55557dc21da3a4531e17981501658a7) |
| `opsgenie.api_key` | [opsgenie.api_key](data-sources--alert_receiver--reference--group-001.md#canonical-5c3ae211464fb4ecf4e1185d6b82b95832742f55949c571261a483ba9d027ef7) |
| `opsgenie.api_key.blindfold_secret_info` | [opsgenie.api_key.blindfold_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-f0c9bb7acbe30ac9bee46f7113de7157126ea89407140d1702ea936ab70afe06) |
| `opsgenie.api_key.blindfold_secret_info.decryption_provider` | [opsgenie.api_key.blindfold_secret_info.decryption_provider](data-sources--alert_receiver--reference--group-001.md#canonical-f9eeb77d69ae2c01471dfeffa94e807a7462bdad79c758f758d3cd1d1e3cf823) |
| `opsgenie.api_key.blindfold_secret_info.location` | [opsgenie.api_key.blindfold_secret_info.location](data-sources--alert_receiver--reference--group-001.md#canonical-5e29dca403633c3348c48a25b2a755fd995139f74d27344db7db797c7c5d3dab) |
| `opsgenie.api_key.blindfold_secret_info.store_provider` | [opsgenie.api_key.blindfold_secret_info.store_provider](data-sources--alert_receiver--reference--group-001.md#canonical-3f25535f70c3322dd8012c35a014185b0e6022353a0a4152ec272d48a0bbebe5) |
| `opsgenie.api_key.clear_secret_info` | [opsgenie.api_key.clear_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-b097bce9fe88f2b8e2c6d869b6a99b47ce227006155eb46aaa6b19815a75a8dd) |
| `opsgenie.api_key.clear_secret_info.provider_ref` | [opsgenie.api_key.clear_secret_info.provider_ref](data-sources--alert_receiver--reference--group-001.md#canonical-752719336c8f6963a8769697c10763926ef7ac0aa1ed5ea98a99c017240d0c7f) |
| `opsgenie.api_key.clear_secret_info.url` | [opsgenie.api_key.clear_secret_info.url](data-sources--alert_receiver--reference--group-001.md#canonical-7d05f56f51ea726b70ff1ecc4fd866664aca2ad6b0db3aa311125f06c090f1f4) |
| `opsgenie.url` | [opsgenie.url](data-sources--alert_receiver--reference--group-001.md#canonical-cf81f122ba8d756e2aab1cd7fc28fe18768538af229e2ffc7026d069b20a393b) |
| `pagerduty` | [pagerduty](data-sources--alert_receiver--reference--group-001.md#canonical-965e652601b561935e9028689a955132e1288a402da28629e4f44149482d20ad) |
| `pagerduty.routing_key` | [pagerduty.routing_key](data-sources--alert_receiver--reference--group-001.md#canonical-8de54a3e55395768513e0eef6a72dc72da8d5580e274b248d869185c77dfd46f) |
| `pagerduty.routing_key.blindfold_secret_info` | [pagerduty.routing_key.blindfold_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-8d797f3c4d556ff4b4ef73930a2b601aecf2de7a5aeb8230f2d2d8216df7d904) |
| `pagerduty.routing_key.blindfold_secret_info.decryption_provider` | [pagerduty.routing_key.blindfold_secret_info.decryption_provider](data-sources--alert_receiver--reference--group-001.md#canonical-9c70629b3a6c242a8739612fef0984f635d6686fc5e4278ed309546c177154b0) |
| `pagerduty.routing_key.blindfold_secret_info.location` | [pagerduty.routing_key.blindfold_secret_info.location](data-sources--alert_receiver--reference--group-001.md#canonical-f89f109fc58c475ca1ded888f89489d4e2563bbc3f43ba2e79208cd62f0bd69c) |
| `pagerduty.routing_key.blindfold_secret_info.store_provider` | [pagerduty.routing_key.blindfold_secret_info.store_provider](data-sources--alert_receiver--reference--group-001.md#canonical-d88615bb7e87df81e5624e52be9c3a4991d6126dc553429a827a98bab57bd01e) |
| `pagerduty.routing_key.clear_secret_info` | [pagerduty.routing_key.clear_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-b6120ffef09ec66d0a081aa70287459ba826670807ff7e48ae24a602a7d4ac8f) |
| `pagerduty.routing_key.clear_secret_info.provider_ref` | [pagerduty.routing_key.clear_secret_info.provider_ref](data-sources--alert_receiver--reference--group-001.md#canonical-64f7f978c7c8b2932d3be08734e6eab5aaa33f2eb7a4e7ecc818504e5697a9ee) |
| `pagerduty.routing_key.clear_secret_info.url` | [pagerduty.routing_key.clear_secret_info.url](data-sources--alert_receiver--reference--group-001.md#canonical-e12381eb7cb85c0dbe24ee4826ccce27897b4dcb329ae802941c8901d2120c71) |
| `pagerduty.url` | [pagerduty.url](data-sources--alert_receiver--reference--group-001.md#canonical-a3f361d6f497b4e8fb5cd9f7616fc89f07a7cd1c1bd3f16c4eab2695674e01e3) |
| `slack` | [slack](data-sources--alert_receiver--reference--group-001.md#canonical-d2a9c6e4f693087a44c9d3ee31bd8dd2e795c83cb39c08a1a7c4b9e27e68f3b6) |
| `slack.channel` | [slack.channel](data-sources--alert_receiver--reference--group-001.md#canonical-af1eb9adecc523aa506e01cc55787111a7408f748eb3c392f4f43d4ab8e9a94e) |
| `slack.url` | [slack.url](data-sources--alert_receiver--reference--group-001.md#canonical-92c56b95bb6f9bb4b98d4109fbfdce09c48084b835a09ce49b3c8886262a0192) |
| `slack.url.blindfold_secret_info` | [slack.url.blindfold_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-5be77721918d4f46454c7f7beebbdd5bcf4ef6e093f5a2b3f09e453e8f67e745) |
| `slack.url.blindfold_secret_info.decryption_provider` | [slack.url.blindfold_secret_info.decryption_provider](data-sources--alert_receiver--reference--group-001.md#canonical-9f50322c26d0b07dce57d4a19be3ee0d3e8cff8976cd6a52135a2ed6619c2975) |
| `slack.url.blindfold_secret_info.location` | [slack.url.blindfold_secret_info.location](data-sources--alert_receiver--reference--group-001.md#canonical-794375f9ac2528cc1199c5457ad3a302ec1809fced3af4e805d78787d2aad553) |
| `slack.url.blindfold_secret_info.store_provider` | [slack.url.blindfold_secret_info.store_provider](data-sources--alert_receiver--reference--group-001.md#canonical-47690f942902aae596e2ca809d660e6bbd78b1b39134eae11874f8e7c0155de7) |
| `slack.url.clear_secret_info` | [slack.url.clear_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-72388240747287cce5c4721e50aeb3d14ca0150f186970d26c8e20d9b1760940) |
| `slack.url.clear_secret_info.provider_ref` | [slack.url.clear_secret_info.provider_ref](data-sources--alert_receiver--reference--group-001.md#canonical-7e52f0f5d45145a2d027c7eb65936c45c9d63b37a970118a0900284f93e72613) |
| `slack.url.clear_secret_info.url` | [slack.url.clear_secret_info.url](data-sources--alert_receiver--reference--group-001.md#canonical-cf361b5c0dd9e0c9f0a23d9e77e2c23bc308495d033f47951dd36f747ce7829e) |
| `sms` | [sms](data-sources--alert_receiver--reference--group-001.md#canonical-221dde72b0adcc56729e87ddc3e8f86053cd2bf697b136a82f25660cabe3b093) |
| `sms.contact_number` | [sms.contact_number](data-sources--alert_receiver--reference--group-001.md#canonical-b5cb315b1adb234e7d01a118d945494d04b28340cd27257e7ad3f3fb39683b27) |
| `webhook` | [webhook](data-sources--alert_receiver--reference--group-001.md#canonical-0a44aa6290abf8699c9f39eb4a1af4a6a092170646528f4a07d74bced6f33f5b) |
| `webhook.http_config` | [webhook.http_config](data-sources--alert_receiver--reference--group-001.md#canonical-ee4112eea95b72eda419a2befcd39264294a47a788c8d01f76ae72c84287094e) |
| `webhook.http_config.auth_token` | [webhook.http_config.auth_token](data-sources--alert_receiver--reference--group-001.md#canonical-1f7112246f71544744f2c8c3086d18e6e67ba5ecb9973b611be32768522844b3) |
| `webhook.http_config.auth_token.token` | [webhook.http_config.auth_token.token](data-sources--alert_receiver--reference--group-001.md#canonical-84019be0692022c576c9612618c63ca6c5ab8aa35dfd504391eea01b131764a4) |
| `webhook.http_config.auth_token.token.blindfold_secret_info` | [webhook.http_config.auth_token.token.blindfold_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-1d24d0169276b852136ac144e4d6fc3eb6c8752233cb368b0590c5b496dc8d69) |
| `webhook.http_config.auth_token.token.blindfold_secret_info.decryption_provider` | [webhook.http_config.auth_token.token.blindfold_secret_info.decryption_provider](data-sources--alert_receiver--reference--group-001.md#canonical-9313d54a132312dc01a3a209faa552b8af3e61f5cb3802820bda5b12c8b7ade5) |
| `webhook.http_config.auth_token.token.blindfold_secret_info.location` | [webhook.http_config.auth_token.token.blindfold_secret_info.location](data-sources--alert_receiver--reference--group-001.md#canonical-615d0bbdf3189025f717780890d24aa756719bf1bd29c7bfbe2c4d9f9a3152f4) |
| `webhook.http_config.auth_token.token.blindfold_secret_info.store_provider` | [webhook.http_config.auth_token.token.blindfold_secret_info.store_provider](data-sources--alert_receiver--reference--group-001.md#canonical-dee3fd25c6a14b27a9c20f1a3527fc969c84ac4c0b6da7ab757029202c3f6f72) |
| `webhook.http_config.auth_token.token.clear_secret_info` | [webhook.http_config.auth_token.token.clear_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-e4a296a6c0d8ea8b178a5c83d3fd1ae881d22e0bf22d819a1510101deb5e7cef) |
| `webhook.http_config.auth_token.token.clear_secret_info.provider_ref` | [webhook.http_config.auth_token.token.clear_secret_info.provider_ref](data-sources--alert_receiver--reference--group-001.md#canonical-ad423780cf601e080dcf9430471592d81a920771f7f7ff9f66a21bfc9024705e) |
| `webhook.http_config.auth_token.token.clear_secret_info.url` | [webhook.http_config.auth_token.token.clear_secret_info.url](data-sources--alert_receiver--reference--group-001.md#canonical-73db7f6fdec49a120e779fd296bc728a41e73cb915f4a419a367d754c55d275c) |
| `webhook.http_config.basic_auth` | [webhook.http_config.basic_auth](data-sources--alert_receiver--reference--group-001.md#canonical-6cc6fa6716de1db4574a4b71f1cb08d60b6d1547de8a62c9ab11da5f5ce5a326) |
| `webhook.http_config.basic_auth.password` | [webhook.http_config.basic_auth.password](data-sources--alert_receiver--reference--group-001.md#canonical-3d1627d087858744d3454ed0412ee68c8dd6248b17b0497b609dc82aee10e759) |
| `webhook.http_config.basic_auth.password.blindfold_secret_info` | [webhook.http_config.basic_auth.password.blindfold_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-3f6aad3fffad17e56901fa84d84be7c935cf485c4aee3a00d9114d50d74ab803) |
| `webhook.http_config.basic_auth.password.blindfold_secret_info.decryption_provider` | [webhook.http_config.basic_auth.password.blindfold_secret_info.decryption_provider](data-sources--alert_receiver--reference--group-001.md#canonical-6c50a94cb57a2a7c9f8c4d3c0c3e2cfc74e7c64ed1ceaf2b2073c2f2cc3b06ae) |
| `webhook.http_config.basic_auth.password.blindfold_secret_info.location` | [webhook.http_config.basic_auth.password.blindfold_secret_info.location](data-sources--alert_receiver--reference--group-001.md#canonical-99e795bf61ae2b41668180dd42043319d9b66b9fc3e11e8b552e1c1d43a1fe5b) |
| `webhook.http_config.basic_auth.password.blindfold_secret_info.store_provider` | [webhook.http_config.basic_auth.password.blindfold_secret_info.store_provider](data-sources--alert_receiver--reference--group-001.md#canonical-0522026ac797a9b1875a7859269768ac5d0a7643053293ea03efd307d43139b7) |
| `webhook.http_config.basic_auth.password.clear_secret_info` | [webhook.http_config.basic_auth.password.clear_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-9dacd42408f22b443e4dfba1a9c01e5615c99bbc507447aba0c634a9654a00f8) |
| `webhook.http_config.basic_auth.password.clear_secret_info.provider_ref` | [webhook.http_config.basic_auth.password.clear_secret_info.provider_ref](data-sources--alert_receiver--reference--group-001.md#canonical-9b24c1473bf19ead5a07c9c05ebd9a1816d74f456751eb9491ee43c05859d4bd) |
| `webhook.http_config.basic_auth.password.clear_secret_info.url` | [webhook.http_config.basic_auth.password.clear_secret_info.url](data-sources--alert_receiver--reference--group-001.md#canonical-48c18cb6b7dbafd3b67b45d136a05ca8f033857df6c575d6cafcd2c7d8461d7f) |
| `webhook.http_config.basic_auth.user_name` | [webhook.http_config.basic_auth.user_name](data-sources--alert_receiver--reference--group-001.md#canonical-046d2b94c6ab8fc01fa779a1b6309f93f67e6ecb850aa9a84e0d68c667a5af35) |
| `webhook.http_config.client_cert_obj` | [webhook.http_config.client_cert_obj](data-sources--alert_receiver--reference--group-001.md#canonical-7ebc31e4511e2f557d9ea77efa21c70b62019d518902e1b2429afbaa5b737a8c) |
| `webhook.http_config.client_cert_obj.use_tls_obj` | [webhook.http_config.client_cert_obj.use_tls_obj](data-sources--alert_receiver--reference--group-001.md#canonical-f881265b1d0a6b1868f008fb1204a930b0fbd2fa8fdf4ef2a1a9b61df8d074d2) |
| `webhook.http_config.client_cert_obj.use_tls_obj.kind` | [webhook.http_config.client_cert_obj.use_tls_obj.kind](data-sources--alert_receiver--reference--group-001.md#canonical-ef6cde67e3679c72560a151f2f1a1a2243e0400c4715b5cfda335669cf3cbc5d) |
| `webhook.http_config.client_cert_obj.use_tls_obj.name` | [webhook.http_config.client_cert_obj.use_tls_obj.name](data-sources--alert_receiver--reference--group-001.md#canonical-44561aaf32ead2357997bfff057b4bce403123ca1547b80a6fb26c4832d08bf1) |
| `webhook.http_config.client_cert_obj.use_tls_obj.namespace` | [webhook.http_config.client_cert_obj.use_tls_obj.namespace](data-sources--alert_receiver--reference--group-001.md#canonical-9279cc5495ab414539175f3ef3ef0aa2d1a6442351b299ea60215710f2b081ef) |
| `webhook.http_config.client_cert_obj.use_tls_obj.tenant` | [webhook.http_config.client_cert_obj.use_tls_obj.tenant](data-sources--alert_receiver--reference--group-001.md#canonical-7e5c1b31f67b3458c45f82aa5e1a1360e155447c1527b32af39ea3fbe1a45e3a) |
| `webhook.http_config.client_cert_obj.use_tls_obj.uid` | [webhook.http_config.client_cert_obj.use_tls_obj.uid](data-sources--alert_receiver--reference--group-001.md#canonical-c974accefef7b7e051c8e7615a0dab20df93e1d479350cc95f36fb3516d97855) |
| `webhook.http_config.enable_http2` | [webhook.http_config.enable_http2](data-sources--alert_receiver--reference--group-001.md#canonical-565db5caa1ca80edb0af3ca6455d8a970e31d723680de423af659ec7f5eeae83) |
| `webhook.http_config.follow_redirects` | [webhook.http_config.follow_redirects](data-sources--alert_receiver--reference--group-001.md#canonical-beb8c75a5df652d47acf7da04adf8c3b5f39b356256c1fa6ad099405a74861c9) |
| `webhook.http_config.no_authorization` | [webhook.http_config.no_authorization](data-sources--alert_receiver--reference--group-001.md#canonical-c2a37440e74b7f64e99bf679bb40db1f4343099a47f38f07751b9fd12321e1b4) |
| `webhook.http_config.no_tls` | [webhook.http_config.no_tls](data-sources--alert_receiver--reference--group-001.md#canonical-6c5737eb3eee895e3c1ad04122ec5034c507afab6eda3aa8bddd4e820dc44361) |
| `webhook.http_config.use_tls` | [webhook.http_config.use_tls](data-sources--alert_receiver--reference--group-001.md#canonical-729b072d40838e5c7d36a715c90775c8a1d840ac6e27957411545179e95ef2b2) |
| `webhook.http_config.use_tls.disable_sni` | [webhook.http_config.use_tls.disable_sni](data-sources--alert_receiver--reference--group-001.md#canonical-99d4f6113f19a90708113ca5f7bda3443f161f72da383fbd92a1b26c464d1fd0) |
| `webhook.http_config.use_tls.max_version` | [webhook.http_config.use_tls.max_version](data-sources--alert_receiver--reference--group-001.md#canonical-5d4e11a7bfda6940c73e70af127d516c22a02398afb0e5154dc9b652f42ee544) |
| `webhook.http_config.use_tls.min_version` | [webhook.http_config.use_tls.min_version](data-sources--alert_receiver--reference--group-001.md#canonical-291f3c42ffd1a27eca659999d7aa27f6842fa76974bae0affa2f926ab5b9de1b) |
| `webhook.http_config.use_tls.sni` | [webhook.http_config.use_tls.sni](data-sources--alert_receiver--reference--group-001.md#canonical-8712d781e530136e653a6c40cb964f03dd6a1e345534795b93162851fb4ff1ce) |
| `webhook.http_config.use_tls.use_server_verification` | [webhook.http_config.use_tls.use_server_verification](data-sources--alert_receiver--reference--group-001.md#canonical-465b85413b81ede1d04f5c1c334479e46f7253d9fc1704e687bfec0dd6eb9751) |
| `webhook.http_config.use_tls.use_server_verification.ca_cert_obj` | [webhook.http_config.use_tls.use_server_verification.ca_cert_obj](data-sources--alert_receiver--reference--group-001.md#canonical-ee8b09f5f34eb3d1b46bca10e727540fdb523a555df895b13a5ce6d4c3778702) |
| `webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca` | [webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca](data-sources--alert_receiver--reference--group-001.md#canonical-d2d169ee1b7eda85830e769b1547a8f1fbb4cf9142fa5b89906bc28d7dc76b91) |
| `webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.kind` | [webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.kind](data-sources--alert_receiver--reference--group-001.md#canonical-481b1e61b4e920db1c6b1995e18de57556b619b4c71abfbef7ee6c19103614a6) |
| `webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.name` | [webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.name](data-sources--alert_receiver--reference--group-001.md#canonical-83a983ce0d26742c590d1d581e6ebc04f2153d66b0989832d0e327b907f38548) |
| `webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.namespace` | [webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.namespace](data-sources--alert_receiver--reference--group-001.md#canonical-53cedae8daad4333f6eed81ab4c740c1506280ec821222d9a7c0126da95cc298) |
| `webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.tenant` | [webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.tenant](data-sources--alert_receiver--reference--group-001.md#canonical-4db14a10eb828f29e2b76bf34f1342a228d33868befb9f7d8615d62dad25e486) |
| `webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.uid` | [webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.uid](data-sources--alert_receiver--reference--group-001.md#canonical-ed04081d938f662fbb480581c1bbb6bc23d5c534a84218013c7f0b7aed9d8bcb) |
| `webhook.http_config.use_tls.volterra_trusted_ca` | [webhook.http_config.use_tls.volterra_trusted_ca](data-sources--alert_receiver--reference--group-001.md#canonical-73e01aaeff96fbd9a2cc7e800f0d42417758d1b69b647eebf135a1301d7c703d) |
| `webhook.url` | [webhook.url](data-sources--alert_receiver--reference--group-001.md#canonical-7cfa4beb05222e82c8dd8e38424802e6875c82bb910caf7f4b79bf1980188a08) |
| `webhook.url.blindfold_secret_info` | [webhook.url.blindfold_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-b47f27c116074fae0eba790c71838ae6f7b1966e39264ce6dcbb385fc0b3bbaa) |
| `webhook.url.blindfold_secret_info.decryption_provider` | [webhook.url.blindfold_secret_info.decryption_provider](data-sources--alert_receiver--reference--group-001.md#canonical-60f0e18396c8f71d41c02e73672eb87768dd510e209bfe735b11f76911c55625) |
| `webhook.url.blindfold_secret_info.location` | [webhook.url.blindfold_secret_info.location](data-sources--alert_receiver--reference--group-001.md#canonical-c78ec6b5d775bad47eee5b686573f9fc56c14eb9433f89f13e69211070bdb877) |
| `webhook.url.blindfold_secret_info.store_provider` | [webhook.url.blindfold_secret_info.store_provider](data-sources--alert_receiver--reference--group-001.md#canonical-573141ea5d363cf762557f840e75f6c085559cd467b2de96f350823fe60ba73d) |
| `webhook.url.clear_secret_info` | [webhook.url.clear_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-cf0e9fca418fd799d37d433219e8218a611b0e807e7bd999e0ae09ff3ff7c13e) |
| `webhook.url.clear_secret_info.provider_ref` | [webhook.url.clear_secret_info.provider_ref](data-sources--alert_receiver--reference--group-001.md#canonical-6ea16f29827b52624536b1049b77da1d047f5ffd35c83342b759feb4af5e8ce2) |
| `webhook.url.clear_secret_info.url` | [webhook.url.clear_secret_info.url](data-sources--alert_receiver--reference--group-001.md#canonical-bae1fcd386bd6e8577321aa31f78abc1d6db35e334e9f94ef89c91857bcc925e) |

<a id="canonical-d267f34fb16c90a0308c4318c0074d3f29fc53a2ec3998d1719a23369534801b"></a>

## Next pages — Property reference / 4cbdf4076975 / 11

- [email](data-sources--alert_receiver--reference--group-001.md#canonical-21a18a279214d46252b2df55a8691c76cca00b0caf61a1e7ee172c4192b4b500)
- [opsgenie](data-sources--alert_receiver--reference--group-001.md#canonical-65e9783467a272c8aef1f2fa069a34e95b84a3a47427010b1b359d211ecd801b)
- [pagerduty](data-sources--alert_receiver--reference--group-001.md#canonical-b4e0c565cef2aed583f5d66d89ae5511e98d68efca125edb833b6cf215d8c1e5)
- [slack](data-sources--alert_receiver--reference--group-001.md#canonical-cff83f266f4e4fce9d31bd4c15d3a18c9aac2fc1bb545e7a39e47b9c6c38343b)
- [sms](data-sources--alert_receiver--reference--group-001.md#canonical-5223bcadfc376ad1be835710290f970f343fcd54aaa3472835c3c119da43455c)
- [webhook](data-sources--alert_receiver--reference--group-001.md#canonical-30aebc9511b4f9fb141ec5e484c794819b7452f6fd388fd46ab47d5adfc81033)
- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)

<a id="canonical-21a18a279214d46252b2df55a8691c76cca00b0caf61a1e7ee172c4192b4b500"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-94f16a0b45e23431d1793bdfc5ff128911d28c478f82622a1004126e67856675"></a>

## email — email / 0d476e26cc5b / 2

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-9e77ba4988d239ee5fcf37742214d6c3d97c59c318f7ff0d06d4ec6a09a48451)
- email

<a id="canonical-34a38c9f017713444b2b336cbcb45ed19a4a47ba911fc8b6a9854e205d67e2ab"></a>

Type: `"single"`. Computed.

\[OneOf: email, opsgenie, pagerduty, slack, sms, webhook\] Email Configuration.

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

- [email](data-sources--alert_receiver--reference--group-001.md#canonical-34a38c9f017713444b2b336cbcb45ed19a4a47ba911fc8b6a9854e205d67e2ab)
- [opsgenie](data-sources--alert_receiver--reference--group-001.md#canonical-9188a7fa450f59adbd4c7bc3efa56887c55557dc21da3a4531e17981501658a7)
- [pagerduty](data-sources--alert_receiver--reference--group-001.md#canonical-965e652601b561935e9028689a955132e1288a402da28629e4f44149482d20ad)
- [slack](data-sources--alert_receiver--reference--group-001.md#canonical-d2a9c6e4f693087a44c9d3ee31bd8dd2e795c83cb39c08a1a7c4b9e27e68f3b6)
- [sms](data-sources--alert_receiver--reference--group-001.md#canonical-221dde72b0adcc56729e87ddc3e8f86053cd2bf697b136a82f25660cabe3b093)
- [webhook](data-sources--alert_receiver--reference--group-001.md#canonical-0a44aa6290abf8699c9f39eb4a1af4a6a092170646528f4a07d74bced6f33f5b)

Select alternatives according to the provider validators above.

<a id="canonical-94fb6037cd9a7cb8e3208cb7c9b32c6369953ebd94dcf78ceb52cf890a93eded"></a>

## Direct properties — email / 0d476e26cc5b / 3

<a id="canonical-cbb29e463d1e4544b00e21e79b42b87b1ae442a5a7e217f9179506a621fbf05b"></a>

<a id="canonical-17d22395b0ded1820feecfe7bde3d8d6721257e3da81db3ac006b7ecffd42c69"></a>

## email property — email / 0d476e26cc5b / 4

Type: `"string"`. Computed.

Email. Email ID of the user.

Upstream description:

Email ID of the user.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "email",
    "formatDescription": "RFC 5322 email address, max 254 characters",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 3,
    "pattern": "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$",
    "validation": {
      "rfc": "RFC 5322"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.email": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.email": "true"
  }
}
```

<a id="canonical-b3de49af011bb95d5fc1f32611b75d337fa52ce2c4ce2876e2a4a1bcdf7da841"></a>

## Next pages — email / 0d476e26cc5b / 5

- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-9e77ba4988d239ee5fcf37742214d6c3d97c59c318f7ff0d06d4ec6a09a48451)
- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)

<a id="canonical-65e9783467a272c8aef1f2fa069a34e95b84a3a47427010b1b359d211ecd801b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b2e737ffb8d0e859e0ae1b5abd23c6ac5e9e1573d26f40abf85c4d021182aaa4"></a>

## opsgenie — opsgenie / 3cfcc7644f02 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-9e77ba4988d239ee5fcf37742214d6c3d97c59c318f7ff0d06d4ec6a09a48451)
- opsgenie

<a id="canonical-9188a7fa450f59adbd4c7bc3efa56887c55557dc21da3a4531e17981501658a7"></a>

Type: `"single"`. Computed.

OpsGenie configuration to send alert notifications.

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

<a id="canonical-7fb5b6b2a76e07ba9655cd93c943b4eb784b8eefe86175db0d58f72b689cc29a"></a>

## Direct properties — opsgenie / 3cfcc7644f02 / 3

- [api_key](data-sources--alert_receiver--reference--group-001.md#canonical-50febb3a45a73e8c66d03fee1483a22c6282c898401e835da3771ea92eaf7f7a): complete subsection reference.

<a id="canonical-cf81f122ba8d756e2aab1cd7fc28fe18768538af229e2ffc7026d069b20a393b"></a>

<a id="canonical-a0d21bd48afe3e87c837150178825f76a92df04079cb7a9b8a24263b49b1dc27"></a>

## url property — opsgenie / 3cfcc7644f02 / 4

Type: `"string"`. Computed.

API URL. URL to send API requests to.

Upstream description:

URL to send API requests to.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 1024,
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

<a id="canonical-a2b576a8cd16c146caf73396716385a5c287380da1ed4206317a8211cdfd2771"></a>

## Next pages — opsgenie / 3cfcc7644f02 / 5

- [opsgenie.api_key](data-sources--alert_receiver--reference--group-001.md#canonical-50febb3a45a73e8c66d03fee1483a22c6282c898401e835da3771ea92eaf7f7a)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-9e77ba4988d239ee5fcf37742214d6c3d97c59c318f7ff0d06d4ec6a09a48451)
- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)

<a id="canonical-50febb3a45a73e8c66d03fee1483a22c6282c898401e835da3771ea92eaf7f7a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-00c89dc9e6e3087cac6cc15aea89131b63b4c74fdd63bae28270fc1ab6a59c0d"></a>

## opsgenie.api_key — opsgenie.api_key / da965390f988 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-9e77ba4988d239ee5fcf37742214d6c3d97c59c318f7ff0d06d4ec6a09a48451)
- [opsgenie](data-sources--alert_receiver--reference--group-001.md#canonical-65e9783467a272c8aef1f2fa069a34e95b84a3a47427010b1b359d211ecd801b)
- opsgenie.api_key

<a id="canonical-5c3ae211464fb4ecf4e1185d6b82b95832742f55949c571261a483ba9d027ef7"></a>

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

<a id="canonical-9c77c775077b771bd42685958d21d3aff11e598c89ff6b6091f93cccc700014a"></a>

## Direct properties — opsgenie.api_key / da965390f988 / 3

- [blindfold_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-9452888a8c2ea5716158def8912ae0a3e27d7c21c163ecce54370d1da2ae6abc): complete subsection reference.

- [clear_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-f744e6a0545baed4e4293e4f76274cfb97e809db31e543382d1f603e3efc0d1b): complete subsection reference.

<a id="canonical-370fd69c99248885f295073332e91f0fed0840fc9f4207a874ace9f94f29e2b1"></a>

## Next pages — opsgenie.api_key / da965390f988 / 4

- [opsgenie.api_key.blindfold_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-9452888a8c2ea5716158def8912ae0a3e27d7c21c163ecce54370d1da2ae6abc)
- [opsgenie.api_key.clear_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-f744e6a0545baed4e4293e4f76274cfb97e809db31e543382d1f603e3efc0d1b)
- [opsgenie](data-sources--alert_receiver--reference--group-001.md#canonical-65e9783467a272c8aef1f2fa069a34e95b84a3a47427010b1b359d211ecd801b)
- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)

<a id="canonical-9452888a8c2ea5716158def8912ae0a3e27d7c21c163ecce54370d1da2ae6abc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ea1230f3c8aa58240c10d254b2b5f27bddc0c81664d5dfe5b94c9cb5e92f18ec"></a>

## opsgenie.api_key.blindfold_secret_info — opsgenie.api_key.blindfold_secret_info / f3b04601b7ba / 2

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-9e77ba4988d239ee5fcf37742214d6c3d97c59c318f7ff0d06d4ec6a09a48451)
- [opsgenie](data-sources--alert_receiver--reference--group-001.md#canonical-65e9783467a272c8aef1f2fa069a34e95b84a3a47427010b1b359d211ecd801b)
- [opsgenie.api_key](data-sources--alert_receiver--reference--group-001.md#canonical-50febb3a45a73e8c66d03fee1483a22c6282c898401e835da3771ea92eaf7f7a)
- opsgenie.api_key.blindfold_secret_info

<a id="canonical-f0c9bb7acbe30ac9bee46f7113de7157126ea89407140d1702ea936ab70afe06"></a>

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

<a id="canonical-414a7a014b3b5ddf3504a3c1c78ac589efda85ef2c6e2607364da36ec9bd7902"></a>

## Direct properties — opsgenie.api_key.blindfold_secret_info / f3b04601b7ba / 3

<a id="canonical-f9eeb77d69ae2c01471dfeffa94e807a7462bdad79c758f758d3cd1d1e3cf823"></a>

<a id="canonical-1e25554fc67a55849a350b0061402059e43fe508098c3433c5a7b03dc58eb867"></a>

## decryption_provider property — opsgenie.api_key.blindfold_secret_info / f3b04601b7ba / 4

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

<a id="canonical-5e29dca403633c3348c48a25b2a755fd995139f74d27344db7db797c7c5d3dab"></a>

<a id="canonical-61354b5282dbdfe246b74418846297ed8762e70c09b664991b0c0773debf5c87"></a>

## location property — opsgenie.api_key.blindfold_secret_info / f3b04601b7ba / 5

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

<a id="canonical-3f25535f70c3322dd8012c35a014185b0e6022353a0a4152ec272d48a0bbebe5"></a>

<a id="canonical-094aeea7575ca503af91217248ec42d975c170e84c9c0fc5ba02ced2b61cd39e"></a>

## store_provider property — opsgenie.api_key.blindfold_secret_info / f3b04601b7ba / 6

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

<a id="canonical-cae5f5e620f0c88cac532546413a93090accec1c49e9a0101958a51aaf39db3a"></a>

## Next pages — opsgenie.api_key.blindfold_secret_info / f3b04601b7ba / 7

- [opsgenie.api_key](data-sources--alert_receiver--reference--group-001.md#canonical-50febb3a45a73e8c66d03fee1483a22c6282c898401e835da3771ea92eaf7f7a)
- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)

<a id="canonical-f744e6a0545baed4e4293e4f76274cfb97e809db31e543382d1f603e3efc0d1b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3686f55e1b87541724dbca5e3878925eac62f499a7d9dd8d07147240636d8d9f"></a>

## opsgenie.api_key.clear_secret_info — opsgenie.api_key.clear_secret_info / 4669e62b9378 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-9e77ba4988d239ee5fcf37742214d6c3d97c59c318f7ff0d06d4ec6a09a48451)
- [opsgenie](data-sources--alert_receiver--reference--group-001.md#canonical-65e9783467a272c8aef1f2fa069a34e95b84a3a47427010b1b359d211ecd801b)
- [opsgenie.api_key](data-sources--alert_receiver--reference--group-001.md#canonical-50febb3a45a73e8c66d03fee1483a22c6282c898401e835da3771ea92eaf7f7a)
- opsgenie.api_key.clear_secret_info

<a id="canonical-b097bce9fe88f2b8e2c6d869b6a99b47ce227006155eb46aaa6b19815a75a8dd"></a>

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

<a id="canonical-624183e2571725970d7a04df9d1c89479f80d744287b55871a3a7eabe07195c5"></a>

## Direct properties — opsgenie.api_key.clear_secret_info / 4669e62b9378 / 3

<a id="canonical-752719336c8f6963a8769697c10763926ef7ac0aa1ed5ea98a99c017240d0c7f"></a>

<a id="canonical-4560deb2675f820d2900499499a0c7cdd9de87d2f4f8e51015b1e0fd645113af"></a>

## provider_ref property — opsgenie.api_key.clear_secret_info / 4669e62b9378 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-7d05f56f51ea726b70ff1ecc4fd866664aca2ad6b0db3aa311125f06c090f1f4"></a>

<a id="canonical-f2eb8df26b04e8b985fecb57f4cbc706689a2a5ad78c3043d5ee6d375175d8ac"></a>

## url property — opsgenie.api_key.clear_secret_info / 4669e62b9378 / 5

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

<a id="canonical-41a77f60695044c1e83c150453af73a6cbf4d7b28ba939341c339044656037d1"></a>

## Next pages — opsgenie.api_key.clear_secret_info / 4669e62b9378 / 6

- [opsgenie.api_key](data-sources--alert_receiver--reference--group-001.md#canonical-50febb3a45a73e8c66d03fee1483a22c6282c898401e835da3771ea92eaf7f7a)
- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)

<a id="canonical-b4e0c565cef2aed583f5d66d89ae5511e98d68efca125edb833b6cf215d8c1e5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-65305831ed87fb20c216a50e4deaf9dfd2ff855c2397ae1cc263bdfa46090c6b"></a>

## pagerduty — pagerduty / acbb830cfae3 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-9e77ba4988d239ee5fcf37742214d6c3d97c59c318f7ff0d06d4ec6a09a48451)
- pagerduty

<a id="canonical-965e652601b561935e9028689a955132e1288a402da28629e4f44149482d20ad"></a>

Type: `"single"`. Computed.

PagerDuty configuration to send alert notifications.

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

<a id="canonical-57ed8b3f06dfa0f9db0bd9e5355551002fea2de16b9f24a3690453714102a8ba"></a>

## Direct properties — pagerduty / acbb830cfae3 / 3

- [routing_key](data-sources--alert_receiver--reference--group-001.md#canonical-860247bb42b454c9eaf1b6b422788a7f83f3f9a78e222c7d28d6147b584ebe93): complete subsection reference.

<a id="canonical-a3f361d6f497b4e8fb5cd9f7616fc89f07a7cd1c1bd3f16c4eab2695674e01e3"></a>

<a id="canonical-332441a04a097cc40833e6c98eb8032719a8b32ba2ff068f1a62bb76898ab3ee"></a>

## url property — pagerduty / acbb830cfae3 / 4

Type: `"string"`. Computed.

Pager Duty URL. URL to send API requests to.

Upstream description:

URL to send API requests to.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 1024,
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

<a id="canonical-908000420d894932d4a108be46be717da3c49b54b6e49153ecb7d54515bff1d6"></a>

## Next pages — pagerduty / acbb830cfae3 / 5

- [pagerduty.routing_key](data-sources--alert_receiver--reference--group-001.md#canonical-860247bb42b454c9eaf1b6b422788a7f83f3f9a78e222c7d28d6147b584ebe93)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-9e77ba4988d239ee5fcf37742214d6c3d97c59c318f7ff0d06d4ec6a09a48451)
- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)

<a id="canonical-860247bb42b454c9eaf1b6b422788a7f83f3f9a78e222c7d28d6147b584ebe93"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-49993972855cd4e3a5fa73569057d976ed1bd0e0e3ee7c9912070838735999d2"></a>

## pagerduty.routing_key — pagerduty.routing_key / c98ae4556869 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-9e77ba4988d239ee5fcf37742214d6c3d97c59c318f7ff0d06d4ec6a09a48451)
- [pagerduty](data-sources--alert_receiver--reference--group-001.md#canonical-b4e0c565cef2aed583f5d66d89ae5511e98d68efca125edb833b6cf215d8c1e5)
- pagerduty.routing_key

<a id="canonical-8de54a3e55395768513e0eef6a72dc72da8d5580e274b248d869185c77dfd46f"></a>

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

<a id="canonical-3843bd23a7f03457dac6c4b4d16d665ad2ca7311a1c22031295846017de12ece"></a>

## Direct properties — pagerduty.routing_key / c98ae4556869 / 3

- [blindfold_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-82b73e21145ee98cb94ae9e4649a2e468f3753d437b1508b6237d320056869f0): complete subsection reference.

- [clear_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-716811bb044d8bfda2d9a2e3f64284b2e4353b20080e5af831a1bc4f9c91b6ec): complete subsection reference.

<a id="canonical-46668d30ef646398459cdebd4216dcd9f2745e08b263131f5cbed83d866bedaf"></a>

## Next pages — pagerduty.routing_key / c98ae4556869 / 4

- [pagerduty.routing_key.blindfold_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-82b73e21145ee98cb94ae9e4649a2e468f3753d437b1508b6237d320056869f0)
- [pagerduty.routing_key.clear_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-716811bb044d8bfda2d9a2e3f64284b2e4353b20080e5af831a1bc4f9c91b6ec)
- [pagerduty](data-sources--alert_receiver--reference--group-001.md#canonical-b4e0c565cef2aed583f5d66d89ae5511e98d68efca125edb833b6cf215d8c1e5)
- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)

<a id="canonical-82b73e21145ee98cb94ae9e4649a2e468f3753d437b1508b6237d320056869f0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1c4b42de5f003b851acb4564ae803e1cc80b615a53daac4874f394097c97d632"></a>

## pagerduty.routing_key.blindfold_secret_info — pagerduty.routing_key.blindfold_secret_info / eb451d92a2c2 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-9e77ba4988d239ee5fcf37742214d6c3d97c59c318f7ff0d06d4ec6a09a48451)
- [pagerduty](data-sources--alert_receiver--reference--group-001.md#canonical-b4e0c565cef2aed583f5d66d89ae5511e98d68efca125edb833b6cf215d8c1e5)
- [pagerduty.routing_key](data-sources--alert_receiver--reference--group-001.md#canonical-860247bb42b454c9eaf1b6b422788a7f83f3f9a78e222c7d28d6147b584ebe93)
- pagerduty.routing_key.blindfold_secret_info

<a id="canonical-8d797f3c4d556ff4b4ef73930a2b601aecf2de7a5aeb8230f2d2d8216df7d904"></a>

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

<a id="canonical-86635be4c064cb588a9027efd98cdeec07bd80178efd40765102b2097baade03"></a>

## Direct properties — pagerduty.routing_key.blindfold_secret_info / eb451d92a2c2 / 3

<a id="canonical-9c70629b3a6c242a8739612fef0984f635d6686fc5e4278ed309546c177154b0"></a>

<a id="canonical-209cc9d3515e86ee48a378a5bf56d40ee7cbdfb1ad195d66aab38fb0240f0920"></a>

## decryption_provider property — pagerduty.routing_key.blindfold_secret_info / eb451d92a2c2 / 4

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

<a id="canonical-f89f109fc58c475ca1ded888f89489d4e2563bbc3f43ba2e79208cd62f0bd69c"></a>

<a id="canonical-e54af21c5cfe4fe994fd729a63744db0a3555b0768c0ebfabba58c30f87acc18"></a>

## location property — pagerduty.routing_key.blindfold_secret_info / eb451d92a2c2 / 5

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

<a id="canonical-d88615bb7e87df81e5624e52be9c3a4991d6126dc553429a827a98bab57bd01e"></a>

<a id="canonical-6ec0a84970dd4322277637fd80960bd2cabfbea1e0b668428961cfa305b2a42d"></a>

## store_provider property — pagerduty.routing_key.blindfold_secret_info / eb451d92a2c2 / 6

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

<a id="canonical-46985da2783147955dcecf40308f5dcec73dee30f112ffae1fcff1e2581aaf05"></a>

## Next pages — pagerduty.routing_key.blindfold_secret_info / eb451d92a2c2 / 7

- [pagerduty.routing_key](data-sources--alert_receiver--reference--group-001.md#canonical-860247bb42b454c9eaf1b6b422788a7f83f3f9a78e222c7d28d6147b584ebe93)
- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)

<a id="canonical-716811bb044d8bfda2d9a2e3f64284b2e4353b20080e5af831a1bc4f9c91b6ec"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-403fcf2becefc35f08556ca68e8430b41044d09528ff509beb3f777eb703ac5b"></a>

## pagerduty.routing_key.clear_secret_info — pagerduty.routing_key.clear_secret_info / d1b7d575ccab / 2

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-9e77ba4988d239ee5fcf37742214d6c3d97c59c318f7ff0d06d4ec6a09a48451)
- [pagerduty](data-sources--alert_receiver--reference--group-001.md#canonical-b4e0c565cef2aed583f5d66d89ae5511e98d68efca125edb833b6cf215d8c1e5)
- [pagerduty.routing_key](data-sources--alert_receiver--reference--group-001.md#canonical-860247bb42b454c9eaf1b6b422788a7f83f3f9a78e222c7d28d6147b584ebe93)
- pagerduty.routing_key.clear_secret_info

<a id="canonical-b6120ffef09ec66d0a081aa70287459ba826670807ff7e48ae24a602a7d4ac8f"></a>

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

<a id="canonical-3407ae411b5906db20cf20422b6b08e27bf20b700bfe42e3d9723560796d38b2"></a>

## Direct properties — pagerduty.routing_key.clear_secret_info / d1b7d575ccab / 3

<a id="canonical-64f7f978c7c8b2932d3be08734e6eab5aaa33f2eb7a4e7ecc818504e5697a9ee"></a>

<a id="canonical-49cf2156cdb4b7cd7b99dd640206993ebe936b0b028e8a7ed2aaabc687e4be6f"></a>

## provider_ref property — pagerduty.routing_key.clear_secret_info / d1b7d575ccab / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-e12381eb7cb85c0dbe24ee4826ccce27897b4dcb329ae802941c8901d2120c71"></a>

<a id="canonical-145857220ab7e676422f3548da488bb1fb0638f7bbc179a56a050046d3f92724"></a>

## url property — pagerduty.routing_key.clear_secret_info / d1b7d575ccab / 5

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

<a id="canonical-d309d20b368bffd1b541bb12797afb6d82545b378d6ee1212f4d472741796a69"></a>

## Next pages — pagerduty.routing_key.clear_secret_info / d1b7d575ccab / 6

- [pagerduty.routing_key](data-sources--alert_receiver--reference--group-001.md#canonical-860247bb42b454c9eaf1b6b422788a7f83f3f9a78e222c7d28d6147b584ebe93)
- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)

<a id="canonical-cff83f266f4e4fce9d31bd4c15d3a18c9aac2fc1bb545e7a39e47b9c6c38343b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-84a881a2871379491fd7fe5a0b579c0fee6cc4c0b753dfde6d3a8beb8b745d66"></a>

## slack — slack / 590e679ad568 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-9e77ba4988d239ee5fcf37742214d6c3d97c59c318f7ff0d06d4ec6a09a48451)
- slack

<a id="canonical-d2a9c6e4f693087a44c9d3ee31bd8dd2e795c83cb39c08a1a7c4b9e27e68f3b6"></a>

Type: `"single"`. Computed.

Slack configuration to send alert notifications.

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

<a id="canonical-ff1eb063c480cab8a4c2c1554b9ccc71bd1a8f7b1d1d9c3b44292384b6b1875d"></a>

## Direct properties — slack / 590e679ad568 / 3

<a id="canonical-af1eb9adecc523aa506e01cc55787111a7408f748eb3c392f4f43d4ab8e9a94e"></a>

<a id="canonical-fbaca8344e8ff5106ffd649c3fc44004f2674bdee89318f002b9dc032babd215"></a>

## channel property — slack / 590e679ad568 / 4

Type: `"string"`. Computed.

Channel or user to send notifications to.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^[a-z0-9-_]{1,80}$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^[a-z0-9-_]{1,80}$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^[a-z0-9-_]{1,80}$"
  }
}
```

- [url](data-sources--alert_receiver--reference--group-001.md#canonical-44a6aa9a6e3832651e5658306c8cb409ffdc5c3e5b48d45473fe64d8df83db4e): complete subsection reference.

<a id="canonical-bf4777b0c2cc88ce23859e720fb67da5bc13aecbc614904db17786e8011b8e7b"></a>

## Next pages — slack / 590e679ad568 / 5

- [slack.url](data-sources--alert_receiver--reference--group-001.md#canonical-44a6aa9a6e3832651e5658306c8cb409ffdc5c3e5b48d45473fe64d8df83db4e)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-9e77ba4988d239ee5fcf37742214d6c3d97c59c318f7ff0d06d4ec6a09a48451)
- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)

<a id="canonical-44a6aa9a6e3832651e5658306c8cb409ffdc5c3e5b48d45473fe64d8df83db4e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-79626b5301a023d63a7b08aebae6a2269923502c2f6757f834a44addb00fdbb6"></a>

## slack.url — slack.url / 915699f71908 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-9e77ba4988d239ee5fcf37742214d6c3d97c59c318f7ff0d06d4ec6a09a48451)
- [slack](data-sources--alert_receiver--reference--group-001.md#canonical-cff83f266f4e4fce9d31bd4c15d3a18c9aac2fc1bb545e7a39e47b9c6c38343b)
- slack.url

<a id="canonical-92c56b95bb6f9bb4b98d4109fbfdce09c48084b835a09ce49b3c8886262a0192"></a>

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

<a id="canonical-ee1689e1c3dfcf8ee7c390d7448ec226f5bc4d2a051f19a56167c2fc8a834e0e"></a>

## Direct properties — slack.url / 915699f71908 / 3

- [blindfold_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-f356f8a1351aa6dd863533e72c704450130967a4b66cb052d37283387faae9ae): complete subsection reference.

- [clear_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-deaa1af5196a4fd96efefd11f94c3ed724e5f85c6b442fc0b8395130f48bbe48): complete subsection reference.

<a id="canonical-dcd12a526fa93f26d563549e15a198c4dd2720c29795a46d611da7da5bd53f06"></a>

## Next pages — slack.url / 915699f71908 / 4

- [slack.url.blindfold_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-f356f8a1351aa6dd863533e72c704450130967a4b66cb052d37283387faae9ae)
- [slack.url.clear_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-deaa1af5196a4fd96efefd11f94c3ed724e5f85c6b442fc0b8395130f48bbe48)
- [slack](data-sources--alert_receiver--reference--group-001.md#canonical-cff83f266f4e4fce9d31bd4c15d3a18c9aac2fc1bb545e7a39e47b9c6c38343b)
- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)

<a id="canonical-f356f8a1351aa6dd863533e72c704450130967a4b66cb052d37283387faae9ae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7d2df72ffc10c3799b5ed6f2c69cc3cf640711c7b37315a4a82631dc2e1b38d3"></a>

## slack.url.blindfold_secret_info — slack.url.blindfold_secret_info / 192dab4a742c / 2

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-9e77ba4988d239ee5fcf37742214d6c3d97c59c318f7ff0d06d4ec6a09a48451)
- [slack](data-sources--alert_receiver--reference--group-001.md#canonical-cff83f266f4e4fce9d31bd4c15d3a18c9aac2fc1bb545e7a39e47b9c6c38343b)
- [slack.url](data-sources--alert_receiver--reference--group-001.md#canonical-44a6aa9a6e3832651e5658306c8cb409ffdc5c3e5b48d45473fe64d8df83db4e)
- slack.url.blindfold_secret_info

<a id="canonical-5be77721918d4f46454c7f7beebbdd5bcf4ef6e093f5a2b3f09e453e8f67e745"></a>

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

<a id="canonical-d8778c04e6fe1daac231c9f411215d2dc3aa8c13748377a215a771c3f792fb16"></a>

## Direct properties — slack.url.blindfold_secret_info / 192dab4a742c / 3

<a id="canonical-9f50322c26d0b07dce57d4a19be3ee0d3e8cff8976cd6a52135a2ed6619c2975"></a>

<a id="canonical-0aca3a0df12237f1c47d8ef8e0b549fe1d2cb6d6608b4d842687e34e91003359"></a>

## decryption_provider property — slack.url.blindfold_secret_info / 192dab4a742c / 4

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

<a id="canonical-794375f9ac2528cc1199c5457ad3a302ec1809fced3af4e805d78787d2aad553"></a>

<a id="canonical-b42c253fe7ef148e2f869043dc1a4168b8f40d9080d351e76437105a4837253a"></a>

## location property — slack.url.blindfold_secret_info / 192dab4a742c / 5

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

<a id="canonical-47690f942902aae596e2ca809d660e6bbd78b1b39134eae11874f8e7c0155de7"></a>

<a id="canonical-3a8af177755fbf563a244d8a39e24bb0eae758e30afb08128dbb080875e4313b"></a>

## store_provider property — slack.url.blindfold_secret_info / 192dab4a742c / 6

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

<a id="canonical-a0040b4395ab509098b945a1de0fecd955d717c1f464a1864d2e440a745a8127"></a>

## Next pages — slack.url.blindfold_secret_info / 192dab4a742c / 7

- [slack.url](data-sources--alert_receiver--reference--group-001.md#canonical-44a6aa9a6e3832651e5658306c8cb409ffdc5c3e5b48d45473fe64d8df83db4e)
- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)

<a id="canonical-deaa1af5196a4fd96efefd11f94c3ed724e5f85c6b442fc0b8395130f48bbe48"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b9e854c790ac647534c7e5d20c0469476ea89f90047d1aa18e0a9d947824dfd9"></a>

## slack.url.clear_secret_info — slack.url.clear_secret_info / cec8cc592a12 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-9e77ba4988d239ee5fcf37742214d6c3d97c59c318f7ff0d06d4ec6a09a48451)
- [slack](data-sources--alert_receiver--reference--group-001.md#canonical-cff83f266f4e4fce9d31bd4c15d3a18c9aac2fc1bb545e7a39e47b9c6c38343b)
- [slack.url](data-sources--alert_receiver--reference--group-001.md#canonical-44a6aa9a6e3832651e5658306c8cb409ffdc5c3e5b48d45473fe64d8df83db4e)
- slack.url.clear_secret_info

<a id="canonical-72388240747287cce5c4721e50aeb3d14ca0150f186970d26c8e20d9b1760940"></a>

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

<a id="canonical-63b4ad65edefee149a2454421471f5ab289d9a9efed835a7e5080ab525148daa"></a>

## Direct properties — slack.url.clear_secret_info / cec8cc592a12 / 3

<a id="canonical-7e52f0f5d45145a2d027c7eb65936c45c9d63b37a970118a0900284f93e72613"></a>

<a id="canonical-2bf2e110462abe78ec480768f88825d2e24d8404d04d925ff8cd8bb234b582bd"></a>

## provider_ref property — slack.url.clear_secret_info / cec8cc592a12 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-cf361b5c0dd9e0c9f0a23d9e77e2c23bc308495d033f47951dd36f747ce7829e"></a>

<a id="canonical-036effc42cb9d545c1c25c2a86c90b56111e1228ab8af0c97da5805cf157a0c3"></a>

## url property — slack.url.clear_secret_info / cec8cc592a12 / 5

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

<a id="canonical-63072fa991df3176a282f22c3da0b2404bc229c34ca94e74df07898070cfcde9"></a>

## Next pages — slack.url.clear_secret_info / cec8cc592a12 / 6

- [slack.url](data-sources--alert_receiver--reference--group-001.md#canonical-44a6aa9a6e3832651e5658306c8cb409ffdc5c3e5b48d45473fe64d8df83db4e)
- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)

<a id="canonical-5223bcadfc376ad1be835710290f970f343fcd54aaa3472835c3c119da43455c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c2922d16c7e77bd23a2e182ba48e7e68275558787086217da07f4d9c89290168"></a>

## sms — sms / c82d18954b33 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-9e77ba4988d239ee5fcf37742214d6c3d97c59c318f7ff0d06d4ec6a09a48451)
- sms

<a id="canonical-221dde72b0adcc56729e87ddc3e8f86053cd2bf697b136a82f25660cabe3b093"></a>

Type: `"single"`. Computed.

SMS Configuration.

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

<a id="canonical-a7c64dfaa39df23b93d6c0f12634d632fde4171388ead497775b0988932bd054"></a>

## Direct properties — sms / c82d18954b33 / 3

<a id="canonical-b5cb315b1adb234e7d01a118d945494d04b28340cd27257e7ad3f3fb39683b27"></a>

<a id="canonical-58992a414d3b8da1c3a09cdabb1d014dfbe362ab1269bc13fe955cf3e65ec85b"></a>

## contact_number property — sms / c82d18954b33 / 4

Type: `"string"`. Computed.

Contact number of the user in ITU E.164 format \[+\]\[country code\]\[subscriber number including
area code\].

Upstream description:

Contact number of the user in ITU E.164 format \[+\]\[country code\]\[subscriber number including
area code\]

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
    "ves.io.schema.rules.string.phone_number": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.phone_number": "true"
  }
}
```

<a id="canonical-6481760f0885a46392c2f6f366dc935c497301c413b629a026d29a9b764d29c2"></a>

## Next pages — sms / c82d18954b33 / 5

- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-9e77ba4988d239ee5fcf37742214d6c3d97c59c318f7ff0d06d4ec6a09a48451)
- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)

<a id="canonical-30aebc9511b4f9fb141ec5e484c794819b7452f6fd388fd46ab47d5adfc81033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-be7f3388116cac307da3498089c94cfc22f80bdb471cf05a0e251b949e674655"></a>

## webhook — webhook / 3a737cd4e55d / 2

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-9e77ba4988d239ee5fcf37742214d6c3d97c59c318f7ff0d06d4ec6a09a48451)
- webhook

<a id="canonical-0a44aa6290abf8699c9f39eb4a1af4a6a092170646528f4a07d74bced6f33f5b"></a>

Type: `"single"`. Computed.

Webhook configuration to send alert notifications.

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

<a id="canonical-81d8228899ef418717dc2cc09ed2a288afdf611b6e56abe1d5b301ce1013931d"></a>

## Direct properties — webhook / 3a737cd4e55d / 3

- [http_config](data-sources--alert_receiver--reference--group-001.md#canonical-bd1503ebab64162c382b91de0969860a066f2e198443b22cffb2211a3a1c2806): complete subsection reference.

- [url](data-sources--alert_receiver--reference--group-001.md#canonical-f89f467192ddf8816aaf8302245110958f22034f09366862618892b0fe1cee24): complete subsection reference.

<a id="canonical-4ba5ddb787dee8222ae8b0ed4b772eb109f2ba71c703df1fe278e7107c1d70bf"></a>

## Next pages — webhook / 3a737cd4e55d / 4

- [webhook.http_config](data-sources--alert_receiver--reference--group-001.md#canonical-bd1503ebab64162c382b91de0969860a066f2e198443b22cffb2211a3a1c2806)
- [webhook.url](data-sources--alert_receiver--reference--group-001.md#canonical-f89f467192ddf8816aaf8302245110958f22034f09366862618892b0fe1cee24)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-9e77ba4988d239ee5fcf37742214d6c3d97c59c318f7ff0d06d4ec6a09a48451)
- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)

<a id="canonical-bd1503ebab64162c382b91de0969860a066f2e198443b22cffb2211a3a1c2806"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cc8decf5b48d0ef9369f986a1a20ad8a049d98e4f31e572a37d1fe2e1f7b7f8e"></a>

## webhook.http_config — webhook.http_config / 790a9ec468c1 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-9e77ba4988d239ee5fcf37742214d6c3d97c59c318f7ff0d06d4ec6a09a48451)
- [webhook](data-sources--alert_receiver--reference--group-001.md#canonical-30aebc9511b4f9fb141ec5e484c794819b7452f6fd388fd46ab47d5adfc81033)
- webhook.http_config

<a id="canonical-ee4112eea95b72eda419a2befcd39264294a47a788c8d01f76ae72c84287094e"></a>

Type: `"single"`. Computed.

HTTP Configuration. Configuration for HTTP endpoint.

Upstream description:

Configuration for HTTP endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-auth_choice": "[\"auth_token\",\"basic_auth\",\"client_cert_obj\",\"no_authorization\"]",
  "x-ves-oneof-field-tls_choice": "[\"no_tls\",\"use_tls\"]"
}
```

<a id="canonical-5696fd29dee80ed4490004176381b1b62e42de849ede42acdc11fd82284b1727"></a>

## Direct properties — webhook.http_config / 790a9ec468c1 / 3

- [auth_token](data-sources--alert_receiver--reference--group-001.md#canonical-e6f92ed9b481addd391586803b376fe0143fa75988acb60a3809f2057d411797): complete subsection reference.

- [basic_auth](data-sources--alert_receiver--reference--group-001.md#canonical-9a6e23edd96fafe5b73e2b56d6c7c7cc137d4c93429a9cbf3a04468d87224180): complete subsection reference.

- [client_cert_obj](data-sources--alert_receiver--reference--group-001.md#canonical-b5029d35648600fab9b94c3ccb71a08ddd0cba20638328b16b1ea0b8252d59ea): complete subsection reference.

<a id="canonical-565db5caa1ca80edb0af3ca6455d8a970e31d723680de423af659ec7f5eeae83"></a>

<a id="canonical-3b5a8815a55b01262a360b8bceaba7f5469624ad3fc2e1dfdf397ff495923b7e"></a>

## enable_http2 property — webhook.http_config / 790a9ec468c1 / 4

Type: `"bool"`. Computed.

Enable HTTP2. Configure to use HTTP2 protocol.

Upstream description:

Configure to use HTTP2 protocol.

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

<a id="canonical-beb8c75a5df652d47acf7da04adf8c3b5f39b356256c1fa6ad099405a74861c9"></a>

<a id="canonical-dcc7c759b13d3ecfc3bff7d4bf5232f3b594cc98c45d96d091301bd6a1436e0d"></a>

## follow_redirects property — webhook.http_config / 790a9ec468c1 / 5

Type: `"bool"`. Computed.

Configure whether HTTP requests follow HTTP 3xx redirects.

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

- [no_authorization](data-sources--alert_receiver--reference--group-001.md#canonical-b133fe7466a8b96fd9f051904e55567ee8670b63a3ee8727d6caf9d8e212f736): complete subsection reference.

- [no_tls](data-sources--alert_receiver--reference--group-001.md#canonical-feb6a539c8a5f94e27f46dca84bb8be937865d28434f8837643ab6bdf0219017): complete subsection reference.

- [use_tls](data-sources--alert_receiver--reference--group-001.md#canonical-f2ce5b5eff735fbfa6d1c5882e4d9c1701d3eff3daacee865fcb3b6c568cbade): complete subsection reference.

<a id="canonical-aa85509132a12133a82857718358fc72b085cdbb1339c0b2eb40dc62c86fdf66"></a>

## Next pages — webhook.http_config / 790a9ec468c1 / 6

- [webhook.http_config.auth_token](data-sources--alert_receiver--reference--group-001.md#canonical-e6f92ed9b481addd391586803b376fe0143fa75988acb60a3809f2057d411797)
- [webhook.http_config.basic_auth](data-sources--alert_receiver--reference--group-001.md#canonical-9a6e23edd96fafe5b73e2b56d6c7c7cc137d4c93429a9cbf3a04468d87224180)
- [webhook.http_config.client_cert_obj](data-sources--alert_receiver--reference--group-001.md#canonical-b5029d35648600fab9b94c3ccb71a08ddd0cba20638328b16b1ea0b8252d59ea)
- [webhook.http_config.no_authorization](data-sources--alert_receiver--reference--group-001.md#canonical-b133fe7466a8b96fd9f051904e55567ee8670b63a3ee8727d6caf9d8e212f736)
- [webhook.http_config.no_tls](data-sources--alert_receiver--reference--group-001.md#canonical-feb6a539c8a5f94e27f46dca84bb8be937865d28434f8837643ab6bdf0219017)
- [webhook.http_config.use_tls](data-sources--alert_receiver--reference--group-001.md#canonical-f2ce5b5eff735fbfa6d1c5882e4d9c1701d3eff3daacee865fcb3b6c568cbade)
- [webhook](data-sources--alert_receiver--reference--group-001.md#canonical-30aebc9511b4f9fb141ec5e484c794819b7452f6fd388fd46ab47d5adfc81033)
- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)

<a id="canonical-e6f92ed9b481addd391586803b376fe0143fa75988acb60a3809f2057d411797"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-20ffd89724f4d8b1a4f84126868f8270ee6cce94100c2863320b414490998970"></a>

## webhook.http_config.auth_token — webhook.http_config.auth_token / f78322c40e7d / 2

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-9e77ba4988d239ee5fcf37742214d6c3d97c59c318f7ff0d06d4ec6a09a48451)
- [webhook](data-sources--alert_receiver--reference--group-001.md#canonical-30aebc9511b4f9fb141ec5e484c794819b7452f6fd388fd46ab47d5adfc81033)
- [webhook.http_config](data-sources--alert_receiver--reference--group-001.md#canonical-bd1503ebab64162c382b91de0969860a066f2e198443b22cffb2211a3a1c2806)
- webhook.http_config.auth_token

<a id="canonical-1f7112246f71544744f2c8c3086d18e6e67ba5ecb9973b611be32768522844b3"></a>

Type: `"single"`. Computed.

Access Token. Authentication Token for access.

Upstream description:

Authentication Token for access.

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

<a id="canonical-39676487e71e5684fc34d9d0011521cfb5f183f1a24034092f8d7bac00f63091"></a>

## Direct properties — webhook.http_config.auth_token / f78322c40e7d / 3

- [token](data-sources--alert_receiver--reference--group-001.md#canonical-23fdfd2074a39d74414302495852c83232632c3dada55d74dbb6bb456a729b60): complete subsection reference.

<a id="canonical-93214f423c2ac8e400c14fd20ebab9e8994131596ccfe06e93973dd1cce0be18"></a>

## Next pages — webhook.http_config.auth_token / f78322c40e7d / 4

- [webhook.http_config.auth_token.token](data-sources--alert_receiver--reference--group-001.md#canonical-23fdfd2074a39d74414302495852c83232632c3dada55d74dbb6bb456a729b60)
- [webhook.http_config](data-sources--alert_receiver--reference--group-001.md#canonical-bd1503ebab64162c382b91de0969860a066f2e198443b22cffb2211a3a1c2806)
- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)

<a id="canonical-23fdfd2074a39d74414302495852c83232632c3dada55d74dbb6bb456a729b60"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4cb6a422f6efdf08ae8ae5628c03d510960c7fff403305d3c74f54be02870f36"></a>

## webhook.http_config.auth_token.token — webhook.http_config.auth_token.token / b5541dc5222a / 2

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-9e77ba4988d239ee5fcf37742214d6c3d97c59c318f7ff0d06d4ec6a09a48451)
- [webhook](data-sources--alert_receiver--reference--group-001.md#canonical-30aebc9511b4f9fb141ec5e484c794819b7452f6fd388fd46ab47d5adfc81033)
- [webhook.http_config](data-sources--alert_receiver--reference--group-001.md#canonical-bd1503ebab64162c382b91de0969860a066f2e198443b22cffb2211a3a1c2806)
- [webhook.http_config.auth_token](data-sources--alert_receiver--reference--group-001.md#canonical-e6f92ed9b481addd391586803b376fe0143fa75988acb60a3809f2057d411797)
- webhook.http_config.auth_token.token

<a id="canonical-84019be0692022c576c9612618c63ca6c5ab8aa35dfd504391eea01b131764a4"></a>

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

<a id="canonical-73d4d2845438b46ae074331530e0b67728d2e32d630c379ba45aaf9eb063df7a"></a>

## Direct properties — webhook.http_config.auth_token.token / b5541dc5222a / 3

- [blindfold_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-cfc789b15742eb6cc40ac6e818a88ebaaffe04b8fe1085c43b8ff2a653ea16ea): complete subsection reference.

- [clear_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-24ff88825df754c1393901cf0d8776573b7afb1c07c8f81567b3128d713803e7): complete subsection reference.

<a id="canonical-9e505e2ef75f877f4e652e55d17122948064d91d1cc1ba5c49c4ab874e45de54"></a>

## Next pages — webhook.http_config.auth_token.token / b5541dc5222a / 4

- [webhook.http_config.auth_token.token.blindfold_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-cfc789b15742eb6cc40ac6e818a88ebaaffe04b8fe1085c43b8ff2a653ea16ea)
- [webhook.http_config.auth_token.token.clear_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-24ff88825df754c1393901cf0d8776573b7afb1c07c8f81567b3128d713803e7)
- [webhook.http_config.auth_token](data-sources--alert_receiver--reference--group-001.md#canonical-e6f92ed9b481addd391586803b376fe0143fa75988acb60a3809f2057d411797)
- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)

<a id="canonical-cfc789b15742eb6cc40ac6e818a88ebaaffe04b8fe1085c43b8ff2a653ea16ea"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-96b302421038966e8b0fe50445329b9e238ab0b6e232f2bc88dc36e7a55f090b"></a>

## webhook.http_config.auth_token.token.blindfold_secret_info — webhook.http_config.auth_token.token.blindfold_secret_info / 4e743c52e541 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-9e77ba4988d239ee5fcf37742214d6c3d97c59c318f7ff0d06d4ec6a09a48451)
- [webhook](data-sources--alert_receiver--reference--group-001.md#canonical-30aebc9511b4f9fb141ec5e484c794819b7452f6fd388fd46ab47d5adfc81033)
- [webhook.http_config](data-sources--alert_receiver--reference--group-001.md#canonical-bd1503ebab64162c382b91de0969860a066f2e198443b22cffb2211a3a1c2806)
- [webhook.http_config.auth_token](data-sources--alert_receiver--reference--group-001.md#canonical-e6f92ed9b481addd391586803b376fe0143fa75988acb60a3809f2057d411797)
- [webhook.http_config.auth_token.token](data-sources--alert_receiver--reference--group-001.md#canonical-23fdfd2074a39d74414302495852c83232632c3dada55d74dbb6bb456a729b60)
- webhook.http_config.auth_token.token.blindfold_secret_info

<a id="canonical-1d24d0169276b852136ac144e4d6fc3eb6c8752233cb368b0590c5b496dc8d69"></a>

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

<a id="canonical-7dbb0d086066c75a20bfc81f470a9df2b14366815abcc60b033a3d60bd019614"></a>

## Direct properties — webhook.http_config.auth_token.token.blindfold_secret_info / 4e743c52e541 / 3

<a id="canonical-9313d54a132312dc01a3a209faa552b8af3e61f5cb3802820bda5b12c8b7ade5"></a>

<a id="canonical-5b952ac33f19985484d6e5621db479ce7fc63719b9f43193b3c04c349738d8e7"></a>

## decryption_provider property — webhook.http_config.auth_token.token.blindfold_secret_info / 4e743c52e541 / 4

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

<a id="canonical-615d0bbdf3189025f717780890d24aa756719bf1bd29c7bfbe2c4d9f9a3152f4"></a>

<a id="canonical-99416a68710a1e745c740866b97188a980d10774662526be75587231e0cb8984"></a>

## location property — webhook.http_config.auth_token.token.blindfold_secret_info / 4e743c52e541 / 5

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

<a id="canonical-dee3fd25c6a14b27a9c20f1a3527fc969c84ac4c0b6da7ab757029202c3f6f72"></a>

<a id="canonical-8f72c0fc7350893b672c44b51e3e3551253ed5c5f39ef77cecc044efca476181"></a>

## store_provider property — webhook.http_config.auth_token.token.blindfold_secret_info / 4e743c52e541 / 6

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

<a id="canonical-4e764a51491b13e34921347bfcb51c177efc021504336f76c3e0d0c90102df98"></a>

## Next pages — webhook.http_config.auth_token.token.blindfold_secret_info / 4e743c52e541 / 7

- [webhook.http_config.auth_token.token](data-sources--alert_receiver--reference--group-001.md#canonical-23fdfd2074a39d74414302495852c83232632c3dada55d74dbb6bb456a729b60)
- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)

<a id="canonical-24ff88825df754c1393901cf0d8776573b7afb1c07c8f81567b3128d713803e7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a34214490aebcade59c970f4fb0ca195cb565dee15ae6c0711ca887057cf0909"></a>

## webhook.http_config.auth_token.token.clear_secret_info — webhook.http_config.auth_token.token.clear_secret_info / 79d5cb99a3a3 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-9e77ba4988d239ee5fcf37742214d6c3d97c59c318f7ff0d06d4ec6a09a48451)
- [webhook](data-sources--alert_receiver--reference--group-001.md#canonical-30aebc9511b4f9fb141ec5e484c794819b7452f6fd388fd46ab47d5adfc81033)
- [webhook.http_config](data-sources--alert_receiver--reference--group-001.md#canonical-bd1503ebab64162c382b91de0969860a066f2e198443b22cffb2211a3a1c2806)
- [webhook.http_config.auth_token](data-sources--alert_receiver--reference--group-001.md#canonical-e6f92ed9b481addd391586803b376fe0143fa75988acb60a3809f2057d411797)
- [webhook.http_config.auth_token.token](data-sources--alert_receiver--reference--group-001.md#canonical-23fdfd2074a39d74414302495852c83232632c3dada55d74dbb6bb456a729b60)
- webhook.http_config.auth_token.token.clear_secret_info

<a id="canonical-e4a296a6c0d8ea8b178a5c83d3fd1ae881d22e0bf22d819a1510101deb5e7cef"></a>

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

<a id="canonical-3d8c965cb7e023f19c7abca4a73c27aa10d6fa7d4b3189b1947bdeaf02abe0ee"></a>

## Direct properties — webhook.http_config.auth_token.token.clear_secret_info / 79d5cb99a3a3 / 3

<a id="canonical-ad423780cf601e080dcf9430471592d81a920771f7f7ff9f66a21bfc9024705e"></a>

<a id="canonical-9fa0bd42e6e7e7420aebed79c75052d00a2f01ac03d3005feaaeeff75c7400ab"></a>

## provider_ref property — webhook.http_config.auth_token.token.clear_secret_info / 79d5cb99a3a3 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-73db7f6fdec49a120e779fd296bc728a41e73cb915f4a419a367d754c55d275c"></a>

<a id="canonical-7b457642bc4980e37ef3cd4f13490867603667eec8a82d79325a0e82c3b80f41"></a>

## url property — webhook.http_config.auth_token.token.clear_secret_info / 79d5cb99a3a3 / 5

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

<a id="canonical-9aa43077c67163d5c40e60b8a8eb7251074100eb0368657374e3e7b795ae9af0"></a>

## Next pages — webhook.http_config.auth_token.token.clear_secret_info / 79d5cb99a3a3 / 6

- [webhook.http_config.auth_token.token](data-sources--alert_receiver--reference--group-001.md#canonical-23fdfd2074a39d74414302495852c83232632c3dada55d74dbb6bb456a729b60)
- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)

<a id="canonical-9a6e23edd96fafe5b73e2b56d6c7c7cc137d4c93429a9cbf3a04468d87224180"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-994b60bf19f099c4e9d0d0de1ce0ece755515f5e4d547c7b68dc98f3aa54c157"></a>

## webhook.http_config.basic_auth — webhook.http_config.basic_auth / e1cb4217e4d0 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-9e77ba4988d239ee5fcf37742214d6c3d97c59c318f7ff0d06d4ec6a09a48451)
- [webhook](data-sources--alert_receiver--reference--group-001.md#canonical-30aebc9511b4f9fb141ec5e484c794819b7452f6fd388fd46ab47d5adfc81033)
- [webhook.http_config](data-sources--alert_receiver--reference--group-001.md#canonical-bd1503ebab64162c382b91de0969860a066f2e198443b22cffb2211a3a1c2806)
- webhook.http_config.basic_auth

<a id="canonical-6cc6fa6716de1db4574a4b71f1cb08d60b6d1547de8a62c9ab11da5f5ce5a326"></a>

Type: `"single"`. Computed.

Authorization parameters to access HTPP alert Receiver Endpoint.

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

<a id="canonical-fff3aa3aa7735279afb111be64b308fbca43abdd5828cba69e944255e480d7b8"></a>

## Direct properties — webhook.http_config.basic_auth / e1cb4217e4d0 / 3

- [password](data-sources--alert_receiver--reference--group-001.md#canonical-c5b88d3ccd5a43f4ce7843b257066be075ddec5eebb4cead0318a37c6fe5c467): complete subsection reference.

<a id="canonical-046d2b94c6ab8fc01fa779a1b6309f93f67e6ecb850aa9a84e0d68c667a5af35"></a>

<a id="canonical-b6b300effae84c9955b1dc410b99caafc3f3025d4b958a95d6903dd3453b4ec1"></a>

## user_name property — webhook.http_config.basic_auth / e1cb4217e4d0 / 4

Type: `"string"`. Computed.

User Name. HTTP Basic Auth User Name.

Upstream description:

HTTP Basic Auth User Name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-22b283c3b97da3b872ec9731d89f46450eacf5cbefa88a3cd64817858faa229c"></a>

## Next pages — webhook.http_config.basic_auth / e1cb4217e4d0 / 5

- [webhook.http_config.basic_auth.password](data-sources--alert_receiver--reference--group-001.md#canonical-c5b88d3ccd5a43f4ce7843b257066be075ddec5eebb4cead0318a37c6fe5c467)
- [webhook.http_config](data-sources--alert_receiver--reference--group-001.md#canonical-bd1503ebab64162c382b91de0969860a066f2e198443b22cffb2211a3a1c2806)
- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)

<a id="canonical-c5b88d3ccd5a43f4ce7843b257066be075ddec5eebb4cead0318a37c6fe5c467"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d2264722959b8973e166b423646444e403fa62a3128f35b2c45b51fb3bede271"></a>

## webhook.http_config.basic_auth.password — webhook.http_config.basic_auth.password / 129fa0da7339 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-9e77ba4988d239ee5fcf37742214d6c3d97c59c318f7ff0d06d4ec6a09a48451)
- [webhook](data-sources--alert_receiver--reference--group-001.md#canonical-30aebc9511b4f9fb141ec5e484c794819b7452f6fd388fd46ab47d5adfc81033)
- [webhook.http_config](data-sources--alert_receiver--reference--group-001.md#canonical-bd1503ebab64162c382b91de0969860a066f2e198443b22cffb2211a3a1c2806)
- [webhook.http_config.basic_auth](data-sources--alert_receiver--reference--group-001.md#canonical-9a6e23edd96fafe5b73e2b56d6c7c7cc137d4c93429a9cbf3a04468d87224180)
- webhook.http_config.basic_auth.password

<a id="canonical-3d1627d087858744d3454ed0412ee68c8dd6248b17b0497b609dc82aee10e759"></a>

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

<a id="canonical-2e8ceb60f78077eeefd71403c9f16f95409916a747406e9e30f0adc412611072"></a>

## Direct properties — webhook.http_config.basic_auth.password / 129fa0da7339 / 3

- [blindfold_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-f84378f5bf50d50c35e56dae47ec52a6946ce1f0464552c0b47eeac3cc349343): complete subsection reference.

- [clear_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-8e04c3d4840e5dc5e2662427842f08e6c00e472587e80813085e8f48d1fe37e2): complete subsection reference.

<a id="canonical-00843d2576fdd7f8fd9c93b9cdbcf78722a014e6b35754e258e3f1e3e1e032bb"></a>

## Next pages — webhook.http_config.basic_auth.password / 129fa0da7339 / 4

- [webhook.http_config.basic_auth.password.blindfold_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-f84378f5bf50d50c35e56dae47ec52a6946ce1f0464552c0b47eeac3cc349343)
- [webhook.http_config.basic_auth.password.clear_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-8e04c3d4840e5dc5e2662427842f08e6c00e472587e80813085e8f48d1fe37e2)
- [webhook.http_config.basic_auth](data-sources--alert_receiver--reference--group-001.md#canonical-9a6e23edd96fafe5b73e2b56d6c7c7cc137d4c93429a9cbf3a04468d87224180)
- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)

<a id="canonical-f84378f5bf50d50c35e56dae47ec52a6946ce1f0464552c0b47eeac3cc349343"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1cc13dbe333913f968274ab59743a73d3589da3f2709b02cff09263fc0601542"></a>

## webhook.http_config.basic_auth.password.blindfold_secret_info — webhook.http_config.basic_auth.password.blindfold_secret_info / e95acbe55a16 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-9e77ba4988d239ee5fcf37742214d6c3d97c59c318f7ff0d06d4ec6a09a48451)
- [webhook](data-sources--alert_receiver--reference--group-001.md#canonical-30aebc9511b4f9fb141ec5e484c794819b7452f6fd388fd46ab47d5adfc81033)
- [webhook.http_config](data-sources--alert_receiver--reference--group-001.md#canonical-bd1503ebab64162c382b91de0969860a066f2e198443b22cffb2211a3a1c2806)
- [webhook.http_config.basic_auth](data-sources--alert_receiver--reference--group-001.md#canonical-9a6e23edd96fafe5b73e2b56d6c7c7cc137d4c93429a9cbf3a04468d87224180)
- [webhook.http_config.basic_auth.password](data-sources--alert_receiver--reference--group-001.md#canonical-c5b88d3ccd5a43f4ce7843b257066be075ddec5eebb4cead0318a37c6fe5c467)
- webhook.http_config.basic_auth.password.blindfold_secret_info

<a id="canonical-3f6aad3fffad17e56901fa84d84be7c935cf485c4aee3a00d9114d50d74ab803"></a>

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

<a id="canonical-6311cd4f0087d0d6246b87216f23c3aa92d799020eacd749abdfee0b4c96fde1"></a>

## Direct properties — webhook.http_config.basic_auth.password.blindfold_secret_info / e95acbe55a16 / 3

<a id="canonical-6c50a94cb57a2a7c9f8c4d3c0c3e2cfc74e7c64ed1ceaf2b2073c2f2cc3b06ae"></a>

<a id="canonical-1b1ea75b15fcb1c54dc96115cf732c80d479cb72fc02a3ac140c96dd7cf395eb"></a>

## decryption_provider property — webhook.http_config.basic_auth.password.blindfold_secret_info / e95acbe55a16 / 4

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

<a id="canonical-99e795bf61ae2b41668180dd42043319d9b66b9fc3e11e8b552e1c1d43a1fe5b"></a>

<a id="canonical-24c598ac71e07400d7b76ae281daa907e38918e4e190b91f998ba91da65e78a7"></a>

## location property — webhook.http_config.basic_auth.password.blindfold_secret_info / e95acbe55a16 / 5

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

<a id="canonical-0522026ac797a9b1875a7859269768ac5d0a7643053293ea03efd307d43139b7"></a>

<a id="canonical-8f649d7d384db52d1e553d8c0b63fe4a371ecaaa7d37ab14167d9d00a07edfdc"></a>

## store_provider property — webhook.http_config.basic_auth.password.blindfold_secret_info / e95acbe55a16 / 6

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

<a id="canonical-3a7155ef77d86b2215ec28e858f592ebbec3a3b7cdc179f31723b375da104f25"></a>

## Next pages — webhook.http_config.basic_auth.password.blindfold_secret_info / e95acbe55a16 / 7

- [webhook.http_config.basic_auth.password](data-sources--alert_receiver--reference--group-001.md#canonical-c5b88d3ccd5a43f4ce7843b257066be075ddec5eebb4cead0318a37c6fe5c467)
- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)

<a id="canonical-8e04c3d4840e5dc5e2662427842f08e6c00e472587e80813085e8f48d1fe37e2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-56e34fd4567172e873ca57decf75da719f0d35d947bdcf06f341090541fb68f2"></a>

## webhook.http_config.basic_auth.password.clear_secret_info — webhook.http_config.basic_auth.password.clear_secret_info / 92985434187f / 2

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-9e77ba4988d239ee5fcf37742214d6c3d97c59c318f7ff0d06d4ec6a09a48451)
- [webhook](data-sources--alert_receiver--reference--group-001.md#canonical-30aebc9511b4f9fb141ec5e484c794819b7452f6fd388fd46ab47d5adfc81033)
- [webhook.http_config](data-sources--alert_receiver--reference--group-001.md#canonical-bd1503ebab64162c382b91de0969860a066f2e198443b22cffb2211a3a1c2806)
- [webhook.http_config.basic_auth](data-sources--alert_receiver--reference--group-001.md#canonical-9a6e23edd96fafe5b73e2b56d6c7c7cc137d4c93429a9cbf3a04468d87224180)
- [webhook.http_config.basic_auth.password](data-sources--alert_receiver--reference--group-001.md#canonical-c5b88d3ccd5a43f4ce7843b257066be075ddec5eebb4cead0318a37c6fe5c467)
- webhook.http_config.basic_auth.password.clear_secret_info

<a id="canonical-9dacd42408f22b443e4dfba1a9c01e5615c99bbc507447aba0c634a9654a00f8"></a>

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

<a id="canonical-9ac35a949ac03384279ad6ffa872a67bf71e76122d2adcbeb763b37e1eaa0ebb"></a>

## Direct properties — webhook.http_config.basic_auth.password.clear_secret_info / 92985434187f / 3

<a id="canonical-9b24c1473bf19ead5a07c9c05ebd9a1816d74f456751eb9491ee43c05859d4bd"></a>

<a id="canonical-ac664c1c30d3d71bf7e68be5135004692b05c3275da31019bf5825e14d2451f0"></a>

## provider_ref property — webhook.http_config.basic_auth.password.clear_secret_info / 92985434187f / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-48c18cb6b7dbafd3b67b45d136a05ca8f033857df6c575d6cafcd2c7d8461d7f"></a>

<a id="canonical-c2728117fd7523cdd753541c10fef2bcd8f4375f6b6247edcc7e3e1679dbccfd"></a>

## url property — webhook.http_config.basic_auth.password.clear_secret_info / 92985434187f / 5

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

<a id="canonical-b8527be747bdc8ff77db0375d4d7937b7d22ab38aad7a1a88e4bd5df6be593e1"></a>

## Next pages — webhook.http_config.basic_auth.password.clear_secret_info / 92985434187f / 6

- [webhook.http_config.basic_auth.password](data-sources--alert_receiver--reference--group-001.md#canonical-c5b88d3ccd5a43f4ce7843b257066be075ddec5eebb4cead0318a37c6fe5c467)
- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)

<a id="canonical-b5029d35648600fab9b94c3ccb71a08ddd0cba20638328b16b1ea0b8252d59ea"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cf4b9875556413b66e3cd661b3e1e8015345b16b9dfbcb5335ff95e6873ff4c4"></a>

## webhook.http_config.client_cert_obj — webhook.http_config.client_cert_obj / aa104c5415d9 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-9e77ba4988d239ee5fcf37742214d6c3d97c59c318f7ff0d06d4ec6a09a48451)
- [webhook](data-sources--alert_receiver--reference--group-001.md#canonical-30aebc9511b4f9fb141ec5e484c794819b7452f6fd388fd46ab47d5adfc81033)
- [webhook.http_config](data-sources--alert_receiver--reference--group-001.md#canonical-bd1503ebab64162c382b91de0969860a066f2e198443b22cffb2211a3a1c2806)
- webhook.http_config.client_cert_obj

<a id="canonical-7ebc31e4511e2f557d9ea77efa21c70b62019d518902e1b2429afbaa5b737a8c"></a>

Type: `"single"`. Computed.

Client Certificate Object. Configuration for client certificate.

Upstream description:

Configuration for client certificate.

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

<a id="canonical-f24824c1990b15abd71f188a1adc44dfd85f1ea029f8523dd914c1713a3d7718"></a>

## Direct properties — webhook.http_config.client_cert_obj / aa104c5415d9 / 3

- [use_tls_obj](data-sources--alert_receiver--reference--group-001.md#canonical-f5f248932bfd5be68863212eb9536189836fc36977020c543249a1d200648f70): complete subsection reference.

<a id="canonical-52be9095cb3549b933efbdbc3591e46a6842580b8b1986165902db6926e7e1df"></a>

## Next pages — webhook.http_config.client_cert_obj / aa104c5415d9 / 4

- [webhook.http_config.client_cert_obj.use_tls_obj](data-sources--alert_receiver--reference--group-001.md#canonical-f5f248932bfd5be68863212eb9536189836fc36977020c543249a1d200648f70)
- [webhook.http_config](data-sources--alert_receiver--reference--group-001.md#canonical-bd1503ebab64162c382b91de0969860a066f2e198443b22cffb2211a3a1c2806)
- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)

<a id="canonical-f5f248932bfd5be68863212eb9536189836fc36977020c543249a1d200648f70"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6a376d6983ac19b425dc901e9597aaa5d8100f9c5346fffa802dfec43da93faa"></a>

## webhook.http_config.client_cert_obj.use_tls_obj — webhook.http_config.client_cert_obj.use_tls_obj / 825ec15df4b2 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-9e77ba4988d239ee5fcf37742214d6c3d97c59c318f7ff0d06d4ec6a09a48451)
- [webhook](data-sources--alert_receiver--reference--group-001.md#canonical-30aebc9511b4f9fb141ec5e484c794819b7452f6fd388fd46ab47d5adfc81033)
- [webhook.http_config](data-sources--alert_receiver--reference--group-001.md#canonical-bd1503ebab64162c382b91de0969860a066f2e198443b22cffb2211a3a1c2806)
- [webhook.http_config.client_cert_obj](data-sources--alert_receiver--reference--group-001.md#canonical-b5029d35648600fab9b94c3ccb71a08ddd0cba20638328b16b1ea0b8252d59ea)
- webhook.http_config.client_cert_obj.use_tls_obj

<a id="canonical-f881265b1d0a6b1868f008fb1204a930b0fbd2fa8fdf4ef2a1a9b61df8d074d2"></a>

Type: `"list"`. Computed.

Certificate Object. Reference to client certificate object.

Upstream description:

Reference to client certificate object.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-9f6994ae3c1a87d64889abf06bc96c1bad4e01ef31768c83c79bde5c8e51f633"></a>

## Direct properties — webhook.http_config.client_cert_obj.use_tls_obj / 825ec15df4b2 / 3

<a id="canonical-ef6cde67e3679c72560a151f2f1a1a2243e0400c4715b5cfda335669cf3cbc5d"></a>

<a id="canonical-99ea6421769f31281307c3e4f6700cb21e6d55e1bd750d05a8979e1a83d9a973"></a>

## kind property — webhook.http_config.client_cert_obj.use_tls_obj / 825ec15df4b2 / 4

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

<a id="canonical-44561aaf32ead2357997bfff057b4bce403123ca1547b80a6fb26c4832d08bf1"></a>

<a id="canonical-91b7101296c9d8bdfdca220ba40ecd44b66318f86d9834371936c5468e06988e"></a>

## name property — webhook.http_config.client_cert_obj.use_tls_obj / 825ec15df4b2 / 5

Type: `"string"`. Computed.

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

<a id="canonical-9279cc5495ab414539175f3ef3ef0aa2d1a6442351b299ea60215710f2b081ef"></a>

<a id="canonical-ad1fdd36923068b9ba309184d982d271a0834d34c9b804920268545b7eedbf96"></a>

## namespace property — webhook.http_config.client_cert_obj.use_tls_obj / 825ec15df4b2 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-7e5c1b31f67b3458c45f82aa5e1a1360e155447c1527b32af39ea3fbe1a45e3a"></a>

<a id="canonical-bbc5dce4822c4196c35eb3b276f1bfc0464681e5eaa685c89192d1fdd64427f0"></a>

## tenant property — webhook.http_config.client_cert_obj.use_tls_obj / 825ec15df4b2 / 7

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

<a id="canonical-c974accefef7b7e051c8e7615a0dab20df93e1d479350cc95f36fb3516d97855"></a>

<a id="canonical-15c0dce2172777a7e01099dc05a3eb9c6f7770efd924468e33d78c6c75f60098"></a>

## uid property — webhook.http_config.client_cert_obj.use_tls_obj / 825ec15df4b2 / 8

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

<a id="canonical-454c66fc30639b8cf7e9ffa0be1f16e6cc3d14ee821b1372f4dc44d9fafb5cef"></a>

## Next pages — webhook.http_config.client_cert_obj.use_tls_obj / 825ec15df4b2 / 9

- [webhook.http_config.client_cert_obj](data-sources--alert_receiver--reference--group-001.md#canonical-b5029d35648600fab9b94c3ccb71a08ddd0cba20638328b16b1ea0b8252d59ea)
- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)

<a id="canonical-b133fe7466a8b96fd9f051904e55567ee8670b63a3ee8727d6caf9d8e212f736"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-77a60f1830f9eba44d9f60999a5d8772e0eb11be707894bb10a6785f1363a47c"></a>

## webhook.http_config.no_authorization — webhook.http_config.no_authorization / 5db9b736316d / 2

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-9e77ba4988d239ee5fcf37742214d6c3d97c59c318f7ff0d06d4ec6a09a48451)
- [webhook](data-sources--alert_receiver--reference--group-001.md#canonical-30aebc9511b4f9fb141ec5e484c794819b7452f6fd388fd46ab47d5adfc81033)
- [webhook.http_config](data-sources--alert_receiver--reference--group-001.md#canonical-bd1503ebab64162c382b91de0969860a066f2e198443b22cffb2211a3a1c2806)
- webhook.http_config.no_authorization

<a id="canonical-c2a37440e74b7f64e99bf679bb40db1f4343099a47f38f07751b9fd12321e1b4"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no authorization.

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

<a id="canonical-e357608991402e9f479f40ff97717ceb46844013d84587abf83e79096c32c5a6"></a>

## Direct properties — webhook.http_config.no_authorization / 5db9b736316d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0e878a265e10a4fc22e6e36ff56602bcaa3025b9a8dfc129f2fa9119b071b66b"></a>

## Next pages — webhook.http_config.no_authorization / 5db9b736316d / 4

- [webhook.http_config](data-sources--alert_receiver--reference--group-001.md#canonical-bd1503ebab64162c382b91de0969860a066f2e198443b22cffb2211a3a1c2806)
- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)

<a id="canonical-feb6a539c8a5f94e27f46dca84bb8be937865d28434f8837643ab6bdf0219017"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8e022b48cf65057f63c6c90d79719bf4e81cf01c810d3cedabb332a0124c36f2"></a>

## webhook.http_config.no_tls — webhook.http_config.no_tls / 2b3b57cf4127 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-9e77ba4988d239ee5fcf37742214d6c3d97c59c318f7ff0d06d4ec6a09a48451)
- [webhook](data-sources--alert_receiver--reference--group-001.md#canonical-30aebc9511b4f9fb141ec5e484c794819b7452f6fd388fd46ab47d5adfc81033)
- [webhook.http_config](data-sources--alert_receiver--reference--group-001.md#canonical-bd1503ebab64162c382b91de0969860a066f2e198443b22cffb2211a3a1c2806)
- webhook.http_config.no_tls

<a id="canonical-6c5737eb3eee895e3c1ad04122ec5034c507afab6eda3aa8bddd4e820dc44361"></a>

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

<a id="canonical-7df4c262fd8a11ba5490d8b805a61cdc255cea3cd208b237837ddefd97493b06"></a>

## Direct properties — webhook.http_config.no_tls / 2b3b57cf4127 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-371799ea084a36e7b85e7f918faf4bb38d06b638711ec5439c2bac80885ced1d"></a>

## Next pages — webhook.http_config.no_tls / 2b3b57cf4127 / 4

- [webhook.http_config](data-sources--alert_receiver--reference--group-001.md#canonical-bd1503ebab64162c382b91de0969860a066f2e198443b22cffb2211a3a1c2806)
- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)

<a id="canonical-f2ce5b5eff735fbfa6d1c5882e4d9c1701d3eff3daacee865fcb3b6c568cbade"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aa1aceb82863f1580e5499f719bfc82a017cd5321becd999d19ed2816aa61122"></a>

## webhook.http_config.use_tls — webhook.http_config.use_tls / b3ee30f4d98b / 2

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-9e77ba4988d239ee5fcf37742214d6c3d97c59c318f7ff0d06d4ec6a09a48451)
- [webhook](data-sources--alert_receiver--reference--group-001.md#canonical-30aebc9511b4f9fb141ec5e484c794819b7452f6fd388fd46ab47d5adfc81033)
- [webhook.http_config](data-sources--alert_receiver--reference--group-001.md#canonical-bd1503ebab64162c382b91de0969860a066f2e198443b22cffb2211a3a1c2806)
- webhook.http_config.use_tls

<a id="canonical-729b072d40838e5c7d36a715c90775c8a1d840ac6e27957411545179e95ef2b2"></a>

Type: `"single"`. Computed.

Configures the token request's TLS settings.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-server_validation_choice": "[\"use_server_verification\",\"volterra_trusted_ca\"]",
  "x-ves-oneof-field-sni_choice": "[\"disable_sni\",\"sni\"]"
}
```

<a id="canonical-661520351c0627458abb8645eef21bcec8081754cd2f5dab56b1d37530895222"></a>

## Direct properties — webhook.http_config.use_tls / b3ee30f4d98b / 3

- [disable_sni](data-sources--alert_receiver--reference--group-001.md#canonical-a6102c6fe6faa3b4ac62a56fac72f778fcd8cc69c4c6bdf1a5324631e686ae41): complete subsection reference.

<a id="canonical-5d4e11a7bfda6940c73e70af127d516c22a02398afb0e5154dc9b652f42ee544"></a>

<a id="canonical-27eb308f156f4b9e1b60a783d942e84f109bea31a774d18cb6f5b7ab43d74255"></a>

## max_version property — webhook.http_config.use_tls / b3ee30f4d98b / 4

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

<a id="canonical-291f3c42ffd1a27eca659999d7aa27f6842fa76974bae0affa2f926ab5b9de1b"></a>

<a id="canonical-37cebe562cce89797c485bcca30aa9e59aaad4f12fa98edb35067ba9c17f8716"></a>

## min_version property — webhook.http_config.use_tls / b3ee30f4d98b / 5

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

<a id="canonical-8712d781e530136e653a6c40cb964f03dd6a1e345534795b93162851fb4ff1ce"></a>

<a id="canonical-027245e748b14f909b62a043028c0dbbe9e3ed26d239b91f05bbb4e99b108056"></a>

## sni property — webhook.http_config.use_tls / b3ee30f4d98b / 6

Type: `"string"`. Computed.

Exclusive with \[disable\_sni\] SNI value to be used.

Upstream description:

Exclusive with \[disable\_sni\] SNI value to be used.

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

- [use_server_verification](data-sources--alert_receiver--reference--group-001.md#canonical-01364fb7f41c532768bab5360c0905a59e2f18e98c9e0b6eae17510edd353a97): complete subsection reference.

- [volterra_trusted_ca](data-sources--alert_receiver--reference--group-001.md#canonical-a22ab1ae5c387a2196279b13e181fcf157fb322d427be59a0a91df255915ec30): complete subsection reference.

<a id="canonical-e28844a345d32d1ddd72de78c47d7797205098eaea81b40b0329ea1f836de6ae"></a>

## Next pages — webhook.http_config.use_tls / b3ee30f4d98b / 7

- [webhook.http_config.use_tls.disable_sni](data-sources--alert_receiver--reference--group-001.md#canonical-a6102c6fe6faa3b4ac62a56fac72f778fcd8cc69c4c6bdf1a5324631e686ae41)
- [webhook.http_config.use_tls.use_server_verification](data-sources--alert_receiver--reference--group-001.md#canonical-01364fb7f41c532768bab5360c0905a59e2f18e98c9e0b6eae17510edd353a97)
- [webhook.http_config.use_tls.volterra_trusted_ca](data-sources--alert_receiver--reference--group-001.md#canonical-a22ab1ae5c387a2196279b13e181fcf157fb322d427be59a0a91df255915ec30)
- [webhook.http_config](data-sources--alert_receiver--reference--group-001.md#canonical-bd1503ebab64162c382b91de0969860a066f2e198443b22cffb2211a3a1c2806)
- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)

<a id="canonical-a6102c6fe6faa3b4ac62a56fac72f778fcd8cc69c4c6bdf1a5324631e686ae41"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1756447e9c0757ccc5cb653822639802be0032fd48abb18e0838cc944f3223b0"></a>

## webhook.http_config.use_tls.disable_sni — webhook.http_config.use_tls.disable_sni / 0f46ade0d276 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-9e77ba4988d239ee5fcf37742214d6c3d97c59c318f7ff0d06d4ec6a09a48451)
- [webhook](data-sources--alert_receiver--reference--group-001.md#canonical-30aebc9511b4f9fb141ec5e484c794819b7452f6fd388fd46ab47d5adfc81033)
- [webhook.http_config](data-sources--alert_receiver--reference--group-001.md#canonical-bd1503ebab64162c382b91de0969860a066f2e198443b22cffb2211a3a1c2806)
- [webhook.http_config.use_tls](data-sources--alert_receiver--reference--group-001.md#canonical-f2ce5b5eff735fbfa6d1c5882e4d9c1701d3eff3daacee865fcb3b6c568cbade)
- webhook.http_config.use_tls.disable_sni

<a id="canonical-99d4f6113f19a90708113ca5f7bda3443f161f72da383fbd92a1b26c464d1fd0"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-f8691c0b234c11fd53f15c33cc0c28fdc38eed190d8d34c67d3fc2d0f192c3ab"></a>

## Direct properties — webhook.http_config.use_tls.disable_sni / 0f46ade0d276 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bb51a063974be0fe6a6b463d1e9c94721cd9766acef3f00f0d9b5634307a0fba"></a>

## Next pages — webhook.http_config.use_tls.disable_sni / 0f46ade0d276 / 4

- [webhook.http_config.use_tls](data-sources--alert_receiver--reference--group-001.md#canonical-f2ce5b5eff735fbfa6d1c5882e4d9c1701d3eff3daacee865fcb3b6c568cbade)
- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)

<a id="canonical-01364fb7f41c532768bab5360c0905a59e2f18e98c9e0b6eae17510edd353a97"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2d25936c6a326c96f7350b8d5d90715c36a9e4f87c4b76792ee1e637ae927456"></a>

## webhook.http_config.use_tls.use_server_verification — webhook.http_config.use_tls.use_server_verification / 65b5f82d85c2 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-9e77ba4988d239ee5fcf37742214d6c3d97c59c318f7ff0d06d4ec6a09a48451)
- [webhook](data-sources--alert_receiver--reference--group-001.md#canonical-30aebc9511b4f9fb141ec5e484c794819b7452f6fd388fd46ab47d5adfc81033)
- [webhook.http_config](data-sources--alert_receiver--reference--group-001.md#canonical-bd1503ebab64162c382b91de0969860a066f2e198443b22cffb2211a3a1c2806)
- [webhook.http_config.use_tls](data-sources--alert_receiver--reference--group-001.md#canonical-f2ce5b5eff735fbfa6d1c5882e4d9c1701d3eff3daacee865fcb3b6c568cbade)
- webhook.http_config.use_tls.use_server_verification

<a id="canonical-465b85413b81ede1d04f5c1c334479e46f7253d9fc1704e687bfec0dd6eb9751"></a>

Type: `"single"`. Computed.

Configuration parameter for use server verification.

Upstream description:

Upstream TLS Validation Context.

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

<a id="canonical-9095a58564a620068f0a111f2e4ad21d331be337dc4fa3589d088f4adc2a609c"></a>

## Direct properties — webhook.http_config.use_tls.use_server_verification / 65b5f82d85c2 / 3

- [ca_cert_obj](data-sources--alert_receiver--reference--group-001.md#canonical-e704ed6b4e7c1d487836a5cfe758d1d6246b87f5f96cf2f46cd30eff0f10f4ff): complete subsection reference.

<a id="canonical-3c6cf87894e7481ac254468a99187eef6648a113da70d0b1a9323129969babcc"></a>

## Next pages — webhook.http_config.use_tls.use_server_verification / 65b5f82d85c2 / 4

- [webhook.http_config.use_tls.use_server_verification.ca_cert_obj](data-sources--alert_receiver--reference--group-001.md#canonical-e704ed6b4e7c1d487836a5cfe758d1d6246b87f5f96cf2f46cd30eff0f10f4ff)
- [webhook.http_config.use_tls](data-sources--alert_receiver--reference--group-001.md#canonical-f2ce5b5eff735fbfa6d1c5882e4d9c1701d3eff3daacee865fcb3b6c568cbade)
- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)

<a id="canonical-e704ed6b4e7c1d487836a5cfe758d1d6246b87f5f96cf2f46cd30eff0f10f4ff"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c4cbeca0f04087a6636950ce7f3b0398f2ae4e7cfc44f7086cbe30f471cdf7a4"></a>

## webhook.http_config.use_tls.use_server_verification.ca_cert_obj — webhook.http_config.use_tls.use_server_verification.ca_cert_obj / 3fd574c8049e / 2

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-9e77ba4988d239ee5fcf37742214d6c3d97c59c318f7ff0d06d4ec6a09a48451)
- [webhook](data-sources--alert_receiver--reference--group-001.md#canonical-30aebc9511b4f9fb141ec5e484c794819b7452f6fd388fd46ab47d5adfc81033)
- [webhook.http_config](data-sources--alert_receiver--reference--group-001.md#canonical-bd1503ebab64162c382b91de0969860a066f2e198443b22cffb2211a3a1c2806)
- [webhook.http_config.use_tls](data-sources--alert_receiver--reference--group-001.md#canonical-f2ce5b5eff735fbfa6d1c5882e4d9c1701d3eff3daacee865fcb3b6c568cbade)
- [webhook.http_config.use_tls.use_server_verification](data-sources--alert_receiver--reference--group-001.md#canonical-01364fb7f41c532768bab5360c0905a59e2f18e98c9e0b6eae17510edd353a97)
- webhook.http_config.use_tls.use_server_verification.ca_cert_obj

<a id="canonical-ee8b09f5f34eb3d1b46bca10e727540fdb523a555df895b13a5ce6d4c3778702"></a>

Type: `"single"`. Computed.

Configuration parameter for ca cert obj.

Upstream description:

Configuration for CA certificate.

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

<a id="canonical-834ed2fa1469a11802c6eaee674b14db00239c8f167853045c0de46927dca6d0"></a>

## Direct properties — webhook.http_config.use_tls.use_server_verification.ca_cert_obj / 3fd574c8049e / 3

- [trusted_ca](data-sources--alert_receiver--reference--group-001.md#canonical-7817a113fa1bab6ffa15eb7c048a99f138c70885c53cfd6ac0881ecbe256fc9e): complete subsection reference.

<a id="canonical-e1c5bede253da9940e79e20a8cc38cf62264885f6cf28ce3cd66c0cfed25043c"></a>

## Next pages — webhook.http_config.use_tls.use_server_verification.ca_cert_obj / 3fd574c8049e / 4

- [webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca](data-sources--alert_receiver--reference--group-001.md#canonical-7817a113fa1bab6ffa15eb7c048a99f138c70885c53cfd6ac0881ecbe256fc9e)
- [webhook.http_config.use_tls.use_server_verification](data-sources--alert_receiver--reference--group-001.md#canonical-01364fb7f41c532768bab5360c0905a59e2f18e98c9e0b6eae17510edd353a97)
- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)

<a id="canonical-7817a113fa1bab6ffa15eb7c048a99f138c70885c53cfd6ac0881ecbe256fc9e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-14541ac8da4779cabf841c0e01797180da5c41a672e1775a4ae3b6eeb050489a"></a>

## webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca — webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca / 03a12075e664 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-9e77ba4988d239ee5fcf37742214d6c3d97c59c318f7ff0d06d4ec6a09a48451)
- [webhook](data-sources--alert_receiver--reference--group-001.md#canonical-30aebc9511b4f9fb141ec5e484c794819b7452f6fd388fd46ab47d5adfc81033)
- [webhook.http_config](data-sources--alert_receiver--reference--group-001.md#canonical-bd1503ebab64162c382b91de0969860a066f2e198443b22cffb2211a3a1c2806)
- [webhook.http_config.use_tls](data-sources--alert_receiver--reference--group-001.md#canonical-f2ce5b5eff735fbfa6d1c5882e4d9c1701d3eff3daacee865fcb3b6c568cbade)
- [webhook.http_config.use_tls.use_server_verification](data-sources--alert_receiver--reference--group-001.md#canonical-01364fb7f41c532768bab5360c0905a59e2f18e98c9e0b6eae17510edd353a97)
- [webhook.http_config.use_tls.use_server_verification.ca_cert_obj](data-sources--alert_receiver--reference--group-001.md#canonical-e704ed6b4e7c1d487836a5cfe758d1d6246b87f5f96cf2f46cd30eff0f10f4ff)
- webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca

<a id="canonical-d2d169ee1b7eda85830e769b1547a8f1fbb4cf9142fa5b89906bc28d7dc76b91"></a>

Type: `"list"`. Computed.

Certificate Object. Reference to client certificate object.

Upstream description:

Reference to client certificate object.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-484c60d4e59c64629799f5f129c3f8e2c262fc3e1f423d82da96f45021f8b7b2"></a>

## Direct properties — webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca / 03a12075e664 / 3

<a id="canonical-481b1e61b4e920db1c6b1995e18de57556b619b4c71abfbef7ee6c19103614a6"></a>

<a id="canonical-ffdd6aeb2fd6845ff3b732fad0fc22e53a8c57270d0e883d4f6a0520d1caf3fc"></a>

## kind property — webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca / 03a12075e664 / 4

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

<a id="canonical-83a983ce0d26742c590d1d581e6ebc04f2153d66b0989832d0e327b907f38548"></a>

<a id="canonical-796c69ba6c9a4f2900cb5d7ed0723b3a80af88ae44527bee808050e064ae6477"></a>

## name property — webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca / 03a12075e664 / 5

Type: `"string"`. Computed.

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

<a id="canonical-53cedae8daad4333f6eed81ab4c740c1506280ec821222d9a7c0126da95cc298"></a>

<a id="canonical-178d29c2fea0ed431138b7ab6ff542640248889ab328712bd3a2e5b623baa21a"></a>

## namespace property — webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca / 03a12075e664 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-4db14a10eb828f29e2b76bf34f1342a228d33868befb9f7d8615d62dad25e486"></a>

<a id="canonical-20b8b5d927538954297c78099ddb5047659ce455741837c47ee33e2ca3f32994"></a>

## tenant property — webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca / 03a12075e664 / 7

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

<a id="canonical-ed04081d938f662fbb480581c1bbb6bc23d5c534a84218013c7f0b7aed9d8bcb"></a>

<a id="canonical-9d942f34859a473ee2af1ab7b8398c410b12b0a318236f88c0dd52bf61435b47"></a>

## uid property — webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca / 03a12075e664 / 8

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

<a id="canonical-e70aefaf6b4f7eeb16518a12c69cebbd17e91372fe72d7e13db932204c860ee6"></a>

## Next pages — webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca / 03a12075e664 / 9

- [webhook.http_config.use_tls.use_server_verification.ca_cert_obj](data-sources--alert_receiver--reference--group-001.md#canonical-e704ed6b4e7c1d487836a5cfe758d1d6246b87f5f96cf2f46cd30eff0f10f4ff)
- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)

<a id="canonical-a22ab1ae5c387a2196279b13e181fcf157fb322d427be59a0a91df255915ec30"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-34851b040225639f815fc35a7cc850cf8bbd46da0828482f2b4d0efc1a3b6464"></a>

## webhook.http_config.use_tls.volterra_trusted_ca — webhook.http_config.use_tls.volterra_trusted_ca / 3583ce45a6ca / 2

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-9e77ba4988d239ee5fcf37742214d6c3d97c59c318f7ff0d06d4ec6a09a48451)
- [webhook](data-sources--alert_receiver--reference--group-001.md#canonical-30aebc9511b4f9fb141ec5e484c794819b7452f6fd388fd46ab47d5adfc81033)
- [webhook.http_config](data-sources--alert_receiver--reference--group-001.md#canonical-bd1503ebab64162c382b91de0969860a066f2e198443b22cffb2211a3a1c2806)
- [webhook.http_config.use_tls](data-sources--alert_receiver--reference--group-001.md#canonical-f2ce5b5eff735fbfa6d1c5882e4d9c1701d3eff3daacee865fcb3b6c568cbade)
- webhook.http_config.use_tls.volterra_trusted_ca

<a id="canonical-73e01aaeff96fbd9a2cc7e800f0d42417758d1b69b647eebf135a1301d7c703d"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for volterra trusted ca.

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

<a id="canonical-8ec6ad5be7a8557f86c10e6cdf70a59e80bdda6a64afde0e5cef758e5c3f2064"></a>

## Direct properties — webhook.http_config.use_tls.volterra_trusted_ca / 3583ce45a6ca / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d2eb468b8fdb73e7c31d6500409208187db74ef6eada3f7147fee034cdd279ff"></a>

## Next pages — webhook.http_config.use_tls.volterra_trusted_ca / 3583ce45a6ca / 4

- [webhook.http_config.use_tls](data-sources--alert_receiver--reference--group-001.md#canonical-f2ce5b5eff735fbfa6d1c5882e4d9c1701d3eff3daacee865fcb3b6c568cbade)
- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)

<a id="canonical-f89f467192ddf8816aaf8302245110958f22034f09366862618892b0fe1cee24"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-306038369c2c4d6c6352fb71575959acdfb25236cd667f30cefb15f437bae7aa"></a>

## webhook.url — webhook.url / 1becd9ab98c6 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-9e77ba4988d239ee5fcf37742214d6c3d97c59c318f7ff0d06d4ec6a09a48451)
- [webhook](data-sources--alert_receiver--reference--group-001.md#canonical-30aebc9511b4f9fb141ec5e484c794819b7452f6fd388fd46ab47d5adfc81033)
- webhook.url

<a id="canonical-7cfa4beb05222e82c8dd8e38424802e6875c82bb910caf7f4b79bf1980188a08"></a>

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

<a id="canonical-51fd93f39b83b035fc86bb2b477a51336453cdd31e2215bb7e603840b50c0bfb"></a>

## Direct properties — webhook.url / 1becd9ab98c6 / 3

- [blindfold_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-55e8016ac00d3b716517a22379e6bdb1a3009cef0d66830197dc93f095476423): complete subsection reference.

- [clear_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-083b5ce744a8d686b70937afb2fd468f0f09699170f7b3274811bec3504509b5): complete subsection reference.

<a id="canonical-85c11092f891074fe4148bcf860f50b9cc1948225acd3941047b912dfc2839d9"></a>

## Next pages — webhook.url / 1becd9ab98c6 / 4

- [webhook.url.blindfold_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-55e8016ac00d3b716517a22379e6bdb1a3009cef0d66830197dc93f095476423)
- [webhook.url.clear_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-083b5ce744a8d686b70937afb2fd468f0f09699170f7b3274811bec3504509b5)
- [webhook](data-sources--alert_receiver--reference--group-001.md#canonical-30aebc9511b4f9fb141ec5e484c794819b7452f6fd388fd46ab47d5adfc81033)
- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)

<a id="canonical-55e8016ac00d3b716517a22379e6bdb1a3009cef0d66830197dc93f095476423"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-51d9f7abd4e512dc4ea867693619059bf230127a2ef61dc101e1f69f85044d22"></a>

## webhook.url.blindfold_secret_info — webhook.url.blindfold_secret_info / 505cd7a00677 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-9e77ba4988d239ee5fcf37742214d6c3d97c59c318f7ff0d06d4ec6a09a48451)
- [webhook](data-sources--alert_receiver--reference--group-001.md#canonical-30aebc9511b4f9fb141ec5e484c794819b7452f6fd388fd46ab47d5adfc81033)
- [webhook.url](data-sources--alert_receiver--reference--group-001.md#canonical-f89f467192ddf8816aaf8302245110958f22034f09366862618892b0fe1cee24)
- webhook.url.blindfold_secret_info

<a id="canonical-b47f27c116074fae0eba790c71838ae6f7b1966e39264ce6dcbb385fc0b3bbaa"></a>

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

<a id="canonical-0903fe79360f1e1105914d20008251238dbe92c93edef03c8846f401f98a9108"></a>

## Direct properties — webhook.url.blindfold_secret_info / 505cd7a00677 / 3

<a id="canonical-60f0e18396c8f71d41c02e73672eb87768dd510e209bfe735b11f76911c55625"></a>

<a id="canonical-9a12fbdd19160f0ee7d7c31a94e25cf32a0f1758a788b8c7d03a73a8e3e1840c"></a>

## decryption_provider property — webhook.url.blindfold_secret_info / 505cd7a00677 / 4

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

<a id="canonical-c78ec6b5d775bad47eee5b686573f9fc56c14eb9433f89f13e69211070bdb877"></a>

<a id="canonical-b0bd37ecdae0fcf026010676c12b451833853ef56d1144bc48a7126176fd3aff"></a>

## location property — webhook.url.blindfold_secret_info / 505cd7a00677 / 5

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

<a id="canonical-573141ea5d363cf762557f840e75f6c085559cd467b2de96f350823fe60ba73d"></a>

<a id="canonical-a1a5f7df352a07d1e6f4065e43a3c6442b11ec626972b6bc46f3e7b51e290ab6"></a>

## store_provider property — webhook.url.blindfold_secret_info / 505cd7a00677 / 6

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

<a id="canonical-78813e8861c9da03d9ec9a2ecaf255e7f7ad93bb7e90d278ba7e061308be9f74"></a>

## Next pages — webhook.url.blindfold_secret_info / 505cd7a00677 / 7

- [webhook.url](data-sources--alert_receiver--reference--group-001.md#canonical-f89f467192ddf8816aaf8302245110958f22034f09366862618892b0fe1cee24)
- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)

<a id="canonical-083b5ce744a8d686b70937afb2fd468f0f09699170f7b3274811bec3504509b5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eec42b467239b8e04e9cc18468a5097e846685d9a2abf4ae657b5020f8fe426e"></a>

## webhook.url.clear_secret_info — webhook.url.clear_secret_info / 514aff14bdbf / 2

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-9e77ba4988d239ee5fcf37742214d6c3d97c59c318f7ff0d06d4ec6a09a48451)
- [webhook](data-sources--alert_receiver--reference--group-001.md#canonical-30aebc9511b4f9fb141ec5e484c794819b7452f6fd388fd46ab47d5adfc81033)
- [webhook.url](data-sources--alert_receiver--reference--group-001.md#canonical-f89f467192ddf8816aaf8302245110958f22034f09366862618892b0fe1cee24)
- webhook.url.clear_secret_info

<a id="canonical-cf0e9fca418fd799d37d433219e8218a611b0e807e7bd999e0ae09ff3ff7c13e"></a>

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

<a id="canonical-f59e3acd8f8317a84969d46bdd1b09837c062401d71f6042e97c57b7133800a9"></a>

## Direct properties — webhook.url.clear_secret_info / 514aff14bdbf / 3

<a id="canonical-6ea16f29827b52624536b1049b77da1d047f5ffd35c83342b759feb4af5e8ce2"></a>

<a id="canonical-2beb1d67f46d355841ad9362e10e6f8ae967f8a34087c73f527db892367e96cb"></a>

## provider_ref property — webhook.url.clear_secret_info / 514aff14bdbf / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-bae1fcd386bd6e8577321aa31f78abc1d6db35e334e9f94ef89c91857bcc925e"></a>

<a id="canonical-8d398dcf2498ddc2443c188b48091e0cc815d044134f9127713d0cb1d75a4a53"></a>

## url property — webhook.url.clear_secret_info / 514aff14bdbf / 5

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

<a id="canonical-32920bc9425721677437ce1aae60eda02b3728a6bf773791e813d249eec88713"></a>

## Next pages — webhook.url.clear_secret_info / 514aff14bdbf / 6

- [webhook.url](data-sources--alert_receiver--reference--group-001.md#canonical-f89f467192ddf8816aaf8302245110958f22034f09366862618892b0fe1cee24)
- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)
