---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-50e1a7495c6733405931b6e1aa64955b3f7b7e9367d434a0681dd464a39d319a"></a>

## Property reference — Property reference / 352114ffb00f / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- Property reference

<a id="canonical-88f66859ada1c3399414be1e3dd60735ca28b8797178b6dc0dbf2071d0a72e75"></a>

## Direct properties — Property reference / 352114ffb00f / 3

<a id="canonical-917780863495b6065efcc90234b596a8c743a23e5c02434f02064fc427c70a0a"></a>

<a id="canonical-4a897745ce856c84b8fced6cabbf0087b22d0b64f4a11a8888c96ca0f5a7c30e"></a>

## annotations property — Property reference / 352114ffb00f / 4

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

<a id="canonical-040bedc227698d8802cdea3aaa7dd1b2c77bb15cbb46e343345d67a47d20b17c"></a>

<a id="canonical-65b133a93e1f71ccf0255bb3cab5ebeb87fe8df18bfa9cdfffd58540f5e47b41"></a>

## description property — Property reference / 352114ffb00f / 5

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

<a id="canonical-357583fd09333e30f36e894673faf179cf8ce7a1d9cd78fab3b86d0589f5b358"></a>

<a id="canonical-f06d4afa51ebe512502ef799d5d41126b536dae22b6a0bf1b838b27e1045579d"></a>

## disable property — Property reference / 352114ffb00f / 6

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

<a id="canonical-7dd6bec24132097fb938a2391b8a3d56474a4c159f589248f609067df6c20d50"></a>

<a id="canonical-76b4b39cdcf06b8b5a7e524535ab8dccb4ef1102851ad027bcb70041b85eb3ab"></a>

## id property — Property reference / 352114ffb00f / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

- [job](resources--workload--reference--group-004.md#canonical-917eb75a2dcdc1639c8e0bf1be98f4f7e4c0d041764379f04763666c3858ea35): complete subsection reference.

<a id="canonical-17b947df499e7bd47576b444a5213f16911306d69d22530d74c4bf8097a8d5c9"></a>

<a id="canonical-ccc446ba8e9cea47fc5e791c9f34126bb4392abe9fbdac9396f987a8bbabf732"></a>

## labels property — Property reference / 352114ffb00f / 8

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

<a id="canonical-ccf8d2ca42212c7e4ea1b0f1b9d76b6e2bb3c39554db2665635950ef2a4efb79"></a>

<a id="canonical-ea67c078e3f704c591e528f29a0907104515fdfc1ee2d9254b03e9f8a084a5cc"></a>

## name property — Property reference / 352114ffb00f / 9

Type: `"string"`. Required.

Name of the Workload. Must be unique within the namespace.

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

<a id="canonical-75daa8efe529ab1a9e2bdda1491c50d510bd37dce5a108fa72a5c9a6f9bc97ef"></a>

<a id="canonical-5e7da54d42e824b7a92393b00267b2e619525f04653acff43464081086f55022"></a>

## namespace property — Property reference / 352114ffb00f / 10

Type: `"string"`. Required.

Namespace where the Workload is created.

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

- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805): complete subsection reference.

- [simple_service](resources--workload--reference--group-016.md#canonical-9361b00a2d9b5209c533e92d6cb37dd367b4eb50395ea573285491ecbf8e6419): complete subsection reference.

- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf): complete subsection reference.

- [timeouts](resources--workload--reference--group-028.md#canonical-923a5edb0b626cfe6c45cb6f5d6d5e0ff2ccab8133ffed2c2cdcdabe8d3e7698): complete subsection reference.
