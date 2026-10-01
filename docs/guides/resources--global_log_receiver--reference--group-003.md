---
page_title: "xcsh_global_log_receiver reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_global_log_receiver reference."
---

# xcsh_global_log_receiver reference

<a id="canonical-75bd7c2ad60e901b0288d497f9b4629175396a73e4f61f1a8f561d9f3fc7ffc3"></a>

## provider_ref property — http_receiver.auth_token.token.clear_secret_info / 7bc0fe7bbefb / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-62aa7c3f33f7b45fe8621ecb4a731227c194b331b592b672fbfa9877cf86139b"></a>

<a id="canonical-0211cdc1bdc99dc7eec7bfa22c62be69bdca1f28fa2fe6ffa4e6078ebf0123a0"></a>

## url property — http_receiver.auth_token.token.clear_secret_info / 7bc0fe7bbefb / 5

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

<a id="canonical-57ec3424219fd34459af19895604e7dbee49fc010bc5dc9631e8a190ca807a19"></a>

## Next pages — http_receiver.auth_token.token.clear_secret_info / 7bc0fe7bbefb / 6

- [http_receiver.auth_token.token](resources--global_log_receiver--reference--group-002.md#canonical-fd7447a2fd26b5480da7e9bfe2820536acfafe4bdfc82b1852e4585cfe927af3)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-12d05cef66a26cbbf70082c850e1f925bf678aefb7db776260e3b301783865a5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-936e20963654d14c43c413c80c88b6796922e56dd0352707ef00b42e72df0dfd"></a>

## http_receiver.batch — http_receiver.batch / 3328831b8399 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-bf2394cf5d8419274cfd823dd262328c8899eadae65313d7b15a65f3204709bf)
- http_receiver.batch

<a id="canonical-265174117249c41d73ed244dee506d563ef19d1d06afbee6c0b10d7e52d9fe27"></a>

Type: `"object"`. single nested block, Optional.

Batch OPTIONS allow tuning for how batches of logs are sent to an endpoint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("max_bytes",
    "max_bytes_disabled"),
  validators.ConflictingObjectAttributes("max_events",
    "max_events_disabled"),
  validators.ConflictingObjectAttributes("timeout_seconds",
    "timeout_seconds_default")}
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
  "x-ves-oneof-field-batch_bytes": "[\"max_bytes\",\"max_bytes_disabled\"]",
  "x-ves-oneof-field-batch_events": "[\"max_events\",\"max_events_disabled\"]",
  "x-ves-oneof-field-batch_timeout": "[\"timeout_seconds\",\"timeout_seconds_default\"]"
}
```

Terraform syntax:

```terraform
batch {
  # Configure direct properties listed below.
}
```

<a id="canonical-66b9212da76ed566c444db99a697f967a8b4f86d861cf766e06ed9bd3a0f02de"></a>

## Direct properties — http_receiver.batch / 3328831b8399 / 3

<a id="canonical-4572099d4d0a9690bff63322630f1a46648a7ce10aa330559f282658253305f3"></a>

<a id="canonical-62001da660482d1f6e7ff0fb59e85afada9050f975b047d8efb7116181ed1339"></a>

## max_bytes property — http_receiver.batch / 3328831b8399 / 4

Type: `"number"`. Optional.

Exclusive with \[max\_bytes\_disabled\] Send batch to endpoint after the batch is equal to or larger
than this many bytes.

Upstream description:

Exclusive with \[max\_bytes\_disabled\] Send batch to endpoint after the batch is equal to or larger
than this many bytes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(4096, 10485760),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 10485760,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 4096
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "4096",
    "ves.io.schema.rules.uint32.lte": "10485760"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "4096",
    "ves.io.schema.rules.uint32.lte": "10485760"
  }
}
```

- [max_bytes_disabled](resources--global_log_receiver--reference--group-003.md#canonical-ba3250035ff4867759abba31dcc1d6f1ce6bffa4d6b07508966ef745dce14c91): complete subsection reference.

<a id="canonical-b10f7e8d121fb5732ec6f366c94dcf440c65951f2609ab22f7e84f158d8aa381"></a>

<a id="canonical-a6e05098708bc3deb3ecdd4bcbb7683f2a46a9e031d07551594d48e62ce6f6ad"></a>

## max_events property — http_receiver.batch / 3328831b8399 / 5

Type: `"number"`. Optional.

Exclusive with \[max\_events\_disabled\] Send batch to endpoint after this many log messages are in
the batch.

Upstream description:

Exclusive with \[max\_events\_disabled\] Send batch to endpoint after this many log messages are in
the batch.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(32, 2000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 32
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "32",
    "ves.io.schema.rules.uint32.lte": "2000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "32",
    "ves.io.schema.rules.uint32.lte": "2000"
  }
}
```

- [max_events_disabled](resources--global_log_receiver--reference--group-003.md#canonical-e827dc910c7e7783499beb07cdfcac4611c16db6047e8f118cbbfdcb1d48d75e): complete subsection reference.

<a id="canonical-d07fe985a7ac7e0b5c480c424e9b3009ecd50d0a31f1703b5c1236d83e88ce11"></a>

<a id="canonical-d2b96e7ec673b81dc8a9609e492ad33d735c0c078a81f39e48421c94b3c7841d"></a>

## timeout_seconds property — http_receiver.batch / 3328831b8399 / 6

Type: `"string"`. Optional.

Exclusive with \[timeout\_seconds\_default\] Send batch to the endpoint after this many seconds.

Upstream description:

Exclusive with \[timeout\_seconds\_default\] Send batch to the endpoint after this many seconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "uint64",
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
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  }
}
```

