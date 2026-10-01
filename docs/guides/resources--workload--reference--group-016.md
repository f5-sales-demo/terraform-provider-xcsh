---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-9c1e355ddba624a77329b567ea0bbf8d879d4656fab84c36dbb4bb1b668fee55"></a>

## service.containers.readiness_check.http_health_check — service.containers.readiness_check.http_health_check / fa25025ac315 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.containers](resources--workload--reference--group-015.md#canonical-3f88422b353a65363a1ce1b0d566b8c6aa769a87c87c23cb23b8978323a9c01a)
- [service.containers.readiness_check](resources--workload--reference--group-015.md#canonical-8c84db60fef3270c2e5c9acf436cd213cc7de6870d6207184ca297b292e72b6f)
- service.containers.readiness_check.http_health_check

<a id="canonical-ed00671b0fe11c37aa4f4cd846c0dfd4f0738adeef6ae46fb71fc3f940e45cb5"></a>

Type: `"object"`. single nested block, Optional.

HTTPHealthCheckType describes a health check based on HTTP GET requests.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("path")}
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
http_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-af02247edda3e1aaad0d189fc6fd7d5db69b49cc19083f6e7cdb98d2d2f94393"></a>

## Direct properties — service.containers.readiness_check.http_health_check / fa25025ac315 / 3

<a id="canonical-da274dd6a77b281f659e8dd0259f08f0b3d59f1ad4aaa8a3b3cd9af8c2243345"></a>

<a id="canonical-44223aaa90a35e5e237e2e60700f133bd3dfb7a84284ba8b4fa31dc2983f60ab"></a>

## headers property — service.containers.readiness_check.http_health_check / fa25025ac315 / 4

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

<a id="canonical-2618f270bedce12b122ac4a335b7a3f22a8d9e11da2225d8cd83c21f3b31d0fd"></a>

<a id="canonical-1266c9c8b7564c5a5f02d540aea74d14e02d851a10c1825716fbc16444e1e20b"></a>

## host_header property — service.containers.readiness_check.http_health_check / fa25025ac315 / 5

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

<a id="canonical-5ae8afd5ce33ecb2fd98de35cb3752f722e3095a4f21eaebbcc0d43826b21714"></a>

<a id="canonical-63c95e9e91be523e4ae4e58284af1fde80a1d07c35f3fa31c99ac0288469a544"></a>

## path property — service.containers.readiness_check.http_health_check / fa25025ac315 / 6

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

