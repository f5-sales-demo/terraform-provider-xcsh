---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-3e9df636e82937eaf0a9c7aa88aaed8ca6faea2a772b3e9e8f56cdbcdc107f24"></a>

## headers property — simple_service.container.readiness_check.http_health_check / 1f3534cddc65 / 4

Type: `["map", "string"]`. Optional.

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

<a id="canonical-2813accbc19c598a6a69a130aec755e13be55b511af81eecbf7771c81f23f00c"></a>

<a id="canonical-392d624617742414b25c0c0b4cc230bf3048a1fde094644b96b0fd4c571a828f"></a>

## host_header property — simple_service.container.readiness_check.http_health_check / 1f3534cddc65 / 5

Type: `"string"`. Optional.

The value of the host header in the HTTP health check request.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(262),
}
```

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

<a id="canonical-6638ff12cb3f303d8479f86acfeadc2215f209e424145636c79884d6b0cdd3c8"></a>

<a id="canonical-5b19e7c9fa7b7a5e4ebbf2d78521477c87643937540ed50cc17c2587e120d650"></a>

## path property — simple_service.container.readiness_check.http_health_check / 1f3534cddc65 / 6

Type: `"string"`. Optional.

Path. Path to access on the HTTP server.

Upstream description:

Path to access on the HTTP server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 2048),
}
```

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