- [timeout_seconds_default](resources--global_log_receiver--reference--group-003.md#canonical-f08d181b76efd4647b33f9526154cebf5c06cfa504ba32f9f063a591c12adc42): complete subsection reference.

<a id="canonical-dbeb4db81a9f9d9d885a3cea526ad5e7f3d3dacfbf2333d88c61fa90f8e6eef7"></a>

## Next pages — http_receiver.batch / 3328831b8399 / 7

- [http_receiver.batch.max_bytes_disabled](resources--global_log_receiver--reference--group-003.md#canonical-ba3250035ff4867759abba31dcc1d6f1ce6bffa4d6b07508966ef745dce14c91)
- [http_receiver.batch.max_events_disabled](resources--global_log_receiver--reference--group-003.md#canonical-e827dc910c7e7783499beb07cdfcac4611c16db6047e8f118cbbfdcb1d48d75e)
- [http_receiver.batch.timeout_seconds_default](resources--global_log_receiver--reference--group-003.md#canonical-f08d181b76efd4647b33f9526154cebf5c06cfa504ba32f9f063a591c12adc42)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-bf2394cf5d8419274cfd823dd262328c8899eadae65313d7b15a65f3204709bf)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-ba3250035ff4867759abba31dcc1d6f1ce6bffa4d6b07508966ef745dce14c91"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-84a05e30d2f981ef9149bfa68864710375ec3b8fb5d985ed2e6bf12f4338e291"></a>

## http_receiver.batch.max_bytes_disabled — http_receiver.batch.max_bytes_disabled / 15b42f4d88ff / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-bf2394cf5d8419274cfd823dd262328c8899eadae65313d7b15a65f3204709bf)
- [http_receiver.batch](resources--global_log_receiver--reference--group-003.md#canonical-12d05cef66a26cbbf70082c850e1f925bf678aefb7db776260e3b301783865a5)
- http_receiver.batch.max_bytes_disabled

<a id="canonical-28c46774b17b04fcb68364411a6caa387ac89f932d58558ec894d5e44c101d03"></a>

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
max_bytes_disabled = {}
```

<a id="canonical-65dbf946d706d1678d03b89a0f5b6da154f0fd13f8573dd1c3aafc8effc59d68"></a>

## Direct properties — http_receiver.batch.max_bytes_disabled / 15b42f4d88ff / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-84001d58181014e6af0d1fad61b4980091ec61dadbf67bb367894cc97d8ae7ce"></a>

## Next pages — http_receiver.batch.max_bytes_disabled / 15b42f4d88ff / 4

- [http_receiver.batch](resources--global_log_receiver--reference--group-003.md#canonical-12d05cef66a26cbbf70082c850e1f925bf678aefb7db776260e3b301783865a5)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-e827dc910c7e7783499beb07cdfcac4611c16db6047e8f118cbbfdcb1d48d75e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-70bbbab02aa215ecdc8a4dbf4ac96115b09d81dbc7657043b6dd133bef5a678e"></a>

## http_receiver.batch.max_events_disabled — http_receiver.batch.max_events_disabled / 42d51582e383 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-bf2394cf5d8419274cfd823dd262328c8899eadae65313d7b15a65f3204709bf)
- [http_receiver.batch](resources--global_log_receiver--reference--group-003.md#canonical-12d05cef66a26cbbf70082c850e1f925bf678aefb7db776260e3b301783865a5)
- http_receiver.batch.max_events_disabled

<a id="canonical-34adbbd77b539b3fcc2c864b59e94ee3aa72cab02577c5122dd58a665b4b38b0"></a>

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
max_events_disabled = {}
```

<a id="canonical-c31d94d7d0e7b8ac108a2242ad445b544a395ebc4bfbcde7b13120d54ed85898"></a>

## Direct properties — http_receiver.batch.max_events_disabled / 42d51582e383 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4ba01c4e3c29f5bfd2576fe85357c5846a0f8f73b48ec6b609a7b0a2ad75331f"></a>

## Next pages — http_receiver.batch.max_events_disabled / 42d51582e383 / 4

- [http_receiver.batch](resources--global_log_receiver--reference--group-003.md#canonical-12d05cef66a26cbbf70082c850e1f925bf678aefb7db776260e3b301783865a5)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-f08d181b76efd4647b33f9526154cebf5c06cfa504ba32f9f063a591c12adc42"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bd6416bf3c332ae1e31b3f97cc581905a365912c950023b430586d01708ffabc"></a>

## http_receiver.batch.timeout_seconds_default — http_receiver.batch.timeout_seconds_default / ce475cc24bf1 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-bf2394cf5d8419274cfd823dd262328c8899eadae65313d7b15a65f3204709bf)
- [http_receiver.batch](resources--global_log_receiver--reference--group-003.md#canonical-12d05cef66a26cbbf70082c850e1f925bf678aefb7db776260e3b301783865a5)
- http_receiver.batch.timeout_seconds_default

<a id="canonical-030081acd975e3483ce32411bd784cc047f8060860bed0bc07d4ce8964a06562"></a>

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
timeout_seconds_default = {}
```

<a id="canonical-bbb26a066d0fda65ffad0d61f0f73035f4d5b595a243cc4de4e5e9d93f6e8b4a"></a>

## Direct properties — http_receiver.batch.timeout_seconds_default / ce475cc24bf1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2e483c5524d5fcc59635152040757c4e8581e089f7a3e29c5c7342d75c0065e1"></a>

## Next pages — http_receiver.batch.timeout_seconds_default / ce475cc24bf1 / 4

- [http_receiver.batch](resources--global_log_receiver--reference--group-003.md#canonical-12d05cef66a26cbbf70082c850e1f925bf678aefb7db776260e3b301783865a5)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-df8d6912b3ad7e684c1a27a96dd1746f6ff2921b6ec104c679b78bdf931bbcfd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-173ad0786401bc8945248f3b7f65a1e821eee6f16ba4dee685ac1472cfd8c04b"></a>

## http_receiver.compression — http_receiver.compression / c51ba0afbeb4 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-bf2394cf5d8419274cfd823dd262328c8899eadae65313d7b15a65f3204709bf)
- http_receiver.compression

<a id="canonical-b83b86b3333e83201fe91a9abe5937fb4de066c4c92b32e0cf43c7ab47de1f6e"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for compression.

Upstream description:

Compression Type.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("compression_default",
    "compression_gzip"),
  validators.ConflictingObjectAttributes("compression_default",
    "compression_none"),
  validators.ConflictingObjectAttributes("compression_gzip",
    "compression_none")}
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
  "x-ves-oneof-field-compression_choice": "[\"compression_default\",\"compression_gzip\",\"compression_none\"]"
}
```

Terraform syntax:

```terraform
compression {
  # Configure direct properties listed below.
}
```

<a id="canonical-35893ce939caa3cc164f1cc21dd272d4dec8e55f160f2eea93b745eb93ffbc76"></a>

## Direct properties — http_receiver.compression / c51ba0afbeb4 / 3

- [compression_default](resources--global_log_receiver--reference--group-003.md#canonical-d974e84bd82d50bb8941ed5fb0374988fa4523de49f08e490373eb60a8250079): complete subsection reference.

- [compression_gzip](resources--global_log_receiver--reference--group-003.md#canonical-274ca8982bc7b1cb9260576e40f1fce75b39a9b66a6b08210794b82eccd66795): complete subsection reference.

- [compression_none](resources--global_log_receiver--reference--group-003.md#canonical-d55aaadc075b26fc3c6ecf9eea4939eca0b92820b728e0c4cc7e65893dd07691): complete subsection reference.

<a id="canonical-91aef9ccdb924e5400bea8ba4f4523b85ded7343ff71ecf1483237c1d9822c3d"></a>

## Next pages — http_receiver.compression / c51ba0afbeb4 / 4

- [http_receiver.compression.compression_default](resources--global_log_receiver--reference--group-003.md#canonical-d974e84bd82d50bb8941ed5fb0374988fa4523de49f08e490373eb60a8250079)
- [http_receiver.compression.compression_gzip](resources--global_log_receiver--reference--group-003.md#canonical-274ca8982bc7b1cb9260576e40f1fce75b39a9b66a6b08210794b82eccd66795)
- [http_receiver.compression.compression_none](resources--global_log_receiver--reference--group-003.md#canonical-d55aaadc075b26fc3c6ecf9eea4939eca0b92820b728e0c4cc7e65893dd07691)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-bf2394cf5d8419274cfd823dd262328c8899eadae65313d7b15a65f3204709bf)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-d974e84bd82d50bb8941ed5fb0374988fa4523de49f08e490373eb60a8250079"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2503d114efb322c748328d47dba57b9941c2fe76e8b7354b873dc16506694e3e"></a>

## http_receiver.compression.compression_default — http_receiver.compression.compression_default / 9fc52a905f3f / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-bf2394cf5d8419274cfd823dd262328c8899eadae65313d7b15a65f3204709bf)
- [http_receiver.compression](resources--global_log_receiver--reference--group-003.md#canonical-df8d6912b3ad7e684c1a27a96dd1746f6ff2921b6ec104c679b78bdf931bbcfd)
- http_receiver.compression.compression_default

<a id="canonical-1f44f0023c891d79e7ef64716a3107e02446701e27abf758a47bf231cfe007ec"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for compression default.

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
compression_default = {}
```

<a id="canonical-6dbac556cb2eb33f6e76d9ba3f1476a26de19a47eae560e0a246aaa24f431dba"></a>

## Direct properties — http_receiver.compression.compression_default / 9fc52a905f3f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-02029c15c373ae96d8eb90786a426b2126ced02470c7b67d153c0dd410eb9c54"></a>

## Next pages — http_receiver.compression.compression_default / 9fc52a905f3f / 4

- [http_receiver.compression](resources--global_log_receiver--reference--group-003.md#canonical-df8d6912b3ad7e684c1a27a96dd1746f6ff2921b6ec104c679b78bdf931bbcfd)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-274ca8982bc7b1cb9260576e40f1fce75b39a9b66a6b08210794b82eccd66795"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-97dc6b51b36948b301115d3bea45ab8784c3fec0407bb2b3c88f1cefa59846ec"></a>

## http_receiver.compression.compression_gzip — http_receiver.compression.compression_gzip / df860fb42a99 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-bf2394cf5d8419274cfd823dd262328c8899eadae65313d7b15a65f3204709bf)
- [http_receiver.compression](resources--global_log_receiver--reference--group-003.md#canonical-df8d6912b3ad7e684c1a27a96dd1746f6ff2921b6ec104c679b78bdf931bbcfd)
- http_receiver.compression.compression_gzip

<a id="canonical-fdee09ed184e1484cc13d23af1018f5e7fc25f224026a6701ba2e9b4f9485835"></a>

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
compression_gzip = {}
```

<a id="canonical-4a32fdd348a72036c652d441c356b38308c625bb4e791ea9c4d9baba49a117c0"></a>

## Direct properties — http_receiver.compression.compression_gzip / df860fb42a99 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6c01810e4e4029d718cc1da434963f0ce7eed91dcc4740db7b24d16b3d84c202"></a>

## Next pages — http_receiver.compression.compression_gzip / df860fb42a99 / 4

- [http_receiver.compression](resources--global_log_receiver--reference--group-003.md#canonical-df8d6912b3ad7e684c1a27a96dd1746f6ff2921b6ec104c679b78bdf931bbcfd)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-d55aaadc075b26fc3c6ecf9eea4939eca0b92820b728e0c4cc7e65893dd07691"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-225568196100039964023ef0ac987438314d875fbdd34f52047e92f757ee3562"></a>

## http_receiver.compression.compression_none — http_receiver.compression.compression_none / 8cc0e0a97123 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-bf2394cf5d8419274cfd823dd262328c8899eadae65313d7b15a65f3204709bf)
- [http_receiver.compression](resources--global_log_receiver--reference--group-003.md#canonical-df8d6912b3ad7e684c1a27a96dd1746f6ff2921b6ec104c679b78bdf931bbcfd)
- http_receiver.compression.compression_none

<a id="canonical-3c57af04e0a21ac228feaabf0fd2cfe370fdcb7db8257a7ce443ad1a2d7f3a37"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for compression none.

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
compression_none = {}
```

<a id="canonical-e17c126a243de130c4e5072fc1b84c4dd5fdb257bd8e73764a3700ad19ed3671"></a>

## Direct properties — http_receiver.compression.compression_none / 8cc0e0a97123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-be385c97a4e49a2523b035242db6a74ef1113ed06d898c4fd0c3c69104571be5"></a>

## Next pages — http_receiver.compression.compression_none / 8cc0e0a97123 / 4

- [http_receiver.compression](resources--global_log_receiver--reference--group-003.md#canonical-df8d6912b3ad7e684c1a27a96dd1746f6ff2921b6ec104c679b78bdf931bbcfd)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-ef42f4caac6e7e7d1d04e43192ea5a08d6b0359f312114466068552ea67c74dc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1962e2a286106de97aacf85a54c97d68caad39e821f45a5453c8181ea82b7a75"></a>

## http_receiver.no_tls — http_receiver.no_tls / d32d4cb616f5 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-bf2394cf5d8419274cfd823dd262328c8899eadae65313d7b15a65f3204709bf)
- http_receiver.no_tls

<a id="canonical-0913081908c9173301d9fd42da6d879fa9a9e65c38c17bd8325e931975bd5e62"></a>

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

<a id="canonical-483d2a5fc3357322f0956229c8d12527c299590b59d0f265b8402c7d293791ea"></a>

## Direct properties — http_receiver.no_tls / d32d4cb616f5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-be360d7aefae69f112aa382ca81dc23f2c8e4e7bc3928b276d3c6766804cc83d"></a>

## Next pages — http_receiver.no_tls / d32d4cb616f5 / 4

- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-bf2394cf5d8419274cfd823dd262328c8899eadae65313d7b15a65f3204709bf)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-980d432686a32927854daafff4cda02c9ab6c131c217c3162a1bfeb407fa095b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4b0c8fe06c65760eb03c0f6e8004b4aaa3a1eb50628fb11471805b207542b81f"></a>

## http_receiver.use_tls — http_receiver.use_tls / d5621f7d340a / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-bf2394cf5d8419274cfd823dd262328c8899eadae65313d7b15a65f3204709bf)
- http_receiver.use_tls

<a id="canonical-da823b430abae4f83379d6bebbefe5b487b8bc28900ca55c0c079586ffcd6676"></a>

Type: `"object"`. single nested block, Optional.

TLS Parameters for client connection to the endpoint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_verify_certificate",
    "enable_verify_certificate"),
  validators.ConflictingObjectAttributes("disable_verify_hostname",
    "enable_verify_hostname"),
  validators.ConflictingObjectAttributes("mtls_disabled",
    "mtls_enable"),
  validators.ConflictingObjectAttributes("no_ca",
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
  "x-ves-oneof-field-ca_choice": "[\"no_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-mtls_choice": "[\"mtls_disabled\",\"mtls_enable\"]",
  "x-ves-oneof-field-verify_certificate": "[\"disable_verify_certificate\",\"enable_verify_certificate\"]",
  "x-ves-oneof-field-verify_hostname": "[\"disable_verify_hostname\",\"enable_verify_hostname\"]"
}
```

Terraform syntax:

```terraform
use_tls {
  # Configure direct properties listed below.
}
```

<a id="canonical-f71598aeb352085e051653fd26e9fe0c8091d7842595884c5abf8552a6d2c285"></a>

## Direct properties — http_receiver.use_tls / d5621f7d340a / 3

- [disable_verify_certificate](resources--global_log_receiver--reference--group-003.md#canonical-9461ad227e29d79b5d96faf06adadb2595ec548edf24cd13b16b0b55a375d9da): complete subsection reference.

- [disable_verify_hostname](resources--global_log_receiver--reference--group-003.md#canonical-6d06f884640428bbcf5847c0190b9066588d48263820d749f52ed1c22ea5fec7): complete subsection reference.

- [enable_verify_certificate](resources--global_log_receiver--reference--group-003.md#canonical-5f889a4f2c848e9cc9aba1608254a035bb37c1de7f03b9f1d63d7109b19c3adf): complete subsection reference.

- [enable_verify_hostname](resources--global_log_receiver--reference--group-003.md#canonical-00f48e482afd0de52255ce4bc298f5cb15b8a1ad3d899be6f85016b598c023e4): complete subsection reference.

- [mtls_disabled](resources--global_log_receiver--reference--group-003.md#canonical-977fb94806c27a70d04a8e13b7876d0ddbd4a510ab6777cf39960bc21fad40eb): complete subsection reference.

- [mtls_enable](resources--global_log_receiver--reference--group-003.md#canonical-360aa9e4eb3d47f78c2734ac8701174eb5a028d2c3475033c37bd90dd7d10724): complete subsection reference.

- [no_ca](resources--global_log_receiver--reference--group-003.md#canonical-9892e568538618feeed135f75cbbbe40c7dfc1914fbc95c69dae2bd3b3d39c80): complete subsection reference.

<a id="canonical-3ac6d04b08336314ffd44d1585bd836c365a6c3366623a8424b71e0e75c03baf"></a>

<a id="canonical-1b39828100ceaf42eac6277afae3fadb3c9ee1b3d54df171c8c4546ec123f154"></a>

## trusted_ca_url property — http_receiver.use_tls / d5621f7d340a / 4

Type: `"string"`. Optional.

Exclusive with \[no\_ca\] The URL or value for trusted Server CA certificate or certificate chain
Certificates in PEM format including the PEM headers.

Upstream description:

Exclusive with \[no\_ca\] The URL or value for trusted Server CA certificate or certificate chain
Certificates in PEM format including the PEM headers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(131072),
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
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

<a id="canonical-c997529b70e10c1b04be6f891b3010c25d7cf9af2a298ee78644034852bb45b0"></a>

## Next pages — http_receiver.use_tls / d5621f7d340a / 5

- [http_receiver.use_tls.disable_verify_certificate](resources--global_log_receiver--reference--group-003.md#canonical-9461ad227e29d79b5d96faf06adadb2595ec548edf24cd13b16b0b55a375d9da)
- [http_receiver.use_tls.disable_verify_hostname](resources--global_log_receiver--reference--group-003.md#canonical-6d06f884640428bbcf5847c0190b9066588d48263820d749f52ed1c22ea5fec7)
- [http_receiver.use_tls.enable_verify_certificate](resources--global_log_receiver--reference--group-003.md#canonical-5f889a4f2c848e9cc9aba1608254a035bb37c1de7f03b9f1d63d7109b19c3adf)
- [http_receiver.use_tls.enable_verify_hostname](resources--global_log_receiver--reference--group-003.md#canonical-00f48e482afd0de52255ce4bc298f5cb15b8a1ad3d899be6f85016b598c023e4)
- [http_receiver.use_tls.mtls_disabled](resources--global_log_receiver--reference--group-003.md#canonical-977fb94806c27a70d04a8e13b7876d0ddbd4a510ab6777cf39960bc21fad40eb)
- [http_receiver.use_tls.mtls_enable](resources--global_log_receiver--reference--group-003.md#canonical-360aa9e4eb3d47f78c2734ac8701174eb5a028d2c3475033c37bd90dd7d10724)
- [http_receiver.use_tls.no_ca](resources--global_log_receiver--reference--group-003.md#canonical-9892e568538618feeed135f75cbbbe40c7dfc1914fbc95c69dae2bd3b3d39c80)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-bf2394cf5d8419274cfd823dd262328c8899eadae65313d7b15a65f3204709bf)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-9461ad227e29d79b5d96faf06adadb2595ec548edf24cd13b16b0b55a375d9da"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3212e2681b3ac8b9d871fccf0ccd2f605c4f9b54e8a5a6779d6d47eb65e8fef3"></a>

## http_receiver.use_tls.disable_verify_certificate — http_receiver.use_tls.disable_verify_certificate / bf09d2de619a / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-bf2394cf5d8419274cfd823dd262328c8899eadae65313d7b15a65f3204709bf)
- [http_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-980d432686a32927854daafff4cda02c9ab6c131c217c3162a1bfeb407fa095b)
- http_receiver.use_tls.disable_verify_certificate

<a id="canonical-6c1587ffa4f5b09118af3ee7cc54b100b2862d28084ae8a53298ef16de45deb7"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable verify certificate.

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
disable_verify_certificate = {}
```

<a id="canonical-2f0baa29444e1eab72b9035aa14699c23bf9c007bc914aacb99a09ae6d98d9db"></a>

## Direct properties — http_receiver.use_tls.disable_verify_certificate / bf09d2de619a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-718f584e97a2c9c73b96f01c27c1e5ca47ad63e888d82fe33217fc5c04aebd43"></a>

## Next pages — http_receiver.use_tls.disable_verify_certificate / bf09d2de619a / 4

- [http_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-980d432686a32927854daafff4cda02c9ab6c131c217c3162a1bfeb407fa095b)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-6d06f884640428bbcf5847c0190b9066588d48263820d749f52ed1c22ea5fec7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0f256202b1635a567688c757b50bd9f593c981bbd784ffd838ebeecb87eecf45"></a>

## http_receiver.use_tls.disable_verify_hostname — http_receiver.use_tls.disable_verify_hostname / 0f53702b567a / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-bf2394cf5d8419274cfd823dd262328c8899eadae65313d7b15a65f3204709bf)
- [http_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-980d432686a32927854daafff4cda02c9ab6c131c217c3162a1bfeb407fa095b)
- http_receiver.use_tls.disable_verify_hostname

<a id="canonical-0a65a6361ac1e28abcc46d1b69b9d65155c3b79d0b73dc9ea41a0e4949c7ba92"></a>

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
disable_verify_hostname = {}
```

<a id="canonical-46d0938cddf68a554727273c4dcccdadc321d3481796907a747be4a0caa47d42"></a>

## Direct properties — http_receiver.use_tls.disable_verify_hostname / 0f53702b567a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bc99b9a99d67aaaab9e324ae71001da3f289b26bd2bb824afc09b78558d5dc3e"></a>

## Next pages — http_receiver.use_tls.disable_verify_hostname / 0f53702b567a / 4

- [http_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-980d432686a32927854daafff4cda02c9ab6c131c217c3162a1bfeb407fa095b)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-5f889a4f2c848e9cc9aba1608254a035bb37c1de7f03b9f1d63d7109b19c3adf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6f16662df7bf7fcd35985131082acd418acbc0a444de12be9b31b119afd505ec"></a>

## http_receiver.use_tls.enable_verify_certificate — http_receiver.use_tls.enable_verify_certificate / fea1f6bd4804 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-bf2394cf5d8419274cfd823dd262328c8899eadae65313d7b15a65f3204709bf)
- [http_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-980d432686a32927854daafff4cda02c9ab6c131c217c3162a1bfeb407fa095b)
- http_receiver.use_tls.enable_verify_certificate

<a id="canonical-e7d6568978db5e8db1944a8e12021140e37eff94f283750f3cedf006f3312030"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable verify certificate.

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
enable_verify_certificate = {}
```

<a id="canonical-397a1dc01c5ecccda15ad7465e6f63393932263c84ad27fa53f6752276270d67"></a>

## Direct properties — http_receiver.use_tls.enable_verify_certificate / fea1f6bd4804 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-927665f821c233b174f853e9c395458da8cc4a0c7567321f7e3495819d6a4c7a"></a>

## Next pages — http_receiver.use_tls.enable_verify_certificate / fea1f6bd4804 / 4

- [http_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-980d432686a32927854daafff4cda02c9ab6c131c217c3162a1bfeb407fa095b)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-00f48e482afd0de52255ce4bc298f5cb15b8a1ad3d899be6f85016b598c023e4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-652dbe70abd3dc8093569cad7998eb082106f7f932f6ddf2102ec9f67b39ef76"></a>

## http_receiver.use_tls.enable_verify_hostname — http_receiver.use_tls.enable_verify_hostname / 1be8eabc66c5 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-bf2394cf5d8419274cfd823dd262328c8899eadae65313d7b15a65f3204709bf)
- [http_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-980d432686a32927854daafff4cda02c9ab6c131c217c3162a1bfeb407fa095b)
- http_receiver.use_tls.enable_verify_hostname

<a id="canonical-9857d6d8423ce08b1dbfaf7c3721358d82393e4b81d46e5bd4222ed0fd8d2c2e"></a>

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
enable_verify_hostname = {}
```

<a id="canonical-f0b8f724d0d5264bcd28df03d189e8c6d74e64ccd3ed1a369ebf54638f67996d"></a>

## Direct properties — http_receiver.use_tls.enable_verify_hostname / 1be8eabc66c5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-984b705ac955d1026cb8e9d90a4a6e46641035fb5a9a2000831b0ebce681d91f"></a>

## Next pages — http_receiver.use_tls.enable_verify_hostname / 1be8eabc66c5 / 4

- [http_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-980d432686a32927854daafff4cda02c9ab6c131c217c3162a1bfeb407fa095b)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-977fb94806c27a70d04a8e13b7876d0ddbd4a510ab6777cf39960bc21fad40eb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f0e0c57202e9298c58c69eaefef1d4ead19fefaacde64b8c1534135e9b783fb9"></a>

## http_receiver.use_tls.mtls_disabled — http_receiver.use_tls.mtls_disabled / 10c14f8c63af / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-bf2394cf5d8419274cfd823dd262328c8899eadae65313d7b15a65f3204709bf)
- [http_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-980d432686a32927854daafff4cda02c9ab6c131c217c3162a1bfeb407fa095b)
- http_receiver.use_tls.mtls_disabled

<a id="canonical-fdddf224b0658e2786c4dc917ade3a70c5c9deb4e6bf0aa1e4cbb2b90d706c70"></a>

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
mtls_disabled = {}
```

<a id="canonical-968a3b1046f51dd743fdefd6c32c48b63b808270c3b8492a70f3a82525ad7578"></a>

## Direct properties — http_receiver.use_tls.mtls_disabled / 10c14f8c63af / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-070433d66fa01297de4a1fdc3f0a02c55e5269ccf96f4ad8fb479cc72fefb691"></a>

## Next pages — http_receiver.use_tls.mtls_disabled / 10c14f8c63af / 4

- [http_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-980d432686a32927854daafff4cda02c9ab6c131c217c3162a1bfeb407fa095b)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-360aa9e4eb3d47f78c2734ac8701174eb5a028d2c3475033c37bd90dd7d10724"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bbec80e584b128024abc04e5c8db5bc9226ac4ce499d4521e1413431bf7045e0"></a>

## http_receiver.use_tls.mtls_enable — http_receiver.use_tls.mtls_enable / 24ac83c3157b / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-bf2394cf5d8419274cfd823dd262328c8899eadae65313d7b15a65f3204709bf)
- [http_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-980d432686a32927854daafff4cda02c9ab6c131c217c3162a1bfeb407fa095b)
- http_receiver.use_tls.mtls_enable

<a id="canonical-ed0f3191afc33ae9591bc93ca15ab527036207533996cabcb3e8453c0f6048d2"></a>

Type: `"object"`. single nested block, Optional.

MTLS Client config allows configuration of mTLS client OPTIONS.

Receipt-pinned upstream constraints:

```json
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
mtls_enable {
  # Configure direct properties listed below.
}
```

<a id="canonical-fa633f7e0f65b6ae0bdf4909c3d069ea7557027affaf2afbef9e9213c992a621"></a>

## Direct properties — http_receiver.use_tls.mtls_enable / 24ac83c3157b / 3

<a id="canonical-1cfe03693aa2f9c35fd567371f986983386e2651e2c178ddb48e2199195a4d1b"></a>

<a id="canonical-3f2014e7c458d92c0d3c59d049628a004a6bafca0528ccce86d3849d3facd150"></a>

## certificate property — http_receiver.use_tls.mtls_enable / 24ac83c3157b / 4

Type: `"string"`. Optional.

Client certificate is PEM-encoded certificate or certificate-chain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(100, 131072),
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

- [key_url](resources--global_log_receiver--reference--group-003.md#canonical-3a81154af6ea05f703d8e5fef45b0309ba44e1ed40d7c194102e775e4b861882): complete subsection reference.

<a id="canonical-6f34c4ddbddcc66d55c65e24d27da0e85bf5481e1af99fd503f5d80a233857a9"></a>

## Next pages — http_receiver.use_tls.mtls_enable / 24ac83c3157b / 5

- [http_receiver.use_tls.mtls_enable.key_url](resources--global_log_receiver--reference--group-003.md#canonical-3a81154af6ea05f703d8e5fef45b0309ba44e1ed40d7c194102e775e4b861882)
- [http_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-980d432686a32927854daafff4cda02c9ab6c131c217c3162a1bfeb407fa095b)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-3a81154af6ea05f703d8e5fef45b0309ba44e1ed40d7c194102e775e4b861882"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d047573b4a39d015ce58c219c06f4ddce31aa8d6e1ca5adb3875f13753ed21b4"></a>

## http_receiver.use_tls.mtls_enable.key_url — http_receiver.use_tls.mtls_enable.key_url / 97b5731f7da1 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-bf2394cf5d8419274cfd823dd262328c8899eadae65313d7b15a65f3204709bf)
- [http_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-980d432686a32927854daafff4cda02c9ab6c131c217c3162a1bfeb407fa095b)
- [http_receiver.use_tls.mtls_enable](resources--global_log_receiver--reference--group-003.md#canonical-360aa9e4eb3d47f78c2734ac8701174eb5a028d2c3475033c37bd90dd7d10724)
- http_receiver.use_tls.mtls_enable.key_url

<a id="canonical-63d19fa58d91c691a326cd544ecfdbf9ad67f000cf8e504ed47c82a824ab8aac"></a>

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
key_url {
  # Configure direct properties listed below.
}
```

<a id="canonical-4ea9ac30ff1118c31165e01f13d688c4a1dfb32135838822bec4cc38bf224dd7"></a>

## Direct properties — http_receiver.use_tls.mtls_enable.key_url / 97b5731f7da1 / 3

- [blindfold_secret_info](resources--global_log_receiver--reference--group-003.md#canonical-8cc4c7cee6df7ddd12cde165e5357f3a02bbdf55a0ed253504ac1c274b090571): complete subsection reference.

- [clear_secret_info](resources--global_log_receiver--reference--group-003.md#canonical-e37ac288535e128c13b47aa9cb5ceec86b2edcd6387de976efc0f81e3772ad7e): complete subsection reference.

<a id="canonical-9a1927cf54be1817731cd613e050cd785bcecc5894195700e13a53b30468c555"></a>

## Next pages — http_receiver.use_tls.mtls_enable.key_url / 97b5731f7da1 / 4

- [http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info](resources--global_log_receiver--reference--group-003.md#canonical-8cc4c7cee6df7ddd12cde165e5357f3a02bbdf55a0ed253504ac1c274b090571)
- [http_receiver.use_tls.mtls_enable.key_url.clear_secret_info](resources--global_log_receiver--reference--group-003.md#canonical-e37ac288535e128c13b47aa9cb5ceec86b2edcd6387de976efc0f81e3772ad7e)
- [http_receiver.use_tls.mtls_enable](resources--global_log_receiver--reference--group-003.md#canonical-360aa9e4eb3d47f78c2734ac8701174eb5a028d2c3475033c37bd90dd7d10724)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-8cc4c7cee6df7ddd12cde165e5357f3a02bbdf55a0ed253504ac1c274b090571"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6baed217965e8569fba16fced769b8ad5c661f86cfced3455edb6e8a06d5a7c6"></a>

## http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info — http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info / f20ca9771b3c / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-bf2394cf5d8419274cfd823dd262328c8899eadae65313d7b15a65f3204709bf)
- [http_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-980d432686a32927854daafff4cda02c9ab6c131c217c3162a1bfeb407fa095b)
- [http_receiver.use_tls.mtls_enable](resources--global_log_receiver--reference--group-003.md#canonical-360aa9e4eb3d47f78c2734ac8701174eb5a028d2c3475033c37bd90dd7d10724)
- [http_receiver.use_tls.mtls_enable.key_url](resources--global_log_receiver--reference--group-003.md#canonical-3a81154af6ea05f703d8e5fef45b0309ba44e1ed40d7c194102e775e4b861882)
- http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info

<a id="canonical-0a8d1948204027f0291065ba458ba1af149837b659057b2c7bc5e2971506543f"></a>

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

<a id="canonical-6e9d1cc222c55b672b7d2e1cc48b45e072462013e7504ec29339e5931318710e"></a>

## Direct properties — http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info / f20ca9771b3c / 3

<a id="canonical-16ba461fe7c6e147eda28f6dccf5e0b3121b3ddcdda4b159f78cd2d5b072cd28"></a>

<a id="canonical-b2e119109b61c19dcea99a19d1d676035dc7643d682cc3c774c07edd05a61607"></a>

## decryption_provider property — http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info / f20ca9771b3c / 4

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

<a id="canonical-d8177e33c399d486470733b2da4945343f29edf28037605013cdcef709ff0391"></a>

<a id="canonical-e58de8acd321683dc01eba80c4a24aa2bc83409a0c44ae560d78b6d45a732b17"></a>

## location property — http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info / f20ca9771b3c / 5

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

<a id="canonical-5502da4075f91726723f0109815182ace83b4edb48aa7835ff4f129308b6586d"></a>

<a id="canonical-bc0baaaf06cf457636109c9c06d951d573272d6f510f9c93e0cfc928a88f6e3b"></a>

## store_provider property — http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info / f20ca9771b3c / 6

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

<a id="canonical-26de7c6e8f9df1e5754441f62eb3849d5fb762f68fd51b932b6a2b069500c7a6"></a>

## Next pages — http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info / f20ca9771b3c / 7

- [http_receiver.use_tls.mtls_enable.key_url](resources--global_log_receiver--reference--group-003.md#canonical-3a81154af6ea05f703d8e5fef45b0309ba44e1ed40d7c194102e775e4b861882)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-e37ac288535e128c13b47aa9cb5ceec86b2edcd6387de976efc0f81e3772ad7e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-30680fac96496b72f30907dbb1bcdfdc8b55eab38b711f59a2e5b7858fd9e6ad"></a>

## http_receiver.use_tls.mtls_enable.key_url.clear_secret_info — http_receiver.use_tls.mtls_enable.key_url.clear_secret_info / f3c05a9233e0 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-bf2394cf5d8419274cfd823dd262328c8899eadae65313d7b15a65f3204709bf)
- [http_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-980d432686a32927854daafff4cda02c9ab6c131c217c3162a1bfeb407fa095b)
- [http_receiver.use_tls.mtls_enable](resources--global_log_receiver--reference--group-003.md#canonical-360aa9e4eb3d47f78c2734ac8701174eb5a028d2c3475033c37bd90dd7d10724)
- [http_receiver.use_tls.mtls_enable.key_url](resources--global_log_receiver--reference--group-003.md#canonical-3a81154af6ea05f703d8e5fef45b0309ba44e1ed40d7c194102e775e4b861882)
- http_receiver.use_tls.mtls_enable.key_url.clear_secret_info

<a id="canonical-09d644c23bb3ab5998307af547a40f163e1f1a11822ff3bd1e26486fb4b2f98e"></a>

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

<a id="canonical-d9c9de690c435527b1bc830e1eb689761d3a7fb29ea42112e509f04b230d9d6f"></a>

## Direct properties — http_receiver.use_tls.mtls_enable.key_url.clear_secret_info / f3c05a9233e0 / 3

<a id="canonical-0d714e528a46c0a2af3194bafedb1b62a9b4912b7cd5743697bdbf2d7ca2242d"></a>

<a id="canonical-8d58e8e39e207583d5bf58e080447c138be030f7fd837b8f8aad67dcf9530812"></a>

## provider_ref property — http_receiver.use_tls.mtls_enable.key_url.clear_secret_info / f3c05a9233e0 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-26c0bb8f9ec97e1cb3176fd691f39b204d68c74a5489a7f63f9438c7546fc2b9"></a>

<a id="canonical-f5a6c8853f5085fa7397d51afe43ad3c28d204ba7692f45e41155901be4dbcba"></a>

## url property — http_receiver.use_tls.mtls_enable.key_url.clear_secret_info / f3c05a9233e0 / 5

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

<a id="canonical-2147d004eba4cde852cd6cf3b61a6c8c78e7958cb7e74d3f8fff20eea54542fb"></a>

## Next pages — http_receiver.use_tls.mtls_enable.key_url.clear_secret_info / f3c05a9233e0 / 6

- [http_receiver.use_tls.mtls_enable.key_url](resources--global_log_receiver--reference--group-003.md#canonical-3a81154af6ea05f703d8e5fef45b0309ba44e1ed40d7c194102e775e4b861882)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-9892e568538618feeed135f75cbbbe40c7dfc1914fbc95c69dae2bd3b3d39c80"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-539225a268757c976fed9d619cbdb2fbecb0b96c13ef12e7826db0d8527118d7"></a>

## http_receiver.use_tls.no_ca — http_receiver.use_tls.no_ca / 905122665590 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-bf2394cf5d8419274cfd823dd262328c8899eadae65313d7b15a65f3204709bf)
- [http_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-980d432686a32927854daafff4cda02c9ab6c131c217c3162a1bfeb407fa095b)
- http_receiver.use_tls.no_ca

<a id="canonical-38cc4af7d34304c708168d1b143c7eeaefdfd8de8724c6b37e2fb50c6576cf62"></a>

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
no_ca = {}
```

<a id="canonical-e15f4ec3880ec88de7ddd41bacd1abab1484d599aae78d38fdb66fd6caf93a80"></a>

## Direct properties — http_receiver.use_tls.no_ca / 905122665590 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a34e41832c6432ee75e4f3a619ed39aee45d3c6fa79e037a94d0136deac5ac2b"></a>

## Next pages — http_receiver.use_tls.no_ca / 905122665590 / 4

- [http_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-980d432686a32927854daafff4cda02c9ab6c131c217c3162a1bfeb407fa095b)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-ee10d6aa6bae7f44f1fe9b11b6baa91894549e346d1cced532d36d362b306e96"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-30d71126a29f1b01d9d2b294cc1f40f82e6ac0830dab4439ff630d2aac636022"></a>

## kafka_receiver — kafka_receiver / 04f74101e9c3 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- kafka_receiver

<a id="canonical-30974c0ff3bb7bbe13c8e4b6fdd227339653da44184623051eff1f0043a6b046"></a>

Type: `"object"`. single nested block, Optional.

Kafka Configuration for Global Log Receiver.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("bootstrap_servers",
    "kafka_topic"),
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
kafka_receiver {
  # Configure direct properties listed below.
}
```

<a id="canonical-59fd6d6050383ae6ac33b39fce4a62a8fd154980251e7b40785803eca53c0f68"></a>

## Direct properties — kafka_receiver / 04f74101e9c3 / 3

- [batch](resources--global_log_receiver--reference--group-003.md#canonical-857739033b743b8ad0dfd8f51198dec65af38d70753991dd5f9039e10a1189b5): complete subsection reference.

<a id="canonical-6991fcf1dde293e7176fc9d2cb81576fd904e807fe54ab5a66d6a8a3e3e59c7b"></a>

<a id="canonical-23ad89da35591ea8a192b30b485e1bfd2e27dd89b7a5b02a3b91488c2aa11096"></a>

## bootstrap_servers property — kafka_receiver / 04f74101e9c3 / 4

Type: `["list", "string"]`. Optional.

List of host:port pairs of the Kafka brokers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 8),
}
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
    "ves.io.schema.rules.repeated.items.string.hostport": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.hostport": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [compression](resources--global_log_receiver--reference--group-003.md#canonical-1ca9494a9e6eb285bbb8a5bab44ed17e05366884a9ea5b51d8bbcb3cee6323be): complete subsection reference.

<a id="canonical-2282ccbe31771db16be8ba1be3539e6545bda85edf555c2c67064a73b55af75c"></a>

<a id="canonical-18608af390dbe22b6de8bdf3a1a62bc8ef06ecbf614921cb87f18cc6cde0a328"></a>

## kafka_topic property — kafka_receiver / 04f74101e9c3 / 5

Type: `"string"`. Optional.

The Kafka topic name to write events to.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(3, 255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "minLength": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 3,
    "pattern": "^[a-zA-Z0-9\\\\._\\\\-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[a-zA-Z0-9\\\\._\\\\-]+$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[a-zA-Z0-9\\\\._\\\\-]+$"
  }
}
```

- [no_tls](resources--global_log_receiver--reference--group-003.md#canonical-7c013a410e875b6e0988c4b653ed8c04de84635c9f23405fb0b1e831599bbd93): complete subsection reference.

- [use_tls](resources--global_log_receiver--reference--group-003.md#canonical-a8ec005d981c013ecfd65ce79982faa1c0e75051f53167339dd024dd81d6a2b1): complete subsection reference.

<a id="canonical-75f619ecc1d3b02b6dbdac0d3e5fc606b56157e5fd8001b5a14415c0a9afa897"></a>

## Next pages — kafka_receiver / 04f74101e9c3 / 6

- [kafka_receiver.batch](resources--global_log_receiver--reference--group-003.md#canonical-857739033b743b8ad0dfd8f51198dec65af38d70753991dd5f9039e10a1189b5)
- [kafka_receiver.compression](resources--global_log_receiver--reference--group-003.md#canonical-1ca9494a9e6eb285bbb8a5bab44ed17e05366884a9ea5b51d8bbcb3cee6323be)
- [kafka_receiver.no_tls](resources--global_log_receiver--reference--group-003.md#canonical-7c013a410e875b6e0988c4b653ed8c04de84635c9f23405fb0b1e831599bbd93)
- [kafka_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-a8ec005d981c013ecfd65ce79982faa1c0e75051f53167339dd024dd81d6a2b1)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-857739033b743b8ad0dfd8f51198dec65af38d70753991dd5f9039e10a1189b5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1f07fab148cb298ffd8f4b3f46a3bd1775eb4c1de715716ebccfe8d0e77b45ab"></a>

## kafka_receiver.batch — kafka_receiver.batch / af458c90f52f / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-ee10d6aa6bae7f44f1fe9b11b6baa91894549e346d1cced532d36d362b306e96)
- kafka_receiver.batch

<a id="canonical-963a6ba94cfc28a6f2efb026eb35c4f2385ae27c8d8b109f8b1da9b3a37e6c54"></a>

Type: `"object"`. single nested block, Optional.

Batch OPTIONS allow tuning for how batches of logs are sent to an endpoint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("max_bytes",
    "max_bytes_disabled"),
  validators.ConflictingObjectAttributes("max_events",
    "max_events_disabled"),
  validators.ConflictingObjectAttributes("timeout_seconds",
    "timeout_seconds_default")}
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
  "x-ves-oneof-field-batch_bytes": "[\"max_bytes\",\"max_bytes_disabled\"]",
  "x-ves-oneof-field-batch_events": "[\"max_events\",\"max_events_disabled\"]",
  "x-ves-oneof-field-batch_timeout": "[\"timeout_seconds\",\"timeout_seconds_default\"]"
}
```

Terraform syntax:

```terraform
batch {
  # Configure direct properties listed below.
}
```

<a id="canonical-95c0d4f5bcdf51f78dfa05edb448ca0793791135400a41a75c993306d6b580b6"></a>

## Direct properties — kafka_receiver.batch / af458c90f52f / 3

<a id="canonical-117dc8605914da4820916b12d4dbfbdca0c72d7f83ac4297177926d989ddbc99"></a>

<a id="canonical-0b91a2e8f1a04ad9d604e1182195ef886f12e2c488c8fda22e5f72b5c020cdf1"></a>

## max_bytes property — kafka_receiver.batch / af458c90f52f / 4

Type: `"number"`. Optional.

Exclusive with \[max\_bytes\_disabled\] Send batch to endpoint after the batch is equal to or larger
than this many bytes.

Upstream description:

Exclusive with \[max\_bytes\_disabled\] Send batch to endpoint after the batch is equal to or larger
than this many bytes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(4096, 10485760),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 10485760,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 4096
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "4096",
    "ves.io.schema.rules.uint32.lte": "10485760"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "4096",
    "ves.io.schema.rules.uint32.lte": "10485760"
  }
}
```

- [max_bytes_disabled](resources--global_log_receiver--reference--group-003.md#canonical-03d11b209c63740ba5ddb019bc7153f8fa21aa9c40500b09918cfc28a641d0ef): complete subsection reference.

<a id="canonical-04f485a24ee8e2ab64a7c8e301d3b0b5c884bd40da07168612adb464a79bc365"></a>

<a id="canonical-13719af4a83273e2ab32ca889be85d86d471656a8ca51319785f2116dcf50608"></a>

## max_events property — kafka_receiver.batch / af458c90f52f / 5

Type: `"number"`. Optional.

Exclusive with \[max\_events\_disabled\] Send batch to endpoint after this many log messages are in
the batch.

Upstream description:

Exclusive with \[max\_events\_disabled\] Send batch to endpoint after this many log messages are in
the batch.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(32, 2000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 32
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "32",
    "ves.io.schema.rules.uint32.lte": "2000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "32",
    "ves.io.schema.rules.uint32.lte": "2000"
  }
}
```

- [max_events_disabled](resources--global_log_receiver--reference--group-003.md#canonical-6c7eb470971eec68afe8ca9c4bc7a1d8fdc8eee809623d26720683e2e99ee290): complete subsection reference.

<a id="canonical-47945ad1af0cba0139bd6b85c370e983ce9ca663a6ee07e1775664bd54b87fe0"></a>

<a id="canonical-5f6e048974014f08196a2f8e831784ab19005b2f29f8c35f35c5c53b616fbfd1"></a>

## timeout_seconds property — kafka_receiver.batch / af458c90f52f / 6

Type: `"string"`. Optional.

Exclusive with \[timeout\_seconds\_default\] Send batch to the endpoint after this many seconds.

Upstream description:

Exclusive with \[timeout\_seconds\_default\] Send batch to the endpoint after this many seconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "uint64",
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
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  }
}
```

- [timeout_seconds_default](resources--global_log_receiver--reference--group-003.md#canonical-2ceb89926db0e6f5860aeab7ed0d50fc468b0ab9a44a8902cb5ff55c00caf126): complete subsection reference.

<a id="canonical-3360e75f2d7176855e4fa434cf5374c9543626b279bcdfab4e774e59459aabd4"></a>

## Next pages — kafka_receiver.batch / af458c90f52f / 7

- [kafka_receiver.batch.max_bytes_disabled](resources--global_log_receiver--reference--group-003.md#canonical-03d11b209c63740ba5ddb019bc7153f8fa21aa9c40500b09918cfc28a641d0ef)
- [kafka_receiver.batch.max_events_disabled](resources--global_log_receiver--reference--group-003.md#canonical-6c7eb470971eec68afe8ca9c4bc7a1d8fdc8eee809623d26720683e2e99ee290)
- [kafka_receiver.batch.timeout_seconds_default](resources--global_log_receiver--reference--group-003.md#canonical-2ceb89926db0e6f5860aeab7ed0d50fc468b0ab9a44a8902cb5ff55c00caf126)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-ee10d6aa6bae7f44f1fe9b11b6baa91894549e346d1cced532d36d362b306e96)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-03d11b209c63740ba5ddb019bc7153f8fa21aa9c40500b09918cfc28a641d0ef"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e672bbb27ba25a7d8025e73c7613c53be7189c175f9f4245ff24196bd25190e4"></a>

## kafka_receiver.batch.max_bytes_disabled — kafka_receiver.batch.max_bytes_disabled / 75e26d758369 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-ee10d6aa6bae7f44f1fe9b11b6baa91894549e346d1cced532d36d362b306e96)
- [kafka_receiver.batch](resources--global_log_receiver--reference--group-003.md#canonical-857739033b743b8ad0dfd8f51198dec65af38d70753991dd5f9039e10a1189b5)
- kafka_receiver.batch.max_bytes_disabled

<a id="canonical-56a6748d774bd3896fd9a9e1dc32aff3ade779c4f52e8406bfdf9a295e411329"></a>

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
max_bytes_disabled = {}
```

<a id="canonical-7d2885264cb769b93b212bff9c62bddcd15992bb88256fa3364f7279455e9abc"></a>

## Direct properties — kafka_receiver.batch.max_bytes_disabled / 75e26d758369 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ffca60aa56360b4fddb1fa2c7ffa296bce6ae3dac947167eb7998e4d32eeaf89"></a>

## Next pages — kafka_receiver.batch.max_bytes_disabled / 75e26d758369 / 4

- [kafka_receiver.batch](resources--global_log_receiver--reference--group-003.md#canonical-857739033b743b8ad0dfd8f51198dec65af38d70753991dd5f9039e10a1189b5)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-6c7eb470971eec68afe8ca9c4bc7a1d8fdc8eee809623d26720683e2e99ee290"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a30d58213e3f338e82528200382fb9b93186012cf7779ff24dd5d2b62246b5e3"></a>

## kafka_receiver.batch.max_events_disabled — kafka_receiver.batch.max_events_disabled / 1b21ab99f9fa / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-ee10d6aa6bae7f44f1fe9b11b6baa91894549e346d1cced532d36d362b306e96)
- [kafka_receiver.batch](resources--global_log_receiver--reference--group-003.md#canonical-857739033b743b8ad0dfd8f51198dec65af38d70753991dd5f9039e10a1189b5)
- kafka_receiver.batch.max_events_disabled

<a id="canonical-2afb15967b8cc8a8d31ca3e3a3e20fcb829ed19513cd9b0c8c2751ed5817f4a7"></a>

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
max_events_disabled = {}
```

<a id="canonical-618d2dde5c32edc72227f86ea5279605e0149f1541d564fdc45828b2165c5f54"></a>

## Direct properties — kafka_receiver.batch.max_events_disabled / 1b21ab99f9fa / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-63f69a19dbd0d4dbc0947670300812e55ad065b0f1df1098c0f6276cc2a1b725"></a>

## Next pages — kafka_receiver.batch.max_events_disabled / 1b21ab99f9fa / 4

- [kafka_receiver.batch](resources--global_log_receiver--reference--group-003.md#canonical-857739033b743b8ad0dfd8f51198dec65af38d70753991dd5f9039e10a1189b5)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-2ceb89926db0e6f5860aeab7ed0d50fc468b0ab9a44a8902cb5ff55c00caf126"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8ddc4564f00045da7607bec6c7d5efe1dfde4786e24b6b78415051e847993c17"></a>

## kafka_receiver.batch.timeout_seconds_default — kafka_receiver.batch.timeout_seconds_default / b353fa2ce1fd / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-ee10d6aa6bae7f44f1fe9b11b6baa91894549e346d1cced532d36d362b306e96)
- [kafka_receiver.batch](resources--global_log_receiver--reference--group-003.md#canonical-857739033b743b8ad0dfd8f51198dec65af38d70753991dd5f9039e10a1189b5)
- kafka_receiver.batch.timeout_seconds_default

<a id="canonical-9ff7d9d5375134e1f689735cc4d40e2df34fc4ee3c2a5e624f89bcf6e200e7a2"></a>

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
timeout_seconds_default = {}
```

<a id="canonical-0f0bd309d37a31de0b7babe4825ad7ba9410ac95149154ea7cd6818504df75a5"></a>

## Direct properties — kafka_receiver.batch.timeout_seconds_default / b353fa2ce1fd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-608db8adaf2a29b0ea97fd1260ddcd2dd32d7599c6c63ce8795b4975892facf8"></a>

## Next pages — kafka_receiver.batch.timeout_seconds_default / b353fa2ce1fd / 4

- [kafka_receiver.batch](resources--global_log_receiver--reference--group-003.md#canonical-857739033b743b8ad0dfd8f51198dec65af38d70753991dd5f9039e10a1189b5)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-1ca9494a9e6eb285bbb8a5bab44ed17e05366884a9ea5b51d8bbcb3cee6323be"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0f2ef815616345c76f7c46d2ed14daea8792cb1193aa6af44f61e81128220886"></a>

## kafka_receiver.compression — kafka_receiver.compression / 61eebbec6736 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-ee10d6aa6bae7f44f1fe9b11b6baa91894549e346d1cced532d36d362b306e96)
- kafka_receiver.compression

<a id="canonical-9deaaf69949c258db0e01ea4908575aec8d42b0f3d575cfd92b2c9feeeeea3d0"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for compression.

Upstream description:

Compression Type.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("compression_default",
    "compression_gzip"),
  validators.ConflictingObjectAttributes("compression_default",
    "compression_none"),
  validators.ConflictingObjectAttributes("compression_gzip",
    "compression_none")}
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
  "x-ves-oneof-field-compression_choice": "[\"compression_default\",\"compression_gzip\",\"compression_none\"]"
}
```

Terraform syntax:

```terraform
compression {
  # Configure direct properties listed below.
}
```

<a id="canonical-6027bef41ae447ad8e591aa836d77f5b48eeb472fb9293165e4fe02cb82d4b4a"></a>

## Direct properties — kafka_receiver.compression / 61eebbec6736 / 3

- [compression_default](resources--global_log_receiver--reference--group-003.md#canonical-40198ce2e8dd01ac4cb4dcdf3e0798e461cbb0ab64ff9379087cd512e0ce767c): complete subsection reference.

- [compression_gzip](resources--global_log_receiver--reference--group-003.md#canonical-cded03b78bb87b0f900ac3a5880e5b1fb966f308288fc1fa45f8db01c1ff5acf): complete subsection reference.

- [compression_none](resources--global_log_receiver--reference--group-003.md#canonical-d44eb9f18330ecb74a92e0a29f4b8ca5210dab0d41e30a90d9281a13895d8d7e): complete subsection reference.

<a id="canonical-7c8293ffb077cf7874ad702f48acdf39dfb251c1ff52bdbec78d52b080dd2081"></a>

## Next pages — kafka_receiver.compression / 61eebbec6736 / 4

- [kafka_receiver.compression.compression_default](resources--global_log_receiver--reference--group-003.md#canonical-40198ce2e8dd01ac4cb4dcdf3e0798e461cbb0ab64ff9379087cd512e0ce767c)
- [kafka_receiver.compression.compression_gzip](resources--global_log_receiver--reference--group-003.md#canonical-cded03b78bb87b0f900ac3a5880e5b1fb966f308288fc1fa45f8db01c1ff5acf)
- [kafka_receiver.compression.compression_none](resources--global_log_receiver--reference--group-003.md#canonical-d44eb9f18330ecb74a92e0a29f4b8ca5210dab0d41e30a90d9281a13895d8d7e)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-ee10d6aa6bae7f44f1fe9b11b6baa91894549e346d1cced532d36d362b306e96)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-40198ce2e8dd01ac4cb4dcdf3e0798e461cbb0ab64ff9379087cd512e0ce767c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3dab58e62f71379320a090ea420856fec16a9ebdd7f6fb31c96b3c6e599c0a0a"></a>

## kafka_receiver.compression.compression_default — kafka_receiver.compression.compression_default / ec3bea39761b / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-ee10d6aa6bae7f44f1fe9b11b6baa91894549e346d1cced532d36d362b306e96)
- [kafka_receiver.compression](resources--global_log_receiver--reference--group-003.md#canonical-1ca9494a9e6eb285bbb8a5bab44ed17e05366884a9ea5b51d8bbcb3cee6323be)
- kafka_receiver.compression.compression_default

<a id="canonical-4752d790f746f1f24332c0dab8615b2bc5e79a942dd9072a72e114f6a02790d7"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for compression default.

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
compression_default = {}
```

<a id="canonical-c7dae2b58be01b2dbd4878084735b3cc6bcc6a203480d5c12d6ab043b11c7b5d"></a>

## Direct properties — kafka_receiver.compression.compression_default / ec3bea39761b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b0ab125a6a66dd27f9b2d74f5d86309a6353a876cf965df05a9c9b46543b7a65"></a>

## Next pages — kafka_receiver.compression.compression_default / ec3bea39761b / 4

- [kafka_receiver.compression](resources--global_log_receiver--reference--group-003.md#canonical-1ca9494a9e6eb285bbb8a5bab44ed17e05366884a9ea5b51d8bbcb3cee6323be)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-cded03b78bb87b0f900ac3a5880e5b1fb966f308288fc1fa45f8db01c1ff5acf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a35245514bdc580577ff56a922bd8998fe9a192bccb05e25bc7a63ee18b33388"></a>

## kafka_receiver.compression.compression_gzip — kafka_receiver.compression.compression_gzip / 123d7f17f958 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-ee10d6aa6bae7f44f1fe9b11b6baa91894549e346d1cced532d36d362b306e96)
- [kafka_receiver.compression](resources--global_log_receiver--reference--group-003.md#canonical-1ca9494a9e6eb285bbb8a5bab44ed17e05366884a9ea5b51d8bbcb3cee6323be)
- kafka_receiver.compression.compression_gzip

<a id="canonical-a4804e8fcd71cff07724ba34ecab165e8283dfe6669061d1befd2783dd42d4df"></a>

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
compression_gzip = {}
```

<a id="canonical-afa9dbdb6132508817807720a463887b41fcb5001d51cf2e6f2d37c67b0dd8f8"></a>

## Direct properties — kafka_receiver.compression.compression_gzip / 123d7f17f958 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f4c3f2e1e0e74120255fdb159246d069c425cb7a747fc7aa3a654f44caae8481"></a>

## Next pages — kafka_receiver.compression.compression_gzip / 123d7f17f958 / 4

- [kafka_receiver.compression](resources--global_log_receiver--reference--group-003.md#canonical-1ca9494a9e6eb285bbb8a5bab44ed17e05366884a9ea5b51d8bbcb3cee6323be)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-d44eb9f18330ecb74a92e0a29f4b8ca5210dab0d41e30a90d9281a13895d8d7e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ace9b7c24b8f345dd9f9f8e2707d5a1a6a7197ae941c2ec9a291bea38879ed66"></a>

## kafka_receiver.compression.compression_none — kafka_receiver.compression.compression_none / 6b84a9e920d2 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-ee10d6aa6bae7f44f1fe9b11b6baa91894549e346d1cced532d36d362b306e96)
- [kafka_receiver.compression](resources--global_log_receiver--reference--group-003.md#canonical-1ca9494a9e6eb285bbb8a5bab44ed17e05366884a9ea5b51d8bbcb3cee6323be)
- kafka_receiver.compression.compression_none

<a id="canonical-f0c69de074e62caade164d81978666e1310c86b93b6de25f93f22941afc4bccd"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for compression none.

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
compression_none = {}
```

<a id="canonical-02359c406eaaf4e8c3c639d3c415a84d61fb7966a1191d372cd4923c5d27e84c"></a>

## Direct properties — kafka_receiver.compression.compression_none / 6b84a9e920d2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-46185e2edad82d181f3b281bc87fd24ea5a3f917c860368e77150f521d29fe97"></a>

## Next pages — kafka_receiver.compression.compression_none / 6b84a9e920d2 / 4

- [kafka_receiver.compression](resources--global_log_receiver--reference--group-003.md#canonical-1ca9494a9e6eb285bbb8a5bab44ed17e05366884a9ea5b51d8bbcb3cee6323be)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-7c013a410e875b6e0988c4b653ed8c04de84635c9f23405fb0b1e831599bbd93"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-78c1bce0e4bbb3cca23730c5ff0d57ac2db60fb56a00c85270bdde4f9eeb3b30"></a>

## kafka_receiver.no_tls — kafka_receiver.no_tls / 82eeb89b6420 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-ee10d6aa6bae7f44f1fe9b11b6baa91894549e346d1cced532d36d362b306e96)
- kafka_receiver.no_tls

<a id="canonical-03ff61fef6f9738b90ff4e2da618cf6a3e6c7601c5c2e15af77e1b2244aea015"></a>

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

<a id="canonical-b81fdb9e54a0ae90914ed163fdf40e585f5dc5108807c6ced9654de010daa259"></a>

## Direct properties — kafka_receiver.no_tls / 82eeb89b6420 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4d04acf184c6b3ed17b4e865891de6ddf397d7918c1b5d9ded0aed7d78b91c2a"></a>

## Next pages — kafka_receiver.no_tls / 82eeb89b6420 / 4

- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-ee10d6aa6bae7f44f1fe9b11b6baa91894549e346d1cced532d36d362b306e96)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-a8ec005d981c013ecfd65ce79982faa1c0e75051f53167339dd024dd81d6a2b1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-455fb6d5e22aa1b6ba50e9c42d6a34a9a72dd03abda3979dd17261af2ac3ef33"></a>

## kafka_receiver.use_tls — kafka_receiver.use_tls / ef52f43191b9 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-ee10d6aa6bae7f44f1fe9b11b6baa91894549e346d1cced532d36d362b306e96)
- kafka_receiver.use_tls

<a id="canonical-d9f92407f496c0c5f4f7581665895ec6aa11e6a1cb7cc5052d3626097f9dbdcd"></a>

Type: `"object"`. single nested block, Optional.

TLS Parameters for client connection to the endpoint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_verify_certificate",
    "enable_verify_certificate"),
  validators.ConflictingObjectAttributes("disable_verify_hostname",
    "enable_verify_hostname"),
  validators.ConflictingObjectAttributes("mtls_disabled",
    "mtls_enable"),
  validators.ConflictingObjectAttributes("no_ca",
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
  "x-ves-oneof-field-ca_choice": "[\"no_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-mtls_choice": "[\"mtls_disabled\",\"mtls_enable\"]",
  "x-ves-oneof-field-verify_certificate": "[\"disable_verify_certificate\",\"enable_verify_certificate\"]",
  "x-ves-oneof-field-verify_hostname": "[\"disable_verify_hostname\",\"enable_verify_hostname\"]"
}
```

Terraform syntax:

```terraform
use_tls {
  # Configure direct properties listed below.
}
```

<a id="canonical-5cb0f1cf754244ab35641b6749eff5da0352d38a3b094f322c0733f35f8b69e9"></a>

## Direct properties — kafka_receiver.use_tls / ef52f43191b9 / 3

- [disable_verify_certificate](resources--global_log_receiver--reference--group-003.md#canonical-71225988dd5527e0f4fda1654fa731a1eebc1acbef3bf10a70192cecf0d0e453): complete subsection reference.

- [disable_verify_hostname](resources--global_log_receiver--reference--group-003.md#canonical-df8c467f005c77f5bba165f0a5e9d7af50c2be059be701662389ddd9b925e099): complete subsection reference.

- [enable_verify_certificate](resources--global_log_receiver--reference--group-003.md#canonical-ecad2042adfccb5ef0170242b433838893f760ad429fba7a7af4546f63442250): complete subsection reference.

- [enable_verify_hostname](resources--global_log_receiver--reference--group-003.md#canonical-22a85a01c431d0eaa90e4d4b7c3158654566778ea2fb9182d1867b7d6581bdb4): complete subsection reference.

- [mtls_disabled](resources--global_log_receiver--reference--group-003.md#canonical-aea33c6211872c9414fc8cce517410c4fc57504cd4dbbfafd4a72b9f8d1ab12a): complete subsection reference.

- [mtls_enable](resources--global_log_receiver--reference--group-003.md#canonical-dee0ce635bae35109413e914afdbfe15bab628bd200b6944c532cb33734d17c3): complete subsection reference.

- [no_ca](resources--global_log_receiver--reference--group-003.md#canonical-43ab25ddc807be9c83daa690a31e0049eb110effb37c21cf45dab55832a8490a): complete subsection reference.

<a id="canonical-56598bd4f356d247a04ae1fe4ed15df04e5a2b24858f427d44ef586123a01d17"></a>

<a id="canonical-77e447a26c1a68216073442a811794ff44a7269e404821d3d5e394d748ebca89"></a>

## trusted_ca_url property — kafka_receiver.use_tls / ef52f43191b9 / 4

Type: `"string"`. Optional.

Exclusive with \[no\_ca\] The URL or value for trusted Server CA certificate or certificate chain
Certificates in PEM format including the PEM headers.

Upstream description:

Exclusive with \[no\_ca\] The URL or value for trusted Server CA certificate or certificate chain
Certificates in PEM format including the PEM headers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(131072),
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
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

<a id="canonical-c2315ff645016e197f2a6bfb209d3af10031278b5c0cb1b5d3646ab590fca9e9"></a>

## Next pages — kafka_receiver.use_tls / ef52f43191b9 / 5

- [kafka_receiver.use_tls.disable_verify_certificate](resources--global_log_receiver--reference--group-003.md#canonical-71225988dd5527e0f4fda1654fa731a1eebc1acbef3bf10a70192cecf0d0e453)
- [kafka_receiver.use_tls.disable_verify_hostname](resources--global_log_receiver--reference--group-003.md#canonical-df8c467f005c77f5bba165f0a5e9d7af50c2be059be701662389ddd9b925e099)
- [kafka_receiver.use_tls.enable_verify_certificate](resources--global_log_receiver--reference--group-003.md#canonical-ecad2042adfccb5ef0170242b433838893f760ad429fba7a7af4546f63442250)
- [kafka_receiver.use_tls.enable_verify_hostname](resources--global_log_receiver--reference--group-003.md#canonical-22a85a01c431d0eaa90e4d4b7c3158654566778ea2fb9182d1867b7d6581bdb4)
- [kafka_receiver.use_tls.mtls_disabled](resources--global_log_receiver--reference--group-003.md#canonical-aea33c6211872c9414fc8cce517410c4fc57504cd4dbbfafd4a72b9f8d1ab12a)
- [kafka_receiver.use_tls.mtls_enable](resources--global_log_receiver--reference--group-003.md#canonical-dee0ce635bae35109413e914afdbfe15bab628bd200b6944c532cb33734d17c3)
- [kafka_receiver.use_tls.no_ca](resources--global_log_receiver--reference--group-003.md#canonical-43ab25ddc807be9c83daa690a31e0049eb110effb37c21cf45dab55832a8490a)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-ee10d6aa6bae7f44f1fe9b11b6baa91894549e346d1cced532d36d362b306e96)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-71225988dd5527e0f4fda1654fa731a1eebc1acbef3bf10a70192cecf0d0e453"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-026b1a025d8dcefacfd5afdc1ecf87a9bf0dd2c91c906423a8743ff67325e2b0"></a>

## kafka_receiver.use_tls.disable_verify_certificate — kafka_receiver.use_tls.disable_verify_certificate / 8ffc379fa8f5 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-ee10d6aa6bae7f44f1fe9b11b6baa91894549e346d1cced532d36d362b306e96)
- [kafka_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-a8ec005d981c013ecfd65ce79982faa1c0e75051f53167339dd024dd81d6a2b1)
- kafka_receiver.use_tls.disable_verify_certificate

<a id="canonical-622d450ee9d37690261226de1a1bdf9a24326629f8cb8d9e62df81e0ebe67fc2"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable verify certificate.

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
disable_verify_certificate = {}
```

<a id="canonical-8dd60a84bdc688bbc422400c6a1aed37ab536e75c4e733e8a13653811f2cfb36"></a>

## Direct properties — kafka_receiver.use_tls.disable_verify_certificate / 8ffc379fa8f5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-688ed8071500c317fc6ed26f294a709f6eec1780c7d316271961fbb129cdbbf3"></a>

## Next pages — kafka_receiver.use_tls.disable_verify_certificate / 8ffc379fa8f5 / 4

- [kafka_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-a8ec005d981c013ecfd65ce79982faa1c0e75051f53167339dd024dd81d6a2b1)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-df8c467f005c77f5bba165f0a5e9d7af50c2be059be701662389ddd9b925e099"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cc73412f34492627627194c80b08dd470b4976dcca333e682ff8c7afb592f6ef"></a>

## kafka_receiver.use_tls.disable_verify_hostname — kafka_receiver.use_tls.disable_verify_hostname / a4430e3d6e91 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-ee10d6aa6bae7f44f1fe9b11b6baa91894549e346d1cced532d36d362b306e96)
- [kafka_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-a8ec005d981c013ecfd65ce79982faa1c0e75051f53167339dd024dd81d6a2b1)
- kafka_receiver.use_tls.disable_verify_hostname

<a id="canonical-9f9e6f5c59b29ca295414a6aa0be43e667e62584defa1ff180441a3b0bf8a94b"></a>

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
disable_verify_hostname = {}
```

<a id="canonical-a9271f6596c5b15d32dc7912696b7aeb911f32170724e064f385df047360acd1"></a>

## Direct properties — kafka_receiver.use_tls.disable_verify_hostname / a4430e3d6e91 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8da6004d183fb876c33872542166a2f26b6026f73e63b485eb60c5e9581990eb"></a>

## Next pages — kafka_receiver.use_tls.disable_verify_hostname / a4430e3d6e91 / 4

- [kafka_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-a8ec005d981c013ecfd65ce79982faa1c0e75051f53167339dd024dd81d6a2b1)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-ecad2042adfccb5ef0170242b433838893f760ad429fba7a7af4546f63442250"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-899b8bc30d9979c7ff96f7ec248d2606dffd3fca20506f99cf141700b6dde3f7"></a>

## kafka_receiver.use_tls.enable_verify_certificate — kafka_receiver.use_tls.enable_verify_certificate / a4f2c9d9f270 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-ee10d6aa6bae7f44f1fe9b11b6baa91894549e346d1cced532d36d362b306e96)
- [kafka_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-a8ec005d981c013ecfd65ce79982faa1c0e75051f53167339dd024dd81d6a2b1)
- kafka_receiver.use_tls.enable_verify_certificate

<a id="canonical-d55a7883cfc079b353d05d171a7323cc1b55a1e9fe6182fa748ad7c6991bc1e4"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable verify certificate.

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
enable_verify_certificate = {}
```

<a id="canonical-c81e0ece1fea78861d61914b1d9583f680f304d4ff956a0b1d4be674419876cc"></a>

## Direct properties — kafka_receiver.use_tls.enable_verify_certificate / a4f2c9d9f270 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-edd7bea3319e2f2de8b6b2b439eead0447929bb7b906a7db0a5cb4853eb798ca"></a>

## Next pages — kafka_receiver.use_tls.enable_verify_certificate / a4f2c9d9f270 / 4

- [kafka_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-a8ec005d981c013ecfd65ce79982faa1c0e75051f53167339dd024dd81d6a2b1)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-22a85a01c431d0eaa90e4d4b7c3158654566778ea2fb9182d1867b7d6581bdb4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c5cd346864b5bad264b1316daea962b19616f0dc97ffc5618f297c7fd65c6340"></a>

## kafka_receiver.use_tls.enable_verify_hostname — kafka_receiver.use_tls.enable_verify_hostname / 5313b43f3595 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-ee10d6aa6bae7f44f1fe9b11b6baa91894549e346d1cced532d36d362b306e96)
- [kafka_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-a8ec005d981c013ecfd65ce79982faa1c0e75051f53167339dd024dd81d6a2b1)
- kafka_receiver.use_tls.enable_verify_hostname

<a id="canonical-519dcb2d3132d21c7c140d13b2be75440fd7cf282156788ed20ceb1bdb66ecd8"></a>

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
enable_verify_hostname = {}
```

<a id="canonical-7de9bb88e9a5e76806290af6a8871327c0ec39b9d4f54286d31500482397ce8d"></a>

## Direct properties — kafka_receiver.use_tls.enable_verify_hostname / 5313b43f3595 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-663e5b748fa7dff2f69eccac588251812e357a3783afdf2fc722fd3dd527f6ba"></a>

## Next pages — kafka_receiver.use_tls.enable_verify_hostname / 5313b43f3595 / 4

- [kafka_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-a8ec005d981c013ecfd65ce79982faa1c0e75051f53167339dd024dd81d6a2b1)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-aea33c6211872c9414fc8cce517410c4fc57504cd4dbbfafd4a72b9f8d1ab12a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ce23fe71a0b47797baa097021845e4a34a2d184f444594220d4e1b16a88e8794"></a>

## kafka_receiver.use_tls.mtls_disabled — kafka_receiver.use_tls.mtls_disabled / a5e92a80784c / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-ee10d6aa6bae7f44f1fe9b11b6baa91894549e346d1cced532d36d362b306e96)
- [kafka_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-a8ec005d981c013ecfd65ce79982faa1c0e75051f53167339dd024dd81d6a2b1)
- kafka_receiver.use_tls.mtls_disabled

<a id="canonical-5c609aa0e27278c5165ab781c2e0fe6a8ed7497c44b4f05809e2f31fd5b2a604"></a>

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
mtls_disabled = {}
```

<a id="canonical-8c0d432a1a87df1b66016a9e061bc45c9ce78980b923853963cc74583a50e7c7"></a>

## Direct properties — kafka_receiver.use_tls.mtls_disabled / a5e92a80784c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8813ffd0c7b4fedd02e932482aff4be3df4f9b5b6a565a22168fe79bf0efa3ae"></a>

## Next pages — kafka_receiver.use_tls.mtls_disabled / a5e92a80784c / 4

- [kafka_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-a8ec005d981c013ecfd65ce79982faa1c0e75051f53167339dd024dd81d6a2b1)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-dee0ce635bae35109413e914afdbfe15bab628bd200b6944c532cb33734d17c3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b795a0d51fd76d2354b1ccf1f8fd865a49c24d6321f0377ec8273ae12ed4f1e5"></a>

## kafka_receiver.use_tls.mtls_enable — kafka_receiver.use_tls.mtls_enable / f6610c3fe212 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-ee10d6aa6bae7f44f1fe9b11b6baa91894549e346d1cced532d36d362b306e96)
- [kafka_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-a8ec005d981c013ecfd65ce79982faa1c0e75051f53167339dd024dd81d6a2b1)
- kafka_receiver.use_tls.mtls_enable

<a id="canonical-5d70aed363e7d721744ff9e3c22126058109a9652fdc783c2b2fcec0d78070ce"></a>

Type: `"object"`. single nested block, Optional.

MTLS Client config allows configuration of mTLS client OPTIONS.

Receipt-pinned upstream constraints:

```json
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
mtls_enable {
  # Configure direct properties listed below.
}
```

<a id="canonical-195ee2e4b5fa3f803e0d3688219ac5c484c0e395620aeb40a175016046587ff4"></a>

## Direct properties — kafka_receiver.use_tls.mtls_enable / f6610c3fe212 / 3

<a id="canonical-14b58c60cd8b5965d94c91762fde4dae3997d8423ea30d7b391018f7f55c3ad6"></a>

<a id="canonical-925afa3f0c5e602027d6a638dbcc6a553655325871ce2a0069d9e889af7c17d0"></a>

## certificate property — kafka_receiver.use_tls.mtls_enable / f6610c3fe212 / 4

Type: `"string"`. Optional.

Client certificate is PEM-encoded certificate or certificate-chain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(100, 131072),
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

- [key_url](resources--global_log_receiver--reference--group-003.md#canonical-63103f3084468e43096c91b1bd9d259f5c0be1a8ef8f990d86767f6a4542d110): complete subsection reference.

<a id="canonical-dc29471df266f5d6874df998148901ad29c8e0244f7a7e7491a9a570bee150a1"></a>

## Next pages — kafka_receiver.use_tls.mtls_enable / f6610c3fe212 / 5

- [kafka_receiver.use_tls.mtls_enable.key_url](resources--global_log_receiver--reference--group-003.md#canonical-63103f3084468e43096c91b1bd9d259f5c0be1a8ef8f990d86767f6a4542d110)
- [kafka_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-a8ec005d981c013ecfd65ce79982faa1c0e75051f53167339dd024dd81d6a2b1)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-63103f3084468e43096c91b1bd9d259f5c0be1a8ef8f990d86767f6a4542d110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5e35f79e5d0683af71ba76047161572a4a4c73f0ea59c0abc37f4fdbe9304c09"></a>

## kafka_receiver.use_tls.mtls_enable.key_url — kafka_receiver.use_tls.mtls_enable.key_url / 31c83329b630 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-ee10d6aa6bae7f44f1fe9b11b6baa91894549e346d1cced532d36d362b306e96)
- [kafka_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-a8ec005d981c013ecfd65ce79982faa1c0e75051f53167339dd024dd81d6a2b1)
- [kafka_receiver.use_tls.mtls_enable](resources--global_log_receiver--reference--group-003.md#canonical-dee0ce635bae35109413e914afdbfe15bab628bd200b6944c532cb33734d17c3)
- kafka_receiver.use_tls.mtls_enable.key_url

<a id="canonical-f08e1c06063c1b1ab3116aa5ec243cf853f0c03763742e7b7a4d2acf72cbf1ff"></a>

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
key_url {
  # Configure direct properties listed below.
}
```

<a id="canonical-4c8cc38286c0cd592e4465c7ced06b11658b4448a7ccfdf7da470a384781362c"></a>

## Direct properties — kafka_receiver.use_tls.mtls_enable.key_url / 31c83329b630 / 3

- [blindfold_secret_info](resources--global_log_receiver--reference--group-003.md#canonical-bb6ca7a22b9fbee1f7def61969250e32a5d48f9e5792ebdf84a8620b3d840fcb): complete subsection reference.

- [clear_secret_info](resources--global_log_receiver--reference--group-003.md#canonical-6625525327d4baeb996ad5d2a3e663008a3f7a1ef894284f7262d4bda5ad0085): complete subsection reference.

<a id="canonical-981956a66b7d9c3f7c56660070f8d5c7bef00ffb24f1b6d08ac33d10f96e8dd9"></a>

## Next pages — kafka_receiver.use_tls.mtls_enable.key_url / 31c83329b630 / 4

- [kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info](resources--global_log_receiver--reference--group-003.md#canonical-bb6ca7a22b9fbee1f7def61969250e32a5d48f9e5792ebdf84a8620b3d840fcb)
- [kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info](resources--global_log_receiver--reference--group-003.md#canonical-6625525327d4baeb996ad5d2a3e663008a3f7a1ef894284f7262d4bda5ad0085)
- [kafka_receiver.use_tls.mtls_enable](resources--global_log_receiver--reference--group-003.md#canonical-dee0ce635bae35109413e914afdbfe15bab628bd200b6944c532cb33734d17c3)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-bb6ca7a22b9fbee1f7def61969250e32a5d48f9e5792ebdf84a8620b3d840fcb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-56db7f2b0fc094e2f49e717974d21393c6305f8c4682f1ad5df301327b428b3c"></a>

## kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info — kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info / 33b1c31d4139 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-ee10d6aa6bae7f44f1fe9b11b6baa91894549e346d1cced532d36d362b306e96)
- [kafka_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-a8ec005d981c013ecfd65ce79982faa1c0e75051f53167339dd024dd81d6a2b1)
- [kafka_receiver.use_tls.mtls_enable](resources--global_log_receiver--reference--group-003.md#canonical-dee0ce635bae35109413e914afdbfe15bab628bd200b6944c532cb33734d17c3)
- [kafka_receiver.use_tls.mtls_enable.key_url](resources--global_log_receiver--reference--group-003.md#canonical-63103f3084468e43096c91b1bd9d259f5c0be1a8ef8f990d86767f6a4542d110)
- kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info

<a id="canonical-3c8bdd47a90070b84b0bfe7fa2f4c4d8812185dbf6d5010cbd506ad6dbd77d43"></a>

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

<a id="canonical-02248a6ca8aa1c93caf84cf3b59f2b5ce75502d19174e32863a019ffdf1e1039"></a>

## Direct properties — kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info / 33b1c31d4139 / 3

<a id="canonical-33468ce842df2c2914005d30288a444846c29af8835b32584aa816aad9721666"></a>

<a id="canonical-10f138eb9323aaf2bc34b67b13cf783637ac2fce9db02b5bcdba62ca4d45ceb1"></a>

## decryption_provider property — kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info / 33b1c31d4139 / 4

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

<a id="canonical-7fce96883dd3ad4054f6bdb5d20c8da8f206100af2b2b6fb0ecb8464fb83b713"></a>

<a id="canonical-ba93733c39211de89f47b6e27100588eac8ddfea44611517115816ecbdd695af"></a>

## location property — kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info / 33b1c31d4139 / 5

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

<a id="canonical-d53af5aa8fc959a377913e83e7dd6b44293074fa0f52d1468428a81d8ee96126"></a>

<a id="canonical-9e4800aac3dc888b3fb0f7492d254f8bba68b21c209e879095926e5688256c7e"></a>

## store_provider property — kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info / 33b1c31d4139 / 6

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

<a id="canonical-33fd155f20626963d7728e58350b7d0218f7a0e559d9a404cefe31b5f0f75440"></a>

## Next pages — kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info / 33b1c31d4139 / 7

- [kafka_receiver.use_tls.mtls_enable.key_url](resources--global_log_receiver--reference--group-003.md#canonical-63103f3084468e43096c91b1bd9d259f5c0be1a8ef8f990d86767f6a4542d110)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-6625525327d4baeb996ad5d2a3e663008a3f7a1ef894284f7262d4bda5ad0085"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-43e3ab6d378eff3f5711911d00707354ab172c714e1ebcf92058ef76c4fe945a"></a>

## kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info — kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info / affb6d0a5d84 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-ee10d6aa6bae7f44f1fe9b11b6baa91894549e346d1cced532d36d362b306e96)
- [kafka_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-a8ec005d981c013ecfd65ce79982faa1c0e75051f53167339dd024dd81d6a2b1)
- [kafka_receiver.use_tls.mtls_enable](resources--global_log_receiver--reference--group-003.md#canonical-dee0ce635bae35109413e914afdbfe15bab628bd200b6944c532cb33734d17c3)
- [kafka_receiver.use_tls.mtls_enable.key_url](resources--global_log_receiver--reference--group-003.md#canonical-63103f3084468e43096c91b1bd9d259f5c0be1a8ef8f990d86767f6a4542d110)
- kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info

<a id="canonical-90327d476ce0cc58b3eec63d068a635e5a518554a4b9a538405b2a2fe7e81b38"></a>

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

<a id="canonical-0633150d836cf8d3e2a73b6aa8effef6871a05cbbc22d55dc828d24df5727998"></a>

## Direct properties — kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info / affb6d0a5d84 / 3

<a id="canonical-0c5b98a9e22e9df1a985ea247ab515e31eaecc9b44da05c8ae6b9e8ff2eaff02"></a>

<a id="canonical-f5511eff99434be710fab79d86bd5d5a165551e7d327c0ab6172a656c7f5f00a"></a>

## provider_ref property — kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info / affb6d0a5d84 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-02022316140abddc673f2ef90fee5a1adc99be8b0555c57c63b490140f5e8f9a"></a>

<a id="canonical-57dd2bbe4384d0e2660b1f9da5b6f3ce71adb5ec38d509727a2dff6d3a94d3f0"></a>

## url property — kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info / affb6d0a5d84 / 5

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

<a id="canonical-8d953fe2d79a228a9431201b2d849b9937528879451c2f9931bce9cf737ff4ef"></a>

## Next pages — kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info / affb6d0a5d84 / 6

- [kafka_receiver.use_tls.mtls_enable.key_url](resources--global_log_receiver--reference--group-003.md#canonical-63103f3084468e43096c91b1bd9d259f5c0be1a8ef8f990d86767f6a4542d110)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-43ab25ddc807be9c83daa690a31e0049eb110effb37c21cf45dab55832a8490a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-322d2801c31c3f1be1c236556e97422ffd451e467c4e9b71cb2be438729729be"></a>

## kafka_receiver.use_tls.no_ca — kafka_receiver.use_tls.no_ca / ded7f3d1ba77 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-ee10d6aa6bae7f44f1fe9b11b6baa91894549e346d1cced532d36d362b306e96)
- [kafka_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-a8ec005d981c013ecfd65ce79982faa1c0e75051f53167339dd024dd81d6a2b1)
- kafka_receiver.use_tls.no_ca

<a id="canonical-35669eabd41c9d9492efedafc66c56f448f71b5c604d81f400bc99ee1a0aa8fd"></a>

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
no_ca = {}
```

<a id="canonical-2c69a55ca0909000f122feba081c80d4b4f5ceb035d228146ae30d8aeaaf8c58"></a>

## Direct properties — kafka_receiver.use_tls.no_ca / ded7f3d1ba77 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-49cb6ef3847c426cb1fd1271ffec2483f532f6016e74ff50c052017ce8930398"></a>

## Next pages — kafka_receiver.use_tls.no_ca / ded7f3d1ba77 / 4

- [kafka_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-a8ec005d981c013ecfd65ce79982faa1c0e75051f53167339dd024dd81d6a2b1)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-5094ca81331d2853ad9a7287be56bcb5f0fc071accea0025c24e2362b7697388"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-60b4da58e769fdb756966bd2ff988a556306c9e142abd1c40daae3f7d362b780"></a>

## new_relic_receiver — new_relic_receiver / e1daff2ce6d2 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- new_relic_receiver

<a id="canonical-6c19c494b563c2d60abdc9410439454b737deb2ef492e68a3bff9eec490830d5"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for new relic receiver.

Upstream description:

Configuration for NewRelic endpoint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("eu",
    "us")}
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
  "x-ves-oneof-field-endpoint_choice": "[\"eu\",\"us\"]"
}
```

Terraform syntax:

```terraform
new_relic_receiver {
  # Configure direct properties listed below.
}
```

<a id="canonical-066577dbaa9682f221960f6803ded7ec05824246bb20304106a284c53bbff4b7"></a>

## Direct properties — new_relic_receiver / e1daff2ce6d2 / 3

- [api_key](resources--global_log_receiver--reference--group-003.md#canonical-21d0bb6c98433b5df1272235657524983aa6ee33411d9d00af349ec8edfdee70): complete subsection reference.

- [eu](resources--global_log_receiver--reference--group-003.md#canonical-15fb24f5ade64bedf6352b056a6e3a30ca4612b6cf094bba660ff11d3dba76d0): complete subsection reference.

- [us](resources--global_log_receiver--reference--group-003.md#canonical-466930c53675c81876d4415f40a0eb14bc3eb85b19db568c53ea7c09d2e2aaa8): complete subsection reference.

<a id="canonical-353c9a872e36621202a02c81fbee07f16c23e9c21506234e68e62b461835692c"></a>

## Next pages — new_relic_receiver / e1daff2ce6d2 / 4

- [new_relic_receiver.api_key](resources--global_log_receiver--reference--group-003.md#canonical-21d0bb6c98433b5df1272235657524983aa6ee33411d9d00af349ec8edfdee70)
- [new_relic_receiver.eu](resources--global_log_receiver--reference--group-003.md#canonical-15fb24f5ade64bedf6352b056a6e3a30ca4612b6cf094bba660ff11d3dba76d0)
- [new_relic_receiver.us](resources--global_log_receiver--reference--group-003.md#canonical-466930c53675c81876d4415f40a0eb14bc3eb85b19db568c53ea7c09d2e2aaa8)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-21d0bb6c98433b5df1272235657524983aa6ee33411d9d00af349ec8edfdee70"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c72cba3011964814bece468bc973eb0aa6fb7bb1b5a450dc8affab4550a51141"></a>

## new_relic_receiver.api_key — new_relic_receiver.api_key / 637c37f461b5 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [new_relic_receiver](resources--global_log_receiver--reference--group-003.md#canonical-5094ca81331d2853ad9a7287be56bcb5f0fc071accea0025c24e2362b7697388)
- new_relic_receiver.api_key

<a id="canonical-2b264cb4868e53777a6ab1fcffb99bb127045f1d6a6e14775fb0295454eb0ca2"></a>

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
api_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-5282c4dc4a334702a82b1a6b9c376eff9a8dcd34b366cbeb1b1afc5e3a13c19b"></a>

## Direct properties — new_relic_receiver.api_key / 637c37f461b5 / 3

- [blindfold_secret_info](resources--global_log_receiver--reference--group-003.md#canonical-369580363aeba5685353d0051b30880d4c9d04e04f833b6409a3066ff727d964): complete subsection reference.

- [clear_secret_info](resources--global_log_receiver--reference--group-003.md#canonical-8f5d58ada5cd42d6056e179b2acecdbf9d401e6971de08561d37f982a4c793d9): complete subsection reference.

<a id="canonical-8c25ba787c9e6f151d668f9425756e6aca970de74f6dd99ee782a215a9bb2c27"></a>

## Next pages — new_relic_receiver.api_key / 637c37f461b5 / 4

- [new_relic_receiver.api_key.blindfold_secret_info](resources--global_log_receiver--reference--group-003.md#canonical-369580363aeba5685353d0051b30880d4c9d04e04f833b6409a3066ff727d964)
- [new_relic_receiver.api_key.clear_secret_info](resources--global_log_receiver--reference--group-003.md#canonical-8f5d58ada5cd42d6056e179b2acecdbf9d401e6971de08561d37f982a4c793d9)
- [new_relic_receiver](resources--global_log_receiver--reference--group-003.md#canonical-5094ca81331d2853ad9a7287be56bcb5f0fc071accea0025c24e2362b7697388)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-369580363aeba5685353d0051b30880d4c9d04e04f833b6409a3066ff727d964"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-22de3cffd0ff872761f82474c10111762e4bf3f6ab0dc7e2cbd34123569d962d"></a>

## new_relic_receiver.api_key.blindfold_secret_info — new_relic_receiver.api_key.blindfold_secret_info / c23cdb4cf8f9 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [new_relic_receiver](resources--global_log_receiver--reference--group-003.md#canonical-5094ca81331d2853ad9a7287be56bcb5f0fc071accea0025c24e2362b7697388)
- [new_relic_receiver.api_key](resources--global_log_receiver--reference--group-003.md#canonical-21d0bb6c98433b5df1272235657524983aa6ee33411d9d00af349ec8edfdee70)
- new_relic_receiver.api_key.blindfold_secret_info

<a id="canonical-8e6141350e7bf8fa5e5a2f7aa87012fee73aa3532be69ba70f96a9f0f381be78"></a>

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

<a id="canonical-547a818994a93205ffafcf397841d7d055a8e3e14281653fbbf8773bbb5ba8a4"></a>

## Direct properties — new_relic_receiver.api_key.blindfold_secret_info / c23cdb4cf8f9 / 3

<a id="canonical-7436a86f5294be1e619d49f552098e3e50cf8e7e83ca35b468ed589b8a12d9ac"></a>

<a id="canonical-43bddfad4a7f8660e31e45dd82fc960fc20e122cafed8ef4fdb73c89f7cfd793"></a>

## decryption_provider property — new_relic_receiver.api_key.blindfold_secret_info / c23cdb4cf8f9 / 4

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

<a id="canonical-ab8178a4dcf5ddef2b7fa41f201ede4e2dbad8f5ec16c4c5711e3615b9ec1763"></a>

<a id="canonical-65c84038b3d121d6e1013a0802f9e099513011240b875f58299577533406a889"></a>

## location property — new_relic_receiver.api_key.blindfold_secret_info / c23cdb4cf8f9 / 5

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

<a id="canonical-5f5a9883fa7bfdee2404edc2844d44985b5e52d908d339340c74d5222e2f1597"></a>

<a id="canonical-ff78f5822bfb1eb428e577a21686f04b8c1896a5422b7fb2ecb67276397f493a"></a>

## store_provider property — new_relic_receiver.api_key.blindfold_secret_info / c23cdb4cf8f9 / 6

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

<a id="canonical-419fe95e752ba1cda81d78b06f0619c9ce5d4e4f2274b64b25deba5de2a3fec4"></a>

## Next pages — new_relic_receiver.api_key.blindfold_secret_info / c23cdb4cf8f9 / 7

- [new_relic_receiver.api_key](resources--global_log_receiver--reference--group-003.md#canonical-21d0bb6c98433b5df1272235657524983aa6ee33411d9d00af349ec8edfdee70)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-8f5d58ada5cd42d6056e179b2acecdbf9d401e6971de08561d37f982a4c793d9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b5d43e4afa9b4a210bdb994072aeb89f5da1a6653646b83b1698b30cfb790183"></a>

## new_relic_receiver.api_key.clear_secret_info — new_relic_receiver.api_key.clear_secret_info / 01d0c934942b / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [new_relic_receiver](resources--global_log_receiver--reference--group-003.md#canonical-5094ca81331d2853ad9a7287be56bcb5f0fc071accea0025c24e2362b7697388)
- [new_relic_receiver.api_key](resources--global_log_receiver--reference--group-003.md#canonical-21d0bb6c98433b5df1272235657524983aa6ee33411d9d00af349ec8edfdee70)
- new_relic_receiver.api_key.clear_secret_info

<a id="canonical-0cd49d3c1d5e41e76e12a6715ae208b036a01bbcd0ca577117291011c5a0336e"></a>

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

<a id="canonical-9875daa96a2af9a73effafd9580ef02143c24795afadf79c4223a5af587a1945"></a>

## Direct properties — new_relic_receiver.api_key.clear_secret_info / 01d0c934942b / 3

<a id="canonical-002a0657116d25c20cf59dda534f57e2449d103079e5e75d269dd98fd2620cca"></a>

<a id="canonical-d1b2793d8b6064a8656559a4b299d009917e4b6f07999fd1a51876ac116eff9c"></a>

## provider_ref property — new_relic_receiver.api_key.clear_secret_info / 01d0c934942b / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-9dfd0bb6291cdb0c350756905f5dcd0726d3558d4d763e31d127db1eeaff583e"></a>

<a id="canonical-4f9106990634e37482a531196dbe13d35f0a79b33f8584206571ad33ac7e512c"></a>

## url property — new_relic_receiver.api_key.clear_secret_info / 01d0c934942b / 5

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

<a id="canonical-1ad4fd1ebf5619f18ab856d5046287246ff82edd090ad274de29be416eaaea3f"></a>

## Next pages — new_relic_receiver.api_key.clear_secret_info / 01d0c934942b / 6

- [new_relic_receiver.api_key](resources--global_log_receiver--reference--group-003.md#canonical-21d0bb6c98433b5df1272235657524983aa6ee33411d9d00af349ec8edfdee70)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-15fb24f5ade64bedf6352b056a6e3a30ca4612b6cf094bba660ff11d3dba76d0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1ce4b318efc9c7f9481f89f711ffc07b54e46af8068a4bfc2a1ca7c2b535d66d"></a>

## new_relic_receiver.eu — new_relic_receiver.eu / 4ea14d92b789 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [new_relic_receiver](resources--global_log_receiver--reference--group-003.md#canonical-5094ca81331d2853ad9a7287be56bcb5f0fc071accea0025c24e2362b7697388)
- new_relic_receiver.eu

<a id="canonical-92b8eb502fc087745662d2f24d107c227c7a8376c7e3024df2db8709854c28ad"></a>

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
eu = {}
```

<a id="canonical-0e71566440e800df07c4ecc60bca3edcb9c069b3fc00e75df92a1c80d42acc2e"></a>

## Direct properties — new_relic_receiver.eu / 4ea14d92b789 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-eb872bd8c3241773b237ca2d5eb11fe86ff4bb235710ca31575f67e0cbb44edb"></a>

## Next pages — new_relic_receiver.eu / 4ea14d92b789 / 4

- [new_relic_receiver](resources--global_log_receiver--reference--group-003.md#canonical-5094ca81331d2853ad9a7287be56bcb5f0fc071accea0025c24e2362b7697388)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-466930c53675c81876d4415f40a0eb14bc3eb85b19db568c53ea7c09d2e2aaa8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-627d6905520d80408f6108af9b2fa33fbb52094370f28d7707dfe441b64bc4c2"></a>

## new_relic_receiver.us — new_relic_receiver.us / 035a8aa7bea8 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [new_relic_receiver](resources--global_log_receiver--reference--group-003.md#canonical-5094ca81331d2853ad9a7287be56bcb5f0fc071accea0025c24e2362b7697388)
- new_relic_receiver.us

<a id="canonical-4351f3cb11d29176bb6e65617ebab6c3c341f487c0aceebb2576538ef01b94e0"></a>

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
us = {}
```

<a id="canonical-7809c9fb7cfeae94e29486d28e1d5b11ecde3b89fbe85e883a39a46ae12dc9ad"></a>

## Direct properties — new_relic_receiver.us / 035a8aa7bea8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e8e306b07d799e1153da0d7c3df469e2cdbb2b9d4e1e4da1d2d2403d5a6179d7"></a>

## Next pages — new_relic_receiver.us / 035a8aa7bea8 / 4

- [new_relic_receiver](resources--global_log_receiver--reference--group-003.md#canonical-5094ca81331d2853ad9a7287be56bcb5f0fc071accea0025c24e2362b7697388)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-94ac97fe81951dfbc63beb0bdb84d4b3cfbea1343ac31d194e4cfe249e881911"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-776a7e2e61d8885d568c38725280e2a8762b05faecec1af2d72cb30b755b8ee4"></a>

## ns_all — ns_all / 660689d808fc / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- ns_all

<a id="canonical-c27f9e5c2cba312b5473ae4b3229b22d564f8299e9f7c4e79e35919e2517a660"></a>

Type: `["object", {}]`. Optional.

\[OneOf: ns\_all, ns\_current, ns\_list\] Enable this option

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

- [ns_all](resources--global_log_receiver--reference--group-003.md#canonical-c27f9e5c2cba312b5473ae4b3229b22d564f8299e9f7c4e79e35919e2517a660)
- [ns_current](resources--global_log_receiver--reference--group-003.md#canonical-31ea42d689d2504cf6aa86775e94506b829f188ffb8cb066ef06b562e932d6c9)
- [ns_list](resources--global_log_receiver--reference--group-003.md#canonical-3d8776edceb7fd01ca4ca000da91ef806f6e280746404127792a14e1d813cd90)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
ns_all = {}
```

<a id="canonical-15587c8e7e16b26861b6131a975f2bc6fb018607d19f2a28c2209131774d7f77"></a>

## Direct properties — ns_all / 660689d808fc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-785fa15bab3225ee6ada6fab68c51312df8ba301088759e5fba9dbdae71a7286"></a>

## Next pages — ns_all / 660689d808fc / 4

- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-5d4a9b8030a880ba327e5064918fb59a3565b4a15a507da773657ce9558b8eaf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-74b9c848480216eb9981198c1afe4bbe88231e1fb09ee6023552407b8c77bd19"></a>

## ns_current — ns_current / 5f2071d17bb0 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- ns_current

<a id="canonical-31ea42d689d2504cf6aa86775e94506b829f188ffb8cb066ef06b562e932d6c9"></a>

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
ns_current = {}
```

<a id="canonical-77987d19716760a9dc4751d7822f4a64c36b9e086d130a80afa34100ad1b9840"></a>

## Direct properties — ns_current / 5f2071d17bb0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e13fc1a9714f76f8097443d3e1e4811656475f12a8fc291955b6c6e2c0af9b42"></a>

## Next pages — ns_current / 5f2071d17bb0 / 4

- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-7b1e2acca773fcf6415527dd41a12a76c20c635cd71ada5012b1e73e227f97e3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3755784c1fd5e453faccbe2e4297c8599ba41a43a4fa1b9299fcea6c5e2a5618"></a>

## ns_list — ns_list / 42e031737151 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- ns_list

<a id="canonical-3d8776edceb7fd01ca4ca000da91ef806f6e280746404127792a14e1d813cd90"></a>

Type: `"object"`. single nested block, Optional.

Namespace List. Namespace List.

Upstream description:

Namespace List.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("namespaces")}
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
ns_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-e91088b74555ad4d7edef1766a97ac467ab0b4cef85b9cd3a05f4a4a14b52d40"></a>

## Direct properties — ns_list / 42e031737151 / 3

<a id="canonical-f57c3d4696e0f5d54322ba7e0b3c80ca4656f4f61894682af2e9107c56cc1c21"></a>

<a id="canonical-c77a990788e51964cbb67270b15b1b2ded656410253ec6af906900047d6e50cd"></a>

## namespaces property — ns_list / 42e031737151 / 4

Type: `["list", "string"]`. Optional.

Namespaces. List of namespaces to stream logs for.

Upstream description:

List of namespaces to stream logs for.

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

<a id="canonical-583acf70e073c5845e5a775088462c406c1edc805fc5d996a9d5348c2bc6ba7a"></a>

## Next pages — ns_list / 42e031737151 / 5

- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-afb382cb3e4d5e62efc324cc347c6ce43343131b41944fdb2dca84647cc5f8be"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-421bc28f17c4c0734150d1358983ad978ff1221532809f2383ec87a0af6dcb92"></a>

## qradar_receiver — qradar_receiver / 50afae6ef7af / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- qradar_receiver

<a id="canonical-fb8d7c7d87103f6fdf99293b8e667c3be11670787d98f3d89f9641bb9f2b4f5c"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for qradar receiver.

Upstream description:

Configuration for IBM QRadar endpoint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("uri"),
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
qradar_receiver {
  # Configure direct properties listed below.
}
```

<a id="canonical-f71c9d331d72f5c016028686aff597dcd0abba7ac8124bddfc25c252de817512"></a>

## Direct properties — qradar_receiver / 50afae6ef7af / 3

- [batch](resources--global_log_receiver--reference--group-003.md#canonical-8644046a1ddae8615df5e57f8e1eb0a9fa845c6109470f0d69963f3864c3c285): complete subsection reference.

- [compression](resources--global_log_receiver--reference--group-003.md#canonical-f176826683b08da7e61ab4029223fc00594d2cfa02c7e455d82115b80bf619fe): complete subsection reference.

- [no_tls](resources--global_log_receiver--reference--group-003.md#canonical-3f30311152f58171cbeaf6297ea7cb93148fa6ae2f8242a0259a3deeb1f59e83): complete subsection reference.

<a id="canonical-05eee3e05588c7e0b159e040592ca066a14aa2f3800f5a898d291c7ad76e60a7"></a>

<a id="canonical-28a891ee51fc7b1b2ec380da10c2b14d51d1e3c6b04bc422443467b703f0d389"></a>

## uri property — qradar_receiver / 50afae6ef7af / 4

Type: `"string"`. Optional.

Log Source Collector URL is the URL of the IBM QRadar Log Source Collector to send logs to,.

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
    "format": "uri",
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
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [use_tls](resources--global_log_receiver--reference--group-003.md#canonical-8380b97319f78370786be39917d8af086bd5a1781bbfcc2f916aa7c688f7a19f): complete subsection reference.

<a id="canonical-d2a6c1b0ed1f88bff4159475124e27c70345277af5cacbf8fb0dd2ee116bd57e"></a>

## Next pages — qradar_receiver / 50afae6ef7af / 5

- [qradar_receiver.batch](resources--global_log_receiver--reference--group-003.md#canonical-8644046a1ddae8615df5e57f8e1eb0a9fa845c6109470f0d69963f3864c3c285)
- [qradar_receiver.compression](resources--global_log_receiver--reference--group-003.md#canonical-f176826683b08da7e61ab4029223fc00594d2cfa02c7e455d82115b80bf619fe)
- [qradar_receiver.no_tls](resources--global_log_receiver--reference--group-003.md#canonical-3f30311152f58171cbeaf6297ea7cb93148fa6ae2f8242a0259a3deeb1f59e83)
- [qradar_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-8380b97319f78370786be39917d8af086bd5a1781bbfcc2f916aa7c688f7a19f)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-8644046a1ddae8615df5e57f8e1eb0a9fa845c6109470f0d69963f3864c3c285"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9be3b7990fd7515a0739f8a726eec4bafb8fdb5dc0be10bbe7d1266dfbf3f43e"></a>

## qradar_receiver.batch — qradar_receiver.batch / bced96f481d8 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [qradar_receiver](resources--global_log_receiver--reference--group-003.md#canonical-afb382cb3e4d5e62efc324cc347c6ce43343131b41944fdb2dca84647cc5f8be)
- qradar_receiver.batch

<a id="canonical-89f6b3ae8d775713f1b99e617e71009649aedc3d46a63069e0dafd1486d82fbc"></a>

Type: `"object"`. single nested block, Optional.

Batch OPTIONS allow tuning for how batches of logs are sent to an endpoint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("max_bytes",
    "max_bytes_disabled"),
  validators.ConflictingObjectAttributes("max_events",
    "max_events_disabled"),
  validators.ConflictingObjectAttributes("timeout_seconds",
    "timeout_seconds_default")}
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
  "x-ves-oneof-field-batch_bytes": "[\"max_bytes\",\"max_bytes_disabled\"]",
  "x-ves-oneof-field-batch_events": "[\"max_events\",\"max_events_disabled\"]",
  "x-ves-oneof-field-batch_timeout": "[\"timeout_seconds\",\"timeout_seconds_default\"]"
}
```

Terraform syntax:

```terraform
batch {
  # Configure direct properties listed below.
}
```

<a id="canonical-cc3cff2dd678ccc6395a3578fc76a9e4740d8736f76fee35cba923016e3f0d91"></a>

## Direct properties — qradar_receiver.batch / bced96f481d8 / 3

<a id="canonical-ade6c2fd590b14093a78241888f0255dbc2a7cffa1309168164075e518a1ba03"></a>

<a id="canonical-0ee4b9c493fa4f251014a264aff7d62dd31223fbf62f7f811a039328cd0e06b6"></a>

## max_bytes property — qradar_receiver.batch / bced96f481d8 / 4

Type: `"number"`. Optional.

Exclusive with \[max\_bytes\_disabled\] Send batch to endpoint after the batch is equal to or larger
than this many bytes.

Upstream description:

Exclusive with \[max\_bytes\_disabled\] Send batch to endpoint after the batch is equal to or larger
than this many bytes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(4096, 10485760),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 10485760,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 4096
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "4096",
    "ves.io.schema.rules.uint32.lte": "10485760"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "4096",
    "ves.io.schema.rules.uint32.lte": "10485760"
  }
}
```

- [max_bytes_disabled](resources--global_log_receiver--reference--group-003.md#canonical-fca0d11650f7f828a284953fab7593f20526a12009e9b89f1c3a762a0b8fb2bf): complete subsection reference.

<a id="canonical-290776540be479258eeb7034db60b29168db272d9e8ec00c7700d24a7ad1b00b"></a>

<a id="canonical-36bb00d526d12299c00ec940660d37f81c114fbfd0670b9f97484aa04fffbc9e"></a>

## max_events property — qradar_receiver.batch / bced96f481d8 / 5

Type: `"number"`. Optional.

Exclusive with \[max\_events\_disabled\] Send batch to endpoint after this many log messages are in
the batch.

Upstream description:

Exclusive with \[max\_events\_disabled\] Send batch to endpoint after this many log messages are in
the batch.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(32, 2000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 32
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "32",
    "ves.io.schema.rules.uint32.lte": "2000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "32",
    "ves.io.schema.rules.uint32.lte": "2000"
  }
}
```

- [max_events_disabled](resources--global_log_receiver--reference--group-003.md#canonical-2d6df74c58f84d9375e482a4392601a043efcee53a8a41c17e1de703554516cd): complete subsection reference.

<a id="canonical-bd39b1eec30cf059139000ee20d11b29a4e50977ef09abaddb5aa593d1b4412a"></a>

<a id="canonical-1c6e634bb187ee75b0d400ec30d87fd1b751b0b9e66d60812e913f6a941798b0"></a>

## timeout_seconds property — qradar_receiver.batch / bced96f481d8 / 6

Type: `"string"`. Optional.

Exclusive with \[timeout\_seconds\_default\] Send batch to the endpoint after this many seconds.

Upstream description:

Exclusive with \[timeout\_seconds\_default\] Send batch to the endpoint after this many seconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "uint64",
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
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  }
}
```

- [timeout_seconds_default](resources--global_log_receiver--reference--group-003.md#canonical-29066e6ded41972a06b7f6b6c428e251ee43c9f2a82304a4fa06d3f234c377f1): complete subsection reference.

<a id="canonical-2dd3cfb952269d77e5ee56d48f07d635ec0b35249c22e88669ca6976938718ed"></a>

## Next pages — qradar_receiver.batch / bced96f481d8 / 7

- [qradar_receiver.batch.max_bytes_disabled](resources--global_log_receiver--reference--group-003.md#canonical-fca0d11650f7f828a284953fab7593f20526a12009e9b89f1c3a762a0b8fb2bf)
- [qradar_receiver.batch.max_events_disabled](resources--global_log_receiver--reference--group-003.md#canonical-2d6df74c58f84d9375e482a4392601a043efcee53a8a41c17e1de703554516cd)
- [qradar_receiver.batch.timeout_seconds_default](resources--global_log_receiver--reference--group-003.md#canonical-29066e6ded41972a06b7f6b6c428e251ee43c9f2a82304a4fa06d3f234c377f1)
- [qradar_receiver](resources--global_log_receiver--reference--group-003.md#canonical-afb382cb3e4d5e62efc324cc347c6ce43343131b41944fdb2dca84647cc5f8be)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-fca0d11650f7f828a284953fab7593f20526a12009e9b89f1c3a762a0b8fb2bf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8bd3d0b8397c9ef0b6066f59d3a43de7660e68a36cc99ff4fcc0c3b388835ca6"></a>

## qradar_receiver.batch.max_bytes_disabled — qradar_receiver.batch.max_bytes_disabled / cfd7ec9511c4 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [qradar_receiver](resources--global_log_receiver--reference--group-003.md#canonical-afb382cb3e4d5e62efc324cc347c6ce43343131b41944fdb2dca84647cc5f8be)
- [qradar_receiver.batch](resources--global_log_receiver--reference--group-003.md#canonical-8644046a1ddae8615df5e57f8e1eb0a9fa845c6109470f0d69963f3864c3c285)
- qradar_receiver.batch.max_bytes_disabled

<a id="canonical-9324f6ee38dd37b5ba7cb5b8d9d2b48158493aefa4d651abf20d74d3a580d058"></a>

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
max_bytes_disabled = {}
```

<a id="canonical-39a55f5447458e3bf093f62f985b9c581ae12049e760abe2913168f89ab708a6"></a>

## Direct properties — qradar_receiver.batch.max_bytes_disabled / cfd7ec9511c4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a605690f6f2ee08ed96f4980db02ea7694be99e0b4d6cb0c8a98d870dad4d743"></a>

## Next pages — qradar_receiver.batch.max_bytes_disabled / cfd7ec9511c4 / 4

- [qradar_receiver.batch](resources--global_log_receiver--reference--group-003.md#canonical-8644046a1ddae8615df5e57f8e1eb0a9fa845c6109470f0d69963f3864c3c285)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-2d6df74c58f84d9375e482a4392601a043efcee53a8a41c17e1de703554516cd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f4a848a8b3ab9d8a74793efb80b1cbdd29ae3d190f1fc6135b9dfd7be5db727b"></a>

## qradar_receiver.batch.max_events_disabled — qradar_receiver.batch.max_events_disabled / bc1158c26222 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [qradar_receiver](resources--global_log_receiver--reference--group-003.md#canonical-afb382cb3e4d5e62efc324cc347c6ce43343131b41944fdb2dca84647cc5f8be)
- [qradar_receiver.batch](resources--global_log_receiver--reference--group-003.md#canonical-8644046a1ddae8615df5e57f8e1eb0a9fa845c6109470f0d69963f3864c3c285)
- qradar_receiver.batch.max_events_disabled

<a id="canonical-9f266a5bc481b831e319e610a657ef7a9a80654212ae2b0f17b8deab69fe3e0f"></a>

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
max_events_disabled = {}
```

<a id="canonical-1aeea5e0247bfd400ffe0c62836ff69354d0efc56cffa891cd55a7564a566898"></a>

## Direct properties — qradar_receiver.batch.max_events_disabled / bc1158c26222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1cdc25fb2aa70197a31f1e5581d2ae17657d5b8e2279c2c8459d8f1328486024"></a>

## Next pages — qradar_receiver.batch.max_events_disabled / bc1158c26222 / 4

- [qradar_receiver.batch](resources--global_log_receiver--reference--group-003.md#canonical-8644046a1ddae8615df5e57f8e1eb0a9fa845c6109470f0d69963f3864c3c285)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-29066e6ded41972a06b7f6b6c428e251ee43c9f2a82304a4fa06d3f234c377f1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1592c21e80fb10529e436a94ec59e23c0daf75fed5d7ad0532f2d492ee9cb693"></a>

## qradar_receiver.batch.timeout_seconds_default — qradar_receiver.batch.timeout_seconds_default / 6c59e5cb5e99 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [qradar_receiver](resources--global_log_receiver--reference--group-003.md#canonical-afb382cb3e4d5e62efc324cc347c6ce43343131b41944fdb2dca84647cc5f8be)
- [qradar_receiver.batch](resources--global_log_receiver--reference--group-003.md#canonical-8644046a1ddae8615df5e57f8e1eb0a9fa845c6109470f0d69963f3864c3c285)
- qradar_receiver.batch.timeout_seconds_default

<a id="canonical-6ef395f56608dfdf7aee92b9de073781e9d241f42c1dec2769d9086c162bc15e"></a>

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
timeout_seconds_default = {}
```

<a id="canonical-5a2449f9778f356012c972ca3bd99359c1c18c23d43025f0bbe98c3936260411"></a>

## Direct properties — qradar_receiver.batch.timeout_seconds_default / 6c59e5cb5e99 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ead91c01359af5bb83c5db514839213eee9b26bb309abe46333c217d9c570fab"></a>

## Next pages — qradar_receiver.batch.timeout_seconds_default / 6c59e5cb5e99 / 4

- [qradar_receiver.batch](resources--global_log_receiver--reference--group-003.md#canonical-8644046a1ddae8615df5e57f8e1eb0a9fa845c6109470f0d69963f3864c3c285)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-f176826683b08da7e61ab4029223fc00594d2cfa02c7e455d82115b80bf619fe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8dd37fbe079f6543e656b2838f0a4d179daf2193149e9ebbfe5ddfbcdba01c80"></a>

## qradar_receiver.compression — qradar_receiver.compression / 2d9aaef99d15 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [qradar_receiver](resources--global_log_receiver--reference--group-003.md#canonical-afb382cb3e4d5e62efc324cc347c6ce43343131b41944fdb2dca84647cc5f8be)
- qradar_receiver.compression

<a id="canonical-306dcb2624983ce2e01787c91252caabccef8af00c84cd570d5b93af0e024a7d"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for compression.

Upstream description:

Compression Type.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("compression_default",
    "compression_gzip"),
  validators.ConflictingObjectAttributes("compression_default",
    "compression_none"),
  validators.ConflictingObjectAttributes("compression_gzip",
    "compression_none")}
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
  "x-ves-oneof-field-compression_choice": "[\"compression_default\",\"compression_gzip\",\"compression_none\"]"
}
```

Terraform syntax:

```terraform
compression {
  # Configure direct properties listed below.
}
```

<a id="canonical-9b0d40fd7be471280707768b2de88fa4e804b5377f4702348f7c71e4c6a37987"></a>

## Direct properties — qradar_receiver.compression / 2d9aaef99d15 / 3

- [compression_default](resources--global_log_receiver--reference--group-003.md#canonical-8b123e17517f29da933f8a310adb31c4e8adfa3c78d395b674f74a1fdec4e8d4): complete subsection reference.

- [compression_gzip](resources--global_log_receiver--reference--group-003.md#canonical-70f8cb6df9de0d48c6b5c76fff93eca3eadd79f99950cd7c61cc8c68c9553fc5): complete subsection reference.

- [compression_none](resources--global_log_receiver--reference--group-003.md#canonical-6876125673ca352bf3b2f75a073bdad506b0499c44d5eaab6073b9c001ea9802): complete subsection reference.

<a id="canonical-2cf6475dcddfef2c31891dcb7f07ed3718cd55ac12d48fd0b5c1e590e6f7f4c8"></a>

## Next pages — qradar_receiver.compression / 2d9aaef99d15 / 4

- [qradar_receiver.compression.compression_default](resources--global_log_receiver--reference--group-003.md#canonical-8b123e17517f29da933f8a310adb31c4e8adfa3c78d395b674f74a1fdec4e8d4)
- [qradar_receiver.compression.compression_gzip](resources--global_log_receiver--reference--group-003.md#canonical-70f8cb6df9de0d48c6b5c76fff93eca3eadd79f99950cd7c61cc8c68c9553fc5)
- [qradar_receiver.compression.compression_none](resources--global_log_receiver--reference--group-003.md#canonical-6876125673ca352bf3b2f75a073bdad506b0499c44d5eaab6073b9c001ea9802)
- [qradar_receiver](resources--global_log_receiver--reference--group-003.md#canonical-afb382cb3e4d5e62efc324cc347c6ce43343131b41944fdb2dca84647cc5f8be)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-8b123e17517f29da933f8a310adb31c4e8adfa3c78d395b674f74a1fdec4e8d4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c66b1a7fb2acaaa84d74462356e80ba56894e1f32906667fa12e5c900de675cc"></a>

## qradar_receiver.compression.compression_default — qradar_receiver.compression.compression_default / dd5247249e6b / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [qradar_receiver](resources--global_log_receiver--reference--group-003.md#canonical-afb382cb3e4d5e62efc324cc347c6ce43343131b41944fdb2dca84647cc5f8be)
- [qradar_receiver.compression](resources--global_log_receiver--reference--group-003.md#canonical-f176826683b08da7e61ab4029223fc00594d2cfa02c7e455d82115b80bf619fe)
- qradar_receiver.compression.compression_default

<a id="canonical-aac82c3a70da25647c2bd2e9fe1a425e0bf70ce0697eee59589fd10202a9106e"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for compression default.

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
compression_default = {}
```

<a id="canonical-279c1b491addbb80a8995086a3986e2e00188c60aa259c89f98697bb93761373"></a>

## Direct properties — qradar_receiver.compression.compression_default / dd5247249e6b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-dd6309172a4d5b49df619846ce886ce8d0b434b4f163f82869aeef1b3649e9ad"></a>

## Next pages — qradar_receiver.compression.compression_default / dd5247249e6b / 4

- [qradar_receiver.compression](resources--global_log_receiver--reference--group-003.md#canonical-f176826683b08da7e61ab4029223fc00594d2cfa02c7e455d82115b80bf619fe)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-70f8cb6df9de0d48c6b5c76fff93eca3eadd79f99950cd7c61cc8c68c9553fc5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8b451222d52bbed000ba2969abc01714827add99dfa9da014a686ed0a609eb3d"></a>

## qradar_receiver.compression.compression_gzip — qradar_receiver.compression.compression_gzip / 3297baebf3ee / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [qradar_receiver](resources--global_log_receiver--reference--group-003.md#canonical-afb382cb3e4d5e62efc324cc347c6ce43343131b41944fdb2dca84647cc5f8be)
- [qradar_receiver.compression](resources--global_log_receiver--reference--group-003.md#canonical-f176826683b08da7e61ab4029223fc00594d2cfa02c7e455d82115b80bf619fe)
- qradar_receiver.compression.compression_gzip

<a id="canonical-c190d0fd5813fb55ac422094ef2acb21bfce6aae9fc4565b2f743ce5a2f04ca0"></a>

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
compression_gzip = {}
```

<a id="canonical-8200f36d907860fb2e26badd5cc7c12607fdcc8baa1771bc44b6ca4c9cebe1fd"></a>

## Direct properties — qradar_receiver.compression.compression_gzip / 3297baebf3ee / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b51b711651946df39b0c62a093476606e7816f6855a8acc0c22c093bfc15551a"></a>

## Next pages — qradar_receiver.compression.compression_gzip / 3297baebf3ee / 4

- [qradar_receiver.compression](resources--global_log_receiver--reference--group-003.md#canonical-f176826683b08da7e61ab4029223fc00594d2cfa02c7e455d82115b80bf619fe)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-6876125673ca352bf3b2f75a073bdad506b0499c44d5eaab6073b9c001ea9802"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3831d244f211beead3f02866622219a4c5c2dfdb26f7aa0eb2744506247edf60"></a>

## qradar_receiver.compression.compression_none — qradar_receiver.compression.compression_none / e88d6ad3e452 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [qradar_receiver](resources--global_log_receiver--reference--group-003.md#canonical-afb382cb3e4d5e62efc324cc347c6ce43343131b41944fdb2dca84647cc5f8be)
- [qradar_receiver.compression](resources--global_log_receiver--reference--group-003.md#canonical-f176826683b08da7e61ab4029223fc00594d2cfa02c7e455d82115b80bf619fe)
- qradar_receiver.compression.compression_none

<a id="canonical-02213be623e582aa51b001c964024a6d375210517c2bf508022d2dc8c54fa766"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for compression none.

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
compression_none = {}
```

<a id="canonical-a87973358277361025cb5c195ecd17b77ab7194baaeb7c1a216aa5830f1a3a34"></a>

## Direct properties — qradar_receiver.compression.compression_none / e88d6ad3e452 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7f09ccb8e2e3e54555b44c6c786e5e31bc83aaeca507143b49306d15b20fde9b"></a>

## Next pages — qradar_receiver.compression.compression_none / e88d6ad3e452 / 4

- [qradar_receiver.compression](resources--global_log_receiver--reference--group-003.md#canonical-f176826683b08da7e61ab4029223fc00594d2cfa02c7e455d82115b80bf619fe)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-3f30311152f58171cbeaf6297ea7cb93148fa6ae2f8242a0259a3deeb1f59e83"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e346253114eb2d38cd6642d6077e8842ac8bba91e5a6c3f1157390038465d69f"></a>

## qradar_receiver.no_tls — qradar_receiver.no_tls / 470779748251 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [qradar_receiver](resources--global_log_receiver--reference--group-003.md#canonical-afb382cb3e4d5e62efc324cc347c6ce43343131b41944fdb2dca84647cc5f8be)
- qradar_receiver.no_tls

<a id="canonical-8545974d89fa8cd5ef85d928bf99dc2f6aabc862d17d94baa22402156e03f81d"></a>

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

<a id="canonical-a7dea9a31b84f4ea05cac55a86a723e79c8fa19ae4879ee045bc37a447b3275b"></a>

## Direct properties — qradar_receiver.no_tls / 470779748251 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-87939e5a6af7230cacd81bc7635fbc130730498612818c118e6130fbd43525b8"></a>

## Next pages — qradar_receiver.no_tls / 470779748251 / 4

- [qradar_receiver](resources--global_log_receiver--reference--group-003.md#canonical-afb382cb3e4d5e62efc324cc347c6ce43343131b41944fdb2dca84647cc5f8be)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-8380b97319f78370786be39917d8af086bd5a1781bbfcc2f916aa7c688f7a19f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f30ec7694b69492bc31639fbb405c0f9e4f18ff8b40df6f4b6615eaaa8dcbf06"></a>

## qradar_receiver.use_tls — qradar_receiver.use_tls / 1736b1abd17f / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [qradar_receiver](resources--global_log_receiver--reference--group-003.md#canonical-afb382cb3e4d5e62efc324cc347c6ce43343131b41944fdb2dca84647cc5f8be)
- qradar_receiver.use_tls

<a id="canonical-2fceb74b19c331681087b72070381ca9c81817ea640f3cd83b5f6c86d55dcb96"></a>

Type: `"object"`. single nested block, Optional.

TLS Parameters for client connection to the endpoint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_verify_certificate",
    "enable_verify_certificate"),
  validators.ConflictingObjectAttributes("disable_verify_hostname",
    "enable_verify_hostname"),
  validators.ConflictingObjectAttributes("mtls_disabled",
    "mtls_enable"),
  validators.ConflictingObjectAttributes("no_ca",
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
  "x-ves-oneof-field-ca_choice": "[\"no_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-mtls_choice": "[\"mtls_disabled\",\"mtls_enable\"]",
  "x-ves-oneof-field-verify_certificate": "[\"disable_verify_certificate\",\"enable_verify_certificate\"]",
  "x-ves-oneof-field-verify_hostname": "[\"disable_verify_hostname\",\"enable_verify_hostname\"]"
}
```

Terraform syntax:

```terraform
use_tls {
  # Configure direct properties listed below.
}
```

<a id="canonical-c8976bc5e46c99a331be8cd2a0ba34955128d5fb467abacee3279d78d0f60d7a"></a>

## Direct properties — qradar_receiver.use_tls / 1736b1abd17f / 3

- [disable_verify_certificate](resources--global_log_receiver--reference--group-003.md#canonical-bf63dfe4f551a495cbafb2f5b16298cb2ba2eb1fe0c451f000bb1594d90419ec): complete subsection reference.

- [disable_verify_hostname](resources--global_log_receiver--reference--group-003.md#canonical-12b9a4a82114c8022aad59dbc7bc29b87c772b4f425ff55df3c29d6b19ef4307): complete subsection reference.

- [enable_verify_certificate](resources--global_log_receiver--reference--group-004.md#canonical-8f0e826f7ca6c66e7b10c83c7bcdbac90dd8db3efb75cb9b218075c3fd13199e): complete subsection reference.

- [enable_verify_hostname](resources--global_log_receiver--reference--group-004.md#canonical-9a9cbfbc2274a4350e1c1d08d6893c65afad9000072e96717ca67608fdae85cc): complete subsection reference.

- [mtls_disabled](resources--global_log_receiver--reference--group-004.md#canonical-2dcc0b6a06e724b7f8ba24996ce7d4f2fe1c119b3a26673a9fdb241c65221c9b): complete subsection reference.

- [mtls_enable](resources--global_log_receiver--reference--group-004.md#canonical-c0ff8cfb35ba02bb45894dd440a666929540a78e3b2fc206dcb62f94a888d75c): complete subsection reference.

- [no_ca](resources--global_log_receiver--reference--group-004.md#canonical-767399ec1e9d93857cdca75e04dde228050045f3486bf56ce091add95778baff): complete subsection reference.

<a id="canonical-e36baa59a1e43aaa6c189a3198b82ab809d3cd6b86cfed48f28f8a7a25f7b3c5"></a>

<a id="canonical-400ca5742290a9bf1edea41a97e5c579977dad8764c693eff0807818004b6f4c"></a>

## trusted_ca_url property — qradar_receiver.use_tls / 1736b1abd17f / 4

Type: `"string"`. Optional.

Exclusive with \[no\_ca\] The URL or value for trusted Server CA certificate or certificate chain
Certificates in PEM format including the PEM headers.

Upstream description:

Exclusive with \[no\_ca\] The URL or value for trusted Server CA certificate or certificate chain
Certificates in PEM format including the PEM headers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(131072),
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
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

<a id="canonical-5f6cd25a0d6d1423de069f019a8cd53b78114eb2e441627f7ee0705521875ee2"></a>

## Next pages — qradar_receiver.use_tls / 1736b1abd17f / 5

- [qradar_receiver.use_tls.disable_verify_certificate](resources--global_log_receiver--reference--group-003.md#canonical-bf63dfe4f551a495cbafb2f5b16298cb2ba2eb1fe0c451f000bb1594d90419ec)
- [qradar_receiver.use_tls.disable_verify_hostname](resources--global_log_receiver--reference--group-003.md#canonical-12b9a4a82114c8022aad59dbc7bc29b87c772b4f425ff55df3c29d6b19ef4307)
- [qradar_receiver.use_tls.enable_verify_certificate](resources--global_log_receiver--reference--group-004.md#canonical-8f0e826f7ca6c66e7b10c83c7bcdbac90dd8db3efb75cb9b218075c3fd13199e)
- [qradar_receiver.use_tls.enable_verify_hostname](resources--global_log_receiver--reference--group-004.md#canonical-9a9cbfbc2274a4350e1c1d08d6893c65afad9000072e96717ca67608fdae85cc)
- [qradar_receiver.use_tls.mtls_disabled](resources--global_log_receiver--reference--group-004.md#canonical-2dcc0b6a06e724b7f8ba24996ce7d4f2fe1c119b3a26673a9fdb241c65221c9b)
- [qradar_receiver.use_tls.mtls_enable](resources--global_log_receiver--reference--group-004.md#canonical-c0ff8cfb35ba02bb45894dd440a666929540a78e3b2fc206dcb62f94a888d75c)
- [qradar_receiver.use_tls.no_ca](resources--global_log_receiver--reference--group-004.md#canonical-767399ec1e9d93857cdca75e04dde228050045f3486bf56ce091add95778baff)
- [qradar_receiver](resources--global_log_receiver--reference--group-003.md#canonical-afb382cb3e4d5e62efc324cc347c6ce43343131b41944fdb2dca84647cc5f8be)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-bf63dfe4f551a495cbafb2f5b16298cb2ba2eb1fe0c451f000bb1594d90419ec"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c0c8625e77d74ca993b827fcbd2e6a3c2c960c655023e69ffab0af0433dd5cea"></a>

## qradar_receiver.use_tls.disable_verify_certificate — qradar_receiver.use_tls.disable_verify_certificate / 12dd6f9422dd / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [qradar_receiver](resources--global_log_receiver--reference--group-003.md#canonical-afb382cb3e4d5e62efc324cc347c6ce43343131b41944fdb2dca84647cc5f8be)
- [qradar_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-8380b97319f78370786be39917d8af086bd5a1781bbfcc2f916aa7c688f7a19f)
- qradar_receiver.use_tls.disable_verify_certificate

<a id="canonical-dea4fd7ebb1a920087e754ac6817a2d536db93476bf91d2046844353587d59a7"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable verify certificate.

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
disable_verify_certificate = {}
```

<a id="canonical-e9d52f535785e77ad80e94687904329701bb8cc58d2ccca99f43fb4fed35a63d"></a>

## Direct properties — qradar_receiver.use_tls.disable_verify_certificate / 12dd6f9422dd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c2c2f506878e0f46a9f3448533d40c5f49d1aaeb3ca0ffe5d763a92ee75d452e"></a>

## Next pages — qradar_receiver.use_tls.disable_verify_certificate / 12dd6f9422dd / 4

- [qradar_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-8380b97319f78370786be39917d8af086bd5a1781bbfcc2f916aa7c688f7a19f)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-12b9a4a82114c8022aad59dbc7bc29b87c772b4f425ff55df3c29d6b19ef4307"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