- [port](resources--workload--reference--group-016.md#canonical-24e4afdbb96808449a7331026613431afeab85861000e3ba6d544937b57a3378): complete subsection reference.

<a id="canonical-10c50052398278630f6c2b25969c4f2203e668d69f2f3455ce512464937a0f2a"></a>

## Next pages — service.containers.readiness_check.http_health_check / fa25025ac315 / 7

- [service.containers.readiness_check.http_health_check.port](resources--workload--reference--group-016.md#canonical-24e4afdbb96808449a7331026613431afeab85861000e3ba6d544937b57a3378)
- [service.containers.readiness_check](resources--workload--reference--group-015.md#canonical-8c84db60fef3270c2e5c9acf436cd213cc7de6870d6207184ca297b292e72b6f)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-24e4afdbb96808449a7331026613431afeab85861000e3ba6d544937b57a3378"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a7bf905eb721c4b00d293c3a208b4229b3993709e52f4738a0890412923349bf"></a>

## service.containers.readiness_check.http_health_check.port — service.containers.readiness_check.http_health_check.port / b22543ffae92 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.containers](resources--workload--reference--group-015.md#canonical-3f88422b353a65363a1ce1b0d566b8c6aa769a87c87c23cb23b8978323a9c01a)
- [service.containers.readiness_check](resources--workload--reference--group-015.md#canonical-8c84db60fef3270c2e5c9acf436cd213cc7de6870d6207184ca297b292e72b6f)
- [service.containers.readiness_check.http_health_check](resources--workload--reference--group-015.md#canonical-82b719c7307729eec6e569e822b334ccfeabf207f05b35e5f10acdfb5d4d5896)
- service.containers.readiness_check.http_health_check.port

<a id="canonical-cb3637fb7ba0e4eaf09d2ec87992f4f275fe1db4af54aabac4e14623b32cb87a"></a>

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

<a id="canonical-599ded0d2550611e3e3c8bd93d770816e184f442c27593d34a90019c8599c341"></a>

## Direct properties — service.containers.readiness_check.http_health_check.port / b22543ffae92 / 3

<a id="canonical-f55f401c8788b684bc452cafc23782aff3e8f23c66ce6555f497f76b3db4c9c5"></a>

<a id="canonical-7600ce29e681a545da3beb5663631d393c2135c1ad974f4917d4032263a8e4f8"></a>

## name property — service.containers.readiness_check.http_health_check.port / b22543ffae92 / 4

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

<a id="canonical-4a4cb399071d4d2fd3a7c507b36c7d1f4a934c4f44b0cd586615ba6587fe62d8"></a>

<a id="canonical-53e72638f0c0d8bd175f44bba7c266d02c00e9243fbaea231e54f68c4473c910"></a>

## num property — service.containers.readiness_check.http_health_check.port / b22543ffae92 / 5

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

<a id="canonical-218594ea89946e280086be90bd100b7fa8bec38c18a1728a4b0632d65a72973a"></a>

## Next pages — service.containers.readiness_check.http_health_check.port / b22543ffae92 / 6

- [service.containers.readiness_check.http_health_check](resources--workload--reference--group-015.md#canonical-82b719c7307729eec6e569e822b334ccfeabf207f05b35e5f10acdfb5d4d5896)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-bae7e7156dfc3765696e12988318f708df2763faadc27866633aa63112a9ba7b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ac2dfba803cd8533337b42237b5bfc8ae107476229a29152bba600053541eeb6"></a>

## service.containers.readiness_check.tcp_health_check — service.containers.readiness_check.tcp_health_check / 0b619d4c90b9 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.containers](resources--workload--reference--group-015.md#canonical-3f88422b353a65363a1ce1b0d566b8c6aa769a87c87c23cb23b8978323a9c01a)
- [service.containers.readiness_check](resources--workload--reference--group-015.md#canonical-8c84db60fef3270c2e5c9acf436cd213cc7de6870d6207184ca297b292e72b6f)
- service.containers.readiness_check.tcp_health_check

<a id="canonical-e6dfc272389a6e40c8e8fca33c292480f5918c80e5e2184339056697aa6a1f43"></a>

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

<a id="canonical-dae6f4da654435384090b3d3e38c017b364f1c2cca0fa3a330075fb2bb7c9808"></a>

## Direct properties — service.containers.readiness_check.tcp_health_check / 0b619d4c90b9 / 3

- [port](resources--workload--reference--group-016.md#canonical-a6d52870797cb8bf62f72a21f05f718de62b25fa3b861b457d282b285db73fa7): complete subsection reference.

<a id="canonical-b42000efbaa1a4918f8d2f2978430fb635cac2861bc1256721d4cca609113d67"></a>

## Next pages — service.containers.readiness_check.tcp_health_check / 0b619d4c90b9 / 4

- [service.containers.readiness_check.tcp_health_check.port](resources--workload--reference--group-016.md#canonical-a6d52870797cb8bf62f72a21f05f718de62b25fa3b861b457d282b285db73fa7)
- [service.containers.readiness_check](resources--workload--reference--group-015.md#canonical-8c84db60fef3270c2e5c9acf436cd213cc7de6870d6207184ca297b292e72b6f)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-a6d52870797cb8bf62f72a21f05f718de62b25fa3b861b457d282b285db73fa7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7fbe4e03b1410bc32991539023b9ebca4ab9377d821368eb2e5e643ac522a4e3"></a>

## service.containers.readiness_check.tcp_health_check.port — service.containers.readiness_check.tcp_health_check.port / 8d4fc5e1cb3d / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.containers](resources--workload--reference--group-015.md#canonical-3f88422b353a65363a1ce1b0d566b8c6aa769a87c87c23cb23b8978323a9c01a)
- [service.containers.readiness_check](resources--workload--reference--group-015.md#canonical-8c84db60fef3270c2e5c9acf436cd213cc7de6870d6207184ca297b292e72b6f)
- [service.containers.readiness_check.tcp_health_check](resources--workload--reference--group-016.md#canonical-bae7e7156dfc3765696e12988318f708df2763faadc27866633aa63112a9ba7b)
- service.containers.readiness_check.tcp_health_check.port

<a id="canonical-1b69e164372f84f796c97173ec2e662b8b0f6816517cbda7feece42eaed48078"></a>

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

<a id="canonical-a4b7439605f2ebe8c0807b9fadc6e014c2d4296fb3b057696d5fc26a7577f3ad"></a>

## Direct properties — service.containers.readiness_check.tcp_health_check.port / 8d4fc5e1cb3d / 3

<a id="canonical-67fe4e3c199696d431db54467ae2eebc8642ebf6c7b5bc822f081d69db72b8dd"></a>

<a id="canonical-4ff87751db0e9b1bf4514aa1aa72beb73d9aab76480dbaae5cf9582a0fd57f6c"></a>

## name property — service.containers.readiness_check.tcp_health_check.port / 8d4fc5e1cb3d / 4

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

<a id="canonical-9d61834c996ac18e041eec735ef867104dcc6043247fe04310e39bd0c0aa7351"></a>

<a id="canonical-b9f047b44efd89eec114a5249b24760d2dec1fd528a92bc944e259ef5300c6b6"></a>

## num property — service.containers.readiness_check.tcp_health_check.port / 8d4fc5e1cb3d / 5

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

<a id="canonical-4e423cdcb241030a89f184d16feca2bf7708f25884a069c5212e6d6ba5c5b331"></a>

## Next pages — service.containers.readiness_check.tcp_health_check.port / 8d4fc5e1cb3d / 6

- [service.containers.readiness_check.tcp_health_check](resources--workload--reference--group-016.md#canonical-bae7e7156dfc3765696e12988318f708df2763faadc27866633aa63112a9ba7b)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-05165ff47537505ed9ace4fa2b50ac0d2124ed1fc4414c6d288c2e658d5e2ec3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8e108da09977993cbb5038c18bccfaee183c878abd2c239e4197552671d905ef"></a>

## service.deploy_options — service.deploy_options / aa208049b40d / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- service.deploy_options

<a id="canonical-832846b5f924a21d9300784b6ac891f7c1cd33357e797e6eb056773af9baf47d"></a>

Type: `"object"`. single nested block, Optional.

Deploy OPTIONS are used to configure the workload deployment OPTIONS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("all_res",
    "default_virtual_sites"),
  validators.ConflictingObjectAttributes("all_res",
    "deploy_ce_sites"),
  validators.ConflictingObjectAttributes("all_res",
    "deploy_ce_virtual_sites"),
  validators.ConflictingObjectAttributes("all_res",
    "deploy_re_sites"),
  validators.ConflictingObjectAttributes("all_res",
    "deploy_re_virtual_sites"),
  validators.ConflictingObjectAttributes("default_virtual_sites",
    "deploy_ce_sites"),
  validators.ConflictingObjectAttributes("default_virtual_sites",
    "deploy_ce_virtual_sites"),
  validators.ConflictingObjectAttributes("default_virtual_sites",
    "deploy_re_sites"),
  validators.ConflictingObjectAttributes("default_virtual_sites",
    "deploy_re_virtual_sites"),
  validators.ConflictingObjectAttributes("deploy_ce_sites",
    "deploy_ce_virtual_sites"),
  validators.ConflictingObjectAttributes("deploy_ce_sites",
    "deploy_re_sites"),
  validators.ConflictingObjectAttributes("deploy_ce_sites",
    "deploy_re_virtual_sites"),
  validators.ConflictingObjectAttributes("deploy_ce_virtual_sites",
    "deploy_re_sites"),
  validators.ConflictingObjectAttributes("deploy_ce_virtual_sites",
    "deploy_re_virtual_sites"),
  validators.ConflictingObjectAttributes("deploy_re_sites",
    "deploy_re_virtual_sites")}
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
  "x-ves-oneof-field-deploy_choice": "[\"all_res\",\"default_virtual_sites\",\"deploy_ce_sites\",\"deploy_ce_virtual_sites\",\"deploy_re_sites\",\"deploy_re_virtual_sites\"]"
}
```

Terraform syntax:

```terraform
deploy_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-c1e528d3d6da40c3fe0a80475edbcff47ca15a37d34f2f645df4f413f77823aa"></a>

## Direct properties — service.deploy_options / aa208049b40d / 3

- [all_res](resources--workload--reference--group-016.md#canonical-f86bd9899375a8d64931bcdf5771cc187e045e55498cd4e85fe22d6eac40bc37): complete subsection reference.

- [default_virtual_sites](resources--workload--reference--group-016.md#canonical-1b268c6c1a5b5681e74f0b62d4b94dd78ce2d412553d044de6c53661417a44fe): complete subsection reference.

- [deploy_ce_sites](resources--workload--reference--group-016.md#canonical-8cdda23ecbd8b0a4afe693a8b8da2a0cdbf039b8155010802897620dd832a13b): complete subsection reference.

- [deploy_ce_virtual_sites](resources--workload--reference--group-016.md#canonical-4a3ccb108bbf4b6f11ead9d960371138448f69f60e24fc9e98af8bca71f40267): complete subsection reference.

- [deploy_re_sites](resources--workload--reference--group-016.md#canonical-3f5bcf3835ef9e50f6b01e38fb785dd9d787df2434acc63af699d0898e5e860c): complete subsection reference.

- [deploy_re_virtual_sites](resources--workload--reference--group-016.md#canonical-8a595d976bf182b2d6306ad10d47bd00245778b64d1589966eda5bfe8517a762): complete subsection reference.

<a id="canonical-6a7f5b8aba7e2997734f3a6d55bdee390fcba5275256099e40d8e36e76b17855"></a>

## Next pages — service.deploy_options / aa208049b40d / 4

- [service.deploy_options.all_res](resources--workload--reference--group-016.md#canonical-f86bd9899375a8d64931bcdf5771cc187e045e55498cd4e85fe22d6eac40bc37)
- [service.deploy_options.default_virtual_sites](resources--workload--reference--group-016.md#canonical-1b268c6c1a5b5681e74f0b62d4b94dd78ce2d412553d044de6c53661417a44fe)
- [service.deploy_options.deploy_ce_sites](resources--workload--reference--group-016.md#canonical-8cdda23ecbd8b0a4afe693a8b8da2a0cdbf039b8155010802897620dd832a13b)
- [service.deploy_options.deploy_ce_virtual_sites](resources--workload--reference--group-016.md#canonical-4a3ccb108bbf4b6f11ead9d960371138448f69f60e24fc9e98af8bca71f40267)
- [service.deploy_options.deploy_re_sites](resources--workload--reference--group-016.md#canonical-3f5bcf3835ef9e50f6b01e38fb785dd9d787df2434acc63af699d0898e5e860c)
- [service.deploy_options.deploy_re_virtual_sites](resources--workload--reference--group-016.md#canonical-8a595d976bf182b2d6306ad10d47bd00245778b64d1589966eda5bfe8517a762)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-f86bd9899375a8d64931bcdf5771cc187e045e55498cd4e85fe22d6eac40bc37"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8c9d84125bd22e909460136089c598c92c141b281e5fa029da498d0dbe4e66ae"></a>

## service.deploy_options.all_res — service.deploy_options.all_res / cf4bd532bc4b / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.deploy_options](resources--workload--reference--group-016.md#canonical-05165ff47537505ed9ace4fa2b50ac0d2124ed1fc4414c6d288c2e658d5e2ec3)
- service.deploy_options.all_res

<a id="canonical-79de8f338c808e68a1ed8b99c1801c92a693c727e413eed50ccaa8b6bb6c8a30"></a>

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
all_res = {}
```

<a id="canonical-21afa42c7e8d196c0cb10b98c43ad58a72c48e831fcbfae735eb9fcccf3c5f8d"></a>

## Direct properties — service.deploy_options.all_res / cf4bd532bc4b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-dcbba43b6a2903d954e91a7d3793a215db55f0057422b7d8e8455c0990a60981"></a>

## Next pages — service.deploy_options.all_res / cf4bd532bc4b / 4

- [service.deploy_options](resources--workload--reference--group-016.md#canonical-05165ff47537505ed9ace4fa2b50ac0d2124ed1fc4414c6d288c2e658d5e2ec3)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-1b268c6c1a5b5681e74f0b62d4b94dd78ce2d412553d044de6c53661417a44fe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4a993240781563b94fbe808560cce0d221d7f8dfbfd7170be83be8691793f660"></a>

## service.deploy_options.default_virtual_sites — service.deploy_options.default_virtual_sites / 0d8f8d9b10e1 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.deploy_options](resources--workload--reference--group-016.md#canonical-05165ff47537505ed9ace4fa2b50ac0d2124ed1fc4414c6d288c2e658d5e2ec3)
- service.deploy_options.default_virtual_sites

<a id="canonical-7998467bf07b0e766bd750aa46d81c000ef0f60f033ba0e5f17e0193c358ab21"></a>

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
default_virtual_sites = {}
```

<a id="canonical-b889d0f02d4f688d78e7addc47dede5680dea82af8845dbbc2d1cded1fc2b807"></a>

## Direct properties — service.deploy_options.default_virtual_sites / 0d8f8d9b10e1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-864b29a17cb693ef60f37dcddde7153060eef95a82be995f8375f217b4076823"></a>

## Next pages — service.deploy_options.default_virtual_sites / 0d8f8d9b10e1 / 4

- [service.deploy_options](resources--workload--reference--group-016.md#canonical-05165ff47537505ed9ace4fa2b50ac0d2124ed1fc4414c6d288c2e658d5e2ec3)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-8cdda23ecbd8b0a4afe693a8b8da2a0cdbf039b8155010802897620dd832a13b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-107a7bfb1ce59e575df8ac80c879ca833ac5c273fd6781be12120e1bfe39e645"></a>

## service.deploy_options.deploy_ce_sites — service.deploy_options.deploy_ce_sites / bb5dc8428de1 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.deploy_options](resources--workload--reference--group-016.md#canonical-05165ff47537505ed9ace4fa2b50ac0d2124ed1fc4414c6d288c2e658d5e2ec3)
- service.deploy_options.deploy_ce_sites

<a id="canonical-9953625bf33a278811793ecea774b4a2bea92ed2c2767eaf2b458f5fb5d4b456"></a>

Type: `"object"`. single nested block, Optional.

Defines a way to deploy a workload on specific Customer sites.

Upstream description:

This defines a way to deploy a workload on specific Customer sites.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("site")}
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
deploy_ce_sites {
  # Configure direct properties listed below.
}
```

<a id="canonical-5970e86a6b1a99b94a246c658fecee606a2c8121057d3bbd53a0b078fea69335"></a>

## Direct properties — service.deploy_options.deploy_ce_sites / bb5dc8428de1 / 3

- [site](resources--workload--reference--group-016.md#canonical-d10026c06fbb5164d1ea440beed0d3b95b00857fe13633ec76a5e3b17ae299d4): complete subsection reference.

<a id="canonical-14dd04e1ccd76fec4fc3fc441c447680027e705ca112958bb17755c825d2aa5d"></a>

## Next pages — service.deploy_options.deploy_ce_sites / bb5dc8428de1 / 4

- [service.deploy_options.deploy_ce_sites.site](resources--workload--reference--group-016.md#canonical-d10026c06fbb5164d1ea440beed0d3b95b00857fe13633ec76a5e3b17ae299d4)
- [service.deploy_options](resources--workload--reference--group-016.md#canonical-05165ff47537505ed9ace4fa2b50ac0d2124ed1fc4414c6d288c2e658d5e2ec3)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-d10026c06fbb5164d1ea440beed0d3b95b00857fe13633ec76a5e3b17ae299d4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-37d4ec1cfecac530dcd2c6097f99f1297c7b6c3819871aaba0911b0051620214"></a>

## service.deploy_options.deploy_ce_sites.site — service.deploy_options.deploy_ce_sites.site / efed938aa665 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.deploy_options](resources--workload--reference--group-016.md#canonical-05165ff47537505ed9ace4fa2b50ac0d2124ed1fc4414c6d288c2e658d5e2ec3)
- [service.deploy_options.deploy_ce_sites](resources--workload--reference--group-016.md#canonical-8cdda23ecbd8b0a4afe693a8b8da2a0cdbf039b8155010802897620dd832a13b)
- service.deploy_options.deploy_ce_sites.site

<a id="canonical-ca5a2f2cbb7b8fc8f4bb436a4f2a0d05c5dcfa2e81d6433bb88e252b37ed757b"></a>

Type: `"object"`. list nested block, Optional.

Which customer sites should this workload be deployed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
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
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-e13f2c542553436af654755a09d4f9685f95eec05d774de04a8eb643ae117e5b"></a>

## Direct properties — service.deploy_options.deploy_ce_sites.site / efed938aa665 / 3

<a id="canonical-70d0e284b18a68c881b4415b5cd7755bcb6907d7223eed45b8e9caacd954aa1a"></a>

<a id="canonical-87ac90fa450a891fa401299695aded7ddacda3a7353efa916aad3c76e7952ac6"></a>

## name property — service.deploy_options.deploy_ce_sites.site / efed938aa665 / 4

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

<a id="canonical-c3a64e42de900f7f31639aa906eb2a5ad1ec1019c60f3a77907093a7b2705e0d"></a>

<a id="canonical-08146c77425859f4d5464d85cb879cd2ad9d2a30013fb37427f07648311733f2"></a>

## namespace property — service.deploy_options.deploy_ce_sites.site / efed938aa665 / 5

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

<a id="canonical-c5a74288f7bf2f153838ef0ea17fff4770d2d76239b37273cadea4dcab5a4a7b"></a>

<a id="canonical-668b410ff2add7e89f5f6ced279d2b9e6a9475c4f08c9661684265227c4e381e"></a>

## tenant property — service.deploy_options.deploy_ce_sites.site / efed938aa665 / 6

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

<a id="canonical-cacdfd075adc795a7756fdccb658562936fd3736da3e545d4ef878cfb86ee55d"></a>

## Next pages — service.deploy_options.deploy_ce_sites.site / efed938aa665 / 7

- [service.deploy_options.deploy_ce_sites](resources--workload--reference--group-016.md#canonical-8cdda23ecbd8b0a4afe693a8b8da2a0cdbf039b8155010802897620dd832a13b)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-4a3ccb108bbf4b6f11ead9d960371138448f69f60e24fc9e98af8bca71f40267"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9c39d5d13ef11b1c037eb0899ba7c8794da6e5c2e7ca997eb87adba06a146de8"></a>

## service.deploy_options.deploy_ce_virtual_sites — service.deploy_options.deploy_ce_virtual_sites / 1418b32d2a07 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.deploy_options](resources--workload--reference--group-016.md#canonical-05165ff47537505ed9ace4fa2b50ac0d2124ed1fc4414c6d288c2e658d5e2ec3)
- service.deploy_options.deploy_ce_virtual_sites

<a id="canonical-e8214dc251f6b80b3fa3064932f14b909abdab367686d4c9a12062954c43f454"></a>

Type: `"object"`. single nested block, Optional.

Defines a way to deploy a workload on specific Customer virtual sites.

Upstream description:

This defines a way to deploy a workload on specific Customer virtual sites.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("virtual_site")}
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
deploy_ce_virtual_sites {
  # Configure direct properties listed below.
}
```

<a id="canonical-60b84aabb436961d9cc898208f4e81dee0d18ea335142b9cc7c6eeaa3517272d"></a>

## Direct properties — service.deploy_options.deploy_ce_virtual_sites / 1418b32d2a07 / 3

- [virtual_site](resources--workload--reference--group-016.md#canonical-6185acb341fd36f71c46f3d0e05ca55c27fdffb2d08b695d989f8d893eb900d4): complete subsection reference.

<a id="canonical-c5674a6a38890bd2b40ba588a3da27fac40695b9ed23da21cbcb36dab900e65d"></a>

## Next pages — service.deploy_options.deploy_ce_virtual_sites / 1418b32d2a07 / 4

- [service.deploy_options.deploy_ce_virtual_sites.virtual_site](resources--workload--reference--group-016.md#canonical-6185acb341fd36f71c46f3d0e05ca55c27fdffb2d08b695d989f8d893eb900d4)
- [service.deploy_options](resources--workload--reference--group-016.md#canonical-05165ff47537505ed9ace4fa2b50ac0d2124ed1fc4414c6d288c2e658d5e2ec3)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-6185acb341fd36f71c46f3d0e05ca55c27fdffb2d08b695d989f8d893eb900d4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a372601c988de43d64820fb933389ba6468d9f81658e7f33ccd000bad3da92a2"></a>

## service.deploy_options.deploy_ce_virtual_sites.virtual_site — service.deploy_options.deploy_ce_virtual_sites.virtual_site / fcdd4b868ff3 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.deploy_options](resources--workload--reference--group-016.md#canonical-05165ff47537505ed9ace4fa2b50ac0d2124ed1fc4414c6d288c2e658d5e2ec3)
- [service.deploy_options.deploy_ce_virtual_sites](resources--workload--reference--group-016.md#canonical-4a3ccb108bbf4b6f11ead9d960371138448f69f60e24fc9e98af8bca71f40267)
- service.deploy_options.deploy_ce_virtual_sites.virtual_site

<a id="canonical-9e6cbbe98bbe7749212a75281e5058076b2e8585877cfa5942b695e194b81adb"></a>

Type: `"object"`. list nested block, Optional.

Which customer virtual sites should this workload be deployed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
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
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-7a9b629aaadc58a077b73fb86e9e09677cfb8e98764904ad6a2a9c9cd3c225e3"></a>

## Direct properties — service.deploy_options.deploy_ce_virtual_sites.virtual_site / fcdd4b868ff3 / 3

<a id="canonical-ab4dc15f64aa6dae39e124bb43f9076b5ab24c07b1999416d35724991e6990aa"></a>

<a id="canonical-a9fca4489af116ec8dc5b0f5b1950eea1abf62e2ed8745a39f4fc86b6f3ca15a"></a>

## name property — service.deploy_options.deploy_ce_virtual_sites.virtual_site / fcdd4b868ff3 / 4

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

<a id="canonical-0086044ff536da2018963979f0749d9221db0879b3fcf1f37f8c938139064242"></a>

<a id="canonical-18ba64a94020c253302d57abe791afc54e28dcb94abe747291897ecd413924cc"></a>

## namespace property — service.deploy_options.deploy_ce_virtual_sites.virtual_site / fcdd4b868ff3 / 5

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

<a id="canonical-ae88cfb74ab845f3f13226a925e13ed39185c30cedf681e2bba45e3391990166"></a>

<a id="canonical-286660813c0a91cdaf3aeaea4d78e409137884bb7c9bfd7bb0722ff79bfb7fc6"></a>

## tenant property — service.deploy_options.deploy_ce_virtual_sites.virtual_site / fcdd4b868ff3 / 6

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

<a id="canonical-c696f47a4a09161a2a0afbb9cb5448a64e8e90c697e046ce03d1a4243b904ee2"></a>

## Next pages — service.deploy_options.deploy_ce_virtual_sites.virtual_site / fcdd4b868ff3 / 7

- [service.deploy_options.deploy_ce_virtual_sites](resources--workload--reference--group-016.md#canonical-4a3ccb108bbf4b6f11ead9d960371138448f69f60e24fc9e98af8bca71f40267)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-3f5bcf3835ef9e50f6b01e38fb785dd9d787df2434acc63af699d0898e5e860c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-43fd9a24b6437ed4747b93f178a007cefa6a5c241c09c643aa402761e1f44437"></a>

## service.deploy_options.deploy_re_sites — service.deploy_options.deploy_re_sites / 75a8bdbf94b3 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.deploy_options](resources--workload--reference--group-016.md#canonical-05165ff47537505ed9ace4fa2b50ac0d2124ed1fc4414c6d288c2e658d5e2ec3)
- service.deploy_options.deploy_re_sites

<a id="canonical-c7eaf9cab44ebb20733700fafb1ade87250d4c79a8229ea47cd4041efe5b707f"></a>

Type: `"object"`. single nested block, Optional.

Defines a way to deploy a workload on specific Regional Edge sites.

Upstream description:

This defines a way to deploy a workload on specific Regional Edge sites.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("site")}
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
deploy_re_sites {
  # Configure direct properties listed below.
}
```

<a id="canonical-c5a485af239ef141d42e3d2390cea51fe7a68bfaa6c56070ac175144455f1717"></a>

## Direct properties — service.deploy_options.deploy_re_sites / 75a8bdbf94b3 / 3

- [site](resources--workload--reference--group-016.md#canonical-8f74cd8ba05aaa74aabab185206618edcf4ff14e3d2a07d1ac3b179986d6cdb6): complete subsection reference.

<a id="canonical-3f7485d777e3d4c5691549b5f8856617284cce5c413aec9c053f14e7faea5118"></a>

## Next pages — service.deploy_options.deploy_re_sites / 75a8bdbf94b3 / 4

- [service.deploy_options.deploy_re_sites.site](resources--workload--reference--group-016.md#canonical-8f74cd8ba05aaa74aabab185206618edcf4ff14e3d2a07d1ac3b179986d6cdb6)
- [service.deploy_options](resources--workload--reference--group-016.md#canonical-05165ff47537505ed9ace4fa2b50ac0d2124ed1fc4414c6d288c2e658d5e2ec3)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-8f74cd8ba05aaa74aabab185206618edcf4ff14e3d2a07d1ac3b179986d6cdb6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9ab1c3642285f6eb164b1f8ef00ca813e50365ad9a98eca33e4d30375a19d913"></a>

## service.deploy_options.deploy_re_sites.site — service.deploy_options.deploy_re_sites.site / 201969486dc5 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.deploy_options](resources--workload--reference--group-016.md#canonical-05165ff47537505ed9ace4fa2b50ac0d2124ed1fc4414c6d288c2e658d5e2ec3)
- [service.deploy_options.deploy_re_sites](resources--workload--reference--group-016.md#canonical-3f5bcf3835ef9e50f6b01e38fb785dd9d787df2434acc63af699d0898e5e860c)
- service.deploy_options.deploy_re_sites.site

<a id="canonical-f928c71842511f72108f8695d5012800faa56c5997515fa491c4856a8ca42269"></a>

Type: `"object"`. list nested block, Optional.

Which regional edge sites should this workload be deployed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
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
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-7cd9fc5e5d4194dacef94fd4ac6e6f2d7ebd6d0e9e328cb3ecdd0c8740b4ba91"></a>

## Direct properties — service.deploy_options.deploy_re_sites.site / 201969486dc5 / 3

<a id="canonical-17d44de5479e2e6d872c5b24ee89af304ccb6e3e49ea75b3ea51e7e49514fed9"></a>

<a id="canonical-96368da9f0bc65a3c09134605f50eb1217495156dd70f3ac10644e40e2070532"></a>

## name property — service.deploy_options.deploy_re_sites.site / 201969486dc5 / 4

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

<a id="canonical-23e49911e3034f5585d5a7095c57af95fd4e6b15e6e2782a878c4b4c73ad8df4"></a>

<a id="canonical-abcd68bbf50f8335033425c45a880ad64f081f856affdb6fb615841e1a716dca"></a>

## namespace property — service.deploy_options.deploy_re_sites.site / 201969486dc5 / 5

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

<a id="canonical-5b83381223b3a9d44942222f4142ab96649f72b80fe3be6a49e3aba5abeb398d"></a>

<a id="canonical-a5913d943554ce832716abb93caf4d8ed1a558c73f8efb49f097f5d3232a9c91"></a>

## tenant property — service.deploy_options.deploy_re_sites.site / 201969486dc5 / 6

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

<a id="canonical-25e04e2bd7cb146dd6ab817e44e8b67a694e78e209ba891cc849647ef1a9c9b9"></a>

## Next pages — service.deploy_options.deploy_re_sites.site / 201969486dc5 / 7

- [service.deploy_options.deploy_re_sites](resources--workload--reference--group-016.md#canonical-3f5bcf3835ef9e50f6b01e38fb785dd9d787df2434acc63af699d0898e5e860c)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-8a595d976bf182b2d6306ad10d47bd00245778b64d1589966eda5bfe8517a762"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e5998bb986796241d8bc43b7714985c76f08a55980b9fbd736ec70d9e5f14fbb"></a>

## service.deploy_options.deploy_re_virtual_sites — service.deploy_options.deploy_re_virtual_sites / fb73a650fd96 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.deploy_options](resources--workload--reference--group-016.md#canonical-05165ff47537505ed9ace4fa2b50ac0d2124ed1fc4414c6d288c2e658d5e2ec3)
- service.deploy_options.deploy_re_virtual_sites

<a id="canonical-2aa895636df54d5ef3628cb805a644782c512c2a5a232c989b74c327f8b00f6b"></a>

Type: `"object"`. single nested block, Optional.

Defines a way to deploy a workload on specific Regional Edge virtual sites.

Upstream description:

This defines a way to deploy a workload on specific Regional Edge virtual sites.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("virtual_site")}
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
deploy_re_virtual_sites {
  # Configure direct properties listed below.
}
```

<a id="canonical-687a3eec80e73334dda7d4d7f77deb8441969089fabd9c83b50bbe7f7e56e030"></a>

## Direct properties — service.deploy_options.deploy_re_virtual_sites / fb73a650fd96 / 3

- [virtual_site](resources--workload--reference--group-016.md#canonical-baa3d3e053cc85acc71419a7c4c6e45b87dc8e6ac53abcd4996cc77b2d3df7e0): complete subsection reference.

<a id="canonical-2abf4b3c2a64c6d09e1da3de35ea5c5f6d19cbbd9375fad33c3f13e741c12135"></a>

## Next pages — service.deploy_options.deploy_re_virtual_sites / fb73a650fd96 / 4

- [service.deploy_options.deploy_re_virtual_sites.virtual_site](resources--workload--reference--group-016.md#canonical-baa3d3e053cc85acc71419a7c4c6e45b87dc8e6ac53abcd4996cc77b2d3df7e0)
- [service.deploy_options](resources--workload--reference--group-016.md#canonical-05165ff47537505ed9ace4fa2b50ac0d2124ed1fc4414c6d288c2e658d5e2ec3)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-baa3d3e053cc85acc71419a7c4c6e45b87dc8e6ac53abcd4996cc77b2d3df7e0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a3b282a1f85c8104bd7c4376d5ede00b0dfd17e70ebe93b70eddc3345bf86b6c"></a>

## service.deploy_options.deploy_re_virtual_sites.virtual_site — service.deploy_options.deploy_re_virtual_sites.virtual_site / 39baf4cf7ae1 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.deploy_options](resources--workload--reference--group-016.md#canonical-05165ff47537505ed9ace4fa2b50ac0d2124ed1fc4414c6d288c2e658d5e2ec3)
- [service.deploy_options.deploy_re_virtual_sites](resources--workload--reference--group-016.md#canonical-8a595d976bf182b2d6306ad10d47bd00245778b64d1589966eda5bfe8517a762)
- service.deploy_options.deploy_re_virtual_sites.virtual_site

<a id="canonical-da6589fa79237c67b32347bf12b2d261f71681d960f61034cf76130db7a8a1cb"></a>

Type: `"object"`. list nested block, Optional.

Which regional edge virtual sites should this workload be deployed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
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
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-fd014f38a66e17f3c6cfff7621780f86acbd94ee98acb58837cd8af4f96b1a8b"></a>

## Direct properties — service.deploy_options.deploy_re_virtual_sites.virtual_site / 39baf4cf7ae1 / 3

<a id="canonical-c3ede1a61ecfb3032382dd4581fdecccc6d1b1c88a4c48353110995d9a98a5fb"></a>

<a id="canonical-bb785777a11f7e320ba0cdb1ee7d6bdbe58770a265bb1edafb374ec8e50a19b7"></a>

## name property — service.deploy_options.deploy_re_virtual_sites.virtual_site / 39baf4cf7ae1 / 4

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

<a id="canonical-b58c9c87e353ad7962287cf0b5b99c7d6edd3cf522e7cd5757b3e8219b2ac796"></a>

<a id="canonical-48b2726353d15944bd25015205d27e778a5012cab46c0baea013446860d67ccf"></a>

## namespace property — service.deploy_options.deploy_re_virtual_sites.virtual_site / 39baf4cf7ae1 / 5

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

<a id="canonical-6af0a4d8a2de294ae320b25be4f3a92fd9d41b1dfc2732eba6d308948776937c"></a>

<a id="canonical-1b5b1d9001095d20cf1860298fe28d1d8f7a93c35c37e180e32709b727a962f4"></a>

## tenant property — service.deploy_options.deploy_re_virtual_sites.virtual_site / 39baf4cf7ae1 / 6

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

<a id="canonical-ffbed8e5b7a37198b45f1361e6dd5da4151c1c5c2dde0578a352f659bbb4a655"></a>

## Next pages — service.deploy_options.deploy_re_virtual_sites.virtual_site / 39baf4cf7ae1 / 7

- [service.deploy_options.deploy_re_virtual_sites](resources--workload--reference--group-016.md#canonical-8a595d976bf182b2d6306ad10d47bd00245778b64d1589966eda5bfe8517a762)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-968fb38b4a3e5a07bd544d6620a48cbb9b651535a5806eb59a7a0bb1caaa9125"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6d824f009e3434874687bab1282bb2611fc277d3fb2831f4cc2c6f049d61e37b"></a>

## service.scale_to_zero — service.scale_to_zero / f6e4f3c477cc / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- service.scale_to_zero

<a id="canonical-c6c399995e06958dbf283a8bd515cd06cd894e30c5a3a0ecf20c2ea105ca1382"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
scale_to_zero = {}
```

<a id="canonical-d4fe51f26a95f17a9de98b24508efed41e117109c6a48337387234c6f4b2541f"></a>

## Direct properties — service.scale_to_zero / f6e4f3c477cc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c41dd511bf02ff029eac285006aab11d88fe5d166df30cbac72c787ae9cbf6a0"></a>

## Next pages — service.scale_to_zero / f6e4f3c477cc / 4

- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-faacd6685afddb473ef64bdf75dbfc0881f5d7702ee057ad8724fc6f0339f0c3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-88a30d21c596803c3454fa30f5559df81a85af13e49b8c635532b7cbeb92e63f"></a>

## service.volumes — service.volumes / f67c19dac250 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- service.volumes

<a id="canonical-3b5ad9968d3cbd3c15653729301abb85331ef3055766d0e0c6993b3d65ef59ff"></a>

Type: `"object"`. list nested block, Optional.

Volumes. Volumes for the service.

Upstream description:

Volumes for the service.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("empty_dir",
    "host_path"),
  validators.ConflictingListObjectAttributes("empty_dir",
    "persistent_volume"),
  validators.ConflictingListObjectAttributes("host_path",
    "persistent_volume")}
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
volumes {
  # Configure direct properties listed below.
}
```

<a id="canonical-0fd053c7b186369323d37de3bd09a68ddb84437a5a60da2514709464640749c6"></a>

## Direct properties — service.volumes / f67c19dac250 / 3

- [empty_dir](resources--workload--reference--group-016.md#canonical-69acc3b0b963fa09d8999143bb52bfafef29754470b5ee54bb3fbaaea878eb00): complete subsection reference.

- [host_path](resources--workload--reference--group-016.md#canonical-b5d7aa2d8cc2c66c3696a5b8c67e6bc5c0c6cfaf281d6c723a2209f673881d2b): complete subsection reference.

<a id="canonical-24e5a53efe72cf2b34e4fabd584a4bb757e218675ee0ca770f7330c6316618f0"></a>

<a id="canonical-c3325059d20d373bc9b1dedb95a12b3fd5c5850dfc6b42624582975ad04f081b"></a>

## name property — service.volumes / f67c19dac250 / 4

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.dns_1123_label": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.dns_1123_label": "true"
  }
}
```

- [persistent_volume](resources--workload--reference--group-016.md#canonical-d5df8ef3ffdec7ca5eeb2b10ee386c06708c26027ffb0fef7e9bd1bb6b5195fd): complete subsection reference.

<a id="canonical-b627276d0a3d49f0192484292becdff007fe8e0648f6f27fe21b6b209ba73821"></a>

## Next pages — service.volumes / f67c19dac250 / 5

- [service.volumes.empty_dir](resources--workload--reference--group-016.md#canonical-69acc3b0b963fa09d8999143bb52bfafef29754470b5ee54bb3fbaaea878eb00)
- [service.volumes.host_path](resources--workload--reference--group-016.md#canonical-b5d7aa2d8cc2c66c3696a5b8c67e6bc5c0c6cfaf281d6c723a2209f673881d2b)
- [service.volumes.persistent_volume](resources--workload--reference--group-016.md#canonical-d5df8ef3ffdec7ca5eeb2b10ee386c06708c26027ffb0fef7e9bd1bb6b5195fd)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-69acc3b0b963fa09d8999143bb52bfafef29754470b5ee54bb3fbaaea878eb00"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-01853013234005d1262a5f869106dfe3c5ba07d220ecef60f1cc243728817076"></a>

## service.volumes.empty_dir — service.volumes.empty_dir / 0692436069ea / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.volumes](resources--workload--reference--group-016.md#canonical-faacd6685afddb473ef64bdf75dbfc0881f5d7702ee057ad8724fc6f0339f0c3)
- service.volumes.empty_dir

<a id="canonical-07c1f4105751b679d58da500c1b38cb94fffb3c592feaa2b298d0efe5b21b07d"></a>

Type: `"object"`. single nested block, Optional.

Volume containing a temporary directory whose lifetime is the same as a replica of a workload.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("size_limit")}
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
empty_dir {
  # Configure direct properties listed below.
}
```

<a id="canonical-a3d8ca7bcdf84dd81f35408bda9aac3f8d1d6d5e374207123a83ec68a3f03f50"></a>

## Direct properties — service.volumes.empty_dir / 0692436069ea / 3

- [mount](resources--workload--reference--group-016.md#canonical-e2749766119cbc462da25ec8bd481c0148be73d66640e4e062536783a5c5bdee): complete subsection reference.

<a id="canonical-1ef805e7589637bee0cdef7a01e5e8a3ad7810fd763216e6a74ee1d36d94deaf"></a>

<a id="canonical-7f98342b9531d9fc7d5d10edcdf52e7d7807a915bc82a60de8a41fdb6c6bae74"></a>

## size_limit property — service.volumes.empty_dir / 0692436069ea / 4

Type: `"number"`. Optional.

Size Limit (in GiB). Configuration parameter for size limit

Upstream description:

Configuration parameter for size limit

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
    "ves.io.schema.rules.double.lte": "10",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.double.lte": "10",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-5db3561e6217fd6228267158a97ef60b14485e4358f20cf6b14985ca261f7548"></a>

## Next pages — service.volumes.empty_dir / 0692436069ea / 5

- [service.volumes.empty_dir.mount](resources--workload--reference--group-016.md#canonical-e2749766119cbc462da25ec8bd481c0148be73d66640e4e062536783a5c5bdee)
- [service.volumes](resources--workload--reference--group-016.md#canonical-faacd6685afddb473ef64bdf75dbfc0881f5d7702ee057ad8724fc6f0339f0c3)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-e2749766119cbc462da25ec8bd481c0148be73d66640e4e062536783a5c5bdee"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1639d39b69e66142dd338e420ef209357f78b6b8a1eaf4a84f0a2cf046707aad"></a>

## service.volumes.empty_dir.mount — service.volumes.empty_dir.mount / 55c7d4b3c4ef / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.volumes](resources--workload--reference--group-016.md#canonical-faacd6685afddb473ef64bdf75dbfc0881f5d7702ee057ad8724fc6f0339f0c3)
- [service.volumes.empty_dir](resources--workload--reference--group-016.md#canonical-69acc3b0b963fa09d8999143bb52bfafef29754470b5ee54bb3fbaaea878eb00)
- service.volumes.empty_dir.mount

<a id="canonical-e0537e18711ae2b2c281ba9e5799f3ebaa18c6aa3d47c1d8e7c6bd9edb42a106"></a>

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

<a id="canonical-9de43a3cc22a29599da7bc4b96607bd2f3c9b55aeed45b84ff2b544ff078cb89"></a>

## Direct properties — service.volumes.empty_dir.mount / 55c7d4b3c4ef / 3

<a id="canonical-05d66e8cc09d0fc3331c75394899539e3065e8d88a7e3c57733eea18fe9b1621"></a>

<a id="canonical-b87e35c4a9853ed8cabb9baed996a1ccfe94acb27908f33347d0f3080a18c6a1"></a>

## mode property — service.volumes.empty_dir.mount / 55c7d4b3c4ef / 4

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

<a id="canonical-96ac1bcc4e24ffdd1a0791262c6990acd236d8fbb4a5bb4a3f342fe198d5120a"></a>

<a id="canonical-691468b19fff10cc3b6ce0f7ee5b65be56153881896c748db41ebfa829f40dc4"></a>

## mount_path property — service.volumes.empty_dir.mount / 55c7d4b3c4ef / 5

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

<a id="canonical-6b38303220571391e34e561ca5c77568ef19d3f9876bda1f144a5e4f6773b3c0"></a>

<a id="canonical-2b550457f4bdabd1df4f59c694a35a5c6fba3b61692f81213addf7d1d2f8e704"></a>

## sub_path property — service.volumes.empty_dir.mount / 55c7d4b3c4ef / 6

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

<a id="canonical-bd80ead151d2932cefd122b4784c85bcbe91cd4c540728e7f602cc03f1af204f"></a>

## Next pages — service.volumes.empty_dir.mount / 55c7d4b3c4ef / 7

- [service.volumes.empty_dir](resources--workload--reference--group-016.md#canonical-69acc3b0b963fa09d8999143bb52bfafef29754470b5ee54bb3fbaaea878eb00)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-b5d7aa2d8cc2c66c3696a5b8c67e6bc5c0c6cfaf281d6c723a2209f673881d2b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3d7b51af8479c12f8edcbe84e0a0a55ee9e071f10a1c8c57a2998b9726dc7f16"></a>

## service.volumes.host_path — service.volumes.host_path / d4c04e859804 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.volumes](resources--workload--reference--group-016.md#canonical-faacd6685afddb473ef64bdf75dbfc0881f5d7702ee057ad8724fc6f0339f0c3)
- service.volumes.host_path

<a id="canonical-1252636d998e04ac1f8c92ea74f5061859387f229b6747dc3590ece30bd69d01"></a>

Type: `"object"`. single nested block, Optional.

Volume containing a host mapped path into the workload.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("path")}
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
host_path {
  # Configure direct properties listed below.
}
```

<a id="canonical-2069e68f5efe06cff3023a696fa995d80a3bb394a7411249f196635de304827d"></a>

## Direct properties — service.volumes.host_path / d4c04e859804 / 3

- [mount](resources--workload--reference--group-016.md#canonical-73774b9285fd1b1259fc744dd21516b338556e4152dca95ea4cc470d2390ed7f): complete subsection reference.

<a id="canonical-64c778c5d76d83daab75425ccdd05404ddbd7fd6ab92e5e1b2ed76b66be4cb73"></a>

<a id="canonical-e15458ee92a6b5a779198c09717fee2a1c9d0e8602fdbac788b2b75bcd49c9a7"></a>

## path property — service.volumes.host_path / d4c04e859804 / 4

Type: `"string"`. Optional.

Path. Path of the directory on the host.

Upstream description:

Path of the directory on the host.

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
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "[^\\\\0]+"
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
    "ves.io.schema.rules.string.pattern": "[^\\\\0]+"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.pattern": "[^\\\\0]+"
  }
}
```

<a id="canonical-bd0e0e7fcea7f63c7e4f2fed62bb63e48e44f3900ebb912277710a2a6542867e"></a>

## Next pages — service.volumes.host_path / d4c04e859804 / 5

- [service.volumes.host_path.mount](resources--workload--reference--group-016.md#canonical-73774b9285fd1b1259fc744dd21516b338556e4152dca95ea4cc470d2390ed7f)
- [service.volumes](resources--workload--reference--group-016.md#canonical-faacd6685afddb473ef64bdf75dbfc0881f5d7702ee057ad8724fc6f0339f0c3)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-73774b9285fd1b1259fc744dd21516b338556e4152dca95ea4cc470d2390ed7f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-731c089fc7dea0b3b34fb5c66172b710f3860548f8c0f92a6622d082ca7a9184"></a>

## service.volumes.host_path.mount — service.volumes.host_path.mount / c8b2ced8283a / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.volumes](resources--workload--reference--group-016.md#canonical-faacd6685afddb473ef64bdf75dbfc0881f5d7702ee057ad8724fc6f0339f0c3)
- [service.volumes.host_path](resources--workload--reference--group-016.md#canonical-b5d7aa2d8cc2c66c3696a5b8c67e6bc5c0c6cfaf281d6c723a2209f673881d2b)
- service.volumes.host_path.mount

<a id="canonical-394502dd8498835ed4d6787c846a968748fd2eb252b19fa8b1d83eff8774a047"></a>

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

<a id="canonical-bff07712c729234871f0fa6ecf277848a8c78076394907450389e8dcff8f416b"></a>

## Direct properties — service.volumes.host_path.mount / c8b2ced8283a / 3

<a id="canonical-547bbd29b54cbd0718d243db5ecde6e415efb22faa68df0e915f20738c6c617e"></a>

<a id="canonical-e756941d1373a61ddb108cdb53028751f0887b0b049836256536f3fb5f7caca9"></a>

## mode property — service.volumes.host_path.mount / c8b2ced8283a / 4

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

<a id="canonical-0392e5c2483eedfa3bcc67fbd00d100e6e951edff02bcb876b6b5f222fd14375"></a>

<a id="canonical-e4cceb24d934bbfa872102fcf84fd8a7d24631944f49dd8b0ea247b5f57d619e"></a>

## mount_path property — service.volumes.host_path.mount / c8b2ced8283a / 5

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

<a id="canonical-e947f2378f6e1212052ac32ed8c7ec22dda3a0c10a3a1e44e1b0c7bb3163d090"></a>

<a id="canonical-cb0455464eea073d788da7d2426b6eb6514c11a6a629993f02d444206a422636"></a>

## sub_path property — service.volumes.host_path.mount / c8b2ced8283a / 6

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

<a id="canonical-19f7812b1a13ce194c55cbfc3c3ef3ef3b1e800be8fb8a98a90b45aa9f3c59c8"></a>

## Next pages — service.volumes.host_path.mount / c8b2ced8283a / 7

- [service.volumes.host_path](resources--workload--reference--group-016.md#canonical-b5d7aa2d8cc2c66c3696a5b8c67e6bc5c0c6cfaf281d6c723a2209f673881d2b)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-d5df8ef3ffdec7ca5eeb2b10ee386c06708c26027ffb0fef7e9bd1bb6b5195fd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bc1b693fe323be9d8f58b9b551c2fd4ead31992f37622907fc6e2ad1463d9693"></a>

## service.volumes.persistent_volume — service.volumes.persistent_volume / e8b2d13c1e4e / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.volumes](resources--workload--reference--group-016.md#canonical-faacd6685afddb473ef64bdf75dbfc0881f5d7702ee057ad8724fc6f0339f0c3)
- service.volumes.persistent_volume

<a id="canonical-3d33f500f531a323f706e63871576cba91308f08138a9cda0753ff2b5400add1"></a>

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

<a id="canonical-1bd22d6bb1a8f14f08f1cc4613e6d10557af642d661178aad14c952fc4c317dd"></a>

## Direct properties — service.volumes.persistent_volume / e8b2d13c1e4e / 3

- [mount](resources--workload--reference--group-016.md#canonical-2c9737889c737e1cc8b425c349eb04d2281f8b50094e599d9a3349f5d641b9c6): complete subsection reference.

- [storage](resources--workload--reference--group-016.md#canonical-a537fefc83c6c2512646fff280f28ece8cb3f1ae1ce20bba56b5da158fdf22b6): complete subsection reference.

<a id="canonical-51835591741da34d79b0e252123a7e29875885c682dc347e40075d7857faeb1d"></a>

## Next pages — service.volumes.persistent_volume / e8b2d13c1e4e / 4

- [service.volumes.persistent_volume.mount](resources--workload--reference--group-016.md#canonical-2c9737889c737e1cc8b425c349eb04d2281f8b50094e599d9a3349f5d641b9c6)
- [service.volumes.persistent_volume.storage](resources--workload--reference--group-016.md#canonical-a537fefc83c6c2512646fff280f28ece8cb3f1ae1ce20bba56b5da158fdf22b6)
- [service.volumes](resources--workload--reference--group-016.md#canonical-faacd6685afddb473ef64bdf75dbfc0881f5d7702ee057ad8724fc6f0339f0c3)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-2c9737889c737e1cc8b425c349eb04d2281f8b50094e599d9a3349f5d641b9c6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-890751f928aeeb729da2f4ed4e9103c787ff26729d474d72201cc66562c8cf90"></a>

## service.volumes.persistent_volume.mount — service.volumes.persistent_volume.mount / 9515ad65bcd9 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.volumes](resources--workload--reference--group-016.md#canonical-faacd6685afddb473ef64bdf75dbfc0881f5d7702ee057ad8724fc6f0339f0c3)
- [service.volumes.persistent_volume](resources--workload--reference--group-016.md#canonical-d5df8ef3ffdec7ca5eeb2b10ee386c06708c26027ffb0fef7e9bd1bb6b5195fd)
- service.volumes.persistent_volume.mount

<a id="canonical-a1dac898a6cc6f0a7645ff99da30875d62b18f9c50f4a382d878b9d777937752"></a>

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

<a id="canonical-7b1eccf69826676be4497ee891a12ee7e7d12079f76fc338c8b335d9491ccb9a"></a>

## Direct properties — service.volumes.persistent_volume.mount / 9515ad65bcd9 / 3

<a id="canonical-63cb2200cfa185f9df4e97e1c96bf3612df7f84d6eeb6d4d8e283d45f3960feb"></a>

<a id="canonical-a807802459e782b33153fc2e3aa8a6e3433bc8f75da7c404f7dcaf23fca34b28"></a>

## mode property — service.volumes.persistent_volume.mount / 9515ad65bcd9 / 4

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

<a id="canonical-96cf97f674fc8102298b350e3bd73ebb496ab41fba8e97682858b770e966e096"></a>

<a id="canonical-2e65b9947999784f881023940c4795164a9700a0c865af7298adf0ed24777307"></a>

## mount_path property — service.volumes.persistent_volume.mount / 9515ad65bcd9 / 5

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

<a id="canonical-2724ab2be3bd179833a967151f688fb1204e0806e401aa23b8edac86c74c4868"></a>

<a id="canonical-0a6c777546c29a87f9fae5666920b087f85e9b7f61d4982b9f2124a1e932c5a1"></a>

## sub_path property — service.volumes.persistent_volume.mount / 9515ad65bcd9 / 6

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

<a id="canonical-4752ad17c86fde0c48e4709d3243ebfbace00c70d9d2ad8960b0c7a022074eaf"></a>

## Next pages — service.volumes.persistent_volume.mount / 9515ad65bcd9 / 7

- [service.volumes.persistent_volume](resources--workload--reference--group-016.md#canonical-d5df8ef3ffdec7ca5eeb2b10ee386c06708c26027ffb0fef7e9bd1bb6b5195fd)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-a537fefc83c6c2512646fff280f28ece8cb3f1ae1ce20bba56b5da158fdf22b6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2813d99a78409de72134d3c9abf5f46c8c2d104225157bc8645d9f1158fad103"></a>

## service.volumes.persistent_volume.storage — service.volumes.persistent_volume.storage / 7ba64d82f402 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.volumes](resources--workload--reference--group-016.md#canonical-faacd6685afddb473ef64bdf75dbfc0881f5d7702ee057ad8724fc6f0339f0c3)
- [service.volumes.persistent_volume](resources--workload--reference--group-016.md#canonical-d5df8ef3ffdec7ca5eeb2b10ee386c06708c26027ffb0fef7e9bd1bb6b5195fd)
- service.volumes.persistent_volume.storage

<a id="canonical-5f96c172ed9c51124e7b4b9967cf3ba77f9524ae42fc850ff1f823702f769682"></a>

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

<a id="canonical-bc0846b00cc0a26b2f02f74a2c92d1d092a710af9563948188a02e69475f757e"></a>

## Direct properties — service.volumes.persistent_volume.storage / 7ba64d82f402 / 3

<a id="canonical-19158f7c956643d2cc524a896343c7a027a22ce18233b975a395a9f321f3c492"></a>

<a id="canonical-be13e80d823c8b70a229ea6de0cceae5ee2615782826878ae278dfa164b923e4"></a>

## access_mode property — service.volumes.persistent_volume.storage / 7ba64d82f402 / 4

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

<a id="canonical-917048d2b27c1b2ea4e4f767b0bc64f2b81fd7250e501793673e7acff3b3fa21"></a>

<a id="canonical-d5c8fb2d1dcc4ee05323a84ad943839c291dae6df5aefbaf2a3f23eed64a03c1"></a>

## class_name property — service.volumes.persistent_volume.storage / 7ba64d82f402 / 5

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

- [default](resources--workload--reference--group-016.md#canonical-cc26e49ece6e898c66202fb663d5aa7ab1d5fbe15dc26bb67df582e6136ef3d1): complete subsection reference.

<a id="canonical-4017995b7c9581471599b6c43cbb0e8eb3451eb0fbc2312362cb96b7f5941200"></a>

<a id="canonical-a3887e5b16b32de38d755f8303ae282aef5198ba4f443cac5dd36dd8dd63f1c4"></a>

## storage_size property — service.volumes.persistent_volume.storage / 7ba64d82f402 / 6

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

<a id="canonical-92d44172c7e27eb26aa1615ae9b8121656365520082b787c72f7a918dfa2ecd1"></a>

## Next pages — service.volumes.persistent_volume.storage / 7ba64d82f402 / 7

- [service.volumes.persistent_volume.storage.default](resources--workload--reference--group-016.md#canonical-cc26e49ece6e898c66202fb663d5aa7ab1d5fbe15dc26bb67df582e6136ef3d1)
- [service.volumes.persistent_volume](resources--workload--reference--group-016.md#canonical-d5df8ef3ffdec7ca5eeb2b10ee386c06708c26027ffb0fef7e9bd1bb6b5195fd)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-cc26e49ece6e898c66202fb663d5aa7ab1d5fbe15dc26bb67df582e6136ef3d1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5a93af146ada9f2cd1f83bbfd982b3f675b42e97f5a213898f97fdfc7d125b31"></a>

## service.volumes.persistent_volume.storage.default — service.volumes.persistent_volume.storage.default / cd0bad782061 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.volumes](resources--workload--reference--group-016.md#canonical-faacd6685afddb473ef64bdf75dbfc0881f5d7702ee057ad8724fc6f0339f0c3)
- [service.volumes.persistent_volume](resources--workload--reference--group-016.md#canonical-d5df8ef3ffdec7ca5eeb2b10ee386c06708c26027ffb0fef7e9bd1bb6b5195fd)
- [service.volumes.persistent_volume.storage](resources--workload--reference--group-016.md#canonical-a537fefc83c6c2512646fff280f28ece8cb3f1ae1ce20bba56b5da158fdf22b6)
- service.volumes.persistent_volume.storage.default

<a id="canonical-eeed58d90af2c7899a2416a1759b0b26e3e0242fb4a84166daec1d4fb70842ea"></a>

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

<a id="canonical-e5e15d5810c6f996d9c5125cb1769c32a8fc7037eacb812c24030c500baa1e6a"></a>

## Direct properties — service.volumes.persistent_volume.storage.default / cd0bad782061 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a4c8b8b5762c4086572cd5642d4d7feeb1c6b03fb728d27b90dd319d7d1e7d99"></a>

## Next pages — service.volumes.persistent_volume.storage.default / cd0bad782061 / 4

- [service.volumes.persistent_volume.storage](resources--workload--reference--group-016.md#canonical-a537fefc83c6c2512646fff280f28ece8cb3f1ae1ce20bba56b5da158fdf22b6)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-9361b00a2d9b5209c533e92d6cb37dd367b4eb50395ea573285491ecbf8e6419"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-230a8c8eddc7cd5403087046830178d7f05cd13b2ada5c8186626d30ef14efdb"></a>

## simple_service — simple_service / e302d8cddbd3 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- simple_service

<a id="canonical-cff3ed51e45c2052315515d0cb4605df00bfd02f40a708551fa14bd5c64d543a"></a>

Type: `"object"`. single nested block, Optional.

SimpleService is a service having one container and one replica that is deployed on all Regional
Edges and advertised on Internet via HTTP loadbalancer on default VIP.

Upstream description:

SimpleService is a service having one container and one replica that is deployed on all Regional
Edges and advertised on Internet via HTTP loadbalancer on default VIP.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disabled",
    "enabled"),
  validators.ConflictingObjectAttributes("do_not_advertise",
    "simple_advertise")}
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
  "x-ves-oneof-field-advertise_choice": "[\"do_not_advertise\",\"simple_advertise\"]",
  "x-ves-oneof-field-persistence_choice": "[\"disabled\",\"enabled\"]"
}
```

Terraform syntax:

```terraform
simple_service {
  # Configure direct properties listed below.
}
```

<a id="canonical-01cb9b357357f788783040733148dedb90eb19bd52c0335552c3be5a149db248"></a>

## Direct properties — simple_service / e302d8cddbd3 / 3

- [configuration](resources--workload--reference--group-016.md#canonical-3898acb911cd1464ffad4a15e5db50b80e5c9c97e8ef46fc21e2955127008d77): complete subsection reference.

- [container](resources--workload--reference--group-016.md#canonical-45cc02935ab54e5ce267d41971e17998e8f55556582a81b7f71e3e643ef36327): complete subsection reference.

- [disabled](resources--workload--reference--group-017.md#canonical-542f5dc522026b382425cb8b3be92f1ceb6739b98fc9c2c9d8048f424d6bc1c8): complete subsection reference.

- [do_not_advertise](resources--workload--reference--group-017.md#canonical-03c266f90c295a8b4c2aea495e36645eec3865f573b8a8bdd54ffd16f42497ac): complete subsection reference.

- [enabled](resources--workload--reference--group-017.md#canonical-0a63e2c8b641721f1b34e55cf2ff86a96e0ff98f533b2effad0740e534e1930c): complete subsection reference.

<a id="canonical-4275e8d15db919fc0eebcd15a50719bfadc26106bea8893ad4f3a6de829994f2"></a>

<a id="canonical-67c7ed0002fb3a4e4877ff57c1b50681c1f8e7359cebd94e132a07116514e88a"></a>

## scale_to_zero property — simple_service / e302d8cddbd3 / 4

Type: `"bool"`. Optional.

Scale down replicas of the service to zero.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [simple_advertise](resources--workload--reference--group-017.md#canonical-f4cb6c4627ef2f328600ec5bc5b3170189481279b9545f5640eb0f1fe61e5b22): complete subsection reference.

<a id="canonical-6bac4c9bb2313a6087d3e343a6f2ef2ed9efd3908ac91a4cfaa102f429bbf264"></a>

## Next pages — simple_service / e302d8cddbd3 / 5

- [simple_service.configuration](resources--workload--reference--group-016.md#canonical-3898acb911cd1464ffad4a15e5db50b80e5c9c97e8ef46fc21e2955127008d77)
- [simple_service.container](resources--workload--reference--group-016.md#canonical-45cc02935ab54e5ce267d41971e17998e8f55556582a81b7f71e3e643ef36327)
- [simple_service.disabled](resources--workload--reference--group-017.md#canonical-542f5dc522026b382425cb8b3be92f1ceb6739b98fc9c2c9d8048f424d6bc1c8)
- [simple_service.do_not_advertise](resources--workload--reference--group-017.md#canonical-03c266f90c295a8b4c2aea495e36645eec3865f573b8a8bdd54ffd16f42497ac)
- [simple_service.enabled](resources--workload--reference--group-017.md#canonical-0a63e2c8b641721f1b34e55cf2ff86a96e0ff98f533b2effad0740e534e1930c)
- [simple_service.simple_advertise](resources--workload--reference--group-017.md#canonical-f4cb6c4627ef2f328600ec5bc5b3170189481279b9545f5640eb0f1fe61e5b22)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-3898acb911cd1464ffad4a15e5db50b80e5c9c97e8ef46fc21e2955127008d77"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c3b8d01e5644625a95c6f150b3c37fbd095414b47610f5b4cccd11d7dc56c274"></a>

## simple_service.configuration — simple_service.configuration / e617ba7a1c14 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [simple_service](resources--workload--reference--group-016.md#canonical-9361b00a2d9b5209c533e92d6cb37dd367b4eb50395ea573285491ecbf8e6419)
- simple_service.configuration

<a id="canonical-6ebbd412a5c6c9c6ba235600c1fb1bd338f1cb58defd2b8c7f9b40bfe9f8fad9"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
configuration {
  # Configure direct properties listed below.
}
```

<a id="canonical-011f6dbee29ba2071cf2c0540a8707198bd360cc7675e2a3182852541430d7d7"></a>

## Direct properties — simple_service.configuration / e617ba7a1c14 / 3

- [parameters](resources--workload--reference--group-016.md#canonical-8485377c46c848c03cb49a3071902aae6a2dbf227abef534b54f195a629b2b9c): complete subsection reference.

<a id="canonical-e938f05df8ede00c7637287c279adf60027f6285ffb6b08c99f967f8ec1679cf"></a>

## Next pages — simple_service.configuration / e617ba7a1c14 / 4

- [simple_service.configuration.parameters](resources--workload--reference--group-016.md#canonical-8485377c46c848c03cb49a3071902aae6a2dbf227abef534b54f195a629b2b9c)
- [simple_service](resources--workload--reference--group-016.md#canonical-9361b00a2d9b5209c533e92d6cb37dd367b4eb50395ea573285491ecbf8e6419)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-8485377c46c848c03cb49a3071902aae6a2dbf227abef534b54f195a629b2b9c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-198a221b086dedda9a551dd220a65af396b47b5d5f80956dd65a86c551fc877d"></a>

## simple_service.configuration.parameters — simple_service.configuration.parameters / c95954c6308b / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [simple_service](resources--workload--reference--group-016.md#canonical-9361b00a2d9b5209c533e92d6cb37dd367b4eb50395ea573285491ecbf8e6419)
- [simple_service.configuration](resources--workload--reference--group-016.md#canonical-3898acb911cd1464ffad4a15e5db50b80e5c9c97e8ef46fc21e2955127008d77)
- simple_service.configuration.parameters

<a id="canonical-cb0c4751d79f3863b4890e2fc3b4d2d6cc489f46f30c2a23313b6b4685af64e2"></a>

Type: `"object"`. list nested block, Optional.

Parameters. Parameters for the workload.

Upstream description:

Parameters for the workload.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("env_var",
    "file")}
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
parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-48766fef797b94181e7916987aca8bbb2d4c1a675f498f6dd10221860e6f679b"></a>

## Direct properties — simple_service.configuration.parameters / c95954c6308b / 3

- [env_var](resources--workload--reference--group-016.md#canonical-126b6dc129f6b0d49f521e0c06f5092c205def8bd986c88844d17ee4d3a10d5a): complete subsection reference.

- [file](resources--workload--reference--group-016.md#canonical-eebfcd67f689c5a2a2f7a3a7db8640fc943f494313b377f7c611b16c8a2ddbd2): complete subsection reference.

<a id="canonical-6bf63b34dcb7f90de5fa23ab18a83aa2342c5748fe07cc77af5584a7560db442"></a>

## Next pages — simple_service.configuration.parameters / c95954c6308b / 4

- [simple_service.configuration.parameters.env_var](resources--workload--reference--group-016.md#canonical-126b6dc129f6b0d49f521e0c06f5092c205def8bd986c88844d17ee4d3a10d5a)
- [simple_service.configuration.parameters.file](resources--workload--reference--group-016.md#canonical-eebfcd67f689c5a2a2f7a3a7db8640fc943f494313b377f7c611b16c8a2ddbd2)
- [simple_service.configuration](resources--workload--reference--group-016.md#canonical-3898acb911cd1464ffad4a15e5db50b80e5c9c97e8ef46fc21e2955127008d77)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-126b6dc129f6b0d49f521e0c06f5092c205def8bd986c88844d17ee4d3a10d5a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f761958dc35cd507b88715206cc8ca257d6d03eb0f9fffb18b53a4518bf2d4e7"></a>

## simple_service.configuration.parameters.env_var — simple_service.configuration.parameters.env_var / dcc5447cb912 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [simple_service](resources--workload--reference--group-016.md#canonical-9361b00a2d9b5209c533e92d6cb37dd367b4eb50395ea573285491ecbf8e6419)
- [simple_service.configuration](resources--workload--reference--group-016.md#canonical-3898acb911cd1464ffad4a15e5db50b80e5c9c97e8ef46fc21e2955127008d77)
- [simple_service.configuration.parameters](resources--workload--reference--group-016.md#canonical-8485377c46c848c03cb49a3071902aae6a2dbf227abef534b54f195a629b2b9c)
- simple_service.configuration.parameters.env_var

<a id="canonical-beb83b4e103caaf0e67811fe389009a47556d93499c3cbaa33e6a34115f7fdc5"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
env_var {
  # Configure direct properties listed below.
}
```

<a id="canonical-7b47ed52e1b2fd5a88d2891ae0af412b0b12ab876b3da6728f20b595a933e15b"></a>

## Direct properties — simple_service.configuration.parameters.env_var / dcc5447cb912 / 3

<a id="canonical-183cabb386cd5323204bcd8bdaf2ef691d5b34083afc1aeba7a5040d93174125"></a>

<a id="canonical-32163c1d65a43470c991c190e4ccc7caae1d156713575c236f3b183bc3fef71c"></a>

## name property — simple_service.configuration.parameters.env_var / dcc5447cb912 / 4

Type: `"string"`. Optional.

Name. Name of Environment Variable.

Upstream description:

Name of Environment Variable.

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

<a id="canonical-da0884b43722c232943be2068b07273fd09d8d0e98ffa7aeab2ecb8b30e1a61a"></a>

<a id="canonical-68fe776c6b07b9ba4a5dc5fd1db4dc4d22a26d98e2b9a8ac6c025cd09fbd345c"></a>

## value property — simple_service.configuration.parameters.env_var / dcc5447cb912 / 5

Type: `"string"`. Optional.

Value. Value of Environment Variable.

Upstream description:

Value of Environment Variable.

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

<a id="canonical-624e2db7685029cdb984f72e17498630705aab80c63976400672cdce82889417"></a>

## Next pages — simple_service.configuration.parameters.env_var / dcc5447cb912 / 6

- [simple_service.configuration.parameters](resources--workload--reference--group-016.md#canonical-8485377c46c848c03cb49a3071902aae6a2dbf227abef534b54f195a629b2b9c)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-eebfcd67f689c5a2a2f7a3a7db8640fc943f494313b377f7c611b16c8a2ddbd2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c0d7ea3d01acdd1736a95b2a6ad50d4674813487a1f024a4d29357dcf1970abd"></a>

## simple_service.configuration.parameters.file — simple_service.configuration.parameters.file / dcaa8ac7c7e3 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [simple_service](resources--workload--reference--group-016.md#canonical-9361b00a2d9b5209c533e92d6cb37dd367b4eb50395ea573285491ecbf8e6419)
- [simple_service.configuration](resources--workload--reference--group-016.md#canonical-3898acb911cd1464ffad4a15e5db50b80e5c9c97e8ef46fc21e2955127008d77)
- [simple_service.configuration.parameters](resources--workload--reference--group-016.md#canonical-8485377c46c848c03cb49a3071902aae6a2dbf227abef534b54f195a629b2b9c)
- simple_service.configuration.parameters.file

<a id="canonical-413b2f9a3441a7ec2642ba83c4b1cdd9413c05de6d23e48ca85537f0a49c7c30"></a>

Type: `"object"`. single nested block, Optional.

Configuration File. Configuration File for the workload.

Upstream description:

Configuration File for the workload.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name",
    "volume_name")}
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
file {
  # Configure direct properties listed below.
}
```

<a id="canonical-4bb99aca7c0c480ff94b5991beeb230dbb2fbee59957bc551e449dbf51ef7296"></a>

## Direct properties — simple_service.configuration.parameters.file / dcaa8ac7c7e3 / 3

<a id="canonical-c490f73736ca61ad8d0b653a38805c1c74bf1aebdf3d3eb14d3ed07183b780d9"></a>

<a id="canonical-1f022b35872980773bc2c6cbf922d5e9f2e875b3468092f4d420c99809f5ea9a"></a>

## data property — simple_service.configuration.parameters.file / dcaa8ac7c7e3 / 4

Type: `"string"`. Optional.

Data. File data

Upstream description:

File data

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(16384),
}
```

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

- [mount](resources--workload--reference--group-016.md#canonical-91246abdee11bb8ef34f55c48cea538613094fd28d05ca7ced124d422b84aad2): complete subsection reference.

<a id="canonical-27d73b79fef6749e7106604a7db59ae86eec9514a76d2e8a68c42ce62253a2fb"></a>

<a id="canonical-e9cebdaad6a54e5528d1ab92653f9d8a7e8af960b4f083fdbe55ec638466eee7"></a>

## name property — simple_service.configuration.parameters.file / dcaa8ac7c7e3 / 5

Type: `"string"`. Optional.

Name. Name of the file.

Upstream description:

Name of the file.

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

<a id="canonical-484a29750a022f2dc6172cea802e918e405f5049c3e811187ca8fcf5c40706e4"></a>

<a id="canonical-57460ec71bdb937f0f4667be0ff3da0d50b7606533e9ee2e4843fe923f2b4603"></a>

## volume_name property — simple_service.configuration.parameters.file / dcaa8ac7c7e3 / 6

Type: `"string"`. Optional.

Volume Name. Name of the Volume.

Upstream description:

Name of the Volume.

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

<a id="canonical-898a6e3a8de3b63168a4acc7485b231ab30a3e763774acf70cc2a3a8c8660bbf"></a>

## Next pages — simple_service.configuration.parameters.file / dcaa8ac7c7e3 / 7

- [simple_service.configuration.parameters.file.mount](resources--workload--reference--group-016.md#canonical-91246abdee11bb8ef34f55c48cea538613094fd28d05ca7ced124d422b84aad2)
- [simple_service.configuration.parameters](resources--workload--reference--group-016.md#canonical-8485377c46c848c03cb49a3071902aae6a2dbf227abef534b54f195a629b2b9c)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-91246abdee11bb8ef34f55c48cea538613094fd28d05ca7ced124d422b84aad2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9ce3c47caa7e893ef863fa0f55c3ba94fa98721e3dd8866b202d724e70b7b8c8"></a>

## simple_service.configuration.parameters.file.mount — simple_service.configuration.parameters.file.mount / 1a7feb8eacd3 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [simple_service](resources--workload--reference--group-016.md#canonical-9361b00a2d9b5209c533e92d6cb37dd367b4eb50395ea573285491ecbf8e6419)
- [simple_service.configuration](resources--workload--reference--group-016.md#canonical-3898acb911cd1464ffad4a15e5db50b80e5c9c97e8ef46fc21e2955127008d77)
- [simple_service.configuration.parameters](resources--workload--reference--group-016.md#canonical-8485377c46c848c03cb49a3071902aae6a2dbf227abef534b54f195a629b2b9c)
- [simple_service.configuration.parameters.file](resources--workload--reference--group-016.md#canonical-eebfcd67f689c5a2a2f7a3a7db8640fc943f494313b377f7c611b16c8a2ddbd2)
- simple_service.configuration.parameters.file.mount

<a id="canonical-e6c8af7ba5a61a831b59c2370965dce7c2ad5906f9395224483184d3559db825"></a>

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

<a id="canonical-f18108e995d8c20048bc949c18b005cf28069565827582a6889c2b3dd7139d87"></a>

## Direct properties — simple_service.configuration.parameters.file.mount / 1a7feb8eacd3 / 3

<a id="canonical-031a63c2139e43513983d211d790b7fc32bd7384ad8e60b206076524ed3890a0"></a>

<a id="canonical-b713254c9110c2734d65f2a82be607ca8e62ea37e4e457f1a783bf37c2fee6e5"></a>

## mode property — simple_service.configuration.parameters.file.mount / 1a7feb8eacd3 / 4

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

<a id="canonical-e6f2bf2a7f45f8f67b66f50bec4d1861419c57e50c3be72de67ea26cfb9d685d"></a>

<a id="canonical-98d57bd237d4f609fd9c3b47bf8219f1499066e7c95bb72dfdf3b5d218fee014"></a>

## mount_path property — simple_service.configuration.parameters.file.mount / 1a7feb8eacd3 / 5

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

<a id="canonical-375cfe051a27f0a509c47a24d6b19ceedb306239e39df25c1f03435c8ba2701a"></a>

<a id="canonical-0189dbc1c2ba464389917fca228168e1b1e69a9fbbb18ed870c9cf8e8a49effb"></a>

## sub_path property — simple_service.configuration.parameters.file.mount / 1a7feb8eacd3 / 6

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

<a id="canonical-d823e491cb82a2e0aca4360c3d72b9af9efde661d8d38375ffd33ffc309c9927"></a>

## Next pages — simple_service.configuration.parameters.file.mount / 1a7feb8eacd3 / 7

- [simple_service.configuration.parameters.file](resources--workload--reference--group-016.md#canonical-eebfcd67f689c5a2a2f7a3a7db8640fc943f494313b377f7c611b16c8a2ddbd2)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-45cc02935ab54e5ce267d41971e17998e8f55556582a81b7f71e3e643ef36327"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9172ebbdb01ea57dc74b9e476b578b936f88fce21644ea3facdbe2f592252988"></a>

## simple_service.container — simple_service.container / bfee4abcf144 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [simple_service](resources--workload--reference--group-016.md#canonical-9361b00a2d9b5209c533e92d6cb37dd367b4eb50395ea573285491ecbf8e6419)
- simple_service.container

<a id="canonical-00f9c01affe0eb773a93e5ed03c936ea7c46356fddec5eabbc0018134d3e0a27"></a>

Type: `"object"`. single nested block, Optional.

ContainerType configures the container information.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name"),
  validators.ConflictingObjectAttributes("custom_flavor",
    "default_flavor"),
  validators.ConflictingObjectAttributes("custom_flavor",
    "flavor"),
  validators.ConflictingObjectAttributes("default_flavor",
    "flavor")}
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
  "x-ves-oneof-field-flavor_choice": "[\"custom_flavor\",\"default_flavor\",\"flavor\"]"
}
```

Terraform syntax:

```terraform
container {
  # Configure direct properties listed below.
}
```

<a id="canonical-a68b16f26a202c68abe65bd30e9d0c219b26e965cae360a3cccfcf7f81ac4ff2"></a>

## Direct properties — simple_service.container / bfee4abcf144 / 3

<a id="canonical-7869280600f9b8d290f226827d7ec0287bb637985754ee88f0f91e2f14082f7e"></a>

<a id="canonical-e14925859de30b345c4ba13fe2d3b08c41d1cdbb135c2022430c7273596cae73"></a>

## args property — simple_service.container / bfee4abcf144 / 4

Type: `["list", "string"]`. Optional.

Arguments to the entrypoint. Overrides the docker image's CMD.

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

<a id="canonical-6b5710d682a35fd5c2b7cc5ae52869134e930d81ba5bb2585b1157252887419f"></a>

<a id="canonical-c8defd6ec33aec8067987f39546f455ec9c0a3c4da8535d15d2b486e52ac10a6"></a>

## command property — simple_service.container / bfee4abcf144 / 5

Type: `["list", "string"]`. Optional.

Command to execute. Overrides the docker image's ENTRYPOINT.

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

- [custom_flavor](resources--workload--reference--group-016.md#canonical-51681123272957d68ae19cbd24fb7a155a10cc4f6798ca20b5574db1a2912e75): complete subsection reference.

- [default_flavor](resources--workload--reference--group-016.md#canonical-49eadf8a20688cec53657bc3336728ad173556a4ca6edb79fd5df18061bcda1c): complete subsection reference.

<a id="canonical-2e7d86c1c00501d3f9cfb83457ca7a12229bb3c776ddd956ed35f7a27f6f1f88"></a>

<a id="canonical-bd6569d0a8ea4479c491d48ced4420f2de96d6e9f8fa3a5574f1bdfa689f86b9"></a>

## flavor property — simple_service.container / bfee4abcf144 / 6

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("CONTAINER_FLAVOR_TYPE_TINY",
    "CONTAINER_FLAVOR_TYPE_MEDIUM",
    "CONTAINER_FLAVOR_TYPE_LARGE"),
}
```

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

- [image](resources--workload--reference--group-016.md#canonical-4de20ea9c2692c5e2993daf744534b6667d662cc2ed3a62f35ba5b5fe8c73dc7): complete subsection reference.

<a id="canonical-a21e2da5e9904d5b787138a1683a1692c88689a80b81fb053b53b7f66629a47c"></a>

<a id="canonical-2800726609451dcfae984618ddaa2bd119b45b5a40654f8297d207867c9bf694"></a>

## init_container property — simple_service.container / bfee4abcf144 / 7

Type: `"bool"`. Optional.

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

- [liveness_check](resources--workload--reference--group-016.md#canonical-5b0e1ac8cd0430ccbb717b0994598b4c3d4597a08492e9d9cbda3ba2563ebe56): complete subsection reference.

<a id="canonical-8c45fac7bc0c23f94be6e9eb3c74f8a220d15d76cd57814b1d9920dc0e3dd1cf"></a>

<a id="canonical-643e971d06060ab6a4dba72f2ff9497efc1fec1cda5b140a289353331f4f8ba8"></a>

## name property — simple_service.container / bfee4abcf144 / 8

Type: `"string"`. Optional.

Name. Name of the container.

Upstream description:

Name of the container.

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

- [readiness_check](resources--workload--reference--group-016.md#canonical-e664c46fdda4fd0b6f3fc712c7f3e8e7cf0c53ba7b845568a2edbfb7e8c2207b): complete subsection reference.

<a id="canonical-cdfad8126dfbd19eb32455b966e2e826be005b13a543c21da5985b03252f586e"></a>

## Next pages — simple_service.container / bfee4abcf144 / 9

- [simple_service.container.custom_flavor](resources--workload--reference--group-016.md#canonical-51681123272957d68ae19cbd24fb7a155a10cc4f6798ca20b5574db1a2912e75)
- [simple_service.container.default_flavor](resources--workload--reference--group-016.md#canonical-49eadf8a20688cec53657bc3336728ad173556a4ca6edb79fd5df18061bcda1c)
- [simple_service.container.image](resources--workload--reference--group-016.md#canonical-4de20ea9c2692c5e2993daf744534b6667d662cc2ed3a62f35ba5b5fe8c73dc7)
- [simple_service.container.liveness_check](resources--workload--reference--group-016.md#canonical-5b0e1ac8cd0430ccbb717b0994598b4c3d4597a08492e9d9cbda3ba2563ebe56)
- [simple_service.container.readiness_check](resources--workload--reference--group-016.md#canonical-e664c46fdda4fd0b6f3fc712c7f3e8e7cf0c53ba7b845568a2edbfb7e8c2207b)
- [simple_service](resources--workload--reference--group-016.md#canonical-9361b00a2d9b5209c533e92d6cb37dd367b4eb50395ea573285491ecbf8e6419)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-51681123272957d68ae19cbd24fb7a155a10cc4f6798ca20b5574db1a2912e75"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f4ac307908b3d2be28fd19e9a84880b17acb6ad290b7913d7c25964ddea44c14"></a>

## simple_service.container.custom_flavor — simple_service.container.custom_flavor / 3dbc40f1a196 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [simple_service](resources--workload--reference--group-016.md#canonical-9361b00a2d9b5209c533e92d6cb37dd367b4eb50395ea573285491ecbf8e6419)
- [simple_service.container](resources--workload--reference--group-016.md#canonical-45cc02935ab54e5ce267d41971e17998e8f55556582a81b7f71e3e643ef36327)
- simple_service.container.custom_flavor

<a id="canonical-2df80efd2e0590ec2744dbdd86302ed74c2d991914e3e92da631fcd4c7d08b97"></a>

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
custom_flavor {
  # Configure direct properties listed below.
}
```

<a id="canonical-a848640da9e246ab806e64d84075c05fec3a9a2b9d441b6a7cf5872bd264e0e1"></a>

## Direct properties — simple_service.container.custom_flavor / 3dbc40f1a196 / 3

<a id="canonical-f2ea2dddad259138680e907ec4dd6ea981c9b55e8cb79a6a77f1e89f7025d8f9"></a>

<a id="canonical-e06f7a36bd15edecd5271708c28026ad9680ddf63c317084c62c3c4d040a9d7c"></a>

## name property — simple_service.container.custom_flavor / 3dbc40f1a196 / 4

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

<a id="canonical-91583cac306ec21355ebb4f2e36c752cd2cfaf81228ef16e844321f10c47b610"></a>

<a id="canonical-215a260b4d30fa2d3ecca0795fa805c94b46d8b60fca58da9756630395f6ca55"></a>

## namespace property — simple_service.container.custom_flavor / 3dbc40f1a196 / 5

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

<a id="canonical-08a1b10ce1d63815e05cf6f70585bd833b9b6b40113050ee870793461b2686ee"></a>

<a id="canonical-4e0c45d7fb69a1b64a599a8f5716976f564201309485db1e01a7ddaeaa3e8366"></a>

## tenant property — simple_service.container.custom_flavor / 3dbc40f1a196 / 6

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

<a id="canonical-9f42cfc43218348497c5217ac68d4a150b252033658ebe24d34221d0b5f0fd44"></a>

## Next pages — simple_service.container.custom_flavor / 3dbc40f1a196 / 7

- [simple_service.container](resources--workload--reference--group-016.md#canonical-45cc02935ab54e5ce267d41971e17998e8f55556582a81b7f71e3e643ef36327)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-49eadf8a20688cec53657bc3336728ad173556a4ca6edb79fd5df18061bcda1c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0c035fdae906fc8d1a1eaba26f94507ccfa41d78e9c9b664323aa969b627ba19"></a>

## simple_service.container.default_flavor — simple_service.container.default_flavor / f255dc656cf6 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [simple_service](resources--workload--reference--group-016.md#canonical-9361b00a2d9b5209c533e92d6cb37dd367b4eb50395ea573285491ecbf8e6419)
- [simple_service.container](resources--workload--reference--group-016.md#canonical-45cc02935ab54e5ce267d41971e17998e8f55556582a81b7f71e3e643ef36327)
- simple_service.container.default_flavor

<a id="canonical-c978015bd7aa9e4f59e70d7228f7ebf19b1079847f9c9aa9d850fd32c279271c"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default_flavor = {}
```

<a id="canonical-c2b8de31b2411250e97cbb0a2c8f339ed849677dacf13e0cb26d2b431c74d798"></a>

## Direct properties — simple_service.container.default_flavor / f255dc656cf6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b8cc89ac6eb5f0a435aebb2852da373f727b386bfcf916213103cadf8a5c966d"></a>

## Next pages — simple_service.container.default_flavor / f255dc656cf6 / 4

- [simple_service.container](resources--workload--reference--group-016.md#canonical-45cc02935ab54e5ce267d41971e17998e8f55556582a81b7f71e3e643ef36327)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-4de20ea9c2692c5e2993daf744534b6667d662cc2ed3a62f35ba5b5fe8c73dc7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ef467757198cbe035db639e80265f68080de75b5937470b431af99d113db7cc0"></a>

## simple_service.container.image — simple_service.container.image / 91f3e15da487 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [simple_service](resources--workload--reference--group-016.md#canonical-9361b00a2d9b5209c533e92d6cb37dd367b4eb50395ea573285491ecbf8e6419)
- [simple_service.container](resources--workload--reference--group-016.md#canonical-45cc02935ab54e5ce267d41971e17998e8f55556582a81b7f71e3e643ef36327)
- simple_service.container.image

<a id="canonical-56fa976189d9f107c364c29091c2d9eb948ca5bc31112aa19717ac98482b551c"></a>

Type: `"object"`. single nested block, Optional.

ImageType configures the image to use, how to pull the image, and the associated secrets to use if
any.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name"),
  validators.ConflictingObjectAttributes("container_registry",
    "public")}
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
  "x-ves-oneof-field-registry_choice": "[\"container_registry\",\"public\"]"
}
```

Terraform syntax:

```terraform
image {
  # Configure direct properties listed below.
}
```

<a id="canonical-1b6f8fbf29492968576f9791105905fa740ac42fbb70507bd64e943d03341b2a"></a>

## Direct properties — simple_service.container.image / 91f3e15da487 / 3

- [container_registry](resources--workload--reference--group-016.md#canonical-60f567edf40c46c6542c4ce5b51b4a19c96ad18065447f8a43ae89b591e7da89): complete subsection reference.

<a id="canonical-f6c47ffa850f8ca1a27998d54a95cea1a78111b110e9d223c980b3c18dcbb9ae"></a>

<a id="canonical-30ad6638e12d5b6b1705e30bf1ae45fec49a46f5e3790e24c7136e71586ced28"></a>

## name property — simple_service.container.image / 91f3e15da487 / 4

Type: `"string"`. Optional.

Name is a container image which are usually given a name such as alpine, ubuntu, or
quay.I/O/etcd:0.13. The format is registry/image:tag or registry/image@image-digest. If registry is
not specified, the Docker public registry is assumed.

Upstream description:

Name is a container image which are usually given a name such as alpine, ubuntu, or
quay.I/O/etcd:0.13. The format is registry/image:tag or registry/image@image-digest. If registry is
not specified, the Docker public registry is assumed. If tag is not specified, latest is assumed.

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

- [public](resources--workload--reference--group-016.md#canonical-c54e042ee9864633be5abc21c8548165775ef212fc0e6650508855915d34bb4a): complete subsection reference.

<a id="canonical-3688246933b4d972d59ef6be2af7b564b8ba8100cc03b895afce96f6ee36143c"></a>

<a id="canonical-cc5f264c560ef112222b0bf7b834c88dc779eb77db561157047d01944d3894b2"></a>

## pull_policy property — simple_service.container.image / 91f3e15da487 / 5

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("IMAGE_PULL_POLICY_DEFAULT",
    "IMAGE_PULL_POLICY_IF_NOT_PRESENT",
    "IMAGE_PULL_POLICY_ALWAYS",
    "IMAGE_PULL_POLICY_NEVER"),
}
```

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

<a id="canonical-edfb84c30440c4290f96b572c0b03fc4b8ec0760bfbcfdb51b69f5065f85bac0"></a>

## Next pages — simple_service.container.image / 91f3e15da487 / 6

- [simple_service.container.image.container_registry](resources--workload--reference--group-016.md#canonical-60f567edf40c46c6542c4ce5b51b4a19c96ad18065447f8a43ae89b591e7da89)
- [simple_service.container.image.public](resources--workload--reference--group-016.md#canonical-c54e042ee9864633be5abc21c8548165775ef212fc0e6650508855915d34bb4a)
- [simple_service.container](resources--workload--reference--group-016.md#canonical-45cc02935ab54e5ce267d41971e17998e8f55556582a81b7f71e3e643ef36327)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-60f567edf40c46c6542c4ce5b51b4a19c96ad18065447f8a43ae89b591e7da89"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-681f83f935c14771d7ddfab40ca51ef7cb7f5328eee84935dcca4395c513a608"></a>

## simple_service.container.image.container_registry — simple_service.container.image.container_registry / 554cb00c423c / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [simple_service](resources--workload--reference--group-016.md#canonical-9361b00a2d9b5209c533e92d6cb37dd367b4eb50395ea573285491ecbf8e6419)
- [simple_service.container](resources--workload--reference--group-016.md#canonical-45cc02935ab54e5ce267d41971e17998e8f55556582a81b7f71e3e643ef36327)
- [simple_service.container.image](resources--workload--reference--group-016.md#canonical-4de20ea9c2692c5e2993daf744534b6667d662cc2ed3a62f35ba5b5fe8c73dc7)
- simple_service.container.image.container_registry

<a id="canonical-68160318b171766cccef5f2490f5fcfb5b546d4200c6d0644fafed84e24cb7e1"></a>

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
container_registry {
  # Configure direct properties listed below.
}
```

<a id="canonical-b9a20985883b310afc297d7731f06be5bcf1ed9c6462f63c02d017df6cf29b51"></a>

## Direct properties — simple_service.container.image.container_registry / 554cb00c423c / 3

<a id="canonical-b69fcd5d9b3a8cc04266b1cf5c895c075c49457312dabfedaa5357e954fcf16f"></a>

<a id="canonical-3a8de96ef4d7a35c338268c661bce45b88112049da4e8e3d511fd3224bd8285e"></a>

## name property — simple_service.container.image.container_registry / 554cb00c423c / 4

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

<a id="canonical-1eaf022f949942f90d291837dd83d57ce2e2e6089958e98ad4f70725a490bbd0"></a>

<a id="canonical-30b291c646d3e193b217b4b486953e2b1a9248f8539ff231646499a50422451b"></a>

## namespace property — simple_service.container.image.container_registry / 554cb00c423c / 5

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

<a id="canonical-3cf2a13cc73a10ea3b921af71c3dfc3356e3c4225b5a657c49cd5a89e89ef8de"></a>

<a id="canonical-6bcb73b9b1c657575cd70259105fe8b51223d5bfb04cf8e1351981d02515fb3a"></a>

## tenant property — simple_service.container.image.container_registry / 554cb00c423c / 6

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

<a id="canonical-0bc29ce2c624df9e3ed1a4a2b96d70ee561ca4ab56cbedb01e77c315406011c8"></a>

## Next pages — simple_service.container.image.container_registry / 554cb00c423c / 7

- [simple_service.container.image](resources--workload--reference--group-016.md#canonical-4de20ea9c2692c5e2993daf744534b6667d662cc2ed3a62f35ba5b5fe8c73dc7)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-c54e042ee9864633be5abc21c8548165775ef212fc0e6650508855915d34bb4a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-42e5b1736ec1046c15a9dd0b739366eef86d72dfe0c59def468696e0ea4caceb"></a>

## simple_service.container.image.public — simple_service.container.image.public / 2b24dc6f4798 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [simple_service](resources--workload--reference--group-016.md#canonical-9361b00a2d9b5209c533e92d6cb37dd367b4eb50395ea573285491ecbf8e6419)
- [simple_service.container](resources--workload--reference--group-016.md#canonical-45cc02935ab54e5ce267d41971e17998e8f55556582a81b7f71e3e643ef36327)
- [simple_service.container.image](resources--workload--reference--group-016.md#canonical-4de20ea9c2692c5e2993daf744534b6667d662cc2ed3a62f35ba5b5fe8c73dc7)
- simple_service.container.image.public

<a id="canonical-57af790578a4d8868cc4f64d8807833ad6058cfabe57d14cc14e23cbd2b7f915"></a>

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
public = {}
```

<a id="canonical-878c1f56f6a03c9253f4b9fb8b9da7461686cf4166d32d1b874315c42f688fd2"></a>

## Direct properties — simple_service.container.image.public / 2b24dc6f4798 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-90b5d84064333c999a8b0de1fe6b861929d8bcfe33790c8fa37c88a1af4e9ff9"></a>

## Next pages — simple_service.container.image.public / 2b24dc6f4798 / 4

- [simple_service.container.image](resources--workload--reference--group-016.md#canonical-4de20ea9c2692c5e2993daf744534b6667d662cc2ed3a62f35ba5b5fe8c73dc7)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-5b0e1ac8cd0430ccbb717b0994598b4c3d4597a08492e9d9cbda3ba2563ebe56"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8865a1826667571f85a1845b3b7589c4ec982ff2e8f0c447f685effdb74238dd"></a>

## simple_service.container.liveness_check — simple_service.container.liveness_check / 88c6b593f035 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [simple_service](resources--workload--reference--group-016.md#canonical-9361b00a2d9b5209c533e92d6cb37dd367b4eb50395ea573285491ecbf8e6419)
- [simple_service.container](resources--workload--reference--group-016.md#canonical-45cc02935ab54e5ce267d41971e17998e8f55556582a81b7f71e3e643ef36327)
- simple_service.container.liveness_check

<a id="canonical-45c112acc38e8fff23325992a787dc2fc017c4f71aa811b0f76a562d5eb734e0"></a>

Type: `"object"`. single nested block, Optional.

HealthCheckType describes a health check to be performed against a container to determine whether it
has started up or is alive or ready to receive traffic.

Upstream description:

HealthCheckType describes a health check to be performed against a container to determine whether it
has started up or is alive or ready to receive traffic.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("healthy_threshold",
    "interval",
    "timeout",
    "unhealthy_threshold"),
  validators.ConflictingObjectAttributes("exec_health_check",
    "http_health_check"),
  validators.ConflictingObjectAttributes("exec_health_check",
    "tcp_health_check"),
  validators.ConflictingObjectAttributes("http_health_check",
    "tcp_health_check")}
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
  "x-ves-oneof-field-health_check_choice": "[\"exec_health_check\",\"http_health_check\",\"tcp_health_check\"]"
}
```

Terraform syntax:

```terraform
liveness_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-4a9b6eab5c6d083af3714110f76b3672b2495b7dfc0b11b1e7c2e8f9296bf616"></a>

## Direct properties — simple_service.container.liveness_check / 88c6b593f035 / 3

- [exec_health_check](resources--workload--reference--group-016.md#canonical-40be668ddb41ddf1c2642afdc7debd0d537925335251310441c024ba65777588): complete subsection reference.

<a id="canonical-64ce493a641dd84fb311aa3db9783db3f5ab0013af909a17555da290bed18c73"></a>

<a id="canonical-e1a990d522675728b41cb61fda12aa574e39a0972c4aa1b2a9c617c32ba7e112"></a>

## healthy_threshold property — simple_service.container.liveness_check / 88c6b593f035 / 4

Type: `"number"`. Optional.

Number of consecutive successful responses after having failed before declaring healthy. In other
words, this is the number of healthy health checks required before marking healthy. Note that during
startup and liveliness, only a single successful health check is required to mark a container..

Upstream description:

Number of consecutive successful responses after having failed before declaring healthy. In other
words, this is the number of healthy health checks required before marking healthy. Note that during
startup and liveliness, only a single successful health check is required to mark a container
healthy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 16),
}
```

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

- [http_health_check](resources--workload--reference--group-016.md#canonical-1b5a5885450fa4bbed852c8fb247160dd868668ee6d71af5a87dd10d038305b1): complete subsection reference.

<a id="canonical-26993e7c3e54f8dc881a43341ca40dd80c52fba14eb23f0a48384eb68d1950b6"></a>

<a id="canonical-3f1752a27c4cd741777b1acc56921cc4f7d50a00247b76a6aec817c8a85a445a"></a>

## initial_delay property — simple_service.container.liveness_check / 88c6b593f035 / 5

Type: `"number"`. Optional.

Number of seconds after the container has started before health checks are initiated.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(600),
}
```

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

<a id="canonical-648900db64d11e86adcdd817bd1d569c65860b88a5834eb63e5e8d296a56e945"></a>

<a id="canonical-75a4688b26d1a97b53fa8818215f26825968c40cd4d5b3ec2ffc565066f3ae4f"></a>

## interval property — simple_service.container.liveness_check / 88c6b593f035 / 6

Type: `"number"`. Optional.

Time interval in seconds between two health check requests.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 600),
}
```

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

- [tcp_health_check](resources--workload--reference--group-016.md#canonical-79159d31a9b54d91a37b1e0ccc85462dc208e68c5c402d7ddc9a96366c6abfd5): complete subsection reference.

<a id="canonical-568d965b003ed1234229efdf2a4a0b7b0f4f3752c4ad58f8232b506633d6ee7c"></a>

<a id="canonical-887fd14c05843e5594ae6d0816508fa4bec6a3f176925d8e1ce4289daf2b73dc"></a>

## timeout property — simple_service.container.liveness_check / 88c6b593f035 / 7

Type: `"number"`. Optional.

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure.

Upstream description:

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 600),
}
```

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

<a id="canonical-daa955d2eaf03cab1909983fb724ebce58a5bd2213f3d8c8f2b862fd45a0f966"></a>

<a id="canonical-ab6c6b847bdde46e11679bd84b00fc94c486a51d2bf10565de7d9fc708c1e435"></a>

## unhealthy_threshold property — simple_service.container.liveness_check / 88c6b593f035 / 8

Type: `"number"`. Optional.

Number of consecutive failed responses before declaring unhealthy. In other words, this is the
number of unhealthy health checks required before a container is marked unhealthy.

Upstream description:

Number of consecutive failed responses before declaring unhealthy. In other words, this is the
number of unhealthy health checks required before a container is marked unhealthy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 16),
}
```

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

<a id="canonical-21ae26957fa0ac83909c8da4d2041e22f12bb76ca0c2b4a6139780dae5e620d2"></a>

## Next pages — simple_service.container.liveness_check / 88c6b593f035 / 9

- [simple_service.container.liveness_check.exec_health_check](resources--workload--reference--group-016.md#canonical-40be668ddb41ddf1c2642afdc7debd0d537925335251310441c024ba65777588)
- [simple_service.container.liveness_check.http_health_check](resources--workload--reference--group-016.md#canonical-1b5a5885450fa4bbed852c8fb247160dd868668ee6d71af5a87dd10d038305b1)
- [simple_service.container.liveness_check.tcp_health_check](resources--workload--reference--group-016.md#canonical-79159d31a9b54d91a37b1e0ccc85462dc208e68c5c402d7ddc9a96366c6abfd5)
- [simple_service.container](resources--workload--reference--group-016.md#canonical-45cc02935ab54e5ce267d41971e17998e8f55556582a81b7f71e3e643ef36327)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-40be668ddb41ddf1c2642afdc7debd0d537925335251310441c024ba65777588"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-abeb9e12986392c574ec7313cd8263a55f8356621796314fbecbfa3f4f20adc5"></a>

## simple_service.container.liveness_check.exec_health_check — simple_service.container.liveness_check.exec_health_check / 15d5cb275f2b / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [simple_service](resources--workload--reference--group-016.md#canonical-9361b00a2d9b5209c533e92d6cb37dd367b4eb50395ea573285491ecbf8e6419)
- [simple_service.container](resources--workload--reference--group-016.md#canonical-45cc02935ab54e5ce267d41971e17998e8f55556582a81b7f71e3e643ef36327)
- [simple_service.container.liveness_check](resources--workload--reference--group-016.md#canonical-5b0e1ac8cd0430ccbb717b0994598b4c3d4597a08492e9d9cbda3ba2563ebe56)
- simple_service.container.liveness_check.exec_health_check

<a id="canonical-305445df30929a7145894799312dfbfbf27bba4d2bc68570a7ce5a614b8313fc"></a>

Type: `"object"`. single nested block, Optional.

ExecHealthCheckType describes a health check based on 'run in container' action. Exit status of 0 is
treated as live/healthy and non-zero is unhealthy.

Upstream description:

ExecHealthCheckType describes a health check based on "run in container" action. Exit status of 0 is
treated as live/healthy and non-zero is unhealthy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("command")}
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
exec_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-ec9d2a1adb707b733c77d15b4891c2488c940bb220ccf622eecb74ac32f89cc7"></a>

## Direct properties — simple_service.container.liveness_check.exec_health_check / 15d5cb275f2b / 3

<a id="canonical-81f1adc158fa533c29c46b163cfe383efdacd4214b4797ae60c03b297b3494c0"></a>

<a id="canonical-9ae0c6a7bceb0c57d06c8c87d439545e35d69042e8aaa9f6a0ac032cf2feb427"></a>

## command property — simple_service.container.liveness_check.exec_health_check / 15d5cb275f2b / 4

Type: `["list", "string"]`. Optional.

Command is the command line to execute inside the container, the working directory for the command
is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a
shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to..

Upstream description:

Command is the command line to execute inside the container, the working directory for the command
is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a
shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to
explicitly call out to that shell.

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

<a id="canonical-c211e0dec6f67b185dbc7f0eb3baa70b178461a3e3fa94d47ff1506ca8954494"></a>

## Next pages — simple_service.container.liveness_check.exec_health_check / 15d5cb275f2b / 5

- [simple_service.container.liveness_check](resources--workload--reference--group-016.md#canonical-5b0e1ac8cd0430ccbb717b0994598b4c3d4597a08492e9d9cbda3ba2563ebe56)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-1b5a5885450fa4bbed852c8fb247160dd868668ee6d71af5a87dd10d038305b1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-13ec196736f8d6cc474cd5088f7963de13547c529c710d7f5419c7434b922c34"></a>

## simple_service.container.liveness_check.http_health_check — simple_service.container.liveness_check.http_health_check / 57d17d188ea1 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [simple_service](resources--workload--reference--group-016.md#canonical-9361b00a2d9b5209c533e92d6cb37dd367b4eb50395ea573285491ecbf8e6419)
- [simple_service.container](resources--workload--reference--group-016.md#canonical-45cc02935ab54e5ce267d41971e17998e8f55556582a81b7f71e3e643ef36327)
- [simple_service.container.liveness_check](resources--workload--reference--group-016.md#canonical-5b0e1ac8cd0430ccbb717b0994598b4c3d4597a08492e9d9cbda3ba2563ebe56)
- simple_service.container.liveness_check.http_health_check

<a id="canonical-69860f17b6062eecc470e5c1c5dc1d073ea69400af6d9dcf053f5f7c5ec0369f"></a>

Type: `"object"`. single nested block, Optional.

HTTPHealthCheckType describes a health check based on HTTP GET requests.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("path")}
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
http_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-60ff3ca451176fe1a90fcf9bc4b887d26e06f5a1774dfd316895c894a3823c87"></a>

## Direct properties — simple_service.container.liveness_check.http_health_check / 57d17d188ea1 / 3

<a id="canonical-30f760eac9b080bba7e48f667b5bc748e10262583e2f0d438ef9c5912a9404c0"></a>

<a id="canonical-d55d97bf49d4f7fb5405169c2cdb6cc474ca64493d05d9a7f40c526f03f906fc"></a>

## headers property — simple_service.container.liveness_check.http_health_check / 57d17d188ea1 / 4

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

<a id="canonical-09580d7547b7b82f4ef4b963777e8a301c8f54b371684896fd4f56693b8785eb"></a>

<a id="canonical-ec08a4af42ee35ff5182f21f76cfa12d71c0ec3f4c92b1310cae63b5d5df9013"></a>

## host_header property — simple_service.container.liveness_check.http_health_check / 57d17d188ea1 / 5

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

<a id="canonical-e2ab0f15af377ef04c8f0505db71b67b59d9069489eeab43afb5fc057a56742d"></a>

<a id="canonical-288c9fdada8c855d731fb5cf7a6bcd6cc8571cdc4c88bd3f722df4575de84b6d"></a>

## path property — simple_service.container.liveness_check.http_health_check / 57d17d188ea1 / 6

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

- [port](resources--workload--reference--group-016.md#canonical-3e1862de67ba271361a6945db0f1155e42863201e3e697d2dbfdf35b98c9c86d): complete subsection reference.

<a id="canonical-6f176d8eded32466a464822c2c846e114321fc56adeaeafcf447e3cd6b61c49b"></a>

## Next pages — simple_service.container.liveness_check.http_health_check / 57d17d188ea1 / 7

- [simple_service.container.liveness_check.http_health_check.port](resources--workload--reference--group-016.md#canonical-3e1862de67ba271361a6945db0f1155e42863201e3e697d2dbfdf35b98c9c86d)
- [simple_service.container.liveness_check](resources--workload--reference--group-016.md#canonical-5b0e1ac8cd0430ccbb717b0994598b4c3d4597a08492e9d9cbda3ba2563ebe56)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-3e1862de67ba271361a6945db0f1155e42863201e3e697d2dbfdf35b98c9c86d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-893c5cd8d4841825ebe8db39f94579cd3eb6ff82f0b351412579bcc6aed5e84f"></a>

## simple_service.container.liveness_check.http_health_check.port — simple_service.container.liveness_check.http_health_check.port / 1c4892edc39c / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [simple_service](resources--workload--reference--group-016.md#canonical-9361b00a2d9b5209c533e92d6cb37dd367b4eb50395ea573285491ecbf8e6419)
- [simple_service.container](resources--workload--reference--group-016.md#canonical-45cc02935ab54e5ce267d41971e17998e8f55556582a81b7f71e3e643ef36327)
- [simple_service.container.liveness_check](resources--workload--reference--group-016.md#canonical-5b0e1ac8cd0430ccbb717b0994598b4c3d4597a08492e9d9cbda3ba2563ebe56)
- [simple_service.container.liveness_check.http_health_check](resources--workload--reference--group-016.md#canonical-1b5a5885450fa4bbed852c8fb247160dd868668ee6d71af5a87dd10d038305b1)
- simple_service.container.liveness_check.http_health_check.port

<a id="canonical-3dc5853a2b08bd7d0a1d84d08f6d9a25a6428fa835b6be79b865155b0336cd5e"></a>

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

<a id="canonical-f6a59ad622a22fb3dc0d67bf7449250b56e495c7d0a1396cc5fc026b2deabeb6"></a>

## Direct properties — simple_service.container.liveness_check.http_health_check.port / 1c4892edc39c / 3

<a id="canonical-0af2ecbc6d8b844d76ea6a8dacd5597d0a4a99306b2ef496ad01796f14c9b7de"></a>

<a id="canonical-0f7a49382a5479a63dc5e319b84788eabe1f79db117cc91f1f45b1912f250dd7"></a>

## name property — simple_service.container.liveness_check.http_health_check.port / 1c4892edc39c / 4

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

<a id="canonical-d1a0af07822329ecb6580716bcf920d2ec6a199ac580608fea69859e0795f128"></a>

<a id="canonical-cd19f754cc9046a76db4a7b858835d74c41e650b795f8359bcf06f00ccd43747"></a>

## num property — simple_service.container.liveness_check.http_health_check.port / 1c4892edc39c / 5

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

<a id="canonical-9bc0bd36dc1e940aa171a9526eda9c40a66cb99c5e83ae358d1c9646aa5aa8fb"></a>

## Next pages — simple_service.container.liveness_check.http_health_check.port / 1c4892edc39c / 6

- [simple_service.container.liveness_check.http_health_check](resources--workload--reference--group-016.md#canonical-1b5a5885450fa4bbed852c8fb247160dd868668ee6d71af5a87dd10d038305b1)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-79159d31a9b54d91a37b1e0ccc85462dc208e68c5c402d7ddc9a96366c6abfd5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eb10f8066104068dca324e0af6a77cd5509c87adb2c7a00ff69f9746b1ea3c19"></a>

## simple_service.container.liveness_check.tcp_health_check — simple_service.container.liveness_check.tcp_health_check / 92424be2eab4 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [simple_service](resources--workload--reference--group-016.md#canonical-9361b00a2d9b5209c533e92d6cb37dd367b4eb50395ea573285491ecbf8e6419)
- [simple_service.container](resources--workload--reference--group-016.md#canonical-45cc02935ab54e5ce267d41971e17998e8f55556582a81b7f71e3e643ef36327)
- [simple_service.container.liveness_check](resources--workload--reference--group-016.md#canonical-5b0e1ac8cd0430ccbb717b0994598b4c3d4597a08492e9d9cbda3ba2563ebe56)
- simple_service.container.liveness_check.tcp_health_check

<a id="canonical-77a1c4432252dd22181523d369f699057802ecd6d84025d6f80a176c92627c40"></a>

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

<a id="canonical-ebea6ddf7d2ba09238f8285e65e3d529781469b517a7b24588af18ee18151164"></a>

## Direct properties — simple_service.container.liveness_check.tcp_health_check / 92424be2eab4 / 3

- [port](resources--workload--reference--group-016.md#canonical-0e8bd99a6769c43f1b0bdd6109e8a6941378761818fce9c92df540d2d53719cc): complete subsection reference.

<a id="canonical-fe6f545132e0eb8a40a8f251ff3800f2e52141471dbe5ff249d37d47bf007de6"></a>

## Next pages — simple_service.container.liveness_check.tcp_health_check / 92424be2eab4 / 4

- [simple_service.container.liveness_check.tcp_health_check.port](resources--workload--reference--group-016.md#canonical-0e8bd99a6769c43f1b0bdd6109e8a6941378761818fce9c92df540d2d53719cc)
- [simple_service.container.liveness_check](resources--workload--reference--group-016.md#canonical-5b0e1ac8cd0430ccbb717b0994598b4c3d4597a08492e9d9cbda3ba2563ebe56)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-0e8bd99a6769c43f1b0bdd6109e8a6941378761818fce9c92df540d2d53719cc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e322ca5054b80980ce2c316fcd7dd58bcc5fc4380b49d4366008a22a7a2af8b6"></a>

## simple_service.container.liveness_check.tcp_health_check.port — simple_service.container.liveness_check.tcp_health_check.port / bf33da36de8a / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [simple_service](resources--workload--reference--group-016.md#canonical-9361b00a2d9b5209c533e92d6cb37dd367b4eb50395ea573285491ecbf8e6419)
- [simple_service.container](resources--workload--reference--group-016.md#canonical-45cc02935ab54e5ce267d41971e17998e8f55556582a81b7f71e3e643ef36327)
- [simple_service.container.liveness_check](resources--workload--reference--group-016.md#canonical-5b0e1ac8cd0430ccbb717b0994598b4c3d4597a08492e9d9cbda3ba2563ebe56)
- [simple_service.container.liveness_check.tcp_health_check](resources--workload--reference--group-016.md#canonical-79159d31a9b54d91a37b1e0ccc85462dc208e68c5c402d7ddc9a96366c6abfd5)
- simple_service.container.liveness_check.tcp_health_check.port

<a id="canonical-5959ed4ab1c977b9d855344a4a57d8889c847a1f43779665c4cc50d11de58bea"></a>

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

<a id="canonical-c9abcf434c568af8adfd4142e8672022e145277a537d74585728dead871d9d99"></a>

## Direct properties — simple_service.container.liveness_check.tcp_health_check.port / bf33da36de8a / 3

<a id="canonical-af12ecb547694d9ddf27ca899bae7b42fb1527ac12fd49532bb6422a8b2d3c85"></a>

<a id="canonical-9293176e9ce170311e4fe92a0f1b078ef6d9c9aee9420ab9ecfce129fff34ca7"></a>

## name property — simple_service.container.liveness_check.tcp_health_check.port / bf33da36de8a / 4

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

<a id="canonical-c74812f3fe34b3e7278df4f85cbf0a7f49fd28ce72165814b9aa9f4713fe9509"></a>

<a id="canonical-a57aeb86bd96602d3bcf68f66e81f1e3b48e5d85c2e86be9650b865dd2775ab9"></a>

## num property — simple_service.container.liveness_check.tcp_health_check.port / bf33da36de8a / 5

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

<a id="canonical-4527cd129234ab9f1a5b675c446bd922962d75c5e9ad59ea2060d578f5c9ca95"></a>

## Next pages — simple_service.container.liveness_check.tcp_health_check.port / bf33da36de8a / 6

- [simple_service.container.liveness_check.tcp_health_check](resources--workload--reference--group-016.md#canonical-79159d31a9b54d91a37b1e0ccc85462dc208e68c5c402d7ddc9a96366c6abfd5)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-e664c46fdda4fd0b6f3fc712c7f3e8e7cf0c53ba7b845568a2edbfb7e8c2207b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a2a9a456d45d7919fd2154f4eb711eab7f16960fd47938bdff15d950e0919929"></a>

## simple_service.container.readiness_check — simple_service.container.readiness_check / abffe41c8406 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [simple_service](resources--workload--reference--group-016.md#canonical-9361b00a2d9b5209c533e92d6cb37dd367b4eb50395ea573285491ecbf8e6419)
- [simple_service.container](resources--workload--reference--group-016.md#canonical-45cc02935ab54e5ce267d41971e17998e8f55556582a81b7f71e3e643ef36327)
- simple_service.container.readiness_check

<a id="canonical-76a5b24df064f3f99cb290d23c74684048a39c0003312ebb70b6f7c88caf77bd"></a>

Type: `"object"`. single nested block, Optional.

HealthCheckType describes a health check to be performed against a container to determine whether it
has started up or is alive or ready to receive traffic.

Upstream description:

HealthCheckType describes a health check to be performed against a container to determine whether it
has started up or is alive or ready to receive traffic.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("healthy_threshold",
    "interval",
    "timeout",
    "unhealthy_threshold"),
  validators.ConflictingObjectAttributes("exec_health_check",
    "http_health_check"),
  validators.ConflictingObjectAttributes("exec_health_check",
    "tcp_health_check"),
  validators.ConflictingObjectAttributes("http_health_check",
    "tcp_health_check")}
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
  "x-ves-oneof-field-health_check_choice": "[\"exec_health_check\",\"http_health_check\",\"tcp_health_check\"]"
}
```

Terraform syntax:

```terraform
readiness_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-f79cc3e5359febbe16cf743ae8f98c07a332db6538d7b9c1622a8f8f0797dba0"></a>

## Direct properties — simple_service.container.readiness_check / abffe41c8406 / 3

- [exec_health_check](resources--workload--reference--group-016.md#canonical-aecf62c3f2b96c101fd10425d4db61344e5173596649b8e62d544d02975f48bd): complete subsection reference.

<a id="canonical-7117c5a5c441f5cf55408bc4f5bdbf8081406d9e5aea9c08123255c5b8ceb79e"></a>

<a id="canonical-60cc31bae9c1e7731494129205957f2dfffe007673026a84533447ab422b0252"></a>

## healthy_threshold property — simple_service.container.readiness_check / abffe41c8406 / 4

Type: `"number"`. Optional.

Number of consecutive successful responses after having failed before declaring healthy. In other
words, this is the number of healthy health checks required before marking healthy. Note that during
startup and liveliness, only a single successful health check is required to mark a container..

Upstream description:

Number of consecutive successful responses after having failed before declaring healthy. In other
words, this is the number of healthy health checks required before marking healthy. Note that during
startup and liveliness, only a single successful health check is required to mark a container
healthy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 16),
}
```

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

- [http_health_check](resources--workload--reference--group-016.md#canonical-07cf7bcd54b17cc339f97dc79d951912b363bd34c5b6eaf3cd153e5255693e7a): complete subsection reference.

<a id="canonical-c71529ce55eda61bb3b69c804ecd186f660f2fba747236ab32fdd6c56c27c848"></a>

<a id="canonical-bc58e8f05ceff0eb1c8f243d72863aa77a92f74102406d4569c0dfa1cce24c4d"></a>

## initial_delay property — simple_service.container.readiness_check / abffe41c8406 / 5

Type: `"number"`. Optional.

Number of seconds after the container has started before health checks are initiated.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(600),
}
```

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

<a id="canonical-c7dfd9e2c3983c60aa4f81aa63cdae9ef9d509292988cb5345b1e33ebf748db8"></a>

<a id="canonical-9a8bd9474bd5f3d7c6f08d980635def1c782657b74824ee0a93cb22388458006"></a>

## interval property — simple_service.container.readiness_check / abffe41c8406 / 6

Type: `"number"`. Optional.

Time interval in seconds between two health check requests.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 600),
}
```

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

- [tcp_health_check](resources--workload--reference--group-017.md#canonical-987a30fd88ab095c6c245d9e41180a6a5df816b1c0a93d984407eb1e24dd2085): complete subsection reference.

<a id="canonical-d1e071a3972fbedd0186144dfef3b1c991159e32cf8ccfd04f586d848dcf9287"></a>

<a id="canonical-32d909a70d73bb84112d98830a6cb728972dc226c0d15f507cb764471d473b78"></a>

## timeout property — simple_service.container.readiness_check / abffe41c8406 / 7

Type: `"number"`. Optional.

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure.

Upstream description:

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 600),
}
```

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

<a id="canonical-69d4e102a9d365a5efbd8f1674b88b092b9b0f8a65f7eeea2170f55dcf88a131"></a>

<a id="canonical-956bf4efa15ff8d70aab23f7a81aa80e9ff47afcd8f2471e5862b72c84008744"></a>

## unhealthy_threshold property — simple_service.container.readiness_check / abffe41c8406 / 8

Type: `"number"`. Optional.

Number of consecutive failed responses before declaring unhealthy. In other words, this is the
number of unhealthy health checks required before a container is marked unhealthy.

Upstream description:

Number of consecutive failed responses before declaring unhealthy. In other words, this is the
number of unhealthy health checks required before a container is marked unhealthy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 16),
}
```

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

<a id="canonical-1ef4361deb2e86bf9be8ccae7610afca8353902bbc05bf9d3dd11d5280e2c2e5"></a>

## Next pages — simple_service.container.readiness_check / abffe41c8406 / 9

- [simple_service.container.readiness_check.exec_health_check](resources--workload--reference--group-016.md#canonical-aecf62c3f2b96c101fd10425d4db61344e5173596649b8e62d544d02975f48bd)
- [simple_service.container.readiness_check.http_health_check](resources--workload--reference--group-016.md#canonical-07cf7bcd54b17cc339f97dc79d951912b363bd34c5b6eaf3cd153e5255693e7a)
- [simple_service.container.readiness_check.tcp_health_check](resources--workload--reference--group-017.md#canonical-987a30fd88ab095c6c245d9e41180a6a5df816b1c0a93d984407eb1e24dd2085)
- [simple_service.container](resources--workload--reference--group-016.md#canonical-45cc02935ab54e5ce267d41971e17998e8f55556582a81b7f71e3e643ef36327)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-aecf62c3f2b96c101fd10425d4db61344e5173596649b8e62d544d02975f48bd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-af2589d9549ece9bc752644dd9b2c64c0c0edda3440a46ffd8da60c1dfe0bf94"></a>

## simple_service.container.readiness_check.exec_health_check — simple_service.container.readiness_check.exec_health_check / dc9509fcc839 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [simple_service](resources--workload--reference--group-016.md#canonical-9361b00a2d9b5209c533e92d6cb37dd367b4eb50395ea573285491ecbf8e6419)
- [simple_service.container](resources--workload--reference--group-016.md#canonical-45cc02935ab54e5ce267d41971e17998e8f55556582a81b7f71e3e643ef36327)
- [simple_service.container.readiness_check](resources--workload--reference--group-016.md#canonical-e664c46fdda4fd0b6f3fc712c7f3e8e7cf0c53ba7b845568a2edbfb7e8c2207b)
- simple_service.container.readiness_check.exec_health_check

<a id="canonical-9557cff36f4dd2be5910bccf411166e22a4cfb971f0179a20b4836e384a9587c"></a>

Type: `"object"`. single nested block, Optional.

ExecHealthCheckType describes a health check based on 'run in container' action. Exit status of 0 is
treated as live/healthy and non-zero is unhealthy.

Upstream description:

ExecHealthCheckType describes a health check based on "run in container" action. Exit status of 0 is
treated as live/healthy and non-zero is unhealthy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("command")}
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
exec_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-f6d47a3a4d6bcdbf6d091485274e5d59848365f62989b68deb8c41a0cb88840f"></a>

## Direct properties — simple_service.container.readiness_check.exec_health_check / dc9509fcc839 / 3

<a id="canonical-8da4003ed78c72e80ddb7c637040779558b7a1976d21ea1ab334add450a7e06c"></a>

<a id="canonical-681c795134253ad82e816685345733c124f0cdd7b0df83d33f77361e80e6974a"></a>

## command property — simple_service.container.readiness_check.exec_health_check / dc9509fcc839 / 4

Type: `["list", "string"]`. Optional.

Command is the command line to execute inside the container, the working directory for the command
is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a
shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to..

Upstream description:

Command is the command line to execute inside the container, the working directory for the command
is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a
shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to
explicitly call out to that shell.

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

<a id="canonical-148599fc9a2976fcf113a25c5c19559f7d3a4222969e9003928cccdf8daf9265"></a>

## Next pages — simple_service.container.readiness_check.exec_health_check / dc9509fcc839 / 5

- [simple_service.container.readiness_check](resources--workload--reference--group-016.md#canonical-e664c46fdda4fd0b6f3fc712c7f3e8e7cf0c53ba7b845568a2edbfb7e8c2207b)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-07cf7bcd54b17cc339f97dc79d951912b363bd34c5b6eaf3cd153e5255693e7a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d48e3205d1955f1492f035ac6c1589de4c9ecdff74c65aca0919c57ebccbf64c"></a>

## simple_service.container.readiness_check.http_health_check — simple_service.container.readiness_check.http_health_check / 1f3534cddc65 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [simple_service](resources--workload--reference--group-016.md#canonical-9361b00a2d9b5209c533e92d6cb37dd367b4eb50395ea573285491ecbf8e6419)
- [simple_service.container](resources--workload--reference--group-016.md#canonical-45cc02935ab54e5ce267d41971e17998e8f55556582a81b7f71e3e643ef36327)
- [simple_service.container.readiness_check](resources--workload--reference--group-016.md#canonical-e664c46fdda4fd0b6f3fc712c7f3e8e7cf0c53ba7b845568a2edbfb7e8c2207b)
- simple_service.container.readiness_check.http_health_check

<a id="canonical-d0a1cfd49468145fc48dff8f7b1e3e1b7ca9c8f01bb15b356e904fdd48bfb44b"></a>

Type: `"object"`. single nested block, Optional.

HTTPHealthCheckType describes a health check based on HTTP GET requests.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("path")}
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
http_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-5c14320cca20207b914a75bf48e70c1adc9e0bcae717c60e57c7ca8b63553267"></a>

## Direct properties — simple_service.container.readiness_check.http_health_check / 1f3534cddc65 / 3

<a id="canonical-5655d13255f08b08a15efa7e1dd329ad9e0442cc6f0229a883ffc999e4024c20"></a>
