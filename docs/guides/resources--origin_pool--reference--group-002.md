---
page_title: "xcsh_origin_pool reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_origin_pool reference."
---

# xcsh_origin_pool reference

<a id="canonical-e733611e652d7e32cc0e8f1b0420191577fcdd8260f728d1323288ac7774c4b4"></a>

## origin_servers.consul_service — origin_servers.consul_service / c5004db889ee / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- origin_servers.consul_service

<a id="canonical-bf5d8bd66042946317ef306d10cb327b81bc31ac2cdf4fd0b5c8a9edeb14cbcd"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with HashiCorp Consul service name and site information.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("service_name"),
  validators.ConflictingObjectAttributes("inside_network",
    "outside_network")}
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
  "x-ves-oneof-field-network_choice": "[\"inside_network\",\"outside_network\"]"
}
```

Terraform syntax:

```terraform
consul_service {
  # Configure direct properties listed below.
}
```

<a id="canonical-33369e2fb972ec20010f7674a8ad8d93118d647c6bd03e58ff1ef347a29be882"></a>

## Direct properties — origin_servers.consul_service / c5004db889ee / 3

- [inside_network](resources--origin_pool--reference--group-002.md#canonical-4a42c57ccac366b39e7deca79afa06cea723c823e726411619ac998d1aaab535): complete subsection reference.

- [outside_network](resources--origin_pool--reference--group-002.md#canonical-1863be90595401f3dae332c3defe999e9f168b945c146138b2d7906463ff6ddf): complete subsection reference.

<a id="canonical-fd653f0d5a6d35504ac6864c8fc0e556003b620622f1fc3e614d8c3cba12d203"></a>

<a id="canonical-7a41b3d9a9104c19408b21fc64b7ab72c07d1023d3689917c3fbce898bdafba2"></a>

## service_name property — origin_servers.consul_service / c5004db889ee / 4

Type: `"string"`. Optional.

Consul service name of this origin server will be listed, including cluster-ID. The format is
servicename:cluster-ID.

Upstream description:

Consul service name of this origin server will be listed, including cluster-ID. The format is
servicename:cluster-ID.

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

- [site_locator](resources--origin_pool--reference--group-002.md#canonical-76461b5602eb417d542d6ffdd9b5f1c271620490392f757ff894c9201d470d1b): complete subsection reference.

- [snat_pool](resources--origin_pool--reference--group-002.md#canonical-1a7b15204832e1a54a2bafcef9441a3296ea89736780c0d9139b7b9a37193e23): complete subsection reference.

<a id="canonical-dbcee336afe0c2a041ab7cd924a934f427578386900e6f86960bb733a0f29860"></a>

## Next pages — origin_servers.consul_service / c5004db889ee / 5

- [origin_servers.consul_service.inside_network](resources--origin_pool--reference--group-002.md#canonical-4a42c57ccac366b39e7deca79afa06cea723c823e726411619ac998d1aaab535)
- [origin_servers.consul_service.outside_network](resources--origin_pool--reference--group-002.md#canonical-1863be90595401f3dae332c3defe999e9f168b945c146138b2d7906463ff6ddf)
- [origin_servers.consul_service.site_locator](resources--origin_pool--reference--group-002.md#canonical-76461b5602eb417d542d6ffdd9b5f1c271620490392f757ff894c9201d470d1b)
- [origin_servers.consul_service.snat_pool](resources--origin_pool--reference--group-002.md#canonical-1a7b15204832e1a54a2bafcef9441a3296ea89736780c0d9139b7b9a37193e23)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-4a42c57ccac366b39e7deca79afa06cea723c823e726411619ac998d1aaab535"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6cdb23cd3d684acf8b5205623346885acc0e45164a512ed209814be4c7f9b916"></a>

## origin_servers.consul_service.inside_network — origin_servers.consul_service.inside_network / d523f489a7a2 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- [origin_servers.consul_service](resources--origin_pool--reference--group-001.md#canonical-e7953cdd2d65579ba3e25ce6ffa198ed7447a59c86e7fccde05f11a418f1b53d)
- origin_servers.consul_service.inside_network

<a id="canonical-540eceffddbc3ea2cc8076910438777def6b70f45d683e0808f5f6eff8fd507a"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for inside network.

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
inside_network = {}
```

<a id="canonical-7ba50b57ff0fc3ad3f56e2c2c9dce621653c558d46c796bd32810ab150ca46aa"></a>

## Direct properties — origin_servers.consul_service.inside_network / d523f489a7a2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e9e678b5966b0c5b2d2fa3f04e867abc50cac48a02ad59d3ffc93a2caef95178"></a>

## Next pages — origin_servers.consul_service.inside_network / d523f489a7a2 / 4