- [port](resources--workload--reference--group-017.md#canonical-90f3b2dc06e006f06a1686361261890864f496eeec517392f40938fad93c3c20): complete subsection reference.

<a id="canonical-529ea6995742a5b37916c547ea41005a0bb236526e64280188588ddba7289e01"></a>

## Next pages — simple_service.container.readiness_check.http_health_check / 1f3534cddc65 / 7

- [simple_service.container.readiness_check.http_health_check.port](resources--workload--reference--group-017.md#canonical-90f3b2dc06e006f06a1686361261890864f496eeec517392f40938fad93c3c20)
- [simple_service.container.readiness_check](resources--workload--reference--group-016.md#canonical-e664c46fdda4fd0b6f3fc712c7f3e8e7cf0c53ba7b845568a2edbfb7e8c2207b)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-90f3b2dc06e006f06a1686361261890864f496eeec517392f40938fad93c3c20"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-56a9c0b7b8596e2fa6b911f299d06159948944171d314a9bfcf988ddcf800acf"></a>

## simple_service.container.readiness_check.http_health_check.port — simple_service.container.readiness_check.http_health_check.port / 57c6757f62c3 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [simple_service](resources--workload--reference--group-016.md#canonical-9361b00a2d9b5209c533e92d6cb37dd367b4eb50395ea573285491ecbf8e6419)
- [simple_service.container](resources--workload--reference--group-016.md#canonical-45cc02935ab54e5ce267d41971e17998e8f55556582a81b7f71e3e643ef36327)
- [simple_service.container.readiness_check](resources--workload--reference--group-016.md#canonical-e664c46fdda4fd0b6f3fc712c7f3e8e7cf0c53ba7b845568a2edbfb7e8c2207b)
- [simple_service.container.readiness_check.http_health_check](resources--workload--reference--group-016.md#canonical-07cf7bcd54b17cc339f97dc79d951912b363bd34c5b6eaf3cd153e5255693e7a)
- simple_service.container.readiness_check.http_health_check.port

<a id="canonical-a726d2d06e2c7299694da2da9b2b06b4407b1449523145dabe4f32603c1fed8a"></a>

Type: `"object"`. single nested block, Optional.

Port. Port

Upstream description:

Port

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("name",
    "num")}
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
  "x-ves-oneof-field-port_choice": "[\"name\",\"num\"]"
}
```

Terraform syntax:

```terraform
port {
  # Configure direct properties listed below.
}
```

<a id="canonical-77a78cd505672a585f9521d800f9b0967dcc45795de71d1a7e37414084aec5c5"></a>

## Direct properties — simple_service.container.readiness_check.http_health_check.port / 57c6757f62c3 / 3

<a id="canonical-5d0a54ba887b56b80c259f2b7ed6ceae6d37c7a9c91bb7c576ffa6b80b89ae46"></a>

<a id="canonical-675e36ac83955692224b21207f18de01a811a2c5e869699140de5fcbd79cb0e7"></a>

## name property — simple_service.container.readiness_check.http_health_check.port / 57c6757f62c3 / 4

Type: `"string"`. Optional.

Port Name. Exclusive with \[num\] Port Name.

Upstream description:

Exclusive with \[num\] Port Name.

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

<a id="canonical-ac1e779afcba695f5e03a32339e3a4e0ff7c0d8ae12a1819772cc024be87a4a2"></a>

<a id="canonical-298bf881f1fddafbf18d5d7f9bc4d746515c81a1201d9140e0a077fca3fa61e9"></a>

## num property — simple_service.container.readiness_check.http_health_check.port / 57c6757f62c3 / 5

Type: `"number"`. Optional.

Port Number. Exclusive with \[name\] Port number.

Upstream description:

Exclusive with \[name\] Port number.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
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

<a id="canonical-c93869458a24c660555ab0ced3520ca7be6a78b14c56429f78d6b7c28811cada"></a>

## Next pages — simple_service.container.readiness_check.http_health_check.port / 57c6757f62c3 / 6

- [simple_service.container.readiness_check.http_health_check](resources--workload--reference--group-016.md#canonical-07cf7bcd54b17cc339f97dc79d951912b363bd34c5b6eaf3cd153e5255693e7a)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-987a30fd88ab095c6c245d9e41180a6a5df816b1c0a93d984407eb1e24dd2085"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eef2aa8678dec4d447a28d6ea4fd0c7d7c222b040c71359c79a42fe2b10fd955"></a>

## simple_service.container.readiness_check.tcp_health_check — simple_service.container.readiness_check.tcp_health_check / 525ec3e6f5cc / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [simple_service](resources--workload--reference--group-016.md#canonical-9361b00a2d9b5209c533e92d6cb37dd367b4eb50395ea573285491ecbf8e6419)
- [simple_service.container](resources--workload--reference--group-016.md#canonical-45cc02935ab54e5ce267d41971e17998e8f55556582a81b7f71e3e643ef36327)
- [simple_service.container.readiness_check](resources--workload--reference--group-016.md#canonical-e664c46fdda4fd0b6f3fc712c7f3e8e7cf0c53ba7b845568a2edbfb7e8c2207b)
- simple_service.container.readiness_check.tcp_health_check

<a id="canonical-570c1e405f8fd6176715da0c888828d5df2bbd62d688c949f0d8c208cb63be72"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
tcp_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-24b3ab7423ce156d74c56ff2a06bfbf5b1400eb6c1e5be737a973fac89d19c27"></a>

## Direct properties — simple_service.container.readiness_check.tcp_health_check / 525ec3e6f5cc / 3

- [port](resources--workload--reference--group-017.md#canonical-aba3f737d2614900f97f12bd2720ca30c0426a7a2d5c1c0841371dc857c7fbad): complete subsection reference.

<a id="canonical-2f0210036c3914237cf4e12cc865babdc24f9b72a509545edd82913591304c56"></a>

## Next pages — simple_service.container.readiness_check.tcp_health_check / 525ec3e6f5cc / 4

- [simple_service.container.readiness_check.tcp_health_check.port](resources--workload--reference--group-017.md#canonical-aba3f737d2614900f97f12bd2720ca30c0426a7a2d5c1c0841371dc857c7fbad)
- [simple_service.container.readiness_check](resources--workload--reference--group-016.md#canonical-e664c46fdda4fd0b6f3fc712c7f3e8e7cf0c53ba7b845568a2edbfb7e8c2207b)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-aba3f737d2614900f97f12bd2720ca30c0426a7a2d5c1c0841371dc857c7fbad"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1549352cc1ab8d5b2585d771e7878491ab9087c41392febe30793dcbec108495"></a>

## simple_service.container.readiness_check.tcp_health_check.port — simple_service.container.readiness_check.tcp_health_check.port / 9b683a18b495 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [simple_service](resources--workload--reference--group-016.md#canonical-9361b00a2d9b5209c533e92d6cb37dd367b4eb50395ea573285491ecbf8e6419)
- [simple_service.container](resources--workload--reference--group-016.md#canonical-45cc02935ab54e5ce267d41971e17998e8f55556582a81b7f71e3e643ef36327)
- [simple_service.container.readiness_check](resources--workload--reference--group-016.md#canonical-e664c46fdda4fd0b6f3fc712c7f3e8e7cf0c53ba7b845568a2edbfb7e8c2207b)
- [simple_service.container.readiness_check.tcp_health_check](resources--workload--reference--group-017.md#canonical-987a30fd88ab095c6c245d9e41180a6a5df816b1c0a93d984407eb1e24dd2085)
- simple_service.container.readiness_check.tcp_health_check.port

<a id="canonical-8ee258de8623ef67348dadb5594d4eab41daae962c596da7661b07b9575021d7"></a>

Type: `"object"`. single nested block, Optional.

Port. Port

Upstream description:

Port

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("name",
    "num")}
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
  "x-ves-oneof-field-port_choice": "[\"name\",\"num\"]"
}
```

Terraform syntax:

```terraform
port {
  # Configure direct properties listed below.
}
```

<a id="canonical-dc4ed5ea0ca1c8c69926907fd13026366b67369fdee71e86a17416a861d6c0b3"></a>

## Direct properties — simple_service.container.readiness_check.tcp_health_check.port / 9b683a18b495 / 3

<a id="canonical-759340e98eef425244b8a644c1d77f36d92fdb167936251bad9202bd6f0f725d"></a>

<a id="canonical-f9e60ead995caa59dfe4a132507e842a6d2f8229a50790f8b700524bc8450bc2"></a>

## name property — simple_service.container.readiness_check.tcp_health_check.port / 9b683a18b495 / 4

Type: `"string"`. Optional.

Port Name. Exclusive with \[num\] Port Name.

Upstream description:

Exclusive with \[num\] Port Name.

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

<a id="canonical-13e9a2f5ea10a5eca1fbd167ae536ac846e23a8f2dcf1f7b3ef39611926da538"></a>

<a id="canonical-6fd70b520aa3fb3ca24d826c9e6829022ee48d908100ee41b1c2f4fef4550ae0"></a>

## num property — simple_service.container.readiness_check.tcp_health_check.port / 9b683a18b495 / 5

Type: `"number"`. Optional.

Port Number. Exclusive with \[name\] Port number.

Upstream description:

Exclusive with \[name\] Port number.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
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

<a id="canonical-157615349d08a37cf38c9f0ec1109bc3e10f9d70afa5677f0999be8185f455fa"></a>

## Next pages — simple_service.container.readiness_check.tcp_health_check.port / 9b683a18b495 / 6

- [simple_service.container.readiness_check.tcp_health_check](resources--workload--reference--group-017.md#canonical-987a30fd88ab095c6c245d9e41180a6a5df816b1c0a93d984407eb1e24dd2085)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-542f5dc522026b382425cb8b3be92f1ceb6739b98fc9c2c9d8048f424d6bc1c8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a1139a259c23a3059bc7eb0990f7e18370869004c77bed657222def65528c29a"></a>

## simple_service.disabled — simple_service.disabled / 495ab8bf8aae / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [simple_service](resources--workload--reference--group-016.md#canonical-9361b00a2d9b5209c533e92d6cb37dd367b4eb50395ea573285491ecbf8e6419)
- simple_service.disabled

<a id="canonical-7d1005404c88a74f8622b0900378c6afca781f1ce8660ef7d114abb82781692a"></a>

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
disabled = {}
```

<a id="canonical-2737d972c1aafb3ce97999f1428509be429400ba70b2a9e1ce5535d79fe6205c"></a>

## Direct properties — simple_service.disabled / 495ab8bf8aae / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8abb55e383f5aae774384a86ccfe8fb8ef2880641bdd054e5c628968bd7cae3c"></a>

## Next pages — simple_service.disabled / 495ab8bf8aae / 4

- [simple_service](resources--workload--reference--group-016.md#canonical-9361b00a2d9b5209c533e92d6cb37dd367b4eb50395ea573285491ecbf8e6419)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-03c266f90c295a8b4c2aea495e36645eec3865f573b8a8bdd54ffd16f42497ac"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6034a21b378124e481c46032f0968d3bb51afec67392824aabe640dfbc5e075b"></a>

## simple_service.do_not_advertise — simple_service.do_not_advertise / 689105705110 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [simple_service](resources--workload--reference--group-016.md#canonical-9361b00a2d9b5209c533e92d6cb37dd367b4eb50395ea573285491ecbf8e6419)
- simple_service.do_not_advertise

<a id="canonical-e17e10819254cb7fdd0864d23646667b01f2ba070cfaa3b46653e19fe441cc9f"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
do_not_advertise = {}
```

<a id="canonical-f381e88c7e6845a68317a52f3750c97a52940e4d6a3e68633172bbfc7094bf31"></a>

## Direct properties — simple_service.do_not_advertise / 689105705110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6de421ad687f8529b5c38f43796248268dc57ac3594790f12a62f7e0a3b9f377"></a>

## Next pages — simple_service.do_not_advertise / 689105705110 / 4

- [simple_service](resources--workload--reference--group-016.md#canonical-9361b00a2d9b5209c533e92d6cb37dd367b4eb50395ea573285491ecbf8e6419)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-0a63e2c8b641721f1b34e55cf2ff86a96e0ff98f533b2effad0740e534e1930c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7decde091468b1f7366c639037210005ba612ecefdb5bae2c789a9ef021a83a3"></a>

## simple_service.enabled — simple_service.enabled / 70ec62e2a239 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [simple_service](resources--workload--reference--group-016.md#canonical-9361b00a2d9b5209c533e92d6cb37dd367b4eb50395ea573285491ecbf8e6419)
- simple_service.enabled

<a id="canonical-476dbf3b471371b580c58191e2782a9e78f96e0cbb2d3df2c68a99f2c58bfaef"></a>

Type: `"object"`. single nested block, Optional.

Persistent storage volume configuration for the workload.

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
enabled {
  # Configure direct properties listed below.
}
```

<a id="canonical-24fef113ac0b0f3dca7b8da563b5a00a035ccb38cb0e7ab05723625eccc14b38"></a>

## Direct properties — simple_service.enabled / 70ec62e2a239 / 3

<a id="canonical-9a9765f8f8bdd86d270b6f8799b79cbada31f43b52fe50622bf25eaed2514519"></a>

<a id="canonical-8382608e7d289f1cae7470665c03b543e1c8be6d04cc1bda24909d911d1419b8"></a>

## name property — simple_service.enabled / 70ec62e2a239 / 4

Type: `"string"`. Optional.

Name. Name of the volume.

Upstream description:

Name of the volume.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
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
    "pattern": "^[a-z0-9]([-a-z0-9]*[a-z0-9])?$",
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
    "ves.io.schema.rules.string.dns_1123_label": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.dns_1123_label": "true"
  }
}
```

- [persistent_volume](resources--workload--reference--group-017.md#canonical-d783e46a45f7f8c26421d73c38a53bb9c4778d7c56f47baf917ce86e57add0b8): complete subsection reference.

<a id="canonical-6220a5370023f98440a73dfcf66e8039d860cc22505fae56b882d2bc0a339587"></a>

## Next pages — simple_service.enabled / 70ec62e2a239 / 5

- [simple_service.enabled.persistent_volume](resources--workload--reference--group-017.md#canonical-d783e46a45f7f8c26421d73c38a53bb9c4778d7c56f47baf917ce86e57add0b8)
- [simple_service](resources--workload--reference--group-016.md#canonical-9361b00a2d9b5209c533e92d6cb37dd367b4eb50395ea573285491ecbf8e6419)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-d783e46a45f7f8c26421d73c38a53bb9c4778d7c56f47baf917ce86e57add0b8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8f16be02cdcd560dde2d8bdd0ae7cec53f8f8ffe1c2bf5062bdbf2568b370833"></a>

## simple_service.enabled.persistent_volume — simple_service.enabled.persistent_volume / 1dc28e260494 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [simple_service](resources--workload--reference--group-016.md#canonical-9361b00a2d9b5209c533e92d6cb37dd367b4eb50395ea573285491ecbf8e6419)
- [simple_service.enabled](resources--workload--reference--group-017.md#canonical-0a63e2c8b641721f1b34e55cf2ff86a96e0ff98f533b2effad0740e534e1930c)
- simple_service.enabled.persistent_volume

<a id="canonical-593658512ca45d4678919142c4188ddee5351ffde84651b7d7961507c32c10dd"></a>

Type: `"object"`. single nested block, Optional.

Volume containing the Persistent Storage for the workload.

Receipt-pinned upstream constraints:

```json
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
persistent_volume {
  # Configure direct properties listed below.
}
```

<a id="canonical-1ac329f4d11b79d4b1f7d2417747510b91e982d97698f3a30dc04324e5554304"></a>

## Direct properties — simple_service.enabled.persistent_volume / 1dc28e260494 / 3

- [mount](resources--workload--reference--group-017.md#canonical-95637d3ba92fde468ca753f8c366ce7cfbda36caba1f951a1034e4d40ae60873): complete subsection reference.

- [storage](resources--workload--reference--group-017.md#canonical-2636c90c147e9c916570c80f2a52aa977521bed7724742f05141a7c22902b6f4): complete subsection reference.

<a id="canonical-1942e814d1ef41061e8629487d219fd11987e51b4035415d2b1a4c3b243eafa7"></a>

## Next pages — simple_service.enabled.persistent_volume / 1dc28e260494 / 4

- [simple_service.enabled.persistent_volume.mount](resources--workload--reference--group-017.md#canonical-95637d3ba92fde468ca753f8c366ce7cfbda36caba1f951a1034e4d40ae60873)
- [simple_service.enabled.persistent_volume.storage](resources--workload--reference--group-017.md#canonical-2636c90c147e9c916570c80f2a52aa977521bed7724742f05141a7c22902b6f4)
- [simple_service.enabled](resources--workload--reference--group-017.md#canonical-0a63e2c8b641721f1b34e55cf2ff86a96e0ff98f533b2effad0740e534e1930c)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-95637d3ba92fde468ca753f8c366ce7cfbda36caba1f951a1034e4d40ae60873"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fbab0889e6435e3b657337907c7d7c229265bccc942446da5dd59708108760fc"></a>

## simple_service.enabled.persistent_volume.mount — simple_service.enabled.persistent_volume.mount / b870a3053b2b / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [simple_service](resources--workload--reference--group-016.md#canonical-9361b00a2d9b5209c533e92d6cb37dd367b4eb50395ea573285491ecbf8e6419)
- [simple_service.enabled](resources--workload--reference--group-017.md#canonical-0a63e2c8b641721f1b34e55cf2ff86a96e0ff98f533b2effad0740e534e1930c)
- [simple_service.enabled.persistent_volume](resources--workload--reference--group-017.md#canonical-d783e46a45f7f8c26421d73c38a53bb9c4778d7c56f47baf917ce86e57add0b8)
- simple_service.enabled.persistent_volume.mount

<a id="canonical-2839d6e0073e3605fd6b491ca99ad2f6e0a304d03bdbf42ebc062b18ff4a6725"></a>

Type: `"object"`. single nested block, Optional.

Volume mount describes how volume is mounted inside a workload.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("mount_path")}
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
mount {
  # Configure direct properties listed below.
}
```

<a id="canonical-ae9e67b77ca4e245c2ef0d76225bb846b3536e1bd597e62dbea9ec5a37207e00"></a>

## Direct properties — simple_service.enabled.persistent_volume.mount / b870a3053b2b / 3

<a id="canonical-1e126309eb3cbafc3676b8dd9a2ad389016f51877116d773c88d1f1fc66957b7"></a>

<a id="canonical-3ec367e843c70180f3d24d9532a715427792ee5381351a1f1d01b7cd44895790"></a>

## mode property — simple_service.enabled.persistent_volume.mount / b870a3053b2b / 4

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("VOLUME_MOUNT_READ_ONLY",
    "VOLUME_MOUNT_READ_WRITE"),
}
```

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

<a id="canonical-c493b3067824442129db4eb8b905a1afb99daa2597aec12c164dea5efcaa4f6e"></a>

<a id="canonical-82bd11b81b25c2d4f2125731aff3c5abc0a09bca11442cfab70c38feee4bba0d"></a>

## mount_path property — simple_service.enabled.persistent_volume.mount / b870a3053b2b / 5

Type: `"string"`. Optional.

Path within the workload container at which the volume should be mounted. Must not contain ':'.

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

<a id="canonical-25c48ed5f8a2e71714fffb5b8b43436947f321a1e7fa9f75d1fa3c7c70624695"></a>

<a id="canonical-50bb1a22857ba37fa14f770ab5b6603116bbe65e9ff928167be098560541b567"></a>

## sub_path property — simple_service.enabled.persistent_volume.mount / b870a3053b2b / 6

Type: `"string"`. Optional.

Path within the volume from which the workload's volume should be mounted. Defaults to '' (volume's
root).

Upstream description:

Path within the volume from which the workload's volume should be mounted. Defaults to "" (volume's
root).

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

<a id="canonical-c52be4ed857fb97a20cfd931418bfa5f27d9f2cdf1eb15c2845b7cd5bc281a15"></a>

## Next pages — simple_service.enabled.persistent_volume.mount / b870a3053b2b / 7

- [simple_service.enabled.persistent_volume](resources--workload--reference--group-017.md#canonical-d783e46a45f7f8c26421d73c38a53bb9c4778d7c56f47baf917ce86e57add0b8)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-2636c90c147e9c916570c80f2a52aa977521bed7724742f05141a7c22902b6f4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c25cb64f5bdc02c13c39c79cfcf2d0a206a6dde6f4fecc6147aba1e0dfa08fea"></a>

## simple_service.enabled.persistent_volume.storage — simple_service.enabled.persistent_volume.storage / 0e44dfaea7c4 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [simple_service](resources--workload--reference--group-016.md#canonical-9361b00a2d9b5209c533e92d6cb37dd367b4eb50395ea573285491ecbf8e6419)
- [simple_service.enabled](resources--workload--reference--group-017.md#canonical-0a63e2c8b641721f1b34e55cf2ff86a96e0ff98f533b2effad0740e534e1930c)
- [simple_service.enabled.persistent_volume](resources--workload--reference--group-017.md#canonical-d783e46a45f7f8c26421d73c38a53bb9c4778d7c56f47baf917ce86e57add0b8)
- simple_service.enabled.persistent_volume.storage

<a id="canonical-4041e37ac294ff3d5bca962a767a816be0de84df638d3b3495b6c4a72edd6cce"></a>

Type: `"object"`. single nested block, Optional.

Persistent storage configuration is used to configure Persistent Volume Claim (PVC).

Upstream description:

Persistent storage configuration is used to configure Persistent Volume Claim (PVC)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("storage_size"),
  validators.ConflictingObjectAttributes("class_name",
    "default")}
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
  "x-ves-oneof-field-class_name_choice": "[\"class_name\",\"default\"]"
}
```

Terraform syntax:

```terraform
storage {
  # Configure direct properties listed below.
}
```

<a id="canonical-616cb93798df36f308dbfbdb07254c8f3f9dc9aeb73de03c309d9b21e1f45d00"></a>

## Direct properties — simple_service.enabled.persistent_volume.storage / 0e44dfaea7c4 / 3

<a id="canonical-66b2a095f677bcce7d79b23c5ae67ec389b673e91baa0d548dbc25002a29ac71"></a>

<a id="canonical-b1fc76c675c686dbd2da42a84e629d9f970563a7cef0b91abd0792342820ed31"></a>

## access_mode property — simple_service.enabled.persistent_volume.storage / 0e44dfaea7c4 / 4

Type: `"string"`. Optional.

\[Enum:
ACCESS\_MODE\_READ\_WRITE\_ONCE|ACCESS\_MODE\_READ\_WRITE\_MANY|ACCESS\_MODE\_READ\_ONLY\_MANY\]
Persistence storage access mode is used to configure access mode for persistent storage -
ACCESS\_MODE\_READ\_WRITE\_ONCE: Read Write Once Read Write Once is used to mount persistent storage
in read/write mode to exactly 1 host - ACCESS\_MODE\_READ\_WRITE\_MANY: Read Write Many Read Write
Many is used.. Possible values are \`ACCESS\_MODE\_READ\_WRITE\_ONCE\`,
\`ACCESS\_MODE\_READ\_WRITE\_MANY\`, \`ACCESS\_MODE\_READ\_ONLY\_MANY\`. Defaults to
\`ACCESS\_MODE\_READ\_WRITE\_ONCE\`.

Upstream description:

Persistence storage access mode is used to configure access mode for persistent storage

&#8203;- ACCESS\_MODE\_READ\_WRITE\_ONCE: Read Write Once

Read Write Once is used to mount persistent storage in read/write mode to exactly 1 host &#8203;-
ACCESS\_MODE\_READ\_WRITE\_MANY: Read Write Many

Read Write Many is used to mount persistent storage in read/write mode to many hosts &#8203;-
ACCESS\_MODE\_READ\_ONLY\_MANY: Read Only Many

Read Only Many is used to mount persistent storage in read-only mode to many hosts.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ACCESS_MODE_READ_WRITE_ONCE",
    "ACCESS_MODE_READ_WRITE_MANY",
    "ACCESS_MODE_READ_ONLY_MANY"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "ACCESS_MODE_READ_WRITE_ONCE",
  "enum": [
    "ACCESS_MODE_READ_WRITE_ONCE",
    "ACCESS_MODE_READ_WRITE_MANY",
    "ACCESS_MODE_READ_ONLY_MANY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-fb37ed5f0f1dfa2bf3aa237806fa01453ea60d08b72cc86a86c0c25ef399adc3"></a>

<a id="canonical-9c4e8776e5ae6fa75ca58d697b7cca206835945796785ff2c615b45d00653577"></a>

## class_name property — simple_service.enabled.persistent_volume.storage / 0e44dfaea7c4 / 5

Type: `"string"`. Optional.

Exclusive with \[default\] Use the specified class name.

Upstream description:

Exclusive with \[default\] Use the specified class name.

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

- [default](resources--workload--reference--group-017.md#canonical-c649287d693569676a6d997f0cba78e33612356692ebeb092119667fede16290): complete subsection reference.

<a id="canonical-de7c53586bd25f7cfff911940bfd96345dc5a69d1f01341bb6e7df17a9e8e701"></a>

<a id="canonical-06e86d15afde54108c129631822b8f1ee3bb575ec98b52c7898aaae538dbf208"></a>

## storage_size property — simple_service.enabled.persistent_volume.storage / 0e44dfaea7c4 / 6

Type: `"number"`. Optional.

Size (in GiB). Size in GiB of the persistent storage.

Upstream description:

Size in GiB of the persistent storage.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.double.gte": "0.004",
    "ves.io.schema.rules.double.lte": "256",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.double.gte": "0.004",
    "ves.io.schema.rules.double.lte": "256",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-27d47fe43f605bd70e6a742b71e37869ed488515088d527c9040543373b1d35e"></a>

## Next pages — simple_service.enabled.persistent_volume.storage / 0e44dfaea7c4 / 7

- [simple_service.enabled.persistent_volume.storage.default](resources--workload--reference--group-017.md#canonical-c649287d693569676a6d997f0cba78e33612356692ebeb092119667fede16290)
- [simple_service.enabled.persistent_volume](resources--workload--reference--group-017.md#canonical-d783e46a45f7f8c26421d73c38a53bb9c4778d7c56f47baf917ce86e57add0b8)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-c649287d693569676a6d997f0cba78e33612356692ebeb092119667fede16290"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-89c3d002984674c0ce853ee477ec8486055578bfffb95e84aefe7db0b5de7683"></a>

## simple_service.enabled.persistent_volume.storage.default — simple_service.enabled.persistent_volume.storage.default / de99dfc90390 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [simple_service](resources--workload--reference--group-016.md#canonical-9361b00a2d9b5209c533e92d6cb37dd367b4eb50395ea573285491ecbf8e6419)
- [simple_service.enabled](resources--workload--reference--group-017.md#canonical-0a63e2c8b641721f1b34e55cf2ff86a96e0ff98f533b2effad0740e534e1930c)
- [simple_service.enabled.persistent_volume](resources--workload--reference--group-017.md#canonical-d783e46a45f7f8c26421d73c38a53bb9c4778d7c56f47baf917ce86e57add0b8)
- [simple_service.enabled.persistent_volume.storage](resources--workload--reference--group-017.md#canonical-2636c90c147e9c916570c80f2a52aa977521bed7724742f05141a7c22902b6f4)
- simple_service.enabled.persistent_volume.storage.default

<a id="canonical-aabaa2eff68de88533c2e7a0f0f5b7fa8dccd64cf6341df26dce9e58c5c738f8"></a>

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
default = {}
```

<a id="canonical-5b9a9d8db014bbb19d393e85144d1418bb01377562696eca37d7c916df99ebac"></a>

## Direct properties — simple_service.enabled.persistent_volume.storage.default / de99dfc90390 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-47378c9d5dc2a9c0430359b320cb09d6c18a1f05758217f5a55c406e8b4f1f8c"></a>

## Next pages — simple_service.enabled.persistent_volume.storage.default / de99dfc90390 / 4

- [simple_service.enabled.persistent_volume.storage](resources--workload--reference--group-017.md#canonical-2636c90c147e9c916570c80f2a52aa977521bed7724742f05141a7c22902b6f4)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-f4cb6c4627ef2f328600ec5bc5b3170189481279b9545f5640eb0f1fe61e5b22"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e9fdf5a73ea0334f60124cf4ff63f7a68cfe9eadba960081c27e814a3b15e50a"></a>

## simple_service.simple_advertise — simple_service.simple_advertise / 3f0cb9ad5be3 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [simple_service](resources--workload--reference--group-016.md#canonical-9361b00a2d9b5209c533e92d6cb37dd367b4eb50395ea573285491ecbf8e6419)
- simple_service.simple_advertise

<a id="canonical-dc06e9b93e52221f9e85099100eb0a5f648eea6bb3e026ddb62876a96c748792"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for simple advertise.

Upstream description:

Advertise OPTIONS for Simple Service.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("domains",
    "service_port")}
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
simple_advertise {
  # Configure direct properties listed below.
}
```

<a id="canonical-8042db538a5e6e0f1730133034b45af8dc561630a3ec7f70fb6406b1bd56797f"></a>

## Direct properties — simple_service.simple_advertise / 3f0cb9ad5be3 / 3

<a id="canonical-1a181ad41806bce9d2b755fd59a6512efd4088f056306cb398b42478523b21ed"></a>

<a id="canonical-8dafea763829c02b2f6e11a32b3b503f0de78ebb23a248a086a1ac2a1500b841"></a>

## domains property — simple_service.simple_advertise / 3f0cb9ad5be3 / 4

Type: `["list", "string"]`. Optional.

List of Domains (host/authority header) that will be matched to Load Balancer. Wildcard hosts are
supported in the suffix or prefix form Supported Domains and search order: 1. Exact Domain names:
www&#46;example.com. 2.

Upstream description:

A list of Domains (host/authority header) that will be matched to Load Balancer. Wildcard hosts are
supported in the suffix or prefix form

Supported Domains and search order: &#8203;1. Exact Domain names: www&#46;example.com. &#8203;2.
Domains starting with a Wildcard: \*.example.com.

Not supported Domains: &#8203;- Just a Wildcard: \* &#8203;- A Wildcard and TLD with no root Domain:
\*.com. &#8203;- A Wildcard not matching a whole DNS label. E.g. \*.example.com and
\*.bar.example.com are valid Wildcards however \*bar.example.com, \*-bar.example.com, and
bar\*.example.com are all invalid.

Additional notes: A Wildcard will not match empty string. E.g. \*.example.com will match
bar.example.com and baz-bar.example.com but not .example.com. The longest Wildcards match first.
Only a single virtual host in the entire route configuration can match on \*. Also a Domain must be
unique across all virtual hosts within an advertise policy.

Domains are also used for SNI matching if the Load Balancer type is HTTPS. Domains also indicate the
list of names for which DNS resolution will be automatically resolved to IP addresses by the system.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 32),
}
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

<a id="canonical-c00141e18741a47821ea4cbe4d83308128cf9cc92d52793f05cb1bce30429d47"></a>

<a id="canonical-f15a8d08b1802ae4bd9c567d66dc8c110ee78fc61e6fac793e32975e1ecc4e29"></a>

## service_port property — simple_service.simple_advertise / 3f0cb9ad5be3 / 5

Type: `"number"`. Optional.

Service port to advertise on Internet via HTTP loadbalancer using port 80.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1024, 65535),
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
    "minimum": 1024
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1024",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1024",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-00280cf795c50a71b653ed8045ebdda17e667f92d7e795ebe1cd072436a05b44"></a>

## Next pages — simple_service.simple_advertise / 3f0cb9ad5be3 / 6

- [simple_service](resources--workload--reference--group-016.md#canonical-9361b00a2d9b5209c533e92d6cb37dd367b4eb50395ea573285491ecbf8e6419)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4aee9bab2b8e3cba4f06b4ab9a762a7b07c502a53072c08f11a11ec76e9826c1"></a>

## stateful_service — stateful_service / 1c8e4d648c5b / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- stateful_service

<a id="canonical-dde20583724223a21f496bb6678aa9dd98fcca44af835bc3f76c1b54b8dbaca1"></a>

Type: `"object"`. single nested block, Optional.

StatefulService maintains per replica state and each replica has its own persistent storage. Each
replica has a unique network identity and stable storage. Stateful service are used for distributed
stateful applications like cassandra, mongodb, redis, etc.

Upstream description:

StatefulService maintains per replica state and each replica has its own persistent storage. Each
replica has a unique network identity and stable storage. Stateful service are used for distributed
stateful applications like cassandra, mongodb, redis, etc.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("containers",
    "persistent_volumes"),
  validators.ConflictingObjectAttributes("num_replicas",
    "scale_to_zero")}
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
  "x-ves-oneof-field-scaling_choice": "[\"num_replicas\",\"scale_to_zero\"]"
}
```

Terraform syntax:

```terraform
stateful_service {
  # Configure direct properties listed below.
}
```

<a id="canonical-5f8c477e30589839eb7d142b330a6721181714a299d33ce4d485e34839ede1ba"></a>

## Direct properties — stateful_service / 1c8e4d648c5b / 3

- [advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b): complete subsection reference.

- [configuration](resources--workload--reference--group-027.md#canonical-e31b4a03a4f30f5493ac332aaa87bd13b05025c51cfc90f92120fe6852c538ec): complete subsection reference.

- [containers](resources--workload--reference--group-028.md#canonical-e99bf3bb456823083e4129ffac6f5e01d1780828fbcb8d3560f97df96d0905f1): complete subsection reference.

- [deploy_options](resources--workload--reference--group-028.md#canonical-2513a93c02711ec77866b205feb6413e7dd2117ba4c06380bf5e0b29878808c3): complete subsection reference.

<a id="canonical-73569a7d7bfe8566f2e37b6920911a1c8d370e05806cda09af855215db5ced98"></a>

<a id="canonical-4f12bcfb9dca90d4aec88de8c46dd7f27e7e01434772ef47cb73389c61bb9f0d"></a>

## num_replicas property — stateful_service / 1c8e4d648c5b / 4

Type: `"number"`. Optional.

Exclusive with \[scale\_to\_zero\] Number of replicas of service to spawn per site.

Upstream description:

Exclusive with \[scale\_to\_zero\] Number of replicas of service to spawn per site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 5),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 5,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.int32.gt": "0",
    "ves.io.schema.rules.int32.lte": "5"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gt": "0",
    "ves.io.schema.rules.int32.lte": "5"
  }
}
```

- [persistent_volumes](resources--workload--reference--group-028.md#canonical-a203e04acd7f749603be1522b30eccf5c29f4b281dddb74e04ca8905b0310f7e): complete subsection reference.

- [scale_to_zero](resources--workload--reference--group-028.md#canonical-29a60e695fa2ddf2da5fe0b273aff907f4dfd3d5b85fe204b155210cd4f6636c): complete subsection reference.

- [volumes](resources--workload--reference--group-028.md#canonical-38e7ef0ae08e8226d79f01ebfd032f90bd70887814b91e9ee1610bebdab59227): complete subsection reference.

<a id="canonical-13be8a637617da486aec5b7018663d5d493e9e4980d6d3b21462120f440de574"></a>

## Next pages — stateful_service / 1c8e4d648c5b / 5

- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.configuration](resources--workload--reference--group-027.md#canonical-e31b4a03a4f30f5493ac332aaa87bd13b05025c51cfc90f92120fe6852c538ec)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-e99bf3bb456823083e4129ffac6f5e01d1780828fbcb8d3560f97df96d0905f1)
- [stateful_service.deploy_options](resources--workload--reference--group-028.md#canonical-2513a93c02711ec77866b205feb6413e7dd2117ba4c06380bf5e0b29878808c3)
- [stateful_service.persistent_volumes](resources--workload--reference--group-028.md#canonical-a203e04acd7f749603be1522b30eccf5c29f4b281dddb74e04ca8905b0310f7e)
- [stateful_service.scale_to_zero](resources--workload--reference--group-028.md#canonical-29a60e695fa2ddf2da5fe0b273aff907f4dfd3d5b85fe204b155210cd4f6636c)
- [stateful_service.volumes](resources--workload--reference--group-028.md#canonical-38e7ef0ae08e8226d79f01ebfd032f90bd70887814b91e9ee1610bebdab59227)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-84f4f11485d54113c9cacfb5803025450de205c3f99a230df3216aaa661e3c9d"></a>

## stateful_service.advertise_options — stateful_service.advertise_options / 8bd1dec78af3 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- stateful_service.advertise_options

<a id="canonical-ebf0520d3c982cd593766c8334de1b3f7a02543b3135109060256fe3e0329551"></a>

Type: `"object"`. single nested block, Optional.

Advertise OPTIONS are used to configure how and where to advertise the workload using load
balancers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("advertise_custom",
    "advertise_in_cluster"),
  validators.ConflictingObjectAttributes("advertise_custom",
    "advertise_on_public"),
  validators.ConflictingObjectAttributes("advertise_custom",
    "do_not_advertise"),
  validators.ConflictingObjectAttributes("advertise_in_cluster",
    "advertise_on_public"),
  validators.ConflictingObjectAttributes("advertise_in_cluster",
    "do_not_advertise"),
  validators.ConflictingObjectAttributes("advertise_on_public",
    "do_not_advertise")}
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
  "x-ves-oneof-field-advertise_choice": "[\"advertise_custom\",\"advertise_in_cluster\",\"advertise_on_public\",\"do_not_advertise\"]"
}
```

Terraform syntax:

```terraform
advertise_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-03bc495590140c3c336a622030322b80d7226d2e31c30b320faf229d0ad2c98c"></a>

## Direct properties — stateful_service.advertise_options / 8bd1dec78af3 / 3

- [advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02): complete subsection reference.

- [advertise_in_cluster](resources--workload--reference--group-020.md#canonical-a380f5505dd498de6e06d125c0596bf9c21d808f953ae1192a4909b606d18f77): complete subsection reference.

- [advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022): complete subsection reference.

- [do_not_advertise](resources--workload--reference--group-027.md#canonical-e2e6e4fb40b69a4c016aa35e437f910b4287d4c46c70aa350e4215c33f66bbc4): complete subsection reference.

<a id="canonical-d95c51fa57a1843cb88bb69b5ea05bd4f4b33b9bdf3b35fb8abe5adcac70b178"></a>

## Next pages — stateful_service.advertise_options / 8bd1dec78af3 / 4

- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_in_cluster](resources--workload--reference--group-020.md#canonical-a380f5505dd498de6e06d125c0596bf9c21d808f953ae1192a4909b606d18f77)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.do_not_advertise](resources--workload--reference--group-027.md#canonical-e2e6e4fb40b69a4c016aa35e437f910b4287d4c46c70aa350e4215c33f66bbc4)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f2b67ce860aa3342718835d3e731943d58231d7e6c953ffa3342193f9e28bbcf"></a>

## stateful_service.advertise_options.advertise_custom — stateful_service.advertise_options.advertise_custom / 67a50990859a / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- stateful_service.advertise_options.advertise_custom

<a id="canonical-6586d1b64e0e6861e78558c02e9c3f60e1fb6cd1602ca6adfbbbfba9842e9fd6"></a>

Type: `"object"`. single nested block, Optional.

Advertise this workload via loadbalancer on specific sites.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("advertise_where",
    "ports")}
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
advertise_custom {
  # Configure direct properties listed below.
}
```

<a id="canonical-850acdd0c40292d97e31b293f12e60ebb709a2604329fbec1269c2d14b3a93e2"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom / 67a50990859a / 3

- [advertise_where](resources--workload--reference--group-017.md#canonical-62885340e8e99ef20f29c8fcd6414c1b749156dbdaa537b73a279ba1dad7c7f4): complete subsection reference.

- [ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad): complete subsection reference.

<a id="canonical-689ed1bd35e80f9ed42b00c19d54c56975fc5a199a27fe8558ade205c65999fb"></a>

## Next pages — stateful_service.advertise_options.advertise_custom / 67a50990859a / 4

- [stateful_service.advertise_options.advertise_custom.advertise_where](resources--workload--reference--group-017.md#canonical-62885340e8e99ef20f29c8fcd6414c1b749156dbdaa537b73a279ba1dad7c7f4)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-62885340e8e99ef20f29c8fcd6414c1b749156dbdaa537b73a279ba1dad7c7f4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1729289e510648a3362745088a364d6a99a0fa346ca421f1ccf7d456fd5e7758"></a>

## stateful_service.advertise_options.advertise_custom.advertise_where — stateful_service.advertise_options.advertise_custom.advertise_where / 5597eaf58c55 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- stateful_service.advertise_options.advertise_custom.advertise_where

<a id="canonical-c247e3ca548e1cfbd5a945d987c41726da62d1691335ee9d22cf9694e8cf23ff"></a>

Type: `"object"`. list nested block, Optional.

Where should this load balancer be available.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("site",
    "virtual_site"),
  validators.ConflictingListObjectAttributes("site",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("virtual_site",
    "vk8s_service")}
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
advertise_where {
  # Configure direct properties listed below.
}
```

<a id="canonical-f8096bcf9d96192593b7f493e994e536415a2dd0bdce3042cdfc7a910cad296c"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.advertise_where / 5597eaf58c55 / 3

- [site](resources--workload--reference--group-017.md#canonical-ab944c2f0f8e9cf644a09a32891373c42d99e6cdcac73086ccfba9eee9902c1e): complete subsection reference.

- [virtual_site](resources--workload--reference--group-017.md#canonical-0d1db9ca7e3cf1694c653f46488737be119f5506b069786092a8b0b97165f603): complete subsection reference.

- [vk8s_service](resources--workload--reference--group-017.md#canonical-d7850e4e94ea5d72019371a5edbe7df7d281fbec80fef6c77f5392bdc7e1bcc7): complete subsection reference.

<a id="canonical-1a9910a3c56205165e6db79ad72eb6cb6c23b74a157d4cdffecd4a3ba3c6c41c"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.advertise_where / 5597eaf58c55 / 4

- [stateful_service.advertise_options.advertise_custom.advertise_where.site](resources--workload--reference--group-017.md#canonical-ab944c2f0f8e9cf644a09a32891373c42d99e6cdcac73086ccfba9eee9902c1e)
- [stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site](resources--workload--reference--group-017.md#canonical-0d1db9ca7e3cf1694c653f46488737be119f5506b069786092a8b0b97165f603)
- [stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service](resources--workload--reference--group-017.md#canonical-d7850e4e94ea5d72019371a5edbe7df7d281fbec80fef6c77f5392bdc7e1bcc7)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-ab944c2f0f8e9cf644a09a32891373c42d99e6cdcac73086ccfba9eee9902c1e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fde635a7ac9d43068fdf5ebccc3dd65cbe58384d07c02436e7742cba86645e56"></a>

## stateful_service.advertise_options.advertise_custom.advertise_where.site — stateful_service.advertise_options.advertise_custom.advertise_where.site / 37445057a304 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.advertise_where](resources--workload--reference--group-017.md#canonical-62885340e8e99ef20f29c8fcd6414c1b749156dbdaa537b73a279ba1dad7c7f4)
- stateful_service.advertise_options.advertise_custom.advertise_where.site

<a id="canonical-cc41a7e31083f450c691135b99808f2d0d6eebf09d3b3c77d13fb0399524f291"></a>

Type: `"object"`. single nested block, Optional.

Defines a reference to a CE site along with network type and an optional IP address where a load
balancer could be advertised.

Upstream description:

This defines a reference to a CE site along with network type and an optional IP address where a
load balancer could be advertised.

Receipt-pinned upstream constraints:

```json
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
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-677a0899ee155a24a9ab99986d40367b2459880bbb0541b9be1ea720475fd3c5"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.advertise_where.site / 37445057a304 / 3

<a id="canonical-9db72e8140f062458337817d522477607f730ab368c60ddce3f8663bc9331a44"></a>

<a id="canonical-60e98c066e32c315870f650e8a003b30539703923ad1c94375ae48c459df0134"></a>

## ip property — stateful_service.advertise_options.advertise_custom.advertise_where.site / 37445057a304 / 4

Type: `"string"`. Optional.

Use given IP address as VIP on the site.

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

<a id="canonical-3b10107490ff8aab9e26bc5cf2f813027640b07c486f0a175e1763ff45572f3e"></a>

<a id="canonical-2a8fdf38abcdcfda1e9b3fba2aeda725e76084f2609ee80c7d48b944b3718238"></a>

## network property — stateful_service.advertise_options.advertise_custom.advertise_where.site / 37445057a304 / 5

Type: `"string"`. Optional.

\[Enum:
SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE|SITE\_NETWORK\_INSIDE|SITE\_NETWORK\_OUTSIDE|SITE\_NETWORK\_SERVICE|SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_IP\_FABRIC\]
Defines network types to be used on site All inside and outside networks. All inside and outside
networks with internet VIP support. All inside networks. Possible values are
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`, \`SITE\_NETWORK\_INSIDE\`, \`SITE\_NETWORK\_OUTSIDE\`,
\`SITE\_NETWORK\_SERVICE\`, \`SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP\`,
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP\`, \`SITE\_NETWORK\_IP\_FABRIC\`.
Defaults to \`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`.

Upstream description:

This defines network types to be used on site

All inside and outside networks. All inside and outside networks with internet VIP support. All
inside networks. All outside networks. All outside networks with internet VIP support. VK8s service
network. &#8203;- SITE\_NETWORK\_IP\_FABRIC: VER IP Fabric network for the site

This Virtual network type is used for exposing virtual host on IP Fabric network on the VER site or
for endpoint in IP Fabric network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "SITE_NETWORK_INSIDE_AND_OUTSIDE",
  "enum": [
    "SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [site](resources--workload--reference--group-017.md#canonical-bb251972e8fb77b0c32441f496f2da18c60b3a9dc1fd12b8ff1ccbd44b4c8915): complete subsection reference.

<a id="canonical-fde2dcf88620d45092cdadd64450f64c5ef1bc3368fe1e64e0fc589c502c0725"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.advertise_where.site / 37445057a304 / 6

- [stateful_service.advertise_options.advertise_custom.advertise_where.site.site](resources--workload--reference--group-017.md#canonical-bb251972e8fb77b0c32441f496f2da18c60b3a9dc1fd12b8ff1ccbd44b4c8915)
- [stateful_service.advertise_options.advertise_custom.advertise_where](resources--workload--reference--group-017.md#canonical-62885340e8e99ef20f29c8fcd6414c1b749156dbdaa537b73a279ba1dad7c7f4)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-bb251972e8fb77b0c32441f496f2da18c60b3a9dc1fd12b8ff1ccbd44b4c8915"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-663cec56a7bb847dcd7afae10ca6839469e662bd6206e9d3d8c625f047285fac"></a>

## stateful_service.advertise_options.advertise_custom.advertise_where.site.site — stateful_service.advertise_options.advertise_custom.advertise_where.site.site / 5773faf2dd4f / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.advertise_where](resources--workload--reference--group-017.md#canonical-62885340e8e99ef20f29c8fcd6414c1b749156dbdaa537b73a279ba1dad7c7f4)
- [stateful_service.advertise_options.advertise_custom.advertise_where.site](resources--workload--reference--group-017.md#canonical-ab944c2f0f8e9cf644a09a32891373c42d99e6cdcac73086ccfba9eee9902c1e)
- stateful_service.advertise_options.advertise_custom.advertise_where.site.site

<a id="canonical-bc0137845165896a907c6dc9e88eecada640f99bc33cf0e36f399755c5d7a2b6"></a>

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
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-604c584c3a4a49c79c9ce09b264ae660762190801955ebbd3b33b65ab553f80b"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.advertise_where.site.site / 5773faf2dd4f / 3

<a id="canonical-300308518d09d2b731f7dfe599bda907740d8c1b826cda9656a7b2ce0da1c731"></a>

<a id="canonical-8ea7f3c3b8ed8e0ffde55c49dfd925057670d731c575ffc2b553c8b90d7e80f7"></a>

## name property — stateful_service.advertise_options.advertise_custom.advertise_where.site.site / 5773faf2dd4f / 4

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

<a id="canonical-0c392b1efc361a4f8b72dc350b4f1a6fee1b7a97a557cd23c10d343f8c972948"></a>

<a id="canonical-56e22fda5f0d99173476cbfd7f5075ce56d832d06110a104509c662e3571568f"></a>

## namespace property — stateful_service.advertise_options.advertise_custom.advertise_where.site.site / 5773faf2dd4f / 5

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

<a id="canonical-c0125d37906fad66ab14b6023f32bd94fea311c3004a24b1a407d520c0e604a3"></a>

<a id="canonical-6ba0a0020f772e5db7f0dbdd469135b5e58dd55719c7ab757d601226de0248c6"></a>

## tenant property — stateful_service.advertise_options.advertise_custom.advertise_where.site.site / 5773faf2dd4f / 6

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

<a id="canonical-55759b47f8f7274b5657454f55d8c93f1f6f68ae94fd6ea893baa29ce03dfac5"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.advertise_where.site.site / 5773faf2dd4f / 7

- [stateful_service.advertise_options.advertise_custom.advertise_where.site](resources--workload--reference--group-017.md#canonical-ab944c2f0f8e9cf644a09a32891373c42d99e6cdcac73086ccfba9eee9902c1e)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-0d1db9ca7e3cf1694c653f46488737be119f5506b069786092a8b0b97165f603"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fc9fecc93571e08d8cd047baa97bc91b1575905fac7df2feed3ff8e53a1377c9"></a>

## stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site — stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site / 95c4febdf089 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.advertise_where](resources--workload--reference--group-017.md#canonical-62885340e8e99ef20f29c8fcd6414c1b749156dbdaa537b73a279ba1dad7c7f4)
- stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site

<a id="canonical-28fdfaa267a57700e49a0e63ec56f4bce6695f0b6548e6ab1bd3abb90b3d6590"></a>

Type: `"object"`. single nested block, Optional.

Defines a reference to a customer site virtual site along with network type where a load balancer
could be advertised.

Upstream description:

This defines a reference to a customer site virtual site along with network type where a load
balancer could be advertised.

Receipt-pinned upstream constraints:

```json
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
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-e2c07b314838123584103bbf93b830f17cb8b8683e58d55aeeca00e26bd2c75b"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site / 95c4febdf089 / 3

<a id="canonical-0129f253091f0c4ed8fbd7319bbc68487d6463f3e4980d680b70c07f68365182"></a>

<a id="canonical-0df0f2c15f301e0b5af20acff34a40d006b98f3f3bdfff0a52524f3165034104"></a>

## network property — stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site / 95c4febdf089 / 4

Type: `"string"`. Optional.

\[Enum:
SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE|SITE\_NETWORK\_INSIDE|SITE\_NETWORK\_OUTSIDE|SITE\_NETWORK\_SERVICE|SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_IP\_FABRIC\]
Defines network types to be used on site All inside and outside networks. All inside and outside
networks with internet VIP support. All inside networks. Possible values are
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`, \`SITE\_NETWORK\_INSIDE\`, \`SITE\_NETWORK\_OUTSIDE\`,
\`SITE\_NETWORK\_SERVICE\`, \`SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP\`,
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP\`, \`SITE\_NETWORK\_IP\_FABRIC\`.
Defaults to \`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`.

Upstream description:

This defines network types to be used on site

All inside and outside networks. All inside and outside networks with internet VIP support. All
inside networks. All outside networks. All outside networks with internet VIP support. VK8s service
network. &#8203;- SITE\_NETWORK\_IP\_FABRIC: VER IP Fabric network for the site

This Virtual network type is used for exposing virtual host on IP Fabric network on the VER site or
for endpoint in IP Fabric network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "SITE_NETWORK_INSIDE_AND_OUTSIDE",
  "enum": [
    "SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [virtual_site](resources--workload--reference--group-017.md#canonical-7f9a2f2a52ac928f817fe1f4c768ba6a5027b320fd9c534c3f19b371cfada77f): complete subsection reference.

<a id="canonical-890f5a17df5d3c5fdfadf0dddd26baceab9f94bbade99e95bc80c80bfc2c0531"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site / 95c4febdf089 / 5

- [stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site.virtual_site](resources--workload--reference--group-017.md#canonical-7f9a2f2a52ac928f817fe1f4c768ba6a5027b320fd9c534c3f19b371cfada77f)
- [stateful_service.advertise_options.advertise_custom.advertise_where](resources--workload--reference--group-017.md#canonical-62885340e8e99ef20f29c8fcd6414c1b749156dbdaa537b73a279ba1dad7c7f4)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-7f9a2f2a52ac928f817fe1f4c768ba6a5027b320fd9c534c3f19b371cfada77f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7d715ae45b4eb1b5fee5994e6814722a16d96a176e8a489976fb10431e58a41f"></a>

## stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site.virtual_site — stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site / 644e5fe3e6b7 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.advertise_where](resources--workload--reference--group-017.md#canonical-62885340e8e99ef20f29c8fcd6414c1b749156dbdaa537b73a279ba1dad7c7f4)
- [stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site](resources--workload--reference--group-017.md#canonical-0d1db9ca7e3cf1694c653f46488737be119f5506b069786092a8b0b97165f603)
- stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site.virtual_site

<a id="canonical-d6b4af611ae2f876571797047cfcb30a01cea7cb7c83171fb99dd300bca1c40a"></a>

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
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-48cc6f208cf49a43d12fe1932e17dc7bed71b5f6c021a22d6741af7c4f123c91"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site / 644e5fe3e6b7 / 3

<a id="canonical-4daa2e8faeacf09694329ba4e211fee52636900df49389c218d1a304ce17b1c0"></a>

<a id="canonical-79d225eaaeed0a122e9c2f8153555c9da8a39e10b47f32ea599e103412402a0a"></a>

## name property — stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site / 644e5fe3e6b7 / 4

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

<a id="canonical-c790d1878123076c81510f0e3a9620ebbce75a82f9cd3b6c445875efffc91f97"></a>

<a id="canonical-baac68336caa0a16cac37bd413e90bd4a3651365f6487c4186b363b8095ff751"></a>

## namespace property — stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site / 644e5fe3e6b7 / 5

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

<a id="canonical-64e2135f16fd789ce3d0373ef0cd4c678017350e62f87acbcb98601fd7dfc5b6"></a>

<a id="canonical-d58bed512119d0f6bd3f874e1f78cb4b8b74437e665e609c67505a222105c5d4"></a>

## tenant property — stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site / 644e5fe3e6b7 / 6

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

<a id="canonical-9a2dbadd4863ea3128ae73cf4bb0779db4ad7ca3d481662c43fef8bdc1927fcf"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site / 644e5fe3e6b7 / 7

- [stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site](resources--workload--reference--group-017.md#canonical-0d1db9ca7e3cf1694c653f46488737be119f5506b069786092a8b0b97165f603)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-d7850e4e94ea5d72019371a5edbe7df7d281fbec80fef6c77f5392bdc7e1bcc7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6ae75a0288aa975827c6764698457a783528c5a2cf50ebf253a538f5755fcef0"></a>

## stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service — stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service / 9ee8a677a11b / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.advertise_where](resources--workload--reference--group-017.md#canonical-62885340e8e99ef20f29c8fcd6414c1b749156dbdaa537b73a279ba1dad7c7f4)
- stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service

<a id="canonical-932af461249d86e904ef8cca19c0b013ed25e1103906ecd2e20c05463627df32"></a>

Type: `"object"`. single nested block, Optional.

Defines a reference to a RE site or virtual site where a load balancer could be advertised in the
vK8s service network.

Upstream description:

This defines a reference to a RE site or virtual site where a load balancer could be advertised in
the vK8s service network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("site",
    "virtual_site")}
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
  "x-ves-oneof-field-choice": "[\"site\",\"virtual_site\"]"
}
```

Terraform syntax:

```terraform
vk8s_service {
  # Configure direct properties listed below.
}
```

<a id="canonical-fad9519086813acaaae24d090277fb7bd68b159da17e4a698dd473790d1ae53c"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service / 9ee8a677a11b / 3

- [site](resources--workload--reference--group-017.md#canonical-884131970451172a4a8398ecaa70e6d847c02ad6d9e9b10714eecb5f740bae5d): complete subsection reference.

- [virtual_site](resources--workload--reference--group-017.md#canonical-c344dcab0fd75a1243e18a69ad24b6e4fd3af0614f20a14ce15c2b28ee9f3c69): complete subsection reference.

<a id="canonical-60b60382bac62ea3a02dc3c0bd9040ef8a2d04b5e765693c0b5759a95a3d135a"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service / 9ee8a677a11b / 4

- [stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service.site](resources--workload--reference--group-017.md#canonical-884131970451172a4a8398ecaa70e6d847c02ad6d9e9b10714eecb5f740bae5d)
- [stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service.virtual_site](resources--workload--reference--group-017.md#canonical-c344dcab0fd75a1243e18a69ad24b6e4fd3af0614f20a14ce15c2b28ee9f3c69)
- [stateful_service.advertise_options.advertise_custom.advertise_where](resources--workload--reference--group-017.md#canonical-62885340e8e99ef20f29c8fcd6414c1b749156dbdaa537b73a279ba1dad7c7f4)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-884131970451172a4a8398ecaa70e6d847c02ad6d9e9b10714eecb5f740bae5d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fdcdf0a3c5c7bdded2d6d21b220b0d680995b7a08588ce88a4a37c1dc2c02007"></a>

## stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service.site — stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service / 0c9c9bdf7970 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.advertise_where](resources--workload--reference--group-017.md#canonical-62885340e8e99ef20f29c8fcd6414c1b749156dbdaa537b73a279ba1dad7c7f4)
- [stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service](resources--workload--reference--group-017.md#canonical-d7850e4e94ea5d72019371a5edbe7df7d281fbec80fef6c77f5392bdc7e1bcc7)
- stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service.site

<a id="canonical-89e7d61e238fe005b206bc06aaf9cdf9b01a39d0b1cb7bcaff3fe6f2e77a6b9d"></a>

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
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-86a2d345b38d07548ed5105dd25b676b8469dfa72f9d7e341b2de8b4a2ec1b2d"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service / 0c9c9bdf7970 / 3

<a id="canonical-42ce2918ad64b5213a6427424dc79d9a0f3f457ffab35d19534b441989f59da1"></a>

<a id="canonical-3994c05e7ac9e6bbf0ef4b21cd4e433749d0a95b76ce58c6f621331006015c08"></a>

## name property — stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service / 0c9c9bdf7970 / 4

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

<a id="canonical-4eb925711a74e8c1dced9322f58a505680bee556ddba14a2de91ebaa13d06426"></a>

<a id="canonical-04f7f3a5b14349108658332a6fb49b254c561b93251518fe11d97ddb7f98aec4"></a>

## namespace property — stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service / 0c9c9bdf7970 / 5

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

<a id="canonical-933fa145fce2d1b5954c7d9a012ce1813dce0e17dd9ee12c30645cc15bd4ce23"></a>

<a id="canonical-88b1a18b716579d3abf34244ff55bf1ab880ea6d85328db8c56a4042fd201a25"></a>

## tenant property — stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service / 0c9c9bdf7970 / 6

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

<a id="canonical-2ae918082ddfb48e3e67fc400bba2b85f82ee8feafaf27b5eb6f913137f7e062"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service / 0c9c9bdf7970 / 7

- [stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service](resources--workload--reference--group-017.md#canonical-d7850e4e94ea5d72019371a5edbe7df7d281fbec80fef6c77f5392bdc7e1bcc7)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-c344dcab0fd75a1243e18a69ad24b6e4fd3af0614f20a14ce15c2b28ee9f3c69"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7fab3c9d1b064b26762b377ba1a8b31291fd93094fc0badeba6eac54eefa1c43"></a>

## stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service.virtual_site — stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service / 6ba0853bd577 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.advertise_where](resources--workload--reference--group-017.md#canonical-62885340e8e99ef20f29c8fcd6414c1b749156dbdaa537b73a279ba1dad7c7f4)
- [stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service](resources--workload--reference--group-017.md#canonical-d7850e4e94ea5d72019371a5edbe7df7d281fbec80fef6c77f5392bdc7e1bcc7)
- stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service.virtual_site

<a id="canonical-1bc8080dc36f26e2c0cd7bc366ca9a8d72d614f074f2a0b35828637716bb3c02"></a>

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
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-9c9698b492eabac34244c9b32d1cc7f2355f26362cfd03b986382d0315319a7b"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service / 6ba0853bd577 / 3

<a id="canonical-79f0f7c66fa6830316953146aff96cd0c50d674f9ab327b88e33fd4e3ab39094"></a>

<a id="canonical-b20e88b171a5b9b1f3dcab9ff54cd49131120e30fb0971c60b2daacabc496117"></a>

## name property — stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service / 6ba0853bd577 / 4

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

<a id="canonical-53f1b7cdcb7b5a3f924270ecfa6cd927cbfaca2dfa5f0e4217ffd00817e33aeb"></a>

<a id="canonical-15c7b5150f106ff461b2ec2e35a4451ef437130f0fa05a1ba8adb775fa51b05a"></a>

## namespace property — stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service / 6ba0853bd577 / 5

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

<a id="canonical-d300de7e2d72e114a0a51ad6c919180db7cff5063170bfade394c23acbdc221a"></a>

<a id="canonical-ddafd5f269ed35bb4e7ada54d309f640732b9fb777b1d53dda273453ec16d98f"></a>

## tenant property — stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service / 6ba0853bd577 / 6

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

<a id="canonical-a98ee18030660f34ead5fd129d0e8969cf2ba59930f4084304b7cb5c656c520f"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service / 6ba0853bd577 / 7

- [stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service](resources--workload--reference--group-017.md#canonical-d7850e4e94ea5d72019371a5edbe7df7d281fbec80fef6c77f5392bdc7e1bcc7)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-93a24235c62c770180649f05048bb8103f5f79e67718dd0500060d741a4c8122"></a>

## stateful_service.advertise_options.advertise_custom.ports — stateful_service.advertise_options.advertise_custom.ports / 5700d79e27bb / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- stateful_service.advertise_options.advertise_custom.ports

<a id="canonical-94dc43c808a6cca8b8e0a6842d85a0adf70334c2a76308b2ffd7018876b81294"></a>

Type: `"object"`. list nested block, Optional.

Ports. Ports to advertise.

Upstream description:

Ports to advertise.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("http_loadbalancer",
    "tcp_loadbalancer")}
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

Terraform syntax:

```terraform
ports {
  # Configure direct properties listed below.
}
```

<a id="canonical-babcc57e2ddfb5ef175dc6e3acade723808ad6bd744e2732e663569e100dda7f"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports / 5700d79e27bb / 3

- [http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882): complete subsection reference.

- [port](resources--workload--reference--group-020.md#canonical-9346adde0cfe4a128d86ed815352cce79b915dbabfc95db14d85f1bcf36002eb): complete subsection reference.

- [tcp_loadbalancer](resources--workload--reference--group-020.md#canonical-0846336c6bc25f0f9b5ca92bf4fc57d0fca71df4dd26f175e981080630a50f3d): complete subsection reference.

<a id="canonical-44b578139edd011b3f74fdcf436e8c975c3eacdb5452993c38f323e6b966be7c"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports / 5700d79e27bb / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.port](resources--workload--reference--group-020.md#canonical-9346adde0cfe4a128d86ed815352cce79b915dbabfc95db14d85f1bcf36002eb)
- [stateful_service.advertise_options.advertise_custom.ports.tcp_loadbalancer](resources--workload--reference--group-020.md#canonical-0846336c6bc25f0f9b5ca92bf4fc57d0fca71df4dd26f175e981080630a50f3d)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7a83706309a7bd02c7bc434ffd7a35f2b95342ff23d86e54618e5bffceede90b"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer / 59f16827f174 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer

<a id="canonical-ef1e9a5d867fa47fbf4f5a8c70f92666e7ef1033909b498dd21af16fee9c76c2"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for http loadbalancer.

Upstream description:

HTTP/HTTPS Load balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("domains"),
  validators.ConflictingObjectAttributes("default_route",
    "specific_routes"),
  validators.ConflictingObjectAttributes("http",
    "https"),
  validators.ConflictingObjectAttributes("http",
    "https_auto_cert"),
  validators.ConflictingObjectAttributes("https",
    "https_auto_cert")}
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
  "x-ves-oneof-field-loadbalancer_type": "[\"http\",\"https\",\"https_auto_cert\"]",
  "x-ves-oneof-field-route_choice": "[\"default_route\",\"specific_routes\"]"
}
```

Terraform syntax:

```terraform
http_loadbalancer {
  # Configure direct properties listed below.
}
```

<a id="canonical-9d02e0262e6167eadaa4a485c11a183dda2bf17ec4f5905deb3197e56835f9e7"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer / 59f16827f174 / 3

- [default_route](resources--workload--reference--group-017.md#canonical-8d583f1134632f8bd8257e1cd4034ae146003e5cb69f89850ac1930fd06d94e1): complete subsection reference.

<a id="canonical-da08ee8840d71286971bdc24d1518ebe93b56f2d8ade1ee01fdcf73566b408c5"></a>

<a id="canonical-ab3aa38d5990262d7f7da3ea54a586d0bfa339505648469dcc42486f952a202a"></a>

## domains property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer / 59f16827f174 / 4

Type: `["list", "string"]`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 32),
}
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

- [http](resources--workload--reference--group-017.md#canonical-0106c8e4449e963d5280c12451e8503308a4aba76db34acb949ad17779b80dc9): complete subsection reference.

- [https](resources--workload--reference--group-017.md#canonical-1622054a074c74fd54c41bd170891da7f136459269bb1c252d02e71dde5f37ff): complete subsection reference.

- [https_auto_cert](resources--workload--reference--group-019.md#canonical-ce66ed1d7b48dc7c09733c25f662bdffbe9a3d5d17d8942dcd5ecabb4e3fbe11): complete subsection reference.

- [specific_routes](resources--workload--reference--group-019.md#canonical-12085629394a59606f007a8c16c0628b4b5c34522811388cb590c912a04fbeae): complete subsection reference.

<a id="canonical-70477e2d23cfc18a07ebdeb627fb36ad51d8959c5a08185a4ca2002777d00fc2"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer / 59f16827f174 / 5

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route](resources--workload--reference--group-017.md#canonical-8d583f1134632f8bd8257e1cd4034ae146003e5cb69f89850ac1930fd06d94e1)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http](resources--workload--reference--group-017.md#canonical-0106c8e4449e963d5280c12451e8503308a4aba76db34acb949ad17779b80dc9)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-017.md#canonical-1622054a074c74fd54c41bd170891da7f136459269bb1c252d02e71dde5f37ff)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-019.md#canonical-ce66ed1d7b48dc7c09733c25f662bdffbe9a3d5d17d8942dcd5ecabb4e3fbe11)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-019.md#canonical-12085629394a59606f007a8c16c0628b4b5c34522811388cb590c912a04fbeae)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-8d583f1134632f8bd8257e1cd4034ae146003e5cb69f89850ac1930fd06d94e1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e2173f7eb37841d907eb6ba9ba68c6e6bec1f714826c2ae7a5f12ba04ba9d809"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.defa / 1e9f134c49b6 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route

<a id="canonical-86588966827fd3bd31684a796f02fc030299109d24a1beb44f3d128296011fd9"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for default route.

Upstream description:

Default route matching all APIs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("auto_host_rewrite",
    "disable_host_rewrite"),
  validators.ConflictingObjectAttributes("auto_host_rewrite",
    "host_rewrite"),
  validators.ConflictingObjectAttributes("disable_host_rewrite",
    "host_rewrite")}
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
  "x-ves-oneof-field-host_rewrite_params": "[\"auto_host_rewrite\",\"disable_host_rewrite\",\"host_rewrite\"]"
}
```

Terraform syntax:

```terraform
default_route {
  # Configure direct properties listed below.
}
```

<a id="canonical-a6fc7d7e567ac5d221970fc94cf0f415c62b1b9484eb409d3c71233220afff0a"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.defa / 1e9f134c49b6 / 3

- [auto_host_rewrite](resources--workload--reference--group-017.md#canonical-d24f73a412c5216e73506466ced95459a6e0abef94f9acdc3a2bd4c6820ba903): complete subsection reference.

- [disable_host_rewrite](resources--workload--reference--group-017.md#canonical-af2a5fed7b9d985376f3a31b038f14d13e14daeebabc1df236c6b33c439e53a7): complete subsection reference.

<a id="canonical-db1673bc3449fbb3a706b3a99e975b4585b70d6e58d37de3b08a73afc5650539"></a>

<a id="canonical-6a7104b046ad9d4e6480361d31880632408a7ad64a08277eb901c46cea8d615f"></a>

## host_rewrite property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.defa / 1e9f134c49b6 / 4

Type: `"string"`. Optional.

Exclusive with \[auto\_host\_rewrite disable\_host\_rewrite\] Host header will be swapped with this
value.

Upstream description:

Exclusive with \[auto\_host\_rewrite disable\_host\_rewrite\] Host header will be swapped with this
value.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

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

<a id="canonical-f67075e2eba49cf38297943ca15dc81e50b6d8f6fe2a640b84a20ed3700edfe3"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.defa / 1e9f134c49b6 / 5

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route.auto_host_rewrite](resources--workload--reference--group-017.md#canonical-d24f73a412c5216e73506466ced95459a6e0abef94f9acdc3a2bd4c6820ba903)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route.disable_host_rewrite](resources--workload--reference--group-017.md#canonical-af2a5fed7b9d985376f3a31b038f14d13e14daeebabc1df236c6b33c439e53a7)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-d24f73a412c5216e73506466ced95459a6e0abef94f9acdc3a2bd4c6820ba903"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0dd1fd8c983a17f3ffa8f29fb5128e5de55b587d1454e8a942e0060fbadda93a"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route.auto_host_rewrite — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.defa / 478249c56b74 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route](resources--workload--reference--group-017.md#canonical-8d583f1134632f8bd8257e1cd4034ae146003e5cb69f89850ac1930fd06d94e1)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route.auto_host_rewrite

<a id="canonical-750f209e6c67a80200cd220e0a222aa0427e3b80274fdacac8034c08e6a0bc9f"></a>

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
auto_host_rewrite = {}
```

<a id="canonical-4f8b7ef68ffad75530507dc43f5ccb27eb489bbab74cb4c7a78c0b8c32a8470c"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.defa / 478249c56b74 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7167ecb648ea16e9ba989519e5ba190b0bd3af7c6f2ead2481a7f78746377d74"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.defa / 478249c56b74 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route](resources--workload--reference--group-017.md#canonical-8d583f1134632f8bd8257e1cd4034ae146003e5cb69f89850ac1930fd06d94e1)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-af2a5fed7b9d985376f3a31b038f14d13e14daeebabc1df236c6b33c439e53a7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1a58f6c3753cd235b61f4ff9b76725d94a3d2abf53390e9938b963d2cdf1f7bc"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route.disable_host_rewrite — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.defa / 3e14ca296dfc / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route](resources--workload--reference--group-017.md#canonical-8d583f1134632f8bd8257e1cd4034ae146003e5cb69f89850ac1930fd06d94e1)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route.disable_host_rewrite

<a id="canonical-bd92b23c1569d8bc2eaa28bf8c0fe280d5e1ba961c03876ca19239d07a256b51"></a>

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
disable_host_rewrite = {}
```

<a id="canonical-80dd1d24d8ceff8cc5ca336f27518c926fe040dc85190c6f88a9604b611b0e6a"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.defa / 3e14ca296dfc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-556d3902b7c03bb8969912c50816a30e0c02fd8771ed07518113221e398045ea"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.defa / 3e14ca296dfc / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route](resources--workload--reference--group-017.md#canonical-8d583f1134632f8bd8257e1cd4034ae146003e5cb69f89850ac1930fd06d94e1)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-0106c8e4449e963d5280c12451e8503308a4aba76db34acb949ad17779b80dc9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1a0efdcd905815c37fc40210dcc60258db78941f9c0172044fe948a5bb8229e1"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 627d3d58f84c / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http

<a id="canonical-54119862606c88bf27227dedfdfe599af55ce749a8a26462e3e304de4f446764"></a>

Type: `"object"`. single nested block, Optional.

HTTP Choice. Choice for selecting HTTP proxy.

Upstream description:

Choice for selecting HTTP proxy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("port",
    "port_ranges")}
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
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]"
}
```

Terraform syntax:

```terraform
http {
  # Configure direct properties listed below.
}
```

<a id="canonical-6c8e53cc05ebe1870eb1c180d02b01f53d1e51db55f9716db614e4ad306ea1ad"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 627d3d58f84c / 3

<a id="canonical-2a8f238fe59da06cea591bd0fd7a9da711e0af77cfa64bc52e307ba22b429a96"></a>

<a id="canonical-d63aa288ac135fd79cde5b922a4fab17aafb7d421f06a3d71de5ae12d3aeea7d"></a>

## dns_volterra_managed property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 627d3d58f84c / 4

Type: `"bool"`. Optional.

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

<a id="canonical-ef7d9e77ddd4d2262f2ceaa8f874adea1ee9ffa4a218942ed587996d7dda5ebd"></a>

<a id="canonical-669ce00ec360e752b96a39327d80eb790b2d8288c06d990bc0f88e9b4aa91198"></a>

## port property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 627d3d58f84c / 5

Type: `"number"`. Optional.

Exclusive with \[port\_ranges\] HTTP port to Listen.

Upstream description:

Exclusive with \[port\_ranges\] HTTP port to Listen.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
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

<a id="canonical-5e6c36d054251278bb6ca5fbb21dff13a9cceb578a2de14bdf76d1c8fd0ebca6"></a>

<a id="canonical-c34520b482aa3732422476454a46f34b5200169ada90b74d3dff5d3a66b26bee"></a>

## port_ranges property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 627d3d58f84c / 6

Type: `"string"`. Optional.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by "-".

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 512),
}
```

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

<a id="canonical-a396bcf77300506cae938e7226c48c066562a99af3dd5babdd0b5632b0ab0577"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 627d3d58f84c / 7

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-1622054a074c74fd54c41bd170891da7f136459269bb1c252d02e71dde5f37ff"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-41c54567054ad7ffed38483eb534712721f88a5dcd0558851b6850c85a232318"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / a2af254c289a / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https

<a id="canonical-18d342e5dde3ebd410d1b305bcc47ec996e32f6d016cbde793ae04eaa21ab282"></a>

Type: `"object"`. single nested block, Optional.

Choice for selecting HTTP proxy with bring your own certificates.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("append_server_name",
    "default_header"),
  validators.ConflictingObjectAttributes("append_server_name",
    "pass_through"),
  validators.ConflictingObjectAttributes("append_server_name",
    "server_name"),
  validators.ConflictingObjectAttributes("default_header",
    "pass_through"),
  validators.ConflictingObjectAttributes("default_header",
    "server_name"),
  validators.ConflictingObjectAttributes("default_loadbalancer",
    "non_default_loadbalancer"),
  validators.ConflictingObjectAttributes("disable_path_normalize",
    "enable_path_normalize"),
  validators.ConflictingObjectAttributes("pass_through",
    "server_name"),
  validators.ConflictingObjectAttributes("port",
    "port_ranges"),
  validators.ConflictingObjectAttributes("tls_cert_params",
    "tls_parameters")}
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
  "x-ves-oneof-field-default_lb_choice": "[\"default_loadbalancer\",\"non_default_loadbalancer\"]",
  "x-ves-oneof-field-path_normalize_choice": "[\"disable_path_normalize\",\"enable_path_normalize\"]",
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]",
  "x-ves-oneof-field-server_header_choice": "[\"append_server_name\",\"default_header\",\"pass_through\",\"server_name\"]",
  "x-ves-oneof-field-tls_certificates_choice": "[\"tls_cert_params\",\"tls_parameters\"]"
}
```

Terraform syntax:

```terraform
https {
  # Configure direct properties listed below.
}
```

<a id="canonical-a0f132f0be6704725e5249aee16bb7ee7d344069a1d5041c4a56eed407934046"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / a2af254c289a / 3

<a id="canonical-4fff0f0cba896a6b78adc3103af5d4e4cc13a5dbe872b0973987c7e06d9ceb80"></a>

<a id="canonical-946fcb9ee8f1841f993f4f51a877f40ee4a087fe72812c7d3f246a8a9cc5e916"></a>

## add_hsts property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / a2af254c289a / 4

Type: `"bool"`. Optional.

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

<a id="canonical-c67b2c9112a01de2797c5377533a2b4b789cc67e4559a4868d0140444550dbd7"></a>

<a id="canonical-06b2213d2a2702badd4af2f19a62dc973a888a18accc1e42e91b7a19d4110f01"></a>

## append_server_name property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / a2af254c289a / 5

Type: `"string"`. Optional.

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

Upstream description:

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

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

- [coalescing_options](resources--workload--reference--group-017.md#canonical-8eb8b4d6de0bb150bffff9cc1df41f9fedb39c01d92178b00d36a90ad0cf68f6): complete subsection reference.

<a id="canonical-71b4e96242adc1b81a620f83f17c6b78626a1bf7becf0b172a6f44cfdb964483"></a>

<a id="canonical-6986423f03b3dd7b3a42dcd106cf5e0e84bc4af1b08f2d974418dc36b8c176ee"></a>

## connection_idle_timeout property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / a2af254c289a / 6

Type: `"number"`. Optional.

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed.

Upstream description:

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed. Note
that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive. This is
specified in milliseconds. The default value is 2 minutes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(600000),
}
```

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

- [default_header](resources--workload--reference--group-017.md#canonical-03afbccdf5faf7a7c0999d543afc01008648f433bf814a50cb46c9a639510b6f): complete subsection reference.

- [default_loadbalancer](resources--workload--reference--group-017.md#canonical-b1324d87fc99c373c1fea17b1f3e0b838f345b4a7dda4c8fe28d62755b96d688): complete subsection reference.

- [disable_path_normalize](resources--workload--reference--group-017.md#canonical-1b542a69c0cbdd8ac8c712829cfc38de3ad362354fad44926a8f085e242aa852): complete subsection reference.

- [enable_path_normalize](resources--workload--reference--group-017.md#canonical-9c7fb4289d4697b150d5635e88c665671c06aaae59d914487a73f31fbacf911e): complete subsection reference.

- [http_protocol_options](resources--workload--reference--group-017.md#canonical-e29813cfdfb724ac3f8dba080a9b7e1cce755c4571663f2d848240792fd7ea85): complete subsection reference.

<a id="canonical-e335b918fdcddfaea2655a41826888742ddb2771c234afcac4d5cb585317c410"></a>

<a id="canonical-632a264927062f2be6a18c68ddecc8d98d218d0958bd9c6d8c81e6e3a331f22d"></a>

## http_redirect property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / a2af254c289a / 7

Type: `"bool"`. Optional.

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

- [non_default_loadbalancer](resources--workload--reference--group-018.md#canonical-53eed5c191be0dde178d215db47e72d8b0233cd6759ea27be1498c1e4cca818a): complete subsection reference.

- [pass_through](resources--workload--reference--group-018.md#canonical-dae166a8aa6c357a4ece08b831d4990704538dcbf9225272b2cf57c4ed128d0d): complete subsection reference.

<a id="canonical-c9ab7123c635fd449d41bfeaac8cf97d332583f39f695f2d1c73925b2e8b8b58"></a>

<a id="canonical-411771246b65b2bbba7e8ddc9939b56677ac6fe201ed83d3424836d1b5fe867c"></a>

## port property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / a2af254c289a / 8

Type: `"number"`. Optional.

Exclusive with \[port\_ranges\] HTTPS port to Listen.

Upstream description:

Exclusive with \[port\_ranges\] HTTPS port to Listen.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
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

<a id="canonical-5d457c74b32174ed85215ea07c9e61d0dc42601bd8bd52cd938068798393b2d5"></a>

<a id="canonical-663c52f10d14968497d33d30ec066d118ec1a250dbdb4044273fb7ef48215d0e"></a>

## port_ranges property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / a2af254c289a / 9

Type: `"string"`. Optional.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by "-".

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 512),
}
```

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

<a id="canonical-0fcd1ee7c29bd8d2a0da9a02ccdbba0d45a948d5100aba996b588aa50cb49164"></a>

<a id="canonical-6725b903c89521c1bc73305aa3c27ede0f20ff2b396b92e43a8c2312aeaa48ca"></a>

## server_name property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / a2af254c289a / 10

Type: `"string"`. Optional.

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

Upstream description:

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

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

- [tls_cert_params](resources--workload--reference--group-018.md#canonical-7ced8ab09b844f7830bb0fba04aea9424ecb6c162fe02a3148f571c4e282b669): complete subsection reference.

- [tls_parameters](resources--workload--reference--group-018.md#canonical-88e031cfad4357b09e79f29ec26027a3ddb61f3d2d19a38b6294452c2c1da6a9): complete subsection reference.

<a id="canonical-b073853bfe13fd024bd5447e9d3b996be72c4bf9b75d5a9341e4735e0d3249c4"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / a2af254c289a / 11

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-017.md#canonical-8eb8b4d6de0bb150bffff9cc1df41f9fedb39c01d92178b00d36a90ad0cf68f6)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default_header](resources--workload--reference--group-017.md#canonical-03afbccdf5faf7a7c0999d543afc01008648f433bf814a50cb46c9a639510b6f)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default_loadbalancer](resources--workload--reference--group-017.md#canonical-b1324d87fc99c373c1fea17b1f3e0b838f345b4a7dda4c8fe28d62755b96d688)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.disable_path_normalize](resources--workload--reference--group-017.md#canonical-1b542a69c0cbdd8ac8c712829cfc38de3ad362354fad44926a8f085e242aa852)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.enable_path_normalize](resources--workload--reference--group-017.md#canonical-9c7fb4289d4697b150d5635e88c665671c06aaae59d914487a73f31fbacf911e)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-017.md#canonical-e29813cfdfb724ac3f8dba080a9b7e1cce755c4571663f2d848240792fd7ea85)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.non_default_loadbalancer](resources--workload--reference--group-018.md#canonical-53eed5c191be0dde178d215db47e72d8b0233cd6759ea27be1498c1e4cca818a)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.pass_through](resources--workload--reference--group-018.md#canonical-dae166a8aa6c357a4ece08b831d4990704538dcbf9225272b2cf57c4ed128d0d)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-018.md#canonical-7ced8ab09b844f7830bb0fba04aea9424ecb6c162fe02a3148f571c4e282b669)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-018.md#canonical-88e031cfad4357b09e79f29ec26027a3ddb61f3d2d19a38b6294452c2c1da6a9)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-8eb8b4d6de0bb150bffff9cc1df41f9fedb39c01d92178b00d36a90ad0cf68f6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ff1dda3cf225913bb7b657f7139a048dad94cd05c5caa2e3c2d1b4b7bee539f1"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 91bcbd0ed28f / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-017.md#canonical-1622054a074c74fd54c41bd170891da7f136459269bb1c252d02e71dde5f37ff)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options

<a id="canonical-71ecf2229f7f292f0bb41875df19494be2bd8821af37cb6e115cd90331961a02"></a>

Type: `"object"`. single nested block, Optional.

TLS connection coalescing configuration (not compatible with mTLS).

Upstream description:

TLS connection coalescing configuration (not compatible with mTLS)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_coalescing",
    "strict_coalescing")}
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
  "x-ves-oneof-field-coalescing_choice": "[\"default_coalescing\",\"strict_coalescing\"]"
}
```

Terraform syntax:

```terraform
coalescing_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-d8f0a01ae974670d049e7400fe53b83506939b5f85fb19ac2b63fc2537f90f6d"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 91bcbd0ed28f / 3

- [default_coalescing](resources--workload--reference--group-017.md#canonical-0709ca9a2e17bef713e9dd2b3074fb9f13aebeb28740f5c7f88ede6af4717220): complete subsection reference.

- [strict_coalescing](resources--workload--reference--group-017.md#canonical-8456fe94afa1f9ade15a933fd4107c21fd28d4cbc99a66799745f413542648c1): complete subsection reference.

<a id="canonical-e0bf57724d17b8425ab7deb0ad6a31c2aeb8465ebd2cfcdedd9eabd09d28618e"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 91bcbd0ed28f / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options.default_coalescing](resources--workload--reference--group-017.md#canonical-0709ca9a2e17bef713e9dd2b3074fb9f13aebeb28740f5c7f88ede6af4717220)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options.strict_coalescing](resources--workload--reference--group-017.md#canonical-8456fe94afa1f9ade15a933fd4107c21fd28d4cbc99a66799745f413542648c1)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-017.md#canonical-1622054a074c74fd54c41bd170891da7f136459269bb1c252d02e71dde5f37ff)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-0709ca9a2e17bef713e9dd2b3074fb9f13aebeb28740f5c7f88ede6af4717220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bb61e950659746c74d58b09f446f34fd31955a1937ffbd16529305d6b781a0a4"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options.default_coalescing — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 5cbc15fe9f7f / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-017.md#canonical-1622054a074c74fd54c41bd170891da7f136459269bb1c252d02e71dde5f37ff)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-017.md#canonical-8eb8b4d6de0bb150bffff9cc1df41f9fedb39c01d92178b00d36a90ad0cf68f6)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options.default_coalescing

<a id="canonical-44dcbac7e2f692619ffb11f85baf9f7730d0c0abe2b7a37ccf79aa2732591d0b"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default_coalescing = {}
```

<a id="canonical-96b105058a4808c89f29dd8da9b05dff9df686e8eca5d98ddaf8fc23aa804d66"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 5cbc15fe9f7f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fa0fa70580d841ce8481e7f14aac8c58960e8c89980e0151c6c0c9adb47285da"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 5cbc15fe9f7f / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-017.md#canonical-8eb8b4d6de0bb150bffff9cc1df41f9fedb39c01d92178b00d36a90ad0cf68f6)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-8456fe94afa1f9ade15a933fd4107c21fd28d4cbc99a66799745f413542648c1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d047b7a93f14040788a90999457c2b47967ce2773ad5ffb88515c30776a79448"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options.strict_coalescing — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 2c88f171db05 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-017.md#canonical-1622054a074c74fd54c41bd170891da7f136459269bb1c252d02e71dde5f37ff)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-017.md#canonical-8eb8b4d6de0bb150bffff9cc1df41f9fedb39c01d92178b00d36a90ad0cf68f6)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options.strict_coalescing

<a id="canonical-5677b87ea8c2e4bf4931c3e99b0b75bd8fc32951de00037b95efac90d3d8945e"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
strict_coalescing = {}
```

<a id="canonical-7feed3dd2d3b7d92006ca80c970c7ad270ae470419b9acb4b70f37aadb803dc1"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 2c88f171db05 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ecf9328d729482fed89cdd9f381b00aa721b4f0411ed26750338a8314474779e"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 2c88f171db05 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-017.md#canonical-8eb8b4d6de0bb150bffff9cc1df41f9fedb39c01d92178b00d36a90ad0cf68f6)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-03afbccdf5faf7a7c0999d543afc01008648f433bf814a50cb46c9a639510b6f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3b5031e88abd002b425935f8f1b537dcb8cc44542792d4b0fbfab9f166becbf9"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default_header — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 8b6dedba3d6e / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-017.md#canonical-1622054a074c74fd54c41bd170891da7f136459269bb1c252d02e71dde5f37ff)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default_header

<a id="canonical-75e4f5257d8caaed5387dd9669c45fbdeb9e1c87be4555ffac6a16ad4994c1e2"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default_header = {}
```

<a id="canonical-dc6f2396c786f3a7bd2ee6068c4ebc460a2196450e159db8de96d1d074c66a4f"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 8b6dedba3d6e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e023c5b29ddf9ebb428592dc07f47ea5ab189c30977c336a125aa96d0a823bab"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 8b6dedba3d6e / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-017.md#canonical-1622054a074c74fd54c41bd170891da7f136459269bb1c252d02e71dde5f37ff)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-b1324d87fc99c373c1fea17b1f3e0b838f345b4a7dda4c8fe28d62755b96d688"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1c115e28118771240895eb68e2e4ccc75826b6bcff1ad89f0739d6f025beb423"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default_loadbalancer — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 952d6af1a2b6 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-017.md#canonical-1622054a074c74fd54c41bd170891da7f136459269bb1c252d02e71dde5f37ff)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default_loadbalancer

<a id="canonical-5087344d3cf83d2571f2d4b803226a1ccaf4f0145299acbfdb8bc6e404a1bb97"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default_loadbalancer = {}
```

<a id="canonical-e5936eb4c5923fe8f97c9d5cb1fe622213a12af9e6f3ab0037dc713119cf6fae"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 952d6af1a2b6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-57131bf94811d6e575b00f99d53f6f800aa0e2fa63666c3f7b6307cd8e2de3b2"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 952d6af1a2b6 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-017.md#canonical-1622054a074c74fd54c41bd170891da7f136459269bb1c252d02e71dde5f37ff)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-1b542a69c0cbdd8ac8c712829cfc38de3ad362354fad44926a8f085e242aa852"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-db247c12db5f6e3b314fca343d4d37ab4f1ae0b05cce2133fd4b21dfe99fbe17"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.disable_path_normalize — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 896230e657d4 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-017.md#canonical-1622054a074c74fd54c41bd170891da7f136459269bb1c252d02e71dde5f37ff)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.disable_path_normalize

<a id="canonical-3ee28cbb7f7a32052da34a4c1f60dfb33d1a5ef0d6e808a11227f8fc2f50786a"></a>

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
disable_path_normalize = {}
```

<a id="canonical-1a691bebd8d6acc900d0536496c4b584a5e9390176989e12ecf5040598ec2547"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 896230e657d4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8c4f7cf7facbd2c5d81a0a4c6d2644ffee4748fd37f8f63b95ed0d07f6a4f6b6"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 896230e657d4 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-017.md#canonical-1622054a074c74fd54c41bd170891da7f136459269bb1c252d02e71dde5f37ff)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-9c7fb4289d4697b150d5635e88c665671c06aaae59d914487a73f31fbacf911e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-621825c76a10d7fe9f31661e4f2d0c140b4a2456a012c351d476730dc561e077"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.enable_path_normalize — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / dd7b0372cb45 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-017.md#canonical-1622054a074c74fd54c41bd170891da7f136459269bb1c252d02e71dde5f37ff)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.enable_path_normalize

<a id="canonical-24c7ee67ac3e7e2e68792ce0b1f007f7be2708e6d85d297988bac3aaf7d3d41b"></a>

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
enable_path_normalize = {}
```

<a id="canonical-86d16742c5de960bb174f996f6ecf2773a75eaa92fb11fd9a0d29afc5380e81d"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / dd7b0372cb45 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bccf8983fe069b75a1fed0adc4fe00e1ad4dc7d2062b147d5cff4d9e9f6358a2"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / dd7b0372cb45 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-017.md#canonical-1622054a074c74fd54c41bd170891da7f136459269bb1c252d02e71dde5f37ff)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-e29813cfdfb724ac3f8dba080a9b7e1cce755c4571663f2d848240792fd7ea85"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-91ed868c2ca0e68f7c7817aadd67cad1166b3697164c0fd3f7ddb267b4074cd7"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 9af27be5c6e0 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-017.md#canonical-1622054a074c74fd54c41bd170891da7f136459269bb1c252d02e71dde5f37ff)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options

<a id="canonical-a5135bda50c9044e8df9d841ad2260368591f0dda4fcfcd1ebd5eca9054911b9"></a>

Type: `"object"`. single nested block, Optional.

HTTP protocol configuration OPTIONS for downstream connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("http_protocol_enable_v1_only",
    "http_protocol_enable_v1_v2"),
  validators.ConflictingObjectAttributes("http_protocol_enable_v1_only",
    "http_protocol_enable_v2_only"),
  validators.ConflictingObjectAttributes("http_protocol_enable_v1_v2",
    "http_protocol_enable_v2_only")}
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
  "x-ves-oneof-field-http_protocol_choice": "[\"http_protocol_enable_v1_only\",\"http_protocol_enable_v1_v2\",\"http_protocol_enable_v2_only\"]"
}
```

Terraform syntax:

```terraform
http_protocol_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-e22c362a4858bb0ddaf605439ac164f0b3ce6b37245e829f98362092034741cc"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 9af27be5c6e0 / 3

- [http_protocol_enable_v1_only](resources--workload--reference--group-017.md#canonical-11663ad1b3589ca7a4bdd67a44644931bcd750d4a2b1c8e7bdd32b80671601fe): complete subsection reference.

- [http_protocol_enable_v1_v2](resources--workload--reference--group-018.md#canonical-6f1b60ff91446b9810d04bb5059916a5813439db4d46bc0eeb351073e1ccde45): complete subsection reference.

- [http_protocol_enable_v2_only](resources--workload--reference--group-018.md#canonical-9898c85668d448c21997aabb8ed23df19552a9932ac87f747d2e08826801f34b): complete subsection reference.

<a id="canonical-ce186876feaa79418f306b01841c82523eb051f5f43991fa255063fbe165b6ae"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 9af27be5c6e0 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-017.md#canonical-11663ad1b3589ca7a4bdd67a44644931bcd750d4a2b1c8e7bdd32b80671601fe)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2](resources--workload--reference--group-018.md#canonical-6f1b60ff91446b9810d04bb5059916a5813439db4d46bc0eeb351073e1ccde45)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only](resources--workload--reference--group-018.md#canonical-9898c85668d448c21997aabb8ed23df19552a9932ac87f747d2e08826801f34b)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-017.md#canonical-1622054a074c74fd54c41bd170891da7f136459269bb1c252d02e71dde5f37ff)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-11663ad1b3589ca7a4bdd67a44644931bcd750d4a2b1c8e7bdd32b80671601fe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d28f46621d0b6f8ccd629748fe369a5ee73ff234cd6de976b8552492fa6faf74"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / d71777973ceb / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-017.md#canonical-1622054a074c74fd54c41bd170891da7f136459269bb1c252d02e71dde5f37ff)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-017.md#canonical-e29813cfdfb724ac3f8dba080a9b7e1cce755c4571663f2d848240792fd7ea85)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-3423eb54f2938253a7e1b11ebd6f1f00b8ea2f79d3ad702a5293c2f5f5264542"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
http_protocol_enable_v1_only {
  # Configure direct properties listed below.
}
```

<a id="canonical-76b84d05c81b412b6e842365256576718181a5e35ab827a3afd756a8a813b33e"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / d71777973ceb / 3

- [header_transformation](resources--workload--reference--group-017.md#canonical-c86711867b914cd37efd496ed15aed5ef34f36ac2fc9d287e35c91e66baae8b3): complete subsection reference.

<a id="canonical-812b966ec2d70bf879bcfc81e36f3eca2693c52436923242bf0eea685bc2eaa2"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / d71777973ceb / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-017.md#canonical-c86711867b914cd37efd496ed15aed5ef34f36ac2fc9d287e35c91e66baae8b3)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-017.md#canonical-e29813cfdfb724ac3f8dba080a9b7e1cce755c4571663f2d848240792fd7ea85)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-c86711867b914cd37efd496ed15aed5ef34f36ac2fc9d287e35c91e66baae8b3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f30ff287b5447129b4185c49275572362da454bc71d60ea9ddba143ce6320c37"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / a4ac81a064bd / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-017.md#canonical-1622054a074c74fd54c41bd170891da7f136459269bb1c252d02e71dde5f37ff)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-017.md#canonical-e29813cfdfb724ac3f8dba080a9b7e1cce755c4571663f2d848240792fd7ea85)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-017.md#canonical-11663ad1b3589ca7a4bdd67a44644931bcd750d4a2b1c8e7bdd32b80671601fe)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-b4d86f01d88e3f2bc48c82206b55400614e48a0a1b99fc41f1496fc77bd0db3b"></a>

Type: `"object"`. single nested block, Optional.

Header Transformation OPTIONS for HTTP/1.1 request/response headers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_header_transformation",
    "preserve_case_header_transformation"),
  validators.ConflictingObjectAttributes("default_header_transformation",
    "proper_case_header_transformation"),
  validators.ConflictingObjectAttributes("preserve_case_header_transformation",
    "proper_case_header_transformation")}
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
  "x-ves-oneof-field-header_transformation_choice": "[\"default_header_transformation\",\"preserve_case_header_transformation\",\"proper_case_header_transformation\"]"
}
```

Terraform syntax:

```terraform
header_transformation {
  # Configure direct properties listed below.
}
```

<a id="canonical-fc8f27d7344a920962dda30f3839dd3e1d883414e86d4b168af07c7a8d3c6690"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / a4ac81a064bd / 3

- [default_header_transformation](resources--workload--reference--group-017.md#canonical-04b573364fbd174ed685f1e7f35acaf1bdeadebd31e2bef4e60e313850789884): complete subsection reference.

- [preserve_case_header_transformation](resources--workload--reference--group-018.md#canonical-1adae2cc4658aef4a0053d49c7e7fb8c1fb39b5b7a9e39c0cdd34a690b7e92ad): complete subsection reference.

- [proper_case_header_transformation](resources--workload--reference--group-018.md#canonical-ed0b98048cbaa4bc9f97119ca566f4bedc756db301257db468192cad97b1d3cc): complete subsection reference.

<a id="canonical-25d0b822034887f61e5e540ba327316aacfaf27d61be49d208c04cd5944d65ae"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / a4ac81a064bd / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](resources--workload--reference--group-017.md#canonical-04b573364fbd174ed685f1e7f35acaf1bdeadebd31e2bef4e60e313850789884)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](resources--workload--reference--group-018.md#canonical-1adae2cc4658aef4a0053d49c7e7fb8c1fb39b5b7a9e39c0cdd34a690b7e92ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](resources--workload--reference--group-018.md#canonical-ed0b98048cbaa4bc9f97119ca566f4bedc756db301257db468192cad97b1d3cc)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-017.md#canonical-11663ad1b3589ca7a4bdd67a44644931bcd750d4a2b1c8e7bdd32b80671601fe)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-04b573364fbd174ed685f1e7f35acaf1bdeadebd31e2bef4e60e313850789884"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
