---
page_title: "xcsh_discovery reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_discovery reference."
---

# xcsh_discovery reference

<a id="canonical-89b3e09314ef5fec837c9034f68d6bb66048b1e8e127c21bebfbd751a57f9668"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0537a5bda60b5126aea7c93ee6e67aace0da051daa99a7ae5e7d6fda35aad03d"></a>

## Property reference — Property reference / 94c550cf146d / 2

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)
- Property reference

<a id="canonical-80b7440a98629026f47ad68df5a776fda45cfcb6a65c6b06a3911425b84549ea"></a>

## Direct properties — Property reference / 94c550cf146d / 3

<a id="canonical-07ecf9fdc0f538fb61668c90ad15c7acc92df262ecfe0d136d4620fb007fd728"></a>

<a id="canonical-877dc113a702597256b50c710ca44b1487d3d5208414c0dd381f68f437264491"></a>

## annotations property — Property reference / 94c550cf146d / 4

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

<a id="canonical-6349768a695e1d4b13d5b9e417c94b17aa577605c839836bbabc8e25a4d9817f"></a>

<a id="canonical-11f2eedf80614ac698ada959ef63b09fdb4ac8857089cce30453cc022ec475bf"></a>

## cluster_id property — Property reference / 94c550cf146d / 5

Type: `"string"`. Computed.

\[OneOf: cluster\_id, no\_cluster\_id; Default: no\_cluster\_id\] Exclusive with \[no\_cluster\_id\]
Specify identifier for discovery cluster. This identifier can be specified in endpoint object to
discover only from this discovery object.

Upstream description:

Exclusive with \[no\_cluster\_id\] Specify identifier for discovery cluster. This identifier can be
specified in endpoint object to discover only from this discovery object.

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

OneOf alternatives in this subsection:

- [cluster_id](data-sources--discovery--reference--group-001.md#canonical-6349768a695e1d4b13d5b9e417c94b17aa577605c839836bbabc8e25a4d9817f)
- [no_cluster_id](data-sources--discovery--reference--group-001.md#canonical-bc76a7116b66e700f48d50b0d53306bb8e2a03efd8228c74ee271680471c69b2)

Select alternatives according to the provider validators above.

<a id="canonical-1dc72ad8265d0d6f8cd06384189c63c00ca6437f687cf6cfb934d70664e13b55"></a>

<a id="canonical-6bafe8761d17bb1cb3fff30eeabad406450d0935b1247e26fd588fa54ca6b9a6"></a>

## description property — Property reference / 94c550cf146d / 6

Type: `"string"`. Computed.

Description of the Discovery.

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

- [discovery_consul](data-sources--discovery--reference--group-001.md#canonical-4393fc86fa2ced1c5d3db9a1bb2d4bd5460a72aa837331742b19239501c44294): complete subsection reference.

- [discovery_k8s](data-sources--discovery--reference--group-001.md#canonical-3c1f55627ad4d68af88df52467484fc60a52655cab4af4210a019b3b79d1b17a): complete subsection reference.

<a id="canonical-3fe6cbe70d98040df256d9d44a6f5c6fea97c464e6ff7035f53749c78e9a8d32"></a>

<a id="canonical-565dd0557b76edd89d62f347e3595d5d4ef43a20f595dc85d9f01f8e4035062b"></a>

## id property — Property reference / 94c550cf146d / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-ebaa16e2be580bc8ef8bca65d1322a99cc2e3f4c4605ccce680ac1d31c3083da"></a>

<a id="canonical-235c01fd8fe8fad173229b758f24ac818cbe3f58916ffb9f7e124d12a8d9de94"></a>

## labels property — Property reference / 94c550cf146d / 8

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

<a id="canonical-efc4a2874030edc499c5eb4f7e559983b099da47ace2b724e12f707bcaee770e"></a>

<a id="canonical-c414687c5b0805d457c7348779297bcef63ecad9143cd8f5b3b1432609e3cc3e"></a>

## name property — Property reference / 94c550cf146d / 9

Type: `"string"`. Required.

Name of the Discovery.

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

<a id="canonical-32b669e6f69aebabb5bbc4af7e9438f92b1613da68289affce502c803211c32d"></a>

<a id="canonical-cc33a09f81e510d6ba3b542cc16a839d33846b23d05a75db78c63709adb87cd8"></a>

## namespace property — Property reference / 94c550cf146d / 10

Type: `"string"`. Required.

Namespace where the Discovery exists.

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

- [no_cluster_id](data-sources--discovery--reference--group-001.md#canonical-cab961f6582c0f7f0205c82d8b10a272b9e0fe4b779f443dc0de9bb85e6b059c): complete subsection reference.

- [where](data-sources--discovery--reference--group-001.md#canonical-86e8aa7a75f491571c4a99862108276e04b016f3b24020940f1e1b752bd38094): complete subsection reference.

<a id="canonical-bb4f405ad40ab7c7297b9fae5e15a848b153c2b1ad1c10942d91582a836d25ba"></a>

## All schema paths — Property reference / 94c550cf146d / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--discovery--reference--group-001.md#canonical-07ecf9fdc0f538fb61668c90ad15c7acc92df262ecfe0d136d4620fb007fd728) |
| `cluster_id` | [cluster_id](data-sources--discovery--reference--group-001.md#canonical-6349768a695e1d4b13d5b9e417c94b17aa577605c839836bbabc8e25a4d9817f) |
| `description` | [description](data-sources--discovery--reference--group-001.md#canonical-1dc72ad8265d0d6f8cd06384189c63c00ca6437f687cf6cfb934d70664e13b55) |
| `discovery_consul` | [discovery_consul](data-sources--discovery--reference--group-001.md#canonical-c1d9546703b5f7b256c015fc7cab2f48155c68c079e1020ab239ae4d0795cb99) |
| `discovery_consul.access_info` | [discovery_consul.access_info](data-sources--discovery--reference--group-001.md#canonical-ba12dd1beadf83ba3cfbf9deb29f01fab9ec1cefba82b10ec50062cc8eca490d) |
| `discovery_consul.access_info.connection_info` | [discovery_consul.access_info.connection_info](data-sources--discovery--reference--group-001.md#canonical-5456f17afb0d25aa7750df9abec8dc58d2ec8e1f26435c5601415828fb58eedd) |
| `discovery_consul.access_info.connection_info.api_server` | [discovery_consul.access_info.connection_info.api_server](data-sources--discovery--reference--group-001.md#canonical-69f0c3e02ff0e08a0e266cc6594132749cd204c70899b55393bdf84bfe7ad480) |
| `discovery_consul.access_info.connection_info.tls_info` | [discovery_consul.access_info.connection_info.tls_info](data-sources--discovery--reference--group-001.md#canonical-50ed6c6fb50989c508aa76b2fda47ae513a22612927c4b32f846289c285fc3dc) |
| `discovery_consul.access_info.connection_info.tls_info.certificate` | [discovery_consul.access_info.connection_info.tls_info.certificate](data-sources--discovery--reference--group-001.md#canonical-ee897d9c4f480ed2b42ed3cee9e7c6e485b7088bfda90101007121d7e87991ef) |
| `discovery_consul.access_info.connection_info.tls_info.key_url` | [discovery_consul.access_info.connection_info.tls_info.key_url](data-sources--discovery--reference--group-001.md#canonical-c5e65ce87690c1898dd75d217b0e95f97ae2ad5066ac8cee9032622b1b452a7e) |
| `discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info` | [discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info](data-sources--discovery--reference--group-001.md#canonical-a910145e2051d4753e799141d70d85fb11d12e66d2b28051da91bdda31d56a3e) |
| `discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info.decryption_provider` | [discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info.decryption_provider](data-sources--discovery--reference--group-001.md#canonical-fcf0a7c14a5886ee85338f117c2aafb1f294e81ca0239b4d78eb5d37d18171c4) |
| `discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info.location` | [discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info.location](data-sources--discovery--reference--group-001.md#canonical-070fca5a417d810d3aefb317f7200c746d8c8ffc74262a330ed3dbeaf8542681) |
| `discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info.store_provider` | [discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info.store_provider](data-sources--discovery--reference--group-001.md#canonical-84e7296e0061ef95aa0f5b2261d6e00c59facfef3b48f14bcf367db00f7aabb9) |
| `discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info` | [discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info](data-sources--discovery--reference--group-001.md#canonical-21684ee32f794c44a302d9214c4974feef83bef0c6530b50f0dbce70f98c9aef) |
| `discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info.provider_ref` | [discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info.provider_ref](data-sources--discovery--reference--group-001.md#canonical-436fcdc85009be9d6a1ea0e111dfc3258e5b46fd58c3ca753bb432f359959f6d) |
| `discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info.url` | [discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info.url](data-sources--discovery--reference--group-001.md#canonical-63f6721cc80db35076930a3db1beaa33c4a65f442a1058f022a23362485c3a6d) |
| `discovery_consul.access_info.connection_info.tls_info.server_name` | [discovery_consul.access_info.connection_info.tls_info.server_name](data-sources--discovery--reference--group-001.md#canonical-c01a2b8b64fd17c88c5c26f51fd5a252f5e3ae1aaacbf7b8b59043f09bbedf64) |
| `discovery_consul.access_info.connection_info.tls_info.trusted_ca_url` | [discovery_consul.access_info.connection_info.tls_info.trusted_ca_url](data-sources--discovery--reference--group-001.md#canonical-37ed59cd45678f7d911cf5a5a09909a5072f7cc8b64fb9145d9c8a4ca2e40b35) |
| `discovery_consul.access_info.http_basic_auth_info` | [discovery_consul.access_info.http_basic_auth_info](data-sources--discovery--reference--group-001.md#canonical-daf091cf7f13cf207dc3d0cbd1196a3437ff74d803a33ba2f967941cdadcd5ea) |
| `discovery_consul.access_info.http_basic_auth_info.passwd_url` | [discovery_consul.access_info.http_basic_auth_info.passwd_url](data-sources--discovery--reference--group-001.md#canonical-45d482293334f47f327af159ed19cc596d1e6181baec335c6965c187d3c9015d) |
| `discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info` | [discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info](data-sources--discovery--reference--group-001.md#canonical-870eb4df657ac1ec91a5236f243018cf184ca926df73ad0b0f2e2c2623d82ca0) |
| `discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info.decryption_provider` | [discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info.decryption_provider](data-sources--discovery--reference--group-001.md#canonical-cac5539856c3c967ee43a3cc3d74195053451e74acea44a1ddf12614e89c9d09) |
| `discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info.location` | [discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info.location](data-sources--discovery--reference--group-001.md#canonical-bb3e16006197b4bbf929bd9e3c8e85a2ce4646cfcca13413ca1f8407ea9e9e2d) |
| `discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info.store_provider` | [discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info.store_provider](data-sources--discovery--reference--group-001.md#canonical-0206eb23a3184298efc9a15e6326bbc336fbadb9700a334a05bc79a13bc39646) |
| `discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info` | [discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info](data-sources--discovery--reference--group-001.md#canonical-85f0e32f63c9e9b138bb006fdbd5be4d6d8d8fad916873a4cc32684d1089e9d8) |
| `discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info.provider_ref` | [discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info.provider_ref](data-sources--discovery--reference--group-001.md#canonical-aa749059cd51d4e22462277e2f77a54ed06da1deee699c5ec1148cdf1432c83d) |
| `discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info.url` | [discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info.url](data-sources--discovery--reference--group-001.md#canonical-9162e59b4621777503a8f4524b6a65620364aec46b6613259b975d2dda8d7ce5) |
| `discovery_consul.access_info.http_basic_auth_info.user_name` | [discovery_consul.access_info.http_basic_auth_info.user_name](data-sources--discovery--reference--group-001.md#canonical-620c15e78dfc41e80c4d7650ffb488d9ab8101079b7aa74cec75f65737490d95) |
| `discovery_consul.publish_info` | [discovery_consul.publish_info](data-sources--discovery--reference--group-001.md#canonical-4ba15100bd73a4c1ee8ee737f199ab6e7760554cfe9ac895703c10fbe3a2b0dc) |
| `discovery_consul.publish_info.disable_spec` | [discovery_consul.publish_info.disable_spec](data-sources--discovery--reference--group-001.md#canonical-9bb4fa3c193891fb854d79b30b15e2f23f6801571ab98e53ba4b62d6ccb097ee) |
| `discovery_consul.publish_info.publish` | [discovery_consul.publish_info.publish](data-sources--discovery--reference--group-001.md#canonical-3b8199c0a5d8aeddde0dde0a4cbde830b5814061e14c676995eed51408e8ac70) |
| `discovery_k8s` | [discovery_k8s](data-sources--discovery--reference--group-001.md#canonical-4eaf3060e9826836c611ea80e9e98f58914ccec81dfe4d1950164e297efae8c5) |
| `discovery_k8s.access_info` | [discovery_k8s.access_info](data-sources--discovery--reference--group-001.md#canonical-a4055a0f637a30c771288b2567a4a6e98bba2e77b5f8828f38b2aa02a33ee068) |
| `discovery_k8s.access_info.connection_info` | [discovery_k8s.access_info.connection_info](data-sources--discovery--reference--group-001.md#canonical-b5dbff5ac61711d4898511ea2e593dff82d07c122bebbf5e957e313a36ff4713) |
| `discovery_k8s.access_info.connection_info.api_server` | [discovery_k8s.access_info.connection_info.api_server](data-sources--discovery--reference--group-001.md#canonical-3f2f36e58cd0e02fb0b87bf6266f75c4bda2877c314b202f51daef22c210d3e9) |
| `discovery_k8s.access_info.connection_info.tls_info` | [discovery_k8s.access_info.connection_info.tls_info](data-sources--discovery--reference--group-001.md#canonical-5ad99e2afab1f9fdbbc931621c0bd2a941902586794c58c6c48dbd47e1fc86c4) |
| `discovery_k8s.access_info.connection_info.tls_info.certificate` | [discovery_k8s.access_info.connection_info.tls_info.certificate](data-sources--discovery--reference--group-001.md#canonical-3532e4310911a6e3960bbd948beb13eb4fb5755b1c1e81599e5a78d5241acc0c) |
| `discovery_k8s.access_info.connection_info.tls_info.key_url` | [discovery_k8s.access_info.connection_info.tls_info.key_url](data-sources--discovery--reference--group-001.md#canonical-d9964ed0c6412650ad5f61e6b1f56afdb51ddefabcdcc4623a74886c00652aeb) |
| `discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info` | [discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info](data-sources--discovery--reference--group-001.md#canonical-5fd69de3b521d70c693a9636a5c973b6775577aeb36078851f5ba4b1ecd6b870) |
| `discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info.decryption_provider` | [discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info.decryption_provider](data-sources--discovery--reference--group-001.md#canonical-6d2ea5603638c17505e2d44858e5ae74c4815b569967e6b2364ad5b57ac2cb4d) |
| `discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info.location` | [discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info.location](data-sources--discovery--reference--group-001.md#canonical-8b23f448cc427d79ab6bfecd7b272cb273d43e36603c5280f59aca544b372241) |
| `discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info.store_provider` | [discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info.store_provider](data-sources--discovery--reference--group-001.md#canonical-eceb2d200373821c13d8aa9a04f173d48775f7a3a4cb065ea4a6cedf0dbed251) |
| `discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info` | [discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info](data-sources--discovery--reference--group-001.md#canonical-15a4d81b2a04b69b6a9084d5fb061e8c8dd2fd48bb94bf5389af23f564af8c5a) |
| `discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info.provider_ref` | [discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info.provider_ref](data-sources--discovery--reference--group-001.md#canonical-3c0adb32ac3dea0d9442668251d5702238b008a32b64cf495756e90153080296) |
| `discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info.url` | [discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info.url](data-sources--discovery--reference--group-001.md#canonical-c6ea6980ff46ef77da53a441bae1c76e61caea183744da08a320c6ecf6cb2010) |
| `discovery_k8s.access_info.connection_info.tls_info.server_name` | [discovery_k8s.access_info.connection_info.tls_info.server_name](data-sources--discovery--reference--group-001.md#canonical-2a42410f3cdc2755a94ed256fedfa0d3e735189f0738ec3a72e6c68b98637a4f) |
| `discovery_k8s.access_info.connection_info.tls_info.trusted_ca_url` | [discovery_k8s.access_info.connection_info.tls_info.trusted_ca_url](data-sources--discovery--reference--group-001.md#canonical-469357d116017e15949f62f2f2117ccb2ef0069fc782356c578573c2e05af635) |
| `discovery_k8s.access_info.isolated` | [discovery_k8s.access_info.isolated](data-sources--discovery--reference--group-001.md#canonical-e51e81e4bc801e7a7ca67a4069b1f8f551138c9a95a8f52d02716eaad1ca9555) |
| `discovery_k8s.access_info.kubeconfig_url` | [discovery_k8s.access_info.kubeconfig_url](data-sources--discovery--reference--group-001.md#canonical-64ff47ccc9b1bb26f85aa9a4b0adbc1b784582233b48fca4afc67663de2c8a7b) |
| `discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info` | [discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info](data-sources--discovery--reference--group-001.md#canonical-93e88151f78c8098e1a58c23da5b2c2ca0943ea22f316117614d433522d1b3e2) |
| `discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info.decryption_provider` | [discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info.decryption_provider](data-sources--discovery--reference--group-001.md#canonical-a3a462a66273d34009a90fbd83b4fb24e79bce861b22d45cc7a440d24f712669) |
| `discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info.location` | [discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info.location](data-sources--discovery--reference--group-001.md#canonical-046f9d713bb4b98d7289dd347ecc06874f5434816a864ed684c4f0606312e5d9) |
| `discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info.store_provider` | [discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info.store_provider](data-sources--discovery--reference--group-001.md#canonical-166c53879e9d0726758debf8f784c297a853ba5b528d8c64156758cb646d16e5) |
| `discovery_k8s.access_info.kubeconfig_url.clear_secret_info` | [discovery_k8s.access_info.kubeconfig_url.clear_secret_info](data-sources--discovery--reference--group-001.md#canonical-2dfc28e191786c3d84d4b77fd827baa59345d88fea197c8a46eead0d607e0767) |
| `discovery_k8s.access_info.kubeconfig_url.clear_secret_info.provider_ref` | [discovery_k8s.access_info.kubeconfig_url.clear_secret_info.provider_ref](data-sources--discovery--reference--group-001.md#canonical-0e06cb3bd6e48d3a4a134089fb64a049e5e50c04a4844d4cc69c37aaf0c18899) |
| `discovery_k8s.access_info.kubeconfig_url.clear_secret_info.url` | [discovery_k8s.access_info.kubeconfig_url.clear_secret_info.url](data-sources--discovery--reference--group-001.md#canonical-70b7070aef32a158bde4795a4ae8b08c1d23c3154a7b4a95b50ef05e304ce83b) |
| `discovery_k8s.access_info.reachable` | [discovery_k8s.access_info.reachable](data-sources--discovery--reference--group-001.md#canonical-dcafd875b728ebc359991f1a0f51906c9a7fe55212ef9d13047ebf30994dbd49) |
| `discovery_k8s.default_all` | [discovery_k8s.default_all](data-sources--discovery--reference--group-001.md#canonical-e4e67154b5c84e4c9ad2ae490f2dcb89c18cc205e36ab1d3af9aa35acb3dba4c) |
| `discovery_k8s.namespace_mapping` | [discovery_k8s.namespace_mapping](data-sources--discovery--reference--group-001.md#canonical-81ea92d29cb22d48ef84853ff533a4a3df57ba01002c2257d88e5e86ec22d586) |
| `discovery_k8s.namespace_mapping.items` | [discovery_k8s.namespace_mapping.items](data-sources--discovery--reference--group-001.md#canonical-9e9598f66e03e6fab91440d72b944a505ee80ea4cbac945b0911bfcf227559d0) |
| `discovery_k8s.namespace_mapping.items.namespace` | [discovery_k8s.namespace_mapping.items.namespace](data-sources--discovery--reference--group-001.md#canonical-38381323b3bb0c89e48c6546b3b0193394b2c9d2ff60ad9f86d93422eba881ae) |
| `discovery_k8s.namespace_mapping.items.namespace_regex` | [discovery_k8s.namespace_mapping.items.namespace_regex](data-sources--discovery--reference--group-001.md#canonical-2aea89c21e70ba10e28bb8dde08414a1721a0f419ac70fa8dcb2d8778d39c6b5) |
| `discovery_k8s.publish_info` | [discovery_k8s.publish_info](data-sources--discovery--reference--group-001.md#canonical-192a3a7fae1a6a9129f657c9b1c421ff251074118a21d10da85c94a53c2be19a) |
| `discovery_k8s.publish_info.disable_spec` | [discovery_k8s.publish_info.disable_spec](data-sources--discovery--reference--group-001.md#canonical-52b8f7175856174ef88ae639e721c45d23c135176a1e3da9de9b33eb4fed5b30) |
| `discovery_k8s.publish_info.dns_delegation` | [discovery_k8s.publish_info.dns_delegation](data-sources--discovery--reference--group-001.md#canonical-49f48ef2bf0985d6827e277621213ea49868fe1982490c32aa193b0e738ce2ae) |
| `discovery_k8s.publish_info.dns_delegation.dns_mode` | [discovery_k8s.publish_info.dns_delegation.dns_mode](data-sources--discovery--reference--group-001.md#canonical-2d5894580f32d3d26e8c8668ebdea16ba28032768be85322014356b828253b19) |
| `discovery_k8s.publish_info.dns_delegation.subdomain` | [discovery_k8s.publish_info.dns_delegation.subdomain](data-sources--discovery--reference--group-001.md#canonical-1c667bf1dfe58e8b2b8a07f34335bfe70d68cbba9824669895b0d1deecd27743) |
| `discovery_k8s.publish_info.publish` | [discovery_k8s.publish_info.publish](data-sources--discovery--reference--group-001.md#canonical-2272daaf79426f53c460ded46ce5f3cfdce2ed2976fa7778e6cce203e6f41ec0) |
| `discovery_k8s.publish_info.publish.namespace` | [discovery_k8s.publish_info.publish.namespace](data-sources--discovery--reference--group-001.md#canonical-ebd927b0d328a2f1883113c55afada96ca8459c43b2cfaa5c821370393e78dd7) |
| `discovery_k8s.publish_info.publish_fqdns` | [discovery_k8s.publish_info.publish_fqdns](data-sources--discovery--reference--group-001.md#canonical-3c8b9c87aa16aae76cff5b00b3a4a6a17103a032e05bef85694ee90d1b16efa2) |
| `id` | [id](data-sources--discovery--reference--group-001.md#canonical-3fe6cbe70d98040df256d9d44a6f5c6fea97c464e6ff7035f53749c78e9a8d32) |
| `labels` | [labels](data-sources--discovery--reference--group-001.md#canonical-ebaa16e2be580bc8ef8bca65d1322a99cc2e3f4c4605ccce680ac1d31c3083da) |
| `name` | [name](data-sources--discovery--reference--group-001.md#canonical-efc4a2874030edc499c5eb4f7e559983b099da47ace2b724e12f707bcaee770e) |
| `namespace` | [namespace](data-sources--discovery--reference--group-001.md#canonical-32b669e6f69aebabb5bbc4af7e9438f92b1613da68289affce502c803211c32d) |
| `no_cluster_id` | [no_cluster_id](data-sources--discovery--reference--group-001.md#canonical-bc76a7116b66e700f48d50b0d53306bb8e2a03efd8228c74ee271680471c69b2) |
| `where` | [where](data-sources--discovery--reference--group-001.md#canonical-412f8957dcafc21637e2f3135813f3e29211c723a4b2e45916d1f8f210f80c9f) |
| `where.site` | [where.site](data-sources--discovery--reference--group-001.md#canonical-be31b01b7e0ba49133be07099144e678e12d337122c7b7b13426fd153d6ec825) |
| `where.site.disable_internet_vip` | [where.site.disable_internet_vip](data-sources--discovery--reference--group-001.md#canonical-8c3705fb0452ba438b37203e22bb6f5ce00f00e4a7d8b810d20de19e24ecc33f) |
| `where.site.enable_internet_vip` | [where.site.enable_internet_vip](data-sources--discovery--reference--group-001.md#canonical-42218dbffdb319e12db9c73b3be22008ad8bd258f04d84a8e1aca408d741c9cb) |
| `where.site.network_type` | [where.site.network_type](data-sources--discovery--reference--group-001.md#canonical-c505429a491c24c9d5a054ccdfeeb73b01e403c965c1bc2f7bddc69008a5e059) |
| `where.site.ref` | [where.site.ref](data-sources--discovery--reference--group-001.md#canonical-0714638296654be35a83072e86fecd2d78fd28093a3fb89b93fed18d3c4c852f) |
| `where.site.ref.kind` | [where.site.ref.kind](data-sources--discovery--reference--group-001.md#canonical-1c17d4ace85d6a62bfc1732948879311d2db243c30f14a91f2a1535e239c78e4) |
| `where.site.ref.name` | [where.site.ref.name](data-sources--discovery--reference--group-001.md#canonical-ae21c70c0ca97266edb8e2ec339e9145b107f71a3bdf27a7e10568ffdd5034cf) |
| `where.site.ref.namespace` | [where.site.ref.namespace](data-sources--discovery--reference--group-001.md#canonical-8d9d9b86ea6d40ba25811b646e5315ba53e0596dbd90e6814a88caf70de0cfa2) |
| `where.site.ref.tenant` | [where.site.ref.tenant](data-sources--discovery--reference--group-001.md#canonical-1419baeb4b93f03524ed4b8e757229fe74e0a6a9d69fbb1915bef012caa4d33c) |
| `where.site.ref.uid` | [where.site.ref.uid](data-sources--discovery--reference--group-001.md#canonical-f6327f377e379351bb1c0c07fb8152e2bb1bab0f06aacbcfe28a8025a261795b) |
| `where.virtual_network` | [where.virtual_network](data-sources--discovery--reference--group-001.md#canonical-2b3126429bb09f20d3825ee12fdd9812ccf18bfcb309fea263f052fee9fadd62) |
| `where.virtual_network.ref` | [where.virtual_network.ref](data-sources--discovery--reference--group-001.md#canonical-75078d1d10cb0092b83a7b12900c29cc3b758c714dbe1e659af7d38564926527) |
| `where.virtual_network.ref.kind` | [where.virtual_network.ref.kind](data-sources--discovery--reference--group-001.md#canonical-9609e98183bd88f5cee50619e83769e6198a6569b794ed498687cd93a5ba1555) |
| `where.virtual_network.ref.name` | [where.virtual_network.ref.name](data-sources--discovery--reference--group-001.md#canonical-124e5a2331fc15954e1b880e62ab6d0675e980b59ba98e69c9ccbc94e46dd1ec) |
| `where.virtual_network.ref.namespace` | [where.virtual_network.ref.namespace](data-sources--discovery--reference--group-001.md#canonical-992a4d08843ae283c3fd17fb428f602deed8aa4948ff146dd45a4626039065ae) |
| `where.virtual_network.ref.tenant` | [where.virtual_network.ref.tenant](data-sources--discovery--reference--group-001.md#canonical-1a39c07a6d89b4a9d9243647445c3f2d2d95a7076b7190a2a146c901e12ffd44) |
| `where.virtual_network.ref.uid` | [where.virtual_network.ref.uid](data-sources--discovery--reference--group-001.md#canonical-b3eb9a71d419ead07f15836c4daa902de6f7f1b73d3c369560dd91622e9b4d9a) |
| `where.virtual_site` | [where.virtual_site](data-sources--discovery--reference--group-001.md#canonical-d96be290ba2a7a660ca6220f0712df65ee278686d6e49f2135969d5f3f258e77) |
| `where.virtual_site.disable_internet_vip` | [where.virtual_site.disable_internet_vip](data-sources--discovery--reference--group-002.md#canonical-d9104db4861bcae21fcf55a77095b3d1d6c5db6a60512269dbf0c7063b9ccb46) |
| `where.virtual_site.enable_internet_vip` | [where.virtual_site.enable_internet_vip](data-sources--discovery--reference--group-002.md#canonical-c569b7f8739477b7566c50ea03b9744cd4692fd27a8e6e89490e2f54c7fc1983) |
| `where.virtual_site.network_type` | [where.virtual_site.network_type](data-sources--discovery--reference--group-001.md#canonical-f823c97152688c3a87f4cbf0d7086dcb37081bedd08fc6557d1f608a00b853c9) |
| `where.virtual_site.ref` | [where.virtual_site.ref](data-sources--discovery--reference--group-002.md#canonical-030ac37ee703ac2d369a2be44223ddf8824ed4f6a1864bf80e2621f4c1d6ef07) |
| `where.virtual_site.ref.kind` | [where.virtual_site.ref.kind](data-sources--discovery--reference--group-002.md#canonical-6c01f9b77285ce4730972df0d39e1ca49122cd072b35f9d7f888295addf04f31) |
| `where.virtual_site.ref.name` | [where.virtual_site.ref.name](data-sources--discovery--reference--group-002.md#canonical-994947f7d3ceb1a0c6a10d9f75c0807ad7da7ad3715ff63df35fc40d7536ebe0) |
| `where.virtual_site.ref.namespace` | [where.virtual_site.ref.namespace](data-sources--discovery--reference--group-002.md#canonical-36af18a6a0a671c12c65f58437c00cca7eb9a5f45222f1289532d40aa599c9f8) |
| `where.virtual_site.ref.tenant` | [where.virtual_site.ref.tenant](data-sources--discovery--reference--group-002.md#canonical-b85e11bac54c66e5593f2af60eab2a9cc6787d3c08ef8ffe84b3503bb3c66205) |
| `where.virtual_site.ref.uid` | [where.virtual_site.ref.uid](data-sources--discovery--reference--group-002.md#canonical-e671acb41888bda6f18111d4d2cda8e0c30a155881207228ceaf1c9bd115a15a) |

<a id="canonical-9bd8134a0e9d49902145fc0b2952bcc5beff2904d8e9b9b0b32ede5ea4ab99ba"></a>

## Next pages — Property reference / 94c550cf146d / 12

- [discovery_consul](data-sources--discovery--reference--group-001.md#canonical-4393fc86fa2ced1c5d3db9a1bb2d4bd5460a72aa837331742b19239501c44294)
- [discovery_k8s](data-sources--discovery--reference--group-001.md#canonical-3c1f55627ad4d68af88df52467484fc60a52655cab4af4210a019b3b79d1b17a)
- [no_cluster_id](data-sources--discovery--reference--group-001.md#canonical-cab961f6582c0f7f0205c82d8b10a272b9e0fe4b779f443dc0de9bb85e6b059c)
- [where](data-sources--discovery--reference--group-001.md#canonical-86e8aa7a75f491571c4a99862108276e04b016f3b24020940f1e1b752bd38094)
- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)

<a id="canonical-4393fc86fa2ced1c5d3db9a1bb2d4bd5460a72aa837331742b19239501c44294"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a024f255e584a8ffa42b3c0011e836c77937537d9a6a036255c5618192d38c3c"></a>

## discovery_consul — discovery_consul / 458b64fe7426 / 2

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-89b3e09314ef5fec837c9034f68d6bb66048b1e8e127c21bebfbd751a57f9668)
- discovery_consul

<a id="canonical-c1d9546703b5f7b256c015fc7cab2f48155c68c079e1020ab239ae4d0795cb99"></a>

Type: `"single"`. Computed.

\[OneOf: discovery\_consul, discovery\_k8s\] Discovery configuration for Hashicorp Consul.

Upstream description:

Discovery configuration for Hashicorp Consul.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-namespace_mapping_choice": "[]"
}
```

OneOf alternatives in this subsection:

- [discovery_consul](data-sources--discovery--reference--group-001.md#canonical-c1d9546703b5f7b256c015fc7cab2f48155c68c079e1020ab239ae4d0795cb99)
- [discovery_k8s](data-sources--discovery--reference--group-001.md#canonical-4eaf3060e9826836c611ea80e9e98f58914ccec81dfe4d1950164e297efae8c5)

Select alternatives according to the provider validators above.

<a id="canonical-ffe72db7a8304a54c19c5de395e1ee74477639f4ed4e79bf4ead37b8736653f3"></a>

## Direct properties — discovery_consul / 458b64fe7426 / 3

- [access_info](data-sources--discovery--reference--group-001.md#canonical-dbd74d13ab2d8d75be2a2a15efdec4d79bf0109f7f3c27be65fa908df1835561): complete subsection reference.

- [publish_info](data-sources--discovery--reference--group-001.md#canonical-caee273d7046670f01d1f5af2c509b030a9ff71d6c10c2156e6275edf310be77): complete subsection reference.

<a id="canonical-63d59b0962081b3dc6b2359a8fcf0dbf9e07301209a9b7e1b81dfda7d5429a92"></a>

## Next pages — discovery_consul / 458b64fe7426 / 4

- [discovery_consul.access_info](data-sources--discovery--reference--group-001.md#canonical-dbd74d13ab2d8d75be2a2a15efdec4d79bf0109f7f3c27be65fa908df1835561)
- [discovery_consul.publish_info](data-sources--discovery--reference--group-001.md#canonical-caee273d7046670f01d1f5af2c509b030a9ff71d6c10c2156e6275edf310be77)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-89b3e09314ef5fec837c9034f68d6bb66048b1e8e127c21bebfbd751a57f9668)
- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)

<a id="canonical-dbd74d13ab2d8d75be2a2a15efdec4d79bf0109f7f3c27be65fa908df1835561"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7970c054641e0d65e21d41e000a25e86a02432b86ff4c31baf928339727d4ac8"></a>

## discovery_consul.access_info — discovery_consul.access_info / 414089f9ebb2 / 2

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-89b3e09314ef5fec837c9034f68d6bb66048b1e8e127c21bebfbd751a57f9668)
- [discovery_consul](data-sources--discovery--reference--group-001.md#canonical-4393fc86fa2ced1c5d3db9a1bb2d4bd5460a72aa837331742b19239501c44294)
- discovery_consul.access_info

<a id="canonical-ba12dd1beadf83ba3cfbf9deb29f01fab9ec1cefba82b10ec50062cc8eca490d"></a>

Type: `"single"`. Computed.

Hashicorp Consul API server information.

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

<a id="canonical-ec2c384d6a07dd074042d09426a02149500ff650b1c8eb824f46b2284a506d31"></a>

## Direct properties — discovery_consul.access_info / 414089f9ebb2 / 3

- [connection_info](data-sources--discovery--reference--group-001.md#canonical-7b2984e1958a890f9a424de227be812f67094bf5a554de2722d30456024c5542): complete subsection reference.

- [http_basic_auth_info](data-sources--discovery--reference--group-001.md#canonical-e0847ac5fa8b1bd05a95904b7a877edd387f39ab562f0c39ce3167b0dbba7b83): complete subsection reference.

<a id="canonical-3a95bfd0f3aad294c24ea6c4e73c4ffd5392426781363afc0399f39126b775db"></a>

## Next pages — discovery_consul.access_info / 414089f9ebb2 / 4

- [discovery_consul.access_info.connection_info](data-sources--discovery--reference--group-001.md#canonical-7b2984e1958a890f9a424de227be812f67094bf5a554de2722d30456024c5542)
- [discovery_consul.access_info.http_basic_auth_info](data-sources--discovery--reference--group-001.md#canonical-e0847ac5fa8b1bd05a95904b7a877edd387f39ab562f0c39ce3167b0dbba7b83)
- [discovery_consul](data-sources--discovery--reference--group-001.md#canonical-4393fc86fa2ced1c5d3db9a1bb2d4bd5460a72aa837331742b19239501c44294)
- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)

<a id="canonical-7b2984e1958a890f9a424de227be812f67094bf5a554de2722d30456024c5542"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-28d2bc3b81ec77f33a83510095dfcdc3ab9efe7674aef8e850950fdc84e5a066"></a>

## discovery_consul.access_info.connection_info — discovery_consul.access_info.connection_info / 0f1a3b5f3e19 / 2

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-89b3e09314ef5fec837c9034f68d6bb66048b1e8e127c21bebfbd751a57f9668)
- [discovery_consul](data-sources--discovery--reference--group-001.md#canonical-4393fc86fa2ced1c5d3db9a1bb2d4bd5460a72aa837331742b19239501c44294)
- [discovery_consul.access_info](data-sources--discovery--reference--group-001.md#canonical-dbd74d13ab2d8d75be2a2a15efdec4d79bf0109f7f3c27be65fa908df1835561)
- discovery_consul.access_info.connection_info

<a id="canonical-5456f17afb0d25aa7750df9abec8dc58d2ec8e1f26435c5601415828fb58eedd"></a>

Type: `"single"`. Computed.

Configuration details to access discovery service REST API.

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

<a id="canonical-00201e654d721c0d814ea6b90ac762c6935919372d6955ba9e56920a6bb26f9b"></a>

## Direct properties — discovery_consul.access_info.connection_info / 0f1a3b5f3e19 / 3

<a id="canonical-69f0c3e02ff0e08a0e266cc6594132749cd204c70899b55393bdf84bfe7ad480"></a>

<a id="canonical-15b204f6e50cc8cd3ef1bbc2dade8242f76573356b686af05fd1a5fb8ed3cea9"></a>

## api_server property — discovery_consul.access_info.connection_info / 0f1a3b5f3e19 / 4

Type: `"string"`. Computed.

API server must be a fully qualified domain string and port specified as host:port pair.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  }
}
```

- [tls_info](data-sources--discovery--reference--group-001.md#canonical-d808d8f649712258acdc3df1710d90f26c24ee600924a604b1eab386d83b67e8): complete subsection reference.

<a id="canonical-671d9b4ffa18e552fc45f547e278a13d564aa8ff61217df007043095d77fa374"></a>

## Next pages — discovery_consul.access_info.connection_info / 0f1a3b5f3e19 / 5

- [discovery_consul.access_info.connection_info.tls_info](data-sources--discovery--reference--group-001.md#canonical-d808d8f649712258acdc3df1710d90f26c24ee600924a604b1eab386d83b67e8)
- [discovery_consul.access_info](data-sources--discovery--reference--group-001.md#canonical-dbd74d13ab2d8d75be2a2a15efdec4d79bf0109f7f3c27be65fa908df1835561)
- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)

<a id="canonical-d808d8f649712258acdc3df1710d90f26c24ee600924a604b1eab386d83b67e8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-97d874dc8988c50dfb15d7a183dd52339f110baa881f61717762f14dbece6e8d"></a>

## discovery_consul.access_info.connection_info.tls_info — discovery_consul.access_info.connection_info.tls_info / 3619734d008d / 2

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-89b3e09314ef5fec837c9034f68d6bb66048b1e8e127c21bebfbd751a57f9668)
- [discovery_consul](data-sources--discovery--reference--group-001.md#canonical-4393fc86fa2ced1c5d3db9a1bb2d4bd5460a72aa837331742b19239501c44294)
- [discovery_consul.access_info](data-sources--discovery--reference--group-001.md#canonical-dbd74d13ab2d8d75be2a2a15efdec4d79bf0109f7f3c27be65fa908df1835561)
- [discovery_consul.access_info.connection_info](data-sources--discovery--reference--group-001.md#canonical-7b2984e1958a890f9a424de227be812f67094bf5a554de2722d30456024c5542)
- discovery_consul.access_info.connection_info.tls_info

<a id="canonical-50ed6c6fb50989c508aa76b2fda47ae513a22612927c4b32f846289c285fc3dc"></a>

Type: `"single"`. Computed.

TLS config for client of discovery service.

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

<a id="canonical-970c783fc6f2a3a9ff738f769c099769b781d8fb35b22d2d5dcbc5c3b6a585a9"></a>

## Direct properties — discovery_consul.access_info.connection_info.tls_info / 3619734d008d / 3

<a id="canonical-ee897d9c4f480ed2b42ed3cee9e7c6e485b7088bfda90101007121d7e87991ef"></a>

<a id="canonical-be99c30a06baff98069ecf1551700c009578065a48b34ab6a61113bf331f3a7c"></a>

## certificate property — discovery_consul.access_info.connection_info.tls_info / 3619734d008d / 4

Type: `"string"`. Computed.

Client certificate is PEM-encoded certificate or certificate-chain.

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
    "formatDescription": "PEM-encoded X.509 certificate, max 5MB",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 100,
    "pattern": "^-----BEGIN CERTIFICATE-----\\n.*\\n-----END CERTIFICATE-----$",
    "validation": {
      "standard": "PEM"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [key_url](data-sources--discovery--reference--group-001.md#canonical-75113c054fae908c5d56057c33de8b3793573cd498b1574a3d154b145f0c5669): complete subsection reference.

<a id="canonical-c01a2b8b64fd17c88c5c26f51fd5a252f5e3ae1aaacbf7b8b59043f09bbedf64"></a>

<a id="canonical-2f43504ef2ceff8ed9b520268f467394462aa37fe77db136292cdb4a843e8e2c"></a>

## server_name property — discovery_consul.access_info.connection_info.tls_info / 3619734d008d / 5

Type: `"string"`. Computed.

ServerName is passed to the server for SNI and is used in the client to check server certificates
against. If ServerName is empty, the hostname used to contact the server is used.

Upstream description:

ServerName is passed to the server for SNI and is used in the client to check server certificates
against. If ServerName is empty, the hostname used to contact the server is used.

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

<a id="canonical-37ed59cd45678f7d911cf5a5a09909a5072f7cc8b64fb9145d9c8a4ca2e40b35"></a>

<a id="canonical-68afb437df4c473d768ac749a86addfbd4189759f858ce15e57aff2a4823b069"></a>

## trusted_ca_url property — discovery_consul.access_info.connection_info.tls_info / 3619734d008d / 6

Type: `"string"`. Computed.

The URL or value for trusted Server CA certificate or certificate chain Certificates in PEM format
including the PEM headers.

Upstream description:

The URL or value for trusted Server CA certificate or certificate chain Certificates in PEM format
including the PEM headers.

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
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-c743d4cab3131882f62a79eacc87a18debc8cf85d4a34c3c64956f76aea18c4c"></a>

## Next pages — discovery_consul.access_info.connection_info.tls_info / 3619734d008d / 7

- [discovery_consul.access_info.connection_info.tls_info.key_url](data-sources--discovery--reference--group-001.md#canonical-75113c054fae908c5d56057c33de8b3793573cd498b1574a3d154b145f0c5669)
- [discovery_consul.access_info.connection_info](data-sources--discovery--reference--group-001.md#canonical-7b2984e1958a890f9a424de227be812f67094bf5a554de2722d30456024c5542)
- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)

<a id="canonical-75113c054fae908c5d56057c33de8b3793573cd498b1574a3d154b145f0c5669"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ce7dc38f6ef535dd17cd0bf3ebaa21b289a403b91f7d506a032d66b4f4583f98"></a>

## discovery_consul.access_info.connection_info.tls_info.key_url — discovery_consul.access_info.connection_info.tls_info.key_url / 77a50711f59a / 2

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-89b3e09314ef5fec837c9034f68d6bb66048b1e8e127c21bebfbd751a57f9668)
- [discovery_consul](data-sources--discovery--reference--group-001.md#canonical-4393fc86fa2ced1c5d3db9a1bb2d4bd5460a72aa837331742b19239501c44294)
- [discovery_consul.access_info](data-sources--discovery--reference--group-001.md#canonical-dbd74d13ab2d8d75be2a2a15efdec4d79bf0109f7f3c27be65fa908df1835561)
- [discovery_consul.access_info.connection_info](data-sources--discovery--reference--group-001.md#canonical-7b2984e1958a890f9a424de227be812f67094bf5a554de2722d30456024c5542)
- [discovery_consul.access_info.connection_info.tls_info](data-sources--discovery--reference--group-001.md#canonical-d808d8f649712258acdc3df1710d90f26c24ee600924a604b1eab386d83b67e8)
- discovery_consul.access_info.connection_info.tls_info.key_url

<a id="canonical-c5e65ce87690c1898dd75d217b0e95f97ae2ad5066ac8cee9032622b1b452a7e"></a>

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

<a id="canonical-676b79a00cac44746bf0dd5aff478afd6adf34b4a089684de3c2c5845f3f5d9b"></a>

## Direct properties — discovery_consul.access_info.connection_info.tls_info.key_url / 77a50711f59a / 3

- [blindfold_secret_info](data-sources--discovery--reference--group-001.md#canonical-e520f7e756143a3078203f0fbc44d2a472e06df176f4da1546599fc077c9193c): complete subsection reference.

- [clear_secret_info](data-sources--discovery--reference--group-001.md#canonical-ef135e436b48133bef0bf51210e795517a6b24012595754b816001504d13b49b): complete subsection reference.

<a id="canonical-e4bda5642871c602e4d3e8403115db8a0d7af0db8c596192ce4862c5cec9d247"></a>

## Next pages — discovery_consul.access_info.connection_info.tls_info.key_url / 77a50711f59a / 4

- [discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info](data-sources--discovery--reference--group-001.md#canonical-e520f7e756143a3078203f0fbc44d2a472e06df176f4da1546599fc077c9193c)
- [discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info](data-sources--discovery--reference--group-001.md#canonical-ef135e436b48133bef0bf51210e795517a6b24012595754b816001504d13b49b)
- [discovery_consul.access_info.connection_info.tls_info](data-sources--discovery--reference--group-001.md#canonical-d808d8f649712258acdc3df1710d90f26c24ee600924a604b1eab386d83b67e8)
- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)

<a id="canonical-e520f7e756143a3078203f0fbc44d2a472e06df176f4da1546599fc077c9193c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-db0aba3d56c71b80fc653cfcd6b2114369ba005b181e97d2323059affe5cb299"></a>

## discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info — discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_i / 02e101a2d4ee / 2

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-89b3e09314ef5fec837c9034f68d6bb66048b1e8e127c21bebfbd751a57f9668)
- [discovery_consul](data-sources--discovery--reference--group-001.md#canonical-4393fc86fa2ced1c5d3db9a1bb2d4bd5460a72aa837331742b19239501c44294)
- [discovery_consul.access_info](data-sources--discovery--reference--group-001.md#canonical-dbd74d13ab2d8d75be2a2a15efdec4d79bf0109f7f3c27be65fa908df1835561)
- [discovery_consul.access_info.connection_info](data-sources--discovery--reference--group-001.md#canonical-7b2984e1958a890f9a424de227be812f67094bf5a554de2722d30456024c5542)
- [discovery_consul.access_info.connection_info.tls_info](data-sources--discovery--reference--group-001.md#canonical-d808d8f649712258acdc3df1710d90f26c24ee600924a604b1eab386d83b67e8)
- [discovery_consul.access_info.connection_info.tls_info.key_url](data-sources--discovery--reference--group-001.md#canonical-75113c054fae908c5d56057c33de8b3793573cd498b1574a3d154b145f0c5669)
- discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info

<a id="canonical-a910145e2051d4753e799141d70d85fb11d12e66d2b28051da91bdda31d56a3e"></a>

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

<a id="canonical-8d3a90b3a22917965fc10c20eb4edd05b09de5c70749770f1a796123b4dd54ce"></a>

## Direct properties — discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_i / 02e101a2d4ee / 3

<a id="canonical-fcf0a7c14a5886ee85338f117c2aafb1f294e81ca0239b4d78eb5d37d18171c4"></a>

<a id="canonical-32f510249793f2cad1a4dd3cc7ee28203778530b13fe5591baffca4e7d895285"></a>

## decryption_provider property — discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_i / 02e101a2d4ee / 4

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

<a id="canonical-070fca5a417d810d3aefb317f7200c746d8c8ffc74262a330ed3dbeaf8542681"></a>

<a id="canonical-d167337a550233a1b4677ee3fc0960bc660f2a3cb209fa8562faefb97c92da78"></a>

## location property — discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_i / 02e101a2d4ee / 5

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

<a id="canonical-84e7296e0061ef95aa0f5b2261d6e00c59facfef3b48f14bcf367db00f7aabb9"></a>

<a id="canonical-abe409a73774e3cc69d3e930f7d354545a290c5a990cb1ea73afcbda77364b90"></a>

## store_provider property — discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_i / 02e101a2d4ee / 6

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

<a id="canonical-477b2b0d303edf7aa988727cb0fb51917c24e1129ad81537e6cfbb3f676386b3"></a>

## Next pages — discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_i / 02e101a2d4ee / 7

- [discovery_consul.access_info.connection_info.tls_info.key_url](data-sources--discovery--reference--group-001.md#canonical-75113c054fae908c5d56057c33de8b3793573cd498b1574a3d154b145f0c5669)
- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)

<a id="canonical-ef135e436b48133bef0bf51210e795517a6b24012595754b816001504d13b49b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-526f6491da90396d4027af3eca921f4c7edc7afa1418d23b7da4dd445a22cdef"></a>

## discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info — discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info / 42f3732721f1 / 2

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-89b3e09314ef5fec837c9034f68d6bb66048b1e8e127c21bebfbd751a57f9668)
- [discovery_consul](data-sources--discovery--reference--group-001.md#canonical-4393fc86fa2ced1c5d3db9a1bb2d4bd5460a72aa837331742b19239501c44294)
- [discovery_consul.access_info](data-sources--discovery--reference--group-001.md#canonical-dbd74d13ab2d8d75be2a2a15efdec4d79bf0109f7f3c27be65fa908df1835561)
- [discovery_consul.access_info.connection_info](data-sources--discovery--reference--group-001.md#canonical-7b2984e1958a890f9a424de227be812f67094bf5a554de2722d30456024c5542)
- [discovery_consul.access_info.connection_info.tls_info](data-sources--discovery--reference--group-001.md#canonical-d808d8f649712258acdc3df1710d90f26c24ee600924a604b1eab386d83b67e8)
- [discovery_consul.access_info.connection_info.tls_info.key_url](data-sources--discovery--reference--group-001.md#canonical-75113c054fae908c5d56057c33de8b3793573cd498b1574a3d154b145f0c5669)
- discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info

<a id="canonical-21684ee32f794c44a302d9214c4974feef83bef0c6530b50f0dbce70f98c9aef"></a>

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

<a id="canonical-594975716536533ff29a86a21a354dbb940ea59e342e04f9705fd0dc43fa8759"></a>

## Direct properties — discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info / 42f3732721f1 / 3

<a id="canonical-436fcdc85009be9d6a1ea0e111dfc3258e5b46fd58c3ca753bb432f359959f6d"></a>

<a id="canonical-314e2401d512a161c3f8bdfe9c9d996fb653de5633154f3e38027b1630ea643c"></a>

## provider_ref property — discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info / 42f3732721f1 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-63f6721cc80db35076930a3db1beaa33c4a65f442a1058f022a23362485c3a6d"></a>

<a id="canonical-e6849bec823ccf2bc6c02e2831aa5655e725701004f0b5d6e01ce346718b5f32"></a>

## url property — discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info / 42f3732721f1 / 5

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

<a id="canonical-328d7082edd9aab125ccebd2fd8535af16654cb10249e881cc629659cea9dc5b"></a>

## Next pages — discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info / 42f3732721f1 / 6

- [discovery_consul.access_info.connection_info.tls_info.key_url](data-sources--discovery--reference--group-001.md#canonical-75113c054fae908c5d56057c33de8b3793573cd498b1574a3d154b145f0c5669)
- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)

<a id="canonical-e0847ac5fa8b1bd05a95904b7a877edd387f39ab562f0c39ce3167b0dbba7b83"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cdafe4c2a712b129f7ca067149e556f8238d15cf97a5fe616b434936cded7281"></a>

## discovery_consul.access_info.http_basic_auth_info — discovery_consul.access_info.http_basic_auth_info / c3799bd2ef7a / 2

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-89b3e09314ef5fec837c9034f68d6bb66048b1e8e127c21bebfbd751a57f9668)
- [discovery_consul](data-sources--discovery--reference--group-001.md#canonical-4393fc86fa2ced1c5d3db9a1bb2d4bd5460a72aa837331742b19239501c44294)
- [discovery_consul.access_info](data-sources--discovery--reference--group-001.md#canonical-dbd74d13ab2d8d75be2a2a15efdec4d79bf0109f7f3c27be65fa908df1835561)
- discovery_consul.access_info.http_basic_auth_info

<a id="canonical-daf091cf7f13cf207dc3d0cbd1196a3437ff74d803a33ba2f967941cdadcd5ea"></a>

Type: `"single"`. Computed.

Authentication parameters to access Hashicorp Consul.

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

<a id="canonical-f7cdace98f27d2e5d80e2f16841e78b8b8d2fe08ed2b545b92c5ab29bd5ef60e"></a>

## Direct properties — discovery_consul.access_info.http_basic_auth_info / c3799bd2ef7a / 3

- [passwd_url](data-sources--discovery--reference--group-001.md#canonical-0dfd0c2e1e1babc7dc67014f2969ef4567e7ef312f265bdc99385726039b70ec): complete subsection reference.

<a id="canonical-620c15e78dfc41e80c4d7650ffb488d9ab8101079b7aa74cec75f65737490d95"></a>

<a id="canonical-9204161f4c7c92403ea53133e294923d008c49e6d37bb70a424173f28eb7ade1"></a>

## user_name property — discovery_consul.access_info.http_basic_auth_info / c3799bd2ef7a / 4

Type: `"string"`. Computed.

User Name. Username in consul.

Upstream description:

Username in consul.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-35c6918014cfe920e1aa65f4687d6be3f99c56d55e7f7a4fa1d303e42c0baef0"></a>

## Next pages — discovery_consul.access_info.http_basic_auth_info / c3799bd2ef7a / 5

- [discovery_consul.access_info.http_basic_auth_info.passwd_url](data-sources--discovery--reference--group-001.md#canonical-0dfd0c2e1e1babc7dc67014f2969ef4567e7ef312f265bdc99385726039b70ec)
- [discovery_consul.access_info](data-sources--discovery--reference--group-001.md#canonical-dbd74d13ab2d8d75be2a2a15efdec4d79bf0109f7f3c27be65fa908df1835561)
- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)

<a id="canonical-0dfd0c2e1e1babc7dc67014f2969ef4567e7ef312f265bdc99385726039b70ec"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9d62e0ca740b21dae7af9e0bc78141cca210e498e997ef69412f61a629377241"></a>

## discovery_consul.access_info.http_basic_auth_info.passwd_url — discovery_consul.access_info.http_basic_auth_info.passwd_url / 8ad0ba111f04 / 2

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-89b3e09314ef5fec837c9034f68d6bb66048b1e8e127c21bebfbd751a57f9668)
- [discovery_consul](data-sources--discovery--reference--group-001.md#canonical-4393fc86fa2ced1c5d3db9a1bb2d4bd5460a72aa837331742b19239501c44294)
- [discovery_consul.access_info](data-sources--discovery--reference--group-001.md#canonical-dbd74d13ab2d8d75be2a2a15efdec4d79bf0109f7f3c27be65fa908df1835561)
- [discovery_consul.access_info.http_basic_auth_info](data-sources--discovery--reference--group-001.md#canonical-e0847ac5fa8b1bd05a95904b7a877edd387f39ab562f0c39ce3167b0dbba7b83)
- discovery_consul.access_info.http_basic_auth_info.passwd_url

<a id="canonical-45d482293334f47f327af159ed19cc596d1e6181baec335c6965c187d3c9015d"></a>

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

<a id="canonical-d5991e118154e505def742d8268f81b5e00b3e7a5675fa10f8dc0cd352c22619"></a>

## Direct properties — discovery_consul.access_info.http_basic_auth_info.passwd_url / 8ad0ba111f04 / 3

- [blindfold_secret_info](data-sources--discovery--reference--group-001.md#canonical-ab79d5f963a3beebe2d058c9f9c54a697b2a61b69a4c80d2409971103e95c0a9): complete subsection reference.

- [clear_secret_info](data-sources--discovery--reference--group-001.md#canonical-fa03072e8c22883956b2f2eb09be343238de698a0fd8a4a427cf0e2255275434): complete subsection reference.

<a id="canonical-7bd0e51ebf56ee887fdff20e4e30817d54739e5022d044a7387c028b89bd8ec4"></a>

## Next pages — discovery_consul.access_info.http_basic_auth_info.passwd_url / 8ad0ba111f04 / 4

- [discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info](data-sources--discovery--reference--group-001.md#canonical-ab79d5f963a3beebe2d058c9f9c54a697b2a61b69a4c80d2409971103e95c0a9)
- [discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info](data-sources--discovery--reference--group-001.md#canonical-fa03072e8c22883956b2f2eb09be343238de698a0fd8a4a427cf0e2255275434)
- [discovery_consul.access_info.http_basic_auth_info](data-sources--discovery--reference--group-001.md#canonical-e0847ac5fa8b1bd05a95904b7a877edd387f39ab562f0c39ce3167b0dbba7b83)
- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)

<a id="canonical-ab79d5f963a3beebe2d058c9f9c54a697b2a61b69a4c80d2409971103e95c0a9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f6c4ca1bc3f0c599b92cabcb0ce76770f60b7832fe30c8776d7293b182ca14e2"></a>

## discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info — discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_in / adaf3041d945 / 2

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-89b3e09314ef5fec837c9034f68d6bb66048b1e8e127c21bebfbd751a57f9668)
- [discovery_consul](data-sources--discovery--reference--group-001.md#canonical-4393fc86fa2ced1c5d3db9a1bb2d4bd5460a72aa837331742b19239501c44294)
- [discovery_consul.access_info](data-sources--discovery--reference--group-001.md#canonical-dbd74d13ab2d8d75be2a2a15efdec4d79bf0109f7f3c27be65fa908df1835561)
- [discovery_consul.access_info.http_basic_auth_info](data-sources--discovery--reference--group-001.md#canonical-e0847ac5fa8b1bd05a95904b7a877edd387f39ab562f0c39ce3167b0dbba7b83)
- [discovery_consul.access_info.http_basic_auth_info.passwd_url](data-sources--discovery--reference--group-001.md#canonical-0dfd0c2e1e1babc7dc67014f2969ef4567e7ef312f265bdc99385726039b70ec)
- discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info

<a id="canonical-870eb4df657ac1ec91a5236f243018cf184ca926df73ad0b0f2e2c2623d82ca0"></a>

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

<a id="canonical-9472614e71afb3fcccb5e3986e92c7894de2d436ed0c48ff3e1971481d741b1d"></a>

## Direct properties — discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_in / adaf3041d945 / 3

<a id="canonical-cac5539856c3c967ee43a3cc3d74195053451e74acea44a1ddf12614e89c9d09"></a>

<a id="canonical-dbbdc0caaa5ea6e5524091180c2790eecd155d6c45b9ab45fce69a460e261f86"></a>

## decryption_provider property — discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_in / adaf3041d945 / 4

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

<a id="canonical-bb3e16006197b4bbf929bd9e3c8e85a2ce4646cfcca13413ca1f8407ea9e9e2d"></a>

<a id="canonical-ecc5440b88432906c7bba0af6c6734f3eb178b93e4b9ef9d50b159dbdf0a4351"></a>

## location property — discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_in / adaf3041d945 / 5

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

<a id="canonical-0206eb23a3184298efc9a15e6326bbc336fbadb9700a334a05bc79a13bc39646"></a>

<a id="canonical-0e51775591e4fe0474596c7b2e145ec14387a55c7dec990bdfbf6c5b62587d63"></a>

## store_provider property — discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_in / adaf3041d945 / 6

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

<a id="canonical-849fc1a90042c6932b268167dd23d066ed82631d4b0c0e517cd09fbfae37f946"></a>

## Next pages — discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_in / adaf3041d945 / 7

- [discovery_consul.access_info.http_basic_auth_info.passwd_url](data-sources--discovery--reference--group-001.md#canonical-0dfd0c2e1e1babc7dc67014f2969ef4567e7ef312f265bdc99385726039b70ec)
- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)

<a id="canonical-fa03072e8c22883956b2f2eb09be343238de698a0fd8a4a427cf0e2255275434"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f5b988d596253d9abeb725c1fd228e76a111e8ac3ee17a8633f8f5ae3293cb50"></a>

## discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info — discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info / 3a87d9905bff / 2

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-89b3e09314ef5fec837c9034f68d6bb66048b1e8e127c21bebfbd751a57f9668)
- [discovery_consul](data-sources--discovery--reference--group-001.md#canonical-4393fc86fa2ced1c5d3db9a1bb2d4bd5460a72aa837331742b19239501c44294)
- [discovery_consul.access_info](data-sources--discovery--reference--group-001.md#canonical-dbd74d13ab2d8d75be2a2a15efdec4d79bf0109f7f3c27be65fa908df1835561)
- [discovery_consul.access_info.http_basic_auth_info](data-sources--discovery--reference--group-001.md#canonical-e0847ac5fa8b1bd05a95904b7a877edd387f39ab562f0c39ce3167b0dbba7b83)
- [discovery_consul.access_info.http_basic_auth_info.passwd_url](data-sources--discovery--reference--group-001.md#canonical-0dfd0c2e1e1babc7dc67014f2969ef4567e7ef312f265bdc99385726039b70ec)
- discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info

<a id="canonical-85f0e32f63c9e9b138bb006fdbd5be4d6d8d8fad916873a4cc32684d1089e9d8"></a>

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

<a id="canonical-9ffa2a03a274ab6730a91fc7837a82c92a47f7c7b829900530e0064b30f1226c"></a>

## Direct properties — discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info / 3a87d9905bff / 3

<a id="canonical-aa749059cd51d4e22462277e2f77a54ed06da1deee699c5ec1148cdf1432c83d"></a>

<a id="canonical-3def486584dae7f0b2993395d2d54b66e0c8613ecd6b5d5a6df4983c1db8b174"></a>

## provider_ref property — discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info / 3a87d9905bff / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-9162e59b4621777503a8f4524b6a65620364aec46b6613259b975d2dda8d7ce5"></a>

<a id="canonical-cb1959ed0288765972a3719f22abfd7e307ef150f2473ebf75d7648b9ff57cbc"></a>

## url property — discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info / 3a87d9905bff / 5

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

<a id="canonical-b84cd60b09adc3ac378efaca16ffc716faf273fd9fea482074ff9d4e84664e1e"></a>

## Next pages — discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info / 3a87d9905bff / 6

- [discovery_consul.access_info.http_basic_auth_info.passwd_url](data-sources--discovery--reference--group-001.md#canonical-0dfd0c2e1e1babc7dc67014f2969ef4567e7ef312f265bdc99385726039b70ec)
- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)

<a id="canonical-caee273d7046670f01d1f5af2c509b030a9ff71d6c10c2156e6275edf310be77"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-322f0cb7aaf3623e9a3d71bd0085de7d33f173062ae1bc6c9d7b142e13599112"></a>

## discovery_consul.publish_info — discovery_consul.publish_info / a25afd4627da / 2

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-89b3e09314ef5fec837c9034f68d6bb66048b1e8e127c21bebfbd751a57f9668)
- [discovery_consul](data-sources--discovery--reference--group-001.md#canonical-4393fc86fa2ced1c5d3db9a1bb2d4bd5460a72aa837331742b19239501c44294)
- discovery_consul.publish_info

<a id="canonical-4ba15100bd73a4c1ee8ee737f199ab6e7760554cfe9ac895703c10fbe3a2b0dc"></a>

Type: `"single"`. Computed.

Configuration parameter for publish info.

Upstream description:

Consul Configuration to publish VIPs.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-publish_choice": "[\"disable\",\"publish\"]"
}
```

<a id="canonical-986884a0b3db498a1abd35991073ea27c933417c26a2ecf6ef678fca270097ae"></a>

## Direct properties — discovery_consul.publish_info / a25afd4627da / 3

- [disable_spec](data-sources--discovery--reference--group-001.md#canonical-28779f35173bb9c6ccad28c87fe258885bb4499746662eb11bd32df0cb7c437b): complete subsection reference.

- [publish](data-sources--discovery--reference--group-001.md#canonical-d781ea62f89403f44067191f69c8b31fad3f59fe6c44e202173c57b7b170c051): complete subsection reference.

<a id="canonical-4c5da818fa1b8a0b968c2910e1551b2581b7620adabfc0be1775749e7364ae1d"></a>

## Next pages — discovery_consul.publish_info / a25afd4627da / 4

- [discovery_consul.publish_info.disable_spec](data-sources--discovery--reference--group-001.md#canonical-28779f35173bb9c6ccad28c87fe258885bb4499746662eb11bd32df0cb7c437b)
- [discovery_consul.publish_info.publish](data-sources--discovery--reference--group-001.md#canonical-d781ea62f89403f44067191f69c8b31fad3f59fe6c44e202173c57b7b170c051)
- [discovery_consul](data-sources--discovery--reference--group-001.md#canonical-4393fc86fa2ced1c5d3db9a1bb2d4bd5460a72aa837331742b19239501c44294)
- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)

<a id="canonical-28779f35173bb9c6ccad28c87fe258885bb4499746662eb11bd32df0cb7c437b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8301d7036c27a8ffdc79d56455229921e3542f71b391250bcf87fc1b4255dfb9"></a>

## discovery_consul.publish_info.disable_spec — discovery_consul.publish_info.disable_spec / fd6933d078cc / 2

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-89b3e09314ef5fec837c9034f68d6bb66048b1e8e127c21bebfbd751a57f9668)
- [discovery_consul](data-sources--discovery--reference--group-001.md#canonical-4393fc86fa2ced1c5d3db9a1bb2d4bd5460a72aa837331742b19239501c44294)
- [discovery_consul.publish_info](data-sources--discovery--reference--group-001.md#canonical-caee273d7046670f01d1f5af2c509b030a9ff71d6c10c2156e6275edf310be77)
- discovery_consul.publish_info.disable_spec

<a id="canonical-9bb4fa3c193891fb854d79b30b15e2f23f6801571ab98e53ba4b62d6ccb097ee"></a>

Type: `["object", {}]`. Computed.

Enable this option

<a id="canonical-2dd3417d8e3c3588037fc35c5ecc0891cd45aede467cf8b34f77c555d2f2a9cd"></a>

## Direct properties — discovery_consul.publish_info.disable_spec / fd6933d078cc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-00dacd076d01fed13695b55190eed42df1a414c140b2eaa5e5d1fd22f7f8439b"></a>

## Next pages — discovery_consul.publish_info.disable_spec / fd6933d078cc / 4

- [discovery_consul.publish_info](data-sources--discovery--reference--group-001.md#canonical-caee273d7046670f01d1f5af2c509b030a9ff71d6c10c2156e6275edf310be77)
- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)

<a id="canonical-d781ea62f89403f44067191f69c8b31fad3f59fe6c44e202173c57b7b170c051"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fe916ebb36b2ffd4f3962f5b505513c54092430fad66ce510e104824b938153f"></a>

## discovery_consul.publish_info.publish — discovery_consul.publish_info.publish / e587afbe5c21 / 2

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-89b3e09314ef5fec837c9034f68d6bb66048b1e8e127c21bebfbd751a57f9668)
- [discovery_consul](data-sources--discovery--reference--group-001.md#canonical-4393fc86fa2ced1c5d3db9a1bb2d4bd5460a72aa837331742b19239501c44294)
- [discovery_consul.publish_info](data-sources--discovery--reference--group-001.md#canonical-caee273d7046670f01d1f5af2c509b030a9ff71d6c10c2156e6275edf310be77)
- discovery_consul.publish_info.publish

<a id="canonical-3b8199c0a5d8aeddde0dde0a4cbde830b5814061e14c676995eed51408e8ac70"></a>

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

<a id="canonical-1ba7181a4203ea520d53bb1f0ba724a2601e0edb263391e6190659a7a2260841"></a>

## Direct properties — discovery_consul.publish_info.publish / e587afbe5c21 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-09abd32ab684e9fdafa2633697d62a95a71ecde7875a9ed01fc5651382fbc94e"></a>

## Next pages — discovery_consul.publish_info.publish / e587afbe5c21 / 4

- [discovery_consul.publish_info](data-sources--discovery--reference--group-001.md#canonical-caee273d7046670f01d1f5af2c509b030a9ff71d6c10c2156e6275edf310be77)
- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)

<a id="canonical-3c1f55627ad4d68af88df52467484fc60a52655cab4af4210a019b3b79d1b17a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ee75c089e7edf8802f23616b1a4349ca84ed824815d0e8afd81e5d9f68b8de3d"></a>

## discovery_k8s — discovery_k8s / 0ab24fee0b17 / 2

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-89b3e09314ef5fec837c9034f68d6bb66048b1e8e127c21bebfbd751a57f9668)
- discovery_k8s

<a id="canonical-4eaf3060e9826836c611ea80e9e98f58914ccec81dfe4d1950164e297efae8c5"></a>

Type: `"single"`. Computed.

Configuration parameter for discovery k8s.

Upstream description:

Discovery configuration for K8s.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-namespace_mapping_choice": "[\"default_all\",\"namespace_mapping\"]"
}
```

<a id="canonical-31c36d3ed2b8d21da76c38e9fe3caaeded7653aabe841569d92d826bec7e6696"></a>

## Direct properties — discovery_k8s / 0ab24fee0b17 / 3

- [access_info](data-sources--discovery--reference--group-001.md#canonical-1588a26e7c162a7f538ee12ef2d0df7a8709f0dcba9a5d7247ee91954a5da1ef): complete subsection reference.

- [default_all](data-sources--discovery--reference--group-001.md#canonical-15326be21e5d6f24d2ccb480b59d2172ad8551beaf05b81dc4d28fd4b49c734c): complete subsection reference.

- [namespace_mapping](data-sources--discovery--reference--group-001.md#canonical-f31f19eafefa2e69b558d06d3bfbc3deac8da1185d9da7f44f2ee640b8a86809): complete subsection reference.

- [publish_info](data-sources--discovery--reference--group-001.md#canonical-8074630174ac5afaea680e91dcf3e042a32b8c5459a26c9262928afb515aa8d6): complete subsection reference.

<a id="canonical-f227e1ad19c66bd7d64e6ecdc54772e90eee7eb448a5026b76c1878165b47a1b"></a>

## Next pages — discovery_k8s / 0ab24fee0b17 / 4

- [discovery_k8s.access_info](data-sources--discovery--reference--group-001.md#canonical-1588a26e7c162a7f538ee12ef2d0df7a8709f0dcba9a5d7247ee91954a5da1ef)
- [discovery_k8s.default_all](data-sources--discovery--reference--group-001.md#canonical-15326be21e5d6f24d2ccb480b59d2172ad8551beaf05b81dc4d28fd4b49c734c)
- [discovery_k8s.namespace_mapping](data-sources--discovery--reference--group-001.md#canonical-f31f19eafefa2e69b558d06d3bfbc3deac8da1185d9da7f44f2ee640b8a86809)
- [discovery_k8s.publish_info](data-sources--discovery--reference--group-001.md#canonical-8074630174ac5afaea680e91dcf3e042a32b8c5459a26c9262928afb515aa8d6)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-89b3e09314ef5fec837c9034f68d6bb66048b1e8e127c21bebfbd751a57f9668)
- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)

<a id="canonical-1588a26e7c162a7f538ee12ef2d0df7a8709f0dcba9a5d7247ee91954a5da1ef"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-45654c0c16ef8d45fd75f094cf29ae23337746f8a1c26dbc6bd0a3466323ae23"></a>

## discovery_k8s.access_info — discovery_k8s.access_info / f5d665471fe8 / 2

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-89b3e09314ef5fec837c9034f68d6bb66048b1e8e127c21bebfbd751a57f9668)
- [discovery_k8s](data-sources--discovery--reference--group-001.md#canonical-3c1f55627ad4d68af88df52467484fc60a52655cab4af4210a019b3b79d1b17a)
- discovery_k8s.access_info

<a id="canonical-a4055a0f637a30c771288b2567a4a6e98bba2e77b5f8828f38b2aa02a33ee068"></a>

Type: `"single"`. Computed.

Configuration parameter for access info.

Upstream description:

K8s API server access.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-config_type": "[\"connection_info\",\"kubeconfig_url\"]",
  "x-ves-oneof-field-k8s_pod_network_choice": "[\"isolated\",\"reachable\"]"
}
```

<a id="canonical-4300d4c516f22f151f7d2e2761c1b42796e946a8e11f85f122e9e7ed340dfdda"></a>

## Direct properties — discovery_k8s.access_info / f5d665471fe8 / 3

- [connection_info](data-sources--discovery--reference--group-001.md#canonical-cf528bf67d19dd94711429a6033be208abc81d71865842940b57917cf5542565): complete subsection reference.

- [isolated](data-sources--discovery--reference--group-001.md#canonical-00a7c186b22920d048334b9409c9d6d82bbeb211d1f2f837ef0309a790b36905): complete subsection reference.

- [kubeconfig_url](data-sources--discovery--reference--group-001.md#canonical-4f2861d3c2ae251a0b6d5e0cf356ad8be678bf17b4be719b879bb71cd04aba75): complete subsection reference.

- [reachable](data-sources--discovery--reference--group-001.md#canonical-761ed7238d5983b8566841a2653c299eb34b551321dc9189008d37ef61b3c8f8): complete subsection reference.

<a id="canonical-56f71944417cd0157bed4a451d6242eb043f7c8f136aa51f0870287797a60b79"></a>

## Next pages — discovery_k8s.access_info / f5d665471fe8 / 4

- [discovery_k8s.access_info.connection_info](data-sources--discovery--reference--group-001.md#canonical-cf528bf67d19dd94711429a6033be208abc81d71865842940b57917cf5542565)
- [discovery_k8s.access_info.isolated](data-sources--discovery--reference--group-001.md#canonical-00a7c186b22920d048334b9409c9d6d82bbeb211d1f2f837ef0309a790b36905)
- [discovery_k8s.access_info.kubeconfig_url](data-sources--discovery--reference--group-001.md#canonical-4f2861d3c2ae251a0b6d5e0cf356ad8be678bf17b4be719b879bb71cd04aba75)
- [discovery_k8s.access_info.reachable](data-sources--discovery--reference--group-001.md#canonical-761ed7238d5983b8566841a2653c299eb34b551321dc9189008d37ef61b3c8f8)
- [discovery_k8s](data-sources--discovery--reference--group-001.md#canonical-3c1f55627ad4d68af88df52467484fc60a52655cab4af4210a019b3b79d1b17a)
- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)

<a id="canonical-cf528bf67d19dd94711429a6033be208abc81d71865842940b57917cf5542565"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-65fa5b5928d589565e2babae598b2bcefec33accbac2443268737b94fef1c08d"></a>

## discovery_k8s.access_info.connection_info — discovery_k8s.access_info.connection_info / 25567f0163ab / 2

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-89b3e09314ef5fec837c9034f68d6bb66048b1e8e127c21bebfbd751a57f9668)
- [discovery_k8s](data-sources--discovery--reference--group-001.md#canonical-3c1f55627ad4d68af88df52467484fc60a52655cab4af4210a019b3b79d1b17a)
- [discovery_k8s.access_info](data-sources--discovery--reference--group-001.md#canonical-1588a26e7c162a7f538ee12ef2d0df7a8709f0dcba9a5d7247ee91954a5da1ef)
- discovery_k8s.access_info.connection_info

<a id="canonical-b5dbff5ac61711d4898511ea2e593dff82d07c122bebbf5e957e313a36ff4713"></a>

Type: `"single"`. Computed.

Configuration details to access discovery service REST API.

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

<a id="canonical-87380564f399e1d461262b8081a7762dafaa5f26c0d3e7e9221ad6dfc3d3ce8d"></a>

## Direct properties — discovery_k8s.access_info.connection_info / 25567f0163ab / 3

<a id="canonical-3f2f36e58cd0e02fb0b87bf6266f75c4bda2877c314b202f51daef22c210d3e9"></a>

<a id="canonical-23649f8759bf0193d79fc6f6f0c3836df99b175922571410f8ea33d6f854fe7d"></a>

## api_server property — discovery_k8s.access_info.connection_info / 25567f0163ab / 4

Type: `"string"`. Computed.

API server must be a fully qualified domain string and port specified as host:port pair.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  }
}
```

- [tls_info](data-sources--discovery--reference--group-001.md#canonical-e758ff6d020bbc37d0b203b0f263ddb3dd82265ef99f1aadd9f27c9b5d196b5e): complete subsection reference.

<a id="canonical-c1ea56d37d84558c6d1344dd761090e69c944d42e0f60c9445b16f03cd590f8d"></a>

## Next pages — discovery_k8s.access_info.connection_info / 25567f0163ab / 5

- [discovery_k8s.access_info.connection_info.tls_info](data-sources--discovery--reference--group-001.md#canonical-e758ff6d020bbc37d0b203b0f263ddb3dd82265ef99f1aadd9f27c9b5d196b5e)
- [discovery_k8s.access_info](data-sources--discovery--reference--group-001.md#canonical-1588a26e7c162a7f538ee12ef2d0df7a8709f0dcba9a5d7247ee91954a5da1ef)
- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)

<a id="canonical-e758ff6d020bbc37d0b203b0f263ddb3dd82265ef99f1aadd9f27c9b5d196b5e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e71efa8ee1143c81aff20f2a1871cfa5e20b8616cf025628432838137985ccaf"></a>

## discovery_k8s.access_info.connection_info.tls_info — discovery_k8s.access_info.connection_info.tls_info / 0b3a4f1ca472 / 2

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-89b3e09314ef5fec837c9034f68d6bb66048b1e8e127c21bebfbd751a57f9668)
- [discovery_k8s](data-sources--discovery--reference--group-001.md#canonical-3c1f55627ad4d68af88df52467484fc60a52655cab4af4210a019b3b79d1b17a)
- [discovery_k8s.access_info](data-sources--discovery--reference--group-001.md#canonical-1588a26e7c162a7f538ee12ef2d0df7a8709f0dcba9a5d7247ee91954a5da1ef)
- [discovery_k8s.access_info.connection_info](data-sources--discovery--reference--group-001.md#canonical-cf528bf67d19dd94711429a6033be208abc81d71865842940b57917cf5542565)
- discovery_k8s.access_info.connection_info.tls_info

<a id="canonical-5ad99e2afab1f9fdbbc931621c0bd2a941902586794c58c6c48dbd47e1fc86c4"></a>

Type: `"single"`. Computed.

TLS config for client of discovery service.

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

<a id="canonical-db7e941380a565f0658f4851e79217d71912e4e043b1913c57e56fd88939077a"></a>

## Direct properties — discovery_k8s.access_info.connection_info.tls_info / 0b3a4f1ca472 / 3

<a id="canonical-3532e4310911a6e3960bbd948beb13eb4fb5755b1c1e81599e5a78d5241acc0c"></a>

<a id="canonical-09e59f4ef66d4d9cce2f25f0526d0e27ecdcd9d83c1d1be5e0ea2ff1d23e5e21"></a>

## certificate property — discovery_k8s.access_info.connection_info.tls_info / 0b3a4f1ca472 / 4

Type: `"string"`. Computed.

Client certificate is PEM-encoded certificate or certificate-chain.

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
    "formatDescription": "PEM-encoded X.509 certificate, max 5MB",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 100,
    "pattern": "^-----BEGIN CERTIFICATE-----\\n.*\\n-----END CERTIFICATE-----$",
    "validation": {
      "standard": "PEM"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [key_url](data-sources--discovery--reference--group-001.md#canonical-05e0f1390172ef927178168b0a96f75c1d8804f04e79c8bd42532aabf722781e): complete subsection reference.

<a id="canonical-2a42410f3cdc2755a94ed256fedfa0d3e735189f0738ec3a72e6c68b98637a4f"></a>

<a id="canonical-8f7f48fb4bd86b917d8d2ab7626cb383190d2bed9655df499c28e5ac7f74500d"></a>

## server_name property — discovery_k8s.access_info.connection_info.tls_info / 0b3a4f1ca472 / 5

Type: `"string"`. Computed.

ServerName is passed to the server for SNI and is used in the client to check server certificates
against. If ServerName is empty, the hostname used to contact the server is used.

Upstream description:

ServerName is passed to the server for SNI and is used in the client to check server certificates
against. If ServerName is empty, the hostname used to contact the server is used.

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

<a id="canonical-469357d116017e15949f62f2f2117ccb2ef0069fc782356c578573c2e05af635"></a>

<a id="canonical-efdd4978bc6ff9f984f8062186a96c814c0d5e83a8c6abb801af94e4ebffe539"></a>

## trusted_ca_url property — discovery_k8s.access_info.connection_info.tls_info / 0b3a4f1ca472 / 6

Type: `"string"`. Computed.

The URL or value for trusted Server CA certificate or certificate chain Certificates in PEM format
including the PEM headers.

Upstream description:

The URL or value for trusted Server CA certificate or certificate chain Certificates in PEM format
including the PEM headers.

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
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-f20b4e5e4c921830fde2de980b3323c36ce1bf4a8ae5dafcdbf41ea56d3a9a3e"></a>

## Next pages — discovery_k8s.access_info.connection_info.tls_info / 0b3a4f1ca472 / 7

- [discovery_k8s.access_info.connection_info.tls_info.key_url](data-sources--discovery--reference--group-001.md#canonical-05e0f1390172ef927178168b0a96f75c1d8804f04e79c8bd42532aabf722781e)
- [discovery_k8s.access_info.connection_info](data-sources--discovery--reference--group-001.md#canonical-cf528bf67d19dd94711429a6033be208abc81d71865842940b57917cf5542565)
- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)

<a id="canonical-05e0f1390172ef927178168b0a96f75c1d8804f04e79c8bd42532aabf722781e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5454028992720a2785840973b924ec39c07704a0fe82c92aabe7cf17ed4df0f8"></a>

## discovery_k8s.access_info.connection_info.tls_info.key_url — discovery_k8s.access_info.connection_info.tls_info.key_url / 9c9564221b24 / 2

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-89b3e09314ef5fec837c9034f68d6bb66048b1e8e127c21bebfbd751a57f9668)
- [discovery_k8s](data-sources--discovery--reference--group-001.md#canonical-3c1f55627ad4d68af88df52467484fc60a52655cab4af4210a019b3b79d1b17a)
- [discovery_k8s.access_info](data-sources--discovery--reference--group-001.md#canonical-1588a26e7c162a7f538ee12ef2d0df7a8709f0dcba9a5d7247ee91954a5da1ef)
- [discovery_k8s.access_info.connection_info](data-sources--discovery--reference--group-001.md#canonical-cf528bf67d19dd94711429a6033be208abc81d71865842940b57917cf5542565)
- [discovery_k8s.access_info.connection_info.tls_info](data-sources--discovery--reference--group-001.md#canonical-e758ff6d020bbc37d0b203b0f263ddb3dd82265ef99f1aadd9f27c9b5d196b5e)
- discovery_k8s.access_info.connection_info.tls_info.key_url

<a id="canonical-d9964ed0c6412650ad5f61e6b1f56afdb51ddefabcdcc4623a74886c00652aeb"></a>

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

<a id="canonical-109101a60b9a248804a280c5563f8d959784e0c5514c8004344adaa6f6ac826e"></a>

## Direct properties — discovery_k8s.access_info.connection_info.tls_info.key_url / 9c9564221b24 / 3

- [blindfold_secret_info](data-sources--discovery--reference--group-001.md#canonical-9c64a101a8fc9da651ff3ff404ffacb4757428d2df5f34b7015a7ac23666e016): complete subsection reference.

- [clear_secret_info](data-sources--discovery--reference--group-001.md#canonical-b13aa4a3e42faed234254beb190fd2a452ca02cab11df36fcc4c6f1967deae37): complete subsection reference.

<a id="canonical-c186bc3bf8b885e6d6fba7551f6cca0c5df7955436505131be735e78831ffcd9"></a>

## Next pages — discovery_k8s.access_info.connection_info.tls_info.key_url / 9c9564221b24 / 4

- [discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info](data-sources--discovery--reference--group-001.md#canonical-9c64a101a8fc9da651ff3ff404ffacb4757428d2df5f34b7015a7ac23666e016)
- [discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info](data-sources--discovery--reference--group-001.md#canonical-b13aa4a3e42faed234254beb190fd2a452ca02cab11df36fcc4c6f1967deae37)
- [discovery_k8s.access_info.connection_info.tls_info](data-sources--discovery--reference--group-001.md#canonical-e758ff6d020bbc37d0b203b0f263ddb3dd82265ef99f1aadd9f27c9b5d196b5e)
- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)

<a id="canonical-9c64a101a8fc9da651ff3ff404ffacb4757428d2df5f34b7015a7ac23666e016"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2f5df7720a3723e525c8b74ca1bf9a518c76fcb41c56d384abe5a25db1195770"></a>

## discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info — discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info / 8e3043aea739 / 2

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-89b3e09314ef5fec837c9034f68d6bb66048b1e8e127c21bebfbd751a57f9668)
- [discovery_k8s](data-sources--discovery--reference--group-001.md#canonical-3c1f55627ad4d68af88df52467484fc60a52655cab4af4210a019b3b79d1b17a)
- [discovery_k8s.access_info](data-sources--discovery--reference--group-001.md#canonical-1588a26e7c162a7f538ee12ef2d0df7a8709f0dcba9a5d7247ee91954a5da1ef)
- [discovery_k8s.access_info.connection_info](data-sources--discovery--reference--group-001.md#canonical-cf528bf67d19dd94711429a6033be208abc81d71865842940b57917cf5542565)
- [discovery_k8s.access_info.connection_info.tls_info](data-sources--discovery--reference--group-001.md#canonical-e758ff6d020bbc37d0b203b0f263ddb3dd82265ef99f1aadd9f27c9b5d196b5e)
- [discovery_k8s.access_info.connection_info.tls_info.key_url](data-sources--discovery--reference--group-001.md#canonical-05e0f1390172ef927178168b0a96f75c1d8804f04e79c8bd42532aabf722781e)
- discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info

<a id="canonical-5fd69de3b521d70c693a9636a5c973b6775577aeb36078851f5ba4b1ecd6b870"></a>

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

<a id="canonical-7d4353f7eb774b2c83fae606de4c823ea65be125de3dd875b4e2124ed05028dc"></a>

## Direct properties — discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info / 8e3043aea739 / 3

<a id="canonical-6d2ea5603638c17505e2d44858e5ae74c4815b569967e6b2364ad5b57ac2cb4d"></a>

<a id="canonical-c309360e290a31670808c810f7bdafb6fb094501207cc4a139de9ca306265904"></a>

## decryption_provider property — discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info / 8e3043aea739 / 4

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

<a id="canonical-8b23f448cc427d79ab6bfecd7b272cb273d43e36603c5280f59aca544b372241"></a>

<a id="canonical-da55235674b82b5da6726f53b75918a24295e4c9c91bcfbfe2d2c33596db406d"></a>

## location property — discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info / 8e3043aea739 / 5

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

<a id="canonical-eceb2d200373821c13d8aa9a04f173d48775f7a3a4cb065ea4a6cedf0dbed251"></a>

<a id="canonical-b872efafe4e275f37d51a1d93c2d29698e573892cdc239ff29750c9509aa568d"></a>

## store_provider property — discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info / 8e3043aea739 / 6

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

<a id="canonical-d113c033094051a2b209819dc9ac60bdc479c4778e2ddedd55bacf9c9107caad"></a>

## Next pages — discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info / 8e3043aea739 / 7

- [discovery_k8s.access_info.connection_info.tls_info.key_url](data-sources--discovery--reference--group-001.md#canonical-05e0f1390172ef927178168b0a96f75c1d8804f04e79c8bd42532aabf722781e)
- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)

<a id="canonical-b13aa4a3e42faed234254beb190fd2a452ca02cab11df36fcc4c6f1967deae37"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5301b145c7d7820c506263c4f403233ec33d219797dceae6b1172b0b53233537"></a>

## discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info — discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info / 5b5229c9049d / 2

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-89b3e09314ef5fec837c9034f68d6bb66048b1e8e127c21bebfbd751a57f9668)
- [discovery_k8s](data-sources--discovery--reference--group-001.md#canonical-3c1f55627ad4d68af88df52467484fc60a52655cab4af4210a019b3b79d1b17a)
- [discovery_k8s.access_info](data-sources--discovery--reference--group-001.md#canonical-1588a26e7c162a7f538ee12ef2d0df7a8709f0dcba9a5d7247ee91954a5da1ef)
- [discovery_k8s.access_info.connection_info](data-sources--discovery--reference--group-001.md#canonical-cf528bf67d19dd94711429a6033be208abc81d71865842940b57917cf5542565)
- [discovery_k8s.access_info.connection_info.tls_info](data-sources--discovery--reference--group-001.md#canonical-e758ff6d020bbc37d0b203b0f263ddb3dd82265ef99f1aadd9f27c9b5d196b5e)
- [discovery_k8s.access_info.connection_info.tls_info.key_url](data-sources--discovery--reference--group-001.md#canonical-05e0f1390172ef927178168b0a96f75c1d8804f04e79c8bd42532aabf722781e)
- discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info

<a id="canonical-15a4d81b2a04b69b6a9084d5fb061e8c8dd2fd48bb94bf5389af23f564af8c5a"></a>

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

<a id="canonical-88a918b0f2487994f0016937e8972c321f4cb3d8ede68a5621696b22c7737535"></a>

## Direct properties — discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info / 5b5229c9049d / 3

<a id="canonical-3c0adb32ac3dea0d9442668251d5702238b008a32b64cf495756e90153080296"></a>

<a id="canonical-6d53f217f489f2ab6b5e9121dc1978b68671ddf08d7b373d76d2ba897db23e38"></a>

## provider_ref property — discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info / 5b5229c9049d / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-c6ea6980ff46ef77da53a441bae1c76e61caea183744da08a320c6ecf6cb2010"></a>

<a id="canonical-d9e52242b4212dd0d782cf053a3e940becd3a0bdbb087e292d9a2c067b39e2e0"></a>

## url property — discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info / 5b5229c9049d / 5

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

<a id="canonical-d0cd5ff73a7d36b479dad063497d87e13bd44661009967991ca18d34e6c3ea6e"></a>

## Next pages — discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info / 5b5229c9049d / 6

- [discovery_k8s.access_info.connection_info.tls_info.key_url](data-sources--discovery--reference--group-001.md#canonical-05e0f1390172ef927178168b0a96f75c1d8804f04e79c8bd42532aabf722781e)
- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)

<a id="canonical-00a7c186b22920d048334b9409c9d6d82bbeb211d1f2f837ef0309a790b36905"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0dfa886943464449add9b611ca8f49e6de9cf588edd828a92952918544214c64"></a>

## discovery_k8s.access_info.isolated — discovery_k8s.access_info.isolated / 99b8968f5df5 / 2

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-89b3e09314ef5fec837c9034f68d6bb66048b1e8e127c21bebfbd751a57f9668)
- [discovery_k8s](data-sources--discovery--reference--group-001.md#canonical-3c1f55627ad4d68af88df52467484fc60a52655cab4af4210a019b3b79d1b17a)
- [discovery_k8s.access_info](data-sources--discovery--reference--group-001.md#canonical-1588a26e7c162a7f538ee12ef2d0df7a8709f0dcba9a5d7247ee91954a5da1ef)
- discovery_k8s.access_info.isolated

<a id="canonical-e51e81e4bc801e7a7ca67a4069b1f8f551138c9a95a8f52d02716eaad1ca9555"></a>

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

<a id="canonical-b1a400f79d6d35f66963e777407f1903e7317f5af9fe87e3b38364a93518fcc9"></a>

## Direct properties — discovery_k8s.access_info.isolated / 99b8968f5df5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-03b139c457b68d821e1c101e52868684978d75229702a2c73b6aa9009c3c1fb9"></a>

## Next pages — discovery_k8s.access_info.isolated / 99b8968f5df5 / 4

- [discovery_k8s.access_info](data-sources--discovery--reference--group-001.md#canonical-1588a26e7c162a7f538ee12ef2d0df7a8709f0dcba9a5d7247ee91954a5da1ef)
- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)

<a id="canonical-4f2861d3c2ae251a0b6d5e0cf356ad8be678bf17b4be719b879bb71cd04aba75"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-93fa3ddf972933f9b2a81cc8a860fbae6a9ecc595ed8a821ca9d56433b06b9c6"></a>

## discovery_k8s.access_info.kubeconfig_url — discovery_k8s.access_info.kubeconfig_url / 570f989935be / 2

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-89b3e09314ef5fec837c9034f68d6bb66048b1e8e127c21bebfbd751a57f9668)
- [discovery_k8s](data-sources--discovery--reference--group-001.md#canonical-3c1f55627ad4d68af88df52467484fc60a52655cab4af4210a019b3b79d1b17a)
- [discovery_k8s.access_info](data-sources--discovery--reference--group-001.md#canonical-1588a26e7c162a7f538ee12ef2d0df7a8709f0dcba9a5d7247ee91954a5da1ef)
- discovery_k8s.access_info.kubeconfig_url

<a id="canonical-64ff47ccc9b1bb26f85aa9a4b0adbc1b784582233b48fca4afc67663de2c8a7b"></a>

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

<a id="canonical-383c8cd407e28d25466198c62e22bfa5ac8e6cdef04284c6b2185d2af87de3ad"></a>

## Direct properties — discovery_k8s.access_info.kubeconfig_url / 570f989935be / 3

- [blindfold_secret_info](data-sources--discovery--reference--group-001.md#canonical-e6974ad7a4c955d461dbb8df0bf71cde64c62e0d50ac63c0a048aacf764fbc71): complete subsection reference.

- [clear_secret_info](data-sources--discovery--reference--group-001.md#canonical-e23289a379404f6f5632c114177e0cb1cb0ab67141f297ab257ea7a5a40901da): complete subsection reference.

<a id="canonical-02ad38d2dd652f3ea637aff367d2d8f1179fb0b0a6aa3f25427a1ae216206cfe"></a>

## Next pages — discovery_k8s.access_info.kubeconfig_url / 570f989935be / 4

- [discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info](data-sources--discovery--reference--group-001.md#canonical-e6974ad7a4c955d461dbb8df0bf71cde64c62e0d50ac63c0a048aacf764fbc71)
- [discovery_k8s.access_info.kubeconfig_url.clear_secret_info](data-sources--discovery--reference--group-001.md#canonical-e23289a379404f6f5632c114177e0cb1cb0ab67141f297ab257ea7a5a40901da)
- [discovery_k8s.access_info](data-sources--discovery--reference--group-001.md#canonical-1588a26e7c162a7f538ee12ef2d0df7a8709f0dcba9a5d7247ee91954a5da1ef)
- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)

<a id="canonical-e6974ad7a4c955d461dbb8df0bf71cde64c62e0d50ac63c0a048aacf764fbc71"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a73b304e9f4c91f1610581ad182d8b17095423a4ddd827bf81c6925f5feda54c"></a>

## discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info — discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info / c6ca8e8489ce / 2

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-89b3e09314ef5fec837c9034f68d6bb66048b1e8e127c21bebfbd751a57f9668)
- [discovery_k8s](data-sources--discovery--reference--group-001.md#canonical-3c1f55627ad4d68af88df52467484fc60a52655cab4af4210a019b3b79d1b17a)
- [discovery_k8s.access_info](data-sources--discovery--reference--group-001.md#canonical-1588a26e7c162a7f538ee12ef2d0df7a8709f0dcba9a5d7247ee91954a5da1ef)
- [discovery_k8s.access_info.kubeconfig_url](data-sources--discovery--reference--group-001.md#canonical-4f2861d3c2ae251a0b6d5e0cf356ad8be678bf17b4be719b879bb71cd04aba75)
- discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info

<a id="canonical-93e88151f78c8098e1a58c23da5b2c2ca0943ea22f316117614d433522d1b3e2"></a>

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

<a id="canonical-8088b77ddadd5fabd64afd1d11e2351f81a556fb12bb9c9afea5748f0f813f21"></a>

## Direct properties — discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info / c6ca8e8489ce / 3

<a id="canonical-a3a462a66273d34009a90fbd83b4fb24e79bce861b22d45cc7a440d24f712669"></a>

<a id="canonical-57783ec20932fed7dec7650977411dd677109517bc07289e61e046f65d02948e"></a>

## decryption_provider property — discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info / c6ca8e8489ce / 4

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

<a id="canonical-046f9d713bb4b98d7289dd347ecc06874f5434816a864ed684c4f0606312e5d9"></a>

<a id="canonical-82d14a24c5991d899020e7ab5e05a50e980ba6bb64cae8ca672be4cd33d02359"></a>

## location property — discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info / c6ca8e8489ce / 5

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

<a id="canonical-166c53879e9d0726758debf8f784c297a853ba5b528d8c64156758cb646d16e5"></a>

<a id="canonical-585f31d98eacda92579a40e790c7bd0f049027497639635cf6f7b4e1e05ac18a"></a>

## store_provider property — discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info / c6ca8e8489ce / 6

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

<a id="canonical-648fd23fbbe2e46f8c6a186a77a052993f8eb24fe9ff8dea21ea5428c8e91de2"></a>

## Next pages — discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info / c6ca8e8489ce / 7

- [discovery_k8s.access_info.kubeconfig_url](data-sources--discovery--reference--group-001.md#canonical-4f2861d3c2ae251a0b6d5e0cf356ad8be678bf17b4be719b879bb71cd04aba75)
- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)

<a id="canonical-e23289a379404f6f5632c114177e0cb1cb0ab67141f297ab257ea7a5a40901da"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0e6cdcb41fa0efe686985e87463fd26cc5e1535eab6f4d6f6cd2cd6d648b7d3e"></a>

## discovery_k8s.access_info.kubeconfig_url.clear_secret_info — discovery_k8s.access_info.kubeconfig_url.clear_secret_info / 945211a32a77 / 2

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-89b3e09314ef5fec837c9034f68d6bb66048b1e8e127c21bebfbd751a57f9668)
- [discovery_k8s](data-sources--discovery--reference--group-001.md#canonical-3c1f55627ad4d68af88df52467484fc60a52655cab4af4210a019b3b79d1b17a)
- [discovery_k8s.access_info](data-sources--discovery--reference--group-001.md#canonical-1588a26e7c162a7f538ee12ef2d0df7a8709f0dcba9a5d7247ee91954a5da1ef)
- [discovery_k8s.access_info.kubeconfig_url](data-sources--discovery--reference--group-001.md#canonical-4f2861d3c2ae251a0b6d5e0cf356ad8be678bf17b4be719b879bb71cd04aba75)
- discovery_k8s.access_info.kubeconfig_url.clear_secret_info

<a id="canonical-2dfc28e191786c3d84d4b77fd827baa59345d88fea197c8a46eead0d607e0767"></a>

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

<a id="canonical-786a16c72e2af801fd8a63d8a32952b837b03dadefb961f55d7da3199845c9a6"></a>

## Direct properties — discovery_k8s.access_info.kubeconfig_url.clear_secret_info / 945211a32a77 / 3

<a id="canonical-0e06cb3bd6e48d3a4a134089fb64a049e5e50c04a4844d4cc69c37aaf0c18899"></a>

<a id="canonical-fff3d3ac3bfc25af2d47655521cba1c1018b943f592ac5bb9c6be4e11fdc8782"></a>

## provider_ref property — discovery_k8s.access_info.kubeconfig_url.clear_secret_info / 945211a32a77 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-70b7070aef32a158bde4795a4ae8b08c1d23c3154a7b4a95b50ef05e304ce83b"></a>

<a id="canonical-301a59cfe3fad3997f006e76fc531885aa42863b596bb8c93458f26949941808"></a>

## url property — discovery_k8s.access_info.kubeconfig_url.clear_secret_info / 945211a32a77 / 5

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

<a id="canonical-589adc15df094bf7f1447b60988ea77b18ce4eb351a8ee72a3c3f3c9c3fa9e1b"></a>

## Next pages — discovery_k8s.access_info.kubeconfig_url.clear_secret_info / 945211a32a77 / 6

- [discovery_k8s.access_info.kubeconfig_url](data-sources--discovery--reference--group-001.md#canonical-4f2861d3c2ae251a0b6d5e0cf356ad8be678bf17b4be719b879bb71cd04aba75)
- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)

<a id="canonical-761ed7238d5983b8566841a2653c299eb34b551321dc9189008d37ef61b3c8f8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b121d44d1805868b1a246952680d60312701b8c8cc180f62a14156052aa4b0a7"></a>

## discovery_k8s.access_info.reachable — discovery_k8s.access_info.reachable / 65cc97768f0d / 2

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-89b3e09314ef5fec837c9034f68d6bb66048b1e8e127c21bebfbd751a57f9668)
- [discovery_k8s](data-sources--discovery--reference--group-001.md#canonical-3c1f55627ad4d68af88df52467484fc60a52655cab4af4210a019b3b79d1b17a)
- [discovery_k8s.access_info](data-sources--discovery--reference--group-001.md#canonical-1588a26e7c162a7f538ee12ef2d0df7a8709f0dcba9a5d7247ee91954a5da1ef)
- discovery_k8s.access_info.reachable

<a id="canonical-dcafd875b728ebc359991f1a0f51906c9a7fe55212ef9d13047ebf30994dbd49"></a>

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

<a id="canonical-05592a03cbffea37d1d0d5aa30114511eeb5ec96d7f2467666d8ab9b1fe846da"></a>

## Direct properties — discovery_k8s.access_info.reachable / 65cc97768f0d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1f7ad30fd77bcb88d12e54965f666db217655094fed94a36652bd436d9d74338"></a>

## Next pages — discovery_k8s.access_info.reachable / 65cc97768f0d / 4

- [discovery_k8s.access_info](data-sources--discovery--reference--group-001.md#canonical-1588a26e7c162a7f538ee12ef2d0df7a8709f0dcba9a5d7247ee91954a5da1ef)
- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)

<a id="canonical-15326be21e5d6f24d2ccb480b59d2172ad8551beaf05b81dc4d28fd4b49c734c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-76b55e4741b5e713f6bfbde0e1e32c747df8a96ebb07455ebcbccb2e6d73a88d"></a>

## discovery_k8s.default_all — discovery_k8s.default_all / e1564d9c4b38 / 2

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-89b3e09314ef5fec837c9034f68d6bb66048b1e8e127c21bebfbd751a57f9668)
- [discovery_k8s](data-sources--discovery--reference--group-001.md#canonical-3c1f55627ad4d68af88df52467484fc60a52655cab4af4210a019b3b79d1b17a)
- discovery_k8s.default_all

<a id="canonical-e4e67154b5c84e4c9ad2ae490f2dcb89c18cc205e36ab1d3af9aa35acb3dba4c"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default all.

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

<a id="canonical-fb40bb3b91e8ed000e1286ab5774c6351258bf5261ae1842f5957bb878444c9a"></a>

## Direct properties — discovery_k8s.default_all / e1564d9c4b38 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-51919261d96b78f225a3ee0d53d7af9da8a8e743003850fca220beffe906231d"></a>

## Next pages — discovery_k8s.default_all / e1564d9c4b38 / 4

- [discovery_k8s](data-sources--discovery--reference--group-001.md#canonical-3c1f55627ad4d68af88df52467484fc60a52655cab4af4210a019b3b79d1b17a)
- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)

<a id="canonical-f31f19eafefa2e69b558d06d3bfbc3deac8da1185d9da7f44f2ee640b8a86809"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-74c5e173026ff408b76411a63572f3fe30666880a455ce26d53c141b717c8369"></a>

## discovery_k8s.namespace_mapping — discovery_k8s.namespace_mapping / 691101aa083d / 2

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-89b3e09314ef5fec837c9034f68d6bb66048b1e8e127c21bebfbd751a57f9668)
- [discovery_k8s](data-sources--discovery--reference--group-001.md#canonical-3c1f55627ad4d68af88df52467484fc60a52655cab4af4210a019b3b79d1b17a)
- discovery_k8s.namespace_mapping

<a id="canonical-81ea92d29cb22d48ef84853ff533a4a3df57ba01002c2257d88e5e86ec22d586"></a>

Type: `"single"`. Computed.

Select the mapping between K8s namespaces from which services will be discovered and App Namespace
to which the discovered services will be shared.

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

<a id="canonical-d50df1578c21f0761489a231b6a563880abce89212d24bd1e0bca4e00878b211"></a>

## Direct properties — discovery_k8s.namespace_mapping / 691101aa083d / 3

- [items](data-sources--discovery--reference--group-001.md#canonical-ac6872c0e0b51b1e1a6ee23da0f4f3da63315e6d1e1974fa0417649896b53538): complete subsection reference.

<a id="canonical-d2bf0196193cd2c5de0b612befc406a4f734f2cf0862085581044b9df744b83a"></a>

## Next pages — discovery_k8s.namespace_mapping / 691101aa083d / 4

- [discovery_k8s.namespace_mapping.items](data-sources--discovery--reference--group-001.md#canonical-ac6872c0e0b51b1e1a6ee23da0f4f3da63315e6d1e1974fa0417649896b53538)
- [discovery_k8s](data-sources--discovery--reference--group-001.md#canonical-3c1f55627ad4d68af88df52467484fc60a52655cab4af4210a019b3b79d1b17a)
- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)

<a id="canonical-ac6872c0e0b51b1e1a6ee23da0f4f3da63315e6d1e1974fa0417649896b53538"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2eea14d9f4f0df98eb7c6dcd0f3a97aaf1ba41545867e82c1e4b2c2c7375cc61"></a>

## discovery_k8s.namespace_mapping.items — discovery_k8s.namespace_mapping.items / bce8e641ce05 / 2

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-89b3e09314ef5fec837c9034f68d6bb66048b1e8e127c21bebfbd751a57f9668)
- [discovery_k8s](data-sources--discovery--reference--group-001.md#canonical-3c1f55627ad4d68af88df52467484fc60a52655cab4af4210a019b3b79d1b17a)
- [discovery_k8s.namespace_mapping](data-sources--discovery--reference--group-001.md#canonical-f31f19eafefa2e69b558d06d3bfbc3deac8da1185d9da7f44f2ee640b8a86809)
- discovery_k8s.namespace_mapping.items

<a id="canonical-9e9598f66e03e6fab91440d72b944a505ee80ea4cbac945b0911bfcf227559d0"></a>

Type: `"list"`. Computed.

Map K8s namespace(s) to App Namespaces. In Shared Configuration, Discovered Services can only be
mapped to a single App Namespace, which is determined by the first matched regex.

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

<a id="canonical-77b45d920c76692414a53fba1d788cc47bc4cff518239c21d3b36787033c2f0a"></a>

## Direct properties — discovery_k8s.namespace_mapping.items / bce8e641ce05 / 3

<a id="canonical-38381323b3bb0c89e48c6546b3b0193394b2c9d2ff60ad9f86d93422eba881ae"></a>

<a id="canonical-4b38728074d4c497175a43e75ef91703a174cd25e6a5c6193fc5636f4112b444"></a>

## namespace property — discovery_k8s.namespace_mapping.items / bce8e641ce05 / 4

Type: `"string"`. Computed.

F5XC Application Namespaces. Select a namespace.

Upstream description:

Select a namespace.

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

<a id="canonical-2aea89c21e70ba10e28bb8dde08414a1721a0f419ac70fa8dcb2d8778d39c6b5"></a>

<a id="canonical-fdc0dd195badd6b0b901937ac89444055c866a1210c21dd9e1767b990adec8ec"></a>

## namespace_regex property — discovery_k8s.namespace_mapping.items / bce8e641ce05 / 5

Type: `"string"`. Computed.

The regex here will be used to match K8s namespace(s).

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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-538693fb3ad04bda57c67eae5f4eda3bd06c0e9390c36fa0ff71809415538240"></a>

## Next pages — discovery_k8s.namespace_mapping.items / bce8e641ce05 / 6

- [discovery_k8s.namespace_mapping](data-sources--discovery--reference--group-001.md#canonical-f31f19eafefa2e69b558d06d3bfbc3deac8da1185d9da7f44f2ee640b8a86809)
- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)

<a id="canonical-8074630174ac5afaea680e91dcf3e042a32b8c5459a26c9262928afb515aa8d6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-269cccfa7e86edce07cd67df062d51f12a844c72059dcf03bc78727a047ba4e7"></a>

## discovery_k8s.publish_info — discovery_k8s.publish_info / c6e0d56ba376 / 2

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-89b3e09314ef5fec837c9034f68d6bb66048b1e8e127c21bebfbd751a57f9668)
- [discovery_k8s](data-sources--discovery--reference--group-001.md#canonical-3c1f55627ad4d68af88df52467484fc60a52655cab4af4210a019b3b79d1b17a)
- discovery_k8s.publish_info

<a id="canonical-192a3a7fae1a6a9129f657c9b1c421ff251074118a21d10da85c94a53c2be19a"></a>

Type: `"single"`. Computed.

Configuration parameter for publish info.

Upstream description:

K8s Configuration to publish VIPs.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-publish_choice": "[\"disable\",\"dns_delegation\",\"publish\",\"publish_fqdns\"]"
}
```

<a id="canonical-d8f1be49f7536f107bab5eb9850d03f8879dcb2f55f1a88351b064ad7de806ef"></a>

## Direct properties — discovery_k8s.publish_info / c6e0d56ba376 / 3

- [disable_spec](data-sources--discovery--reference--group-001.md#canonical-3a23f524cd5fb3d5694c29d72ff1ce5d7463160a7c9dcf94b64775ecc7f95e11): complete subsection reference.

- [dns_delegation](data-sources--discovery--reference--group-001.md#canonical-09505fcacded8129fd26a52f8ec7c1bb09456dc729ec0d2704e96fe3b69996af): complete subsection reference.

- [publish](data-sources--discovery--reference--group-001.md#canonical-2df243e4bfb6ced19720de56b009ca0429243ea2877ceee2cdfe2cd4f2d8913f): complete subsection reference.

- [publish_fqdns](data-sources--discovery--reference--group-001.md#canonical-ae600f475e938be4f78f11f84ea895c8bae5a27d551ec050ac49e44242e8a346): complete subsection reference.

<a id="canonical-9c59f90dd0eef4185ed3bcd47e5dbb668cf9ac6e46680f18ce8a932da0412db4"></a>

## Next pages — discovery_k8s.publish_info / c6e0d56ba376 / 4

- [discovery_k8s.publish_info.disable_spec](data-sources--discovery--reference--group-001.md#canonical-3a23f524cd5fb3d5694c29d72ff1ce5d7463160a7c9dcf94b64775ecc7f95e11)
- [discovery_k8s.publish_info.dns_delegation](data-sources--discovery--reference--group-001.md#canonical-09505fcacded8129fd26a52f8ec7c1bb09456dc729ec0d2704e96fe3b69996af)
- [discovery_k8s.publish_info.publish](data-sources--discovery--reference--group-001.md#canonical-2df243e4bfb6ced19720de56b009ca0429243ea2877ceee2cdfe2cd4f2d8913f)
- [discovery_k8s.publish_info.publish_fqdns](data-sources--discovery--reference--group-001.md#canonical-ae600f475e938be4f78f11f84ea895c8bae5a27d551ec050ac49e44242e8a346)
- [discovery_k8s](data-sources--discovery--reference--group-001.md#canonical-3c1f55627ad4d68af88df52467484fc60a52655cab4af4210a019b3b79d1b17a)
- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)

<a id="canonical-3a23f524cd5fb3d5694c29d72ff1ce5d7463160a7c9dcf94b64775ecc7f95e11"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5b92028ccb8f59f7c9dbaec5378d4a6ff5a0dbd11eeb80f079815abfed2fcc48"></a>

## discovery_k8s.publish_info.disable_spec — discovery_k8s.publish_info.disable_spec / 9875952b27a1 / 2

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-89b3e09314ef5fec837c9034f68d6bb66048b1e8e127c21bebfbd751a57f9668)
- [discovery_k8s](data-sources--discovery--reference--group-001.md#canonical-3c1f55627ad4d68af88df52467484fc60a52655cab4af4210a019b3b79d1b17a)
- [discovery_k8s.publish_info](data-sources--discovery--reference--group-001.md#canonical-8074630174ac5afaea680e91dcf3e042a32b8c5459a26c9262928afb515aa8d6)
- discovery_k8s.publish_info.disable_spec

<a id="canonical-52b8f7175856174ef88ae639e721c45d23c135176a1e3da9de9b33eb4fed5b30"></a>

Type: `["object", {}]`. Computed.

Enable this option

<a id="canonical-176e46c86782af337e0f7c2c5ca48870fd1a7707634bc5e80154149cd944d56c"></a>

## Direct properties — discovery_k8s.publish_info.disable_spec / 9875952b27a1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e74fdebdd06e8e072ccf1b169d20fe0d2410f7b6f3b6f6537354497c677f869a"></a>

## Next pages — discovery_k8s.publish_info.disable_spec / 9875952b27a1 / 4

- [discovery_k8s.publish_info](data-sources--discovery--reference--group-001.md#canonical-8074630174ac5afaea680e91dcf3e042a32b8c5459a26c9262928afb515aa8d6)
- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)

<a id="canonical-09505fcacded8129fd26a52f8ec7c1bb09456dc729ec0d2704e96fe3b69996af"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7ac562e52942248f34d3e72699c99fefeaf8eb41289890dfb9279caa7213ca57"></a>

## discovery_k8s.publish_info.dns_delegation — discovery_k8s.publish_info.dns_delegation / b07e9171d2c0 / 2

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-89b3e09314ef5fec837c9034f68d6bb66048b1e8e127c21bebfbd751a57f9668)
- [discovery_k8s](data-sources--discovery--reference--group-001.md#canonical-3c1f55627ad4d68af88df52467484fc60a52655cab4af4210a019b3b79d1b17a)
- [discovery_k8s.publish_info](data-sources--discovery--reference--group-001.md#canonical-8074630174ac5afaea680e91dcf3e042a32b8c5459a26c9262928afb515aa8d6)
- discovery_k8s.publish_info.dns_delegation

<a id="canonical-49f48ef2bf0985d6827e277621213ea49868fe1982490c32aa193b0e738ce2ae"></a>

Type: `"single"`. Computed.

Configuration parameter for dns delegation.

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

<a id="canonical-86c048d6168124fa9063538c73189444ea4d49153c783f6ef46ed7d0b16634ee"></a>

## Direct properties — discovery_k8s.publish_info.dns_delegation / b07e9171d2c0 / 3

<a id="canonical-2d5894580f32d3d26e8c8668ebdea16ba28032768be85322014356b828253b19"></a>

<a id="canonical-36ac70c4af00ecc23dbd1cb967570213cc8df39acef77b14c7d5902e987c16fe"></a>

## dns_mode property — discovery_k8s.publish_info.dns_delegation / b07e9171d2c0 / 4

Type: `"string"`. Computed.

\[Enum: CORE\_DNS|KUBE\_DNS\] Two modes are possible CoreDNS: Whether external K8s cluster is
running core-DNS KubeDNS: External K8s cluster is running kube-DNS. Possible values are
\`CORE\_DNS\`, \`KUBE\_DNS\`. Defaults to \`CORE\_DNS\`.

Upstream description:

Two modes are possible

CoreDNS: Whether external K8s cluster is running core-DNS KubeDNS: External K8s cluster is running
kube-DNS.

Receipt-pinned upstream constraints:

```json
{
  "default": "CORE_DNS",
  "enum": [
    "CORE_DNS",
    "KUBE_DNS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1c667bf1dfe58e8b2b8a07f34335bfe70d68cbba9824669895b0d1deecd27743"></a>

<a id="canonical-c825ee1d90cd3f7dafdaed45a9424c619851eda5885d7f512daba3ad49567406"></a>

## subdomain property — discovery_k8s.publish_info.dns_delegation / b07e9171d2c0 / 5

Type: `"string"`. Computed.

The DNS subdomain for which F5XC will respond to DNS queries.

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

<a id="canonical-66ce979c70bceb2af810cc807973986f883543cbbb27fec5b81f7ecd6063e8a4"></a>

## Next pages — discovery_k8s.publish_info.dns_delegation / b07e9171d2c0 / 6

- [discovery_k8s.publish_info](data-sources--discovery--reference--group-001.md#canonical-8074630174ac5afaea680e91dcf3e042a32b8c5459a26c9262928afb515aa8d6)
- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)

<a id="canonical-2df243e4bfb6ced19720de56b009ca0429243ea2877ceee2cdfe2cd4f2d8913f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9b4bc4923df43ed9d4007c52fe07c6d2edd8624ed844e98ac47fb19ffb38e0f9"></a>

## discovery_k8s.publish_info.publish — discovery_k8s.publish_info.publish / e48600dcb1fb / 2

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-89b3e09314ef5fec837c9034f68d6bb66048b1e8e127c21bebfbd751a57f9668)
- [discovery_k8s](data-sources--discovery--reference--group-001.md#canonical-3c1f55627ad4d68af88df52467484fc60a52655cab4af4210a019b3b79d1b17a)
- [discovery_k8s.publish_info](data-sources--discovery--reference--group-001.md#canonical-8074630174ac5afaea680e91dcf3e042a32b8c5459a26c9262928afb515aa8d6)
- discovery_k8s.publish_info.publish

<a id="canonical-2272daaf79426f53c460ded46ce5f3cfdce2ed2976fa7778e6cce203e6f41ec0"></a>

Type: `"single"`. Computed.

K8SPublishType.

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

<a id="canonical-acd545b1af9dfb2e100c9bf8eef4e6af091379b9a85bc24481329b5125e88c52"></a>

## Direct properties — discovery_k8s.publish_info.publish / e48600dcb1fb / 3

<a id="canonical-ebd927b0d328a2f1883113c55afada96ca8459c43b2cfaa5c821370393e78dd7"></a>

<a id="canonical-cc99c134564d4c2135a0543265e4e71522c4c0f3ed2ff36456c0d11e77b3e1fb"></a>

## namespace property — discovery_k8s.publish_info.publish / e48600dcb1fb / 4

Type: `"string"`. Computed.

The namespace where the service/endpoints need to be created if it's not included in the domain. The
external K8s administrator needs to ensure that the namespace exists.

Upstream description:

The namespace where the service/endpoints need to be created if it's not included in the domain. The
external K8s administrator needs to ensure that the namespace exists.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-36dcfd8082d29e3a8f342ef3a6078c79c06ce36ce93cb4b3d955e6d670435395"></a>

## Next pages — discovery_k8s.publish_info.publish / e48600dcb1fb / 5

- [discovery_k8s.publish_info](data-sources--discovery--reference--group-001.md#canonical-8074630174ac5afaea680e91dcf3e042a32b8c5459a26c9262928afb515aa8d6)
- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)

<a id="canonical-ae600f475e938be4f78f11f84ea895c8bae5a27d551ec050ac49e44242e8a346"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-75a143dc90a0f32230f0680ba84fc6bc0093ae1378b5aed273bed3988eb58464"></a>

## discovery_k8s.publish_info.publish_fqdns — discovery_k8s.publish_info.publish_fqdns / 5f88b56b9784 / 2

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-89b3e09314ef5fec837c9034f68d6bb66048b1e8e127c21bebfbd751a57f9668)
- [discovery_k8s](data-sources--discovery--reference--group-001.md#canonical-3c1f55627ad4d68af88df52467484fc60a52655cab4af4210a019b3b79d1b17a)
- [discovery_k8s.publish_info](data-sources--discovery--reference--group-001.md#canonical-8074630174ac5afaea680e91dcf3e042a32b8c5459a26c9262928afb515aa8d6)
- discovery_k8s.publish_info.publish_fqdns

<a id="canonical-3c8b9c87aa16aae76cff5b00b3a4a6a17103a032e05bef85694ee90d1b16efa2"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for publish fqdns.

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

<a id="canonical-7624ac25ea5f38c94b568943fed4d73eb84f67c634ea90945c4b02feb8de8d29"></a>

## Direct properties — discovery_k8s.publish_info.publish_fqdns / 5f88b56b9784 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5827ef3abc069e039a63e86196da49c4048917461d351699c9dac7718c87b5a5"></a>

## Next pages — discovery_k8s.publish_info.publish_fqdns / 5f88b56b9784 / 4

- [discovery_k8s.publish_info](data-sources--discovery--reference--group-001.md#canonical-8074630174ac5afaea680e91dcf3e042a32b8c5459a26c9262928afb515aa8d6)
- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)

<a id="canonical-cab961f6582c0f7f0205c82d8b10a272b9e0fe4b779f443dc0de9bb85e6b059c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0d770e3604dd410857ba0c03dad0f8d696dde1eda10c44823938c9009ba52ea2"></a>

## no_cluster_id — no_cluster_id / 01f2ee8c2acc / 2

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-89b3e09314ef5fec837c9034f68d6bb66048b1e8e127c21bebfbd751a57f9668)
- no_cluster_id

<a id="canonical-bc76a7116b66e700f48d50b0d53306bb8e2a03efd8228c74ee271680471c69b2"></a>

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

<a id="canonical-1aae64aa9c92c183a0915da4ea03835950cd29ae926c3eb2813c19ac0ae4acbe"></a>

## Direct properties — no_cluster_id / 01f2ee8c2acc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0c5776d8cf6d58322ba3a95c85988c9625f592f9fd31dbe54aab997443b239b2"></a>

## Next pages — no_cluster_id / 01f2ee8c2acc / 4

- [Property reference](data-sources--discovery--reference--group-001.md#canonical-89b3e09314ef5fec837c9034f68d6bb66048b1e8e127c21bebfbd751a57f9668)
- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)

<a id="canonical-86e8aa7a75f491571c4a99862108276e04b016f3b24020940f1e1b752bd38094"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0c5932712c303cf306cc53c07186f88933629038e3821729a575f9fdd76ddc25"></a>

## where — where / f6d23f8847b5 / 2

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-89b3e09314ef5fec837c9034f68d6bb66048b1e8e127c21bebfbd751a57f9668)
- where

<a id="canonical-412f8957dcafc21637e2f3135813f3e29211c723a4b2e45916d1f8f210f80c9f"></a>

Type: `"single"`. Computed.

NetworkSiteRefSelector defines a union of reference to site or reference to virtual\_network or
reference to virtual\_site It is used to determine virtual network using following rules \* Direct
reference to virtual\_network object \* Site local network when referring to site object \* All site
local..

Upstream description:

NetworkSiteRefSelector defines a union of reference to site or reference to virtual\_network or
reference to virtual\_site It is used to determine virtual network using following rules \* Direct
reference to virtual\_network object \* Site local network when referring to site object \* All site
local networks for sites selected by referring to virtual\_site object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ref_or_selector": "[\"site\",\"virtual_network\",\"virtual_site\"]"
}
```

<a id="canonical-e282cee9792ea3986b217fb27dda0525d0019ad439457b1e87e1cb910d809314"></a>

## Direct properties — where / f6d23f8847b5 / 3

- [site](data-sources--discovery--reference--group-001.md#canonical-5d56699cb08276324472b54d8dad817bce1c7511ad8ed36b5e76dfedc368d266): complete subsection reference.

- [virtual_network](data-sources--discovery--reference--group-001.md#canonical-8042369f31b82ef4efd1e1e08aa7c5376469f4c4d96cb59b0f03957a64faccfb): complete subsection reference.

- [virtual_site](data-sources--discovery--reference--group-001.md#canonical-9fad6288637564585ba5e19b4645c139a7c02678af28eed7d5f04b6a3aa9535a): complete subsection reference.

<a id="canonical-61ab730dada1c76612fc36e191e4beda8f9e9c50bbf731158d6e4113cf7c5713"></a>

## Next pages — where / f6d23f8847b5 / 4

- [where.site](data-sources--discovery--reference--group-001.md#canonical-5d56699cb08276324472b54d8dad817bce1c7511ad8ed36b5e76dfedc368d266)
- [where.virtual_network](data-sources--discovery--reference--group-001.md#canonical-8042369f31b82ef4efd1e1e08aa7c5376469f4c4d96cb59b0f03957a64faccfb)
- [where.virtual_site](data-sources--discovery--reference--group-001.md#canonical-9fad6288637564585ba5e19b4645c139a7c02678af28eed7d5f04b6a3aa9535a)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-89b3e09314ef5fec837c9034f68d6bb66048b1e8e127c21bebfbd751a57f9668)
- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)

<a id="canonical-5d56699cb08276324472b54d8dad817bce1c7511ad8ed36b5e76dfedc368d266"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-34401fa932f045d6d143ac3e53f07dcb51fea8d0ce00fcdf4b977ebc5f9533b2"></a>

## where.site — where.site / 1495f9e0a46e / 2

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-89b3e09314ef5fec837c9034f68d6bb66048b1e8e127c21bebfbd751a57f9668)
- [where](data-sources--discovery--reference--group-001.md#canonical-86e8aa7a75f491571c4a99862108276e04b016f3b24020940f1e1b752bd38094)
- where.site

<a id="canonical-be31b01b7e0ba49133be07099144e678e12d337122c7b7b13426fd153d6ec825"></a>

Type: `"single"`. Computed.

Specifies a direct reference to a site configuration object.

Upstream description:

This specifies a direct reference to a site configuration object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-internet_vip_choice": "[\"disable_internet_vip\",\"enable_internet_vip\"]"
}
```

<a id="canonical-d3325b27b2fe7bd02c28f38cc1c4751d1238c976b4443d1e52b0ecccbec28769"></a>

## Direct properties — where.site / 1495f9e0a46e / 3

- [disable_internet_vip](data-sources--discovery--reference--group-001.md#canonical-aeb0233507696b9cb18c223a8bf77bc141a2ab4027a2f72511a308ce599ef82d): complete subsection reference.

- [enable_internet_vip](data-sources--discovery--reference--group-001.md#canonical-a8706eee0aa864b6da572208791fee81a1e9f0422a844aef2c68a5edce1b8cbe): complete subsection reference.

<a id="canonical-c505429a491c24c9d5a054ccdfeeb73b01e403c965c1bc2f7bddc69008a5e059"></a>

<a id="canonical-0efc1ff0e3f08c93b32f40df42797af4b98e89377ef27e3ddf7f0bd417c1785b"></a>

## network_type property — where.site / 1495f9e0a46e / 4

Type: `"string"`. Computed.

\[Enum:
VIRTUAL\_NETWORK\_SITE\_LOCAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE|VIRTUAL\_NETWORK\_PER\_SITE|VIRTUAL\_NETWORK\_PUBLIC|VIRTUAL\_NETWORK\_GLOBAL|VIRTUAL\_NETWORK\_SITE\_SERVICE|VIRTUAL\_NETWORK\_VER\_INTERNAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE|VIRTUAL\_NETWORK\_IP\_AUTO|VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK|VIRTUAL\_NETWORK\_SRV6\_NETWORK|VIRTUAL\_NETWORK\_IP\_FABRIC|VIRTUAL\_NETWORK\_SEGMENT|VIRTUAL\_NETWORK\_MANAGEMENT\]
Different types of virtual networks understood by the system Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside) network. This is an insecure
network and is connected to public internet via NAT Gateways/firwalls Virtual-network of this type
is local to.. Possible values are \`VIRTUAL\_NETWORK\_SITE\_LOCAL\`,
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\`, \`VIRTUAL\_NETWORK\_PER\_SITE\`,
\`VIRTUAL\_NETWORK\_PUBLIC\`, \`VIRTUAL\_NETWORK\_GLOBAL\`, \`VIRTUAL\_NETWORK\_SITE\_SERVICE\`,
\`VIRTUAL\_NETWORK\_VER\_INTERNAL\`, \`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE\`,
\`VIRTUAL\_NETWORK\_IP\_AUTO\`, \`VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK\`,
\`VIRTUAL\_NETWORK\_SRV6\_NETWORK\`, \`VIRTUAL\_NETWORK\_IP\_FABRIC\`,
\`VIRTUAL\_NETWORK\_SEGMENT\`, \`VIRTUAL\_NETWORK\_MANAGEMENT\`. Defaults to
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\`.

Upstream description:

Different types of virtual networks understood by the system

Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside)
network. This is an insecure network and is connected to public internet via NAT Gateways/firwalls
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on CE sites. This network is created automatically and present on all sites
Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE is a private network inside site. It
is a secure network and is not connected to public network. Virtual-network of this type is local to
every site. Two virtual networks of this type on different sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on CE sites. This network is created during provisioning of site User defined per-site
virtual network. Scope of this virtual network is limited to the site. This is not yet supported
Virtual-network of type VIRTUAL\_NETWORK\_PUBLIC directly connects to the public internet.
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on RE sites only It is an internally created by the system. They must not be created by
user Virtual Networks with global scope across different sites in F5XC domain. An example global
virtual-network called "AIN Network" is created for every tenant. For F5 Distributed Cloud fabric

Constraints: It is currently only supported as internally created by the system. VK8s service
network for a given tenant. Used to advertise a virtual host only to vk8s pods for that tenant
Constraints: It is an internally created by the system. Must not be created by user VER internal
network for the site. It can only be used for virtual hosts with SMA\_PROXY type proxy Constraints:
It is an internally created by the system. Must not be created by user Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE represents both VIRTUAL\_NETWORK\_SITE\_LOCAL and
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE

Constraints: This network type is only meaningful in an advertise policy When virtual-network of
type VIRTUAL\_NETWORK\_IP\_AUTO is selected for an endpoint, VER will try to determine the network
based on the provided IP address

Constraints: This network type is only meaningful in an endpoint

VoltADN Private Network is used on F5 Distributed Cloud RE(s) to connect to customer private
networks This network is created by opening a support ticket

This network is per site srv6 network VER IP Fabric network for the site. This Virtual network type
is used for exposing virtual host on IP Fabric network on the VER site or for endpoint in IP Fabric
network Constraints: It is an internally created by the system. Must not be created by user
Virtual-network of type VIRTUAL\_NETWORK\_SEGMENT for segment interface Virtual-network of type
VIRTUAL\_NETWORK\_MANAGEMENT is used for management purposes.

Receipt-pinned upstream constraints:

```json
{
  "default": "VIRTUAL_NETWORK_SITE_LOCAL",
  "enum": [
    "VIRTUAL_NETWORK_SITE_LOCAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE",
    "VIRTUAL_NETWORK_PER_SITE",
    "VIRTUAL_NETWORK_PUBLIC",
    "VIRTUAL_NETWORK_GLOBAL",
    "VIRTUAL_NETWORK_SITE_SERVICE",
    "VIRTUAL_NETWORK_VER_INTERNAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE",
    "VIRTUAL_NETWORK_IP_AUTO",
    "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK",
    "VIRTUAL_NETWORK_SRV6_NETWORK",
    "VIRTUAL_NETWORK_IP_FABRIC",
    "VIRTUAL_NETWORK_SEGMENT",
    "VIRTUAL_NETWORK_MANAGEMENT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [ref](data-sources--discovery--reference--group-001.md#canonical-1dad1ccc90a897fc903d4f8fe1740420c45839c16599c3129594d410de66c276): complete subsection reference.

<a id="canonical-799318c2b90e65cb1e2f00516f3b6c312ec09acb8cff91078fc1fe0f4e58042e"></a>

## Next pages — where.site / 1495f9e0a46e / 5

- [where.site.disable_internet_vip](data-sources--discovery--reference--group-001.md#canonical-aeb0233507696b9cb18c223a8bf77bc141a2ab4027a2f72511a308ce599ef82d)
- [where.site.enable_internet_vip](data-sources--discovery--reference--group-001.md#canonical-a8706eee0aa864b6da572208791fee81a1e9f0422a844aef2c68a5edce1b8cbe)
- [where.site.ref](data-sources--discovery--reference--group-001.md#canonical-1dad1ccc90a897fc903d4f8fe1740420c45839c16599c3129594d410de66c276)
- [where](data-sources--discovery--reference--group-001.md#canonical-86e8aa7a75f491571c4a99862108276e04b016f3b24020940f1e1b752bd38094)
- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)

<a id="canonical-aeb0233507696b9cb18c223a8bf77bc141a2ab4027a2f72511a308ce599ef82d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3a196cc866cb526c3b0a58219435e922067705e229c59e50f5f3cd01dc13eded"></a>

## where.site.disable_internet_vip — where.site.disable_internet_vip / 34b21f403886 / 2

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-89b3e09314ef5fec837c9034f68d6bb66048b1e8e127c21bebfbd751a57f9668)
- [where](data-sources--discovery--reference--group-001.md#canonical-86e8aa7a75f491571c4a99862108276e04b016f3b24020940f1e1b752bd38094)
- [where.site](data-sources--discovery--reference--group-001.md#canonical-5d56699cb08276324472b54d8dad817bce1c7511ad8ed36b5e76dfedc368d266)
- where.site.disable_internet_vip

<a id="canonical-8c3705fb0452ba438b37203e22bb6f5ce00f00e4a7d8b810d20de19e24ecc33f"></a>

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

<a id="canonical-2f08ca899e5858866621d204452c3f5ba47a6d02f22285116f348560e60f4e28"></a>

## Direct properties — where.site.disable_internet_vip / 34b21f403886 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cd14e08c6ba626f10c5f8fc1c568dea970ffc1ab8ffb8d9771e300048c55af54"></a>

## Next pages — where.site.disable_internet_vip / 34b21f403886 / 4

- [where.site](data-sources--discovery--reference--group-001.md#canonical-5d56699cb08276324472b54d8dad817bce1c7511ad8ed36b5e76dfedc368d266)
- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)

<a id="canonical-a8706eee0aa864b6da572208791fee81a1e9f0422a844aef2c68a5edce1b8cbe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-364a286d7bfe373978afb3a4288dc0ae0a9fc445b3b0b37c8a028d16a974dd85"></a>

## where.site.enable_internet_vip — where.site.enable_internet_vip / 0e6d81903e19 / 2

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-89b3e09314ef5fec837c9034f68d6bb66048b1e8e127c21bebfbd751a57f9668)
- [where](data-sources--discovery--reference--group-001.md#canonical-86e8aa7a75f491571c4a99862108276e04b016f3b24020940f1e1b752bd38094)
- [where.site](data-sources--discovery--reference--group-001.md#canonical-5d56699cb08276324472b54d8dad817bce1c7511ad8ed36b5e76dfedc368d266)
- where.site.enable_internet_vip

<a id="canonical-42218dbffdb319e12db9c73b3be22008ad8bd258f04d84a8e1aca408d741c9cb"></a>

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

<a id="canonical-f592122e6dd19d532e2b559eaeb9a70bea7c6344d90bb525437813b58cec6799"></a>

## Direct properties — where.site.enable_internet_vip / 0e6d81903e19 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9c00fa070079d4ee82cfbf1c03b09065e2f9be76033eab875b89da920cac0071"></a>

## Next pages — where.site.enable_internet_vip / 0e6d81903e19 / 4

- [where.site](data-sources--discovery--reference--group-001.md#canonical-5d56699cb08276324472b54d8dad817bce1c7511ad8ed36b5e76dfedc368d266)
- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)

<a id="canonical-1dad1ccc90a897fc903d4f8fe1740420c45839c16599c3129594d410de66c276"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-538a3f1be784a62d128be7302055e35a140b7d3e1a0223707752c40edb720df9"></a>

## where.site.ref — where.site.ref / cb0b8f4f6bae / 2

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-89b3e09314ef5fec837c9034f68d6bb66048b1e8e127c21bebfbd751a57f9668)
- [where](data-sources--discovery--reference--group-001.md#canonical-86e8aa7a75f491571c4a99862108276e04b016f3b24020940f1e1b752bd38094)
- [where.site](data-sources--discovery--reference--group-001.md#canonical-5d56699cb08276324472b54d8dad817bce1c7511ad8ed36b5e76dfedc368d266)
- where.site.ref

<a id="canonical-0714638296654be35a83072e86fecd2d78fd28093a3fb89b93fed18d3c4c852f"></a>

Type: `"list"`. Computed.

Reference. A site direct reference.

Upstream description:

A site direct reference.

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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-8c4601c570ff10166360f472fee05b84f37d63c807e63bf45ba2e3120d715178"></a>

## Direct properties — where.site.ref / cb0b8f4f6bae / 3

<a id="canonical-1c17d4ace85d6a62bfc1732948879311d2db243c30f14a91f2a1535e239c78e4"></a>

<a id="canonical-8cd3aaf8d96b114a142a4901326fb7ab23633a4e0bf5ff0b09391505af69c4cb"></a>

## kind property — where.site.ref / cb0b8f4f6bae / 4

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

<a id="canonical-ae21c70c0ca97266edb8e2ec339e9145b107f71a3bdf27a7e10568ffdd5034cf"></a>

<a id="canonical-92dbb5b48b396221d6bdcb527cb2d1bd72afffdf3bd3dca1970d235d41fba5e3"></a>

## name property — where.site.ref / cb0b8f4f6bae / 5

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

<a id="canonical-8d9d9b86ea6d40ba25811b646e5315ba53e0596dbd90e6814a88caf70de0cfa2"></a>

<a id="canonical-391d64f9d71e9155acc67c403ea2e4bc8c0c60a8b4b6a90f6cbd45802b30fe60"></a>

## namespace property — where.site.ref / cb0b8f4f6bae / 6

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

<a id="canonical-1419baeb4b93f03524ed4b8e757229fe74e0a6a9d69fbb1915bef012caa4d33c"></a>

<a id="canonical-623bd960125455dacd75ee70a52f910d4422ae874fb3fa6ea3b7cdb92ba3b0d9"></a>

## tenant property — where.site.ref / cb0b8f4f6bae / 7

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

<a id="canonical-f6327f377e379351bb1c0c07fb8152e2bb1bab0f06aacbcfe28a8025a261795b"></a>

<a id="canonical-808993158cbefb78de79d03df172ea2a832f5bdab1dcf2ea91cca868f51e8c16"></a>

## uid property — where.site.ref / cb0b8f4f6bae / 8

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

<a id="canonical-c00f443c0db68491dc396c79ffa8c6de4ffde67ddfcb71bff440d1a56dd7acd3"></a>

## Next pages — where.site.ref / cb0b8f4f6bae / 9

- [where.site](data-sources--discovery--reference--group-001.md#canonical-5d56699cb08276324472b54d8dad817bce1c7511ad8ed36b5e76dfedc368d266)
- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)

<a id="canonical-8042369f31b82ef4efd1e1e08aa7c5376469f4c4d96cb59b0f03957a64faccfb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6181e954d7e8b363ecc87264924937e073d3d86de16ce486c21f4f6f390d1981"></a>

## where.virtual_network — where.virtual_network / e07f56def38c / 2

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-89b3e09314ef5fec837c9034f68d6bb66048b1e8e127c21bebfbd751a57f9668)
- [where](data-sources--discovery--reference--group-001.md#canonical-86e8aa7a75f491571c4a99862108276e04b016f3b24020940f1e1b752bd38094)
- where.virtual_network

<a id="canonical-2b3126429bb09f20d3825ee12fdd9812ccf18bfcb309fea263f052fee9fadd62"></a>

Type: `"single"`. Computed.

Specifies a direct reference to a network configuration object.

Upstream description:

This specifies a direct reference to a network configuration object.

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

<a id="canonical-c78322e4c535625155680cf46c6b4513ae6dcbaaf9d338051c2e3d193ab266e9"></a>

## Direct properties — where.virtual_network / e07f56def38c / 3

- [ref](data-sources--discovery--reference--group-001.md#canonical-29c7fa2be7ea24fd66e783ab16b4d450115b8764fbd785ca9d730a6a2e4a5479): complete subsection reference.

<a id="canonical-429f865a9b25ac992533dce4cd0d9039b93454f1c594f298bfa9b57953b77954"></a>

## Next pages — where.virtual_network / e07f56def38c / 4

- [where.virtual_network.ref](data-sources--discovery--reference--group-001.md#canonical-29c7fa2be7ea24fd66e783ab16b4d450115b8764fbd785ca9d730a6a2e4a5479)
- [where](data-sources--discovery--reference--group-001.md#canonical-86e8aa7a75f491571c4a99862108276e04b016f3b24020940f1e1b752bd38094)
- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)

<a id="canonical-29c7fa2be7ea24fd66e783ab16b4d450115b8764fbd785ca9d730a6a2e4a5479"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a91c0306cae52d927277ebdf468dfdc96931c1e9f5085c5c84647035217a258a"></a>

## where.virtual_network.ref — where.virtual_network.ref / 10708a1f628e / 2

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-89b3e09314ef5fec837c9034f68d6bb66048b1e8e127c21bebfbd751a57f9668)
- [where](data-sources--discovery--reference--group-001.md#canonical-86e8aa7a75f491571c4a99862108276e04b016f3b24020940f1e1b752bd38094)
- [where.virtual_network](data-sources--discovery--reference--group-001.md#canonical-8042369f31b82ef4efd1e1e08aa7c5376469f4c4d96cb59b0f03957a64faccfb)
- where.virtual_network.ref

<a id="canonical-75078d1d10cb0092b83a7b12900c29cc3b758c714dbe1e659af7d38564926527"></a>

Type: `"list"`. Computed.

Reference. A virtual network direct reference.

Upstream description:

A virtual network direct reference.

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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-e47087bcb3d212c5c5f01ae5b858035ba04bdd104fb3ddd347e1c416958bef3c"></a>

## Direct properties — where.virtual_network.ref / 10708a1f628e / 3

<a id="canonical-9609e98183bd88f5cee50619e83769e6198a6569b794ed498687cd93a5ba1555"></a>

<a id="canonical-f74062e3e226cc499264e30aa43b0942cf6b0c1a30f6201efb0f6eeb8aabec9e"></a>

## kind property — where.virtual_network.ref / 10708a1f628e / 4

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

<a id="canonical-124e5a2331fc15954e1b880e62ab6d0675e980b59ba98e69c9ccbc94e46dd1ec"></a>

<a id="canonical-3c4ec9028b75a1e0c0ddd8b3fe21c32de66d8b61ca0f68a9a2723284d89fad5c"></a>

## name property — where.virtual_network.ref / 10708a1f628e / 5

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

<a id="canonical-992a4d08843ae283c3fd17fb428f602deed8aa4948ff146dd45a4626039065ae"></a>

<a id="canonical-c1240ac9595e8aad705fb6d479962636d89c67099508177a92fe3fe91674c730"></a>

## namespace property — where.virtual_network.ref / 10708a1f628e / 6

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

<a id="canonical-1a39c07a6d89b4a9d9243647445c3f2d2d95a7076b7190a2a146c901e12ffd44"></a>

<a id="canonical-fc70325ab7eaf282c0ccd0b88b8dd9b8694be145d7312874e2d691221ce5c678"></a>

## tenant property — where.virtual_network.ref / 10708a1f628e / 7

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

<a id="canonical-b3eb9a71d419ead07f15836c4daa902de6f7f1b73d3c369560dd91622e9b4d9a"></a>

<a id="canonical-353394c4ede657511e00e32cd42d0ed15ad6f4ed4285f480d77d925236075cd0"></a>

## uid property — where.virtual_network.ref / 10708a1f628e / 8

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

<a id="canonical-32b4f49515a3ecd7d7f820fed85da66f41cc435e5590d8a36ccbd24b120ead61"></a>

## Next pages — where.virtual_network.ref / 10708a1f628e / 9

- [where.virtual_network](data-sources--discovery--reference--group-001.md#canonical-8042369f31b82ef4efd1e1e08aa7c5376469f4c4d96cb59b0f03957a64faccfb)
- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)

<a id="canonical-9fad6288637564585ba5e19b4645c139a7c02678af28eed7d5f04b6a3aa9535a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6f3a6359b6d2d2fadfa7ad06c97813ac0931f2b4ec182663271aa42b164f3167"></a>

## where.virtual_site — where.virtual_site / c572bd858956 / 2

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-89b3e09314ef5fec837c9034f68d6bb66048b1e8e127c21bebfbd751a57f9668)
- [where](data-sources--discovery--reference--group-001.md#canonical-86e8aa7a75f491571c4a99862108276e04b016f3b24020940f1e1b752bd38094)
- where.virtual_site

<a id="canonical-d96be290ba2a7a660ca6220f0712df65ee278686d6e49f2135969d5f3f258e77"></a>

Type: `"single"`. Computed.

Virtual Site. A reference to virtual\_site object.

Upstream description:

A reference to virtual\_site object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-internet_vip_choice": "[\"disable_internet_vip\",\"enable_internet_vip\"]"
}
```

<a id="canonical-9add0c601cc3b1fc128d41816d18b80b180e4f76619150b90666cb7fd321b034"></a>

## Direct properties — where.virtual_site / c572bd858956 / 3

- [disable_internet_vip](data-sources--discovery--reference--group-002.md#canonical-b407e42215329b167a2e7a5086fd503c6ee783711dc21e63b995a1ae7fafec1a): complete subsection reference.

- [enable_internet_vip](data-sources--discovery--reference--group-002.md#canonical-8876a4d5abf6ebb6e64729ab6013f0f8f412e4b8f34fd37e401538bc852b92ea): complete subsection reference.

<a id="canonical-f823c97152688c3a87f4cbf0d7086dcb37081bedd08fc6557d1f608a00b853c9"></a>

<a id="canonical-323d59e3bdb22a17079db10d4f935e743b52141c00d23ffcf3516a83b9f69c4c"></a>

## network_type property — where.virtual_site / c572bd858956 / 4

Type: `"string"`. Computed.

\[Enum:
VIRTUAL\_NETWORK\_SITE\_LOCAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE|VIRTUAL\_NETWORK\_PER\_SITE|VIRTUAL\_NETWORK\_PUBLIC|VIRTUAL\_NETWORK\_GLOBAL|VIRTUAL\_NETWORK\_SITE\_SERVICE|VIRTUAL\_NETWORK\_VER\_INTERNAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE|VIRTUAL\_NETWORK\_IP\_AUTO|VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK|VIRTUAL\_NETWORK\_SRV6\_NETWORK|VIRTUAL\_NETWORK\_IP\_FABRIC|VIRTUAL\_NETWORK\_SEGMENT|VIRTUAL\_NETWORK\_MANAGEMENT\]
Different types of virtual networks understood by the system Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside) network. This is an insecure
network and is connected to public internet via NAT Gateways/firwalls Virtual-network of this type
is local to.. Possible values are \`VIRTUAL\_NETWORK\_SITE\_LOCAL\`,
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\`, \`VIRTUAL\_NETWORK\_PER\_SITE\`,
\`VIRTUAL\_NETWORK\_PUBLIC\`, \`VIRTUAL\_NETWORK\_GLOBAL\`, \`VIRTUAL\_NETWORK\_SITE\_SERVICE\`,
\`VIRTUAL\_NETWORK\_VER\_INTERNAL\`, \`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE\`,
\`VIRTUAL\_NETWORK\_IP\_AUTO\`, \`VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK\`,
\`VIRTUAL\_NETWORK\_SRV6\_NETWORK\`, \`VIRTUAL\_NETWORK\_IP\_FABRIC\`,
\`VIRTUAL\_NETWORK\_SEGMENT\`, \`VIRTUAL\_NETWORK\_MANAGEMENT\`. Defaults to
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\`.

Upstream description:

Different types of virtual networks understood by the system

Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside)
network. This is an insecure network and is connected to public internet via NAT Gateways/firwalls
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on CE sites. This network is created automatically and present on all sites
Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE is a private network inside site. It
is a secure network and is not connected to public network. Virtual-network of this type is local to
every site. Two virtual networks of this type on different sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on CE sites. This network is created during provisioning of site User defined per-site
virtual network. Scope of this virtual network is limited to the site. This is not yet supported
Virtual-network of type VIRTUAL\_NETWORK\_PUBLIC directly connects to the public internet.
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on RE sites only It is an internally created by the system. They must not be created by
user Virtual Networks with global scope across different sites in F5XC domain. An example global
virtual-network called "AIN Network" is created for every tenant. For F5 Distributed Cloud fabric

Constraints: It is currently only supported as internally created by the system. VK8s service
network for a given tenant. Used to advertise a virtual host only to vk8s pods for that tenant
Constraints: It is an internally created by the system. Must not be created by user VER internal
network for the site. It can only be used for virtual hosts with SMA\_PROXY type proxy Constraints:
It is an internally created by the system. Must not be created by user Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE represents both VIRTUAL\_NETWORK\_SITE\_LOCAL and
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE

Constraints: This network type is only meaningful in an advertise policy When virtual-network of
type VIRTUAL\_NETWORK\_IP\_AUTO is selected for an endpoint, VER will try to determine the network
based on the provided IP address

Constraints: This network type is only meaningful in an endpoint

VoltADN Private Network is used on F5 Distributed Cloud RE(s) to connect to customer private
networks This network is created by opening a support ticket

This network is per site srv6 network VER IP Fabric network for the site. This Virtual network type
is used for exposing virtual host on IP Fabric network on the VER site or for endpoint in IP Fabric
network Constraints: It is an internally created by the system. Must not be created by user
Virtual-network of type VIRTUAL\_NETWORK\_SEGMENT for segment interface Virtual-network of type
VIRTUAL\_NETWORK\_MANAGEMENT is used for management purposes.

Receipt-pinned upstream constraints:

```json
{
  "default": "VIRTUAL_NETWORK_SITE_LOCAL",
  "enum": [
    "VIRTUAL_NETWORK_SITE_LOCAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE",
    "VIRTUAL_NETWORK_PER_SITE",
    "VIRTUAL_NETWORK_PUBLIC",
    "VIRTUAL_NETWORK_GLOBAL",
    "VIRTUAL_NETWORK_SITE_SERVICE",
    "VIRTUAL_NETWORK_VER_INTERNAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE",
    "VIRTUAL_NETWORK_IP_AUTO",
    "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK",
    "VIRTUAL_NETWORK_SRV6_NETWORK",
    "VIRTUAL_NETWORK_IP_FABRIC",
    "VIRTUAL_NETWORK_SEGMENT",
    "VIRTUAL_NETWORK_MANAGEMENT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [ref](data-sources--discovery--reference--group-002.md#canonical-9646d5cad7404290af7064e00576aa6878886d07eea06a846ac08f27ca657745): complete subsection reference.
