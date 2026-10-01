---
page_title: "xcsh_route reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_route reference."
---

# xcsh_route reference

<a id="canonical-39ddf071f404b7b188726a97d60a9e775e004148d3c0c4d1da27baa0f2170c8b"></a>

## Next pages — routes.waf_type / 0de7aa022c42 / 4

- [routes.waf_type.app_firewall](data-sources--route--reference--group-003.md#canonical-88030fa2b72a81bf0ddc2942630f991c9f90521d8af7e716da3c29f0c8d2f4fe)
- [routes.waf_type.disable_waf](data-sources--route--reference--group-003.md#canonical-7884b5c1d9635487ccd9a56662ae6550d3826a3e6bf08605f814f7179377afcb)
- [routes.waf_type.inherit_waf](data-sources--route--reference--group-003.md#canonical-91e3ce8ef5fd47a5eb45350c087bb01614958033ebdf213ea5329a739dec83a3)
- [routes](data-sources--route--reference--group-001.md#canonical-19922725b52fed5452e47b4b9e3033774938f366da9529fca3df5d26e12a8f49)
- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)

<a id="canonical-88030fa2b72a81bf0ddc2942630f991c9f90521d8af7e716da3c29f0c8d2f4fe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-25aff3836a9dd4f10001f03f495aca0c8ee53af2b5b75d4e972f16d680f476ba"></a>

## routes.waf_type.app_firewall — routes.waf_type.app_firewall / 32dbf267693d / 2

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)
- [Property reference](data-sources--route--reference--group-001.md#canonical-a6c77d69b7834e989d425e83281601a25261f73228caae96eeda7871db19dfbf)
- [routes](data-sources--route--reference--group-001.md#canonical-19922725b52fed5452e47b4b9e3033774938f366da9529fca3df5d26e12a8f49)
- [routes.waf_type](data-sources--route--reference--group-002.md#canonical-7362ef3ec8de94401ce424467749d400c77686ad5c2a78680790aca10481fc58)
- routes.waf_type.app_firewall

<a id="canonical-08b7da3bcdf4f42058fd372a654de5fdc4e1bed93787ada684356d253c7344fc"></a>

Type: `"single"`. Computed.

List of references to the app\_firewall configuration objects.

Upstream description:

A list of references to the app\_firewall configuration objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-dd4d79e56adea97d31862775fad3bcfe87642d8bafe6c26ab81dfd1f53a73a00"></a>

## Direct properties — routes.waf_type.app_firewall / 32dbf267693d / 3

- [app_firewall](data-sources--route--reference--group-003.md#canonical-5ddd5d6ed8d1c12cf65bb804c53febd00ef422e51146f014188e46782a347b41): complete subsection reference.

<a id="canonical-c270e0ef908515124eda5f9824c5dae53c44b1b1739646a5e67d425f75ad94f1"></a>

## Next pages — routes.waf_type.app_firewall / 32dbf267693d / 4

- [routes.waf_type.app_firewall.app_firewall](data-sources--route--reference--group-003.md#canonical-5ddd5d6ed8d1c12cf65bb804c53febd00ef422e51146f014188e46782a347b41)
- [routes.waf_type](data-sources--route--reference--group-002.md#canonical-7362ef3ec8de94401ce424467749d400c77686ad5c2a78680790aca10481fc58)
- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)

<a id="canonical-5ddd5d6ed8d1c12cf65bb804c53febd00ef422e51146f014188e46782a347b41"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9aa86ca846ac8b7b56c82ec399760e33950f1fcb5c7290c1878bf8f00e069414"></a>

## routes.waf_type.app_firewall.app_firewall — routes.waf_type.app_firewall.app_firewall / d6e80cc472eb / 2

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)
- [Property reference](data-sources--route--reference--group-001.md#canonical-a6c77d69b7834e989d425e83281601a25261f73228caae96eeda7871db19dfbf)
- [routes](data-sources--route--reference--group-001.md#canonical-19922725b52fed5452e47b4b9e3033774938f366da9529fca3df5d26e12a8f49)
- [routes.waf_type](data-sources--route--reference--group-002.md#canonical-7362ef3ec8de94401ce424467749d400c77686ad5c2a78680790aca10481fc58)
- [routes.waf_type.app_firewall](data-sources--route--reference--group-003.md#canonical-88030fa2b72a81bf0ddc2942630f991c9f90521d8af7e716da3c29f0c8d2f4fe)
- routes.waf_type.app_firewall.app_firewall

<a id="canonical-ff7eb447a7555ded0623d37c2b379b9df32a081d3c87e2c65006a3a2a2f052fd"></a>

Type: `"list"`. Computed.

References to an Application Firewall configuration object.

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
    "ves.io.schema.rules.repeated.num_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.num_items": "1"
  }
}
```

<a id="canonical-1b4c9e527e2a155aad760e8acdc800613b36ff9a9547968f79d0454456c604e9"></a>

## Direct properties — routes.waf_type.app_firewall.app_firewall / d6e80cc472eb / 3

<a id="canonical-999fc33da0f95410d9775965c47cb3bf51c04a20eefd199d72b713bb5d354b69"></a>

<a id="canonical-aedc81ac921d11f9202e2728ec197e4f93f4603afbae8769140c7594c5a96bd5"></a>

## kind property — routes.waf_type.app_firewall.app_firewall / d6e80cc472eb / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-d4626025295332003c78381472dd925c45f1802c996bf11549e74ea358fc2ce7"></a>

<a id="canonical-74aaf1173be6a0f62a81f3d4ab96356bd91226c9f3be057a2e25b8b01a3cd2b8"></a>

## name property — routes.waf_type.app_firewall.app_firewall / d6e80cc472eb / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-969aa455b217b7fe4cf0500ef68fba943ab363c17f3949505eebb7189ecace14"></a>

<a id="canonical-b1988c52fbe409fed66c40a7e5867fe5b6f1485e1023d36d79c8083d7c20078b"></a>

## namespace property — routes.waf_type.app_firewall.app_firewall / d6e80cc472eb / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-2e71e4a96cc0fb91702c2a70ebc10f3e2c221ed83db894495226ad79e56be2b6"></a>

<a id="canonical-8800ac809950e0c0d60c12eedad879059757c90b376957f08c5502ac6e92f303"></a>

## tenant property — routes.waf_type.app_firewall.app_firewall / d6e80cc472eb / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-b0c9bf5328dfb73ba4b0ec27dbac1e0e670f3d595e2d5729a06e2ff7b8c0fc42"></a>

<a id="canonical-8c79479dd914418ce7f9322fbb920d44d92053fa899f0b81201906749bfaffb7"></a>

## uid property — routes.waf_type.app_firewall.app_firewall / d6e80cc472eb / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-f2fe593a5e26d21539b03240cdb52291fc4f27825a562244b7b90061823a0e4d"></a>

## Next pages — routes.waf_type.app_firewall.app_firewall / d6e80cc472eb / 9

- [routes.waf_type.app_firewall](data-sources--route--reference--group-003.md#canonical-88030fa2b72a81bf0ddc2942630f991c9f90521d8af7e716da3c29f0c8d2f4fe)
- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)

<a id="canonical-7884b5c1d9635487ccd9a56662ae6550d3826a3e6bf08605f814f7179377afcb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-191fd353f1d9e3a21d2ecffeb7f11d6c899149237cccef239762380f16f18f77"></a>

## routes.waf_type.disable_waf — routes.waf_type.disable_waf / 3a8781e65a52 / 2

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)
- [Property reference](data-sources--route--reference--group-001.md#canonical-a6c77d69b7834e989d425e83281601a25261f73228caae96eeda7871db19dfbf)
- [routes](data-sources--route--reference--group-001.md#canonical-19922725b52fed5452e47b4b9e3033774938f366da9529fca3df5d26e12a8f49)
- [routes.waf_type](data-sources--route--reference--group-002.md#canonical-7362ef3ec8de94401ce424467749d400c77686ad5c2a78680790aca10481fc58)
- routes.waf_type.disable_waf

<a id="canonical-75c7cb34fa00beeb2daba6031f0e4eb54c14d719a6d3ae1fe3fcc857772c3469"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable waf.

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

<a id="canonical-99622f774c55130f4e4da2744e67de79420909cdcdfe1d3208fcd52557a8687f"></a>

## Direct properties — routes.waf_type.disable_waf / 3a8781e65a52 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0826c266b80a2c110a947033e3de450ffc3b74b7ff88fb958f81e3dab995e445"></a>

## Next pages — routes.waf_type.disable_waf / 3a8781e65a52 / 4

- [routes.waf_type](data-sources--route--reference--group-002.md#canonical-7362ef3ec8de94401ce424467749d400c77686ad5c2a78680790aca10481fc58)
- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)

<a id="canonical-91e3ce8ef5fd47a5eb45350c087bb01614958033ebdf213ea5329a739dec83a3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0fa9f566fe844f7c1d5dd305fe3de09d83ad108b11192824a94e106201328aa1"></a>

## routes.waf_type.inherit_waf — routes.waf_type.inherit_waf / 78cb7112ca48 / 2

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)
- [Property reference](data-sources--route--reference--group-001.md#canonical-a6c77d69b7834e989d425e83281601a25261f73228caae96eeda7871db19dfbf)
- [routes](data-sources--route--reference--group-001.md#canonical-19922725b52fed5452e47b4b9e3033774938f366da9529fca3df5d26e12a8f49)
- [routes.waf_type](data-sources--route--reference--group-002.md#canonical-7362ef3ec8de94401ce424467749d400c77686ad5c2a78680790aca10481fc58)
- routes.waf_type.inherit_waf

<a id="canonical-17ddb98cd621b57d1bfd18867c4b38ee45bdbfe3bcdc848418aa814027ba6f10"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for inherit waf.

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

<a id="canonical-d96f513a5a2256653a1dabfb5a9b4a0ebffbd892b4eb5a33a52164a0d8549a5c"></a>

## Direct properties — routes.waf_type.inherit_waf / 78cb7112ca48 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-18f1be7973ce5a1ed3c65558a6336581c9742b177ee8b69e5f77c141a77ef48b"></a>

## Next pages — routes.waf_type.inherit_waf / 78cb7112ca48 / 4

- [routes.waf_type](data-sources--route--reference--group-002.md#canonical-7362ef3ec8de94401ce424467749d400c77686ad5c2a78680790aca10481fc58)
- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)
