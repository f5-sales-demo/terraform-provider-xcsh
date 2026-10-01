---
page_title: "xcsh_cloud_link reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cloud_link reference."
---

# xcsh_cloud_link reference

<a id="canonical-f9f572eb4dd4355ccf1d4d4a87719ef48acf695e515e7a8097f8aa7e0a96fa35"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f34f25713738bc190fe596ae8477d015dd4d9c4e9934fb7d7c6a5839b5f2ff41"></a>

## Property reference — Property reference / e21a50164c17 / 2

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md#canonical-e2bde7dd950eff273e290001dfe725260a3cdddc3ece9f8132c786c2291ffa6f)
- Property reference

<a id="canonical-5a0916a0b97f3a46ffd3ccbc4ccbced2217be84e91780fa1571cdf3ee42412e0"></a>

## Direct properties — Property reference / e21a50164c17 / 3

<a id="canonical-ca8cf22cd8c841feb6c3fede64fed3075aa71d343a8abea9ffc2393b94e6b406"></a>

<a id="canonical-c69959a18f7eeed439ccf78f8bb756b84301742a5b1c5892572a96d380537fbe"></a>

## annotations property — Property reference / e21a50164c17 / 4

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

- [aws](resources--cloud_link--reference--group-001.md#canonical-9a19c7c88eed6a7a655520fb70180ac858a289742be5f8497055a556855b0686): complete subsection reference.

<a id="canonical-e480eb1067e48ee819e4e9315f9f6af189de26079b5577299da2afddce3b86a3"></a>

<a id="canonical-07caf7313edcd106e03c83783b8ab7a08d5443131eaa415e367d0bb7a5b2455e"></a>

## description property — Property reference / e21a50164c17 / 5

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

<a id="canonical-51e94521e6897f3d22acf6b4e6b9815ba6991cc285b76a68d68cb5480615e016"></a>

<a id="canonical-92c3b58e36b58cacf8c032a8bd07e3c717307115bd72bfb4eec99a7c08f89e80"></a>

## disable property — Property reference / e21a50164c17 / 6

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

- [disabled](resources--cloud_link--reference--group-001.md#canonical-4a7582a37c7b581ae3e077db76831503a8dc9e2b39d7a5f66d2f268f2aa7411c): complete subsection reference.

- [enabled](resources--cloud_link--reference--group-001.md#canonical-b8705c14c9d8e6bec591d18c74d04c6d9d70acb0c574cd6d29120416eed8a236): complete subsection reference.

- [gcp](resources--cloud_link--reference--group-001.md#canonical-98b04cd1a263fa94de21346302cdfec25d91055cbcf9c3b463608298b2f20fd9): complete subsection reference.

<a id="canonical-feb76e9237cf55d84cb4bee42c5868de122bade27ec94e37507c1cb7c1a51973"></a>

<a id="canonical-126bb6bd94c305825e3d753d341c72b63352195c5d756245a92eac6f213ae334"></a>

## id property — Property reference / e21a50164c17 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-3ce6a8335538981f612dbb6875e2ace2d659f72f2c88b48cca513c0658f2633e"></a>

<a id="canonical-9727c947046c0ea9adbcd319906eb305c79ca3bb90d828b67c86a2e52bb900e8"></a>

## labels property — Property reference / e21a50164c17 / 8

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

<a id="canonical-e901e5f25a6229875d83004cbf615c118032237004648bed16ce2c143ba0f0b8"></a>

<a id="canonical-e56d096d3ca398a1d9afe46e5bec7c78c0b72d0e4b5008b1f4643379368092ea"></a>

## name property — Property reference / e21a50164c17 / 9

Type: `"string"`. Required.

Name of the Cloud Link. Must be unique within the namespace.

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

<a id="canonical-30f02e590f391b823359661ff99180bc55d8ad66767325ca873a0eeb85f9d5be"></a>

<a id="canonical-ad6a593a6b471ef16afa428ad27cf4a807bd411bf34d51a2ac980193927ed558"></a>

## namespace property — Property reference / e21a50164c17 / 10

Type: `"string"`. Required.

Namespace where the Cloud Link is created.

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

- [timeouts](resources--cloud_link--reference--group-001.md#canonical-407980008cac2d543b4e7b2f130f22cee93d886cd60141abd6ce82624b6f9f80): complete subsection reference.

<a id="canonical-e3969c0abd5dde0eebac2b47c9500efb9900a64ab06f213359fcb6f7911a5b13"></a>

## All schema paths — Property reference / e21a50164c17 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--cloud_link--reference--group-001.md#canonical-ca8cf22cd8c841feb6c3fede64fed3075aa71d343a8abea9ffc2393b94e6b406) |
| `aws` | [aws](resources--cloud_link--reference--group-001.md#canonical-42bb50301ad66808b11fcc68e8ed8382587a5ae1987542bc7181fce818529fdc) |
| `aws.aws_cred` | [aws.aws_cred](resources--cloud_link--reference--group-001.md#canonical-45c98749a6f5a59251ffbd2346339d7c13f8527bb218146b9d83a1ea5a6b5275) |
| `aws.aws_cred.name` | [aws.aws_cred.name](resources--cloud_link--reference--group-001.md#canonical-aaeb3116b5e7f6346cfbf4ce0fbb2c3575fb82a44f90e1d4259b5b0c279b20d5) |
| `aws.aws_cred.namespace` | [aws.aws_cred.namespace](resources--cloud_link--reference--group-001.md#canonical-ed2a96c11e5185eed6aec214582a6c569a1624a8dc03027df4c3197d9637577f) |
| `aws.aws_cred.tenant` | [aws.aws_cred.tenant](resources--cloud_link--reference--group-001.md#canonical-73530f86c61659897a203ab01b3470e1a791466604f0f27e84a94d494f078ae0) |
| `aws.byoc` | [aws.byoc](resources--cloud_link--reference--group-001.md#canonical-f35ea3c1aee90dba1810e6f70da5b3ed2795f5e5ec1f56fcb0000abd710659a2) |
| `aws.byoc.connections` | [aws.byoc.connections](resources--cloud_link--reference--group-001.md#canonical-87a7a6424812cb22d05880c53ed32bc53a91f58784dd512c9cadfc836d90502f) |
| `aws.byoc.connections.auth_key` | [aws.byoc.connections.auth_key](resources--cloud_link--reference--group-001.md#canonical-9b77a033e3c5355f37a038aac22960886a6b93ed79e2b62abe04c6755da6da2b) |
| `aws.byoc.connections.auth_key.blindfold_secret_info` | [aws.byoc.connections.auth_key.blindfold_secret_info](resources--cloud_link--reference--group-001.md#canonical-2f6b6d55f74c166aea45299ce41efa21fef327268eff5146aaf62a3294924259) |
| `aws.byoc.connections.auth_key.blindfold_secret_info.decryption_provider` | [aws.byoc.connections.auth_key.blindfold_secret_info.decryption_provider](resources--cloud_link--reference--group-001.md#canonical-5a0c6530c8f42483f4437d61140c4a6aabdbade33c484f953bcf1de95d5ddf85) |
| `aws.byoc.connections.auth_key.blindfold_secret_info.location` | [aws.byoc.connections.auth_key.blindfold_secret_info.location](resources--cloud_link--reference--group-001.md#canonical-23b13746229b0663c0841af90cd16ede2de6a778f635e3d6deddf5d7adf2f015) |
| `aws.byoc.connections.auth_key.blindfold_secret_info.store_provider` | [aws.byoc.connections.auth_key.blindfold_secret_info.store_provider](resources--cloud_link--reference--group-001.md#canonical-5a6194b516f8734b56b941a6b03081b1b0ba309bba273dbbda205d980a58d9c3) |
| `aws.byoc.connections.auth_key.clear_secret_info` | [aws.byoc.connections.auth_key.clear_secret_info](resources--cloud_link--reference--group-001.md#canonical-79361ecd727ed040c21a0d317213a816e4a1e907b4fc7332162aaa0eb5953552) |
| `aws.byoc.connections.auth_key.clear_secret_info.provider_ref` | [aws.byoc.connections.auth_key.clear_secret_info.provider_ref](resources--cloud_link--reference--group-001.md#canonical-fa7a83498b7712fd99a2758e68858e0d54c0fb77f796ace77393443b51f7a62f) |
| `aws.byoc.connections.auth_key.clear_secret_info.url` | [aws.byoc.connections.auth_key.clear_secret_info.url](resources--cloud_link--reference--group-001.md#canonical-2447158e18137c3b4e950e020d83a698393d7f1980d97d942fb36d0c534c319c) |
| `aws.byoc.connections.bgp_asn` | [aws.byoc.connections.bgp_asn](resources--cloud_link--reference--group-001.md#canonical-0e9db001357088d73d65dbc3def181e894d602bacf6141d8d3b66e5c2b501bc6) |
| `aws.byoc.connections.connection_id` | [aws.byoc.connections.connection_id](resources--cloud_link--reference--group-001.md#canonical-77da4421e8990d7681532bdb36aacf9b2bb583ff9fe8aabaca22a26bfe652e8e) |
| `aws.byoc.connections.ipv4` | [aws.byoc.connections.ipv4](resources--cloud_link--reference--group-001.md#canonical-8eeababf30a295bb72d456c01cbadcbaf8b3437a96581896238cb0745e8aca6e) |
| `aws.byoc.connections.ipv4.aws_router_peer_address` | [aws.byoc.connections.ipv4.aws_router_peer_address](resources--cloud_link--reference--group-001.md#canonical-16879842cd10cd862563a3c7f66948e58dacc13f63b19e9c2de83f239b20469b) |
| `aws.byoc.connections.ipv4.router_peer_address` | [aws.byoc.connections.ipv4.router_peer_address](resources--cloud_link--reference--group-001.md#canonical-26d20bedfa771a790472401799bc2b4f773fc98486fc12205411d2eb6028254f) |
| `aws.byoc.connections.metadata` | [aws.byoc.connections.metadata](resources--cloud_link--reference--group-001.md#canonical-f46f7a4cb3a480809f2ad3b94fee07db43806a45e6fd41058d9583fbdb4a86e5) |
| `aws.byoc.connections.metadata.description_spec` | [aws.byoc.connections.metadata.description_spec](resources--cloud_link--reference--group-001.md#canonical-59503ac7c22e1bffbdd079366ca0b73a521dc0dcb022410b492890010a662171) |
| `aws.byoc.connections.metadata.name` | [aws.byoc.connections.metadata.name](resources--cloud_link--reference--group-001.md#canonical-78744c5a8ec6c0adabb611c64f8387c6f5ccb9a9fbca2def0beef751a87a3606) |
| `aws.byoc.connections.region` | [aws.byoc.connections.region](resources--cloud_link--reference--group-001.md#canonical-7c12f02a3e90dcd9a81a129af338a0fd359176938009d5e14219c8183e98532b) |
| `aws.byoc.connections.system_generated_name` | [aws.byoc.connections.system_generated_name](resources--cloud_link--reference--group-001.md#canonical-b48c67fe9c604a466fcff0dc9a41a2b02e18037e9a55a602e94e2922d5b2d5a8) |
| `aws.byoc.connections.tags` | [aws.byoc.connections.tags](resources--cloud_link--reference--group-001.md#canonical-21faf9b2f9817d2f17d92896a70bb5f437b2a7b79fcdafd1167e55a3f2a0c74d) |
| `aws.byoc.connections.user_assigned_name` | [aws.byoc.connections.user_assigned_name](resources--cloud_link--reference--group-001.md#canonical-9a879d834611c3b9e1dd36590c7f2ea945cd7c8f6db1c8a9f9aff787e936b592) |
| `aws.byoc.connections.virtual_interface_type` | [aws.byoc.connections.virtual_interface_type](resources--cloud_link--reference--group-001.md#canonical-5c4dfb41a6fb176b1a1db82552e7b752c536fc7e26ee0dd4fd56ede292191e32) |
| `aws.byoc.connections.vlan` | [aws.byoc.connections.vlan](resources--cloud_link--reference--group-001.md#canonical-512beac77628a844a61fa0495db7ceb6f801a5d6ac7f82ebbdfe5db2ed050a96) |
| `aws.custom_asn` | [aws.custom_asn](resources--cloud_link--reference--group-001.md#canonical-d1ad23ba41747eb34a58a268e234fcc73223ae7b82ee1d8d572635b72ca6e9f0) |
| `description` | [description](resources--cloud_link--reference--group-001.md#canonical-e480eb1067e48ee819e4e9315f9f6af189de26079b5577299da2afddce3b86a3) |
| `disable` | [disable](resources--cloud_link--reference--group-001.md#canonical-51e94521e6897f3d22acf6b4e6b9815ba6991cc285b76a68d68cb5480615e016) |
| `disabled` | [disabled](resources--cloud_link--reference--group-001.md#canonical-702d2f35a7f3d57f00d8eb6b000fce9c61dd97e4e23eae3e8c188c6a32ee707d) |
| `enabled` | [enabled](resources--cloud_link--reference--group-001.md#canonical-4548c7067b83df24918124a71cae1f86634fbf6d2d693a4613e3f1ec52671905) |
| `enabled.cloudlink_network_name` | [enabled.cloudlink_network_name](resources--cloud_link--reference--group-001.md#canonical-d6717732848727f266a9a724dcede0a8df649107fb059bb503bf797bc8453578) |
| `gcp` | [gcp](resources--cloud_link--reference--group-001.md#canonical-08893a2fb138c2038185a105036b37180ec884db77427d7f03d89ff6bb94aa21) |
| `gcp.byoc` | [gcp.byoc](resources--cloud_link--reference--group-001.md#canonical-23568d4584b2b0b92cea6555c8dd7fb9b802d469aff38bb9b5396847c6bfd981) |
| `gcp.byoc.connections` | [gcp.byoc.connections](resources--cloud_link--reference--group-001.md#canonical-7d6abe8456acd336dc26952be25799df2fd11c0ba5c781ee124c2a976d248e8e) |
| `gcp.byoc.connections.interconnect_attachment_name` | [gcp.byoc.connections.interconnect_attachment_name](resources--cloud_link--reference--group-001.md#canonical-78edf06e2eb2494bd34f87e3a5a47f1b5f92f40908aabe5d8d726d11dbfec194) |
| `gcp.byoc.connections.metadata` | [gcp.byoc.connections.metadata](resources--cloud_link--reference--group-001.md#canonical-be3de5a6640ff5870472efa66c93ac74b98d8a203850e57b44070ceaf9a364d8) |
| `gcp.byoc.connections.metadata.description_spec` | [gcp.byoc.connections.metadata.description_spec](resources--cloud_link--reference--group-001.md#canonical-dab8630b39d7a0f2b979f9886414b00e4fc95ef65771ae755abad1b067d066a7) |
| `gcp.byoc.connections.metadata.name` | [gcp.byoc.connections.metadata.name](resources--cloud_link--reference--group-001.md#canonical-435098ed9e69e11bcb21c218984afb32446719dfab969ccd09e21f738ed97ca9) |
| `gcp.byoc.connections.project` | [gcp.byoc.connections.project](resources--cloud_link--reference--group-001.md#canonical-4017db3e815ce6d06028c125f3b2e24d85099b2c4d564f600f2967961d8e8b64) |
| `gcp.byoc.connections.region` | [gcp.byoc.connections.region](resources--cloud_link--reference--group-001.md#canonical-4f33c39eaec28c4bee3374ccda360a57fe7fcc1d9975b64c89b410d154c757e0) |
| `gcp.byoc.connections.same_as_credential` | [gcp.byoc.connections.same_as_credential](resources--cloud_link--reference--group-001.md#canonical-1a84269eb9c19588aaf88767a2037167df6e0b49154fd2d87b7b87bfdb038019) |
| `gcp.gcp_cred` | [gcp.gcp_cred](resources--cloud_link--reference--group-001.md#canonical-02b0cf1064e23607b9644c6fe500cf85441de84725bd2569ad037517bcff2526) |
| `gcp.gcp_cred.name` | [gcp.gcp_cred.name](resources--cloud_link--reference--group-001.md#canonical-b03ae26c2cdee37996a07c96264acfa81611538ca48ee91f79f38ca606b7ae80) |
| `gcp.gcp_cred.namespace` | [gcp.gcp_cred.namespace](resources--cloud_link--reference--group-001.md#canonical-cbad7213c4dd1bc922f31ac0e9e6ae8c5b6e8a22c2464e5bc00e150e77f9a60d) |
| `gcp.gcp_cred.tenant` | [gcp.gcp_cred.tenant](resources--cloud_link--reference--group-001.md#canonical-960187c666731c79461e5925abc97a067370040d1dd3d4676087cd8c93595035) |
| `id` | [id](resources--cloud_link--reference--group-001.md#canonical-feb76e9237cf55d84cb4bee42c5868de122bade27ec94e37507c1cb7c1a51973) |
| `labels` | [labels](resources--cloud_link--reference--group-001.md#canonical-3ce6a8335538981f612dbb6875e2ace2d659f72f2c88b48cca513c0658f2633e) |
| `name` | [name](resources--cloud_link--reference--group-001.md#canonical-e901e5f25a6229875d83004cbf615c118032237004648bed16ce2c143ba0f0b8) |
| `namespace` | [namespace](resources--cloud_link--reference--group-001.md#canonical-30f02e590f391b823359661ff99180bc55d8ad66767325ca873a0eeb85f9d5be) |
| `timeouts` | [timeouts](resources--cloud_link--reference--group-001.md#canonical-51a4f5daddbbed2aacc8cf0eead62606f49de452f1c165a3fffcd5626be340ef) |
| `timeouts.create` | [timeouts.create](resources--cloud_link--reference--group-001.md#canonical-65038fbd75e19dcb633ca0ccdc2069a8b2418b6bd89a349322745809bb2e6124) |
| `timeouts.delete` | [timeouts.delete](resources--cloud_link--reference--group-001.md#canonical-63f07ef8527de6aaa7357ef34d6f90fb16161cae2f80889c8924f230619e6a4d) |
| `timeouts.read` | [timeouts.read](resources--cloud_link--reference--group-001.md#canonical-117db8d285456a4e8edf44857120bf9afecd474b6110a1512436b4b124b1100a) |
| `timeouts.update` | [timeouts.update](resources--cloud_link--reference--group-001.md#canonical-6172a5a11be2275f73910304dd4806a077680268f6cf6adbc97601e993e8b21f) |

<a id="canonical-c7bd1718b976d60f80a91e7498cad0fe882d3bca705e2ad28bd5b7fb5463c747"></a>

## Next pages — Property reference / e21a50164c17 / 12

- [aws](resources--cloud_link--reference--group-001.md#canonical-9a19c7c88eed6a7a655520fb70180ac858a289742be5f8497055a556855b0686)
- [disabled](resources--cloud_link--reference--group-001.md#canonical-4a7582a37c7b581ae3e077db76831503a8dc9e2b39d7a5f66d2f268f2aa7411c)
- [enabled](resources--cloud_link--reference--group-001.md#canonical-b8705c14c9d8e6bec591d18c74d04c6d9d70acb0c574cd6d29120416eed8a236)
- [gcp](resources--cloud_link--reference--group-001.md#canonical-98b04cd1a263fa94de21346302cdfec25d91055cbcf9c3b463608298b2f20fd9)
- [timeouts](resources--cloud_link--reference--group-001.md#canonical-407980008cac2d543b4e7b2f130f22cee93d886cd60141abd6ce82624b6f9f80)
- [xcsh_cloud_link](../resources/cloud_link.md#canonical-e2bde7dd950eff273e290001dfe725260a3cdddc3ece9f8132c786c2291ffa6f)

<a id="canonical-9a19c7c88eed6a7a655520fb70180ac858a289742be5f8497055a556855b0686"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b25359c901e9f3b956d957c8e45c2a7b2a3b65518df601f42cf10d8be7926033"></a>

## aws — aws / ce2e79704f9c / 2

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md#canonical-e2bde7dd950eff273e290001dfe725260a3cdddc3ece9f8132c786c2291ffa6f)
- [Property reference](resources--cloud_link--reference--group-001.md#canonical-f9f572eb4dd4355ccf1d4d4a87719ef48acf695e515e7a8097f8aa7e0a96fa35)
- aws

<a id="canonical-42bb50301ad66808b11fcc68e8ed8382587a5ae1987542bc7181fce818529fdc"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: aws, gcp\] Amazon Web Services(AWS) CloudLink Provider. CloudLink for AWS Cloud Provider.

Upstream description:

CloudLink for AWS Cloud Provider.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-cloud_link_type": "[\"byoc\"]",
  "x-ves-oneof-field-direct_connect_gateway_asn_choice": "[\"custom_asn\"]"
}
```

OneOf alternatives in this subsection:

- [aws](resources--cloud_link--reference--group-001.md#canonical-42bb50301ad66808b11fcc68e8ed8382587a5ae1987542bc7181fce818529fdc)
- [gcp](resources--cloud_link--reference--group-001.md#canonical-08893a2fb138c2038185a105036b37180ec884db77427d7f03d89ff6bb94aa21)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
aws {
  # Configure direct properties listed below.
}
```

<a id="canonical-b1cf477fde143efcc3adb070c8fb082a7cff83449dbaaa073567213288199512"></a>

## Direct properties — aws / ce2e79704f9c / 3

- [aws_cred](resources--cloud_link--reference--group-001.md#canonical-3b58ee1146261908c0a7cac4ae26a9e3b507c4e3b1ad5a81cd7ed8245b23cee3): complete subsection reference.

- [byoc](resources--cloud_link--reference--group-001.md#canonical-afe5d15581ebc162c703a9ee3b5efdf26e44dde02c8ac2b97147d43a27718b62): complete subsection reference.

<a id="canonical-d1ad23ba41747eb34a58a268e234fcc73223ae7b82ee1d8d572635b72ca6e9f0"></a>

<a id="canonical-5d85ddb45473c5eaf7f3a69d86891b9d50ebaebc9ffa5a183b3cdc985e9a948a"></a>

## custom_asn property — aws / ce2e79704f9c / 4

Type: `"number"`. Optional.

Exclusive with \[\] F5XC will use custom ASN to create a Direct Connect Gateway.

Upstream description:

Exclusive with \[\] F5XC will use custom ASN to create a Direct Connect Gateway.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  validators.Int64RangeSetValidator(
    validators.Int64Range{Minimum: 64512, Maximum: 65534},
    validators.Int64Range{Minimum: 4200000000, Maximum: 4294967294},
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
    "maximum": 4294967294,
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
    "ves.io.schema.rules.uint32.ranges": "64512-65534, 4200000000-4294967294"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "64512-65534, 4200000000-4294967294"
  }
}
```

<a id="canonical-fa5b704d3feeae59bccfec7a5df2bd188f0dd511a6405150740c851b9a4698e3"></a>

## Next pages — aws / ce2e79704f9c / 5

- [aws.aws_cred](resources--cloud_link--reference--group-001.md#canonical-3b58ee1146261908c0a7cac4ae26a9e3b507c4e3b1ad5a81cd7ed8245b23cee3)
- [aws.byoc](resources--cloud_link--reference--group-001.md#canonical-afe5d15581ebc162c703a9ee3b5efdf26e44dde02c8ac2b97147d43a27718b62)
- [Property reference](resources--cloud_link--reference--group-001.md#canonical-f9f572eb4dd4355ccf1d4d4a87719ef48acf695e515e7a8097f8aa7e0a96fa35)
- [xcsh_cloud_link](../resources/cloud_link.md#canonical-e2bde7dd950eff273e290001dfe725260a3cdddc3ece9f8132c786c2291ffa6f)

<a id="canonical-3b58ee1146261908c0a7cac4ae26a9e3b507c4e3b1ad5a81cd7ed8245b23cee3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e08e340d4c0dc61887cee56c2a3a40c536ea6cce3982b1ca2b64da7071392a53"></a>

## aws.aws_cred — aws.aws_cred / 2ec3279df646 / 2

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md#canonical-e2bde7dd950eff273e290001dfe725260a3cdddc3ece9f8132c786c2291ffa6f)
- [Property reference](resources--cloud_link--reference--group-001.md#canonical-f9f572eb4dd4355ccf1d4d4a87719ef48acf695e515e7a8097f8aa7e0a96fa35)
- [aws](resources--cloud_link--reference--group-001.md#canonical-9a19c7c88eed6a7a655520fb70180ac858a289742be5f8497055a556855b0686)
- aws.aws_cred

<a id="canonical-45c98749a6f5a59251ffbd2346339d7c13f8527bb218146b9d83a1ea5a6b5275"></a>

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
aws_cred {
  # Configure direct properties listed below.
}
```

<a id="canonical-786e0bdd564c0e16b9b1db646a01f471c0562530787b4a1ab09f9359b0e8cd25"></a>

## Direct properties — aws.aws_cred / 2ec3279df646 / 3

<a id="canonical-aaeb3116b5e7f6346cfbf4ce0fbb2c3575fb82a44f90e1d4259b5b0c279b20d5"></a>

<a id="canonical-a424792984addf533091e9ecdc4f01b6d474677de45846cbb670131e2fbc363e"></a>

## name property — aws.aws_cred / 2ec3279df646 / 4

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

<a id="canonical-ed2a96c11e5185eed6aec214582a6c569a1624a8dc03027df4c3197d9637577f"></a>

<a id="canonical-17c26f8a7782f31da88735fadde25abb086a57f0dcb2b22ba2bf037dd7660851"></a>

## namespace property — aws.aws_cred / 2ec3279df646 / 5

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

<a id="canonical-73530f86c61659897a203ab01b3470e1a791466604f0f27e84a94d494f078ae0"></a>

<a id="canonical-7163046c2e4e338371b7c4b218367b3b14e8bc4d9cef03d4f9a49caf83c369fa"></a>

## tenant property — aws.aws_cred / 2ec3279df646 / 6

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

<a id="canonical-b975cbf3eefea4c98f104c33a0f2a8882a3d375d0842c221497d7bc2f244c9ff"></a>

## Next pages — aws.aws_cred / 2ec3279df646 / 7

- [aws](resources--cloud_link--reference--group-001.md#canonical-9a19c7c88eed6a7a655520fb70180ac858a289742be5f8497055a556855b0686)
- [xcsh_cloud_link](../resources/cloud_link.md#canonical-e2bde7dd950eff273e290001dfe725260a3cdddc3ece9f8132c786c2291ffa6f)

<a id="canonical-afe5d15581ebc162c703a9ee3b5efdf26e44dde02c8ac2b97147d43a27718b62"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-937ec31c57b028f60d35730b9074b32473c37df9800e37c53fff1ef90c180f15"></a>

## aws.byoc — aws.byoc / 6d9fd148440b / 2

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md#canonical-e2bde7dd950eff273e290001dfe725260a3cdddc3ece9f8132c786c2291ffa6f)
- [Property reference](resources--cloud_link--reference--group-001.md#canonical-f9f572eb4dd4355ccf1d4d4a87719ef48acf695e515e7a8097f8aa7e0a96fa35)
- [aws](resources--cloud_link--reference--group-001.md#canonical-9a19c7c88eed6a7a655520fb70180ac858a289742be5f8497055a556855b0686)
- aws.byoc

<a id="canonical-f35ea3c1aee90dba1810e6f70da5b3ed2795f5e5ec1f56fcb0000abd710659a2"></a>

Type: `"object"`. single nested block, Optional.

Bring Your Own Connections. List of Bring You Own Connection.

Upstream description:

List of Bring You Own Connection.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("connections")}
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
byoc {
  # Configure direct properties listed below.
}
```

<a id="canonical-a7c30c036f0fd58d2bf07913746c847cc097486747787063d3f0cc28c2543fd3"></a>

## Direct properties — aws.byoc / 6d9fd148440b / 3

- [connections](resources--cloud_link--reference--group-001.md#canonical-4b72943812ee33c937d369fe92c2fa132f37fa6358e94366197f044bfdb30981): complete subsection reference.

<a id="canonical-4570641bec9c40b251d5ee4ee6b8115d84f2affe99ca086fc34f7dc6f7039f02"></a>

## Next pages — aws.byoc / 6d9fd148440b / 4

- [aws.byoc.connections](resources--cloud_link--reference--group-001.md#canonical-4b72943812ee33c937d369fe92c2fa132f37fa6358e94366197f044bfdb30981)
- [aws](resources--cloud_link--reference--group-001.md#canonical-9a19c7c88eed6a7a655520fb70180ac858a289742be5f8497055a556855b0686)
- [xcsh_cloud_link](../resources/cloud_link.md#canonical-e2bde7dd950eff273e290001dfe725260a3cdddc3ece9f8132c786c2291ffa6f)

<a id="canonical-4b72943812ee33c937d369fe92c2fa132f37fa6358e94366197f044bfdb30981"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-77c87f1a7ba959f826331c7800c6eaaa2607672f4e7ac2def9b4746d85e105c1"></a>

## aws.byoc.connections — aws.byoc.connections / 572246e91c18 / 2

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md#canonical-e2bde7dd950eff273e290001dfe725260a3cdddc3ece9f8132c786c2291ffa6f)
- [Property reference](resources--cloud_link--reference--group-001.md#canonical-f9f572eb4dd4355ccf1d4d4a87719ef48acf695e515e7a8097f8aa7e0a96fa35)
- [aws](resources--cloud_link--reference--group-001.md#canonical-9a19c7c88eed6a7a655520fb70180ac858a289742be5f8497055a556855b0686)
- [aws.byoc](resources--cloud_link--reference--group-001.md#canonical-afe5d15581ebc162c703a9ee3b5efdf26e44dde02c8ac2b97147d43a27718b62)
- aws.byoc.connections

<a id="canonical-87a7a6424812cb22d05880c53ed32bc53a91f58784dd512c9cadfc836d90502f"></a>

Type: `"object"`. list nested block, Optional.

List of Bring You Own Connections. These AWS Direct Connect connections are not managed by F5XC but
will be used for connecting sites and REs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("bgp_asn",
    "connection_id",
    "region",
    "vlan"),
  validators.ConflictingListObjectAttributes("system_generated_name",
    "user_assigned_name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 10,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 10,
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
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
connections {
  # Configure direct properties listed below.
}
```

<a id="canonical-5a86fe2db11ac22293613101439acf1208c50b9b992b724daaa1029ce2f50744"></a>

## Direct properties — aws.byoc.connections / 572246e91c18 / 3

- [auth_key](resources--cloud_link--reference--group-001.md#canonical-f6e476d3b7064c8ab9052ff1458f0a5324400bad8658a444927b999985789e0d): complete subsection reference.

<a id="canonical-0e9db001357088d73d65dbc3def181e894d602bacf6141d8d3b66e5c2b501bc6"></a>

<a id="canonical-502c840519b25894e99197e7e3529acb09786793e804772dfb91a775b5a208a7"></a>

## bgp_asn property — aws.byoc.connections / 572246e91c18 / 4

Type: `"number"`. Optional.

The Border Gateway Protocol (BGP) Autonomous System Number (ASN) of your on-premises router for the
new virtual interface to be configured on AWS.

Upstream description:

The Border Gateway Protocol (BGP) Autonomous System Number (ASN) of your on-premises router for the
new virtual interface to be configured on AWS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 2147483647),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2147483647,
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
    "ves.io.schema.rules.uint32.lte": "2147483647"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  }
}
```

<a id="canonical-77da4421e8990d7681532bdb36aacf9b2bb583ff9fe8aabaca22a26bfe652e8e"></a>

<a id="canonical-8e2c15de3488df24daac501cd80fef4cab99f81662b3067cd722e5df58e490d3"></a>

## connection_id property — aws.byoc.connections / 572246e91c18 / 5

Type: `"string"`. Optional.

ID of the existing AWS Direct Connect Connection.

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
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^(dxcon-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(dxcon-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(dxcon-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

- [ipv4](resources--cloud_link--reference--group-001.md#canonical-322d3c9c1646fb6ce3bce45f4e7e9ec0d0ee273ae5d3ab2faf21df4cf836ffc8): complete subsection reference.

- [metadata](resources--cloud_link--reference--group-001.md#canonical-d13db5818d86beb05a5b14b868bce2c51fc4d9581869af4546016029204548bf): complete subsection reference.

<a id="canonical-7c12f02a3e90dcd9a81a129af338a0fd359176938009d5e14219c8183e98532b"></a>

<a id="canonical-13eafcf36caa12a42962ed7701ecf88657d3c3d9fa8a0f320c9bc7981e74db20"></a>

## region property — aws.byoc.connections / 572246e91c18 / 6

Type: `"string"`. Optional.

\[Enum:
ap-northeast-1|ap-southeast-1|eu-central-1|eu-west-1|eu-west-3|sa-east-1|us-east-1|us-east-2|us-west-2|ca-central-1|af-south-1|ap-east-1|ap-south-1|ap-northeast-2|ap-southeast-2|eu-south-1|eu-north-1|eu-west-2|me-south-1|us-west-1|ap-southeast-3\]
Region. Region where the connection is setup. Possible values are \`ap-northeast-1\`,
\`ap-southeast-1\`, \`eu-central-1\`, \`eu-west-1\`, \`eu-west-3\`, \`sa-east-1\`, \`us-east-1\`,
\`us-east-2\`, \`us-west-2\`, \`ca-central-1\`, \`af-south-1\`, \`ap-east-1\`, \`ap-south-1\`,
\`ap-northeast-2\`, \`ap-southeast-2\`, \`eu-south-1\`, \`eu-north-1\`, \`eu-west-2\`,
\`me-south-1\`, \`us-west-1\`, \`ap-southeast-3\`.

Upstream description:

Region where the connection is setup.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ap-northeast-1",
    "ap-southeast-1",
    "eu-central-1",
    "eu-west-1",
    "eu-west-3",
    "sa-east-1",
    "us-east-1",
    "us-east-2",
    "us-west-2",
    "ca-central-1",
    "af-south-1",
    "ap-east-1",
    "ap-south-1",
    "ap-northeast-2",
    "ap-southeast-2",
    "eu-south-1",
    "eu-north-1",
    "eu-west-2",
    "me-south-1",
    "us-west-1",
    "ap-southeast-3"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "ap-northeast-1",
    "ap-southeast-1",
    "eu-central-1",
    "eu-west-1",
    "eu-west-3",
    "sa-east-1",
    "us-east-1",
    "us-east-2",
    "us-west-2",
    "ca-central-1",
    "af-south-1",
    "ap-east-1",
    "ap-south-1",
    "ap-northeast-2",
    "ap-southeast-2",
    "eu-south-1",
    "eu-north-1",
    "eu-west-2",
    "me-south-1",
    "us-west-1",
    "ap-southeast-3"
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"ap-northeast-1\\\",\\\"ap-southeast-1\\\",\\\"eu-central-1\\\",\\\"eu-west-1\\\",\\\"eu-west-3\\\",\\\"sa-east-1\\\",\\\"us-east-1\\\",\\\"us-east-2\\\",\\\"us-west-2\\\",\\\"ca-central-1\\\",\\\"af-south-1\\\",\\\"ap-east-1\\\",\\\"ap-south-1\\\",\\\"ap-northeast-2\\\",\\\"ap-southeast-2\\\",\\\"eu-south-1\\\",\\\"eu-north-1\\\",\\\"eu-west-2\\\",\\\"me-south-1\\\",\\\"us-west-1\\\",\\\"ap-southeast-3\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"ap-northeast-1\\\",\\\"ap-southeast-1\\\",\\\"eu-central-1\\\",\\\"eu-west-1\\\",\\\"eu-west-3\\\",\\\"sa-east-1\\\",\\\"us-east-1\\\",\\\"us-east-2\\\",\\\"us-west-2\\\",\\\"ca-central-1\\\",\\\"af-south-1\\\",\\\"ap-east-1\\\",\\\"ap-south-1\\\",\\\"ap-northeast-2\\\",\\\"ap-southeast-2\\\",\\\"eu-south-1\\\",\\\"eu-north-1\\\",\\\"eu-west-2\\\",\\\"me-south-1\\\",\\\"us-west-1\\\",\\\"ap-southeast-3\\\"]"
  }
}
```

- [system_generated_name](resources--cloud_link--reference--group-001.md#canonical-d66ca7488d949f476c00c361d6f84e91dbd27e722cafc9ea685df0ffafb31ab9): complete subsection reference.

<a id="canonical-21faf9b2f9817d2f17d92896a70bb5f437b2a7b79fcdafd1167e55a3f2a0c74d"></a>

<a id="canonical-99790088a115e3e9dd8800807cbfe0a8bdbd78616142fc6e5e7f0bbd1351d7bf"></a>

## tags property — aws.byoc.connections / 572246e91c18 / 7

Type: `["map", "string"]`. Optional.

AWS Tags is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in AWS console. Specified tags will be added to Virtual
interface along with any F5XC specific tags.

Upstream description:

AWS Tags is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in AWS console. Specified tags will be added to Virtual
interface along with any F5XC specific tags.

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
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "256",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "256",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-9a879d834611c3b9e1dd36590c7f2ea945cd7c8f6db1c8a9f9aff787e936b592"></a>

<a id="canonical-67789ddc3aab9ddc5f8eaee69d094e517a110db657f24e91138163b64bf03b35"></a>

## user_assigned_name property — aws.byoc.connections / 572246e91c18 / 8

Type: `"string"`. Optional.

Exclusive with \[system\_generated\_name\] User is managing the AWS resource name.

Upstream description:

Exclusive with \[system\_generated\_name\] User is managing the AWS resource name.

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

<a id="canonical-5c4dfb41a6fb176b1a1db82552e7b752c536fc7e26ee0dd4fd56ede292191e32"></a>

<a id="canonical-9f52c2c0c2c845a8724bde71b045bb5c8becf87cfa5df5bccdb1f0c6aa4a8768"></a>

## virtual_interface_type property — aws.byoc.connections / 572246e91c18 / 9

Type: `"string"`. Optional.

\[Enum: PRIVATE\] Defines the type of virtual interface that needs to be configured on AWS -
PRIVATE: Private A private virtual interface should be used to access an Amazon VPC using private IP
addresses. - TRANSIT: Transit A transit virtual interface is a VLAN that transports traffic from a
Direct Connect.. The only possible value is \`PRIVATE\`. Defaults to \`PRIVATE\`.

Upstream description:

Defines the type of virtual interface that needs to be configured on AWS

&#8203;- PRIVATE: Private

A private virtual interface should be used to access an Amazon VPC using private IP addresses.
&#8203;- TRANSIT: Transit

A transit virtual interface is a VLAN that transports traffic from a Direct Connect gateway to one
or more transit gateways.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("PRIVATE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "PRIVATE",
  "enum": [
    "PRIVATE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-512beac77628a844a61fa0495db7ceb6f801a5d6ac7f82ebbdfe5db2ed050a96"></a>

<a id="canonical-f2e9aed219657800906df61e9596059510e4513c38763d406b073c2e03154393"></a>

## vlan property — aws.byoc.connections / 572246e91c18 / 10

Type: `"number"`. Optional.

Virtual Local Area Network number for the new virtual interface to be configured on the AWS. This
tag is required for any traffic traversing the AWS Direct Connect connection.

Upstream description:

Virtual Local Area Network number for the new virtual interface to be configured on the AWS. This
tag is required for any traffic traversing the AWS Direct Connect connection.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 4094),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 4094,
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
    "ves.io.schema.rules.uint32.lte": "4094"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "4094"
  }
}
```

<a id="canonical-e1c3a80e0aa9754c9d009a55a0fed10a81865b29c469efe8e8cebcc492586f92"></a>

## Next pages — aws.byoc.connections / 572246e91c18 / 11

- [aws.byoc.connections.auth_key](resources--cloud_link--reference--group-001.md#canonical-f6e476d3b7064c8ab9052ff1458f0a5324400bad8658a444927b999985789e0d)
- [aws.byoc.connections.ipv4](resources--cloud_link--reference--group-001.md#canonical-322d3c9c1646fb6ce3bce45f4e7e9ec0d0ee273ae5d3ab2faf21df4cf836ffc8)
- [aws.byoc.connections.metadata](resources--cloud_link--reference--group-001.md#canonical-d13db5818d86beb05a5b14b868bce2c51fc4d9581869af4546016029204548bf)
- [aws.byoc.connections.system_generated_name](resources--cloud_link--reference--group-001.md#canonical-d66ca7488d949f476c00c361d6f84e91dbd27e722cafc9ea685df0ffafb31ab9)
- [aws.byoc](resources--cloud_link--reference--group-001.md#canonical-afe5d15581ebc162c703a9ee3b5efdf26e44dde02c8ac2b97147d43a27718b62)
- [xcsh_cloud_link](../resources/cloud_link.md#canonical-e2bde7dd950eff273e290001dfe725260a3cdddc3ece9f8132c786c2291ffa6f)

<a id="canonical-f6e476d3b7064c8ab9052ff1458f0a5324400bad8658a444927b999985789e0d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-661d72c7b310d709ff467b6ae363de1d09c88a4ec1ac1313e9ce3df3084890c2"></a>

## aws.byoc.connections.auth_key — aws.byoc.connections.auth_key / 35558ff51088 / 2

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md#canonical-e2bde7dd950eff273e290001dfe725260a3cdddc3ece9f8132c786c2291ffa6f)
- [Property reference](resources--cloud_link--reference--group-001.md#canonical-f9f572eb4dd4355ccf1d4d4a87719ef48acf695e515e7a8097f8aa7e0a96fa35)
- [aws](resources--cloud_link--reference--group-001.md#canonical-9a19c7c88eed6a7a655520fb70180ac858a289742be5f8497055a556855b0686)
- [aws.byoc](resources--cloud_link--reference--group-001.md#canonical-afe5d15581ebc162c703a9ee3b5efdf26e44dde02c8ac2b97147d43a27718b62)
- [aws.byoc.connections](resources--cloud_link--reference--group-001.md#canonical-4b72943812ee33c937d369fe92c2fa132f37fa6358e94366197f044bfdb30981)
- aws.byoc.connections.auth_key

<a id="canonical-9b77a033e3c5355f37a038aac22960886a6b93ed79e2b62abe04c6755da6da2b"></a>

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
auth_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-4ed71eda7f7e2008625a70dd33d7d797c6d4a763b44bdcd8c61a4a5a1e0f3ee5"></a>

## Direct properties — aws.byoc.connections.auth_key / 35558ff51088 / 3

- [blindfold_secret_info](resources--cloud_link--reference--group-001.md#canonical-519de8755a41bee52040be8516f5e1a2525579ad32de0cbf101ad375f901c744): complete subsection reference.

- [clear_secret_info](resources--cloud_link--reference--group-001.md#canonical-b9106c9dc7b1cd359d94937c41d867345ded311501d9478a4eb90392111df36c): complete subsection reference.

<a id="canonical-5f466c8f55c163f0834d6a80ae92d2a9ec6b85c227b8d7e88d08c4678e472881"></a>

## Next pages — aws.byoc.connections.auth_key / 35558ff51088 / 4

- [aws.byoc.connections.auth_key.blindfold_secret_info](resources--cloud_link--reference--group-001.md#canonical-519de8755a41bee52040be8516f5e1a2525579ad32de0cbf101ad375f901c744)
- [aws.byoc.connections.auth_key.clear_secret_info](resources--cloud_link--reference--group-001.md#canonical-b9106c9dc7b1cd359d94937c41d867345ded311501d9478a4eb90392111df36c)
- [aws.byoc.connections](resources--cloud_link--reference--group-001.md#canonical-4b72943812ee33c937d369fe92c2fa132f37fa6358e94366197f044bfdb30981)
- [xcsh_cloud_link](../resources/cloud_link.md#canonical-e2bde7dd950eff273e290001dfe725260a3cdddc3ece9f8132c786c2291ffa6f)

<a id="canonical-519de8755a41bee52040be8516f5e1a2525579ad32de0cbf101ad375f901c744"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e0ce072d8fe0189fc1ac838d8bec1dfb05a465d0b14333b4f41e82db0fa95598"></a>

## aws.byoc.connections.auth_key.blindfold_secret_info — aws.byoc.connections.auth_key.blindfold_secret_info / 84f87328a380 / 2

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md#canonical-e2bde7dd950eff273e290001dfe725260a3cdddc3ece9f8132c786c2291ffa6f)
- [Property reference](resources--cloud_link--reference--group-001.md#canonical-f9f572eb4dd4355ccf1d4d4a87719ef48acf695e515e7a8097f8aa7e0a96fa35)
- [aws](resources--cloud_link--reference--group-001.md#canonical-9a19c7c88eed6a7a655520fb70180ac858a289742be5f8497055a556855b0686)
- [aws.byoc](resources--cloud_link--reference--group-001.md#canonical-afe5d15581ebc162c703a9ee3b5efdf26e44dde02c8ac2b97147d43a27718b62)
- [aws.byoc.connections](resources--cloud_link--reference--group-001.md#canonical-4b72943812ee33c937d369fe92c2fa132f37fa6358e94366197f044bfdb30981)
- [aws.byoc.connections.auth_key](resources--cloud_link--reference--group-001.md#canonical-f6e476d3b7064c8ab9052ff1458f0a5324400bad8658a444927b999985789e0d)
- aws.byoc.connections.auth_key.blindfold_secret_info

<a id="canonical-2f6b6d55f74c166aea45299ce41efa21fef327268eff5146aaf62a3294924259"></a>

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

<a id="canonical-cbc56a853cd42a91e8952fc55965d7ea806cbb93aad8df12d18cb8692ba0c0b2"></a>

## Direct properties — aws.byoc.connections.auth_key.blindfold_secret_info / 84f87328a380 / 3

<a id="canonical-5a0c6530c8f42483f4437d61140c4a6aabdbade33c484f953bcf1de95d5ddf85"></a>

<a id="canonical-f9b2c24b49e754c81b2e93762c44e6afb056cd1f2c9a11f8e95b0a6dcf3ae91d"></a>

## decryption_provider property — aws.byoc.connections.auth_key.blindfold_secret_info / 84f87328a380 / 4

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

<a id="canonical-23b13746229b0663c0841af90cd16ede2de6a778f635e3d6deddf5d7adf2f015"></a>

<a id="canonical-4205309abc861c441c535f97c8b3e2c610fdcb4b0e83b0c94c2c6a1c8f0ac9ac"></a>

## location property — aws.byoc.connections.auth_key.blindfold_secret_info / 84f87328a380 / 5

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

<a id="canonical-5a6194b516f8734b56b941a6b03081b1b0ba309bba273dbbda205d980a58d9c3"></a>

<a id="canonical-abc8f3c5a192805504b11785ee488221a16d5945e3f83cf156bce86eb9c9f1aa"></a>

## store_provider property — aws.byoc.connections.auth_key.blindfold_secret_info / 84f87328a380 / 6

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

<a id="canonical-52627d1a5d55860f57c89f5962d89358be96415499fa8108398f2749de156401"></a>

## Next pages — aws.byoc.connections.auth_key.blindfold_secret_info / 84f87328a380 / 7

- [aws.byoc.connections.auth_key](resources--cloud_link--reference--group-001.md#canonical-f6e476d3b7064c8ab9052ff1458f0a5324400bad8658a444927b999985789e0d)
- [xcsh_cloud_link](../resources/cloud_link.md#canonical-e2bde7dd950eff273e290001dfe725260a3cdddc3ece9f8132c786c2291ffa6f)

<a id="canonical-b9106c9dc7b1cd359d94937c41d867345ded311501d9478a4eb90392111df36c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f5416e184e5e0cebb8d581a30dfe7eb1c54b5f295bd4b452b02d1b24ce50842a"></a>

## aws.byoc.connections.auth_key.clear_secret_info — aws.byoc.connections.auth_key.clear_secret_info / c0be9e4592c9 / 2

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md#canonical-e2bde7dd950eff273e290001dfe725260a3cdddc3ece9f8132c786c2291ffa6f)
- [Property reference](resources--cloud_link--reference--group-001.md#canonical-f9f572eb4dd4355ccf1d4d4a87719ef48acf695e515e7a8097f8aa7e0a96fa35)
- [aws](resources--cloud_link--reference--group-001.md#canonical-9a19c7c88eed6a7a655520fb70180ac858a289742be5f8497055a556855b0686)
- [aws.byoc](resources--cloud_link--reference--group-001.md#canonical-afe5d15581ebc162c703a9ee3b5efdf26e44dde02c8ac2b97147d43a27718b62)
- [aws.byoc.connections](resources--cloud_link--reference--group-001.md#canonical-4b72943812ee33c937d369fe92c2fa132f37fa6358e94366197f044bfdb30981)
- [aws.byoc.connections.auth_key](resources--cloud_link--reference--group-001.md#canonical-f6e476d3b7064c8ab9052ff1458f0a5324400bad8658a444927b999985789e0d)
- aws.byoc.connections.auth_key.clear_secret_info

<a id="canonical-79361ecd727ed040c21a0d317213a816e4a1e907b4fc7332162aaa0eb5953552"></a>

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

<a id="canonical-21d93c656103fa80221f301c56cbc41f813d1bbce9e1696ead871e9b49742a30"></a>

## Direct properties — aws.byoc.connections.auth_key.clear_secret_info / c0be9e4592c9 / 3

<a id="canonical-fa7a83498b7712fd99a2758e68858e0d54c0fb77f796ace77393443b51f7a62f"></a>

<a id="canonical-e2d3ed884e2450df16f5b7cc1895958503e0b535c72d2dfe9fd095fb4832d6aa"></a>

## provider_ref property — aws.byoc.connections.auth_key.clear_secret_info / c0be9e4592c9 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2447158e18137c3b4e950e020d83a698393d7f1980d97d942fb36d0c534c319c"></a>

<a id="canonical-c1dd8dab6f99082b885c66cd81da0dff98a8506e4d2124b3e9e6803d27f0cb81"></a>

## url property — aws.byoc.connections.auth_key.clear_secret_info / c0be9e4592c9 / 5

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

<a id="canonical-7bc50cc3d8ebe9cbb014f82c5018746e837697aa7ee12cd78412ca6284d9101a"></a>

## Next pages — aws.byoc.connections.auth_key.clear_secret_info / c0be9e4592c9 / 6

- [aws.byoc.connections.auth_key](resources--cloud_link--reference--group-001.md#canonical-f6e476d3b7064c8ab9052ff1458f0a5324400bad8658a444927b999985789e0d)
- [xcsh_cloud_link](../resources/cloud_link.md#canonical-e2bde7dd950eff273e290001dfe725260a3cdddc3ece9f8132c786c2291ffa6f)

<a id="canonical-322d3c9c1646fb6ce3bce45f4e7e9ec0d0ee273ae5d3ab2faf21df4cf836ffc8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f6ca018229180a27a21ac40bde693ff80b828fd9b24f8b63f4ae54cea652733e"></a>

## aws.byoc.connections.ipv4 — aws.byoc.connections.ipv4 / 82c061f2660a / 2

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md#canonical-e2bde7dd950eff273e290001dfe725260a3cdddc3ece9f8132c786c2291ffa6f)
- [Property reference](resources--cloud_link--reference--group-001.md#canonical-f9f572eb4dd4355ccf1d4d4a87719ef48acf695e515e7a8097f8aa7e0a96fa35)
- [aws](resources--cloud_link--reference--group-001.md#canonical-9a19c7c88eed6a7a655520fb70180ac858a289742be5f8497055a556855b0686)
- [aws.byoc](resources--cloud_link--reference--group-001.md#canonical-afe5d15581ebc162c703a9ee3b5efdf26e44dde02c8ac2b97147d43a27718b62)
- [aws.byoc.connections](resources--cloud_link--reference--group-001.md#canonical-4b72943812ee33c937d369fe92c2fa132f37fa6358e94366197f044bfdb30981)
- aws.byoc.connections.ipv4

<a id="canonical-8eeababf30a295bb72d456c01cbadcbaf8b3437a96581896238cb0745e8aca6e"></a>

Type: `"object"`. single nested block, Optional.

Configure BGP IPv4 peering for endpoints.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("aws_router_peer_address",
    "router_peer_address")}
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
ipv4 {
  # Configure direct properties listed below.
}
```

<a id="canonical-7a27ae442eddce4f3b852f14906671a50519fb455c706e5f101d4e29ac6d33b3"></a>

## Direct properties — aws.byoc.connections.ipv4 / 82c061f2660a / 3

<a id="canonical-16879842cd10cd862563a3c7f66948e58dacc13f63b19e9c2de83f239b20469b"></a>

<a id="canonical-6d2ad142abf466fdac0b7bcae9122510ec3544600d1ab03e6e688a8949dd9a24"></a>

## aws_router_peer_address property — aws.byoc.connections.ipv4 / 82c061f2660a / 4

Type: `"string"`. Optional.

The BGP peer IP configured on the AWS endpoint.

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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "32",
    "ves.io.schema.rules.string.min_ip_prefix_length": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "32",
    "ves.io.schema.rules.string.min_ip_prefix_length": "1"
  }
}
```

<a id="canonical-26d20bedfa771a790472401799bc2b4f773fc98486fc12205411d2eb6028254f"></a>

<a id="canonical-fffd4f419459268fec2bfce8b879dba31762bc7ce7416272e7639ce1fc3215f7"></a>

## router_peer_address property — aws.byoc.connections.ipv4 / 82c061f2660a / 5

Type: `"string"`. Optional.

The BGP peer IP configured on your (customer) endpoint.

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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "32",
    "ves.io.schema.rules.string.min_ip_prefix_length": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "32",
    "ves.io.schema.rules.string.min_ip_prefix_length": "1"
  }
}
```

<a id="canonical-0080c43fb9a912338270428751410e6f9136c47c240733f794786ace0bdb20d0"></a>

## Next pages — aws.byoc.connections.ipv4 / 82c061f2660a / 6

- [aws.byoc.connections](resources--cloud_link--reference--group-001.md#canonical-4b72943812ee33c937d369fe92c2fa132f37fa6358e94366197f044bfdb30981)
- [xcsh_cloud_link](../resources/cloud_link.md#canonical-e2bde7dd950eff273e290001dfe725260a3cdddc3ece9f8132c786c2291ffa6f)

<a id="canonical-d13db5818d86beb05a5b14b868bce2c51fc4d9581869af4546016029204548bf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9ef57b72bff09f7d717f16e5d822f345aeaa28dd16d7ac3c74d4d5b3fc3929cd"></a>

## aws.byoc.connections.metadata — aws.byoc.connections.metadata / 533524049d2e / 2

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md#canonical-e2bde7dd950eff273e290001dfe725260a3cdddc3ece9f8132c786c2291ffa6f)
- [Property reference](resources--cloud_link--reference--group-001.md#canonical-f9f572eb4dd4355ccf1d4d4a87719ef48acf695e515e7a8097f8aa7e0a96fa35)
- [aws](resources--cloud_link--reference--group-001.md#canonical-9a19c7c88eed6a7a655520fb70180ac858a289742be5f8497055a556855b0686)
- [aws.byoc](resources--cloud_link--reference--group-001.md#canonical-afe5d15581ebc162c703a9ee3b5efdf26e44dde02c8ac2b97147d43a27718b62)
- [aws.byoc.connections](resources--cloud_link--reference--group-001.md#canonical-4b72943812ee33c937d369fe92c2fa132f37fa6358e94366197f044bfdb30981)
- aws.byoc.connections.metadata

<a id="canonical-f46f7a4cb3a480809f2ad3b94fee07db43806a45e6fd41058d9583fbdb4a86e5"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-f1a5e58a90790e11ee023299d694884cdbd29afdd937fcb8e2d040bb17d99a1b"></a>

## Direct properties — aws.byoc.connections.metadata / 533524049d2e / 3

<a id="canonical-59503ac7c22e1bffbdd079366ca0b73a521dc0dcb022410b492890010a662171"></a>

<a id="canonical-c853af8427c0a881067716b71fae41b84362d1150229fbafba2c44bda12cd7d2"></a>

## description_spec property — aws.byoc.connections.metadata / 533524049d2e / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-78744c5a8ec6c0adabb611c64f8387c6f5ccb9a9fbca2def0beef751a87a3606"></a>

<a id="canonical-34b5a59e50f1f0b694aeec140b8d594bec2e868848de5d7074fbfc9048ef07f8"></a>

## name property — aws.byoc.connections.metadata / 533524049d2e / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
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
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-22e3797be572c25649b65fc93bbad92726733fbabb42884f36f30748bff18e27"></a>

## Next pages — aws.byoc.connections.metadata / 533524049d2e / 6

- [aws.byoc.connections](resources--cloud_link--reference--group-001.md#canonical-4b72943812ee33c937d369fe92c2fa132f37fa6358e94366197f044bfdb30981)
- [xcsh_cloud_link](../resources/cloud_link.md#canonical-e2bde7dd950eff273e290001dfe725260a3cdddc3ece9f8132c786c2291ffa6f)

<a id="canonical-d66ca7488d949f476c00c361d6f84e91dbd27e722cafc9ea685df0ffafb31ab9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-38847948b1393601c8db2bd3057fe566671524522e314a979a4b7f68298cafa3"></a>

## aws.byoc.connections.system_generated_name — aws.byoc.connections.system_generated_name / 17405a21946f / 2

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md#canonical-e2bde7dd950eff273e290001dfe725260a3cdddc3ece9f8132c786c2291ffa6f)
- [Property reference](resources--cloud_link--reference--group-001.md#canonical-f9f572eb4dd4355ccf1d4d4a87719ef48acf695e515e7a8097f8aa7e0a96fa35)
- [aws](resources--cloud_link--reference--group-001.md#canonical-9a19c7c88eed6a7a655520fb70180ac858a289742be5f8497055a556855b0686)
- [aws.byoc](resources--cloud_link--reference--group-001.md#canonical-afe5d15581ebc162c703a9ee3b5efdf26e44dde02c8ac2b97147d43a27718b62)
- [aws.byoc.connections](resources--cloud_link--reference--group-001.md#canonical-4b72943812ee33c937d369fe92c2fa132f37fa6358e94366197f044bfdb30981)
- aws.byoc.connections.system_generated_name

<a id="canonical-b48c67fe9c604a466fcff0dc9a41a2b02e18037e9a55a602e94e2922d5b2d5a8"></a>

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
system_generated_name = {}
```

<a id="canonical-44320358c0714493a59074aa8e794559882d840fc0321dbc28521ba23bdd024a"></a>

## Direct properties — aws.byoc.connections.system_generated_name / 17405a21946f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a6da4c4065b73d3e7abb00f028cc73694f43ba31e99d5c1ccf877047c009526b"></a>

## Next pages — aws.byoc.connections.system_generated_name / 17405a21946f / 4

- [aws.byoc.connections](resources--cloud_link--reference--group-001.md#canonical-4b72943812ee33c937d369fe92c2fa132f37fa6358e94366197f044bfdb30981)
- [xcsh_cloud_link](../resources/cloud_link.md#canonical-e2bde7dd950eff273e290001dfe725260a3cdddc3ece9f8132c786c2291ffa6f)

<a id="canonical-4a7582a37c7b581ae3e077db76831503a8dc9e2b39d7a5f66d2f268f2aa7411c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-51efe496fa800da973e989c4ddb8700097ebb5017f06b79a82dd8c05a981a6e2"></a>

## disabled — disabled / 29c8e7f16612 / 2

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md#canonical-e2bde7dd950eff273e290001dfe725260a3cdddc3ece9f8132c786c2291ffa6f)
- [Property reference](resources--cloud_link--reference--group-001.md#canonical-f9f572eb4dd4355ccf1d4d4a87719ef48acf695e515e7a8097f8aa7e0a96fa35)
- disabled

<a id="canonical-702d2f35a7f3d57f00d8eb6b000fce9c61dd97e4e23eae3e8c188c6a32ee707d"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disabled, enabled\] Enable this option

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

- [disabled](resources--cloud_link--reference--group-001.md#canonical-702d2f35a7f3d57f00d8eb6b000fce9c61dd97e4e23eae3e8c188c6a32ee707d)
- [enabled](resources--cloud_link--reference--group-001.md#canonical-4548c7067b83df24918124a71cae1f86634fbf6d2d693a4613e3f1ec52671905)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disabled = {}
```

<a id="canonical-740357aa96e130908a798cfe513d0a12f3cb920cf0bf7a795f0d740957ebdaae"></a>

## Direct properties — disabled / 29c8e7f16612 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c95ad10776b532843675cb3f42a00fd21bb06f6912ad34e799565aff86568fe8"></a>

## Next pages — disabled / 29c8e7f16612 / 4

- [Property reference](resources--cloud_link--reference--group-001.md#canonical-f9f572eb4dd4355ccf1d4d4a87719ef48acf695e515e7a8097f8aa7e0a96fa35)
- [xcsh_cloud_link](../resources/cloud_link.md#canonical-e2bde7dd950eff273e290001dfe725260a3cdddc3ece9f8132c786c2291ffa6f)

<a id="canonical-b8705c14c9d8e6bec591d18c74d04c6d9d70acb0c574cd6d29120416eed8a236"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-82a22be46893098802ddc1f76a96f1d6f43b59f725d013abf923188760d89c70"></a>

## enabled — enabled / 4c4c761b0f7b / 2

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md#canonical-e2bde7dd950eff273e290001dfe725260a3cdddc3ece9f8132c786c2291ffa6f)
- [Property reference](resources--cloud_link--reference--group-001.md#canonical-f9f572eb4dd4355ccf1d4d4a87719ef48acf695e515e7a8097f8aa7e0a96fa35)
- enabled

<a id="canonical-4548c7067b83df24918124a71cae1f86634fbf6d2d693a4613e3f1ec52671905"></a>

Type: `"object"`. single nested block, Optional.

CloudLink ADN Network Config.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cloudlink_network_name")}
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

<a id="canonical-0429d4e6508f9c3491bca64d4b3f81d788614bd2a020363731c93666a3f37558"></a>

## Direct properties — enabled / 4c4c761b0f7b / 3

<a id="canonical-d6717732848727f266a9a724dcede0a8df649107fb059bb503bf797bc8453578"></a>

<a id="canonical-459f9f1cb8c1f4ec832353aa19eb991f0ce86d8cef3b215c410bb6136f12ca95"></a>

## cloudlink_network_name property — enabled / 4c4c761b0f7b / 4

Type: `"string"`. Optional.

Establish private connectivity with the F5 Distributed Cloud Global Network using a Private ADN
network. To provision a Private ADN network, please contact F5 Distributed Cloud support.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-4eefe4d55d9e5e0d5453ab990092570f088b5337848540ca313651299fd584c7"></a>

## Next pages — enabled / 4c4c761b0f7b / 5

- [Property reference](resources--cloud_link--reference--group-001.md#canonical-f9f572eb4dd4355ccf1d4d4a87719ef48acf695e515e7a8097f8aa7e0a96fa35)
- [xcsh_cloud_link](../resources/cloud_link.md#canonical-e2bde7dd950eff273e290001dfe725260a3cdddc3ece9f8132c786c2291ffa6f)

<a id="canonical-98b04cd1a263fa94de21346302cdfec25d91055cbcf9c3b463608298b2f20fd9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9e122f648814bab3410685bd554e807167348cb4a337047840685cc9de4e4b7a"></a>

## gcp — gcp / 72557e49a1a8 / 2

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md#canonical-e2bde7dd950eff273e290001dfe725260a3cdddc3ece9f8132c786c2291ffa6f)
- [Property reference](resources--cloud_link--reference--group-001.md#canonical-f9f572eb4dd4355ccf1d4d4a87719ef48acf695e515e7a8097f8aa7e0a96fa35)
- gcp

<a id="canonical-08893a2fb138c2038185a105036b37180ec884db77427d7f03d89ff6bb94aa21"></a>

Type: `"object"`. single nested block, Optional.

Google Cloud Platform (GCP) CloudLink Provider. CloudLink for GCP Cloud Provider.

Upstream description:

CloudLink for GCP Cloud Provider.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-cloud_link_type": "[\"byoc\"]"
}
```

Terraform syntax:

```terraform
gcp {
  # Configure direct properties listed below.
}
```

<a id="canonical-14fd9345ba4d2a7caa9eb50620908cdb4f75314f63eecba4e6f3ed772e12054e"></a>

## Direct properties — gcp / 72557e49a1a8 / 3

- [byoc](resources--cloud_link--reference--group-001.md#canonical-e338639ca36e96d6bf0f5483fc7e1fc3e49b4946ef053d151ee966be7ed8847f): complete subsection reference.

- [gcp_cred](resources--cloud_link--reference--group-001.md#canonical-dbb5eb32808aa37a77eb7dad058d46d3db79e79b5044106c4ed056fc309480b2): complete subsection reference.

<a id="canonical-0bc5b314040777556207264a90d356262bbcb506c3c3dd3c6f263dbcb2c0af4d"></a>

## Next pages — gcp / 72557e49a1a8 / 4

- [gcp.byoc](resources--cloud_link--reference--group-001.md#canonical-e338639ca36e96d6bf0f5483fc7e1fc3e49b4946ef053d151ee966be7ed8847f)
- [gcp.gcp_cred](resources--cloud_link--reference--group-001.md#canonical-dbb5eb32808aa37a77eb7dad058d46d3db79e79b5044106c4ed056fc309480b2)
- [Property reference](resources--cloud_link--reference--group-001.md#canonical-f9f572eb4dd4355ccf1d4d4a87719ef48acf695e515e7a8097f8aa7e0a96fa35)
- [xcsh_cloud_link](../resources/cloud_link.md#canonical-e2bde7dd950eff273e290001dfe725260a3cdddc3ece9f8132c786c2291ffa6f)

<a id="canonical-e338639ca36e96d6bf0f5483fc7e1fc3e49b4946ef053d151ee966be7ed8847f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1c7282d713db202b8a18e7236963acbae7781fedce117e574ee4a1c60f3fc307"></a>

## gcp.byoc — gcp.byoc / 993ec1100ce9 / 2

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md#canonical-e2bde7dd950eff273e290001dfe725260a3cdddc3ece9f8132c786c2291ffa6f)
- [Property reference](resources--cloud_link--reference--group-001.md#canonical-f9f572eb4dd4355ccf1d4d4a87719ef48acf695e515e7a8097f8aa7e0a96fa35)
- [gcp](resources--cloud_link--reference--group-001.md#canonical-98b04cd1a263fa94de21346302cdfec25d91055cbcf9c3b463608298b2f20fd9)
- gcp.byoc

<a id="canonical-23568d4584b2b0b92cea6555c8dd7fb9b802d469aff38bb9b5396847c6bfd981"></a>

Type: `"object"`. single nested block, Optional.

GCP Bring Your Own Connections. List of GCP Bring You Own Connections.

Upstream description:

List of GCP Bring You Own Connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("connections")}
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
byoc {
  # Configure direct properties listed below.
}
```

<a id="canonical-c6d74582bbe503b4edf0b78336b751eb82975a0c34d08cab504c7dee703078b0"></a>

## Direct properties — gcp.byoc / 993ec1100ce9 / 3

- [connections](resources--cloud_link--reference--group-001.md#canonical-f0c351486e5af898d9b72bce3bf1819e6398294226f768d6b8ad212511d8e621): complete subsection reference.

<a id="canonical-370f204ac568e408d98f88bc0ec571d693ec5b43925d2d0bea633140c3d118f6"></a>

## Next pages — gcp.byoc / 993ec1100ce9 / 4

- [gcp.byoc.connections](resources--cloud_link--reference--group-001.md#canonical-f0c351486e5af898d9b72bce3bf1819e6398294226f768d6b8ad212511d8e621)
- [gcp](resources--cloud_link--reference--group-001.md#canonical-98b04cd1a263fa94de21346302cdfec25d91055cbcf9c3b463608298b2f20fd9)
- [xcsh_cloud_link](../resources/cloud_link.md#canonical-e2bde7dd950eff273e290001dfe725260a3cdddc3ece9f8132c786c2291ffa6f)

<a id="canonical-f0c351486e5af898d9b72bce3bf1819e6398294226f768d6b8ad212511d8e621"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-285dc0a8a297af559a5d1694100e169b6250d30bac738d205e753c79b042145b"></a>

## gcp.byoc.connections — gcp.byoc.connections / a7129f601f98 / 2

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md#canonical-e2bde7dd950eff273e290001dfe725260a3cdddc3ece9f8132c786c2291ffa6f)
- [Property reference](resources--cloud_link--reference--group-001.md#canonical-f9f572eb4dd4355ccf1d4d4a87719ef48acf695e515e7a8097f8aa7e0a96fa35)
- [gcp](resources--cloud_link--reference--group-001.md#canonical-98b04cd1a263fa94de21346302cdfec25d91055cbcf9c3b463608298b2f20fd9)
- [gcp.byoc](resources--cloud_link--reference--group-001.md#canonical-e338639ca36e96d6bf0f5483fc7e1fc3e49b4946ef053d151ee966be7ed8847f)
- gcp.byoc.connections

<a id="canonical-7d6abe8456acd336dc26952be25799df2fd11c0ba5c781ee124c2a976d248e8e"></a>

Type: `"object"`. list nested block, Optional.

Each 'Bring Your Own Connection' represents a virtual connection that the customer has provisioned
in the Cloud (.

Upstream description:

Each 'Bring Your Own Connection' represents a virtual connection that the customer has provisioned
in the Cloud (example: AWS Direct Connect). F5XC will orchestrate networking resources in the cloud
to facilitate seamless private connectivity.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("interconnect_attachment_name",
    "region"),
  validators.ConflictingListObjectAttributes("project",
    "same_as_credential")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 10,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 10,
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
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
connections {
  # Configure direct properties listed below.
}
```

<a id="canonical-8d9e601dc48718e7c9f3fb00b2ce87957726c7f3d7336412ebec882b4f285ba8"></a>

## Direct properties — gcp.byoc.connections / a7129f601f98 / 3

<a id="canonical-78edf06e2eb2494bd34f87e3a5a47f1b5f92f40908aabe5d8d726d11dbfec194"></a>

<a id="canonical-3f7ce795b509ccea805dcb0f90b5479783e4c62c0adf5c3426ced1e0d3c3eb1d"></a>

## interconnect_attachment_name property — gcp.byoc.connections / a7129f601f98 / 4

Type: `"string"`. Optional.

Name of already-existing GCP Cloud Interconnect Attachment.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 63,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 63,
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
    "ves.io.schema.rules.string.max_len": "63",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "63",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [metadata](resources--cloud_link--reference--group-001.md#canonical-6d3d50c6cb675f2757193ac5fe44e800f9f5cb7b0ba899e7620a6b1dc59b14e4): complete subsection reference.

<a id="canonical-4017db3e815ce6d06028c125f3b2e24d85099b2c4d564f600f2967961d8e8b64"></a>

<a id="canonical-8fb765f3063d7f1a41182dc0f5c75ba06bee05ecb4683423740997ad78ac7ef4"></a>

## project property — gcp.byoc.connections / a7129f601f98 / 5

Type: `"string"`. Optional.

Exclusive with \[same\_as\_credential\] Specify a GCP Project for the interconnect attachment.

Upstream description:

Exclusive with \[same\_as\_credential\] Specify a GCP Project for the interconnect attachment.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 30),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 30,
  "minLength": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 30,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "30",
    "ves.io.schema.rules.string.min_len": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "30",
    "ves.io.schema.rules.string.min_len": "4"
  }
}
```

<a id="canonical-4f33c39eaec28c4bee3374ccda360a57fe7fcc1d9975b64c89b410d154c757e0"></a>

<a id="canonical-91095cac0fb2695484dc40191c24e9398fe0fa93d0a91d4e579f857737dc743b"></a>

## region property — gcp.byoc.connections / a7129f601f98 / 6

Type: `"string"`. Optional.

\[Enum:
asia-east1|asia-east2|asia-northeast1|asia-northeast2|asia-northeast3|asia-southeast1|asia-southeast2|europe-central2|europe-north1|europe-west1|europe-west2|europe-west3|europe-west4|europe-west6|europe-west8|europe-west9|europe-west10|europe-west12|europe-southwest1|me-west1|me-central1|me-central2|northamerica-northeast1|northamerica-northeast2|us-central1|us-east1|us-east4|us-east5|us-south1|us-west1|us-west2|us-west3|us-west4|southamerica-east1|southamerica-west1|australia-southeast1|australia-southeast2|asia-south1|asia-south2\]
GCP Region in which the GCP Cloud Interconnect attachment is configured. Possible values are
\`asia-east1\`, \`asia-east2\`, \`asia-northeast1\`, \`asia-northeast2\`, \`asia-northeast3\`,
\`asia-southeast1\`, \`asia-southeast2\`, \`europe-central2\`, \`europe-north1\`, \`europe-west1\`,
\`europe-west2\`, \`europe-west3\`, \`europe-west4\`, \`europe-west6\`, \`europe-west8\`,
\`europe-west9\`, \`europe-west10\`, \`europe-west12\`, \`europe-southwest1\`, \`me-west1\`,
\`me-central1\`, \`me-central2\`, \`northamerica-northeast1\`, \`northamerica-northeast2\`,
\`us-central1\`, \`us-east1\`, \`us-east4\`, \`us-east5\`, \`us-south1\`, \`us-west1\`,
\`us-west2\`, \`us-west3\`, \`us-west4\`, \`southamerica-east1\`, \`southamerica-west1\`,
\`australia-southeast1\`, \`australia-southeast2\`, \`asia-south1\`, \`asia-south2\`.

Upstream description:

GCP Region in which the GCP Cloud Interconnect attachment is configured.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("asia-east1",
    "asia-east2",
    "asia-northeast1",
    "asia-northeast2",
    "asia-northeast3",
    "asia-southeast1",
    "asia-southeast2",
    "europe-central2",
    "europe-north1",
    "europe-west1",
    "europe-west2",
    "europe-west3",
    "europe-west4",
    "europe-west6",
    "europe-west8",
    "europe-west9",
    "europe-west10",
    "europe-west12",
    "europe-southwest1",
    "me-west1",
    "me-central1",
    "me-central2",
    "northamerica-northeast1",
    "northamerica-northeast2",
    "us-central1",
    "us-east1",
    "us-east4",
    "us-east5",
    "us-south1",
    "us-west1",
    "us-west2",
    "us-west3",
    "us-west4",
    "southamerica-east1",
    "southamerica-west1",
    "australia-southeast1",
    "australia-southeast2",
    "asia-south1",
    "asia-south2"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "asia-east1",
    "asia-east2",
    "asia-northeast1",
    "asia-northeast2",
    "asia-northeast3",
    "asia-southeast1",
    "asia-southeast2",
    "europe-central2",
    "europe-north1",
    "europe-west1",
    "europe-west2",
    "europe-west3",
    "europe-west4",
    "europe-west6",
    "europe-west8",
    "europe-west9",
    "europe-west10",
    "europe-west12",
    "europe-southwest1",
    "me-west1",
    "me-central1",
    "me-central2",
    "northamerica-northeast1",
    "northamerica-northeast2",
    "us-central1",
    "us-east1",
    "us-east4",
    "us-east5",
    "us-south1",
    "us-west1",
    "us-west2",
    "us-west3",
    "us-west4",
    "southamerica-east1",
    "southamerica-west1",
    "australia-southeast1",
    "australia-southeast2",
    "asia-south1",
    "asia-south2"
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"asia-east1\\\",\\\"asia-east2\\\",\\\"asia-northeast1\\\",\\\"asia-northeast2\\\",\\\"asia-northeast3\\\",\\\"asia-southeast1\\\",\\\"asia-southeast2\\\",\\\"europe-central2\\\",\\\"europe-north1\\\",\\\"europe-west1\\\",\\\"europe-west2\\\",\\\"europe-west3\\\",\\\"europe-west4\\\",\\\"europe-west6\\\",\\\"europe-west8\\\",\\\"europe-west9\\\",\\\"europe-west10\\\",\\\"europe-west12\\\",\\\"europe-southwest1\\\",\\\"me-west1\\\",\\\"me-central1\\\",\\\"me-central2\\\",\\\"northamerica-northeast1\\\",\\\"northamerica-northeast2\\\",\\\"us-central1\\\",\\\"us-east1\\\",\\\"us-east4\\\",\\\"us-east5\\\",\\\"us-south1\\\",\\\"us-west1\\\",\\\"us-west2\\\",\\\"us-west3\\\",\\\"us-west4\\\",\\\"southamerica-east1\\\",\\\"southamerica-west1\\\",\\\"australia-southeast1\\\",\\\"australia-southeast2\\\",\\\"asia-south1\\\",\\\"asia-south2\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"asia-east1\\\",\\\"asia-east2\\\",\\\"asia-northeast1\\\",\\\"asia-northeast2\\\",\\\"asia-northeast3\\\",\\\"asia-southeast1\\\",\\\"asia-southeast2\\\",\\\"europe-central2\\\",\\\"europe-north1\\\",\\\"europe-west1\\\",\\\"europe-west2\\\",\\\"europe-west3\\\",\\\"europe-west4\\\",\\\"europe-west6\\\",\\\"europe-west8\\\",\\\"europe-west9\\\",\\\"europe-west10\\\",\\\"europe-west12\\\",\\\"europe-southwest1\\\",\\\"me-west1\\\",\\\"me-central1\\\",\\\"me-central2\\\",\\\"northamerica-northeast1\\\",\\\"northamerica-northeast2\\\",\\\"us-central1\\\",\\\"us-east1\\\",\\\"us-east4\\\",\\\"us-east5\\\",\\\"us-south1\\\",\\\"us-west1\\\",\\\"us-west2\\\",\\\"us-west3\\\",\\\"us-west4\\\",\\\"southamerica-east1\\\",\\\"southamerica-west1\\\",\\\"australia-southeast1\\\",\\\"australia-southeast2\\\",\\\"asia-south1\\\",\\\"asia-south2\\\"]"
  }
}
```

- [same_as_credential](resources--cloud_link--reference--group-001.md#canonical-22f4eb34d12ec3d2d5e0bfcbe8e75b52f5b621659a13ebe64f1df38959f8e997): complete subsection reference.

<a id="canonical-1d872eb3d6ed15f978187d4338bd39ec3b99af2af230347dbfb99a4e43952241"></a>

## Next pages — gcp.byoc.connections / a7129f601f98 / 7

- [gcp.byoc.connections.metadata](resources--cloud_link--reference--group-001.md#canonical-6d3d50c6cb675f2757193ac5fe44e800f9f5cb7b0ba899e7620a6b1dc59b14e4)
- [gcp.byoc.connections.same_as_credential](resources--cloud_link--reference--group-001.md#canonical-22f4eb34d12ec3d2d5e0bfcbe8e75b52f5b621659a13ebe64f1df38959f8e997)
- [gcp.byoc](resources--cloud_link--reference--group-001.md#canonical-e338639ca36e96d6bf0f5483fc7e1fc3e49b4946ef053d151ee966be7ed8847f)
- [xcsh_cloud_link](../resources/cloud_link.md#canonical-e2bde7dd950eff273e290001dfe725260a3cdddc3ece9f8132c786c2291ffa6f)

<a id="canonical-6d3d50c6cb675f2757193ac5fe44e800f9f5cb7b0ba899e7620a6b1dc59b14e4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7f5f24b806f4e3724826ad6246ef58e1e0c26c3d532b63563c98b3c62f06d451"></a>

## gcp.byoc.connections.metadata — gcp.byoc.connections.metadata / 0ce907c1ea48 / 2

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md#canonical-e2bde7dd950eff273e290001dfe725260a3cdddc3ece9f8132c786c2291ffa6f)
- [Property reference](resources--cloud_link--reference--group-001.md#canonical-f9f572eb4dd4355ccf1d4d4a87719ef48acf695e515e7a8097f8aa7e0a96fa35)
- [gcp](resources--cloud_link--reference--group-001.md#canonical-98b04cd1a263fa94de21346302cdfec25d91055cbcf9c3b463608298b2f20fd9)
- [gcp.byoc](resources--cloud_link--reference--group-001.md#canonical-e338639ca36e96d6bf0f5483fc7e1fc3e49b4946ef053d151ee966be7ed8847f)
- [gcp.byoc.connections](resources--cloud_link--reference--group-001.md#canonical-f0c351486e5af898d9b72bce3bf1819e6398294226f768d6b8ad212511d8e621)
- gcp.byoc.connections.metadata

<a id="canonical-be3de5a6640ff5870472efa66c93ac74b98d8a203850e57b44070ceaf9a364d8"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-33941f4c5cfc2c706a4053bf140723f89c9485a88cb4814df61d09d8fd9a34a7"></a>

## Direct properties — gcp.byoc.connections.metadata / 0ce907c1ea48 / 3

<a id="canonical-dab8630b39d7a0f2b979f9886414b00e4fc95ef65771ae755abad1b067d066a7"></a>

<a id="canonical-d88916126ca4a5d8b9585b09ae8583ab9907110267008666eb1f8e9ff615d7ad"></a>

## description_spec property — gcp.byoc.connections.metadata / 0ce907c1ea48 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-435098ed9e69e11bcb21c218984afb32446719dfab969ccd09e21f738ed97ca9"></a>

<a id="canonical-5dcc845e0cdfc81bab9fd6e8047e3898cdcc9d1c4517946334fe1125473e8f8a"></a>

## name property — gcp.byoc.connections.metadata / 0ce907c1ea48 / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
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
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-19477bac7d2a504a131aa5dd2bcd14678664784470e60564405cee4b237e5b27"></a>

## Next pages — gcp.byoc.connections.metadata / 0ce907c1ea48 / 6

- [gcp.byoc.connections](resources--cloud_link--reference--group-001.md#canonical-f0c351486e5af898d9b72bce3bf1819e6398294226f768d6b8ad212511d8e621)
- [xcsh_cloud_link](../resources/cloud_link.md#canonical-e2bde7dd950eff273e290001dfe725260a3cdddc3ece9f8132c786c2291ffa6f)

<a id="canonical-22f4eb34d12ec3d2d5e0bfcbe8e75b52f5b621659a13ebe64f1df38959f8e997"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f6bb7fca1d425270bbbbb0fd68df4109e972b3f4b3af18818970079801942ce2"></a>

## gcp.byoc.connections.same_as_credential — gcp.byoc.connections.same_as_credential / e7bbf4a77708 / 2

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md#canonical-e2bde7dd950eff273e290001dfe725260a3cdddc3ece9f8132c786c2291ffa6f)
- [Property reference](resources--cloud_link--reference--group-001.md#canonical-f9f572eb4dd4355ccf1d4d4a87719ef48acf695e515e7a8097f8aa7e0a96fa35)
- [gcp](resources--cloud_link--reference--group-001.md#canonical-98b04cd1a263fa94de21346302cdfec25d91055cbcf9c3b463608298b2f20fd9)
- [gcp.byoc](resources--cloud_link--reference--group-001.md#canonical-e338639ca36e96d6bf0f5483fc7e1fc3e49b4946ef053d151ee966be7ed8847f)
- [gcp.byoc.connections](resources--cloud_link--reference--group-001.md#canonical-f0c351486e5af898d9b72bce3bf1819e6398294226f768d6b8ad212511d8e621)
- gcp.byoc.connections.same_as_credential

<a id="canonical-1a84269eb9c19588aaf88767a2037167df6e0b49154fd2d87b7b87bfdb038019"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for same as credential.

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
same_as_credential = {}
```

<a id="canonical-52d1e0f0718b5bde8c3af9ceda155b7be37c82b154d502a7395becbd91a2687e"></a>

## Direct properties — gcp.byoc.connections.same_as_credential / e7bbf4a77708 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0d179bc935a64342a76760839d4ee37867b09e83ea3434d500df6654896d5a38"></a>

## Next pages — gcp.byoc.connections.same_as_credential / e7bbf4a77708 / 4

- [gcp.byoc.connections](resources--cloud_link--reference--group-001.md#canonical-f0c351486e5af898d9b72bce3bf1819e6398294226f768d6b8ad212511d8e621)
- [xcsh_cloud_link](../resources/cloud_link.md#canonical-e2bde7dd950eff273e290001dfe725260a3cdddc3ece9f8132c786c2291ffa6f)

<a id="canonical-dbb5eb32808aa37a77eb7dad058d46d3db79e79b5044106c4ed056fc309480b2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8a8b7ae6d1ee58a0a87db56893640aa5ab2ffd31582aefde258191f5f586389d"></a>

## gcp.gcp_cred — gcp.gcp_cred / c0d6b26f8fa5 / 2

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md#canonical-e2bde7dd950eff273e290001dfe725260a3cdddc3ece9f8132c786c2291ffa6f)
- [Property reference](resources--cloud_link--reference--group-001.md#canonical-f9f572eb4dd4355ccf1d4d4a87719ef48acf695e515e7a8097f8aa7e0a96fa35)
- [gcp](resources--cloud_link--reference--group-001.md#canonical-98b04cd1a263fa94de21346302cdfec25d91055cbcf9c3b463608298b2f20fd9)
- gcp.gcp_cred

<a id="canonical-02b0cf1064e23607b9644c6fe500cf85441de84725bd2569ad037517bcff2526"></a>

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
gcp_cred {
  # Configure direct properties listed below.
}
```

<a id="canonical-231316479d8bbc5673c2bce60e268bfafb8e52c1e90307312f1734231e95d759"></a>

## Direct properties — gcp.gcp_cred / c0d6b26f8fa5 / 3

<a id="canonical-b03ae26c2cdee37996a07c96264acfa81611538ca48ee91f79f38ca606b7ae80"></a>

<a id="canonical-44c8485a0b27a3475ffd429a8acf215c06f74241f035b0e5bef583590e519a82"></a>

## name property — gcp.gcp_cred / c0d6b26f8fa5 / 4

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

<a id="canonical-cbad7213c4dd1bc922f31ac0e9e6ae8c5b6e8a22c2464e5bc00e150e77f9a60d"></a>

<a id="canonical-3bd1e10aa33fd68a574be9a98678295b9b086f99baefdf69e017941a703d58b2"></a>

## namespace property — gcp.gcp_cred / c0d6b26f8fa5 / 5

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

<a id="canonical-960187c666731c79461e5925abc97a067370040d1dd3d4676087cd8c93595035"></a>

<a id="canonical-e6ab09bbddeb8133600bc5e374e5ef162f94b7b9d003cb8a0f88404577fac550"></a>

## tenant property — gcp.gcp_cred / c0d6b26f8fa5 / 6

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

<a id="canonical-83400244d20cb7a10b83e625234b362246a75f7c702848f5b8a9b702979439fe"></a>

## Next pages — gcp.gcp_cred / c0d6b26f8fa5 / 7

- [gcp](resources--cloud_link--reference--group-001.md#canonical-98b04cd1a263fa94de21346302cdfec25d91055cbcf9c3b463608298b2f20fd9)
- [xcsh_cloud_link](../resources/cloud_link.md#canonical-e2bde7dd950eff273e290001dfe725260a3cdddc3ece9f8132c786c2291ffa6f)

<a id="canonical-407980008cac2d543b4e7b2f130f22cee93d886cd60141abd6ce82624b6f9f80"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eb81abf0a469ffd6e408f5b5104da5beea280d69e202b36fccb92c6df69c67e7"></a>

## timeouts — timeouts / 2a46a1e31fb7 / 2

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md#canonical-e2bde7dd950eff273e290001dfe725260a3cdddc3ece9f8132c786c2291ffa6f)
- [Property reference](resources--cloud_link--reference--group-001.md#canonical-f9f572eb4dd4355ccf1d4d4a87719ef48acf695e515e7a8097f8aa7e0a96fa35)
- timeouts

<a id="canonical-51a4f5daddbbed2aacc8cf0eead62606f49de452f1c165a3fffcd5626be340ef"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-08231be27fa840e84ac3bf713a451c2b2fe64365b035b221761184bce31a4a15"></a>

## Direct properties — timeouts / 2a46a1e31fb7 / 3

<a id="canonical-65038fbd75e19dcb633ca0ccdc2069a8b2418b6bd89a349322745809bb2e6124"></a>

<a id="canonical-e0013bf4f9582054c7692394693697c370d8c0d778c8b6bfa7728b6f0ee802e5"></a>

## create property — timeouts / 2a46a1e31fb7 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-63f07ef8527de6aaa7357ef34d6f90fb16161cae2f80889c8924f230619e6a4d"></a>

<a id="canonical-8e5926b2a79b895a3d86b1dd196fcc9d9e29b3338c5ea9faa2dc9e4aebbdf1f2"></a>

## delete property — timeouts / 2a46a1e31fb7 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-117db8d285456a4e8edf44857120bf9afecd474b6110a1512436b4b124b1100a"></a>

<a id="canonical-8d04503c86370a0bc02090537593293a6f0e2528b8ab7b244268027dc53c1078"></a>

## read property — timeouts / 2a46a1e31fb7 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-6172a5a11be2275f73910304dd4806a077680268f6cf6adbc97601e993e8b21f"></a>

<a id="canonical-6e06428cf8993b14ad708cd7eeefd608b0e7d0afb5243a00eef00caef6e33419"></a>

## update property — timeouts / 2a46a1e31fb7 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-e50039b443b8f224d560554f74e140b32a8b59531b9526878b933576a99787e9"></a>

## Next pages — timeouts / 2a46a1e31fb7 / 8

- [Property reference](resources--cloud_link--reference--group-001.md#canonical-f9f572eb4dd4355ccf1d4d4a87719ef48acf695e515e7a8097f8aa7e0a96fa35)
- [xcsh_cloud_link](../resources/cloud_link.md#canonical-e2bde7dd950eff273e290001dfe725260a3cdddc3ece9f8132c786c2291ffa6f)
