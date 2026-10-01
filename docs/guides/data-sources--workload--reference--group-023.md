---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-579532e9800f50f73a8beea7708a4fba89b01ded7429ddd72818cb5e719e9753"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 3226130dc64a / 6

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls.crl](data-sources--workload--reference--group-023.md#canonical-72c9ae85c6d7c1c9441815f4287691df1e2e2984d0e9850bb106edcba3233093)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls.no_crl](data-sources--workload--reference--group-023.md#canonical-65f7b89b229ae56736233e983d6ee84a3ea49139621115cd4fc0526e39a3d497)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca](data-sources--workload--reference--group-023.md#canonical-c25f679cf13053c21e8478bd715145451461b80c6055eba199793e7b405917b3)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_disabled](data-sources--workload--reference--group-023.md#canonical-212ea2af65c31256d594f079baeb4d219b208bb2e8bbab3c9c9819f3a2431ca3)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options](data-sources--workload--reference--group-023.md#canonical-633e70a684fd7390b2777fe958b675b0d21a9abc3e7ee2ad6a239c93f3c92472)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-d839a30793886c61fd06909d6937c45fe1a29205269cf1c2d29af93120a0ce44)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-72c9ae85c6d7c1c9441815f4287691df1e2e2984d0e9850bb106edcba3233093"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-560fae5e442e6c72b231e0288d334ec5f27ee43afacdd3f1625f867c5d679e31"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls.crl — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 89a3645702ec / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-d839a30793886c61fd06909d6937c45fe1a29205269cf1c2d29af93120a0ce44)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-022.md#canonical-feafac3c30f911a98096a4b3d0da61b9fe41e35c73a4efce957bc606373292cd)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls.crl

<a id="canonical-5664fdb44516af89ea1b0d7acd885a37556fe2fcb18455a4a8583427b0ee3112"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-dd14dfaa32c4f1868086f48972c402e80c12e9f8ed563af615e1203b8f845fbd"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 89a3645702ec / 3

<a id="canonical-508f1b370a81ea981c9ac5c751407ff4a59f6fbf51760ffe121111100bf862d1"></a>

<a id="canonical-c3df465093fd66cf7914aa3b0518c69d10ac0d924b839d85c3e054fdeb688930"></a>

## name property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 89a3645702ec / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-2c59c73e521fdc975e1990705304dd6943637a94eb2c09616cb3084fc1c16959"></a>

<a id="canonical-55a83faea23819e972832ba34aff61f349cfe2997d2f258e83130610ab0c1354"></a>

## namespace property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 89a3645702ec / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-4f512b82ed54341472148cf84995702b139a7f5a881d49e39155e5319b921a83"></a>

<a id="canonical-c5c540eaf41124b117a438c7f9da4a1974f0e426b09973b4819df5bf6aaeecac"></a>

## tenant property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 89a3645702ec / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-334a0e6ee2b2250b1a51e6e7e5916dfc60cd2776dbc4828cf75b1ad5a8bed678"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 89a3645702ec / 7

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-022.md#canonical-feafac3c30f911a98096a4b3d0da61b9fe41e35c73a4efce957bc606373292cd)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-65f7b89b229ae56736233e983d6ee84a3ea49139621115cd4fc0526e39a3d497"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e4024777bf7bf90d5239014740711c31c390f643bb5390acc5483e7d7e296f21"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls.no_crl — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 743627e536af / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-d839a30793886c61fd06909d6937c45fe1a29205269cf1c2d29af93120a0ce44)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-022.md#canonical-feafac3c30f911a98096a4b3d0da61b9fe41e35c73a4efce957bc606373292cd)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls.no_crl

<a id="canonical-5568d9a2f406d6f4238437be0e59535be8d673dc348fc7463e159a31641b24a9"></a>

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

<a id="canonical-805f2473dc4ed9dac0b8dfe577192134b54e9913ef7ae5abef24df8ed2e75a00"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 743627e536af / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-08e3e903be3afa39fce16bcae5b3ebfb35c11badc71344188a27e7fadd5ff9fd"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 743627e536af / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-022.md#canonical-feafac3c30f911a98096a4b3d0da61b9fe41e35c73a4efce957bc606373292cd)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-c25f679cf13053c21e8478bd715145451461b80c6055eba199793e7b405917b3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8fa95115d8796a5504400aaaa24dc135061f52c22441f1d021623b5d656dedd5"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / b7a2c63a2f35 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-d839a30793886c61fd06909d6937c45fe1a29205269cf1c2d29af93120a0ce44)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-022.md#canonical-feafac3c30f911a98096a4b3d0da61b9fe41e35c73a4efce957bc606373292cd)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca

<a id="canonical-153d57c971b662344240f6bb7b99d1c236e27bea61f64b40f99da0c19d4998d1"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-ae07c5c6e7de05ea09f3669d585ebb94278e181f78058e9244f0e93c8656dc93"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / b7a2c63a2f35 / 3

<a id="canonical-7d43ee40d092705d4f290fbfb8e2e25f079432e113d3859dfb29f2e2c98c8b6c"></a>

<a id="canonical-6bd01f0791941d1cf89a4d137b87d066f4e2f852fea33fd76457d670e8f7e2f7"></a>

## name property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / b7a2c63a2f35 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-7b91a2d718c2df713d1ff0a63213cff8d8cbacd40ef06755c6f2fbf7d9d0de9c"></a>

<a id="canonical-2ff71967956bf51c20cd39c4b4811165344829321affd7f1ee78f4f12d1ddbf6"></a>

## namespace property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / b7a2c63a2f35 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-c18e846815ada1dc074535cd39c6306b303a01a9ca8df421e5791e071dd29fd1"></a>

<a id="canonical-ee778dab31246c80f17472c2b05c11f10d8122a62da0f5f6f443976cda0bfc2f"></a>

## tenant property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / b7a2c63a2f35 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-3ee23ee2d1ba7b0074735793056d9561e5a7958710ea7357c802e166d7787c09"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / b7a2c63a2f35 / 7

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-022.md#canonical-feafac3c30f911a98096a4b3d0da61b9fe41e35c73a4efce957bc606373292cd)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-212ea2af65c31256d594f079baeb4d219b208bb2e8bbab3c9c9819f3a2431ca3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c5af5e8b7e92116ab950b5dda4bdb84544ff8691b5d49871ef615376647ff2a4"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_disabled — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / e0025a2a0d1e / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-d839a30793886c61fd06909d6937c45fe1a29205269cf1c2d29af93120a0ce44)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-022.md#canonical-feafac3c30f911a98096a4b3d0da61b9fe41e35c73a4efce957bc606373292cd)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_disabled

<a id="canonical-28bfb4da9f01c85621b0a3c6e969452d07c1f03c2f9a24a45161e27194c5cd37"></a>

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

<a id="canonical-a3573bb2f7307163689186fed86354cca9310c34657b2d34602e36495df5b0b1"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / e0025a2a0d1e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c12e9d0e4e5cb65533a08d4878d52dacefc15d1004ca86fa5d4883701f65d984"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / e0025a2a0d1e / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-022.md#canonical-feafac3c30f911a98096a4b3d0da61b9fe41e35c73a4efce957bc606373292cd)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-633e70a684fd7390b2777fe958b675b0d21a9abc3e7ee2ad6a239c93f3c92472"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7192ee45254f4b504ca298e7f4ba829342716a7e9f642a3f5da09a323ace2184"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / a97e3949e0b2 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-d839a30793886c61fd06909d6937c45fe1a29205269cf1c2d29af93120a0ce44)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-022.md#canonical-feafac3c30f911a98096a4b3d0da61b9fe41e35c73a4efce957bc606373292cd)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options

<a id="canonical-2fcfdb2c9dd08cc6a6bda1c9760f8c68c22efe9921c437bb5de75ab1d02d3375"></a>

Type: `"single"`. Computed.

X-Forwarded-Client-Cert header elements to be added to requests.

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

<a id="canonical-adecdb32cb13b0f83b138defa37b32dfbb2bfe3b65a3be25f4487356d50b1554"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / a97e3949e0b2 / 3

<a id="canonical-5112cc74b20ccce01371b9377855323a6940b8016938360a5c57a18318af261d"></a>

<a id="canonical-9c61b70eba1e980d203fb74a9ae0f8e34fa4db5fe2ec0d611235589e59dc9837"></a>

## xfcc_header_elements property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / a97e3949e0b2 / 4

Type: `["list", "string"]`. Computed.

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

<a id="canonical-23f0b469c297e935263694d1fd93b5cd1ba42dbc11cc69ae56f970b4b2dd272c"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / a97e3949e0b2 / 5

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-022.md#canonical-feafac3c30f911a98096a4b3d0da61b9fe41e35c73a4efce957bc606373292cd)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-20d5fcf752fc72102d70fd1c28f153fc2f63ff15340f7a13c67250ba8c1a0917"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-83f57d1fd2ecfff37c54ca3ff2fb35fac35cf916e84bb99cb521891abae9a622"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 6fe7375cd71e / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes

<a id="canonical-b5654d1cb4dc3439c627d568ebb49392e3c40bb4684d53edc8913cbc9d191455"></a>

Type: `"single"`. Computed.

