---
page_title: "xcsh_authentication reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_authentication reference."
---

# xcsh_authentication reference

<a id="canonical-c4df54315d421dd70758e3db0bb63b6e69f2aee4a7ec3a11548b735bedc072c3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-df01d0159e680fed2ec3f64640732abb293671c318dee97c70efe61ad6300fc6"></a>

## Property reference — Property reference / e70deceb3bbb / 2

Breadcrumbs:

- [xcsh_authentication](../resources/authentication.md#canonical-417d73493652f5a1771ef46f280f23ffd5c91667377a08f64262e5ac9ba40724)
- Property reference

<a id="canonical-2f9d5bb54140ae1b04c42a5779b730e3fee1fd5e6eff06f38721495941821644"></a>

## Direct properties — Property reference / e70deceb3bbb / 3

<a id="canonical-e7f081deb3b1fc8f649559fa9636fc1d2a1b03f53cc5e23b1ae8312b40d31642"></a>

<a id="canonical-b14eff2f06f8e8b6ac2495a4fd38e1006dada240cf03832398f5514c16884673"></a>

## annotations property — Property reference / e70deceb3bbb / 4

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata.

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

- [cookie_params](resources--authentication--reference--group-001.md#canonical-388b68d284e492b110f52b383cfd53ce4a47a7ae27b24501356f1b1a4c640894): complete subsection reference.

<a id="canonical-7e3142b644a47a20441b3d72c365c848592e3b323a11f39b965a1b5b5ea7d050"></a>

<a id="canonical-23ad0e0eb2e0c3d29b007ba57b4cc1f0559d48008ebb80ea461761e150f4c3ea"></a>

## description property — Property reference / e70deceb3bbb / 5

Type: `"string"`. Optional.

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

<a id="canonical-c82df17cd88b9934a4ab2ad666755ddbb487a06ca2504dc8a091a08559dbb002"></a>

<a id="canonical-ad4243facc4257fefa15b293637a07d568d0737c8014ebca8500c4057bf4247b"></a>

## disable property — Property reference / e70deceb3bbb / 6

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Upstream description:

A value of true will administratively disable the object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-88f9d191847f0e5db9cd039cb2b5ccd37dabb19e9d5cba10e8aa6f06f7896931"></a>

<a id="canonical-73e2e0aeab5ba745fbd5945f81e74d8f640dadc1bf81d1ebcb987fd9e6168652"></a>

## id property — Property reference / e70deceb3bbb / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-2ac067a04b3f8bf16d53ffc1fbfb61b38a97643af13f966a19b750300066d9d8"></a>

<a id="canonical-4326fa0ab1d8b8f0ad2839d8023984cddcd39489f86d286b80926ef000f58b40"></a>

## labels property — Property reference / e70deceb3bbb / 8

Type: `["map", "string"]`. Optional.

Labels is a user defined key value map that can be attached to resources for organization and
filtering.

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

<a id="canonical-4b0869525274f8b2a85b6028b2696f6cf9322cec10d0694c22d6a2613aad993d"></a>

<a id="canonical-f4bb657faf29272be2fa2f08f406525b85e58fa0bba68ca32d8dade9eff35771"></a>

## name property — Property reference / e70deceb3bbb / 9

Type: `"string"`. Required.

Name of the Authentication. Must be unique within the namespace.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NameValidator(),
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

<a id="canonical-16d676e6b5c7edd899c1f1fedb3f8dd212906005911ec226b40cbbac9dd16cfa"></a>

<a id="canonical-bb23fb7b5518ccf80fa5359d48881af6a4620210f291c9eb5af2fcd39ac3ffcf"></a>

## namespace property — Property reference / e70deceb3bbb / 10

Type: `"string"`. Required.

Namespace where the Authentication is created.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NamespaceValidator(),
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

- [oidc_auth](resources--authentication--reference--group-001.md#canonical-2e567deb2f14580d6e05093b8e3a34311034e624bcff1f38e036f6d78ebe3faa): complete subsection reference.

- [timeouts](resources--authentication--reference--group-001.md#canonical-f3d9b4a248b762990240bfd0ec12afe3d19dc70df6b11e56bb750677650dbc0c): complete subsection reference.

<a id="canonical-eb27d1a76d79c65a2a106eb44930b132642b85f50b4c2bbab356f93f4da0fa99"></a>

## All schema paths — Property reference / e70deceb3bbb / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--authentication--reference--group-001.md#canonical-e7f081deb3b1fc8f649559fa9636fc1d2a1b03f53cc5e23b1ae8312b40d31642) |
| `cookie_params` | [cookie_params](resources--authentication--reference--group-001.md#canonical-519eef6b1059effa4acea52a4a974b2c80ab512e9e688d98d5c088cb97c0d450) |
| `cookie_params.auth_hmac` | [cookie_params.auth_hmac](resources--authentication--reference--group-001.md#canonical-d86a5d2d32b9220ca91ab190396f006d9eb3e12e0b560a1d6a747ba0f803b275) |
| `cookie_params.auth_hmac.prim_key` | [cookie_params.auth_hmac.prim_key](resources--authentication--reference--group-001.md#canonical-2eff77495e12ae843e251a04e31a675ac250853c3a85e1f0df63a33f64706b14) |
| `cookie_params.auth_hmac.prim_key.blindfold_secret_info` | [cookie_params.auth_hmac.prim_key.blindfold_secret_info](resources--authentication--reference--group-001.md#canonical-c49415e7ea62183b3d3b591d6b7bcca278d310ec5e22bc0065df9be2e4faddd8) |
| `cookie_params.auth_hmac.prim_key.blindfold_secret_info.decryption_provider` | [cookie_params.auth_hmac.prim_key.blindfold_secret_info.decryption_provider](resources--authentication--reference--group-001.md#canonical-fc392e116a148d62bcb618aa127b43b8b75851f9e85aaf874951f60117b562cf) |
| `cookie_params.auth_hmac.prim_key.blindfold_secret_info.location` | [cookie_params.auth_hmac.prim_key.blindfold_secret_info.location](resources--authentication--reference--group-001.md#canonical-6b7bd8dc970517f2bb376cf987cb24993d09fff47071f0ff444a165b52b3d591) |
| `cookie_params.auth_hmac.prim_key.blindfold_secret_info.store_provider` | [cookie_params.auth_hmac.prim_key.blindfold_secret_info.store_provider](resources--authentication--reference--group-001.md#canonical-f71dd2be76253e0f04289bddae73ca035a629c3bc5340a5304fcff29025a7110) |
| `cookie_params.auth_hmac.prim_key.clear_secret_info` | [cookie_params.auth_hmac.prim_key.clear_secret_info](resources--authentication--reference--group-001.md#canonical-a08b3215cef078d9a332229a9b73c36b7f99a224327e568ebab3084b64664678) |
| `cookie_params.auth_hmac.prim_key.clear_secret_info.provider_ref` | [cookie_params.auth_hmac.prim_key.clear_secret_info.provider_ref](resources--authentication--reference--group-001.md#canonical-6ddd685a3d062e8509c03bc318fa182fecaf6fd1f1067d48cfcce8c531df1c04) |
| `cookie_params.auth_hmac.prim_key.clear_secret_info.url` | [cookie_params.auth_hmac.prim_key.clear_secret_info.url](resources--authentication--reference--group-001.md#canonical-f77515edb7ae55202367f2dc24fe947e5106a3952b1d661c7634442c4aa524bd) |
| `cookie_params.auth_hmac.prim_key_expiry` | [cookie_params.auth_hmac.prim_key_expiry](resources--authentication--reference--group-001.md#canonical-a389a96209c574a53fd143c2a5959b202dcbcfd0d797c67121b626f691874427) |
| `cookie_params.auth_hmac.sec_key` | [cookie_params.auth_hmac.sec_key](resources--authentication--reference--group-001.md#canonical-c30b00e0dd7ae068303e58c0646e9f9959296867119c7c25a5b59b85e01ecd65) |
| `cookie_params.auth_hmac.sec_key.blindfold_secret_info` | [cookie_params.auth_hmac.sec_key.blindfold_secret_info](resources--authentication--reference--group-001.md#canonical-35b3a8e73b9af5b15ecb01d1f54739e76209d88ec8ddd3da1d9be24d5387756b) |
| `cookie_params.auth_hmac.sec_key.blindfold_secret_info.decryption_provider` | [cookie_params.auth_hmac.sec_key.blindfold_secret_info.decryption_provider](resources--authentication--reference--group-001.md#canonical-fc1871bce4f1be36c35fa619fa12f209d51f9100953dd6794f441a1b2d37effb) |
| `cookie_params.auth_hmac.sec_key.blindfold_secret_info.location` | [cookie_params.auth_hmac.sec_key.blindfold_secret_info.location](resources--authentication--reference--group-001.md#canonical-a835ed14ae044385eb23c2e9cbc44a6edc7f8b5d9d9f3cc3cb284e81dac833cf) |
| `cookie_params.auth_hmac.sec_key.blindfold_secret_info.store_provider` | [cookie_params.auth_hmac.sec_key.blindfold_secret_info.store_provider](resources--authentication--reference--group-001.md#canonical-6e2113e2be731a69972c9167a5afb82dae8058134d0cc14a31dcf50112d4e361) |
| `cookie_params.auth_hmac.sec_key.clear_secret_info` | [cookie_params.auth_hmac.sec_key.clear_secret_info](resources--authentication--reference--group-001.md#canonical-dfe02a41712801ca962b88cffd6ce94c65fadfcc734adbf8ce504abab7a1cfac) |
| `cookie_params.auth_hmac.sec_key.clear_secret_info.provider_ref` | [cookie_params.auth_hmac.sec_key.clear_secret_info.provider_ref](resources--authentication--reference--group-001.md#canonical-b8ab64af16b7666b9e9d09d503c709b9a9cf47a36cbce11f99db043417e1f55b) |
| `cookie_params.auth_hmac.sec_key.clear_secret_info.url` | [cookie_params.auth_hmac.sec_key.clear_secret_info.url](resources--authentication--reference--group-001.md#canonical-9cf40f4085dae15ed47aa0b55c436b832b63dea8433570b875bcc06a20065796) |
| `cookie_params.auth_hmac.sec_key_expiry` | [cookie_params.auth_hmac.sec_key_expiry](resources--authentication--reference--group-001.md#canonical-414991a052b0300da29311bab12268e66e6a490c20960958acc18c79d62b3f1e) |
| `cookie_params.cookie_expiry` | [cookie_params.cookie_expiry](resources--authentication--reference--group-001.md#canonical-77e66dbb11cfe9fccefc4d370eb73478d927f39b5d58681c8ac79f2bb8a26abe) |
| `cookie_params.cookie_refresh_interval` | [cookie_params.cookie_refresh_interval](resources--authentication--reference--group-001.md#canonical-61f6d6e04fce575db1dbdec8f44daaed0552308442e91ef972b0660acd12b7e7) |
| `cookie_params.kms_key_hmac` | [cookie_params.kms_key_hmac](resources--authentication--reference--group-001.md#canonical-c0f75baed8f4e4b8c48d2310b75ed97a11af3a6989763fedca4d2b3379692b5e) |
| `cookie_params.session_expiry` | [cookie_params.session_expiry](resources--authentication--reference--group-001.md#canonical-f462d28d4af048b27bcf0300965960f53ff5f83a748e436e98788e5c96b124c7) |
| `description` | [description](resources--authentication--reference--group-001.md#canonical-7e3142b644a47a20441b3d72c365c848592e3b323a11f39b965a1b5b5ea7d050) |
| `disable` | [disable](resources--authentication--reference--group-001.md#canonical-c82df17cd88b9934a4ab2ad666755ddbb487a06ca2504dc8a091a08559dbb002) |
| `id` | [id](resources--authentication--reference--group-001.md#canonical-88f9d191847f0e5db9cd039cb2b5ccd37dabb19e9d5cba10e8aa6f06f7896931) |
| `labels` | [labels](resources--authentication--reference--group-001.md#canonical-2ac067a04b3f8bf16d53ffc1fbfb61b38a97643af13f966a19b750300066d9d8) |
| `name` | [name](resources--authentication--reference--group-001.md#canonical-4b0869525274f8b2a85b6028b2696f6cf9322cec10d0694c22d6a2613aad993d) |
| `namespace` | [namespace](resources--authentication--reference--group-001.md#canonical-16d676e6b5c7edd899c1f1fedb3f8dd212906005911ec226b40cbbac9dd16cfa) |
| `oidc_auth` | [oidc_auth](resources--authentication--reference--group-001.md#canonical-48e7a126ba16299db6a46ba8a0438e02329942fcf8bdbc4434f5d0c9b23a2872) |
| `oidc_auth.client_secret` | [oidc_auth.client_secret](resources--authentication--reference--group-001.md#canonical-f1471e4b85507fe4e773849820ec63d5815e9c6d81e4c0f01344352a51e080a9) |
| `oidc_auth.client_secret.blindfold_secret_info` | [oidc_auth.client_secret.blindfold_secret_info](resources--authentication--reference--group-001.md#canonical-d6acc683f52976c946e4f4217c905aa53ae573ea25c5345d81c9ccb4bc6e15c5) |
| `oidc_auth.client_secret.blindfold_secret_info.decryption_provider` | [oidc_auth.client_secret.blindfold_secret_info.decryption_provider](resources--authentication--reference--group-001.md#canonical-ce025cc8d5aa1c67d4935b431b89a33b85b6f7b7486c52f9fef7efd7ecda5c6a) |
| `oidc_auth.client_secret.blindfold_secret_info.location` | [oidc_auth.client_secret.blindfold_secret_info.location](resources--authentication--reference--group-001.md#canonical-d3f81d3eb539972689ec8ee5bfea7199097116d76edcfa7fe1d6274faca818c4) |
| `oidc_auth.client_secret.blindfold_secret_info.store_provider` | [oidc_auth.client_secret.blindfold_secret_info.store_provider](resources--authentication--reference--group-001.md#canonical-7fe64fe54c5e2e71af85ab8d1ff6d4541d9dd6f6afd87c5386df90ea56fd9960) |
| `oidc_auth.client_secret.clear_secret_info` | [oidc_auth.client_secret.clear_secret_info](resources--authentication--reference--group-001.md#canonical-e3a89c2b38fc576de42abc97a929fded33d64337edef90e2e0b996bc6c8e7d3f) |
| `oidc_auth.client_secret.clear_secret_info.provider_ref` | [oidc_auth.client_secret.clear_secret_info.provider_ref](resources--authentication--reference--group-001.md#canonical-28269df732c34025b664bec2622c58d4662d9e50f9454d85b19ca3d7f5504bbd) |
| `oidc_auth.client_secret.clear_secret_info.url` | [oidc_auth.client_secret.clear_secret_info.url](resources--authentication--reference--group-001.md#canonical-8a36e6c67b1aeea91b42cf901f8f96077afde596cd79650bb624320462aac533) |
| `oidc_auth.oidc_auth_params` | [oidc_auth.oidc_auth_params](resources--authentication--reference--group-001.md#canonical-f3fd11a1cf7a6c471b0531bff5f0076cd92adcb337cbd76f87c8d0c05bde8d3b) |
| `oidc_auth.oidc_auth_params.auth_endpoint_url` | [oidc_auth.oidc_auth_params.auth_endpoint_url](resources--authentication--reference--group-001.md#canonical-26544726cc764047b66623bc4664595208866df67cff71ded28202077f93978f) |
| `oidc_auth.oidc_auth_params.end_session_endpoint_url` | [oidc_auth.oidc_auth_params.end_session_endpoint_url](resources--authentication--reference--group-001.md#canonical-3a04cfda7585adcfe708377cd81ca88951d2d2fea7017ea47bfaf09c9695e89f) |
| `oidc_auth.oidc_auth_params.token_endpoint_url` | [oidc_auth.oidc_auth_params.token_endpoint_url](resources--authentication--reference--group-001.md#canonical-51057b65909a42872849ef6e2f4229bbae9af8cde3f2e221a69e46483cf844aa) |
| `oidc_auth.oidc_client_id` | [oidc_auth.oidc_client_id](resources--authentication--reference--group-001.md#canonical-0dc6ede03efb3f0597cd6aef018f246b1b5edd819770fb315a0ceb0f8278bfac) |
| `oidc_auth.oidc_well_known_config_url` | [oidc_auth.oidc_well_known_config_url](resources--authentication--reference--group-001.md#canonical-722122c8ee0935f80128a13ec05761e74d8aca5143cbb04539fc24f3d0b56e58) |
| `timeouts` | [timeouts](resources--authentication--reference--group-001.md#canonical-76b7c92e264865a56bfa33ebebf977419bd5694207e59e7448b3735977fd3966) |
| `timeouts.create` | [timeouts.create](resources--authentication--reference--group-001.md#canonical-f97bf9951fea525c16a6585a973e19d61636ca409031ce84f4b36e86b37fbaa7) |
| `timeouts.delete` | [timeouts.delete](resources--authentication--reference--group-001.md#canonical-0665dc32e2090665888270c7aeec01404783fca42d7711598737fd1ba03a5a2e) |
| `timeouts.read` | [timeouts.read](resources--authentication--reference--group-001.md#canonical-6f2dad3306717c3c6fe5076a39303565cef6319fe214614a085ea33889c9f3df) |
| `timeouts.update` | [timeouts.update](resources--authentication--reference--group-001.md#canonical-3875566ac7dec4595a1a72a336d0af9ab6d4ae28cdbd2a5714ffeadd1c018eae) |

<a id="canonical-5bd7926141703677ae37cc13e9ee546a2ceb32d92d9a2d95c1270021c4adcbb0"></a>

## Next pages — Property reference / e70deceb3bbb / 12

- [cookie_params](resources--authentication--reference--group-001.md#canonical-388b68d284e492b110f52b383cfd53ce4a47a7ae27b24501356f1b1a4c640894)
- [oidc_auth](resources--authentication--reference--group-001.md#canonical-2e567deb2f14580d6e05093b8e3a34311034e624bcff1f38e036f6d78ebe3faa)
- [timeouts](resources--authentication--reference--group-001.md#canonical-f3d9b4a248b762990240bfd0ec12afe3d19dc70df6b11e56bb750677650dbc0c)
- [xcsh_authentication](../resources/authentication.md#canonical-417d73493652f5a1771ef46f280f23ffd5c91667377a08f64262e5ac9ba40724)

<a id="canonical-388b68d284e492b110f52b383cfd53ce4a47a7ae27b24501356f1b1a4c640894"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-516cb121723523a5a83588c3e538daeb5374df2f0dd24f76439077b422ce6882"></a>

## cookie_params — cookie_params / 542c873f1f87 / 2

Breadcrumbs:

- [xcsh_authentication](../resources/authentication.md#canonical-417d73493652f5a1771ef46f280f23ffd5c91667377a08f64262e5ac9ba40724)
- [Property reference](resources--authentication--reference--group-001.md#canonical-c4df54315d421dd70758e3db0bb63b6e69f2aee4a7ec3a11548b735bedc072c3)
- cookie_params

<a id="canonical-519eef6b1059effa4acea52a4a974b2c80ab512e9e688d98d5c088cb97c0d450"></a>

Type: `"object"`. single nested block, Optional.

Specifies different cookie related config parameters for authentication.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("auth_hmac",
    "kms_key_hmac")}
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
  "x-ves-oneof-field-secret_choice": "[\"auth_hmac\",\"kms_key_hmac\"]"
}
```

Terraform syntax:

```terraform
cookie_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-33a5e3e973825460afbbbcfd8be595fd35f4d7461f3c3263db8b8c20ec0741b2"></a>

## Direct properties — cookie_params / 542c873f1f87 / 3

- [auth_hmac](resources--authentication--reference--group-001.md#canonical-4ccf76209bd9ae38fb2902738a78c7f7dc8594989e9d18ad6c4a9e56adb4f931): complete subsection reference.

<a id="canonical-77e66dbb11cfe9fccefc4d370eb73478d927f39b5d58681c8ac79f2bb8a26abe"></a>

<a id="canonical-ef17fb686d9f55a2cb2c9dd79774ad97b06f9e76de5bf38a17ab5d15e4739bf2"></a>

## cookie_expiry property — cookie_params / 542c873f1f87 / 4

Type: `"number"`. Optional.

Specifies in seconds max duration of the allocated cookie. This maps to “Max-Age” attribute in the
session cookie. This will act as an expiry duration on the client side after which client will not
be setting the cookie as part of the request.

Upstream description:

Specifies in seconds max duration of the allocated cookie. This maps to “Max-Age” attribute in the
session cookie. This will act as an expiry duration on the client side after which client will not
be setting the cookie as part of the request. Default cookie expiry is 3600 seconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(86400),
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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

<a id="canonical-61f6d6e04fce575db1dbdec8f44daaed0552308442e91ef972b0660acd12b7e7"></a>

<a id="canonical-7b8df49f2c2c92a2602d071e612ad94a64524fd6b12774a3c2f0bc47d2fbee85"></a>

## cookie_refresh_interval property — cookie_params / 542c873f1f87 / 5

Type: `"number"`. Optional.

Specifies in seconds refresh interval for session cookie. This is used to keep the active user
active and reduce RE-login. When an incoming cookie's session expiry is still valid, and time to
expire falls behind this interval, RE-issue a cookie with new expiry and with the same original
session..

Upstream description:

Specifies in seconds refresh interval for session cookie. This is used to keep the active user
active and reduce RE-login. When an incoming cookie's session expiry is still valid, and time to
expire falls behind this interval, RE-issue a cookie with new expiry and with the same original
session expiry. Default refresh interval is 3000 seconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(86400),
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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

- [kms_key_hmac](resources--authentication--reference--group-001.md#canonical-839a32944466cddbe8406764bc6097e2c3a9f11b665c9c2895a455405ecf45c4): complete subsection reference.

<a id="canonical-f462d28d4af048b27bcf0300965960f53ff5f83a748e436e98788e5c96b124c7"></a>

<a id="canonical-a16d53ab3922fe664582248fa5e8325bcc1b993954dbac2763809a9a217da260"></a>

## session_expiry property — cookie_params / 542c873f1f87 / 6

Type: `"number"`. Optional.

Specifies in seconds max lifetime of an authenticated session after which the user will be forced to
login again. Default session expiry is 86400 seconds(24 hours).

Upstream description:

Specifies in seconds max lifetime of an authenticated session after which the user will be forced to
login again. Default session expiry is 86400 seconds(24 hours).

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(1296000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1296000,
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
    "ves.io.schema.rules.uint32.lte": "1296000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "1296000"
  }
}
```

<a id="canonical-9751da14928f80446ff8f6cd3bb62a2a9050a3b85395f40e83bac759cefa7e47"></a>

## Next pages — cookie_params / 542c873f1f87 / 7

- [cookie_params.auth_hmac](resources--authentication--reference--group-001.md#canonical-4ccf76209bd9ae38fb2902738a78c7f7dc8594989e9d18ad6c4a9e56adb4f931)
- [cookie_params.kms_key_hmac](resources--authentication--reference--group-001.md#canonical-839a32944466cddbe8406764bc6097e2c3a9f11b665c9c2895a455405ecf45c4)
- [Property reference](resources--authentication--reference--group-001.md#canonical-c4df54315d421dd70758e3db0bb63b6e69f2aee4a7ec3a11548b735bedc072c3)
- [xcsh_authentication](../resources/authentication.md#canonical-417d73493652f5a1771ef46f280f23ffd5c91667377a08f64262e5ac9ba40724)

<a id="canonical-4ccf76209bd9ae38fb2902738a78c7f7dc8594989e9d18ad6c4a9e56adb4f931"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-013fab46f5fa02b0f977bbced151643d61ae5e9383f96f0dc2edc505d3a3b9ab"></a>

## cookie_params.auth_hmac — cookie_params.auth_hmac / a2e132a2b4a3 / 2

Breadcrumbs:

- [xcsh_authentication](../resources/authentication.md#canonical-417d73493652f5a1771ef46f280f23ffd5c91667377a08f64262e5ac9ba40724)
- [Property reference](resources--authentication--reference--group-001.md#canonical-c4df54315d421dd70758e3db0bb63b6e69f2aee4a7ec3a11548b735bedc072c3)
- [cookie_params](resources--authentication--reference--group-001.md#canonical-388b68d284e492b110f52b383cfd53ce4a47a7ae27b24501356f1b1a4c640894)
- cookie_params.auth_hmac

<a id="canonical-d86a5d2d32b9220ca91ab190396f006d9eb3e12e0b560a1d6a747ba0f803b275"></a>

Type: `"object"`. single nested block, Optional.

HMAC primary and secondary keys to be used for hashing the Cookie. Each key also have an associated
expiry timestamp, beyond which key is invalid.

Upstream description:

HMAC primary and secondary keys to be used for hashing the Cookie. Each key also have an associated
expiry timestamp, beyond which key is invalid.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("prim_key_expiry",
    "sec_key_expiry")}
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
auth_hmac {
  # Configure direct properties listed below.
}
```

<a id="canonical-16db0fed833cc228d0c3f01f32fcd496fabd99dc513f5a900461c7a86b8f4647"></a>

## Direct properties — cookie_params.auth_hmac / a2e132a2b4a3 / 3

- [prim_key](resources--authentication--reference--group-001.md#canonical-ce65bd0d2e6dff20c84e468e4817de3bf0d18e8eb9178e12326f09cb69c14c6b): complete subsection reference.

<a id="canonical-a389a96209c574a53fd143c2a5959b202dcbcfd0d797c67121b626f691874427"></a>

<a id="canonical-dad05bcda32066bd04d8010989ffd1899b00fb7be5e81f6a85c8425d2825e004"></a>

## prim_key_expiry property — cookie_params.auth_hmac / a2e132a2b4a3 / 4

Type: `"string"`. Optional.

HMAC Primary Key Expiry. Primary HMAC Key Expiry time.

Upstream description:

Primary HMAC Key Expiry time.

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

- [sec_key](resources--authentication--reference--group-001.md#canonical-3cbabb805df13286253db6e6b48c8ed787a2afd6940f44f7d440b1304ee2ea0f): complete subsection reference.

<a id="canonical-414991a052b0300da29311bab12268e66e6a490c20960958acc18c79d62b3f1e"></a>

<a id="canonical-776daffafbbc80f415b11871119e7bcd2b2298d778b4ebd2badd52f21fdb0148"></a>

## sec_key_expiry property — cookie_params.auth_hmac / a2e132a2b4a3 / 5

Type: `"string"`. Optional.

HMAC Secondary Key Expiry. Secondary HMAC Key Expiry time.

Upstream description:

Secondary HMAC Key Expiry time.

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

<a id="canonical-375c238223533cd0d13bd6993ec74030463c3d1c6de38d9e11c3ddb6d509be77"></a>

## Next pages — cookie_params.auth_hmac / a2e132a2b4a3 / 6

- [cookie_params.auth_hmac.prim_key](resources--authentication--reference--group-001.md#canonical-ce65bd0d2e6dff20c84e468e4817de3bf0d18e8eb9178e12326f09cb69c14c6b)
- [cookie_params.auth_hmac.sec_key](resources--authentication--reference--group-001.md#canonical-3cbabb805df13286253db6e6b48c8ed787a2afd6940f44f7d440b1304ee2ea0f)
- [cookie_params](resources--authentication--reference--group-001.md#canonical-388b68d284e492b110f52b383cfd53ce4a47a7ae27b24501356f1b1a4c640894)
- [xcsh_authentication](../resources/authentication.md#canonical-417d73493652f5a1771ef46f280f23ffd5c91667377a08f64262e5ac9ba40724)

<a id="canonical-ce65bd0d2e6dff20c84e468e4817de3bf0d18e8eb9178e12326f09cb69c14c6b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-86db90e59feab520392de94a77130414806a9a5d452f7482c2ec1441ee519402"></a>

## cookie_params.auth_hmac.prim_key — cookie_params.auth_hmac.prim_key / 293feada71d2 / 2

Breadcrumbs:

- [xcsh_authentication](../resources/authentication.md#canonical-417d73493652f5a1771ef46f280f23ffd5c91667377a08f64262e5ac9ba40724)
- [Property reference](resources--authentication--reference--group-001.md#canonical-c4df54315d421dd70758e3db0bb63b6e69f2aee4a7ec3a11548b735bedc072c3)
- [cookie_params](resources--authentication--reference--group-001.md#canonical-388b68d284e492b110f52b383cfd53ce4a47a7ae27b24501356f1b1a4c640894)
- [cookie_params.auth_hmac](resources--authentication--reference--group-001.md#canonical-4ccf76209bd9ae38fb2902738a78c7f7dc8594989e9d18ad6c4a9e56adb4f931)
- cookie_params.auth_hmac.prim_key

<a id="canonical-2eff77495e12ae843e251a04e31a675ac250853c3a85e1f0df63a33f64706b14"></a>

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
prim_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-5a02adb896487d926b219358651fed58bf5674683b69409232fd893695ca4431"></a>

## Direct properties — cookie_params.auth_hmac.prim_key / 293feada71d2 / 3

- [blindfold_secret_info](resources--authentication--reference--group-001.md#canonical-8bf40ccb677f923021ee2eb42032777561e614aab481fc00411fed83e603c720): complete subsection reference.

- [clear_secret_info](resources--authentication--reference--group-001.md#canonical-14e97b4b771c2e9820c30a21480571fa45fbc19a2b61698cbc46fa38ed949f64): complete subsection reference.

<a id="canonical-57abd71d58c52adfc7534ce278791ec3eaf771eb22ece0317fa1b0676bbaac48"></a>

## Next pages — cookie_params.auth_hmac.prim_key / 293feada71d2 / 4

- [cookie_params.auth_hmac.prim_key.blindfold_secret_info](resources--authentication--reference--group-001.md#canonical-8bf40ccb677f923021ee2eb42032777561e614aab481fc00411fed83e603c720)
- [cookie_params.auth_hmac.prim_key.clear_secret_info](resources--authentication--reference--group-001.md#canonical-14e97b4b771c2e9820c30a21480571fa45fbc19a2b61698cbc46fa38ed949f64)
- [cookie_params.auth_hmac](resources--authentication--reference--group-001.md#canonical-4ccf76209bd9ae38fb2902738a78c7f7dc8594989e9d18ad6c4a9e56adb4f931)
- [xcsh_authentication](../resources/authentication.md#canonical-417d73493652f5a1771ef46f280f23ffd5c91667377a08f64262e5ac9ba40724)

<a id="canonical-8bf40ccb677f923021ee2eb42032777561e614aab481fc00411fed83e603c720"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7b393eb57893e85d002c0c2ba924fb02e9c85eb26d766879593b70aa87f8dc2b"></a>

## cookie_params.auth_hmac.prim_key.blindfold_secret_info — cookie_params.auth_hmac.prim_key.blindfold_secret_info / 907170c3745f / 2

Breadcrumbs:

- [xcsh_authentication](../resources/authentication.md#canonical-417d73493652f5a1771ef46f280f23ffd5c91667377a08f64262e5ac9ba40724)
- [Property reference](resources--authentication--reference--group-001.md#canonical-c4df54315d421dd70758e3db0bb63b6e69f2aee4a7ec3a11548b735bedc072c3)
- [cookie_params](resources--authentication--reference--group-001.md#canonical-388b68d284e492b110f52b383cfd53ce4a47a7ae27b24501356f1b1a4c640894)
- [cookie_params.auth_hmac](resources--authentication--reference--group-001.md#canonical-4ccf76209bd9ae38fb2902738a78c7f7dc8594989e9d18ad6c4a9e56adb4f931)
- [cookie_params.auth_hmac.prim_key](resources--authentication--reference--group-001.md#canonical-ce65bd0d2e6dff20c84e468e4817de3bf0d18e8eb9178e12326f09cb69c14c6b)
- cookie_params.auth_hmac.prim_key.blindfold_secret_info

<a id="canonical-c49415e7ea62183b3d3b591d6b7bcca278d310ec5e22bc0065df9be2e4faddd8"></a>

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

<a id="canonical-f96fe79d0f27cbfb987a3b6fa07288741702c02f2f101e5a51f765e0a19db048"></a>

## Direct properties — cookie_params.auth_hmac.prim_key.blindfold_secret_info / 907170c3745f / 3

<a id="canonical-fc392e116a148d62bcb618aa127b43b8b75851f9e85aaf874951f60117b562cf"></a>

<a id="canonical-92da67624be7f50051b09b718864a58bb4b254c7c6039bbca3e9d07958266073"></a>

## decryption_provider property — cookie_params.auth_hmac.prim_key.blindfold_secret_info / 907170c3745f / 4

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

<a id="canonical-6b7bd8dc970517f2bb376cf987cb24993d09fff47071f0ff444a165b52b3d591"></a>

<a id="canonical-b7e18098cfcf80b01a14345e482a3d1ad5eda719293f522ec4ebcdf3c6d34a9d"></a>

## location property — cookie_params.auth_hmac.prim_key.blindfold_secret_info / 907170c3745f / 5

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

<a id="canonical-f71dd2be76253e0f04289bddae73ca035a629c3bc5340a5304fcff29025a7110"></a>

<a id="canonical-234e0f67e25e23383a9fc2d719ec77e2169a97019890559de4c7c714e9e6b282"></a>

## store_provider property — cookie_params.auth_hmac.prim_key.blindfold_secret_info / 907170c3745f / 6

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

<a id="canonical-7f580498743fee0f555fd67fe2b473f4fb08ae750d32c32ccf5d5381a09f5e67"></a>

## Next pages — cookie_params.auth_hmac.prim_key.blindfold_secret_info / 907170c3745f / 7

- [cookie_params.auth_hmac.prim_key](resources--authentication--reference--group-001.md#canonical-ce65bd0d2e6dff20c84e468e4817de3bf0d18e8eb9178e12326f09cb69c14c6b)
- [xcsh_authentication](../resources/authentication.md#canonical-417d73493652f5a1771ef46f280f23ffd5c91667377a08f64262e5ac9ba40724)

<a id="canonical-14e97b4b771c2e9820c30a21480571fa45fbc19a2b61698cbc46fa38ed949f64"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-97709ca7c83a740426223e01feeba069d86d950c7d4100d19cb58e2e192c84ff"></a>

## cookie_params.auth_hmac.prim_key.clear_secret_info — cookie_params.auth_hmac.prim_key.clear_secret_info / ccd7cccbe23f / 2

Breadcrumbs:

- [xcsh_authentication](../resources/authentication.md#canonical-417d73493652f5a1771ef46f280f23ffd5c91667377a08f64262e5ac9ba40724)
- [Property reference](resources--authentication--reference--group-001.md#canonical-c4df54315d421dd70758e3db0bb63b6e69f2aee4a7ec3a11548b735bedc072c3)
- [cookie_params](resources--authentication--reference--group-001.md#canonical-388b68d284e492b110f52b383cfd53ce4a47a7ae27b24501356f1b1a4c640894)
- [cookie_params.auth_hmac](resources--authentication--reference--group-001.md#canonical-4ccf76209bd9ae38fb2902738a78c7f7dc8594989e9d18ad6c4a9e56adb4f931)
- [cookie_params.auth_hmac.prim_key](resources--authentication--reference--group-001.md#canonical-ce65bd0d2e6dff20c84e468e4817de3bf0d18e8eb9178e12326f09cb69c14c6b)
- cookie_params.auth_hmac.prim_key.clear_secret_info

<a id="canonical-a08b3215cef078d9a332229a9b73c36b7f99a224327e568ebab3084b64664678"></a>

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

<a id="canonical-8c02e0a2edfa1d31da1a22e0cfaab6a4bcd07debff4ddafc342a3313e1fff350"></a>

## Direct properties — cookie_params.auth_hmac.prim_key.clear_secret_info / ccd7cccbe23f / 3

<a id="canonical-6ddd685a3d062e8509c03bc318fa182fecaf6fd1f1067d48cfcce8c531df1c04"></a>

<a id="canonical-26cc2e0c1095c2bbba3d3804bc0483aa6a450a9117d679e2368ac6861e1a783b"></a>

## provider_ref property — cookie_params.auth_hmac.prim_key.clear_secret_info / ccd7cccbe23f / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-f77515edb7ae55202367f2dc24fe947e5106a3952b1d661c7634442c4aa524bd"></a>

<a id="canonical-14899cc45d879b6b73a838aa6d900b9a89177592c9c315aeb761cbdd1e813ce7"></a>

## url property — cookie_params.auth_hmac.prim_key.clear_secret_info / ccd7cccbe23f / 5

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

<a id="canonical-fafccfa41ab09dbf97cc6331a3fb8483b4ce185c530da15862678760336158b4"></a>

## Next pages — cookie_params.auth_hmac.prim_key.clear_secret_info / ccd7cccbe23f / 6

- [cookie_params.auth_hmac.prim_key](resources--authentication--reference--group-001.md#canonical-ce65bd0d2e6dff20c84e468e4817de3bf0d18e8eb9178e12326f09cb69c14c6b)
- [xcsh_authentication](../resources/authentication.md#canonical-417d73493652f5a1771ef46f280f23ffd5c91667377a08f64262e5ac9ba40724)

<a id="canonical-3cbabb805df13286253db6e6b48c8ed787a2afd6940f44f7d440b1304ee2ea0f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c037849c3fa30f0fd3425621550ad75f86282d0401b8d2ffb1c49c3486da963b"></a>

## cookie_params.auth_hmac.sec_key — cookie_params.auth_hmac.sec_key / 69fe78c6c8e1 / 2

Breadcrumbs:

- [xcsh_authentication](../resources/authentication.md#canonical-417d73493652f5a1771ef46f280f23ffd5c91667377a08f64262e5ac9ba40724)
- [Property reference](resources--authentication--reference--group-001.md#canonical-c4df54315d421dd70758e3db0bb63b6e69f2aee4a7ec3a11548b735bedc072c3)
- [cookie_params](resources--authentication--reference--group-001.md#canonical-388b68d284e492b110f52b383cfd53ce4a47a7ae27b24501356f1b1a4c640894)
- [cookie_params.auth_hmac](resources--authentication--reference--group-001.md#canonical-4ccf76209bd9ae38fb2902738a78c7f7dc8594989e9d18ad6c4a9e56adb4f931)
- cookie_params.auth_hmac.sec_key

<a id="canonical-c30b00e0dd7ae068303e58c0646e9f9959296867119c7c25a5b59b85e01ecd65"></a>

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
sec_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-3f0fe3b4a8414b7759fe81f9223c61c1178d5b19df152a25335b08cbdfc3ab62"></a>

## Direct properties — cookie_params.auth_hmac.sec_key / 69fe78c6c8e1 / 3

- [blindfold_secret_info](resources--authentication--reference--group-001.md#canonical-ab77b47d3b05d492038f3c593714653d1c9367ab7696f86482f826ad3a6b3389): complete subsection reference.

- [clear_secret_info](resources--authentication--reference--group-001.md#canonical-6cb2bbe741c4b080be15c278f5be1bacf4fa48d8d227a8d9b724185baa5f445f): complete subsection reference.

<a id="canonical-4ad8f8f7fe3bbc7bf1b9875b802cb7b70c04761f7da2ca1590829946a8f6af9b"></a>

## Next pages — cookie_params.auth_hmac.sec_key / 69fe78c6c8e1 / 4

- [cookie_params.auth_hmac.sec_key.blindfold_secret_info](resources--authentication--reference--group-001.md#canonical-ab77b47d3b05d492038f3c593714653d1c9367ab7696f86482f826ad3a6b3389)
- [cookie_params.auth_hmac.sec_key.clear_secret_info](resources--authentication--reference--group-001.md#canonical-6cb2bbe741c4b080be15c278f5be1bacf4fa48d8d227a8d9b724185baa5f445f)
- [cookie_params.auth_hmac](resources--authentication--reference--group-001.md#canonical-4ccf76209bd9ae38fb2902738a78c7f7dc8594989e9d18ad6c4a9e56adb4f931)
- [xcsh_authentication](../resources/authentication.md#canonical-417d73493652f5a1771ef46f280f23ffd5c91667377a08f64262e5ac9ba40724)

<a id="canonical-ab77b47d3b05d492038f3c593714653d1c9367ab7696f86482f826ad3a6b3389"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-df082952303418d2e8573328424a494871286db672401ab4c03f9b458da7b6a8"></a>

## cookie_params.auth_hmac.sec_key.blindfold_secret_info — cookie_params.auth_hmac.sec_key.blindfold_secret_info / d5350f1d1707 / 2

Breadcrumbs:

- [xcsh_authentication](../resources/authentication.md#canonical-417d73493652f5a1771ef46f280f23ffd5c91667377a08f64262e5ac9ba40724)
- [Property reference](resources--authentication--reference--group-001.md#canonical-c4df54315d421dd70758e3db0bb63b6e69f2aee4a7ec3a11548b735bedc072c3)
- [cookie_params](resources--authentication--reference--group-001.md#canonical-388b68d284e492b110f52b383cfd53ce4a47a7ae27b24501356f1b1a4c640894)
- [cookie_params.auth_hmac](resources--authentication--reference--group-001.md#canonical-4ccf76209bd9ae38fb2902738a78c7f7dc8594989e9d18ad6c4a9e56adb4f931)
- [cookie_params.auth_hmac.sec_key](resources--authentication--reference--group-001.md#canonical-3cbabb805df13286253db6e6b48c8ed787a2afd6940f44f7d440b1304ee2ea0f)
- cookie_params.auth_hmac.sec_key.blindfold_secret_info

<a id="canonical-35b3a8e73b9af5b15ecb01d1f54739e76209d88ec8ddd3da1d9be24d5387756b"></a>

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

<a id="canonical-38575dd64b9971a7a3edc1ee324dc46a3dadfa2c2ad0b7f362207a31a6370488"></a>

## Direct properties — cookie_params.auth_hmac.sec_key.blindfold_secret_info / d5350f1d1707 / 3

<a id="canonical-fc1871bce4f1be36c35fa619fa12f209d51f9100953dd6794f441a1b2d37effb"></a>

<a id="canonical-2e5293b29a067a1daa38715f92eecd65d7c9fba43582285dce717e7e8d734568"></a>

## decryption_provider property — cookie_params.auth_hmac.sec_key.blindfold_secret_info / d5350f1d1707 / 4

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

<a id="canonical-a835ed14ae044385eb23c2e9cbc44a6edc7f8b5d9d9f3cc3cb284e81dac833cf"></a>

<a id="canonical-882d488ef803960f651a29cfdfcaf7c14e37a611c6c040af38972f6b187de9d2"></a>

## location property — cookie_params.auth_hmac.sec_key.blindfold_secret_info / d5350f1d1707 / 5

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

<a id="canonical-6e2113e2be731a69972c9167a5afb82dae8058134d0cc14a31dcf50112d4e361"></a>

<a id="canonical-f6ef952ddd82c33ae56ae866673a92c85691d498cbd4b8a8dc36976efb21fad6"></a>

## store_provider property — cookie_params.auth_hmac.sec_key.blindfold_secret_info / d5350f1d1707 / 6

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

<a id="canonical-9bfe8e2cc5490ba7f203afb0bb521e5b2fd9fd6eeba942bec2064edc7efc8284"></a>

## Next pages — cookie_params.auth_hmac.sec_key.blindfold_secret_info / d5350f1d1707 / 7

- [cookie_params.auth_hmac.sec_key](resources--authentication--reference--group-001.md#canonical-3cbabb805df13286253db6e6b48c8ed787a2afd6940f44f7d440b1304ee2ea0f)
- [xcsh_authentication](../resources/authentication.md#canonical-417d73493652f5a1771ef46f280f23ffd5c91667377a08f64262e5ac9ba40724)

<a id="canonical-6cb2bbe741c4b080be15c278f5be1bacf4fa48d8d227a8d9b724185baa5f445f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a0031dba076060b5f4f7ec58d8502bc3c913cef360e32fe9941fc0dbd36529f2"></a>

## cookie_params.auth_hmac.sec_key.clear_secret_info — cookie_params.auth_hmac.sec_key.clear_secret_info / de0d5a90f202 / 2

Breadcrumbs:

- [xcsh_authentication](../resources/authentication.md#canonical-417d73493652f5a1771ef46f280f23ffd5c91667377a08f64262e5ac9ba40724)
- [Property reference](resources--authentication--reference--group-001.md#canonical-c4df54315d421dd70758e3db0bb63b6e69f2aee4a7ec3a11548b735bedc072c3)
- [cookie_params](resources--authentication--reference--group-001.md#canonical-388b68d284e492b110f52b383cfd53ce4a47a7ae27b24501356f1b1a4c640894)
- [cookie_params.auth_hmac](resources--authentication--reference--group-001.md#canonical-4ccf76209bd9ae38fb2902738a78c7f7dc8594989e9d18ad6c4a9e56adb4f931)
- [cookie_params.auth_hmac.sec_key](resources--authentication--reference--group-001.md#canonical-3cbabb805df13286253db6e6b48c8ed787a2afd6940f44f7d440b1304ee2ea0f)
- cookie_params.auth_hmac.sec_key.clear_secret_info

<a id="canonical-dfe02a41712801ca962b88cffd6ce94c65fadfcc734adbf8ce504abab7a1cfac"></a>

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

<a id="canonical-078489912819ef27649c25e32d52d8a0fa2b36d45d6801c153ef1075b6bf8b3e"></a>

## Direct properties — cookie_params.auth_hmac.sec_key.clear_secret_info / de0d5a90f202 / 3

<a id="canonical-b8ab64af16b7666b9e9d09d503c709b9a9cf47a36cbce11f99db043417e1f55b"></a>

<a id="canonical-a0f185bf5c640fe6efd1e6d0ab9f1e0b245924a0c386dc900fd1a1ec6836f8c4"></a>

## provider_ref property — cookie_params.auth_hmac.sec_key.clear_secret_info / de0d5a90f202 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-9cf40f4085dae15ed47aa0b55c436b832b63dea8433570b875bcc06a20065796"></a>

<a id="canonical-9f2378aa7fcff98f17357b531e336ef079d9aebd54753b13db9dfac6f8fb0c84"></a>

## url property — cookie_params.auth_hmac.sec_key.clear_secret_info / de0d5a90f202 / 5

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

<a id="canonical-6e4a55d1ea4379b50dbe12663bd3510c87411d82109e8aad54ebf5ccd61442c5"></a>

## Next pages — cookie_params.auth_hmac.sec_key.clear_secret_info / de0d5a90f202 / 6

- [cookie_params.auth_hmac.sec_key](resources--authentication--reference--group-001.md#canonical-3cbabb805df13286253db6e6b48c8ed787a2afd6940f44f7d440b1304ee2ea0f)
- [xcsh_authentication](../resources/authentication.md#canonical-417d73493652f5a1771ef46f280f23ffd5c91667377a08f64262e5ac9ba40724)

<a id="canonical-839a32944466cddbe8406764bc6097e2c3a9f11b665c9c2895a455405ecf45c4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d38018e18729e15ebb30d5e6ea7a57ee70df46ad1a04a5300dc47a9c51247fc0"></a>

## cookie_params.kms_key_hmac — cookie_params.kms_key_hmac / 45ced01c2d6f / 2

Breadcrumbs:

- [xcsh_authentication](../resources/authentication.md#canonical-417d73493652f5a1771ef46f280f23ffd5c91667377a08f64262e5ac9ba40724)
- [Property reference](resources--authentication--reference--group-001.md#canonical-c4df54315d421dd70758e3db0bb63b6e69f2aee4a7ec3a11548b735bedc072c3)
- [cookie_params](resources--authentication--reference--group-001.md#canonical-388b68d284e492b110f52b383cfd53ce4a47a7ae27b24501356f1b1a4c640894)
- cookie_params.kms_key_hmac

<a id="canonical-c0f75baed8f4e4b8c48d2310b75ed97a11af3a6989763fedca4d2b3379692b5e"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for kms key hmac.

Upstream description:

Reference to KMS Key Object.

Receipt-pinned upstream constraints:

```json
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
kms_key_hmac = {}
```

<a id="canonical-c4d1c13d2b6610abc5e6a586d9b46c171de8dfd6ffe4540a03c9eae5ec403871"></a>

## Direct properties — cookie_params.kms_key_hmac / 45ced01c2d6f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fc8d297a742a485256bcf4e054dffd444099478fa3455a4969c27be22b7ebdd5"></a>

## Next pages — cookie_params.kms_key_hmac / 45ced01c2d6f / 4

- [cookie_params](resources--authentication--reference--group-001.md#canonical-388b68d284e492b110f52b383cfd53ce4a47a7ae27b24501356f1b1a4c640894)
- [xcsh_authentication](../resources/authentication.md#canonical-417d73493652f5a1771ef46f280f23ffd5c91667377a08f64262e5ac9ba40724)

<a id="canonical-2e567deb2f14580d6e05093b8e3a34311034e624bcff1f38e036f6d78ebe3faa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6d232ce87693c9c8028afd7ec80fec96cdd88bb0a206b9091a680a4d76fa72a5"></a>

## oidc_auth — oidc_auth / 1604e08b5037 / 2

Breadcrumbs:

- [xcsh_authentication](../resources/authentication.md#canonical-417d73493652f5a1771ef46f280f23ffd5c91667377a08f64262e5ac9ba40724)
- [Property reference](resources--authentication--reference--group-001.md#canonical-c4df54315d421dd70758e3db0bb63b6e69f2aee4a7ec3a11548b735bedc072c3)
- oidc_auth

<a id="canonical-48e7a126ba16299db6a46ba8a0438e02329942fcf8bdbc4434f5d0c9b23a2872"></a>

Type: `"object"`. single nested block, Optional.

OIDCAuthType.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("oidc_client_id"),
  validators.ConflictingObjectAttributes("oidc_auth_params",
    "oidc_well_known_config_url")}
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
  "x-ves-oneof-field-auth_params_choice": "[\"oidc_auth_params\",\"oidc_well_known_config_url\"]"
}
```

Terraform syntax:

```terraform
oidc_auth {
  # Configure direct properties listed below.
}
```

<a id="canonical-d237f52dedea0d7fe2d7897453e4f4e697deb24450c95ffb88f573aa2efecb56"></a>

## Direct properties — oidc_auth / 1604e08b5037 / 3

- [client_secret](resources--authentication--reference--group-001.md#canonical-78b1c4f283db96d1d7e8764c81017d9f54977729a4e58e2582299daebf6d7067): complete subsection reference.

- [oidc_auth_params](resources--authentication--reference--group-001.md#canonical-9997e4fc1e5cf5b720e5199a0a3a43ac96def4ba91b3b31e8ce53fb67b78086e): complete subsection reference.

<a id="canonical-0dc6ede03efb3f0597cd6aef018f246b1b5edd819770fb315a0ceb0f8278bfac"></a>

<a id="canonical-84c4705f3a56e7798e184a0da1312f946a9b62d72f172c14035885c30c81ceb0"></a>

## oidc_client_id property — oidc_auth / 1604e08b5037 / 4

Type: `"string"`. Optional.

Client ID used while sending the Authorization Request to OIDC server.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-722122c8ee0935f80128a13ec05761e74d8aca5143cbb04539fc24f3d0b56e58"></a>

<a id="canonical-4c31f6cfef805a1699e56cca8e1717dbb6273b41e0a6dc4f09822b644ea73d20"></a>

## oidc_well_known_config_url property — oidc_auth / 1604e08b5037 / 5

Type: `"string"`. Optional.

Exclusive with \[oidc\_auth\_params\] An OIDC well-known configuration URL that will be used to
fetch authentication related endpoints.

Upstream description:

Exclusive with \[oidc\_auth\_params\] An OIDC well-known configuration URL that will be used to
fetch authentication related endpoints.

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
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-289e70c96e45409ad514faafdde100f3bd3f3f3dee45cfa96eb159dd78f682cf"></a>

## Next pages — oidc_auth / 1604e08b5037 / 6

- [oidc_auth.client_secret](resources--authentication--reference--group-001.md#canonical-78b1c4f283db96d1d7e8764c81017d9f54977729a4e58e2582299daebf6d7067)
- [oidc_auth.oidc_auth_params](resources--authentication--reference--group-001.md#canonical-9997e4fc1e5cf5b720e5199a0a3a43ac96def4ba91b3b31e8ce53fb67b78086e)
- [Property reference](resources--authentication--reference--group-001.md#canonical-c4df54315d421dd70758e3db0bb63b6e69f2aee4a7ec3a11548b735bedc072c3)
- [xcsh_authentication](../resources/authentication.md#canonical-417d73493652f5a1771ef46f280f23ffd5c91667377a08f64262e5ac9ba40724)

<a id="canonical-78b1c4f283db96d1d7e8764c81017d9f54977729a4e58e2582299daebf6d7067"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d44c4269f67ddd2b372b49d5958d7213a6fbe400ff927cdac9228e142182fa9f"></a>

## oidc_auth.client_secret — oidc_auth.client_secret / d26ecbb62f38 / 2

Breadcrumbs:

- [xcsh_authentication](../resources/authentication.md#canonical-417d73493652f5a1771ef46f280f23ffd5c91667377a08f64262e5ac9ba40724)
- [Property reference](resources--authentication--reference--group-001.md#canonical-c4df54315d421dd70758e3db0bb63b6e69f2aee4a7ec3a11548b735bedc072c3)
- [oidc_auth](resources--authentication--reference--group-001.md#canonical-2e567deb2f14580d6e05093b8e3a34311034e624bcff1f38e036f6d78ebe3faa)
- oidc_auth.client_secret

<a id="canonical-f1471e4b85507fe4e773849820ec63d5815e9c6d81e4c0f01344352a51e080a9"></a>

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
client_secret {
  # Configure direct properties listed below.
}
```

<a id="canonical-c45251a1342726ad164cc228f4d79f2db4d7dc4ee3d43e6727b670c119719619"></a>

## Direct properties — oidc_auth.client_secret / d26ecbb62f38 / 3

- [blindfold_secret_info](resources--authentication--reference--group-001.md#canonical-e7ad5cade66fe517e3eec794084e47e358762b0095123dcb9258e54d562fd131): complete subsection reference.

- [clear_secret_info](resources--authentication--reference--group-001.md#canonical-35ed25c7c32c3271e612bf2a974be888e9ebfe583023a379e35b5fdc3a14fd79): complete subsection reference.

<a id="canonical-333b709ae070f6aa6cb0295ae1a27ae141a8e02fedd9f9fb19a0b049c2908305"></a>

## Next pages — oidc_auth.client_secret / d26ecbb62f38 / 4

- [oidc_auth.client_secret.blindfold_secret_info](resources--authentication--reference--group-001.md#canonical-e7ad5cade66fe517e3eec794084e47e358762b0095123dcb9258e54d562fd131)
- [oidc_auth.client_secret.clear_secret_info](resources--authentication--reference--group-001.md#canonical-35ed25c7c32c3271e612bf2a974be888e9ebfe583023a379e35b5fdc3a14fd79)
- [oidc_auth](resources--authentication--reference--group-001.md#canonical-2e567deb2f14580d6e05093b8e3a34311034e624bcff1f38e036f6d78ebe3faa)
- [xcsh_authentication](../resources/authentication.md#canonical-417d73493652f5a1771ef46f280f23ffd5c91667377a08f64262e5ac9ba40724)

<a id="canonical-e7ad5cade66fe517e3eec794084e47e358762b0095123dcb9258e54d562fd131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9fe975888adf5855f01a89f81a73bb89c00d84050b19f71a6f3765c570b26c79"></a>

## oidc_auth.client_secret.blindfold_secret_info — oidc_auth.client_secret.blindfold_secret_info / 572d03e0992b / 2

Breadcrumbs:

- [xcsh_authentication](../resources/authentication.md#canonical-417d73493652f5a1771ef46f280f23ffd5c91667377a08f64262e5ac9ba40724)
- [Property reference](resources--authentication--reference--group-001.md#canonical-c4df54315d421dd70758e3db0bb63b6e69f2aee4a7ec3a11548b735bedc072c3)
- [oidc_auth](resources--authentication--reference--group-001.md#canonical-2e567deb2f14580d6e05093b8e3a34311034e624bcff1f38e036f6d78ebe3faa)
- [oidc_auth.client_secret](resources--authentication--reference--group-001.md#canonical-78b1c4f283db96d1d7e8764c81017d9f54977729a4e58e2582299daebf6d7067)
- oidc_auth.client_secret.blindfold_secret_info

<a id="canonical-d6acc683f52976c946e4f4217c905aa53ae573ea25c5345d81c9ccb4bc6e15c5"></a>

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

<a id="canonical-7f2ac999af408aca05761d11e752688699ea60b716cf0cbed33dd2d2a30a9ddb"></a>

## Direct properties — oidc_auth.client_secret.blindfold_secret_info / 572d03e0992b / 3

<a id="canonical-ce025cc8d5aa1c67d4935b431b89a33b85b6f7b7486c52f9fef7efd7ecda5c6a"></a>

<a id="canonical-23e0ea2e87f6717f1af4d45273a8a37cb7ba40b8696e2e01e339acd84c625b3f"></a>

## decryption_provider property — oidc_auth.client_secret.blindfold_secret_info / 572d03e0992b / 4

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

<a id="canonical-d3f81d3eb539972689ec8ee5bfea7199097116d76edcfa7fe1d6274faca818c4"></a>

<a id="canonical-deafc102ad6660aef216f499101260bc46979a6e3ef6c68a0ba3c56268ff5c28"></a>

## location property — oidc_auth.client_secret.blindfold_secret_info / 572d03e0992b / 5

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

<a id="canonical-7fe64fe54c5e2e71af85ab8d1ff6d4541d9dd6f6afd87c5386df90ea56fd9960"></a>

<a id="canonical-946d1319674aa949b659d71ea370b33722f27868f199dc7cd6bd2193e3933560"></a>

## store_provider property — oidc_auth.client_secret.blindfold_secret_info / 572d03e0992b / 6

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

<a id="canonical-4ac07a300eeff75230a30a217fe1ef6da07fd669b04ac5e2a487714ac21091d8"></a>

## Next pages — oidc_auth.client_secret.blindfold_secret_info / 572d03e0992b / 7

- [oidc_auth.client_secret](resources--authentication--reference--group-001.md#canonical-78b1c4f283db96d1d7e8764c81017d9f54977729a4e58e2582299daebf6d7067)
- [xcsh_authentication](../resources/authentication.md#canonical-417d73493652f5a1771ef46f280f23ffd5c91667377a08f64262e5ac9ba40724)

<a id="canonical-35ed25c7c32c3271e612bf2a974be888e9ebfe583023a379e35b5fdc3a14fd79"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cb87ccccd07dd2b4901a00581d4ab45ebe679d7fe4044ef43a54eda3caad0d76"></a>

## oidc_auth.client_secret.clear_secret_info — oidc_auth.client_secret.clear_secret_info / d47efb676d8a / 2

Breadcrumbs:

- [xcsh_authentication](../resources/authentication.md#canonical-417d73493652f5a1771ef46f280f23ffd5c91667377a08f64262e5ac9ba40724)
- [Property reference](resources--authentication--reference--group-001.md#canonical-c4df54315d421dd70758e3db0bb63b6e69f2aee4a7ec3a11548b735bedc072c3)
- [oidc_auth](resources--authentication--reference--group-001.md#canonical-2e567deb2f14580d6e05093b8e3a34311034e624bcff1f38e036f6d78ebe3faa)
- [oidc_auth.client_secret](resources--authentication--reference--group-001.md#canonical-78b1c4f283db96d1d7e8764c81017d9f54977729a4e58e2582299daebf6d7067)
- oidc_auth.client_secret.clear_secret_info

<a id="canonical-e3a89c2b38fc576de42abc97a929fded33d64337edef90e2e0b996bc6c8e7d3f"></a>

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

<a id="canonical-15f89134be0955650b3c598068b8ca49e31169ddaec045ade355bfbc1469277c"></a>

## Direct properties — oidc_auth.client_secret.clear_secret_info / d47efb676d8a / 3

<a id="canonical-28269df732c34025b664bec2622c58d4662d9e50f9454d85b19ca3d7f5504bbd"></a>

<a id="canonical-5509d788adf032c82b2a6798c993eda05f6dd3211cbd01f35ec95c6a7f9649d0"></a>

## provider_ref property — oidc_auth.client_secret.clear_secret_info / d47efb676d8a / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-8a36e6c67b1aeea91b42cf901f8f96077afde596cd79650bb624320462aac533"></a>

<a id="canonical-447645052a365637b13555f42875680fc9079b5bab32e195c7ef0274defd8ab7"></a>

## url property — oidc_auth.client_secret.clear_secret_info / d47efb676d8a / 5

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

<a id="canonical-09b2be93823e670afdc2e130c4e868ce0216335dfc759754b2792dcca9e376e7"></a>

## Next pages — oidc_auth.client_secret.clear_secret_info / d47efb676d8a / 6

- [oidc_auth.client_secret](resources--authentication--reference--group-001.md#canonical-78b1c4f283db96d1d7e8764c81017d9f54977729a4e58e2582299daebf6d7067)
- [xcsh_authentication](../resources/authentication.md#canonical-417d73493652f5a1771ef46f280f23ffd5c91667377a08f64262e5ac9ba40724)

<a id="canonical-9997e4fc1e5cf5b720e5199a0a3a43ac96def4ba91b3b31e8ce53fb67b78086e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5a253aef405659317c407fe20143138bfdb6b18d31962e6e874598b73bc3f30d"></a>

## oidc_auth.oidc_auth_params — oidc_auth.oidc_auth_params / 5356ef6bd544 / 2

Breadcrumbs:

- [xcsh_authentication](../resources/authentication.md#canonical-417d73493652f5a1771ef46f280f23ffd5c91667377a08f64262e5ac9ba40724)
- [Property reference](resources--authentication--reference--group-001.md#canonical-c4df54315d421dd70758e3db0bb63b6e69f2aee4a7ec3a11548b735bedc072c3)
- [oidc_auth](resources--authentication--reference--group-001.md#canonical-2e567deb2f14580d6e05093b8e3a34311034e624bcff1f38e036f6d78ebe3faa)
- oidc_auth.oidc_auth_params

<a id="canonical-f3fd11a1cf7a6c471b0531bff5f0076cd92adcb337cbd76f87c8d0c05bde8d3b"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for oidc auth params.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("auth_endpoint_url",
    "end_session_endpoint_url",
    "token_endpoint_url")}
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
oidc_auth_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-defa61a1cef05c74acd92bc393efdf6ca1013895c6c594d2e145fb041715d65f"></a>

## Direct properties — oidc_auth.oidc_auth_params / 5356ef6bd544 / 3

<a id="canonical-26544726cc764047b66623bc4664595208866df67cff71ded28202077f93978f"></a>

<a id="canonical-5c79e8e8371c7502af94e2e8cb539342efa495f33dcd16da757e6b9362369480"></a>

## auth_endpoint_url property — oidc_auth.oidc_auth_params / 5356ef6bd544 / 4

Type: `"string"`. Optional.

URL of the authorization server's authorization endpoint.

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
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
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
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-3a04cfda7585adcfe708377cd81ca88951d2d2fea7017ea47bfaf09c9695e89f"></a>

<a id="canonical-8953aacb296225844da06119ad53eac520e3e874e61a02bac9f0c94e850bcd32"></a>

## end_session_endpoint_url property — oidc_auth.oidc_auth_params / 5356ef6bd544 / 5

Type: `"string"`. Optional.

URL of the authorization server's Logout endpoint.

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
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
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
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-51057b65909a42872849ef6e2f4229bbae9af8cde3f2e221a69e46483cf844aa"></a>

<a id="canonical-22f0baf3e3376ac1c1a72bbf9869ca373399538ade6df3bf8f8e9bdf507c0350"></a>

## token_endpoint_url property — oidc_auth.oidc_auth_params / 5356ef6bd544 / 6

Type: `"string"`. Optional.

URL of the authorization server's Token endpoint.

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
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
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
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-faf0159ce9ef21f0bfbb5d625db8cd124a9d04ca0f3577a4e2cb8a0a80b9629a"></a>

## Next pages — oidc_auth.oidc_auth_params / 5356ef6bd544 / 7

- [oidc_auth](resources--authentication--reference--group-001.md#canonical-2e567deb2f14580d6e05093b8e3a34311034e624bcff1f38e036f6d78ebe3faa)
- [xcsh_authentication](../resources/authentication.md#canonical-417d73493652f5a1771ef46f280f23ffd5c91667377a08f64262e5ac9ba40724)

<a id="canonical-f3d9b4a248b762990240bfd0ec12afe3d19dc70df6b11e56bb750677650dbc0c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5449ae1cafb8c8891ba175cedf17c93b2338d4a9011b144fbb06afaccbac5346"></a>

## timeouts — timeouts / 65cb8a90953f / 2

Breadcrumbs:

- [xcsh_authentication](../resources/authentication.md#canonical-417d73493652f5a1771ef46f280f23ffd5c91667377a08f64262e5ac9ba40724)
- [Property reference](resources--authentication--reference--group-001.md#canonical-c4df54315d421dd70758e3db0bb63b6e69f2aee4a7ec3a11548b735bedc072c3)
- timeouts

<a id="canonical-76b7c92e264865a56bfa33ebebf977419bd5694207e59e7448b3735977fd3966"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-8f0ff3db5d95b013bdadda79d2bfed1ad179491fffdf568c5d9796539731152b"></a>

## Direct properties — timeouts / 65cb8a90953f / 3

<a id="canonical-f97bf9951fea525c16a6585a973e19d61636ca409031ce84f4b36e86b37fbaa7"></a>

<a id="canonical-f0999fcd7f4b46c9523db574266fc26c28a92bd2f4cf84fd4c0031c8cc8504d6"></a>

## create property — timeouts / 65cb8a90953f / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0665dc32e2090665888270c7aeec01404783fca42d7711598737fd1ba03a5a2e"></a>

<a id="canonical-d2f8ffd75efa6ffa369cec1cb3d845af2bc668c01a23451d8e8ae3a85010b8f8"></a>

## delete property — timeouts / 65cb8a90953f / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-6f2dad3306717c3c6fe5076a39303565cef6319fe214614a085ea33889c9f3df"></a>

<a id="canonical-298827308b894d5449d358e3c4543b7a5917671901a9403817cda75dabc8f8e2"></a>

## read property — timeouts / 65cb8a90953f / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-3875566ac7dec4595a1a72a336d0af9ab6d4ae28cdbd2a5714ffeadd1c018eae"></a>

<a id="canonical-3d6a23cd4ca675f5c71b6bcd3c866789318613bb6d05601dc7c5d4585fb1551c"></a>

## update property — timeouts / 65cb8a90953f / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-f16f499f2fb89ec2fce78b2dc9e32a8e367049cacf839ef7f0937503af5c8a1d"></a>

## Next pages — timeouts / 65cb8a90953f / 8

- [Property reference](resources--authentication--reference--group-001.md#canonical-c4df54315d421dd70758e3db0bb63b6e69f2aee4a7ec3a11548b735bedc072c3)
- [xcsh_authentication](../resources/authentication.md#canonical-417d73493652f5a1771ef46f280f23ffd5c91667377a08f64262e5ac9ba40724)
