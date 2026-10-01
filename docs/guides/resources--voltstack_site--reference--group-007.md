---
page_title: "xcsh_voltstack_site reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_voltstack_site reference."
---

# xcsh_voltstack_site reference

<a id="canonical-571f6a42d9532c7d574ce46bbce1886aa792ecada661e4e31ce925d67be18527"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 09017c499fb9 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-b84c8fd249287c751da748ad56a43c4930e50e9f2150d940fb34e1375474c21b)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--reference--group-006.md#canonical-631abb4ba0a2d2f35af76b0131e8a361ca50b9eef89361ed9dee41d900b74a04)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key](resources--voltstack_site--reference--group-006.md#canonical-bc80dd9d42d6825d0ec6d1e650739edaa480568a7018f14a7813e8c45d0a36fa)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info

<a id="canonical-badb90131d619f07ea9c4b01b76a3a1586f4d76b23827c82e4e9214dfb72ca3a"></a>

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

<a id="canonical-3bd5706dbb97403ac6d99165f9858c0321869c480cc1796682cad3b7c78d2330"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 09017c499fb9 / 3

<a id="canonical-dbd3d95ed330806efacf7084d8141edb581f70b4b42c0a9bef7b278e5f434ca3"></a>

<a id="canonical-b9a22231bdebc9a471642547c6484fa81acec50cdcef8038ba5f59af7f2f9995"></a>

## provider_ref property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 09017c499fb9 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-383b5d250711ea18c410725cfcded2ccef2ca596ba3ae7c6c3428710be60c598"></a>

<a id="canonical-2b44568fa2b88dccc21d2b1d3ebe7f1a86491eff5f2f030272676e46dc2fc627"></a>

## url property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 09017c499fb9 / 5

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

