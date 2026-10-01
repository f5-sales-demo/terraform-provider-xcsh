---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-3c0151ad0155f5da1ac5c73aede46e8aa8cb37785aa55fb0b1ebacff8bf61968"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / e7c887fd0df1 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-024.md#canonical-181a1518b8314bdeef4b5e79a2b8aacc9f75d1d5ed91dcbc4a205f9ad5b37385)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-024.md#canonical-a309460de31d5a477944352a93673e8a090b1d6c2342355139ca840878a2f2e6)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-024.md#canonical-b87765252f8d022952915dc691ece596862e75e46755df054811175f08b62bac)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-024.md#canonical-8ec02f648334ffabd08a1a24cb655feb3f94360a06e3130a289d4248a06127c7)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-97dcd243f228e249fcfa9da9c4403b30e821d7d4989686a1154dc2987d4e0e46"></a>

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

<a id="canonical-a5dabf1545b546efd7b82aff0d8c074d49a7fcee722ddbf6b016683ef3a4ea0a"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / e7c887fd0df1 / 3

- [header_transformation](resources--workload--reference--group-025.md#canonical-d2acfeb449d068421d08f9e5d93b7ec539e5cea9cf37d137a9f5a6dd25dce493): complete subsection reference.

<a id="canonical-31be5c908a8d4f26e9dc150c4a032a2649c0023dddcc06ae489bda1bfca5f7db"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / e7c887fd0df1 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-025.md#canonical-d2acfeb449d068421d08f9e5d93b7ec539e5cea9cf37d137a9f5a6dd25dce493)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-024.md#canonical-8ec02f648334ffabd08a1a24cb655feb3f94360a06e3130a289d4248a06127c7)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-d2acfeb449d068421d08f9e5d93b7ec539e5cea9cf37d137a9f5a6dd25dce493"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f7a52359d87aa686a86318702695f891d64bea9b86d167b5c86c45472cb0f012"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 0fe231d38651 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-024.md#canonical-181a1518b8314bdeef4b5e79a2b8aacc9f75d1d5ed91dcbc4a205f9ad5b37385)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-024.md#canonical-a309460de31d5a477944352a93673e8a090b1d6c2342355139ca840878a2f2e6)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-024.md#canonical-b87765252f8d022952915dc691ece596862e75e46755df054811175f08b62bac)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-024.md#canonical-8ec02f648334ffabd08a1a24cb655feb3f94360a06e3130a289d4248a06127c7)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-024.md#canonical-92465093303585154e382d027fedc271e7ad828b7b464e54c366e157a8a8de55)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-03328b4f00b4917f78027711ed7de3c8b5cbf2212f072849ade7cb242ded133a"></a>

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

<a id="canonical-ec3db4812893d801e97da12539ddd260ee50f3aded119071e522d39aaf26582d"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 0fe231d38651 / 3

- [default_header_transformation](resources--workload--reference--group-025.md#canonical-79fde7e1ebd5ffd52ea0c3201d3832a0b405676f6279e0746352413eb0cf8340): complete subsection reference.

- [preserve_case_header_transformation](resources--workload--reference--group-025.md#canonical-ca2996ea7972aa462ab7037908285244be5bc89caf602a89d8857e914a11b997): complete subsection reference.

- [proper_case_header_transformation](resources--workload--reference--group-025.md#canonical-271133419ba23d0dd010d9c6c52cf5a19d2465ecb789bbcc6f33e75051e21cd6): complete subsection reference.

<a id="canonical-d038a6aaf1ead3dd3c1835e3a28b7dd3d6bb539d45ed2a67391e6d2072533390"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 0fe231d38651 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](resources--workload--reference--group-025.md#canonical-79fde7e1ebd5ffd52ea0c3201d3832a0b405676f6279e0746352413eb0cf8340)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](resources--workload--reference--group-025.md#canonical-ca2996ea7972aa462ab7037908285244be5bc89caf602a89d8857e914a11b997)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](resources--workload--reference--group-025.md#canonical-271133419ba23d0dd010d9c6c52cf5a19d2465ecb789bbcc6f33e75051e21cd6)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-024.md#canonical-92465093303585154e382d027fedc271e7ad828b7b464e54c366e157a8a8de55)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-79fde7e1ebd5ffd52ea0c3201d3832a0b405676f6279e0746352413eb0cf8340"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-647c8c4c467d9727dcec22e083f20219e7e073ca8ee002ed89f666721ae740b1"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 7d733c7e194b / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-024.md#canonical-181a1518b8314bdeef4b5e79a2b8aacc9f75d1d5ed91dcbc4a205f9ad5b37385)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-024.md#canonical-a309460de31d5a477944352a93673e8a090b1d6c2342355139ca840878a2f2e6)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-024.md#canonical-b87765252f8d022952915dc691ece596862e75e46755df054811175f08b62bac)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-024.md#canonical-8ec02f648334ffabd08a1a24cb655feb3f94360a06e3130a289d4248a06127c7)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-024.md#canonical-92465093303585154e382d027fedc271e7ad828b7b464e54c366e157a8a8de55)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-025.md#canonical-d2acfeb449d068421d08f9e5d93b7ec539e5cea9cf37d137a9f5a6dd25dce493)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-fb02cfadb9ce79f4f53db2b3266363580f8d39b25aa92db4f95d876284d48872"></a>

Type: `["object", {}]`. Optional.

Use the platform's current default HTTP header transformation behavior.

Receipt-pinned upstream constraints:

```json
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
default_header_transformation = {}
```

<a id="canonical-73e7c08daec247699a1c78459501ca506519bfa6122e0a42bd4e8be8a3564e53"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 7d733c7e194b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-eefebd454307dfad6329bef59a0aea7756ca52ab0e75d9cdb525a84ffd87d477"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 7d733c7e194b / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-025.md#canonical-d2acfeb449d068421d08f9e5d93b7ec539e5cea9cf37d137a9f5a6dd25dce493)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-ca2996ea7972aa462ab7037908285244be5bc89caf602a89d8857e914a11b997"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ac573a1c33edb3711e272ea9e2ca7c453da31ffcdc43daa18d32a050aa347237"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / e0b9b70eede0 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-024.md#canonical-181a1518b8314bdeef4b5e79a2b8aacc9f75d1d5ed91dcbc4a205f9ad5b37385)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-024.md#canonical-a309460de31d5a477944352a93673e8a090b1d6c2342355139ca840878a2f2e6)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-024.md#canonical-b87765252f8d022952915dc691ece596862e75e46755df054811175f08b62bac)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-024.md#canonical-8ec02f648334ffabd08a1a24cb655feb3f94360a06e3130a289d4248a06127c7)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-024.md#canonical-92465093303585154e382d027fedc271e7ad828b7b464e54c366e157a8a8de55)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-025.md#canonical-d2acfeb449d068421d08f9e5d93b7ec539e5cea9cf37d137a9f5a6dd25dce493)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-4e8e9e39327af2d14eb9155919d0b27b2b2b33479982bc05c933923eb8c61c32"></a>

Type: `["object", {}]`. Optional.

Preserve HTTP header-name case when upstream case must remain unchanged.

Receipt-pinned upstream constraints:

```json
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
preserve_case_header_transformation = {}
```

<a id="canonical-fe521c8e2b730e11dc12cc801908ca4ebdd4015dddf664bab0966a7ccad7b6ec"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / e0b9b70eede0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2250d8099b93f0a4ec1a2c25dc3293a0e9414e8075530e694204bb6aa66359ca"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / e0b9b70eede0 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-025.md#canonical-d2acfeb449d068421d08f9e5d93b7ec539e5cea9cf37d137a9f5a6dd25dce493)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-271133419ba23d0dd010d9c6c52cf5a19d2465ecb789bbcc6f33e75051e21cd6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7f5ebe9108663109f141e1e38deaf26e0566786609969a77613f5db03fa3e7df"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 543e2a07ae52 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-024.md#canonical-181a1518b8314bdeef4b5e79a2b8aacc9f75d1d5ed91dcbc4a205f9ad5b37385)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-024.md#canonical-a309460de31d5a477944352a93673e8a090b1d6c2342355139ca840878a2f2e6)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-024.md#canonical-b87765252f8d022952915dc691ece596862e75e46755df054811175f08b62bac)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-024.md#canonical-8ec02f648334ffabd08a1a24cb655feb3f94360a06e3130a289d4248a06127c7)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-024.md#canonical-92465093303585154e382d027fedc271e7ad828b7b464e54c366e157a8a8de55)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-025.md#canonical-d2acfeb449d068421d08f9e5d93b7ec539e5cea9cf37d137a9f5a6dd25dce493)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-93f4f51c1fe6c3841c1e95adef5506925f09d3e9d823ff531a8571b12dea4211"></a>

Type: `["object", {}]`. Optional.

Transform HTTP header names to proper case when explicit transformation is required.

Receipt-pinned upstream constraints:

```json
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
proper_case_header_transformation = {}
```

<a id="canonical-1ac504eb47a6102d4777220b90435ee4ddd9a39d2d844df3176fd763a2ffda25"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 543e2a07ae52 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d0fb02d7845cd5fbfd2f92ea469c118c16700d6520a8cc3910d421ec2eb1a00e"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 543e2a07ae52 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-025.md#canonical-d2acfeb449d068421d08f9e5d93b7ec539e5cea9cf37d137a9f5a6dd25dce493)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-15441c05ad8def32778d9ef4d5922b2072fdc86634838ba94a24aa885b666301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8c6750d1325fcb609fa1bd1df0c4c3919823413d9c21071e3e5a4e33d755ab5b"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2 — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 94a8c577d2ee / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-024.md#canonical-181a1518b8314bdeef4b5e79a2b8aacc9f75d1d5ed91dcbc4a205f9ad5b37385)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-024.md#canonical-a309460de31d5a477944352a93673e8a090b1d6c2342355139ca840878a2f2e6)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-024.md#canonical-b87765252f8d022952915dc691ece596862e75e46755df054811175f08b62bac)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-024.md#canonical-8ec02f648334ffabd08a1a24cb655feb3f94360a06e3130a289d4248a06127c7)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-9567cdb2c27cc4766ff4e16806e9f88ebe9d208bb3c7c7ffde9782a8dda8a33a"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for http protocol enable v1 v2.

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
http_protocol_enable_v1_v2 = {}
```

<a id="canonical-bb5699db7d5f5361d5fa71e0c5a23ffb3bd1b3221905935a928f4411c29367a0"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 94a8c577d2ee / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6785888e7783a49b3fc46a0d5844ec859006047cd90b66e7034c59a16ad58893"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 94a8c577d2ee / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-024.md#canonical-8ec02f648334ffabd08a1a24cb655feb3f94360a06e3130a289d4248a06127c7)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-ff148f42c7de06c827e69e8040208dcfeb1f0b90105212ab6d7b77792ac87c52"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-50017a0830f3788df6b76217dd706274de0927be69824dbd97952745ce59877e"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 7b16c6646285 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-024.md#canonical-181a1518b8314bdeef4b5e79a2b8aacc9f75d1d5ed91dcbc4a205f9ad5b37385)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-024.md#canonical-a309460de31d5a477944352a93673e8a090b1d6c2342355139ca840878a2f2e6)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-024.md#canonical-b87765252f8d022952915dc691ece596862e75e46755df054811175f08b62bac)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-024.md#canonical-8ec02f648334ffabd08a1a24cb655feb3f94360a06e3130a289d4248a06127c7)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-194be731ad9039f62286670b419972dfa017775d1d2fd20294926a5761d0c06c"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for http protocol enable v2 only.

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
http_protocol_enable_v2_only = {}
```

<a id="canonical-5998e36278b15c8acb251083cdf7e23384b55c928d54061d60a564be37313c16"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 7b16c6646285 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d6568e703502671612e733bad65ac8bcfebc28d94df9e501ce7eef7e8988409c"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 7b16c6646285 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-024.md#canonical-8ec02f648334ffabd08a1a24cb655feb3f94360a06e3130a289d4248a06127c7)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-852102dc74be9c177a92828aa9e940c8f6f9416323b1a52561778bfb60a87789"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-12e5fe358fa7e64b80a17da781094c27070ea2e13ca83259397f1a7e537e512c"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.non_default_loadbalancer — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 9e4be5ef1575 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-024.md#canonical-181a1518b8314bdeef4b5e79a2b8aacc9f75d1d5ed91dcbc4a205f9ad5b37385)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-024.md#canonical-a309460de31d5a477944352a93673e8a090b1d6c2342355139ca840878a2f2e6)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-024.md#canonical-b87765252f8d022952915dc691ece596862e75e46755df054811175f08b62bac)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.non_default_loadbalancer

<a id="canonical-7da3605a539b50750e6d68bded79ede09c19c6639ad5992f5dfdcd268f2bcfb5"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for non default loadbalancer.

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
non_default_loadbalancer = {}
```

<a id="canonical-01e5c5dad3aab0a9abc795d627b5dda87d51d96627c2223b83d45d0bee28c55f"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 9e4be5ef1575 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b9d918a19a71c95923a250a6144305ede5ae6951a6b2c0a6ba255b1558dfd577"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 9e4be5ef1575 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-024.md#canonical-b87765252f8d022952915dc691ece596862e75e46755df054811175f08b62bac)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-7b295371e956977ac8a72d8e065faf5d1796e294e82de61df7928f7f6f1e86ba"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-60b1e612fe40a8c6535e6034a098182682ca428a938d7b064ef67fe756e8ef9f"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.pass_through — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 0521a7883598 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-024.md#canonical-181a1518b8314bdeef4b5e79a2b8aacc9f75d1d5ed91dcbc4a205f9ad5b37385)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-024.md#canonical-a309460de31d5a477944352a93673e8a090b1d6c2342355139ca840878a2f2e6)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-024.md#canonical-b87765252f8d022952915dc691ece596862e75e46755df054811175f08b62bac)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.pass_through

<a id="canonical-315a83539197e14202f0e9148cc7f06360dcd36cb9443fb753d8b5bcf9378569"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for pass through.

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
pass_through = {}
```

<a id="canonical-3c3222a9dedce0db9b68c59858c7fac461a985652f0bd7ca5d9cddab4f38d79b"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 0521a7883598 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2982365732fa7a40557649121f12d66f5aebb864c3f600567ad8cfbbe2f31930"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 0521a7883598 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-024.md#canonical-b87765252f8d022952915dc691ece596862e75e46755df054811175f08b62bac)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-3c628a7bc1cdae78492582bbcf546221a1af0a6a9879bb9cc9756246e42ba22d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5e85d86943f7d19ba50636b6153fff176b04ada86c398299d525d84a5ce56fcd"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 85aea024a30e / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-024.md#canonical-181a1518b8314bdeef4b5e79a2b8aacc9f75d1d5ed91dcbc4a205f9ad5b37385)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-024.md#canonical-a309460de31d5a477944352a93673e8a090b1d6c2342355139ca840878a2f2e6)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-024.md#canonical-b87765252f8d022952915dc691ece596862e75e46755df054811175f08b62bac)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params

<a id="canonical-2e87d8f49a8b838c2bd045893e7ea10ea96d2c212e4afa23c0578f900c0878e2"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for tls cert params.

Upstream description:

Select TLS Parameters and Certificates.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("certificates"),
  validators.ConflictingObjectAttributes("no_mtls",
    "use_mtls")}
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
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

Terraform syntax:

```terraform
tls_cert_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-a2471ad5f441129e216129d4aa891648c7e52ec2dbd6feda0e7c9f8014ce3a33"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 85aea024a30e / 3

- [certificates](resources--workload--reference--group-025.md#canonical-78e3ef87c00bad2226a8291d60c061d35e64e7f7b4e4c89291c7c1a865bed9b7): complete subsection reference.

- [no_mtls](resources--workload--reference--group-025.md#canonical-f5c429eece9afff07184051a65044005af040af255845dc32f835aa1a360cde6): complete subsection reference.

- [tls_config](resources--workload--reference--group-025.md#canonical-c5e9027a9924ad008528dc71fbbf429d17d91e3b254331e90baf9d109f92af45): complete subsection reference.

- [use_mtls](resources--workload--reference--group-025.md#canonical-4b6b4795a6f2233496dc880bd1e62b3368648a7664d327459eca0169422c8c2b): complete subsection reference.

<a id="canonical-fe4f12cc39c2170891a75cd2380682e910b7c55baf1ea856a2af6d8d14049a67"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 85aea024a30e / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.certificates](resources--workload--reference--group-025.md#canonical-78e3ef87c00bad2226a8291d60c061d35e64e7f7b4e4c89291c7c1a865bed9b7)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.no_mtls](resources--workload--reference--group-025.md#canonical-f5c429eece9afff07184051a65044005af040af255845dc32f835aa1a360cde6)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-025.md#canonical-c5e9027a9924ad008528dc71fbbf429d17d91e3b254331e90baf9d109f92af45)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-025.md#canonical-4b6b4795a6f2233496dc880bd1e62b3368648a7664d327459eca0169422c8c2b)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-024.md#canonical-b87765252f8d022952915dc691ece596862e75e46755df054811175f08b62bac)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-78e3ef87c00bad2226a8291d60c061d35e64e7f7b4e4c89291c7c1a865bed9b7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b25e0899dafd80dfbaead4ab09aad39934ce90db190570b4e73752c5db68e796"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.certificates — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 5c8c2b3d0009 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-024.md#canonical-181a1518b8314bdeef4b5e79a2b8aacc9f75d1d5ed91dcbc4a205f9ad5b37385)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-024.md#canonical-a309460de31d5a477944352a93673e8a090b1d6c2342355139ca840878a2f2e6)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-024.md#canonical-b87765252f8d022952915dc691ece596862e75e46755df054811175f08b62bac)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-025.md#canonical-3c628a7bc1cdae78492582bbcf546221a1af0a6a9879bb9cc9756246e42ba22d)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.certificates

<a id="canonical-099f8806cc8ca89de5f3e04ff453611ee7d0e832e157183ee0828847bcd0fb4b"></a>

Type: `"object"`. list nested block, Optional.

Select one or more certificates with any domain names.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
certificates {
  # Configure direct properties listed below.
}
```

<a id="canonical-b788d88df7d07a4318056580ed59c8abc83795e33ec0eb566ae51ef2cafc4397"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 5c8c2b3d0009 / 3

<a id="canonical-ff350b4bcef63ea290c54a5f4c5033e9d7a1ed225325d023e9901a046c0c4c27"></a>

<a id="canonical-ee1039ba1409da70e77ef6a1107e5fff9169e28bde57d2b3136d7ef0eaf39336"></a>

## name property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 5c8c2b3d0009 / 4

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

<a id="canonical-25084a3ab6216ef2fc6fb30886d737f07912ca60dfad9e2fd3b31d8ab6d57f42"></a>

<a id="canonical-7d01ddce2a6aa7f29fb6762fb2cc8a63bc4eefc2808bd7aa023653c446c4515a"></a>

## namespace property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 5c8c2b3d0009 / 5

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

<a id="canonical-4fcec1b526a7394da0b64f4bf74aecf44ccf0c2a1721900fb18cb633512becda"></a>

<a id="canonical-48e2a351d7b4abee27ff3089165f7830db39c01e396b49dcfc5db15b51eb14fd"></a>

## tenant property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 5c8c2b3d0009 / 6

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

<a id="canonical-d6f1532bf8afc282329b6d5c7358ef7621b154f33fea11f60acaeb4ee852981d"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 5c8c2b3d0009 / 7

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-025.md#canonical-3c628a7bc1cdae78492582bbcf546221a1af0a6a9879bb9cc9756246e42ba22d)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-f5c429eece9afff07184051a65044005af040af255845dc32f835aa1a360cde6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-60cd308b65d4052ed266aa039aa362c10e19c9a023412a9a20524812b11ef82e"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.no_mtls — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / a20e6a04b450 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-024.md#canonical-181a1518b8314bdeef4b5e79a2b8aacc9f75d1d5ed91dcbc4a205f9ad5b37385)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-024.md#canonical-a309460de31d5a477944352a93673e8a090b1d6c2342355139ca840878a2f2e6)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-024.md#canonical-b87765252f8d022952915dc691ece596862e75e46755df054811175f08b62bac)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-025.md#canonical-3c628a7bc1cdae78492582bbcf546221a1af0a6a9879bb9cc9756246e42ba22d)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.no_mtls

<a id="canonical-d635a4f81081830f87fb243c0272109c99cf04ddf9b1aac3ff36588b618fab3c"></a>

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
no_mtls = {}
```

<a id="canonical-4b72374b2a14f37025b04776c36f7b75f10c06c71152e957f3e47374439bf25a"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / a20e6a04b450 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0d22987ba97f87c325cdc8bdba72fe2829443b87a7e0d8241accc43928981d0b"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / a20e6a04b450 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-025.md#canonical-3c628a7bc1cdae78492582bbcf546221a1af0a6a9879bb9cc9756246e42ba22d)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-c5e9027a9924ad008528dc71fbbf429d17d91e3b254331e90baf9d109f92af45"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7e673f1ef43e99a0c77b6f5361c232391a8757623024b867a0561dc54ae17a75"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 7318c8f273b0 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-024.md#canonical-181a1518b8314bdeef4b5e79a2b8aacc9f75d1d5ed91dcbc4a205f9ad5b37385)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-024.md#canonical-a309460de31d5a477944352a93673e8a090b1d6c2342355139ca840878a2f2e6)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-024.md#canonical-b87765252f8d022952915dc691ece596862e75e46755df054811175f08b62bac)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-025.md#canonical-3c628a7bc1cdae78492582bbcf546221a1af0a6a9879bb9cc9756246e42ba22d)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config

<a id="canonical-51ff00bbd634ec6078dc38f8a9b12002178340fbed8aa8dd9ea7c8c475a4b6b6"></a>

Type: `"object"`. single nested block, Optional.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_security",
    "default_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "low_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("default_security",
    "low_security"),
  validators.ConflictingObjectAttributes("default_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("low_security",
    "medium_security")}
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
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

Terraform syntax:

```terraform
tls_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-1e07f2dede9ff44e5d0414a5541b96c7b9436dfe45c5523df9d3c1cf1298c215"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 7318c8f273b0 / 3

- [custom_security](resources--workload--reference--group-025.md#canonical-96e9ee576e4bb78948288b679468bbc161a1e963f374c1db80d76dbc5a936eb7): complete subsection reference.

- [default_security](resources--workload--reference--group-025.md#canonical-823cb1b0a14f4f44f966beccf0a7e848693bcbbeccc2a9ad1ba02d49dd6dfdd2): complete subsection reference.

- [low_security](resources--workload--reference--group-025.md#canonical-f8f17378ce79fac1b842bd6210ddf5ba721fb37b3e91e0425641a3881ee5c30e): complete subsection reference.

- [medium_security](resources--workload--reference--group-025.md#canonical-bb88497f3cf2308e9629b5ab59c79b74e3e6b9f2f5ab05cfb3a13bc44db40855): complete subsection reference.

<a id="canonical-7d810c0a6ff15373a96daa5acdc57c2dcc5d4e9df13f9cdf62e963154eb27688"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 7318c8f273b0 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.custom_security](resources--workload--reference--group-025.md#canonical-96e9ee576e4bb78948288b679468bbc161a1e963f374c1db80d76dbc5a936eb7)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.default_security](resources--workload--reference--group-025.md#canonical-823cb1b0a14f4f44f966beccf0a7e848693bcbbeccc2a9ad1ba02d49dd6dfdd2)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.low_security](resources--workload--reference--group-025.md#canonical-f8f17378ce79fac1b842bd6210ddf5ba721fb37b3e91e0425641a3881ee5c30e)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.medium_security](resources--workload--reference--group-025.md#canonical-bb88497f3cf2308e9629b5ab59c79b74e3e6b9f2f5ab05cfb3a13bc44db40855)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-025.md#canonical-3c628a7bc1cdae78492582bbcf546221a1af0a6a9879bb9cc9756246e42ba22d)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-96e9ee576e4bb78948288b679468bbc161a1e963f374c1db80d76dbc5a936eb7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dc8c34948110a50baf68ffebd02abcac21e4e73dffda4e5c2254a6dcdccf1541"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.custom_security — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 28e3c407d07e / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-024.md#canonical-181a1518b8314bdeef4b5e79a2b8aacc9f75d1d5ed91dcbc4a205f9ad5b37385)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-024.md#canonical-a309460de31d5a477944352a93673e8a090b1d6c2342355139ca840878a2f2e6)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-024.md#canonical-b87765252f8d022952915dc691ece596862e75e46755df054811175f08b62bac)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-025.md#canonical-3c628a7bc1cdae78492582bbcf546221a1af0a6a9879bb9cc9756246e42ba22d)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-025.md#canonical-c5e9027a9924ad008528dc71fbbf429d17d91e3b254331e90baf9d109f92af45)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.custom_security

<a id="canonical-41764934a3d8b22395b36ef1e7cbea406f196ddbf64f505d21857a83a591985a"></a>

Type: `"object"`. single nested block, Optional.

Defines TLS protocol config including min/max versions and allowed ciphers.

Upstream description:

This defines TLS protocol config including min/max versions and allowed ciphers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cipher_suites")}
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
custom_security {
  # Configure direct properties listed below.
}
```

<a id="canonical-d1febe92a1a04b57834bce15497e7b3111f37a4a0779d903db3612d71ff40228"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 28e3c407d07e / 3

<a id="canonical-d80e1d6dbce4d3f6eba70836c7fa6d78f9519cd4c651b63653d8989dafd218c6"></a>

<a id="canonical-9e003ecc2ea25df382e6890f865c3194828dfb6e2d7e8b683f96da38f2be4127"></a>

## cipher_suites property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 28e3c407d07e / 4

Type: `["list", "string"]`. Optional.

The TLS listener will only support the specified cipher list.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-ae9f6f36b24a13bb8202c99fa16c145ac5d637a46cdc91f62e55b2712dda3620"></a>

<a id="canonical-f5afebf54cce806a7829dfb9f7793e0870e98551f131338dc977c2b0a25b2080"></a>

## max_version property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 28e3c407d07e / 5

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

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

<a id="canonical-5ccf4f099b930030cdb387ca060700e5269dde2f682034a89bc7ec7efaa963c7"></a>

<a id="canonical-6f83544b28082ea859877db35fe64300e3a6319492577804ab27ab2552768694"></a>

## min_version property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 28e3c407d07e / 6

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

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

<a id="canonical-a479732cf89e19244ab045ff59a2a5da7e3c50859d37055ae35c81899ab48633"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 28e3c407d07e / 7

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-025.md#canonical-c5e9027a9924ad008528dc71fbbf429d17d91e3b254331e90baf9d109f92af45)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-823cb1b0a14f4f44f966beccf0a7e848693bcbbeccc2a9ad1ba02d49dd6dfdd2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bbbb600b8d17e1403ae593b1d14e05809c18e219752fc4588c1f5f214481150c"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.default_security — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 4285606a3e91 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-024.md#canonical-181a1518b8314bdeef4b5e79a2b8aacc9f75d1d5ed91dcbc4a205f9ad5b37385)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-024.md#canonical-a309460de31d5a477944352a93673e8a090b1d6c2342355139ca840878a2f2e6)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-024.md#canonical-b87765252f8d022952915dc691ece596862e75e46755df054811175f08b62bac)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-025.md#canonical-3c628a7bc1cdae78492582bbcf546221a1af0a6a9879bb9cc9756246e42ba22d)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-025.md#canonical-c5e9027a9924ad008528dc71fbbf429d17d91e3b254331e90baf9d109f92af45)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.default_security

<a id="canonical-5ff02168b112e5880b211bff2c0047f77b1b3d914716265cf49c9b243ade238c"></a>

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
default_security = {}
```

<a id="canonical-701ac3c3ed8305bbdc9936b4ffad18ed394b41274d1b302ed3cccb43cbc6335c"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 4285606a3e91 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-906555b689b2875b93f34095c6158fa028607cc1191edee47fa5f35e4a839efb"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 4285606a3e91 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-025.md#canonical-c5e9027a9924ad008528dc71fbbf429d17d91e3b254331e90baf9d109f92af45)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-f8f17378ce79fac1b842bd6210ddf5ba721fb37b3e91e0425641a3881ee5c30e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-20543e09c7f80b7d8a2578eba8e1feeb57def5bf766f0b16d2e17cc3101b1995"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.low_security — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 5789298536cc / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-024.md#canonical-181a1518b8314bdeef4b5e79a2b8aacc9f75d1d5ed91dcbc4a205f9ad5b37385)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-024.md#canonical-a309460de31d5a477944352a93673e8a090b1d6c2342355139ca840878a2f2e6)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-024.md#canonical-b87765252f8d022952915dc691ece596862e75e46755df054811175f08b62bac)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-025.md#canonical-3c628a7bc1cdae78492582bbcf546221a1af0a6a9879bb9cc9756246e42ba22d)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-025.md#canonical-c5e9027a9924ad008528dc71fbbf429d17d91e3b254331e90baf9d109f92af45)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.low_security

<a id="canonical-d945defcdb6ccbd9d9524b8c551dc46cbfa3869fee8c15c3ee6c36cd971b19d2"></a>

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
low_security = {}
```

<a id="canonical-ee08cb5caacdf64c88d0dacbf1735ef663d0b77dc9e65c1d0bc8361743800968"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 5789298536cc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9a7901b10791c4c9ed9414ac0bd02365e796c43f47aacd01beba607cf50565e2"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 5789298536cc / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-025.md#canonical-c5e9027a9924ad008528dc71fbbf429d17d91e3b254331e90baf9d109f92af45)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-bb88497f3cf2308e9629b5ab59c79b74e3e6b9f2f5ab05cfb3a13bc44db40855"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8b185575b9d9cf8b0c1653eabf71e8c7eef621552d293b893dbaf6e19ddb0866"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.medium_security — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / e9ee208426ea / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-024.md#canonical-181a1518b8314bdeef4b5e79a2b8aacc9f75d1d5ed91dcbc4a205f9ad5b37385)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-024.md#canonical-a309460de31d5a477944352a93673e8a090b1d6c2342355139ca840878a2f2e6)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-024.md#canonical-b87765252f8d022952915dc691ece596862e75e46755df054811175f08b62bac)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-025.md#canonical-3c628a7bc1cdae78492582bbcf546221a1af0a6a9879bb9cc9756246e42ba22d)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-025.md#canonical-c5e9027a9924ad008528dc71fbbf429d17d91e3b254331e90baf9d109f92af45)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.medium_security

<a id="canonical-83bfc73bcdd3c034c2bfd95348acd0c2a121a4ef4b62a39d354146a1d202614f"></a>

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
medium_security = {}
```

<a id="canonical-5d1c0663825c13099410ac47c418a3cf226ba6b5f74be1bc71ed32f879240a99"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / e9ee208426ea / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e4a67186fd80d34978ebcfe236d0ceeebeb2db0235db0fef7fe3965d5852db0d"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / e9ee208426ea / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-025.md#canonical-c5e9027a9924ad008528dc71fbbf429d17d91e3b254331e90baf9d109f92af45)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-4b6b4795a6f2233496dc880bd1e62b3368648a7664d327459eca0169422c8c2b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b23d1d184477a62232907edcbad17d37b38f2888a10e1d04c518c3b1e0680927"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 9c16cb05733f / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-024.md#canonical-181a1518b8314bdeef4b5e79a2b8aacc9f75d1d5ed91dcbc4a205f9ad5b37385)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-024.md#canonical-a309460de31d5a477944352a93673e8a090b1d6c2342355139ca840878a2f2e6)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-024.md#canonical-b87765252f8d022952915dc691ece596862e75e46755df054811175f08b62bac)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-025.md#canonical-3c628a7bc1cdae78492582bbcf546221a1af0a6a9879bb9cc9756246e42ba22d)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls

<a id="canonical-f7933ee33c03673020310e9305531978c55bb47072ba1d9f9a4ad964f7d05706"></a>

Type: `"object"`. single nested block, Optional.

Validation context for downstream client TLS connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("crl",
    "no_crl"),
  validators.ConflictingObjectAttributes("trusted_ca",
    "trusted_ca_url"),
  validators.ConflictingObjectAttributes("xfcc_disabled",
    "xfcc_options")}
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
  "x-ves-oneof-field-crl_choice": "[\"crl\",\"no_crl\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-xfcc_header": "[\"xfcc_disabled\",\"xfcc_options\"]"
}
```

Terraform syntax:

```terraform
use_mtls {
  # Configure direct properties listed below.
}
```

<a id="canonical-f574a2cfef257b363c0d70eb1ba46898e4a2e20964a7ebbd524e1964f8791d3d"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 9c16cb05733f / 3

<a id="canonical-63236dde64333bd10081b88dc63c544734cea2b0d798aedd5511b0132b340caf"></a>

<a id="canonical-c015c315b244e277af758122792f50f0a0fb294835bbe07fea8a887b1b7e7bd5"></a>

## client_certificate_optional property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 9c16cb05733f / 4

Type: `"bool"`. Optional.

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated.

Upstream description:

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated. If the client
does not provide a certificate, the connection will be accepted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [crl](resources--workload--reference--group-025.md#canonical-206e14e743c7a80978d86f819931977b30764d0bb43260b8e21988ffee248ecd): complete subsection reference.

- [no_crl](resources--workload--reference--group-025.md#canonical-f05b1225a48eab661e7b44c763c7dba91f07e8ff7315085fb3b53bf44c2afeff): complete subsection reference.

- [trusted_ca](resources--workload--reference--group-025.md#canonical-33abe851d3bfb228ccfa68a1d9d78c744e21d04c688e656bd7cda112bb89fa57): complete subsection reference.

<a id="canonical-aeb279f3b0bd9c769604b26fd16e91223a5af8775746c30c31ff8767e927408a"></a>

<a id="canonical-837a4e143039a7d9cae24502f50e4591c562d1cecc69ad7850ab42ee18f4145b"></a>

## trusted_ca_url property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 9c16cb05733f / 5

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Upstream description:

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

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
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

- [xfcc_disabled](resources--workload--reference--group-025.md#canonical-c841cb951a1c4f9d7501b5f4a236ba845ed4c7d507aa5f4153aaf6e5db31bb32): complete subsection reference.

- [xfcc_options](resources--workload--reference--group-025.md#canonical-69786fb447c3ee5e0def78812708df796e9f9e9848d7fb7f2d060a9cdc4b7864): complete subsection reference.

<a id="canonical-ce49c9bebb20589d5121093b0327bc3e8393d55e07353e06186bc9618230b051"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 9c16cb05733f / 6

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.crl](resources--workload--reference--group-025.md#canonical-206e14e743c7a80978d86f819931977b30764d0bb43260b8e21988ffee248ecd)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.no_crl](resources--workload--reference--group-025.md#canonical-f05b1225a48eab661e7b44c763c7dba91f07e8ff7315085fb3b53bf44c2afeff)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca](resources--workload--reference--group-025.md#canonical-33abe851d3bfb228ccfa68a1d9d78c744e21d04c688e656bd7cda112bb89fa57)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_disabled](resources--workload--reference--group-025.md#canonical-c841cb951a1c4f9d7501b5f4a236ba845ed4c7d507aa5f4153aaf6e5db31bb32)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options](resources--workload--reference--group-025.md#canonical-69786fb447c3ee5e0def78812708df796e9f9e9848d7fb7f2d060a9cdc4b7864)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-025.md#canonical-3c628a7bc1cdae78492582bbcf546221a1af0a6a9879bb9cc9756246e42ba22d)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-206e14e743c7a80978d86f819931977b30764d0bb43260b8e21988ffee248ecd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c00e6a5bf10d3862b86d4ed5355e4eafd4f3b5297306dcd433d51197dcf5029e"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.crl — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / dc771878a095 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-024.md#canonical-181a1518b8314bdeef4b5e79a2b8aacc9f75d1d5ed91dcbc4a205f9ad5b37385)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-024.md#canonical-a309460de31d5a477944352a93673e8a090b1d6c2342355139ca840878a2f2e6)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-024.md#canonical-b87765252f8d022952915dc691ece596862e75e46755df054811175f08b62bac)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-025.md#canonical-3c628a7bc1cdae78492582bbcf546221a1af0a6a9879bb9cc9756246e42ba22d)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-025.md#canonical-4b6b4795a6f2233496dc880bd1e62b3368648a7664d327459eca0169422c8c2b)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.crl

<a id="canonical-c6ccb50b8a6f8a2a6a0cdc5b10a450297475836f0bbd86b19a1280b203a402f1"></a>

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
crl {
  # Configure direct properties listed below.
}
```

<a id="canonical-da1139e35cbc283013463422f7f99489ed58aed6a45bcbc7ee92c6e0a6505041"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / dc771878a095 / 3

<a id="canonical-802720feb992e67de1475bc2ddfdd9ff7655341f5c095a8288e2d2de4996838a"></a>

<a id="canonical-bdb8a890c093c892feca5ecb0e49702e15358df79136abfba81bccd818a3fa76"></a>

## name property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / dc771878a095 / 4

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

<a id="canonical-0204ecff34bc814af0ea60a13ac8dbb59b7f1daf8c35bcbb8c7ccc4bd278b252"></a>

<a id="canonical-d58fcc2d0b5ef3be6f937b6efd22f28ccec7e67584ed07315da105cffc661a40"></a>

## namespace property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / dc771878a095 / 5

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

<a id="canonical-f41529d9afd7401bc9953f1865aa9090778a2d6f0590a057e450a77b3bf345b0"></a>

<a id="canonical-de866587a754e0bca3b18704647fada3188b2c4bb9b48aaac71f326a9b935e57"></a>

## tenant property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / dc771878a095 / 6

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

<a id="canonical-29389d02037740f9b472e1236e9617bfe8adec6bb9d4a665aa9014f3b5535d3b"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / dc771878a095 / 7

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-025.md#canonical-4b6b4795a6f2233496dc880bd1e62b3368648a7664d327459eca0169422c8c2b)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-f05b1225a48eab661e7b44c763c7dba91f07e8ff7315085fb3b53bf44c2afeff"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-74406f68fde989d563b388225f0138ff78d57fa201df961bc0c858efff3b86fa"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.no_crl — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 529f78539b5c / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-024.md#canonical-181a1518b8314bdeef4b5e79a2b8aacc9f75d1d5ed91dcbc4a205f9ad5b37385)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-024.md#canonical-a309460de31d5a477944352a93673e8a090b1d6c2342355139ca840878a2f2e6)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-024.md#canonical-b87765252f8d022952915dc691ece596862e75e46755df054811175f08b62bac)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-025.md#canonical-3c628a7bc1cdae78492582bbcf546221a1af0a6a9879bb9cc9756246e42ba22d)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-025.md#canonical-4b6b4795a6f2233496dc880bd1e62b3368648a7664d327459eca0169422c8c2b)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.no_crl

<a id="canonical-cf1c06a246b12258b8112f582b4c7cad30d2909aa0cca32f7fee99b249f1f061"></a>

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
no_crl = {}
```

<a id="canonical-6714ed86203a57566b7b3a4de217c5b487af132840927b954bceeaefc521cdc1"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 529f78539b5c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cdd5d9c8282d76e35e7c5a30cb8c3dc1321d23db02bb814bc0c885b8e26130fb"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 529f78539b5c / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-025.md#canonical-4b6b4795a6f2233496dc880bd1e62b3368648a7664d327459eca0169422c8c2b)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-33abe851d3bfb228ccfa68a1d9d78c744e21d04c688e656bd7cda112bb89fa57"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a16cbb4847bc5bd8b09e3999bec3e08248c9143dc477441b878ee18240479216"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / e2786a5e4eaf / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-024.md#canonical-181a1518b8314bdeef4b5e79a2b8aacc9f75d1d5ed91dcbc4a205f9ad5b37385)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-024.md#canonical-a309460de31d5a477944352a93673e8a090b1d6c2342355139ca840878a2f2e6)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-024.md#canonical-b87765252f8d022952915dc691ece596862e75e46755df054811175f08b62bac)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-025.md#canonical-3c628a7bc1cdae78492582bbcf546221a1af0a6a9879bb9cc9756246e42ba22d)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-025.md#canonical-4b6b4795a6f2233496dc880bd1e62b3368648a7664d327459eca0169422c8c2b)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca

<a id="canonical-39fada2b6b67d4a95d4fcbfbde21e2ed2b5b244316b9841dedffcc5bb131a2fe"></a>

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
trusted_ca {
  # Configure direct properties listed below.
}
```

<a id="canonical-56180fe3d05a9359c5b8ac46556724dc70aa3755c08c01348f0183178f2e39c9"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / e2786a5e4eaf / 3

<a id="canonical-f6d00f67ee4786977a27d2ceeaf2bfcefb5ef7bbb4f8b3b10f2552e4a7892e2d"></a>

<a id="canonical-737da0b64cff8cd5d5ef4e630398a34f779bc3e5c2548e11229bf1a372e1bbbd"></a>

## name property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / e2786a5e4eaf / 4

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

<a id="canonical-2d4fda43dd4bb9bd67bc6c166efad04abdca77f1ad74f787733d07b2d0dbf97a"></a>

<a id="canonical-37e9d5c97c9ab64dfe4d3058ef0f6ad17e246b6f4c7e97def4516956115463da"></a>

## namespace property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / e2786a5e4eaf / 5

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

<a id="canonical-db550e2ae590ddd6bad3873cda13c680e9c7f4560f2ecfeef0fed1530ec1d65e"></a>

<a id="canonical-b290eb9960f5a0961622720d0a4acc0f9ad025eced7b1b390c3b40559f4a8f3c"></a>

## tenant property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / e2786a5e4eaf / 6

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

<a id="canonical-2c7c2cbfc341a46e13e2ae1d024e0efd83ba1377e018af96e482e9faac26714b"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / e2786a5e4eaf / 7

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-025.md#canonical-4b6b4795a6f2233496dc880bd1e62b3368648a7664d327459eca0169422c8c2b)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-c841cb951a1c4f9d7501b5f4a236ba845ed4c7d507aa5f4153aaf6e5db31bb32"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9f123ea9dd64e504166f1fbe03994cac88c6b14cbfa1500a58aba4041ad89612"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_disabled — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 102d8422be88 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-024.md#canonical-181a1518b8314bdeef4b5e79a2b8aacc9f75d1d5ed91dcbc4a205f9ad5b37385)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-024.md#canonical-a309460de31d5a477944352a93673e8a090b1d6c2342355139ca840878a2f2e6)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-024.md#canonical-b87765252f8d022952915dc691ece596862e75e46755df054811175f08b62bac)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-025.md#canonical-3c628a7bc1cdae78492582bbcf546221a1af0a6a9879bb9cc9756246e42ba22d)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-025.md#canonical-4b6b4795a6f2233496dc880bd1e62b3368648a7664d327459eca0169422c8c2b)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_disabled

<a id="canonical-736357d613f2abcf70571666a54a0257b278c94aa9b059a07c5aaa35976717aa"></a>

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
xfcc_disabled = {}
```

<a id="canonical-544a0033f43d072cfcd17983ac1bc57aa7e9228255470889d3080885bf333e6c"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 102d8422be88 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d068e58279d289c9c4cbe09280e8396ff56a78667a2dd8c8f0da6e8ea1470f1c"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 102d8422be88 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-025.md#canonical-4b6b4795a6f2233496dc880bd1e62b3368648a7664d327459eca0169422c8c2b)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-69786fb447c3ee5e0def78812708df796e9f9e9848d7fb7f2d060a9cdc4b7864"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-596758bcb143f857a17b20488c241d789f6c67f5bc6eed43d85ac2f68e6f9afa"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 688bfdec8489 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-024.md#canonical-181a1518b8314bdeef4b5e79a2b8aacc9f75d1d5ed91dcbc4a205f9ad5b37385)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-024.md#canonical-a309460de31d5a477944352a93673e8a090b1d6c2342355139ca840878a2f2e6)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-024.md#canonical-b87765252f8d022952915dc691ece596862e75e46755df054811175f08b62bac)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-025.md#canonical-3c628a7bc1cdae78492582bbcf546221a1af0a6a9879bb9cc9756246e42ba22d)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-025.md#canonical-4b6b4795a6f2233496dc880bd1e62b3368648a7664d327459eca0169422c8c2b)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options

<a id="canonical-fc9859d7d2afcf6ff38e2e88095076fb6114d556027c7b367ffb4a68009f2a02"></a>

Type: `"object"`. single nested block, Optional.

X-Forwarded-Client-Cert header elements to be added to requests.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("xfcc_header_elements")}
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
xfcc_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-3dfda02ea1ab0127a281ed11baaa6bb148cb9f491e9535f933c651be21120e9a"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 688bfdec8489 / 3

<a id="canonical-fcfd00a63d065b13f0e8a734d5190436a2be1a898bd625a7da7805e99c56fe30"></a>

<a id="canonical-ebe7bd97bd41f5f32aa4b4fae77ebcb2ec9991bc70a9eefa2a28c64f6518a0e9"></a>

## xfcc_header_elements property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 688bfdec8489 / 4

Type: `["list", "string"]`. Optional.

\[Enum: XFCC\_NONE|XFCC\_CERT|XFCC\_CHAIN|XFCC\_SUBJECT|XFCC\_URI|XFCC\_DNS\]
X-Forwarded-Client-Cert header elements to be added to requests. Possible values are \`XFCC\_NONE\`,
\`XFCC\_CERT\`, \`XFCC\_CHAIN\`, \`XFCC\_SUBJECT\`, \`XFCC\_URI\`, \`XFCC\_DNS\`. Defaults to
\`XFCC\_NONE\`.

Upstream description:

X-Forwarded-Client-Cert header elements to be added to requests.

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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  }
}
```

<a id="canonical-f06b4e328b5f8393ac78b8c1a50c772b7042b4e7b68885b9ae71d261c9349987"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 688bfdec8489 / 5

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-025.md#canonical-4b6b4795a6f2233496dc880bd1e62b3368648a7664d327459eca0169422c8c2b)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-a8beac99657fd081b6488810eaaffaf6ca1d8f3a81b91771a1ebcb59baf6428d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-25b7ed0eec85e2b35d40a356e4e672af712080979b1b64fc33ad35a59cc1c00c"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 52bc8806b49b / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-024.md#canonical-181a1518b8314bdeef4b5e79a2b8aacc9f75d1d5ed91dcbc4a205f9ad5b37385)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-024.md#canonical-a309460de31d5a477944352a93673e8a090b1d6c2342355139ca840878a2f2e6)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-024.md#canonical-b87765252f8d022952915dc691ece596862e75e46755df054811175f08b62bac)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters

<a id="canonical-e64bce32215ae75e994f8b6c77c09b42149a07af16adaf4eccb46104847504cf"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for tls parameters.

Upstream description:

Inline TLS parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("tls_certificates"),
  validators.ConflictingObjectAttributes("no_mtls",
    "use_mtls")}
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
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

Terraform syntax:

```terraform
tls_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-ab465c7d04d5e5f82863c6b13d6492acc00f2da224ce9df414f655549e8f87c6"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 52bc8806b49b / 3

- [no_mtls](resources--workload--reference--group-025.md#canonical-3f3e649fd68dc530401e0bef931a1ddb84404e4ceb340c81ce5444eb88d4a7e5): complete subsection reference.

- [tls_certificates](resources--workload--reference--group-025.md#canonical-cebe02ca1ebf24bdbd590ffaec07d6074e348ae8fdf28cc27398c0ea071b933c): complete subsection reference.

- [tls_config](resources--workload--reference--group-025.md#canonical-a94ee613158932bb5c9558c469570b93f889c5164771a837ffc7d0c469e15c86): complete subsection reference.

- [use_mtls](resources--workload--reference--group-026.md#canonical-3e6d344c978c9f34e7adda74cf47dd6c1119580937563c05f07d5f40daaa838a): complete subsection reference.

<a id="canonical-3c8bea1bc440c67ad30d81d79948751e3e039b1a7123cf380e24fdbfdb64037d"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 52bc8806b49b / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.no_mtls](resources--workload--reference--group-025.md#canonical-3f3e649fd68dc530401e0bef931a1ddb84404e4ceb340c81ce5444eb88d4a7e5)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-025.md#canonical-cebe02ca1ebf24bdbd590ffaec07d6074e348ae8fdf28cc27398c0ea071b933c)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-025.md#canonical-a94ee613158932bb5c9558c469570b93f889c5164771a837ffc7d0c469e15c86)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-026.md#canonical-3e6d344c978c9f34e7adda74cf47dd6c1119580937563c05f07d5f40daaa838a)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-024.md#canonical-b87765252f8d022952915dc691ece596862e75e46755df054811175f08b62bac)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-3f3e649fd68dc530401e0bef931a1ddb84404e4ceb340c81ce5444eb88d4a7e5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f3d74c2e6e8e2cdc0e2e9c68d6a04b960ac082b3df2e72fac88cc6d51e08d88c"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.no_mtls — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 31859163d60b / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-024.md#canonical-181a1518b8314bdeef4b5e79a2b8aacc9f75d1d5ed91dcbc4a205f9ad5b37385)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-024.md#canonical-a309460de31d5a477944352a93673e8a090b1d6c2342355139ca840878a2f2e6)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-024.md#canonical-b87765252f8d022952915dc691ece596862e75e46755df054811175f08b62bac)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-025.md#canonical-a8beac99657fd081b6488810eaaffaf6ca1d8f3a81b91771a1ebcb59baf6428d)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.no_mtls

<a id="canonical-cf0006f9c3c894ae8f5713135669d6b31a25e65a54f86ec3d2ec3f9f4a6198dd"></a>

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
no_mtls = {}
```

<a id="canonical-c4d1d3e1f216ab726d998d5b84ebed25f09d071ec5455abd41ce2004c0e0f709"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 31859163d60b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e527ab5fd7fcbddcb2ee467d2b4408465d1030022eee80338d67ca96b91b8a00"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 31859163d60b / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-025.md#canonical-a8beac99657fd081b6488810eaaffaf6ca1d8f3a81b91771a1ebcb59baf6428d)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-cebe02ca1ebf24bdbd590ffaec07d6074e348ae8fdf28cc27398c0ea071b933c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bc3d01f2e6ceaeb1709fddf26ce7bcf98d75846a503b19b0e100fc1803c6077c"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / bd01a89cfeb6 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-024.md#canonical-181a1518b8314bdeef4b5e79a2b8aacc9f75d1d5ed91dcbc4a205f9ad5b37385)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-024.md#canonical-a309460de31d5a477944352a93673e8a090b1d6c2342355139ca840878a2f2e6)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-024.md#canonical-b87765252f8d022952915dc691ece596862e75e46755df054811175f08b62bac)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-025.md#canonical-a8beac99657fd081b6488810eaaffaf6ca1d8f3a81b91771a1ebcb59baf6428d)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates

<a id="canonical-7659fe97574c61dd7e4790d8a17f973814e51fa1b788fe55e0d8dbde13699b1d"></a>

Type: `"object"`. list nested block, Optional.

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

Upstream description:

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("certificate_url"),
  validators.ConflictingListObjectAttributes("custom_hash_algorithms",
    "disable_ocsp_stapling"),
  validators.ConflictingListObjectAttributes("custom_hash_algorithms",
    "use_system_defaults"),
  validators.ConflictingListObjectAttributes("disable_ocsp_stapling",
    "use_system_defaults")}
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
tls_certificates {
  # Configure direct properties listed below.
}
```

<a id="canonical-e2c16106f934b0f06b58675bdb0c26d8a9ad277a10dd7cfae8b38cf77fd56c8b"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / bd01a89cfeb6 / 3

<a id="canonical-8345a078bd4d76ebc151b94a6e98ed1c4709dba629357bc291dbf5333c9ccf99"></a>

<a id="canonical-fd5382f3fef201a5cc924ea8369b831155a95af5f76a09b3448be0dfd2067a35"></a>

## certificate_url property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / bd01a89cfeb6 / 4

Type: `"string"`. Optional.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Upstream description:

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

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
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [custom_hash_algorithms](resources--workload--reference--group-025.md#canonical-eefe54c984c3228c3cf5262c5c4f5908493136732614f85883b60ef1f1be7807): complete subsection reference.

<a id="canonical-a18955bd8810d5a7917eaf0225b343e10fa3df66e6e28c1c3fb0c6767f3e2b0d"></a>

<a id="canonical-2ee6d242ac52a4a9c1f081f824db831478752d0dfc8fe41dbc94fe72782adba8"></a>

## description_spec property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / bd01a89cfeb6 / 5

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--workload--reference--group-025.md#canonical-7d598b4a65a8e1185817df5e66330b7fef55a4bf9d433c857d85310d2d1ebb08): complete subsection reference.

- [private_key](resources--workload--reference--group-025.md#canonical-618ed527f5e22a95f053dc2d0e96a0587bcabf4d30905c2dcd451864da8cad33): complete subsection reference.

- [use_system_defaults](resources--workload--reference--group-025.md#canonical-2758e3c3ab5bcfca1d3ec8a7540b945b0cccce5b4b7e42ef90a32d05984911ab): complete subsection reference.

<a id="canonical-14da95e33e2b5af2238e6020b8c528463d22739127c313bbbdb5bb23c9598303"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / bd01a89cfeb6 / 6

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms](resources--workload--reference--group-025.md#canonical-eefe54c984c3228c3cf5262c5c4f5908493136732614f85883b60ef1f1be7807)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.disable_ocsp_stapling](resources--workload--reference--group-025.md#canonical-7d598b4a65a8e1185817df5e66330b7fef55a4bf9d433c857d85310d2d1ebb08)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](resources--workload--reference--group-025.md#canonical-618ed527f5e22a95f053dc2d0e96a0587bcabf4d30905c2dcd451864da8cad33)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.use_system_defaults](resources--workload--reference--group-025.md#canonical-2758e3c3ab5bcfca1d3ec8a7540b945b0cccce5b4b7e42ef90a32d05984911ab)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-025.md#canonical-a8beac99657fd081b6488810eaaffaf6ca1d8f3a81b91771a1ebcb59baf6428d)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-eefe54c984c3228c3cf5262c5c4f5908493136732614f85883b60ef1f1be7807"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-667bd945c885e08ec51dfc979fedadb5cf84bfdcb52070cb324750771ae23ef4"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / f46de3104437 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-024.md#canonical-181a1518b8314bdeef4b5e79a2b8aacc9f75d1d5ed91dcbc4a205f9ad5b37385)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-024.md#canonical-a309460de31d5a477944352a93673e8a090b1d6c2342355139ca840878a2f2e6)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-024.md#canonical-b87765252f8d022952915dc691ece596862e75e46755df054811175f08b62bac)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-025.md#canonical-a8beac99657fd081b6488810eaaffaf6ca1d8f3a81b91771a1ebcb59baf6428d)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-025.md#canonical-cebe02ca1ebf24bdbd590ffaec07d6074e348ae8fdf28cc27398c0ea071b933c)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms

<a id="canonical-bb89d0e82ac2b37557d436b73756a95b07c948149954521aa8da5434fa896b95"></a>

Type: `"object"`. single nested block, Optional.

Specifies the hash algorithms to be used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("hash_algorithms")}
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
custom_hash_algorithms {
  # Configure direct properties listed below.
}
```

<a id="canonical-e078c35def6ec4d3c9146a304a7449d04133da7afda4c4ab535462c8664c0ae8"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / f46de3104437 / 3

<a id="canonical-77e32b9395a9fd80e6c413f9fe1e8b49d8f1426a927d6bb1214fe524681a3b60"></a>

<a id="canonical-9ae3ce8848ef27d934950e7bc7449f1345406d1136f37876d07cefba79bac954"></a>

## hash_algorithms property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / f46de3104437 / 4

Type: `["list", "string"]`. Optional.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

Upstream description:

Ordered list of hash algorithms to be used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 4),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-7bc6c711271d1d1e89c5e84d3d50360ee2d6d023398bf6c1141b1e804fa259df"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / f46de3104437 / 5

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-025.md#canonical-cebe02ca1ebf24bdbd590ffaec07d6074e348ae8fdf28cc27398c0ea071b933c)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-7d598b4a65a8e1185817df5e66330b7fef55a4bf9d433c857d85310d2d1ebb08"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-08111d9e7e5fb0242a9d75b8585bb2320f4c2aa258b378e10b1034a2da15c533"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.disable_ocsp_stapling — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / f0dd357436c9 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-024.md#canonical-181a1518b8314bdeef4b5e79a2b8aacc9f75d1d5ed91dcbc4a205f9ad5b37385)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-024.md#canonical-a309460de31d5a477944352a93673e8a090b1d6c2342355139ca840878a2f2e6)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-024.md#canonical-b87765252f8d022952915dc691ece596862e75e46755df054811175f08b62bac)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-025.md#canonical-a8beac99657fd081b6488810eaaffaf6ca1d8f3a81b91771a1ebcb59baf6428d)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-025.md#canonical-cebe02ca1ebf24bdbd590ffaec07d6074e348ae8fdf28cc27398c0ea071b933c)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.disable_ocsp_stapling

<a id="canonical-84cf65804db39fc405a6094b73fdba306a1f38c83bc30fc232e11f6a97e4b611"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable ocsp stapling.

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
disable_ocsp_stapling = {}
```

<a id="canonical-ae6a4c34425367725b648d475a09dae69f531a7b3f44a72e5b2bce5c32fac524"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / f0dd357436c9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-184c2ca84762198ffd0ae5caa3e9d50b610b4bbe2b59d62e3c0f4d4cb61b6399"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / f0dd357436c9 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-025.md#canonical-cebe02ca1ebf24bdbd590ffaec07d6074e348ae8fdf28cc27398c0ea071b933c)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-618ed527f5e22a95f053dc2d0e96a0587bcabf4d30905c2dcd451864da8cad33"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ddd77e88caf2403cebd6262ab13ba7412804378b1383ed227a50c1f57a91c4b5"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 29b35cc8d110 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-024.md#canonical-181a1518b8314bdeef4b5e79a2b8aacc9f75d1d5ed91dcbc4a205f9ad5b37385)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-024.md#canonical-a309460de31d5a477944352a93673e8a090b1d6c2342355139ca840878a2f2e6)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-024.md#canonical-b87765252f8d022952915dc691ece596862e75e46755df054811175f08b62bac)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-025.md#canonical-a8beac99657fd081b6488810eaaffaf6ca1d8f3a81b91771a1ebcb59baf6428d)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-025.md#canonical-cebe02ca1ebf24bdbd590ffaec07d6074e348ae8fdf28cc27398c0ea071b933c)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key

<a id="canonical-9485c4704081acab53ed874556ccef3495163e135469942ae7898f29feb1db18"></a>

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
private_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-f2364fed77ef5415a8c6b587bda9fc27a90ff3695cbcab9ca20d258f1ed22e51"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 29b35cc8d110 / 3

- [blindfold_secret_info](resources--workload--reference--group-025.md#canonical-487e6220a095e3aae4b59c4fb19793464fc67dafe22b3aef764fd864e2c76c0c): complete subsection reference.

- [clear_secret_info](resources--workload--reference--group-025.md#canonical-eea79922f2c039df9ba42bfa40759cd6210b9b2b403ecc1aeffbc40ded3874ae): complete subsection reference.

<a id="canonical-c8853b954be4dec263eaf29591ed70eb3311cf4a0db50671b1535e0af51d0312"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 29b35cc8d110 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info](resources--workload--reference--group-025.md#canonical-487e6220a095e3aae4b59c4fb19793464fc67dafe22b3aef764fd864e2c76c0c)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info](resources--workload--reference--group-025.md#canonical-eea79922f2c039df9ba42bfa40759cd6210b9b2b403ecc1aeffbc40ded3874ae)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-025.md#canonical-cebe02ca1ebf24bdbd590ffaec07d6074e348ae8fdf28cc27398c0ea071b933c)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-487e6220a095e3aae4b59c4fb19793464fc67dafe22b3aef764fd864e2c76c0c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4c5e7307a777c0f3fc80d29f6c88d5c0b1b208d6dadfd7967813258db77d3ca2"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 3af7a834582b / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-024.md#canonical-181a1518b8314bdeef4b5e79a2b8aacc9f75d1d5ed91dcbc4a205f9ad5b37385)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-024.md#canonical-a309460de31d5a477944352a93673e8a090b1d6c2342355139ca840878a2f2e6)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-024.md#canonical-b87765252f8d022952915dc691ece596862e75e46755df054811175f08b62bac)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-025.md#canonical-a8beac99657fd081b6488810eaaffaf6ca1d8f3a81b91771a1ebcb59baf6428d)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-025.md#canonical-cebe02ca1ebf24bdbd590ffaec07d6074e348ae8fdf28cc27398c0ea071b933c)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](resources--workload--reference--group-025.md#canonical-618ed527f5e22a95f053dc2d0e96a0587bcabf4d30905c2dcd451864da8cad33)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-bcec6549d179ef75b30d39535ad8c8a56f4a9939fd4cf168905a5f0706ad4ee0"></a>

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

<a id="canonical-b86609d2d3656e4f09d50183360cc38916e4faf1cb95f7f7aa97304d10329718"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 3af7a834582b / 3

<a id="canonical-124efa92f7d5bbe2c2fb95cad74985eca951c30c34ad97962540fc3c00de8ea8"></a>

<a id="canonical-b82c7fcd9d5a9cd5b53cdf7507c1ad04424a5696648465faae7c3848a34292c6"></a>

## decryption_provider property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 3af7a834582b / 4

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

<a id="canonical-a574d990a028fc033e38bc85502f4ff6746945a09a29e903b4ce23b952576706"></a>

<a id="canonical-a3b0a287ca39fafb7fa0dd4ee010541de0bb3ace69dddae50ee06d190e08059c"></a>

## location property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 3af7a834582b / 5

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

<a id="canonical-94b57d3d7a9cac3c4fc8a0c77124a2de00c495cb476a07d99c4108a8fa0df97a"></a>

<a id="canonical-84d3f13c270a63ea399e5fb66ab89858b355377595304ba9590d201d568ea550"></a>

## store_provider property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 3af7a834582b / 6

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

<a id="canonical-12b573c1da501da896a196b3befa78c55d7614c808045e7672f064216645d074"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 3af7a834582b / 7

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](resources--workload--reference--group-025.md#canonical-618ed527f5e22a95f053dc2d0e96a0587bcabf4d30905c2dcd451864da8cad33)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-eea79922f2c039df9ba42bfa40759cd6210b9b2b403ecc1aeffbc40ded3874ae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a83706d6934037e4ee188a3502a9400137d1a0374a5e074b9792768d007b2e56"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 9f0d836d47e2 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-024.md#canonical-181a1518b8314bdeef4b5e79a2b8aacc9f75d1d5ed91dcbc4a205f9ad5b37385)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-024.md#canonical-a309460de31d5a477944352a93673e8a090b1d6c2342355139ca840878a2f2e6)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-024.md#canonical-b87765252f8d022952915dc691ece596862e75e46755df054811175f08b62bac)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-025.md#canonical-a8beac99657fd081b6488810eaaffaf6ca1d8f3a81b91771a1ebcb59baf6428d)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-025.md#canonical-cebe02ca1ebf24bdbd590ffaec07d6074e348ae8fdf28cc27398c0ea071b933c)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](resources--workload--reference--group-025.md#canonical-618ed527f5e22a95f053dc2d0e96a0587bcabf4d30905c2dcd451864da8cad33)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info

<a id="canonical-856c9844cc12955dab51f5bc333735da14ee44382dad87eedac6ba9564fd9c2d"></a>

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

<a id="canonical-e179c546bc921018dd1fd2b67756b39aca8f04b63626eb4e574f49d3af5abe9a"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 9f0d836d47e2 / 3

<a id="canonical-eeedd6b3c15b239f4821ea4877ae90b805773b85af128228770bdbf7f87d7abb"></a>

<a id="canonical-fc630b9c84299fe40ff56bb5629b971405ddd52b033ded6b3cb527c1edfd8e76"></a>

## provider_ref property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 9f0d836d47e2 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-709fa2dc0ac43020bb87152e30e35229e8b37f1cc412e9724f7e6599b363fdaa"></a>

<a id="canonical-909b9721a53d26a176107303c022337ccba9f9e6f80bafb5c872f49942eb9d2b"></a>

## url property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 9f0d836d47e2 / 5

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

<a id="canonical-b0d48326aa6cc6ee50423dd5b5203c49baf0b8d84ad358419d3241ebb0c634f2"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 9f0d836d47e2 / 6

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](resources--workload--reference--group-025.md#canonical-618ed527f5e22a95f053dc2d0e96a0587bcabf4d30905c2dcd451864da8cad33)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-2758e3c3ab5bcfca1d3ec8a7540b945b0cccce5b4b7e42ef90a32d05984911ab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fc85b3201e42eab23b08894264f3d1d5fd2b90204ed28c229c8849d3e13defe3"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.use_system_defaults — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 53afcf59631c / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-024.md#canonical-181a1518b8314bdeef4b5e79a2b8aacc9f75d1d5ed91dcbc4a205f9ad5b37385)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-024.md#canonical-a309460de31d5a477944352a93673e8a090b1d6c2342355139ca840878a2f2e6)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-024.md#canonical-b87765252f8d022952915dc691ece596862e75e46755df054811175f08b62bac)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-025.md#canonical-a8beac99657fd081b6488810eaaffaf6ca1d8f3a81b91771a1ebcb59baf6428d)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-025.md#canonical-cebe02ca1ebf24bdbd590ffaec07d6074e348ae8fdf28cc27398c0ea071b933c)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.use_system_defaults

<a id="canonical-7b1895a003978de71014a6c0ef7d4d940346431ad6804e663d98c0a3e9afe9c6"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for use system defaults.

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
use_system_defaults = {}
```

<a id="canonical-7d303349c8385965b0cb3eab53849fd41ea7cc2b1ce65b1a5b7ffb28252d8287"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 53afcf59631c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5c0094ea7c00fc3c4c28a8e33218553d5541f4b0a8f42c93fbeb75b80672bcb5"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 53afcf59631c / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-025.md#canonical-cebe02ca1ebf24bdbd590ffaec07d6074e348ae8fdf28cc27398c0ea071b933c)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-a94ee613158932bb5c9558c469570b93f889c5164771a837ffc7d0c469e15c86"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-daec9dac5870c9b65028ef5d5fc936bc8b36fd82ec9142678404d4b551cf4a7e"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 330fd2a8b7a6 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-024.md#canonical-181a1518b8314bdeef4b5e79a2b8aacc9f75d1d5ed91dcbc4a205f9ad5b37385)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-024.md#canonical-a309460de31d5a477944352a93673e8a090b1d6c2342355139ca840878a2f2e6)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-024.md#canonical-b87765252f8d022952915dc691ece596862e75e46755df054811175f08b62bac)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-025.md#canonical-a8beac99657fd081b6488810eaaffaf6ca1d8f3a81b91771a1ebcb59baf6428d)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config

<a id="canonical-ed5fea328e52ee3373a2fbf1c53d279ed6c0e3135f5463a131a2a318b22ebf92"></a>

Type: `"object"`. single nested block, Optional.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_security",
    "default_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "low_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("default_security",
    "low_security"),
  validators.ConflictingObjectAttributes("default_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("low_security",
    "medium_security")}
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
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

Terraform syntax:

```terraform
tls_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-f8905178802f03821c0719598a76c863f65a2af59a6f2afbc584d33b6c1f0914"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 330fd2a8b7a6 / 3

- [custom_security](resources--workload--reference--group-025.md#canonical-0ac66cee4b2151f7aa95ff4d64e7959a28a65ab97365bd52ec24f0b0ff21c8a7): complete subsection reference.

- [default_security](resources--workload--reference--group-025.md#canonical-25b62dc0dfe600af6f97902ca03a5dfcec7d0b708ebb02760c1c4b3bf50c2514): complete subsection reference.

- [low_security](resources--workload--reference--group-025.md#canonical-2a8043cd2b255d0ba5df2939fee3153ee8d91fa1238ead406ad9a52620975a8c): complete subsection reference.

- [medium_security](resources--workload--reference--group-026.md#canonical-d682e8c1112c788c962b6d04b14a4932687e7ceba5336090a4af3a2b0f0910a5): complete subsection reference.

<a id="canonical-626b5c0c9440e77f38d41503ac48e5ca8acab957ab8521ee7c98bc3f6c201309"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 330fd2a8b7a6 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.custom_security](resources--workload--reference--group-025.md#canonical-0ac66cee4b2151f7aa95ff4d64e7959a28a65ab97365bd52ec24f0b0ff21c8a7)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.default_security](resources--workload--reference--group-025.md#canonical-25b62dc0dfe600af6f97902ca03a5dfcec7d0b708ebb02760c1c4b3bf50c2514)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.low_security](resources--workload--reference--group-025.md#canonical-2a8043cd2b255d0ba5df2939fee3153ee8d91fa1238ead406ad9a52620975a8c)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.medium_security](resources--workload--reference--group-026.md#canonical-d682e8c1112c788c962b6d04b14a4932687e7ceba5336090a4af3a2b0f0910a5)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-025.md#canonical-a8beac99657fd081b6488810eaaffaf6ca1d8f3a81b91771a1ebcb59baf6428d)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-0ac66cee4b2151f7aa95ff4d64e7959a28a65ab97365bd52ec24f0b0ff21c8a7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ce1f48a3da3211cd2ab215eb25c06b3cfd16773690d660e8bf4266235a79955c"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.custom_security — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 061a5822c19b / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-024.md#canonical-181a1518b8314bdeef4b5e79a2b8aacc9f75d1d5ed91dcbc4a205f9ad5b37385)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-024.md#canonical-a309460de31d5a477944352a93673e8a090b1d6c2342355139ca840878a2f2e6)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-024.md#canonical-b87765252f8d022952915dc691ece596862e75e46755df054811175f08b62bac)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-025.md#canonical-a8beac99657fd081b6488810eaaffaf6ca1d8f3a81b91771a1ebcb59baf6428d)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-025.md#canonical-a94ee613158932bb5c9558c469570b93f889c5164771a837ffc7d0c469e15c86)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.custom_security

<a id="canonical-2b2d67b9aaed561cb4cab5c089f180c3c34494f253f422e0c9ccb6d32f5c42a0"></a>

Type: `"object"`. single nested block, Optional.

Defines TLS protocol config including min/max versions and allowed ciphers.

Upstream description:

This defines TLS protocol config including min/max versions and allowed ciphers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cipher_suites")}
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
custom_security {
  # Configure direct properties listed below.
}
```

<a id="canonical-ac7d16c5dd97c99a2b056ebeacb88b1c7dd3c20a95d407ab0308904090d7cb2d"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 061a5822c19b / 3

<a id="canonical-b85b4c7ce9695f5b9bd36e99de57537ba67976ad017881738af00b272369ac40"></a>

<a id="canonical-6b74951ac434b6ae6941ff68ad9bb2e6db164d3e6fb7fda032538f5656b25413"></a>

## cipher_suites property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 061a5822c19b / 4

Type: `["list", "string"]`. Optional.

The TLS listener will only support the specified cipher list.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-c36bdce9afc9e8824345422f59bc03e4e69c0fe123cd14d51eeb13968a348155"></a>

<a id="canonical-67755d0359ecee95df3cd6096d7caeddc4ce73ade18d52c692b8601efb9707db"></a>

## max_version property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 061a5822c19b / 5

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

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

<a id="canonical-1ae46d99a3a983a242c6e3b52ee8c3c12e311e07d898d07a4632df7bdefcd014"></a>

<a id="canonical-6f3203e5c0ea5eb5672768743dbfa06598055ee3bd638822533f0d10ca19e20c"></a>

## min_version property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 061a5822c19b / 6

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

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

<a id="canonical-6c82b69a61aa9326fc2996399ed794daae3ec46bb2cd8cee8023fa0f3b40ec3b"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 061a5822c19b / 7

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-025.md#canonical-a94ee613158932bb5c9558c469570b93f889c5164771a837ffc7d0c469e15c86)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-25b62dc0dfe600af6f97902ca03a5dfcec7d0b708ebb02760c1c4b3bf50c2514"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-beb35825b6efac2387fa761d872311f80890e86ed51bd76dc02c1fff4897bf7d"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.default_security — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 45c52921aac8 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-024.md#canonical-181a1518b8314bdeef4b5e79a2b8aacc9f75d1d5ed91dcbc4a205f9ad5b37385)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-024.md#canonical-a309460de31d5a477944352a93673e8a090b1d6c2342355139ca840878a2f2e6)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-024.md#canonical-b87765252f8d022952915dc691ece596862e75e46755df054811175f08b62bac)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-025.md#canonical-a8beac99657fd081b6488810eaaffaf6ca1d8f3a81b91771a1ebcb59baf6428d)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-025.md#canonical-a94ee613158932bb5c9558c469570b93f889c5164771a837ffc7d0c469e15c86)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.default_security

<a id="canonical-e9c985b96c81441c1e8f0f7c320dd7fa186217f9703fcbe092604cea5badd206"></a>

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
default_security = {}
```

<a id="canonical-68dd02a07bc55ce8f5d405ae0946e75f6c9fc727febbc7048eb364e490cc8e9d"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 45c52921aac8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b6162c89989eca900867bd96fb26fd81e587f910a4fdaf0e58042bfeb8a86620"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 45c52921aac8 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-025.md#canonical-a94ee613158932bb5c9558c469570b93f889c5164771a837ffc7d0c469e15c86)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-2a8043cd2b255d0ba5df2939fee3153ee8d91fa1238ead406ad9a52620975a8c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
