---
page_title: "xcsh_nfv_service reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nfv_service reference."
---

# xcsh_nfv_service reference

<a id="canonical-f650e598db143c54deea47db9bef16a837d3b3bd1f54b5583e6985e032da5990"></a>

## Next pages — palo_alto_fw_service.pan_ami_bundle2 / 93eea51f1fe6 / 4

- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-1ca81a9ddb9348c620d239771df702d0b475bf02453eb0a2a936c6227d30b60c)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-e1a6b7b002ad6cbdb482c5e45ad67e792c6a1ccc39ae457d5fda0a4476aba660"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-05029db969bdab9d0105354aec45849f97e3f208ab67612f1825d05dc777efcf"></a>

## palo_alto_fw_service.panorama_server — palo_alto_fw_service.panorama_server / 668e963258c5 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-1ca81a9ddb9348c620d239771df702d0b475bf02453eb0a2a936c6227d30b60c)
- palo_alto_fw_service.panorama_server

<a id="canonical-f6db7a352ae40b4857608d9a0b3f02f5d3dd93aad8cc38d39ff38cb070980b26"></a>

Type: `"single"`. Computed.

Configuration parameter for panorama server.

Upstream description:

Panorama Server Type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-f66a315e72d736c84992006cf20276a666dbe23d08ca50e3389e192c1607692e"></a>

## Direct properties — palo_alto_fw_service.panorama_server / 668e963258c5 / 3