<a id="canonical-077a906539181c8de7621c8c296dec64770347aa4c2d6b8746f2a201b4706d6c"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 09017c499fb9 / 6

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key](resources--voltstack_site--reference--group-006.md#canonical-bc80dd9d42d6825d0ec6d1e650739edaa480568a7018f14a7813e8c45d0a36fa)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-ddc300ee675f6654917a2597aaa201b4f5df02043b47315e45e3081cd991d4c8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3619c4a231c7b0c44bcc07d09893e74462e3eda5b8f508ab7032ed403d4618ae"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.no_chap — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / d863b5572546 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-b84c8fd249287c751da748ad56a43c4930e50e9f2150d940fb34e1375474c21b)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--reference--group-006.md#canonical-631abb4ba0a2d2f35af76b0131e8a361ca50b9eef89361ed9dee41d900b74a04)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.no_chap

<a id="canonical-15ad78b85c13a4b13b51b37b7e225d054f09a47e8c8352d2400c3d75d89c21e6"></a>

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
no_chap = {}
```

<a id="canonical-19e01aa20c8229649ad70081d4c7ec349ef5bb60bbc0f7a162deaddb3d239000"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / d863b5572546 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-291b4365884fc639bad8f3bb05c664c8f5d8f9b17b9df70eef7f37315533bbb9"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / d863b5572546 / 4

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--reference--group-006.md#canonical-631abb4ba0a2d2f35af76b0131e8a361ca50b9eef89361ed9dee41d900b74a04)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-9edae627e94b200f98ab5d8f9076889619e850d6fcf6336daa0d1f7b9dccaf6d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c482b76b75186e66d3501fabd00e93a25c2c75b32b652c128a42987c802c9c7d"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 657e5adb9d8a / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-b84c8fd249287c751da748ad56a43c4930e50e9f2150d940fb34e1375474c21b)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--reference--group-006.md#canonical-631abb4ba0a2d2f35af76b0131e8a361ca50b9eef89361ed9dee41d900b74a04)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password

<a id="canonical-b9d7f706d1b169281e3bda17734ccffb05b82225c5240b74c5cef387318d125e"></a>

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
password {
  # Configure direct properties listed below.
}
```

<a id="canonical-389fa46c58014e9cdff27535d31191e13779a40c660567ec9934ea7eaeba4ec8"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 657e5adb9d8a / 3

- [blindfold_secret_info](resources--voltstack_site--reference--group-007.md#canonical-089330650007dd248934bcfae48d5f8055bcb75ecded09f718cd54b64eb601b5): complete subsection reference.

- [clear_secret_info](resources--voltstack_site--reference--group-007.md#canonical-ed598338277b4cd9c26398a42e04958d8b7adffa5234cd3236fa9935286acb1b): complete subsection reference.

<a id="canonical-e9682a1bdb11dbde1a99159f55ed895e84006ce159306236e21aeef1e6c389a9"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 657e5adb9d8a / 4

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info](resources--voltstack_site--reference--group-007.md#canonical-089330650007dd248934bcfae48d5f8055bcb75ecded09f718cd54b64eb601b5)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info](resources--voltstack_site--reference--group-007.md#canonical-ed598338277b4cd9c26398a42e04958d8b7adffa5234cd3236fa9935286acb1b)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--reference--group-006.md#canonical-631abb4ba0a2d2f35af76b0131e8a361ca50b9eef89361ed9dee41d900b74a04)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-089330650007dd248934bcfae48d5f8055bcb75ecded09f718cd54b64eb601b5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8d0ec2e9eb227d1ab3c34f045ac6d5960a2dbece3ce861735c0e854085708423"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 2885ed5f7fe0 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-b84c8fd249287c751da748ad56a43c4930e50e9f2150d940fb34e1375474c21b)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--reference--group-006.md#canonical-631abb4ba0a2d2f35af76b0131e8a361ca50b9eef89361ed9dee41d900b74a04)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password](resources--voltstack_site--reference--group-007.md#canonical-9edae627e94b200f98ab5d8f9076889619e850d6fcf6336daa0d1f7b9dccaf6d)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info

<a id="canonical-da4cdb368a18712c00bccc0aa7fe1589bfac544f81b88b61fb738e78c6afdfc0"></a>

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

<a id="canonical-01a661cbc62cbf3c63140a70e5a4053256484af142a015caaa3ba14b66b88805"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 2885ed5f7fe0 / 3

<a id="canonical-8b31a90edebc48c1694bc4b111e0a71722b65a92b374bbb9044e36dd13496207"></a>

<a id="canonical-989daa5fa3083ff6788f19e06edfee7b61241476f351c35f252a64292f01a41e"></a>

## decryption_provider property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 2885ed5f7fe0 / 4

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

<a id="canonical-40db04a4b78720e127fc6f457f5a490cd09f099b6d9310de7396e6186363e34b"></a>

<a id="canonical-4785033da010fbdfca0f05806a4f95cdf56cd34e60593483ed8ece0877d6d4dc"></a>

## location property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 2885ed5f7fe0 / 5

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

<a id="canonical-48e2585151f6e0ab4925e4a254e9addf9162a0a2628da905025473fc00062628"></a>

<a id="canonical-ee17cde84832c81567e2eb1991adb595d6acce0c630a8e144fe34332f1810863"></a>

## store_provider property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 2885ed5f7fe0 / 6

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

<a id="canonical-c5e8429aa69bb791be097c4af992f696e8576b90f5f6314ebe83ab14526f5297"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 2885ed5f7fe0 / 7

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password](resources--voltstack_site--reference--group-007.md#canonical-9edae627e94b200f98ab5d8f9076889619e850d6fcf6336daa0d1f7b9dccaf6d)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-ed598338277b4cd9c26398a42e04958d8b7adffa5234cd3236fa9935286acb1b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0c32f98e75e1b0371a60a52cb1cf8a014e7bd3964fdc734b648decea271307a2"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / c022b1d10873 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-b84c8fd249287c751da748ad56a43c4930e50e9f2150d940fb34e1375474c21b)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--reference--group-006.md#canonical-631abb4ba0a2d2f35af76b0131e8a361ca50b9eef89361ed9dee41d900b74a04)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password](resources--voltstack_site--reference--group-007.md#canonical-9edae627e94b200f98ab5d8f9076889619e850d6fcf6336daa0d1f7b9dccaf6d)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info

<a id="canonical-40ddb3025cdac65c477e74c84ba71667a9392c60d08590b88b72a74b84b263c1"></a>

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

<a id="canonical-65d048d4fd7ee066b07a01641422f0b30a0d0a12fb186391d622994fedefbcab"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / c022b1d10873 / 3

<a id="canonical-b4defde914373200fd10c580ea2eba623d3b93aba23e440e54e8fb6e3afb466f"></a>

<a id="canonical-dfd14797e59a3cffde316d0563165ec4b369594625b12e31342aa393f13c8422"></a>

## provider_ref property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / c022b1d10873 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2e253b598b1b570b3b796d00e8ac45df09f637d5ee8d8939112fc9ec257d3523"></a>

<a id="canonical-47c0c30cfa2bb73303318695a72aaebea427f2e2bbe0819c070443cf18123422"></a>

## url property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / c022b1d10873 / 5

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

<a id="canonical-db5a98173bc2504767d4628f631b460229878fb8c6d88750acd005e297096528"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / c022b1d10873 / 6

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password](resources--voltstack_site--reference--group-007.md#canonical-9edae627e94b200f98ab5d8f9076889619e850d6fcf6336daa0d1f7b9dccaf6d)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-8987b76c1b3b57fede2c84f453c17907fe933fe1ebcd9a0bea1ee6534c5da46a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9c4b06b2f07316e2e8012aa576170e8d4ece25635c4ba9ad1e01f041a9af4932"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 7c117ea21341 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-b84c8fd249287c751da748ad56a43c4930e50e9f2150d940fb34e1375474c21b)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--reference--group-006.md#canonical-631abb4ba0a2d2f35af76b0131e8a361ca50b9eef89361ed9dee41d900b74a04)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage

<a id="canonical-cbbde46af73df6ed329cbcab89b5b74ea96166114c67a8aaafd90b069443d727"></a>

Type: `"object"`. list nested block, Optional.

List of Virtual Storage Pool definitions which are referred back by Storage Class label match
selection.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
storage {
  # Configure direct properties listed below.
}
```

<a id="canonical-84fd493d22bcb5908cf8b996161eb73e92da260f7783e0f8515c44d7690d51f1"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 7c117ea21341 / 3

<a id="canonical-f1065169a3bcb3ac3f96d69cd121d8a1aa27dee501819ee29f140e638c8eb4f1"></a>

<a id="canonical-851dd179bffeabcf4f96def1dcce74e9e951c74f5d7bf47c0af57d508c285cec"></a>

## labels property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 7c117ea21341 / 4

Type: `["map", "string"]`. Optional.

List of labels for Storage Device used in NetApp ONTAP. It is used for storage class label match
selection.

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
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [volume_defaults](resources--voltstack_site--reference--group-007.md#canonical-62916bff510d72b933f4949d739251722c90c75fd2a77388c3031d572fd3f9c2): complete subsection reference.

<a id="canonical-e16c8e95f0cf3ca9af4a00d843deb1ce920ac444a53557cff1ae67a68fed1097"></a>

<a id="canonical-f7c3c2332d8bb58050b4d00ad5886b240f05d4fbd6023638b3f9d7f60e5cee8b"></a>

## zone property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 7c117ea21341 / 5

Type: `"string"`. Optional.

Virtual Pool Zone. Virtual Storage Pool zone definition.

Upstream description:

Virtual Storage Pool zone definition.

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

<a id="canonical-ae2627c4c7b52ccd2088f73a15aac87d29e10651df558a8682940998235b0575"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 7c117ea21341 / 6

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults](resources--voltstack_site--reference--group-007.md#canonical-62916bff510d72b933f4949d739251722c90c75fd2a77388c3031d572fd3f9c2)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--reference--group-006.md#canonical-631abb4ba0a2d2f35af76b0131e8a361ca50b9eef89361ed9dee41d900b74a04)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-62916bff510d72b933f4949d739251722c90c75fd2a77388c3031d572fd3f9c2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9d1cb0173cb033e513327b5d43f957c80c91f412c8b2b075eef5f9be04252c2d"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 0f6a8dae41c7 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-b84c8fd249287c751da748ad56a43c4930e50e9f2150d940fb34e1375474c21b)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--reference--group-006.md#canonical-631abb4ba0a2d2f35af76b0131e8a361ca50b9eef89361ed9dee41d900b74a04)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage](resources--voltstack_site--reference--group-007.md#canonical-8987b76c1b3b57fede2c84f453c17907fe933fe1ebcd9a0bea1ee6534c5da46a)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults

<a id="canonical-ae2abe34e8cd5e72add5ab4ab5969b8869371da2f5b979edb733b1b0bfd27578"></a>

Type: `"object"`. single nested block, Optional.

It controls how each volume is provisioned by default using these OPTIONS in a special section of
the configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("adaptive_qos_policy",
    "no_qos"),
  validators.ConflictingObjectAttributes("adaptive_qos_policy",
    "qos_policy"),
  validators.ConflictingObjectAttributes("no_qos",
    "qos_policy")}
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
  "x-ves-oneof-field-qos_policy_choice": "[\"adaptive_qos_policy\",\"no_qos\",\"qos_policy\"]"
}
```

Terraform syntax:

```terraform
volume_defaults {
  # Configure direct properties listed below.
}
```

<a id="canonical-a44177bf535f67887d3f4989c8e664d9eca03ab11067134502aa2c5a26b55183"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 0f6a8dae41c7 / 3

<a id="canonical-3685311abd3da4e51fd673466635116518737bc6de099bbaad87831103c0b21c"></a>

<a id="canonical-693b6b67d1224c6ac693dbfb9b51b69cf445515fdfc132c368a5832cb37098ec"></a>

## adaptive_qos_policy property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 0f6a8dae41c7 / 4

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Exclusive with \[no\_qos qos\_policy\] Enter Adaptive QoS Policy Name.

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
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-50a7a930334102842ec6121d54caee337ee7f75037092e4b5e06cfd5ed5ad116"></a>

<a id="canonical-402a2a03e9a8458ef6e926ad00a508eaac533d5c6254b8c0379ce61a6c50bdeb"></a>

## encryption property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 0f6a8dae41c7 / 5

Type: `"bool"`. Optional.

Enable Encryption. Enable NetApp volume encryption.

Upstream description:

Enable NetApp volume encryption.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0c309aa33f0f94dbe1933a91040adc08f6c435ad6f45c8a1fd9e3bb3fd3b4aa7"></a>

<a id="canonical-c49b28071b4a8c58f397debe75a7ae78f1d7ab4e66a011d5506ab2cd36e84f1a"></a>

## export_policy property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 0f6a8dae41c7 / 6

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Export policy to use.

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

- [no_qos](resources--voltstack_site--reference--group-007.md#canonical-e772a6664396a4e9c6d13b52572ffecf202100ef8d035f568ee457e73a835692): complete subsection reference.

<a id="canonical-a5014b8e85c707db9d35dee6f017dccbc6958b2bd3444ddf9ad33f6b7ce6af4f"></a>

<a id="canonical-9d96cd2b4f340745fd733aa842693fe7770773fa86626dc1b7a5ccf5b22ba4e8"></a>

## qos_policy property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 0f6a8dae41c7 / 7

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Exclusive with \[adaptive\_qos\_policy no\_qos\] Enter QoS Policy Name.

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
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-086e907125cff1c5a247027d46a4e22bd1506169498a74298a3315547a17b231"></a>

<a id="canonical-9dda97b153a15d94d64387f29d69be3dcdd8ce442f48f5ffd71e7c2cc2c9fe4b"></a>

## security_style property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 0f6a8dae41c7 / 8

Type: `"string"`. Optional.

Security Style. Security style for new volumes.

Upstream description:

Security style for new volumes.

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

<a id="canonical-e03ddfb364d2642001345c62b230f86c0e843039eed92139296c88d3e78ab25d"></a>

<a id="canonical-a81abb60657b7e36cd23e38d118fbccfa584d858340ea67d913acdce602b7d7a"></a>

## snapshot_dir property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 0f6a8dae41c7 / 9

Type: `"bool"`. Optional.

Access to Snapshot Directory. Access to the .snapshot directory.

Upstream description:

Access to the .snapshot directory.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-56bdcc95d7ee8a95a99bcfc49e94938fa1ea6d6c2f8232df9c95f64846b18fa5"></a>

<a id="canonical-63aadae7e185edb8bb08a2dda48d5bb16da1c98037f78ef3017e3291ab284b22"></a>

## snapshot_policy property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 0f6a8dae41c7 / 10

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Snapshot policy to use.

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

<a id="canonical-7b02ded62c0606a6ac0ef15fe0ef72c67b952719c7810053cd00a2b13eac1572"></a>

<a id="canonical-e19725a7317e8ad0c0809f928fbc3616498ce4022661b9a9170c624e32a24ccb"></a>

## snapshot_reserve property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 0f6a8dae41c7 / 11

Type: `"string"`. Optional.

Percentage of volume reserved for snapshots. '0' if snapshot policy is 'none', else ''.

Upstream description:

Percentage of volume reserved for snapshots. "0" if snapshot policy is "none", else ""

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

<a id="canonical-8d3c2685b082c6c6b0a0d4a3597be0b9af370334bd1edb15dbe3d5883a8f0afc"></a>

<a id="canonical-ffc897e3648bc300ee52510a3f20a79cee018f3bac22122213ff38f528680dba"></a>

## space_reserve property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 0f6a8dae41c7 / 12

Type: `"string"`. Optional.

\[Enum: none|thick\] Space reservation mode; “none” (thin) or “volume” (thick). Possible values are
\`none\`, \`thick\`.

Upstream description:

Space reservation mode; “none” (thin) or “volume” (thick)

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("none",
    "thick"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "none",
    "thick"
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
    "ves.io.schema.rules.string.in": "[\\\"none\\\",\\\"thick\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"none\\\",\\\"thick\\\"]"
  }
}
```

<a id="canonical-7072d405bb3a3c5a5e10b621259676c488d68c4b485aa8fa4e7348af93e8c944"></a>

<a id="canonical-0db413eb059227f64b7b85dd95f9bff4ff305bc0d2f9afcbc6fcd3f0da72bf81"></a>

## split_on_clone property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 0f6a8dae41c7 / 13

Type: `"bool"`. Optional.

Split a clone from its parent upon creation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-7679174228af12d2a25128a47fa3535876eefaa1a2005e3d3bab26f2b802c6b5"></a>

<a id="canonical-947751f3239903f5a448635625a1c429ec3596cc66b14fe98d3931d1dd476c7e"></a>

## tiering_policy property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 0f6a8dae41c7 / 14

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Tiering policy to use. "none" is default.

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

<a id="canonical-238cb87597b3fa536a8c0c5f68924946f6e3209de7be8342f838f177b1d5c8f2"></a>

<a id="canonical-9ad5f9bfd61eeecabf2756eea6f68f45fd2ebdd997e84a3fb2c654f5f9cb8e18"></a>

## unix_permissions property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 0f6a8dae41c7 / 15

Type: `"number"`. Optional.

Unix permission mode for new volumes. All allowed 777.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-55a94796ff229d9f094a3fbf29b9074263a3495bb22767da459c0de940f67cef"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 0f6a8dae41c7 / 16

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.no_qos](resources--voltstack_site--reference--group-007.md#canonical-e772a6664396a4e9c6d13b52572ffecf202100ef8d035f568ee457e73a835692)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage](resources--voltstack_site--reference--group-007.md#canonical-8987b76c1b3b57fede2c84f453c17907fe933fe1ebcd9a0bea1ee6534c5da46a)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-e772a6664396a4e9c6d13b52572ffecf202100ef8d035f568ee457e73a835692"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d0794264b78c17cfa76a95f8217499114c08f2999c213cc87c47e2d33ddc1ef4"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.no_qos — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 6aad893fd1e6 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-b84c8fd249287c751da748ad56a43c4930e50e9f2150d940fb34e1375474c21b)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--reference--group-006.md#canonical-631abb4ba0a2d2f35af76b0131e8a361ca50b9eef89361ed9dee41d900b74a04)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage](resources--voltstack_site--reference--group-007.md#canonical-8987b76c1b3b57fede2c84f453c17907fe933fe1ebcd9a0bea1ee6534c5da46a)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults](resources--voltstack_site--reference--group-007.md#canonical-62916bff510d72b933f4949d739251722c90c75fd2a77388c3031d572fd3f9c2)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.no_qos

<a id="canonical-3f84deddc70fb576903d9aaee800d45969ec35c96bd6eb08396097331c28e67d"></a>

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
no_qos = {}
```

<a id="canonical-6ac39c27215d8d726ffacb62b7a4ebc1fd552834139cbdf452809cf9dfce8f56"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 6aad893fd1e6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4507544e0db0eab6d0ce9e65d53b872f0da4ef5628926054c74b9af8e118a3a5"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 6aad893fd1e6 / 4

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults](resources--voltstack_site--reference--group-007.md#canonical-62916bff510d72b933f4949d739251722c90c75fd2a77388c3031d572fd3f9c2)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-72930290386fbf5a47dda8bd8be3f15dfd77c8d7af59e1bb7cac4cdc40c893e9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d03e7017b611713a05a96971a56ed270bf376f0627a2dcba3f8f076470bd48e5"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / f5899ea317e9 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-b84c8fd249287c751da748ad56a43c4930e50e9f2150d940fb34e1375474c21b)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--reference--group-006.md#canonical-631abb4ba0a2d2f35af76b0131e8a361ca50b9eef89361ed9dee41d900b74a04)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap

<a id="canonical-2816725bd5f070833a328689b70d4220003dafaab95975e351a2628874beffff"></a>

Type: `"object"`. single nested block, Optional.

Device NetApp Backend ONTAP SAN CHAP configuration OPTIONS for enabled CHAP.

Receipt-pinned upstream constraints:

```json
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
use_chap {
  # Configure direct properties listed below.
}
```

<a id="canonical-fdb5c26aa8e198fd89832d213a579e66b2e9013e3e5209430c02ca279a664c84"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / f5899ea317e9 / 3

- [chap_initiator_secret](resources--voltstack_site--reference--group-007.md#canonical-62fbaf2fd45b2b6ef9ba5dacea4f4c01329fdb6335c5fdbc81ef042d83bcfbbe): complete subsection reference.

- [chap_target_initiator_secret](resources--voltstack_site--reference--group-007.md#canonical-80c3458fccb6fed4fb2fccb4ef4667331272c46020f4d6f0b9a880a139c91c7b): complete subsection reference.

<a id="canonical-1c51fd93b387add2db518af886781b203e0b63f8f6a5f7166dd6a1aeaa73c426"></a>

<a id="canonical-8e029aa82dea0677398d38bfbc581ed3a1ef6515b6d4874faeee1de01ab40836"></a>

## chap_target_username property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / f5899ea317e9 / 4

Type: `"string"`. Optional.

Target username. Required if useCHAP=true.

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

<a id="canonical-d7139a6c534f8adf25cd4bf7c1a879aa43713c0fffbc3585c9a237b41d6bb982"></a>

<a id="canonical-8be88ceba0ff0877506d899f8c33706a361aa659cc537f88a2377d59f5a21f70"></a>

## chap_username property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / f5899ea317e9 / 5

Type: `"string"`. Optional.

Inbound username. Required if useCHAP=true.

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

<a id="canonical-6dcc40cef8bc878404544441a44ae5ed8bc9f4db41a30a05729b365701eb79d5"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / f5899ea317e9 / 6

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret](resources--voltstack_site--reference--group-007.md#canonical-62fbaf2fd45b2b6ef9ba5dacea4f4c01329fdb6335c5fdbc81ef042d83bcfbbe)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret](resources--voltstack_site--reference--group-007.md#canonical-80c3458fccb6fed4fb2fccb4ef4667331272c46020f4d6f0b9a880a139c91c7b)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--reference--group-006.md#canonical-631abb4ba0a2d2f35af76b0131e8a361ca50b9eef89361ed9dee41d900b74a04)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-62fbaf2fd45b2b6ef9ba5dacea4f4c01329fdb6335c5fdbc81ef042d83bcfbbe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6be22554074810af8930161e2f4844584d47ef2d1192366d0a140aae8898e162"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / a5343405396c / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-b84c8fd249287c751da748ad56a43c4930e50e9f2150d940fb34e1375474c21b)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--reference--group-006.md#canonical-631abb4ba0a2d2f35af76b0131e8a361ca50b9eef89361ed9dee41d900b74a04)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](resources--voltstack_site--reference--group-007.md#canonical-72930290386fbf5a47dda8bd8be3f15dfd77c8d7af59e1bb7cac4cdc40c893e9)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret

<a id="canonical-a7fe6b8124420f2129dbf14f6026a99af5b20122a6d50425466208bd518b4c52"></a>

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
chap_initiator_secret {
  # Configure direct properties listed below.
}
```

<a id="canonical-bdb15ca0a8df3f661ea22577f35ac340aec49078a7778d8da3f8ab69a20231c5"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / a5343405396c / 3

- [blindfold_secret_info](resources--voltstack_site--reference--group-007.md#canonical-876a5f24010d2003ef235930a0e5e9b4c64933bf0f95e5e4aeaa3f56d9adf310): complete subsection reference.

- [clear_secret_info](resources--voltstack_site--reference--group-007.md#canonical-0de35ee818dd0a51f4577fff77f0248a37cb78c0e201befd1badc2514da5dd8b): complete subsection reference.

<a id="canonical-a713689ab09724434b628ea2695b6e14cdc2155507e20035129acba52f5ffc22"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / a5343405396c / 4

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info](resources--voltstack_site--reference--group-007.md#canonical-876a5f24010d2003ef235930a0e5e9b4c64933bf0f95e5e4aeaa3f56d9adf310)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info](resources--voltstack_site--reference--group-007.md#canonical-0de35ee818dd0a51f4577fff77f0248a37cb78c0e201befd1badc2514da5dd8b)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](resources--voltstack_site--reference--group-007.md#canonical-72930290386fbf5a47dda8bd8be3f15dfd77c8d7af59e1bb7cac4cdc40c893e9)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-876a5f24010d2003ef235930a0e5e9b4c64933bf0f95e5e4aeaa3f56d9adf310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a0289d118b4466373b2f558eed0f7e67370461fb5a9b8a90c2d55f3cf2d666d5"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / cfa3cbb3f2af / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-b84c8fd249287c751da748ad56a43c4930e50e9f2150d940fb34e1375474c21b)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--reference--group-006.md#canonical-631abb4ba0a2d2f35af76b0131e8a361ca50b9eef89361ed9dee41d900b74a04)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](resources--voltstack_site--reference--group-007.md#canonical-72930290386fbf5a47dda8bd8be3f15dfd77c8d7af59e1bb7cac4cdc40c893e9)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret](resources--voltstack_site--reference--group-007.md#canonical-62fbaf2fd45b2b6ef9ba5dacea4f4c01329fdb6335c5fdbc81ef042d83bcfbbe)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info

<a id="canonical-f3e7720ad7864625e3b2929ae785649a564c5861d859d50faf06d231fa46997e"></a>

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

<a id="canonical-c90b6e5189c170505611d421b56b637fd5f1727bf88c5ebf1833064112210592"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / cfa3cbb3f2af / 3

<a id="canonical-c464d96fbacc1c5da381be4a34946efb1f7f14f29bec5dbb9e54cc5564232cb0"></a>

<a id="canonical-cbc360e9329d7203456b021efaa9ca01d95a94eebb856ce2f7a86d0200bb0ce6"></a>

## decryption_provider property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / cfa3cbb3f2af / 4

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

<a id="canonical-1ec22741155ce7f87560d4c611c41303754abfb8df27861d04e7f5b76e10e5be"></a>

<a id="canonical-29b9f1fc61722cbd32e33093badb16cf0c48c47860d62f02f9a85e27b09b5517"></a>

## location property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / cfa3cbb3f2af / 5

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

<a id="canonical-e3224caf897e883984c6cd08f3c6c349ffc4c3c88c7e72c8c7f0e3c9306373fc"></a>

<a id="canonical-1b5c9fe21f886efe12d6707acde4cdf736f34d0d0262fd9eb3a150690d678bf3"></a>

## store_provider property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / cfa3cbb3f2af / 6

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

<a id="canonical-45d2390ebf710dc82d6e257e389dc81f6bdbc1444db0172d1cff7636627ee7c1"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / cfa3cbb3f2af / 7

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret](resources--voltstack_site--reference--group-007.md#canonical-62fbaf2fd45b2b6ef9ba5dacea4f4c01329fdb6335c5fdbc81ef042d83bcfbbe)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-0de35ee818dd0a51f4577fff77f0248a37cb78c0e201befd1badc2514da5dd8b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ff4914b3d02af970ae5980706acdd7761e160237e35fb4535fe430d3c69b545c"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / e3db3898e7cd / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-b84c8fd249287c751da748ad56a43c4930e50e9f2150d940fb34e1375474c21b)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--reference--group-006.md#canonical-631abb4ba0a2d2f35af76b0131e8a361ca50b9eef89361ed9dee41d900b74a04)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](resources--voltstack_site--reference--group-007.md#canonical-72930290386fbf5a47dda8bd8be3f15dfd77c8d7af59e1bb7cac4cdc40c893e9)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret](resources--voltstack_site--reference--group-007.md#canonical-62fbaf2fd45b2b6ef9ba5dacea4f4c01329fdb6335c5fdbc81ef042d83bcfbbe)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info

<a id="canonical-f3d35020a1fb0db06e8c5b798104df414ab307e70e11bd6b3f64db2c501b6623"></a>

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

<a id="canonical-27c097ddc62cacb2cd5b3bf8215dbaf448cafec3e11518fdd2c0d7452790f007"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / e3db3898e7cd / 3

<a id="canonical-9d9ad96b91c0ee78419cdeda66cdd895f28089a7042b32793ae1c65018265b6d"></a>

<a id="canonical-a82accbd89fe9c0de360d2dfe58f56afbb14dfccf4457d1b0a9c8fc8d5c0627e"></a>

## provider_ref property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / e3db3898e7cd / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-da6e13208a6e1929c35f0f11ee4ba74407dbddf456697098c923e4d1d8d9e837"></a>

<a id="canonical-a9fd07b9a607cd153efe8d56e5417b3bd7e1c601d8ef02d40a83d51b783345c3"></a>

## url property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / e3db3898e7cd / 5

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

<a id="canonical-f561f9218b129809e4e7dc56da5771525aec675096b02b6b7d82f6258532d94b"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / e3db3898e7cd / 6

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret](resources--voltstack_site--reference--group-007.md#canonical-62fbaf2fd45b2b6ef9ba5dacea4f4c01329fdb6335c5fdbc81ef042d83bcfbbe)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-80c3458fccb6fed4fb2fccb4ef4667331272c46020f4d6f0b9a880a139c91c7b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-009bfe81eb3d799698ce945076e55f59a0c21d1cf028f10854cd299dd0f44351"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / c19660086f60 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-b84c8fd249287c751da748ad56a43c4930e50e9f2150d940fb34e1375474c21b)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--reference--group-006.md#canonical-631abb4ba0a2d2f35af76b0131e8a361ca50b9eef89361ed9dee41d900b74a04)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](resources--voltstack_site--reference--group-007.md#canonical-72930290386fbf5a47dda8bd8be3f15dfd77c8d7af59e1bb7cac4cdc40c893e9)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret

<a id="canonical-eda92a6e1c99a76852abae7d6bd79f44341cda3642b86bbb46448421632b1c43"></a>

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
chap_target_initiator_secret {
  # Configure direct properties listed below.
}
```

<a id="canonical-e470adf57e1e0d736ee41b58e9c0d1c809cd9aef950cc1ec2b74a72e5427ef34"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / c19660086f60 / 3

- [blindfold_secret_info](resources--voltstack_site--reference--group-007.md#canonical-8e262ac89677dd59594ad65b4a712812b7ed7a8108e92c59081fd87eaace7124): complete subsection reference.

- [clear_secret_info](resources--voltstack_site--reference--group-007.md#canonical-2efe9d6db821ada3c9d79ef7135910e9a0a2d2115c84a995149a580d291afcb6): complete subsection reference.

<a id="canonical-ec3de844afa248841e2b65852ed36b255cb06264e674415b92bf66d4284df310"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / c19660086f60 / 4

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info](resources--voltstack_site--reference--group-007.md#canonical-8e262ac89677dd59594ad65b4a712812b7ed7a8108e92c59081fd87eaace7124)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info](resources--voltstack_site--reference--group-007.md#canonical-2efe9d6db821ada3c9d79ef7135910e9a0a2d2115c84a995149a580d291afcb6)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](resources--voltstack_site--reference--group-007.md#canonical-72930290386fbf5a47dda8bd8be3f15dfd77c8d7af59e1bb7cac4cdc40c893e9)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-8e262ac89677dd59594ad65b4a712812b7ed7a8108e92c59081fd87eaace7124"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9f0457bb6e6d87b984a150ff3fb8682ce9239a909584b23f5e85367fb329c4ac"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 6acd900c12a7 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-b84c8fd249287c751da748ad56a43c4930e50e9f2150d940fb34e1375474c21b)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--reference--group-006.md#canonical-631abb4ba0a2d2f35af76b0131e8a361ca50b9eef89361ed9dee41d900b74a04)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](resources--voltstack_site--reference--group-007.md#canonical-72930290386fbf5a47dda8bd8be3f15dfd77c8d7af59e1bb7cac4cdc40c893e9)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret](resources--voltstack_site--reference--group-007.md#canonical-80c3458fccb6fed4fb2fccb4ef4667331272c46020f4d6f0b9a880a139c91c7b)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info

<a id="canonical-e3f3b45e5997b909b9cad32a6622193dfc93a2680845d1761c2fc55039c89997"></a>

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

<a id="canonical-f63f907e5f6c5d582e3002ede1499e3388de64f312bdaa48ef71e43f5ad50bf2"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 6acd900c12a7 / 3

<a id="canonical-f32a543e6321b36df4905c3c4a532ede8f3340ea0b5d7c8b1f11ac4fb97cb427"></a>

<a id="canonical-ccf8cf232b5fb126983e715ce38f08330ce588cf3315c47367ab2cf295e7eafe"></a>

## decryption_provider property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 6acd900c12a7 / 4

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

<a id="canonical-e3b48d58a29481b4d195701774ca3fb2bc14619509105540b494d5a9919b9f6b"></a>

<a id="canonical-877d7964facf5f3a3663dbd711b3573f6ba6ce11f36741685ba39e06e3c2ee57"></a>

## location property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 6acd900c12a7 / 5

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

<a id="canonical-7de10487fd3a14e2fec297cd01aacf36ccc2cc2ceef5c506081e1649b6ad9fcf"></a>

<a id="canonical-b0592b7b876f1b2f548423be6b4a238926c1b41c443448b1ddd7d9c653f3b6e2"></a>

## store_provider property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 6acd900c12a7 / 6

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

<a id="canonical-b7919516d629efb1eac95bf7e2c1f6796a4572be57d6a9390c673e8c2a620341"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 6acd900c12a7 / 7

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret](resources--voltstack_site--reference--group-007.md#canonical-80c3458fccb6fed4fb2fccb4ef4667331272c46020f4d6f0b9a880a139c91c7b)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-2efe9d6db821ada3c9d79ef7135910e9a0a2d2115c84a995149a580d291afcb6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4f815b345917e3072f411ea4270d1b636a73cc8b8d8dde2286c6ed3682d3fd5d"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / b2c7ff5eb194 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-b84c8fd249287c751da748ad56a43c4930e50e9f2150d940fb34e1375474c21b)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--reference--group-006.md#canonical-631abb4ba0a2d2f35af76b0131e8a361ca50b9eef89361ed9dee41d900b74a04)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](resources--voltstack_site--reference--group-007.md#canonical-72930290386fbf5a47dda8bd8be3f15dfd77c8d7af59e1bb7cac4cdc40c893e9)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret](resources--voltstack_site--reference--group-007.md#canonical-80c3458fccb6fed4fb2fccb4ef4667331272c46020f4d6f0b9a880a139c91c7b)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info

<a id="canonical-64a510bcf5d56127e598fe7a9386d26953082eb3dbfaf103b4f49de0f594682c"></a>

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

<a id="canonical-01e8ef54c79efac4be5d08524dd6b98ebc5bf0a31e9fc612eb264369e21c6057"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / b2c7ff5eb194 / 3

<a id="canonical-d9abb1e425b1b58e0bcc4b27b3b02c1cdf97fcb6833932bc9f5e1900e3ec04ad"></a>

<a id="canonical-16acd46bc2936137fd90ceb80dcc8a6632b4af109572b769ab4a53ea07c552f6"></a>

## provider_ref property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / b2c7ff5eb194 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-306fe39c941396a93d518c17f65c2e1b0a1791c2477bd51eca9bea2e3673727d"></a>

<a id="canonical-e63f2851c42cfac8ba3013ff56ba2d8bd9caeb5d67b6bdc75f203432b85e751d"></a>

## url property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / b2c7ff5eb194 / 5

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

<a id="canonical-76268d5333973df3f6802b37e150bde40dc7cc8e15de46ac2c5a2de3c1d4e5ef"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / b2c7ff5eb194 / 6

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret](resources--voltstack_site--reference--group-007.md#canonical-80c3458fccb6fed4fb2fccb4ef4667331272c46020f4d6f0b9a880a139c91c7b)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-571b06ed2c72749baa5e3ab02d7db10f83d9d762b9cc97ccd21fa4fc922d4e6a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bbeeb240738ecf76f0df65e56b27ab516ce499dcd8062578c937e246194b91cc"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / e37f00cb5c41 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-b84c8fd249287c751da748ad56a43c4930e50e9f2150d940fb34e1375474c21b)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--reference--group-006.md#canonical-631abb4ba0a2d2f35af76b0131e8a361ca50b9eef89361ed9dee41d900b74a04)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults

<a id="canonical-c431237b4ca5798bc47244b9e3bba65254b19017113e42cc7143095f5956fe48"></a>

Type: `"object"`. single nested block, Optional.

It controls how each volume is provisioned by default using these OPTIONS in a special section of
the configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("adaptive_qos_policy",
    "no_qos"),
  validators.ConflictingObjectAttributes("adaptive_qos_policy",
    "qos_policy"),
  validators.ConflictingObjectAttributes("no_qos",
    "qos_policy")}
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
  "x-ves-oneof-field-qos_policy_choice": "[\"adaptive_qos_policy\",\"no_qos\",\"qos_policy\"]"
}
```

Terraform syntax:

```terraform
volume_defaults {
  # Configure direct properties listed below.
}
```

<a id="canonical-fc03e8dc93b2a24b77af2daa0b843e2bba5a12ff2d4c3e66ac9b65f239f10e55"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / e37f00cb5c41 / 3

<a id="canonical-3a18cf52e5ca619e5e47324cf112bcd0c14208802fb0982c4e70be7199163ad9"></a>

<a id="canonical-0cef87301c326ea02947860d680954bd6e20eff6176a147b358796808cf0422a"></a>

## adaptive_qos_policy property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / e37f00cb5c41 / 4

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Exclusive with \[no\_qos qos\_policy\] Enter Adaptive QoS Policy Name.

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
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1c4759844a4b7cf211f25bd61f5acfc03ca658b4c18c49a2a90cff5cd1983536"></a>

<a id="canonical-737153351cb5fb7b02844018a90025c3c8415bcd1318f3a00aa8a865e131ac87"></a>

## encryption property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / e37f00cb5c41 / 5

Type: `"bool"`. Optional.

Enable Encryption. Enable NetApp volume encryption.

Upstream description:

Enable NetApp volume encryption.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-c485d1fe702c0f6d1a73c5f50a223e3024ac29c23accd4632f52000cc87ed74d"></a>

<a id="canonical-d2f5d0fc62923e9e389643a261a009431a574fd44ed69cd4e3120744f740e415"></a>

## export_policy property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / e37f00cb5c41 / 6

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Export policy to use.

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

- [no_qos](resources--voltstack_site--reference--group-007.md#canonical-8eb558c111244565e4bb5d53525eace91ecd862f27b1ab03294834e09ef198bd): complete subsection reference.

<a id="canonical-f9c19ed423e9e6b6465284ebe5d526390da0bb12f377db7a96a20cccb6f70a0c"></a>

<a id="canonical-977cab6a9924de4265518a9548ff082fb9452f7442d1b948a93efd46a6969e67"></a>

## qos_policy property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / e37f00cb5c41 / 7

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Exclusive with \[adaptive\_qos\_policy no\_qos\] Enter QoS Policy Name.

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
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-b6fbdd6d76750588afae1a47c664cedb38409cae7ff80b8f55d38acdf5907394"></a>

<a id="canonical-181525ed275219bbee480a56495589516229cd458acf1ca9938dda1246044218"></a>

## security_style property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / e37f00cb5c41 / 8

Type: `"string"`. Optional.

Security Style. Security style for new volumes.

Upstream description:

Security style for new volumes.

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

<a id="canonical-b3e5b03768a92b799dd9ea484f86681763d2a18485e2333e8b3538448d18c6f0"></a>

<a id="canonical-c1d62e302c35ceebaab431ccf8fed7b9cd4670bcdfcfc11da52ece3e661c9080"></a>

## snapshot_dir property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / e37f00cb5c41 / 9

Type: `"bool"`. Optional.

Access to Snapshot Directory. Access to the .snapshot directory.

Upstream description:

Access to the .snapshot directory.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-d0e078895868962801398f5178807a13143623e16af80a56b10192eaf484c074"></a>

<a id="canonical-fd362dcc5c5dfeef7d89c4045146d447928117b6190c1c0427e51caa7e458056"></a>

## snapshot_policy property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / e37f00cb5c41 / 10

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Snapshot policy to use.

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

<a id="canonical-fb715670b8fe7dff45e8de574363d6b4b4d95aaedf98170463407aa8b4ca2b62"></a>

<a id="canonical-f1d80ca4f42e5ae96acf09d3b9acdd27c2ccf038d8cebf3cc4e4f6a83fa3ca0b"></a>

## snapshot_reserve property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / e37f00cb5c41 / 11

Type: `"string"`. Optional.

Percentage of volume reserved for snapshots. '0' if snapshot policy is 'none', else ''.

Upstream description:

Percentage of volume reserved for snapshots. "0" if snapshot policy is "none", else ""

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

<a id="canonical-37e940a6bad280bb5db23f85cd00557d61536cca906235689681d6f254945d5b"></a>

<a id="canonical-9159387e6e7b4a2889f60ebf47ca5989ad197c3c68dc8e95bdf9fec15b2c69e4"></a>

## space_reserve property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / e37f00cb5c41 / 12

Type: `"string"`. Optional.

\[Enum: none|thick\] Space reservation mode; “none” (thin) or “volume” (thick). Possible values are
\`none\`, \`thick\`.

Upstream description:

Space reservation mode; “none” (thin) or “volume” (thick)

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("none",
    "thick"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "none",
    "thick"
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
    "ves.io.schema.rules.string.in": "[\\\"none\\\",\\\"thick\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"none\\\",\\\"thick\\\"]"
  }
}
```

<a id="canonical-c4dbcd5ce6d89b5bf0da59a576ca1e202af3c120e6dd6a2353603237bbaa5d16"></a>

<a id="canonical-5a5a08a286c6eb471894ea73e9a840fbb0b72d5410bc52f12f94e42d65065d93"></a>

## split_on_clone property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / e37f00cb5c41 / 13

Type: `"bool"`. Optional.

Split a clone from its parent upon creation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-35fef5225162826226bda4bb0d56ff213c3e30192368d462c046838e4cf56780"></a>

<a id="canonical-0113584a8b7e63b5a9bf758947ee6413511eeb6bc28a68c09e4823be7b447cf8"></a>

## tiering_policy property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / e37f00cb5c41 / 14

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Tiering policy to use. "none" is default.

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

<a id="canonical-05365c6c643e6dd022c3bf6d07ccea4eb36f581aee7c58c4625228c63950f0e1"></a>

<a id="canonical-c3d396dd8b0340b54f06c2d103e0e61c80397377ce94d72a7073bf74834e9b37"></a>

## unix_permissions property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / e37f00cb5c41 / 15

Type: `"number"`. Optional.

Unix permission mode for new volumes. All allowed 777.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-30a21da5188a53ce51e1584ffcc27a6ebb9b83138e0359954665b78ded10c65e"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / e37f00cb5c41 / 16

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.no_qos](resources--voltstack_site--reference--group-007.md#canonical-8eb558c111244565e4bb5d53525eace91ecd862f27b1ab03294834e09ef198bd)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--reference--group-006.md#canonical-631abb4ba0a2d2f35af76b0131e8a361ca50b9eef89361ed9dee41d900b74a04)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-8eb558c111244565e4bb5d53525eace91ecd862f27b1ab03294834e09ef198bd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6b9afd2bff13550c2f5b1f2e99ee843ec4bfe82b9b9be530617490103b9d4fa8"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.no_qos — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 770b749a408c / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-b84c8fd249287c751da748ad56a43c4930e50e9f2150d940fb34e1375474c21b)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--reference--group-006.md#canonical-631abb4ba0a2d2f35af76b0131e8a361ca50b9eef89361ed9dee41d900b74a04)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults](resources--voltstack_site--reference--group-007.md#canonical-571b06ed2c72749baa5e3ab02d7db10f83d9d762b9cc97ccd21fa4fc922d4e6a)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.no_qos

<a id="canonical-95b8fa21a5665280816e2c245c70bd16ca0536ad4a0ffb04823c425222c409c4"></a>

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
no_qos = {}
```

<a id="canonical-4d9d4ac66b9d8a61790741daec63365e2108269f0f2cb677e407d7c72f073360"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 770b749a408c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e684641c5f6dac5a39f4ac3f8b722f4a79f2fe83500817a85f7c4a524ad1c9c8"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 770b749a408c / 4

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults](resources--voltstack_site--reference--group-007.md#canonical-571b06ed2c72749baa5e3ab02d7db10f83d9d762b9cc97ccd21fa4fc922d4e6a)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-de9ccfeb4756e83ff6cc51ab35fb939749876ea24008d3dcd033ebad9d591290"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-01355aca0cf7a4f4c9f734752e4bf859bb9dee65470664b4ba883f41c33cc46e"></a>

## custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / ce075d67ec78 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator

<a id="canonical-80c43cb2c10cd305b951ae4e890a6938d85995631f373f377c691f74049f59cb"></a>

Type: `"object"`. single nested block, Optional.

Device configuration for Pure Storage Service Orchestrator.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cluster_id")}
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
pure_service_orchestrator {
  # Configure direct properties listed below.
}
```

<a id="canonical-34076224cd3965aca7c4335c01533c9c49fea6d170744724a8da2845fa3d8845"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / ce075d67ec78 / 3

- [arrays](resources--voltstack_site--reference--group-007.md#canonical-10e5e15a5e7a3f4cc11d72b2d663f32fce0d18582ae0fe8cc49686085cfa5b65): complete subsection reference.

<a id="canonical-aef0f8926f3f1985b2d813b00dbab6c3a90de3026430329edfc15068f0774fb3"></a>

<a id="canonical-c4b6a0dcebdefe7d6c21011e03c6c1e18c72631c696a9b05cec290ea9f187039"></a>

## cluster_id property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / ce075d67ec78 / 4

Type: `"string"`. Optional.

ClusterID is added as a prefix for all volumes created by this PSO installation. ClusterID is also
used to identify the volumes used by the datastore, pso-db. ClusterID MUST BE UNIQUE for multiple
K8s clusters running on top of the same storage arrays.

Upstream description:

ClusterID is added as a prefix for all volumes created by this PSO installation. ClusterID is also
used to identify the volumes used by the datastore, pso-db. ClusterID MUST BE UNIQUE for multiple
K8s clusters running on top of the same storage arrays. Characters allowed: alphanumeric and
underscores.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 22),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 22,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 22,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z0-9_]*$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "22",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.pattern": "^[a-z0-9_]*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "22",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.pattern": "^[a-z0-9_]*$"
  }
}
```

<a id="canonical-5cb6f00fd1d0bcfe50bbc0d8d95b1cfe819db07b2725f169729ebcf355961e13"></a>

<a id="canonical-c94a9245828885b85a7851282d57f89c435e215458c948c8c3a50a78b0f7a015"></a>

## enable_storage_topology property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / ce075d67ec78 / 5

Type: `"bool"`. Optional.

Option is to enable/disable the csi topology feature for pso-csi.

Upstream description:

This option is to enable/disable the csi topology feature for pso-csi.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-4331697cc7ebbd3947e363dca0f49e2815ce36576ee99fed0eb61f8241d0580a"></a>

<a id="canonical-91be8f2a67dbce2a3f73f64e09100e3e6a54ea2dbd40fce9fb0589ccef97db7b"></a>

## enable_strict_topology property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / ce075d67ec78 / 6

Type: `"bool"`. Optional.

Option is to enable/disable the strict csi topology feature for pso-csi.

Upstream description:

This option is to enable/disable the strict csi topology feature for pso-csi.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-4912b79a3998dc79d2e0a8233ea235e7e6b9cb805ac56b3fd0c6602be1ff9d13"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / ce075d67ec78 / 7

- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](resources--voltstack_site--reference--group-007.md#canonical-10e5e15a5e7a3f4cc11d72b2d663f32fce0d18582ae0fe8cc49686085cfa5b65)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-10e5e15a5e7a3f4cc11d72b2d663f32fce0d18582ae0fe8cc49686085cfa5b65"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7e033fec7e6c8ef86b490b3af99ca5fa04481820ed744ca516b611a085742e22"></a>

## custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / a9e38d7fa396 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator](resources--voltstack_site--reference--group-007.md#canonical-de9ccfeb4756e83ff6cc51ab35fb939749876ea24008d3dcd033ebad9d591290)
- custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays

<a id="canonical-ddd68f2780dc3006598f43b36481b5ebb38900efbbea73102d448a4025922b69"></a>

Type: `"object"`. single nested block, Optional.

Arrays Configuration. Device configuration for PSO Arrays.

Upstream description:

Device configuration for PSO Arrays.

Receipt-pinned upstream constraints:

```json
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
arrays {
  # Configure direct properties listed below.
}
```

<a id="canonical-49fde2d01b3160402305d999a78e86f48975f6360be21525a9c993efc0556a79"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / a9e38d7fa396 / 3

- [flash_array](resources--voltstack_site--reference--group-007.md#canonical-cbc35f9b467a358eb6bafedd445841bd301be40f065c94072f748f2d456b5639): complete subsection reference.

- [flash_blade](resources--voltstack_site--reference--group-007.md#canonical-e57c3afa5a734e7ced7465483aa31cdf6884ccae71ca5e0ef40f1fe9d7a98397): complete subsection reference.

<a id="canonical-17323215a6c9fb69d24a61e9fcc88f757a1303c850f141ddfe758b1de274aa98"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / a9e38d7fa396 / 4

- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array](resources--voltstack_site--reference--group-007.md#canonical-cbc35f9b467a358eb6bafedd445841bd301be40f065c94072f748f2d456b5639)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](resources--voltstack_site--reference--group-007.md#canonical-e57c3afa5a734e7ced7465483aa31cdf6884ccae71ca5e0ef40f1fe9d7a98397)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator](resources--voltstack_site--reference--group-007.md#canonical-de9ccfeb4756e83ff6cc51ab35fb939749876ea24008d3dcd033ebad9d591290)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-cbc35f9b467a358eb6bafedd445841bd301be40f065c94072f748f2d456b5639"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1c2466fc776e4059bdf383ed5864c63f027f0c8b10c80fb724a6fd1ff123a471"></a>

## custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / d2f2e6f0fdb0 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator](resources--voltstack_site--reference--group-007.md#canonical-de9ccfeb4756e83ff6cc51ab35fb939749876ea24008d3dcd033ebad9d591290)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](resources--voltstack_site--reference--group-007.md#canonical-10e5e15a5e7a3f4cc11d72b2d663f32fce0d18582ae0fe8cc49686085cfa5b65)
- custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array

<a id="canonical-03d6bba40b2c2b1a8ddcb88cc7dc37c206f5b6dd8caa340b34b75637d06adbb7"></a>

Type: `"object"`. single nested block, Optional.

Specify what storage flash arrays should be managed the plugin.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("default_fs_type",
    "flash_arrays",
    "iscsi_login_timeout",
    "san_type")}
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
flash_array {
  # Configure direct properties listed below.
}
```

<a id="canonical-d818626b734d61d5f10190e89d8477278b2a2946dde27666e1d8e3c8acb59cc5"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / d2f2e6f0fdb0 / 3

<a id="canonical-fc807eb189cdfa07c9fc16327781c84cc3f5652822d39be51f01051fb33157b2"></a>

<a id="canonical-6cee1d5fcc94b708e05f93615795bb77d282bf7ad93eaaa1a51033bd60be3614"></a>

## default_fs_opt property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / d2f2e6f0fdb0 / 4

Type: `"string"`. Optional.

Block volume default mkfs OPTIONS. Not recommended to change!

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
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-b5faff3fb2eca1af83350e2608685a0f740e47e5b10eee42ea7f0c6d2879d28e"></a>

<a id="canonical-d43147afda253c1ccac703e216c2bdcabd53733cd782fb7b751eae541103dbdb"></a>

## default_fs_type property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / d2f2e6f0fdb0 / 5

Type: `"string"`. Optional.

\[Enum: xfs|ext4\] Block volume default filesystem type. Not recommended to change!. Possible values
are \`xfs\`, \`ext4\`.

Upstream description:

Block volume default filesystem type. Not recommended to change!

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("xfs",
    "ext4"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "xfs",
    "ext4"
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"xfs\\\",\\\"ext4\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"xfs\\\",\\\"ext4\\\"]"
  }
}
```

<a id="canonical-95c3c047a0d79e0ac79fe2be012aa5bf40add226ba03edc8575b0d448a413322"></a>

<a id="canonical-fac97909d2ee68b825e3a5c414387ad94272741aa396cacef8660c4a108936b3"></a>

## default_mount_opts property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / d2f2e6f0fdb0 / 6

Type: `["list", "string"]`. Optional.

Block volume default filesystem mount OPTIONS. Not recommended to change!

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(4),
}
```

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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-fee60b49f88db53565eacffde198506a2517cf974b360ea4a3c05cf2e177d1d6"></a>

<a id="canonical-15524f3075caceeee1310fb147dac8aa86cf389ee8ebd84ae5dc875c31af3464"></a>

## disable_preempt_attachments property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / d2f2e6f0fdb0 / 7

Type: `"bool"`. Optional.

Disable Preempt Attachments. Enable/Disable attachment preemption!

Upstream description:

Enable/Disable attachment preemption!

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [flash_arrays](resources--voltstack_site--reference--group-007.md#canonical-fc4d372e87b619fd42260d782f283520691f53f6fdc161b967cafb0e81483612): complete subsection reference.

<a id="canonical-1dc3078c77136924fe9e07317def8c3e964ab75e3e339111856179b6296aade6"></a>

<a id="canonical-dd93a7848998f00e5cf1cc1d08b98e4081c5070648130b0f796bb98b617b3435"></a>

## iscsi_login_timeout property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / d2f2e6f0fdb0 / 8

Type: `"number"`. Optional.

ISCSI login timeout in seconds. Not recommended to change!

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 100),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 100,
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
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

<a id="canonical-2008cc5803b9924c1e168245461dc55acb112dd844175ad3f4ef756761289cf1"></a>

<a id="canonical-3ca3b876f616c7507b232414a988afc123846debde716be3296faf669e7d597a"></a>

## san_type property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / d2f2e6f0fdb0 / 9

Type: `"string"`. Optional.

\[Enum: ISCSI|FC\] Block volume access protocol, either ISCSI or FC. Possible values are \`ISCSI\`,
\`FC\`.

Upstream description:

Block volume access protocol, either ISCSI or FC.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ISCSI",
    "FC"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "ISCSI",
    "FC"
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"ISCSI\\\",\\\"FC\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"ISCSI\\\",\\\"FC\\\"]"
  }
}
```

<a id="canonical-7e3581c244170cb037f97010bad4078f5123dc8c8f4f6997c9fdf8b4212ce555"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / d2f2e6f0fdb0 / 10

- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays](resources--voltstack_site--reference--group-007.md#canonical-fc4d372e87b619fd42260d782f283520691f53f6fdc161b967cafb0e81483612)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](resources--voltstack_site--reference--group-007.md#canonical-10e5e15a5e7a3f4cc11d72b2d663f32fce0d18582ae0fe8cc49686085cfa5b65)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-fc4d372e87b619fd42260d782f283520691f53f6fdc161b967cafb0e81483612"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3ce899f5e9496b2c3959c61344f0721f0f401c42966248648637f021c8bbb02f"></a>

## custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / ad06da4259cf / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator](resources--voltstack_site--reference--group-007.md#canonical-de9ccfeb4756e83ff6cc51ab35fb939749876ea24008d3dcd033ebad9d591290)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](resources--voltstack_site--reference--group-007.md#canonical-10e5e15a5e7a3f4cc11d72b2d663f32fce0d18582ae0fe8cc49686085cfa5b65)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array](resources--voltstack_site--reference--group-007.md#canonical-cbc35f9b467a358eb6bafedd445841bd301be40f065c94072f748f2d456b5639)
- custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays

<a id="canonical-f2fb0c24af64c7afcdf099e144b44e7a6f8be5879403f18cff1b05b3811d7085"></a>

Type: `"object"`. list nested block, Optional.

For FlashArrays you must set the 'mgmt\_endpoint' and 'api\_token'.

Upstream description:

For FlashArrays you must set the "mgmt\_endpoint" and "api\_token"

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("mgmt_dns_name",
    "mgmt_ip")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
flash_arrays {
  # Configure direct properties listed below.
}
```

<a id="canonical-49fdf4482ac2bbae2d74e506f4db9ebbab5160640011c6018477c31697d71bd1"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / ad06da4259cf / 3

- [api_token](resources--voltstack_site--reference--group-007.md#canonical-6fc408f703f1b9b56a42ffbf8268178168c8f8592275e7dd06a6095f2fad357e): complete subsection reference.

<a id="canonical-ec31fea978af266e40dd5aa5fff25632a80b0406e5ccb3d589a4436c0b2ab4f8"></a>

<a id="canonical-656dbf6b181e1b924ba4e09854f717fdeacd93a9d5c65dffbda11a55590db587"></a>

## labels property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / ad06da4259cf / 4

Type: `["map", "string"]`. Optional.

Specifies labels optional, and can be any key-value pair for use with the PSO 'fleet' provisioner.

Upstream description:

The labels are optional, and can be any key-value pair for use with the PSO "fleet" provisioner.

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
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-b6f7fce4f2ff4a9657c3ef70b30bd3e9522b7b970e8a1273be609eff9b77bfec"></a>

<a id="canonical-36519f27ab88d155729707dd6a5d2c1d2d227267d66c0e131c708da161f14330"></a>

## mgmt_dns_name property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / ad06da4259cf / 5

Type: `"string"`. Optional.

Exclusive with \[mgmt\_ip\] Management Endpoint's IP address is discovered using DNS name
resolution. The name given here is fully qualified domain name.

Upstream description:

Exclusive with \[mgmt\_ip\] Management Endpoint's IP address is discovered using DNS name
resolution. The name given here is fully qualified domain name.

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

<a id="canonical-b5d7f0b1c0ce21a26f9d47fc7984278bda23cffdeea689b58b1c8fb1b66bf8ca"></a>

<a id="canonical-482a368df118749e3dac76b8ed8a54a88c9f1477940b52020b04eff79641ca8c"></a>

## mgmt_ip property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / ad06da4259cf / 6

Type: `"string"`. Optional.

Exclusive with \[mgmt\_dns\_name\] Management Endpoint is reachable at the given IP address.

Upstream description:

Exclusive with \[mgmt\_dns\_name\] Management Endpoint is reachable at the given IP address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPValidator(),
}
```

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-4849d6aa4ff96326d15d2fbc4ca93a2596cccf6f5cbfe3473c7b61f691291916"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / ad06da4259cf / 7

- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token](resources--voltstack_site--reference--group-007.md#canonical-6fc408f703f1b9b56a42ffbf8268178168c8f8592275e7dd06a6095f2fad357e)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array](resources--voltstack_site--reference--group-007.md#canonical-cbc35f9b467a358eb6bafedd445841bd301be40f065c94072f748f2d456b5639)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-6fc408f703f1b9b56a42ffbf8268178168c8f8592275e7dd06a6095f2fad357e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-84319b28e5b4e1d3551b3373649821030d69602e8ec7c6fe934bf095b07a7e06"></a>

## custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / 99da1c47ebce / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator](resources--voltstack_site--reference--group-007.md#canonical-de9ccfeb4756e83ff6cc51ab35fb939749876ea24008d3dcd033ebad9d591290)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](resources--voltstack_site--reference--group-007.md#canonical-10e5e15a5e7a3f4cc11d72b2d663f32fce0d18582ae0fe8cc49686085cfa5b65)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array](resources--voltstack_site--reference--group-007.md#canonical-cbc35f9b467a358eb6bafedd445841bd301be40f065c94072f748f2d456b5639)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays](resources--voltstack_site--reference--group-007.md#canonical-fc4d372e87b619fd42260d782f283520691f53f6fdc161b967cafb0e81483612)
- custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token

<a id="canonical-2c209bbde4604b962ee472314229872b0421b4bfba316b9a4593a3a3228fa2d5"></a>

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
api_token {
  # Configure direct properties listed below.
}
```

<a id="canonical-648a41a00f067dec261546a303bb3998a72f1ab858988cce8ed7b6f5b5d37f7b"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / 99da1c47ebce / 3

- [blindfold_secret_info](resources--voltstack_site--reference--group-007.md#canonical-fbd3909b778fa3a7f3e6758e79f729b51bc25e07e0a0ed59462fadaad144ae9c): complete subsection reference.

- [clear_secret_info](resources--voltstack_site--reference--group-007.md#canonical-423a4add3990689c5cd9074b81ae98e9a77c9aea60b813d7f21f4f82f1ea3788): complete subsection reference.

<a id="canonical-2867c9e79bea17be646dc435f6f9ef7e4cdcbb6868d3de6bf0e7701953cedf35"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / 99da1c47ebce / 4

- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info](resources--voltstack_site--reference--group-007.md#canonical-fbd3909b778fa3a7f3e6758e79f729b51bc25e07e0a0ed59462fadaad144ae9c)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info](resources--voltstack_site--reference--group-007.md#canonical-423a4add3990689c5cd9074b81ae98e9a77c9aea60b813d7f21f4f82f1ea3788)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays](resources--voltstack_site--reference--group-007.md#canonical-fc4d372e87b619fd42260d782f283520691f53f6fdc161b967cafb0e81483612)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-fbd3909b778fa3a7f3e6758e79f729b51bc25e07e0a0ed59462fadaad144ae9c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ac03ea4d403c0cb02d5e79b7b6cdde65c0664c319945c170b065b1d63f302a64"></a>

## custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / 13fb4837b17a / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator](resources--voltstack_site--reference--group-007.md#canonical-de9ccfeb4756e83ff6cc51ab35fb939749876ea24008d3dcd033ebad9d591290)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](resources--voltstack_site--reference--group-007.md#canonical-10e5e15a5e7a3f4cc11d72b2d663f32fce0d18582ae0fe8cc49686085cfa5b65)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array](resources--voltstack_site--reference--group-007.md#canonical-cbc35f9b467a358eb6bafedd445841bd301be40f065c94072f748f2d456b5639)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays](resources--voltstack_site--reference--group-007.md#canonical-fc4d372e87b619fd42260d782f283520691f53f6fdc161b967cafb0e81483612)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token](resources--voltstack_site--reference--group-007.md#canonical-6fc408f703f1b9b56a42ffbf8268178168c8f8592275e7dd06a6095f2fad357e)
- custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info

<a id="canonical-92aadd9c33d58d9bfecb511cf6209944f162bcf2a7f15bd788b2b15804124958"></a>

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

<a id="canonical-34a979be84f2d0076977eca4d2d22087d2708bc308e21f1d1d3b5cc4c6c105d2"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / 13fb4837b17a / 3

<a id="canonical-329876a02eabec110ba88015d08dcaaccb513d29a59c02ff27bef05500b64f88"></a>

<a id="canonical-c4471fc8a1f5b299ba3e291b84c8552cf528a60fba651ee155cd12e1cab76b8a"></a>

## decryption_provider property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / 13fb4837b17a / 4

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

<a id="canonical-8b58b67dbb571a2e53fb0d4a80eeddf73b333caa9e2afd9aa95fa4374366adf0"></a>

<a id="canonical-600c75b40aaf18af59315aafea944491a76215b4c3b46a312d65e4e6ee01faaf"></a>

## location property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / 13fb4837b17a / 5

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

<a id="canonical-b045efe8cdc408f19c6dfbf78ca9f29114a8133532a0f5e1dbc56d604f93191b"></a>

<a id="canonical-d37e275b0251b9cd5d3c5c2d54bf6bf65bf2344cd35aefbb4973c38e90dfef97"></a>

## store_provider property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / 13fb4837b17a / 6

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

<a id="canonical-f07d0d5e1551322e89955d2305ab1202eb05816611807a8655f9f01ec616e54c"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / 13fb4837b17a / 7

- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token](resources--voltstack_site--reference--group-007.md#canonical-6fc408f703f1b9b56a42ffbf8268178168c8f8592275e7dd06a6095f2fad357e)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-423a4add3990689c5cd9074b81ae98e9a77c9aea60b813d7f21f4f82f1ea3788"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c8218ea49eaaaa1f24d753b0694609fe27667402e19c65a01539393b8dfff967"></a>

## custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / abde7751d867 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator](resources--voltstack_site--reference--group-007.md#canonical-de9ccfeb4756e83ff6cc51ab35fb939749876ea24008d3dcd033ebad9d591290)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](resources--voltstack_site--reference--group-007.md#canonical-10e5e15a5e7a3f4cc11d72b2d663f32fce0d18582ae0fe8cc49686085cfa5b65)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array](resources--voltstack_site--reference--group-007.md#canonical-cbc35f9b467a358eb6bafedd445841bd301be40f065c94072f748f2d456b5639)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays](resources--voltstack_site--reference--group-007.md#canonical-fc4d372e87b619fd42260d782f283520691f53f6fdc161b967cafb0e81483612)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token](resources--voltstack_site--reference--group-007.md#canonical-6fc408f703f1b9b56a42ffbf8268178168c8f8592275e7dd06a6095f2fad357e)
- custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info

<a id="canonical-459a5fc5963e0117482584524e02d7a3595af7e893eaa3dd1e66a19294e24d5e"></a>

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

<a id="canonical-39d728a95dc461453155703ba60c681677e04a6d8ea27d6e3fffe09108864cca"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / abde7751d867 / 3

<a id="canonical-a6b6274afbe910a3637eb2aeb78fcf98bbe45b2b1e18043d939a8c9aecf97035"></a>

<a id="canonical-f6a7065e504de7c6659aa64dee487fb31d5df6a868f05ace7a70185c8897a00d"></a>

## provider_ref property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / abde7751d867 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-682a3524c8974bd91d124e57f3faa5f56f8024c3863fc49067aaed4dddf3ce80"></a>

<a id="canonical-015d3a6ae5c982826836c87d2561eb27378cedf428749e4436af5dcbd5044c4f"></a>

## url property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / abde7751d867 / 5

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

<a id="canonical-669c96bb570dcd852778592ae5b3d2f10a7488717c832d1aacb5b7c2e02d82e9"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / abde7751d867 / 6

- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token](resources--voltstack_site--reference--group-007.md#canonical-6fc408f703f1b9b56a42ffbf8268178168c8f8592275e7dd06a6095f2fad357e)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-e57c3afa5a734e7ced7465483aa31cdf6884ccae71ca5e0ef40f1fe9d7a98397"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e89ad6ceff589d4c89a5600455053e55accab0817ea84b47619e9eee90cfa871"></a>

## custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / fbf818bc55b4 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator](resources--voltstack_site--reference--group-007.md#canonical-de9ccfeb4756e83ff6cc51ab35fb939749876ea24008d3dcd033ebad9d591290)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](resources--voltstack_site--reference--group-007.md#canonical-10e5e15a5e7a3f4cc11d72b2d663f32fce0d18582ae0fe8cc49686085cfa5b65)
- custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade

<a id="canonical-fa6e7ac5d35c6b39645d831132cb8107d551d7fed5881c381823a1d29fbf3558"></a>

Type: `"object"`. single nested block, Optional.

Specify what storage flash blades should be managed the plugin.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("flash_blades")}
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
flash_blade {
  # Configure direct properties listed below.
}
```

<a id="canonical-ea3fbc67772b2a97c6f1337a0777de4215a5c648c25660eea85b02b564848448"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / fbf818bc55b4 / 3

<a id="canonical-1c8bd16efaffa2c1a9d6a605c21860fdcd29f5b6405ba9c53f5c397be4bfea58"></a>

<a id="canonical-722af33877f2d30174c0f80ec93983722882ca59ac3ea80ae5b9107e4f45b0e3"></a>

## enable_snapshot_directory property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / fbf818bc55b4 / 4

Type: `"bool"`. Optional.

Enable Snapshot Directory. Enable/Disable FlashBlade snapshots.

Upstream description:

Enable/Disable FlashBlade snapshots.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-42d2c147a3afb4731be1398b7eabc70be50ae09a75649939bbbac749886a96d6"></a>

<a id="canonical-34ed00a60a5e787961edeb5f32a18e27634c3645ee06fb885f357778f3ccc440"></a>

## export_rules property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / fbf818bc55b4 / 5

Type: `"string"`. Optional.

NFS Export Rules. NFS Export rules.

Upstream description:

NFS Export rules.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 250),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 250,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 250,
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
    "ves.io.schema.rules.string.max_len": "250",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "250",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [flash_blades](resources--voltstack_site--reference--group-007.md#canonical-91b7878b328b931ec30d7b79d830bcf39f838e7379a242aafd1f2da857a31459): complete subsection reference.

<a id="canonical-09a2156219016fb44ac1b2d70e6f43a340306915fdac53701ff5784a5f8ef30c"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / fbf818bc55b4 / 6

- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades](resources--voltstack_site--reference--group-007.md#canonical-91b7878b328b931ec30d7b79d830bcf39f838e7379a242aafd1f2da857a31459)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](resources--voltstack_site--reference--group-007.md#canonical-10e5e15a5e7a3f4cc11d72b2d663f32fce0d18582ae0fe8cc49686085cfa5b65)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-91b7878b328b931ec30d7b79d830bcf39f838e7379a242aafd1f2da857a31459"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2321dc78f70d7a438c426d97f59aa0084d4815e666e3c1f87facb897f031cdcc"></a>

## custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / 9b80276acbed / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator](resources--voltstack_site--reference--group-007.md#canonical-de9ccfeb4756e83ff6cc51ab35fb939749876ea24008d3dcd033ebad9d591290)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](resources--voltstack_site--reference--group-007.md#canonical-10e5e15a5e7a3f4cc11d72b2d663f32fce0d18582ae0fe8cc49686085cfa5b65)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](resources--voltstack_site--reference--group-007.md#canonical-e57c3afa5a734e7ced7465483aa31cdf6884ccae71ca5e0ef40f1fe9d7a98397)
- custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades

<a id="canonical-1132c17ea32011c5e4e8b0dc9ad66285afd200062b6f6f638eeed226a808f744"></a>

Type: `"object"`. list nested block, Optional.

For FlashBlades you must set the 'mgmt\_endpoint', 'api\_token' and nfs\_endpoint.

Upstream description:

For FlashBlades you must set the "mgmt\_endpoint", "api\_token" and nfs\_endpoint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("mgmt_dns_name",
    "mgmt_ip"),
  validators.ConflictingListObjectAttributes("nfs_endpoint_dns_name",
    "nfs_endpoint_ip")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
flash_blades {
  # Configure direct properties listed below.
}
```

<a id="canonical-efc28b0a3b878febf3dbf7415734765828a9f90f03bb77f7e99542cbb31f6814"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / 9b80276acbed / 3

- [api_token](resources--voltstack_site--reference--group-007.md#canonical-7d8686112dc91045553727ed8ef6ba38af48e78ea4ec5d304002be06c05636d5): complete subsection reference.

<a id="canonical-f326704afe2bfa11fea77ff7fb53e3f42917764cca3133d5c326a34ff90bba29"></a>

<a id="canonical-e89a7bd01ad39083ca3fa731b3f754ab161fc214f1a766f56e8a5acecfdd21b3"></a>

## labels property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / 9b80276acbed / 4

Type: `["map", "string"]`. Optional.

Specifies labels optional, and can be any key-value pair for use with the PSO 'fleet' provisioner.

Upstream description:

The labels are optional, and can be any key-value pair for use with the PSO "fleet" provisioner.

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
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-92375abca0488eb2868f09ac57be4b6ddbe6ed6fc29323aaea8eb76af713749c"></a>

<a id="canonical-1092154ffab08ff39f1339c07c675ac9bb95705886e3f2fae798a7b0c7c5d8c3"></a>

## mgmt_dns_name property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / 9b80276acbed / 5

Type: `"string"`. Optional.

Exclusive with \[mgmt\_ip\] Management Endpoint's IP address is discovered using DNS name
resolution. The name given here is fully qualified domain name.

Upstream description:

Exclusive with \[mgmt\_ip\] Management Endpoint's IP address is discovered using DNS name
resolution. The name given here is fully qualified domain name.

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

<a id="canonical-47ed1eff5eabbb605b3dcce7ff962c5972fd168b5862a8832721754b1d00a5fc"></a>

<a id="canonical-0e503bd6d55ae40b0d07b1f5bf17d3bc3aab9b2d1485add4bb54e5544efa891b"></a>

## mgmt_ip property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / 9b80276acbed / 6

Type: `"string"`. Optional.

Exclusive with \[mgmt\_dns\_name\] Management Endpoint is reachable at the given IP address.

Upstream description:

Exclusive with \[mgmt\_dns\_name\] Management Endpoint is reachable at the given IP address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPValidator(),
}
```

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-8a5890383de0f11e00e6b7604213972cce52540d284cc950c112017a08100958"></a>

<a id="canonical-c27de089dee470df0ac80601d636af815eae19f51c3b1c568c5331b943f53843"></a>

## nfs_endpoint_dns_name property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / 9b80276acbed / 7

Type: `"string"`. Optional.

Exclusive with \[nfs\_endpoint\_ip\] Endpoint's IP address is discovered using DNS name resolution.
The name given here is fully qualified domain name.

Upstream description:

Exclusive with \[nfs\_endpoint\_ip\] Endpoint's IP address is discovered using DNS name resolution.
The name given here is fully qualified domain name.

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

<a id="canonical-d2206fb8a370d9a4e00b22cd7777940d88aecc65156a647206a38ee0632fbdfc"></a>

<a id="canonical-e113819d2b6656e6263ca413aa8535577ef9dd72a87678e9b30f80109aa7b016"></a>

## nfs_endpoint_ip property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / 9b80276acbed / 8

Type: `"string"`. Optional.

Exclusive with \[nfs\_endpoint\_dns\_name\] Endpoint is reachable at the given IP address.

Upstream description:

Exclusive with \[nfs\_endpoint\_dns\_name\] Endpoint is reachable at the given IP address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPValidator(),
}
```

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-2be104491fc99ad7a51e0d6fe7bb64732cf46a89aeca1e2e8b14bf7d8bed54ae"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / 9b80276acbed / 9

- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token](resources--voltstack_site--reference--group-007.md#canonical-7d8686112dc91045553727ed8ef6ba38af48e78ea4ec5d304002be06c05636d5)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](resources--voltstack_site--reference--group-007.md#canonical-e57c3afa5a734e7ced7465483aa31cdf6884ccae71ca5e0ef40f1fe9d7a98397)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-7d8686112dc91045553727ed8ef6ba38af48e78ea4ec5d304002be06c05636d5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-355ad66ca41f4aa3c63726333ea8ab395c9d25bb9152bf9f9f65f38db5312115"></a>

## custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / 9cce73560f7a / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator](resources--voltstack_site--reference--group-007.md#canonical-de9ccfeb4756e83ff6cc51ab35fb939749876ea24008d3dcd033ebad9d591290)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](resources--voltstack_site--reference--group-007.md#canonical-10e5e15a5e7a3f4cc11d72b2d663f32fce0d18582ae0fe8cc49686085cfa5b65)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](resources--voltstack_site--reference--group-007.md#canonical-e57c3afa5a734e7ced7465483aa31cdf6884ccae71ca5e0ef40f1fe9d7a98397)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades](resources--voltstack_site--reference--group-007.md#canonical-91b7878b328b931ec30d7b79d830bcf39f838e7379a242aafd1f2da857a31459)
- custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token

<a id="canonical-26de5ea5bef1f81a8db31948ce60cca2c226fea38b684615129b9f3fadbf1c5b"></a>

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
api_token {
  # Configure direct properties listed below.
}
```

<a id="canonical-f4704f2935b7aba0e82a16b115b8e4bbf4ac26b0fdca1953ff9a53de0ca800f8"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / 9cce73560f7a / 3

- [blindfold_secret_info](resources--voltstack_site--reference--group-007.md#canonical-5eab9f4134b79435beea18abfea855d5a585b88d2ff1af34a5ce019976e56199): complete subsection reference.

- [clear_secret_info](resources--voltstack_site--reference--group-007.md#canonical-3b5cf2f3495d944499c893e7600f86cd8b5ec26d18407f2cd0669d83fb6aa4ec): complete subsection reference.

<a id="canonical-fe5c4705af4d083836661d347f93ff4cd429ef5c0c2f464d27917f22a4ed0e7a"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / 9cce73560f7a / 4

- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info](resources--voltstack_site--reference--group-007.md#canonical-5eab9f4134b79435beea18abfea855d5a585b88d2ff1af34a5ce019976e56199)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info](resources--voltstack_site--reference--group-007.md#canonical-3b5cf2f3495d944499c893e7600f86cd8b5ec26d18407f2cd0669d83fb6aa4ec)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades](resources--voltstack_site--reference--group-007.md#canonical-91b7878b328b931ec30d7b79d830bcf39f838e7379a242aafd1f2da857a31459)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-5eab9f4134b79435beea18abfea855d5a585b88d2ff1af34a5ce019976e56199"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b44bd7dd82e7f66b1036b2151e80e9f1c154695af859f2ee4248bbfa2771e1fe"></a>

## custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / 95e9ccbff7cd / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator](resources--voltstack_site--reference--group-007.md#canonical-de9ccfeb4756e83ff6cc51ab35fb939749876ea24008d3dcd033ebad9d591290)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](resources--voltstack_site--reference--group-007.md#canonical-10e5e15a5e7a3f4cc11d72b2d663f32fce0d18582ae0fe8cc49686085cfa5b65)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](resources--voltstack_site--reference--group-007.md#canonical-e57c3afa5a734e7ced7465483aa31cdf6884ccae71ca5e0ef40f1fe9d7a98397)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades](resources--voltstack_site--reference--group-007.md#canonical-91b7878b328b931ec30d7b79d830bcf39f838e7379a242aafd1f2da857a31459)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token](resources--voltstack_site--reference--group-007.md#canonical-7d8686112dc91045553727ed8ef6ba38af48e78ea4ec5d304002be06c05636d5)
- custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info

<a id="canonical-ffa76222fcf798e39a958bb45142eaa352a14ee911a5d871025d538f4b93fcec"></a>

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

<a id="canonical-bbf5a0beb0e0fcc0295e1f0dd99150552a8280508b189dca0660900512f15d2b"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / 95e9ccbff7cd / 3

<a id="canonical-33d97b483ace47b3cdd33a669a88dfa5bd90b30df7b67a4504b21c4af83af7d8"></a>

<a id="canonical-f5b71ad8b12393f0f38ddb31bc530c4becdad2d4fe4aad1f350069583ec36985"></a>

## decryption_provider property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / 95e9ccbff7cd / 4

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

<a id="canonical-e22321fe0d8b10b3b3e02541d93fee3efff84d4c49c9a771e477f6c9efd24efd"></a>

<a id="canonical-c4b532c59cce66836f697b8917092cd64e94ace3a01bef7b19ad104d702be4bf"></a>

## location property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / 95e9ccbff7cd / 5

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

<a id="canonical-40f9b98c299466ab9dee1d53e5e8ee8164dbf97bf8cb051318a64c15e45d46fb"></a>

<a id="canonical-c84f999020e4f62c33708572e7ce386a2436b191ce49e87ecd0973190594d1d7"></a>

## store_provider property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / 95e9ccbff7cd / 6

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

<a id="canonical-4724a1cea55af85a2604930043ee8626379db9a53113c5a6f310707cf906381c"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / 95e9ccbff7cd / 7

- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token](resources--voltstack_site--reference--group-007.md#canonical-7d8686112dc91045553727ed8ef6ba38af48e78ea4ec5d304002be06c05636d5)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-3b5cf2f3495d944499c893e7600f86cd8b5ec26d18407f2cd0669d83fb6aa4ec"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1f26bc97b2d229f035b28e1478bbabab35d00bb4ba130e7b0c81c0879d75f428"></a>

## custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / dc2f0499e126 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator](resources--voltstack_site--reference--group-007.md#canonical-de9ccfeb4756e83ff6cc51ab35fb939749876ea24008d3dcd033ebad9d591290)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](resources--voltstack_site--reference--group-007.md#canonical-10e5e15a5e7a3f4cc11d72b2d663f32fce0d18582ae0fe8cc49686085cfa5b65)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](resources--voltstack_site--reference--group-007.md#canonical-e57c3afa5a734e7ced7465483aa31cdf6884ccae71ca5e0ef40f1fe9d7a98397)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades](resources--voltstack_site--reference--group-007.md#canonical-91b7878b328b931ec30d7b79d830bcf39f838e7379a242aafd1f2da857a31459)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token](resources--voltstack_site--reference--group-007.md#canonical-7d8686112dc91045553727ed8ef6ba38af48e78ea4ec5d304002be06c05636d5)
- custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info

<a id="canonical-4f118ac780be07d07c9966a93ad795982c1ae75f680531ca9f68f1c636b19f82"></a>

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

<a id="canonical-75d644354eb7f18d33862416fd1f953d469c57d422e19e0b4df6bcd3cd2e4d44"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / dc2f0499e126 / 3

<a id="canonical-9313699ccf4799e16a510ee4e3512392ec84207ac3b0a2b4d9169895cd3e997e"></a>

<a id="canonical-707a28c086c48105e51207e9d46f125d5c8055a1a3553c5c9b0b06ca4cbf678a"></a>

## provider_ref property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / dc2f0499e126 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-5773172efa8567013add6948f4884382e03c32f0e9cc7277192c25c75cc03e2e"></a>

<a id="canonical-98a131ac0ca1522a11b81e8c3b52b13598d49e9d80e55c4a3a040a6d3d296ed2"></a>

## url property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / dc2f0499e126 / 5

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

<a id="canonical-ed40cee1f5cd0864a54c6264c37e2dcaaa01216abef83926fc6e48c5246410e2"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / dc2f0499e126 / 6

- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token](resources--voltstack_site--reference--group-007.md#canonical-7d8686112dc91045553727ed8ef6ba38af48e78ea4ec5d304002be06c05636d5)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-b08755996fc0dd1370f5b1a65cec228ba0784ab7c0614c94d52d969da31ee93f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ba576540deaf8b3559dffc81afa8b9506dce11a14a938d6d68abd9f93b086722"></a>

## custom_storage_config.storage_interface_list — custom_storage_config.storage_interface_list / f9cab0a8e4f5 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- custom_storage_config.storage_interface_list

<a id="canonical-7bd52d28e8a6901fe958f79121898834fcf8653fae186c36d30356747f111dbf"></a>

Type: `"object"`. single nested block, Optional.

Configure storage interfaces for this App Stack site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("storage_interfaces")}
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
storage_interface_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-94f144f20aad4b708c139bb357abfaf544669d73bed5cf35c8b7f2f71b023725"></a>

## Direct properties — custom_storage_config.storage_interface_list / f9cab0a8e4f5 / 3

- [storage_interfaces](resources--voltstack_site--reference--group-007.md#canonical-bf287d05b921bc8d012366ba095e45ba24f608b6090f78573cb548333e6d0adf): complete subsection reference.

<a id="canonical-40657bcba2fc8ae258c0585524876d374a509f6d422fd7f0330ba4f4a75dd842"></a>

## Next pages — custom_storage_config.storage_interface_list / f9cab0a8e4f5 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces](resources--voltstack_site--reference--group-007.md#canonical-bf287d05b921bc8d012366ba095e45ba24f608b6090f78573cb548333e6d0adf)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-bf287d05b921bc8d012366ba095e45ba24f608b6090f78573cb548333e6d0adf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2cc63ae29813a9e4ec1796fb648eeff44b71a3eaa5aebd80ea51dc9e687dd464"></a>

## custom_storage_config.storage_interface_list.storage_interfaces — custom_storage_config.storage_interface_list.storage_interfaces / aaf0e0837129 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_interface_list](resources--voltstack_site--reference--group-007.md#canonical-b08755996fc0dd1370f5b1a65cec228ba0784ab7c0614c94d52d969da31ee93f)
- custom_storage_config.storage_interface_list.storage_interfaces

<a id="canonical-33a8a61490cab053bfc97fceb8dc5ce824d27f61f6d20591ea085f31a4bde7aa"></a>

Type: `"object"`. list nested block, Optional.

Configure storage interfaces for this App Stack site.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
storage_interfaces {
  # Configure direct properties listed below.
}
```

<a id="canonical-7ab4b4c61b724496d412d25328b360955bc20ce28dac2237b15a8bfb330816ea"></a>

## Direct properties — custom_storage_config.storage_interface_list.storage_interfaces / aaf0e0837129 / 3

<a id="canonical-446e324aef73085451e33c7b9b07d7b9e2613a94d7574aa2bc01e05f3d01e285"></a>

<a id="canonical-b09965221bbf49c78353d6dc37ace4307b1eec25ebad53207bc8ab558383ea96"></a>

## description_spec property — custom_storage_config.storage_interface_list.storage_interfaces / aaf0e0837129 / 4

Type: `"string"`. Optional.

Interface Description. Description for this Interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

- [labels](resources--voltstack_site--reference--group-007.md#canonical-c8f2fd6f471d3669277fcdd7d9894542f32548353b1dbb3bd7784d31ee398bcd): complete subsection reference.

- [storage_interface](resources--voltstack_site--reference--group-007.md#canonical-4b5f5157b0b360ef7220411b551874fb7f25ef49cf6925bd490169d12cc557c1): complete subsection reference.

<a id="canonical-daa2bb449afa346d05501b65339df53c823092624c9b7056af8e6b3d7e37cdf4"></a>

## Next pages — custom_storage_config.storage_interface_list.storage_interfaces / aaf0e0837129 / 5

- [custom_storage_config.storage_interface_list.storage_interfaces.labels](resources--voltstack_site--reference--group-007.md#canonical-c8f2fd6f471d3669277fcdd7d9894542f32548353b1dbb3bd7784d31ee398bcd)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](resources--voltstack_site--reference--group-007.md#canonical-4b5f5157b0b360ef7220411b551874fb7f25ef49cf6925bd490169d12cc557c1)
- [custom_storage_config.storage_interface_list](resources--voltstack_site--reference--group-007.md#canonical-b08755996fc0dd1370f5b1a65cec228ba0784ab7c0614c94d52d969da31ee93f)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-c8f2fd6f471d3669277fcdd7d9894542f32548353b1dbb3bd7784d31ee398bcd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-10df94e8cfe1b5647201a34c178d95dc8195da099e21c30921c2c80d5e8156bf"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.labels — custom_storage_config.storage_interface_list.storage_interfaces.labels / e9e34b5a48c9 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_interface_list](resources--voltstack_site--reference--group-007.md#canonical-b08755996fc0dd1370f5b1a65cec228ba0784ab7c0614c94d52d969da31ee93f)
- [custom_storage_config.storage_interface_list.storage_interfaces](resources--voltstack_site--reference--group-007.md#canonical-bf287d05b921bc8d012366ba095e45ba24f608b6090f78573cb548333e6d0adf)
- custom_storage_config.storage_interface_list.storage_interfaces.labels

<a id="canonical-68cc0de482ee21e8ed5569e934d451b94733e8fe3238164fcc92ee3d5525d33f"></a>

Type: `"object"`. single nested block, Optional.

Add Labels for this Interface, these labels can be used in firewall policy.

Receipt-pinned upstream constraints:

```json
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
labels {}
```

<a id="canonical-084a3605485259c01d20b25e5ca0bfaf30d83d7d93789c0de607d146b06c065c"></a>

## Direct properties — custom_storage_config.storage_interface_list.storage_interfaces.labels / e9e34b5a48c9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-73f7a760ffd2fd24f6488609cf13c26e8e455185a099e6c2fd050439e7b2934d"></a>

## Next pages — custom_storage_config.storage_interface_list.storage_interfaces.labels / e9e34b5a48c9 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces](resources--voltstack_site--reference--group-007.md#canonical-bf287d05b921bc8d012366ba095e45ba24f608b6090f78573cb548333e6d0adf)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-4b5f5157b0b360ef7220411b551874fb7f25ef49cf6925bd490169d12cc557c1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eed55b57e4ab9afabcf0ff6733b45bab49c4c05e5e226273e72b95f83ec57074"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface — custom_storage_config.storage_interface_list.storage_interfaces.storage_interfac / 3912a58e37ca / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_interface_list](resources--voltstack_site--reference--group-007.md#canonical-b08755996fc0dd1370f5b1a65cec228ba0784ab7c0614c94d52d969da31ee93f)
- [custom_storage_config.storage_interface_list.storage_interfaces](resources--voltstack_site--reference--group-007.md#canonical-bf287d05b921bc8d012366ba095e45ba24f608b6090f78573cb548333e6d0adf)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface

<a id="canonical-972ad1c8399c7eb7fd712431530cc6ef44ae11640794917e9778f8f8dd28c9ce"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for storage interface.

Upstream description:

Ethernet Interface Configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("device"),
  validators.ConflictingObjectAttributes("cluster",
    "node"),
  validators.ConflictingObjectAttributes("dhcp_client",
    "dhcp_server"),
  validators.ConflictingObjectAttributes("dhcp_client",
    "static_ip"),
  validators.ConflictingObjectAttributes("dhcp_server",
    "static_ip"),
  validators.ConflictingObjectAttributes("ipv6_auto_config",
    "no_ipv6_address"),
  validators.ConflictingObjectAttributes("ipv6_auto_config",
    "static_ipv6_address"),
  validators.ConflictingObjectAttributes("is_primary",
    "not_primary"),
  validators.ConflictingObjectAttributes("monitor",
    "monitor_disabled"),
  validators.ConflictingObjectAttributes("no_ipv6_address",
    "static_ipv6_address"),
  validators.ConflictingObjectAttributes("site_local_inside_network",
    "site_local_network"),
  validators.ConflictingObjectAttributes("site_local_inside_network",
    "storage_network"),
  validators.ConflictingObjectAttributes("site_local_network",
    "storage_network"),
  validators.ConflictingObjectAttributes("untagged",
    "vlan_id")}
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
  "x-ves-oneof-field-address_choice": "[\"dhcp_client\",\"dhcp_server\",\"static_ip\"]",
  "x-ves-oneof-field-ipv6_address_choice": "[\"ipv6_auto_config\",\"no_ipv6_address\",\"static_ipv6_address\"]",
  "x-ves-oneof-field-monitoring_choice": "[\"monitor\",\"monitor_disabled\"]",
  "x-ves-oneof-field-network_choice": "[\"segment_network\",\"site_local_inside_network\",\"site_local_network\",\"storage_network\"]",
  "x-ves-oneof-field-node_choice": "[\"cluster\",\"node\"]",
  "x-ves-oneof-field-primary_choice": "[\"is_primary\",\"not_primary\"]",
  "x-ves-oneof-field-vlan_choice": "[\"untagged\",\"vlan_id\"]"
}
```

Terraform syntax:

```terraform
storage_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-6d35dfac4496e6f59aa8c728a7721d43a6e9c5957276ea9105f4576889c3fea4"></a>

## Direct properties — custom_storage_config.storage_interface_list.storage_interfaces.storage_interfac / 3912a58e37ca / 3

- [cluster](resources--voltstack_site--reference--group-007.md#canonical-72c772d2e20d365a3a7b2c08514698c888e218b5201ba6d6a29390e00f57b279): complete subsection reference.

<a id="canonical-60f26cd66f642fb87207631075f1bcda1b6c3e5b2311ab4be94aecbfad26a982"></a>

<a id="canonical-7707279c151beb1ac532e91c5130a0e19377923de8f43a7b4faf4e5e7ccc24b5"></a>

## device property — custom_storage_config.storage_interface_list.storage_interfaces.storage_interfac / 3912a58e37ca / 4

Type: `"string"`. Optional.

Interface configuration for the ethernet device.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
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
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [dhcp_client](resources--voltstack_site--reference--group-008.md#canonical-c41c30b61699cdf5ec5b557d9f7867f844012b3a7fb25579d66e19b182c72977): complete subsection reference.

- [dhcp_server](resources--voltstack_site--reference--group-008.md#canonical-42351ab44d60398e2668611fab445c4832736b6368b6cf90be8aa3eab3d497a9): complete subsection reference.

- [ipv6_auto_config](resources--voltstack_site--reference--group-008.md#canonical-ed04c83a22ec1d7ccf65addae42b4ad40540f66767c9f4c074de4d8df4bdaa9d): complete subsection reference.

- [is_primary](resources--voltstack_site--reference--group-008.md#canonical-e233f46c595cc7f4b5037e9682a7dcb551e29bfcce2d2b25d248908259e4c7fd): complete subsection reference.

- [monitor](resources--voltstack_site--reference--group-008.md#canonical-78695bdbbca52ad601659fadf7eff66535b67168aec2870316f845a845289521): complete subsection reference.

- [monitor_disabled](resources--voltstack_site--reference--group-008.md#canonical-74948c6464e130a31bd2bf77a8d2955a24e4516bdf05427dd9c303819fd79d4e): complete subsection reference.

<a id="canonical-a8a900a6f9ebfccc5340e10525ca389e7395e3592b4dd585321af3afae927252"></a>

<a id="canonical-fdd3131e669fe75ad13c2b9edfffeb3e6763675e257dd09be06903cd298b357b"></a>

## mtu property — custom_storage_config.storage_interface_list.storage_interfaces.storage_interfac / 3912a58e37ca / 5

Type: `"number"`. Optional.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

Upstream description:

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  validators.Int64RangeSetValidator(
    validators.Int64Range{Minimum: 0, Maximum: 0},
    validators.Int64Range{Minimum: 512, Maximum: 9000},
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
    "maximum": 9000,
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
    "ves.io.schema.rules.uint32.ranges": "0,512-9000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,512-9000"
  }
}
```

- [no_ipv6_address](resources--voltstack_site--reference--group-008.md#canonical-7b536a807e687aa57dac244178981854b631c44c52bd3db5536d1baa7b4c56ec): complete subsection reference.

<a id="canonical-43074b8b11590b1d104019cec05a8655b3bd4fcca479f3f30cd1f29c7ddc3f80"></a>

<a id="canonical-5b2f04b8dd02b1d73b4f7694ddce204b3c217723ff0750c8550e324fc576b9f7"></a>

## node property — custom_storage_config.storage_interface_list.storage_interfaces.storage_interfac / 3912a58e37ca / 6

Type: `"string"`. Optional.

Exclusive with \[cluster\] Configuration will apply to a device on the given node.

Upstream description:

Exclusive with \[cluster\] Configuration will apply to a device on the given node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
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
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [not_primary](resources--voltstack_site--reference--group-008.md#canonical-8c072367f1eeb349045b99160eff276e367d485425707affb2dc11ad805b6c41): complete subsection reference.

<a id="canonical-841ac7de64ec62ae8e4143d3fe611501924216b863af9f4adfbdf2e81c0da2ae"></a>

<a id="canonical-a738242cb2149b48b0324a5270daa8f48648ede6206514b40022b0a33ab2fb13"></a>

## priority property — custom_storage_config.storage_interface_list.storage_interfaces.storage_interfac / 3912a58e37ca / 7

Type: `"number"`. Optional.

Priority of the network interface when multiple network interfaces are present in outside network
Greater the value, higher the priority.

Upstream description:

Priority of the network interface when multiple network interfaces are present in outside network
Greater the value, higher the priority.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 255,
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
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

- [site_local_inside_network](resources--voltstack_site--reference--group-008.md#canonical-897be0fab3772adc1eb2714cf5c98904cbd5bf19f3ac1b999d7d01f1059b38e1): complete subsection reference.

- [site_local_network](resources--voltstack_site--reference--group-008.md#canonical-b80b9c6d38e7be998f5c943fa1872add6a441ff52cc78c558bbaa102506e6bc0): complete subsection reference.

- [static_ip](resources--voltstack_site--reference--group-008.md#canonical-9f3d104bbb8c26bcdf8a3b9eed0197735f69dd880beab2f21345ce60b2a6ba75): complete subsection reference.

- [static_ipv6_address](resources--voltstack_site--reference--group-008.md#canonical-e0c288edcbf3f2bfeaaabb09a0c7a36882f2b2d6e936ee90ab993855ff239937): complete subsection reference.

- [storage_network](resources--voltstack_site--reference--group-008.md#canonical-0e33e60da28a142dba8ab77d9e1511ccc83e80388d5ca51cf964e52f1179b44c): complete subsection reference.

- [untagged](resources--voltstack_site--reference--group-008.md#canonical-2f218dc0f157efa1d4e18e4d9019be1ce84de78ba55f755f2081d6739a6b8cd8): complete subsection reference.

<a id="canonical-30415b4b66fc4acb73b05f1af5c2f55cc46259031217100e91d3f7feb392fc29"></a>

<a id="canonical-c95ee674d65fbf4ac9c5c7b68d226c9a1ac84866e80b5bbd04ca662f2f26225c"></a>

## vlan_id property — custom_storage_config.storage_interface_list.storage_interfaces.storage_interfac / 3912a58e37ca / 8

Type: `"number"`. Optional.

Exclusive with \[untagged\] Configure a VLAN tagged ethernet interface.

Upstream description:

Exclusive with \[untagged\] Configure a VLAN tagged ethernet interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 4095),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 4095,
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "4095"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "4095"
  }
}
```

<a id="canonical-26aeb64d76672fa913300b99822ea1f1212b5402a6ee8fb0d5ff7ed86e12afda"></a>

## Next pages — custom_storage_config.storage_interface_list.storage_interfaces.storage_interfac / 3912a58e37ca / 9

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.cluster](resources--voltstack_site--reference--group-007.md#canonical-72c772d2e20d365a3a7b2c08514698c888e218b5201ba6d6a29390e00f57b279)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_client](resources--voltstack_site--reference--group-008.md#canonical-c41c30b61699cdf5ec5b557d9f7867f844012b3a7fb25579d66e19b182c72977)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server](resources--voltstack_site--reference--group-008.md#canonical-42351ab44d60398e2668611fab445c4832736b6368b6cf90be8aa3eab3d497a9)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config](resources--voltstack_site--reference--group-008.md#canonical-ed04c83a22ec1d7ccf65addae42b4ad40540f66767c9f4c074de4d8df4bdaa9d)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.is_primary](resources--voltstack_site--reference--group-008.md#canonical-e233f46c595cc7f4b5037e9682a7dcb551e29bfcce2d2b25d248908259e4c7fd)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.monitor](resources--voltstack_site--reference--group-008.md#canonical-78695bdbbca52ad601659fadf7eff66535b67168aec2870316f845a845289521)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.monitor_disabled](resources--voltstack_site--reference--group-008.md#canonical-74948c6464e130a31bd2bf77a8d2955a24e4516bdf05427dd9c303819fd79d4e)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.no_ipv6_address](resources--voltstack_site--reference--group-008.md#canonical-7b536a807e687aa57dac244178981854b631c44c52bd3db5536d1baa7b4c56ec)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.not_primary](resources--voltstack_site--reference--group-008.md#canonical-8c072367f1eeb349045b99160eff276e367d485425707affb2dc11ad805b6c41)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.site_local_inside_network](resources--voltstack_site--reference--group-008.md#canonical-897be0fab3772adc1eb2714cf5c98904cbd5bf19f3ac1b999d7d01f1059b38e1)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.site_local_network](resources--voltstack_site--reference--group-008.md#canonical-b80b9c6d38e7be998f5c943fa1872add6a441ff52cc78c558bbaa102506e6bc0)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ip](resources--voltstack_site--reference--group-008.md#canonical-9f3d104bbb8c26bcdf8a3b9eed0197735f69dd880beab2f21345ce60b2a6ba75)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ipv6_address](resources--voltstack_site--reference--group-008.md#canonical-e0c288edcbf3f2bfeaaabb09a0c7a36882f2b2d6e936ee90ab993855ff239937)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.storage_network](resources--voltstack_site--reference--group-008.md#canonical-0e33e60da28a142dba8ab77d9e1511ccc83e80388d5ca51cf964e52f1179b44c)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.untagged](resources--voltstack_site--reference--group-008.md#canonical-2f218dc0f157efa1d4e18e4d9019be1ce84de78ba55f755f2081d6739a6b8cd8)
- [custom_storage_config.storage_interface_list.storage_interfaces](resources--voltstack_site--reference--group-007.md#canonical-bf287d05b921bc8d012366ba095e45ba24f608b6090f78573cb548333e6d0adf)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-72c772d2e20d365a3a7b2c08514698c888e218b5201ba6d6a29390e00f57b279"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