Defines various OPTIONS to define a route.

Upstream description:

This defines various OPTIONS to define a route.

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

<a id="canonical-25c077e9629311ca54d26fdd1536f592bc9c4fe53f7b846199a33f06958f98d3"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 6fe7375cd71e / 3

- [routes](data-sources--workload--reference--group-023.md#canonical-c07077f21b2816e8db5b47273a7b26ad2b3fb24187445b00b277c2108076cf6c): complete subsection reference.

<a id="canonical-3c4f74f55b4c4f82737bb4cd6323b9131f31a838715150e6597c1a350d90b60e"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 6fe7375cd71e / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-023.md#canonical-c07077f21b2816e8db5b47273a7b26ad2b3fb24187445b00b277c2108076cf6c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-c07077f21b2816e8db5b47273a7b26ad2b3fb24187445b00b277c2108076cf6c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f6879e04946719d13930fccb8a0902f68c3a9af9417ef2c06a6c58970f91cec0"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / de7210ec7d9f / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-023.md#canonical-20d5fcf752fc72102d70fd1c28f153fc2f63ff15340f7a13c67250ba8c1a0917)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes

<a id="canonical-f0281281a17f2c7e1c66769c0f19003eaf2dfbe5b16a248fc556956d6cb0e14e"></a>

Type: `"list"`. Computed.

Routes. Routes for this loadbalancer.

Upstream description:

Routes for this loadbalancer.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

<a id="canonical-7bacc608a42554e05ce7b2c8e39b3dffd1c13e1961c9086a1d892b8fce274075"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / de7210ec7d9f / 3

- [custom_route_object](data-sources--workload--reference--group-023.md#canonical-1b7e5a96c1a657424e601bef879e7586a1805dcbd4be59afee3f08c70b873257): complete subsection reference.

- [direct_response_route](data-sources--workload--reference--group-023.md#canonical-9150ada40cd3a65b4f474f3e76d0ed50d03b850009b246f2012864d9ca9f77f1): complete subsection reference.

- [redirect_route](data-sources--workload--reference--group-023.md#canonical-5e027771e4626c74f12b688baf95b40804037d7cdf062e04cac370092fc01f51): complete subsection reference.

- [simple_route](data-sources--workload--reference--group-023.md#canonical-aa1a4e20c821310874f143b2cde560cdb2dfe7962517bb612441fa2e0ccdae55): complete subsection reference.

<a id="canonical-beefe695d16e9e332a2bf8201d346756e3550d278797520c687ee8ea430ff936"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / de7210ec7d9f / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.custom_route_object](data-sources--workload--reference--group-023.md#canonical-1b7e5a96c1a657424e601bef879e7586a1805dcbd4be59afee3f08c70b873257)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route](data-sources--workload--reference--group-023.md#canonical-9150ada40cd3a65b4f474f3e76d0ed50d03b850009b246f2012864d9ca9f77f1)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-023.md#canonical-5e027771e4626c74f12b688baf95b40804037d7cdf062e04cac370092fc01f51)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route](data-sources--workload--reference--group-023.md#canonical-aa1a4e20c821310874f143b2cde560cdb2dfe7962517bb612441fa2e0ccdae55)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-023.md#canonical-20d5fcf752fc72102d70fd1c28f153fc2f63ff15340f7a13c67250ba8c1a0917)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-1b7e5a96c1a657424e601bef879e7586a1805dcbd4be59afee3f08c70b873257"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-44ac69ab6fa76e577efac14f7b9e70efb57197560f38ba5ceed36902d4031917"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.custom_route_object — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 244c21bb721e / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-023.md#canonical-20d5fcf752fc72102d70fd1c28f153fc2f63ff15340f7a13c67250ba8c1a0917)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-023.md#canonical-c07077f21b2816e8db5b47273a7b26ad2b3fb24187445b00b277c2108076cf6c)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.custom_route_object

<a id="canonical-380324796792a665303483c798f16fd079befb5cd8dcfe415ecbefe7c505e4dd"></a>

Type: `"single"`. Computed.

Custom route uses a route object created outside of this view.

Upstream description:

A custom route uses a route object created outside of this view.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-caching": "[\"caching_disable\",\"caching_inherit\"]"
}
```

<a id="canonical-b1da57c030b7f53b48eda7a2cc7c7fb4e73696f6e3621c544fa395b1bdbcb473"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 244c21bb721e / 3

- [caching_disable](data-sources--workload--reference--group-023.md#canonical-ac8b019a69d76ec05f131f56e3081cfa0bd53a65ab6a55b5ca05e29bbcb01eae): complete subsection reference.

- [caching_inherit](data-sources--workload--reference--group-023.md#canonical-6914d3d611f9d78adf82547b856bb6f2f58745c8554318c125ad5e41803c45c9): complete subsection reference.

- [route_ref](data-sources--workload--reference--group-023.md#canonical-d27de944453d41ec078deb671ed4ef617e8998336619c4b109ceb9caee97202d): complete subsection reference.

<a id="canonical-004f53d69a83467fd79ec89e70b0a23d3141915ba7e73cad45857d1d7a16d019"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 244c21bb721e / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_disable](data-sources--workload--reference--group-023.md#canonical-ac8b019a69d76ec05f131f56e3081cfa0bd53a65ab6a55b5ca05e29bbcb01eae)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_inherit](data-sources--workload--reference--group-023.md#canonical-6914d3d611f9d78adf82547b856bb6f2f58745c8554318c125ad5e41803c45c9)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref](data-sources--workload--reference--group-023.md#canonical-d27de944453d41ec078deb671ed4ef617e8998336619c4b109ceb9caee97202d)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-023.md#canonical-c07077f21b2816e8db5b47273a7b26ad2b3fb24187445b00b277c2108076cf6c)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-ac8b019a69d76ec05f131f56e3081cfa0bd53a65ab6a55b5ca05e29bbcb01eae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e418a49ac673433521f411af19029e2826d7e4a606f00225f1a3a4576058c8cf"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_disable — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / a633b20acb62 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-023.md#canonical-20d5fcf752fc72102d70fd1c28f153fc2f63ff15340f7a13c67250ba8c1a0917)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-023.md#canonical-c07077f21b2816e8db5b47273a7b26ad2b3fb24187445b00b277c2108076cf6c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.custom_route_object](data-sources--workload--reference--group-023.md#canonical-1b7e5a96c1a657424e601bef879e7586a1805dcbd4be59afee3f08c70b873257)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_disable

<a id="canonical-2901755136522efe576751ddf7063356725af26f0bc7257e42a40561a398a8a3"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for caching disable.

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

<a id="canonical-2c0da78f405e455d48e7e1ae3d95161c1eab0653bc681f40e6716e432433c099"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / a633b20acb62 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5e932870f214ffdc7203fd99ce37b80d61bee8a28994755484a6dc842c59898c"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / a633b20acb62 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.custom_route_object](data-sources--workload--reference--group-023.md#canonical-1b7e5a96c1a657424e601bef879e7586a1805dcbd4be59afee3f08c70b873257)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-6914d3d611f9d78adf82547b856bb6f2f58745c8554318c125ad5e41803c45c9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c434e8c5dcf9468ca13011a08df2d5009b78967bac422e81f262a5c82b5ec94f"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_inherit — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 8d2b3bfa6739 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-023.md#canonical-20d5fcf752fc72102d70fd1c28f153fc2f63ff15340f7a13c67250ba8c1a0917)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-023.md#canonical-c07077f21b2816e8db5b47273a7b26ad2b3fb24187445b00b277c2108076cf6c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.custom_route_object](data-sources--workload--reference--group-023.md#canonical-1b7e5a96c1a657424e601bef879e7586a1805dcbd4be59afee3f08c70b873257)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_inherit

<a id="canonical-7c66841af2d48cd4a9c93f25d634bb89ceb7842fcdd2640ea0157d99519c4d92"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for caching inherit.

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

<a id="canonical-47ae0850569f75ead5799d1eef9328ddc6c58b6db8a8f5744049d2b42c39ad5d"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 8d2b3bfa6739 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d1f95bd1132bbbef21e6236b95d2fff58ebd7a5661fccb1183b368711362678d"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 8d2b3bfa6739 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.custom_route_object](data-sources--workload--reference--group-023.md#canonical-1b7e5a96c1a657424e601bef879e7586a1805dcbd4be59afee3f08c70b873257)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-d27de944453d41ec078deb671ed4ef617e8998336619c4b109ceb9caee97202d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-85bbaa3d92eaed285a7de9b577a42bc75762475b33998405a7f162b9d4a20a0a"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 1c6aee13a751 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-023.md#canonical-20d5fcf752fc72102d70fd1c28f153fc2f63ff15340f7a13c67250ba8c1a0917)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-023.md#canonical-c07077f21b2816e8db5b47273a7b26ad2b3fb24187445b00b277c2108076cf6c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.custom_route_object](data-sources--workload--reference--group-023.md#canonical-1b7e5a96c1a657424e601bef879e7586a1805dcbd4be59afee3f08c70b873257)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref

<a id="canonical-a19eb6d2775cda01d7548ff3b2ec416f20512b6c15a7c1a2eae9624df5cfb69c"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-4bf51d6a65fd8110ec674762269a28de936846eb9a211efce4b4029446a78e69"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 1c6aee13a751 / 3

<a id="canonical-3b3a429f88d8f5624a351321eefdc1741e87341a6e4cce8e1ab4315c982fe13b"></a>

<a id="canonical-41d71458c6b4fd0a32cda826dafbde56eefc8a234ac89990053c704c05858726"></a>

## name property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 1c6aee13a751 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-74c59d7025d5c682a003c8b1da15574c4e511ccf9cf4f2dd79cab298bf2f4000"></a>

<a id="canonical-72bb509349c6c2de13364ca2698a4e024387ef8cbf3e61b055f3254aff52c4c3"></a>

## namespace property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 1c6aee13a751 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-568e01665a53e4fb36acfce420a9fa18a4a735df46c97c69a3e747af1a7eb4ea"></a>

<a id="canonical-778079d045ff0e96fec155094051711555baf7ec87ece2f224d36c792d916803"></a>

## tenant property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 1c6aee13a751 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-e136cd3afb14083abeac83e19a9d2117a972259b48299b3c81279d4152f1cb85"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 1c6aee13a751 / 7

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.custom_route_object](data-sources--workload--reference--group-023.md#canonical-1b7e5a96c1a657424e601bef879e7586a1805dcbd4be59afee3f08c70b873257)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-9150ada40cd3a65b4f474f3e76d0ed50d03b850009b246f2012864d9ca9f77f1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-61420f3250e65cc7d9856ade6fa0b61cfce2364bca7fbfeffeafda3a77a91bf9"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 9c11326906cc / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-023.md#canonical-20d5fcf752fc72102d70fd1c28f153fc2f63ff15340f7a13c67250ba8c1a0917)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-023.md#canonical-c07077f21b2816e8db5b47273a7b26ad2b3fb24187445b00b277c2108076cf6c)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route

<a id="canonical-cea4c81f3e8503914f1a1091fb9bb10c4c7a751ecba6a27756fcad41eb827889"></a>

Type: `"single"`. Computed.

Direct response route matches on path, incoming header, incoming port and/or HTTP method and
responds directly to the matching traffic.

Upstream description:

A direct response route matches on path, incoming header, incoming port and/or HTTP method and
responds directly to the matching traffic.

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

<a id="canonical-967f76e375458ba6f479d47eb72be5f1e31b7016677b9fedf3b10356e9e295db"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 9c11326906cc / 3

- [headers](data-sources--workload--reference--group-023.md#canonical-69a56ccce607a77b681f38d7b4dbc732ba3ff1ace362cb1df470298ab287741e): complete subsection reference.

<a id="canonical-8b1faf0746db25502a90de1a6a91552f1db83ddfe2a1136007c5ea777a76f0c5"></a>

<a id="canonical-0d1f64369054e3dbdb87920c16c0977e0a09020b6e30a3f698c0761bafc33fdb"></a>

## http_method property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 9c11326906cc / 4

Type: `"string"`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

Upstream description:

Specifies the HTTP method used to access a resource.

Any HTTP Method.

Receipt-pinned upstream constraints:

```json
{
  "default": "ANY",
  "enum": [
    "ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [incoming_port](data-sources--workload--reference--group-023.md#canonical-901ff71732c9dd8d7767a708a0a7822db3c3cd9bbc95eebd0fe538a378b075ee): complete subsection reference.

- [path](data-sources--workload--reference--group-023.md#canonical-80a590333d2a668e5bb3962a9ecfb4d2c65b04af470698b3df881dd9666e0759): complete subsection reference.

- [route_direct_response](data-sources--workload--reference--group-023.md#canonical-8bd6f5d922342aac2bdb5a026de55c7c9494cac3c1ceca19566f0284af3fd36c): complete subsection reference.

<a id="canonical-6366e0664c3a2c76bcf8f1f4166ecc50a32ae2d4680ff11361b8679be0a5efdc"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 9c11326906cc / 5

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route.headers](data-sources--workload--reference--group-023.md#canonical-69a56ccce607a77b681f38d7b4dbc732ba3ff1ace362cb1df470298ab287741e)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port](data-sources--workload--reference--group-023.md#canonical-901ff71732c9dd8d7767a708a0a7822db3c3cd9bbc95eebd0fe538a378b075ee)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route.path](data-sources--workload--reference--group-023.md#canonical-80a590333d2a668e5bb3962a9ecfb4d2c65b04af470698b3df881dd9666e0759)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response](data-sources--workload--reference--group-023.md#canonical-8bd6f5d922342aac2bdb5a026de55c7c9494cac3c1ceca19566f0284af3fd36c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-023.md#canonical-c07077f21b2816e8db5b47273a7b26ad2b3fb24187445b00b277c2108076cf6c)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-69a56ccce607a77b681f38d7b4dbc732ba3ff1ace362cb1df470298ab287741e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ccaad538995b33bd259d0d7645cb0199bf9b9e6c88956dedf1b9dd1ab943927c"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route.headers — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 197b6debe618 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-023.md#canonical-20d5fcf752fc72102d70fd1c28f153fc2f63ff15340f7a13c67250ba8c1a0917)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-023.md#canonical-c07077f21b2816e8db5b47273a7b26ad2b3fb24187445b00b277c2108076cf6c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route](data-sources--workload--reference--group-023.md#canonical-9150ada40cd3a65b4f474f3e76d0ed50d03b850009b246f2012864d9ca9f77f1)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route.headers

<a id="canonical-8d65c633c30139acd6428804cc781cd6fbbcc33422b48371245000d408109337"></a>

Type: `"list"`. Computed.

Headers. List of (key, value) headers.

Upstream description:

List of (key, value) headers.

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
    "minItems": 0,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-e340ad52de3c7893fe73e7593506daadc12a586d14620478870ff0750798b63d"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 197b6debe618 / 3

<a id="canonical-54da66c43480cf48e7a098867d81d7153cbeae2822b70e39bf6e3e8194506c1e"></a>

<a id="canonical-a34b815491ee8aa9888971c06937f4eabf0564261980fbf828dc7ac5111378db"></a>

## exact property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 197b6debe618 / 4

Type: `"string"`. Computed.

Exclusive with \[presence regex\] Header value to match exactly.

Upstream description:

Exclusive with \[presence regex\] Header value to match exactly.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-236a6a69c91c0bb00e0711cfdf3929b02b8ea5b0dc68cd5cd01685b7517f44a7"></a>

<a id="canonical-a569195dee942a6f42be16a5b30b696b6554edcb3e919d5277f766cdd722ee37"></a>

## invert_match property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 197b6debe618 / 5

Type: `"bool"`. Computed.

Invert the result of the match to detect missing header or non-matching value.

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

<a id="canonical-17df933ead9e1cbaf71edb7b4514237d9302c9f28cf28a0fc8b02699af90b847"></a>

<a id="canonical-382b07d3265d9739656287b9b43fb945ba5b27fb26a504e84e9b391876d2b55a"></a>

## name property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 197b6debe618 / 6

Type: `"string"`. Computed.

Name. Name of the header.

Upstream description:

Name of the header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-b88be7052a2e8537ad902a9f1b2e6dddbd1862b7e662340a60d13872c854cc02"></a>

<a id="canonical-2ee46baa9224242bc79e2a88264aabe5ac7ac28190a5ef4738eac624bb72684f"></a>

## presence property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 197b6debe618 / 7

Type: `"bool"`. Computed.

Exclusive with \[exact regex\] If true, check for presence of header.

Upstream description:

Exclusive with \[exact regex\] If true, check for presence of header.

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

<a id="canonical-55ed4debcc0c84a7ede7e5137742bb139c40868c977e6b13ca877f59dd1378ab"></a>

<a id="canonical-6b8571d1c863cf0a0ff99029550362ec9d6daf73ba490c46975ad0c3be3262b1"></a>

## regex property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 197b6debe618 / 8

Type: `"string"`. Computed.

Exclusive with \[exact presence\] Regex match of the header value in re2 format.

Upstream description:

Exclusive with \[exact presence\] Regex match of the header value in re2 format.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-63fe02b20716259b5c679f0af4f689fd0660ae2e4a69db0136a5daf46bc8c9c5"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 197b6debe618 / 9

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route](data-sources--workload--reference--group-023.md#canonical-9150ada40cd3a65b4f474f3e76d0ed50d03b850009b246f2012864d9ca9f77f1)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-901ff71732c9dd8d7767a708a0a7822db3c3cd9bbc95eebd0fe538a378b075ee"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c0af47a30ea496c4549b50fe32e9588077f28eadbeae6d6bc92eab2c16aa5009"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 67ec448d3482 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-023.md#canonical-20d5fcf752fc72102d70fd1c28f153fc2f63ff15340f7a13c67250ba8c1a0917)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-023.md#canonical-c07077f21b2816e8db5b47273a7b26ad2b3fb24187445b00b277c2108076cf6c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route](data-sources--workload--reference--group-023.md#canonical-9150ada40cd3a65b4f474f3e76d0ed50d03b850009b246f2012864d9ca9f77f1)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port

<a id="canonical-2d16c799033ab30d7a4b05887137493596200ef1ea464085f33586f966b5563a"></a>

Type: `"single"`. Computed.

Port match of the request can be a range or a specific port.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_match": "[\"no_port_match\",\"port\",\"port_ranges\"]"
}
```

<a id="canonical-a13f46effe341495eb63dede91e07facd4f055e047996a4b7acc88187c52c8db"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 67ec448d3482 / 3

- [no_port_match](data-sources--workload--reference--group-023.md#canonical-a379afab7d478865f5ec6c4ae2fd3eeef0a1ee18cb8744972a14c94b1d21eb10): complete subsection reference.

<a id="canonical-ff056e1963259008d8e8c3fd596a7529993fffb9ef86ef129e522e29b7818f13"></a>

<a id="canonical-b581074c500a8e88defa98a690539cb704d1212a0316080b8240ceccdd0ff312"></a>

## port property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 67ec448d3482 / 4

Type: `"number"`. Computed.

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

Upstream description:

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

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

<a id="canonical-59583c38d16f3e6c61309e669a5a28943cd49062bc3284975a86915f30b0aeee"></a>

<a id="canonical-ace8437ca9e9af8212f33025d6443f5ca46d3939dbd3ae9aabca6c3478942093"></a>

## port_ranges property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 67ec448d3482 / 5

Type: `"string"`. Computed.

Exclusive with \[no\_port\_match port\] Port range to match.

Upstream description:

Exclusive with \[no\_port\_match port\] Port range to match.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 32,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 32,
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
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  }
}
```

<a id="canonical-fcd91de5c05bc51c735c1f2ec74d34579edd3bb3e2f593a63837a8630368cb78"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 67ec448d3482 / 6

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port.no_port_match](data-sources--workload--reference--group-023.md#canonical-a379afab7d478865f5ec6c4ae2fd3eeef0a1ee18cb8744972a14c94b1d21eb10)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route](data-sources--workload--reference--group-023.md#canonical-9150ada40cd3a65b4f474f3e76d0ed50d03b850009b246f2012864d9ca9f77f1)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-a379afab7d478865f5ec6c4ae2fd3eeef0a1ee18cb8744972a14c94b1d21eb10"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-30b67aac3630dc2e154caeb0b2c17b8aad7e6bf404a2abedd2c3ede314575ec2"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port.no_port_match — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / fb0cb76f9f43 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-023.md#canonical-20d5fcf752fc72102d70fd1c28f153fc2f63ff15340f7a13c67250ba8c1a0917)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-023.md#canonical-c07077f21b2816e8db5b47273a7b26ad2b3fb24187445b00b277c2108076cf6c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route](data-sources--workload--reference--group-023.md#canonical-9150ada40cd3a65b4f474f3e76d0ed50d03b850009b246f2012864d9ca9f77f1)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port](data-sources--workload--reference--group-023.md#canonical-901ff71732c9dd8d7767a708a0a7822db3c3cd9bbc95eebd0fe538a378b075ee)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port.no_port_match

<a id="canonical-9513341265d555cf91dee425e51e6fec461a74bbf534745b88b2a10d3a47b725"></a>

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

<a id="canonical-6581efe1c0cc0691d7fd58f94ca97d04994da71017eda81b538cda466684a125"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / fb0cb76f9f43 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-72671eb67a0ffe6494a30491fb8fd99ea5a98d22cf723db00716d61618c26655"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / fb0cb76f9f43 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port](data-sources--workload--reference--group-023.md#canonical-901ff71732c9dd8d7767a708a0a7822db3c3cd9bbc95eebd0fe538a378b075ee)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-80a590333d2a668e5bb3962a9ecfb4d2c65b04af470698b3df881dd9666e0759"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7332853788c993cf822ad7860ee3de0404d42896f13663efd3b33b6551c3534f"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route.path — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 89e574513e35 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-023.md#canonical-20d5fcf752fc72102d70fd1c28f153fc2f63ff15340f7a13c67250ba8c1a0917)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-023.md#canonical-c07077f21b2816e8db5b47273a7b26ad2b3fb24187445b00b277c2108076cf6c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route](data-sources--workload--reference--group-023.md#canonical-9150ada40cd3a65b4f474f3e76d0ed50d03b850009b246f2012864d9ca9f77f1)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route.path

<a id="canonical-eda90c667d4ab826091e89977d325003110b2e2412ee99e615f9de4a8e62edf2"></a>

Type: `"single"`. Computed.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

<a id="canonical-d0ade4a0b09501f4837fbd96066ff7e8fa3bb3cd30dedc47e620ad5b0cfdcfca"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 89e574513e35 / 3

<a id="canonical-78796656795ccb3cfdd2ca8055ac8c04ce66b1c730d716d2fe19f32734756a16"></a>

<a id="canonical-51d13a85f17e86d649357c9e889983f11f8875030030b290f66592c669192797"></a>

## path property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 89e574513e35 / 4

Type: `"string"`. Computed.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

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
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-9489f25b86ca4d82ab0b9d4f1b612da0bb185f74c0178a45fa3a70937dab7f21"></a>

<a id="canonical-67f107e4227a693f9a82adc58d1e2999eafb091484d1851778c26fc802459396"></a>

## prefix property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 89e574513e35 / 5

Type: `"string"`. Computed.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0712256765f01d2f9b0414fbecfd469b7e4158968d7f63eb90e49cedeb4e82f3"></a>

<a id="canonical-0c3eb0fb46f030d830371b5c13e3ecb47488c6c614c7525573f8bb1bfacd2d18"></a>

## regex property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 89e574513e35 / 6

Type: `"string"`. Computed.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-d1351fe76a0ede1e8a86bb93f3272b6cbbdf90406fdd9df11a5e7373148c0758"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 89e574513e35 / 7

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route](data-sources--workload--reference--group-023.md#canonical-9150ada40cd3a65b4f474f3e76d0ed50d03b850009b246f2012864d9ca9f77f1)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-8bd6f5d922342aac2bdb5a026de55c7c9494cac3c1ceca19566f0284af3fd36c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6ef3ebfef28cceb2e7f424cd76c7deb8015a9a70bb77d7f8c4b1ddd395c872c4"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / d40f1b90d2f4 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-023.md#canonical-20d5fcf752fc72102d70fd1c28f153fc2f63ff15340f7a13c67250ba8c1a0917)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-023.md#canonical-c07077f21b2816e8db5b47273a7b26ad2b3fb24187445b00b277c2108076cf6c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route](data-sources--workload--reference--group-023.md#canonical-9150ada40cd3a65b4f474f3e76d0ed50d03b850009b246f2012864d9ca9f77f1)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response

<a id="canonical-bc53cb86e2b9037f718b4f8fc159cfe308ce33d49b658bd7be5921805273659b"></a>

Type: `"single"`. Computed.

Send this direct response in case of route match action is direct response.

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

<a id="canonical-bad72429a343d3d136a75439d0c812ba6f342d8eaed33464affdf94e28f5df3c"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / d40f1b90d2f4 / 3

<a id="canonical-b4e86df656b5ebeebaf7abd2a7f5b07d59c1e18defd96f87b0a96940798bca56"></a>

<a id="canonical-cda1c4ef4c3f4ad9e163b808426acd3d4a10c5285b9189bede66b9648696ae00"></a>

## response_body_encoded property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / d40f1b90d2f4 / 4

Type: `"string"`. Computed.

Response body to send. Currently supported URL schemes is string:/// for which message should be
encoded in Base64 format. The message can be either plain text or HTML.

Upstream description:

Response body to send. Currently supported URL schemes is string:/// for which message should be
encoded in Base64 format. The message can be either plain text or HTML. E.g. "&lt;p&gt; Access
Denied &lt;/p&gt;". Base64 encoded string URL for this is
string:///PHA+IEFjY2VzcyBEZW5pZWQgPC9wPg==.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 65536,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 65536
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 65536,
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
    "ves.io.schema.rules.string.max_bytes": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-26ed305b16c4736130808d6039d4f681403c0c0b1579fb63528ce1571c5d6be2"></a>

<a id="canonical-d36ab83390c57a522dc3da47adde3f1bcf9904f5e0253ff0065a35b3757f0afd"></a>

## response_code property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / d40f1b90d2f4 / 5

Type: `"number"`. Computed.

Response Code. Response code to send.

Upstream description:

Response code to send.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 599,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 100
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "100",
    "ves.io.schema.rules.uint32.lte": "599"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "100",
    "ves.io.schema.rules.uint32.lte": "599"
  }
}
```

<a id="canonical-52455cb25e9e1f911d990d0a4e2128afabf2afa95e314cec8cb5f1a8631f74a4"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / d40f1b90d2f4 / 6

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route](data-sources--workload--reference--group-023.md#canonical-9150ada40cd3a65b4f474f3e76d0ed50d03b850009b246f2012864d9ca9f77f1)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-5e027771e4626c74f12b688baf95b40804037d7cdf062e04cac370092fc01f51"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-78973b90882be99138e3d571e99d32cefb9b08fa6a7855e25091f8225315fd91"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 89db968c2b13 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-023.md#canonical-20d5fcf752fc72102d70fd1c28f153fc2f63ff15340f7a13c67250ba8c1a0917)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-023.md#canonical-c07077f21b2816e8db5b47273a7b26ad2b3fb24187445b00b277c2108076cf6c)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route

<a id="canonical-54548fb2ac645141706eb80b963be04b59b8858a4d55bfc02cffabd16007874f"></a>

Type: `"single"`. Computed.

Redirect route matches on path, incoming header, incoming port and/or HTTP method and redirects the
matching traffic to a different URL.

Upstream description:

A redirect route matches on path, incoming header, incoming port and/or HTTP method and redirects
the matching traffic to a different URL.

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

<a id="canonical-1f405c01d47cc530cfbae0c799ef6d885f702e3f6f83ba524922fb76d323ef21"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 89db968c2b13 / 3

- [headers](data-sources--workload--reference--group-023.md#canonical-692822229cc0f3a4b0b4f3690f2e55800b80cd7eb64fe6e9da4d22c205223988): complete subsection reference.

<a id="canonical-9752b15e8d941c32592d9ea65ec035c7b1f0bf516d9963c0f1875f200385fc69"></a>

<a id="canonical-d98e9fb05c86ef7cbafd44a3575db94d0fbf4fa98667a2ebe782c8cc74c7e215"></a>

## http_method property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 89db968c2b13 / 4

Type: `"string"`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

Upstream description:

Specifies the HTTP method used to access a resource.

Any HTTP Method.

Receipt-pinned upstream constraints:

```json
{
  "default": "ANY",
  "enum": [
    "ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [incoming_port](data-sources--workload--reference--group-023.md#canonical-d5d2a60576f7cae6df63feb6af6e4a080e1c5bab7f3881965e40fbfcaa8b9846): complete subsection reference.

- [path](data-sources--workload--reference--group-023.md#canonical-8c26b960282ba754c0053b69529d59e33140c2370ae91250cd34699f0b4c6f98): complete subsection reference.

- [route_redirect](data-sources--workload--reference--group-023.md#canonical-bb94e98fb2e8e2337fc10061a8168cd0ff5d21f8c79af03eda8815afafdf366e): complete subsection reference.

<a id="canonical-c4d0301f8e0f9643fcd8e543c7682fe405e524e3fb8b761d7e7f19f545cab707"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 89db968c2b13 / 5

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.headers](data-sources--workload--reference--group-023.md#canonical-692822229cc0f3a4b0b4f3690f2e55800b80cd7eb64fe6e9da4d22c205223988)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port](data-sources--workload--reference--group-023.md#canonical-d5d2a60576f7cae6df63feb6af6e4a080e1c5bab7f3881965e40fbfcaa8b9846)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.path](data-sources--workload--reference--group-023.md#canonical-8c26b960282ba754c0053b69529d59e33140c2370ae91250cd34699f0b4c6f98)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](data-sources--workload--reference--group-023.md#canonical-bb94e98fb2e8e2337fc10061a8168cd0ff5d21f8c79af03eda8815afafdf366e)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-023.md#canonical-c07077f21b2816e8db5b47273a7b26ad2b3fb24187445b00b277c2108076cf6c)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-692822229cc0f3a4b0b4f3690f2e55800b80cd7eb64fe6e9da4d22c205223988"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cefd9ac389ab824b3430044ecc095ecd3b8def6d665ff09258d3c07bb76697e9"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.headers — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 3ffb44278b64 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-023.md#canonical-20d5fcf752fc72102d70fd1c28f153fc2f63ff15340f7a13c67250ba8c1a0917)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-023.md#canonical-c07077f21b2816e8db5b47273a7b26ad2b3fb24187445b00b277c2108076cf6c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-023.md#canonical-5e027771e4626c74f12b688baf95b40804037d7cdf062e04cac370092fc01f51)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.headers

<a id="canonical-f3714f946eb93ed39fe66d066fd9eba70a2bc61c3232de90dec0db385d11ff56"></a>

Type: `"list"`. Computed.

Headers. List of (key, value) headers.

Upstream description:

List of (key, value) headers.

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
    "minItems": 0,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-8fa7c221b231235676c00ffac977bebed0d9133c2d7ed4355dfea1f5a05f3657"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 3ffb44278b64 / 3

<a id="canonical-ae5025b079906b47d8eec39f9db9392354d773d89027d846fe0ad04f21fff7e9"></a>

<a id="canonical-182c654b4355e41fc8429ccae2402393cc4e09f14c4b10dc1cbbc25aa93a0487"></a>

## exact property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 3ffb44278b64 / 4

Type: `"string"`. Computed.

Exclusive with \[presence regex\] Header value to match exactly.

Upstream description:

Exclusive with \[presence regex\] Header value to match exactly.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-e51b5bc845e7387646fb77556add99838031e93ddc3f36e0ce317a7d2a70a661"></a>

<a id="canonical-7f186d5b44b8e63f3a35f849203e23c94249576011f6fe9c24de87061e925929"></a>

## invert_match property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 3ffb44278b64 / 5

Type: `"bool"`. Computed.

Invert the result of the match to detect missing header or non-matching value.

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

<a id="canonical-9c8dd545d12f90a0e4b71bc946263e4ca6f5965aecdfcaadda0b715dfb01ce43"></a>

<a id="canonical-d9c05ce303b951d6cb82b5bb362a1834cdf58cef3e87f09945fb855f479e7346"></a>

## name property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 3ffb44278b64 / 6

Type: `"string"`. Computed.

Name. Name of the header.

Upstream description:

Name of the header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-663214dfcad9e906573b22c3b119f03130f6ce2bb77b13703ccd3575e8d5bb17"></a>

<a id="canonical-eedd2c2e9f28cb7d398df3d76127811bea0193843ac299e1138c0d3f99efbedc"></a>

## presence property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 3ffb44278b64 / 7

Type: `"bool"`. Computed.

Exclusive with \[exact regex\] If true, check for presence of header.

Upstream description:

Exclusive with \[exact regex\] If true, check for presence of header.

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

<a id="canonical-3636dc3949bfe3602ea62eb6e25c63798274ef3cdb46eb74002c84e0bda44502"></a>

<a id="canonical-fe0a62674387861802635a79c9b3ff254348d03bd2e639a9c8668729bda38950"></a>

## regex property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 3ffb44278b64 / 8

Type: `"string"`. Computed.

Exclusive with \[exact presence\] Regex match of the header value in re2 format.

Upstream description:

Exclusive with \[exact presence\] Regex match of the header value in re2 format.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-4125253bcc5e983f37d254a11ac545498485d0f8f0b78843471c21767d0c551d"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 3ffb44278b64 / 9

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-023.md#canonical-5e027771e4626c74f12b688baf95b40804037d7cdf062e04cac370092fc01f51)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-d5d2a60576f7cae6df63feb6af6e4a080e1c5bab7f3881965e40fbfcaa8b9846"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6f911b4f960a8c9875983854ce5c2d0760a6d187aec5026abc9ba6df782b705b"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / b6d14680fe42 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-023.md#canonical-20d5fcf752fc72102d70fd1c28f153fc2f63ff15340f7a13c67250ba8c1a0917)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-023.md#canonical-c07077f21b2816e8db5b47273a7b26ad2b3fb24187445b00b277c2108076cf6c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-023.md#canonical-5e027771e4626c74f12b688baf95b40804037d7cdf062e04cac370092fc01f51)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port

<a id="canonical-85e4b3221594824c0dc51564bafc85b1eee59574f1156e81b95336d5938cf9cd"></a>

Type: `"single"`. Computed.

Port match of the request can be a range or a specific port.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_match": "[\"no_port_match\",\"port\",\"port_ranges\"]"
}
```

<a id="canonical-39f02390866f8f5028a993a609c6c0089aff9c492c964eb49ee68e04e929668d"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / b6d14680fe42 / 3

- [no_port_match](data-sources--workload--reference--group-023.md#canonical-91bcc6518c3a5a027bc588d007330e07f9309cf77ccbf01e8900163c49e12355): complete subsection reference.

<a id="canonical-21ee2a23ce6d6f64fb34d5694772c54687db85e3a7e7994a732797411061743e"></a>

<a id="canonical-d756916e9c5a57684759dba70596174e0b4969a0946d01e34a77a0fdeabc1911"></a>

## port property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / b6d14680fe42 / 4

Type: `"number"`. Computed.

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

Upstream description:

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

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

<a id="canonical-9b0a317d176d1e0d6cf8fddbe078108bd4ce6d0106a40085d78648bab3119349"></a>

<a id="canonical-079e6652389b3eb2bf0a983cc38a9a1a7225ba5eed6f0c6431759df15b9652e3"></a>

## port_ranges property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / b6d14680fe42 / 5

Type: `"string"`. Computed.

Exclusive with \[no\_port\_match port\] Port range to match.

Upstream description:

Exclusive with \[no\_port\_match port\] Port range to match.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 32,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 32,
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
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  }
}
```

<a id="canonical-0b322dd7a7478411744645e0da7fac44a70752e457d6bfeddecb716266c5eb5f"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / b6d14680fe42 / 6

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port.no_port_match](data-sources--workload--reference--group-023.md#canonical-91bcc6518c3a5a027bc588d007330e07f9309cf77ccbf01e8900163c49e12355)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-023.md#canonical-5e027771e4626c74f12b688baf95b40804037d7cdf062e04cac370092fc01f51)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-91bcc6518c3a5a027bc588d007330e07f9309cf77ccbf01e8900163c49e12355"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b28cc1d317a94af986a8efed62ac07bc3119677cfe9bcd175812770a96f645cb"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port.no_port_match — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / f736c36ccd9c / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-023.md#canonical-20d5fcf752fc72102d70fd1c28f153fc2f63ff15340f7a13c67250ba8c1a0917)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-023.md#canonical-c07077f21b2816e8db5b47273a7b26ad2b3fb24187445b00b277c2108076cf6c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-023.md#canonical-5e027771e4626c74f12b688baf95b40804037d7cdf062e04cac370092fc01f51)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port](data-sources--workload--reference--group-023.md#canonical-d5d2a60576f7cae6df63feb6af6e4a080e1c5bab7f3881965e40fbfcaa8b9846)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port.no_port_match

<a id="canonical-307d43587a35568e2524a9e89d5f7b3e46ff645b60770982931c44dfe77f8b40"></a>

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

<a id="canonical-fcb56956b0f247d20ee4ccdd51c760a98b6b8a88ef68635c440d8bc7b798f879"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / f736c36ccd9c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1ee81d1d3e4dbca71b5cb6c2e0aba6b961816455fd9f685b42059726f11a61c1"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / f736c36ccd9c / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port](data-sources--workload--reference--group-023.md#canonical-d5d2a60576f7cae6df63feb6af6e4a080e1c5bab7f3881965e40fbfcaa8b9846)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-8c26b960282ba754c0053b69529d59e33140c2370ae91250cd34699f0b4c6f98"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-16129aefd17215521d2b2abd85cbf92e132e576fbf5b79453ed3a31485332916"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.path — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 387cceaf2fa3 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-023.md#canonical-20d5fcf752fc72102d70fd1c28f153fc2f63ff15340f7a13c67250ba8c1a0917)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-023.md#canonical-c07077f21b2816e8db5b47273a7b26ad2b3fb24187445b00b277c2108076cf6c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-023.md#canonical-5e027771e4626c74f12b688baf95b40804037d7cdf062e04cac370092fc01f51)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.path

<a id="canonical-783e830c4de97d48633227f141c78dc9861d130537b55adbd4fbc39ba5c25c1e"></a>

Type: `"single"`. Computed.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

<a id="canonical-c08d460098fe4777f0fb0c8b6f6c4b04c63b0a25b437c01e6de83f60c84033f7"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 387cceaf2fa3 / 3

<a id="canonical-999d31d978266004a8c81ce4edf80cf50ca7795bc5face2fc3e19e3c20a47b33"></a>

<a id="canonical-0402dd91927199676dcd4fdf1aea72cb94e0ebe834d564db2a5edb7e39d034ef"></a>

## path property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 387cceaf2fa3 / 4

Type: `"string"`. Computed.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

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
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-f542cb92b956cca4d207a1b69f70f86f49d31bd4956767d6e2a9e6aa0667d714"></a>

<a id="canonical-8228061dcad6275b77340fc71e460c18942e647fdd2fdb98d0244144194b41bd"></a>

## prefix property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 387cceaf2fa3 / 5

Type: `"string"`. Computed.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-c7e67f8927a2522812afd886db3caf09bea9f3cc824f47a80f62fbdab9387902"></a>

<a id="canonical-8152f3b1ad9a9945bac99905df99d2a6dda830211f987cabc05f77a7c4e4d423"></a>

## regex property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 387cceaf2fa3 / 6

Type: `"string"`. Computed.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-eb4715aaaaa10245f9ebaace14257f68ec8ea5cf104ceddd8c72719b39973e5e"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 387cceaf2fa3 / 7

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-023.md#canonical-5e027771e4626c74f12b688baf95b40804037d7cdf062e04cac370092fc01f51)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-bb94e98fb2e8e2337fc10061a8168cd0ff5d21f8c79af03eda8815afafdf366e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5b6a00df63166d6c2d49b65174124f01f5b932ad34c79ac00728859e65099b42"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 4789d67b21a2 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-023.md#canonical-20d5fcf752fc72102d70fd1c28f153fc2f63ff15340f7a13c67250ba8c1a0917)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-023.md#canonical-c07077f21b2816e8db5b47273a7b26ad2b3fb24187445b00b277c2108076cf6c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-023.md#canonical-5e027771e4626c74f12b688baf95b40804037d7cdf062e04cac370092fc01f51)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect

<a id="canonical-0956dfcea6ac2efd8342fb3287840c1f33857c119866bc065c4acef84047bd27"></a>

Type: `"single"`. Computed.

Route redirect parameters when match action is redirect.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-query_params": "[\"remove_all_params\",\"replace_params\",\"retain_all_params\"]",
  "x-ves-oneof-field-redirect_path_choice": "[\"path_redirect\",\"prefix_rewrite\"]"
}
```

<a id="canonical-5239eeff31025c2f5bc8f43896cb457473815d28066664acb3334a9a43bcb100"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 4789d67b21a2 / 3

<a id="canonical-a04b23491be78c654105fc6b1464eb90e70bb26a9b0e05c51fc40b29ce411f1f"></a>

<a id="canonical-80cdd4879e0a15966c6ffafc34989409c00f42a8cb5bd553eac1995a3298e8f2"></a>

## host_redirect property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 4789d67b21a2 / 4

Type: `"string"`. Computed.

Swap host part of incoming URL in redirect URL.

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

<a id="canonical-4049f070053f2a9d6a812039b8bf19173fbc84e92ac7b386099d3d97783941f1"></a>

<a id="canonical-f9613cbb86663e85c801348c469e0066986d6b019b7e341a6c21be1988f79454"></a>

## path_redirect property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 4789d67b21a2 / 5

Type: `"string"`. Computed.

Exclusive with \[prefix\_rewrite\] swap path part of incoming URL in redirect URL.

Upstream description:

Exclusive with \[prefix\_rewrite\] swap path part of incoming URL in redirect URL.

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-2e60dbf71ea7d1073d94344512387bba81b86c0429f24fefb523abc379c7b9e8"></a>

<a id="canonical-6185cf69cf8dcb98f6ffa3cdc654dff402c5ea436fb07d6df5d68699fd8be0dd"></a>

## prefix_rewrite property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 4789d67b21a2 / 6

Type: `"string"`. Computed.

Exclusive with \[path\_redirect\] In Redirect response, the matched prefix (or path) should be
swapped with this value. This option allows redirect URLs be dynamically created based on the
request.

Upstream description:

Exclusive with \[path\_redirect\] In Redirect response, the matched prefix (or path) should be
swapped with this value. This option allows redirect URLs be dynamically created based on the
request.

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-399fc9e1495131483fe8605623bac289b5d694b45fc17dc34cc83981c5d1e5f7"></a>

<a id="canonical-4b2cfb323f5847522f45e9007ce5c76ef75a37fb42e345d40b5c54a8fa2c78ca"></a>

## proto_redirect property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 4789d67b21a2 / 7

Type: `"string"`. Computed.

\[Enum: incoming-proto|http|https\] Swap protocol part of incoming URL in redirect URL The protocol
can be swapped with either HTTP or HTTPS When incoming-proto option is specified, swapping of
protocol is not done. Possible values are \`incoming-proto\`, \`http\`, \`https\`.

Upstream description:

Swap protocol part of incoming URL in redirect URL The protocol can be swapped with either HTTP or
HTTPS When incoming-proto option is specified, swapping of protocol is not done.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "incoming-proto",
    "http",
    "https"
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
    "ves.io.schema.rules.string.in": "[\\\"incoming-proto\\\",\\\"http\\\",\\\"https\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"incoming-proto\\\",\\\"http\\\",\\\"https\\\"]"
  }
}
```

- [remove_all_params](data-sources--workload--reference--group-023.md#canonical-383614c1fc189fefdd85b9eacf6e7bdd5eb94c0b17523f0dac7a4ed7b906afb5): complete subsection reference.

<a id="canonical-9810cb1a19dcc07fb3f33b1a5fe42694f2780174d65cf7c17e8cc0c0c4b86f1d"></a>

<a id="canonical-19779c4058b6cc966ab87501158ff25929075162ec53a37198a6ea1ac35e0a77"></a>

## replace_params property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 4789d67b21a2 / 8

Type: `"string"`. Computed.

Exclusive with \[remove\_all\_params retain\_all\_params\].

Upstream description:

Exclusive with \[remove\_all\_params retain\_all\_params\]

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

<a id="canonical-d93f6feddc9a835e4262e9e3096a912968d480ed83b565107b313ecc17e4bc01"></a>

<a id="canonical-b63930c446d80443ac4427659f7fcbaf6fbec52edcb39028c74ba9470336827c"></a>

## response_code property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 4789d67b21a2 / 9

Type: `"number"`. Computed.

The HTTP status code to use in the redirect response.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 599,
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
    "ves.io.schema.rules.uint32.lte": "599"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "599"
  }
}
```

- [retain_all_params](data-sources--workload--reference--group-023.md#canonical-2d20e8e4704d8e15b7c27eaf586fee03a09ae9742a34831e44191960813b21d5): complete subsection reference.

<a id="canonical-77c92e850ce14d6444dc4ce529c254a260872cd64f1d5a9ff23275851409ee17"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 4789d67b21a2 / 10

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.remove_all_params](data-sources--workload--reference--group-023.md#canonical-383614c1fc189fefdd85b9eacf6e7bdd5eb94c0b17523f0dac7a4ed7b906afb5)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.retain_all_params](data-sources--workload--reference--group-023.md#canonical-2d20e8e4704d8e15b7c27eaf586fee03a09ae9742a34831e44191960813b21d5)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-023.md#canonical-5e027771e4626c74f12b688baf95b40804037d7cdf062e04cac370092fc01f51)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-383614c1fc189fefdd85b9eacf6e7bdd5eb94c0b17523f0dac7a4ed7b906afb5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6145a33722b05803b6357476e9711aa8265fb65068aa2609dbe40ef7be25f02a"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.remove_all_params — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 5b4e96fb8f18 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-023.md#canonical-20d5fcf752fc72102d70fd1c28f153fc2f63ff15340f7a13c67250ba8c1a0917)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-023.md#canonical-c07077f21b2816e8db5b47273a7b26ad2b3fb24187445b00b277c2108076cf6c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-023.md#canonical-5e027771e4626c74f12b688baf95b40804037d7cdf062e04cac370092fc01f51)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](data-sources--workload--reference--group-023.md#canonical-bb94e98fb2e8e2337fc10061a8168cd0ff5d21f8c79af03eda8815afafdf366e)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.remove_all_params

<a id="canonical-a6e2e9e1f4b384aae2a4035f17deb05c1149eed78e2d5e576d731acf6fff549d"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for remove all params.

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

<a id="canonical-18f2a56d317b11dffe558c8f0938f5278ed9106dd6ac50170ffd58d8e08c86b3"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 5b4e96fb8f18 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e9a1ede54fff0fe1efdf013f401a111915023d6702d8a8afd142c8f18afb56f1"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 5b4e96fb8f18 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](data-sources--workload--reference--group-023.md#canonical-bb94e98fb2e8e2337fc10061a8168cd0ff5d21f8c79af03eda8815afafdf366e)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-2d20e8e4704d8e15b7c27eaf586fee03a09ae9742a34831e44191960813b21d5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-88b5bc900c2bb3426d8a8ab044eecba8856bd8959f37153843ded22318a950ef"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.retain_all_params — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / d47b002e36aa / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-023.md#canonical-20d5fcf752fc72102d70fd1c28f153fc2f63ff15340f7a13c67250ba8c1a0917)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-023.md#canonical-c07077f21b2816e8db5b47273a7b26ad2b3fb24187445b00b277c2108076cf6c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-023.md#canonical-5e027771e4626c74f12b688baf95b40804037d7cdf062e04cac370092fc01f51)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](data-sources--workload--reference--group-023.md#canonical-bb94e98fb2e8e2337fc10061a8168cd0ff5d21f8c79af03eda8815afafdf366e)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.retain_all_params

<a id="canonical-3af7b89a3ab1908f9850c7ddea3f474219e20687c8a07a722eedb0bc001d4f97"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for retain all params.

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

<a id="canonical-f37d1701b18c22135a99c35b5ba5a0d86e472ccc786c1c51eb13543aa2037bae"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / d47b002e36aa / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-62ab8fcfefa513648793b3feb5fdb819123b592fa2390c41799a4a5f3fddcbdb"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / d47b002e36aa / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](data-sources--workload--reference--group-023.md#canonical-bb94e98fb2e8e2337fc10061a8168cd0ff5d21f8c79af03eda8815afafdf366e)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-aa1a4e20c821310874f143b2cde560cdb2dfe7962517bb612441fa2e0ccdae55"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f6a5db45cc74dda620689cddc2814f047e4d410a816589451d5413df69649ad1"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 2cb4fc0de0ea / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-023.md#canonical-20d5fcf752fc72102d70fd1c28f153fc2f63ff15340f7a13c67250ba8c1a0917)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-023.md#canonical-c07077f21b2816e8db5b47273a7b26ad2b3fb24187445b00b277c2108076cf6c)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route

<a id="canonical-22eebd303fa5342116438d2ebd446b9bffda3a3e2c18f0b99346c275a4c20d56"></a>

Type: `"single"`. Computed.

Simple route matches on path and/or HTTP method and forwards the matching traffic to the default
origin pool specified outside.

Upstream description:

A simple route matches on path and/or HTTP method and forwards the matching traffic to the default
origin pool specified outside.

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

<a id="canonical-5be55c54249ca137f37363b0b0136dd5cac58bb928edae043b38bbe07b564ac2"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 2cb4fc0de0ea / 3

- [auto_host_rewrite](data-sources--workload--reference--group-023.md#canonical-3f3929df38d0ce9741883baf0e9282b358c8a22107ccead654dbcef102ec8aad): complete subsection reference.

- [disable_host_rewrite](data-sources--workload--reference--group-023.md#canonical-adf8c3218940e78b1c407c92f87ed05bb749ff1aadab5738d88fde6722b215ff): complete subsection reference.

<a id="canonical-15a5fee375bd15328166f603ad0307dd0ac623836253edb11b8238a9ec928bd7"></a>

<a id="canonical-b25735162cf8e854fa64cd0342490d3091e286e2e3d2a60ed859d0e3e229e5a6"></a>

## host_rewrite property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 2cb4fc0de0ea / 4

Type: `"string"`. Computed.

Exclusive with \[auto\_host\_rewrite disable\_host\_rewrite\] Host header will be swapped with this
value.

Upstream description:

Exclusive with \[auto\_host\_rewrite disable\_host\_rewrite\] Host header will be swapped with this
value.

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

<a id="canonical-da8f90952694a15bfdf1aa9f49c21b4b40f4992f0efb307e548df9fbcb94fa20"></a>

<a id="canonical-860ac6e6e418288ff3836ee8e68018edf353cd89ec2d21f9db2b3c151f030388"></a>

## http_method property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 2cb4fc0de0ea / 5

Type: `"string"`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

Upstream description:

Specifies the HTTP method used to access a resource.

Any HTTP Method.

Receipt-pinned upstream constraints:

```json
{
  "default": "ANY",
  "enum": [
    "ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [path](data-sources--workload--reference--group-023.md#canonical-416f8eac7281e26f6b22b2c015bf37247a0c806b7674d273c9f1d34f6d66215d): complete subsection reference.

<a id="canonical-00b182a9788b62c63f96bd517f8e0e981bfa097c8c430c94c4d5a69fb23f6c41"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 2cb4fc0de0ea / 6

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route.auto_host_rewrite](data-sources--workload--reference--group-023.md#canonical-3f3929df38d0ce9741883baf0e9282b358c8a22107ccead654dbcef102ec8aad)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite](data-sources--workload--reference--group-023.md#canonical-adf8c3218940e78b1c407c92f87ed05bb749ff1aadab5738d88fde6722b215ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route.path](data-sources--workload--reference--group-023.md#canonical-416f8eac7281e26f6b22b2c015bf37247a0c806b7674d273c9f1d34f6d66215d)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-023.md#canonical-c07077f21b2816e8db5b47273a7b26ad2b3fb24187445b00b277c2108076cf6c)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-3f3929df38d0ce9741883baf0e9282b358c8a22107ccead654dbcef102ec8aad"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-76e75fad9d1b68aca25f27b6e63625bbceaad0ec222b409ab777000b1bc0f3b6"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route.auto_host_rewrite — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 2614040f4e45 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-023.md#canonical-20d5fcf752fc72102d70fd1c28f153fc2f63ff15340f7a13c67250ba8c1a0917)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-023.md#canonical-c07077f21b2816e8db5b47273a7b26ad2b3fb24187445b00b277c2108076cf6c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route](data-sources--workload--reference--group-023.md#canonical-aa1a4e20c821310874f143b2cde560cdb2dfe7962517bb612441fa2e0ccdae55)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route.auto_host_rewrite

<a id="canonical-ce8ce49146a40a3ea603a458b6b38b9da2dd580ef422f43ef1662edb2c246c16"></a>

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

<a id="canonical-654fd7523209bed1ecb8937db358f3d9a3d1e37f1d21800a07bb5fa113dcab2e"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 2614040f4e45 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cdc4aaecaa62eaf2df315d19438b5a9afe6ab171fb6889c8cfab5151eea3a574"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 2614040f4e45 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route](data-sources--workload--reference--group-023.md#canonical-aa1a4e20c821310874f143b2cde560cdb2dfe7962517bb612441fa2e0ccdae55)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-adf8c3218940e78b1c407c92f87ed05bb749ff1aadab5738d88fde6722b215ff"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ab2f982a8582c5995efb2c25e61b59d7f8d82548909d86c7d218bab68bc503ab"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / b2a38d6fe625 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-023.md#canonical-20d5fcf752fc72102d70fd1c28f153fc2f63ff15340f7a13c67250ba8c1a0917)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-023.md#canonical-c07077f21b2816e8db5b47273a7b26ad2b3fb24187445b00b277c2108076cf6c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route](data-sources--workload--reference--group-023.md#canonical-aa1a4e20c821310874f143b2cde560cdb2dfe7962517bb612441fa2e0ccdae55)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite

<a id="canonical-de692886d9bded360332fc71014e3724a9a763eb87700ad6b5ab2c3acf983c4b"></a>

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

<a id="canonical-2ea0eec972dcdef832c061654db0628b524c1ccc42d576c94eafaf4d696bc23c"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / b2a38d6fe625 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-705ced8bd6b72e8b82eecdfa16f75ed06fb2c260e079cd87f80426fd43f055fb"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / b2a38d6fe625 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route](data-sources--workload--reference--group-023.md#canonical-aa1a4e20c821310874f143b2cde560cdb2dfe7962517bb612441fa2e0ccdae55)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-416f8eac7281e26f6b22b2c015bf37247a0c806b7674d273c9f1d34f6d66215d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1dd857195c97a96df99afb4d6c5e46307d28e6772c17d811949e035e0e21ae0d"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route.path — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 80aa6e76a9a1 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-023.md#canonical-20d5fcf752fc72102d70fd1c28f153fc2f63ff15340f7a13c67250ba8c1a0917)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-023.md#canonical-c07077f21b2816e8db5b47273a7b26ad2b3fb24187445b00b277c2108076cf6c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route](data-sources--workload--reference--group-023.md#canonical-aa1a4e20c821310874f143b2cde560cdb2dfe7962517bb612441fa2e0ccdae55)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route.path

<a id="canonical-aa1114650356cfeeaa138add08e879566ea2cda5a451204cf85da7ba4efefb0f"></a>

Type: `"single"`. Computed.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

<a id="canonical-3bf08f950fb4479dbe98890c832aeb715179b09d7f8774563d8a0bdc13a35112"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 80aa6e76a9a1 / 3

<a id="canonical-eb5f7bb105668a2e901425ee766a45a04f4c2b3ee6c48d6f07e67862b7570c86"></a>

<a id="canonical-9b36c92d632af443ec7d724e1bafccd359a10ce63d7abab22fe409d4d4dc619d"></a>

## path property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 80aa6e76a9a1 / 4

Type: `"string"`. Computed.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

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
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-6b56e0279c41ea66f8f2b3d2616dfd2bc2f23782ad7e46c7db1f48e469d5e6c3"></a>

<a id="canonical-a6a994bd78e165bd57443f3ed5707eae34603afa366ee3d4d582e06154b3d977"></a>

## prefix property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 80aa6e76a9a1 / 5

Type: `"string"`. Computed.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-c14dc19939a9c4cfbe809652d1f4d72b790b415658e7600687004632f0f305ae"></a>

<a id="canonical-2ebfc8de50b2f2a8f694455391534673642f0f27f8c658ec4b4188961cae99d7"></a>

## regex property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 80aa6e76a9a1 / 6

Type: `"string"`. Computed.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-7b2a1869d9b9609c36e788c42f1a8cdbc3d0c916c1865cedcf34fe839322cb07"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 80aa6e76a9a1 / 7

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route](data-sources--workload--reference--group-023.md#canonical-aa1a4e20c821310874f143b2cde560cdb2dfe7962517bb612441fa2e0ccdae55)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-8e6beb1e6ec643dfb4a332587fc2573c1c206ef42d39c64e3c513be4b816197c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-06a17714c9867aa986c9cf7e084c5e91f75a822c86ce1c0d7d1677aaae18d944"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port / 0e94e836ac06 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port

<a id="canonical-4de58de65a6501b37dca161480f1781d7c557009e5ec4fbff25b4c62cfdbb12a"></a>

Type: `"single"`. Computed.

Port. Port of the workload.

Upstream description:

Port of the workload.

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

<a id="canonical-4aeabc9609f7e50035f9913188f4fb24ae24c5ad64d00431e3955e4dd3611e4f"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port / 0e94e836ac06 / 3

- [info](data-sources--workload--reference--group-023.md#canonical-9c095dfe4c3c20909f92aee108a806384006e775db738e36461723a66638c9a2): complete subsection reference.

<a id="canonical-499091dbdf021229c887f554092d94fd5d8b8a91dd13972628cbba328b664391"></a>

<a id="canonical-51241d7bed6227d0297814d58d2dbc4a4d1a681aa9a4fdb4e4b9f2f11b2c112d"></a>

## name property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port / 0e94e836ac06 / 4

Type: `"string"`. Computed.

Name. Name of the Port.

Upstream description:

Name of the Port.

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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.iana_svc_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.iana_svc_name": "true"
  }
}
```

<a id="canonical-b91567e353250d8703bf7da2adc56fed0f9977d2de53a7b0e3f094950bc60366"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port / 0e94e836ac06 / 5

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port.info](data-sources--workload--reference--group-023.md#canonical-9c095dfe4c3c20909f92aee108a806384006e775db738e36461723a66638c9a2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-9c095dfe4c3c20909f92aee108a806384006e775db738e36461723a66638c9a2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2a61366715416846615d8d42b5d32bb298d55be4f4018976f4115685e54f3892"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port.info — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port.in / 49eba95535b2 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port](data-sources--workload--reference--group-023.md#canonical-8e6beb1e6ec643dfb4a332587fc2573c1c206ef42d39c64e3c513be4b816197c)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port.info

<a id="canonical-ebb2ba84791551f057832b970acddfb4ed3fe52119671b209cd55a5da070fd83"></a>

Type: `"single"`. Computed.

Port Information. Port information.

Upstream description:

Port information.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-target_port_choice": "[\"same_as_port\",\"target_port\"]"
}
```

<a id="canonical-b51c506a2bc31093a48cd948ee873d064a748c84f54306e313f4cd91840c554a"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port.in / 49eba95535b2 / 3

<a id="canonical-3141093b66457d9ad04ee138e556a8e81e954f871260940fcdac3a49345cdc64"></a>

<a id="canonical-6b93b872444497d0b9dd9c0207e90e8768e93d4cc2099aba90019f072c724eee"></a>

## port property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port.in / 49eba95535b2 / 4

Type: `"number"`. Computed.

Port. Port the workload can be reached on.

Upstream description:

Port the workload can be reached on.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-3c0c8de859ce1149f59677fc0f2ac71c44bdfce1704de34f779240ed215812c1"></a>

<a id="canonical-9a95c9af3481ab1faa93022fd0a473110df516dd939f2a8aef6a4b99162c9af2"></a>

## protocol property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port.in / 49eba95535b2 / 5

Type: `"string"`. Computed.

\[Enum: PROTOCOL\_TCP|PROTOCOL\_HTTP|PROTOCOL\_HTTP2|PROTOCOL\_TLS\_WITH\_SNI|PROTOCOL\_UDP\] Type
of protocol - PROTOCOL\_TCP: TCP TCP - PROTOCOL\_HTTP: HTTP HTTP - PROTOCOL\_HTTP2: HTTP2 HTTP2 -
PROTOCOL\_TLS\_WITH\_SNI: TLS with SNI TLS with SNI - PROTOCOL\_UDP: UDP UDP. Possible values are
\`PROTOCOL\_TCP\`, \`PROTOCOL\_HTTP\`, \`PROTOCOL\_HTTP2\`, \`PROTOCOL\_TLS\_WITH\_SNI\`,
\`PROTOCOL\_UDP\`. Defaults to \`PROTOCOL\_TCP\`.

Upstream description:

Type of protocol

&#8203;- PROTOCOL\_TCP: TCP

TCP &#8203;- PROTOCOL\_HTTP: HTTP

HTTP &#8203;- PROTOCOL\_HTTP2: HTTP2

HTTP2 &#8203;- PROTOCOL\_TLS\_WITH\_SNI: TLS with SNI

TLS with SNI &#8203;- PROTOCOL\_UDP: UDP

UDP.

Receipt-pinned upstream constraints:

```json
{
  "default": "PROTOCOL_TCP",
  "enum": [
    "PROTOCOL_TCP",
    "PROTOCOL_HTTP",
    "PROTOCOL_HTTP2",
    "PROTOCOL_TLS_WITH_SNI",
    "PROTOCOL_UDP"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [same_as_port](data-sources--workload--reference--group-023.md#canonical-ce2c6d0493cfca4fed613947fc4443828b45b1384db3043ab3c610aedcdd2ff2): complete subsection reference.

<a id="canonical-53ec9c910bc3c31f9c323d7ad6113bd83d3f0c007655bbfb23b317f473584ce5"></a>

<a id="canonical-1fb3998c0a8e425fbf00c4b516d76e1b9e5b98b3cd70c22cc16c6df1f783bab1"></a>

## target_port property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port.in / 49eba95535b2 / 6

Type: `"number"`. Computed.

Exclusive with \[same\_as\_port\] Port the workload is listening on.

Upstream description:

Exclusive with \[same\_as\_port\] Port the workload is listening on.

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
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-175421eddf11110028b3c457a2752ddaa340a67a8a1d0cdbd7d1c09e6ae48f88"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port.in / 49eba95535b2 / 7

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port.info.same_as_port](data-sources--workload--reference--group-023.md#canonical-ce2c6d0493cfca4fed613947fc4443828b45b1384db3043ab3c610aedcdd2ff2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port](data-sources--workload--reference--group-023.md#canonical-8e6beb1e6ec643dfb4a332587fc2573c1c206ef42d39c64e3c513be4b816197c)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-ce2c6d0493cfca4fed613947fc4443828b45b1384db3043ab3c610aedcdd2ff2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