- [authorization_key](data-sources--nfv_service--reference--group-004.md#canonical-0242fb55ed58e4071a901b27c643692ee4589f22df6ddb64f6b2475ec6748759): complete subsection reference.

<a id="canonical-e07cc48c251ef5305a6604222a929ad16ed2e9b80b82def33af86eed77d819ad"></a>

<a id="canonical-f85da6f31c56ae5ef23f1ec9d93237e7af714d6eec9ff91fb81cf54023d2a63c"></a>

## device_group_name property — palo_alto_fw_service.panorama_server / 668e963258c5 / 4

Type: `"string"`. Computed.

Device Group Name. Device Group Name.

Upstream description:

Device Group Name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-eb9b8a4c7f21bc8b114410f1e5266e3d0f5ebea5347241faf2a11a340d6262ed"></a>

<a id="canonical-728c36189601010e2f26fdde5cfd49b7b6713f96ddb205e06f42af921ed2fdd0"></a>

## server property — palo_alto_fw_service.panorama_server / 668e963258c5 / 5

Type: `"string"`. Computed.

Panorama Server Address to which the firewall should connect to.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-fb367e3c5894cdd2018afde322234da88c0c378ecf826e2b28fa1b78116ad5fc"></a>

<a id="canonical-fbd160840f65cdf881a206f9f5af428fea19f07a649f752b2f0af4ac34feec86"></a>

## template_stack_name property — palo_alto_fw_service.panorama_server / 668e963258c5 / 6

Type: `"string"`. Computed.

Template stack name. Template Stack Name.

Upstream description:

Template Stack Name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-5d874a4626be77e4c570f19e29280a4ec6b5c17383a7bb51f7bcab1298311496"></a>

## Next pages — palo_alto_fw_service.panorama_server / 668e963258c5 / 7

- [palo_alto_fw_service.panorama_server.authorization_key](data-sources--nfv_service--reference--group-004.md#canonical-0242fb55ed58e4071a901b27c643692ee4589f22df6ddb64f6b2475ec6748759)
- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-1ca81a9ddb9348c620d239771df702d0b475bf02453eb0a2a936c6227d30b60c)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-0242fb55ed58e4071a901b27c643692ee4589f22df6ddb64f6b2475ec6748759"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d130604d057fd79b7e851d9256d629ac12d5efb035a04ce0bf317f9e364343db"></a>

## palo_alto_fw_service.panorama_server.authorization_key — palo_alto_fw_service.panorama_server.authorization_key / 815eb934b7e9 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-1ca81a9ddb9348c620d239771df702d0b475bf02453eb0a2a936c6227d30b60c)
- [palo_alto_fw_service.panorama_server](data-sources--nfv_service--reference--group-004.md#canonical-e1a6b7b002ad6cbdb482c5e45ad67e792c6a1ccc39ae457d5fda0a4476aba660)
- palo_alto_fw_service.panorama_server.authorization_key

<a id="canonical-7087b7a61ce087440641a34f89da34388592b3099522bef5504188a347e41833"></a>

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

<a id="canonical-50d9ee23e6c8fa55d0815509eb5e1031968a3424941c2bf3f14c02d0ea38dcaf"></a>

## Direct properties — palo_alto_fw_service.panorama_server.authorization_key / 815eb934b7e9 / 3

- [blindfold_secret_info](data-sources--nfv_service--reference--group-004.md#canonical-597186743dbc7a7f0310a087d1546ed1993150c4bdaa173e06638cbcad4c3edb): complete subsection reference.

- [clear_secret_info](data-sources--nfv_service--reference--group-004.md#canonical-2f41bcdcfc9827138cf41946398e77000a55bb081ef016de212629a820e73b43): complete subsection reference.

<a id="canonical-c210fc8f57628cbe055a619405d06c6b29129ebed9b86017a2887b79e9147d73"></a>

## Next pages — palo_alto_fw_service.panorama_server.authorization_key / 815eb934b7e9 / 4

- [palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info](data-sources--nfv_service--reference--group-004.md#canonical-597186743dbc7a7f0310a087d1546ed1993150c4bdaa173e06638cbcad4c3edb)
- [palo_alto_fw_service.panorama_server.authorization_key.clear_secret_info](data-sources--nfv_service--reference--group-004.md#canonical-2f41bcdcfc9827138cf41946398e77000a55bb081ef016de212629a820e73b43)
- [palo_alto_fw_service.panorama_server](data-sources--nfv_service--reference--group-004.md#canonical-e1a6b7b002ad6cbdb482c5e45ad67e792c6a1ccc39ae457d5fda0a4476aba660)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-597186743dbc7a7f0310a087d1546ed1993150c4bdaa173e06638cbcad4c3edb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5870fe97c1b5b76047e1708b1a73142a91b28ac0cadcbe48a16f7644ba9e8d63"></a>

## palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info — palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info / 1e8c6937b5e9 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-1ca81a9ddb9348c620d239771df702d0b475bf02453eb0a2a936c6227d30b60c)
- [palo_alto_fw_service.panorama_server](data-sources--nfv_service--reference--group-004.md#canonical-e1a6b7b002ad6cbdb482c5e45ad67e792c6a1ccc39ae457d5fda0a4476aba660)
- [palo_alto_fw_service.panorama_server.authorization_key](data-sources--nfv_service--reference--group-004.md#canonical-0242fb55ed58e4071a901b27c643692ee4589f22df6ddb64f6b2475ec6748759)
- palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info

<a id="canonical-ce5ddadb96c4e802f9b962a466b0dfba8b8fecbbfb6738fbaaceb71fbb5a51b4"></a>

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

<a id="canonical-1456b5b69c8c8cdfed36e5ec7e49fa1ebde306d3005a33a85dfff3b18c178b5f"></a>

## Direct properties — palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info / 1e8c6937b5e9 / 3

<a id="canonical-e5086876927f7f872a39128513b1e224af6bdf136be0a90cce9be99d336b962d"></a>

<a id="canonical-019687566eea1abddd1f1c1cd1fd53b07fa7778ce05e78e30b24ccfc1be41e7a"></a>

## decryption_provider property — palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info / 1e8c6937b5e9 / 4

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

<a id="canonical-76207d0b8f1dea416285800d5e750f151fc39035f48eda3ff837a66e401fbcc6"></a>

<a id="canonical-0a53f935666e80e50fae8c9eb13bad4f5c263b08a2e0bf5b527937617b3c87b0"></a>

## location property — palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info / 1e8c6937b5e9 / 5

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

<a id="canonical-fc7822902ef656890c511162923621130b1e96b3e7e9950c409a44e9428e4882"></a>

<a id="canonical-b70df28deddd37a9ca8d347603a338b1c631c2a1915cd4337e00b97fc4bb3158"></a>

## store_provider property — palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info / 1e8c6937b5e9 / 6

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

<a id="canonical-b6247d1449cb08ba515a83a844afdb29da2ccd28a37075f03ad7e5d43dfc53fd"></a>

## Next pages — palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info / 1e8c6937b5e9 / 7

- [palo_alto_fw_service.panorama_server.authorization_key](data-sources--nfv_service--reference--group-004.md#canonical-0242fb55ed58e4071a901b27c643692ee4589f22df6ddb64f6b2475ec6748759)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-2f41bcdcfc9827138cf41946398e77000a55bb081ef016de212629a820e73b43"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-45cafdb6028f72a90c0be41bc4b8a2a60b9dfbb05d038843e190edde9cec739b"></a>

## palo_alto_fw_service.panorama_server.authorization_key.clear_secret_info — palo_alto_fw_service.panorama_server.authorization_key.clear_secret_info / 7cf5b19865e6 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-1ca81a9ddb9348c620d239771df702d0b475bf02453eb0a2a936c6227d30b60c)
- [palo_alto_fw_service.panorama_server](data-sources--nfv_service--reference--group-004.md#canonical-e1a6b7b002ad6cbdb482c5e45ad67e792c6a1ccc39ae457d5fda0a4476aba660)
- [palo_alto_fw_service.panorama_server.authorization_key](data-sources--nfv_service--reference--group-004.md#canonical-0242fb55ed58e4071a901b27c643692ee4589f22df6ddb64f6b2475ec6748759)
- palo_alto_fw_service.panorama_server.authorization_key.clear_secret_info

<a id="canonical-fdc86e5a89e70ea638a477ef56eed52b973ae0f0ed3788ac37f0d71b952770ec"></a>

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

<a id="canonical-45143a86a4915312d2be5d232c0e2257eb7b0d2f3022f2994ddba59713b08fb6"></a>

## Direct properties — palo_alto_fw_service.panorama_server.authorization_key.clear_secret_info / 7cf5b19865e6 / 3

<a id="canonical-b8877be8dce08bc69f65ba9ede3806ef022d24c2c46146c7cc8d63fea205e6ed"></a>

<a id="canonical-744e8e4abe8b773af10fbdc7a08d06708a9c81f8a081629439447434dd6e2603"></a>

## provider_ref property — palo_alto_fw_service.panorama_server.authorization_key.clear_secret_info / 7cf5b19865e6 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-cf495fb23ac807e2ecbbd460459c9c0c9cb7c104e12de0698fdc21588384f2e7"></a>

<a id="canonical-b2392432cffcbb2fbdbe2a0193f3e2a846fc889a1f730755db67d147ed0704c6"></a>

## url property — palo_alto_fw_service.panorama_server.authorization_key.clear_secret_info / 7cf5b19865e6 / 5

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

<a id="canonical-bf8a4131deb41a6208903181050f7a61ddaa66228fdfda54955fff5f7dc59333"></a>

## Next pages — palo_alto_fw_service.panorama_server.authorization_key.clear_secret_info / 7cf5b19865e6 / 6

- [palo_alto_fw_service.panorama_server.authorization_key](data-sources--nfv_service--reference--group-004.md#canonical-0242fb55ed58e4071a901b27c643692ee4589f22df6ddb64f6b2475ec6748759)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-b844e50e5ec67e22eb6280a754515913fecc9f21e564c553eeac9822b82d6daa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-914568753fa73fe516a50c9ff5671a50441cf081fad9d7a6c2b03b975f491780"></a>

## palo_alto_fw_service.service_nodes — palo_alto_fw_service.service_nodes / 1ac910cf8f06 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-1ca81a9ddb9348c620d239771df702d0b475bf02453eb0a2a936c6227d30b60c)
- palo_alto_fw_service.service_nodes

<a id="canonical-d52ffbccbca930d63feb98fcaf26bd89750e45a621e15f19d7ac3c79bedc6949"></a>

Type: `"single"`. Computed.

Configuration parameter for service nodes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-5d3e60c616312eb5105725e7d78be384c029cfc28f3e8ce04db78aae98e23337"></a>

## Direct properties — palo_alto_fw_service.service_nodes / 1ac910cf8f06 / 3

- [nodes](data-sources--nfv_service--reference--group-004.md#canonical-0e9fccb6f1864e82f7866453377cde5ecc1ed614f32deae2680af1a76c9f64d2): complete subsection reference.

<a id="canonical-8fe8a39be87e6e3d65dc8a69ac6287a95fc12c440e63fbe6001213e43dee78da"></a>

## Next pages — palo_alto_fw_service.service_nodes / 1ac910cf8f06 / 4

- [palo_alto_fw_service.service_nodes.nodes](data-sources--nfv_service--reference--group-004.md#canonical-0e9fccb6f1864e82f7866453377cde5ecc1ed614f32deae2680af1a76c9f64d2)
- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-1ca81a9ddb9348c620d239771df702d0b475bf02453eb0a2a936c6227d30b60c)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-0e9fccb6f1864e82f7866453377cde5ecc1ed614f32deae2680af1a76c9f64d2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9d66aeead1cc79a30659c4aad3d658c75f9b0cffa21130caf4383c55dab834c0"></a>

## palo_alto_fw_service.service_nodes.nodes — palo_alto_fw_service.service_nodes.nodes / 22e7f8e547e9 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-1ca81a9ddb9348c620d239771df702d0b475bf02453eb0a2a936c6227d30b60c)
- [palo_alto_fw_service.service_nodes](data-sources--nfv_service--reference--group-004.md#canonical-b844e50e5ec67e22eb6280a754515913fecc9f21e564c553eeac9822b82d6daa)
- palo_alto_fw_service.service_nodes.nodes

<a id="canonical-89a396ded9969282bfcae6b87a1f87e828a9da251b65d7cdae4247aa1b44f2c9"></a>

Type: `"list"`. Computed.

Palo Alto Networks AZ Nodes. Configuration parameter for nodes

Upstream description:

Configuration parameter for nodes

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 2,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 2,
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
    "ves.io.schema.rules.repeated.max_items": "2",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "2",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-74970206a2fcb898dc4cc5edd282b452f2299a379f63f2476dce3c394f52f498"></a>

## Direct properties — palo_alto_fw_service.service_nodes.nodes / 22e7f8e547e9 / 3

<a id="canonical-668fcc8b1d83f3f22798be801ea32eed23b2aec8f5b429fb02d4d953cfa6e977"></a>

<a id="canonical-2e95fcbb735f0b6505d17735216ffc2175dce58eff648e4c1617c35d6888722c"></a>

## aws_az_name property — palo_alto_fw_service.service_nodes.nodes / 22e7f8e547e9 / 4

Type: `"string"`. Computed.

AWS availability zone, must be consistent with the selected AWS region. It is recommended that AZ is
one of the AZ for sites.

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
    "pattern": "^([a-z]{2})-([a-z0-9]{4,20})-([a-z0-9]{2})$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^([a-z]{2})-([a-z0-9]{4,20})-([a-z0-9]{2})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^([a-z]{2})-([a-z0-9]{4,20})-([a-z0-9]{2})$"
  }
}
```

- [mgmt_subnet](data-sources--nfv_service--reference--group-004.md#canonical-39cf7caf789527f8242b1322147d03610959914f964493e7864c963ef286b71c): complete subsection reference.

<a id="canonical-176272e380ee6430841d508dc4f18522d3c906922a6424ac82942a47e588c183"></a>

<a id="canonical-86843dc19e87881d337c91fa6dcf30a0a0a7a73b04dc9ac0425cb2dd496c7808"></a>

## node_name property — palo_alto_fw_service.service_nodes.nodes / 22e7f8e547e9 / 5

Type: `"string"`. Computed.

Node Name will be used to assign as hostname to the service.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
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

- [reserved_mgmt_subnet](data-sources--nfv_service--reference--group-004.md#canonical-1e43326f6697853495d0f3f0b524b4a424a4441f171c5830c8734b7da8284273): complete subsection reference.

<a id="canonical-bed76ab82860fcad034fbb3bde6a029ee2194ec58c9224acb8405a1754be3823"></a>

## Next pages — palo_alto_fw_service.service_nodes.nodes / 22e7f8e547e9 / 6

- [palo_alto_fw_service.service_nodes.nodes.mgmt_subnet](data-sources--nfv_service--reference--group-004.md#canonical-39cf7caf789527f8242b1322147d03610959914f964493e7864c963ef286b71c)
- [palo_alto_fw_service.service_nodes.nodes.reserved_mgmt_subnet](data-sources--nfv_service--reference--group-004.md#canonical-1e43326f6697853495d0f3f0b524b4a424a4441f171c5830c8734b7da8284273)
- [palo_alto_fw_service.service_nodes](data-sources--nfv_service--reference--group-004.md#canonical-b844e50e5ec67e22eb6280a754515913fecc9f21e564c553eeac9822b82d6daa)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-39cf7caf789527f8242b1322147d03610959914f964493e7864c963ef286b71c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-959addaa4a0d631f77fd781cbc918426599ebdc7c4f2b6ae3410e8d200d94065"></a>

## palo_alto_fw_service.service_nodes.nodes.mgmt_subnet — palo_alto_fw_service.service_nodes.nodes.mgmt_subnet / 84dbab681161 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-1ca81a9ddb9348c620d239771df702d0b475bf02453eb0a2a936c6227d30b60c)
- [palo_alto_fw_service.service_nodes](data-sources--nfv_service--reference--group-004.md#canonical-b844e50e5ec67e22eb6280a754515913fecc9f21e564c553eeac9822b82d6daa)
- [palo_alto_fw_service.service_nodes.nodes](data-sources--nfv_service--reference--group-004.md#canonical-0e9fccb6f1864e82f7866453377cde5ecc1ed614f32deae2680af1a76c9f64d2)
- palo_alto_fw_service.service_nodes.nodes.mgmt_subnet

<a id="canonical-b5e618de74f566bbb1216eafff1c9b27ba3446737df2192bab91b9651b31a07d"></a>

Type: `"single"`. Computed.

Configuration parameter for mgmt subnet.

Upstream description:

Parameters for AWS subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"existing_subnet_id\",\"subnet_param\"]"
}
```

<a id="canonical-46049989b60478e24179bef1801a0e7766b93e2343b78366847167dc467240ad"></a>

## Direct properties — palo_alto_fw_service.service_nodes.nodes.mgmt_subnet / 84dbab681161 / 3

<a id="canonical-ee97acd810ab2c844cf45f1a8574783af7d16841494c6b01c6757438cb1ae5e1"></a>

<a id="canonical-0df38e5a0dfb2f98de2e412f5531860bd8240a4a584b32b888f0cc9a34cece77"></a>

## existing_subnet_id property — palo_alto_fw_service.service_nodes.nodes.mgmt_subnet / 84dbab681161 / 4

Type: `"string"`. Computed.

Exclusive with \[subnet\_param\] Information about existing subnet ID.

Upstream description:

Exclusive with \[subnet\_param\] Information about existing subnet ID.

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
    },
    "pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

- [subnet_param](data-sources--nfv_service--reference--group-004.md#canonical-3f09e906a025178edba56865f641dc1c38b73e12f2deea8ff4c6ebefd482c796): complete subsection reference.

<a id="canonical-7552f2606522fcc81172cee8eb140f142c4abab9886c244efb874bee9b7e99f8"></a>

## Next pages — palo_alto_fw_service.service_nodes.nodes.mgmt_subnet / 84dbab681161 / 5

- [palo_alto_fw_service.service_nodes.nodes.mgmt_subnet.subnet_param](data-sources--nfv_service--reference--group-004.md#canonical-3f09e906a025178edba56865f641dc1c38b73e12f2deea8ff4c6ebefd482c796)
- [palo_alto_fw_service.service_nodes.nodes](data-sources--nfv_service--reference--group-004.md#canonical-0e9fccb6f1864e82f7866453377cde5ecc1ed614f32deae2680af1a76c9f64d2)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-3f09e906a025178edba56865f641dc1c38b73e12f2deea8ff4c6ebefd482c796"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1821fef7b63bbe0af104d1492b515d50f8d68ee4528c1b938d1cfc5eb897a67d"></a>

## palo_alto_fw_service.service_nodes.nodes.mgmt_subnet.subnet_param — palo_alto_fw_service.service_nodes.nodes.mgmt_subnet.subnet_param / 94cf6ce70042 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-1ca81a9ddb9348c620d239771df702d0b475bf02453eb0a2a936c6227d30b60c)
- [palo_alto_fw_service.service_nodes](data-sources--nfv_service--reference--group-004.md#canonical-b844e50e5ec67e22eb6280a754515913fecc9f21e564c553eeac9822b82d6daa)
- [palo_alto_fw_service.service_nodes.nodes](data-sources--nfv_service--reference--group-004.md#canonical-0e9fccb6f1864e82f7866453377cde5ecc1ed614f32deae2680af1a76c9f64d2)
- [palo_alto_fw_service.service_nodes.nodes.mgmt_subnet](data-sources--nfv_service--reference--group-004.md#canonical-39cf7caf789527f8242b1322147d03610959914f964493e7864c963ef286b71c)
- palo_alto_fw_service.service_nodes.nodes.mgmt_subnet.subnet_param

<a id="canonical-8137adcaa39053b44102c1a09914f99fa2daf7197dd2f626fb517fdc704b8ebf"></a>

Type: `"single"`. Computed.

Parameters for creating a new cloud subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-55ee181ac011b3a6d24ff9dde6016cd6d5055982f04308e56c25fdc4ca9fa0c2"></a>

## Direct properties — palo_alto_fw_service.service_nodes.nodes.mgmt_subnet.subnet_param / 94cf6ce70042 / 3

<a id="canonical-70026c5c42fcf3756a5c2a6625a04ec5f9dfdfeae169fb63ab87f4e280535859"></a>

<a id="canonical-c1ebec2f9a17951b8cbd172305624fd2ae416dbba1139871d10445641fdcfedf"></a>

## ipv4 property — palo_alto_fw_service.service_nodes.nodes.mgmt_subnet.subnet_param / 94cf6ce70042 / 4

Type: `"string"`. Computed.

IPv4 Subnet. IPv4 subnet prefix for this subnet.

Upstream description:

IPv4 subnet prefix for this subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "ipv4",
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  }
}
```

<a id="canonical-804742fb197515a3835614cd9cd3898d8cc27b7af5f9d5a281bf1de34e749075"></a>

## Next pages — palo_alto_fw_service.service_nodes.nodes.mgmt_subnet.subnet_param / 94cf6ce70042 / 5

- [palo_alto_fw_service.service_nodes.nodes.mgmt_subnet](data-sources--nfv_service--reference--group-004.md#canonical-39cf7caf789527f8242b1322147d03610959914f964493e7864c963ef286b71c)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-1e43326f6697853495d0f3f0b524b4a424a4441f171c5830c8734b7da8284273"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3d3c42515fc78cbc32be6b299cc93b863955c1855419cfaa591ec854b2a3f543"></a>

## palo_alto_fw_service.service_nodes.nodes.reserved_mgmt_subnet — palo_alto_fw_service.service_nodes.nodes.reserved_mgmt_subnet / 48dd0974fc7d / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-1ca81a9ddb9348c620d239771df702d0b475bf02453eb0a2a936c6227d30b60c)
- [palo_alto_fw_service.service_nodes](data-sources--nfv_service--reference--group-004.md#canonical-b844e50e5ec67e22eb6280a754515913fecc9f21e564c553eeac9822b82d6daa)
- [palo_alto_fw_service.service_nodes.nodes](data-sources--nfv_service--reference--group-004.md#canonical-0e9fccb6f1864e82f7866453377cde5ecc1ed614f32deae2680af1a76c9f64d2)
- palo_alto_fw_service.service_nodes.nodes.reserved_mgmt_subnet

<a id="canonical-ea05d0a89bc1d338ef2e57a04db338d492c86f7e44ce7c8e175b5666e702a218"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for reserved mgmt subnet.

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

<a id="canonical-38d38a22eeeedd11dda7d00f7050cf08e2df1e6d832789f73cc5a9b791f45634"></a>

## Direct properties — palo_alto_fw_service.service_nodes.nodes.reserved_mgmt_subnet / 48dd0974fc7d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d5fe34fd01c280ca1778f499c42d278bd1b46fca49efe5056c1eec9d15618931"></a>

## Next pages — palo_alto_fw_service.service_nodes.nodes.reserved_mgmt_subnet / 48dd0974fc7d / 4

- [palo_alto_fw_service.service_nodes.nodes](data-sources--nfv_service--reference--group-004.md#canonical-0e9fccb6f1864e82f7866453377cde5ecc1ed614f32deae2680af1a76c9f64d2)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
