---
page_title: "xcsh_certificate reference"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_certificate reference."
---

# xcsh_certificate reference

<a id="canonical-6e191e632e125ee6cf3e18d3e730166767b3b520ab1f5c6d147e575ffbede87c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4eff0ca1e8fb75a0c34bf2d70c88a0ed92051fff7e6618a1cfc0691e75e6a7e1"></a>

## Property reference — Property reference / df91404af0d7 / 2

Breadcrumbs:

- [xcsh_certificate](../resources/certificate.md#canonical-843c7ef2d56e85f1ba6b346483224fce2e27fa0dc0213cd0e259331947e5281c)
- Property reference

<a id="canonical-78683b2d8843d6c64308ae37e91ca2f5e312464bb05e85159578abb19d90d576"></a>

## Direct properties — Property reference / df91404af0d7 / 3

<a id="canonical-38a354c2500381250bdfa3542585b468121a38608e96f4cc7c7ba1a6fa744422"></a>

<a id="canonical-e4606bb699f23aa3d6a443039272b80ceb857cdf968c2d06ffdd7ca35011ef25"></a>

## annotations property — Property reference / df91404af0d7 / 4

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

- [certificate_chain](resources--certificate--reference--group-001.md#canonical-4ce00beeaeca753bf40bf153c80e3c73fe88090767a9268c8211471b9bc758f2): complete subsection reference.

<a id="canonical-18b9c9d3018dfa7d907a17dcfff0a875f0bc4fd44d8c9c91bd286340f3b4e497"></a>

<a id="canonical-335bf068f18caebad332a49c70ab1b55a42bbcc7294cb569cec95c409a03f20a"></a>

## certificate_url property — Property reference / df91404af0d7 / 5

Type: `"string"`. Required.

Certificate. Certificate or certificate chain in PEM format including the PEM headers.

Upstream description:

Certificate. Certificate or certificate chain in PEM format including the PEM headers.

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

- [custom_hash_algorithms](resources--certificate--reference--group-001.md#canonical-dd35da17fca7ce614baa82a836abf60594cb67b39a1df956d42103bbec25afe5): complete subsection reference.

<a id="canonical-905e8a0e073f0a9a5131b21e86d722a5351a2bbdcabbd61624451df610e27dc8"></a>

<a id="canonical-dfc9500a2e1450006be57fdc1afa4331628b91cd0448915ffbf646e810a34a71"></a>

## description property — Property reference / df91404af0d7 / 6

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

<a id="canonical-f2161fa14573f23e80f871a8e03bd1307c1d4cb1417fdad88b4fda450bddee4e"></a>

<a id="canonical-e91bb17fafa9bb9fadc7ba043d66d98e40118ae5233e424a17804a7d660fcdc9"></a>

## disable property — Property reference / df91404af0d7 / 7

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

- [disable_ocsp_stapling](resources--certificate--reference--group-001.md#canonical-ee98fcbd6a7c5294e6e7a467370541de3c9e695a32f199c5778784ad2fc88adf): complete subsection reference.

<a id="canonical-3b00ed83238bdf2ac86fc59297feb7544eac54ede8ad4fb5b10d3aa01a957a9d"></a>

<a id="canonical-2bfb308cadefa2bd65cfa2de22a9501be67341a5824b866d5f226ae0cbbf50f0"></a>

## id property — Property reference / df91404af0d7 / 8

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-4d7feaff747365673e8d962fbe3c8bc3d06fbf0d81e872e44be23d43a4e99434"></a>

<a id="canonical-60ab86a8d3da296813eb352820a6cc70963f1c3231120e2397e726e50e76c24a"></a>

## labels property — Property reference / df91404af0d7 / 9

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

<a id="canonical-1fc575ee736774b0444af2e108632d55bee9a24c904c1f6dfd12a3d30541e9a2"></a>

<a id="canonical-61a0dd858289210b6720d1f37a828e13bb623354a5719462089380dab9fb5294"></a>

## name property — Property reference / df91404af0d7 / 10

Type: `"string"`. Required.

Name of the Certificate. Must be unique within the namespace.

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

<a id="canonical-02b680c6153b719d3adbeaff0a8f98d5cdef37eb5eb6d97c3122fe1e3ecfe077"></a>

<a id="canonical-ae491327ce9904d662267a007da696c851ca2c8ef6e62956f0046d3dd9a34462"></a>

## namespace property — Property reference / df91404af0d7 / 11

Type: `"string"`. Required.

Namespace where the Certificate is created.

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

- [private_key](resources--certificate--reference--group-001.md#canonical-8e1aeaf2f75010b910dac72222e02d4aef67ae0727fe8b096bb549169e62257a): complete subsection reference.

- [timeouts](resources--certificate--reference--group-001.md#canonical-0acd593526a416499551f53e53d0d47348147a7e1c5d224854f97d055395082a): complete subsection reference.

- [use_system_defaults](resources--certificate--reference--group-001.md#canonical-4578c307b8c2fdeeb97d97669b6314c04314f99c2790d41d3c637d375083bc7a): complete subsection reference.

<a id="canonical-94bf18054e181526f32dcd3282c01d1d8f3a2ce3acc7678bd2134a9bcec1b71a"></a>

## All schema paths — Property reference / df91404af0d7 / 12

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--certificate--reference--group-001.md#canonical-38a354c2500381250bdfa3542585b468121a38608e96f4cc7c7ba1a6fa744422) |
| `certificate_chain` | [certificate_chain](resources--certificate--reference--group-001.md#canonical-7cc605dae2785377af1cc216447ef39e5db5d02e1dbf4d7c67ed5d9245cb92da) |
| `certificate_chain.name` | [certificate_chain.name](resources--certificate--reference--group-001.md#canonical-d04b4e0a81e539db3e5a70bba9b60da822c66e196c67af676e7898c6191f3437) |
| `certificate_chain.namespace` | [certificate_chain.namespace](resources--certificate--reference--group-001.md#canonical-f31636ba69f98fc536455a7582a3d230218c6cf396708500551e0d77553663e6) |
| `certificate_chain.tenant` | [certificate_chain.tenant](resources--certificate--reference--group-001.md#canonical-d9f3d29287a26a9cfeeb2b1253bfe6237d350248fc80533234c08f7823b2887e) |
| `certificate_url` | [certificate_url](resources--certificate--reference--group-001.md#canonical-18b9c9d3018dfa7d907a17dcfff0a875f0bc4fd44d8c9c91bd286340f3b4e497) |
| `custom_hash_algorithms` | [custom_hash_algorithms](resources--certificate--reference--group-001.md#canonical-3e10cfc3c833c2203ba630fd474e4b9d2a7b977c7447bc85706949feab6d7755) |
| `custom_hash_algorithms.hash_algorithms` | [custom_hash_algorithms.hash_algorithms](resources--certificate--reference--group-001.md#canonical-d4c3fb5cf4120771cb0c967fffb1fa73dd34ea31c41a974387610c27392f0d37) |
| `description` | [description](resources--certificate--reference--group-001.md#canonical-905e8a0e073f0a9a5131b21e86d722a5351a2bbdcabbd61624451df610e27dc8) |
| `disable` | [disable](resources--certificate--reference--group-001.md#canonical-f2161fa14573f23e80f871a8e03bd1307c1d4cb1417fdad88b4fda450bddee4e) |
| `disable_ocsp_stapling` | [disable_ocsp_stapling](resources--certificate--reference--group-001.md#canonical-d690b8b13cc58bf7f02de258e2f427dd44836084f25bd8bd58f643984c950fdc) |
| `id` | [id](resources--certificate--reference--group-001.md#canonical-3b00ed83238bdf2ac86fc59297feb7544eac54ede8ad4fb5b10d3aa01a957a9d) |
| `labels` | [labels](resources--certificate--reference--group-001.md#canonical-4d7feaff747365673e8d962fbe3c8bc3d06fbf0d81e872e44be23d43a4e99434) |
| `name` | [name](resources--certificate--reference--group-001.md#canonical-1fc575ee736774b0444af2e108632d55bee9a24c904c1f6dfd12a3d30541e9a2) |
| `namespace` | [namespace](resources--certificate--reference--group-001.md#canonical-02b680c6153b719d3adbeaff0a8f98d5cdef37eb5eb6d97c3122fe1e3ecfe077) |
| `private_key` | [private_key](resources--certificate--reference--group-001.md#canonical-1071d528055f1eb374284ded5e07000c8e8e071f125b13b15c0cdce1814e54bf) |
| `private_key.blindfold_secret_info` | [private_key.blindfold_secret_info](resources--certificate--reference--group-001.md#canonical-ad55a11619efb16db86d71dbb8be79188c2310b356ba0248a94aea9c9824517c) |
| `private_key.blindfold_secret_info.decryption_provider` | [private_key.blindfold_secret_info.decryption_provider](resources--certificate--reference--group-001.md#canonical-ac0885d104ac7ac10bab5dce4125b8331e8876624d4c7badaf95b8ab2df6dc49) |
| `private_key.blindfold_secret_info.location` | [private_key.blindfold_secret_info.location](resources--certificate--reference--group-001.md#canonical-4d2a4c083d91191117c71d427849f09cec9a59f6e9005a1acbd9f592b7bd98ee) |
| `private_key.blindfold_secret_info.store_provider` | [private_key.blindfold_secret_info.store_provider](resources--certificate--reference--group-001.md#canonical-4adb7c33fc6388366b917c5e048b8cb06ed7f78d01542d68cafe2e6a08c09522) |
| `private_key.clear_secret_info` | [private_key.clear_secret_info](resources--certificate--reference--group-001.md#canonical-2e40f25d20ab13cd50f9d3295f5348ea5a9625cf33a297f220989faa24daa622) |
| `private_key.clear_secret_info.provider_ref` | [private_key.clear_secret_info.provider_ref](resources--certificate--reference--group-001.md#canonical-7aca5acd3f8cdcfa77660a5f59402a87ff09a647ed4724f528ed47f2751fe237) |
| `private_key.clear_secret_info.url` | [private_key.clear_secret_info.url](resources--certificate--reference--group-001.md#canonical-ce2db1b815f97579827535386791e81039814a6e260305132f56997f022b02bc) |
| `timeouts` | [timeouts](resources--certificate--reference--group-001.md#canonical-00cdfe6c460a75540d36187eaa4be82adef4792db1dcf327248bbb0a1cb3c6e6) |
| `timeouts.create` | [timeouts.create](resources--certificate--reference--group-001.md#canonical-936d2cfc385148e3040912141aac8e2510dc405bda8b9be80d24cd460a3d1f80) |
| `timeouts.delete` | [timeouts.delete](resources--certificate--reference--group-001.md#canonical-42a2746bf82376a38ce2dde0337f480c2eac1258c420f3b1cbfd8a39345854ff) |
| `timeouts.read` | [timeouts.read](resources--certificate--reference--group-001.md#canonical-5fd40209355d05f26ee58d1abc5fcdf3a4d77886eed70266d5ab184afb32eb9b) |
| `timeouts.update` | [timeouts.update](resources--certificate--reference--group-001.md#canonical-5816f0e1809a41160f150294ed30fedfece62b40ef3beacf929214542da931bc) |
| `use_system_defaults` | [use_system_defaults](resources--certificate--reference--group-001.md#canonical-d1b0cb2d787186968818bb0283394c6259d5070301e550ad83220050b639c115) |

<a id="canonical-296cdaacc0da1e1a1f4d6cb58200446588612e72dfb153c09146d27aa47f2746"></a>

## Next pages — Property reference / df91404af0d7 / 13

- [certificate_chain](resources--certificate--reference--group-001.md#canonical-4ce00beeaeca753bf40bf153c80e3c73fe88090767a9268c8211471b9bc758f2)
- [custom_hash_algorithms](resources--certificate--reference--group-001.md#canonical-dd35da17fca7ce614baa82a836abf60594cb67b39a1df956d42103bbec25afe5)
- [disable_ocsp_stapling](resources--certificate--reference--group-001.md#canonical-ee98fcbd6a7c5294e6e7a467370541de3c9e695a32f199c5778784ad2fc88adf)
- [private_key](resources--certificate--reference--group-001.md#canonical-8e1aeaf2f75010b910dac72222e02d4aef67ae0727fe8b096bb549169e62257a)
- [timeouts](resources--certificate--reference--group-001.md#canonical-0acd593526a416499551f53e53d0d47348147a7e1c5d224854f97d055395082a)
- [use_system_defaults](resources--certificate--reference--group-001.md#canonical-4578c307b8c2fdeeb97d97669b6314c04314f99c2790d41d3c637d375083bc7a)
- [xcsh_certificate](../resources/certificate.md#canonical-843c7ef2d56e85f1ba6b346483224fce2e27fa0dc0213cd0e259331947e5281c)

<a id="canonical-4ce00beeaeca753bf40bf153c80e3c73fe88090767a9268c8211471b9bc758f2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-441b97866140cb0861e22655d55496b00aa54d8bb937c8235e69b8673d0002f0"></a>

## certificate_chain — certificate_chain / 50d63f39b82c / 2

Breadcrumbs:

- [xcsh_certificate](../resources/certificate.md#canonical-843c7ef2d56e85f1ba6b346483224fce2e27fa0dc0213cd0e259331947e5281c)
- [Property reference](resources--certificate--reference--group-001.md#canonical-6e191e632e125ee6cf3e18d3e730166767b3b520ab1f5c6d147e575ffbede87c)
- certificate_chain

<a id="canonical-7cc605dae2785377af1cc216447ef39e5db5d02e1dbf4d7c67ed5d9245cb92da"></a>

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
certificate_chain {
  # Configure direct properties listed below.
}
```

<a id="canonical-bc2a1ecf457bf9496927947237069975df7fd7243ab6db554c452bd683089cbe"></a>

## Direct properties — certificate_chain / 50d63f39b82c / 3

<a id="canonical-d04b4e0a81e539db3e5a70bba9b60da822c66e196c67af676e7898c6191f3437"></a>

<a id="canonical-397adce47acdcec8963e1e1f422bc1916e196c3832ed0bed6acfa09cf0ce98d2"></a>

## name property — certificate_chain / 50d63f39b82c / 4

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

<a id="canonical-f31636ba69f98fc536455a7582a3d230218c6cf396708500551e0d77553663e6"></a>

<a id="canonical-37aeef657a946a1286f0147bfad9a0f383c84cd98362a1fd7312efe90ae6178a"></a>

## namespace property — certificate_chain / 50d63f39b82c / 5

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

<a id="canonical-d9f3d29287a26a9cfeeb2b1253bfe6237d350248fc80533234c08f7823b2887e"></a>

<a id="canonical-e943d8950e564cd212f80598a789eba9ec941e2eca347760f473d781131ad8f8"></a>

## tenant property — certificate_chain / 50d63f39b82c / 6

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

<a id="canonical-fe165bd4fdea994b1e01985337f20dead47bbd544151cb70240cc1c6934f11be"></a>

## Next pages — certificate_chain / 50d63f39b82c / 7

- [Property reference](resources--certificate--reference--group-001.md#canonical-6e191e632e125ee6cf3e18d3e730166767b3b520ab1f5c6d147e575ffbede87c)
- [xcsh_certificate](../resources/certificate.md#canonical-843c7ef2d56e85f1ba6b346483224fce2e27fa0dc0213cd0e259331947e5281c)

<a id="canonical-dd35da17fca7ce614baa82a836abf60594cb67b39a1df956d42103bbec25afe5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9c6858bb5182bc5a72bad8bcf5ca149786459ffebe527061ce0cb08bd0eaa4ac"></a>

## custom_hash_algorithms — custom_hash_algorithms / e8475ac881b6 / 2

Breadcrumbs:

- [xcsh_certificate](../resources/certificate.md#canonical-843c7ef2d56e85f1ba6b346483224fce2e27fa0dc0213cd0e259331947e5281c)
- [Property reference](resources--certificate--reference--group-001.md#canonical-6e191e632e125ee6cf3e18d3e730166767b3b520ab1f5c6d147e575ffbede87c)
- custom_hash_algorithms

<a id="canonical-3e10cfc3c833c2203ba630fd474e4b9d2a7b977c7447bc85706949feab6d7755"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: custom\_hash\_algorithms, disable\_ocsp\_stapling, use\_system\_defaults; Default:
use\_system\_defaults\] Specifies the hash algorithms to be used.

Upstream description:

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

OneOf alternatives in this subsection:

- [custom_hash_algorithms](resources--certificate--reference--group-001.md#canonical-3e10cfc3c833c2203ba630fd474e4b9d2a7b977c7447bc85706949feab6d7755)
- [disable_ocsp_stapling](resources--certificate--reference--group-001.md#canonical-d690b8b13cc58bf7f02de258e2f427dd44836084f25bd8bd58f643984c950fdc)
- [use_system_defaults](resources--certificate--reference--group-001.md#canonical-d1b0cb2d787186968818bb0283394c6259d5070301e550ad83220050b639c115)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
custom_hash_algorithms {
  # Configure direct properties listed below.
}
```

<a id="canonical-57ccebf5bb40dbfdf157c61ba3c7a6390770bca71b0cb58ca402ba1fd11f9147"></a>

## Direct properties — custom_hash_algorithms / e8475ac881b6 / 3

<a id="canonical-d4c3fb5cf4120771cb0c967fffb1fa73dd34ea31c41a974387610c27392f0d37"></a>

<a id="canonical-8924969ab87bb10f6d55acbd1a51c31f2263240cdb465fa0e6fb543e83d1f20d"></a>

## hash_algorithms property — custom_hash_algorithms / e8475ac881b6 / 4

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

<a id="canonical-c15bcce0803d7dfe5d4d2370ffa6343d1974849074a467bba1fbca22d46dd94b"></a>

## Next pages — custom_hash_algorithms / e8475ac881b6 / 5

- [Property reference](resources--certificate--reference--group-001.md#canonical-6e191e632e125ee6cf3e18d3e730166767b3b520ab1f5c6d147e575ffbede87c)
- [xcsh_certificate](../resources/certificate.md#canonical-843c7ef2d56e85f1ba6b346483224fce2e27fa0dc0213cd0e259331947e5281c)

<a id="canonical-ee98fcbd6a7c5294e6e7a467370541de3c9e695a32f199c5778784ad2fc88adf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-639bfd88aaf09d74447818031d3c3499105ab4e1cdc8bcb20b5a5d70a43f38d7"></a>

## disable_ocsp_stapling — disable_ocsp_stapling / 8d8164e87821 / 2

Breadcrumbs:

- [xcsh_certificate](../resources/certificate.md#canonical-843c7ef2d56e85f1ba6b346483224fce2e27fa0dc0213cd0e259331947e5281c)
- [Property reference](resources--certificate--reference--group-001.md#canonical-6e191e632e125ee6cf3e18d3e730166767b3b520ab1f5c6d147e575ffbede87c)
- disable_ocsp_stapling

<a id="canonical-d690b8b13cc58bf7f02de258e2f427dd44836084f25bd8bd58f643984c950fdc"></a>

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

<a id="canonical-ac2f627c8cadbb6866d909469ca049b75304b81c490d3695a77b3acca229f52c"></a>

## Direct properties — disable_ocsp_stapling / 8d8164e87821 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4198015e86c0fcaf2609a0193c5becf8591734b15ed9789702aa51107c633540"></a>

## Next pages — disable_ocsp_stapling / 8d8164e87821 / 4

- [Property reference](resources--certificate--reference--group-001.md#canonical-6e191e632e125ee6cf3e18d3e730166767b3b520ab1f5c6d147e575ffbede87c)
- [xcsh_certificate](../resources/certificate.md#canonical-843c7ef2d56e85f1ba6b346483224fce2e27fa0dc0213cd0e259331947e5281c)

<a id="canonical-8e1aeaf2f75010b910dac72222e02d4aef67ae0727fe8b096bb549169e62257a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bba6c3985e1755321dd0ea34b9480953a9fd8aa59897f3776febbcc6f4be334b"></a>

## private_key — private_key / 304a28227938 / 2

Breadcrumbs:

- [xcsh_certificate](../resources/certificate.md#canonical-843c7ef2d56e85f1ba6b346483224fce2e27fa0dc0213cd0e259331947e5281c)
- [Property reference](resources--certificate--reference--group-001.md#canonical-6e191e632e125ee6cf3e18d3e730166767b3b520ab1f5c6d147e575ffbede87c)
- private_key

<a id="canonical-1071d528055f1eb374284ded5e07000c8e8e071f125b13b15c0cdce1814e54bf"></a>

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

<a id="canonical-159babbdfe42c9292a72f2777b1837a394a906b58a168e029b02a9da1b4abc1f"></a>

## Direct properties — private_key / 304a28227938 / 3

- [blindfold_secret_info](resources--certificate--reference--group-001.md#canonical-6671f5d8e7dc6a68540619f79d5bd8ad88b37b663cd1c285583387e9c10fb46f): complete subsection reference.

- [clear_secret_info](resources--certificate--reference--group-001.md#canonical-c270908703b45866948ee8d4bda9f10a48d40e21975638a877c06812a1bdacd2): complete subsection reference.

<a id="canonical-fe75a7fefbe3d06a11ff4860998fcb9eac90bc8ec83a9d632747398738a7438b"></a>

## Next pages — private_key / 304a28227938 / 4

- [private_key.blindfold_secret_info](resources--certificate--reference--group-001.md#canonical-6671f5d8e7dc6a68540619f79d5bd8ad88b37b663cd1c285583387e9c10fb46f)
- [private_key.clear_secret_info](resources--certificate--reference--group-001.md#canonical-c270908703b45866948ee8d4bda9f10a48d40e21975638a877c06812a1bdacd2)
- [Property reference](resources--certificate--reference--group-001.md#canonical-6e191e632e125ee6cf3e18d3e730166767b3b520ab1f5c6d147e575ffbede87c)
- [xcsh_certificate](../resources/certificate.md#canonical-843c7ef2d56e85f1ba6b346483224fce2e27fa0dc0213cd0e259331947e5281c)

<a id="canonical-6671f5d8e7dc6a68540619f79d5bd8ad88b37b663cd1c285583387e9c10fb46f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-64f1df444955e66c761c4242af56d789ef1c6d7254063d44a74d41a6cad6ed8e"></a>

## private_key.blindfold_secret_info — private_key.blindfold_secret_info / 6d34bd2cd150 / 2

Breadcrumbs:

- [xcsh_certificate](../resources/certificate.md#canonical-843c7ef2d56e85f1ba6b346483224fce2e27fa0dc0213cd0e259331947e5281c)
- [Property reference](resources--certificate--reference--group-001.md#canonical-6e191e632e125ee6cf3e18d3e730166767b3b520ab1f5c6d147e575ffbede87c)
- [private_key](resources--certificate--reference--group-001.md#canonical-8e1aeaf2f75010b910dac72222e02d4aef67ae0727fe8b096bb549169e62257a)
- private_key.blindfold_secret_info

<a id="canonical-ad55a11619efb16db86d71dbb8be79188c2310b356ba0248a94aea9c9824517c"></a>

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

<a id="canonical-63bc21d464447e4360b9a561417cc71a07c808699a74b22dbdadd29014d0ed09"></a>

## Direct properties — private_key.blindfold_secret_info / 6d34bd2cd150 / 3

<a id="canonical-ac0885d104ac7ac10bab5dce4125b8331e8876624d4c7badaf95b8ab2df6dc49"></a>

<a id="canonical-98d9f92b9e93693448f04a67b81504809e60955c920ea193349c0f70e44aff38"></a>

## decryption_provider property — private_key.blindfold_secret_info / 6d34bd2cd150 / 4

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

<a id="canonical-4d2a4c083d91191117c71d427849f09cec9a59f6e9005a1acbd9f592b7bd98ee"></a>

<a id="canonical-195d14e0b92e7f58603dec167fc670fac6919f25d76c1698d50d5e850f2bdbd1"></a>

## location property — private_key.blindfold_secret_info / 6d34bd2cd150 / 5

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

<a id="canonical-4adb7c33fc6388366b917c5e048b8cb06ed7f78d01542d68cafe2e6a08c09522"></a>

<a id="canonical-a396bfacd7c92b4edcc55d7f61ebe2d015bb35989ae563d57e27296dc0b0d896"></a>

## store_provider property — private_key.blindfold_secret_info / 6d34bd2cd150 / 6

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

<a id="canonical-8c8f0fd7fefd176a62b7983ffe7f7fe17f1ea97d390c495e6567ad55e32b9252"></a>

## Next pages — private_key.blindfold_secret_info / 6d34bd2cd150 / 7

- [private_key](resources--certificate--reference--group-001.md#canonical-8e1aeaf2f75010b910dac72222e02d4aef67ae0727fe8b096bb549169e62257a)
- [xcsh_certificate](../resources/certificate.md#canonical-843c7ef2d56e85f1ba6b346483224fce2e27fa0dc0213cd0e259331947e5281c)

<a id="canonical-c270908703b45866948ee8d4bda9f10a48d40e21975638a877c06812a1bdacd2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7fe688491d2817512e8d694f95af86e83c78d3f1ab66e90bcd88f077c0c06fec"></a>

## private_key.clear_secret_info — private_key.clear_secret_info / 6078a0feb22a / 2

Breadcrumbs:

- [xcsh_certificate](../resources/certificate.md#canonical-843c7ef2d56e85f1ba6b346483224fce2e27fa0dc0213cd0e259331947e5281c)
- [Property reference](resources--certificate--reference--group-001.md#canonical-6e191e632e125ee6cf3e18d3e730166767b3b520ab1f5c6d147e575ffbede87c)
- [private_key](resources--certificate--reference--group-001.md#canonical-8e1aeaf2f75010b910dac72222e02d4aef67ae0727fe8b096bb549169e62257a)
- private_key.clear_secret_info

<a id="canonical-2e40f25d20ab13cd50f9d3295f5348ea5a9625cf33a297f220989faa24daa622"></a>

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

<a id="canonical-652ed76752453cf5c8b5cf0c122e00d893fda6d2c76c6b07651bb13e70e17569"></a>

## Direct properties — private_key.clear_secret_info / 6078a0feb22a / 3

<a id="canonical-7aca5acd3f8cdcfa77660a5f59402a87ff09a647ed4724f528ed47f2751fe237"></a>

<a id="canonical-4f6435ef6b8cdaf31767999de06a7ccb1b4989fd5130b760de4f605206901261"></a>

## provider_ref property — private_key.clear_secret_info / 6078a0feb22a / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-ce2db1b815f97579827535386791e81039814a6e260305132f56997f022b02bc"></a>

<a id="canonical-b6af0926985006297fa2fc3e3e2c20ce138d69c101a97a450a50289852e61cf4"></a>

## url property — private_key.clear_secret_info / 6078a0feb22a / 5

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

<a id="canonical-5aa5ec369cdffcea379ac25ac24f422cab05cd9803b83aadf9577be956402e6f"></a>

## Next pages — private_key.clear_secret_info / 6078a0feb22a / 6

- [private_key](resources--certificate--reference--group-001.md#canonical-8e1aeaf2f75010b910dac72222e02d4aef67ae0727fe8b096bb549169e62257a)
- [xcsh_certificate](../resources/certificate.md#canonical-843c7ef2d56e85f1ba6b346483224fce2e27fa0dc0213cd0e259331947e5281c)

<a id="canonical-0acd593526a416499551f53e53d0d47348147a7e1c5d224854f97d055395082a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6f01089e523970388dd4270df8e2c0836506f6413fb1946340b198941197c043"></a>

## timeouts — timeouts / 699c2acec6a0 / 2

Breadcrumbs:

- [xcsh_certificate](../resources/certificate.md#canonical-843c7ef2d56e85f1ba6b346483224fce2e27fa0dc0213cd0e259331947e5281c)
- [Property reference](resources--certificate--reference--group-001.md#canonical-6e191e632e125ee6cf3e18d3e730166767b3b520ab1f5c6d147e575ffbede87c)
- timeouts

<a id="canonical-00cdfe6c460a75540d36187eaa4be82adef4792db1dcf327248bbb0a1cb3c6e6"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-b476a75228f8a57edba3399b9c77416a08cc05d4b4f0b423def798049551cd77"></a>

## Direct properties — timeouts / 699c2acec6a0 / 3

<a id="canonical-936d2cfc385148e3040912141aac8e2510dc405bda8b9be80d24cd460a3d1f80"></a>

<a id="canonical-9f1acd64237f63ede7dad14be99eca9e40dffadc86042f15d2e6bc34dc5db56e"></a>

## create property — timeouts / 699c2acec6a0 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-42a2746bf82376a38ce2dde0337f480c2eac1258c420f3b1cbfd8a39345854ff"></a>

<a id="canonical-cd5437f73c4bf191fbbab93ba7c78832646254a3488ea20bc1caefb7a631f1b6"></a>

## delete property — timeouts / 699c2acec6a0 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-5fd40209355d05f26ee58d1abc5fcdf3a4d77886eed70266d5ab184afb32eb9b"></a>

<a id="canonical-db4ecdfce499d841d9c9ec20f8d75aa1fa68e12103956776612ea0f50411a3dd"></a>

## read property — timeouts / 699c2acec6a0 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-5816f0e1809a41160f150294ed30fedfece62b40ef3beacf929214542da931bc"></a>

<a id="canonical-0f39a353397305c6fc92338f98050abf22679e2b6de27d80185a7e325f1bacf5"></a>

## update property — timeouts / 699c2acec6a0 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-7f4ce781c44e48495f0394ea919fd9e37c67e71af46b2293d63f11c46b1fab36"></a>

## Next pages — timeouts / 699c2acec6a0 / 8

- [Property reference](resources--certificate--reference--group-001.md#canonical-6e191e632e125ee6cf3e18d3e730166767b3b520ab1f5c6d147e575ffbede87c)
- [xcsh_certificate](../resources/certificate.md#canonical-843c7ef2d56e85f1ba6b346483224fce2e27fa0dc0213cd0e259331947e5281c)

<a id="canonical-4578c307b8c2fdeeb97d97669b6314c04314f99c2790d41d3c637d375083bc7a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a69c41047bc30aa8db1c3ae229493d6b002e939f35dc07b682986b6f5b7d09e0"></a>

## use_system_defaults — use_system_defaults / 5f9b1e548f2d / 2

Breadcrumbs:

- [xcsh_certificate](../resources/certificate.md#canonical-843c7ef2d56e85f1ba6b346483224fce2e27fa0dc0213cd0e259331947e5281c)
- [Property reference](resources--certificate--reference--group-001.md#canonical-6e191e632e125ee6cf3e18d3e730166767b3b520ab1f5c6d147e575ffbede87c)
- use_system_defaults

<a id="canonical-d1b0cb2d787186968818bb0283394c6259d5070301e550ad83220050b639c115"></a>

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

<a id="canonical-fdaf608ccfde3eb93ac167a5f89a9443c2193cfeeaf7e772612c47313e1819db"></a>

## Direct properties — use_system_defaults / 5f9b1e548f2d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7ac08f047e63b3ff2dcb4e7e4ea6fac2be3a7f5b9642870a89c72812734824e8"></a>

## Next pages — use_system_defaults / 5f9b1e548f2d / 4

- [Property reference](resources--certificate--reference--group-001.md#canonical-6e191e632e125ee6cf3e18d3e730166767b3b520ab1f5c6d147e575ffbede87c)
- [xcsh_certificate](../resources/certificate.md#canonical-843c7ef2d56e85f1ba6b346483224fce2e27fa0dc0213cd0e259331947e5281c)