- [origin_servers.consul_service](resources--origin_pool--reference--group-001.md#canonical-e7953cdd2d65579ba3e25ce6ffa198ed7447a59c86e7fccde05f11a418f1b53d)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-1863be90595401f3dae332c3defe999e9f168b945c146138b2d7906463ff6ddf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-01e35082805f20824500f4ab2aba6bb72d58f4d0ab70571b448167bc27f53190"></a>

## origin_servers.consul_service.outside_network — origin_servers.consul_service.outside_network / faca34b8cd2d / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- [origin_servers.consul_service](resources--origin_pool--reference--group-001.md#canonical-e7953cdd2d65579ba3e25ce6ffa198ed7447a59c86e7fccde05f11a418f1b53d)
- origin_servers.consul_service.outside_network

<a id="canonical-ae04953eb23f15b858ed3eec07a9f6f2af80d44dc20f00dcff0a47fb3df7d1a8"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for outside network.

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
outside_network = {}
```

<a id="canonical-29cb3449e2a55a23222512cf697b8152ace6177199b35f6e6aef1b3a92510b67"></a>

## Direct properties — origin_servers.consul_service.outside_network / faca34b8cd2d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-dc881881bd59ed20ae2a681784afa3ef41723eb2fb90e002f3abe17576fc4eb0"></a>

## Next pages — origin_servers.consul_service.outside_network / faca34b8cd2d / 4

- [origin_servers.consul_service](resources--origin_pool--reference--group-001.md#canonical-e7953cdd2d65579ba3e25ce6ffa198ed7447a59c86e7fccde05f11a418f1b53d)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-76461b5602eb417d542d6ffdd9b5f1c271620490392f757ff894c9201d470d1b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-409469e82cacdcf794b80506a927b941f6da58fd1abf284a5443127b39d21f3c"></a>

## origin_servers.consul_service.site_locator — origin_servers.consul_service.site_locator / 697b935a4d1e / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- [origin_servers.consul_service](resources--origin_pool--reference--group-001.md#canonical-e7953cdd2d65579ba3e25ce6ffa198ed7447a59c86e7fccde05f11a418f1b53d)
- origin_servers.consul_service.site_locator

<a id="canonical-b4ade0bd48844219b6cfc73d1378472a73d4c422da4f3295adc499bd27d361c4"></a>

Type: `"object"`. single nested block, Optional.

Message defines a reference to a site or virtual site object.

Upstream description:

This message defines a reference to a site or virtual site object.

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
site_locator {
  # Configure direct properties listed below.
}
```

<a id="canonical-d75bba96d5ee2682d0e431383ac55df5229f89e7f64f126d2743db2623f9018f"></a>

## Direct properties — origin_servers.consul_service.site_locator / 697b935a4d1e / 3

- [site](resources--origin_pool--reference--group-002.md#canonical-c10b368f4ba5c96dfada451628abe2ed0ec919cc076bd336f463be61ed7a5e8f): complete subsection reference.

- [virtual_site](resources--origin_pool--reference--group-002.md#canonical-304fb0e3e1e414df1480668ab5dbc03e81eecd0ebefc543326520b3bbd1885fc): complete subsection reference.

<a id="canonical-004ed12f693c8f14bb5dbbe0b68804ee43d199909abf32961fca36e6c1254168"></a>

## Next pages — origin_servers.consul_service.site_locator / 697b935a4d1e / 4

- [origin_servers.consul_service.site_locator.site](resources--origin_pool--reference--group-002.md#canonical-c10b368f4ba5c96dfada451628abe2ed0ec919cc076bd336f463be61ed7a5e8f)
- [origin_servers.consul_service.site_locator.virtual_site](resources--origin_pool--reference--group-002.md#canonical-304fb0e3e1e414df1480668ab5dbc03e81eecd0ebefc543326520b3bbd1885fc)
- [origin_servers.consul_service](resources--origin_pool--reference--group-001.md#canonical-e7953cdd2d65579ba3e25ce6ffa198ed7447a59c86e7fccde05f11a418f1b53d)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-c10b368f4ba5c96dfada451628abe2ed0ec919cc076bd336f463be61ed7a5e8f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-26924a62b99aeb23e7be72b283b689cbe08760255a902096fb7d725092021ccc"></a>

## origin_servers.consul_service.site_locator.site — origin_servers.consul_service.site_locator.site / fd9db3934827 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- [origin_servers.consul_service](resources--origin_pool--reference--group-001.md#canonical-e7953cdd2d65579ba3e25ce6ffa198ed7447a59c86e7fccde05f11a418f1b53d)
- [origin_servers.consul_service.site_locator](resources--origin_pool--reference--group-002.md#canonical-76461b5602eb417d542d6ffdd9b5f1c271620490392f757ff894c9201d470d1b)
- origin_servers.consul_service.site_locator.site

<a id="canonical-1283d35fa0c67f378ab125a6e0b7dc239a89ffc8890e1056e525de055fb5030c"></a>

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

<a id="canonical-0cec1da085ceabe3e954dda663657649f403ecc251aa84ac7d09c647571ee022"></a>

## Direct properties — origin_servers.consul_service.site_locator.site / fd9db3934827 / 3

<a id="canonical-2c7dd3770855afa108bc3cf50f1dae7f952aa15bc8b8a75be3e2834d13e39446"></a>

<a id="canonical-a7ec66d2b5c2e83fe7c9794c1be12a2ba7bbb4b6aefa41afedc1076291659fbe"></a>

## name property — origin_servers.consul_service.site_locator.site / fd9db3934827 / 4

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

<a id="canonical-78e24096f0cce7a0dcdac8dba9ad9d7a9998111209730a1aab0646e97ecc7daa"></a>

<a id="canonical-eab29f01041a9d5f34f07fe1f6847a1c6a36cee41006abcfcabf2bb9452fa463"></a>

## namespace property — origin_servers.consul_service.site_locator.site / fd9db3934827 / 5

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

<a id="canonical-6ab849b8a33c78fc34cba9bf9224accfce32c9d6b6e8b2edb3119b0d858f1fc2"></a>

<a id="canonical-94f1b902612036146c01912937a39e4ae35cf95a38b2f504849b85e0736db9ac"></a>

## tenant property — origin_servers.consul_service.site_locator.site / fd9db3934827 / 6

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

<a id="canonical-f74bb8719b39ecbfbd27ac68c3d61aba84d2658cf14d6147266100d3cd013ade"></a>

## Next pages — origin_servers.consul_service.site_locator.site / fd9db3934827 / 7

- [origin_servers.consul_service.site_locator](resources--origin_pool--reference--group-002.md#canonical-76461b5602eb417d542d6ffdd9b5f1c271620490392f757ff894c9201d470d1b)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-304fb0e3e1e414df1480668ab5dbc03e81eecd0ebefc543326520b3bbd1885fc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8c14a365b04581767cb57ab2a81c22da99ee7e913c57e521368fc6123f1c731a"></a>

## origin_servers.consul_service.site_locator.virtual_site — origin_servers.consul_service.site_locator.virtual_site / 18551d84deee / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- [origin_servers.consul_service](resources--origin_pool--reference--group-001.md#canonical-e7953cdd2d65579ba3e25ce6ffa198ed7447a59c86e7fccde05f11a418f1b53d)
- [origin_servers.consul_service.site_locator](resources--origin_pool--reference--group-002.md#canonical-76461b5602eb417d542d6ffdd9b5f1c271620490392f757ff894c9201d470d1b)
- origin_servers.consul_service.site_locator.virtual_site

<a id="canonical-49bc0070edca322e346f277f37e3cd5a85c017c289ab58ac0d02fccfa3582856"></a>

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

<a id="canonical-2bbfd40106e400b2c40d9b66ff32818cd5a329149ce2ff56d07b498ddc22ef2c"></a>

## Direct properties — origin_servers.consul_service.site_locator.virtual_site / 18551d84deee / 3

<a id="canonical-134f49b25a1bba16a7762ea30eb4ed38d114344d88420fb468fe01612901b230"></a>

<a id="canonical-c7ddf91a030c1e72b592195d5a6fa0b31ce61975972b58f9006e1bceffba4086"></a>

## name property — origin_servers.consul_service.site_locator.virtual_site / 18551d84deee / 4

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

<a id="canonical-e11c62cc0590dab52b5ce8e1c6c7a238ca57f55e5d35eaec97e0e6f54680f25f"></a>

<a id="canonical-96df2de0ad300ad80ae37ef49a50c11527d1b39eade92da79f1c5c90312788ce"></a>

## namespace property — origin_servers.consul_service.site_locator.virtual_site / 18551d84deee / 5

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

<a id="canonical-3c358daa92a4638e7e857dd0511715b7240b1d724545823465df21f568da1c5e"></a>

<a id="canonical-65b340b2a84a955389dff555b42ae67d6443023db9d0ca52e73722c4958546df"></a>

## tenant property — origin_servers.consul_service.site_locator.virtual_site / 18551d84deee / 6

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

<a id="canonical-c7bd6a95b2dacae350e7d5a43cf2bab65e816309caef73b9f8f9f033a2c78564"></a>

## Next pages — origin_servers.consul_service.site_locator.virtual_site / 18551d84deee / 7

- [origin_servers.consul_service.site_locator](resources--origin_pool--reference--group-002.md#canonical-76461b5602eb417d542d6ffdd9b5f1c271620490392f757ff894c9201d470d1b)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-1a7b15204832e1a54a2bafcef9441a3296ea89736780c0d9139b7b9a37193e23"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fb72bbae84ea3829658d65c0b14a0a55ddb8c567c44c7f91e0e8eca2ba8f612b"></a>

## origin_servers.consul_service.snat_pool — origin_servers.consul_service.snat_pool / 334537b0b072 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- [origin_servers.consul_service](resources--origin_pool--reference--group-001.md#canonical-e7953cdd2d65579ba3e25ce6ffa198ed7447a59c86e7fccde05f11a418f1b53d)
- origin_servers.consul_service.snat_pool

<a id="canonical-f45b62c68cfeb0b8eb377807de1cd77e058c708438cf963b8bcc1ad4184396be"></a>

Type: `"object"`. single nested block, Optional.

SNAT Pool. SNAT Pool configuration.

Upstream description:

SNAT Pool configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_snat_pool",
    "snat_pool")}
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
  "x-ves-oneof-field-snat_pool_choice": "[\"no_snat_pool\",\"snat_pool\"]"
}
```

Terraform syntax:

```terraform
snat_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-ecdf14410d7155f998d33cdea21a0e6112fe2fd74fe5ada454b50585cae5fb85"></a>

## Direct properties — origin_servers.consul_service.snat_pool / 334537b0b072 / 3

- [no_snat_pool](resources--origin_pool--reference--group-002.md#canonical-75c623b10477b679cba11a1f73a70b233d863ad2084dca4d8e0f318428be97f8): complete subsection reference.

- [snat_pool](resources--origin_pool--reference--group-002.md#canonical-1423da2f61e512ae3064ad95aad210ccf2521205e27ac12044f197328b36f972): complete subsection reference.

<a id="canonical-5cfe2dc8dd582e060a5a57ad30fb08799a0f3c90175d2a6db8a2ee8c8d0a3dd3"></a>

## Next pages — origin_servers.consul_service.snat_pool / 334537b0b072 / 4

- [origin_servers.consul_service.snat_pool.no_snat_pool](resources--origin_pool--reference--group-002.md#canonical-75c623b10477b679cba11a1f73a70b233d863ad2084dca4d8e0f318428be97f8)
- [origin_servers.consul_service.snat_pool.snat_pool](resources--origin_pool--reference--group-002.md#canonical-1423da2f61e512ae3064ad95aad210ccf2521205e27ac12044f197328b36f972)
- [origin_servers.consul_service](resources--origin_pool--reference--group-001.md#canonical-e7953cdd2d65579ba3e25ce6ffa198ed7447a59c86e7fccde05f11a418f1b53d)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-75c623b10477b679cba11a1f73a70b233d863ad2084dca4d8e0f318428be97f8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e3fc8a21b95b8bc0ba3bed8f9db883f30d48e9c90beccce977c8d402c7c25254"></a>

## origin_servers.consul_service.snat_pool.no_snat_pool — origin_servers.consul_service.snat_pool.no_snat_pool / 2ef28bbb45d6 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- [origin_servers.consul_service](resources--origin_pool--reference--group-001.md#canonical-e7953cdd2d65579ba3e25ce6ffa198ed7447a59c86e7fccde05f11a418f1b53d)
- [origin_servers.consul_service.snat_pool](resources--origin_pool--reference--group-002.md#canonical-1a7b15204832e1a54a2bafcef9441a3296ea89736780c0d9139b7b9a37193e23)
- origin_servers.consul_service.snat_pool.no_snat_pool

<a id="canonical-65146b3867f400aafe696cfe3850fc25bdb046ac6a0d018df51aa51d5770a369"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no snat pool.

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
no_snat_pool = {}
```

<a id="canonical-6ebc3b543ed502b463e768270d14eb03e0dcd8b169ec5810a6eacf19204a3e5f"></a>

## Direct properties — origin_servers.consul_service.snat_pool.no_snat_pool / 2ef28bbb45d6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-dc3418af59a5875c82e4cffee7e12be4d6f3663d4957a28e7cbbf5c484edcdcb"></a>

## Next pages — origin_servers.consul_service.snat_pool.no_snat_pool / 2ef28bbb45d6 / 4

- [origin_servers.consul_service.snat_pool](resources--origin_pool--reference--group-002.md#canonical-1a7b15204832e1a54a2bafcef9441a3296ea89736780c0d9139b7b9a37193e23)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-1423da2f61e512ae3064ad95aad210ccf2521205e27ac12044f197328b36f972"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9750ee571132b90d494c40fd0a9ebf416443937bf5516f471f61b16dd7f4b5a5"></a>

## origin_servers.consul_service.snat_pool.snat_pool — origin_servers.consul_service.snat_pool.snat_pool / 6e3d747f69fc / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- [origin_servers.consul_service](resources--origin_pool--reference--group-001.md#canonical-e7953cdd2d65579ba3e25ce6ffa198ed7447a59c86e7fccde05f11a418f1b53d)
- [origin_servers.consul_service.snat_pool](resources--origin_pool--reference--group-002.md#canonical-1a7b15204832e1a54a2bafcef9441a3296ea89736780c0d9139b7b9a37193e23)
- origin_servers.consul_service.snat_pool.snat_pool

<a id="canonical-25866dce6896ba9da923277acfc55e21dad742ff6276874c393bc29708b4bf40"></a>

Type: `"object"`. single nested block, Optional.

List of IPv4 prefixes that represent an endpoint.

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
snat_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-9cb4faaf19870f61d89c0960c908e3ee7e04a2d1e824b11d48e2fe2c99c1abd8"></a>

## Direct properties — origin_servers.consul_service.snat_pool.snat_pool / 6e3d747f69fc / 3

<a id="canonical-88611a1e5f9ed9876c7072c00ea1dbc1a2e8531d31d3559e86e0939ed0fb2ec8"></a>

<a id="canonical-59d81ba9080ab23e74e42c42d35020b3e8097c0a2efd9f5a53893d4be80dc04b"></a>

## prefixes property — origin_servers.consul_service.snat_pool.snat_pool / 6e3d747f69fc / 4

Type: `["list", "string"]`. Optional.

List of IPv4 prefixes that represent an endpoint.

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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-db49c4bdf4048fbd8386c86ca1b9bb22104b6555f50c13360fee38976a36139c"></a>

## Next pages — origin_servers.consul_service.snat_pool.snat_pool / 6e3d747f69fc / 5

- [origin_servers.consul_service.snat_pool](resources--origin_pool--reference--group-002.md#canonical-1a7b15204832e1a54a2bafcef9441a3296ea89736780c0d9139b7b9a37193e23)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-4707dd4589e4425912ca819bdba0f57eaa059d66a65782ad2e52450b8c12e6f9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c50a5022e2953e8643b4f32c47d4d5bbce6457378611a340093daf2226bfcc52"></a>

## origin_servers.custom_endpoint_object — origin_servers.custom_endpoint_object / 9bd57e7fdd03 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- origin_servers.custom_endpoint_object

<a id="canonical-722a6f33a3ea2749dafa392f2bc314fc4f553f9e07d654e5d971d65338b5e171"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with a reference to endpoint object.

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
custom_endpoint_object {
  # Configure direct properties listed below.
}
```

<a id="canonical-e6d12c0fcf66f3a92f306e4db62f014fdacc8cb521e43a8d3fae5ca9baedc77e"></a>

## Direct properties — origin_servers.custom_endpoint_object / 9bd57e7fdd03 / 3

- [endpoint](resources--origin_pool--reference--group-002.md#canonical-3b5a1ed5b3d6903ee6933e0a56e171055ca0740a13ffca0883dd05855e83d4da): complete subsection reference.

<a id="canonical-b70fdd85ccf3ec50d84d79ff4f837577c9c61593dce6bb9666f43794348aa661"></a>

## Next pages — origin_servers.custom_endpoint_object / 9bd57e7fdd03 / 4

- [origin_servers.custom_endpoint_object.endpoint](resources--origin_pool--reference--group-002.md#canonical-3b5a1ed5b3d6903ee6933e0a56e171055ca0740a13ffca0883dd05855e83d4da)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-3b5a1ed5b3d6903ee6933e0a56e171055ca0740a13ffca0883dd05855e83d4da"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-775e87af69336d8bd60ed4d6c009891b05e9425ae9679b19ae1a0972895f47d9"></a>

## origin_servers.custom_endpoint_object.endpoint — origin_servers.custom_endpoint_object.endpoint / f11b4a5f7a6f / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- [origin_servers.custom_endpoint_object](resources--origin_pool--reference--group-002.md#canonical-4707dd4589e4425912ca819bdba0f57eaa059d66a65782ad2e52450b8c12e6f9)
- origin_servers.custom_endpoint_object.endpoint

<a id="canonical-a81bd5af47328b0ba2600822e2823fac1084cd286c7ddbb3cadb49fc6b713d59"></a>

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
endpoint {
  # Configure direct properties listed below.
}
```

<a id="canonical-92a360f59a7517a78ae79fc975a7383d3afb41b0f39ab6d4357e367e6ceef079"></a>

## Direct properties — origin_servers.custom_endpoint_object.endpoint / f11b4a5f7a6f / 3

<a id="canonical-03e165f02fd524e60e9393437558b9b089b3ab08b5f0e6012d92fec8191ae8d9"></a>

<a id="canonical-2f307903df971c33506281abfb3ae1fc600d95e2335cf70779c4786b9bcce6af"></a>

## name property — origin_servers.custom_endpoint_object.endpoint / f11b4a5f7a6f / 4

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

<a id="canonical-674455876f4f17b283c3099a5fa280a57f5175a96fa6e03730575c2e360d8496"></a>

<a id="canonical-a631443a1f732f2e1dd6ce9ee7fa3342b1a887cbda231814ce625344b6a8f40d"></a>

## namespace property — origin_servers.custom_endpoint_object.endpoint / f11b4a5f7a6f / 5

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

<a id="canonical-fcb66f47914386d9b69e938f50be5d9538f02a8de039c44bdcc45686a4461f29"></a>

<a id="canonical-933246303b327bede0beebce0a18cc80bdf02eb5ef616236f368e21b3792ca50"></a>

## tenant property — origin_servers.custom_endpoint_object.endpoint / f11b4a5f7a6f / 6

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

<a id="canonical-9cbe7b8e85016f31a175ae07cfb1350e0e73167f331e3e0d9f9d4dc7c3ff9d17"></a>

## Next pages — origin_servers.custom_endpoint_object.endpoint / f11b4a5f7a6f / 7

- [origin_servers.custom_endpoint_object](resources--origin_pool--reference--group-002.md#canonical-4707dd4589e4425912ca819bdba0f57eaa059d66a65782ad2e52450b8c12e6f9)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-807eedd9067c3254887005dfeaa77d3a564cb3cacb261d63c3ed9cb49a836cf0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0cdec1b8652d675fd6da972942011c7ddfbac9272e31914f000c9d80baabe63f"></a>

## origin_servers.k8s_service — origin_servers.k8s_service / f5190b4825a3 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- origin_servers.k8s_service

<a id="canonical-4cc36d072e3b836f3963a4006394f79d05773937f6f93137562620853f55262f"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with K8s service name and site information.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("inside_network",
    "outside_network"),
  validators.ConflictingObjectAttributes("inside_network",
    "vk8s_networks"),
  validators.ConflictingObjectAttributes("outside_network",
    "vk8s_networks")}
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
  "x-ves-oneof-field-network_choice": "[\"inside_network\",\"outside_network\",\"vk8s_networks\"]",
  "x-ves-oneof-field-service_info": "[\"service_name\"]"
}
```

Terraform syntax:

```terraform
k8s_service {
  # Configure direct properties listed below.
}
```

<a id="canonical-bea995db80ad66ea3b701d9308ada6cf4fea19922711f07f7b97f4ab5c67bd12"></a>

## Direct properties — origin_servers.k8s_service / f5190b4825a3 / 3

- [inside_network](resources--origin_pool--reference--group-002.md#canonical-013ff0765402bbd2711ed544a2799121019703ca5a6dce627b8b0cbf743b1e85): complete subsection reference.

- [outside_network](resources--origin_pool--reference--group-002.md#canonical-8e3f268a53548041b3c7f9c1d2e9ce97d7b480abcf749a7e89b6a422c5d5008a): complete subsection reference.

<a id="canonical-356706e2a9910b0a6fc13a0b8729a852116f168e187474761622561f6ce8b1b3"></a>

<a id="canonical-263f5e49d08eec011e51e62ca8f2173b4f4e826f7bf07114831ea1ff404e5154"></a>

## protocol property — origin_servers.k8s_service / f5190b4825a3 / 4

Type: `"string"`. Optional.

\[Enum: PROTOCOL\_TCP|PROTOCOL\_UDP\] Type of protocol - PROTOCOL\_TCP: TCP - PROTOCOL\_UDP: UDP.
Possible values are \`PROTOCOL\_TCP\`, \`PROTOCOL\_UDP\`. Defaults to \`PROTOCOL\_TCP\`.

Upstream description:

Type of protocol

&#8203;- PROTOCOL\_TCP: TCP

&#8203;- PROTOCOL\_UDP: UDP.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("PROTOCOL_TCP",
    "PROTOCOL_UDP"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "PROTOCOL_TCP",
  "enum": [
    "PROTOCOL_TCP",
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

<a id="canonical-d90bd5cdec234deaa49aaadee54f0268019d2f893292dd563355ac649c77f322"></a>

<a id="canonical-81bea2b76bc571a1a52511b2979a33b3d659ae240ce7e23e6574a96f40337b28"></a>

## service_name property — origin_servers.k8s_service / f5190b4825a3 / 5

Type: `"string"`. Optional.

Exclusive with \[\] K8s service name of the origin server will be listed, including the namespace
and cluster-ID. For vK8s services, you need to enter a string with the format
servicename.namespace:example-namespace'frontend', namespace is 'speedtest' and cluster-ID is
'prod', then you will enter..

Upstream description:

Exclusive with \[\] K8s service name of the origin server will be listed, including the namespace
and cluster-ID. For vK8s services, you need to enter a string with the format
servicename.namespace:example-namespace"frontend", namespace is "speedtest" and cluster-ID is
"prod", then you will enter "frontend.speedtest:prod". Both namespace and cluster-ID are optional.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ves_service_namespace_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ves_service_namespace_name": "true"
  }
}
```

- [site_locator](resources--origin_pool--reference--group-002.md#canonical-fec8180f6858352f0a73bf4bdd71d1ea4f6444468b306a3e52c57e0a976c6978): complete subsection reference.

- [snat_pool](resources--origin_pool--reference--group-002.md#canonical-a1b37e91760eefcb13c5b47a96707eaf7eb332c254fe8c0da41bec0be2669f92): complete subsection reference.

- [vk8s_networks](resources--origin_pool--reference--group-002.md#canonical-84590abf9cd0356d75e014f13d186b60d98fe9a6532904b7b487563c9812986c): complete subsection reference.

<a id="canonical-4c11742dcf0dbbf99d11a2be579572f03ccf3677913b686949246024ea1d2a6f"></a>

## Next pages — origin_servers.k8s_service / f5190b4825a3 / 6

- [origin_servers.k8s_service.inside_network](resources--origin_pool--reference--group-002.md#canonical-013ff0765402bbd2711ed544a2799121019703ca5a6dce627b8b0cbf743b1e85)
- [origin_servers.k8s_service.outside_network](resources--origin_pool--reference--group-002.md#canonical-8e3f268a53548041b3c7f9c1d2e9ce97d7b480abcf749a7e89b6a422c5d5008a)
- [origin_servers.k8s_service.site_locator](resources--origin_pool--reference--group-002.md#canonical-fec8180f6858352f0a73bf4bdd71d1ea4f6444468b306a3e52c57e0a976c6978)
- [origin_servers.k8s_service.snat_pool](resources--origin_pool--reference--group-002.md#canonical-a1b37e91760eefcb13c5b47a96707eaf7eb332c254fe8c0da41bec0be2669f92)
- [origin_servers.k8s_service.vk8s_networks](resources--origin_pool--reference--group-002.md#canonical-84590abf9cd0356d75e014f13d186b60d98fe9a6532904b7b487563c9812986c)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-013ff0765402bbd2711ed544a2799121019703ca5a6dce627b8b0cbf743b1e85"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b4596afea247571d09fdd9da78ae0587f4a5489d954356b0ee58ed30cf4eabc7"></a>

## origin_servers.k8s_service.inside_network — origin_servers.k8s_service.inside_network / 7c5695749ba5 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- [origin_servers.k8s_service](resources--origin_pool--reference--group-002.md#canonical-807eedd9067c3254887005dfeaa77d3a564cb3cacb261d63c3ed9cb49a836cf0)
- origin_servers.k8s_service.inside_network

<a id="canonical-4bc32f8a429bf8d320f06fe432f84355781c6d641025e83f4ac211efb7bcba3d"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for inside network.

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
inside_network = {}
```

<a id="canonical-d86cc239cf8519bdd89d88f17fcf9ec82ab99700823d371cef64c312a6f55a17"></a>

## Direct properties — origin_servers.k8s_service.inside_network / 7c5695749ba5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3debd383ba1bc8b8227e6c24e8bd9bc0a20e9f705bf2b9de7b79d6a1c2de4274"></a>

## Next pages — origin_servers.k8s_service.inside_network / 7c5695749ba5 / 4

- [origin_servers.k8s_service](resources--origin_pool--reference--group-002.md#canonical-807eedd9067c3254887005dfeaa77d3a564cb3cacb261d63c3ed9cb49a836cf0)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-8e3f268a53548041b3c7f9c1d2e9ce97d7b480abcf749a7e89b6a422c5d5008a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2c797ec6a6e0d988f176888aafb294cd5c4cacf2f927230ac15e6797bc0bcea8"></a>

## origin_servers.k8s_service.outside_network — origin_servers.k8s_service.outside_network / b483427cc243 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- [origin_servers.k8s_service](resources--origin_pool--reference--group-002.md#canonical-807eedd9067c3254887005dfeaa77d3a564cb3cacb261d63c3ed9cb49a836cf0)
- origin_servers.k8s_service.outside_network

<a id="canonical-209c35de8baadf74ac5fbc62a5a8d7850ffa035c8763beb97c48ebeed19fd304"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for outside network.

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
outside_network = {}
```

<a id="canonical-7708b0bdc37154d91fcdc7da318f0de18b3c9f973d223b050e3638e467bcb1c3"></a>

## Direct properties — origin_servers.k8s_service.outside_network / b483427cc243 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7747998e416e79e34d99f63bfd5b3fe435896079005805e9f02d27cc2f341cbc"></a>

## Next pages — origin_servers.k8s_service.outside_network / b483427cc243 / 4

- [origin_servers.k8s_service](resources--origin_pool--reference--group-002.md#canonical-807eedd9067c3254887005dfeaa77d3a564cb3cacb261d63c3ed9cb49a836cf0)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-fec8180f6858352f0a73bf4bdd71d1ea4f6444468b306a3e52c57e0a976c6978"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c60bb12c57f131f54bab9ad8de1c7464f03206f9741828863ca1de9be233a62a"></a>

## origin_servers.k8s_service.site_locator — origin_servers.k8s_service.site_locator / c244fee394fa / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- [origin_servers.k8s_service](resources--origin_pool--reference--group-002.md#canonical-807eedd9067c3254887005dfeaa77d3a564cb3cacb261d63c3ed9cb49a836cf0)
- origin_servers.k8s_service.site_locator

<a id="canonical-24511af17e95e4c39a461030ee8a80cc3b7a0e894d32cc50c7cdef56710ad4be"></a>

Type: `"object"`. single nested block, Optional.

Message defines a reference to a site or virtual site object.

Upstream description:

This message defines a reference to a site or virtual site object.

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
site_locator {
  # Configure direct properties listed below.
}
```

<a id="canonical-73e2f0c3a8864f190ee8c7f8f76f5e96d518141952aeafa3b7677325ebe9d310"></a>

## Direct properties — origin_servers.k8s_service.site_locator / c244fee394fa / 3

- [site](resources--origin_pool--reference--group-002.md#canonical-d07e5e01f640cbb8cdc44b31fc32382c0f364af88480ffc4eb15a531f89558cd): complete subsection reference.

- [virtual_site](resources--origin_pool--reference--group-002.md#canonical-0ac865239cf1eb3cd171b1cdbedf0d1e375df9d4d012b44fafc42be973c29849): complete subsection reference.

<a id="canonical-a2ef4f65cad3db8888ab566d2e9c26136066d92ab1dc2a3f88ff0f82e719e438"></a>

## Next pages — origin_servers.k8s_service.site_locator / c244fee394fa / 4

- [origin_servers.k8s_service.site_locator.site](resources--origin_pool--reference--group-002.md#canonical-d07e5e01f640cbb8cdc44b31fc32382c0f364af88480ffc4eb15a531f89558cd)
- [origin_servers.k8s_service.site_locator.virtual_site](resources--origin_pool--reference--group-002.md#canonical-0ac865239cf1eb3cd171b1cdbedf0d1e375df9d4d012b44fafc42be973c29849)
- [origin_servers.k8s_service](resources--origin_pool--reference--group-002.md#canonical-807eedd9067c3254887005dfeaa77d3a564cb3cacb261d63c3ed9cb49a836cf0)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-d07e5e01f640cbb8cdc44b31fc32382c0f364af88480ffc4eb15a531f89558cd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-da7a279be67083ef195c5d177c57503eb08fa541a1d348918e98b9a629158a65"></a>

## origin_servers.k8s_service.site_locator.site — origin_servers.k8s_service.site_locator.site / 26fbebe063da / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- [origin_servers.k8s_service](resources--origin_pool--reference--group-002.md#canonical-807eedd9067c3254887005dfeaa77d3a564cb3cacb261d63c3ed9cb49a836cf0)
- [origin_servers.k8s_service.site_locator](resources--origin_pool--reference--group-002.md#canonical-fec8180f6858352f0a73bf4bdd71d1ea4f6444468b306a3e52c57e0a976c6978)
- origin_servers.k8s_service.site_locator.site

<a id="canonical-de98207845ff9afe786dba67d5eeb362c67d98d522cc23dcddfb0f41bfebc6ca"></a>

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

<a id="canonical-4db50ac4e6507fdf567b0a8d18ffdd0722d37afba2df38ea70e94ab5a6956352"></a>

## Direct properties — origin_servers.k8s_service.site_locator.site / 26fbebe063da / 3

<a id="canonical-6dcb4caa518c1bae5ccb8e4f00e1c4ec010ad1acf5719ccdce53bf6b571c6bc6"></a>

<a id="canonical-0a2a150d8d2db8d50d76f2b6c5969cba802d6683b9c4cc86a189ce0ff19ca833"></a>

## name property — origin_servers.k8s_service.site_locator.site / 26fbebe063da / 4

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

<a id="canonical-801f1e4681b1603c215908c3579fc78e706417121ac01d7c02fe6f9c05431a68"></a>

<a id="canonical-b4deb279c0abb2d973f0191f0d5fcce5914c9e1a7fea73e9f5156a05e6e10468"></a>

## namespace property — origin_servers.k8s_service.site_locator.site / 26fbebe063da / 5

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

<a id="canonical-a4101f3b356aedb695c6bfa7a04807f00ee83ebde582664fbf0b4bd3b88b5e26"></a>

<a id="canonical-3f332e01ecd705e87bdb4e56038b5807d4b62a0dde495447634c372f7503af09"></a>

## tenant property — origin_servers.k8s_service.site_locator.site / 26fbebe063da / 6

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

<a id="canonical-19d99cb31406efb0ac584ac327161451d8e9d18570f830098d84a83c11a870f0"></a>

## Next pages — origin_servers.k8s_service.site_locator.site / 26fbebe063da / 7

- [origin_servers.k8s_service.site_locator](resources--origin_pool--reference--group-002.md#canonical-fec8180f6858352f0a73bf4bdd71d1ea4f6444468b306a3e52c57e0a976c6978)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-0ac865239cf1eb3cd171b1cdbedf0d1e375df9d4d012b44fafc42be973c29849"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-99cf730d573eb17cb8625ad7a8bf6987d241587c0bd977bb4b908bef795a941c"></a>

## origin_servers.k8s_service.site_locator.virtual_site — origin_servers.k8s_service.site_locator.virtual_site / d3393cbbd159 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- [origin_servers.k8s_service](resources--origin_pool--reference--group-002.md#canonical-807eedd9067c3254887005dfeaa77d3a564cb3cacb261d63c3ed9cb49a836cf0)
- [origin_servers.k8s_service.site_locator](resources--origin_pool--reference--group-002.md#canonical-fec8180f6858352f0a73bf4bdd71d1ea4f6444468b306a3e52c57e0a976c6978)
- origin_servers.k8s_service.site_locator.virtual_site

<a id="canonical-a269e26df9220ee07488cdf47cb035e08b612288d1ec4e39387ca101f678bf1d"></a>

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

<a id="canonical-09149a9ecfc439f8ce207ae3f2b2c0cb63a354449cd98be64b69aed40da97111"></a>

## Direct properties — origin_servers.k8s_service.site_locator.virtual_site / d3393cbbd159 / 3

<a id="canonical-55fd434ee75ed6b9deb1e21b5bcbfbf8a15d738ef11f761525749453c5d6ec2e"></a>

<a id="canonical-f13cf7e1160b41510d95186b0ae6025c8e2a2300334425cef0bf98fd1377b933"></a>

## name property — origin_servers.k8s_service.site_locator.virtual_site / d3393cbbd159 / 4

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

<a id="canonical-36315ceb239d1acc4110673cc4a483090dc86d5bb1f7925662886c6a154c9d06"></a>

<a id="canonical-323715f919d6f04e00eea1a39ec41e2a5f075477f2aa5a2e5407d4dcd8773cfb"></a>

## namespace property — origin_servers.k8s_service.site_locator.virtual_site / d3393cbbd159 / 5

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

<a id="canonical-97e6fd5e5d826832b836f76255a771a41d51bb79e1ccc3a82298237d63a8b930"></a>

<a id="canonical-99165c567feb0ce99f52d01e50a06674bc04908d1d4c957f7145567c6313ef98"></a>

## tenant property — origin_servers.k8s_service.site_locator.virtual_site / d3393cbbd159 / 6

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

<a id="canonical-4dcceccf994e1aedecbd1e62ca7b4130d8ccbdb0a05bd78ce10d8c76086186b5"></a>

## Next pages — origin_servers.k8s_service.site_locator.virtual_site / d3393cbbd159 / 7

- [origin_servers.k8s_service.site_locator](resources--origin_pool--reference--group-002.md#canonical-fec8180f6858352f0a73bf4bdd71d1ea4f6444468b306a3e52c57e0a976c6978)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-a1b37e91760eefcb13c5b47a96707eaf7eb332c254fe8c0da41bec0be2669f92"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d9574707d5a4d1126e588b3c1f6c0739b40a18a56506288c33f1b035a83dcc82"></a>

## origin_servers.k8s_service.snat_pool — origin_servers.k8s_service.snat_pool / b7fc9436db72 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- [origin_servers.k8s_service](resources--origin_pool--reference--group-002.md#canonical-807eedd9067c3254887005dfeaa77d3a564cb3cacb261d63c3ed9cb49a836cf0)
- origin_servers.k8s_service.snat_pool

<a id="canonical-77fb38331d336a35ec65521c63972927f194ef30063f660cf1dbaad94edaa74f"></a>

Type: `"object"`. single nested block, Optional.

SNAT Pool. SNAT Pool configuration.

Upstream description:

SNAT Pool configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_snat_pool",
    "snat_pool")}
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
  "x-ves-oneof-field-snat_pool_choice": "[\"no_snat_pool\",\"snat_pool\"]"
}
```

Terraform syntax:

```terraform
snat_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-d7c99128fe2efed5fc4959cadb1a3df25226fd711be90e4d0ff892f4b1702f9a"></a>

## Direct properties — origin_servers.k8s_service.snat_pool / b7fc9436db72 / 3

- [no_snat_pool](resources--origin_pool--reference--group-002.md#canonical-d831bc02f2a22f559c92339f640d96267da17df3cbf7ccfbb2c99e495d04c7db): complete subsection reference.

- [snat_pool](resources--origin_pool--reference--group-002.md#canonical-234095206a9dc0b973af98506a0d5692cd4b7792a63dac6fd01ac1b1e9de359c): complete subsection reference.

<a id="canonical-9ece2a397e5d17ffaca8a31a7f88385b3d75e3a3838efa7fdf2244bb3c0a4bd7"></a>

## Next pages — origin_servers.k8s_service.snat_pool / b7fc9436db72 / 4

- [origin_servers.k8s_service.snat_pool.no_snat_pool](resources--origin_pool--reference--group-002.md#canonical-d831bc02f2a22f559c92339f640d96267da17df3cbf7ccfbb2c99e495d04c7db)
- [origin_servers.k8s_service.snat_pool.snat_pool](resources--origin_pool--reference--group-002.md#canonical-234095206a9dc0b973af98506a0d5692cd4b7792a63dac6fd01ac1b1e9de359c)
- [origin_servers.k8s_service](resources--origin_pool--reference--group-002.md#canonical-807eedd9067c3254887005dfeaa77d3a564cb3cacb261d63c3ed9cb49a836cf0)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-d831bc02f2a22f559c92339f640d96267da17df3cbf7ccfbb2c99e495d04c7db"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e6dbf463170df1bb91ddd3a8a57355b0a92920086788a0c8d1529f3acc80ba3f"></a>

## origin_servers.k8s_service.snat_pool.no_snat_pool — origin_servers.k8s_service.snat_pool.no_snat_pool / ae58f2ae32bd / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- [origin_servers.k8s_service](resources--origin_pool--reference--group-002.md#canonical-807eedd9067c3254887005dfeaa77d3a564cb3cacb261d63c3ed9cb49a836cf0)
- [origin_servers.k8s_service.snat_pool](resources--origin_pool--reference--group-002.md#canonical-a1b37e91760eefcb13c5b47a96707eaf7eb332c254fe8c0da41bec0be2669f92)
- origin_servers.k8s_service.snat_pool.no_snat_pool

<a id="canonical-e4133e83c926000ffb9985337e03d2e65fbd226c7a9adfc6a8773734d8885c63"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no snat pool.

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
no_snat_pool = {}
```

<a id="canonical-d0739a0bbc929ec0e15463658f0140603a940ea89044346c7b0b89b6c55bc0fd"></a>

## Direct properties — origin_servers.k8s_service.snat_pool.no_snat_pool / ae58f2ae32bd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-036d55913a4aa06cf2d13311249053a63dc6a1716e213a26cef92ff549e808dd"></a>

## Next pages — origin_servers.k8s_service.snat_pool.no_snat_pool / ae58f2ae32bd / 4

- [origin_servers.k8s_service.snat_pool](resources--origin_pool--reference--group-002.md#canonical-a1b37e91760eefcb13c5b47a96707eaf7eb332c254fe8c0da41bec0be2669f92)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-234095206a9dc0b973af98506a0d5692cd4b7792a63dac6fd01ac1b1e9de359c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cbc487798df871e955f65b29a6e54063746e6a10f18dda363f5c81c75c17d41e"></a>

## origin_servers.k8s_service.snat_pool.snat_pool — origin_servers.k8s_service.snat_pool.snat_pool / 86d4a514f141 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- [origin_servers.k8s_service](resources--origin_pool--reference--group-002.md#canonical-807eedd9067c3254887005dfeaa77d3a564cb3cacb261d63c3ed9cb49a836cf0)
- [origin_servers.k8s_service.snat_pool](resources--origin_pool--reference--group-002.md#canonical-a1b37e91760eefcb13c5b47a96707eaf7eb332c254fe8c0da41bec0be2669f92)
- origin_servers.k8s_service.snat_pool.snat_pool

<a id="canonical-b10aa12219775810e3f7abbbf722b8e4919e11a797318f420756d478e2d8108d"></a>

Type: `"object"`. single nested block, Optional.

List of IPv4 prefixes that represent an endpoint.

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
snat_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-7ac27025be741f957cee51c6ed3d9df2f349fb37d48fb58e4d4805fe3f07e806"></a>

## Direct properties — origin_servers.k8s_service.snat_pool.snat_pool / 86d4a514f141 / 3

<a id="canonical-2e8de694041bb17a4fc323a0aa6311f4744a4aaa0fca7feb0ba7e9ea5b99ce1b"></a>

<a id="canonical-30ae3686c642b669fd8f5509165f924cce54d59784fe6b6e0e1b9697a1ff10ff"></a>

## prefixes property — origin_servers.k8s_service.snat_pool.snat_pool / 86d4a514f141 / 4

Type: `["list", "string"]`. Optional.

List of IPv4 prefixes that represent an endpoint.

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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-f4383bf3e2130711ee7d8fcba6e2f7b69a8dab60111d9aa1203298b8fa85c9d7"></a>

## Next pages — origin_servers.k8s_service.snat_pool.snat_pool / 86d4a514f141 / 5

- [origin_servers.k8s_service.snat_pool](resources--origin_pool--reference--group-002.md#canonical-a1b37e91760eefcb13c5b47a96707eaf7eb332c254fe8c0da41bec0be2669f92)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-84590abf9cd0356d75e014f13d186b60d98fe9a6532904b7b487563c9812986c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b0690ca860040e4852fd34c2f29436c9ec64ebb92c103b4050fedb1fac956874"></a>

## origin_servers.k8s_service.vk8s_networks — origin_servers.k8s_service.vk8s_networks / b986b90fb745 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- [origin_servers.k8s_service](resources--origin_pool--reference--group-002.md#canonical-807eedd9067c3254887005dfeaa77d3a564cb3cacb261d63c3ed9cb49a836cf0)
- origin_servers.k8s_service.vk8s_networks

<a id="canonical-eecf2bd28d2c5c29d426b5660a620aa0359c03e8a185038510282c61a9e23688"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for vk8s networks.

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
vk8s_networks = {}
```

<a id="canonical-4d577c9cd48fc2b52b780275f4eb55d07a560b71a5f9bf0d3a236c90aac25a59"></a>

## Direct properties — origin_servers.k8s_service.vk8s_networks / b986b90fb745 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-87c28c490ba460d088b2d05f8389d0a4f61401b9bc5f11de47314121613570c9"></a>

## Next pages — origin_servers.k8s_service.vk8s_networks / b986b90fb745 / 4

- [origin_servers.k8s_service](resources--origin_pool--reference--group-002.md#canonical-807eedd9067c3254887005dfeaa77d3a564cb3cacb261d63c3ed9cb49a836cf0)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-2c6da3442d0b6ec67494cb089009cadef7abf00026bdc05ae26a0c0b45758b43"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5794cdf5142b18b06be10e55f52105fa56fe77d16fd7415c6a3166c941886ea1"></a>

## origin_servers.private_ip — origin_servers.private_ip / 4066f628c564 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- origin_servers.private_ip

<a id="canonical-1e683ee35788385ef9e0a366ace6cd6aa447b671f350432ef3503bc033caa643"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with private or public IP address and site information.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("inside_network",
    "outside_network"),
  validators.ConflictingObjectAttributes("inside_network",
    "segment"),
  validators.ConflictingObjectAttributes("outside_network",
    "segment")}
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
  "x-ves-oneof-field-network_choice": "[\"inside_network\",\"outside_network\",\"segment\"]",
  "x-ves-oneof-field-private_ip_choice": "[\"ip\"]"
}
```

Terraform syntax:

```terraform
private_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-a63f6555886186794a2ef6be575c34a4358577177afad998262dbffd5f55bbc8"></a>

## Direct properties — origin_servers.private_ip / 4066f628c564 / 3

- [inside_network](resources--origin_pool--reference--group-002.md#canonical-bf3f60ff881a191cbfe7048babc617ff5de1f691fa17d438631d2b19e8ad8860): complete subsection reference.

<a id="canonical-dde746a7871eabf587417e36556e9e9ac730f68e8e7e0d25ab79840fb174c7b9"></a>

<a id="canonical-256d3c68ed615332e50a71b7d9c1626e9eb3d99f63c7941fb77389497f084d22"></a>

## ip property — origin_servers.private_ip / 4066f628c564 / 4

Type: `"string"`. Optional.

IP. Exclusive with \[\] Private IPv4 address.

Upstream description:

Exclusive with \[\] Private IPv4 address.

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

- [outside_network](resources--origin_pool--reference--group-002.md#canonical-bedb2928c1fbf75a50a19779afb07db6b8e28a715c4af1a89801c6b3985ea0dc): complete subsection reference.

- [segment](resources--origin_pool--reference--group-002.md#canonical-fc529d1654d9b7d09b73057885d8377dbb32c7e6b0cc772ae0c235bd003215bc): complete subsection reference.

- [site_locator](resources--origin_pool--reference--group-002.md#canonical-403982637789ddf9be76de1351897d3e6d1395e93595289c72ad4e0a530fc4a0): complete subsection reference.

- [snat_pool](resources--origin_pool--reference--group-002.md#canonical-a3e9c73d075ef6944f06e1449607f229e16310f514d20b75f90b54014a79a046): complete subsection reference.

<a id="canonical-2791bd56dad6130772d7b976b998c48dedae139b0e78c7e69c070d955152aa15"></a>

## Next pages — origin_servers.private_ip / 4066f628c564 / 5

- [origin_servers.private_ip.inside_network](resources--origin_pool--reference--group-002.md#canonical-bf3f60ff881a191cbfe7048babc617ff5de1f691fa17d438631d2b19e8ad8860)
- [origin_servers.private_ip.outside_network](resources--origin_pool--reference--group-002.md#canonical-bedb2928c1fbf75a50a19779afb07db6b8e28a715c4af1a89801c6b3985ea0dc)
- [origin_servers.private_ip.segment](resources--origin_pool--reference--group-002.md#canonical-fc529d1654d9b7d09b73057885d8377dbb32c7e6b0cc772ae0c235bd003215bc)
- [origin_servers.private_ip.site_locator](resources--origin_pool--reference--group-002.md#canonical-403982637789ddf9be76de1351897d3e6d1395e93595289c72ad4e0a530fc4a0)
- [origin_servers.private_ip.snat_pool](resources--origin_pool--reference--group-002.md#canonical-a3e9c73d075ef6944f06e1449607f229e16310f514d20b75f90b54014a79a046)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-bf3f60ff881a191cbfe7048babc617ff5de1f691fa17d438631d2b19e8ad8860"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8fb38a56ed6c1b40aab18997fc3067673bac8908d2b916796f84e71f1d7ce1e0"></a>

## origin_servers.private_ip.inside_network — origin_servers.private_ip.inside_network / a0f1227054fa / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- [origin_servers.private_ip](resources--origin_pool--reference--group-002.md#canonical-2c6da3442d0b6ec67494cb089009cadef7abf00026bdc05ae26a0c0b45758b43)
- origin_servers.private_ip.inside_network

<a id="canonical-44d9e9329af27e38c264066bb9569ff8efa1260bd5a083d90a203a218395954b"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for inside network.

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
inside_network = {}
```

<a id="canonical-f5b8dd42d3f07d867b8b03b0d5bc13f5e228f9173ecd311fa944bf641f646ab8"></a>

## Direct properties — origin_servers.private_ip.inside_network / a0f1227054fa / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fa68a079a9c229712bad00b4706a95ffa6c06274de86d79d6ecf5f48ea81524e"></a>

## Next pages — origin_servers.private_ip.inside_network / a0f1227054fa / 4

- [origin_servers.private_ip](resources--origin_pool--reference--group-002.md#canonical-2c6da3442d0b6ec67494cb089009cadef7abf00026bdc05ae26a0c0b45758b43)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-bedb2928c1fbf75a50a19779afb07db6b8e28a715c4af1a89801c6b3985ea0dc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-65c5e7931aeff536c8f390c1a394d35fcb9b0cffde6e383d2901558b0315ce24"></a>

## origin_servers.private_ip.outside_network — origin_servers.private_ip.outside_network / e7f2d4c7a664 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- [origin_servers.private_ip](resources--origin_pool--reference--group-002.md#canonical-2c6da3442d0b6ec67494cb089009cadef7abf00026bdc05ae26a0c0b45758b43)
- origin_servers.private_ip.outside_network

<a id="canonical-c86d92076fac7c55c1f9d1dfb32578c1d8e6bc4d9644516a291f6b30cc1e74b9"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for outside network.

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
outside_network = {}
```

<a id="canonical-16b9b72721cbeaf6f9a2a4b85e8cda1efeb056e593557014de5e51dc30b2f740"></a>

## Direct properties — origin_servers.private_ip.outside_network / e7f2d4c7a664 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0b002d76989be7f6c78410cb753ca9300e82b788dd450c28928b697ea6281d14"></a>

## Next pages — origin_servers.private_ip.outside_network / e7f2d4c7a664 / 4

- [origin_servers.private_ip](resources--origin_pool--reference--group-002.md#canonical-2c6da3442d0b6ec67494cb089009cadef7abf00026bdc05ae26a0c0b45758b43)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-fc529d1654d9b7d09b73057885d8377dbb32c7e6b0cc772ae0c235bd003215bc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d110eede7e63a912adc08c5dfc0cbd3798edddb0b126a15c8c1ae2bc5c511b72"></a>

## origin_servers.private_ip.segment — origin_servers.private_ip.segment / 6a4472630f23 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- [origin_servers.private_ip](resources--origin_pool--reference--group-002.md#canonical-2c6da3442d0b6ec67494cb089009cadef7abf00026bdc05ae26a0c0b45758b43)
- origin_servers.private_ip.segment

<a id="canonical-1a87e0c6dfebeb8a6494e3e957494237de3635f050c384b854586ed3907667f4"></a>

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
segment {
  # Configure direct properties listed below.
}
```

<a id="canonical-5993295692f7e28d6706af15a8930311b2894870ad2fee7dc04512ae1b384de3"></a>

## Direct properties — origin_servers.private_ip.segment / 6a4472630f23 / 3

<a id="canonical-6b09762dada6a9bbe039141b10c40ff174c5a1ee1c866b0559ff36be65f06f7e"></a>

<a id="canonical-b8157e51f074104316b7f2ea702fb1c45037d42424b4290bd32eaf4ceca3f605"></a>

## name property — origin_servers.private_ip.segment / 6a4472630f23 / 4

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

<a id="canonical-022cb5b3c00a2c719a585a24b0935c5674cdaff1cbd16abcd622e7d0c7b70af1"></a>

<a id="canonical-1f034eb4195dfd03ad928727d46854406b2179d8ac984e5376e9a2f8328b2c99"></a>

## namespace property — origin_servers.private_ip.segment / 6a4472630f23 / 5

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

<a id="canonical-a5b7619cfc4b755e049c2fb5ffea1623e4d63e6ba266b9e024da3f9ada3657b4"></a>

<a id="canonical-efdf52094a3c94a570bb23784100ddf5bfdcb3e93aba0c9feb46f87c0ec8006a"></a>

## tenant property — origin_servers.private_ip.segment / 6a4472630f23 / 6

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

<a id="canonical-f43c138c9bfea1de5f3f38fb2bab1308c95c0fd53a9752cfb524560f18097adf"></a>

## Next pages — origin_servers.private_ip.segment / 6a4472630f23 / 7

- [origin_servers.private_ip](resources--origin_pool--reference--group-002.md#canonical-2c6da3442d0b6ec67494cb089009cadef7abf00026bdc05ae26a0c0b45758b43)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-403982637789ddf9be76de1351897d3e6d1395e93595289c72ad4e0a530fc4a0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0f6268cc0b62e0a7b31e159ac1662ff6e1546da4bed38d621478bd34f6ea48a6"></a>

## origin_servers.private_ip.site_locator — origin_servers.private_ip.site_locator / 02190617d50f / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- [origin_servers.private_ip](resources--origin_pool--reference--group-002.md#canonical-2c6da3442d0b6ec67494cb089009cadef7abf00026bdc05ae26a0c0b45758b43)
- origin_servers.private_ip.site_locator

<a id="canonical-b0b7442b5f15d7ee4a872faf21b86f61b43065a1765293bb469c9fe85b474373"></a>

Type: `"object"`. single nested block, Optional.

Message defines a reference to a site or virtual site object.

Upstream description:

This message defines a reference to a site or virtual site object.

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
site_locator {
  # Configure direct properties listed below.
}
```

<a id="canonical-25255269c5c59bcab684d2be2f9cad2de52fd9d1db0096330c63482db9823451"></a>

## Direct properties — origin_servers.private_ip.site_locator / 02190617d50f / 3

- [site](resources--origin_pool--reference--group-002.md#canonical-0f0e8c8e6fbd7d21697a62ea0b7be283ef817b162434d0c858f682b6721de090): complete subsection reference.

- [virtual_site](resources--origin_pool--reference--group-002.md#canonical-bf70af535672586823296e467b07c921abb8f7686b4e0658bb7a181d6ef76af9): complete subsection reference.

<a id="canonical-f2a2098d7e090188190b5e035ace4bb553d9868784d0fb938d3e485d2cd74967"></a>

## Next pages — origin_servers.private_ip.site_locator / 02190617d50f / 4

- [origin_servers.private_ip.site_locator.site](resources--origin_pool--reference--group-002.md#canonical-0f0e8c8e6fbd7d21697a62ea0b7be283ef817b162434d0c858f682b6721de090)
- [origin_servers.private_ip.site_locator.virtual_site](resources--origin_pool--reference--group-002.md#canonical-bf70af535672586823296e467b07c921abb8f7686b4e0658bb7a181d6ef76af9)
- [origin_servers.private_ip](resources--origin_pool--reference--group-002.md#canonical-2c6da3442d0b6ec67494cb089009cadef7abf00026bdc05ae26a0c0b45758b43)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-0f0e8c8e6fbd7d21697a62ea0b7be283ef817b162434d0c858f682b6721de090"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0eb2d03e144424fcee36a20ce7c771d9e3d544a8dc6f35506d86590ec6200117"></a>

## origin_servers.private_ip.site_locator.site — origin_servers.private_ip.site_locator.site / 9e3371671c49 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- [origin_servers.private_ip](resources--origin_pool--reference--group-002.md#canonical-2c6da3442d0b6ec67494cb089009cadef7abf00026bdc05ae26a0c0b45758b43)
- [origin_servers.private_ip.site_locator](resources--origin_pool--reference--group-002.md#canonical-403982637789ddf9be76de1351897d3e6d1395e93595289c72ad4e0a530fc4a0)
- origin_servers.private_ip.site_locator.site

<a id="canonical-bd22e9cbcd50d5d0e415c28d7470de00712657f0385803ce87e156905f2cda37"></a>

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

<a id="canonical-34f193630900857af9dd3479ae08035f8b359095d25a01b91511cb2735b1a264"></a>

## Direct properties — origin_servers.private_ip.site_locator.site / 9e3371671c49 / 3

<a id="canonical-27a120f7b92510cf16ec0a95f6985c757c0ee44368d97ae57d665ef7d9187c76"></a>

<a id="canonical-c7fc418461a67799d5419d9a018e249031d842927fcb20223ecb43fdee592b34"></a>

## name property — origin_servers.private_ip.site_locator.site / 9e3371671c49 / 4

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

<a id="canonical-34fb467206054e3e9ff9fd01f2f4da665a2e3344d91a8c8805142181fda62422"></a>

<a id="canonical-ff3fbde52d704ddd7e416c88e2935ddd63087fd89574dde0a13480ff62e3e67a"></a>

## namespace property — origin_servers.private_ip.site_locator.site / 9e3371671c49 / 5

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

<a id="canonical-ff706414df112f3eac19c45c8e54f5d89192cec2d948e50c0da3c1824ac8b295"></a>

<a id="canonical-9892e4c81a56d1611e84ac8dd61841e149e7e63e7bb53dacd0069727edaa14de"></a>

## tenant property — origin_servers.private_ip.site_locator.site / 9e3371671c49 / 6

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

<a id="canonical-ab2c3ba8322c1e1b7b6e6b7b2c48709cee880491e1aae9651faeffd2a59b6908"></a>

## Next pages — origin_servers.private_ip.site_locator.site / 9e3371671c49 / 7

- [origin_servers.private_ip.site_locator](resources--origin_pool--reference--group-002.md#canonical-403982637789ddf9be76de1351897d3e6d1395e93595289c72ad4e0a530fc4a0)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-bf70af535672586823296e467b07c921abb8f7686b4e0658bb7a181d6ef76af9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d6ecb77d9243a48937c4f08eaefc17a2d9ded43215d3b69dc2cbaa703dd89eec"></a>

## origin_servers.private_ip.site_locator.virtual_site — origin_servers.private_ip.site_locator.virtual_site / 1a46b25a0597 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- [origin_servers.private_ip](resources--origin_pool--reference--group-002.md#canonical-2c6da3442d0b6ec67494cb089009cadef7abf00026bdc05ae26a0c0b45758b43)
- [origin_servers.private_ip.site_locator](resources--origin_pool--reference--group-002.md#canonical-403982637789ddf9be76de1351897d3e6d1395e93595289c72ad4e0a530fc4a0)
- origin_servers.private_ip.site_locator.virtual_site

<a id="canonical-aab2d25f74faa8dea78cb9e823f2f2dd1a44d1b123ccee54ac26fa3c9925ba6a"></a>

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

<a id="canonical-0588a8751691978878400d4f56a87011138e29d650c261c92ecb4cab3c2b7571"></a>

## Direct properties — origin_servers.private_ip.site_locator.virtual_site / 1a46b25a0597 / 3

<a id="canonical-47908698d7d0f70ba3bc47706e61752c46638713a403904bae4f52cc7de4a185"></a>

<a id="canonical-de525608fbe14424668296791b15996002e873c9a225bbc4088638a5b3048365"></a>

## name property — origin_servers.private_ip.site_locator.virtual_site / 1a46b25a0597 / 4

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

<a id="canonical-a0a086e3cf9f4d8217a683f0532d48e1ce032f4a8174b8ef209f4d401f443299"></a>

<a id="canonical-7856977aba2ccfd820498a8004e6c64626beb516a7641f85223a3f6f1170942c"></a>

## namespace property — origin_servers.private_ip.site_locator.virtual_site / 1a46b25a0597 / 5

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

<a id="canonical-3b200b4ab9c3d42b6974ab6dfa7d940be2e0baac144ad61622d1a6e528d0b6e1"></a>

<a id="canonical-ab068124e2b57a13756afff1fea767c6fe522ead266ad1750de7947bc5d8edcf"></a>

## tenant property — origin_servers.private_ip.site_locator.virtual_site / 1a46b25a0597 / 6

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

<a id="canonical-9ca79da482d6191c692a0307d4166f2ba482ce2e524115f5bb30794caf400be6"></a>

## Next pages — origin_servers.private_ip.site_locator.virtual_site / 1a46b25a0597 / 7

- [origin_servers.private_ip.site_locator](resources--origin_pool--reference--group-002.md#canonical-403982637789ddf9be76de1351897d3e6d1395e93595289c72ad4e0a530fc4a0)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-a3e9c73d075ef6944f06e1449607f229e16310f514d20b75f90b54014a79a046"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d805544e9d9a83d4b91d3d4c3efa86b805d0a4453fabb2c3bb0834ec83f2f45c"></a>

## origin_servers.private_ip.snat_pool — origin_servers.private_ip.snat_pool / 5866bd3cb2a7 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- [origin_servers.private_ip](resources--origin_pool--reference--group-002.md#canonical-2c6da3442d0b6ec67494cb089009cadef7abf00026bdc05ae26a0c0b45758b43)
- origin_servers.private_ip.snat_pool

<a id="canonical-6c29470f1ac3ba3bafcf81989b04dec33d03a1cecd98f3a5240e3837d4007baa"></a>

Type: `"object"`. single nested block, Optional.

SNAT Pool. SNAT Pool configuration.

Upstream description:

SNAT Pool configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_snat_pool",
    "snat_pool")}
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
  "x-ves-oneof-field-snat_pool_choice": "[\"no_snat_pool\",\"snat_pool\"]"
}
```

Terraform syntax:

```terraform
snat_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-d534abdac454512d834cae49255db16f1c5abe0e9bcd050bb6f5ef0369b25d36"></a>

## Direct properties — origin_servers.private_ip.snat_pool / 5866bd3cb2a7 / 3

- [no_snat_pool](resources--origin_pool--reference--group-002.md#canonical-ab6f9d350a4628a2b7f1210aa815a20b4c1d59135a750775406ef96a7585f5f7): complete subsection reference.

- [snat_pool](resources--origin_pool--reference--group-002.md#canonical-2b9e345b77d63daf4c022ef140a9552326208f5418e487e34a617a50fda68f8e): complete subsection reference.

<a id="canonical-f78901866ad99ba75092c62e68f8bb87c2b6efa27887483b3bd5eadd91aa32aa"></a>

## Next pages — origin_servers.private_ip.snat_pool / 5866bd3cb2a7 / 4

- [origin_servers.private_ip.snat_pool.no_snat_pool](resources--origin_pool--reference--group-002.md#canonical-ab6f9d350a4628a2b7f1210aa815a20b4c1d59135a750775406ef96a7585f5f7)
- [origin_servers.private_ip.snat_pool.snat_pool](resources--origin_pool--reference--group-002.md#canonical-2b9e345b77d63daf4c022ef140a9552326208f5418e487e34a617a50fda68f8e)
- [origin_servers.private_ip](resources--origin_pool--reference--group-002.md#canonical-2c6da3442d0b6ec67494cb089009cadef7abf00026bdc05ae26a0c0b45758b43)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-ab6f9d350a4628a2b7f1210aa815a20b4c1d59135a750775406ef96a7585f5f7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-98c75693a5a5cebc7fed1dde91656e81caecd58e18f4a1fb012d8bbb7b1c2944"></a>

## origin_servers.private_ip.snat_pool.no_snat_pool — origin_servers.private_ip.snat_pool.no_snat_pool / 01e7b36675bc / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- [origin_servers.private_ip](resources--origin_pool--reference--group-002.md#canonical-2c6da3442d0b6ec67494cb089009cadef7abf00026bdc05ae26a0c0b45758b43)
- [origin_servers.private_ip.snat_pool](resources--origin_pool--reference--group-002.md#canonical-a3e9c73d075ef6944f06e1449607f229e16310f514d20b75f90b54014a79a046)
- origin_servers.private_ip.snat_pool.no_snat_pool

<a id="canonical-6700f16253d3439c437e941e88b05d685356df542755de6cbc112ad25e7aa9f2"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no snat pool.

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
no_snat_pool = {}
```

<a id="canonical-38042c8cc0f8b20021bee2b5d713a78e9444636b27a5c58d72f70ba0948ad2fd"></a>

## Direct properties — origin_servers.private_ip.snat_pool.no_snat_pool / 01e7b36675bc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6c795b77ba0c83b7d0e68f6aca8567295456cc0ad2c09467c74f96471d5f7901"></a>

## Next pages — origin_servers.private_ip.snat_pool.no_snat_pool / 01e7b36675bc / 4

- [origin_servers.private_ip.snat_pool](resources--origin_pool--reference--group-002.md#canonical-a3e9c73d075ef6944f06e1449607f229e16310f514d20b75f90b54014a79a046)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-2b9e345b77d63daf4c022ef140a9552326208f5418e487e34a617a50fda68f8e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-75f72b48c3f7ac64c9685114769439c4957d0acc8f0e2132c766bdfac7f0b34f"></a>

## origin_servers.private_ip.snat_pool.snat_pool — origin_servers.private_ip.snat_pool.snat_pool / 279902c1f8c2 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- [origin_servers.private_ip](resources--origin_pool--reference--group-002.md#canonical-2c6da3442d0b6ec67494cb089009cadef7abf00026bdc05ae26a0c0b45758b43)
- [origin_servers.private_ip.snat_pool](resources--origin_pool--reference--group-002.md#canonical-a3e9c73d075ef6944f06e1449607f229e16310f514d20b75f90b54014a79a046)
- origin_servers.private_ip.snat_pool.snat_pool

<a id="canonical-0f1080232b6c72ab201fbb30f4412277d08a1979c06b4abfeff77d8615294f23"></a>

Type: `"object"`. single nested block, Optional.

List of IPv4 prefixes that represent an endpoint.

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
snat_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-4b5bd8e5789894b5c4ca3fa2a6c3cfa150122db0153b5f70b040372781df1e69"></a>

## Direct properties — origin_servers.private_ip.snat_pool.snat_pool / 279902c1f8c2 / 3

<a id="canonical-b4f99cdbffb1fb4d25630fd762eeea320e9c5cd4775e1a6cb736a41e9344f783"></a>

<a id="canonical-e378099205925337b8a30b1ffbe090db4cc9c62da5edff912753c3108273a23f"></a>

## prefixes property — origin_servers.private_ip.snat_pool.snat_pool / 279902c1f8c2 / 4

Type: `["list", "string"]`. Optional.

List of IPv4 prefixes that represent an endpoint.

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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-850bbd9603720bdf92f2ea6f9ec86d0f3bcf09dcaa306e8bf6f2b352313e8837"></a>

## Next pages — origin_servers.private_ip.snat_pool.snat_pool / 279902c1f8c2 / 5

- [origin_servers.private_ip.snat_pool](resources--origin_pool--reference--group-002.md#canonical-a3e9c73d075ef6944f06e1449607f229e16310f514d20b75f90b54014a79a046)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-5f29de2edbca1bf266d4175c8611bfdf1235ff2e40f228e3a30fc1e25889003b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-95f4280f8179e4627df62e20cfb6c37d03f34bd432e019284f5febe135ede907"></a>

## origin_servers.private_name — origin_servers.private_name / 86d6f68b3b54 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- origin_servers.private_name

<a id="canonical-cac0931ea8a860ed4ecd5bada698762abb726b9716e0be57b7b242e1c7d368c6"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with private or public DNS name and site information.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("dns_name"),
  validators.ConflictingObjectAttributes("inside_network",
    "outside_network"),
  validators.ConflictingObjectAttributes("inside_network",
    "segment"),
  validators.ConflictingObjectAttributes("outside_network",
    "segment")}
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
  "x-ves-oneof-field-network_choice": "[\"inside_network\",\"outside_network\",\"segment\"]"
}
```

Terraform syntax:

```terraform
private_name {
  # Configure direct properties listed below.
}
```

<a id="canonical-0ed445a3a538a09d1b8d42d95a464b4909b14eec8f27254e9496c6c52765ef44"></a>

## Direct properties — origin_servers.private_name / 86d6f68b3b54 / 3

<a id="canonical-0ce08da9969c53ad44ad88ce0ce37314bc8cf80480ff7d84407f0c7e31bc7bab"></a>

<a id="canonical-a0d9af72d83747a8917d48db3db1886cbf341d73fe6f2dd800d3f2e47e7b8a7b"></a>

## dns_name property — origin_servers.private_name / 86d6f68b3b54 / 4

Type: `"string"`. Optional.

DNS Name. DNS Name

Upstream description:

DNS Name

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$`),
    ""),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "network",
    "characterSet": {
      "allowed": "[a-z0-9.-]",
      "description": "Dot-separated DNS labels"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "formatDescription": "RFC 1123 FQDN: lowercase, dot-separated labels, max 253 chars total",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.95,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
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

- [inside_network](resources--origin_pool--reference--group-002.md#canonical-a36d2fa2be1f0d15ca3eadc2a77d764e5a1c78fe8b226fd0cbbdfc5a35bb4efd): complete subsection reference.

- [outside_network](resources--origin_pool--reference--group-002.md#canonical-d88773cd95172d0dec2ea487ce35270966104bcf0c85da7d1648a8d96dd61258): complete subsection reference.

<a id="canonical-41e47e31ed71fce7467d096222d035ba381f0ccba41d2022f08b81bd07b481ec"></a>

<a id="canonical-1c5a9f330485521b4d015c22c593616c362dd07d84bbb952aeb39efb869bba08"></a>

## refresh_interval property — origin_servers.private_name / 86d6f68b3b54 / 5

Type: `"number"`. Optional.

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Upstream description:

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  validators.Int64RangeSetValidator(
    validators.Int64Range{Minimum: 0, Maximum: 0},
    validators.Int64Range{Minimum: 10, Maximum: 604800},
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
    "maximum": 604800,
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
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  }
}
```

- [segment](resources--origin_pool--reference--group-002.md#canonical-60ea7301c863cc557b1dd21f578bfbe240ab28ec6b84e041fbabdded489ebd15): complete subsection reference.

- [site_locator](resources--origin_pool--reference--group-002.md#canonical-26ceb0dc7ee963e75d00a0f6a258f7ab1a32bd90ecc760578cdc16c56ee837ea): complete subsection reference.

- [snat_pool](resources--origin_pool--reference--group-002.md#canonical-395b2554be38a6f0ebb7d2f98b1ae82855b7de5cbf1d39b0b477e32b962dde23): complete subsection reference.

<a id="canonical-fd113ed33034e9e850c3fd161428552e8d3a9895d50e109982e125cd507a6dc1"></a>

## Next pages — origin_servers.private_name / 86d6f68b3b54 / 6

- [origin_servers.private_name.inside_network](resources--origin_pool--reference--group-002.md#canonical-a36d2fa2be1f0d15ca3eadc2a77d764e5a1c78fe8b226fd0cbbdfc5a35bb4efd)
- [origin_servers.private_name.outside_network](resources--origin_pool--reference--group-002.md#canonical-d88773cd95172d0dec2ea487ce35270966104bcf0c85da7d1648a8d96dd61258)
- [origin_servers.private_name.segment](resources--origin_pool--reference--group-002.md#canonical-60ea7301c863cc557b1dd21f578bfbe240ab28ec6b84e041fbabdded489ebd15)
- [origin_servers.private_name.site_locator](resources--origin_pool--reference--group-002.md#canonical-26ceb0dc7ee963e75d00a0f6a258f7ab1a32bd90ecc760578cdc16c56ee837ea)
- [origin_servers.private_name.snat_pool](resources--origin_pool--reference--group-002.md#canonical-395b2554be38a6f0ebb7d2f98b1ae82855b7de5cbf1d39b0b477e32b962dde23)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-a36d2fa2be1f0d15ca3eadc2a77d764e5a1c78fe8b226fd0cbbdfc5a35bb4efd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ec491bb58a14288fd316d3b148e4d49c8f00ffc72d1af4eff48e124075f0b4e5"></a>

## origin_servers.private_name.inside_network — origin_servers.private_name.inside_network / 0374eededd04 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- [origin_servers.private_name](resources--origin_pool--reference--group-002.md#canonical-5f29de2edbca1bf266d4175c8611bfdf1235ff2e40f228e3a30fc1e25889003b)
- origin_servers.private_name.inside_network

<a id="canonical-bad88409213fccb90ae3551ca3c95d1c9656693e0c182e2088a3fb4a102099dd"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for inside network.

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
inside_network = {}
```

<a id="canonical-c7a865fd977393e3da1b346b238bdbe4b02d943b51fbbde45bb3ce10c546b962"></a>

## Direct properties — origin_servers.private_name.inside_network / 0374eededd04 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-86a268dafb83942c324a6a9e4665233577e152db625fe67f915a3a612b41762c"></a>

## Next pages — origin_servers.private_name.inside_network / 0374eededd04 / 4

- [origin_servers.private_name](resources--origin_pool--reference--group-002.md#canonical-5f29de2edbca1bf266d4175c8611bfdf1235ff2e40f228e3a30fc1e25889003b)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-d88773cd95172d0dec2ea487ce35270966104bcf0c85da7d1648a8d96dd61258"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-59fdac83462d498d8c8143d8e31d901276b587cc6a6dfa46a7cdac9c35a10ff6"></a>

## origin_servers.private_name.outside_network — origin_servers.private_name.outside_network / 9d3e61698f15 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- [origin_servers.private_name](resources--origin_pool--reference--group-002.md#canonical-5f29de2edbca1bf266d4175c8611bfdf1235ff2e40f228e3a30fc1e25889003b)
- origin_servers.private_name.outside_network

<a id="canonical-07b4373dd86cb4c9b876536febf5a449936d669c88bc1bc007edc3f285c5046d"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for outside network.

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
outside_network = {}
```

<a id="canonical-fb326ab75181b7e687ce815c1b155fecb086bf453da741b987296b94a1a82425"></a>

## Direct properties — origin_servers.private_name.outside_network / 9d3e61698f15 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0478877335512fbf569235496718c504e40bd1dee6e4b41eabae67b8c954c732"></a>

## Next pages — origin_servers.private_name.outside_network / 9d3e61698f15 / 4

- [origin_servers.private_name](resources--origin_pool--reference--group-002.md#canonical-5f29de2edbca1bf266d4175c8611bfdf1235ff2e40f228e3a30fc1e25889003b)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-60ea7301c863cc557b1dd21f578bfbe240ab28ec6b84e041fbabdded489ebd15"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a110519946708e656c431728f73b55e9c758f568e830425a47f0849be6899584"></a>

## origin_servers.private_name.segment — origin_servers.private_name.segment / 29342b9d447c / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- [origin_servers.private_name](resources--origin_pool--reference--group-002.md#canonical-5f29de2edbca1bf266d4175c8611bfdf1235ff2e40f228e3a30fc1e25889003b)
- origin_servers.private_name.segment

<a id="canonical-5f44b1ed95084cbb9e57465f510de16e746e12e6ff21e85a8007f48320e11787"></a>

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
segment {
  # Configure direct properties listed below.
}
```

<a id="canonical-c534527c8669db169eff339ef838cf7c0217956df6b4b457088ada6a14ae7956"></a>

## Direct properties — origin_servers.private_name.segment / 29342b9d447c / 3

<a id="canonical-dcc30c1a89676e67a30c75fc76e9b9769ef317bf03a1d30e44d4b85e335c1278"></a>

<a id="canonical-313f4173f7a8691181f716b35d8bc2285307d5e35897c97900e7f17c6fb62a0d"></a>

## name property — origin_servers.private_name.segment / 29342b9d447c / 4

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

<a id="canonical-a78ba70e724860151c1573d43e0dec12c7e152391b887893a012aa8277cde738"></a>

<a id="canonical-83f721d529d4a017917ee55f98420f69739cd0ff4046d9736b27d293ba4ef5b2"></a>

## namespace property — origin_servers.private_name.segment / 29342b9d447c / 5

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

<a id="canonical-087b31e7008bad871cad30e1b82a54de451014d5e49f76893b03c5be2cd2529a"></a>

<a id="canonical-2057ac622a72c3d4addb38cb1635054d185105aecf6be13b87d627ab7454ce8f"></a>

## tenant property — origin_servers.private_name.segment / 29342b9d447c / 6

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

<a id="canonical-5eda4c76abc9755dd0885b393f0e0872a47f444a30782e317ae3468f57019fd9"></a>

## Next pages — origin_servers.private_name.segment / 29342b9d447c / 7

- [origin_servers.private_name](resources--origin_pool--reference--group-002.md#canonical-5f29de2edbca1bf266d4175c8611bfdf1235ff2e40f228e3a30fc1e25889003b)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-26ceb0dc7ee963e75d00a0f6a258f7ab1a32bd90ecc760578cdc16c56ee837ea"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7ab5dc025405d59c23c69a6ec44e25bc1d3d2810be47afedf94bcdb13cb09cc1"></a>

## origin_servers.private_name.site_locator — origin_servers.private_name.site_locator / 154490aa12f1 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- [origin_servers.private_name](resources--origin_pool--reference--group-002.md#canonical-5f29de2edbca1bf266d4175c8611bfdf1235ff2e40f228e3a30fc1e25889003b)
- origin_servers.private_name.site_locator

<a id="canonical-48068257e132a5c4ef6f21d48f7063e93c689f5057cb6b94b637ac9bd78d3332"></a>

Type: `"object"`. single nested block, Optional.

Message defines a reference to a site or virtual site object.

Upstream description:

This message defines a reference to a site or virtual site object.

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
site_locator {
  # Configure direct properties listed below.
}
```

<a id="canonical-6f5929f88c8dab402f9e106a43bb17d23941e9ed75bd057eb430ba3ad44cf52d"></a>

## Direct properties — origin_servers.private_name.site_locator / 154490aa12f1 / 3

- [site](resources--origin_pool--reference--group-002.md#canonical-e213514862369e17da311a6ca04ba9509bf22bafc16d0e1fd63807282f0b449c): complete subsection reference.

- [virtual_site](resources--origin_pool--reference--group-002.md#canonical-346117a43ae59ce69b14ee6f36b6ce5fadcba1b281f2bab32ba7971dcb51d112): complete subsection reference.

<a id="canonical-c51aba407f2b33f2b1bf996c8b36cf78c380a261fa56d2c2fc013008c412da32"></a>

## Next pages — origin_servers.private_name.site_locator / 154490aa12f1 / 4

- [origin_servers.private_name.site_locator.site](resources--origin_pool--reference--group-002.md#canonical-e213514862369e17da311a6ca04ba9509bf22bafc16d0e1fd63807282f0b449c)
- [origin_servers.private_name.site_locator.virtual_site](resources--origin_pool--reference--group-002.md#canonical-346117a43ae59ce69b14ee6f36b6ce5fadcba1b281f2bab32ba7971dcb51d112)
- [origin_servers.private_name](resources--origin_pool--reference--group-002.md#canonical-5f29de2edbca1bf266d4175c8611bfdf1235ff2e40f228e3a30fc1e25889003b)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-e213514862369e17da311a6ca04ba9509bf22bafc16d0e1fd63807282f0b449c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-43f3b99bb6144909a042d6b69c09b5d4475209ca73119913d10e3e00c79864de"></a>

## origin_servers.private_name.site_locator.site — origin_servers.private_name.site_locator.site / af957a1a64d3 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- [origin_servers.private_name](resources--origin_pool--reference--group-002.md#canonical-5f29de2edbca1bf266d4175c8611bfdf1235ff2e40f228e3a30fc1e25889003b)
- [origin_servers.private_name.site_locator](resources--origin_pool--reference--group-002.md#canonical-26ceb0dc7ee963e75d00a0f6a258f7ab1a32bd90ecc760578cdc16c56ee837ea)
- origin_servers.private_name.site_locator.site

<a id="canonical-b453632d3b5064536ddde51bfa0d20dd78a02c0824e1a949ef893b7036680baf"></a>

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

<a id="canonical-5af56d97bebf9a2a0df25607564e1de762fd833ffe16e82dcd64d66cdd66ed64"></a>

## Direct properties — origin_servers.private_name.site_locator.site / af957a1a64d3 / 3

<a id="canonical-6ffbc1239345816cf22b5942dd2796114c1543cde154db1bfa4e80920b4d7a3b"></a>

<a id="canonical-5dece9e0a5f68a118ce7d7ac7b2e45bd59e9562cbb032ed2f8159355a7a33507"></a>

## name property — origin_servers.private_name.site_locator.site / af957a1a64d3 / 4

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

<a id="canonical-8c52c68519dadf8713ce2aff2258ef3315c306420827072cf276718772a756b6"></a>

<a id="canonical-3df65d6d42f11c810b2e54ea2d9a30a0e58162a653be5787e35a3915e6ce4d72"></a>

## namespace property — origin_servers.private_name.site_locator.site / af957a1a64d3 / 5

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

<a id="canonical-dad70ef825aa2825d2f02e32ffd33fa766897faaeaf6ec2faeec0eeb4fe474cb"></a>

<a id="canonical-4ad4d1aed98451ffff58d9e39bbf2c11259d7b214e2795d8c79aca10b05d9dcb"></a>

## tenant property — origin_servers.private_name.site_locator.site / af957a1a64d3 / 6

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

<a id="canonical-bacd7785ff0809ad3c580e722e02e4266752dbf3ab63504b0eaffd8931db63dc"></a>

## Next pages — origin_servers.private_name.site_locator.site / af957a1a64d3 / 7

- [origin_servers.private_name.site_locator](resources--origin_pool--reference--group-002.md#canonical-26ceb0dc7ee963e75d00a0f6a258f7ab1a32bd90ecc760578cdc16c56ee837ea)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-346117a43ae59ce69b14ee6f36b6ce5fadcba1b281f2bab32ba7971dcb51d112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d635a37f323e1802e4bdfb60c616acfe3d90468f42e6f80c34f5b86f18e73c27"></a>

## origin_servers.private_name.site_locator.virtual_site — origin_servers.private_name.site_locator.virtual_site / 17f50feb83ff / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- [origin_servers.private_name](resources--origin_pool--reference--group-002.md#canonical-5f29de2edbca1bf266d4175c8611bfdf1235ff2e40f228e3a30fc1e25889003b)
- [origin_servers.private_name.site_locator](resources--origin_pool--reference--group-002.md#canonical-26ceb0dc7ee963e75d00a0f6a258f7ab1a32bd90ecc760578cdc16c56ee837ea)
- origin_servers.private_name.site_locator.virtual_site

<a id="canonical-348a7140d5a4fc4900d6c0c1df538ffebece08085f08fad45754eff833ae9d9a"></a>

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

<a id="canonical-c71a149b4f49b33348ede7dc501dd4f2dd72e131c06e9123cbbe55e91a921706"></a>

## Direct properties — origin_servers.private_name.site_locator.virtual_site / 17f50feb83ff / 3

<a id="canonical-ec235a5ff88ebe29c3978ffd105211e3077825e74f0a81200ee73e1cc38809a2"></a>

<a id="canonical-3454c65a4b1c0d616927e0f41e4bba07762bbefae1d4f1b2d2b7c25c42e67a04"></a>

## name property — origin_servers.private_name.site_locator.virtual_site / 17f50feb83ff / 4

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

<a id="canonical-65897b90931e6191eb37f2ac88a9e6bc47fe44f69f59ab8a2d7298f4713a3e2a"></a>

<a id="canonical-bfff12fb3a6fe7c6b53a4b123a6dcf08eb1b18ea0574342ff5ecaa5f8b15d049"></a>

## namespace property — origin_servers.private_name.site_locator.virtual_site / 17f50feb83ff / 5

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

<a id="canonical-ff63d2601a69a5a94fbb57358104d6467b459829cdaefe7a48ba29ff0e375ac7"></a>

<a id="canonical-a860e22d880dd45c107d27f8ea562ffd75e241698a1832d4cac285f215277161"></a>

## tenant property — origin_servers.private_name.site_locator.virtual_site / 17f50feb83ff / 6

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

<a id="canonical-9b8a626a2d42641faf9294136628c989e2e2539fd3aeeb9088277abe86f00b17"></a>

## Next pages — origin_servers.private_name.site_locator.virtual_site / 17f50feb83ff / 7

- [origin_servers.private_name.site_locator](resources--origin_pool--reference--group-002.md#canonical-26ceb0dc7ee963e75d00a0f6a258f7ab1a32bd90ecc760578cdc16c56ee837ea)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-395b2554be38a6f0ebb7d2f98b1ae82855b7de5cbf1d39b0b477e32b962dde23"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-125e5e0b80a87696be2e8c7dcb4eab6edd977692203e7c516325769cbbffdac2"></a>

## origin_servers.private_name.snat_pool — origin_servers.private_name.snat_pool / 9d46d9bd77df / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- [origin_servers.private_name](resources--origin_pool--reference--group-002.md#canonical-5f29de2edbca1bf266d4175c8611bfdf1235ff2e40f228e3a30fc1e25889003b)
- origin_servers.private_name.snat_pool

<a id="canonical-46ea5444232de51039fd93f759d4f94c067b19c65f703b180872d1fbd761511f"></a>

Type: `"object"`. single nested block, Optional.

SNAT Pool. SNAT Pool configuration.

Upstream description:

SNAT Pool configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_snat_pool",
    "snat_pool")}
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
  "x-ves-oneof-field-snat_pool_choice": "[\"no_snat_pool\",\"snat_pool\"]"
}
```

Terraform syntax:

```terraform
snat_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-c5fca792bff6740c4f3c8110785febc74aa56508758dd857a7993cabf31cdacc"></a>

## Direct properties — origin_servers.private_name.snat_pool / 9d46d9bd77df / 3

- [no_snat_pool](resources--origin_pool--reference--group-002.md#canonical-dc2e92e4c875101ebd8dfb936c2f4c52aba1481c28816632f479fd6fe058bd96): complete subsection reference.

- [snat_pool](resources--origin_pool--reference--group-002.md#canonical-a977d25f5dbcda909b50a1c143f19f01810b42d455468b0efea17828dfcf6089): complete subsection reference.

<a id="canonical-3d3f7f07c4038f720d98f4c3601637d63d8038c3daedeb2e109483a9cc37f8b7"></a>

## Next pages — origin_servers.private_name.snat_pool / 9d46d9bd77df / 4

- [origin_servers.private_name.snat_pool.no_snat_pool](resources--origin_pool--reference--group-002.md#canonical-dc2e92e4c875101ebd8dfb936c2f4c52aba1481c28816632f479fd6fe058bd96)
- [origin_servers.private_name.snat_pool.snat_pool](resources--origin_pool--reference--group-002.md#canonical-a977d25f5dbcda909b50a1c143f19f01810b42d455468b0efea17828dfcf6089)
- [origin_servers.private_name](resources--origin_pool--reference--group-002.md#canonical-5f29de2edbca1bf266d4175c8611bfdf1235ff2e40f228e3a30fc1e25889003b)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-dc2e92e4c875101ebd8dfb936c2f4c52aba1481c28816632f479fd6fe058bd96"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-69031abfae3a51f891686c307b7fad3ce1f51a220224076f70b6c45276b0bdfa"></a>

## origin_servers.private_name.snat_pool.no_snat_pool — origin_servers.private_name.snat_pool.no_snat_pool / 97878df29fba / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- [origin_servers.private_name](resources--origin_pool--reference--group-002.md#canonical-5f29de2edbca1bf266d4175c8611bfdf1235ff2e40f228e3a30fc1e25889003b)
- [origin_servers.private_name.snat_pool](resources--origin_pool--reference--group-002.md#canonical-395b2554be38a6f0ebb7d2f98b1ae82855b7de5cbf1d39b0b477e32b962dde23)
- origin_servers.private_name.snat_pool.no_snat_pool

<a id="canonical-178f6d27948ecb0ba6979e87924b25adc765d40398cc373b2ff3b3ba2b7507d9"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no snat pool.

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
no_snat_pool = {}
```

<a id="canonical-676bbf6037661e2e20691d8d36b9e36a77d0517487ece48917abafe68452b4f9"></a>

## Direct properties — origin_servers.private_name.snat_pool.no_snat_pool / 97878df29fba / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-243a93a5cb196dbc28ac5f7d92fd6883948838433e3ff5b2c01fc9524006ba8d"></a>

## Next pages — origin_servers.private_name.snat_pool.no_snat_pool / 97878df29fba / 4

- [origin_servers.private_name.snat_pool](resources--origin_pool--reference--group-002.md#canonical-395b2554be38a6f0ebb7d2f98b1ae82855b7de5cbf1d39b0b477e32b962dde23)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-a977d25f5dbcda909b50a1c143f19f01810b42d455468b0efea17828dfcf6089"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0708391b3a7752507215dcaa8fc5972041bc4f71ec55e11c5d8ae3dd3b67acaa"></a>

## origin_servers.private_name.snat_pool.snat_pool — origin_servers.private_name.snat_pool.snat_pool / b13c097b14c2 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- [origin_servers.private_name](resources--origin_pool--reference--group-002.md#canonical-5f29de2edbca1bf266d4175c8611bfdf1235ff2e40f228e3a30fc1e25889003b)
- [origin_servers.private_name.snat_pool](resources--origin_pool--reference--group-002.md#canonical-395b2554be38a6f0ebb7d2f98b1ae82855b7de5cbf1d39b0b477e32b962dde23)
- origin_servers.private_name.snat_pool.snat_pool

<a id="canonical-4629448c3ad00f4559b859ca2b593ef693a00eead7de02d6d8075b570d997f01"></a>

Type: `"object"`. single nested block, Optional.

List of IPv4 prefixes that represent an endpoint.

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
snat_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-5bea05972415fb83f9ec5e6be22c2451972b34afb786da6c6702b4fd0827b338"></a>

## Direct properties — origin_servers.private_name.snat_pool.snat_pool / b13c097b14c2 / 3

<a id="canonical-0a943bd15bda44430af3858613ecfa8c6ab8d76a48a0d95e0ff0e468be8020e8"></a>

<a id="canonical-33f9a52e27eb966386b63ae2e42950e8798923a93223d712cb03114f5cd55a3c"></a>

## prefixes property — origin_servers.private_name.snat_pool.snat_pool / b13c097b14c2 / 4

Type: `["list", "string"]`. Optional.

List of IPv4 prefixes that represent an endpoint.

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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-cd9605ebddc61bd18758655df3600fe176a4d27edf5c6dc36a74f076509c5b38"></a>

## Next pages — origin_servers.private_name.snat_pool.snat_pool / b13c097b14c2 / 5

- [origin_servers.private_name.snat_pool](resources--origin_pool--reference--group-002.md#canonical-395b2554be38a6f0ebb7d2f98b1ae82855b7de5cbf1d39b0b477e32b962dde23)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-86892a0b8abf53409d722e2d600702358eeb6ed794eb16cede78c0fd203677c6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-104025d5be542f58430a225396e499e01e88813b0982484e715fcb390eaef2b2"></a>

## origin_servers.public_ip — origin_servers.public_ip / cd06b8d49a5f / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- origin_servers.public_ip

<a id="canonical-bc493ef6b56befe695eb3e4962b5b785ace5850f7b092275fa7a27eddc3c7590"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with public IP address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-public_ip_choice": "[\"ip\"]"
}
```

Terraform syntax:

```terraform
public_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-4daa32787a6a8e88a417b8e022438b178e3488dc57ad030c70e06077f69fc21a"></a>

## Direct properties — origin_servers.public_ip / cd06b8d49a5f / 3

<a id="canonical-315ece931a77f943a0a7e78202ee0e183a2d95fe8cf067cf1d4bf14c69c91b9e"></a>

<a id="canonical-1a7e3611045486e79373fdcd513f57dfb1925fa892a1a485f15b026f2519800f"></a>

## ip property — origin_servers.public_ip / cd06b8d49a5f / 4

Type: `"string"`. Optional.

Public IPv4. Exclusive with \[\] Public IPv4 address.

Upstream description:

Exclusive with \[\] Public IPv4 address.

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

<a id="canonical-21354007df7d007812160bc5850e2883477616796f918bdc49671ed85d1645b5"></a>

## Next pages — origin_servers.public_ip / cd06b8d49a5f / 5

- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-116ce7a9a539106140d695639f56717a3f212c77d02f1eac26ab9efaa3734e7b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5de681207203827c52a148a64ca906028b37055565eb16a35a9bdda9157cb45e"></a>

## origin_servers.public_name — origin_servers.public_name / d29e75624d6b / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- origin_servers.public_name

<a id="canonical-2ac9f93881a630c4cac32514c0fd0313b7c46513a69187474672a6b7e2d8b5bc"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with public DNS name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("dns_name")}
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
public_name {
  # Configure direct properties listed below.
}
```

<a id="canonical-3382e78098de52b42348cf179dccd9d555a580234a8909a1eddc880662b97e60"></a>

## Direct properties — origin_servers.public_name / d29e75624d6b / 3

<a id="canonical-956b384aaf91f37038cd01af534efd898eb6cd3958c1e351823f9b966fd167ee"></a>

<a id="canonical-9b5b98e3803c22da6f3fcfc1677ce3882a8badcb683158956fac7f2b58e06a9c"></a>

## dns_name property — origin_servers.public_name / d29e75624d6b / 4

Type: `"string"`. Optional.

DNS Name. DNS Name

Upstream description:

DNS Name

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
    "characterSet": {
      "allowed": "[a-z0-9.-]",
      "description": "Dot-separated DNS labels"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "formatDescription": "RFC 1123 FQDN: lowercase, dot-separated labels, max 253 chars total",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
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

<a id="canonical-c94bbdb691fcbe68524a58d9dfc318fbe7d53071999ca46b4b571e644510a22a"></a>

<a id="canonical-4304d55d403b58ad7b4a60762c5e51b3c08b6e159bb92303739dff78f64481c1"></a>

## refresh_interval property — origin_servers.public_name / d29e75624d6b / 5

Type: `"number"`. Optional.

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Upstream description:

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  validators.Int64RangeSetValidator(
    validators.Int64Range{Minimum: 0, Maximum: 0},
    validators.Int64Range{Minimum: 10, Maximum: 604800},
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
    "maximum": 604800,
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
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  }
}
```

<a id="canonical-5a55dbcae7f5af802606d91e2566c9115de61c25381cd4c6b5507ec5f93658df"></a>

## Next pages — origin_servers.public_name / d29e75624d6b / 6

- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-9c6aea85b0e4c5054d34baa041360d879809b176747d1ec0c14e7d1b8beb9df7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e919e88b75d917f4d4af7c8619dc2d3f51f48a216d87e53e8120a24fe9ad0de5"></a>

## origin_servers.vn_private_ip — origin_servers.vn_private_ip / 82f1c9d88431 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- origin_servers.vn_private_ip

<a id="canonical-66c53a132a7f5352b43080f4c539c11c3f5d98ca5983171459fdfb679acfa1aa"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with IP on Virtual Network.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-virtual_network_ip_choice": "[\"ip\"]"
}
```

Terraform syntax:

```terraform
vn_private_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-8596901fe77d639f86e83378f50f6ddd6287f84392bccca0e5586b97879295d5"></a>

## Direct properties — origin_servers.vn_private_ip / 82f1c9d88431 / 3

<a id="canonical-33347e5af6084f36f080e2d54a9859d7ae8b669e1dcb1a01d18856e036c472cb"></a>

<a id="canonical-f4c9a8bf23da464473bc32ec43960de2773f8b21207ff060de2a096ef7635c9c"></a>

## ip property — origin_servers.vn_private_ip / 82f1c9d88431 / 4

Type: `"string"`. Optional.

IPv4. Exclusive with \[\] IPv4 address.

Upstream description:

Exclusive with \[\] IPv4 address.

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

- [virtual_network](resources--origin_pool--reference--group-002.md#canonical-da40c77a17125339e3fdc81f55d2308e08c8880604e8c5bef4de7fa5ed0be1bc): complete subsection reference.

<a id="canonical-f60b7d706a1d4d6ab1415c2c62030d5941387ab8efc3f760f197a16dc25eca03"></a>

## Next pages — origin_servers.vn_private_ip / 82f1c9d88431 / 5

- [origin_servers.vn_private_ip.virtual_network](resources--origin_pool--reference--group-002.md#canonical-da40c77a17125339e3fdc81f55d2308e08c8880604e8c5bef4de7fa5ed0be1bc)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-da40c77a17125339e3fdc81f55d2308e08c8880604e8c5bef4de7fa5ed0be1bc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ca967a58e2d47d3cdad078851e096be5decbc38984c5b3e9d93faba18b1bd95e"></a>

## origin_servers.vn_private_ip.virtual_network — origin_servers.vn_private_ip.virtual_network / d20d2c68f2ec / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- [origin_servers.vn_private_ip](resources--origin_pool--reference--group-002.md#canonical-9c6aea85b0e4c5054d34baa041360d879809b176747d1ec0c14e7d1b8beb9df7)
- origin_servers.vn_private_ip.virtual_network

<a id="canonical-039b0ea2ea020688e310705175ee00e092efd4df2cdf28f0f2970db3e2fb5c5c"></a>

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
virtual_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-a5b65f298e357254679d02434fbf563e7893bc42aec555dd7ea63d8c500dfd94"></a>

## Direct properties — origin_servers.vn_private_ip.virtual_network / d20d2c68f2ec / 3

<a id="canonical-008b25925619102ff21f64000ba4a844215fc82183853ea47c318b118290eebd"></a>

<a id="canonical-34eb686f80fbcceab630d9b645fd38151efa6a75259c04a4294218ed4aa17a38"></a>

## name property — origin_servers.vn_private_ip.virtual_network / d20d2c68f2ec / 4

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

<a id="canonical-2bad1689a77e515554cf53e6a29b48466e6b93b50a81f0b86357c43afe36fe18"></a>

<a id="canonical-2e40600320eec686dffc85af9cea51036e9529432f1731fb4d6bd03ab84b7c0b"></a>

## namespace property — origin_servers.vn_private_ip.virtual_network / d20d2c68f2ec / 5

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

<a id="canonical-15eb1e303a855ee36633106025d09ac6db8593ba8c13ececadef427d18edfe9f"></a>

<a id="canonical-3eeaa5a91ada5a92f87e681363acd1f7b6964ec946f9ba83cdb03f3101bda54d"></a>

## tenant property — origin_servers.vn_private_ip.virtual_network / d20d2c68f2ec / 6

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

<a id="canonical-54f2361e38c2915f10abed09d31c492587967373ecfb6c52b2e750deaf0c4a44"></a>

## Next pages — origin_servers.vn_private_ip.virtual_network / d20d2c68f2ec / 7

- [origin_servers.vn_private_ip](resources--origin_pool--reference--group-002.md#canonical-9c6aea85b0e4c5054d34baa041360d879809b176747d1ec0c14e7d1b8beb9df7)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-4790c620301dfed425ab1a0395b5baee9c0e9940049d0b38fa5a5ed992615e56"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c346eec8c1230c1f4bb0f6c8e6d465b088b0ee81184144cbb8e08f5467c4ee10"></a>

## origin_servers.vn_private_name — origin_servers.vn_private_name / 9a3f43b893f9 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- origin_servers.vn_private_name

<a id="canonical-5ff6ebc37855c1b55bab1719486a4cb87a505a36d4d079305c8b0d237e5df01d"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with DNS name on Virtual Network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("dns_name")}
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
vn_private_name {
  # Configure direct properties listed below.
}
```

<a id="canonical-0e70c1f4a3de1fe8098e39508fec4af5929ea4a57627d9af766f3e2c88137c08"></a>

## Direct properties — origin_servers.vn_private_name / 9a3f43b893f9 / 3

<a id="canonical-c230a39ef5c4565ea779a73a220a341021d716124800b64488748640b2a60f80"></a>

<a id="canonical-7cf27ef1d5b5ba24332a2b0fb6a053ecc1cecfb110049450ea00057fdf77350f"></a>

## dns_name property — origin_servers.vn_private_name / 9a3f43b893f9 / 4

Type: `"string"`. Optional.

DNS Name. DNS Name

Upstream description:

DNS Name

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$`),
    ""),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "network",
    "characterSet": {
      "allowed": "[a-z0-9.-]",
      "description": "Dot-separated DNS labels"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "formatDescription": "RFC 1123 FQDN: lowercase, dot-separated labels, max 253 chars total",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.95,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
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

- [private_network](resources--origin_pool--reference--group-002.md#canonical-f78bc6743e651275089952fc8428743e1d1054bc70345e38eb79a054e89e6600): complete subsection reference.

<a id="canonical-a3f5b71099c2ca733e0f25c1b7f4121ff9774bfa75229ed81e9a87fa9c7dbbbb"></a>

## Next pages — origin_servers.vn_private_name / 9a3f43b893f9 / 5

- [origin_servers.vn_private_name.private_network](resources--origin_pool--reference--group-002.md#canonical-f78bc6743e651275089952fc8428743e1d1054bc70345e38eb79a054e89e6600)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-f78bc6743e651275089952fc8428743e1d1054bc70345e38eb79a054e89e6600"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e47056de69feb5e6e34c46650c36f2e63a333f968909727b91a690c2cc1e5830"></a>

## origin_servers.vn_private_name.private_network — origin_servers.vn_private_name.private_network / 1a0a50b07602 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- [origin_servers.vn_private_name](resources--origin_pool--reference--group-002.md#canonical-4790c620301dfed425ab1a0395b5baee9c0e9940049d0b38fa5a5ed992615e56)
- origin_servers.vn_private_name.private_network

<a id="canonical-5fcf92e681c0d64a45c65b72012593c9bbf458930aa5d2426b9406e602a684ec"></a>

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
private_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-06c7dc8a73700062e04c80aa38afa8f338878247aaa60598a015d79888951efb"></a>

## Direct properties — origin_servers.vn_private_name.private_network / 1a0a50b07602 / 3

<a id="canonical-251847ef6c4773eda3f378812c9c185b0babfed849c69511dc476fc40aee398c"></a>

<a id="canonical-33e754bf6c5259c1dbba69e4bf6496af6bae38fe3c068b1cf5cdf3e1071ab5c9"></a>

## name property — origin_servers.vn_private_name.private_network / 1a0a50b07602 / 4

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

<a id="canonical-c13596a18905004c6b7398499384af1f0d322c5283f8cd4c47c4236c3f6f6e30"></a>

<a id="canonical-9ea36d139648a2bbd7f8b81fcd2bd85056731df2176e5daa663883a26fe7eb17"></a>

## namespace property — origin_servers.vn_private_name.private_network / 1a0a50b07602 / 5

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

<a id="canonical-e38c952d151a6c7f2d8b92cd202f98de9aee065c62ba4ba696ed7ac2eab26bb8"></a>

<a id="canonical-8fbe0efb067674ed365d1482471858c92e32cf7934e3304b8bb6be47414f2bfe"></a>

## tenant property — origin_servers.vn_private_name.private_network / 1a0a50b07602 / 6

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

<a id="canonical-d224e486824d13f9d0aecca01895e698e7f9ee056085ad8fec80db876926db9a"></a>

## Next pages — origin_servers.vn_private_name.private_network / 1a0a50b07602 / 7

- [origin_servers.vn_private_name](resources--origin_pool--reference--group-002.md#canonical-4790c620301dfed425ab1a0395b5baee9c0e9940049d0b38fa5a5ed992615e56)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-d4a1eba2ebd5b3aefffcc5c49f834145c4fab929d7024ff96905398f2406158f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e354de66881b428f990b31692bcba463b5b3ea91bcc7b4fcfc4edea14ea3abd2"></a>

## same_as_endpoint_port — same_as_endpoint_port / 1581f4d41079 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- same_as_endpoint_port

<a id="canonical-83612ed238eaabdb14d2ae66732f0b85ad5ce12a3af85da13e3e876582fa0658"></a>

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
same_as_endpoint_port = {}
```

<a id="canonical-12433de8cdc78a505c1e966f85e065b569473b9420c0f5316b69a12660abd04f"></a>

## Direct properties — same_as_endpoint_port / 1581f4d41079 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5d4870fb1774b9e5968f64181d07ed1c502c8d0ffd460641277fd46a6652c622"></a>

## Next pages — same_as_endpoint_port / 1581f4d41079 / 4

- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-7ae7ce0a89d4f66924dedb8dfdf24bdabfc1c4ceecaf7e928bf3c2815469bcce"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ea3c4634161e795c7572550c41d92fa6643d21e152f811a7659dd80f251d20ce"></a>

## timeouts — timeouts / 897210e6f558 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- timeouts

<a id="canonical-00b59cfc8389916908c08b416bc90e9669ff9c56fd9a0f8395f82fc5e547c57d"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-4fed19033c35c36d687ce49cfc8b2aeec5af59abfb1b3216b1990727dbdc1d04"></a>

## Direct properties — timeouts / 897210e6f558 / 3

<a id="canonical-3a25e579d6cc337c1c535600bb5a5dd98645d0edb4f17d3438cb8485a978edde"></a>

<a id="canonical-20d1c656737f5f0fd8c4cb3bdb18e1c922b91066c9e50840e046a121832aaba9"></a>

## create property — timeouts / 897210e6f558 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-5b7a596926eb1935c1123ba766458e40aad2d71e6d2e73030d7ceba3e628c57b"></a>

<a id="canonical-1219a70ba294e3efc50245cfd5ac9841d96e573bde9786516d64ce3d7add8bdd"></a>

## delete property — timeouts / 897210e6f558 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-164ad8c4429ef13ac7e81f756085680fe07432b196580ebaae7d1d932fd9621b"></a>

<a id="canonical-a5ab979637fce449578426e91d3336f7d4131a22c772516d3c52444709d36853"></a>

## read property — timeouts / 897210e6f558 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-14dd3f317d46bac30ac3cc2d3b8f5150fa07e802f41a3467a15bce1d35fccbd8"></a>

<a id="canonical-14db15886be426abd132a9c26ca87a6c4a8fbc88651304c05f8e3de71696387b"></a>

## update property — timeouts / 897210e6f558 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-072eb2a89b486f2a7356ea7c5e99a7b90fcb46aee2dbe16152bc3699072e275b"></a>

## Next pages — timeouts / 897210e6f558 / 8

- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-94f24e1eba33680896ec00dce0eb25ca3c9a2d809ad8bfb56c911d0bc8905f38"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ec9bf5285f97e39ce4f1250b834cda05caa72d8f563b407ffe9b73be43e9305b"></a>

## upstream_conn_pool_reuse_type — upstream_conn_pool_reuse_type / 0a591b3622a9 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- upstream_conn_pool_reuse_type

<a id="canonical-885fb983433829a1ec57ecde5e41de6772c9483c958d2f1dcf40b58091de1622"></a>

Type: `"object"`. single nested block, Optional.

Select upstream connection pool reuse state for every downstream connection. This configuration
choice is for HTTP(S) LB only.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_conn_pool_reuse",
    "enable_conn_pool_reuse")}
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
  "x-ves-oneof-field-map_downstream_to_upstream_conn_pool_type": "[\"disable_conn_pool_reuse\",\"enable_conn_pool_reuse\"]"
}
```

Terraform syntax:

```terraform
upstream_conn_pool_reuse_type {
  # Configure direct properties listed below.
}
```

<a id="canonical-11c22c77a24916778f76637345f6b127c69c496ad066ea98bddb4dcd141adb51"></a>

## Direct properties — upstream_conn_pool_reuse_type / 0a591b3622a9 / 3

- [disable_conn_pool_reuse](resources--origin_pool--reference--group-002.md#canonical-3fa3b165c26d2278c043348761092a50d904a5b83c67cc60489fe96b367c1abb): complete subsection reference.

- [enable_conn_pool_reuse](resources--origin_pool--reference--group-002.md#canonical-dda954c14a5e4c4183a3ae7591c6c7c4c1deafc64fef2cead779befc197f7631): complete subsection reference.

<a id="canonical-b587c9d37519d1bc12e735fd5c2bff0ab7a1ee18b4bd8ee9eced70eef9fee964"></a>

## Next pages — upstream_conn_pool_reuse_type / 0a591b3622a9 / 4

- [upstream_conn_pool_reuse_type.disable_conn_pool_reuse](resources--origin_pool--reference--group-002.md#canonical-3fa3b165c26d2278c043348761092a50d904a5b83c67cc60489fe96b367c1abb)
- [upstream_conn_pool_reuse_type.enable_conn_pool_reuse](resources--origin_pool--reference--group-002.md#canonical-dda954c14a5e4c4183a3ae7591c6c7c4c1deafc64fef2cead779befc197f7631)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-3fa3b165c26d2278c043348761092a50d904a5b83c67cc60489fe96b367c1abb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b2002f1206be68e8e9857b28a8b2f4e44bbbe5ac203b22d9cd9663f642d9909e"></a>

## upstream_conn_pool_reuse_type.disable_conn_pool_reuse — upstream_conn_pool_reuse_type.disable_conn_pool_reuse / 5cb52a424040 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [upstream_conn_pool_reuse_type](resources--origin_pool--reference--group-002.md#canonical-94f24e1eba33680896ec00dce0eb25ca3c9a2d809ad8bfb56c911d0bc8905f38)
- upstream_conn_pool_reuse_type.disable_conn_pool_reuse

<a id="canonical-72fac9903e12f1d1f04b145b02857d415c3de6822142310971f0fcc9c0a62d4c"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable conn pool reuse.

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
disable_conn_pool_reuse = {}
```

<a id="canonical-c1794aece506b67b7b3309698138fa33853001822fea028d249acdce45a25509"></a>

## Direct properties — upstream_conn_pool_reuse_type.disable_conn_pool_reuse / 5cb52a424040 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-288c74705fe10bb74ed2ad01e9867df2364e5a85d411efbf8747a6134fd564cf"></a>

## Next pages — upstream_conn_pool_reuse_type.disable_conn_pool_reuse / 5cb52a424040 / 4

- [upstream_conn_pool_reuse_type](resources--origin_pool--reference--group-002.md#canonical-94f24e1eba33680896ec00dce0eb25ca3c9a2d809ad8bfb56c911d0bc8905f38)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-dda954c14a5e4c4183a3ae7591c6c7c4c1deafc64fef2cead779befc197f7631"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-54b4bcd5d689198396633a92a89b6dae7f8ee4290197a05e62c85828546310bc"></a>

## upstream_conn_pool_reuse_type.enable_conn_pool_reuse — upstream_conn_pool_reuse_type.enable_conn_pool_reuse / 6b570acca0e4 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [upstream_conn_pool_reuse_type](resources--origin_pool--reference--group-002.md#canonical-94f24e1eba33680896ec00dce0eb25ca3c9a2d809ad8bfb56c911d0bc8905f38)
- upstream_conn_pool_reuse_type.enable_conn_pool_reuse

<a id="canonical-914460952669ffb6a52ec16dbbdd4fe3d5431886f236e055408d05b20f63e8d1"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable conn pool reuse.

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
enable_conn_pool_reuse = {}
```

<a id="canonical-3354de7daa0c21172bfbcbc031f218c035867cd1e2d3668a3f826ffb47dbb3de"></a>

## Direct properties — upstream_conn_pool_reuse_type.enable_conn_pool_reuse / 6b570acca0e4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-83947483c7235ddb539386be7c2a56dc68e6a47440cf05afec851f6a6de66a3c"></a>

## Next pages — upstream_conn_pool_reuse_type.enable_conn_pool_reuse / 6b570acca0e4 / 4

- [upstream_conn_pool_reuse_type](resources--origin_pool--reference--group-002.md#canonical-94f24e1eba33680896ec00dce0eb25ca3c9a2d809ad8bfb56c911d0bc8905f38)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-18c27b6d6857cf831ee951a3437c0ab831adbcbcb92c8fd70efb90f76fba4fdf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0f57443944e27581aaefd60065addc80fcc1c1c9d88214b407351aea635473f0"></a>

## use_tls — use_tls / 6f46ee02d5f8 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- use_tls

<a id="canonical-da6a539edede34014702388d9e0abd5ba500bd242fc3a4a40f407fdf4ae46046"></a>

Type: `"object"`. single nested block, Optional.

TLS Parameters for Origin Servers. Upstream TLS Parameters.

Upstream description:

Upstream TLS Parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_session_key_caching",
    "disable_session_key_caching"),
  validators.ConflictingObjectAttributes("default_session_key_caching",
    "max_session_keys"),
  validators.ConflictingObjectAttributes("disable_session_key_caching",
    "max_session_keys"),
  validators.ConflictingObjectAttributes("disable_sni",
    "sni"),
  validators.ConflictingObjectAttributes("disable_sni",
    "use_host_header_as_sni"),
  validators.ConflictingObjectAttributes("no_mtls",
    "use_mtls"),
  validators.ConflictingObjectAttributes("no_mtls",
    "use_mtls_obj"),
  validators.ConflictingObjectAttributes("skip_server_verification",
    "use_server_verification"),
  validators.ConflictingObjectAttributes("skip_server_verification",
    "volterra_trusted_ca"),
  validators.ConflictingObjectAttributes("sni",
    "use_host_header_as_sni"),
  validators.ConflictingObjectAttributes("use_mtls",
    "use_mtls_obj"),
  validators.ConflictingObjectAttributes("use_server_verification",
    "volterra_trusted_ca")}
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
  "x-ves-oneof-field-max_session_keys_type": "[\"default_session_key_caching\",\"disable_session_key_caching\",\"max_session_keys\"]",
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\",\"use_mtls_obj\"]",
  "x-ves-oneof-field-server_validation_choice": "[\"skip_server_verification\",\"use_server_verification\",\"volterra_trusted_ca\"]",
  "x-ves-oneof-field-sni_choice": "[\"disable_sni\",\"sni\",\"use_host_header_as_sni\"]"
}
```

Terraform syntax:

```terraform
use_tls {
  # Configure direct properties listed below.
}
```

<a id="canonical-13fdaf15051ef9b9594b0bababdb9b0750b34ab83073fc20106a366e3b94097d"></a>

## Direct properties — use_tls / 6f46ee02d5f8 / 3

- [default_session_key_caching](resources--origin_pool--reference--group-003.md#canonical-db57aadc97b42dba89563240a91f6af47cc17e420b4d7893a0e0b8c69095750e): complete subsection reference.

- [disable_session_key_caching](resources--origin_pool--reference--group-003.md#canonical-2d518a92dd087f658f61307220cf744714465756ac60518a3f10c735c3896556): complete subsection reference.

- [disable_sni](resources--origin_pool--reference--group-003.md#canonical-b5b7d33a1c24c0132e9e622b828b15c3ae409345b1aa93ba56b2dd9883dc17d3): complete subsection reference.

<a id="canonical-44fb1c637240b30606d9859d2b11546d31bd40ce512bc43cbd045707a0e71d28"></a>

<a id="canonical-ea3f84995884f53b6a6e44d2c738bc258aed393431c5e341432a85092f7f2013"></a>

## max_session_keys property — use_tls / 6f46ee02d5f8 / 4

Type: `"number"`. Optional.

Exclusive with \[default\_session\_key\_caching disable\_session\_key\_caching\] Number of session
keys that are cached.

Upstream description:

Exclusive with \[default\_session\_key\_caching disable\_session\_key\_caching\]

Number of session keys that are cached.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(2, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 2
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "64"
  }
}
```

- [no_mtls](resources--origin_pool--reference--group-003.md#canonical-48de757ca201189218824316bbd23f0667a07e3ace00bac6187f161d9aa35529): complete subsection reference.

- [skip_server_verification](resources--origin_pool--reference--group-003.md#canonical-12f882e6573470126ee15f0a7306222599f694c38b2a7952bb364bb049223f61): complete subsection reference.

<a id="canonical-68805d3a775008d51de2db1476c038b589d9fc6063322a25af75808f74a65fa0"></a>

<a id="canonical-b941230e34e58ee616a4b07100b53a889ac4f29a27610d9c381128c44e2042e9"></a>

## sni property — use_tls / 6f46ee02d5f8 / 5

Type: `"string"`. Optional.

Exclusive with \[disable\_sni use\_host\_header\_as\_sni\] SNI value to be used.

Upstream description:

Exclusive with \[disable\_sni use\_host\_header\_as\_sni\] SNI value to be used.

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

- [tls_config](resources--origin_pool--reference--group-003.md#canonical-56a9eaa2d7444586b13cd43e728c690562529568863e6fa1625bdcc3e182ea5d): complete subsection reference.

- [use_host_header_as_sni](resources--origin_pool--reference--group-003.md#canonical-f89b3835729201ab6bf3817b0e742a65db04a6aabcaff540a54ee8954ee9d28a): complete subsection reference.

- [use_mtls](resources--origin_pool--reference--group-003.md#canonical-58aaa3c6b66d808f551ad41b0ac3978d46b44a6122f283547e690bfcea5d4845): complete subsection reference.

- [use_mtls_obj](resources--origin_pool--reference--group-003.md#canonical-50025ee6e2504d3ccf7005acea544e797f91c7565591c23faad18c3469f04194): complete subsection reference.

- [use_server_verification](resources--origin_pool--reference--group-003.md#canonical-c0f1fd7fb6a8c6bd28b606ef7816026802d315812d220663bc1a5539c0c74402): complete subsection reference.

- [volterra_trusted_ca](resources--origin_pool--reference--group-003.md#canonical-3320a35d9af7db62bdfa410461c6c31b0acb87e9b009de36d0b36da595233754): complete subsection reference.
