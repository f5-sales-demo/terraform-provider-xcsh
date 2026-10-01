---
page_title: "xcsh_bigip_http_proxy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bigip_http_proxy reference."
---

# xcsh_bigip_http_proxy reference

<a id="canonical-f7b9d6606bc8565cf6a0e7abd9c21678e5b551276e02d51a74729b1514f1308e"></a>

## namespace property — proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site / 71de83fa7095 / 5

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

<a id="canonical-2d799fba917c211ba6ff4df94d8d04df88c29164838aee9968975cde1d73f277"></a>

<a id="canonical-e842c45336754bf0aeed683c924727ae15eecbdfd06a6424ec8a00472541b57c"></a>

## tenant property — proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site / 71de83fa7095 / 6

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

<a id="canonical-c67bed7cbbdeec02c67df48be90ee327b5aa57f147848db691948db979f9de26"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site / 71de83fa7095 / 7

- [proxy_advertisement.advertise_custom.advertise_where.vk8s_service](resources--bigip_http_proxy--reference--group-002.md#canonical-e11a4d6a07eb2cc00ffc458854833379b43949e4abab740e6a8c423949a6674e)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-a1324d8ef4347390b9f1e100b6ece47ce241b1c260e0eb27d90921134eca1262"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-22b2dbdb4607cbd29d30922e72400648c867bc7eaa54a880fca21dc14c488098"></a>

## proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site — proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site / 057186367dc3 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_advertisement](resources--bigip_http_proxy--reference--group-002.md#canonical-2a89bae749fbef43ff1ed28b601d6404c949ab646437965ef3776fba8218ad5a)
- [proxy_advertisement.advertise_custom](resources--bigip_http_proxy--reference--group-002.md#canonical-03c906c318cb23d88fa6dc4c79d25c56ba8d805b6ce1151d6c3faeb58afd76fc)
- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--reference--group-002.md#canonical-14ff5e5a12ebe72a0729ad520c70dc8b1e16ed514b4a72e99b09379e44f920cf)
- [proxy_advertisement.advertise_custom.advertise_where.vk8s_service](resources--bigip_http_proxy--reference--group-002.md#canonical-e11a4d6a07eb2cc00ffc458854833379b43949e4abab740e6a8c423949a6674e)
- proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site

<a id="canonical-a5d8c5c3598c9a72dcd703a621be8e42c70dee85707ed16b308415c46658101c"></a>

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

<a id="canonical-1f8f01f07e4f56cd9da8e9aab4ee1aeb7f9b2be81a93c31dd84138d252a39af5"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site / 057186367dc3 / 3

<a id="canonical-8bd4f4c16f8b5a7674a4faee079e7139b6e52458198524060d6645c532824b3a"></a>

<a id="canonical-a20a9f420523197ea7f1999735a386e0a2cf92e2e3d6465091e02cd88b689925"></a>

## name property — proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site / 057186367dc3 / 4

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

<a id="canonical-8669ae18b658509fc3e58c1567b0b9f8e0c468ae3b2900611ab306a02e201d87"></a>

<a id="canonical-0fa26e42f0dd740e226547b6cb5e421892ae414be5b501fc96629cf2d6607601"></a>

## namespace property — proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site / 057186367dc3 / 5

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

<a id="canonical-38982c279fb87666a347dd2f6327e5f5b8e50904a6fae3bd2e8aa4df7f9e33a6"></a>

<a id="canonical-c5e18e1464db604569e58b0fbff28299aa34510853062cfbb634468fcba39106"></a>

## tenant property — proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site / 057186367dc3 / 6

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

<a id="canonical-5b538aafbd773825df21a1efd42e5325be5b517c36d253024a8459b0cf730352"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site / 057186367dc3 / 7

- [proxy_advertisement.advertise_custom.advertise_where.vk8s_service](resources--bigip_http_proxy--reference--group-002.md#canonical-e11a4d6a07eb2cc00ffc458854833379b43949e4abab740e6a8c423949a6674e)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-7e3d9d10858cdc4cd01a24952d34d04710876f9cb534696e3342242cd698fd1a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4c4d53305d7ba91c8fbbe94ed850a8e4a51916698d4cab9b23055ade5729f1e3"></a>

## proxy_advertisement.do_not_advertise — proxy_advertisement.do_not_advertise / ca3b480ba64e / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_advertisement](resources--bigip_http_proxy--reference--group-002.md#canonical-2a89bae749fbef43ff1ed28b601d6404c949ab646437965ef3776fba8218ad5a)
- proxy_advertisement.do_not_advertise

<a id="canonical-48c812065580a2ea07b05dd8c53411a5e8a56996f7519141fc161d47ec6fb1e4"></a>

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

<a id="canonical-d9719801b553666153dc9717cfbe02854df36a371d6c9331df56136ac9362675"></a>

## Direct properties — proxy_advertisement.do_not_advertise / ca3b480ba64e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9ffea0e4cb68b067098f94f7457ba0c73d0243a98a1146be59c0f4842cf28255"></a>

## Next pages — proxy_advertisement.do_not_advertise / ca3b480ba64e / 4

- [proxy_advertisement](resources--bigip_http_proxy--reference--group-002.md#canonical-2a89bae749fbef43ff1ed28b601d6404c949ab646437965ef3776fba8218ad5a)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c4982e3002edfd179917e844d748cfa314a4426ba5b145601c58ca475fac988c"></a>

## proxy_config — proxy_config / f137c760b3b1 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- proxy_config

<a id="canonical-393817bce832cba65ce788bb455b93d2361a3acf7d70c0c209d16ce7fe906499"></a>

Type: `"object"`. single nested block, Optional.

HTTP/HTTPS Load Balancer. HTTP/HTTPS Load balancer.

Upstream description:

HTTP/HTTPS Load balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("domains"),
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
  "x-ves-oneof-field-loadbalancer_type": "[\"http\",\"https\",\"https_auto_cert\"]"
}
```

Terraform syntax:

```terraform
proxy_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-fe9efb989d90594fe6b5f67121c170f8d2c0f42aa22d14dd495b93e6c129d951"></a>

## Direct properties — proxy_config / f137c760b3b1 / 3

<a id="canonical-3b37bc515b8a8d14272c48d565783cc5ff6c219707ca2c5f4f09ea8c71a80009"></a>

<a id="canonical-ce5609910ee5d67879f84821784d7dbcd4852acb1f82a5171b0a8dd0e4423971"></a>

## domains property — proxy_config / f137c760b3b1 / 4

Type: `["list", "string"]`. Optional.

List of domains (host/authority header) that will be matched to loadbalancer. Wildcard hosts are
supported in the suffix or prefix form Domain search order: 1. Exact domain names: \`\`.

Upstream description:

A list of domains (host/authority header) that will be matched to loadbalancer. Wildcard hosts are
supported in the suffix or prefix form

Domain search order: &#8203;1. Exact domain names: \`\`www&#46;example.com\`\`. &#8203;2. Prefix
domain wildcards: \`\`\*.example.com\`\` or \`\`\*-bar.example.com\`\`. &#8203;3. Special wildcard
\`\`\*\`\` matching any domain.

Wildcard will not match empty string. E.g. \`\`\*-bar.example.com\`\` will match
\`\`baz-bar.example.com\`\` but not \`\`-bar.example.com\`\`. The longest wildcards match first.

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

- [http](resources--bigip_http_proxy--reference--group-003.md#canonical-d8687e630241c8177883e5aed4912e2d7f7fa76e042a9779a9f447f7d73ea296): complete subsection reference.

- [https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655): complete subsection reference.

- [https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-a89b29ce5d64616933b19dfd941b49ab97fb290e7c1567333a820ab327fa5df6): complete subsection reference.

<a id="canonical-8ac3294c997ebf2967e88013d231cb9755b79ccd2c4b11c8baa98a70d1fb7ba4"></a>

## Next pages — proxy_config / f137c760b3b1 / 5

- [proxy_config.http](resources--bigip_http_proxy--reference--group-003.md#canonical-d8687e630241c8177883e5aed4912e2d7f7fa76e042a9779a9f447f7d73ea296)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-a89b29ce5d64616933b19dfd941b49ab97fb290e7c1567333a820ab327fa5df6)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-d8687e630241c8177883e5aed4912e2d7f7fa76e042a9779a9f447f7d73ea296"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c39201c6c73eecda897a72349d7dc8610f89c29c6ea40ffda2d861428ef309aa"></a>

## proxy_config.http — proxy_config.http / 5c6b8f4435fd / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- proxy_config.http

<a id="canonical-8fdb3823d250b342a9c3491ea57433fa1cfd2500dafdee5de3bbd686e65f0fb7"></a>

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

<a id="canonical-6ed673511d9ec58ffec8b52dca2a4aff5a308da1c901e5760635c7cacc49d090"></a>

## Direct properties — proxy_config.http / 5c6b8f4435fd / 3

<a id="canonical-5c57f3767667a2a27383c68908b9fab5a911285d6efc2b104c9d64e053bab797"></a>

<a id="canonical-f3594bd2e143fc99dfbd347873e513583a45c418f27283d7bc682ef3a0c24138"></a>

## dns_volterra_managed property — proxy_config.http / 5c6b8f4435fd / 4

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

<a id="canonical-0ef545bfee9afd0dbe5e8efc58d11fc4e0f9ef0223eb5bc7064ba88ffbae689d"></a>

<a id="canonical-fed1a0979e3fa1f688a4869e88b40f695b4ae2bdb9569ec572f18f45ad814125"></a>

## port property — proxy_config.http / 5c6b8f4435fd / 5

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

<a id="canonical-96028dd8a68d42d01baa5b1c7160ec78ade35f0f109aa51588b8e5784d0e01af"></a>

<a id="canonical-c54a1cdfe406936689f0d5de6982f503c96b9523e1ecc5c16fca7bcd913ee70b"></a>

## port_ranges property — proxy_config.http / 5c6b8f4435fd / 6

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

<a id="canonical-67c9e69716beda8c4ed806947ac93b04a7faaa6e00918c39c6c4806b7bd8b42a"></a>

## Next pages — proxy_config.http / 5c6b8f4435fd / 7

- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-af713cdb41ddc6287202a40f3affc6c1352cf612f5a43cae22647af01b9cb3d5"></a>

## proxy_config.https — proxy_config.https / 9a635d52ce42 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- proxy_config.https

<a id="canonical-d7ba2d6b1c65e156882a61906e192ab13cac2784851176fbfe6470d06d3748b1"></a>

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

<a id="canonical-6359a547f16f1fc0f922eed48028b4480c8a1b9d739eb86bf01b9f57e1051b3a"></a>

## Direct properties — proxy_config.https / 9a635d52ce42 / 3

<a id="canonical-05bc32e787c935e6e0af7bbb56b42607056d87f02cfd0b2d5c751a90834d9008"></a>

<a id="canonical-a84698c54b1bb70379327144ec31975642ca535996a371a4de4be92a02709a95"></a>

## add_hsts property — proxy_config.https / 9a635d52ce42 / 4

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

<a id="canonical-3c03a18964c054e9682e6280faafd9fc270cb4fd5733ab6f186d75568031cc0b"></a>

<a id="canonical-964c6df07b4d1f05e9bd6b210de33e73263909711e0960aed78ac6aa8b84f9bb"></a>

## append_server_name property — proxy_config.https / 9a635d52ce42 / 5

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

- [coalescing_options](resources--bigip_http_proxy--reference--group-003.md#canonical-281c3500456d9e1b2541b735df813a8712ef0dfdfec4375183190b13d1e7583f): complete subsection reference.

<a id="canonical-ed67099a35b039de36d96b1d7597e070ac6250adb46254ae8d1f587a9109a00c"></a>

<a id="canonical-3887ae846003d8def861eb087598cb92b227f35c84730dd6b183f90e6a2411b5"></a>

## connection_idle_timeout property — proxy_config.https / 9a635d52ce42 / 6

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

- [default_header](resources--bigip_http_proxy--reference--group-003.md#canonical-977f3843be123c631179fac89316a67dfd380f0e5a369290780fb38b8e91242e): complete subsection reference.

- [default_loadbalancer](resources--bigip_http_proxy--reference--group-003.md#canonical-5007f8764bf4b43e268bb25251f5abba2d21f1cff7fa04d7042a07e4cc8f1eca): complete subsection reference.

- [disable_path_normalize](resources--bigip_http_proxy--reference--group-003.md#canonical-49a22d9a7a36911612c707dc895cb4c8d615ae671e0a9ec91865f97b1e1f749d): complete subsection reference.

- [enable_path_normalize](resources--bigip_http_proxy--reference--group-003.md#canonical-629d86b3f1b2fbf2565410dde83f9fae730d3ab1eb3dd0ff6d8166f8706f14d8): complete subsection reference.

- [http_protocol_options](resources--bigip_http_proxy--reference--group-003.md#canonical-604ec47f9724ee98fbe6f227f219fa4eaac8145347b62c87cdbdeef5f0984317): complete subsection reference.

<a id="canonical-45b3acd581d120085b3fed2d4ce43f254c9027c201499eb83f685b61869e7383"></a>

<a id="canonical-b2b08be2904d8c7eac56d5b56cefbe4c87766f93f977ebeb6b40144655a421a2"></a>

## http_redirect property — proxy_config.https / 9a635d52ce42 / 7

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

- [non_default_loadbalancer](resources--bigip_http_proxy--reference--group-003.md#canonical-1a2189c142e8575942c881f77f36e23be9b4a082bc0437c9ffa4804fe8eff103): complete subsection reference.

- [pass_through](resources--bigip_http_proxy--reference--group-003.md#canonical-b4ab5c9199515937cf6af77ff8d62d8d5299643598ed6c99f7bd0892debaede4): complete subsection reference.

<a id="canonical-77a526a48086f81d72a0b81a2040624b091efd2a1bb2986bddcc79b7df36b1b8"></a>

<a id="canonical-719f8a88827bfb0f9505ce772bcf7ec8f898545f3edbb29af5535f44857059bd"></a>

## port property — proxy_config.https / 9a635d52ce42 / 8

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

<a id="canonical-dc2102357c1f30f1df1ca5dea847ba193c5756a883087c1d10eb293c611f1de8"></a>

<a id="canonical-2539b0dc57c9e180996f6b61a49372091a27148244dd9c873fa94f7216c6b4fa"></a>

## port_ranges property — proxy_config.https / 9a635d52ce42 / 9

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

<a id="canonical-cdfa91fd445288e4601e1918915db6081b9357e3bab22d712b70883744a4eb9f"></a>

<a id="canonical-53bf62c15488b619e38dfd276dc06c826d428f3e076713ce44dbfa9b3f6ca0a2"></a>

## server_name property — proxy_config.https / 9a635d52ce42 / 10

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

- [tls_cert_params](resources--bigip_http_proxy--reference--group-003.md#canonical-5bd61375e16f524597e4f2a07869259ce316df6e104e564e7919a5e5a7f76a12): complete subsection reference.

- [tls_parameters](resources--bigip_http_proxy--reference--group-003.md#canonical-1a2fff832c914deeb21855b2b6d870103e691b9c714316eef87fcff21a04a187): complete subsection reference.

<a id="canonical-c4cfa85c5c1b4ba04d636396fc9ab4cea3525e3afd11ee5f295a2925e448377e"></a>

## Next pages — proxy_config.https / 9a635d52ce42 / 11

- [proxy_config.https.coalescing_options](resources--bigip_http_proxy--reference--group-003.md#canonical-281c3500456d9e1b2541b735df813a8712ef0dfdfec4375183190b13d1e7583f)
- [proxy_config.https.default_header](resources--bigip_http_proxy--reference--group-003.md#canonical-977f3843be123c631179fac89316a67dfd380f0e5a369290780fb38b8e91242e)
- [proxy_config.https.default_loadbalancer](resources--bigip_http_proxy--reference--group-003.md#canonical-5007f8764bf4b43e268bb25251f5abba2d21f1cff7fa04d7042a07e4cc8f1eca)
- [proxy_config.https.disable_path_normalize](resources--bigip_http_proxy--reference--group-003.md#canonical-49a22d9a7a36911612c707dc895cb4c8d615ae671e0a9ec91865f97b1e1f749d)
- [proxy_config.https.enable_path_normalize](resources--bigip_http_proxy--reference--group-003.md#canonical-629d86b3f1b2fbf2565410dde83f9fae730d3ab1eb3dd0ff6d8166f8706f14d8)
- [proxy_config.https.http_protocol_options](resources--bigip_http_proxy--reference--group-003.md#canonical-604ec47f9724ee98fbe6f227f219fa4eaac8145347b62c87cdbdeef5f0984317)
- [proxy_config.https.non_default_loadbalancer](resources--bigip_http_proxy--reference--group-003.md#canonical-1a2189c142e8575942c881f77f36e23be9b4a082bc0437c9ffa4804fe8eff103)
- [proxy_config.https.pass_through](resources--bigip_http_proxy--reference--group-003.md#canonical-b4ab5c9199515937cf6af77ff8d62d8d5299643598ed6c99f7bd0892debaede4)
- [proxy_config.https.tls_cert_params](resources--bigip_http_proxy--reference--group-003.md#canonical-5bd61375e16f524597e4f2a07869259ce316df6e104e564e7919a5e5a7f76a12)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-003.md#canonical-1a2fff832c914deeb21855b2b6d870103e691b9c714316eef87fcff21a04a187)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-281c3500456d9e1b2541b735df813a8712ef0dfdfec4375183190b13d1e7583f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-49523d097eb7d19faf76fd02bfcf128feeb0d6e6bda86e7126c44c872225d82b"></a>

## proxy_config.https.coalescing_options — proxy_config.https.coalescing_options / ac34b3a240bf / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- proxy_config.https.coalescing_options

<a id="canonical-5b613686c29e6e3e67cea41fb7f261da9c867d208b9de70de63ed947f3b85d1f"></a>

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

<a id="canonical-816bf9a4f2ac3ef195b60904ef7b60fd065eec237001927f3f00804689ef8614"></a>

## Direct properties — proxy_config.https.coalescing_options / ac34b3a240bf / 3

- [default_coalescing](resources--bigip_http_proxy--reference--group-003.md#canonical-764c466fa5e369a3fa9706ff30d5c7a73e2553ab0a8823e73b323b38cb3979ae): complete subsection reference.

- [strict_coalescing](resources--bigip_http_proxy--reference--group-003.md#canonical-992fa2bf1ada6ca81ec7b4192fef982f1b0fa94ca3d6f76d7c7ec6445d7c69d6): complete subsection reference.

<a id="canonical-6a56e3b48bb25d68cbe74642146d5cae95e01426fc56e973fdfa58f144c15e99"></a>

## Next pages — proxy_config.https.coalescing_options / ac34b3a240bf / 4

- [proxy_config.https.coalescing_options.default_coalescing](resources--bigip_http_proxy--reference--group-003.md#canonical-764c466fa5e369a3fa9706ff30d5c7a73e2553ab0a8823e73b323b38cb3979ae)
- [proxy_config.https.coalescing_options.strict_coalescing](resources--bigip_http_proxy--reference--group-003.md#canonical-992fa2bf1ada6ca81ec7b4192fef982f1b0fa94ca3d6f76d7c7ec6445d7c69d6)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-764c466fa5e369a3fa9706ff30d5c7a73e2553ab0a8823e73b323b38cb3979ae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-48e320fdad707b4d35ec7929f6cdd5f54287ec0aef9f5cbb1cd3640c8febd9f2"></a>

## proxy_config.https.coalescing_options.default_coalescing — proxy_config.https.coalescing_options.default_coalescing / 1e128e29fae2 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- [proxy_config.https.coalescing_options](resources--bigip_http_proxy--reference--group-003.md#canonical-281c3500456d9e1b2541b735df813a8712ef0dfdfec4375183190b13d1e7583f)
- proxy_config.https.coalescing_options.default_coalescing

<a id="canonical-578508cb3a88168d0287502e5811cfd22662681e12c9ce24c39b74ba8d74886f"></a>

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

<a id="canonical-079b667f6b8c6db84b4162f32a07fe8dd5c27a231266dd6ea3281c7a48458d43"></a>

## Direct properties — proxy_config.https.coalescing_options.default_coalescing / 1e128e29fae2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-db8927033c25473b339217247e8a01f88c42dde3940955b95b98de31c6f89270"></a>

## Next pages — proxy_config.https.coalescing_options.default_coalescing / 1e128e29fae2 / 4

- [proxy_config.https.coalescing_options](resources--bigip_http_proxy--reference--group-003.md#canonical-281c3500456d9e1b2541b735df813a8712ef0dfdfec4375183190b13d1e7583f)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-992fa2bf1ada6ca81ec7b4192fef982f1b0fa94ca3d6f76d7c7ec6445d7c69d6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fb3e7229be6c2f31a53137796b0c2b5c3a6f339cd5ebe47e1e11a340d4288472"></a>

## proxy_config.https.coalescing_options.strict_coalescing — proxy_config.https.coalescing_options.strict_coalescing / 3db0d53e950b / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- [proxy_config.https.coalescing_options](resources--bigip_http_proxy--reference--group-003.md#canonical-281c3500456d9e1b2541b735df813a8712ef0dfdfec4375183190b13d1e7583f)
- proxy_config.https.coalescing_options.strict_coalescing

<a id="canonical-73aa8b2e7167f857e5a481420a7ea43aebf799380efb952b442f5b377d7ce051"></a>

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

<a id="canonical-b7df30f1e8743e6727e0dc4819def4b69414f8587d5c9fbf7a588e8ebe72e61b"></a>

## Direct properties — proxy_config.https.coalescing_options.strict_coalescing / 3db0d53e950b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5dcb128da304fb03876bec6a08f65c097c75768ba928e7552552f41872041a54"></a>

## Next pages — proxy_config.https.coalescing_options.strict_coalescing / 3db0d53e950b / 4

- [proxy_config.https.coalescing_options](resources--bigip_http_proxy--reference--group-003.md#canonical-281c3500456d9e1b2541b735df813a8712ef0dfdfec4375183190b13d1e7583f)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-977f3843be123c631179fac89316a67dfd380f0e5a369290780fb38b8e91242e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a11848928a53627e1b45d022431db49e7a2dc244b869ef3df6955bd3b106f6e0"></a>

## proxy_config.https.default_header — proxy_config.https.default_header / 1123eebe01ab / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- proxy_config.https.default_header

<a id="canonical-bf3f08497c044314c76f8983ad5ed14cc85b29a25b4f6b43631c6f2a78c76f14"></a>

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

<a id="canonical-419f5551fcfaafae91c771ed84fd327e0bad562a6a425ab32445a61f27daec1c"></a>

## Direct properties — proxy_config.https.default_header / 1123eebe01ab / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1d18e9bd9170cf713e1458a801e112cb81f82751cbb2c0be3db6c8b486e94134"></a>

## Next pages — proxy_config.https.default_header / 1123eebe01ab / 4

- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-5007f8764bf4b43e268bb25251f5abba2d21f1cff7fa04d7042a07e4cc8f1eca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f2fea20881dbbd405cb3a42ab4e076f542042920490b37c6e6879678e81f234a"></a>

## proxy_config.https.default_loadbalancer — proxy_config.https.default_loadbalancer / cc0ee2d38fc8 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- proxy_config.https.default_loadbalancer

<a id="canonical-358e63cdf62cb159386bd96955df9b390095f9579237eb2845c7fa226f53517a"></a>

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

<a id="canonical-e3009b78535bbcb9b25ef1e3b7095250cba6c961c6b304a6fa570c405740dda9"></a>

## Direct properties — proxy_config.https.default_loadbalancer / cc0ee2d38fc8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6be448a4dbe61174be5af5ddf5a6486b514b87f9e395e796f969b9b50d8a7ce0"></a>

## Next pages — proxy_config.https.default_loadbalancer / cc0ee2d38fc8 / 4

- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-49a22d9a7a36911612c707dc895cb4c8d615ae671e0a9ec91865f97b1e1f749d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c5fc3fcb752ec9caf2ba52e63d9a00fde31184a2067e3ec2b3abee3965d9b58f"></a>

## proxy_config.https.disable_path_normalize — proxy_config.https.disable_path_normalize / 6f458c64a50c / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- proxy_config.https.disable_path_normalize

<a id="canonical-ebb206f6f8d94deed75db42ebcc3754b2b04183787c284207cdb3a5d0fba4624"></a>

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

<a id="canonical-cd2450eb3e4a317bf9aa57365423f3609a01c9cf10657cb9eafde03165be1672"></a>

## Direct properties — proxy_config.https.disable_path_normalize / 6f458c64a50c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-251e9fd5fc2f495782e9673d2f7faeb02dad80dbca40268a9ceaa5cee60c1dcb"></a>

## Next pages — proxy_config.https.disable_path_normalize / 6f458c64a50c / 4

- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-629d86b3f1b2fbf2565410dde83f9fae730d3ab1eb3dd0ff6d8166f8706f14d8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1bc91690efdf52112b8a5d14ae0b1c4cfb085f2f8bc075a1988612203cd88bef"></a>

## proxy_config.https.enable_path_normalize — proxy_config.https.enable_path_normalize / d3d4e7c2e890 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- proxy_config.https.enable_path_normalize

<a id="canonical-c2ad5915f33a2d7f830ec3f2300d94ef7900c8c44abc4140519529526be9082a"></a>

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

<a id="canonical-679f6d25a37456bd07008358e2cc24da97bb762627420bf581007bd72b31a53e"></a>

## Direct properties — proxy_config.https.enable_path_normalize / d3d4e7c2e890 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-492fc9420ba93184007d6e5057b174a23b9d6e863e8283ab921a6201994cb348"></a>

## Next pages — proxy_config.https.enable_path_normalize / d3d4e7c2e890 / 4

- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-604ec47f9724ee98fbe6f227f219fa4eaac8145347b62c87cdbdeef5f0984317"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e3da8ecc925fdeeb9bea8c931a5854145a023c4e0cb230a88e9c212a48f704a5"></a>

## proxy_config.https.http_protocol_options — proxy_config.https.http_protocol_options / 3c01fe04d2c8 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- proxy_config.https.http_protocol_options

<a id="canonical-824bcf954b96ad918b0c9af7e2aa4add26fefe25847a4b0274203c9cd6c604f3"></a>

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

<a id="canonical-6e7e5b18ad46811d4b1b87394a804c05bd83cef2115e9aa99cd65266876b8363"></a>

## Direct properties — proxy_config.https.http_protocol_options / 3c01fe04d2c8 / 3

- [http_protocol_enable_v1_only](resources--bigip_http_proxy--reference--group-003.md#canonical-371c8d85a94dd3162833b618142f2e756ada0f5e617a6f3daea608ad6bbc97ae): complete subsection reference.

- [http_protocol_enable_v1_v2](resources--bigip_http_proxy--reference--group-003.md#canonical-2432b6deabc4d742066033ef4194e5b1b2c0c8350482c941b8fcbf6437db4259): complete subsection reference.

- [http_protocol_enable_v2_only](resources--bigip_http_proxy--reference--group-003.md#canonical-d50cfbd8b345481601791ecab456ec522476a2ef460bea32044c3edf7884f82d): complete subsection reference.

<a id="canonical-3c2c81fc2cd98e86ca3be11ff060ea7e7171e01f2dd70fbb2051b681c327c721"></a>

## Next pages — proxy_config.https.http_protocol_options / 3c01fe04d2c8 / 4

- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only](resources--bigip_http_proxy--reference--group-003.md#canonical-371c8d85a94dd3162833b618142f2e756ada0f5e617a6f3daea608ad6bbc97ae)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_v2](resources--bigip_http_proxy--reference--group-003.md#canonical-2432b6deabc4d742066033ef4194e5b1b2c0c8350482c941b8fcbf6437db4259)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v2_only](resources--bigip_http_proxy--reference--group-003.md#canonical-d50cfbd8b345481601791ecab456ec522476a2ef460bea32044c3edf7884f82d)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-371c8d85a94dd3162833b618142f2e756ada0f5e617a6f3daea608ad6bbc97ae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fa17edd7cac6a531b1eab11bf87b9620910423190f911d45762bb0add4536462"></a>

## proxy_config.https.http_protocol_options.http_protocol_enable_v1_only — proxy_config.https.http_protocol_options.http_protocol_enable_v1_only / 8f66a1b884cb / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- [proxy_config.https.http_protocol_options](resources--bigip_http_proxy--reference--group-003.md#canonical-604ec47f9724ee98fbe6f227f219fa4eaac8145347b62c87cdbdeef5f0984317)
- proxy_config.https.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-9c27a040f5225d7406699d05ac7f11128ab725b2a8c34d870d737abffc30d35f"></a>

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

<a id="canonical-c2909571c360cc1a8b9181ba0311e8c9ce0ffd6de8615456c88fe85781a8be22"></a>

## Direct properties — proxy_config.https.http_protocol_options.http_protocol_enable_v1_only / 8f66a1b884cb / 3

- [header_transformation](resources--bigip_http_proxy--reference--group-003.md#canonical-aef57451735671479f3a383f97b3a92892332170f60ddbbb847a39c7f26fd69c): complete subsection reference.

<a id="canonical-531044374a5fc8fae9051ef2ca87184c04c472505b9dda07c644225bcc039c95"></a>

## Next pages — proxy_config.https.http_protocol_options.http_protocol_enable_v1_only / 8f66a1b884cb / 4

- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--bigip_http_proxy--reference--group-003.md#canonical-aef57451735671479f3a383f97b3a92892332170f60ddbbb847a39c7f26fd69c)
- [proxy_config.https.http_protocol_options](resources--bigip_http_proxy--reference--group-003.md#canonical-604ec47f9724ee98fbe6f227f219fa4eaac8145347b62c87cdbdeef5f0984317)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-aef57451735671479f3a383f97b3a92892332170f60ddbbb847a39c7f26fd69c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b4053f7d845156ca276153050adb2cc7773ba139466c28567387bda61760531c"></a>

## proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation — proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_tra / be6074802f0b / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- [proxy_config.https.http_protocol_options](resources--bigip_http_proxy--reference--group-003.md#canonical-604ec47f9724ee98fbe6f227f219fa4eaac8145347b62c87cdbdeef5f0984317)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only](resources--bigip_http_proxy--reference--group-003.md#canonical-371c8d85a94dd3162833b618142f2e756ada0f5e617a6f3daea608ad6bbc97ae)
- proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-00f5e5b19490e694c36965a009c8dbe63ab9aa9a2d116ad761306fda3b3145ca"></a>

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

<a id="canonical-377ac49d2c3f1a3bdfbd832299ccea75318fe7591fabba7b49624a2252967dd8"></a>

## Direct properties — proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_tra / be6074802f0b / 3

- [default_header_transformation](resources--bigip_http_proxy--reference--group-003.md#canonical-cc2a91ad71fdca428f93453f547f222d35bfb8424ac8f7bc66ad53ca17a58b2d): complete subsection reference.

- [preserve_case_header_transformation](resources--bigip_http_proxy--reference--group-003.md#canonical-5759bb1e0c2ed0dd42d954e061d08f7afde0ba994bbd000af358946601745f03): complete subsection reference.

- [proper_case_header_transformation](resources--bigip_http_proxy--reference--group-003.md#canonical-6b638dc1f7822be48a72c613db7e52b7e6cd175d4f2d199357ecf8918d650473): complete subsection reference.

<a id="canonical-44f467587663f9d3960b413fef527fdfdfef94a1f405170fbf100478ca91eae1"></a>

## Next pages — proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_tra / be6074802f0b / 4

- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](resources--bigip_http_proxy--reference--group-003.md#canonical-cc2a91ad71fdca428f93453f547f222d35bfb8424ac8f7bc66ad53ca17a58b2d)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](resources--bigip_http_proxy--reference--group-003.md#canonical-5759bb1e0c2ed0dd42d954e061d08f7afde0ba994bbd000af358946601745f03)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](resources--bigip_http_proxy--reference--group-003.md#canonical-6b638dc1f7822be48a72c613db7e52b7e6cd175d4f2d199357ecf8918d650473)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only](resources--bigip_http_proxy--reference--group-003.md#canonical-371c8d85a94dd3162833b618142f2e756ada0f5e617a6f3daea608ad6bbc97ae)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-cc2a91ad71fdca428f93453f547f222d35bfb8424ac8f7bc66ad53ca17a58b2d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-29bf6c80d7aa13c772d324e5f56cfdd5aee7c8d5e8110869ffde66352287e92e"></a>

## proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation — proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_tra / ec53685ba18b / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- [proxy_config.https.http_protocol_options](resources--bigip_http_proxy--reference--group-003.md#canonical-604ec47f9724ee98fbe6f227f219fa4eaac8145347b62c87cdbdeef5f0984317)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only](resources--bigip_http_proxy--reference--group-003.md#canonical-371c8d85a94dd3162833b618142f2e756ada0f5e617a6f3daea608ad6bbc97ae)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--bigip_http_proxy--reference--group-003.md#canonical-aef57451735671479f3a383f97b3a92892332170f60ddbbb847a39c7f26fd69c)
- proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-c96a1ff50b5a3fd8703a608d205b987855f3b5fd867d2e81731609fcee3311ba"></a>

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

<a id="canonical-82e5ad7dcd8013646cdb045ec244da8788ed9007387530d0b4678f33b5fdb70a"></a>

## Direct properties — proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_tra / ec53685ba18b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-84f473f8498e26fd0ee6f08ea2b99ce347681ed9fa8967923526262bd12cffee"></a>

## Next pages — proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_tra / ec53685ba18b / 4

- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--bigip_http_proxy--reference--group-003.md#canonical-aef57451735671479f3a383f97b3a92892332170f60ddbbb847a39c7f26fd69c)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-5759bb1e0c2ed0dd42d954e061d08f7afde0ba994bbd000af358946601745f03"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7f179ab0cab7dc4ad3b587fb096f33900924edca155d8f4f9cbdfc90d5b7365f"></a>

## proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation — proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_tra / d3da58a6f0d7 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- [proxy_config.https.http_protocol_options](resources--bigip_http_proxy--reference--group-003.md#canonical-604ec47f9724ee98fbe6f227f219fa4eaac8145347b62c87cdbdeef5f0984317)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only](resources--bigip_http_proxy--reference--group-003.md#canonical-371c8d85a94dd3162833b618142f2e756ada0f5e617a6f3daea608ad6bbc97ae)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--bigip_http_proxy--reference--group-003.md#canonical-aef57451735671479f3a383f97b3a92892332170f60ddbbb847a39c7f26fd69c)
- proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-ff1c876852b8c36ed77296e3ab06229569462443d8ca686db332a24ed16657be"></a>

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

<a id="canonical-f6e73ca76c1673752b7fe3b480ba0fc804fc4422e1e4a1b95cea8d8ad4c3f351"></a>

## Direct properties — proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_tra / d3da58a6f0d7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-18032957fa8a7bd739271aaa0d62500a014c94a214f77cc8986b3c73953f8b3a"></a>

## Next pages — proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_tra / d3da58a6f0d7 / 4

- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--bigip_http_proxy--reference--group-003.md#canonical-aef57451735671479f3a383f97b3a92892332170f60ddbbb847a39c7f26fd69c)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-6b638dc1f7822be48a72c613db7e52b7e6cd175d4f2d199357ecf8918d650473"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2df9f52579d488c2148171d636d619dbfd54d7cb125eb0c7f49e82495da39b81"></a>

## proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation — proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_tra / 9b7cbae57f0e / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- [proxy_config.https.http_protocol_options](resources--bigip_http_proxy--reference--group-003.md#canonical-604ec47f9724ee98fbe6f227f219fa4eaac8145347b62c87cdbdeef5f0984317)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only](resources--bigip_http_proxy--reference--group-003.md#canonical-371c8d85a94dd3162833b618142f2e756ada0f5e617a6f3daea608ad6bbc97ae)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--bigip_http_proxy--reference--group-003.md#canonical-aef57451735671479f3a383f97b3a92892332170f60ddbbb847a39c7f26fd69c)
- proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-080bd63feeac5d47c6e62e47db22fbe113f2aba65c087135d6e9a1b8bd3ad922"></a>

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

<a id="canonical-07803969accccf7ab0216c18ea745f8beb910019452c3bb26bf4f2d2214fd256"></a>

## Direct properties — proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_tra / 9b7cbae57f0e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-10527c8c93c448f01ce9c4afd1133518335dbaf2314e12c0b2cf039dceb36540"></a>

## Next pages — proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_tra / 9b7cbae57f0e / 4

- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--bigip_http_proxy--reference--group-003.md#canonical-aef57451735671479f3a383f97b3a92892332170f60ddbbb847a39c7f26fd69c)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-2432b6deabc4d742066033ef4194e5b1b2c0c8350482c941b8fcbf6437db4259"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-77076b4b256f685470f80ced46d206b43beebe7e8cb76c893c34b0735c52bd86"></a>

## proxy_config.https.http_protocol_options.http_protocol_enable_v1_v2 — proxy_config.https.http_protocol_options.http_protocol_enable_v1_v2 / 8f4eb3004434 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- [proxy_config.https.http_protocol_options](resources--bigip_http_proxy--reference--group-003.md#canonical-604ec47f9724ee98fbe6f227f219fa4eaac8145347b62c87cdbdeef5f0984317)
- proxy_config.https.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-7a5e1e3908558626883b41a15505a2240bd9f29749aab03d17744bde520fad26"></a>

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

<a id="canonical-32ddc6a2c8667d8f3d197d5946d8bba6d74fae0b4309c3cd1baa4f0af912898a"></a>

## Direct properties — proxy_config.https.http_protocol_options.http_protocol_enable_v1_v2 / 8f4eb3004434 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a64709fbd0294f21b3c93c1d4ab611225d0598460c24852a991379c2ef59bb65"></a>

## Next pages — proxy_config.https.http_protocol_options.http_protocol_enable_v1_v2 / 8f4eb3004434 / 4

- [proxy_config.https.http_protocol_options](resources--bigip_http_proxy--reference--group-003.md#canonical-604ec47f9724ee98fbe6f227f219fa4eaac8145347b62c87cdbdeef5f0984317)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-d50cfbd8b345481601791ecab456ec522476a2ef460bea32044c3edf7884f82d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-281a380d4d001754f42bb21f57c576c57703b96a42307f634f1cfe474f491609"></a>

## proxy_config.https.http_protocol_options.http_protocol_enable_v2_only — proxy_config.https.http_protocol_options.http_protocol_enable_v2_only / 02dd087ca6ba / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- [proxy_config.https.http_protocol_options](resources--bigip_http_proxy--reference--group-003.md#canonical-604ec47f9724ee98fbe6f227f219fa4eaac8145347b62c87cdbdeef5f0984317)
- proxy_config.https.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-31469aaaf9db9f1d864fa0b93a13e2ccef31142c48de49b76042d2c4a598e70d"></a>

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

<a id="canonical-a81705e9fc16e28b5aabf3ef0626f1d7527eb35ca094e268c7a390113e578304"></a>

## Direct properties — proxy_config.https.http_protocol_options.http_protocol_enable_v2_only / 02dd087ca6ba / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-980baa430609610d4a0fea0f14d48372116ff27baaa4d793db713b390b114b44"></a>

## Next pages — proxy_config.https.http_protocol_options.http_protocol_enable_v2_only / 02dd087ca6ba / 4

- [proxy_config.https.http_protocol_options](resources--bigip_http_proxy--reference--group-003.md#canonical-604ec47f9724ee98fbe6f227f219fa4eaac8145347b62c87cdbdeef5f0984317)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-1a2189c142e8575942c881f77f36e23be9b4a082bc0437c9ffa4804fe8eff103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a0c9c431e5a5d200766537ea768c19aa2c461782dee4a5761777c2107e56e215"></a>

## proxy_config.https.non_default_loadbalancer — proxy_config.https.non_default_loadbalancer / 9734ed0a7c36 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- proxy_config.https.non_default_loadbalancer

<a id="canonical-d91791e83ae838d417e3f509e6f03d1c23682b0f099288cc1953ff13f021839a"></a>

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

<a id="canonical-81328f21ce208402ba4f9ebe6e7978c210bbf7d56d014590a0ddb4510eaf49ef"></a>

## Direct properties — proxy_config.https.non_default_loadbalancer / 9734ed0a7c36 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-21f4eed1d3fe32b8ff24eb9488eabd2ad793f0cf86252ba82dac046ccfafdced"></a>

## Next pages — proxy_config.https.non_default_loadbalancer / 9734ed0a7c36 / 4

- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-b4ab5c9199515937cf6af77ff8d62d8d5299643598ed6c99f7bd0892debaede4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7b149e83e0f73f411399a7aac91287d4d02724dad9404e1a9245c30d026cec7f"></a>

## proxy_config.https.pass_through — proxy_config.https.pass_through / a2314b80f7f2 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- proxy_config.https.pass_through

<a id="canonical-90cd7a4ee5dcdf12446444d1f275527772c139b2914c0d412bd3c090318590dd"></a>

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

<a id="canonical-c7cc77f51a0c43773e404cf31cad3ffbe4c0c38eb70c09c5fc7f892251f5f289"></a>

## Direct properties — proxy_config.https.pass_through / a2314b80f7f2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1f1f18773c6cff8203a2644f5072cf5a4b21ed88cbf6273e8a4cf5c35940d073"></a>

## Next pages — proxy_config.https.pass_through / a2314b80f7f2 / 4

- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-5bd61375e16f524597e4f2a07869259ce316df6e104e564e7919a5e5a7f76a12"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-11fd3e3dc57d1370e9c61ee403106dd27ae227cb3035ad069f34c8e6e37de792"></a>

## proxy_config.https.tls_cert_params — proxy_config.https.tls_cert_params / 4e8268b86b98 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- proxy_config.https.tls_cert_params

<a id="canonical-2ff80792778940a864793bb40570db1a909590583c99e208bd885c4414d38a3c"></a>

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

<a id="canonical-dbf0d6e0939fdc58442ecc3c6d2ec40ef7d6bd5eb0607c7fafda445eec2dec03"></a>

## Direct properties — proxy_config.https.tls_cert_params / 4e8268b86b98 / 3

- [certificates](resources--bigip_http_proxy--reference--group-003.md#canonical-47886e0f5470255383b21f82c3013168d3a217dfb4b7f7877c19ad950e85cac0): complete subsection reference.

- [no_mtls](resources--bigip_http_proxy--reference--group-003.md#canonical-cb3e6e0460e32c95f444ae7592d4540821e155a0fd5e795d8dab247f9f46f3ab): complete subsection reference.

- [tls_config](resources--bigip_http_proxy--reference--group-003.md#canonical-f6c55dd7229f01ea832113eef2f6d767f5cf620fe5dd26314f30bb9cbd91fcb8): complete subsection reference.

- [use_mtls](resources--bigip_http_proxy--reference--group-003.md#canonical-ab992cfa86ac9a98a0354633f9492528ef66db2e377b634e36ba8faac4b6a1f4): complete subsection reference.

<a id="canonical-57dd9191cfbdf2d51b257fc9ba04319ec21319cdda1def40bd17c6e679aed426"></a>

## Next pages — proxy_config.https.tls_cert_params / 4e8268b86b98 / 4

- [proxy_config.https.tls_cert_params.certificates](resources--bigip_http_proxy--reference--group-003.md#canonical-47886e0f5470255383b21f82c3013168d3a217dfb4b7f7877c19ad950e85cac0)
- [proxy_config.https.tls_cert_params.no_mtls](resources--bigip_http_proxy--reference--group-003.md#canonical-cb3e6e0460e32c95f444ae7592d4540821e155a0fd5e795d8dab247f9f46f3ab)
- [proxy_config.https.tls_cert_params.tls_config](resources--bigip_http_proxy--reference--group-003.md#canonical-f6c55dd7229f01ea832113eef2f6d767f5cf620fe5dd26314f30bb9cbd91fcb8)
- [proxy_config.https.tls_cert_params.use_mtls](resources--bigip_http_proxy--reference--group-003.md#canonical-ab992cfa86ac9a98a0354633f9492528ef66db2e377b634e36ba8faac4b6a1f4)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-47886e0f5470255383b21f82c3013168d3a217dfb4b7f7877c19ad950e85cac0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-65d19039207ccb2723c63b896dcd2875924828eb11d254b36bcd0d26bd5fb9d7"></a>

## proxy_config.https.tls_cert_params.certificates — proxy_config.https.tls_cert_params.certificates / 8d9f0ccc22cc / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- [proxy_config.https.tls_cert_params](resources--bigip_http_proxy--reference--group-003.md#canonical-5bd61375e16f524597e4f2a07869259ce316df6e104e564e7919a5e5a7f76a12)
- proxy_config.https.tls_cert_params.certificates

<a id="canonical-bd4999eed78144989cd4db9b0b74a387f4521b12d9f2e02a562e374cafd97805"></a>

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

<a id="canonical-2c6011b4ce5ec38f7898b75e4133da4a3fc4280a449638d3c67e039344509ea8"></a>

## Direct properties — proxy_config.https.tls_cert_params.certificates / 8d9f0ccc22cc / 3

<a id="canonical-c872796d1428e512135213bedc145496003aa8031b796c4a39df40d5910ac789"></a>

<a id="canonical-8bb9f3a8cb93116dd3cd56ff4692ce5c8462d666961b52a331e76645f6ac9add"></a>

## name property — proxy_config.https.tls_cert_params.certificates / 8d9f0ccc22cc / 4

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

<a id="canonical-a9f92a396f4cb8c05e6b2c7a0ebf42117ff79b7258ac403eee0d8558677fdc50"></a>

<a id="canonical-aef158cd1a179d626b2c11c11d72d8bafee26fbde7733cf71f462f550a99972b"></a>

## namespace property — proxy_config.https.tls_cert_params.certificates / 8d9f0ccc22cc / 5

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

<a id="canonical-74a247680179575538231b6ec156ddaf4338640e88b3f6e868c4f68a9d45bde0"></a>

<a id="canonical-edcc43eeb1e57b8321194f7d48b29896dfe233febae7cb775e4624fded17409c"></a>

## tenant property — proxy_config.https.tls_cert_params.certificates / 8d9f0ccc22cc / 6

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

<a id="canonical-fca5d2b450a1bfd2f1d63bbc62e7cf9d6a14451ac31e3a23ffc5a5a1ae72a358"></a>

## Next pages — proxy_config.https.tls_cert_params.certificates / 8d9f0ccc22cc / 7

- [proxy_config.https.tls_cert_params](resources--bigip_http_proxy--reference--group-003.md#canonical-5bd61375e16f524597e4f2a07869259ce316df6e104e564e7919a5e5a7f76a12)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-cb3e6e0460e32c95f444ae7592d4540821e155a0fd5e795d8dab247f9f46f3ab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-611d8f0236b75bea78846f61ca0b6016766215e23a84eb1c0432e5a3b81d934d"></a>

## proxy_config.https.tls_cert_params.no_mtls — proxy_config.https.tls_cert_params.no_mtls / 68285fddf50a / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- [proxy_config.https.tls_cert_params](resources--bigip_http_proxy--reference--group-003.md#canonical-5bd61375e16f524597e4f2a07869259ce316df6e104e564e7919a5e5a7f76a12)
- proxy_config.https.tls_cert_params.no_mtls

<a id="canonical-f9dd2a1fb59cc16f0651e7e8e3eecf0047de27be6e01448d017986ee242a6313"></a>

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

<a id="canonical-435d2d88286c894a152f282b15a914451269782038fad5f4a948e6c18e0adbc2"></a>

## Direct properties — proxy_config.https.tls_cert_params.no_mtls / 68285fddf50a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-63bc53336108013ad657febdebe39e6fe3cb7a03c34a58de86ce05d6a6f27dc7"></a>

## Next pages — proxy_config.https.tls_cert_params.no_mtls / 68285fddf50a / 4

- [proxy_config.https.tls_cert_params](resources--bigip_http_proxy--reference--group-003.md#canonical-5bd61375e16f524597e4f2a07869259ce316df6e104e564e7919a5e5a7f76a12)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-f6c55dd7229f01ea832113eef2f6d767f5cf620fe5dd26314f30bb9cbd91fcb8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aa6442f83070b1e6f6b2295923232a82e7782c56f34bccce4859a8147b1e0eae"></a>

## proxy_config.https.tls_cert_params.tls_config — proxy_config.https.tls_cert_params.tls_config / a3b92fb69f5e / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- [proxy_config.https.tls_cert_params](resources--bigip_http_proxy--reference--group-003.md#canonical-5bd61375e16f524597e4f2a07869259ce316df6e104e564e7919a5e5a7f76a12)
- proxy_config.https.tls_cert_params.tls_config

<a id="canonical-9fb2c7ab0788fab94ea7d8c76732185c46695fd51a235fe60cdc109bc5954993"></a>

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

<a id="canonical-106cf60924f0199c2079f6e0d9585e4e2087e24d49e52dfd1b1698a1ecaf40e9"></a>

## Direct properties — proxy_config.https.tls_cert_params.tls_config / a3b92fb69f5e / 3

- [custom_security](resources--bigip_http_proxy--reference--group-003.md#canonical-14aa890e3df7808e3c58cdaa5e618576d77627703d1ec5de969f8cdee8e3c3f4): complete subsection reference.

- [default_security](resources--bigip_http_proxy--reference--group-003.md#canonical-0ac5d56f7938ba85c0ab411923665bbb0d0b05894e3e28e02a1cb011d7deef1e): complete subsection reference.

- [low_security](resources--bigip_http_proxy--reference--group-003.md#canonical-78814a404d5ab7bfdf67cf6e2bd1aa85d3b518366772ccf092850bf7594beb8a): complete subsection reference.

- [medium_security](resources--bigip_http_proxy--reference--group-003.md#canonical-d936b8e5c2b915b82790504e6ece8d621632a1bfc3ea844ccaf242cdee95ce43): complete subsection reference.

<a id="canonical-03423b36e9cf2db890327e36e3111c28d14c60901e55fb53ee4fadffd4328151"></a>

## Next pages — proxy_config.https.tls_cert_params.tls_config / a3b92fb69f5e / 4

- [proxy_config.https.tls_cert_params.tls_config.custom_security](resources--bigip_http_proxy--reference--group-003.md#canonical-14aa890e3df7808e3c58cdaa5e618576d77627703d1ec5de969f8cdee8e3c3f4)
- [proxy_config.https.tls_cert_params.tls_config.default_security](resources--bigip_http_proxy--reference--group-003.md#canonical-0ac5d56f7938ba85c0ab411923665bbb0d0b05894e3e28e02a1cb011d7deef1e)
- [proxy_config.https.tls_cert_params.tls_config.low_security](resources--bigip_http_proxy--reference--group-003.md#canonical-78814a404d5ab7bfdf67cf6e2bd1aa85d3b518366772ccf092850bf7594beb8a)
- [proxy_config.https.tls_cert_params.tls_config.medium_security](resources--bigip_http_proxy--reference--group-003.md#canonical-d936b8e5c2b915b82790504e6ece8d621632a1bfc3ea844ccaf242cdee95ce43)
- [proxy_config.https.tls_cert_params](resources--bigip_http_proxy--reference--group-003.md#canonical-5bd61375e16f524597e4f2a07869259ce316df6e104e564e7919a5e5a7f76a12)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-14aa890e3df7808e3c58cdaa5e618576d77627703d1ec5de969f8cdee8e3c3f4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9c62217c55387b87009e3fd136e280d6b3f775149ce86de7a35cf8d8ebdf99e6"></a>

## proxy_config.https.tls_cert_params.tls_config.custom_security — proxy_config.https.tls_cert_params.tls_config.custom_security / e11c4a4d4ef3 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- [proxy_config.https.tls_cert_params](resources--bigip_http_proxy--reference--group-003.md#canonical-5bd61375e16f524597e4f2a07869259ce316df6e104e564e7919a5e5a7f76a12)
- [proxy_config.https.tls_cert_params.tls_config](resources--bigip_http_proxy--reference--group-003.md#canonical-f6c55dd7229f01ea832113eef2f6d767f5cf620fe5dd26314f30bb9cbd91fcb8)
- proxy_config.https.tls_cert_params.tls_config.custom_security

<a id="canonical-0dce79ec83f4d8288e8581aa705ec725f968025d5fa961bcbdc8415edab5a4f9"></a>

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

<a id="canonical-2d821e75d9aee17396e3b71b57cf2fd611ffde1fef38690f73f2aa8b41fd553d"></a>

## Direct properties — proxy_config.https.tls_cert_params.tls_config.custom_security / e11c4a4d4ef3 / 3

<a id="canonical-901aec354adb9f20385cd521bda246b07eaf1bcee8f76211c469d44a812e6210"></a>

<a id="canonical-9749afc66665a0c1790059333e1106f76f0277687216303958bad2e685c86e6f"></a>

## cipher_suites property — proxy_config.https.tls_cert_params.tls_config.custom_security / e11c4a4d4ef3 / 4

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

<a id="canonical-9d52aecd74dbe4ed038eb3b4254d61832cb4ce7efcb83c36e0c8120e0ea8051c"></a>

<a id="canonical-09f966e9f52b22cbb4614335e3b69a1b11282454df09068d269fead82f631c71"></a>

## max_version property — proxy_config.https.tls_cert_params.tls_config.custom_security / e11c4a4d4ef3 / 5

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

<a id="canonical-98aeeac38c86dac31d61f285ff7729c53affbfcf38e63343758ad0ae66aba485"></a>

<a id="canonical-2a2ccae99fe8879050fc991ef03bd3607b53d962084fd94d79905ee22aa1e737"></a>

## min_version property — proxy_config.https.tls_cert_params.tls_config.custom_security / e11c4a4d4ef3 / 6

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

<a id="canonical-90d5ac2f2d12ce7cdbfc5add1b4169c857caa1f93d0a66169292b8cb3f5cb624"></a>

## Next pages — proxy_config.https.tls_cert_params.tls_config.custom_security / e11c4a4d4ef3 / 7

- [proxy_config.https.tls_cert_params.tls_config](resources--bigip_http_proxy--reference--group-003.md#canonical-f6c55dd7229f01ea832113eef2f6d767f5cf620fe5dd26314f30bb9cbd91fcb8)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-0ac5d56f7938ba85c0ab411923665bbb0d0b05894e3e28e02a1cb011d7deef1e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f59c690c2763b0548f49016f1c492251773e8fc409127313249d42df5b8ad196"></a>

## proxy_config.https.tls_cert_params.tls_config.default_security — proxy_config.https.tls_cert_params.tls_config.default_security / 7a89eac486c5 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- [proxy_config.https.tls_cert_params](resources--bigip_http_proxy--reference--group-003.md#canonical-5bd61375e16f524597e4f2a07869259ce316df6e104e564e7919a5e5a7f76a12)
- [proxy_config.https.tls_cert_params.tls_config](resources--bigip_http_proxy--reference--group-003.md#canonical-f6c55dd7229f01ea832113eef2f6d767f5cf620fe5dd26314f30bb9cbd91fcb8)
- proxy_config.https.tls_cert_params.tls_config.default_security

<a id="canonical-c7f0d15f0d46ef80f81f8e648a599ec4a0fdb5bd76c4e45f26183cf50dc1a13b"></a>

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

<a id="canonical-3f6bfecca5745080f836c128d1b2a2ab59df68c9ba6280a4c55a377b3005e580"></a>

## Direct properties — proxy_config.https.tls_cert_params.tls_config.default_security / 7a89eac486c5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-04ccc856ca464a0dbffdddf799c174807c9dc327229be12f2d2c7bde622d14ce"></a>

## Next pages — proxy_config.https.tls_cert_params.tls_config.default_security / 7a89eac486c5 / 4

- [proxy_config.https.tls_cert_params.tls_config](resources--bigip_http_proxy--reference--group-003.md#canonical-f6c55dd7229f01ea832113eef2f6d767f5cf620fe5dd26314f30bb9cbd91fcb8)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-78814a404d5ab7bfdf67cf6e2bd1aa85d3b518366772ccf092850bf7594beb8a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8bf0f30c0b26dd0c81a0bb24ff8712456877c7816536ba1fea2e0f84d1a61c9d"></a>

## proxy_config.https.tls_cert_params.tls_config.low_security — proxy_config.https.tls_cert_params.tls_config.low_security / 3f9bcfa41044 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- [proxy_config.https.tls_cert_params](resources--bigip_http_proxy--reference--group-003.md#canonical-5bd61375e16f524597e4f2a07869259ce316df6e104e564e7919a5e5a7f76a12)
- [proxy_config.https.tls_cert_params.tls_config](resources--bigip_http_proxy--reference--group-003.md#canonical-f6c55dd7229f01ea832113eef2f6d767f5cf620fe5dd26314f30bb9cbd91fcb8)
- proxy_config.https.tls_cert_params.tls_config.low_security

<a id="canonical-84676e3211481fe9133b502c29b3edb00f5537ee9b6349d982b7627be8b2c3d3"></a>

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

<a id="canonical-d0d90dac3365c0200fa1ff1bef21e90f5f33ebc1843f3e8c5a767052ca937874"></a>

## Direct properties — proxy_config.https.tls_cert_params.tls_config.low_security / 3f9bcfa41044 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cb32664d7235d228f6ca97beabcdd8648d4a824104d98a5b9f40e764dbed13d4"></a>

## Next pages — proxy_config.https.tls_cert_params.tls_config.low_security / 3f9bcfa41044 / 4

- [proxy_config.https.tls_cert_params.tls_config](resources--bigip_http_proxy--reference--group-003.md#canonical-f6c55dd7229f01ea832113eef2f6d767f5cf620fe5dd26314f30bb9cbd91fcb8)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-d936b8e5c2b915b82790504e6ece8d621632a1bfc3ea844ccaf242cdee95ce43"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3e5e81fc2bbbcacbc218f1baf926e6a20d28f9ed4b375f0a550d2cc9e391864f"></a>

## proxy_config.https.tls_cert_params.tls_config.medium_security — proxy_config.https.tls_cert_params.tls_config.medium_security / 0192da04fb20 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- [proxy_config.https.tls_cert_params](resources--bigip_http_proxy--reference--group-003.md#canonical-5bd61375e16f524597e4f2a07869259ce316df6e104e564e7919a5e5a7f76a12)
- [proxy_config.https.tls_cert_params.tls_config](resources--bigip_http_proxy--reference--group-003.md#canonical-f6c55dd7229f01ea832113eef2f6d767f5cf620fe5dd26314f30bb9cbd91fcb8)
- proxy_config.https.tls_cert_params.tls_config.medium_security

<a id="canonical-baf50061acafe10e948a9aa4e8104b4e212d499bed258cd65e1a42626d0e201c"></a>

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

<a id="canonical-b8263a9910e08b30905b37f4017ff1223a5b03c5d3e529619dd1f6cf266dbe9c"></a>

## Direct properties — proxy_config.https.tls_cert_params.tls_config.medium_security / 0192da04fb20 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3b2f685c5871123c28bec7b2628e66679268a08cb28b0e7a2888fdf532649255"></a>

## Next pages — proxy_config.https.tls_cert_params.tls_config.medium_security / 0192da04fb20 / 4

- [proxy_config.https.tls_cert_params.tls_config](resources--bigip_http_proxy--reference--group-003.md#canonical-f6c55dd7229f01ea832113eef2f6d767f5cf620fe5dd26314f30bb9cbd91fcb8)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-ab992cfa86ac9a98a0354633f9492528ef66db2e377b634e36ba8faac4b6a1f4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-79f055d99578069f681dbceccb59f278e64b44d4025d5997219c3ac9ea085583"></a>

## proxy_config.https.tls_cert_params.use_mtls — proxy_config.https.tls_cert_params.use_mtls / 0f7c85ad7c54 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- [proxy_config.https.tls_cert_params](resources--bigip_http_proxy--reference--group-003.md#canonical-5bd61375e16f524597e4f2a07869259ce316df6e104e564e7919a5e5a7f76a12)
- proxy_config.https.tls_cert_params.use_mtls

<a id="canonical-82cf13559eda168fe8119e873f36047a234e1f8fcc024866f88db78d68a66b01"></a>

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

<a id="canonical-5f1407a41629ad5f49bf02d53e4076a651cb0ceaea1a33d93c36864d38a320f1"></a>

## Direct properties — proxy_config.https.tls_cert_params.use_mtls / 0f7c85ad7c54 / 3

<a id="canonical-c7ca7a8fcb8bbfa08375a628cd84a87f6dfafc1f74e21af6363413a773cff1df"></a>

<a id="canonical-13a894c35bdff5012f411780379215c4b14ae4ae4f0cc1f3fc424a98afca12b9"></a>

## client_certificate_optional property — proxy_config.https.tls_cert_params.use_mtls / 0f7c85ad7c54 / 4

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

- [crl](resources--bigip_http_proxy--reference--group-003.md#canonical-814e0492dbda15472fa56e0d94d48d2ad3eb9c8270e2c8aec3c7e10fce504748): complete subsection reference.

- [no_crl](resources--bigip_http_proxy--reference--group-003.md#canonical-b3f9002228d311a018f359e1d63eae37806622d44ccdffb2eac474c912b9bc30): complete subsection reference.

- [trusted_ca](resources--bigip_http_proxy--reference--group-003.md#canonical-b27c060f02fe2c7313698d492bc0e06ac435622841040be1b9b23057df23fd17): complete subsection reference.

<a id="canonical-03a2db73dc3fecfefad3906ec81f92f423033822f2da0854a4cb13361acf2a2c"></a>

<a id="canonical-33e9401182e05ad726d4398164d27fe2a22783235816e834e32347c5db6d447b"></a>

## trusted_ca_url property — proxy_config.https.tls_cert_params.use_mtls / 0f7c85ad7c54 / 5

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

- [xfcc_disabled](resources--bigip_http_proxy--reference--group-003.md#canonical-e54d1b43cac002402ed4c4dd1e62daae8ad0b4795dec551194e3999c67d5f5ac): complete subsection reference.

- [xfcc_options](resources--bigip_http_proxy--reference--group-003.md#canonical-5848f4611b1b29e0d8c3bb03956e36b43f0f8a78d9e11eb91b0aeede25d215e9): complete subsection reference.

<a id="canonical-005077cfd1fb8630cff22b598d3639c1a85cf2f1e432a384d0cfe5329d05a888"></a>

## Next pages — proxy_config.https.tls_cert_params.use_mtls / 0f7c85ad7c54 / 6

- [proxy_config.https.tls_cert_params.use_mtls.crl](resources--bigip_http_proxy--reference--group-003.md#canonical-814e0492dbda15472fa56e0d94d48d2ad3eb9c8270e2c8aec3c7e10fce504748)
- [proxy_config.https.tls_cert_params.use_mtls.no_crl](resources--bigip_http_proxy--reference--group-003.md#canonical-b3f9002228d311a018f359e1d63eae37806622d44ccdffb2eac474c912b9bc30)
- [proxy_config.https.tls_cert_params.use_mtls.trusted_ca](resources--bigip_http_proxy--reference--group-003.md#canonical-b27c060f02fe2c7313698d492bc0e06ac435622841040be1b9b23057df23fd17)
- [proxy_config.https.tls_cert_params.use_mtls.xfcc_disabled](resources--bigip_http_proxy--reference--group-003.md#canonical-e54d1b43cac002402ed4c4dd1e62daae8ad0b4795dec551194e3999c67d5f5ac)
- [proxy_config.https.tls_cert_params.use_mtls.xfcc_options](resources--bigip_http_proxy--reference--group-003.md#canonical-5848f4611b1b29e0d8c3bb03956e36b43f0f8a78d9e11eb91b0aeede25d215e9)
- [proxy_config.https.tls_cert_params](resources--bigip_http_proxy--reference--group-003.md#canonical-5bd61375e16f524597e4f2a07869259ce316df6e104e564e7919a5e5a7f76a12)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-814e0492dbda15472fa56e0d94d48d2ad3eb9c8270e2c8aec3c7e10fce504748"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-be38f18905c17180c47cb17a84de7c9a18707559bd2a501c789f36e1b3892d4c"></a>

## proxy_config.https.tls_cert_params.use_mtls.crl — proxy_config.https.tls_cert_params.use_mtls.crl / f40e58e43596 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- [proxy_config.https.tls_cert_params](resources--bigip_http_proxy--reference--group-003.md#canonical-5bd61375e16f524597e4f2a07869259ce316df6e104e564e7919a5e5a7f76a12)
- [proxy_config.https.tls_cert_params.use_mtls](resources--bigip_http_proxy--reference--group-003.md#canonical-ab992cfa86ac9a98a0354633f9492528ef66db2e377b634e36ba8faac4b6a1f4)
- proxy_config.https.tls_cert_params.use_mtls.crl

<a id="canonical-502b24de89ca1e1b1a0068a6e97a427a6284d44c7212ccae3b837a9ee016b6ef"></a>

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

<a id="canonical-a3d040d80b8d307ed4d7ccaf962ef38d2c5f19c13d45bda23e9b9804c22d955c"></a>

## Direct properties — proxy_config.https.tls_cert_params.use_mtls.crl / f40e58e43596 / 3

<a id="canonical-df4943326cd7324c37dc0d6cc7c1c36900ae2a7f741b239bfe89f366cada7436"></a>

<a id="canonical-222ccc017615e2421d2a9b93d092c59d52f50c45eb317e8aaab5551f1f8a74bb"></a>

## name property — proxy_config.https.tls_cert_params.use_mtls.crl / f40e58e43596 / 4

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

<a id="canonical-393a408547a811624866cd08b5f4c024fe9d3da658cf0fc6a7363312b6bbcd5c"></a>

<a id="canonical-16ac55d85a435c94a8f1a04b2ae0113bca942717b0e86aa3d6b830fe1389eff4"></a>

## namespace property — proxy_config.https.tls_cert_params.use_mtls.crl / f40e58e43596 / 5

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

<a id="canonical-b73b1d254c29abe7e88b4eb2701bf6fef9a94b20a6952a9b451c2201bf59d932"></a>

<a id="canonical-2d9c04f355c7df090d4f0bfaba23a4993d4de232d96a711a6ff15aeb65bb6322"></a>

## tenant property — proxy_config.https.tls_cert_params.use_mtls.crl / f40e58e43596 / 6

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

<a id="canonical-cbfb70550db07d5c671318604c0c6895eabb12908f7f6e4a14592302700688de"></a>

## Next pages — proxy_config.https.tls_cert_params.use_mtls.crl / f40e58e43596 / 7

- [proxy_config.https.tls_cert_params.use_mtls](resources--bigip_http_proxy--reference--group-003.md#canonical-ab992cfa86ac9a98a0354633f9492528ef66db2e377b634e36ba8faac4b6a1f4)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-b3f9002228d311a018f359e1d63eae37806622d44ccdffb2eac474c912b9bc30"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4b362793333f10867385025a6fa44528d3a7d6785c00779caf656a3f6a4517b5"></a>

## proxy_config.https.tls_cert_params.use_mtls.no_crl — proxy_config.https.tls_cert_params.use_mtls.no_crl / faf920466ef1 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- [proxy_config.https.tls_cert_params](resources--bigip_http_proxy--reference--group-003.md#canonical-5bd61375e16f524597e4f2a07869259ce316df6e104e564e7919a5e5a7f76a12)
- [proxy_config.https.tls_cert_params.use_mtls](resources--bigip_http_proxy--reference--group-003.md#canonical-ab992cfa86ac9a98a0354633f9492528ef66db2e377b634e36ba8faac4b6a1f4)
- proxy_config.https.tls_cert_params.use_mtls.no_crl

<a id="canonical-c81784c97c6f7c0c0621f7143e3ba2c13ec2229fe146c2732588c97b6da28729"></a>

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

<a id="canonical-2734a4610219db4fa07c8d57196d984bcd654d65e349a9882ad28003b85a0d46"></a>

## Direct properties — proxy_config.https.tls_cert_params.use_mtls.no_crl / faf920466ef1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9f26ea0582e5ab2f74ffe3c770c6cde83e93a9f3970050878e8b845b6957fdb8"></a>

## Next pages — proxy_config.https.tls_cert_params.use_mtls.no_crl / faf920466ef1 / 4

- [proxy_config.https.tls_cert_params.use_mtls](resources--bigip_http_proxy--reference--group-003.md#canonical-ab992cfa86ac9a98a0354633f9492528ef66db2e377b634e36ba8faac4b6a1f4)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-b27c060f02fe2c7313698d492bc0e06ac435622841040be1b9b23057df23fd17"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-84ed99f1916a5337a71d9a29513f3fe8701bba65dba8d1b7317414e8e05caad7"></a>

## proxy_config.https.tls_cert_params.use_mtls.trusted_ca — proxy_config.https.tls_cert_params.use_mtls.trusted_ca / ea32536fcd13 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- [proxy_config.https.tls_cert_params](resources--bigip_http_proxy--reference--group-003.md#canonical-5bd61375e16f524597e4f2a07869259ce316df6e104e564e7919a5e5a7f76a12)
- [proxy_config.https.tls_cert_params.use_mtls](resources--bigip_http_proxy--reference--group-003.md#canonical-ab992cfa86ac9a98a0354633f9492528ef66db2e377b634e36ba8faac4b6a1f4)
- proxy_config.https.tls_cert_params.use_mtls.trusted_ca

<a id="canonical-80b9e158851e1e884d68dcbf7226316688a4edb27d02313740a8556113bfedae"></a>

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

<a id="canonical-7a46e453a146fab8f0022a242642e49b992e29ab13cf79c9550efb98768f2be3"></a>

## Direct properties — proxy_config.https.tls_cert_params.use_mtls.trusted_ca / ea32536fcd13 / 3

<a id="canonical-438a70b70d09cad647641536e0d48ef4d22160f992296d0d9efb457921bcdcb9"></a>

<a id="canonical-e9b961dbf3629a103e8d89395530cdf7f09233a6730631cd53497bc508d75eae"></a>

## name property — proxy_config.https.tls_cert_params.use_mtls.trusted_ca / ea32536fcd13 / 4

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

<a id="canonical-7811c4df08ae8ccdedee402ccbbfd76d04aee57c9a507af4ab3b890bb2ac8118"></a>

<a id="canonical-8fa3e5c9e164c3b92ebddacf729271828226a38568166a3b31adb70efee427ab"></a>

## namespace property — proxy_config.https.tls_cert_params.use_mtls.trusted_ca / ea32536fcd13 / 5

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

<a id="canonical-9af38733057be6492361f13851c2261252f049bd863e76508e346e4bde740cf1"></a>

<a id="canonical-e9d6ec2027d97c6c4d7e822118b836e2b3cd553a712976203178b52a27d5988b"></a>

## tenant property — proxy_config.https.tls_cert_params.use_mtls.trusted_ca / ea32536fcd13 / 6

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

<a id="canonical-882702a07c709e5733354f0e1fa96d957e1ece002ddc000df87ec9f7fa506d0f"></a>

## Next pages — proxy_config.https.tls_cert_params.use_mtls.trusted_ca / ea32536fcd13 / 7

- [proxy_config.https.tls_cert_params.use_mtls](resources--bigip_http_proxy--reference--group-003.md#canonical-ab992cfa86ac9a98a0354633f9492528ef66db2e377b634e36ba8faac4b6a1f4)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-e54d1b43cac002402ed4c4dd1e62daae8ad0b4795dec551194e3999c67d5f5ac"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9ab3d96d772b9624a736dab33ba8718df07961fdf64bc516ed8bb2c010e8bd40"></a>

## proxy_config.https.tls_cert_params.use_mtls.xfcc_disabled — proxy_config.https.tls_cert_params.use_mtls.xfcc_disabled / 1628172c78d3 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- [proxy_config.https.tls_cert_params](resources--bigip_http_proxy--reference--group-003.md#canonical-5bd61375e16f524597e4f2a07869259ce316df6e104e564e7919a5e5a7f76a12)
- [proxy_config.https.tls_cert_params.use_mtls](resources--bigip_http_proxy--reference--group-003.md#canonical-ab992cfa86ac9a98a0354633f9492528ef66db2e377b634e36ba8faac4b6a1f4)
- proxy_config.https.tls_cert_params.use_mtls.xfcc_disabled

<a id="canonical-153f22e9f40f3cdd74f99dd54735075e1c10f55a57a98ee4e0c076fae3ef79dd"></a>

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

<a id="canonical-0a8f21131241398d59644fee684e8a63c441467bbace3eb99a63bb01149f4933"></a>

## Direct properties — proxy_config.https.tls_cert_params.use_mtls.xfcc_disabled / 1628172c78d3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e7816707f69a8c95d31f8a47bc3abf0ae126a0dc39a9fec0bc1d9430a220440b"></a>

## Next pages — proxy_config.https.tls_cert_params.use_mtls.xfcc_disabled / 1628172c78d3 / 4

- [proxy_config.https.tls_cert_params.use_mtls](resources--bigip_http_proxy--reference--group-003.md#canonical-ab992cfa86ac9a98a0354633f9492528ef66db2e377b634e36ba8faac4b6a1f4)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-5848f4611b1b29e0d8c3bb03956e36b43f0f8a78d9e11eb91b0aeede25d215e9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cda70932f8c35d8280b804da4d2516c2527e0979e9f011cc80883d7c25b34522"></a>

## proxy_config.https.tls_cert_params.use_mtls.xfcc_options — proxy_config.https.tls_cert_params.use_mtls.xfcc_options / 86c0c5bc122b / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- [proxy_config.https.tls_cert_params](resources--bigip_http_proxy--reference--group-003.md#canonical-5bd61375e16f524597e4f2a07869259ce316df6e104e564e7919a5e5a7f76a12)
- [proxy_config.https.tls_cert_params.use_mtls](resources--bigip_http_proxy--reference--group-003.md#canonical-ab992cfa86ac9a98a0354633f9492528ef66db2e377b634e36ba8faac4b6a1f4)
- proxy_config.https.tls_cert_params.use_mtls.xfcc_options

<a id="canonical-3e9acc1797065511707f9e5ee32f7819d1b60f5587645633b55db80f43abae58"></a>

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

<a id="canonical-d47ca4a403cec20740639814699242c77fce6e53653985413fcac92aabaa50ce"></a>

## Direct properties — proxy_config.https.tls_cert_params.use_mtls.xfcc_options / 86c0c5bc122b / 3

<a id="canonical-a6bc1484321bac4a7d584c15120471adb1dc5d46c6840dec1bc6721781710f72"></a>

<a id="canonical-35526d7f8e5d916253b82adab38802f32e6b3a9831fecc1a90eb3e8d6704eef5"></a>

## xfcc_header_elements property — proxy_config.https.tls_cert_params.use_mtls.xfcc_options / 86c0c5bc122b / 4

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

<a id="canonical-434eb83d374f3f375601049dcb65183476f38cd529b23d3c5bd07c3215fb1129"></a>

## Next pages — proxy_config.https.tls_cert_params.use_mtls.xfcc_options / 86c0c5bc122b / 5

- [proxy_config.https.tls_cert_params.use_mtls](resources--bigip_http_proxy--reference--group-003.md#canonical-ab992cfa86ac9a98a0354633f9492528ef66db2e377b634e36ba8faac4b6a1f4)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-1a2fff832c914deeb21855b2b6d870103e691b9c714316eef87fcff21a04a187"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6c3b3827404f6425f235ad480062ad2b44b249eaee6f06432076b95f5b984869"></a>

## proxy_config.https.tls_parameters — proxy_config.https.tls_parameters / bc9131b2e67e / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- proxy_config.https.tls_parameters

<a id="canonical-18498408501693d229a28f817fd49a988b6de53602355e564ab4a852dd2acfdf"></a>

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

<a id="canonical-67e59cb083b3c2b889c1418c7143dee50869f4315e57f03a2c01f5e677d0887c"></a>

## Direct properties — proxy_config.https.tls_parameters / bc9131b2e67e / 3

- [no_mtls](resources--bigip_http_proxy--reference--group-003.md#canonical-ac0c3eb98c8cb31c354cfb15b66431ca1a5acf2bef46ebbee07fb03c2d26b299): complete subsection reference.

- [tls_certificates](resources--bigip_http_proxy--reference--group-003.md#canonical-17e23908f5dd270a57d0f0d511aeb709933f657b78ee076266687a44b9d80651): complete subsection reference.

- [tls_config](resources--bigip_http_proxy--reference--group-003.md#canonical-8d3d19560fe789477a691a271848845c6d01c63fd3ae309a100ca27dbb7ba220): complete subsection reference.

- [use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-7c66469fcd6b8d3805bd3f8b1b7b58d8a3f9725abb801b3be3c1a43c83ba8830): complete subsection reference.

<a id="canonical-b758fcdbaaea0df527a17b26ee26a9c6b0fd28f9f08680dd2312898ba7470482"></a>

## Next pages — proxy_config.https.tls_parameters / bc9131b2e67e / 4

- [proxy_config.https.tls_parameters.no_mtls](resources--bigip_http_proxy--reference--group-003.md#canonical-ac0c3eb98c8cb31c354cfb15b66431ca1a5acf2bef46ebbee07fb03c2d26b299)
- [proxy_config.https.tls_parameters.tls_certificates](resources--bigip_http_proxy--reference--group-003.md#canonical-17e23908f5dd270a57d0f0d511aeb709933f657b78ee076266687a44b9d80651)
- [proxy_config.https.tls_parameters.tls_config](resources--bigip_http_proxy--reference--group-003.md#canonical-8d3d19560fe789477a691a271848845c6d01c63fd3ae309a100ca27dbb7ba220)
- [proxy_config.https.tls_parameters.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-7c66469fcd6b8d3805bd3f8b1b7b58d8a3f9725abb801b3be3c1a43c83ba8830)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-ac0c3eb98c8cb31c354cfb15b66431ca1a5acf2bef46ebbee07fb03c2d26b299"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-85d4529e1d601c6832b7f9ca2bb1714993a05c055e3b553de52422d64ec4985c"></a>

## proxy_config.https.tls_parameters.no_mtls — proxy_config.https.tls_parameters.no_mtls / ece1d4559514 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-003.md#canonical-1a2fff832c914deeb21855b2b6d870103e691b9c714316eef87fcff21a04a187)
- proxy_config.https.tls_parameters.no_mtls

<a id="canonical-f4303b0268f2e4a4f4ebaabad8a2d7538740a04d2ea2de06ea8195d2ad14904a"></a>

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

<a id="canonical-876acb5d958a758d2556db087f3ccf298093ac1d33a3208fef4c562e77ccff49"></a>

## Direct properties — proxy_config.https.tls_parameters.no_mtls / ece1d4559514 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-83038d552b7790429fabd4c173f4d60faf098c6dd8ff2f88a4fa7b6d283932be"></a>

## Next pages — proxy_config.https.tls_parameters.no_mtls / ece1d4559514 / 4

- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-003.md#canonical-1a2fff832c914deeb21855b2b6d870103e691b9c714316eef87fcff21a04a187)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-17e23908f5dd270a57d0f0d511aeb709933f657b78ee076266687a44b9d80651"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e7d2aa76abb9fd0c261c11286ceb0eabbf39941db4590161634100e5faade364"></a>

## proxy_config.https.tls_parameters.tls_certificates — proxy_config.https.tls_parameters.tls_certificates / 9500eb0af4f8 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-003.md#canonical-1a2fff832c914deeb21855b2b6d870103e691b9c714316eef87fcff21a04a187)
- proxy_config.https.tls_parameters.tls_certificates

<a id="canonical-28c4ef3f7b1d0f50063efed6b8fb5bdbe6928a2a162502e4e1432dfe1a4b51f4"></a>

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

<a id="canonical-0a39e82f5aa22165021495d79782a0f59df494483bc205cf8b96e3b98d01723c"></a>

## Direct properties — proxy_config.https.tls_parameters.tls_certificates / 9500eb0af4f8 / 3

<a id="canonical-0e9d063d2d26348ccbea12fd56101ea7deb8481038aa0e9ee5e277fb07edcff1"></a>

<a id="canonical-b30931c57f23f12a19849328d9684019f7007d4c88627cec43dd8af2be30379b"></a>

## certificate_url property — proxy_config.https.tls_parameters.tls_certificates / 9500eb0af4f8 / 4

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

- [custom_hash_algorithms](resources--bigip_http_proxy--reference--group-003.md#canonical-834406f1c3e610d00e681cad4aeb8918a745b4ed83495e92ebc133be95c009c7): complete subsection reference.

<a id="canonical-dfd3d17b361c5e5cdffbacd6ae84eb201e2312c6079de8e2aead80b4e9f1b9ce"></a>

<a id="canonical-5a5fe017b16b48a8810c40bd405385ce817039eff0b0914b0fffe97b3328b2e9"></a>

## description_spec property — proxy_config.https.tls_parameters.tls_certificates / 9500eb0af4f8 / 5

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--bigip_http_proxy--reference--group-003.md#canonical-d4b89a03c3bd1ddba395ccf58e610d80781c9cf22387440e29817c74501cfbcc): complete subsection reference.

- [private_key](resources--bigip_http_proxy--reference--group-003.md#canonical-651b77361281bd598c894e0035a49cb75b2b84a8ec1791e1ef3902a8d2f2e2b2): complete subsection reference.

- [use_system_defaults](resources--bigip_http_proxy--reference--group-003.md#canonical-4e0eb2a89cefeddce9ec0eb3d4d2c4a524fc7cdb34c59b1cd6adc4d73eaa807c): complete subsection reference.

<a id="canonical-464bda16a3a20de8f0c7428fa173e7c53c7b3d7e19fd609ae60bd4269a7b3b49"></a>

## Next pages — proxy_config.https.tls_parameters.tls_certificates / 9500eb0af4f8 / 6

- [proxy_config.https.tls_parameters.tls_certificates.custom_hash_algorithms](resources--bigip_http_proxy--reference--group-003.md#canonical-834406f1c3e610d00e681cad4aeb8918a745b4ed83495e92ebc133be95c009c7)
- [proxy_config.https.tls_parameters.tls_certificates.disable_ocsp_stapling](resources--bigip_http_proxy--reference--group-003.md#canonical-d4b89a03c3bd1ddba395ccf58e610d80781c9cf22387440e29817c74501cfbcc)
- [proxy_config.https.tls_parameters.tls_certificates.private_key](resources--bigip_http_proxy--reference--group-003.md#canonical-651b77361281bd598c894e0035a49cb75b2b84a8ec1791e1ef3902a8d2f2e2b2)
- [proxy_config.https.tls_parameters.tls_certificates.use_system_defaults](resources--bigip_http_proxy--reference--group-003.md#canonical-4e0eb2a89cefeddce9ec0eb3d4d2c4a524fc7cdb34c59b1cd6adc4d73eaa807c)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-003.md#canonical-1a2fff832c914deeb21855b2b6d870103e691b9c714316eef87fcff21a04a187)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-834406f1c3e610d00e681cad4aeb8918a745b4ed83495e92ebc133be95c009c7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2757a33f350eba3a0bee63603ccf7acd43f1c975ab8341b590fcacb094676964"></a>

## proxy_config.https.tls_parameters.tls_certificates.custom_hash_algorithms — proxy_config.https.tls_parameters.tls_certificates.custom_hash_algorithms / 10da06294aa9 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-003.md#canonical-1a2fff832c914deeb21855b2b6d870103e691b9c714316eef87fcff21a04a187)
- [proxy_config.https.tls_parameters.tls_certificates](resources--bigip_http_proxy--reference--group-003.md#canonical-17e23908f5dd270a57d0f0d511aeb709933f657b78ee076266687a44b9d80651)
- proxy_config.https.tls_parameters.tls_certificates.custom_hash_algorithms

<a id="canonical-bfb937f013fbe490efc3c2371d149f1b4b5a4f34f92db2e49e59b9a3f7a6f48c"></a>

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

<a id="canonical-1cf53d1ecf9fcd403ab49e53aee47f6c48d3e77bc70d902345506850471f29c8"></a>

## Direct properties — proxy_config.https.tls_parameters.tls_certificates.custom_hash_algorithms / 10da06294aa9 / 3

<a id="canonical-50a83b298ff5b34ce82487dbc6699367c2adadb6451ad4229d16b238bdcead6d"></a>

<a id="canonical-e2533569f2f96e077c39abea76ebc3a9d9ddd1f67813d34a385cda8620f180ed"></a>

## hash_algorithms property — proxy_config.https.tls_parameters.tls_certificates.custom_hash_algorithms / 10da06294aa9 / 4

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

<a id="canonical-f1f7bcd65c8295da84e58ab78f2d8a7ec8edfa1fb761da9ee4a40c31f480aa73"></a>

## Next pages — proxy_config.https.tls_parameters.tls_certificates.custom_hash_algorithms / 10da06294aa9 / 5

- [proxy_config.https.tls_parameters.tls_certificates](resources--bigip_http_proxy--reference--group-003.md#canonical-17e23908f5dd270a57d0f0d511aeb709933f657b78ee076266687a44b9d80651)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-d4b89a03c3bd1ddba395ccf58e610d80781c9cf22387440e29817c74501cfbcc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-08aaaabcdc0355ffbb2b224d2356659caa33cd95262c6b42d17da00d7a6f6e48"></a>

## proxy_config.https.tls_parameters.tls_certificates.disable_ocsp_stapling — proxy_config.https.tls_parameters.tls_certificates.disable_ocsp_stapling / e855fd4823ce / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-003.md#canonical-1a2fff832c914deeb21855b2b6d870103e691b9c714316eef87fcff21a04a187)
- [proxy_config.https.tls_parameters.tls_certificates](resources--bigip_http_proxy--reference--group-003.md#canonical-17e23908f5dd270a57d0f0d511aeb709933f657b78ee076266687a44b9d80651)
- proxy_config.https.tls_parameters.tls_certificates.disable_ocsp_stapling

<a id="canonical-ef2e62d7f08ed175aeba1ad3b5de280c77dd2f27d6f516783c3e4b303677a4a3"></a>

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

<a id="canonical-a419315874ba5c20282b7f5d51d27dff059c458dd9a7563611a947acc419329d"></a>

## Direct properties — proxy_config.https.tls_parameters.tls_certificates.disable_ocsp_stapling / e855fd4823ce / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-234eb00e1b7beb6982f28d11243943084c8c846cba7788559be9d5356adc9bc1"></a>

## Next pages — proxy_config.https.tls_parameters.tls_certificates.disable_ocsp_stapling / e855fd4823ce / 4

- [proxy_config.https.tls_parameters.tls_certificates](resources--bigip_http_proxy--reference--group-003.md#canonical-17e23908f5dd270a57d0f0d511aeb709933f657b78ee076266687a44b9d80651)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-651b77361281bd598c894e0035a49cb75b2b84a8ec1791e1ef3902a8d2f2e2b2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d9f01e8315b58ffa2172cebc61139554effee51fc0545e0bdd82ab094e200b3b"></a>

## proxy_config.https.tls_parameters.tls_certificates.private_key — proxy_config.https.tls_parameters.tls_certificates.private_key / 4590c7f19a10 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-003.md#canonical-1a2fff832c914deeb21855b2b6d870103e691b9c714316eef87fcff21a04a187)
- [proxy_config.https.tls_parameters.tls_certificates](resources--bigip_http_proxy--reference--group-003.md#canonical-17e23908f5dd270a57d0f0d511aeb709933f657b78ee076266687a44b9d80651)
- proxy_config.https.tls_parameters.tls_certificates.private_key

<a id="canonical-c77e33afd614778021b9f115f9c1c655255fb73f2b55292f79197111900511a9"></a>

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

<a id="canonical-fdcf5c96877357eccf50199949b37b8b71196f963f16d444daa63d55250321c6"></a>

## Direct properties — proxy_config.https.tls_parameters.tls_certificates.private_key / 4590c7f19a10 / 3

- [blindfold_secret_info](resources--bigip_http_proxy--reference--group-003.md#canonical-2a7aa0f429ca26185c3a1abf8ad744f85de8644418f62888a9e51e1d8a565000): complete subsection reference.

- [clear_secret_info](resources--bigip_http_proxy--reference--group-003.md#canonical-03597980da80a84fd75d30d56c8c986a7bc941672ce358253180aa5304a6ad58): complete subsection reference.

<a id="canonical-e77a25b96a18ee09cf3bc9bfdc6a7b9d5cdca2c6c25cda2c865e41ef6ed7ec3b"></a>

## Next pages — proxy_config.https.tls_parameters.tls_certificates.private_key / 4590c7f19a10 / 4

- [proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info](resources--bigip_http_proxy--reference--group-003.md#canonical-2a7aa0f429ca26185c3a1abf8ad744f85de8644418f62888a9e51e1d8a565000)
- [proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info](resources--bigip_http_proxy--reference--group-003.md#canonical-03597980da80a84fd75d30d56c8c986a7bc941672ce358253180aa5304a6ad58)
- [proxy_config.https.tls_parameters.tls_certificates](resources--bigip_http_proxy--reference--group-003.md#canonical-17e23908f5dd270a57d0f0d511aeb709933f657b78ee076266687a44b9d80651)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-2a7aa0f429ca26185c3a1abf8ad744f85de8644418f62888a9e51e1d8a565000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-663785ac49c1c054c808c625d5bbd3044fe56c76b5d790f76b90dfb684f857af"></a>

## proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info — proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_ / 46138b2e7c44 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-003.md#canonical-1a2fff832c914deeb21855b2b6d870103e691b9c714316eef87fcff21a04a187)
- [proxy_config.https.tls_parameters.tls_certificates](resources--bigip_http_proxy--reference--group-003.md#canonical-17e23908f5dd270a57d0f0d511aeb709933f657b78ee076266687a44b9d80651)
- [proxy_config.https.tls_parameters.tls_certificates.private_key](resources--bigip_http_proxy--reference--group-003.md#canonical-651b77361281bd598c894e0035a49cb75b2b84a8ec1791e1ef3902a8d2f2e2b2)
- proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-c2b3d64a85a37e3ca9aef94de67e22e403fee32369dba423940b67b673b85718"></a>

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

<a id="canonical-8c9b7e50e4717dfee8d17483a57be304be177bab429c98427691bae07eedb61c"></a>

## Direct properties — proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_ / 46138b2e7c44 / 3

<a id="canonical-268eb8209778de1f9e0f2c658e731158e40dcaa33688384feb0aca0a2e80f5d2"></a>

<a id="canonical-a18b09c52a90c3b55471c3451ba2acf817da61f9fd26478495a8e07a203765bb"></a>

## decryption_provider property — proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_ / 46138b2e7c44 / 4

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

<a id="canonical-196cf5cf4a6bd640b04dc8d627d1a76dad203fa8c3b55e57ee5bee2d25e742b1"></a>

<a id="canonical-a6b0e4d0014b1c5739364ef4447e62768eafa1c7742820f4f07f3fd60b03970f"></a>

## location property — proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_ / 46138b2e7c44 / 5

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

<a id="canonical-250a35845114e5a7179e65bb07d4cff6af3f69a125d9a050b2e3ec0390beded4"></a>

<a id="canonical-c00784111a9266636b5e1c7e817a2844ab74a877cea0ab7cda3670dc909b5290"></a>

## store_provider property — proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_ / 46138b2e7c44 / 6

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

<a id="canonical-304f3c6c635093ede683bf14c1638b2073d830572c295145e4dbbe3793ce898a"></a>

## Next pages — proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_ / 46138b2e7c44 / 7

- [proxy_config.https.tls_parameters.tls_certificates.private_key](resources--bigip_http_proxy--reference--group-003.md#canonical-651b77361281bd598c894e0035a49cb75b2b84a8ec1791e1ef3902a8d2f2e2b2)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-03597980da80a84fd75d30d56c8c986a7bc941672ce358253180aa5304a6ad58"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e620fce24d57aa33dd783b2b732802700abea6c602fedf6caa47920bb62ea16c"></a>

## proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info — proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info / 3d1f803fa36d / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-003.md#canonical-1a2fff832c914deeb21855b2b6d870103e691b9c714316eef87fcff21a04a187)
- [proxy_config.https.tls_parameters.tls_certificates](resources--bigip_http_proxy--reference--group-003.md#canonical-17e23908f5dd270a57d0f0d511aeb709933f657b78ee076266687a44b9d80651)
- [proxy_config.https.tls_parameters.tls_certificates.private_key](resources--bigip_http_proxy--reference--group-003.md#canonical-651b77361281bd598c894e0035a49cb75b2b84a8ec1791e1ef3902a8d2f2e2b2)
- proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info

<a id="canonical-d1e5237a84a4a3cfd416f1125ed9ae80598f22ddf1726b7d37004789606771b9"></a>

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

<a id="canonical-0354fef313e86cdde8c960add1c3ee0761f3284de8065fe80d19171ef08bf120"></a>

## Direct properties — proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info / 3d1f803fa36d / 3

<a id="canonical-148e0b4a8b1c1a70a3ced2358e236a3263ba04028c46bc87c52437aba1fa53b6"></a>

<a id="canonical-90faa3c30f40fd7ed6a00cec63602b4a227f68e7299e464a84eafc3d66505160"></a>

## provider_ref property — proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info / 3d1f803fa36d / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-8b899a3ff4ac4aebfbe1c8278a4c5168ca7f1cee9ce1fca20d50ee2eecece695"></a>

<a id="canonical-18083257f01f84582f59f50559ca7fe88145043dd73ba7e25016491ae62f9c9a"></a>

## url property — proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info / 3d1f803fa36d / 5

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

<a id="canonical-229529b0f4f297b5a33133d8991d26ea340264058233d48cd0b7b087427941f5"></a>

## Next pages — proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info / 3d1f803fa36d / 6

- [proxy_config.https.tls_parameters.tls_certificates.private_key](resources--bigip_http_proxy--reference--group-003.md#canonical-651b77361281bd598c894e0035a49cb75b2b84a8ec1791e1ef3902a8d2f2e2b2)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-4e0eb2a89cefeddce9ec0eb3d4d2c4a524fc7cdb34c59b1cd6adc4d73eaa807c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2792d19f438e4e4d9671e0a1bc8832f5e3e7bc85c616f7690fe3fe88a4a94c23"></a>

## proxy_config.https.tls_parameters.tls_certificates.use_system_defaults — proxy_config.https.tls_parameters.tls_certificates.use_system_defaults / 54bfef654eef / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-003.md#canonical-1a2fff832c914deeb21855b2b6d870103e691b9c714316eef87fcff21a04a187)
- [proxy_config.https.tls_parameters.tls_certificates](resources--bigip_http_proxy--reference--group-003.md#canonical-17e23908f5dd270a57d0f0d511aeb709933f657b78ee076266687a44b9d80651)
- proxy_config.https.tls_parameters.tls_certificates.use_system_defaults

<a id="canonical-f08d2fe1755a51720d01252eb692896f4295e54a6e8b22cfae790d1e151a124d"></a>

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

<a id="canonical-98802a7e97a8e39466a37e203675bfe383d9d726429fa3f48a23b81d397031ff"></a>

## Direct properties — proxy_config.https.tls_parameters.tls_certificates.use_system_defaults / 54bfef654eef / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-735c09c723137ef257e2bc00c56053de30d7428419a20bd008a5ff1824427d7a"></a>

## Next pages — proxy_config.https.tls_parameters.tls_certificates.use_system_defaults / 54bfef654eef / 4

- [proxy_config.https.tls_parameters.tls_certificates](resources--bigip_http_proxy--reference--group-003.md#canonical-17e23908f5dd270a57d0f0d511aeb709933f657b78ee076266687a44b9d80651)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-8d3d19560fe789477a691a271848845c6d01c63fd3ae309a100ca27dbb7ba220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-12f15e4432b20aa1bc10c1e8d58176946cfe0cd987427751de849b88036db2f7"></a>

## proxy_config.https.tls_parameters.tls_config — proxy_config.https.tls_parameters.tls_config / 5be0f750d4c8 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-003.md#canonical-1a2fff832c914deeb21855b2b6d870103e691b9c714316eef87fcff21a04a187)
- proxy_config.https.tls_parameters.tls_config

<a id="canonical-93ecc36fa57cb54f9e2d0e9ebde281af2bace84a712d0e725200797ef21838ab"></a>

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

<a id="canonical-c9f720a9f839c6a477bc3a9dc8f808058e0ec2829dcd40aba27be9e2289b4954"></a>

## Direct properties — proxy_config.https.tls_parameters.tls_config / 5be0f750d4c8 / 3

- [custom_security](resources--bigip_http_proxy--reference--group-003.md#canonical-4f4b16161560b96190970f27c434f13c55c7584692c1e1363cf963cb58350f47): complete subsection reference.

- [default_security](resources--bigip_http_proxy--reference--group-003.md#canonical-41c8c6a7a6f50b2bd121886f3a8938577e216d8001ef55aff881647ccb61e210): complete subsection reference.

- [low_security](resources--bigip_http_proxy--reference--group-003.md#canonical-fcd338cb941f66326d694837ca0eab0dce37e0eda4a28b68d95992c18f799e92): complete subsection reference.

- [medium_security](resources--bigip_http_proxy--reference--group-003.md#canonical-4c34d4d1d1c78885cd72fcb0e89623f118e578a61bb4ffe2e2bfede51bd7f241): complete subsection reference.

<a id="canonical-64693358ff02d3cd590b54ca88bdf54fb0e77c1d91e3be4ff72142dfdd4d299f"></a>

## Next pages — proxy_config.https.tls_parameters.tls_config / 5be0f750d4c8 / 4

- [proxy_config.https.tls_parameters.tls_config.custom_security](resources--bigip_http_proxy--reference--group-003.md#canonical-4f4b16161560b96190970f27c434f13c55c7584692c1e1363cf963cb58350f47)
- [proxy_config.https.tls_parameters.tls_config.default_security](resources--bigip_http_proxy--reference--group-003.md#canonical-41c8c6a7a6f50b2bd121886f3a8938577e216d8001ef55aff881647ccb61e210)
- [proxy_config.https.tls_parameters.tls_config.low_security](resources--bigip_http_proxy--reference--group-003.md#canonical-fcd338cb941f66326d694837ca0eab0dce37e0eda4a28b68d95992c18f799e92)
- [proxy_config.https.tls_parameters.tls_config.medium_security](resources--bigip_http_proxy--reference--group-003.md#canonical-4c34d4d1d1c78885cd72fcb0e89623f118e578a61bb4ffe2e2bfede51bd7f241)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-003.md#canonical-1a2fff832c914deeb21855b2b6d870103e691b9c714316eef87fcff21a04a187)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-4f4b16161560b96190970f27c434f13c55c7584692c1e1363cf963cb58350f47"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a16ab4f39a138cf96dd4bb5e410973059678076b2293d0ac7273379fdd0ca204"></a>

## proxy_config.https.tls_parameters.tls_config.custom_security — proxy_config.https.tls_parameters.tls_config.custom_security / 3167896de489 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-003.md#canonical-1a2fff832c914deeb21855b2b6d870103e691b9c714316eef87fcff21a04a187)
- [proxy_config.https.tls_parameters.tls_config](resources--bigip_http_proxy--reference--group-003.md#canonical-8d3d19560fe789477a691a271848845c6d01c63fd3ae309a100ca27dbb7ba220)
- proxy_config.https.tls_parameters.tls_config.custom_security

<a id="canonical-a40d1133e7c8c8606beebf1f5153afb67aca392c74be4734ce5b60017381cd7c"></a>

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

<a id="canonical-2031ac7d421f9c4fea22e5a3c7ab1bb4eca8089948a37e11e3b13e5fd074bbdd"></a>

## Direct properties — proxy_config.https.tls_parameters.tls_config.custom_security / 3167896de489 / 3

<a id="canonical-fb3d6bf292be22848ddabf3b3a5923f95c1cf3dd206bf0f6b434bcab51821573"></a>

<a id="canonical-121030fc3ac2b5675cc78394ae296fbd21558d23ebe7174ceb7ef8a4d96ea3cd"></a>

## cipher_suites property — proxy_config.https.tls_parameters.tls_config.custom_security / 3167896de489 / 4

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

<a id="canonical-af2ac28dca91fd30f505363e35731419202134b0a4160c91473389abf8821558"></a>

<a id="canonical-6703ed3efdc8062f8a2dd711d51351daba8e365aba9a0a09c47d06995c785681"></a>

## max_version property — proxy_config.https.tls_parameters.tls_config.custom_security / 3167896de489 / 5

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

<a id="canonical-114553bf2c513d42647d6cbf803ed14021240a0233d4ab823c36421064f6c210"></a>

<a id="canonical-c57b236b26fbb51586eab9c29ed72ea600d654700a9dd893b4d80c3449915b1c"></a>

## min_version property — proxy_config.https.tls_parameters.tls_config.custom_security / 3167896de489 / 6

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

<a id="canonical-2f9af7662b8420d2174eaf7c6738280b5a5d4c2ece00d8ddf06ac67ee8780a47"></a>

## Next pages — proxy_config.https.tls_parameters.tls_config.custom_security / 3167896de489 / 7

- [proxy_config.https.tls_parameters.tls_config](resources--bigip_http_proxy--reference--group-003.md#canonical-8d3d19560fe789477a691a271848845c6d01c63fd3ae309a100ca27dbb7ba220)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-41c8c6a7a6f50b2bd121886f3a8938577e216d8001ef55aff881647ccb61e210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b51b089fce54fbaa1e5ded92d69766e006ec9760150e7b26005471ea1de2f910"></a>

## proxy_config.https.tls_parameters.tls_config.default_security — proxy_config.https.tls_parameters.tls_config.default_security / d2d0ca7eb163 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-003.md#canonical-1a2fff832c914deeb21855b2b6d870103e691b9c714316eef87fcff21a04a187)
- [proxy_config.https.tls_parameters.tls_config](resources--bigip_http_proxy--reference--group-003.md#canonical-8d3d19560fe789477a691a271848845c6d01c63fd3ae309a100ca27dbb7ba220)
- proxy_config.https.tls_parameters.tls_config.default_security

<a id="canonical-4afd0f8fa2def1390f28512d0526535bc25801d43fe9311496644bdcd06de2f1"></a>

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

<a id="canonical-9b39b859d0d0030dc8bd7712e9d4282925d65fb0383f7b1763f494e9b254ea2d"></a>

## Direct properties — proxy_config.https.tls_parameters.tls_config.default_security / d2d0ca7eb163 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8f7ce3f39832509978228ab1615f2ac51fabfafaa42b6fc1af3eb2fed739de24"></a>

## Next pages — proxy_config.https.tls_parameters.tls_config.default_security / d2d0ca7eb163 / 4

- [proxy_config.https.tls_parameters.tls_config](resources--bigip_http_proxy--reference--group-003.md#canonical-8d3d19560fe789477a691a271848845c6d01c63fd3ae309a100ca27dbb7ba220)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-fcd338cb941f66326d694837ca0eab0dce37e0eda4a28b68d95992c18f799e92"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3605a09f0bfeb694c7fed00f5fe020309eacd81bb42d7915a715c4973a8cb12a"></a>

## proxy_config.https.tls_parameters.tls_config.low_security — proxy_config.https.tls_parameters.tls_config.low_security / e90842c91fb0 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-003.md#canonical-1a2fff832c914deeb21855b2b6d870103e691b9c714316eef87fcff21a04a187)
- [proxy_config.https.tls_parameters.tls_config](resources--bigip_http_proxy--reference--group-003.md#canonical-8d3d19560fe789477a691a271848845c6d01c63fd3ae309a100ca27dbb7ba220)
- proxy_config.https.tls_parameters.tls_config.low_security

<a id="canonical-2642c54c589756b71c250c311fbac0dbdc7e33b39523b54535b7f2618663d480"></a>

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

<a id="canonical-47311faac155ca0efc5c3b8cec5030cf7af5a2ec0b4295f4d2e738862d9c6a7a"></a>

## Direct properties — proxy_config.https.tls_parameters.tls_config.low_security / e90842c91fb0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b2e7680566b58a3dbc07f8fca95e2afcc11800ad8c2b09622e7ad99478a4b8d1"></a>

## Next pages — proxy_config.https.tls_parameters.tls_config.low_security / e90842c91fb0 / 4

- [proxy_config.https.tls_parameters.tls_config](resources--bigip_http_proxy--reference--group-003.md#canonical-8d3d19560fe789477a691a271848845c6d01c63fd3ae309a100ca27dbb7ba220)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-4c34d4d1d1c78885cd72fcb0e89623f118e578a61bb4ffe2e2bfede51bd7f241"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
