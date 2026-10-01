---
page_title: "xcsh_network_firewall reference"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_network_firewall reference."
---

# xcsh_network_firewall reference

<a id="canonical-3788cba45b5e266537cf6289bb39431b1592d6bdb952ff53cafe1b93ac1711d3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-43b88e7239a64baeb437e64edac22e2dec8b6067883de03114e2cb6bf3163aa1"></a>

## Property reference — Property reference / dc3c10cb1258 / 2

Breadcrumbs:

- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-9b62b76cf7e299087b5f4928cbe980ffbde2f654430a3985cbcf52e5b4f58a56)
- Property reference

<a id="canonical-c56e72ee78ead89fd9e3825b34cf889b647980029e27389fb5e43899d445e293"></a>

## Direct properties — Property reference / dc3c10cb1258 / 3

- [active_enhanced_firewall_policies](data-sources--network_firewall--reference--group-001.md#canonical-c1db3a9936cb42c65172736c3fc1ae492099d197ab2210144fa69b0d58707bcc): complete subsection reference.

- [active_fast_acls](data-sources--network_firewall--reference--group-001.md#canonical-d4354aa962ed2150fd0c3f1333ec9704d2988dad097399766a9954a5d1c8c012): complete subsection reference.

- [active_forward_proxy_policies](data-sources--network_firewall--reference--group-001.md#canonical-afd575d646ce1b00dd26dfce3151ed9e0095d30d344c02b5f6206458da435264): complete subsection reference.

- [active_network_policies](data-sources--network_firewall--reference--group-001.md#canonical-305e9030fb3795b1b3800369e84572869e50f0adfb0d495b02a3110ffa45955b): complete subsection reference.

<a id="canonical-7bb229d650231c9206cadcc688d27aeb0a5f072c2a1b623c37d1c9a43d1c9bc6"></a>

<a id="canonical-34b541c491b1d26cd3e1d929a143f43d1f3e186b00da644da863b8972a3794f2"></a>

## annotations property — Property reference / dc3c10cb1258 / 4

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

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

<a id="canonical-60cc4f098840260bbfa5d536c517a749838b6e9bc6e9b3c28f74aab1f4bf07fa"></a>

<a id="canonical-848b96ecb765563f75140cdf9b3506d73e22a33904c2bc95ac5210ec18c6b997"></a>

## description property — Property reference / dc3c10cb1258 / 5

Type: `"string"`. Computed.

Description of the NetworkFirewall.

Upstream description:

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

- [disable_fast_acl](data-sources--network_firewall--reference--group-001.md#canonical-564dd5ff59989a71d5bb181e2f97423323974e1f4bc126c34c630ece5eb89aee): complete subsection reference.

- [disable_forward_proxy_policy](data-sources--network_firewall--reference--group-001.md#canonical-1009a248731ea6605183fff5b12b8b466028a679527532ab20e3b1878d05b5f7): complete subsection reference.

- [disable_network_policy](data-sources--network_firewall--reference--group-001.md#canonical-13109d6287ba0d99ceaf95506b396c1388015b320689d328da9935261769de4b): complete subsection reference.

<a id="canonical-c05fa4a679223a1b9c19906e7c28e5c6dbb2055cf1224b11f2a0b928481fa9f1"></a>

<a id="canonical-2678b74c28181b0d6e0f0dce55c8ebdc39090fafcd9ed3c9b6efc3e8ca479325"></a>

## id property — Property reference / dc3c10cb1258 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-2355c877ebb9de9625f75300eaccb48040ed4aaf79494e680e7df2a62e360005"></a>

<a id="canonical-0e0cb39880035bb7bd8c36f0102b85c4700fd439580c07e6f4866363c0b4a485"></a>

## labels property — Property reference / dc3c10cb1258 / 7

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

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

<a id="canonical-f4ddcd8dae0302c1de6b1bd49d8e2c89cf25bcf718d13bc00039639586f526de"></a>

<a id="canonical-bd7c69b871061f3946d138835c24ca2eccd0757ddb68bbfa28b87809fabd8001"></a>

## name property — Property reference / dc3c10cb1258 / 8

Type: `"string"`. Required.

Name of the NetworkFirewall.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

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

<a id="canonical-507cbaaed414983405fd508fa8321e38c5b1ca4bc1a19e1747ef28f228376b9a"></a>

<a id="canonical-0ddc2347a1ea590d7858e828f708e9c8b07fc46246893f67fd4e89504bcfe43d"></a>

## namespace property — Property reference / dc3c10cb1258 / 9

Type: `"string"`. Optional, Computed.

Namespace where the NetworkFirewall exists.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

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

<a id="canonical-55d89e40f69762d4a3548b9d7fbb85b8631935c3e3952a0872a79b8a1a4a665c"></a>

## All schema paths — Property reference / dc3c10cb1258 / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `active_enhanced_firewall_policies` | [active_enhanced_firewall_policies](data-sources--network_firewall--reference--group-001.md#canonical-c4066807d33c540280fcfa7b1b01ac56a71da9e7e9438fd036c144a9542a69e3) |
| `active_enhanced_firewall_policies.enhanced_firewall_policies` | [active_enhanced_firewall_policies.enhanced_firewall_policies](data-sources--network_firewall--reference--group-001.md#canonical-8c6bcc3e47d8c281c5654d5e47ab43065d3a5986ae4b14988de5303458d1e78f) |
| `active_enhanced_firewall_policies.enhanced_firewall_policies.name` | [active_enhanced_firewall_policies.enhanced_firewall_policies.name](data-sources--network_firewall--reference--group-001.md#canonical-f501111bd596e470b7316e92e0e82fe11f7c46877cbfcfc526718886ac036a7c) |
| `active_enhanced_firewall_policies.enhanced_firewall_policies.namespace` | [active_enhanced_firewall_policies.enhanced_firewall_policies.namespace](data-sources--network_firewall--reference--group-001.md#canonical-3442d49a1089d282fb53c1cb774a50a668d1caf1eca3871413041f722f8b6156) |
| `active_enhanced_firewall_policies.enhanced_firewall_policies.tenant` | [active_enhanced_firewall_policies.enhanced_firewall_policies.tenant](data-sources--network_firewall--reference--group-001.md#canonical-eb1b60f760d17bace423a3900d4a7aaf8383960d616ae3108dc6f76802aa0479) |
| `active_fast_acls` | [active_fast_acls](data-sources--network_firewall--reference--group-001.md#canonical-53fb941de7de8fbc7fcd8445f36d1a55a8094088fa492ecbbeab4670a5be4cb9) |
| `active_fast_acls.fast_acls` | [active_fast_acls.fast_acls](data-sources--network_firewall--reference--group-001.md#canonical-60b847aeb6e7d623d23f8b7a97fbc77af547046ce6c1463f86fed70534589afe) |
| `active_fast_acls.fast_acls.name` | [active_fast_acls.fast_acls.name](data-sources--network_firewall--reference--group-001.md#canonical-644dbb85ee0efe754f4d57cd2613b8ed64a11d480602b1107b75f8a2fba1bd33) |
| `active_fast_acls.fast_acls.namespace` | [active_fast_acls.fast_acls.namespace](data-sources--network_firewall--reference--group-001.md#canonical-6905ad1eac09cd02f7b642ba84238fcc411f91431485317d6e833218e8726916) |
| `active_fast_acls.fast_acls.tenant` | [active_fast_acls.fast_acls.tenant](data-sources--network_firewall--reference--group-001.md#canonical-3a15215bd51962428b32b6fe9502b3f9733700cb94e3662e4e7c57a6fc8e2b1a) |
| `active_forward_proxy_policies` | [active_forward_proxy_policies](data-sources--network_firewall--reference--group-001.md#canonical-09cda56c4701ff9cbecfecfb0e749f5f4f584a4f74bf3e54ca99e67d216392f5) |
| `active_forward_proxy_policies.forward_proxy_policies` | [active_forward_proxy_policies.forward_proxy_policies](data-sources--network_firewall--reference--group-001.md#canonical-0eea5ad4402949090632dce3c370b75325452e80ba6444752f61a28c95673302) |
| `active_forward_proxy_policies.forward_proxy_policies.name` | [active_forward_proxy_policies.forward_proxy_policies.name](data-sources--network_firewall--reference--group-001.md#canonical-d72bcb080b31a806f8dc04f6e1c4770e99951ef8942e587165916281b571ccef) |
| `active_forward_proxy_policies.forward_proxy_policies.namespace` | [active_forward_proxy_policies.forward_proxy_policies.namespace](data-sources--network_firewall--reference--group-001.md#canonical-168759d8a73776755d12c9c9327facdc604d2d6b204108f09f4eff7db529bcdd) |
| `active_forward_proxy_policies.forward_proxy_policies.tenant` | [active_forward_proxy_policies.forward_proxy_policies.tenant](data-sources--network_firewall--reference--group-001.md#canonical-f1a91a78f05de2246634d8024be6ea200e0427b78ac49e3e630a3c168857364e) |
| `active_network_policies` | [active_network_policies](data-sources--network_firewall--reference--group-001.md#canonical-b628c71c77dd61693a019895ebf5e44714be80532a65441eedfe87f5aa695818) |
| `active_network_policies.network_policies` | [active_network_policies.network_policies](data-sources--network_firewall--reference--group-001.md#canonical-cc9d95ea056917dde22181974a91d0fa7f3a4677cf3447af26e1df9e894a5ffb) |
| `active_network_policies.network_policies.name` | [active_network_policies.network_policies.name](data-sources--network_firewall--reference--group-001.md#canonical-b6c6880e43ab7d721146f40c5a9ebb7ec196272613fe5436be6b4e316091136e) |
| `active_network_policies.network_policies.namespace` | [active_network_policies.network_policies.namespace](data-sources--network_firewall--reference--group-001.md#canonical-f7651208b70f112a90e2a42a5c02651447246d68e2f300199e7110bc2fcd7c1c) |
| `active_network_policies.network_policies.tenant` | [active_network_policies.network_policies.tenant](data-sources--network_firewall--reference--group-001.md#canonical-ffc54afbc144e20e9902c06650adc9095ba5ccd592c4294d7ae34fe365109ba4) |
| `annotations` | [annotations](data-sources--network_firewall--reference--group-001.md#canonical-7bb229d650231c9206cadcc688d27aeb0a5f072c2a1b623c37d1c9a43d1c9bc6) |
| `description` | [description](data-sources--network_firewall--reference--group-001.md#canonical-60cc4f098840260bbfa5d536c517a749838b6e9bc6e9b3c28f74aab1f4bf07fa) |
| `disable_fast_acl` | [disable_fast_acl](data-sources--network_firewall--reference--group-001.md#canonical-0cf0a26aa833fe6092d0eb000dfaf254f79ee86878b935d21b7ced27a29c65bc) |
| `disable_forward_proxy_policy` | [disable_forward_proxy_policy](data-sources--network_firewall--reference--group-001.md#canonical-cdbb308e1853445eb0041820c490f1981a8d4ad8eee0e546049e023a110ec2b5) |
| `disable_network_policy` | [disable_network_policy](data-sources--network_firewall--reference--group-001.md#canonical-f2112a5d1ae43ac487b21224e2c126a99dcb7bbee8ad4d717b322e9b4d9b8812) |
| `id` | [id](data-sources--network_firewall--reference--group-001.md#canonical-c05fa4a679223a1b9c19906e7c28e5c6dbb2055cf1224b11f2a0b928481fa9f1) |
| `labels` | [labels](data-sources--network_firewall--reference--group-001.md#canonical-2355c877ebb9de9625f75300eaccb48040ed4aaf79494e680e7df2a62e360005) |
| `name` | [name](data-sources--network_firewall--reference--group-001.md#canonical-f4ddcd8dae0302c1de6b1bd49d8e2c89cf25bcf718d13bc00039639586f526de) |
| `namespace` | [namespace](data-sources--network_firewall--reference--group-001.md#canonical-507cbaaed414983405fd508fa8321e38c5b1ca4bc1a19e1747ef28f228376b9a) |

<a id="canonical-d032c49345b421df71cc47c653af350d75db91c19b9ff12fb7c44c25c2fc9695"></a>

## Next pages — Property reference / dc3c10cb1258 / 11

- [active_enhanced_firewall_policies](data-sources--network_firewall--reference--group-001.md#canonical-c1db3a9936cb42c65172736c3fc1ae492099d197ab2210144fa69b0d58707bcc)
- [active_fast_acls](data-sources--network_firewall--reference--group-001.md#canonical-d4354aa962ed2150fd0c3f1333ec9704d2988dad097399766a9954a5d1c8c012)
- [active_forward_proxy_policies](data-sources--network_firewall--reference--group-001.md#canonical-afd575d646ce1b00dd26dfce3151ed9e0095d30d344c02b5f6206458da435264)
- [active_network_policies](data-sources--network_firewall--reference--group-001.md#canonical-305e9030fb3795b1b3800369e84572869e50f0adfb0d495b02a3110ffa45955b)
- [disable_fast_acl](data-sources--network_firewall--reference--group-001.md#canonical-564dd5ff59989a71d5bb181e2f97423323974e1f4bc126c34c630ece5eb89aee)
- [disable_forward_proxy_policy](data-sources--network_firewall--reference--group-001.md#canonical-1009a248731ea6605183fff5b12b8b466028a679527532ab20e3b1878d05b5f7)
- [disable_network_policy](data-sources--network_firewall--reference--group-001.md#canonical-13109d6287ba0d99ceaf95506b396c1388015b320689d328da9935261769de4b)
- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-9b62b76cf7e299087b5f4928cbe980ffbde2f654430a3985cbcf52e5b4f58a56)

<a id="canonical-c1db3a9936cb42c65172736c3fc1ae492099d197ab2210144fa69b0d58707bcc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bbe07c4407f487615ee9cfb18678b25009532c8ba1fbb0aacdcc09cdeeb90d7a"></a>

## active_enhanced_firewall_policies — active_enhanced_firewall_policies / d1ff08c636a1 / 2

Breadcrumbs:

- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-9b62b76cf7e299087b5f4928cbe980ffbde2f654430a3985cbcf52e5b4f58a56)
- [Property reference](data-sources--network_firewall--reference--group-001.md#canonical-3788cba45b5e266537cf6289bb39431b1592d6bdb952ff53cafe1b93ac1711d3)
- active_enhanced_firewall_policies

<a id="canonical-c4066807d33c540280fcfa7b1b01ac56a71da9e7e9438fd036c144a9542a69e3"></a>

Type: `"single"`. Computed.

\[OneOf: active\_enhanced\_firewall\_policies, active\_network\_policies, disable\_network\_policy;
Default: disable\_network\_policy\] List of Enhanced Firewall Policies These policies use
session-based rules and provide all OPTIONS available under firewall policies with an additional
option for service insertion.

Upstream description:

List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

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

- [active_enhanced_firewall_policies](data-sources--network_firewall--reference--group-001.md#canonical-c4066807d33c540280fcfa7b1b01ac56a71da9e7e9438fd036c144a9542a69e3)
- [active_network_policies](data-sources--network_firewall--reference--group-001.md#canonical-b628c71c77dd61693a019895ebf5e44714be80532a65441eedfe87f5aa695818)
- [disable_network_policy](data-sources--network_firewall--reference--group-001.md#canonical-f2112a5d1ae43ac487b21224e2c126a99dcb7bbee8ad4d717b322e9b4d9b8812)

Select alternatives according to the provider validators above.

<a id="canonical-4ad2143a8dd3aaccf726d4ab7e21f2678146820e02eb00ad511286fb7525e551"></a>

## Direct properties — active_enhanced_firewall_policies / d1ff08c636a1 / 3

- [enhanced_firewall_policies](data-sources--network_firewall--reference--group-001.md#canonical-ee62251a13cfea75dfe86dca81e1ed021303aca31550bc23d8400752a12d53bb): complete subsection reference.

<a id="canonical-af7b4c8ced962e44736e62f2b1029829e783cbf549adecd387afd6b5dfb83ccd"></a>

## Next pages — active_enhanced_firewall_policies / d1ff08c636a1 / 4

- [active_enhanced_firewall_policies.enhanced_firewall_policies](data-sources--network_firewall--reference--group-001.md#canonical-ee62251a13cfea75dfe86dca81e1ed021303aca31550bc23d8400752a12d53bb)
- [Property reference](data-sources--network_firewall--reference--group-001.md#canonical-3788cba45b5e266537cf6289bb39431b1592d6bdb952ff53cafe1b93ac1711d3)
- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-9b62b76cf7e299087b5f4928cbe980ffbde2f654430a3985cbcf52e5b4f58a56)

<a id="canonical-ee62251a13cfea75dfe86dca81e1ed021303aca31550bc23d8400752a12d53bb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5e186b2cba2402cfb4dbb460b30703ad26aa189d744e5e7db97d635d85341be8"></a>

## active_enhanced_firewall_policies.enhanced_firewall_policies — active_enhanced_firewall_policies.enhanced_firewall_policies / b20089dd780e / 2

Breadcrumbs:

- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-9b62b76cf7e299087b5f4928cbe980ffbde2f654430a3985cbcf52e5b4f58a56)
- [Property reference](data-sources--network_firewall--reference--group-001.md#canonical-3788cba45b5e266537cf6289bb39431b1592d6bdb952ff53cafe1b93ac1711d3)
- [active_enhanced_firewall_policies](data-sources--network_firewall--reference--group-001.md#canonical-c1db3a9936cb42c65172736c3fc1ae492099d197ab2210144fa69b0d58707bcc)
- active_enhanced_firewall_policies.enhanced_firewall_policies

<a id="canonical-8c6bcc3e47d8c281c5654d5e47ab43065d3a5986ae4b14988de5303458d1e78f"></a>

Type: `"list"`. Computed.

Ordered List of Enhanced Firewall Policies active.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-99d6ef000874f898fbed2fdac9c02004c58de43799ba59c3bfcf7e52e7b4bb35"></a>

## Direct properties — active_enhanced_firewall_policies.enhanced_firewall_policies / b20089dd780e / 3

<a id="canonical-f501111bd596e470b7316e92e0e82fe11f7c46877cbfcfc526718886ac036a7c"></a>

<a id="canonical-c79ce174fc638541e7b1f9cfd8ddbe96d24d46ccbdf1f90116448171fc6c8c43"></a>

## name property — active_enhanced_firewall_policies.enhanced_firewall_policies / b20089dd780e / 4

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

<a id="canonical-3442d49a1089d282fb53c1cb774a50a668d1caf1eca3871413041f722f8b6156"></a>

<a id="canonical-b308a4c6d8fbb446d67b46f868c82878f4f8e3f10bd0c8f14e7c838ec3c3db03"></a>

## namespace property — active_enhanced_firewall_policies.enhanced_firewall_policies / b20089dd780e / 5

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

<a id="canonical-eb1b60f760d17bace423a3900d4a7aaf8383960d616ae3108dc6f76802aa0479"></a>

<a id="canonical-e14637fa66335aaaf8dd28a1499bf57cf4fd75769cad86f8045f1d20d082b21f"></a>

## tenant property — active_enhanced_firewall_policies.enhanced_firewall_policies / b20089dd780e / 6

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

<a id="canonical-2a229ede07aba541d53bee71be72066cf08e11bc86bb964cf855bab17e946695"></a>

## Next pages — active_enhanced_firewall_policies.enhanced_firewall_policies / b20089dd780e / 7

- [active_enhanced_firewall_policies](data-sources--network_firewall--reference--group-001.md#canonical-c1db3a9936cb42c65172736c3fc1ae492099d197ab2210144fa69b0d58707bcc)
- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-9b62b76cf7e299087b5f4928cbe980ffbde2f654430a3985cbcf52e5b4f58a56)

<a id="canonical-d4354aa962ed2150fd0c3f1333ec9704d2988dad097399766a9954a5d1c8c012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0e9e96a7f5dda0f48efd1a766e43533554e52a19765c183539fadc5b3b0ab023"></a>

## active_fast_acls — active_fast_acls / 6c5f3b09439e / 2

Breadcrumbs:

- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-9b62b76cf7e299087b5f4928cbe980ffbde2f654430a3985cbcf52e5b4f58a56)
- [Property reference](data-sources--network_firewall--reference--group-001.md#canonical-3788cba45b5e266537cf6289bb39431b1592d6bdb952ff53cafe1b93ac1711d3)
- active_fast_acls

<a id="canonical-53fb941de7de8fbc7fcd8445f36d1a55a8094088fa492ecbbeab4670a5be4cb9"></a>

Type: `"single"`. Computed.

\[OneOf: active\_fast\_acls, disable\_fast\_acl; Default: disable\_fast\_acl\] Configuration
parameter for active fast acls.

Upstream description:

List of Fast ACL(s).

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

- [active_fast_acls](data-sources--network_firewall--reference--group-001.md#canonical-53fb941de7de8fbc7fcd8445f36d1a55a8094088fa492ecbbeab4670a5be4cb9)
- [disable_fast_acl](data-sources--network_firewall--reference--group-001.md#canonical-0cf0a26aa833fe6092d0eb000dfaf254f79ee86878b935d21b7ced27a29c65bc)

Select alternatives according to the provider validators above.

<a id="canonical-8add93be5d8ca0e1e88d8927538291d51b414bb3c4c4c8c04f61b37d39773084"></a>

## Direct properties — active_fast_acls / 6c5f3b09439e / 3

- [fast_acls](data-sources--network_firewall--reference--group-001.md#canonical-8e2da80b3645c5ca01c6060b9ed395a5a15d36a778ba1f6a67de3dc8c64370cd): complete subsection reference.

<a id="canonical-184c0125be91817078d1cb6d747ff53816e708ebe8d013d37f631d1aa480300e"></a>

## Next pages — active_fast_acls / 6c5f3b09439e / 4

- [active_fast_acls.fast_acls](data-sources--network_firewall--reference--group-001.md#canonical-8e2da80b3645c5ca01c6060b9ed395a5a15d36a778ba1f6a67de3dc8c64370cd)
- [Property reference](data-sources--network_firewall--reference--group-001.md#canonical-3788cba45b5e266537cf6289bb39431b1592d6bdb952ff53cafe1b93ac1711d3)
- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-9b62b76cf7e299087b5f4928cbe980ffbde2f654430a3985cbcf52e5b4f58a56)

<a id="canonical-8e2da80b3645c5ca01c6060b9ed395a5a15d36a778ba1f6a67de3dc8c64370cd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-11ce2752cc78e68a902659889b4fcce8a942775f5b79cd00b13572ee49f54f1d"></a>

## active_fast_acls.fast_acls — active_fast_acls.fast_acls / 2aaaaaa4079e / 2

Breadcrumbs:

- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-9b62b76cf7e299087b5f4928cbe980ffbde2f654430a3985cbcf52e5b4f58a56)
- [Property reference](data-sources--network_firewall--reference--group-001.md#canonical-3788cba45b5e266537cf6289bb39431b1592d6bdb952ff53cafe1b93ac1711d3)
- [active_fast_acls](data-sources--network_firewall--reference--group-001.md#canonical-d4354aa962ed2150fd0c3f1333ec9704d2988dad097399766a9954a5d1c8c012)
- active_fast_acls.fast_acls

<a id="canonical-60b847aeb6e7d623d23f8b7a97fbc77af547046ce6c1463f86fed70534589afe"></a>

Type: `"list"`. Computed.

Ordered List of Fast ACL(s) active for this network firewall.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-9ea7c2c3f4cdbf3953845f3f3f7aa9655edc5ad335818c9540cc2d529112143b"></a>

## Direct properties — active_fast_acls.fast_acls / 2aaaaaa4079e / 3

<a id="canonical-644dbb85ee0efe754f4d57cd2613b8ed64a11d480602b1107b75f8a2fba1bd33"></a>

<a id="canonical-5dcdb3372462a2ad1e23bb45547337da5ddf93c52a7ce4d1236242baae359e27"></a>

## name property — active_fast_acls.fast_acls / 2aaaaaa4079e / 4

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

<a id="canonical-6905ad1eac09cd02f7b642ba84238fcc411f91431485317d6e833218e8726916"></a>

<a id="canonical-1dd6289b5265c76db61c9ed64e275fbf6d6d49545c03eca7087b90dfa85321d2"></a>

## namespace property — active_fast_acls.fast_acls / 2aaaaaa4079e / 5

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

<a id="canonical-3a15215bd51962428b32b6fe9502b3f9733700cb94e3662e4e7c57a6fc8e2b1a"></a>

<a id="canonical-6d62b909f37a8e4ec9483fd562df0965bad2cc0af115836e41f1661362ddb59f"></a>

## tenant property — active_fast_acls.fast_acls / 2aaaaaa4079e / 6

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

<a id="canonical-555319f6b5133281ee98b5cd1a5359ade48bef13f9e0374234632b200b6ffe84"></a>

## Next pages — active_fast_acls.fast_acls / 2aaaaaa4079e / 7

- [active_fast_acls](data-sources--network_firewall--reference--group-001.md#canonical-d4354aa962ed2150fd0c3f1333ec9704d2988dad097399766a9954a5d1c8c012)
- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-9b62b76cf7e299087b5f4928cbe980ffbde2f654430a3985cbcf52e5b4f58a56)

<a id="canonical-afd575d646ce1b00dd26dfce3151ed9e0095d30d344c02b5f6206458da435264"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0c6172ce81ac137e0776bd0413c21a0776b3a9f4cb698037b9634bab0282172c"></a>

## active_forward_proxy_policies — active_forward_proxy_policies / ab3eb9af5d50 / 2

Breadcrumbs:

- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-9b62b76cf7e299087b5f4928cbe980ffbde2f654430a3985cbcf52e5b4f58a56)
- [Property reference](data-sources--network_firewall--reference--group-001.md#canonical-3788cba45b5e266537cf6289bb39431b1592d6bdb952ff53cafe1b93ac1711d3)
- active_forward_proxy_policies

<a id="canonical-09cda56c4701ff9cbecfecfb0e749f5f4f584a4f74bf3e54ca99e67d216392f5"></a>

Type: `"single"`. Computed.

\[OneOf: active\_forward\_proxy\_policies, disable\_forward\_proxy\_policy; Default:
disable\_forward\_proxy\_policy\] Ordered List of Forward Proxy Policies active.

Upstream description:

Ordered List of Forward Proxy Policies active.

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

- [active_forward_proxy_policies](data-sources--network_firewall--reference--group-001.md#canonical-09cda56c4701ff9cbecfecfb0e749f5f4f584a4f74bf3e54ca99e67d216392f5)
- [disable_forward_proxy_policy](data-sources--network_firewall--reference--group-001.md#canonical-cdbb308e1853445eb0041820c490f1981a8d4ad8eee0e546049e023a110ec2b5)

Select alternatives according to the provider validators above.

<a id="canonical-31e0bf4dc992a70e3acd919711c0256247c460e83da7f7fa5e424c22c7d3f4a2"></a>

## Direct properties — active_forward_proxy_policies / ab3eb9af5d50 / 3

- [forward_proxy_policies](data-sources--network_firewall--reference--group-001.md#canonical-4dce6774dadc244eb74c3c763afbaa70905cef7b9f35577acc15e4846112b433): complete subsection reference.

<a id="canonical-182f752419259febd7a04aa3df8c944c7e54f252b608eced58359b5e0966b77b"></a>

## Next pages — active_forward_proxy_policies / ab3eb9af5d50 / 4

- [active_forward_proxy_policies.forward_proxy_policies](data-sources--network_firewall--reference--group-001.md#canonical-4dce6774dadc244eb74c3c763afbaa70905cef7b9f35577acc15e4846112b433)
- [Property reference](data-sources--network_firewall--reference--group-001.md#canonical-3788cba45b5e266537cf6289bb39431b1592d6bdb952ff53cafe1b93ac1711d3)
- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-9b62b76cf7e299087b5f4928cbe980ffbde2f654430a3985cbcf52e5b4f58a56)

<a id="canonical-4dce6774dadc244eb74c3c763afbaa70905cef7b9f35577acc15e4846112b433"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-26933f4988ea8a9cac40eabe1da2782630492325cf32d34f98f0118cf50d31e0"></a>

## active_forward_proxy_policies.forward_proxy_policies — active_forward_proxy_policies.forward_proxy_policies / 64bf42dc2964 / 2

Breadcrumbs:

- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-9b62b76cf7e299087b5f4928cbe980ffbde2f654430a3985cbcf52e5b4f58a56)
- [Property reference](data-sources--network_firewall--reference--group-001.md#canonical-3788cba45b5e266537cf6289bb39431b1592d6bdb952ff53cafe1b93ac1711d3)
- [active_forward_proxy_policies](data-sources--network_firewall--reference--group-001.md#canonical-afd575d646ce1b00dd26dfce3151ed9e0095d30d344c02b5f6206458da435264)
- active_forward_proxy_policies.forward_proxy_policies

<a id="canonical-0eea5ad4402949090632dce3c370b75325452e80ba6444752f61a28c95673302"></a>

Type: `"list"`. Computed.

Ordered List of Forward Proxy Policies active.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-3170926a22d77b00948d7ab835962e7145a210730234dcd2bcd21840c13aefbc"></a>

## Direct properties — active_forward_proxy_policies.forward_proxy_policies / 64bf42dc2964 / 3

<a id="canonical-d72bcb080b31a806f8dc04f6e1c4770e99951ef8942e587165916281b571ccef"></a>

<a id="canonical-b99cf0b4285fdb9f4b997821110c5271490be6803d1a724fb8ec3323cb011ba8"></a>

## name property — active_forward_proxy_policies.forward_proxy_policies / 64bf42dc2964 / 4

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

<a id="canonical-168759d8a73776755d12c9c9327facdc604d2d6b204108f09f4eff7db529bcdd"></a>

<a id="canonical-4fbf9c5bca124278b04d1b13ae4f8f5a29825c12e1481412bb506cd2490620cf"></a>

## namespace property — active_forward_proxy_policies.forward_proxy_policies / 64bf42dc2964 / 5

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

<a id="canonical-f1a91a78f05de2246634d8024be6ea200e0427b78ac49e3e630a3c168857364e"></a>

<a id="canonical-dbbd742bed567839769e8d38304383e4583700614e0306e64e9cbe9eb4312384"></a>

## tenant property — active_forward_proxy_policies.forward_proxy_policies / 64bf42dc2964 / 6

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

<a id="canonical-b324c9342569be1e93e08e7e2aaf9761e3c9726a24baa956e677303c1f07f611"></a>

## Next pages — active_forward_proxy_policies.forward_proxy_policies / 64bf42dc2964 / 7

- [active_forward_proxy_policies](data-sources--network_firewall--reference--group-001.md#canonical-afd575d646ce1b00dd26dfce3151ed9e0095d30d344c02b5f6206458da435264)
- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-9b62b76cf7e299087b5f4928cbe980ffbde2f654430a3985cbcf52e5b4f58a56)

<a id="canonical-305e9030fb3795b1b3800369e84572869e50f0adfb0d495b02a3110ffa45955b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-927fd2ab379e2f5b3f24d9c2c4dba5581b48af2e8f77d7d095da62cceed54828"></a>

## active_network_policies — active_network_policies / 40412f3e7a4e / 2

Breadcrumbs:

- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-9b62b76cf7e299087b5f4928cbe980ffbde2f654430a3985cbcf52e5b4f58a56)
- [Property reference](data-sources--network_firewall--reference--group-001.md#canonical-3788cba45b5e266537cf6289bb39431b1592d6bdb952ff53cafe1b93ac1711d3)
- active_network_policies

<a id="canonical-b628c71c77dd61693a019895ebf5e44714be80532a65441eedfe87f5aa695818"></a>

Type: `"single"`. Computed.

Configuration parameter for active network policies.

Upstream description:

List of firewall policy views.

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

<a id="canonical-7d69e50529edee5b21f67a81e85e03b8a63bf29a06226c984f069174beedf92a"></a>

## Direct properties — active_network_policies / 40412f3e7a4e / 3

- [network_policies](data-sources--network_firewall--reference--group-001.md#canonical-8f5816363500d8406d9f5e6a33547bc4921b1668d8287cc3fe848dcc12a63399): complete subsection reference.

<a id="canonical-81f4172efc581d66917eda64c5e4aba0b5ff734b42c97be418339997c7d42425"></a>

## Next pages — active_network_policies / 40412f3e7a4e / 4

- [active_network_policies.network_policies](data-sources--network_firewall--reference--group-001.md#canonical-8f5816363500d8406d9f5e6a33547bc4921b1668d8287cc3fe848dcc12a63399)
- [Property reference](data-sources--network_firewall--reference--group-001.md#canonical-3788cba45b5e266537cf6289bb39431b1592d6bdb952ff53cafe1b93ac1711d3)
- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-9b62b76cf7e299087b5f4928cbe980ffbde2f654430a3985cbcf52e5b4f58a56)

<a id="canonical-8f5816363500d8406d9f5e6a33547bc4921b1668d8287cc3fe848dcc12a63399"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-00e8c9f480c04175a93c70cfecaf29802c50a4ab882ac8d4f9db6880f10efd90"></a>

## active_network_policies.network_policies — active_network_policies.network_policies / b09d6f02e292 / 2

Breadcrumbs:

- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-9b62b76cf7e299087b5f4928cbe980ffbde2f654430a3985cbcf52e5b4f58a56)
- [Property reference](data-sources--network_firewall--reference--group-001.md#canonical-3788cba45b5e266537cf6289bb39431b1592d6bdb952ff53cafe1b93ac1711d3)
- [active_network_policies](data-sources--network_firewall--reference--group-001.md#canonical-305e9030fb3795b1b3800369e84572869e50f0adfb0d495b02a3110ffa45955b)
- active_network_policies.network_policies

<a id="canonical-cc9d95ea056917dde22181974a91d0fa7f3a4677cf3447af26e1df9e894a5ffb"></a>

Type: `"list"`. Computed.

Ordered List of Firewall Policies active for this network firewall.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-4f608260ecfe1fffce5b7d81d87f030d80dc2e0eb4f55a4fd1e59bf685a5816b"></a>

## Direct properties — active_network_policies.network_policies / b09d6f02e292 / 3

<a id="canonical-b6c6880e43ab7d721146f40c5a9ebb7ec196272613fe5436be6b4e316091136e"></a>

<a id="canonical-8901c16a0bf53f84a3f77b9034706a206763ad2966452b3302cb120c42131abb"></a>

## name property — active_network_policies.network_policies / b09d6f02e292 / 4

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

<a id="canonical-f7651208b70f112a90e2a42a5c02651447246d68e2f300199e7110bc2fcd7c1c"></a>

<a id="canonical-510bbf34babd59723a03e27905af55a48829a343c3d07c43c39540f5b7fcafaa"></a>

## namespace property — active_network_policies.network_policies / b09d6f02e292 / 5

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

<a id="canonical-ffc54afbc144e20e9902c06650adc9095ba5ccd592c4294d7ae34fe365109ba4"></a>

<a id="canonical-83814e844ae864fd0a9ea6d7df91d9aca1c872401a8070f1da12e2c512da6e02"></a>

## tenant property — active_network_policies.network_policies / b09d6f02e292 / 6

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

<a id="canonical-c67b43c7579b19b9855c8225dd5a004bcb5dc2841b93192ac33a969112f60f31"></a>

## Next pages — active_network_policies.network_policies / b09d6f02e292 / 7

- [active_network_policies](data-sources--network_firewall--reference--group-001.md#canonical-305e9030fb3795b1b3800369e84572869e50f0adfb0d495b02a3110ffa45955b)
- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-9b62b76cf7e299087b5f4928cbe980ffbde2f654430a3985cbcf52e5b4f58a56)

<a id="canonical-564dd5ff59989a71d5bb181e2f97423323974e1f4bc126c34c630ece5eb89aee"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-61c45295640bccf789a0d71867e8bf20e53cd197331c87d1b6a8cea77621d684"></a>

## disable_fast_acl — disable_fast_acl / 86b64cdf2f88 / 2

Breadcrumbs:

- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-9b62b76cf7e299087b5f4928cbe980ffbde2f654430a3985cbcf52e5b4f58a56)
- [Property reference](data-sources--network_firewall--reference--group-001.md#canonical-3788cba45b5e266537cf6289bb39431b1592d6bdb952ff53cafe1b93ac1711d3)
- disable_fast_acl

<a id="canonical-0cf0a26aa833fe6092d0eb000dfaf254f79ee86878b935d21b7ced27a29c65bc"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable fast acl. Defaults to \`map\[\]\`. Server applies default when
omitted.

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

<a id="canonical-e9f9609ea41686e89b411e6d7d7b60ce53a2b176cda92839dd11f735417acbb3"></a>

## Direct properties — disable_fast_acl / 86b64cdf2f88 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1fba91ac4fc6e7b0dd063e3ed4b0bdcc6866f97ef887437a210c22f1e7762793"></a>

## Next pages — disable_fast_acl / 86b64cdf2f88 / 4

- [Property reference](data-sources--network_firewall--reference--group-001.md#canonical-3788cba45b5e266537cf6289bb39431b1592d6bdb952ff53cafe1b93ac1711d3)
- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-9b62b76cf7e299087b5f4928cbe980ffbde2f654430a3985cbcf52e5b4f58a56)

<a id="canonical-1009a248731ea6605183fff5b12b8b466028a679527532ab20e3b1878d05b5f7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e83d28dad3d51476837323a75dd192e72dbccd6f9c4fcd9ed3a568abad4d885f"></a>

## disable_forward_proxy_policy — disable_forward_proxy_policy / 43e2ade10290 / 2

Breadcrumbs:

- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-9b62b76cf7e299087b5f4928cbe980ffbde2f654430a3985cbcf52e5b4f58a56)
- [Property reference](data-sources--network_firewall--reference--group-001.md#canonical-3788cba45b5e266537cf6289bb39431b1592d6bdb952ff53cafe1b93ac1711d3)
- disable_forward_proxy_policy

<a id="canonical-cdbb308e1853445eb0041820c490f1981a8d4ad8eee0e546049e023a110ec2b5"></a>

Type: `["object", {}]`. Computed.

Policy configuration for this feature. Defaults to \`map\[\]\`. Server applies default when omitted.

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

<a id="canonical-c4b768878eb831b0afb89623f683cce65e15e5f29331311a8c939358e87e0ed3"></a>

## Direct properties — disable_forward_proxy_policy / 43e2ade10290 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d19f79d96640e98bfca68230c7def93bddd9dcb2e1268afc1dfd85cce7edb074"></a>

## Next pages — disable_forward_proxy_policy / 43e2ade10290 / 4

- [Property reference](data-sources--network_firewall--reference--group-001.md#canonical-3788cba45b5e266537cf6289bb39431b1592d6bdb952ff53cafe1b93ac1711d3)
- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-9b62b76cf7e299087b5f4928cbe980ffbde2f654430a3985cbcf52e5b4f58a56)

<a id="canonical-13109d6287ba0d99ceaf95506b396c1388015b320689d328da9935261769de4b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1e5c03c8410393b795943db6262ae24fb6056c6913c0badad069a37a39e4fae0"></a>

## disable_network_policy — disable_network_policy / 24638e9b9f92 / 2

Breadcrumbs:

- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-9b62b76cf7e299087b5f4928cbe980ffbde2f654430a3985cbcf52e5b4f58a56)
- [Property reference](data-sources--network_firewall--reference--group-001.md#canonical-3788cba45b5e266537cf6289bb39431b1592d6bdb952ff53cafe1b93ac1711d3)
- disable_network_policy

<a id="canonical-f2112a5d1ae43ac487b21224e2c126a99dcb7bbee8ad4d717b322e9b4d9b8812"></a>

Type: `["object", {}]`. Computed.

Policy configuration for this feature. Defaults to \`map\[\]\`. Server applies default when omitted.

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

<a id="canonical-a2d2e6cb832afaf846e70281c791cc20df67614a83dbb9741b4e981da4b977e8"></a>

## Direct properties — disable_network_policy / 24638e9b9f92 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bdd0b759cf89bbbf852cb5f5e7324ab56fb5249677cd77d5fba40d56d68c7364"></a>

## Next pages — disable_network_policy / 24638e9b9f92 / 4

- [Property reference](data-sources--network_firewall--reference--group-001.md#canonical-3788cba45b5e266537cf6289bb39431b1592d6bdb952ff53cafe1b93ac1711d3)
- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-9b62b76cf7e299087b5f4928cbe980ffbde2f654430a3985cbcf52e5b4f58a56)
